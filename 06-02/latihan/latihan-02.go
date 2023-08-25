package latihan

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func PrintPhone(phone []string, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Strings(phone)
	for i := 0; i <= len(phone)-1; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d. %s\n", i+1, phone[i])
	}
}
