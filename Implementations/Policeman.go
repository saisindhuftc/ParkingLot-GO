package Implementations

import "fmt"

type Policeman struct{}

func PolicemanConstruct() *Policeman {
	return &Policeman{}
}

func (p *Policeman) notifyFull(parkingLotId int) {
	fmt.Printf("Policeman notified: Parking lot with ID %d is full.\n", parkingLotId)
}

func (p *Policeman) notifyAvailable(parkingLotId int) {
	fmt.Printf("Policeman notified: Parking lot with ID %d has available slots.\n", parkingLotId)
}
