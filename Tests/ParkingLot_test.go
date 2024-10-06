package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Exceptions"
	"ParkingLot_go/Implementations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExceptionNewParkingLotIsEmpty(t *testing.T) {
	owner := &Implementations.Owner{}

	assert.PanicsWithValue(t, Exceptions.ErrCannotCreateParkingLotException, func() {
		Implementations.ParkingLotConstruct(0, owner)
	})
}

func TestExceptionForNegativeParkingSlots(t *testing.T) {
	owner := &Implementations.Owner{}

	assert.PanicsWithValue(t, Exceptions.ErrCannotCreateParkingLotException, func() {
		Implementations.ParkingLotConstruct(-1, owner)
	})
}

func TestCreateParkingLotWith5Slots(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)

	assert.NotNil(t, parkingLot)
}

func TestParkWith5Slots(t *testing.T) {
	owner := &Implementations.Owner{}
	car := &Implementations.Car{RegistrationNumber: "AP-9876", Color: Enums.BLACK}

	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	ticket, _ := parkingLot.Park(car)

	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(*car))
}

func TestCannotParkSameCarTwice(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Park(car)
	assert.EqualError(t, err, Exceptions.ErrCarAlreadyParked.Error())
}

func TestParkingLotWithOneSlotIsFullWhenCarParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(1, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}

	ticket, _ := parkingLot.Park(car)

	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(*car))
}

func TestParkingLotWithTwoSlotsIsNotFullWhenOneCarParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(2, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1431", Color: Enums.BLUE}

	ticket, _ := parkingLot.Park(car)

	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(*car))
	assert.False(t, parkingLot.IsFull())
}

func TestParkInFullParkingLot(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(1, owner)
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-5678", Color: Enums.BLUE}

	_, err := parkingLot.Park(firstCar)
	assert.NoError(t, err)

	_, err = parkingLot.Park(secondCar)
	assert.EqualError(t, err, Exceptions.ErrParkingLotIsFull.Error())
}

func TestParkInNearestAvailableSlot(t *testing.T) {
	owner := &Implementations.Owner{}
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-9999", Color: Enums.BLUE}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)

	parkingLot.Park(firstCar)
	parkingLot.Park(secondCar)

	assert.True(t, parkingLot.IsCarAlreadyParked(*firstCar))
	assert.True(t, parkingLot.IsCarAlreadyParked(*secondCar))
}

func TestParkInNearestAvailableSlotAfterUnparking(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-5678", Color: Enums.BLUE}
	thirdCar := &Implementations.Car{RegistrationNumber: "AP-9999", Color: Enums.GREEN}
	firstCarTicket, _ := parkingLot.Park(firstCar)

	parkingLot.Park(secondCar)
	parkingLot.Unpark(firstCarTicket)
	parkingLot.Park(thirdCar)

	assert.True(t, parkingLot.IsCarAlreadyParked(*thirdCar))
}

func TestIsCarAlreadyParkedForNonParkedCar(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1432", Color: Enums.YELLOW}

	assert.False(t, parkingLot.IsCarAlreadyParked(*car))
}

func TestIsParkingLotFull(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(1, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-4321", Color: Enums.BLUE}

	parkingLot.Park(car)

	assert.True(t, parkingLot.IsFull())
}

func TestIsParkingLotNotFull(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-9876", Color: Enums.GREEN}

	parkingLot.Park(car)

	assert.False(t, parkingLot.IsFull())
}

func TestUnpark(t *testing.T) {
	owner := &Implementations.Owner{}
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)

	ticket, _ := parkingLot.Park(car)
	unparkedCar, err := parkingLot.Unpark(ticket)

	assert.Nil(t, err)
	assert.Equal(t, car, unparkedCar)
}

func TestUnparkCarThatIsNotParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	invalidTicket := &Implementations.Ticket{}
	_, err := parkingLot.Unpark(invalidTicket)

	assert.NotNil(t, err)
	assert.Equal(t, Exceptions.ErrInvalidTicket, err)
}

func TestUnparkCarFromEmptyParkingLot(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	invalidTicket := &Implementations.Ticket{}

	_, err := parkingLot.Unpark(invalidTicket)

	assert.NotNil(t, err)
	assert.Equal(t, Exceptions.ErrInvalidTicket, err)
}

func TestCountCarsByColor(t *testing.T) {
	owner := &Implementations.Owner{}
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-9999", Color: Enums.RED}
	thirdCar := &Implementations.Car{RegistrationNumber: "AP-0001", Color: Enums.BLUE}

	parkingLot := Implementations.ParkingLotConstruct(5, owner)

	parkingLot.Park(firstCar)
	parkingLot.Park(secondCar)
	parkingLot.Park(thirdCar)

	assert.Equal(t, 2, parkingLot.CountCarsByColor(Enums.RED))
}

func TestCountCarsByRedColorIsNotFoundInParkingLot(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(1, owner)
	count := parkingLot.CountCarsByColor(Enums.RED)

	assert.Equal(t, 0, count)
}

func TestCountCarsByColorNotPresent(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(1, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.BLUE}

	parkingLot.Park(car)

	assert.Equal(t, 0, parkingLot.CountCarsByColor(Enums.YELLOW))
}

func TestIsCarWithRegistrationNumberParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.RED}

	parkingLot.Park(car)
	isParked, _ := parkingLot.IsCarWithRegistrationNumberParked("AP-1234")

	assert.True(t, isParked)
}

func TestIsCarWithoutRegistrationNumberCannotParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.ParkingLotConstruct(5, owner)
	car := &Implementations.Car{RegistrationNumber: "", Color: Enums.RED}
	_, err := parkingLot.IsCarWithRegistrationNumberParked(car.RegistrationNumber)

	assert.NotNil(t, err)
	assert.Equal(t, Exceptions.ErrCarNeedsRegistrationNumber, err)
}
