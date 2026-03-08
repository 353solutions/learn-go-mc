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

func NewVehicle(lat, lng float64) (Vehicle, error) {
	if lat < -90 || lat > 90 {
		return Vehicle{}, fmt.Errorf("invalid lat: %#v", lat)
	}

	if lng < -180 || lng > 180 {
		return Vehicle{}, fmt.Errorf("invalid lng: %#v", lng)
	}

	return Vehicle{Lat: lat, Lng: lng}, nil
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

}
