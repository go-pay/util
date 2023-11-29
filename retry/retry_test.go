package retry

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	err := Retry(func() error {
		log.Println("retry func")
		return errors.New("please retry")
	}, 3, 2*time.Second)
	if err != nil {
		log.Println(err)
	}
}
