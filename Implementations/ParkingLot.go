package Implementations

import (
	"ParkingLot_go/Enums"
	"errors"
)

type ParkingLot struct {
	totalSlots int
	slots      []*Slot
	ID         bool
}

// Equal method to compare two ParkingLot instances
func (p *ParkingLot) Equal(other ParkingLot) bool {
	return p.ID == other.ID // or compare based on other relevant fields
}

func NewParkingLot(totalSlots int) (*ParkingLot, error) {
	if totalSlots <= 0 {
		return nil, errors.New("parking lot size must be positive")
	}
	lot := &ParkingLot{
		totalSlots: totalSlots,
		slots:      make([]*Slot, totalSlots),
	}
	for i := 0; i < totalSlots; i++ {
		lot.slots[i] = &Slot{}
	}
	return lot, nil
}

func (p *ParkingLot) findNearestSlot() (*Slot, error) {
	for _, slot := range p.slots {
		if slot.IsFree() {
			return slot, nil
		}
	}
	return nil, errors.New("parking lot is full")
}

func (p *ParkingLot) Park(car Car) (*Ticket, error) {
	if p.IsFull() {
		return nil, errors.New("parking lot is full")
	}
	if p.IsCarAlreadyParked(car) {
		return nil, errors.New("car is already parked")
	}
	slot, err := p.findNearestSlot()
	if err != nil {
		return nil, err
	}
	return slot.Park(car)
}

func (p *ParkingLot) Unpark(ticket *Ticket) (*Car, error) {
	for _, slot := range p.slots {
		if car, err := slot.Unpark(ticket); err == nil {
			return car, nil
		}
	}
	return nil, errors.New("car not found in the parking lot")
}

func (p *ParkingLot) IsCarAlreadyParked(car Car) bool {
	for _, slot := range p.slots {
		if slot.CheckingCarInParkingSlot(car) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) IsFull() bool {
	for _, slot := range p.slots {
		if slot.IsFree() {
			return false
		}
	}
	return true
}

func (p *ParkingLot) CountCarsByColor(color Enums.Color) int {
	count := 0
	for _, slot := range p.slots {
		if slot.HasCarOfColor(color) {
			count++
		}
	}
	return count
}

func (p *ParkingLot) IsCarWithRegistrationNumberParked(registrationNumber string) (bool, error) {
	for _, slot := range p.slots {
		if slot.HasCarWithRegistrationNumber(registrationNumber) {
			return true, nil
		}
	}
	return false, errors.New("car needs registration number")
}
