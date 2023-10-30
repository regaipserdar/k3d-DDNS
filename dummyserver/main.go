package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "time"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*")

    r.GET("/", func(c *gin.Context) {
        hostname, _ := os.Hostname()
        podIP := c.ClientIP()
        currentTime := time.Now().Format("2006-01-02 15:04:05")

        data := gin.H{
            "PodHostname":  hostname,
            "PodIPAddress": podIP,
            "PodCurrentTime": currentTime,
        }

        c.HTML(http.StatusOK, "index.html", data)
    })

    r.Run(":8080")
}
