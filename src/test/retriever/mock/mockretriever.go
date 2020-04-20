package mock

import "fmt"

type MyRetriever struct {
	Context string
}

func (s *MyRetriever) String() string {
	return fmt.Sprintf("Strainger:{context: %s}", s.Context)
}

func (r *MyRetriever) Post(url string, data map[string]string) string {
	r.Context = data["context"]
	return "ok"
}

func (r *MyRetriever) Get(url string) string {
	return r.Context
}
