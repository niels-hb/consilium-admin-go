package random

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/niels-hb/consilium-admin/models"
)

func GetRandomIntInRange(min int, max int) (int, error) {
	if !validateParameterRange(min, max) {
		return 0, fmt.Errorf("%v (max) < %v (min)", max, min)
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Intn((max+1)-min) + min, nil
}

func GetRandomTime() *time.Time {
	random := time.Unix(faker.UnixTime(), 0)

	return &random
}

func GetRandomCategory() string {
	idx, _ := GetRandomIntInRange(0, len(models.GetCategories())-1)

	return models.GetCategories()[idx]
}

func GetRandomScheduleType() string {
	idx, _ := GetRandomIntInRange(0, len(models.GetScheduleTypes())-1)

	return models.GetScheduleTypes()[idx]
}

func validateParameterRange(min int, max int) bool {
	return max >= min
}
