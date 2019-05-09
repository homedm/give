package main

import (
	"github.com/homedm/give/commands"
	_ "rsc.io/quote"
)

func main() {
	//fmt.Println(quote.Hello())
	commands.Run()
}
