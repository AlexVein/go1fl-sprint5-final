package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

var (
	ErrInvalidArgumentsCount = errors.New("invalid arguments count")
	ErrInvalidFormat         = errors.New("invalid format")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("%w: %s", ErrInvalidArgumentsCount, datastring)
	}

	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, data[0])
	}

	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, data[1])
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var result string
	var distance float64
	var walkingSpentCalories float64
	var runningSpentCalories float64
	var spentCalories float64
	var err error

	distance = spentenergy.Distance(ds.Steps, ds.Height)

	walkingSpentCalories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	runningSpentCalories, err = spentenergy.RunningSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	spentCalories = walkingSpentCalories + runningSpentCalories

	result = fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, spentCalories)

	return result, err
}
