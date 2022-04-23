package db

import "time"

type Team struct {
	TeamNumber int    `json:"teamNumber"`
	Nickname   string `json:"nickname"`
}

const (
	MatchTypePractice = 0
	MatchTypeQual     = 1
)

const (
	ClimbLevelNone      = 0
	ClimbLevelLow       = 1
	ClimbLevelMid       = 2
	ClimbLevelHigh      = 3
	ClimbLevelTraversal = 4
)

type TeamMatchStat struct {
	ID              int       `json:"id"`
	TeamNumber      int       `json:"teamNumber"`
	MatchType       int       `json:"matchType"`
	MatchNumber     int       `json:"matchNumber"`
	ScoutName       string    `json:"scoutName"`
	SubmitDatetime  time.Time `json:"submitDatetime"`
	Taxi            bool      `json:"taxi"`
	AutoCargoLow    int       `json:"autoCargoLow"`
	AutoCargoHigh   int       `json:"autoCargoHigh"`
	TeleopCargoLow  int       `json:"teleopCargoLow"`
	TeleopCargoHigh int       `json:"teleopCargoHigh"`
	ClimbLevel      int       `json:"climbLevel"`
	PlayedDefense   bool      `json:"playedDefense"`
	Comments        string    `json:"comments"`
}
