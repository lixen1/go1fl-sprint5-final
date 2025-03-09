package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	splitDatastring := strings.Split(datastring, ",")

	if len(splitDatastring) != 2 {
		return fmt.Errorf("неверное количество аргументов")
	}

	ds.Steps, err = strconv.Atoi(splitDatastring[0])
	if err != nil {
		return err
	}

	ds.Duration, err = time.ParseDuration(splitDatastring[1])
	if err != nil {
		return err
	}

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() string {
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, spentenergy.Distance(ds.Steps), spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration))
}
