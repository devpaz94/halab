package dto

type Bath struct {
	ID   string
	Name string
	Size string
	URL  string
}

type User struct {
	ID           string
	Name         string
	Email        string
	ProfilePhoto string
}

type UserSessions struct {
	UserId     string
	OAuthToken string
	LoginDate  string
	LastSeen   string
}
