package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	en := gin.Default()
	en.GET("/Registation/:dept/:name", func(d *gin.Context) {
		deparment := d.Param("dept")
		name := d.Param("name")
		//fmt.Println(deparment,name)
		d.String(200, "Welcome to %s deparment %s", deparment, name)
	})
	en.Run(":8080")

}
