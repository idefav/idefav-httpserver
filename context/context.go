package context

type Interface interface {
	PostForm(key string) string
	RequestBody(body interface{}) (interface{}, error)
	Query(key string) string
	Status(code int)
	Header(key string) string
	SetHeader(key string, value string)
	Success(code string, msg string, data interface{})
	Fail(code string, msg string)
}
