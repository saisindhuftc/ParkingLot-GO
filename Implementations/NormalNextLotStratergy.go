package Implementations

import (
	"errors"
)

type NormalNextLotStrategy struct{}

func (n *NormalNextLotStrategy) GetNextLot(assignedParkingLots []*ParkingLot) (*ParkingLot, error) {
	for _, lot := range assignedParkingLots {
		if !lot.IsFull() {
			return lot, nil
		}
	}
	return nil, errors.New("all parking lots are full")
}
