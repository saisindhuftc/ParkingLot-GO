package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Exceptions"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicemanNotifiedWhenParkingLotIsFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(1)
	policeman := Implementations.PolicemanConstruct()
	owner.RegisterNotifiable(parkingLot, policeman)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, err := parkingLot.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())
}

func TestRegisterNotifiableToPolicemen(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	policeman := Implementations.PolicemanConstruct()
	parkingLot := owner.CreateParkingLot(1)
	attendent := Implementations.AttendentConstructDefault()

	err := owner.AssignParkingLotToAttendent(attendent, parkingLot)
	assert.NoError(t, err)
	owner.RegisterNotifiable(parkingLot, policeman)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, err = attendent.Park(car)
	assert.NoError(t, err)
}

func TestPolicemanNotifyFullAllParkingLotsAreFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(1)
	policeman := Implementations.PolicemanConstruct()

	owner.RegisterNotifiable(firstLot, policeman)
	owner.RegisterNotifiable(secondLot, policeman)

	_, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-1432", Color: Enums.YELLOW})
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())
}

func TestPolicemanNotifyFullSomeParkingLotsAreFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(2)
	policeman := Implementations.PolicemanConstruct()

	owner.RegisterNotifiable(firstLot, policeman)
	owner.RegisterNotifiable(secondLot, policeman)

	_, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())
}

func TestPolicemanNotifiedWhenParkingLotAvailable(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(1)
	policeman := Implementations.PolicemanConstruct()

	owner.RegisterNotifiable(parkingLot, policeman)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	ticket, err := parkingLot.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Unpark(ticket)
	assert.NoError(t, err)

	_, err = parkingLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)
}

func TestPolicemanNotifyAvailableSecondParkingLotIsAvailable(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(1)
	policeman := Implementations.PolicemanConstruct()

	owner.RegisterNotifiable(firstLot, policeman)
	owner.RegisterNotifiable(secondLot, policeman)

	_, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	secondTicket, err := secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = secondLot.Unpark(secondTicket)
	assert.NoError(t, err)

	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.NoError(t, err)
}

func TestPolicemanNotifyFullSomeParkingLotsAreAvailable(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(2)
	policeman := Implementations.PolicemanConstruct()

	owner.RegisterNotifiable(firstLot, policeman)
	owner.RegisterNotifiable(secondLot, policeman)

	firstTicket, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())

	_, err = firstLot.Unpark(firstTicket)
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-8888", Color: Enums.BLACK})
	assert.NoError(t, err)
}
