package main

import (
	"fmt"
)

type user struct {
	mobno int
	place string
	or    order
}

type userarray struct {
	ur [100]user
}

type order struct {
	prod [6]product
}

type product struct {
	qty quantity
}

type quantity struct {
	no int
}

func (ur *[]userarray) adduser() bool {
	var n int
	var mob int
	fmt.Printf("\n Enter the user ID: ")
	fmt.Scanln(&n)
	fmt.Printf("\n enter the mobile number: ")
	fmt.Scanln(&mob)
	for i := 0; i < 100; i++ {
		if mob == ur[i].mobno {
			fmt.Printf("\n User already present, Not successful")
			return false
		}
	}
	ur[n].mobno = mob
	fmt.Printf("\n enter the place: ")
	fmt.Scanln(&ur[n].place)
	fmt.Printf("\n User added successfully")
	return true
}

func listofproducts() {
	fmt.Printf("\n 1.pen \n 2.pencil \n 3.ball \n 4. mobile \n 5. bowl ")
}

func (ur *userarray) Placeorder() {
	var n int
	fmt.Printf("Enter the user id: \n")
	fmt.Scanln(&n)
	for i := 1; i < 6; i++ {
		fmt.Printf("Enter the %d product quantity: \n", i)
		fmt.Scanln(&ur[n].or.prod[i].qty.no)
	}
}

func (ur *userarray) Changeorder() {
	var n, i, j int
	fmt.Printf("Enter the user id: \n")
	fmt.Scanln(&n)
	fmt.Printf("\n Enter the product ID and quantity you want to change")
	fmt.Printf("\n Product ID: ")
	fmt.Scanln(&i)
	fmt.Printf("\n Quantity: ")
	fmt.Scanln(&j)
	ur[n].or.prod[i].qty.no = j
}

func (ur *userarray) Cancelorder() {
	var n int
	fmt.Printf("Enter the user id: \n")
	fmt.Scanln(&n)
	for i := 1; i < 6; i++ {
		ur[n].or.prod[i].qty.no = 0
	}
}

func (ur *userarray) Getorder() {
	var n int
	fmt.Printf("Enter the user id: \n")
	fmt.Scanln(&n)
	for i := 1; i < 6; i++ {
		fmt.Printf("\n The %d product quantity is : %d ", i, ur[n].or.prod[i].qty.no)
	}
}

func main() {
	adduser()
}
