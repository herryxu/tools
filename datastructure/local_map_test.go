
package datastructure

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheExpireTime(t *testing.T) {
	cache := NewLocalMap()

	// set value with expire
	cache.Set("key1", "value1", 0)
	cache.Set("key2", "value2", 0)
	localMap()
	// get value
	value1, ok1 := cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)

	value2, ok2 := cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)

	// wait for ....
	time.Sleep(2 * time.Second)

	// get value again
	ok1 = cache.CheckAndSet("key1", "value11111", 2*time.Second)
	value2, _ = cache.Get("key1")
	fmt.Println("Key1:", value2, ok1)

	value2, ok2 = cache.Get("key2")
	fmt.Println("Key2:", value2, ok2)
}
func localMap() {
	cache := NewLocalMap()
	// get value
	value1, ok1 := cache.Get("key1")
	fmt.Println("Key1:", value1, ok1)
}
