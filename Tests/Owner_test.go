package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Exceptions"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingOwner(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	assert.NotPanics(t, func() {
		owner.CreateParkingLot(2)
	})
}

func TestCreatingParkingLotWithZeroTotalSlots(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	assert.Panics(t, func() {
		owner.CreateParkingLot(0)
	})
}

func TestExceptionWhenOwnerAssignNotOwnedParkingLot(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	otherOwner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := otherOwner.CreateParkingLot(3)

	err := owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	assert.EqualError(t, err, "this parking lot is not owned by the owner")
}

func TestOwnerAssignParkingLotToAttendent(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, firstParkingLot))
}

func TestExceptionAssignParkingLotIsNotOwnedByOwner(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	secondOwner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := secondOwner.CreateParkingLot(3)

	err := owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	assert.EqualError(t, err, "this parking lot is not owned by the owner")
}

func TestOwnerAssignMultipleParkingLotToSingleAttendent(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(3)
	secondParkingLot := owner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, firstParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, secondParkingLot))
}

func TestOwnerAssignMultipleParkingLotToMultipleAttendent(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstAttendent := Implementations.AttendentConstructDefault()
	secondAttendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(3)
	secondParkingLot := owner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToAttendent(firstAttendent, firstParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(secondAttendent, secondParkingLot))
}

func TestSingleAttendentCannotHaveMultipleOwners(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	otherOwner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(3)
	secondParkingLot := otherOwner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, firstParkingLot))
	err := otherOwner.AssignParkingLotToAttendent(attendent, secondParkingLot)
	assert.EqualError(t, err, "this parking lot is not owned by the owner")
}

func TestOwnerAssignParkingLotToSelf(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(2)
	secondParkingLot := owner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToSelf(secondParkingLot))
	assert.NoError(t, owner.AssignParkingLotToSelf(firstParkingLot))
}

func TestOwnerExceptionWhenNoParkingLotAssignedToIt(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	_, err := owner.Park(firstCar)
	assert.Error(t, err)
}

func TestOwnerParkingACar(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(2)
	firstCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	assert.NoError(t, owner.AssignParkingLotToSelf(firstParkingLot))
	_, err := owner.Park(firstCar)
	assert.NoError(t, err)
}

func TestParkingOwnerParkingCarInFullParkingLot(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(1)
	firstCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	assert.NoError(t, owner.AssignParkingLotToSelf(firstParkingLot))
	_, err := owner.Park(firstCar)
	assert.NoError(t, err)

	_, err = owner.Park(secondCar)
	assert.EqualError(t, err, "all parking lots are full")
}

func TestOwnerUnParkTheCar(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(2)
	firstCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	assert.NoError(t, owner.AssignParkingLotToSelf(firstParkingLot))

	ticket, err := owner.Park(firstCar)
	assert.NoError(t, err)

	actualCar, err := owner.Unpark(ticket)
	assert.NoError(t, err)
	assert.Equal(t, firstCar, actualCar)
}

func TestOwnerUnParkingCarParkedInNonAssignedParkingLot(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(2)
	secondParkingLot := owner.CreateParkingLot(3)
	firstCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	attendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	assert.NoError(t, attendent.Assign(secondParkingLot, owner))
	assert.NoError(t, owner.AssignParkingLotToSelf(firstParkingLot))

	ticket, err := attendent.Park(firstCar)
	assert.NoError(t, err)

	_, err = owner.Unpark(ticket)
	assert.EqualError(t, err, "car not found")
}

func TestOwnerNotifiedWhenParkingLotFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(1)
	assert.NoError(t, owner.AssignParkingLotToSelf(parkingLot))

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, err := owner.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.Equal(t, Exceptions.ErrParkingLotIsFull, err) // Check for the specific error
}

func TestOwnerNotifyFullAllParkingLotsAreFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(1)
	ownerSpy := &Implementations.Owner{}
	*ownerSpy = *owner

	firstLot.RegisterNotifiable(ownerSpy)
	secondLot.RegisterNotifiable(ownerSpy)

	_, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.EqualError(t, err, "parking lot is full")
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-9998", Color: Enums.YELLOW})
	assert.EqualError(t, err, "parking lot is full")
}

func TestOwnerNotifyFullSomeParkingLotsAreFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(2)
	ownerSpy := &Implementations.Owner{}
	*ownerSpy = *owner

	firstLot.RegisterNotifiable(ownerSpy)
	secondLot.RegisterNotifiable(ownerSpy)

	_, err := firstLot.Park(&Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED})
	assert.NoError(t, err)
	_, err = secondLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)

	_, err = firstLot.Park(&Implementations.Car{LicensePlate: "AP-9999", Color: Enums.GREEN})
	assert.EqualError(t, err, "parking lot is full")
}

func TestNotifyFullWhenParkingLotIsFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	ownerSpy := &Implementations.Owner{}
	*ownerSpy = *owner
	parkingLot := owner.CreateParkingLot(1)

	firstCar := &Implementations.Car{LicensePlate: "UP81", Color: Enums.BLUE}

	assert.NoError(t, owner.AssignParkingLotToSelf(parkingLot))

	_, err := owner.Park(firstCar)
	assert.NoError(t, err)
}

func TestOwnerNotifiedWhenParkingLotIsAvailable(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(1)
	ownerSpy := &Implementations.Owner{}
	*ownerSpy = *owner

	parkingLot.RegisterNotifiable(ownerSpy)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	ticket, err := parkingLot.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Unpark(ticket)
	assert.NoError(t, err)
	_, err = parkingLot.Park(&Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE})
	assert.NoError(t, err)
}
