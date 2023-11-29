package geohash

import (
	"log"
	"testing"
)

func TestEncode(t *testing.T) {

	geohash := Encode(31.2851847116, 121.5571761131, 10)
	log.Println(geohash) // wtw3yp71rm
}
