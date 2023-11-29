package lru

import (
	"log"
	"testing"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(2)

	cache.Put("1", "one")
	log.Println(cache.Get("1"))

	log.Println("===============")
	cache.Put("2", "two")
	log.Println(cache.Get("1"))
	log.Println(cache.Get("2"))

	log.Println("===============")
	cache.Put("3", "three")
	log.Println(cache.Get("1")) // nil
	log.Println(cache.Get("2"))
	log.Println(cache.Get("3"))

	log.Println("===============")
	cache.Put("1", "one")
	log.Println(cache.Get("2")) // nil
	log.Println(cache.Get("3"))
	log.Println(cache.Get("1"))

	log.Println("===============")
	cache.Put("2", "two")
	log.Println(cache.Get("1"))
	log.Println(cache.Get("2"))
	log.Println(cache.Get("3")) // nil
}
