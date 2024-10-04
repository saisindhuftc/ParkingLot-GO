package Implementations

import "fmt"

type Policeman struct{}

func NewPoliceman() *Policeman {
	return &Policeman{}
}

func (p *Policeman) NotifyFull(parkingLotId int) {
	fmt.Printf("Policeman notified: Parking lot with ID %d is full.\n", parkingLotId)
}

func (p *Policeman) NotifyAvailable(parkingLotId int) {
	fmt.Printf("Policeman notified: Parking lot with ID %d has available slots.\n", parkingLotId)
}
