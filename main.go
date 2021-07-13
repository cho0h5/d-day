package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now().In(loc)
	end := time.Date(2022, 11, 17, 0, 0, 0, 0, loc)
	fmt.Println(now)
	fmt.Println(end)
	fmt.Println()
	fmt.Println(end.Sub(now).Hours() / 24)
}
