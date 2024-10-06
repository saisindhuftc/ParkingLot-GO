package Implementations

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Exceptions"
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

func ParkingLotConstruct(totalSlots int, owner *Owner) *ParkingLot {
	if totalSlots <= 0 {
		panic(Exceptions.ErrCannotCreateParkingLotException)
	}
	if owner == nil {
		panic(Exceptions.ErrParkingLotAlreadyAssigned)
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
		lot.slots[i] = SlotConstruct()
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
	return nil, Exceptions.ErrParkingLotIsFull
}

func (parkinglot *ParkingLot) Park(car *Car) (*Ticket, error) {
	if parkinglot.IsFull() {
		return nil, Exceptions.ErrParkingLotIsFull
	}
	if parkinglot.IsCarAlreadyParked(*car) {
		return nil, Exceptions.ErrCarAlreadyParked
	}
	slot, _ := parkinglot.findNearestSlot()
	ticket, _ := slot.Park(*car)
	if parkinglot.IsFull() {
		parkinglot.isFull = true
		parkinglot.notifyFull()
	}
	return ticket, nil
}

func (parkinglot *ParkingLot) Unpark(ticket *Ticket) (*Car, error) {
	for _, slot := range parkinglot.slots {
		car, err := slot.Unpark(ticket)
		if err == nil {
			if parkinglot.isFull && !parkinglot.IsFull() {
				parkinglot.isFull = false
				parkinglot.notifyAvailable()
			}
			if car != nil {
				return car, nil
			}
		}
	}
	return nil, Exceptions.ErrInvalidTicket
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
	if registrationNumber == "" {
		return false, Exceptions.ErrCarNeedsRegistrationNumber
	}
	for _, slot := range parkinglot.slots {
		if slot.HasCarWithRegistrationNumber(registrationNumber) {
			return true, nil
		}
	}
	return false, nil
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

func (parkinglot *ParkingLot) GetParkingLotId() int {
	return parkinglot.ParkingLotId
}
