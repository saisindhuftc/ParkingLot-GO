package Implementations

type Attendable interface {
	Assign(parkingLot *ParkingLot) error
	Park(car *Car) (*Ticket, error)
	CheckIfCarIsAlreadyParked(car *Car) error
	Unpark(ticket *Ticket) (*Car, error)
}
