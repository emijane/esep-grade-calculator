package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Errorf("Expected C; got %s", got)
	}
}

func TestGetGradeD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 60, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Errorf("Expected D; got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	for _, tc := range []struct {
		kind GradeType
		want string
	}{
		{Assignment, "assignment"}, {Exam, "exam"}, {Essay, "essay"},
	} {
		if got := tc.kind.String(); got != tc.want {
			t.Errorf("Expected %s; got %s", tc.want, got)
		}
	}
}

func TestPassFailEvaluation(t *testing.T) {
	for _, tc := range []struct {
		score int
		want  string
	}{
		{70, "Pass"}, {69, "Fail"},
	} {
		gc := NewGradeCalculator(PassFail)
		gc.AddGrade("assignment", tc.score, Assignment)
		gc.AddGrade("exam", tc.score, Exam)
		gc.AddGrade("essay", tc.score, Essay)
		if got := gc.GetFinalGrade(); got != tc.want {
			t.Errorf("Expected %s for %d; got %s", tc.want, tc.score, got)
		}
	}
}
