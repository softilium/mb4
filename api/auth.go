package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/db"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/schema"
	"github.com/softilium/mb4/ent/user"
)

type credentialsPair struct {
	//	ID       int    `json:"id",omitempty`
	Password string `json:"password",omitempty`
	UserID   string `json:"userId",omitempty`
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "user", "")
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func Auth_Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var creds credentialsPair
	//var user db.DB.  User
	var err error

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Printf("Unable to extract login/password from: %s", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(creds.UserID) == "" {
		log.Printf("Empty UserID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creds.UserID = strings.TrimSpace(strings.ToLower(creds.UserID))

	u, err := db.DB.User.Query().
		Where(user.And(user.UserNameEQ(creds.UserID), user.PasswordHashEQ(db.PasswordHash(creds.Password)))).
		Only(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	} else {
		newhash := db.PasswordHash(creds.Password)
		if u.PasswordHash != newhash {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

	}

	setNewAuthCookies(w, u)

}

func Auth_Logout(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
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

func Auth_RefreshToken(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.C.RefreshSecretKey), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userId := claims["jti"].(string)

	u, err := db.DB.User.Query().
		Where(user.UserNameEQ(userId)).
		Only(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	accessTokenString, err := createToken(u, config.C.AccessSecretKey, config.C.AccessTokenLife)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: accessTokenString,
		//Expires:  expiration,
		HttpOnly: true,
	})

	refreshTokenString, err := createToken(u, config.C.RefreshSecretKey, -1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: refreshTokenString,
		//Expires:  expiration,
		HttpOnly: true,
		Path:     "/auth/refresh",
	})

}

func Auth_Me(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	token_is_ok := false

	token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		token_is_ok = true
		return []byte(config.C.AccessSecretKey), nil
	})

	if !token_is_ok {
		log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		w.WriteHeader(419) // "timeout" it is signal to use refresh token
		return
	}

	userId := claims["jti"].(string)

	u, err := db.DB.User.Query().
		Where(user.UserNameEQ(userId)).
		Only(context.Background())
	if err != nil {
		Auth_Logout(w, r)
		return
	}

	userInfo := credentialsPair{UserID: u.UserName}
	result, err := json.Marshal(userInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func Auth_Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var creds credentialsPair
	var err error

	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Printf("Unable to extract login/password from: %s", r.Body)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = db.DB.User.Query().
		Where(user.UserNameEQ(creds.UserID)).
		Only(context.Background())
	if err == nil {
		w.WriteHeader(http.StatusFound)
		return
	}

	u, err := db.DB.User.Create().
		SetAuthType(schema.Auth_Type_email).
		SetUserName(strings.TrimSpace(strings.ToLower(creds.UserID))).
		SetPasswordHash(db.PasswordHash(creds.Password)).
		Save(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setNewAuthCookies(w, u)

	w.WriteHeader(http.StatusOK)

}

func Auth_ResetPassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)

}

func Auth_UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)

}

func createToken(User *ent.User, secret string, minutesLife int) (string, error) {

	claims := jwt.StandardClaims{}
	if minutesLife > 0 {
		expiration := time.Now().Add(time.Minute * time.Duration(minutesLife))
		claims.ExpiresAt = expiration.Unix()
	}
	claims.Id = User.UserName

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenAsString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenAsString, nil

}

func setNewAuthCookies(w http.ResponseWriter, u *ent.User) {

	accessTokenString, err := createToken(u, config.C.AccessSecretKey, config.C.AccessTokenLife)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: accessTokenString,
		//Expires:  expiration,
		HttpOnly: true,
	})

	refreshTokenString, err := createToken(u, config.C.RefreshSecretKey, -1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: refreshTokenString,
		//Expires:  expiration,
		HttpOnly: true,
		Path:     "/auth/refresh",
	})

}
