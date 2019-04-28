package infrastructure

import (
  "log"
  )

var Router *gin.Engine

func init() {
  mstItemController := controllers.NewMstItemController(NewMySQLHandler())

  router := gin.Default()
  authorized := router.Group("/v1/")
  {
    authorized.GET("/mstItems/", func(c *gin.Context) { mstItemController.Index(c) }
    authorized.POST("/mstItems/", func(c *gin.Context) { mstItemController.Create(c) }
  }
  log.Print("router setup complete") 
  Router = router
}
