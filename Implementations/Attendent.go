package Implementations

import "errors"

type Attendent struct {
	AssignedParkingLots []*ParkingLot
	ParkedCars          []*Car
	NextLotStrategy     NextLotStrategy
	AssignedOwner       *Owner
}

func AttendentConstruct(strategy NextLotStrategy) *Attendent {
	return &Attendent{
		AssignedParkingLots: []*ParkingLot{},
		ParkedCars:          []*Car{},
		NextLotStrategy:     strategy,
	}
}

func AttendentConstructDefault() *Attendent {
	return &Attendent{
		AssignedParkingLots: []*ParkingLot{},
		ParkedCars:          []*Car{},
		NextLotStrategy:     &NormalNextLotStrategy{},
	}
}

func (attendent *Attendent) Assign(parkingLot *ParkingLot, owner *Owner) error {
	if attendent.AssignedOwner != nil && attendent.AssignedOwner != owner {
		return errors.New("this parking lot is not owned by the owner")
	}
	for _, lot := range attendent.AssignedParkingLots {
		if lot == parkingLot {
			return errors.New("parking lot already assigned")
		}
	}
	if attendent.AssignedOwner == nil {
		attendent.AssignedOwner = owner
	}
	attendent.AssignedParkingLots = append(attendent.AssignedParkingLots, parkingLot)
	return nil
}

func (attendent *Attendent) Park(car *Car) (*Ticket, error) {
	if len(attendent.AssignedParkingLots) == 0 {
		return nil, errors.New("no parking lot assigned")
	}
	if err := attendent.CheckIfCarIsAlreadyParked(car); err != nil {
		return nil, err
	}

	selectedLot, err := attendent.NextLotStrategy.GetNextLot(attendent.AssignedParkingLots)
	if err != nil {
		return nil, err
	}

	ticket, _ := selectedLot.Park(car)
	attendent.ParkedCars = append(attendent.ParkedCars, car) // Remember to add parked car
	return ticket, nil
}

func (attendent *Attendent) CheckIfCarIsAlreadyParked(car *Car) error {
	for _, parkedCar := range attendent.ParkedCars {
		if parkedCar == car {
			return errors.New("car already assigned to this parking lot")
		}
	}
	return nil
}

func (attendent *Attendent) Unpark(ticket *Ticket) (*Car, error) {
	for _, lot := range attendent.AssignedParkingLots {
		unparkedCar, err := lot.Unpark(ticket)
		if err == nil {
			// Remove the car from the parked cars list
			for i, parkedCar := range attendent.ParkedCars {
				if parkedCar == unparkedCar {
					attendent.ParkedCars = append(attendent.ParkedCars[:i], attendent.ParkedCars[i+1:]...)
					break
				}
			}
			return unparkedCar, nil
		}
	}
	return nil, errors.New("car not found")
}
