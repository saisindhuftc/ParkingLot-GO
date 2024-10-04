package Implementations

import (
	"ParkingLot_go/Enums"
	"errors"
	"github.com/google/uuid"
	"math/big"
)

type ParkingLot struct {
	totalSlots   int
	slots        []*Slot
	ParkingLotId int
	notifiables  []Notifiable
	Owner        *Owner
	isFull       bool
}

func NewParkingLot(totalSlots int, owner *Owner) *ParkingLot {
	if totalSlots <= 0 {
		panic(errors.New("parking lot size must be positive"))
	}
	if owner == nil {
		panic(errors.New("parking lot cannot be created without an owner"))
	}
	uuidValue := uuid.New()
	lot := &ParkingLot{
		totalSlots:   totalSlots,
		Owner:        owner,
		ParkingLotId: uuidToInt(uuidValue),
		notifiables:  []Notifiable{},
		slots:        make([]*Slot, totalSlots),
	}
	for i := 0; i < totalSlots; i++ {
		lot.slots[i] = NewSlot()
	}
	return lot
}

func uuidToInt(u uuid.UUID) int {
	i := new(big.Int)
	i.SetString(u.String(), 16)
	return int(i.Int64())
}

func (parkinglot *ParkingLot) findNearestSlot() (*Slot, error) {
	for _, slot := range parkinglot.slots {
		if slot.IsFree() {
			return slot, nil
		}
	}
	return nil, errors.New("parking lot is full")
}

func (parkinglot *ParkingLot) Park(car *Car) *Ticket {
	if parkinglot.IsFull() {
		panic(errors.New("parking lot is full"))
	}
	if parkinglot.IsCarAlreadyParked(*car) {
		panic(errors.New("car is already parked"))
	}
	slot, _ := parkinglot.findNearestSlot()
	ticket, _ := slot.Park(*car)
	if parkinglot.IsFull() {
		parkinglot.isFull = true
		parkinglot.notifyFull()
	}
	return ticket
}

func (parkinglot *ParkingLot) Unpark(ticket *Ticket) (*Car, error) {
	for _, slot := range parkinglot.slots {
		car, err := slot.Unpark(ticket)
		if err == nil {
			if !parkinglot.IsFull() {
				parkinglot.isFull = false // Notify availability if it was previously full
				parkinglot.notifyAvailable()
			}
			if car != nil {
				return car, nil
			}
		}
	}
	return nil, errors.New("car not found in the parking lot")
}

func (parkinglot *ParkingLot) IsCarAlreadyParked(car Car) bool {
	for _, slot := range parkinglot.slots {
		if slot.CheckingCarInParkingSlot(car) {
			return true
		}
	}
	return false
}

func (parkinglot *ParkingLot) IsFull() bool {
	for _, slot := range parkinglot.slots {
		if slot.IsFree() {
			return false
		}
	}
	return true
}

func (parkinglot *ParkingLot) CountCarsByColor(color Enums.Color) int {
	count := 0
	for _, slot := range parkinglot.slots {
		if slot.HasCarOfColor(color) {
			count++
		}
	}
	return count
}

func (parkinglot *ParkingLot) IsCarWithRegistrationNumberParked(registrationNumber string) (bool, error) {
	for _, slot := range parkinglot.slots {
		if slot.HasCarWithRegistrationNumber(registrationNumber) {
			return true, nil
		}
	}
	return false, errors.New("car needs registration number")
}

func (parkinglot *ParkingLot) CountParkedCars() int {
	count := 0
	for _, slot := range parkinglot.slots {
		if !slot.IsFree() {
			count++
		}
	}
	return count
}

func (parkinglot *ParkingLot) notifyFull() {
	for _, notifiable := range parkinglot.notifiables {
		notifiable.notifyFull(parkinglot.ParkingLotId)
	}
}

func (parkinglot *ParkingLot) notifyAvailable() {
	for _, notifiable := range parkinglot.notifiables {
		notifiable.notifyAvailable(parkinglot.ParkingLotId)
	}
}

func (parkinglot *ParkingLot) RegisterNotifiable(notifiable Notifiable) {
	parkinglot.notifiables = append(parkinglot.notifiables, notifiable)
}
