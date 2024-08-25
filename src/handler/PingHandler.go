package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"imran4u/goTest/src/model/response"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	filePath := "src/json/response/ping-response.json"
	// filePath := "src/json/response/response.json" // wrong file path

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Successfully Opened ${filePath}")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(string(byteValue))

	var pingResponse response.PingResponse

	marshingError := json.Unmarshal(byteValue, &pingResponse)
	if marshingError != nil {
		fmt.Println(marshingError)
		c.JSON(http.StatusInternalServerError, gin.H{"marshingError": marshingError.Error()})
		return
	}

	c.JSON(http.StatusOK, pingResponse)
}
