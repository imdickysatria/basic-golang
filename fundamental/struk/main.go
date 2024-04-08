package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}


type Group struct {
	name string
	admin User
	users []User
	isAvailable bool
}
func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Riyan Dicky"
	user.LastName = "Satria"
	user.Email = "dickysatria62@gmail.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Riyan Dicky"
	user2.LastName = "Satria"
	user2.Email = "dickysatria62@gmail.com"
	user2.IsActive = true


	user3 := User{ID: 3, FirstName: "Jack", LastName: "Kalaku", IsActive: false}

	user4 := User{
		ID: 4,
		FirstName: "Oks",
		LastName: "Bat",
		Email: "dicky@gmail.com",
		IsActive: false,
	}

	user5 := User{5,"try","suo","ahah@gmail.com", true}

	fmt.Println(user.FirstName)
	fmt.Println(user2.LastName)
	fmt.Println(user3)
	fmt.Println(user4)
	fmt.Println(user5)

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

	fmt.Println("================")

	users := []User{user3, user2}
	group := Group{"Gamer", user, users, true}

	displayGroup(group)


	fmt.Println("====================")

	result := user.display()
	fmt.Println(result)

	fmt.Println(user2.display())


	fmt.Println("=====================")

	group.displayGroup()
	// fmt.Println(result2)


}

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email: %s",user.FirstName, user.LastName, user.Email)
}

func (group Group) displayGroup() {
	fmt.Printf("Name : %s",group.name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.users))
	fmt.Println("")
	fmt.Println("Users name :")
	for _, user := range group.users {
		fmt.Println(user.FirstName)
	}
	
}

func displayGroup(group Group){
	fmt.Printf("Name : %s",group.name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.users))
	fmt.Println("")
	fmt.Println("Users name :")
	for _, user := range group.users {
		fmt.Println(user.FirstName)
	}
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email: %s",user.FirstName, user.LastName, user.Email)
	return result
}