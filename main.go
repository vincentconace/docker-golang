package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	taskpkg "github.com/vincentconace/golang-docker/task"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	r := gin.Default()

	r.GET("/task", func(c *gin.Context) {
		task := &taskpkg.Task{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Text:      "Hello",
			Completed: false,
		}

		taskresult, _ := taskpkg.SaveTask(task)
		fmt.Println("creado exitosamente")
		c.JSON(http.StatusOK, taskresult)
	})

	r.Run()
}
