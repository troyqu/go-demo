package lb

import (
	"fmt"
	"runtime"
)

func ShowGoutin(){
	runtime.GOMAXPROCS(1)
	for i := 0; i<10; i++ {
		//fmt.Println("show i %d ",i)
		i := i
		//fmt.Println("show i 2 %d ",i)
		go func() {
			fmt.Println(i)
		}()
	}
	var ch = make(chan int)
	<- ch
}