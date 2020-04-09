package main

import (
	"fmt"

	ion "gitlab.fixstars.com/ion/ion-go"
)

func main() {
	_, _ = ion.NewBuilder()

	fmt.Println("vim-go")
}
