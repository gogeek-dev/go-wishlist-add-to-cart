package gowishlistaddtocart

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type TblProduct struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Title         string
	Description   string
	Price         float32
	Image         string
	ImagePath     string
	OverallRating int
	WishlistCount int
	CreatedDate   string
	ModifiedDate  string
	Value         int `gorm:"-"`
	Productcount  int `gorm:"-"`
}

type TblUser struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	FirstName     string
	MobileNumber  string
	EmailId       string
	Password      string
	CreatedDate   string
	ModifiedDate  string
	LastLoginDate string
}

type TblWishlist struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ProductId   int
	UserId      int
	CreatedDate string
}

type TblOrder struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	ProductId   int
	UserId      int
	Quantity    int
	CreatedDate string
}

type TblRatingReview struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	ProductId   int
	UserId      int
	Rating      int
	Review      string
	CreatedDate string
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

var DB *gorm.DB

func AddroutesAndDbMigrate(Db *gorm.DB, P *gin.RouterGroup) *gin.RouterGroup {

	err := Db.AutoMigrate(
		&TblOrder{},
		&TblProduct{},
		&TblUser{},
		&TblWishlist{},
		&TblRatingReview{},
	)

	if err != nil {
		panic(err)
	}

	DB = Db

	P.GET("/index", Productlist)
	P.POST("/addwishlists", Addwishlist)
	P.POST("/addtoorder", Addtoorder)
	P.GET("/wishlist", Wishlist)
	P.GET("/deletewishs", Deletewishlist)
	P.GET("/removeall", Deletewishlistall)
	P.GET("/signout", Logout)
	P.GET("/myorders", Viewmyorders)

	return P
}

func Productlist(c *gin.Context) {
	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]

	userid, _ := c.Get("userid")

	fmt.Println("sess userid", userid)

	if userid == nil {

		c.Redirect(301, "/")

	} else {

		var prod []TblProduct

		err := DB.Debug().Table("tbl_products").Find(&prod)

		if err.Error != nil {
			panic(err.Error)
		}

		var ordercount int64

		DB.Debug().Model(TblOrder{}).Count(&ordercount)

		firstname, _ := c.Get("username")

		fmt.Println("ordercount", ordercount)

		c.HTML(200, "products.html", gin.H{"products": prod, "orderscount": ordercount, "name": firstname})
	}

}

func Addwishlist(c *gin.Context) {

	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]

	userid, _ := c.Get("userid")

	if userid == nil {

		c.Redirect(301, "/")

	} else {

		fmt.Println("sess uid", userid)

		pid, _ := strconv.Atoi(c.Request.PostFormValue("product_id"))

		fmt.Println("pid:", pid)

		var count int64

		er := DB.Table("tbl_wishlists").Where("user_id = ? AND product_id = ?", userid, pid).Count(&count)

		if er != nil {
			fmt.Println(er)
		}
		fmt.Println("count", count)

		if count != 1 {

			DB.Create(&TblWishlist{
				ProductId:   pid,
				UserId:      userid.(int),
				CreatedDate: time.Now().Format("Jan 2, 2006 03:04:05 PM"),
			})

			json.NewEncoder(c.Writer).Encode("Added in your wishlist")

		} else {

			json.NewEncoder(c.Writer).Encode("Already in your wishlist")
		}

	}

}

func Addtoorder(c *gin.Context) {

	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]
	userid, _ := c.Get("userid")

	if userid == nil {

		c.Redirect(301, "/")

	} else {

		var order TblOrder

		pid, _ := strconv.Atoi(c.Request.PostFormValue("product_id"))

		order.CreatedDate = time.Now().Format("Jan 2, 2006 03:04:05 PM")

		order.Quantity = 1

		order.ProductId = pid

		order.UserId = userid.(int)

		var count int64

		er := DB.Debug().Table("tbl_orders").Where("user_id = ? AND product_id = ?", userid, pid).Count(&count)

		if er.Error != nil {
			log.Println(er.Error)
		}

		if count != 1 {

			err := DB.Table("tbl_orders").Create(&order)

			if err.Error != nil {

				log.Println(err.Error)

			}

			var ocount int64

			DB.Table("tbl_orders").Where("user_id = ?", userid).Count(&ocount)

			m := map[string]interface{}{
				"msg":   "Added in your Cartlist",
				"count": ocount,
			}

			json.NewEncoder(c.Writer).Encode(m)

		} else {

			m := map[string]interface{}{
				"msg": "This product already in your cartlist",
			}
			json.NewEncoder(c.Writer).Encode(m)
		}
	}

}

