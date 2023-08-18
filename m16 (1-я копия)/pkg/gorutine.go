package pkg

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Amount int

func Accumulate(value int) int {
	var mu = sync.Mutex{}
	for i := 0; i < 10; i++ {
		mu.Lock()
		v := value + rand.Intn(10)
		fmt.Println("Горутина(плюс) - ", i, " Баланс - ", v)
		time.Sleep(time.Millisecond)
		mu.Unlock()
		go func() {
			for i := 0; i < 5; i++ {
				mu.Lock()
				Amount = v - rand.Intn(5)
				if Amount < 0 {
					fmt.Println("Число z превышает число v. Баланс не может быть отрицательный!")
				} else {
					fmt.Println("Горутина(минус) - ", i, " Баланс - ", Amount)
				}
				time.Sleep(time.Millisecond)
				mu.Unlock()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println("Текущий баланс - ", Amount)

	return Amount
}
