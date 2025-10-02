package esepunittests

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

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

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	n := gc.calculateNumericalGrade()
	switch {
	case n >= 90:
		return "A"
	case n >= 80:
		return "B"
	case n >= 70:
		return "C"
	case n >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{Name: name, Grade: grade, Type: Assignment})
	case Exam:
		gc.exams = append(gc.exams, Grade{Name: name, Grade: grade, Type: Exam})
	case Essay:
		gc.essays = append(gc.essays, Grade{Name: name, Grade: grade, Type: Essay})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	const wa, we, ws = 0.50, 0.35, 0.15
	a := float64(computeAverage(gc.assignments))
	e := float64(computeAverage(gc.exams))
	s := float64(computeAverage(gc.essays))
	total := wa*a + we*e + ws*s
	return int(total)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for i := range grades {
		sum += grades[i].Grade
	}
	return sum / len(grades)
}
