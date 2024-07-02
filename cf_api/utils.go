package cfapi

import (
	"log"
	"runtime"
	"sync"
)

func trackSqlVar(sql string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			if sql != "" {
				log.Printf("Sql variable is set to %s", sql)
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	for {
		if sql == "" {
			runtime.ReadTrace()
			log.Panicf("Sql variable is set to %s", sql)
			break
		}
	}
}
