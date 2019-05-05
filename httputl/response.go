package httputl

type Response struct {
	Code   ResponseCode `json:"code"`
	Msg    ResponseMsg  `json:"msg"`
	Result interface{}  `json:"result"`
}

type RespArray struct {
	PageIndex int           `json:"pageIndex"`
	PageCount int           `json:"pageCount"`
	PageSize  int           `json:"pageSize"`
	Total     int           `json:"total"`
	Data      []interface{} `json:"data"`
}

func RespArraySuccess(index, size, total int, data []interface{}) Response {
	count := total / size
	if total % size > 0 {
		count += 1
	}
	return Response{
		Code: RPCD_Success,
		Msg:  RPSTR_SUCCESS,
		Result: RespArray{
			PageIndex: index,
			PageCount: count,
			PageSize:  size,
			Total:     total,
			Data:      data,
		},
	}
}

func RespSuccess(result interface{}) Response {
	return Response{
		Code:   RPCD_Success,
		Msg:    RPSTR_SUCCESS,
		Result: result,
	}
}

func RespFailed(code int, msg string) Response {
	return Response{
		Code: ResponseCode(code),
		Msg:  ResponseMsg(msg),
	}
}

func RespParamNoFound(paramKey string) Response {
	return Response{
		Code: RPCD_ParamNoFound,
		Msg:  RPSTR_ParamNoFound + "paramKey",
	}
}

func RespDefaultFailed() Response {
	return Response{
		Code: RPCD_Failed,
		Msg:  RPSTR_FAILED,
	}
}

func RespDefaultSuccess() Response {
	return Response{
		Code: RPCD_Success,
		Msg:  RPSTR_SUCCESS,
	}
}

func Resp404Failed() Response {
	return Response{
		Code: RPCD_PathNoFound,
		Msg:  RPSTR__PathNoFound,
	}
}

type ResponseCode int

const (
	RPCD_Success ResponseCode = iota
	RPCD_Failed
	RPCD_ServerError
	RPCD_ParamNoFound

	RPCD_PathNoFound
)

type ResponseMsg string

const (
	RPSTR_SUCCESS      ResponseMsg = "成功！"
	RPSTR_FAILED       ResponseMsg = "失败！"
	RPSTR_ServerError  ResponseMsg = "服务器内部错误！"
	RPSTR_ParamNoFound ResponseMsg = "缺少请求参数："
	RPSTR__PathNoFound ResponseMsg = "404，未找到请求路径！"
)
