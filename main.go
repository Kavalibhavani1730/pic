package main

import (
	"fmt"
	"strconv"
	"strings"

)

// package variable declaration
const conferenceTickets int = 50

var conferenceName string = "maryland"
var RemainingTickets uint = 50
var bookings = make([]map[string]string, 0)


func main() {

	fmt.Println("welcome to ticket booking")
	fmt.Println("get your ticket here")

	greeting()

	getFirstNames()
	//var bookings [50]string without slice

	for {

		//if userTickets>remainingTickets{
		//fmt.Printf("we only have %v tickets we cant book which exceeds that",remainingTickets)
		//continue
		//}
		firstname, lastname, email, userTickets := unserinputs()

		isNameValidate, isMailValidate, isTickets := Validateone(firstname, lastname, email, userTickets)

		if isNameValidate && isMailValidate && isTickets {
			RemainingTickets = RemainingTickets - userTickets

			//create a map for a user
			var userData = map[string]string
            userData=["firstname"] = firstname
			userData=["lastname"] = lastname
			userData=["email"] = email
			userData["numberoftickets"]=strconv.FormatUint(uint64(userTickets),10)



			//bookings[0] = firstname + " " + lastname
			bookings = append(bookings,userData )

			fmt.Printf("thank you %v %v for booking %v for %v you will receive an confirmation mail to your %v. \n", firstname, lastname, userTickets, conferenceName, email)
			fmt.Printf("remaining tickets left are %v in %v. \n", RemainingTickets, conferenceName)

			//fmt.Printf("the whole array is %v\n", bookings)
			//fmt.Printf("the first array is %T\n", bookings[0])
			//fmt.Printf("the length of an array is %v\n", len(bookings))
			firstnames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("printing all the bookings %v \n", firstnames)

			//for getting getfirstname
			firstNames := getFirstNames()
			fmt.Printf("the firstname is %v", firstNames)

			if RemainingTickets == 0 {
				fmt.Println("all tickets are booked")
				break
			}
		} else {
			if !isNameValidate {
				fmt.Println("name is not valid")
			}
			if !isMailValidate {
				fmt.Println("mail is not valid")
			}
			if !isTickets {
				fmt.Println("enter valid number of tickets")
			}
		}

	}

}
func greeting() {
	fmt.Printf("welcome to the %v booking application \n", conferenceName)
	fmt.Printf("we have total tickets of %v \n", conferenceTickets)
	fmt.Printf("get your tickets here to attend %v and the tickets available are %v \n ", remainingTickets, conferenceTickets)
}
func getFirstNames() []string {

	firstnames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}

func unserinputs() (string, string, string, uint) {
	var firstname string
	var lastname string
	var userTickets uint
	var email string

	fmt.Println("enter you firstname")
	fmt.Scan(&firstname)
	fmt.Println("enter you lastname")
	fmt.Scan(&lastname)
	fmt.Println("enter your email")
	fmt.Scan(&email)
	fmt.Println("enter your number of tickets")
	fmt.Scan(&userTickets)
	return firstname, lastname, email, userTickets
}
