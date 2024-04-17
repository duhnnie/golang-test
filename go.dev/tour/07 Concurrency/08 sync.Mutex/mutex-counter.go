package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeAccess struct {
	m sync.Mutex
	v map[string]int
}

func (s *SafeAccess) Inc(key string) {
	s.m.Lock()
	fmt.Println("Writting to key...")
	s.v[key]++
	fmt.Printf("Written %v!\n", s.v[key])
	time.Sleep(100 * time.Millisecond)
	s.m.Unlock()
}

func (s *SafeAccess) Value(key string) int {
	s.m.Lock()
	defer s.m.Unlock()

	return s.v[key]
}

func main() {
	s := SafeAccess{v: make(map[string]int)}

	for i := 0; i < 10; i++ {
		go s.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Printf("Value: %v\n", s.Value("somekey"))
}