package common

type NabSyncedItems struct {
	Led   []NabLedCmd
	Ear   []NabEarCmd
	Count int
}

type NabSync struct {
	Count int    `json:"count"`
	ID    string `json:"id"`
}

type NabLedItem struct {
	Color    string `json:"color"`
	Duration uint   `json:"duration"`
}

type NabLedCmd struct {
	Delay    int          `json:"delay"`
	ID       NabLed       `json:"id"`
	Sequence []NabLedItem `json:"sequence"`
	Sync     NabSync      `json:"sync"`
}

type NabEarItem struct {
	Position  uint8 `json:"position"`
	Direction uint8 `json:"direction"`
	Duration  uint  `json:"duration"`
}

type NabEarCmd struct {
	Delay    int          `json:"delay"`
	ID       NabEar       `json:"id"`
	Sequence []NabEarItem `json:"sequence"`
	Sync     NabSync      `json:"sync"`
}

type NabEarEvent struct {
	ID       NabEar `json:"id"`
	Position uint8  `json:"position"`
}

type NabStatus struct {
	Connected bool       `json:"connected"`
	IP        string     `json:"ip"`
	HWVersion NabVersion `json:"hwVersion"`
	FWVersion string     `json:"fwVersion"`
	SWVersion string     `json:"swVersion"`
}

type NabAudioCmd struct {
	State  NabAudio `json:"state"`
	Link   string   `json:"link"`
	Volume uint8    `json:"volume"`
}

type NabAudioEvent struct {
	State NabAudio `json:"state"`
}

type NabRFIDEvent struct {
	Value string `json:"value"`
}

type NabButtonEvent struct {
	ID    NabButton `json:"id"`
	Value uint8     `json:"value"`
}
