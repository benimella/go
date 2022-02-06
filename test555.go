package main

import (
	"fmt"
	"gds/channel"
)

func main() {
	l := channel.NewFindIn()
	ret := l.Exec(l.GetUserList(), l.GetProductList(), l.GetHospitalList())
	for v := range ret {
		fmt.Println(v)
	}
}
