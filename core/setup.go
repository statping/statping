package core

import (
	"fmt"
	"github.com/hunterlong/statup/utils"
	"os"
)

func InsertDefaultComms() {
	emailer := &Communication{
		Method:    "email",
		Removable: false,
		Enabled:   false,
	}
	Create(emailer)
}

func DeleteConfig() {
	err := os.Remove("./config.yml")
	if err != nil {
		utils.Log(3, err)
	}
}

type ErrorResponse struct {
	Error string
}

func LoadSampleData() error {
	fmt.Println("Inserting Sample Data...")
	s1 := &Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s2 := &Service{
		Name:           "Statup.io",
		Domain:         "https://statup.io",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s3 := &Service{
		Name:           "Statup.io SSL Check",
		Domain:         "https://statup.io",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           443,
		Type:           "tcp",
	}
	s4 := &Service{
		Name:           "Github Failing Check",
		Domain:         "https://github.com/thisisnotausernamemaybeitis",
		ExpectedStatus: 200,
		Interval:       15,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	s1.Create()
	s2.Create()
	s3.Create()
	s4.Create()

	checkin := &Checkin{
		Service:  s2.Id,
		Interval: 30,
		Api:      utils.NewSHA1Hash(18),
	}
	checkin.Create()

	for i := 0; i < 20; i++ {
		s1.Check()
		s2.Check()
		s3.Check()
		s4.Check()
	}

	return nil
}
