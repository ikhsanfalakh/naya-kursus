package latihan

import (
	"sync"
	"time"
)

func GetMovies(movies []string, moviesChannel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, movie := range movies {
		moviesChannel <- movie
		time.Sleep(time.Second) // delay untuk simulasi
	}
}
