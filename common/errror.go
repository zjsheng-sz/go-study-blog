package common

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppErr(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var (
	// 系统级错误
	ErrInternal     = NewAppErr(500, "服务器内部错误", nil)
	ErrNotFound     = NewAppErr(404, "资源未找到", nil)
	ErrUnauthorized = NewAppErr(401, "未授权访问", nil)
	ErrForbidden    = NewAppErr(403, "禁止访问", nil)

	// 用户相关错误
	ErrInvalidInput    = NewAppErr(400, "无效的输入参数", nil)
	ErrUserNotFound    = NewAppErr(404, "用户不存在", nil)
	ErrUserExists      = NewAppErr(409, "用户已存在", nil)
	ErrInvalidPassword = NewAppErr(400, "密码错误", nil)

	// 文章相关错误
	ErrPostNotFound     = NewAppErr(404, "文章不存在", nil)
	ErrPostCreateFailed = NewAppErr(500, "创建文章失败", nil)

	// 评论相关错误
	ErrCommentNotFound     = NewAppErr(404, "评论不存在", nil)
	ErrCommentCreateFailed = NewAppErr(500, "创建评论失败", nil)

	// 数据库相关错误
	ErrDBOperation = NewAppErr(500, "数据库操作失败", nil)
)
