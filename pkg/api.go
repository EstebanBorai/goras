package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Api struct {
	username string
	token    string
	client   http.Client
}

func NewApi(username, token string) (*Api, error) {
	var instance *Api = new(Api)

	if len(token) == 0 {
		return nil, errors.New("missing API token")
	}

	instance.username = username
	instance.token = token

	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	instance.client = client

	return instance, nil
}

func (api *Api) GetUserStarts(user string, page int) (GitHubRepositories, error) {
	log.Printf("Fetching starred repositories page: %d\n", page)

	url := fmt.Sprintf("https://api.github.com/users/%s/starred?page=%d", user, page)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var header http.Header = map[string][]string{
		"Authorization": {fmt.Sprintf("token %s", api.token)},
	}

	req.Header = header
	res, err := api.client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		bodyString := string(body[:])

		if err != nil {
			return nil, err
		}

		log.Println(bodyString)

		return nil, fmt.Errorf("received %d from request to %s", res.StatusCode, url)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var starredRepositories GitHubRepositories = make(GitHubRepositories, 0)

	if err = json.Unmarshal([]byte(body), &starredRepositories); err != nil {
		return nil, err
	}

	log.Printf("Fetch success for starred repositories on page: %d\n", page)
	return starredRepositories, nil
}

func (api *Api) DeleteStarredRpository(user string, starredRepository *GitHubRepository) error {
	log.Printf("Removing start for: %s\n", starredRepository.Name)
	url := fmt.Sprintf("https://api.github.com/user/starred/%s/%s", starredRepository.Owner.Login, starredRepository.Name)
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	req.SetBasicAuth(api.username, api.token)
	res, err := api.client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		bodyString := string(body[:])

		if err != nil {
			return err
		}

		log.Println(bodyString)

		return fmt.Errorf("received %d from request to %s", res.StatusCode, url)
	}

	log.Printf("Removed start for: %s with success\n", starredRepository.Name)
	return nil
}
