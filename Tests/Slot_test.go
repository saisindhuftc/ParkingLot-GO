package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Implementations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlotIsInitiallyFree(t *testing.T) {
	slot := Implementations.SlotConstruct()
	assert.True(t, slot.IsFree())
}

func TestParkCarInFreeSlot(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED)

	ticket, err := slot.Park(car)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.False(t, slot.IsFree())
	assert.True(t, slot.CheckingCarInParkingSlot(car))
}

func TestCannotParkCarInOccupiedSlot(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	_, err := slot.Park(car)
	assert.NoError(t, err)

	anotherCar := Implementations.NewCar("AP-5678", Enums.BLUE) // Update to Enums.BLUE
	_, err = slot.Park(anotherCar)
	assert.Error(t, err)
	assert.Equal(t, "slot is already occupied", err.Error())
}

func TestUnparkCarFromOccupiedSlot(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	ticket, err := slot.Park(car)
	assert.NoError(t, err)

	unparkedCar, err := slot.Unpark(ticket)
	assert.NoError(t, err)
	assert.Equal(t, car, *unparkedCar)
	assert.True(t, slot.IsFree())
}

func TestCannotUnparkCarFromFreeSlot(t *testing.T) {
	slot := Implementations.SlotConstruct()
	invalidTicket := Implementations.TicketConstruct()

	_, err := slot.Unpark(invalidTicket)
	assert.Error(t, err)
	assert.Equal(t, "car not found in the slot", err.Error())
}

func TestHasCarOfSameColor(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	_, err := slot.Park(car)
	assert.NoError(t, err)

	assert.True(t, slot.HasCarOfColor(Enums.RED)) // Update to Enums.RED
}

func TestHasCarOfDifferentColor(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	_, err := slot.Park(car)
	assert.NoError(t, err)

	assert.False(t, slot.HasCarOfColor(Enums.BLUE)) // Update to Enums.BLUE
}

func TestHasCarWithRegistrationNumber(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	_, err := slot.Park(car)
	assert.NoError(t, err)

	assert.True(t, slot.HasCarWithRegistrationNumber("AP-1234"))
}

func TestHasCarWithRegistrationNumberThrowsException(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1432", Enums.YELLOW) // Update to Enums.YELLOW

	_, err := slot.Park(car)
	assert.NoError(t, err)

	assert.False(t, slot.HasCarWithRegistrationNumber("AP-5678"))
}

func TestCheckingCarInParkingSlot(t *testing.T) {
	slot := Implementations.SlotConstruct()
	car := Implementations.NewCar("AP-1234", Enums.RED) // Update to Enums.RED

	_, err := slot.Park(car)
	assert.NoError(t, err)

	assert.True(t, slot.CheckingCarInParkingSlot(car))
	assert.False(t, slot.CheckingCarInParkingSlot(Implementations.NewCar("AP-5678", Enums.BLUE))) // Update to Enums.BLUE
}
