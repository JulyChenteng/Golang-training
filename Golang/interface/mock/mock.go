package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	r.Contents = "hello"
	return r.Contents
}
