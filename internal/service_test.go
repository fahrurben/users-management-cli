package internal

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var service *Service

const API_MOCK_1 = "https://api.test.com/1"
const API_MOCK_2 = "https://api.test.com/2"
const RESPONSE_MOCK = `[
	{
			"_id": "64d39b0582ec3cff5fc7f24e",
			"index": 0,
			"guid": "03ee84da-5a54-493f-8438-60bad7ab6e2a",
			"isActive": true,
			"balance": "$2,633.92",
			"tags": [
					"pariatur",
					"qui",
					"ea",
					"culpa",
					"laboris",
					"laboris",
					"minim"
			],
			"friends": [
					{
							"id": 0,
							"name": "Koch Valdez"
					},
					{
							"id": 1,
							"name": "Kramer Bush"
					},
					{
							"id": 2,
							"name": "Townsend Church"
					}
			]
	}
]`

func init() {
	service = &Service{}
}

func TestGetUsers(t *testing.T) {
	urls := []string{
		API_MOCK_1,
		API_MOCK_2,
	}

	httpmock.Activate()
	httpmock.RegisterResponder(
		"GET",
		API_MOCK_1,
		httpmock.NewStringResponder(500, ""),
	)

	httpmock.RegisterResponder(
		"GET",
		API_MOCK_2,
		httpmock.NewStringResponder(200, RESPONSE_MOCK),
	)

	users, err := service.GetUsers(urls)

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))

	user := users[0]
	friendsJsonStr, _ := json.Marshal(user.Friends)

	assert.Equal(t, "64d39b0582ec3cff5fc7f24e", user.Id)
	assert.Equal(t, 0, user.Index)
	assert.Equal(t, "03ee84da-5a54-493f-8438-60bad7ab6e2a", user.Guid)
	assert.Equal(t, true, user.IsActive)
	assert.Equal(t, "$2,633.92", user.Balance)
	assert.Equal(t, "64d39b0582ec3cff5fc7f24e", user.Id)
	assert.Equal(t, "pariatur,qui,ea,culpa,laboris,laboris,minim", strings.Join(user.Tags[:], ","))
	assert.Equal(t, `[{"id":0,"name":"Koch Valdez"},{"id":1,"name":"Kramer Bush"},{"id":2,"name":"Townsend Church"}]`, string(friendsJsonStr))
}

func TestSaveUsers(t *testing.T) {
	urls := []string{
		API_MOCK_1,
		API_MOCK_2,
	}

	httpmock.Activate()
	httpmock.RegisterResponder(
		"GET",
		API_MOCK_1,
		httpmock.NewStringResponder(500, ""),
	)

	httpmock.RegisterResponder(
		"GET",
		API_MOCK_2,
		httpmock.NewStringResponder(200, RESPONSE_MOCK),
	)

	users, err := service.GetUsers(urls)
	assert.Nil(t, err)

	err = service.SaveUsers(users)
	assert.Nil(t, err)

	file, err := os.Open("../data/users.csv")
	assert.Nil(t, err)
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	record := records[0]
	assert.Nil(t, err)
	assert.Equal(t, "64d39b0582ec3cff5fc7f24e", record[0])
	assert.Equal(t, "0", record[1])
	assert.Equal(t, "03ee84da-5a54-493f-8438-60bad7ab6e2a", record[2])
	assert.Equal(t, "true", record[3])
	assert.Equal(t, "$2,633.92", record[4])
	assert.Equal(t, "pariatur,qui,ea,culpa,laboris,laboris,minim", record[5])
}

func TestSearchUser(t *testing.T) {
	urls := []string{
		API_MOCK_1,
		API_MOCK_2,
	}

	httpmock.Activate()
	httpmock.RegisterResponder(
		"GET",
		API_MOCK_1,
		httpmock.NewStringResponder(500, ""),
	)

	httpmock.RegisterResponder(
		"GET",
		API_MOCK_2,
		httpmock.NewStringResponder(200, RESPONSE_MOCK),
	)

	users, err := service.GetUsers(urls)
	assert.Nil(t, err)

	err = service.SaveUsers(users)
	assert.Nil(t, err)

	results, err := service.SearchUsers([]string{"pariatur", "qui"})
	assert.NotNil(t, results)
}
