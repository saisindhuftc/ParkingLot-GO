package Implementations

import (
	"ParkingLot_go/Enums"
)

type Car struct {
	RegistrationNumber string
	Color              Enums.Color
	LicensePlate       string
}

func NewCar(registrationNumber string, color Enums.Color) Car {
	return Car{
		RegistrationNumber: registrationNumber,
		Color:              color,
	}
}

func (c *Car) Equal(other Car) bool {
	return c.LicensePlate == other.LicensePlate
}

func (c *Car) IsColor(color Enums.Color) bool {
	return c.Color == color
}

func (c *Car) HasRegistrationNumber(registrationNumber string) bool {
	return c.RegistrationNumber == registrationNumber
}
