package validator

import (
	"encoding/json"
)

type SectionCompleteness struct {
	Section    string  `json:"section"`
	Percentage float64 `json:"percentage"`
	Filled     int     `json:"filled"`
	Total      int     `json:"total"`
}

type ProfileCompleteness struct {
	Overall  float64               `json:"overall"`
	Sections []SectionCompleteness `json:"sections"`
}

var requiredFields = map[string][]string{
	"mission": {
		"mission_statement",
		"impact_areas",
	},
	"assets": {
		"assets",
	},
	"adversaries": {
		"adversaries",
	},
	"threats": {
		"threats",
	},
	"risks": {
		"risks",
	},
	"mitigations": {
		"mitigations",
	},
}

func CalculateSectionCompleteness(section string, data *string) SectionCompleteness {
	fields, ok := requiredFields[section]
	if !ok {
		return SectionCompleteness{Section: section, Percentage: 0, Filled: 0, Total: 0}
	}

	total := len(fields)
	if data == nil || *data == "" {
		return SectionCompleteness{
			Section:    section,
			Percentage: 0,
			Filled:     0,
			Total:      total,
		}
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(*data), &jsonData); err != nil {
		return SectionCompleteness{
			Section:    section,
			Percentage: 0,
			Filled:     0,
			Total:      total,
		}
	}

	filled := 0
	for _, field := range fields {
		if isFieldFilled(jsonData, field) {
			filled++
		}
	}

	percentage := 0.0
	if total > 0 {
		percentage = float64(filled) / float64(total) * 100
	}

	return SectionCompleteness{
		Section:    section,
		Percentage: percentage,
		Filled:     filled,
		Total:      total,
	}
}

func isFieldFilled(data map[string]interface{}, field string) bool {
	val, ok := data[field]
	if !ok {
		return false
	}

	switch v := val.(type) {
	case nil:
		return false
	case string:
		return v != ""
	case []interface{}:
		return len(v) > 0
	case map[string]interface{}:
		return len(v) > 0
	default:
		return true
	}
}

func CalculateProfileCompleteness(sections map[string]*string) ProfileCompleteness {
	sectionNames := []string{"mission", "assets", "adversaries", "threats", "risks", "mitigations"}

	var sectionResults []SectionCompleteness
	totalPercentage := 0.0

	for _, name := range sectionNames {
		data := sections[name]
		result := CalculateSectionCompleteness(name, data)
		sectionResults = append(sectionResults, result)
		totalPercentage += result.Percentage
	}

	overall := 0.0
	if len(sectionNames) > 0 {
		overall = totalPercentage / float64(len(sectionNames))
	}

	return ProfileCompleteness{
		Overall:  overall,
		Sections: sectionResults,
	}
}
