package query

const (
	Create  = "INSERT INTO usuario (nome, email) VALUES (?, ?)"
	List    = "SELECT id, nome, email FROM usuario"
	GetById = "SELECT id, nome, email FROM usuario WHERE id=?"
	Update  = "UPDATE usuario SET nome=?, email=? WHERE id=?"
	Delete  = "DELETE FROM usuario WHERE id=?"
)
