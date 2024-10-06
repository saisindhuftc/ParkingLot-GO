package Implementations

import (
	"ParkingLot_go/Exceptions"
	"errors"
	"fmt"
)

type Owner struct {
	Attendents       []*Attendent
	OwnerParkingLots []*ParkingLot
	notifiables      []Notifiable
	Attendent
}

func OwnerConstruct() *Owner {
	return &Owner{
		Attendents:       []*Attendent{},
		OwnerParkingLots: []*ParkingLot{},
		Attendent:        *AttendentConstructDefault(),
	}
}

func (owner *Owner) CreateParkingLot(totalSlots int) *ParkingLot {
	if totalSlots <= 0 {
		panic(Exceptions.ErrCannotCreateParkingLotException)
	}
	parkingLot := ParkingLotConstruct(totalSlots, owner)
	parkingLot.RegisterNotifiable(owner)
	owner.OwnerParkingLots = append(owner.OwnerParkingLots, parkingLot)
	return parkingLot
}

func (owner *Owner) AssignParkingLotToAttendent(attendent *Attendent, parkingLot *ParkingLot) error {
	isOwnedByThisOwner := false
	for _, ownerLot := range owner.OwnerParkingLots {
		if ownerLot == parkingLot {
			isOwnedByThisOwner = true
			break
		}
	}
	if !isOwnedByThisOwner {
		return errors.New("this parking lot is not owned by the owner")
	}
	return attendent.Assign(parkingLot, owner)
}

func (owner *Owner) AssignParkingLotToSelf(parkingLot *ParkingLot) error {
	for _, ownerLot := range owner.OwnerParkingLots {
		if ownerLot == parkingLot {
			return owner.Assign(parkingLot, owner)
		}
	}
	return errors.New("this parking lot is not owned by this owner")
}

func contains(lots []*ParkingLot, lot *ParkingLot) bool {
	for _, item := range lots {
		if item == lot {
			return true
		}
	}
	return false
}

func (owner *Owner) RegisterNotifiable(parkingLot *ParkingLot, notifiable Notifiable) {
	parkingLot.RegisterNotifiable(notifiable)
}

func (owner *Owner) notifyFull(parkingLotId int) {
	fmt.Printf("Owner notified: Parking lot with ID %d is full.\n", parkingLotId)
}

func (owner *Owner) notifyAvailable(parkingLotId int) {
	fmt.Printf("Owner notified: Parking lot with ID %d has available slots.\n", parkingLotId)
}
