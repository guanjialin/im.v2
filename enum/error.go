package enum

import "fmt"

type ErrorResp struct {
	Code    int
	Content string
}

func New(code int, content string) ErrorResp {
	return ErrorResp{code, content}
}

func (e *ErrorResp) Error() string {
	return fmt.Sprintf("Code: %d, Content: %s\n",
		e.Code, e.Content)
}

// 基础错误
var (
	ErrorInvalidParameter = New(1, "无效的参数")
)

// 第三方接口错误
var (
	ErrorGithubInvalidCOde = New(1000, "Github返回无效的code")
	ErrorGithubAccessToken = New(1001, "从GitHub获取AccessToken失败")
)
