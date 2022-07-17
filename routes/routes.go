// 在这个文件中注册 URL handler

package routes

import "github.com/gin-gonic/gin"

// Routes 注册 API URL 路由
func Routes(app *gin.Engine) {
	app.GET("/", Index)
}
