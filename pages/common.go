package pages

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/gorilla/sessions"
	"github.com/rs/xid"
	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/user"
)

var (
	templates     map[string]*template.Template
	SessionsStore *sessions.CookieStore
)

type SessionStruct struct {
	User          *ent.User //there will be another session's fields here
	Authenticated bool
	UserName      string
}

func LoadSessionStruct(r *http.Request) SessionStruct {

	data := SessionStruct{User: nil, Authenticated: false, UserName: ""}

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
		data.User = users[0]
		data.Authenticated = true
		data.UserName = data.User.UserName
	}

	return data

}

func init() {
	var err error
	templates, err = amber.CompileDir("pages/", amber.DirOptions{Ext: ".pug"}, amber.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}

	sessionkey := []byte(config.C.SessionCryptKey)
	SessionsStore = sessions.NewCookieStore(sessionkey)
	SessionsStore.Options.HttpOnly = true

}
