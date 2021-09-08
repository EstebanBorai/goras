package pkg

import "fmt"

const PAGE_LIMIT int = 10

type Statistics struct {
	totalStarredRepositories int
	totalDeletedRepositories int
	totalPagesFetched        int
}

type App struct {
	api        Api
	options    Options
	statistics Statistics
}

func NewApp(options Options) *App {
	var instance *App = new(App)

	if err := options.PromptConfirmOptions(); err != nil {
		panic(err)
	}

	api, createApiError := NewApi(options.Token)

	if createApiError != nil {
		panic(createApiError)
	}

	instance.api = *api
	instance.options = options

	return instance
}

func (app *App) Start() error {
	starredRepositories, getUserStartsErr := app.fetchStarredRepos()

	if getUserStartsErr != nil {
		return getUserStartsErr
	}

	deleteStarredRepositoriesErr := app.deleteStarredRepositories(starredRepositories)

	if deleteStarredRepositoriesErr != nil {
		return deleteStarredRepositoriesErr
	}

	fmt.Printf("Total Starred Repositories:\t%d\n", app.statistics.totalStarredRepositories)
	fmt.Printf("Total Unstarted Repositories:\t%d", app.statistics.totalDeletedRepositories)

	return nil
}

func (app *App) fetchStarredRepos() (*GitHubRepositories, error) {
	var loopIteration int = 1
	var starredRepositories GitHubRepositories = make([]GitHubRepository, 0)

	for {
		gitHubStarredRepos, getUserStartsErr := app.api.GetUserStarts(app.options.Username, loopIteration)

		if getUserStartsErr != nil {
			return nil, getUserStartsErr
		}

		repositoryCount := len(gitHubStarredRepos)

		if repositoryCount == 0 || loopIteration == PAGE_LIMIT {
			break
		}

		app.statistics.totalStarredRepositories += repositoryCount

		for _, el := range gitHubStarredRepos {
			starredRepositories = append(gitHubStarredRepos, el)
		}

		loopIteration++
	}

	app.statistics.totalPagesFetched = loopIteration

	return &starredRepositories, nil
}

func (app *App) deleteStarredRepositories(starredRepositories *GitHubRepositories) error {
	for _, el := range *starredRepositories {
		if err := app.api.DeleteStarredRpository(app.options.Username, &el); err != nil {
			return err
		}

		app.statistics.totalDeletedRepositories++
	}

	return nil
}
