package main

import (
	"fmt"

	"github.com/scm1400/learngo/banking"
)

func main() {

	account := banking.Account{Owner: "muice", Balance: 1000}

	fmt.Println(account)
}
