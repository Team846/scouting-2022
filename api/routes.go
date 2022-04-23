package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team846/scouting-2022/api/db"
	"github.com/team846/scouting-2022/api/summary"
)

type Env struct {
	Db             *sql.DB
	MatchDropCount int
}

// getTeams responds with a list of all the teams.
func (e *Env) GetTeams(c *gin.Context) {
	teams, err := db.GetTeams(e.Db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, teams)
}

// postTeams adds a team.
func (e *Env) PostTeams(c *gin.Context) {
	var newTeam db.Team

	if err := c.BindJSON(&newTeam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if err := db.NewTeam(e.Db, newTeam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, newTeam)
}

func (e *Env) GetTeamByNumber(c *gin.Context) {
	teamNumber, err := strconv.Atoi(c.Param("team_number"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	team, err := db.GetTeamByNumber(e.Db, teamNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (e *Env) GetTeamMatchStatsByTeam(c *gin.Context) {
	teamNumber, err := strconv.Atoi(c.Param("team_number"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	stats, err := db.GetTeamMatchStatsByTeam(e.Db, teamNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (e *Env) GetTeamSummaryStats(c *gin.Context) {
	teamNumber, err := strconv.Atoi(c.Param("team_number"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	stats, err := db.GetTeamMatchStatsByTeam(e.Db, teamNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, summary.Stats(stats, e.MatchDropCount))
}
