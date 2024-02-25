package common

type INabReceiver interface {
	SyncPacket(mac string, data NabSyncedItems) error
	Led(mac string, data NabLedCmd) error
	Ear(mac string, data NabEar) error
	RecorderState(mac string, data NabAudio) error
	PlayerState(mac string, data NabAudio) error
	PlayLink(mac string, url string) error
	PlayerVolume(mac string, data uint8) error
	PlayerData(mac string, data []byte) error
	Reboot(mac string) error
}

type INabSender interface {
	Ear(mac string, data NabEarEvent)
	Status(mac string, data NabStatus)
	Rfid(mac string, data string)
	PlayerState(mac string) // may be changed without user input (ex: when sending data/on data end)
	Button(mac string, data NabButton, value uint8)
	RecorderData(mac string, data []byte)
}
