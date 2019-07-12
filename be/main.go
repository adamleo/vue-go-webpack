package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/prometheus/common/log"

	"fmt"
	"net/http"
)

type myForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

func axisFormChecker(f myForm) (bad_results []string) {

	if f.Name == "" {
		bad_results = append(bad_results, "name")
	} else {
		log.Info("name: ", f.Name)
	}

	if f.Description == "" {
		bad_results = append(bad_results, "description")
	} else {
		log.Info("description: ", f.Description)
	}

	log.Info("bad_results: ", bad_results)

	return
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{
		"name":        fakeForm.Name,
		"description": fakeForm.Description,
		"message":     "Project Created!",
	})
}

func axiosHandler(c *gin.Context) {

	var fakeForm myForm

	if err := c.ShouldBind(&fakeForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	bad_results := axisFormChecker(fakeForm)

	response := gin.H{}

	if len(bad_results) != 0 {
		for i := 0; i < len(bad_results); i++ {
			response[bad_results[i]] = "this field needs to be filled."
		}
		c.JSON(420, response)
	} else {
		response = gin.H{
			"name":        fakeForm.Name,
			"description": fakeForm.Description,
			"message":     "Project Created!",
		}
		c.JSON(200, response)
	}
}

func mountedHandler(c *gin.Context) {
	log.Info("Vue is mounted on the frontend.")
	c.String(200, "Vue is mounted on the frontend.")
	// c.JSON(http.StatusOK, gin.H{
	// 	"one": "ping",
	// 	"two": "pong",
	// })
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("be/views/*")
	// router.Use(static.Serve("/js", static.LocalFile("fe/resources/assets/js", true)))
	router.Use(static.Serve("/js", static.LocalFile("fe/public/js", true)))
	router.Use(static.Serve("/css", static.LocalFile("fe/resources/assets/css", true)))

	//router.LoadHTMLFiles("views/template1.html", "views/template2.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hi There!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("pong."))
	})

	router.GET("/mounted", mountedHandler)

	router.POST("/projects", formHandler)

	router.POST("/axios", axiosHandler)

	router.Run(":8080")

}

/*

curl -X GET 127.0.0.1:8080/


*/
