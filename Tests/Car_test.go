package Tests

import (
	"ParkingLot_go/Enums"
	"ParkingLot_go/Implementations"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarCreation(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}

	assert.NotNil(t, car, "Car object should not be nil")
	assert.True(t, car.HasRegistrationNumber("AP-1234"), "Registration number should match")
	assert.True(t, car.IsColor(Enums.Red), "Color should match")
}

func TestCarEquality(t *testing.T) {
	firstCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	secondCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	thirdCar := &Implementations.Car{RegistrationNumber: "AP-5678", Color: Enums.Blue}

	assert.Equal(t, firstCar, secondCar)
	assert.NotEqual(t, firstCar, thirdCar)
}

func TestCarInequalityWithDifferentRegistrationNumber(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	anotherCar := &Implementations.Car{RegistrationNumber: "AP-5678", Color: Enums.Red}

	assert.NotEqual(t, car, anotherCar)
}

func TestCarInequalityWithDifferentColor(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	anotherCar := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Blue}

	assert.NotEqual(t, car, anotherCar)
}

func TestCarInequalityWithDifferentObject(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}
	notACar := "Not a car"

	assert.NotEqual(t, car, notACar)
}

func TestCarInequalityWithNull(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1234", Color: Enums.Red}

	assert.NotEqual(t, car, nil)
}

func TestCarWithColorYellow(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1432", Color: Enums.Yellow}

	assert.True(t, car.IsColor(Enums.Yellow))
}

func TestCarWithIncorrectColorYellow(t *testing.T) {
	car := &Implementations.Car{RegistrationNumber: "AP-1432", Color: Enums.Yellow}

	assert.False(t, car.IsColor(Enums.Blue))
}
