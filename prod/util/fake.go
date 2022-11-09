package util

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type FakeUser struct {
	Fullname string
	Number   string
	Position string
	Email    string
}

// ReportsTo string

func GenFakeUser() FakeUser {

	name := gofakeit.Name()
	numNoFormat := gofakeit.Phone()
	num := fmt.Sprintf("(%s)%s-%s", numNoFormat[0:3], numNoFormat[3:6], numNoFormat[6:10])
	position := gofakeit.JobTitle()
	email := fmt.Sprintf("%s@company.com", strings.ToLower(strings.Join(strings.Split(name, " "), ".")))

	// fmt.Println(name, num, position, email)

	return FakeUser{
		Fullname: name,
		Number:   num,
		Position: position,
		Email:    email,
	}
	// ReportsTo: "",
}
