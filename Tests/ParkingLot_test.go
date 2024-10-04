package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExceptionNewParkingLotIsEmpty(t *testing.T) {
	owner := Implementations.NewOwner()
	assert.PanicsWithError(t, "parking lot size must be positive", func() {
		owner.CreateParkingLot(0)
	})
}

func TestExceptionForNegativeParkingSlots(t *testing.T) {
	owner := Implementations.NewOwner()
	assert.PanicsWithError(t, "parking lot size must be positive", func() {
		owner.CreateParkingLot(-1)
	})
}

func TestCreateParkingLotWith5Slots(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(5)

	assert.NotNil(t, parkingLot)
}

func TestCannotParkSameCarTwice(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(5)
	car := Implementations.NewCar("AP-1234", Enums.Red)

	parkingLot.Park(&car)
	assert.PanicsWithError(t, "car is already parked", func() {
		parkingLot.Park(&car)
	})
}

func TestParkingLotIsFullAfterParkingOneCarInOneSlotLot(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(1)
	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket := parkingLot.Park(&car)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
	assert.True(t, parkingLot.IsFull())
}

func TestParkingLotWithTwoSlotsIsNotFullWithOneCarParked(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(2)
	car := Implementations.NewCar("AP-1431", Enums.Blue)

	ticket := parkingLot.Park(&car)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
	assert.False(t, parkingLot.IsFull())
}

func TestParkingCarInFullParkingLotThrowsError(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(1)
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-5678", Enums.Blue)

	parkingLot.Park(&firstCar)
	assert.PanicsWithError(t, "parking lot is full", func() {
		parkingLot.Park(&secondCar)
	})
}

func TestParkInNearestAvailableSlot(t *testing.T) {
	owner := Implementations.NewOwner()
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-9999", Enums.Blue)
	parkingLot := owner.CreateParkingLot(5)

	parkingLot.Park(&firstCar)
	parkingLot.Park(&secondCar)

	assert.True(t, parkingLot.IsCarAlreadyParked(firstCar))
	assert.True(t, parkingLot.IsCarAlreadyParked(secondCar))
}

func TestParkInNearestAvailableSlotAfterUnparking(t *testing.T) {
	owner := Implementations.NewOwner()
	parkingLot := owner.CreateParkingLot(5)
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-5678", Enums.Blue)
	thirdCar := Implementations.NewCar("AP-9999", Enums.Green)

	firstCarTicket := parkingLot.Park(&firstCar)
	parkingLot.Park(&secondCar)

	_, err := parkingLot.Unpark(firstCarTicket)
	assert.Nil(t, err)

	parkingLot.Park(&thirdCar)
	assert.True(t, parkingLot.IsCarAlreadyParked(thirdCar))
}

func TestUnparkCarThatIsNotParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	invalidTicket := &Implementations.Ticket{}

	_, err := parkingLot.Unpark(invalidTicket)

	assert.NotNil(t, err)
	assert.Equal(t, "car not found in the parking lot", err.Error())
}

func TestUnparkCarFromEmptyParkingLot(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	invalidTicket := &Implementations.Ticket{} // Empty parking lot

	_, err := parkingLot.Unpark(invalidTicket)

	assert.NotNil(t, err)
	assert.Equal(t, "car not found in the parking lot", err.Error())
}

func TestUnpark(t *testing.T) {
	owner := &Implementations.Owner{}
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	parkingLot := Implementations.NewParkingLot(5, owner)
	ticket := parkingLot.Park(car)

	unparkedCar, err := parkingLot.Unpark(ticket)

	assert.Nil(t, err)
	assert.Equal(t, car, unparkedCar)
}

func TestCountCarsByRedColorIsNotFoundInParkingLot(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(1, owner)

	count := parkingLot.CountCarsByColor(Enums.Red)

	assert.Equal(t, 0, count)
}

func TestCountCarsByColorNotPresent(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(1, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Blue}

	parkingLot.Park(car)

	assert.Equal(t, 0, parkingLot.CountCarsByColor(Enums.Yellow))
}

func TestIsCarAlreadyParkedForNonParkedCar(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1432", Color: Enums.Yellow}

	assert.False(t, parkingLot.IsCarAlreadyParked(*car))
}

func TestIsParkingLotFull(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(1, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-4321", Color: Enums.Blue}

	parkingLot.Park(car)

	assert.True(t, parkingLot.IsFull())
}

func TestIsParkingLotNotFull(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-9876", Color: Enums.Green}

	parkingLot.Park(car)

	assert.False(t, parkingLot.IsFull())
}

func TestCountCarsByColor(t *testing.T) {
	owner := &Implementations.Owner{}
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-9999", Color: Enums.Red}
	thirdCar := &Implementations.Car{RegistrationNumber: "AP-0001", Color: Enums.Blue}
	parkingLot := Implementations.NewParkingLot(5, owner)

	parkingLot.Park(firstCar)
	parkingLot.Park(secondCar)
	parkingLot.Park(thirdCar)

	assert.Equal(t, 2, parkingLot.CountCarsByColor(Enums.Red))
}

func TestIsCarWithRegistrationNumberParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}

	parkingLot.Park(car)

	isParked, _ := parkingLot.IsCarWithRegistrationNumberParked("AP-1234")
	assert.True(t, isParked)
}

func TestIsCarWithoutRegistrationNumberCannotParked(t *testing.T) {
	owner := &Implementations.Owner{}
	parkingLot := Implementations.NewParkingLot(5, owner)
	car := &Implementations.Car{RegistrationNumber: "", Color: Enums.Red}

	_, err := parkingLot.IsCarWithRegistrationNumberParked(car.RegistrationNumber)

	assert.NotNil(t, err)
	assert.Equal(t, "car needs registration number", err.Error())
}
