package api

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/droso-hass/openab/internal/common"
	"github.com/droso-hass/openab/internal/utils"
	"github.com/nats-io/nats.go"
)

func (a *API) getReceiver(mac string) common.INabReceiver {
	for _, h := range a.receivers {
		recv := h.FindReceiver(mac)
		if recv != nil {
			return recv
		}
	}
	return nil
}

func (a *API) processSub(m *nats.Msg) {
	sp := strings.Split(m.Subject, ".")
	recv := a.getReceiver(sp[1])
	if recv == nil {
		return
	}
	var err error

	switch sp[2] {
	case "led":
		c := common.NabLedCmd{}
		err = json.Unmarshal(m.Data, &c)
		if err == nil {
			if c.Sync.ID != "" {
				sync, ok := a.syncCmds[c.Sync.ID]
				if !ok {
					sync = common.NabSyncedItems{
						Count: 1,
					}
				} else {
					sync.Count++
				}
				sync.Led = append(sync.Led, c)

				if sync.Count >= c.Sync.Count {
					err = recv.SyncPacket(sync)
					delete(a.syncCmds, c.Sync.ID)
				} else {
					a.syncCmds[c.Sync.ID] = sync
				}
			} else {
				err = recv.Led(c)
			}
		}
	case "ear":
		c := common.NabEarCmd{}
		err = json.Unmarshal(m.Data, &c)
		if err == nil {
			if c.Sync.ID != "" {
				sync, ok := a.syncCmds[c.Sync.ID]
				if !ok {
					sync = common.NabSyncedItems{
						Count: 1,
					}
				} else {
					sync.Count++
				}
				sync.Ear = append(sync.Ear, c)

				if sync.Count >= c.Sync.Count {
					err = recv.SyncPacket(sync)
					delete(a.syncCmds, c.Sync.ID)
				} else {
					a.syncCmds[c.Sync.ID] = sync
				}
			} else {
				err = recv.Ear(c)
			}
		}
	case "player":
		if sp[3] == "command" {
			c := common.NabAudioCmd{}
			err = json.Unmarshal(m.Data, &c)
			if err == nil {
				err = recv.Player(c)
			}
		} else if sp[3] == "data" {
			err = recv.PlayerData(m.Data)
		}
	case "recorder":
		if sp[3] == "command" {
			c := common.NabAudioCmd{}
			err = json.Unmarshal(m.Data, &c)
			if err == nil {
				err = recv.Recorder(c)
			}
		}
	case "reboot":
		err = recv.Reboot()
	}

	if err != nil {
		slog.Warn("error processing command", utils.ErrAttr(err))
	}
}
