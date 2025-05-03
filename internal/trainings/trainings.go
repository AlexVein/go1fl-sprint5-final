package trainings

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
	ErrUnknownTrainingType   = errors.New("неизвестный тип тренировки")
	ErrZeroOrNegativeValue   = errors.New("zero or negative value")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("%w: %s", ErrInvalidArgumentsCount, datastring)
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, data[0])
	}

	if t.Steps <= 0 {
		return fmt.Errorf("%w: %d", ErrZeroOrNegativeValue, t.Steps)
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidFormat, data[2])
	}

	if t.Duration <= 0 {
		return fmt.Errorf("%w: %s", ErrZeroOrNegativeValue, t.Duration)
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var result string
	var distance float64
	var meanSpeed float64
	var spentCalories float64
	var err error

	distance = spentenergy.Distance(t.Steps, t.Height)
	meanSpeed = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("%w: %s", ErrUnknownTrainingType, t.TrainingType)
	}

	result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories)

	return result, err
}
