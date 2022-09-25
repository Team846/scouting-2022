package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/team846/scouting-2022/api"
	"github.com/team846/scouting-2022/api/db"
)

const eventName string = "2022cc"

const reactBuildDir string = "app/build"

func main() {
	db, err := db.Open(eventName)
	if err != nil {
		fmt.Println("database err: ", err)
	}

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(reactBuildDir, true)))
	router.NoRoute(func(c *gin.Context) {
		c.File(reactBuildDir + "/index.html")
	})

	env := &api.Env{Db: db, MatchDropCount: 0}
	apiGroup := router.Group("/api")
	apiGroup.GET("/teams", env.GetTeams)
	apiGroup.POST("/teams", env.PostTeams)
	apiGroup.GET("/team/:team_number", env.GetTeamByNumber)
	apiGroup.GET("/stats/:team_number", env.GetTeamMatchStatsByTeam)
	apiGroup.GET("/stats/summary/:team_number", env.GetTeamSummaryStats)
	router.Run(":8080")
}
