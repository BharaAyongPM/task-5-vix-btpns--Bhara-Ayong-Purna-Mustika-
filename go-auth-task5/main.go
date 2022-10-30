package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

var currentImage *imageupload.Image

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/image", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		currentImage.Write(c.Writer)
	})

	r.GET("/thumbnail", func(c *gin.Context) {
		if currentImage == nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		t, err := imageupload.ThumbnailJPEG(currentImage, 300, 300, 80)

		if err != nil {
			panic(err)
		}

		t.Write(c.Writer)
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			panic(err)
		}

		currentImage = img

		c.Redirect(http.StatusMovedPermanently, "/")
	})
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
