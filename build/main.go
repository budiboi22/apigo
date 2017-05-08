package main

import (
	"fmt"
	"os"
	"runtime"
	"github.com/budiboi22/apigo/app"
)

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	if err := app.Setup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
