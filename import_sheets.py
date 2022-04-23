from enum import Enum
import sqlite3

DATA_FILE_NAME = "scouting_data.txt"
DB_FILE_NAME = "db.sqlite3"

TEAMS = [
    86,
    88,
    118,
    254,
    316,
    364,
    401,
    498,
    599,
    694,
    846,
    857,
    1023,
    1477,
    1506,
    1619,
    1629,
    1731,
    1732,
    1745,
    1756,
    1771,
    1816,
    1828,
    1868,
    1914,
    1918,
    2147,
    2168,
    2438,
    2481,
    2506,
    2556,
    2557,
    2586,
    2590,
    2601,
    2638,
    2687,
    2959,
    2960,
    2974,
    3166,
    3175,
    3256,
    3414,
    3459,
    3476,
    3504,
    3641,
    3937,
    3970,
    4135,
    4329,
    4381,
    4415,
    4786,
    5086,
    5089,
    5114,
    5422,
    5612,
    5705,
    5724,
    5809,
    5903,
    6413,
    6672,
    7541,
    8122,
    8153,
    8546,
    8573,
    8590,
    8841,
    8898,
]


class Match:
    def __init__(self, data_row):
        if len(data_row) < 9:
            raise Exception("bad row")

        self.match_number = int(data_row[1])
        self.team_number = int(data_row[2])
        self.scout_name = data_row[3]
        self.cargo_auto_low = int(data_row[4])
        self.cargo_auto_high = int(data_row[5])
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


if __name__ == "__main__":
    main()
