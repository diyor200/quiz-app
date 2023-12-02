// main.go
package main

import (
	"awesomeProject/internal/controller"
	"github.com/gin-gonic/gin"
)

type quizData struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

var data = []quizData{
	{
		Question: "What is the capital of France?",
		Options:  []string{"Paris", "London", "Berlin", "Madrid"},
		Answer:   0,
	},
	{
		Question: "one",
		Options:  []string{"one", "London", "Berlin", "Madrid"},
		Answer:   0,
	},
	{
		Question: "two",
		Options:  []string{"two", "London", "Berlin", "Madrid"},
		Answer:   0,
	},
	{
		Question: "three",
		Options:  []string{"three", "London", "Berlin", "Madrid"},
		Answer:   0,
	},
}

func main() {
	ctrl := controller.NewController()
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/try-again", ctrl.TryAgain)
	router.GET("/questions", ctrl.Questions)
	router.POST("/submit", ctrl.Submit)
	router.GET("/result", ctrl.Result)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
