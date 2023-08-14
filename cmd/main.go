package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fahrurben/users-management-cli/internal"
)

func main() {
	// Default urls
	urls := []string{
		"https://run.mocky.io/v3/03d2a7bd-f12f-4275-9e9a-84e41f9c2aae",
		"https://run.mocky.io/v3/87931203-8086-43ef-ba16-4c8903d8fa88",
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl != "" {
		urls = strings.Split(apiUrl, ",")
	}

	if len(os.Args) == 1 {
		panic("Wrong arguments")
	}
	commandArg := os.Args[1]

	service := &internal.Service{}

	if commandArg == "save" {
		users, err := service.GetUsers(urls)
		if err != nil {
			panic(err)
		}
		err = service.SaveUsers(users)
		if err != nil {
			panic(err)
		}
		fmt.Println("Data saved succesfully at folder data with filename users.csv")
	} else if commandArg == "search" {
		if len(os.Args) < 2 {
			panic("Wrong arguments")
		}
		tagsArg := os.Args[2]
		if strings.HasPrefix(tagsArg, "--tag=") != true {
			panic("Search must have --tag arguments")
		}

		tagsArg = strings.Replace(tagsArg, "--tag=", "", 1)
		tags := strings.Split(tagsArg, ",")
		users, err := service.SearchUsers(tags)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%-24s %s \n", "User Id", "Balance")
		fmt.Println("------------------------ ---------")
		for _, user := range users {
			fmt.Printf("%-24s %s \n", user.Id, user.Balance)
		}
	}
}
