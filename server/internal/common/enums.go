package common

type MsgType uint8

const (
	MsgTypePlay         MsgType = 0
	MsgTypePlayStart    MsgType = 1
	MsgTypePlayStop     MsgType = 2
	MsgTypeVolume       MsgType = 3
	MsgTypeRecordStart  MsgType = 4
	MsgTypeRecordStop   MsgType = 5
	MsgTypeSetEar       MsgType = 6
	MsgTypeEarData      MsgType = 7
	MsgTypeSetLed       MsgType = 8
	MsgTypeGetImage     MsgType = 9
	MsgTypeImageData    MsgType = 10
	MsgTypeReboot       MsgType = 11
	MsgTypeConnected    MsgType = 12 // mac, version
	MsgTypeDisconnected MsgType = 13
	MsgTypeButton       MsgType = 14 // type (short, double, long) + (value for wheel)
	MsgTypeRfid         MsgType = 15
	MsgTypeRecordData   MsgType = 16
)

type NabAudio uint8

const (
	NabAudioStopped = 0
	NabAudioRunning = 1
)

type NabButton uint8

const (
	NabButtonShort     NabButton = 0
	NabButtonDouble    NabButton = 1
	NabButtonLongStart NabButton = 2
	NabButtonLongEnd   NabButton = 3
	NabButtonVolume    NabButton = 4
)

type NabLed uint8

const (
	NabLedNose   NabLed = 0
	NabLedLeft   NabLed = 1
	NabLedMiddle NabLed = 2
	NabLedRight  NabLed = 3
	NabLedBottom NabLed = 4
)

type NabEar uint8

const (
	NabEarLeft  NabEar = 0
	NabEarRight NabEar = 1
)

type NabVersion uint8

const (
	NabVersionOriginal NabVersion = 1
	NabVersionTagTag   NabVersion = 2
	NabVersionKarotz   NabVersion = 3
)
