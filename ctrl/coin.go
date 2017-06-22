package ctrl

import "github.com/gin-gonic/gin"

func GetCoinCount(c *gin.Context) {
	c.JSON(200, gin.H{
		"c025":0,
		"c050":0,
		"c1":0,
		"c5":10,
		"c10":10,
	})
}

func SetCoinCount(c *gin.Context) {

}

func EmptyCoin(c *gin.Context) {

}

func PayoutCoin(c *gin.Context) {

}