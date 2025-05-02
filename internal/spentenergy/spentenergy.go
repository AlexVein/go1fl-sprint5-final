package spentenergy

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	//ErrInvalidArgument = errors.New("invalid argument")
	ErrZeroOrNegative = errors.New("zero or negative value")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("%w: steps must be greater than zero", ErrZeroOrNegative)
	}

	if height <= 0 {
		return 0, fmt.Errorf("%w: height must be greater than zero", ErrZeroOrNegative)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("%w: duration must be greater than zero", ErrZeroOrNegative)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("%w: weight must be greater than zero", ErrZeroOrNegative)
	}

	spentCalories := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		log.Println(ErrZeroOrNegative)
		return 0
	}

	if duration <= 0 {
		log.Println(ErrZeroOrNegative)
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
