package integration

import "github.com/ddd/internal/context/log_handler/domain/model/human_logfile"

func PreSendCommand(resultHumanLogFile []human_logfile.HumanLogFileRowable) [][]string {
	var rawData [][]string
	for _, row := range resultHumanLogFile {
		rawData = append(rawData, []string{
			row.GetPlayerWhoKilled(),
			row.GetPlayerWhoDied(),
			row.GetMeanOfDeath(),
		})
	}
	return rawData
}
