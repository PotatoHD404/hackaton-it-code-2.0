package main

import (
	"fmt"
	"github.com/joho/godotenv"
	function "hackaton-it-code-2.0/src/api/handler"
)

type User struct {
	id      int
	name    string
	surname string
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	//ctx := context.Background()
	db := function.NewDB()
	//var users []User
	//query, err := db.Query("SELECT * FROM users;", &user.id)
	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		panic(err.Error())
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.name, &user.surname, &user.id)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	fmt.Println(users)
	//fmt.Println(query)
	//fmt.Println(exec)
	//if err != nil {
	//	return
	//}
	//handler := function.NewHttpHandler()

	//log.Fatal(http.ListenAndServe(":3000", handler))
}
