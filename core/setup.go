package core

import (
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"os"
)

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
	utils.Log(1, "Inserting Sample Data...")
	s1 := &types.Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
	}
	s2 := &types.Service{
		Name:           "Statup Github",
		Domain:         "https://github.com/hunterlong/statup",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	}
	s3 := &types.Service{
		Name:           "JSON Users Test",
		Domain:         "https://jsonplaceholder.typicode.com/users",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	}
	s4 := &types.Service{
		Name:           "JSON API Tester",
		Domain:         "https://jsonplaceholder.typicode.com/posts",
		ExpectedStatus: 201,
		Expected:       `(title)": "((\\"|[statup])*)"`,
		Interval:       30,
		Type:           "http",
		Method:         "POST",
		PostData:       `{ "title": "statup", "body": "bar", "userId": 19999 }`,
		Timeout:        30,
	}
	s5 := &types.Service{
		Name:     "Postgres TCP Check",
		Domain:   "0.0.0.0",
		Interval: 20,
		Type:     "tcp",
		Port:     5432,
		Timeout:  120,
	}
	id, err := CreateService(s1)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = CreateService(s2)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = CreateService(s3)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = CreateService(s4)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = CreateService(s5)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating TCP Service %v: %v", id, err))
	}

	//checkin := &Checkin{
	//	Service:  s2.Id,
	//	Interval: 30,
	//	Api:      utils.NewSHA1Hash(18),
	//}
	//id, err = checkin.Create()
	//if err != nil {
	//	utils.Log(3, fmt.Sprintf("Error creating Checkin %v: %v", id, err))
	//}

	//for i := 0; i < 3; i++ {
	//	s1.Check()
	//	s2.Check()
	//	s3.Check()
	//	s4.Check()
	//}

	utils.Log(1, "Sample data has finished importing")

	return nil
}
