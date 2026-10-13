package esepunittests

type GradeCalculator struct {
	grades     []Grade
	evaluation Evaluation
}

type Evaluation int

const (
	Letter Evaluation = iota
	PassFail
)

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator(evaluations ...Evaluation) *GradeCalculator {
	evaluation := Letter
	if len(evaluations) > 0 {
		evaluation = evaluations[0]
	}
	return &GradeCalculator{grades: make([]Grade, 0), evaluation: evaluation}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()
	if gc.evaluation == PassFail {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{Name: name, Grade: grade, Type: gradeType})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.grades, Assignment)
	exam_average := computeAverage(gc.grades, Exam)
	essay_average := computeAverage(gc.grades, Essay)

	weighted_grade := assignment_average*.5 + exam_average*.35 + essay_average*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, gradeType GradeType) float64 {
	sum, count := 0, 0
	for _, grade := range grades {
		if grade.Type == gradeType {
			sum += grade.Grade
			count++
		}
	}
	return float64(sum) / float64(count)
}
