package resp

import "fmt"

type Err struct {
	Code int
	Msg  string
}

func (e Err) Error() string {
	return fmt.Sprintf("err:%d, msg:%s", e.Code, e.Msg)
}

var (
	InternalErr    = Err{501, "服务器内部错误"}
	ParamMissErr   = Err{1001, "参数缺失"}
	ParamErr       = Err{1002, "参数错误"}
	ResultNotFound = Err{1003, "查无结果"}
)
