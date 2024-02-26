package common

type INabReceiver interface {
	SyncPacket(data NabSyncedItems) error
	Led(data NabLedCmd) error
	Ear(data NabEarCmd) error
	Player(data NabAudioCmd) error
	Recorder(data NabAudioCmd) error
	PlayerData(data []byte) error
	Reboot() error
}

type INabReceiverHander interface {
	FindReceiver(mac string) INabReceiver
}

type INabSender interface {
	Ear(mac string, data NabEarEvent)
	Status(mac string, data NabStatus)
	Rfid(mac string, data NabRFIDEvent)
	PlayerState(mac string, data NabAudioEvent) // may be changed without user input (ex: when sending data/on data end)
	Button(mac string, data NabButtonEvent)
	RecorderData(mac string, data []byte)
}
