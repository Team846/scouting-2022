package db

import (
	"database/sql"
)

func GetTeams(db *sql.DB) ([]Team, error) {
	rows, err := db.Query("SELECT * FROM teams ORDER BY team_number ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []Team{}
	for rows.Next() {
		var team Team
		if err := rows.Scan(&team.TeamNumber, &team.Nickname); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func NewTeam(db *sql.DB, team Team) error {
	result, err := db.Exec("INSERT INTO teams VALUES (?, ?)", team.TeamNumber, team.Nickname)
	if err != nil {
		return err
	}

	if _, err := result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func GetTeamByNumber(db *sql.DB, teamNumber int) (*Team, error) {
	var team Team

	row := db.QueryRow("SELECT * FROM teams WHERE team_number = ?", teamNumber)
	if err := row.Scan(&team.TeamNumber, &team.Nickname); err != nil {
		return nil, err
	}
	return &team, nil
}

func GetTeamMatchStatsByTeam(db *sql.DB, teamNumber int) ([]TeamMatchStat, error) {
	rows, err := db.Query("SELECT * FROM team_match_stats WHERE team_number = ? ORDER BY match_type ASC, match_number ASC", teamNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := []TeamMatchStat{}
	for rows.Next() {
		var stat TeamMatchStat
		if err := rows.Scan(
			&stat.ID,
			&stat.TeamNumber,
			&stat.MatchType,
			&stat.MatchNumber,
			&stat.ScoutName,
			&stat.SubmitDatetime,
			&stat.Taxi,
			&stat.AutoCargoLow,
			&stat.AutoCargoHigh,
			&stat.TeleopCargoLow,
			&stat.TeleopCargoHigh,
			&stat.ClimbLevel,
			&stat.PlayedDefense,
			&stat.Comments,
		); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
