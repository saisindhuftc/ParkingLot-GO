package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExceptionNewParkingLotIsEmpty(t *testing.T) {
	_, err := Implementations.NewParkingLot(0)
	assert.Error(t, err)
	assert.Equal(t, "parking lot size must be positive", err.Error())
}

func TestExceptionForNegativeParkingSlots(t *testing.T) {
	_, err := Implementations.NewParkingLot(-1)
	assert.Error(t, err)
	assert.Equal(t, "parking lot size must be positive", err.Error())
}

func TestCreateParkingLotWith5Slots(t *testing.T) {
	parkingLot, err := Implementations.NewParkingLot(5)
	assert.NoError(t, err)
	assert.NotNil(t, parkingLot)
}

func TestCannotParkSameCarTwice(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	car := Implementations.NewCar("AP-1234", Enums.Red)

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)

	_, err = parkingLot.Park(car)
	assert.Error(t, err)
	assert.Equal(t, "car is already parked", err.Error())
}

func TestParkingLotWithOneSlotIsFullWhenCarParked(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	car := Implementations.NewCar("AP-1234", Enums.Red)
	ticket, err := parkingLot.Park(car)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
}

func TestParkingLotWithTwoSlotsIsNotFullWhenOneCarParked(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(2)
	car := Implementations.NewCar("AP-1431", Enums.Blue)
	ticket, err := parkingLot.Park(car)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
}

func TestParkWith5Slots(t *testing.T) {
	car := Implementations.NewCar("AP-9876", Enums.Black)
	parkingLot, _ := Implementations.NewParkingLot(5)
	ticket, err := parkingLot.Park(car)

	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
}

func TestParkInFullParkingLot(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-5678", Enums.Blue)

	_, err := parkingLot.Park(firstCar)
	assert.NoError(t, err)

	_, err = parkingLot.Park(secondCar)
	assert.Error(t, err)
	assert.Equal(t, "parking lot is full", err.Error())
}

func TestParkInNearestAvailableSlot(t *testing.T) {
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-9999", Enums.Blue)
	parkingLot, _ := Implementations.NewParkingLot(5)

	_, err := parkingLot.Park(firstCar)
	assert.NoError(t, err)

	_, err = parkingLot.Park(secondCar)
	assert.NoError(t, err)

	assert.True(t, parkingLot.IsCarAlreadyParked(firstCar))
	assert.True(t, parkingLot.IsCarAlreadyParked(secondCar))
}

func TestParkInNearestAvailableSlotAfterUnparking(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-5678", Enums.Blue)
	thirdCar := Implementations.NewCar("AP-9999", Enums.Green)
	firstCarTicket, _ := parkingLot.Park(firstCar)

	_, err := parkingLot.Park(secondCar)
	assert.NoError(t, err)

	_, err = parkingLot.Unpark(firstCarTicket)
	assert.NoError(t, err)

	_, err = parkingLot.Park(thirdCar)
	assert.NoError(t, err)

	assert.True(t, parkingLot.IsCarAlreadyParked(thirdCar))
}

func TestUnparkCarThatIsNotParked(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	invalidTicket := Implementations.NewTicket() // Ticket for an empty slot

	_, err := parkingLot.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in the parking lot", err.Error())
}

func TestUnparkCarFromEmptyParkingLot(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	invalidTicket := Implementations.NewTicket() // Empty parking lot

	_, err := parkingLot.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in the parking lot", err.Error())
}

func TestUnpark(t *testing.T) {
	car := Implementations.NewCar("AP-1234", Enums.Red)
	parkingLot, _ := Implementations.NewParkingLot(5)
	ticket, _ := parkingLot.Park(car)
	unparkedCar, err := parkingLot.Unpark(ticket)

	assert.NoError(t, err)
	assert.Equal(t, car, *unparkedCar) // Ensure the correct car is returned
}

func TestCountCarsByRedColorIsNotFoundInParkingLot(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	count := parkingLot.CountCarsByColor(Enums.Red)

	assert.Equal(t, 0, count)
}

func TestCountCarsByColorNotPresent(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	car := Implementations.NewCar("AP-1234", Enums.Blue)

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)
	assert.Equal(t, 0, parkingLot.CountCarsByColor(Enums.Yellow))
}

func TestCountCarsByColor(t *testing.T) {
	firstCar := Implementations.NewCar("AP-1234", Enums.Red)
	secondCar := Implementations.NewCar("AP-9999", Enums.Red)
	thirdCar := Implementations.NewCar("AP-0001", Enums.Blue)
	parkingLot, _ := Implementations.NewParkingLot(5)

	_, err := parkingLot.Park(firstCar)
	assert.NoError(t, err)

	_, err = parkingLot.Park(secondCar)
	assert.NoError(t, err)

	_, err = parkingLot.Park(thirdCar)
	assert.NoError(t, err)

	assert.Equal(t, 2, parkingLot.CountCarsByColor(Enums.Red))
}

func TestIsCarAlreadyParkedForNonParkedCar(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	car := Implementations.NewCar("AP-1432", Enums.Yellow)

	assert.False(t, parkingLot.IsCarAlreadyParked(car)) // Car is not parked
}

func TestIsParkingLotFull(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	car := Implementations.NewCar("AP-4321", Enums.Blue)

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)
	assert.True(t, parkingLot.IsFull()) // Parking lot is full after one car is parked
}

func TestIsParkingLotNotFull(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	car := Implementations.NewCar("AP-9876", Enums.Green)

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)
	assert.False(t, parkingLot.IsFull()) // Parking lot is not full
}

func TestIsCarWithRegistrationNumberParked(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(5)
	car := Implementations.NewCar("AP-1234", Enums.Red)

	_, err := parkingLot.Park(car)
	assert.NoError(t, err)

	isParked, err := parkingLot.IsCarWithRegistrationNumberParked("AP-1234")
	assert.NoError(t, err)
	assert.True(t, isParked)
}
