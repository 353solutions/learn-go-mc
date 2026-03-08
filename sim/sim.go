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

}
