package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbalrestu/perpustakaan-api/controllers"
)

func main() {
	router := gin.Default()
	var buku controllers.BukuController
	router.GET("/bukuGet", buku.Get)
	router.POST("/bukuPOST", buku.Post)
	router.Run()
}
