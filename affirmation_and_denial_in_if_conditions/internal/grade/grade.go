package grade

import "fmt"

// CalculateGrade determines the letter grade based on a numeric score
func CalculateGrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score >= 50 && score < 60:
		return "E"
	default:
		return "F"
	}
}

// DisplayResult shows the appropriate message based on the grade
func DisplayResult(score int) string {
	grade := CalculateGrade(score)

	if grade == "F" {
		return "You have failed the exam"
	}

	return fmt.Sprintf("You have passed with %s grade", grade)
}

// AdminGradeStats returns statistics that only admins should see
func AdminGradeStats(score int) string {
	grade := CalculateGrade(score)

	// Admin can see additional information
	var percentile string
	switch grade {
	case "A":
		percentile = "Top 10%"
	case "B":
		percentile = "Top 25%"
	case "C":
		percentile = "Average"
	case "D", "E":
		percentile = "Below Average"
	default:
		percentile = "Lowest 10%"
	}

	return fmt.Sprintf("Grade: %s, Percentile: %s, Raw Score: %d/100",
		grade, percentile, score)
}
