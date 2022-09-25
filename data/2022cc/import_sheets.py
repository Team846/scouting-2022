from enum import Enum
import sqlite3
import copy

DATA_FILE_NAME = "scouting_data.txt"
DB_FILE_NAME = "db.sqlite3"

TEAMS = [
    114,
    254,
    359,
    498,
    604,
    649,
    694,
    696,
    846,
    971,
    972,
    973,
    1619,
    1678,
    1690,
    1700,
    2046,
    2486,
    2813,
    2910,
    2930,
    3175,
    3256,
    3310,
    3476,
    3478,
    3647,
    3940,
    4255,
    4414,
    4499,
    5104,
    5507,
    5940,
    6036,
    6800,
    7034,
    7157,
    8033,
]

MATCHES = {
    1: [6036, 7157, 8033, 3647, 254, 604],
    2: [1690, 114, 1678, 2910, 1700, 3310],
    3: [6800, 1619, 4499, 3478, 694, 972],
    4: [5507, 5104, 2046, 971, 498, 5940],
    5: [4414, 359, 7034, 3175, 649, 3940],
    6: [3256, 2486, 2813, 696, 973, 846],
    7: [2930, 1678, 8033, 4255, 3476, 972],
    8: [5940, 3478, 7034, 649, 359, 3647],
    9: [498, 4414, 3310, 4499, 694, 3256],
    10: [2046, 973, 6036, 1700, 696, 1690],
    11: [3940, 2813, 604, 114, 2930, 971],
    12: [3476, 7157, 3175, 6800, 2910, 846],
    13: [5104, 4255, 254, 1619, 2486, 5507],
    14: [114, 8033, 694, 696, 3478, 6036],
    15: [2910, 7157, 604, 4499, 1700, 4414],
    16: [846, 2930, 3310, 359, 973, 5104],
    17: [7034, 649, 3476, 498, 1690, 5507],
    18: [254, 6800, 5940, 3256, 3940, 2046],
    19: [971, 2813, 3175, 1678, 3647, 2486],
    20: [1619, 972, 359, 4255, 604, 1700],
    21: [694, 3310, 5507, 114, 649, 2046],
    22: [2486, 6036, 6800, 846, 1678, 4414],
    23: [3478, 973, 498, 254, 972, 3175],
    24: [4255, 2930, 1619, 7157, 971, 7034],
    25: [4499, 2910, 8033, 5940, 2813, 3647],
    26: [3476, 3940, 1690, 5104, 3256, 696],
    27: [3175, 114, 498, 2486, 7157, 4255],
    28: [5507, 254, 846, 972, 2930, 1700],
    29: [696, 604, 649, 6800, 3647, 5104],
    30: [2046, 4499, 3310, 2813, 8033, 1619],
    31: [971, 3478, 1678, 1690, 3256, 4414],
    32: [694, 6036, 359, 2910, 973, 3476],
    33: [3940, 5940, 114, 7034, 3175, 1619],
    34: [5104, 971, 4414, 3478, 254, 2813],
    35: [696, 972, 7157, 3256, 359, 3310],
    36: [973, 4255, 649, 6036, 4499, 5940],
    37: [1700, 3940, 6800, 1690, 8033, 2486],
    38: [2046, 3476, 3647, 604, 2930, 5507],
    39: [694, 1678, 2910, 498, 846, 7034],
    40: [3175, 1700, 6036, 3310, 4255, 8033],
    41: [3647, 972, 3256, 2486, 2046, 359],
    42: [7157, 3940, 1619, 2813, 649, 1678],
    43: [6800, 2930, 973, 3478, 5507, 2910],
    44: [604, 971, 1690, 694, 846, 5940],
    45: [7034, 254, 114, 3476, 5104, 4499],
    46: [696, 498, 2930, 4414, 1619, 6036],
    47: [846, 4255, 3647, 694, 973, 3940],
    48: [4499, 5507, 2813, 3175, 1690, 359],
    49: [649, 4414, 2910, 972, 971, 8033],
    50: [3256, 498, 1678, 604, 5104, 114],
    51: [254, 1700, 2046, 6800, 7034, 696],
    52: [3476, 3478, 2486, 3310, 5940, 7157],
    53: [5104, 694, 2930, 3175, 3256, 8033],
    54: [1690, 3647, 2910, 1619, 649, 254],
    55: [1700, 498, 359, 6800, 971, 3310],
    56: [696, 5507, 5940, 1678, 4499, 7157],
    57: [846, 3940, 2486, 3478, 604, 2046],
    58: [4414, 972, 2813, 4255, 7034, 6036],
    59: [973, 4499, 971, 114, 3476, 6800],
    60: [5940, 359, 1678, 254, 694, 498],
    61: [3175, 3478, 4255, 696, 2910, 2046],
    62: [1619, 846, 3256, 6036, 5507, 3940],
    63: [604, 8033, 4414, 973, 3310, 7034],
    64: [1700, 649, 2486, 2813, 2930, 3476]
}

matches_missing = copy.deepcopy(MATCHES)


class Match:
    def __init__(self, data_row):
        if len(data_row) < 9:
            raise Exception("bad row")

        self.match_number = int(data_row[1])
        self.team_number = int(data_row[2])
        self.scout_name = data_row[3]
        self.cargo_auto_high = int(data_row[4])
        self.cargo_auto_low = int(data_row[5])
        self.cargo_teleop_low = int(data_row[6])
        self.cargo_teleop_high = int(data_row[7])

        raw_climber = data_row[8].strip()
        if raw_climber == "Did not climb":
            self.climber = 0
        elif raw_climber == "Low Rung":
            self.climber = 1
        elif raw_climber == "Mid Rung":
            self.climber = 2
        elif raw_climber == "High Rung":
            self.climber = 3
        elif raw_climber == "Traversal Rung":
            self.climber = 4
        else:
            raise Exception("bad climber value")

        if self.team_number not in MATCHES[self.match_number]:
            print(
                f"team {self.team_number} is not supposed to be in match {self.match_number}")
        elif self.team_number not in matches_missing[self.match_number]:
            print(
                f"duplicate form for team {self.team_number} in match {self.match_number}")
        else:
            matches_missing[self.match_number].remove(self.team_number)


def main():
    data = {}
    con = sqlite3.connect(DB_FILE_NAME)

    con.execute("DELETE FROM team_match_stats;")
    con.commit()

    with open(DATA_FILE_NAME) as data_file:
        for line in data_file.readlines():
            match = Match(line.split('\t'))
            if data.get(match.team_number) == None:
                data[match.team_number] = []
            data[match.team_number].append(match)

    for d in data.values():
        for m in d:
            statement = f"""
            INSERT INTO team_match_stats (
                team_number, match_type, match_number,
                scout_name, submit_datetime,
                taxi, auto_cargo_low, auto_cargo_high, teleop_cargo_low, teleop_cargo_high, climb_level,
                played_defense, comments)
            VALUES (
                {m.team_number}, 1, {m.match_number},
                \"{m.scout_name}\", 0,
                0, {m.cargo_auto_low}, {m.cargo_auto_high}, {m.cargo_teleop_low}, {m.cargo_teleop_high}, {m.climber},
                0, \"na\"
            );
            """

            con.execute(statement)
            con.commit()
            # print(statement)

            if int(m.team_number) not in TEAMS:
                print("Err!", m.team_number)

    for match, teams in matches_missing.items():
        for team in teams:
            print(f"missing team {team} from match {match}")


if __name__ == "__main__":
    main()
