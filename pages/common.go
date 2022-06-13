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
	"github.com/softilium/mb4/ent/user"
)

var (
	templates     map[string]*template.Template
	SessionsStore *sessions.CookieStore
)

type sessionStruct struct {
	Authenticated bool
	UserName      string
	UserId        string // run-time, not in the session data
}

func sessionIsAuth(s *sessions.Session) bool {
	v, ok := s.Values["authenticated"].(bool)
	return !ok || (ok && v)
}

func loadSessionStruct(r *http.Request) sessionStruct {

	data := sessionStruct{Authenticated: false, UserName: ""}

	session, _ := SessionsStore.Get(r, config.C.SessionCookieName)

	if sessionIsAuth(session) {

		userId, ok := session.Values["userId"].(string)
		if !ok {
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
			data.Authenticated = true
			data.UserId = users[0].ID.String()
			data.UserName = users[0].UserName
		}

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
