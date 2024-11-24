package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Pointer principles
var a int = 1

var b *int = &a // or b := &a (pointer points to the memory address of a)

func updateInt(x *int) {
	*x = 25
}

var c string = "Jace"

var d *string = &c //pointer points to the memory address of c

func updateString(x *string) {
	*x = "Awesome Jace"
}

var myNameString string = "jace"

var myNamePointer *string = &myNameString

// Declare
var myNamesSlice []string = []string{"Joe", "Harry", "Luke"}

var myNamesMap map[int]string = map[int]string{1: "Jace", 2: "Henry", 3: "David"}

func myGoodMorningFunc(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func myGoodByeFunc(n string) {
	fmt.Printf("Good bye %v\n", n)
}

func mySimpleGreetingsFunc(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func myComplexGreetingsFunc(n map[int]string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// Create a global waitgroup variable. WaitGroup waits for a collection of goroutines to finish
var wg *sync.WaitGroup = &sync.WaitGroup{}

// var wg sync.WaitGroup

// Main function
func main() {

	/*
	   	var a int = 1
	       var b *int = &a

	   	fmt.Println(a)
	   	fmt.Println(b)  //b is pointing to the memory address of a
	   	fmt.Println(*b) //*b is pointing to the value of a

	   	*b = *b + 2

	   	fmt.Println(a)
	*/

	fmt.Printf("The value of int a is %v", a)
	fmt.Println()
	fmt.Println(b)
	fmt.Println()
	fmt.Println(*b)
	fmt.Println()
	updateInt(b)
	fmt.Println()
	fmt.Printf("The new value of int a is %v", a)
	fmt.Println()
	fmt.Printf("The value of string c is %v", c)
	fmt.Println()
	updateString(d)
	fmt.Println()
	fmt.Printf("The value of new string c is %v", c)
	fmt.Println()
	fmt.Println(myNameString)
	fmt.Println()
	fmt.Println(myNamePointer)
	fmt.Println()
	fmt.Println(*myNamePointer)
	fmt.Println()
	mySimpleGreetingsFunc(myNamesSlice, myGoodMorningFunc)
	fmt.Println()
	mySimpleGreetingsFunc(myNamesSlice, myGoodByeFunc)
	fmt.Println()
	myComplexGreetingsFunc(myNamesMap, myGoodMorningFunc)
	fmt.Println()
	myComplexGreetingsFunc(myNamesMap, myGoodByeFunc)
	fmt.Println()

	//-----------Using struct from user-struct.go--------------------

	//Create a new instance of Address1 for User1
	address1 := createAddress(19, "Good Road", "Good Street", "Good Country")
	fmt.Println("User1 address is", address1)
	//Create a new instance of Address2 for User2
	address2 := createAddress(19, "Great Road", "Great Street", "Great Country")
	fmt.Println("User2 address is", address2)

	//Create a new instance of User1
	user1 := createUser("Jace1", 25, true, "Cake", 5.99, 1.99, address1)
	fmt.Println("User1's details are", user1)
	//Create a new instance of User2
	user2 := createUser("Jace2", 35, false, "Icecream", 7.99, 2.99, address2)
	fmt.Println("User2's details are", user2)

	//Print User1's address invoking printAddress() function
	user1.printAddress()
	//Print User2's address invoking printAddress() function
	user2.printAddress()

	// Print User's name
	fmt.Println("User1's name is", user1.Name)
	// Print User's age
	fmt.Println("User1's name is", user1.Age)
	// Print User1's Remote status
	fmt.Println("The remote status for User1 is", user1.IsRemote)
	// Print User1's Items
	fmt.Println("User1's items are", user1.Items)
	// Print User1's address
	fmt.Println("User1's address is", user1.Address)
	// Print User1's country
	fmt.Println("User1's country is", user1.Address.Country)

	//Add item
	user1.addItem("Pizza", 13.99)
	//Update tip
	user1.updateTip(19.99)
	fmt.Println("User1's details with added ITEM  and updated TIP are", user1)
	fmt.Println()

	//Add item again
	user1.addItem("Pasta", 4.99)
	//Update tip again
	user1.updateTip(23.33)
	fmt.Println("User1's details with added ITEM AGAIN and updated TIP AGAIN are", user1)

	// Creating a ANONYMOUS struct
	job := struct {
		title  string
		salary int
	}{
		title:  "Sofware Engineer",
		salary: 100000,
	}

	fmt.Println("Job", job.title)

	fmt.Println()

	// Serializing User1 JSON Data or User1 struct
	jsonData, _ := json.Marshal(user1)
	fmt.Println(jsonData) //bytes representation of JSON data
	fmt.Println()
	fmt.Println(string(jsonData)) //convert bytes representaion of JSON data into string

	//-----------Using goroutines.go--------------------

	//Calculate the execution time from here
	now := time.Now()

	fmt.Println()

	// Get User Name is going to take 1 second
	userName := getUserByName("John")
	fmt.Println(userName)

	// Create a channel for below two Goroutines which takes in a struct of type Message. The channel has a buffer size of 2.
	ch := make(chan *Message, 2)

	// Add all the goroutines to the the waitgroup
	wg.Add(2)

	// Goroutine to get User Chats is going to take 2 seconds
	go getUserChats("John", ch, wg)

	// Goroutine to get User Friends is going to take 3 seconds
	go getUserFriends("John", ch, wg)

	// Wait untill the above goroutines are done. Blocks until the WaitGroup counter is zero.
	wg.Wait()

	// Close the channel
	close(ch)

	// Consume the message from the channel
	for msg := range ch {
		fmt.Println(msg)
	}

	// Show the execution time
	fmt.Println(time.Since(now))

	//-----------Using interfaces.go--------------------

	fmt.Println()

	// Create a new circle1 and circle2
	c1 := createCircle(3)
	c2 := createCircle(6)

	// Area of circle1 and circle2
	fmt.Println()
	fmt.Printf("The area of circle 1 is %v and the area of circle 2 is %v\n", c1.Area, c2.Area)
	fmt.Println()

	// Update radius for circle1 and circle2

	// Area of circle1 and circle2 with updated radius

	// Circumference of circle1 and circle2
	fmt.Println()
	fmt.Printf("The circumference of circle 1 is %v and the circumference of circle 2 is %v\n", c1.circumference(), c2.circumference())
	fmt.Println()

	// Circumference of circle1 and circle2 using INTERFACE
	fmt.Println()
	fmt.Printf("The circumference of circle 1 USING INTERFACE is %v and the circumference of circle 2 USING INTERFACE is %v\n", calculateCircumference(&c1), calculateCircumference(&c2))
	fmt.Println()

	// Create a new rectangle1 and recatange2
	fmt.Println()
	r1 := createRetangle(3, 3)
	r2 := createRetangle(9, 9)
	fmt.Println()

	// Area of rectangel and reactangle2
	fmt.Printf("The area of rectange 1 is %v and the area of rectange 2 is %v\n", r1.Area, r2.Area)

	// Circumference of rectangle1 and rectangle2
	fmt.Println()
	fmt.Printf("The circumference of rectangle 1 is %v and the circumference of rectangel 2 is %v\n", r1.circumference(), r2.circumference())
	fmt.Println()

	// Circumference of rectangle1 and rectangle2 using INTERFACE
	fmt.Printf("The circumference of rectangle 1 USING INTERFACE is %v and the circumference of rectangle 2 USING INTERFACE is %v\n", calculateCircumference(&r1), calculateCircumference(&r2))
	fmt.Println()

	//-----------Using slices.go--------------------

	//List the integers in the sliceInts
	fmt.Println()
	for k, v := range sliceInts {
		fmt.Printf("The list of integers in sliceInts is%v:%v\n", k, v)
	}
	fmt.Println()

	//List the floats in sliceFloats
	fmt.Println()
	for k, v := range sliceFloats {
		fmt.Printf("The list of floats in sliceFloats is%v:%v\n", k, v)
	}
	fmt.Println()

	//List the strings in sliceStrings
	fmt.Println()
	for k, v := range sliceStrings {
		fmt.Printf("The list of strings in sliceStrings is %v:%v \n", k, v)
	}
	fmt.Println()

	//List of maps in sliceMaps
	fmt.Println()
	for k, v := range sliceMaps {
		fmt.Printf("The list of maps in sliceMaps is %v:%v\n", k, v)
	}
	fmt.Println()

	//List of structs in sliceStructs
	fmt.Println()
	for k, v := range sliceStructs {
		fmt.Printf("The list of structs in sliceStructs is %v:%v\n", k, v)
	}
	fmt.Println()

	//List of interface structs in sliceInterface
	fmt.Println()
	for k, v := range sliceInterface {
		fmt.Printf("The list of structs of type interface in sliceInterface is %v:%v\n", k, v)
	}
	fmt.Println()

	// Apply a function on the list of interface structs in sliceInteface
	fmt.Println()
	for _, v := range sliceInterface {
		printShapeInfo(v)
		fmt.Println("---")
	}

	//-----------Using goroutines2.go--------------------

	// declare a WaitGroup variable
	var wg sync.WaitGroup

	// Make a channel with appropriate buffer size for goroutines to pass data
	c := make(chan string, 2)

	// Add all the goroutines to the the waitgroup
	wg.Add(2)

	// Goroutine to get worker1
	go worker1(c, &wg)

	// Goroutine to get worker2
	go worker2(c, &wg)

	go func() {
		wg.Wait() // Wait untill the above goroutines are done. Blocks until the WaitGroup counter is zero.
		close(c)  // Close the channel
	}()

	// Wait untill the above goroutines are done. Blocks until the WaitGroup counter is zero.
	//wg.Wait()

	// Close the channel
	//close(c)

	// Consume the message from the channel
	for msg := range c {
		fmt.Println(msg)
	}

	// Consume the message from the channel
	//output := <-c
	//fmt.Println("Output from channel is", output)

}
