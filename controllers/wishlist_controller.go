package controllers

import (
	"encoding/json"
	"fmt"
	mysqldb "go_ecommerce/connection"
	"go_ecommerce/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Addwishlist(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]

	if userid == nil {
		c.Redirect(301, "/")
	} else {
		db := mysqldb.SetupDB()
		fmt.Println("sess uid", userid)
		pid := c.Request.PostFormValue("product_id")

		fmt.Println("pid:", pid)

		created_date := time.Now().Format("Jan 2, 2006 03:04:05 PM")

		var count int

		er := db.QueryRow("SELECT count(product_id) cnt FROM tbl_wishlist WHERE user_id = ? AND product_id = ?", userid, pid).Scan(&count)
		if er != nil {
			fmt.Println(er)
		}
		fmt.Println("count", count)
		if count != 1 {

			insWish, err := db.Prepare("INSERT INTO  tbl_wishlist(product_id,user_id,created_date) VALUES(?,?,?)")

			if err != nil {
				fmt.Println(err)
			}
			insWish.Exec(pid, userid, created_date)

			json.NewEncoder(c.Writer).Encode("Added in your wishlist")

		} else {

			json.NewEncoder(c.Writer).Encode("Already in your wishlist")
		}
		defer db.Close()
	}

}

func Wishlist(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]
	fmt.Println("userid", userid)
	if userid == nil {
		c.Redirect(301, "/viewwishlist")
	} else {
		db := mysqldb.SetupDB()

		WishRows, err := db.Query("SELECT product_id,id FROM tbl_wishlist where user_id=?", userid)

		if err != nil {
			fmt.Println(err)
		}
		wishlist := models.Products{}
		res1 := []models.Products{}
		for WishRows.Next() {
			var wid, product_id int
			_ = WishRows.Scan(&product_id, &wid)

			var title, image_path string
			var price float32
			var quantity int
			_ = db.QueryRow("SELECT title,price,image_path FROM tbl_products WHERE id=?", product_id).Scan(&title, &price, &image_path)

			db.QueryRow("SELECT quantity FROM tbl_orders WHERE product_id=? AND user_id=?", product_id, userid).Scan(&quantity)
			wishlist.Title = title

			wishlist.Price = price
			wishlist.Imagepath = image_path
			wishlist.ID = wid
			wishlist.Overallrating = product_id
			wishlist.Value = quantity
			res1 = append(res1, wishlist)
		}
		var ordercount int
		_ = db.QueryRow("SELECT count(id)  FROM tbl_orders WHERE user_id = ? ", userid).Scan(&ordercount)

		fmt.Println("ordercount", ordercount)

		name := session.Values["firstname"]

		c.HTML(200, "wishlist.html", gin.H{"wishlist": res1, "orderscount": ordercount, "name": name})

		defer db.Close()
	}

}

func Deletewishlist(c *gin.Context) {
	db := mysqldb.SetupDB()
	wid := c.Query("wid")
	fmt.Println("wid", wid)
	delWish, err := db.Prepare("DELETE FROM  tbl_wishlist WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delWish.Exec(wid)

	defer db.Close()
	c.Redirect(301, "/wishlist")

}
func Deletewishlistall(c *gin.Context) {
	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]

	fmt.Println("sess user:", userid)

	if userid == nil {
		c.Redirect(301, "/")
	}
	db := mysqldb.SetupDB()
	delWish, err := db.Prepare("DELETE FROM  tbl_wishlist WHERE user_id=?")
	if err != nil {
		panic(err.Error())
	}
	delWish.Exec(userid)

	defer db.Close()
	c.Redirect(301, "/wishlist")
}
