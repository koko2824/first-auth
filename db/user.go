package db

import (
	"context"
	"first-auth/models"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
)


func Insert(user models.User) error {
	return user.Insert(context.Background(), DB, boil.Infer())
}