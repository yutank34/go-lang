package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	nowYear, nowMonth, nowDay := time.Now().Date()
	lastMonth := nowMonth - 1
	if lastMonth < 1 {
		lastMonth = lastMonth + 12
	}
	month := time.Date(nowYear, lastMonth, nowDay, 0, 0, 0, 0, time.UTC)
	year := time.Date(nowYear-1, nowMonth, nowDay, 0, 0, 0, 0, time.UTC)

	monthIssues := make([]*Issue, 0, 10)
	yearIssues := make([]*Issue, 0, 50)
	overYearIssues := make([]*Issue, 0, 100)
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			monthIssues = append(monthIssues, item)
		}
		if item.CreatedAt.After(year) {
			yearIssues = append(yearIssues, item)
		}
		if item.CreatedAt.Before(year) {
			overYearIssues = append(overYearIssues, item)
		}
	}

	fmt.Printf("1ヶ月以内のissue:%v\n", len(monthIssues))
	for _, item := range monthIssues {
		fmt.Printf("#%-5d %9.9s %.55s, %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("1年以内のissue:%v\n", len(yearIssues))
	for _, item := range yearIssues {
		// fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		fmt.Printf("#%-5d %9.9s %.55s, %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("1年以上前のissue:%v\n", len(overYearIssues))
	for _, item := range overYearIssues {
		// fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		fmt.Printf("#%-5d %9.9s %.55s, %v\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
