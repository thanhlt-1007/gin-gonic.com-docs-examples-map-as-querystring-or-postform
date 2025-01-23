package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func postMapHandler(context *gin.Context) {
    ids := context.QueryMap("ids")
    names := context.PostFormMap("names")
    context.JSON(
        http.StatusOK,
        gin.H {
            "ids": ids,
            "names": names,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.POST("/map", postMapHandler)
    engine.Run()
}
