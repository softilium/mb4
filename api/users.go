package api

import (
	"context"
	"encoding/base64"
	"strings"
	"time"

	"net/http"

	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent/user"
	"github.com/softilium/mb4/pages"
)

// UsersLogin supports both username+password form fields and Authorization header
func UsersLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userName := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	userName = strings.ToLower(strings.TrimSpace(userName))
	if len(userName) < 3 {

		authraw := r.Header.Get("Authorization")
		if authraw != "" {
			tokens := strings.Split(authraw, " ")
			if (len(tokens) == 2) && (tokens[0] == "Basic") {

				t1dec, err := base64.StdEncoding.DecodeString(tokens[1])
				if err == nil {
					tokens2 := strings.Split(string(t1dec), ":")
					if len(tokens2) == 2 {
						userName = tokens2[0]
						password = tokens2[1]
					}
				}
			}
		}

		if len(userName) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	u, err := db.DB.User.Query().
		Limit(1).
		Where(user.And(user.UserNameEQ(userName), user.PasswordHashEQ(db.PasswordHash(password)))).
		All(context.Background())
	handleErr(err, w)
	if len(u) == 1 {

		session, _ := pages.SessionsStore.Get(r, config.C.SessionCookieName)
		session.Values["userId"] = u[0].ID.String()
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)

}

//UsersLogout clears session and redirects to login page
func UsersLogout(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, _ := pages.SessionsStore.Get(r, config.C.SessionCookieName)
	session.Values["userId"] = ""
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)

}

//UsersRegister registers new user
func UsersRegister(w http.ResponseWriter, r *http.Request) {

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

	nu, err := db.DB.User.Create().
		SetUserName(strings.TrimSpace(strings.ToLower(userName))).
		SetPasswordHash(db.PasswordHash(password)).
		Save(context.Background())
	handleErr(err, w)

	session, _ := pages.SessionsStore.Get(r, config.C.SessionCookieName)
	session.Values["userId"] = nu.ID.String()
	session.Values["authenticated"] = true
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)

}

// /api/auth/start-invest-accounts-flow
func UsersStartInvestAccountsFlow(w http.ResponseWriter, r *http.Request) {

	session := pages.LoadSessionStruct(r)
	if !session.Authenticated {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	if r.Method == http.MethodPost {

		parNewdate := r.URL.Query().Get("newdate")
		newValue, err := time.Parse("2006-01-02", parNewdate)
		handleErr(err, w)

		_, err = db.DB.User.UpdateOneID(session.UserID).SetStartInvestAccountsFlow(newValue).Save(context.Background())
		handleErr(err, w)
	}

	if r.Method == http.MethodGet {
		if !session.GetUser().StartInvestAccountsFlow.IsZero() {
			res := session.GetUser().StartInvestAccountsFlow.Format("2006-01-02")
			w.Write([]byte(res))
		}
	}
}
