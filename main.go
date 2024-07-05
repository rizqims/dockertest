package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	g := gin.Default()

	g.GET("/", welcomeHandler)
	g.POST("/", welcomeJalmaHandler)

	g.Run(":8080")

}

func welcomeHandler(c *gin.Context){
	ucapan := "selamat pagi bro bro!"
	c.JSON(http.StatusOK, ucapan)
}

type ucapan struct{
	Nama string `json:"nama"`
	Hari string `json:"hari"`
	Waktu string `json:"waktu"`
}

func welcomeJalmaHandler(c *gin.Context) {
	newData := ucapan{}
	err := c.ShouldBindBodyWithJSON(&newData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error JSON"})
		return
	}

	ucapanRes := fmt.Sprintf("Selamat %v %v! semoga di hari %v permasalahanmu bisa selesai!", newData.Waktu, newData.Nama, newData.Hari)
	c.JSON(http.StatusOK, ucapanRes)
}