package pages

import (
	"context"
	"log"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/gorilla/sessions"
	"github.com/rs/xid"
	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/investaccount"
	"github.com/softilium/mb4/ent/user"
)

var (
	SessionsStore *sessions.CookieStore
)

type SessionStruct struct {
	user          *ent.User //there will be another session's fields here
	Authenticated bool      //in pug templated we can use only fields, not funcs. We need to have fields from User here
	UserID        xid.ID    //this is the same as user.ID
	UserIDStr     string
	UserName      string
	UserIsAdmin   bool
	Debug         bool
	Vue           bool
	Echarts       bool
}

func (session *SessionStruct) GetUser() *ent.User {

	if session.user == nil {
		panic("session.user is nil")
	}
	return session.user
}

func (session *SessionStruct) GetInvestAccountXids() ([]xid.ID, error) {

	allUserAccounts, err := db.DB.InvestAccount.Query().Where(investaccount.HasOwnerWith(user.IDEQ(session.UserID))).All(context.Background())
	if err != nil {
		return nil, err
	}

	xids := make([]xid.ID, len(allUserAccounts))
	for i, v := range allUserAccounts {
		xids[i] = v.ID
	}

	return xids, nil

}

func (session *SessionStruct) GetInvestAccountXidsMap() (map[xid.ID]bool, error) {

	list, err := session.GetInvestAccountXids()
	res := make(map[xid.ID]bool)
	if err != nil {
		return res, err
	}
	for _, v := range list {
		res[v] = true
	}

	return res, nil

}

func LoadSessionStruct(r *http.Request) SessionStruct {

	data := SessionStruct{user: nil, Authenticated: false, UserName: "", UserIsAdmin: false, Debug: config.C.Debug}

	session, err := SessionsStore.Get(r, config.C.SessionCookieName)
	if err != nil {
		log.Println(err)
		return data
	}

	userId, ok := session.Values["userId"].(string)
	if !ok || userId == "" {
		return data
	}

	xid, err := xid.FromString(userId)
	if err != nil {
		log.Println(err)
		return data
	}

	users, err := db.DB.User.Query().Limit(1).Where(user.IDEQ(xid)).All(context.Background())
	if err != nil {
		log.Println(err)
		return data
	}

	if len(users) == 1 {
		data.user = users[0]
		data.Authenticated = true
		data.UserName = data.user.UserName
		data.UserIsAdmin = data.user.Admin
		data.UserID = data.user.ID
		data.UserIDStr = data.user.ID.String()
	}

	return data

}

func init() {

	pongo2.DefaultSet.Debug = config.C.Debug

	sessionkey := []byte(config.C.SessionCryptKey)
	SessionsStore = sessions.NewCookieStore(sessionkey)
	SessionsStore.Options.HttpOnly = true

}
