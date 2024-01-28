package v2

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"os"
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

func makeWav(data []byte, outfile string) error {
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

	out, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	_, err = out.Write(header)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}
	return nil
}

func convertRecording() {
	// ffmpeg -i in.wav -acodec pcm_s16le out.wav
}

func convertPlayer() {
	// ffmpeg -i in.wav -vn -ar 44100 -ac 1 -b:a 64k out.mp3
}
