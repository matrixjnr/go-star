package main

import (
	"fmt"
)

// simple user details struct
type user struct {
	firstName string
	lastName  string
	email     string
	age       int
}

var users []user = make([]user, 0)

func addUser(u user) {
	users = append(users, u)
}

func getUserByEmail(email string) user {
	for _, u := range users {
		if u.email == email {
			return u
		}
	}
	return user{}
}

func updateUserByEmail(oldEmail string, u user) {
	for i, user := range users {
		if user.email == oldEmail {
			users[i] = u
			return
		}
	}
}

func deleteUserByEmail(email string) {
	for i, u := range users {
		if u.email == email {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
}

func main() {
	fmt.Println("The source code of life begins here")
	// lets scan user input to add user details
	var email string

	// print welcome to user db
	fmt.Println("Welcome to User DB")

	// it is a continuous loop of user input and printing data until user decides to exit
	for {
		fmt.Println("Menu Commands")
		// list of commands
		fmt.Println("1 - Add User")
		fmt.Println("2 - Update User")
		fmt.Println("3 - Delete User")
		fmt.Println("4 - Get User")
		fmt.Println("5 - Get All Users")
		fmt.Println("q - Quit")
		fmt.Print("Enter Command: ")
		var command string
		fmt.Scan(&command)

		switch command {
		case "q":
			fmt.Println("Exiting User DB")
			return
		case "1":
			enterUserDetails()
		case "2":
			updateData()
		case "3":
			fmt.Println("Enter Email: ")
			fmt.Scanln(&email)
			deleteUserByEmail(email)
		case "4":
			fmt.Println("Enter Email: ")
			fmt.Scanln(&email)
			u := getUserByEmail(email)
			if u.email == "" {
				fmt.Println("User not found")
			} else {
				fmt.Println(u)
			}
		case "5":
			for _, u := range users {
				fmt.Println(u)
			}
		default:
			fmt.Println("Invalid Command")
		}
	}
}

func updateData() {
	var firstName, lastName, email, newEmail string
	var age int

	fmt.Println("Enter Email of User to Update: ")
	fmt.Scanln(&email)

	fmt.Println("Enter New First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter New Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter New Email: ")
	fmt.Scanln(&newEmail)

	fmt.Println("Enter New Age: ")
	fmt.Scanln(&age)

	updateUserByEmail(email, user{
		firstName: firstName,
		lastName:  lastName,
		email:     newEmail,
		age:       age,
	})
}

func enterUserDetails() {
	var firstName, lastName, email string
	var age int

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter Age: ")
	fmt.Scanln(&age)

	addUser(user{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		age:       age,
	})
}