func Wishlist(c *gin.Context) {

	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]
	userid, _ := c.Get("userid")
	fmt.Println("userid", userid)
	if userid == nil {
		c.Redirect(301, "/viewwishlist")
	} else {

		var wishlist []TblWishlist

		err := DB.Table("tbl_wishlists").Where("user_id=?", userid.(int)).Find(&wishlist)

		if err != nil {
			fmt.Println(err)
		}

		var prod1 []TblProduct

		for _, val := range wishlist {

			var pro TblProduct

			er := DB.Table("tbl_products").Where("id=?", val.ProductId).First(&pro)

			if er.Error != nil {
				panic(er.Error)
			}

			var ord TblOrder

			DB.Table("tbl_orders").Select("quantity").Where("product_id=? AND user_id=?", val.ProductId, userid).Find(&ord)

			var prod TblProduct

			prod.ID = pro.ID

			prod.Title = pro.Title

			prod.Price = pro.Price

			prod.ImagePath = pro.ImagePath

			prod.OverallRating = pro.ID

			prod.Value = ord.Quantity

			prod1 = append(prod1, prod)

		}

		var ordercount int64

		DB.Table("tbl_orders").Where("user_id=?", userid.(int)).Count(&ordercount)

		fmt.Println("ordercount", ordercount)

		name, _ := c.Get("username")

		c.HTML(200, "wishlist.html", gin.H{"wishlist": prod1, "orderscount": ordercount, "name": name})

	}

}

func Deletewishlist(c *gin.Context) {

	wid, _ := strconv.Atoi(c.Query("wid"))

	err := DB.Debug().Where("product_id=?", wid).Delete(&TblWishlist{})

	if err.Error != nil {
		panic(err.Error)
	}

	c.Redirect(301, "/wishlist")

}

func Deletewishlistall(c *gin.Context) {

	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]
	userid, _ := c.Get("userid")

	fmt.Println("sess user:", userid)

	if userid == nil {
		c.Redirect(301, "/")
	}

	err := DB.Where("user_id=?", userid.(int)).Delete(TblWishlist{})

	if err != nil {
		panic(err.Error)
	}

	c.Redirect(301, "/wishlist")
}

func Logout(c *gin.Context) {
	time.Sleep(2 * time.Second)
	session, err := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	if err != nil {
		fmt.Println(err)
	}
	session.Values["firstname"] = ""

	session.Values["userid"] = ""

	session.Options.MaxAge = -1

	er := session.Save(c.Request, c.Writer)
	if er != nil {
		fmt.Println(er)
	} else {

		c.HTML(200, "login.html", nil)
	}
}

func Viewmyorders(c *gin.Context) {

	// session, _ := store.Get(c.Request, os.Getenv("SESSION_KEY"))

	// userid := session.Values["userid"]
	userid, _ := c.Get("userid")

	name, _ := c.Get("username")

	if userid == nil {

		c.Redirect(301, "/")

	} else {

		var res2 []TblProduct

		err := DB.Table("tbl_products").Select("tbl_orders.created_date,tbl_orders.id,title,price,image_path").Joins("INNER JOIN tbl_orders on tbl_orders.product_id=tbl_products.id").Where("tbl_orders.user_id=?", userid).Find(&res2)

		if err.Error != nil {
			log.Println(err.Error)
		}

		var ordercount int64

		DB.Table("tbl_orders").Where("user_id = ?", userid.(int)).Count(&ordercount)

		c.HTML(200, "myorder.html", gin.H{"orderlist": res2, "orderscount": ordercount, "name": name})
	}
}
