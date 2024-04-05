package services

import (
	"regexp"
	"strings"
)

type HumanRowMapper struct {
	result map[string]string
}

func NewHumanRowMapper() *HumanRowMapper {
	return &HumanRowMapper{
		result: make(map[string]string),
	}
}

func (h *HumanRowMapper) Map(rawRow string) map[string]string {
	rowMapped := h.result

	// Remove 'greater than' and 'less than' symbols of the <world>
	rawRow = strings.ReplaceAll(rawRow, "<", "")
	rawRow = strings.ReplaceAll(rawRow, ">", "")

	re := regexp.MustCompile("Kill: (\\w+) (\\w+) (\\w+): (.+) killed (.+) by (.+)")
	matches := re.FindStringSubmatch(rawRow)

	if len(matches) > 0 {

		rowMapped = map[string]string{
			"who_killed":     matches[4],
			"who_died":       matches[5],
			"means_of_death": matches[6],
		}

	}

	return rowMapped
}
