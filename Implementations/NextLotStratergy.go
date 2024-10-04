package Implementations

type NextLotStrategy interface {
	GetNextLot(assignedParkingLots []*ParkingLot) (*ParkingLot, error)
}
