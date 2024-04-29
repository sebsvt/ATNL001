package repositories

type User struct {
	ID             int    `db:"id"`
	FirstName      string `db:"firstname"`
	LastName       string `db:"lastname"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	CreatedAt      string `db:"created_at"`
}

type UserRepository interface {
	FromEmail(email string) (*User, error)
	FromID(id int) (*User, error)
	CreateNewUser(newUser User) (*User, error)
}
