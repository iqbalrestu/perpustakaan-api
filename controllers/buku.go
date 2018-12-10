package controllers

import (
	"github.com/gin-gonic/gin"
)

type BukuController struct{}

func (BukuController) Get(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "get success"})
}

func (BukuController) Post(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "post success"})

}

func (BukuController) Put(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "put success"})

}

func (BukuController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "deleted success"})

}

func (BukuController) Patch(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "patch success"})

}

func (BukuController) Head(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "head success"})

}

func (BukuController) Option(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success", "data": "option success"})

}
