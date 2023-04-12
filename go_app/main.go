// Created by: Charlotte Jhu
// Created on: April 2023
//
// This program calculates the volume of a sphere

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function calculates the volume of a sphere
	// variables
	var radius float64

	// input
	fmt.Println("This program calculates the volume of a sphere.")
	fmt.Println()
	fmt.Print("Enter the radius (cm): ")
	fmt.Scanln(&radius)
	fmt.Println()

	// process
	volume := (4.0 / 3) * math.Pi * (math.Pow(radius, 3))

	// output
	fmt.Println("The volume is ", %.2f, volume, " cm³.")
	fmt.Println("\nDone.")
}
