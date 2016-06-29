package main

import (
	"github.com/PiedPiperServer/conf"
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/hi_json", func(c *iris.Context) {

		c.JSON(iris.StatusOK, iris.Map{
			"Name":  "Iris",
			"Born":  "13 March 2016",
			"Stars": 3693,
		})
	})
	iris.Listen(":" + conf.ServiceConfig.HttpPort)
}
