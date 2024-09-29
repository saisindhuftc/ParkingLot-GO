package Exceptions

import "errors"

var (
	ErrCarAlreadyParked           = errors.New("car already parked")
	ErrCarNeedsRegistrationNumber = errors.New("car needs registration number")
	ErrInvalidTicket              = errors.New("invalid ticket")
	ErrParkingLotAlreadyAssigned  = errors.New("parking lot already assigned")
	ErrParkingLotIsFull           = errors.New("parking lot is full")
	ErrSlotIsOccupied             = errors.New("slot is occupied")
)
