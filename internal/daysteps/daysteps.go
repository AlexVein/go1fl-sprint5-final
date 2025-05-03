package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

var (
	ErrInvalidArgumentsCount = errors.New("invalid arguments count")
	ErrInvalidFormat         = errors.New("invalid format")
	ErrUnknownTrainingType   = errors.New("неизвестный тип тренировки")
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
	// TODO: реализовать функцию
}
