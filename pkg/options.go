package pkg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	Token       string
	Username    string
	tokenSource string
}

func NewOptions() (*Options, error) {
	var options *Options = new(Options)

	if token, tokenSource, err := getGitHubUserToken(); err == nil {
		options.Token = token
		options.tokenSource = tokenSource
	} else {
		return nil, err
	}

	if username, err := getGitHubUserName(); err == nil {
		options.Username = username
	} else {
		return nil, err
	}

	return options, nil
}

func getGitHubUserToken() (string, string, error) {
	var gitHubUserToken string = os.Getenv("GITHUB_USER_TOKEN")

	if len(gitHubUserToken) == 0 {
		if token, err := ReadStdin("GitHub User Token: "); err == nil {
			return token, "Stdin Prompt", nil
		} else {
			return "", "", err
		}
	}

	return gitHubUserToken, "Environment Variable", nil
}

func getGitHubUserName() (string, error) {
	var gitHubUserName string = os.Getenv("GITHUB_USERNAME")

	if len(gitHubUserName) == 0 {
		if token, err := ReadStdin("GitHub Username: "); err == nil {
			return token, nil
		} else {
			return "", err
		}
	}

	return gitHubUserName, nil
}

func (opts *Options) displayOptions() {
	fmt.Print("\n\n")
	fmt.Println("===========================")
	fmt.Println("Deathstar Execution Options")
	fmt.Println("===========================")
	fmt.Printf("Username:\t%s\n", opts.Username)
	fmt.Printf("Token Source:\t%s", opts.tokenSource)
	fmt.Print("\n\n")
}

func (opts *Options) PromptConfirmOptions() error {
	fmt.Println("Confirm execution options")

	opts.displayOptions()

	answer, err := ReadStdin("[Y]: Confirm (Any other key will abort the process)")

	if err != nil {
		return err
	}

	if strings.EqualFold(answer, "Y") {
		return nil
	}

	return errors.New("execution options are not confirmed")
}
