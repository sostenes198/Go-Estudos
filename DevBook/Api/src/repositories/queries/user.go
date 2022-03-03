package queries

const (
	UserList              = "SELECT id, name, nick, email, password, createAt FROM users WHERE name LIKE ? OR nick LIKE ?"
	UserGetById           = "SELECT id, name, nick, email, password, createAt FROM users WHERE id=?"
	UserGetToEmailByEmail = "SELECT id, password FROM users WHERE email=?"
	UserCreate            = "INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)"
	UserUpdate            = "UPDATE users set name=?, nick=?, email=? WHERE id=?"
	UserDelete            = "DELETE from users WHERE id=?"
)
