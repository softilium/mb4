package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigData struct {
	AccessTokenLife    int // in minutes
	AccessSecretKey    string
	RefreshSecretKey   string
	DebugMode          bool
	DbType             string
	DbConnectionString string
	ListenAddr         string
	SessionCookieName  string
	SessionCryptKey    string
}

var C ConfigData

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	C.AccessTokenLife, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE"))
	C.AccessSecretKey = os.Getenv("ACCESS_SECRET_KEY")
	C.RefreshSecretKey = os.Getenv("REFRESH_SECRET_KEY")
	C.DebugMode, _ = strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	C.DbConnectionString = os.Getenv("DB_CONNECTION_STRING")
	C.DbType = os.Getenv("DB_TYPE")
	C.ListenAddr = os.Getenv("LISTEN_ADDR")
	C.SessionCookieName = os.Getenv("SESSION_COOKIE_NAME")
	C.SessionCryptKey = os.Getenv("SESSION_CRYPT_KEY")

}
