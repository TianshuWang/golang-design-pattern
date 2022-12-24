package gojsonq

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"testing"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []user{
	{
		Name: "Juan",
		Age:  18,
	},
	{
		Name: "Maria",
		Age:  22,
	},
	{
		Name: "Lucia",
		Age:  20,
	},
	{
		Name: "Matias",
		Age:  60,
	},
	{
		Name: "Lucas",
		Age:  10,
	},
}

func TestGojsonq(t *testing.T) {

	j, err := json.Marshal(users)

	if err != nil {
		fmt.Printf("Got error marshalling users: %v", err)
	}

	sum := gojsonq.New().FromString(string(j)).Where("age", ">=", 18).
		Where("age", "<=", 35).
		Sum("age")
	if sum != 60 {
		t.Fatal("Test gojsonq filter sum fail")
	}

	max := gojsonq.New().FromString(string(j)).Max("age")
	if max != 60 {
		t.Fatal("Test gojsonq max fail")
	}
}
