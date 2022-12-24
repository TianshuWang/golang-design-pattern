package invocation_pipeline

import (
	"log"
	"testing"
)

func TestPipeline(t *testing.T) {
	var users = []user{
		{
			name: "Juan",
			age:  18,
		},
		{
			name: "Maria",
			age:  22,
		},
		{
			name: "Lucia",
			age:  20,
		},
		{
			name: "Matias",
			age:  60,
		},
		{
			name: "Lucas",
			age:  10,
		},
	}

	if sum := SumAge(users, FilterAge, MapAgeToSlice); sum != 60 {
		log.Fatal("Test pipeline fail")
	}

	if sum := SumAge(users, FilterAge, FilterNameStartsWithM, MapAgeToSlice); sum != 22 {
		log.Fatal("Test pipeline fail")
	}
}
