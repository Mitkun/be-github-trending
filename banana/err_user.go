package banana

import "errors"

var (
	UserConflict   = errors.New("Người dùng đã tồn tại")
	UserNotFound   = errors.New("Người dùng không tồn tại")
	SignUpFail     = errors.New("Đăng ký thất bại")
	UserNotUpdated = errors.New("Cập nhật thông tin người dùng thất bại")
)
