# Command Line Quiz

Small Go command-line quiz that reads questions from a CSV file with a time limit.

## Run

```
go run main.go
```

## Options

-csv string
CSV file with problems (default: problems.csv)

-limit int
time limit in seconds (default: 30)

Example:

```
go run main.go -csv=problems.csv -limit=20
```
