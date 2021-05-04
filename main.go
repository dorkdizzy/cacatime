package main

import (
	"fmt"
	"time"

	"github.com/pkg/browser"
)

func main() {
	currentTimeSecond := time.Now().Second()
	if currentTimeSecond == 2 {
		fmt.Println("It is almost caca time, tiens bon !")
	} else if currentTimeSecond == 3 {
		fmt.Println("It is caca time!!!!!")
		browser.OpenURL("https://img.huffingtonpost.com/asset/5c93860b220000c9001c475b.jpeg")
	} else {
		fmt.Println("Je suis desole poto, it is not caca time :((((((")
	}
}
