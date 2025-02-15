package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type post struct {
	Username string `json:"username"`
	Data string `json:"data"`
	Likes int `json:"likes"`
}

// type user struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

type Response struct {
	Status bool `json:"Status"`
}

type Status struct {
	Username string `json:"Username"`
}

// var URL1 = "http://localhost:3000/"
var URL = "https://ternetbackend1.onrender.com/"

func loginUser() (bool, string) {
    var username string
    var password string
    fmt.Print("username: :")
    fmt.Scan(&username)
    fmt.Print("password: :")
    fmt.Scan(&password)
    // fmt.Println(username, password)
	url := URL + "loginuser"
	data := map[string]string{"username": username, "password": password}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}
	return result.Status, username
}

func createUser() (bool, string) {
    var username string
    var password string
    fmt.Print("username: :")
    fmt.Scan(&username)
    fmt.Print("password: :")
    fmt.Scan(&password)
    // fmt.Println(username, password)
	url := URL + "createuser"
	data := map[string]string{"username": username, "password": password}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}
	return result.Status, username
}

func getPosts() []post {
	url := URL + "getposts"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var posts []post
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Println(err)
	}
	return posts
}

func createPost(username string) bool {
	url := URL + "createpost"
	// var data string
    fmt.Print("Write your post: :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dat := map[string]string{"username": username, "data": scanner.Text()}
	jsonDat, err := json.Marshal(dat)
	if err != nil {
		log.Println(err)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonDat))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}
	return result.Status
}

func likePost(username, data string) {
	url := URL + "likepost"
	// var username string
	// var data string
	dat := map[string]string{"username": username, "data": data}
	jsonDat, err := json.Marshal(dat)
	if err != nil {
		log.Println(err)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonDat))
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}
	// return result.Status
}


func main() {
	fmt.Print("\033[H\033[2J")
	help := flag.Bool("help", false, "Lists all the features available.")
	login := flag.Bool("login", false, "Login user.")
	logout := flag.Bool("logout", false, "Logout user.")
	signup := flag.Bool("signup", false, "Creates a new user.")
	newpost := flag.Bool("newpost", false, "Create new post.")
	posts := flag.Bool("browse", false, "Browse posts.")
	flag.Parse()

	switch {

	case *help:
		fmt.Print("\033[H\033[2J")
		fmt.Println("|-------------------------------------------------------|")
		fmt.Println("|  Ternet.                                              |")
		fmt.Println("|                                                       |")
		fmt.Println("|  A CLI based Micro Blogging Tool.                     |")
		fmt.Println("|  Developed in GOLang by Vincent Samuel Paul.          |")
		fmt.Println("|                                                       |")
		fmt.Println("|  Use the below flags.                                 |")
		fmt.Println("|    -help -> to view all the functions available.      |")
		fmt.Println("|    -login -> login to browse, post and like posts.    |")
		fmt.Println("|    -signup -> signup to create a new user.            |")
		fmt.Println("|    -browse -> browse posts from users.                |")
		fmt.Println("|      while browsing posts, enter:                     |")
		fmt.Println("|        'l' to like the post.                          |")
		fmt.Println("|        'x' to exit browsing.                          |")
		fmt.Println("|    -newpost -> create a new post.                     |")
		fmt.Println("|    -logout -> Logout.                                 |")
		fmt.Println("|                                                       |")
		fmt.Println("|  Enjoy using the tool.                                |")
		fmt.Println("|                                                       |")
		fmt.Println("|  Produced by © Vincent Samuel Paul 2025.              |")
		fmt.Println("|-------------------------------------------------------|")

	case *login:
		_, err := os.ReadFile("state.json")
		if err != nil {
			res, username := loginUser()
			if !res {
				fmt.Println("Invalid Credentials.")
			} else {
				state := map[string]string{"username":username}
				stateBytes, err := json.Marshal(state)
				if err != nil {
					log.Println(err)
				}
				err = os.WriteFile("state.json", stateBytes, 0644)
				if err != nil {
					log.Println(err)
				}
				fmt.Println("Logged IN.")
			}
		} else {
			fmt.Println("Logged IN.")
		}

	case *logout:
		err := os.Remove("state.json")
		if err != nil {
			log.Println(err)
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println("----Thank You----")

	case *signup:
		_, err := os.ReadFile("state.json")
		if err != nil {
			res, username := createUser()
			if !res {
				fmt.Println("User already exists.")
			} else {
				state := map[string]string{"username":username}
				stateBytes, err := json.Marshal(state)
				if err != nil {
					log.Println(err)
				}
				err = os.WriteFile("state.json", stateBytes, 0644)
				if err != nil {
					log.Println(err)
				}
				fmt.Println("User created.")
			}
		} else {
			fmt.Println("You're already logged in.")
		}

	case *newpost:
		out, err := os.ReadFile("state.json")
		if err != nil {
			fmt.Println("Login first.")
		} else {
			var user Status
			err := json.Unmarshal(out, &user)
			if err != nil {
				log.Println(err)
			}
			res := createPost(user.Username)
			if res {
				fmt.Println("New post created.")
			} else {
				fmt.Println("Error creating post.")
			}
		}

	case *posts:
		_, err := os.ReadFile("state.json")
		if err != nil {
			fmt.Println("Login first.")
		} else {
			res := getPosts()
			reader := bufio.NewReader(os.Stdin)
			for i, j := range res {
				fmt.Print("\033[H\033[2J")
				fmt.Println("----------")
				fmt.Println("Username: :", j.Username)
				fmt.Println(j.Data)
				fmt.Println("Likes: :", j.Likes)
				fmt.Println("----------")
				char, _, err := reader.ReadRune()
				if err != nil {
					log.Println(err)
				}
				if char == 'x' {
					fmt.Print("\033[H\033[2J")
					fmt.Println("----Thank You----")
					break
				} else if char == 'l' {
					fmt.Print("\033[H\033[2J")
					likePost(j.Username, j.Data)
					fmt.Println("----------")
					fmt.Println("Username: :", j.Username)
					fmt.Println(j.Data)
					fmt.Println("Likes: :", j.Likes+1)
					fmt.Println("----------")
				}
				if i == len(res)-1 {
					fmt.Print("\033[H\033[2J")
					fmt.Println("You've scrolled too far.")
					break
				}
			}
		}
	}

}