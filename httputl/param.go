package httputl

type MethodType int

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
	ANY
)

type Param struct {
	Method  MethodType
	Url     string
	Header  map[string]string
	Body    map[string]interface{}
	Failed  func(int, error)
	Success func(string)
}
