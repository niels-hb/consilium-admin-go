package random

import (
	"testing"

	"github.com/niels-hb/consilium-admin/arrays"
	"github.com/niels-hb/consilium-admin/models"
)

func TestGetRandomIntInRange(t *testing.T) {
	got, _ := GetRandomIntInRange(0, 1)
	if got < 0 || got > 1 {
		t.Errorf("GetRandomIntInRange(0, 1) = %v; want 0 <= x <= 1", got)
	}

	got, err := GetRandomIntInRange(1, 0)
	if err == nil {
		t.Errorf("GetRandomIntInRange(1, 0) = (%v, %v); want (0, 0 (max) < 1 (min))", got, err)
	}

	for i := 0; i < 1000; i++ {
		got, _ = GetRandomIntInRange(100, 1000)
		if got < 100 || got > 1000 {
			t.Errorf("GetRandomIntInRange(100, 1000) = %v; want 100 <= x <= 1000", got)
		}
	}

	for i := 0; i < 1000; i++ {
		got, _ = GetRandomIntInRange(-100, 100)
		if got < -100 || got > 100 {
			t.Errorf("GetRandomIntInRange(-100, 100) = %v; want -100 <= x <= 100", got)
		}
	}

	for i := 0; i < 1000; i++ {
		got, _ = GetRandomIntInRange(-200, -100)
		if got < -200 || got > -100 {
			t.Errorf("GetRandomIntInRange(-200, -100) = %v; want -200 <= x <= -100", got)
		}
	}
}

func BenchmarkGetRandomIntInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomIntInRange(0, 1e6)
	}
}

func TestGetRandomTime(t *testing.T) {
	got := GetRandomTime()
	if got == nil {
		t.Errorf("GetRandomTime() = %v; want valid *time.Time", got)
	}
}

func TestGetRandomCategory(t *testing.T) {
	for i := 0; i < len(models.GetCategories())*10; i++ {
		got := GetRandomCategory()
		if !arrays.Contains(models.GetCategories(), got) {
			t.Errorf("GetRandomCategory() = %v; want any in %v", got, models.GetCategories())
		}
	}
}

func TestGetRandomScheduleType(t *testing.T) {
	for i := 0; i < len(models.GetScheduleTypes())*10; i++ {
		got := GetRandomScheduleType()
		if !arrays.Contains(models.GetScheduleTypes(), got) {
			t.Errorf("GetRandomScheduleType() = %v; want any in %v", got, models.GetScheduleTypes())
		}
	}
}

func TestValidateParameterRange(t *testing.T) {
	got := validateParameterRange(0, 1)
	if !got {
		t.Errorf("validateParameterRange(0, 1) = %v; want true", got)
	}

	got = validateParameterRange(1, 0)
	if got {
		t.Errorf("validateParameterRange(1, 0) = %v; want false", got)
	}

	got = validateParameterRange(0, 0)
	if !got {
		t.Errorf("validateParameterRange(0, 0) = %v; want true", got)
	}
}
