package appli

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Test() {
	fmt.Printf("Test")
}

func CallFizzBuzz(c *gin.Context) {
	int1, err1 := strconv.ParseFloat(c.Param("int1"), 64)
	int2, err2 := strconv.ParseFloat(c.Param("int2"), 64)
	limit, errLimit := strconv.ParseInt(c.Param("limit"), 10, 64)

	if err1 != nil {
		log.Printf("Error parsing first number")
	} else if err2 != nil {
		log.Printf("Error parsing second number")
	} else if errLimit != nil {
		log.Printf(errLimit.Error())

	}

	str1 := c.Param("str1")
	str2 := c.Param("str2")

	var results []string

	for j := int64(1); j <= limit; j++ {
		var num = float64(j)
		if math.Mod(num, int1) == 0 && math.Mod(num, int2) == 0 {
			results = append(results, str1+str2)
		} else if math.Mod(num, int1) == 0 {
			results = append(results, str1)
		} else if math.Mod(num, int2) == 0 {
			results = append(results, str2)
		} else {
			results = append(results, strconv.FormatInt(j, 10))
		}
	}

	c.IndentedJSON(http.StatusOK, results)

}
