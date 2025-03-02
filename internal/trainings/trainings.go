package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	splitDatastring := strings.Split(datastring, ",")

	if len(splitDatastring) != 3 {
		return fmt.Errorf("неверное количество аргументов")
	}

	t.Steps, err = strconv.Atoi(splitDatastring[0])
	if err != nil {
		return err
	}

	if splitDatastring[1] != "Бег" && splitDatastring[1] != "Ходьба" {
		return fmt.Errorf("неизвестный тип тренировки")
	}
	t.TrainingType = splitDatastring[1]

	t.Duration, err = time.ParseDuration(splitDatastring[2])
	if err != nil {
		return err
	}

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() string {
	var spentCalories float64

	if t.TrainingType == "Бег" {
		spentCalories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		spentCalories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "неизвестный тип тренировки"
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч.\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), spentCalories)
}
