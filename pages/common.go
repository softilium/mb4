package pages

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eknkc/amber"
	"github.com/gorilla/sessions"
	"github.com/softilium/mb4/config"
)

var (
	templates     map[string]*template.Template
	sessionsStore *sessions.CookieStore
)

type authInfo struct {
	Authenticated bool
	UserName      string
}

func sessionUserName(s *sessions.Session) string {
	v, ok := s.Values["userName"].(string)
	if !ok {
		return v
	}
	return ""
}

func sessionIsAuth(s *sessions.Session) bool {
	v, ok := s.Values["authenticated"].(bool)
	return !ok || (ok && v)
}

func loadAuthInfo(r *http.Request) authInfo {

	data := authInfo{Authenticated: false, UserName: ""}

	session, _ := sessionsStore.Get(r, config.C.SessionCookieName)

	if sessionIsAuth(session) {
		data.Authenticated = true
		data.UserName = sessionUserName(session)
	}

	return data

}

func init() {
	var err error
	templates, err = amber.CompileDir("pages/", amber.DirOptions{Ext: ".pug"}, amber.DefaultOptions)
	if err != nil {
		log.Fatal(err)
	}

	sessionkey := []byte("super-secret-key")
	sessionsStore = sessions.NewCookieStore(sessionkey)

}
