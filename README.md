# Scouting 2022

Scouting data viewer for champs!

```
# Build frontend
cd app && npm build && cd ..

# Run local server
go run .

# Import teams into database
sqlite3 db.sqlite3 < import_teams.sql

# Import scouting data into database (scouting_data.txt exported from Google sheets)
python3 import_sheets.py
```

Then view dashboard on localhost:8080
