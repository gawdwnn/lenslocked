package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"lenslocked.com/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "godwin"
	dbname = "lenslockedDb_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	
	user := models.User{
		Name:  "Mich James",
		Email: "Mich@gmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}

	userById, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userById)
}

//psql -U postgres -p 5432 -h localhost lenslockedDb_dev
