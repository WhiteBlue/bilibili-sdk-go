package main

import (
	. "github.com/bilibili-service/lib"
	"fmt"
)


func main() {
	r := NewBiliClient()

	back, err := r.GetBangumi("2")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(back)
}