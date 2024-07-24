package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Simple user details struct
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
	// Print welcome to user db
	fmt.Println("Welcome to User DB")

	reader := bufio.NewReader(os.Stdin)

	// It is a continuous loop of user input and printing data until the user decides to exit
	for {
		printMenu()

		fmt.Print("Enter Command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "q":
			fmt.Println("Exiting User DB")
			return
		case "1":
			enterUserDetails(reader)
		case "2":
			updateData(reader)
		case "3":
			fmt.Print("Enter Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
			deleteUserByEmail(email)
		case "4":
			fmt.Print("Enter Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)
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

func printMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1 - Add User")
	fmt.Println("2 - Update User")
	fmt.Println("3 - Delete User")
	fmt.Println("4 - Get User")
	fmt.Println("5 - Get All Users")
	fmt.Println("q - Quit")
}

func updateData(reader *bufio.Reader) {
	var firstName, lastName, email, newEmail string
	var age int

	fmt.Print("Enter Email of User to Update: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter New First Name: ")
	firstName, _ = reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter New Last Name: ")
	lastName, _ = reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter New Email: ")
	newEmail, _ = reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	fmt.Print("Enter New Age: ")
	fmt.Scanf("%d\n", &age)

	updateUserByEmail(email, user{
		firstName: firstName,
		lastName:  lastName,
		email:     newEmail,
		age:       age,
	})
}

func enterUserDetails(reader *bufio.Reader) {
	var firstName, lastName, email string
	var age int

	fmt.Print("Enter First Name: ")
	firstName, _ = reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter Last Name: ")
	lastName, _ = reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter Age: ")
	fmt.Scanf("%d\n", &age)

	addUser(user{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		age:       age,
	})
}
