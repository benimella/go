package channel

// https://cloud.tencent.com/developer/article/1864041

import (
	"sync"
	"time"
)

type FindIn struct{}

func NewFindIn() *FindIn {
	return &FindIn{}
}

type FindInMethod func(args ...interface{}) <-chan interface{}

func (*FindIn) Exec(chs ...<-chan interface{}) chan interface{} {
	ret := make(chan interface{})
	wg := sync.WaitGroup{}
	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan interface{}) {
			defer wg.Done()
			for v := range c {
				ret <- v
			}
		}(ch)
	}
	go func() {
		defer close(ret)
		wg.Wait()
	}()
	return ret
}

func (*FindIn) GetUserList() <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		time.Sleep(time.Second * 3)
		ret <- "get-user-list"
	}()
	return ret
}

func (*FindIn) GetProductList() <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		time.Sleep(time.Second * 5)
		ret <- "get-product-list"
	}()
	return ret
}

func (*FindIn) GetHospitalList() <-chan interface{} {
	ret := make(chan interface{})
	go func() {
		defer close(ret)
		time.Sleep(time.Second * 2)
		ret <- "get-hospital-list"
	}()
	return ret
}
