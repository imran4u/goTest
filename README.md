# goTest

## Using [Go gin library](https://github.com/gin-gonic/gin) 

    1. Run command
    go get -u github.com/gin-gonic/gin

    2. sample router

        package main

        import (
        "net/http"

        "github.com/gin-gonic/gin"
        )

        func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
            "message": "pong",
            })
        })
        r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
        }

    3. Read josn file and give response [ref](https://tutorialedge.net/golang/parsing-json-with-golang)
      
        ```go
        jsonFile, err := os.Open("users.json")
        // if we os.Open returns an error then handle it
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("Successfully Opened users.json")
        ```
