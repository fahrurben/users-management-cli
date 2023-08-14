package internal

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Service struct {
}

func (service *Service) GetUsers(urls []string) ([]User, error) {
	var users []User
	for _, url := range urls {
		arrData, err := service.fetchUsers(url)
		if err != nil {
			continue
		}

		users = append(users, arrData...)
	}
	return users, nil
}

func (service *Service) fetchUsers(url string) ([]User, error) {
	var users []User = make([]User, 0)

	var client = &http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(response.Body)
		return nil, errors.New(string(bodyBytes))
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *Service) SaveUsers(users []User) error {
	records := [][]string{}

	for _, user := range users {
		records = append(records, []string{
			user.Id,
			strconv.Itoa(user.Index),
			user.Guid,
			strconv.FormatBool(user.IsActive),
			user.Balance,
			strings.Join(user.Tags[:], ","),
		})
	}

	csvFile, err := os.Create("../data/users.csv")
	if err != nil {
		return err
	}

	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
