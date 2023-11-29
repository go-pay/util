package snowflake

import (
	"log"
	"testing"
	"time"
)

func TestNewNode(t *testing.T) {
	node, err := NewNode(1)
	if err != nil {
		log.Printf("err:%v.\n", err)
		return
	}
	for i := 0; i < 20; i++ {
		go func() {
			id := node.Generate().Int64()
			log.Printf("id:%d \n", id)
		}()
	}
	time.Sleep(time.Second)
}
