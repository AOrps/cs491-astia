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
	// ReportsTo string
	Email string
}

func GenFakeUser() FakeUser {

	name := gofakeit.Name()
	num := gofakeit.Phone()
	position := gofakeit.JobTitle()
	email := fmt.Sprintf("%s@company.com", strings.Join(strings.Split(name, " "), "."))

	return FakeUser{
		Fullname: name,
		Number:   num,
		Position: position,
		// ReportsTo: "",
		Email: email,
	}
}
