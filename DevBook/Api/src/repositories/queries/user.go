package queries

const (
	UserCreate  = "INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)"
	UserList    = "SELECT id, name, nick, email, password, createAt FROM users WHERE name LIKE ? OR nick LIKE ?"
	UserGetById = "SELECT id, name, nick, email, password, createAt FROM users WHERE id=?"
)
