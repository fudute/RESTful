package users

// User defintion
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Repo of user
type Repo struct {
	Users []User
}

// Add a user to repo
func (repo *Repo) Add(u User) {
	repo.Users = append(repo.Users, u)
}

// GetAll users info
func (repo *Repo) GetAll() []User {
	return repo.Users
}
