package api

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 初始化App
func initApp(r *gin.Engine) {
}
