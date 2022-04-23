CREATE TABLE team_match_stats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_number INTEGER NOT NULL REFERENCES teams(team_number),
    --
    -- Match info
    --
    -- 0=practice match, 1=qual match
    match_type INTEGER NOT NULL,
    match_number INTEGER NOT NULL,
    --
    -- Scout
    --
    scout_name TEXT NOT NULL,
    submit_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    --
    -- Match data
    --
    -- 0/1 = no/yes
    taxi INTEGER NOT NULL,
    auto_cargo_low INTEGER NOT NULL,
    auto_cargo_high INTEGER NOT NULL,
    teleop_cargo_low INTEGER NOT NULL,
    teleop_cargo_high INTEGER NOT NULL,
    -- 0/1/2/3/4 = none/low/mid/high/traversal
    climb_level INTEGER NOT NULL,
    -- 0/1 = no/yes
    played_defense INTEGER NOT NULL,
    comments TEXT
);