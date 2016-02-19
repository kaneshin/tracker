package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kaneshin/tracker"
)

var _ = log.Logger{}
var _ = os.Stdout

const filename = "/tmp/tracker.log"

func init() {
	// Set logger to write for the file
	// f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("error opening file:", err.Error())
	// }
	// log.SetOutput(f)
	// tracker.Logger = log.New(f, "[TRACKER] ", 0)

	// Customize form by the duration
	tracker.Format = func(d time.Duration) string {
		sec := float64(d.Nanoseconds()) / float64(time.Second)
		return fmt.Sprintf("%.3f seconds", sec)
	}

	// Available to set hook function
	tracker.Hook = func(d time.Duration) {
		// Do something
	}
}

func Get(url string) ([]byte, error) {
	defer tracker.Time(time.Now(), "HTTP Request:")

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	Get("http://httpbin.org/get")
	Get("http://httpbin.org/get")
	Get("http://httpbin.org/get")
	Get("http://httpbin.org/get")
	Get("http://httpbin.org/get")
}
