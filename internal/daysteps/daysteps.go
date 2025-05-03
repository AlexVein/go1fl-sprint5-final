package daysteps

import (
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
