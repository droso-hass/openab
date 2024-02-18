package v2

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/drosocode/openab/internal/utils"
)

var RecDataSize = 4096

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	if sdata[0:2] == "08" {
		err := n.handleRecording(sdata[3:])
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(sdata)
}

func (n *NabConn) handleRecording(rawdata string) error {
	if len(n.recData) >= RecDataSize {
		wav, err := makeWav(n.recData)
		if err != nil {
			return err
		}
		filedata, err := convertRecording(wav)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("./rec/%s_%d.wav", strings.Replace(n.mac, ":", "", -1), time.Now().Unix())
		err = utils.WriteFile(filename, filedata)
		if err != nil {
			return err
		}
		n.recData = []byte{}
	}
	h, err := hex.DecodeString(rawdata)
	if err != nil {
		return err
	}
	n.recData = append(n.recData, h...)
	return nil
}
