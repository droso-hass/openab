package v2

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"log"
	"os"
	"os/exec"
)

var ConstRiff []byte
var ConstWaveFact []byte
var ConstData []byte

func init() {
	// from mkriff in std/reclib.mtl
	var err error
	ConstRiff, err = hex.DecodeString("52494646")
	if err != nil {
		log.Fatal(err)
	}
	ConstWaveFact, err = hex.DecodeString("57415645666D74201400000011000100401F0000D70F0000000104000200F9016661637404000000")
	if err != nil {
		log.Fatal(err)
	}
	ConstData, err = hex.DecodeString("64617461")
	if err != nil {
		log.Fatal(err)
	}
}

func makeWav(data []byte) ([]byte, error) {
	l := len(data)
	riffSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(riffSize, uint32(l+52))
	factSize := make([]byte, 4)
	binary.LittleEndian.PutUint32(factSize, uint32((l>>8)*505))
	size := make([]byte, 4)
	binary.LittleEndian.PutUint32(size, uint32(l))

	header := []byte{}
	header = append(header, ConstRiff...)
	header = append(header, riffSize...)
	header = append(header, ConstWaveFact...)
	header = append(header, factSize...)
	header = append(header, ConstData...)
	header = append(header, size...)

	return append(header, data...), nil
}

// ffmpeg -i in.wav -acodec pcm_s16le out.wav
func convertRecording(in []byte) ([]byte, error) {
	cmd := exec.Command("ffmpeg",
		"-hide_banner", "-loglevel", "error",
		"-i", "pipe:0",
		"-acodec", "pcm_s16le",
		"-f", "wav",
		"pipe:1",
	)
	out := new(bytes.Buffer)

	cmd.Stderr = os.Stderr
	cmd.Stdout = out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	_, err = stdin.Write(in)
	if err != nil {
		return nil, err
	}
	err = stdin.Close()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func convertPlayer(input string, ch chan []byte) error {
	// ffmpeg -i in.wav -vn -ar 44100 -ac 1 -b:a 64k out.mp3
	cmd := exec.Command("ffmpeg",
		"-hide_banner", "-loglevel", "error",
		"-i", input,
		"-vn",
		"-ar", "44100",
		"-ac", "1",
		"-b:a", "64k",
		"-f", "mp3",
		"pipe:1",
	)

	cmd.Stderr = os.Stderr
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		for {
			buf := make([]byte, 512)
			n, e := out.Read(buf)
			if e == io.EOF {
				break
			}
			ch <- buf[0:n]
		}
		ch <- nil
	}()
	go cmd.Wait()
	return nil
}
