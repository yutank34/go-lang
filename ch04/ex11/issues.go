package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// repo:golang/go is:open json decoder
func getIssues(query []string) {
	result, err := SearchIssues(query)
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

const createUrl = "https://api.github.com/repos/yutank34/go-lang/issues"

func createIssue(parts []string) {
	issue := new(reqIssue)
	issue.Title = parts[0]
	if len(parts) == 2 {
		issue.Body = parts[1]
	}
	result, err := postCreateIssue(issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("登録完了" + strconv.Itoa(result.Id) + result.Title)
}

type reqIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
type resIssue struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func postCreateIssue(reqBody *reqIssue) (*resIssue, error) {
	b, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", createUrl, bytes.NewBuffer(b))
	req.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result resIssue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return &result, err
	}
	resp.Body.Close()
	return &result, nil
}
