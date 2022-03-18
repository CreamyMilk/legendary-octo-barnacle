package main

import (
	"github.com/CreamyMilk/legendary-octo-barnacle/engine"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/mngnt", engine.ManganatoHandler)
	r.GET("/search/:provider/:titleName", engine.SearchHandler)
	r.GET("/get-chapter/:provider/:chapterId", engine.GetChapterHandler)
	r.GET("/chapters-list/:provider/:titleId", engine.ChaptersListHandler)
}

func main() {
	r := gin.Default()
	Route(r)
	r.Run()
}
