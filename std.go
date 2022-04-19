package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

const message = "Hello Go"

type httpWriter struct {
	client *http.Client
	addr   string
}

func newHttpWriter(addr string) *httpWriter {
	return &httpWriter{
		client: http.DefaultClient,
		addr:   addr,
	}
}

func (w httpWriter) Write(p []byte) (n int, err error) {
	_, err = w.client.Post(w.addr, "text/plain", bytes.NewBuffer(p))
	return
}

func main() {
	file, _ := os.OpenFile("message.txt", os.O_CREATE|os.O_RDWR, 0644)
	buffer := &bytes.Buffer{}

	httpWriter := newHttpWriter("http://localhost:1111/")

	writeMessageEncoded(file, message)
	writeMessageEncoded(buffer, message)
	writeMessageEncoded(httpWriter, message)

	fmt.Printf("Buffer: %v", buffer.String())
}

func writeMessageEncoded(writer io.Writer, message string) {
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	writer.Write([]byte(encodedMessage))
}
