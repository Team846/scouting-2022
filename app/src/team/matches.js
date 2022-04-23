import React, { useEffect, useState } from "react";
import { useParams } from "react-router";
import { Link } from "react-router-dom";
import { Paper, TableBody, TableContainer, TableHead, Table, TableCell, TableRow, IconButton, Collapse } from "@mui/material";
import { KeyboardArrowDown, KeyboardArrowUp } from "@mui/icons-material"
import NotFound from '../404.js';

export default function TeamMatches() {
    const params = useParams();

    const [team, setTeam] = useState({});
    const [matchStats, setMatchStats] = useState([]);
    const [notFound, setNotFound] = useState(false);

    useEffect(() => {
        getTeamByNumber(params.teamNumber)
            .then(data => setTeam(data))
            .catch(_ => setNotFound(true));

        getTeamMatchStatsByTeam(params.teamNumber)
            .then(data => setMatchStats(data));
    }, [params.teamNumber]);

    if (notFound) {
        return <NotFound />;
    }

    return <>
        <Link to="/">{"<<"} All teams</Link>
        <h2>Team {team.teamNumber} - {team.nickname}</h2>
        <TableContainer component={Paper}>
            <Table size="small">
                <TableHead>
                    <TableCell />
                    <TableCell>Match</TableCell>
                    <TableCell>Auto Taxi</TableCell>
                    <TableCell>Auto Low Cargo</TableCell>
                    <TableCell>Auto High Cargo</TableCell>
                    <TableCell>Teleop Low Cargo</TableCell>
                    <TableCell>Teleop High Cargo</TableCell>
                    <TableCell>Climb</TableCell>
                    <TableCell>Played Defense</TableCell>
                </TableHead>
                <TableBody>
                    {matchStats.map(s => <TeamMatchesRow stat={s} />)}
                </TableBody>
            </Table>
        </TableContainer>
    </>
}

function TeamMatchesRow(props) {
    const { stat } = props;

    const [openDetails, setOpenDetails] = useState(false);

    return <>
        <TableRow sx={{ '& > *': { borderBottom: 'unset' } }}>
            <TableCell>
                <IconButton
                    aria-label="expand row"
                    size="small"
                    onClick={() => setOpenDetails(!openDetails)}
                >
                    {openDetails ? <KeyboardArrowUp /> : <KeyboardArrowDown />}
                </IconButton>
            </TableCell>
            <TableCell component="th" scope="row">
                {stat.matchType === 1 ? "Qual" : "Prac"}
                &nbsp;
                {stat.matchNumber}
            </TableCell>
            <TableCell>{stat.taxi ? "Yes" : "No"}</TableCell>
            <TableCell>{stat.autoCargoLow}</TableCell>
            <TableCell>{stat.autoCargoHigh}</TableCell>
            <TableCell>{stat.teleopCargoLow}</TableCell>
            <TableCell>{stat.teleopCargoHigh}</TableCell>
            <TableCell>
                {["None", "Low", "Mid", "High", "Traversal"][stat.climbLevel]}
            </TableCell>
            <TableCell>{stat.playedDefense ? "Yes" : "No"}</TableCell>
        </TableRow>

        <TableRow>
            <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={9}>
                <Collapse in={openDetails} timeout="auto" unmountOnExit>
                    <p>
                        <b>Scout: </b>
                        {stat.scoutName}
                    </p>
                    <p>
                        <b>Date: </b>
                        {stat.submitDatetime}
                    </p>
                    <p>
                        <b>Comments:<br /></b>
                        {stat.comments}
                    </p>
                </Collapse>
            </TableCell>
        </TableRow>
    </>;
}

async function getTeamByNumber(teamNumber) {
    return await fetch(`/api/team/${teamNumber}`)
        .then(res => {
            if (res.ok) {
                return res.json();
            } else {
                throw Error();
            }
        })
}


async function getTeamMatchStatsByTeam(teamNumber) {
    return await fetch(`/api/stats/${teamNumber}`)
        .then(res => {
            if (res.ok) {
                return res.json();
            } else {
                throw Error();
            }
        })
}