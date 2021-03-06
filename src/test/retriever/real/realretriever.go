package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Myretriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (m *Myretriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bytes, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
