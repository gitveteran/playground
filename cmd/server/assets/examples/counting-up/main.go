package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

func main() {
	n := rand.Intn(8) + 2
	fmt.Printf("n: %d\n", n)

	body := js.Global().Get("document").Get("body")
	body.Set("innerHTML", `<ul></ul>`)

	ul := body.Call("querySelector", "ul")
	for i := 0; i < n; i++ {
		ul.Call("insertAdjacentHTML", "beforeEnd", fmt.Sprintf("<li>%d</li>", i))
		time.Sleep(time.Second / 2)
	}
}
