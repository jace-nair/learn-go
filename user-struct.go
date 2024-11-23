package main

import "fmt"

// User struct
type User struct {
	Name     string             `json:"name"`
	Age      int                `json: "age"`
	IsRemote bool               `json: "isRemote"`
	Items    map[string]float64 `json: "items"`
	Tip      float64            `json: "tip"`
	Address
}

// Function to create new user
func createUser(name string, age int, isRemote bool, itemName string, itemPrice float64, tip float64, address Address) User {
	//create an instance of struct using name, empty map and zero tip and returns the newly created myStruct
	u := User{
		Name:     name,
		Age:      age,
		IsRemote: isRemote,
		Items:    map[string]float64{itemName: itemPrice},
		Tip:      tip,
		Address:  address,
	}
	//return the new user
	return u
}

// Receiver function to update user name
func (u *User) updateName(name string) {
	u.Name = name
}

// Receiver function to update user age
func (u *User) updateAge(age int) {
	u.Age = age
}

// Receiver function to update user availability
func (u *User) updateIsRemote(isRemote bool) {
	u.IsRemote = isRemote
}

// Receiver function to add user items
func (u *User) addItem(itemName string, itemPrice float64) {
	u.Items[itemName] = itemPrice
}

// Receiver function to update user tip
func (u *User) updateTip(tip float64) {
	u.Tip = tip
}

// Reciever function to update user address
func (u *User) updateAddress(address Address) {
	u.Address = address
}

// Receiver function to update all user details
func (u *User) updateUser(name string, age int, isRemote bool, itemName string, itemPrice float64, tip float64, address Address) {
	u.Name = name
	u.Age = age
	u.IsRemote = isRemote
	u.Items[itemName] = itemPrice
	u.Tip = tip
	u.Address = address
}

// Reciever function to print user details
func (u *User) printUser() {
	fmt.Println(u)
}

// Address struct used as part of the User struct
type Address struct {
	HouseNumber int    `json: "houseNumber"`
	Street      string `json: "street"`
	City        string `json: "city"`
	Country     string `json: "country"`
}

// Function to create address for new user
func createAddress(houseNumber int, street string, city string, country string) Address {
	a := Address{
		HouseNumber: houseNumber,
		Street:      street,
		City:        city,
		Country:     country,
	}
	//return the new address
	return a
}

// Receiver function to update address house number
func (a *Address) updateHouseNumer(houseNumber int) {
	a.HouseNumber = houseNumber
}

// Receiver function to update address street
func (a *Address) updateStreet(street string) {
	a.Street = street
}

// Receiver function to update address city
func (a *Address) updateCity(city string) {
	a.City = city
}

// Receiver function to update address country
func (a *Address) updateCountry(country string) {
	a.Country = country
}

// Receiver function to update all address details
func (a *Address) updateAddress(houseNumber int, street string, city string, country string) {
	a.HouseNumber = houseNumber
	a.Street = street
	a.City = city
	a.Country = country
}

// Reciever function to print user address
func (a *Address) printAddress() {
	fmt.Println("User address is", a)
}
