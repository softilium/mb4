package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	DbType             string
	DbConnectionString string
	ListenAddr         string
	SessionCookieName  string
	SessionCryptKey    string
	Debug              bool
	AdminLogin         string // initial values for empty database
	AdminPassw         string
}

var C ConfigData

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	C.DbConnectionString = os.Getenv("DB_CONNECTION_STRING")
	C.DbType = os.Getenv("DB_TYPE")
	C.ListenAddr = os.Getenv("LISTEN_ADDR")
	C.SessionCookieName = os.Getenv("SESSION_COOKIE_NAME")
	C.SessionCryptKey = os.Getenv("SESSION_CRYPT_KEY")
	C.Debug = os.Getenv("DEBUG") == "true"
	C.AdminLogin = os.Getenv("ADMIN_LOGIN")
	C.AdminPassw = os.Getenv("ADMIN_PASSW")

}
