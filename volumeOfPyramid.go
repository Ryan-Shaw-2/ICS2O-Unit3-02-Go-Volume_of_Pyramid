// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program calculates the volume of a pyramid

package main

import "fmt"

func main() {
	// This function calculates the volume of a pyramid
	var length int
	var width int
	var height int

	// input
	fmt.Println("This program calculates the volume of a pyramid")
	fmt.Println()
	fmt.Print("Enter the length of the base (cm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width of the base (cm): ")
	fmt.Scanln(&width)
	fmt.Print("Enter the height of the pyramid (cm): ")
	fmt.Scanln(&height)
	fmt.Println()

	// process
	var volume = (length * width * height) / 3

	// output
	fmt.Println("The volume is:", volume, "cm³")
}
