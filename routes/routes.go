package routes

import (
	"database/sql"
	"go_ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*/*.html")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.Home)
	r.POST("/", controllers.Login)
	r.GET("/register", controllers.Registerpage)
	r.POST("/register", controllers.Register)
	r.GET("/index", controllers.Productlist)
	r.POST("/addwishlists", controllers.Addwishlist)
	r.POST("/addtoorder", controllers.Addtoorder)
	r.GET("/wishlist", controllers.Wishlist)
	r.GET("/deletewishs", controllers.Deletewishlist)
	r.GET("/removeall", controllers.Deletewishlistall)
	r.GET("/signout", controllers.Logout)
	r.GET("/myorders", controllers.Viewmyorders)

	return r
}
