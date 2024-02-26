package v2

import "github.com/droso-hass/openab/internal/common"

func (n *NabConn) write(data string) error {
	_, err := n.conn.Write([]byte(data))
	return err
}

func (n *NabConn) SyncPacket(mac string, data common.NabSyncedItems) error {
	return nil
}

func (n *NabConn) Led(mac string, data common.NabLedCmd) error {
	return nil
}

func (n *NabConn) Ear(mac string, data common.NabEar) error {
	return nil
}

func (n *NabConn) RecorderState(mac string, data common.NabAudio) error {
	return nil
}

func (n *NabConn) PlayerState(mac string, data common.NabAudio) error {
	return nil
}

func (n *NabConn) PlayLink(mac string, url string) error {
	return nil
}

func (n *NabConn) PlayerVolume(mac string, data uint8) error {
	return nil
}

func (n *NabConn) PlayerData(mac string, data []byte) error {
	return nil
}

func (n *NabConn) Reboot(mac string) error {
	return nil
}
