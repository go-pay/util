package lru

import (
	"log"
	"testing"
	"time"
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

func TestCacheWithTimeout(t *testing.T) {
	cache := NewCacheWithTimeout(3, time.Second*3)
	//cache := lru.NewCache(3)
	cache.Put("1", "one")
	cache.Put("2", "two")
	cache.Put("3", "three")
	//cache.Set("4", "four")
	//cache.Set("5", "five")
	//cache.Set("6", "six")
	//log.Println("key=1: ", cache.Get("1")) // one
	//log.Println("key=2: ", cache.Get("2")) // nil
	//log.Println("key=3: ", cache.Get("3")) // nil
	//log.Println("key=4: ", cache.Get("4")) // four
	//log.Println("key=5: ", cache.Get("5")) // five
	//log.Println("key=6: ", cache.Get("6")) // six

	log.Println("key=1: ", cache.Get("1")) // one
	log.Println("key=2: ", cache.Get("2")) // two
	log.Println("key=3: ", cache.Get("3")) // three
	cache.Put("4", "four")
	log.Println("====1=====")
	log.Println("key=1: ", cache.Get("1")) // nil
	log.Println("key=2: ", cache.Get("2")) // two
	log.Println("key=3: ", cache.Get("3")) // three
	log.Println("key=4: ", cache.Get("4")) // four
	cache.Put("5", "five")
	log.Println("====2=====")
	log.Println("key=1: ", cache.Get("1")) // nil
	log.Println("key=2: ", cache.Get("2")) // nil
	log.Println("key=3: ", cache.Get("3")) // three
	log.Println("key=4: ", cache.Get("4")) // four
	log.Println("key=5: ", cache.Get("5")) // five
	time.Sleep(time.Second*3 + 1)
	log.Println("====3=====")
	log.Println("key=1: ", cache.Get("1")) // nil
	log.Println("key=2: ", cache.Get("2")) // nil
	log.Println("key=3: ", cache.Get("3")) // nil
	log.Println("key=4: ", cache.Get("4")) // nil
	log.Println("key=5: ", cache.Get("5")) // nil
	cache.Put("5", "five")
	cache.Put("4", "four")
	cache.Put("2", "two")
	log.Println("====4=====")
	log.Println("key=1: ", cache.Get("1")) // nil
	log.Println("key=2: ", cache.Get("2")) // two
	log.Println("key=3: ", cache.Get("3")) // nil
	log.Println("key=4: ", cache.Get("4")) // four
	log.Println("key=5: ", cache.Get("5")) // five
	log.Println("=====5====")
	cache.Put("1", "one")
	log.Println("key=1: ", cache.Get("1")) // one
	log.Println("key=2: ", cache.Get("2")) // nil
	log.Println("key=3: ", cache.Get("3")) // nil
	log.Println("key=4: ", cache.Get("4")) // four
	log.Println("key=5: ", cache.Get("5")) // five
}
