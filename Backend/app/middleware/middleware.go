package middleware

import "github.com/labstack/echo"

// Authenticate 无论什么请求都设置用户为默认用户
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Set("User", "default")
		return next(context)
	}
}
