package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type UserReport struct {
	Fullname string
	Number   string
	Position string
	Email    string
	Status   string
}

func GenFakeReportList(n int) []UserReport {

	rand.Seed(time.Now().Unix())

	statusOptions := []string{"Received", "Clicked", "Credential Log"}

	var userReports []UserReport

	for i := 0; i < n; i++ {
		name := gofakeit.Name()
		numNoFormat := gofakeit.Phone()
		num := fmt.Sprintf("(%s)%s-%s", numNoFormat[0:3], numNoFormat[3:6], numNoFormat[6:10])
		position := gofakeit.JobTitle()
		email := fmt.Sprintf("%s@company.com", strings.ToLower(strings.Join(strings.Split(name, " "), ".")))

		// fmt.Println(name, num, position, email)

		userReport := UserReport{
			Fullname: name,
			Number:   num,
			Position: position,
			Email:    email,
			Status:   statusOptions[rand.Intn(len(statusOptions))],
		}

		userReports = append(userReports, userReport)
	}

	return userReports
	// ReportsTo: "",
}
