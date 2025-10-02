# ESEP Grade Calculator

This library is a simple utility to calculate your grades. It computes the averages for each category (assignments, tests, essays). Then it applies the weights, and converts it to a letter grade.

Runtime requirements:
Golang 1.24+ (Must be installed)

Setup:
git clone https://github.com/<your-username>/esep-grade-calculator
cd esep-grade-calculator/grade-calculator

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

Usage: 
1. Calculate grades for assignments, essays, and exams.
2. Calculate weights for grades, for GPA.
3. Use the calculator to calculate the final letter grade.

Project structure:
grade_calculator.go – main code
grade_calculator_test.go – tests
go.mod – Go module file
README.md