package common

type NabSyncedItems struct {
	Led []NabLedCmd
	Ear []NabEarCmd
}

type NabSync struct {
	Count int `json:"count"`
	ID    int `json:"id"`
}

type NabLedItem struct {
	Color    string `json:"color"`
	Duration int    `json:"duration"`
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
}

type NabEarCmd struct {
	Delay    int        `json:"delay"`
	ID       NabEar     `json:"id"`
	Sequence NabEarItem `json:"sequence"`
	Sync     NabSync    `json:"sync"`
}

type NabEarEvent struct {
	ID       NabButton `json:"id"`
	Position uint8     `json:"position"`
}

type NabStatus struct {
	Connected bool       `json:"connected"`
	IP        string     `json:"ip"`
	HWVersion NabVersion `json:"hwVersion"`
	FWVersion string     `json:"fwVersion"`
	SWVersion string     `json:"swVersion"`
}
