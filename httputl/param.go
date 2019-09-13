package httputl

type MethodType int

const (
	GET MethodType = iota
	POST
	PUT
	DELETE
	ANY
)

type HttpParam struct {
	Method  MethodType
	Url     string
	Header  map[string]string
	Body    map[string]interface{}
	Result *HttpResponse
}

type HttpResponse struct {
	Err error
	Result interface{}
	IsSuccessChan chan bool
}
