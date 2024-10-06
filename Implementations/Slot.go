package Implementations

import (
	"ParkingLot_go/Enums"
	"errors"
)

type Slot struct {
	car    *Car
	ticket *Ticket
}

func SlotConstruct() *Slot {
	return &Slot{
		car:    nil,
		ticket: nil,
	}
}

func (s *Slot) IsFree() bool {
	return s.car == nil
}

func (s *Slot) Park(car Car) (*Ticket, error) {
	if !s.IsFree() {
		return nil, errors.New("slot is already occupied")
	}
	s.car = &car
	s.ticket = TicketConstruct()
	return s.ticket, nil
}

func (s *Slot) Unpark(ticket *Ticket) (*Car, error) {
	if s.IsFree() {
		return nil, errors.New("car not found in the slot")
	}
	if s.ticket.Equals(ticket) {
		car := s.car
		s.car = nil
		s.ticket = nil
		return car, nil
	}
	return nil, errors.New("invalid ticket")
}

func (s *Slot) HasCarOfColor(color Enums.Color) bool {
	return !s.IsFree() && s.car.IsColor(color)
}

func (s *Slot) HasCarWithRegistrationNumber(registrationNumber string) bool {
	return !s.IsFree() && s.car.HasRegistrationNumber(registrationNumber)
}

func (s *Slot) CheckingCarInParkingSlot(car Car) bool {
	return !s.IsFree() && *s.car == car
}
