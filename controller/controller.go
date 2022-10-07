package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	UserCode string `json:"userCode"`
}

func GetCode(c *gin.Context) {
	fmt.Println("Got Code from Front End")
	data := Data{}
	res := c.Bind(&data)
	fmt.Println("User C0de:", data.UserCode)
	c.JSON(http.StatusOK, gin.H{"data": res})

	// a := Data{}
	// b, err := c.GetRawData()
	// if err != nil {
	// 	panic(err)
	// }
	// json.Unmarshal(b, &a)
	// userCode := a.Code
	// fmt.Println("User COde:", userCode)
}
