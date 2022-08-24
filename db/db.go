package db

import (
	"context"
	"fmt"
	"log"

	"crypto/sha512"

	"github.com/softilium/mb4/config"
	"github.com/softilium/mb4/ent"
	"github.com/softilium/mb4/ent/migrate"
	"github.com/softilium/mb4/ent/user"

	_ "github.com/lib/pq"
)

var DB *ent.Client

func PasswordHash(Password string) string {

	v := sha512.Sum512([]byte(Password))
	return fmt.Sprintf("%x", v)

}

func init() {

	var err error

	DB, err = ent.Open(config.C.DbType, config.C.DbConnectionString)
	if err != nil {
		log.Fatalf("Failed opening connection: %v", err)
	}

	err = DB.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(false),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true))
	if err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	if admins, err := DB.User.Query().
		Where(user.UserNameEQ("admin")).
		Limit(1).
		All(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	} else {
		if len(admins) == 0 {
			_, err = DB.User.Create().
				SetUserName("admin").
				SetPasswordHash(PasswordHash("admin")).
				SetAdmin(true).
				Save(context.Background())
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Println("Created admin user")
		}

	}

}
