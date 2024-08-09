package sync

import (
	"sync"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		var res sync.Map
		wg.Add(100)
		for s := 0; s < 100; s++ {
			go func() {
				defer wg.Done()
				res.Store(s, s)
			}()
		}
		wg.Wait()
		aaa := make(map[int]int, 100)
		res.Range(func(key, value interface{}) bool {
			aaa[key.(int)] = value.(int)
			return true
		})

	}
}

func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(100)
		aaa := make(map[int]int, 100)
		var mu sync.Mutex
		for s := 0; s < 100; s++ {
			go func() {
				defer wg.Done()
				defer mu.Unlock()
				mu.Lock()
				aaa[s] = s
			}()
		}
		wg.Wait()
	}
}
