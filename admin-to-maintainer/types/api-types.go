package types

import "time"

// api-types.go maps GitHub API JSON objects to Go structs

// User represents the full user record of a particular user on GitHub
type User struct {
	SiteAdmin         bool       `json:"site_admin,omitempty"`
	Login             string     `json:"login,omitempty"`
	ID                int        `json:"id,omitempty"`
	AvatarURL         string     `json:"avatar_url,omitempty"`
	GravatarID        string     `json:"gravatar_id,omitempty"`
	URL               string     `json:"url,omitempty"`
	Name              string     `json:"name,omitempty"`
	Company           string     `json:"company,omitempty"`
	Blog              string     `json:"blog,omitempty"`
	Location          string     `json:"location,omitempty"`
	Email             string     `json:"email,omitempty"`
	Hireable          bool       `json:"hireable,omitempty"`
	Bio               string     `json:"bio,omitempty"`
	PublicRepos       int        `json:"public_repos,omitempty"`
	Followers         int        `json:"followers,omitempty"`
	Following         int        `json:"following,omitempty"`
	HTMLURL           string     `json:"html_url,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	Type              string     `json:"type,omitempty"`
	FollowingURL      string     `json:"following_url,omitempty"`
	FollowersURL      string     `json:"followers_url,omitempty"`
	GistsURL          string     `json:"gists_url,omitempty"`
	StarredURL        string     `json:"starred_url,omitempty"`
	SubscriptionsURL  string     `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string     `json:"organizations_url,omitempty"`
	ReposURL          string     `json:"repos_url,omitempty"`
	EventsURL         string     `json:"events_url,omitempty"`
	ReceivedEventsURL string     `json:"received_events_url,omitempty"`
}

// Repository represents a respository on GitHub with all associated metadata
type Repository struct {
	ID              int           `json:"id,omitempty"`
	Owner           User          `json:"owner,omitempty"`
	Name            string        `json:"name,omitempty"`
	FullName        string        `json:"full_name,omitempty"`
	Description     string        `json:"description,omitempty"`
	Private         bool          `json:"private"`
	Fork            bool          `json:"fork,omitempty"`
	URL             string        `json:"url,omitempty"`
	HTMLURL         string        `json:"html_url,omitempty"`
	CloneURL        string        `json:"clone_url,omitempty"`
	GitURL          string        `json:"git_url,omitempty"`
	SSHURL          string        `json:"ssh_url,omitempty"`
	SVNURL          string        `json:"svn_url,omitempty"`
	MirrorURL       string        `json:"mirror_url,omitempty"`
	Homepage        string        `json:"homepage,omitempty"`
	Language        string        `json:"language,omitempty"`
	Forks           int           `json:"forks,omitempty"`
	ForksCount      int           `json:"forks_count,omitempty"`
	StargazersCount int           `json:"stargazers_count,omitempty"`
	Watchers        int           `json:"watchers,omitempty"`
	WatchersCount   int           `json:"watchers_count,omitempty"`
	Size            int           `json:"size,omitempty"`
	MasterBranch    string        `json:"master_branch,omitempty"`
	OpenIssues      int           `json:"open_issues,omitempty"`
	PushedAt        *time.Time    `json:"pushed_at,omitempty"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
	Permissions     Permissions   `json:"permissions,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	Parent          *Repository   `json:"parent,omitempty"`
	Source          *Repository   `json:"source,omitempty"`
	HasIssues       bool          `json:"has_issues,omitempty"`
	HasWiki         bool          `json:"has_wiki,omitempty"`
	HasDownloads    bool          `json:"has_downloads,omitempty"`
}

// Permissions represent the permissions as they apply to the accessing url
type Permissions struct {
	Admin bool
	Push  bool
	Pull  bool
}

// Organization represents an organization on GitHub with all associated metadata
type Organization struct {
	Description      string     `json:"description,omitempty"`
	AvatarURL        string     `json:"avatar_url,omitempty"`
	PublicMembersURL string     `json:"public_member_url,omitempty"`
	MembersURL       string     `json:"members_url,omitempty"`
	EventsURL        string     `json:"events_url,omitempty"`
	ReposURL         string     `json:"repos_url,omitempty"`
	URL              string     `json:"url,omitempty"`
	ID               int        `json:"id,omitempty"`
	Login            string     `json:"login,omitempty"`
	Name             string     `json:"name,omitempty"`
	Company          string     `json:"company,omitempty"`
	Blog             string     `json:"blog,omitempty"`
	Location         string     `json:"location,omitempty"`
	Email            string     `json:"email,omitempty"`
	PublicRepos      int        `json:"public_repos,omitempty"`
	PublicGists      int        `json:"public_gists,omitempty"`
	Followers        int        `json:"followers,omitempty"`
	Followering      int        `json:"following,omitempty"`
	HTMLURL          string     `json:"html_url,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	Type             string     `json:"type,omitempty"`

	//Limited Access Fields
	TotalPrivateRepos int    `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos int    `json:"owned_private_repos,omitempty"`
	PrivateGists      int    `json:"private_gists,omitempty"`
	DiskUsage         int    `json:"disk_usage,omitempty"`
	Collaborators     int    `json:"collaborators,omitempty"`
	BillingEmail      string `json:"billing_email,omitempty"`
	Plan              Plan   `json:"plan,omitempty"`
}

// Plan represents the state of an organization's GitHub plan
type Plan struct {
	Name         string `json:"name,omitempty"`
	Space        int    `json:"space,omitempty"`
	PrivateRepos int    `json:"private_repos,omitempty"`
}
