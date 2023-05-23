package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/YuichiNAGAO/sql-boiler_sql-migrate_benkyo/db/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("mysql", "user:mysql@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	// read
	user1, err := models.FindUser(ctx, db, 1)
	fmt.Println(user1)

	// create
	var newuser models.User
	newuser.Name = "test"
	newuser.Email = "test"
	err = newuser.Insert(ctx, db, boil.Infer())

	// user2 := &models.User{
	// 	Name:  "test2",
	// 	Email: "test2",
	// }
	// err = user2.Insert(ctx, db, boil.Infer())

	// update
	user1.Name = "updated name"
	user1.Email = "updated email"
	rowsAff, err := user1.Update(ctx, db, boil.Infer())
	fmt.Println(rowsAff)

	// delete
	allUsers, _ := models.Users().All(ctx, db)
	lastUser := allUsers[len(allUsers)-1]
	fmt.Println(allUsers)
	fmt.Println(lastUser)
	lastUser.Delete(ctx, db)

	// カラムでselect
	userTest, _ := models.Users(models.UserWhere.Name.EQ("test")).One(ctx, db)
	fmt.Println(userTest)
}
