package services

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
	"time"
)

func (s *Service) Samples() []database.DbObject {
	createdOn := time.Now().Add(((-24 * 30) * 3) * time.Hour).UTC()
	s1 := &Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          1,
		GroupId:        1,
		Permalink:      null.NewNullString("google"),
		VerifySSL:      null.NewNullBool(true),
		CreatedAt:      createdOn,
	}
	s2 := &Service{
		Name:           "Statping Github",
		Domain:         "https://github.com/hunterlong/statping",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          2,
		Permalink:      null.NewNullString("statping_github"),
		VerifySSL:      null.NewNullBool(true),
		CreatedAt:      createdOn,
	}
	s3 := &Service{
		Name:           "JSON Users Test",
		Domain:         "https://jsonplaceholder.typicode.com/users",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
		Order:          3,
		Public:         null.NewNullBool(true),
		VerifySSL:      null.NewNullBool(true),
		GroupId:        2,
		CreatedAt:      createdOn,
	}
	s4 := &Service{
		Name:           "JSON API Tester",
		Domain:         "https://jsonplaceholder.typicode.com/posts",
		ExpectedStatus: 201,
		Expected:       null.NewNullString(`(title)": "((\\"|[statping])*)"`),
		Interval:       30,
		Type:           "http",
		Method:         "POST",
		PostData:       null.NewNullString(`{ "title": "statping", "body": "bar", "userId": 19999 }`),
		Timeout:        30,
		Order:          4,
		Public:         null.NewNullBool(true),
		VerifySSL:      null.NewNullBool(true),
		GroupId:        2,
		CreatedAt:      createdOn,
	}
	s5 := &Service{
		Name:      "Google DNS",
		Domain:    "8.8.8.8",
		Interval:  20,
		Type:      "tcp",
		Port:      53,
		Timeout:   120,
		Order:     5,
		Public:    null.NewNullBool(true),
		GroupId:   1,
		CreatedAt: createdOn,
	}

	return []database.DbObject{s1, s2, s3, s4, s5}
}
