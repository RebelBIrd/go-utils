# go-utils
> golang 工具集，持续更新
## fileutl
### Manager 
>文件管理器
- Name 文件名
- Path 文件存储路径
- Write(string) 写文件
- Create() error 创建文件
- Open(block func(error, *os.File)) 打开文件
- IsExist() bool 文件是否存在
- PathExistOrCreate() 文件是否存在，不存在创建
- ReadAll(block func([]byte, error)) 读取文件内容
## httputl
### GetParam(ctx *gin.Context, key string) interface{} 获取gin上传参数
### Response 网络请求结果
- Code 请求返回Code
- Msg 请求返回信息
### RespArray 分页返回结果
- PageIndex 页码
- PageCount 页数
- PageSize 每页数量
- Total 总条数
- Data []interface{} 结果
### RespArraySuccess(index, size, total int, data []interface{}) Response 
>成功返回分页数据
### RespSuccess(result interface{}) Response 
>返回成功结果
### RespFailed(code int, msg string) Response 
>返回失败结果
### RespParamNoFound(paramKey string) Response 
>返回参数没找到结果
### RespDefaultFailed() Response 
>返回默认失败
### RespDefaultSuccess() Response 
>返回默认成功无结果
### Resp404Failed() Response 
>返回404请求
### ResponseCode 
>返回Code
- RPCD_Success 成功
- RPCD_Failed 失败
- RPCD_ServerError 服务器内部错误
- RPCD_ParamNoFound 未找到参数
- RPCD_PathNoFound 未找到路径
### ResponseMsg 
>返回msg
- RPSTR_SUCCESS      ResponseMsg = "成功！"
- RPSTR_FAILED       ResponseMsg = "失败！"
- RPSTR_ServerError  ResponseMsg = "服务器内部错误！"
- RPSTR_ParamNoFound ResponseMsg = "缺少请求参数："
- RPSTR__PathNoFound ResponseMsg = "404，未找到请求路径！"
### DoHttp(methodType MethodType, url string, header map[string]string, body map[string]interface{}, response HttpResponse) 
>客户端网络请求
### HttpParam 
>网络请求参数
- Method MethodType 请求方式
    - GET
    - POST
    - PUT
    - DELETE
- Url 请求地址
- Header 请求头
- Body 请求体
- Result HttpResponse 请求结果
    - Err 错误返回
    - Result 成功返回，需要自定义
    - IsSuccessChan chan bool 请求是否成功
## logutl
- Error(...interface{}) 
>打印红色错误
- Debug(...interface{}) 
>打印路径调试信息
- Info(...interface{}) 
>打印白色内容
- Warning(...interface{}) 
>打印黄色警告
- init() 
>初始化log日志写入，将在当前目录的创建log子目录，生成当前时间的log文件
- Write(msgType, msg string) 
>写入log信息
## maputl
- Struct2Map(obj struct) map[string]interface{} 
>struct转map
- Map2Struct(mp map[string]interface{}, obj interface{}) (err error) 
>map转struct，obj为指针
- GetMapValue(m map[string]interface{}, key string, value interface{}) 
>error 取出map[string]interface{}中的value，注意：只能取基础类型， value为指针
## ormutl
```go
//mysql驱动器，带yaml配置项
type MysqlConf struct {
	Url      string `yaml:"url"`      // 数据库地址
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	Database string `yaml:"database"` // 数据库名
	*xorm.Engine                      // 数据库引擎
}
```
- InitMysql(MysqlConf) 
>初始化mysql
- GetEngine() *MysqlConf
>获取数据库引擎
- MysqlConf.InitTables(beans ...interface{})
>初始化数据库表
## redisutl
```go
// redis配置引擎
type RedisConf struct {
	Url      string `yaml:"url"`      //redis地址
	Username string `yaml:"username"` //redis用户名
	Password string `yaml:"password"` //redis密码
	Database int    `yaml:"database"` //redis数据库
	*redis.Client                     //redis客户端引擎
}
```
- GetRedis() *RedisConf
>获取redis引擎
- InitRedis(conf RedisConf)
>初始化redis引擎
- RedisConf.Save(key string, value interface{})
>保存
- RedisConf.IsExist(key string) bool
>判断key是否存在
- RedisConf.ExpireKey(key string) bool
>刷新redis
## sliceutl
- InArray(key interface{}, sli []interface{}) bool
> key是否在sli数组中
## strutl
- ConnString(strs ...string) string
> 连接string
- GetRandomString(size int) string 
> 获取size位数的随机string
## timeutl
- GetNowTime() string
> 获取当前时间