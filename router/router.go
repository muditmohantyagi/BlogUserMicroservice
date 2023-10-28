package router

import (
	"blog.com/controller"
	"blog.com/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("POST", "GET", "PUT", "DELETE", "UPDATE", "OPTIONS")

	// Register the middleware
	r.Use(cors.New(corsConfig))
	/**Allow origin CORS setting end:**/

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})
	/*-------------Routeing started---------------*/
	user := r.Group("api/v1/user")
	{
		var UserController = new(controller.UserController)
		user.POST("/register", UserController.Register)
		user.POST("/login", UserController.Login)
	}
	userAuth := r.Group("api/v1/user")
	userAuth.Use(middleware.AuthorizeJWT())
	{
		var UserController = new(controller.UserController)
		userAuth.POST("/logout", UserController.Logout)
	}

	return r

}
