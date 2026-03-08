package main

import (
	"fmt"
)

type Vehicle struct {
	Lat float64
	Lng float64
}

/* New/Factory functions form
func NewVehicle(lat, lng float64) Vehicle {
func NewVehicle(lat, lng float64) (Vehicle, error) {
func NewVehicle(lat, lng float64) *Vehicle {
func NewVehicle(lat, lng float64) (*Vehicle, error) {
*/

func validateLatLng(lat, lng float64) error {
	if lat < -90 || lat > 90 {
		return fmt.Errorf("invalid lat: %#v", lat)
	}

	if lng < -180 || lng > 180 {
		return fmt.Errorf("invalid lng: %#v", lng)
	}

	return nil
}

func NewVehicle(lat, lng float64) (*Vehicle, error) {
	if err := validateLatLng(lat, lng); err != nil {
		return nil, err
	}

	return &Vehicle{Lat: lat, Lng: lng}, nil
}

// v is called "the reciever"
func (v *Vehicle) Move(lat, lng float64) error {
	if err := validateLatLng(lat, lng); err != nil {
		return err
	}

	v.Lat = lat
	v.Lng = lng
	return nil
}

type Car struct {
	Vehicle
	LicensePlate string
}

type Mover interface {
	Move(lat, lng float64) error
}

func MoveAll(movers []Mover, lat, lng float64) error {
	for _, m := range movers {
		if err := m.Move(lat, lng); err != nil {
			return err
		}
	}

	return nil
}

/*
- Interface say what we need, not what we provide
- Interfaces are small (stdlib avg < 2), if you have more than 4 - rethink
- Rule of thumb: accept interface, return types
- Start with types, discover interfaces
*/

func main() {
	v := Vehicle{
		Lng: 34.7818,
		Lat: 32.0853,
	}
	fmt.Println(v)
	fmt.Printf("v : %v\n", v)
	fmt.Printf("+v: %+v\n", v)
	fmt.Printf("#v: %#v\n", v)
	v.Lat += 0.02

	fmt.Printf("#v: %#v\n", Vehicle{1, 2})
	fmt.Println(NewVehicle(10, 20))
	fmt.Println(NewVehicle(10, 200))

	var v2 Vehicle
	fmt.Println("v2:", v2)

	if err := v2.Move(40, 50); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("v2 (move): %#v\n", v2)
	}

	c := Car{
		LicensePlate: "G0PH3R",
	}
	fmt.Printf("c: %#v\n", c)
	fmt.Println("c.Lat:", c.Lat)
	if err := c.Move(10, 20); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("c (move): %#v\n", c)
	}

	v3, _ := NewVehicle(1, 2)
	c2 := &Car{LicensePlate: "MVALL"}
	movers := []Mover{v3, c2, &c}
	if err := MoveAll(movers, 55, 37); err != nil {
		fmt.Println("MoveAll ERROR:", err)
	} else {
		fmt.Printf("v3 after MoveAll: %#v\n", v3)
		fmt.Printf("c2 after MoveAll: %#v\n", c2)
	}
}
