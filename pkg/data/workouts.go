package data

func (receiver *workoutManager) Add(workout Workout) error {
	date := workout.Date.Format(timeFormat)

	if _, ok := receiver.workouts[date]; !ok {
		receiver.workouts[date] = make(map[string]string)
	}

	if _, ok := receiver.workouts[date][workout.Type]; ok {
		return errAlreadyExists
	}

	receiver.workouts[date][workout.Type] = workout.Repetitions

	return nil
}

type workoutManager struct {
	workouts map[string]map[string]string
}

const timeFormat = "02.01.2006"
