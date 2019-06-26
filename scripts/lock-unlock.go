package main
import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

type UnsafeCounter struct {
	v map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (u *UnsafeCounter) Inc(key string) {
	u.v[key]++
}

func (u *UnsafeCounter) Value(key string) int {
	return u.v[key]
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	d := SafeCounter{v: make(map[string]int)}
	u := UnsafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go d.Inc("somekey")
		//fmt.Println(d.Value("somekey"))
	//	go u.Inc("onekey")
		//fmt.Println(u.Value("onekey"))
	}
	time.Sleep(time.Second)
	fmt.Println(d.Value("somekey"))
	fmt.Println(u.Value("onekey"))
}
