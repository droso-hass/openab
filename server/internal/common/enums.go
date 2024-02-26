package common

type NabAudio uint8

const (
	NabAudioStopped = 1
	NabAudioRunning = 2
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
