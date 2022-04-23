import { Paper, TableBody, TableContainer, TableHead, Table, TableCell, TableRow } from "@mui/material";
import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";

export default function TeamsList() {
    const [teams, setTeams] = useState([]);

    // Load teams on mount
    useEffect(() => {
        const queryParams = new URLSearchParams(window.location.search);
        const sortBy = queryParams.get('sortBy');

        (async () => {
            const data = await getTeams();

            await Promise.all(data.map(async (d) => {
                const summary = await getTeamSummaryStats(d.teamNumber);
                d.summary = summary;
            }));

            switch (sortBy) {
                case "autoPoints":
                    data.sort((a, b) => b.summary.autoPoints - a.summary.autoPoints);
                    break;
                case "teleopPoints":
                    data.sort((a, b) => b.summary.teleopPoints - a.summary.teleopPoints);
                    break;
                case "climbPoints":
                    data.sort((a, b) => b.summary.climbPoints - a.summary.climbPoints);
                    break;
                case "totalCargoPoints":
                    data.sort((a, b) => b.summary.totalCargoPoints - a.summary.totalCargoPoints);
                    break;
                case "totalPoints":
                    data.sort((a, b) => b.summary.totalPoints - a.summary.totalPoints);
                    break;
                default:
            }

            setTeams(data);
        })();
    }, []);

    return <>
        <b>Sort by:</b>
        <div id="sortByLinks">
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }`}>Team</a>
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }?sortBy=autoPoints`}>Auto Points</a>
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }?sortBy=teleopPoints`}>Teleop Points</a>
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }?sortBy=climbPoints`}>Climb Points</a>
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }?sortBy=totalCargoPoints`}>Total Cargo Points</a>
            <a href={`${window.location.protocol + '//' + window.location.host + window.location.pathname
                }?sortBy=totalPoints`}>Total Points</a>
        </div>
        <TableContainer component={Paper}>
            <Table size="small">
                <TableHead>
                    <TableCell />
                    <TableCell>Team</TableCell>
                    <TableCell>Name</TableCell>
                    <TableCell>Matches Played</TableCell>
                    <TableCell>Auto Points</TableCell>
                    <TableCell>Teleop Points</TableCell>
                    <TableCell>Climb Points</TableCell>
                    <TableCell>Total Cargo Points</TableCell>
                    <TableCell>Total Points</TableCell>
                </TableHead>
                <TableBody>{
                    teams.map((t, i) =>
                    (
                        <TableRow className={t.teamNumber == 846 ? "team846" : ""}>
                            <TableCell>{i + 1}</TableCell>
                            <TableCell component="th" scope="row">
                                <Link to={`/team/${t.teamNumber}`}>{t.teamNumber}</Link>
                            </TableCell>

                            <TableCell>{t.nickname}</TableCell>
                            <TableCell>{t.summary.matchesPlayed}</TableCell>
                            <TableCell>{formatDouble(t.summary.autoPoints)}</TableCell>
                            <TableCell>{formatDouble(t.summary.teleopPoints)}</TableCell>
                            <TableCell>{formatDouble(t.summary.climbPoints)}</TableCell>
                            <TableCell>{formatDouble(t.summary.totalCargoPoints)}</TableCell>
                            <TableCell>{formatDouble(t.summary.totalPoints)}</TableCell>
                        </TableRow>
                    )
                    )}
                </TableBody>
            </Table>
        </TableContainer>
    </>
}

async function getTeams() {
    return await fetch("/api/teams").then(res => {
        if (res.ok) {
            return res.json();
        } else {
            throw Error();
        }
    });
}

async function getTeamSummaryStats(teamNumber) {
    return await fetch(`/api/stats/summary/${teamNumber}`)
        .then(res => {
            if (res.ok) {
                return res.json();
            } else {
                throw Error();
            }
        })
}

function formatDouble(n) {
    return (Math.round(n * 100) / 100).toFixed(2);
}