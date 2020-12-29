package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func main() {
	runtime.GOMAXPROCS(300) // Ограничение 30 горутин
	err := connect()        // Подключение к postgres
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	router = gin.Default() // Используем http соединение
	router.LoadHTMLFiles(cfg.HTML + "index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})
	go router.POST("/Add", addValues)
	go router.POST("/Rmv", deleteValues)
	go router.POST("/Cng", changeValues)
	go router.POST("/Top ", topUsers)
	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

// Функция пост запроса на добавление пользователя
func addValues(c *gin.Context) {
	var err error
	Name := c.PostForm("name")
	Group, err := strconv.Atoi(c.PostForm("group"))
	if err != nil {
		fmt.Println(err.Error())
	}
	Score, err := strconv.Atoi(c.PostForm("score"))
	if err != nil {
		fmt.Println(err.Error())
	}
	err = queryAdd(Name, Group, Score)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.HTML(200, "index.html", gin.H{
		"title": "Home Page",
	})
}

// На удаление пользователя
func deleteValues(c *gin.Context) {
	ID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	err = queryRmv(ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "index.html", gin.H{
		"title": "Home Page",
	})
}

// На изменение очков пользователя
func changeValues(c *gin.Context) {
	ID, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	Score, err := strconv.Atoi(c.PostForm("score"))
	err = queryCng(Score, ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.HTML(200, "index.html", gin.H{
		"title": "Home Page",
	})
}

//Вывод топ x пользователей
func topUsers(c *gin.Context) {
	x, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		fmt.Println(err.Error())
	}
	Group, err := strconv.Atoi(c.PostForm("score"))
	a, err := queryRow(Group)

	a = queryTop(a, x)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < len(a); i++ {

		rows, err := db.Query(`select * from users WHERE "score" = ($1)`, a[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		c.HTML(200, "topUsers.html", gin.H{
			"Array": rows,
		})
	}
}
