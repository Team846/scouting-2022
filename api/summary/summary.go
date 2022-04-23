package summary

import (
	"sort"

	"github.com/team846/scouting-2022/api/db"
)

type TeamMatchStatSummary struct {
	AutoPoints       float64 `json:"autoPoints"`
	TeleopPoints     float64 `json:"teleopPoints"`
	ClimbPoints      float64 `json:"climbPoints"`
	TotalCargoPoints float64 `json:"totalCargoPoints"`
	TotalPoints      float64 `json:"totalPoints"`
	MatchesPlayed    int     `json:"matchesPlayed"`
}

func (a *TeamMatchStatSummary) add(b TeamMatchStatSummary) {
	a.AutoPoints += b.AutoPoints
	a.TeleopPoints += b.TeleopPoints
	a.ClimbPoints += b.ClimbPoints
	a.TotalCargoPoints += b.TotalCargoPoints
	a.TotalPoints += b.TotalPoints
}

func (a *TeamMatchStatSummary) mul(b float64) {
	a.AutoPoints *= b
	a.TeleopPoints *= b
	a.ClimbPoints *= b
	a.TotalCargoPoints *= b
	a.TotalPoints *= b
}

func (a *TeamMatchStatSummary) div(b float64) {
	a.mul(1.0 / b)
}

// TODO add drop n matches, bias recent n matches
func Stats(stats []db.TeamMatchStat, dropCount int) TeamMatchStatSummary {
	if len(stats) == 0 {
		return TeamMatchStatSummary{}
	}

	summaries := make([]TeamMatchStatSummary, len(stats))

	for i, s := range stats {
		// Auto cargo
		summaries[i].AutoPoints += float64(2 * s.AutoCargoLow)
		summaries[i].AutoPoints += float64(4 * s.AutoCargoHigh)

		// Teleop cargo
		summaries[i].TeleopPoints += float64(1 * s.TeleopCargoLow)
		summaries[i].TeleopPoints += float64(2 * s.TeleopCargoHigh)

		// Total cargo
		summaries[i].TotalCargoPoints = summaries[i].AutoPoints + summaries[i].TeleopPoints

		// Auto taxi points
		if s.Taxi {
			summaries[i].AutoPoints += 2
		}

		// Climb points
		climbPointValues := [5]float64{0, 4, 6, 10, 15}
		summaries[i].ClimbPoints = climbPointValues[s.ClimbLevel]

		// Total points
		summaries[i].TotalPoints = summaries[i].AutoPoints + summaries[i].TeleopPoints + summaries[i].ClimbPoints
	}

	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].TotalPoints < summaries[j].TotalPoints
	})
	summaries = summaries[dropCount:]

	summary := TeamMatchStatSummary{}
	for _, s := range summaries {
		summary.add(s)
	}

	summary.div(float64(len(summaries)))
	summary.MatchesPlayed = len(summaries) + dropCount

	return summary
}
