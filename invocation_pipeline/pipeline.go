package invocation_pipeline

import (
	"log"
	"strings"
)

type user struct {
	name string
	age  int
}

func FilterAge(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if u.age >= 18 && u.age <= 35 {
			slice = append(slice, u)
		}
	}
	return slice
}

func FilterNameStartsWithM(users []user) interface{} {
	var slice []user
	for _, u := range users {
		if strings.HasPrefix(u.name, "M") {
			slice = append(slice, u)
		}
	}
	return slice
}

func MapAgeToSlice(users []user) interface{} {
	var slice []int
	for _, u := range users {
		slice = append(slice, u.age)
	}
	return slice
}

func SumAge(users []user, pipes ...func([]user) interface{}) int {
	var ages []int
	var sum int
	for _, f := range pipes {
		result := f(users)
		switch result.(type) {
		case []user:
			users = result.([]user)
		case []int:
			ages = result.([]int)
		}
	}

	if len(ages) == 0 {
		log.Fatal("No MapAgeToSlice method")
	}

	for _, age := range ages {
		sum += age
	}

	return sum
}
