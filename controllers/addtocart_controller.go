package controllers

import (
	"encoding/json"
	"fmt"
	mysqldb "go_ecommerce/connection"
	"go_ecommerce/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Addtoorder(c *gin.Context) {

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

		quantity := 1

		var count int

		_ = db.QueryRow("SELECT count(product_id) cnt FROM tbl_orders WHERE user_id = ? AND product_id = ?", userid, pid).Scan(&count)

		if count != 1 {
			insOrder, err := db.Prepare("INSERT INTO  tbl_orders(product_id,user_id,quantity,created_date) VALUES(?,?,?,?)")

			if err != nil {
				fmt.Println(err)
			}
			insOrder.Exec(pid, userid, quantity, created_date)

			var ordercount int
			_ = db.QueryRow("SELECT count(id)  FROM tbl_orders WHERE user_id = ? ", userid).Scan(&ordercount)

			m := map[string]interface{}{
				"msg":   "Added in your Cartlist",
				"count": ordercount,
			}

			json.NewEncoder(c.Writer).Encode(m)

		} else {

			m := map[string]interface{}{
				"msg": "This product already in your cartlist",
			}
			json.NewEncoder(c.Writer).Encode(m)
		}
		defer db.Close()
	}

}

func Viewmyorders(c *gin.Context) {
	session, _ := store.Get(c.Request, "mysession")

	userid := session.Values["userid"]

	name := session.Values["firstname"]

	if userid == nil {
		c.Redirect(301, "/")
	} else {
		db := mysqldb.SetupDB()

		fmt.Println("sess user:", userid)

		orderRows, err := db.Query("SELECT tbl_orders.created_date	,tbl_orders.id, title,price,image_path FROM tbl_products INNER JOIN tbl_orders ON tbl_orders.product_id=tbl_products.id WHERE tbl_orders.user_id=?", userid)

		if err != nil {
			fmt.Println(err)
		}

		orderlist := models.Products{}
		res2 := []models.Products{}
		for orderRows.Next() {
			var title, image_path, created_date string
			var price float32
			var orderid int
			err = orderRows.Scan(&created_date, &orderid, &title, &price, &image_path)
			if err != nil {
				fmt.Println(err)
			}

			orderlist.Title = title
			orderlist.Price = price
			orderlist.Imagepath = image_path
			orderlist.ID = orderid
			orderlist.Createddate = created_date
			res2 = append(res2, orderlist)
		}
		var ordercount int
		_ = db.QueryRow("SELECT count(id)  FROM tbl_orders WHERE user_id = ? ", userid).Scan(&ordercount)

		c.HTML(200, "myorder.html", gin.H{"orderlist": res2, "orderscount": ordercount, "name": name})

		defer db.Close()
	}
}
