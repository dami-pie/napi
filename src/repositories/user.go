package repositories

import (
	"database/sql"

	"github.com/dami-pie/napi/models"
)

// Aqui receberemos a conexão com um banco e ele irá fazer a interação com o banco em si
type users struct {
	db *sql.DB
}

// Recebe um banco e irá jogar para dentro do struct de usuário
// A conexão é aberta no service e aqui faremos a manipulação
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO userData (userEmail, userGroupId) VALUES(?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Email, user.GroupID)
	if erro != nil {
		return 0, erro
	}

	lastInsertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertedID), nil
}
