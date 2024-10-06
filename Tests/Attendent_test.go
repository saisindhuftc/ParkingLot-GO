package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Exceptions"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests for Assign() in Attendent
func TestAssignParkingLotToAttendent(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(2)
	attendent := Implementations.AttendentConstructDefault()

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, parkingLot))
}

func TestAssignAParkingLotTwice(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(5)
	attendent := Implementations.AttendentConstructDefault()

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, parkingLot))
	assert.Error(t, owner.AssignParkingLotToAttendent(attendent, parkingLot), Exceptions.ErrParkingLotAlreadyAssigned)
}

func TestAssignMultipleParkingLotToSingleAttendent(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstParkingLot := owner.CreateParkingLot(3)
	secondParkingLot := owner.CreateParkingLot(3)
	thirdParkingLot := owner.CreateParkingLot(3)
	fourthParkingLot := owner.CreateParkingLot(3)
	attendent := Implementations.AttendentConstructDefault()

	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, firstParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, secondParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, thirdParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(attendent, fourthParkingLot))
}

func TestAssignMultipleParkingLotToMultipleAttendents(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstAttendent := Implementations.AttendentConstructDefault()
	secondAttendent := Implementations.AttendentConstructDefault()
	thirdAttendent := Implementations.AttendentConstructDefault()
	fourthAttendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(3)
	secondParkingLot := owner.CreateParkingLot(3)
	thirdParkingLot := owner.CreateParkingLot(3)
	fourthParkingLot := owner.CreateParkingLot(3)

	assert.NoError(t, owner.AssignParkingLotToAttendent(firstAttendent, firstParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(secondAttendent, secondParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(thirdAttendent, thirdParkingLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(fourthAttendent, fourthParkingLot))
}

// Tests for Park() in Attendent
func TestParkIfCarIsAlreadyParked(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	parkingLot := owner.CreateParkingLot(5)
	owner.AssignParkingLotToAttendent(attendent, parkingLot)
	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}

	_, err := attendent.Park(car)
	assert.NoError(t, err)
	_, err = attendent.Park(car)
	assert.Error(t, err, Exceptions.ErrCarAlreadyParked)
}

func TestParkIfCarIsAlreadyParkedInAnotherParkingLot(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(1)
	secondParkingLot := owner.CreateParkingLot(5)
	owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	owner.AssignParkingLotToAttendent(attendent, secondParkingLot)
	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	_, _ = attendent.Park(firstCar)
	_, _ = attendent.Park(secondCar)
	_, err := attendent.Park(firstCar)
	assert.Error(t, err, Exceptions.ErrCarAlreadyParked)
}

func TestAttendentParksCarAndReturnsTicket(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(2)
	attendent := Implementations.AttendentConstructDefault()
	owner.AssignParkingLotToAttendent(attendent, parkingLot)
	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(*car))
}

func TestAttendentCannotParkSameCarTwice(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	parkingLot := owner.CreateParkingLot(2)
	attendent := Implementations.AttendentConstructDefault()
	owner.AssignParkingLotToAttendent(attendent, parkingLot)
	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}

	attendent.Park(car)
	_, err := attendent.Park(car)
	assert.Error(t, err)
	assert.Equal(t, "car already assigned to this parking lot", err.Error())
}

func TestAttendentParksCarsSequentially(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(2)
	secondLot := owner.CreateParkingLot(2)
	attendent := Implementations.AttendentConstructDefault()
	owner.AssignParkingLotToAttendent(attendent, firstLot)
	owner.AssignParkingLotToAttendent(attendent, secondLot)
	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	thirdCar := &Implementations.Car{LicensePlate: "AP-9101", Color: Enums.GREEN}

	_, err := attendent.Park(firstCar)
	assert.NoError(t, err)
	_, err = attendent.Park(secondCar)
	assert.NoError(t, err)
	_, err = attendent.Park(thirdCar)
	assert.NoError(t, err)

	assert.Equal(t, 2, firstLot.CountParkedCars())
	assert.Equal(t, 1, secondLot.CountParkedCars())
}

func TestParkIfFirstParkingLotIsFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(1)
	secondParkingLot := owner.CreateParkingLot(1)
	owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	owner.AssignParkingLotToAttendent(attendent, secondParkingLot)
	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	thirdCar := &Implementations.Car{LicensePlate: "AP-9101", Color: Enums.GREEN}

	_, err := attendent.Park(firstCar)
	assert.NoError(t, err)
	_, err = attendent.Park(secondCar)
	assert.NoError(t, err)
	_, err = attendent.Park(thirdCar)
	assert.Error(t, err, Exceptions.ErrParkingLotIsFull)
}

func TestParkIfAllParkingLotsAreFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	firstParkingLot := owner.CreateParkingLot(1)
	secondParkingLot := owner.CreateParkingLot(1)
	owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	owner.AssignParkingLotToAttendent(attendent, secondParkingLot)
	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	thirdCar := &Implementations.Car{LicensePlate: "AP-9101", Color: Enums.GREEN}

	_, err := attendent.Park(firstCar)
	assert.NoError(t, err)
	_, err = attendent.Park(secondCar)
	assert.NoError(t, err)
	_, err = attendent.Park(thirdCar)
	assert.Error(t, err, Exceptions.ErrParkingLotIsFull)
}

func TestUnparkIfTicketIsInvalidForSecondCar(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()

	parkingLot := owner.CreateParkingLot(5)
	owner.AssignParkingLotToAttendent(attendent, parkingLot)

	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	attendent.Park(firstCar)
	ticket, _ := attendent.Park(secondCar)

	// Unpark second car
	attendent.Unpark(ticket)

	// Trying to unpark the same ticket again should return an error
	_, err := attendent.Unpark(ticket)
	assert.Error(t, err, "Expected CarNotFoundException")
}

