package main

import (
	"fmt"
	"github.com/alexrett/idler"
	"time"
)

func main() {
	i := idler.Idle{}
	for {
		fmt.Println(i.GetIdleTime())
		time.Sleep(time.Second * 1)
	}
}
