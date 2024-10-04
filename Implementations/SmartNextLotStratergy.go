package Implementations

import (
	"errors"
)

type SmartNextLotStrategy struct{}

func (s *SmartNextLotStrategy) GetNextLot(assignedParkingLots []*ParkingLot) (*ParkingLot, error) {
	var selectedLot *ParkingLot
	minCars := int(^uint(0) >> 1)

	for _, lot := range assignedParkingLots {
		if !lot.IsFull() && lot.CountParkedCars() < minCars {
			minCars = lot.CountParkedCars()
			selectedLot = lot
		}
	}

	if selectedLot == nil {
		return nil, errors.New("all parking lots are full")
	}
	return selectedLot, nil
}
