package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// TryAgain new quiz
func (c *Controller) TryAgain(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/questions")
}

// Questions starts quiz
func (c *Controller) Questions(ctx *gin.Context) {
	var currentQuestionIndex = 0
	var result = 0
	if currentQuestionIndex < len(data) {
		question := data[currentQuestionIndex]
		ctx.HTML(http.StatusOK, "quiz.html", gin.H{"question": question})
	} else {
		ctx.Redirect(http.StatusFound, "/result")
	}
}

// Submit checks question result
func (c *Controller) Submit(ctx *gin.Context) {

	selectedOptionStr := ctx.PostForm("quiz")
	selectedOption, err := strconv.Atoi(selectedOptionStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid selected option"})
		return
	}

	correctAnswer := data[currentQuestionIndex].Answer

	if selectedOption == correctAnswer {
		result++
		ctx.HTML(http.StatusOK, "result.html", gin.H{"message": "Correct!"})
	} else {
		ctx.HTML(http.StatusOK, "result.html", gin.H{"message": "Incorrect. Try again!"})
	}

	currentQuestionIndex++
}

// Result gets result
func (c *Controller) Result(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "completed.html",
		gin.H{"message": fmt.Sprintf("Quiz completed. Result: %d/%d", result, len(data))})
}
