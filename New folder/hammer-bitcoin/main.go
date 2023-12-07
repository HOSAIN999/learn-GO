// package main

// import (
//
//	"fmt"
//	"math/rand"
//	"time"
//
// )

// type animals interface {
// 	says() string
// 	HowManyLegs() int
// }
// type dog struct {
// 	Name        string
// 	Sound       string
// 	numberOfLeg int
// }

// func (d *dog) says() string {
// 	return d.Sound
// }
// func (d *dog) HowManyLegs() int {
// 	return d.numberOfLeg
// }

// type cat struct {
// 	Name        string
// 	Sound       string
// 	numberOfLeg int
// 	Hastail     bool
// }

//	func (c *cat) HowManyLegs() int {
//		return c.numberOfLeg
//	}
//
//	func (c *cat) says() string {
//		return c.Sound
//	}
//
//	func main() {
//		Dog := dog{
//			Name:        "dog",
//			Sound:       "Woof",
//			numberOfLeg: 4,
//		}
//		Riddle(&Dog)
//		Cat := cat{
//			Name:        "cat",
//			Sound:       "meow",
//			numberOfLeg: 4,
//			Hastail:     true,
//		}
//		Riddle(&Cat)
//	}
//
//	func Riddle(a animals) {
//		riddle := fmt.Sprintf(`this is animal says "%s" and has "%d" leg. What animal is it`, a.says(), a.HowManyLegs())
//		fmt.Println(riddle)
//	}
// type Vehicle struct {
// 	numberOfWheels    int
// 	NumberOfPassenger int
// }
// type Car struct {
// 	Make       string
// 	Model      string
// 	Year       int
// 	IsElectric bool
// 	IsHybrid   bool
// 	Vehicle    Vehicle
// }

//	func (v Vehicle) showDetails() {
//		fmt.Println("Number of Passenger", v.NumberOfPassenger)
//		fmt.Println("Number of Wheels", v.numberOfWheels)
//	}
//
//	func (c Car) show() {
//		fmt.Println("Make:", c.Make)
//		fmt.Println("Model:", c.Model)
//		fmt.Println("Year:", c.Year)
//		fmt.Println("IsElectric:", c.IsElectric)
//		fmt.Println("IsHybrid:", c.IsHybrid)
//		c.Vehicle.showDetails()
//	}
//
//	func main() {
//		suv := Vehicle{
//			numberOfWheels:    4,
//			NumberOfPassenger: 6,
//		}
//		volvoXC90 := Car{
//			Make:       "Volvo",
//			Model:      "XC90 T8",
//			Year:       2021,
//			IsElectric: false,
//			IsHybrid:   true,
//			Vehicle:    suv,
//		}
//		volvoXC90.show()
//		fmt.Println()
//		teslaModelX := Car{
//			Make:       "Tesla",
//			Model:      "Model X",
//			Year:       2022,
//			IsElectric: true,
//			IsHybrid:   true,
//			Vehicle:    suv,
//		}
//		teslaModelX.Vehicle.NumberOfPassenger = 7
//		teslaModelX.Vehicle.numberOfWheels = 6
//		teslaModelX.show()
//	}
