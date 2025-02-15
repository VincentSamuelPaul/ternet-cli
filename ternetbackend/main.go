package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var routes = map[string]string{
	"/home": "GET",
	"/loginuser": "POST",
	"/createuser": "POST",
	"/createpost": "POST",
	"/getposts": "GET",
	"/likeposts": "POST",
}

func getRoutes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, routes)
}

var DB *sql.DB

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	State bool
}

type post struct {
	Username string `json:"username"`
	Data string `json:"data"`
	Likes int `json:"likes"`
}

func loginUser(c *gin.Context) {
	var userObj user
	if err := c.ShouldBindJSON(&userObj); err != nil {
		log.Println(err)
	}
	userObj.State = false
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	query := `SELECT * FROM USERS;`
	data, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()
	var username string
	var password string
	for data.Next() {
		data.Scan(&username, &password)
		if userObj.Username == username && userObj.Password == password {
			userObj.State = true
			break
		}
	}
	c.JSON(http.StatusOK, gin.H {
		"Status": userObj.State,
	})
}

func createUser(c *gin.Context) {
	var userObj user
	if err := c.ShouldBindJSON(&userObj); err != nil {
		log.Println(err)
	}
	userObj.State = true
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	query := `SELECT username, password FROM USERS;`
	data, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()
	var username string
	var password string
	for data.Next() {
		data.Scan(&username, &password)
		if username == userObj.Username {
			userObj.State = false
			break
		}
	}
	if userObj.State {
		query := `INSERT INTO USERS VALUES (` + `"` + userObj.Username + `","` + userObj.Password + `");`
		_, err := DB.Exec(query)
		if err != nil {
			log.Println(err)
		}
	}
	c.JSON(http.StatusOK, gin.H {
		"Status": userObj.State,
	})
}

func createPost(c *gin.Context) {
	var postObj post
	if err := c.ShouldBindJSON(&postObj); err != nil {
		log.Println(err)
	}
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	query := `INSERT INTO POSTS VALUES (` + `"` + postObj.Username + `","` + postObj.Data + `","0");`
	_, err = DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()
	c.JSON(http.StatusOK, gin.H {
		"Status": true,
	})
}

func getPosts(c *gin.Context) {
	var postObj post
	var posts []post
	if err := c.ShouldBindJSON(&postObj); err != nil {
		log.Println(err)
	}
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	query := `SELECT * FROM POSTS;`
	Data, err := DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()
	var username string
	var data string
	var likes int
	for Data.Next() {
		Data.Scan(&username, &data, &likes)
		postObj.Username = username
		postObj.Data = data
		postObj.Likes = likes
		posts = append(posts, postObj)
	}
	c.JSON(http.StatusOK, posts)
}

func likePost(c *gin.Context) {
	var postObj post
	if err := c.ShouldBindJSON(&postObj); err != nil {
		log.Println(err)
	}
	DB, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	query := `UPDATE POSTS SET likes = likes + 1 WHERE username = ` + `"` + postObj.Username + `" AND data = "` + postObj.Data + `";`
	_, err = DB.Exec(query)
	if err != nil {
		log.Println(err)
	}
	defer DB.Close()
	postObj.Likes += 1
	c.JSON(http.StatusOK, gin.H {
		"Status": true,
	})
}

func main() {
	router := gin.Default()
	router.GET("/routes", getRoutes)
	router.POST("/loginuser", loginUser)
	router.POST("/createuser", createUser)
	router.GET("/getposts", getPosts)
	router.POST("/createpost", createPost)
	router.POST("/likepost", likePost)
	router.Run(":3000")
}