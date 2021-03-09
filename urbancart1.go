package main

import (
	"fmt"
)

type user struct {
	mobno int
	place string
	or    order
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

func main() {
	var ur [100]user
	var choice int
LABEL:
	fmt.Printf("\n Hi user! what would you like to do?")
	fmt.Printf("\n 1. adding the user 2. display list of products avaliable 3. place order 4. change your order 5. cancel order 6. get you order summary 7. exit")
	fmt.Printf("\n Enter your choice: ")
	fmt.Scanln(&choice)
	if choice == 1 {
		var n int
		var mob int
		fmt.Printf("\n Enter the user ID: ")
		fmt.Scanln(&n)
		fmt.Printf("\n enter the mobile number: ")
		fmt.Scanln(&mob)
		for i := 0; i < 100; i++ {
			if mob == ur[i].mobno {
				fmt.Printf("\n User already present, Not successful")
				goto LABEL
			}
		}
		ur[n].mobno = mob
		fmt.Printf("\n enter the place: ")
		fmt.Scanln(&ur[n].place)
		fmt.Printf("\n User added successfully")
		goto LABEL
	} else if choice == 2 {
		fmt.Printf("\n 1.pen \n 2.pencil \n 3.ball \n 4. mobile \n 5. bowl ")
		goto LABEL
	} else if choice == 3 {
		var n int
		fmt.Printf("Enter the user id: \n")
		fmt.Scanln(&n)
		for i := 1; i < 6; i++ {
			fmt.Printf("Enter the %d product quantity: \n", i)
			fmt.Scanln(&ur[n].or.prod[i].qty.no)
		}
		goto LABEL
	} else if choice == 4 {
		var n, i, j int
		fmt.Printf("Enter the user id: \n")
		fmt.Scanln(&n)
		fmt.Printf("\n Enter the product ID and quantity you want to change")
		fmt.Printf("\n Product ID: ")
		fmt.Scanln(&i)
		fmt.Printf("\n Quantity: ")
		fmt.Scanln(&j)
		ur[n].or.prod[i].qty.no = j
		goto LABEL
	} else if choice == 5 {
		var n int
		fmt.Printf("Enter the user id: \n")
		fmt.Scanln(&n)
		for i := 1; i < 6; i++ {
			ur[n].or.prod[i].qty.no = 0
		}
		goto LABEL
	} else if choice == 6 {
		var n int
		fmt.Printf("Enter the user id: \n")
		fmt.Scanln(&n)
		for i := 1; i < 6; i++ {
			fmt.Printf("\n The %d product quantity is : %d ", i, ur[n].or.prod[i].qty.no)
		}
		goto LABEL
	} else if choice == 7 {
		fmt.Printf("\n Come back again! exiting")
		goto LABEL1
	} else {
		fmt.Printf("\n Incorrect input entered")
		goto LABEL
	}
LABEL1:
}
