package main

import (
	"fmt"
	"struktur/management"
)


func main() {

	user := management.User{1,"try","suo","ahah@gmail.com", true}
	user2 := management.User{2,"try","suo","ahah@gmail.com", true}


	fmt.Println(user.FirstName)

	fmt.Println("================")

	users := []management.User{user,user2}
	group := management.Group{"Gamer", user, users, true}


	fmt.Println("====================")

	result := user.Display()
	fmt.Println(result)

	fmt.Println(user2.Display())


	fmt.Println("=====================")

	group.DisplayGroup()


}

