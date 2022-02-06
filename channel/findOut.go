package channel

import (
	"fmt"
	"sync"
	"time"
)

type FindOut struct{}

func NewFindOut() *FindOut {
	return &FindOut{}
}

func (*FindOut) Exec(wg *sync.WaitGroup, src <-chan interface{}, dstList ...chan interface{}) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			for i := range dstList {
				close(dstList[i])
			}
		}()
		for v := range src {
			for i := range dstList {
				dstList[i] <- v
			}
		}
	}()
}

func (*FindOut) WriteMysql(wg *sync.WaitGroup) chan interface{} {
	wg.Add(1)
	ret := make(chan interface{})
	go func() {
		defer wg.Done()
		// write execute
		time.Sleep(time.Second * 2)
		for v := range ret {
			fmt.Println("写入mysql", v)
		}
	}()
	return ret
}

func (*FindOut) WriteRedis(wg *sync.WaitGroup) chan interface{} {
	wg.Add(1)
	ret := make(chan interface{})
	go func() {
		defer wg.Done()
		// write execute
		time.Sleep(time.Second * 2)
		for v := range ret {
			fmt.Println("写入redis", v)
		}
	}()
	return ret
}
