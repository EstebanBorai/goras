package pkg

type GitHubRepository struct {
	ID          int32                 `json:"id"`
	NodeID      string                `json:"node_id"`
	Name        string                `json:"name"`
	FullName    string                `json:"full_name"`
	Private     bool                  `json:"private"`
	HtmlURL     string                `json:"html_url"`
	Description string                `json:"description"`
	Fork        bool                  `json:"fork"`
	Owner       GitHubRepositoryOwner `json:"owner"`
}

type GitHubRepositoryOwner struct {
	ID    int32  `json:"id"`
	Login string `json:"login"`
}

type GitHubRepositories []GitHubRepository
