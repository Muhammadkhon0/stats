package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Muhammadkhon0/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   50_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 3400

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "Cafe",
			Status:   types.StatusFail,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Restaurant",
			Status:   types.StatusOk,
		},
	}

	fmt.Println(TotalInCategory(payments, "Cafe"))

	//Output: 5300
}

func TestCategoriesAvg_manyCategory(t *testing.T) {
	payments := []types.Payment{
		{
			Category: "car",
			Amount:   0,
			Status:   types.StatusOk,
		},
		{
			Category: "car",
			Amount:   100,
			Status:   types.StatusInProgress,
		},
		{
			Category: "food",
			Amount:   10000,
			Status:   types.StatusOk,
		},
		{
			Category: "fun",
			Amount:   100,
			Status:   types.StatusOk,
		},
		{
			Category: "fun",
			Amount:   100,
			Status:   types.StatusFail,
		},
	}
	want := map[types.Category]types.Money{
		"car":  50,
		"food": 10000,
		"fun":  100,
	}
	got := CategoriesAvg(payments)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	want := map[types.Category]types.Money{}
	got := CategoriesAvg(payments)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestCategoriesAvg_notFound(t *testing.T) {
	payments := []types.Payment{
		{
			Category: "car",
			Amount:   0,
			Status:   types.StatusFail,
		},
		{
			Category: "car",
			Amount:   100,
			Status:   types.StatusFail,
		},
	}
	want := map[types.Category]types.Money{}
	got := CategoriesAvg(payments)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestPeriodsDynamic_empty(t *testing.T) {
	first, second := map[types.Category]types.Money{}, map[types.Category]types.Money{}
	want := map[types.Category]types.Money{}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
func TestPeriodsDynamic_oneParam(t *testing.T) {
	first, second := map[types.Category]types.Money{
		"car": 10,
	}, map[types.Category]types.Money{
		"food": 5,
	}
	want := map[types.Category]types.Money{
		"car":  -10,
		"food": 5,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestPeriodsDynamic_manyParam(t *testing.T) {
	first, second := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}, map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	}
	want := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}