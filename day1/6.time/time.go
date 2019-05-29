package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now().Format("2006년01-02-15-04-05")
	p(now)
}
