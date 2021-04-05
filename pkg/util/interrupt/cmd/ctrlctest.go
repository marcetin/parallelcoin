package main

import (
	"fmt"
	
	"github.com/p9c/monorepo/pkg/util/interrupt"
)

func main() {
	interrupt.AddHandler(func() {
		fmt.Println("IT'S THE END OF THE WORLD!")
	},
	)
	<-interrupt.HandlersDone
}
