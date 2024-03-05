package v2

import (
	"encoding/binary"
	"encoding/hex"
	"log"
)

var RecDataSize = 4096

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
