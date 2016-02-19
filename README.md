# Tracker

## Installation

```shell
go get github.com/kaneshin/tracker
```


## Usage

### Example

```go
import "github.com/kaneshin/tracker"

func init() {
	// Set logger to write for the file
	f, _ := os.OpenFile("/tmp/tracker.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
	tracker.Logger = log.New(f, "[TRACKER] ", 0)

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
	// To calculate time this function.
	defer tracker.Time(time.Now(), "HTTP Request:")

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	Get("http://httpbin.org/get")
}
```


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
