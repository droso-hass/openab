package v2

import (
	"errors"
	"fmt"
	"strings"

	"github.com/droso-hass/openab/internal/common"
)

var ErrUnknownCmd = errors.New("unknown command")

func (n *NabConn) write(data string) error {
	_, err := n.conn.Write([]byte(data))
	return err
}

func encodeLed(data common.NabLedCmd) string {
	str := strings.Builder{}
	for _, x := range data.Sequence {
		str.WriteString(fmt.Sprintf("%s;%d;", x.Color, x.Duration))
	}
	return fmt.Sprintf("03;%d;%d;%s", data.ID, data.Delay, strings.TrimSuffix(str.String(), ";"))
}

func encodeEar(data common.NabEarCmd) string {
	str := strings.Builder{}
	for _, x := range data.Sequence {
		str.WriteString(fmt.Sprintf("%d;%d;%d", x.Position, x.Direction, x.Duration))
	}
	return fmt.Sprintf("02;%d;%d;%s", data.ID, data.Delay, strings.TrimSuffix(str.String(), ";"))
}

func (n *NabConn) SyncPacket(data common.NabSyncedItems) error {
	str := strings.Builder{}
	for _, x := range data.Led {
		str.WriteString(encodeLed(x))
		str.WriteString("\n")
	}
	for _, x := range data.Ear {
		str.WriteString(encodeEar(x))
		str.WriteString("\n")
	}
	return n.write(strings.TrimSuffix(str.String(), "\n"))
}

func (n *NabConn) Led(data common.NabLedCmd) error {
	return n.write(encodeLed(data))
}

func (n *NabConn) Ear(data common.NabEarCmd) error {
	return n.write(encodeEar(data))
}

func (n *NabConn) Recorder(data common.NabAudioCmd) error {
	if data.State == common.NabAudioRunning {
		return n.write("06;1")
	} else if data.State == common.NabAudioStopped {
		return n.write("06;0")
	}
	return ErrUnknownCmd
}

func (n *NabConn) Player(data common.NabAudioCmd) error {
	if data.State == common.NabAudioRunning {
		return n.write("07;1")
	} else if data.State == common.NabAudioStopped {
		n.stopPlayer()
		return n.write("07;0")
	}

	if data.Link != "" {
		n.playLink(data.Link)
		return nil
	}

	if data.State == 0 && data.Link == "" {
		return n.write(fmt.Sprintf("07;2;%d", data.Volume))
	}

	return ErrUnknownCmd
}

func (n *NabConn) PlayerData(data []byte) error {
	return n.playStream(data)
}

func (n *NabConn) Reboot() error {
	return n.write("01")
}
