package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	mysqldb "go_ecommerce/connection"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Home(c *gin.Context) {

	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {

	db := mysqldb.SetupDB()

	key := hex.EncodeToString(make([]byte, 32))

	emailid := c.Request.PostFormValue("email")

	Password := c.Request.PostFormValue("password")

	var password, first_name string
	var id int

	Userrow := db.QueryRow("SELECT id,first_name,password FROM  tbl_user WHERE email_id=?", emailid)

	err := Userrow.Scan(&id, &first_name, &password)

	if err == nil {

		decrypted := decrypt(password, key)

		if Password != decrypted {

			c.HTML(200, "login.html", gin.H{"data": "*your password is incorrect"})

		} else {
			session, _ := store.Get(c.Request, "mysession")

			session.Values["firstname"] = first_name

			session.Values["userid"] = id

			session.Save(c.Request, c.Writer)

			last_login_date := time.Now().Format("Jan 2, 2006 03:04:05 PM")

			updUser, err := db.Prepare("UPDATE tbl_user SET last_login_date=? WHERE email_id=?")

			if err != nil {
				panic(err.Error())
			}
			updUser.Exec(last_login_date, emailid)

			c.Redirect(301, "/index")
		}

	} else {

		c.HTML(200, "login.html", gin.H{"data": "*you are not registered"})

	}

}

func Registerpage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Register(c *gin.Context) {

	db := mysqldb.SetupDB()

	key := hex.EncodeToString(make([]byte, 32))

	firstname := c.Request.PostFormValue("firstname")

	email := c.Request.PostFormValue("email")

	mobilenumber := c.Request.PostFormValue("mobile")

	password := c.Request.PostFormValue("password")

	encryptpwd := encrypt(password, key)

	created_date := time.Now().Format("Jan 2, 2006 03:04:05 PM")

	var count int
	_ = db.QueryRow("SELECT COUNT(id) FROM tbl_user WHERE email_id= ?", email).Scan(&count)
	if count != 1 {

		insUser, err := db.Prepare("INSERT INTO tbl_user(first_name, email_id,mobile_number,password,created_date) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insUser.Exec(firstname, email, mobilenumber, encryptpwd, created_date)

		time.Sleep(1 * time.Second)

		var userid int
		var username string
		_ = db.QueryRow("SELECT id,first_name FROM  tbl_user WHERE email_id=?", email).Scan(&userid, &username)

		session, _ := store.Get(c.Request, "mysession")

		session.Values["firstname"] = firstname

		session.Values["userid"] = userid

		session.Save(c.Request, c.Writer)

		last_login_date := time.Now().Format("Jan 2, 2006 03:04:05 PM")

		updUser, err := db.Prepare("UPDATE tbl_user SET last_login_date=? WHERE email_id=?")

		if err != nil {
			panic(err.Error())
		}
		updUser.Exec(last_login_date, email)

		c.Redirect(301, "/index")
	} else {
		c.HTML(200, "register.html", gin.H{"data": "*Already used this Email id "})
	}

}

func Logout(c *gin.Context) {
	time.Sleep(2 * time.Second)
	session, err := store.Get(c.Request, "mysession")

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

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Since the key is in string, we need to convert decode it to byteslog.Println(password)
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
