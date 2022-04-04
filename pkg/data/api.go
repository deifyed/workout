package data

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/afero"
)

func (receiver DataClient) Insert(workout Workout) error {
	f, err := ensureOpen(receiver.fs, receiver.dbPath)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seeking before reading workouts: %w", err)
	}

	header, workouts, err := getWorkouts(f)
	if err != nil {
		return fmt.Errorf("acquiring records: %w", err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("seeking before writing records: %w", err)
	}

	if !contains(header, workout.Type) {
		header = append(header, workout.Type)
	}

	workoutMgr := workoutManager{
		workouts: workouts,
	}

	err = workoutMgr.Add(workout)
	if err != nil {
		return fmt.Errorf("adding workout: %w", err)
	}

	updatedRecords, err := toRecords(header, workouts)
	if err != nil {
		return fmt.Errorf("converting to records: %w", err)
	}

	writer := csv.NewWriter(f)

	err = writer.WriteAll(updatedRecords)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func NewClient(fs *afero.Afero, dbPath string) DataClient {
	return DataClient{
		fs:     fs,
		dbPath: dbPath,
	}
}
