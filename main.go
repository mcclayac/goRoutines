package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Go Routines")
	cars := fillCars()
	go showCars(cars, "First Go Routine")
	go showCars(cars, "Second Go Routine")
	showCars(cars, "No Go Routine")
	//showCars(cars, "Fourth Go Routine")
	//showCars(cars, "Fifth Go Routine")

	go func(msg string) {
		fmt.Println(msg)
	} ("going")

	//time.Sleep(2 * time.Second)
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

type Cars map[string] int

func fillCars() map[string] int {
	cars := make(map[string] int)
	cars["Jeep"] = 23
	cars["Buick"] = 11
	cars["Ford"] = 15
	cars["Chevy"] = 9
	cars["Nissan"] = 29
	return cars
}

func showCars(c Cars, m string) {
	for key, i := range c {
		fmt.Fprintf(os.Stdout, "[%[1]v] %[2]v = %[3]v %[4]v \n",m,i,key,c[key])

	}

}
