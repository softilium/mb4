package api

import (
	"context"
	"strings"
	"time"

	"net/http"

	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent/schema"
	"github.com/softilium/mb4/ent/user"
	"github.com/softilium/mb4/pages"
)

func ApiLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userName := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	userName = strings.ToLower(strings.TrimSpace(userName))
	if len(userName) < 3 {
		w.WriteHeader(http.StatusBadRequest)
	}

	u, err := db.DB.User.Query().
		Limit(1).
		Where(user.And(user.UserNameEQ(userName), user.PasswordHashEQ(db.PasswordHash(password)))).
		All(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(u) == 1 {
		session, _ := pages.SessionsStore.Get(r, config.C.SessionCookieName)
		session.Values["userName"] = userName
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)

}

func Auth_Logout(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/auth/refresh",
	})

	w.WriteHeader(http.StatusOK)

}

func ApiRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userName := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	userName = strings.ToLower(strings.TrimSpace(userName))
	if len(userName) < 3 {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err := db.DB.User.Query().
		Where(user.UserNameEQ(userName)).
		Only(context.Background())
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	_, err = db.DB.User.Create().
		SetAuthType(schema.Auth_Type_email).
		SetUserName(strings.TrimSpace(strings.ToLower(userName))).
		SetPasswordHash(db.PasswordHash(password)).
		Save(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, _ := pages.SessionsStore.Get(r, config.C.SessionCookieName)
	session.Values["userName"] = userName
	session.Values["authenticated"] = true
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)

}
