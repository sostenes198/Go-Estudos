package main

import (
	"fmt"
	"implementacao/repository/contracts"
	"implementacao/repository/sql"
	"implementacao/repository/sql/mysql"
	"log"
)

var userRepository contracts.UserRepository

func main() {
	db := sql.Open()
	defer db.Close()
	userRepository = mysql.NewUserRepository(db)

	afterUser, err := userRepository.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(afterUser)

	//user, err := entitity.NewUser(entitity.Params{Id: 1, Name: "Batata", Email: "ASD@ASD.com"})
	if err := userRepository.Delete(1); err != nil {
		log.Fatal(err)
	}

	beforeUser, err := userRepository.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(beforeUser)
}
