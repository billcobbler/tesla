package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/billcobbler/tesla"
)

func main() {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     os.Getenv("TESLA_CLIENT_ID"),
			ClientSecret: os.Getenv("TESLA_CLIENT_SECRET"),
			Email:        os.Getenv("TESLA_USERNAME"),
			Password:     os.Getenv("TESLA_PASSWORD"),
		})
	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", vehicles)

	vehicle := vehicles[0]
	_, err = vehicle.MobileEnabled()
	if err != nil {
		panic(err)
	}

	charge, _ := vehicle.ChargeState()
	climate, _ := vehicle.ClimateState()
	drive, _ := vehicle.DriveState()
	gui, _ := vehicle.GuiSettings()
	car, _ := vehicle.VehicleState()
	fmt.Println("Charge State")
	fmt.Println(prettyPrint(charge))
	fmt.Println("Climate State")
	fmt.Println(prettyPrint(climate))
	fmt.Println("Drive State")
	fmt.Println(prettyPrint(drive))
	fmt.Println("GUI State")
	fmt.Println(prettyPrint(gui))
	fmt.Println("Vehicle State")
	fmt.Println(prettyPrint(car))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
