package chromedriver

import "net/http"

const (
	ChromeDriverPort = `http://localhost:9515/`
)

func Get(url string) string {
	http.Get(url)
}
