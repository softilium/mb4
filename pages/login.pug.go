package pages

import (
	"context"
	"net/http"

	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent/user"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		templates["login"].Execute(w, nil)
	}
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		ph := db.PasswordHash(password)

		users, err := db.DB.User.Query().
			Limit(1).
			Where(user.And(user.UserNameEQ(username), user.PasswordHashEQ(ph))).
			All(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if len(users) == 1 {
			session, _ := sessionsStore.Get(r, config.C.SessionCookieName)
			session.Values["authenticated"] = true
			session.Values["userName"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			errInfo := struct{ ErrorDescr string }{ErrorDescr: "Invalid username or password"}
			templates["login"].Execute(w, errInfo)
		}

	}

}
