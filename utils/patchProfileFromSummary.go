package utils

import (
	"strings"

	"github.com/zakhaev26/recruitments/schemas"
)

func PatchProfileFromSummary(profile *schemas.Profile, summary schemas.Summary) {
	if summary.Name != "" {
		profile.Name = summary.Name
	}
	if summary.Phone != "" {
		profile.Phone = summary.Phone
	}
	if len(summary.Skills) > 0 {
		profile.Skills = strings.Join(summary.Skills, ", ")
	}
	if len(summary.Education) > 0 {
		education := make([]string, len(summary.Education))
		for i, edu := range summary.Education {
			education[i] = edu.Name + " (" + strings.Join(edu.Dates, ", ") + ")"
		}
		profile.Education = strings.Join(education, "; ")
	}
	if len(summary.Experience) > 0 {
		experience := make([]string, len(summary.Experience))
		for i, exp := range summary.Experience {
			experience[i] = exp.Title + " at " + exp.Organization
		}
		profile.Experience = strings.Join(experience, "; ")
	}
}