func TestUnparkIfTicketIsInvalidForSecondParkingLot(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()

	firstParkingLot := owner.CreateParkingLot(1)
	secondParkingLot := owner.CreateParkingLot(5)

	owner.AssignParkingLotToAttendent(attendent, firstParkingLot)
	owner.AssignParkingLotToAttendent(attendent, secondParkingLot)

	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}

	attendent.Park(firstCar)

	ticket, _ := attendent.Park(secondCar)

	attendent.Unpark(ticket)

	_, err := attendent.Unpark(ticket)
	assert.Error(t, err, Exceptions.ErrCarNotFound)
}

func TestUnparkCarWithValidTicket(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	parkingLot := owner.CreateParkingLot(2)
	owner.AssignParkingLotToAttendent(attendent, parkingLot)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	ticket, _ := attendent.Park(car)

	unparkedCar, err := attendent.Unpark(ticket)
	assert.NoError(t, err, "Should not return error when unparking with a valid ticket")
	assert.Equal(t, car, unparkedCar, "Unparked car should match the parked car")
	assert.False(t, parkingLot.IsCarAlreadyParked(*car), "Car should no longer be parked in the lot")
}

func TestUnparkCarWithInvalidTicket(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	parkingLot := owner.CreateParkingLot(2)

	owner.AssignParkingLotToAttendent(attendent, parkingLot)

	invalidTicket := &Implementations.Ticket{}

	_, err := attendent.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found", err.Error())
}

func TestCannotUnparkSameCarTwice(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	attendent := Implementations.AttendentConstructDefault()
	parkingLot := owner.CreateParkingLot(2)

	owner.AssignParkingLotToAttendent(attendent, parkingLot)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	ticket, _ := attendent.Park(car)

	attendent.Unpark(ticket)

	_, err := attendent.Unpark(ticket)
	assert.Error(t, err)
	assert.Equal(t, "car not found", err.Error())
}

func TestSmartAttendentThrowsWhenNoParkingLotAssigned(t *testing.T) {
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})
	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}

	_, err := smartAttendent.Park(car)
	assert.Error(t, err, Exceptions.ErrParkingLotAlreadyAssigned)
}

func TestSmartAttendentAssignsMultipleLots(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(2)
	secondLot := owner.CreateParkingLot(3)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	assert.NoError(t, owner.AssignParkingLotToAttendent(smartAttendent, firstLot))
	assert.NoError(t, owner.AssignParkingLotToAttendent(smartAttendent, secondLot))
}

func TestSmartAttendentParksCarInLotWithFewestCars(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(2)
	secondLot := owner.CreateParkingLot(2)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	owner.AssignParkingLotToAttendent(smartAttendent, firstLot)
	owner.AssignParkingLotToAttendent(smartAttendent, secondLot)

	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, _ = smartAttendent.Park(firstCar)

	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	secondTicket, _ := smartAttendent.Park(secondCar)

	assert.NotNil(t, secondTicket, "Ticket should not be nil for secondCar")
	assert.True(t, secondLot.IsCarAlreadyParked(*secondCar), "secondCar should be parked in secondLot")
}

func TestSmartAttendentThrowsWhenAllLotsFull(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(1)
	secondLot := owner.CreateParkingLot(1)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	owner.AssignParkingLotToAttendent(smartAttendent, firstLot)
	owner.AssignParkingLotToAttendent(smartAttendent, secondLot)

	firstCar := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, _ = smartAttendent.Park(firstCar)

	secondCar := &Implementations.Car{LicensePlate: "AP-5678", Color: Enums.BLUE}
	_, _ = smartAttendent.Park(secondCar)

	thirdCar := &Implementations.Car{LicensePlate: "AP-9012", Color: Enums.GREEN}
	_, err := smartAttendent.Park(thirdCar)
	assert.Error(t, err, Exceptions.ErrParkingLotIsFull)
}

func TestSmartAttendentCannotParkSameCarTwice(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	lot := owner.CreateParkingLot(2)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})
	owner.AssignParkingLotToAttendent(smartAttendent, lot)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	_, _ = smartAttendent.Park(car) // Park the car

	_, err := smartAttendent.Park(car)
	assert.Error(t, err, Exceptions.ErrCarAlreadyParked)
}

func TestSmartAttendentUnparksCarWithValidTicket(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(2)
	secondLot := owner.CreateParkingLot(2)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	owner.AssignParkingLotToAttendent(smartAttendent, firstLot)
	owner.AssignParkingLotToAttendent(smartAttendent, secondLot)

	car := &Implementations.Car{LicensePlate: "AP-1234", Color: Enums.RED}
	ticket, _ := smartAttendent.Park(car)

	unparkedCar, err := smartAttendent.Unpark(ticket)
	assert.NoError(t, err, "Should not return error when unparking with a valid ticket")
	assert.Equal(t, car, unparkedCar, "Unparked car should match the parked car")
	assert.False(t, firstLot.IsCarAlreadyParked(*car), "The car should no longer be parked in the lot")
}

func TestSmartAttendentUnparksCarWithInvalidTicket(t *testing.T) {
	owner := Implementations.OwnerConstruct()
	firstLot := owner.CreateParkingLot(2)
	smartAttendent := Implementations.AttendentConstruct(&Implementations.SmartNextLotStrategy{})

	owner.AssignParkingLotToAttendent(smartAttendent, firstLot)

	invalidTicket := &Implementations.Ticket{}

	_, err := smartAttendent.Unpark(invalidTicket)
	assert.Error(t, err, Exceptions.ErrCarNotFound)
}
