package resolver

import (
	"crypto/sha256"
	"encoding/binary"
	"log"
)

func hashKey(key string) uint16 {
	sh := sha256.New()
	_, err := sh.Write([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	return binary.LittleEndian.Uint16(sh.Sum(nil))
}
