package Implementations

type Notifiable interface {
	notifyFull(parkingLotId int)
	notifyAvailable(parkingLotId int)
}
