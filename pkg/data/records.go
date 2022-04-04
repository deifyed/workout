package data

import (
	"encoding/csv"
	"fmt"
	"io"
)

func getWorkouts(f io.Reader) ([]string, map[string]map[string]string, error) {
	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("parsing: %w", err)
	}

	header, workouts := fromRecords(records)

	return header, workouts, nil
}

func fromRecords(records [][]string) ([]string, map[string]map[string]string) {
	if len(records) == 0 {
		return make([]string, 0), make(map[string]map[string]string)
	}

	workouts := make(map[string]map[string]string)

	header := records[0]

	for _, entry := range records[1:] {
		date := entry[0]
		workouts[date] = make(map[string]string)

		for index, repetitions := range entry[1:] {
			workoutType := header[index+1]

			workouts[date][workoutType] = repetitions
		}
	}

	return header, workouts
}

func toRecords(header []string, w map[string]map[string]string) ([][]string, error) {
	headerIndexMap := make(map[string]int)

	for index, item := range header {
		headerIndexMap[item] = index
	}

	result := make([][]string, 0)

	result = append(result, header)

	for date, workouts := range w {
		entry := make([]string, len(header))

		for i := 0; i < len(entry); i++ {
			entry[i] = "0"
		}

		entry[0] = date

		for workout, repetitions := range workouts {
			entry[headerIndexMap[workout]] = repetitions
		}

		result = append(result, entry)
	}

	return result, nil
}
