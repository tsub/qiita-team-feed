package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Item struct {
	RenderedBody string        `json:"rendered_body"`
	Body         string        `json:"body"`
	CoEditing    bool          `json:"coediting"`
	CreatedAt    time.Time     `json:"created_at"`
	ID           string        `json:"id"`
	Private      bool          `json:"private"`
	Tags         []Tag         `json:"tags"`
	Title        string        `json:"title"`
	UpdatedAt    time.Time     `json:"updated_at"`
	URL          string        `json:"url"`
	User         User          `json:"user"`
}

type Tag struct {
	FollowersCount int    `json:"followers_count"`
	IconURL        string `json:"icon_url"`
	ID             string `json:"id"`
	ItemsCount     int    `json:"items_count"`
}

type User struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	FacebookID        string `json:"facebook_id"`
	FolloweesCount    int    `json:"followees_count"`
	FollowersCount    int    `json:"followers_count"`
	GithubLoginName   string `json:"github_login_name"`
	ID                string `json:"id"`
	ItemsCount        int    `json:"items_count"`
	LinkedinID        string `json:"linkedin_id"`
	Location          string `json:"location"`
	Organization      string `json:"organization"`
	PermanentID       int    `json:"permanent_id"`
	ProfileImageURL   string `json:"profile_image_url"`
	TwitterScreenName string `json:"twitter_screen_name"`
	WebsiteURL        string `json:"website_url"`
}

func main() {
	token := os.Getenv("QIITA_ACCESS_TOKEN")
	team := os.Getenv("QIITA_TEAM_NAME")

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://" + team + ".qiita.com/api/v2/items", nil)
	req.Header.Add("Authorization", "Bearer " + token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	items := new([]Item)
	err = json.NewDecoder(resp.Body).Decode(items)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range *items {
		fmt.Println(item.Title)
	}
}
