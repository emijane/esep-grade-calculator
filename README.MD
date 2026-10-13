# ESEP Grade Calculator

A Go library for computing a final letter grade from assignments, exams, and
essays. It averages each category and applies weights of 50%, 35%, and 15%.
Letter cutoffs are A >= 90, B >= 80, C >= 70, D >= 60, and F < 60.
Pass/Fail evaluation uses C or better as the passing threshold.

## Setup

Install Go 1.24.4 or newer, then run:

```sh
git clone https://github.com/emijane/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```

## Tests and coverage

```sh
go test ./...
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
```

## Usage

Create a calculator with `NewGradeCalculator()` for letter grades or
`NewGradeCalculator(PassFail)` for Pass/Fail. Add each score with
`AddGrade(name, score, gradeType)`, using `Assignment`, `Exam`, or `Essay`.
Call `GetFinalGrade()` to retrieve the result.
