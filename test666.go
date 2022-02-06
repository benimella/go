package main

import (
	"gds/channel"
	"sync"
)

func main() {
	l := channel.NewFindOut()
	src := make(chan interface{})
	wg := &sync.WaitGroup{}
	l.Exec(wg, src, l.WriteMysql(wg), l.WriteRedis(wg))
	//i := 0
	for i := 0; i <= 10; i++ {
		src <- i
	}
	go func() {
		defer close(src)
		wg.Wait()
	}()
	/*for {

	}*/
}
