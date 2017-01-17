package easyjson

var (
	//FormatError json 格式错误
	FormatError = 1
	//ParamError 字段类型错误
	ParamError = 2
)

// Error 格式错误
type Error struct {
	Type  int
	Param string
	Msg   string
}

func (e *Error) Error() string {
	return e.Param + ":" + e.Msg
}
