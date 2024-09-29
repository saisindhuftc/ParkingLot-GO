package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Implementations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignParkingLotToAttendent(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(1)

	err := attendent.Assign(*parkingLot)
	assert.NoError(t, err)
}

func TestCannotAssignSameParkingLotTwice(t *testing.T) {
	parkingLot, _ := Implementations.NewParkingLot(1)
	attendent := Implementations.NewAttendent()

	err := attendent.Assign(*parkingLot)
	assert.NoError(t, err)

	err = attendent.Assign(*parkingLot)
	assert.Error(t, err)
	assert.Equal(t, "parking lot already assigned", err.Error())
}

func TestAttendantParksCarAndReturnsTicket(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(2)
	attendent.Assign(*parkingLot)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.True(t, parkingLot.IsCarAlreadyParked(car))
}

func TestAttendantUnparksCarWithTicket(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(2)
	attendent.Assign(*parkingLot)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)

	unparkedCar, err := attendent.Unpark(ticket)
	assert.NoError(t, err)
	assert.Equal(t, car, *unparkedCar)
	assert.False(t, parkingLot.IsCarAlreadyParked(car))
}

func TestAttendantUnparksCarWithInvalidTicket(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(2)
	attendent.Assign(*parkingLot)

	invalidTicket := Implementations.NewTicket()

	_, err := attendent.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in assigned parking lot", err.Error())
}

func TestUnparkedCarShouldMatchWithParkedCar(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot1, _ := Implementations.NewParkingLot(1)
	parkingLot2, _ := Implementations.NewParkingLot(1)
	attendent.Assign(*parkingLot1)
	attendent.Assign(*parkingLot2)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)

	unparkedCar, err := attendent.Unpark(ticket)
	assert.NoError(t, err)
	assert.Equal(t, car, *unparkedCar)
}

func TestUnparkCarThatIsNotInAssignedParkingLot(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot1, _ := Implementations.NewParkingLot(1)
	parkingLot2, _ := Implementations.NewParkingLot(1)
	attendent.Assign(*parkingLot1)
	attendent.Assign(*parkingLot2)

	invalidTicket := Implementations.NewTicket()

	_, err := attendent.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in assigned parking lot", err.Error())
}

func TestAttendantCannotParkSameCarTwice(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(2)
	attendent.Assign(*parkingLot)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	_, err := attendent.Park(car)
	assert.NoError(t, err)

	_, err = attendent.Park(car)
	assert.Error(t, err)
	assert.Equal(t, "car already assigned to this parking lot", err.Error())
}

func TestParkingSameInDifferentSlots(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot1, _ := Implementations.NewParkingLot(1)
	parkingLot2, _ := Implementations.NewParkingLot(1)
	attendent.Assign(*parkingLot1)
	attendent.Assign(*parkingLot2)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	_, err := attendent.Park(car)
	assert.NoError(t, err)

	_, err = attendent.Park(car)
	assert.Error(t, err)
	assert.Equal(t, "car already assigned to this parking lot", err.Error())
}

func TestUnParkCarInDifferentSlot(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot1, _ := Implementations.NewParkingLot(1)
	parkingLot2, _ := Implementations.NewParkingLot(1)
	attendent.Assign(*parkingLot1)
	attendent.Assign(*parkingLot2)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)

	unparkedCar, err := attendent.Unpark(ticket)
	assert.NoError(t, err)
	assert.Equal(t, car, *unparkedCar)
	assert.False(t, parkingLot1.IsCarAlreadyParked(car))
	assert.False(t, parkingLot2.IsCarAlreadyParked(car))
}

func TestUnParkTheSameCarAgain(t *testing.T) {
	attendent := Implementations.NewAttendent()
	parkingLot, _ := Implementations.NewParkingLot(1)
	attendent.Assign(*parkingLot)

	car := Implementations.NewCar("AP-1234", Enums.Red)

	ticket, err := attendent.Park(car)
	assert.NoError(t, err)

	_, err = attendent.Unpark(ticket)
	assert.NoError(t, err)

	_, err = attendent.Unpark(ticket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in assigned parking lot", err.Error())
}
