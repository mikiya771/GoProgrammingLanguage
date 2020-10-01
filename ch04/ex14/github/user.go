package github

type User struct {
	Login   string
	HTMLURL string
}

const listUsersURL = "https://api.github.com/users"

type UsersListResult struct {
	Users    []*User
	nextLink string
}
