package Implementations

import (
	"errors"
)

type Attendant struct {
	assignedParkingLots []ParkingLot
	parkedCars          []Car
}

func NewAttendent() *Attendant {
	return &Attendant{
		assignedParkingLots: []ParkingLot{},
		parkedCars:          []Car{},
	}
}

func (a *Attendant) Assign(parkingLot ParkingLot) error {
	for _, lot := range a.assignedParkingLots {
		if lot.Equal(parkingLot) {
			return errors.New("parking lot already assigned")
		}
	}
	a.assignedParkingLots = append(a.assignedParkingLots, parkingLot)
	return nil
}

func (a *Attendant) Park(car Car) (*Ticket, error) {
	for _, c := range a.parkedCars {
		if c.Equal(car) { // Use the Equal method for Car
			return nil, errors.New("car already assigned to this parking lot")
		}
	}
	for _, lot := range a.assignedParkingLots {
		if !lot.IsFull() {
			a.parkedCars = append(a.parkedCars, car)
			return lot.Park(car)
		}
	}
	return nil, errors.New("all parking lots are full")
}

func (a *Attendant) Unpark(ticket *Ticket) (*Car, error) {
	for _, lot := range a.assignedParkingLots {
		unparkedCar, err := lot.Unpark(ticket)
		if err == nil {
			for i, c := range a.parkedCars {
				if c.Equal(*unparkedCar) { // Use the Equal method for Car
					a.parkedCars = append(a.parkedCars[:i], a.parkedCars[i+1:]...)
					break
				}
			}
			return unparkedCar, nil
		}
	}
	return nil, errors.New("car not found in assigned parking lot")
}
