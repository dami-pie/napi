package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dami-pie/napi/models"
)

//TODO(will): os repositórios deveriam implementar uma interface?

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (repository UserRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO userdata (userEmail, userGroupId) VALUES(?, ?)",
	)
	if err != nil {
		return 0, fmt.Errorf("create: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}
	defer statement.Close()

	result, err := statement.Exec(user.Email, user.GroupID)
	if err != nil {
		return 0, fmt.Errorf("create: não foi possível executar a declaração SQL: [%w]", err)
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("create: não foi possível recuperar o último ID inserido: [%w]", err)
	}

	return uint64(lastInsertedID), nil
}

func (repository UserRepository) Delete(id uint64) error {
	stmt, err := repository.db.Prepare(
		"DELETE FROM userdata WHERE userID = ?",
	)

	if err != nil {
		return fmt.Errorf("delete: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return fmt.Errorf("delete: não foi possível executar a declaração SQL [%w]", err)
	}

	return nil
}

func (repository UserRepository) Get(id uint64) (models.User, error) {
	stmt, err := repository.db.Prepare(
		"SELECT userEmail, userGroupId FROM userdata WHERE userID = ?",
	)

	if err != nil {
		return models.User{}, fmt.Errorf("get: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.Email, &user.GroupID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("get: não foram encontrados usuários com o id %d: [%w]", id, err)
		}

		return models.User{}, fmt.Errorf("get: não foi possível bindar a linha da tabela no modelo: [%w]", err)
	}

	return user, nil
}

func (repository UserRepository) GetByEmail(email string) (models.User, error) {
	stmt, err := repository.db.Prepare(
		"SELECT userEmail, userGroupId FROM userdata WHERE userEmail = ?",
	)

	if err != nil {
		return models.User{}, fmt.Errorf("get: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(email).Scan(&user.Email, &user.GroupID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("get: não foram encontrados usuários com o email %s: [%w]", email, err)
		}

		return models.User{}, fmt.Errorf("get: não foi possível bindar a linha da tabela no modelo: [%w]", err)
	}

	return user, nil
}

func (repository UserRepository) GetAll() ([]models.User, error) {
	stmt, err := repository.db.Prepare(
		"SELECT userEmail, userGroupId FROM userdata",
	)

	if err != nil {
		return nil, fmt.Errorf("get all: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("get all: não foi possível consultar a tabela: [%w]", err)
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Email, &user.GroupID); err != nil {
			return nil, fmt.Errorf("get all: erro ao bindar variáveis da linha da tabela ao modelo: [%w]", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, fmt.Errorf("get all: erros encontrados após bindar linhas da tabela aos modelos: [%w]", err)
	}

	return users, nil
}

//TODO(will): será que essa função é mesmo necessária?

func (repository UserRepository) Exists(id uint64) (bool, error) {
	stmt, err := repository.db.Prepare(
		"SELECT 1 FROM userdata WHERE userID = ?",
	)

	if err != nil {
		return false, fmt.Errorf("exists: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	var exists bool
	row := stmt.QueryRow(id).Scan(&exists)

	if errors.Is(row, sql.ErrNoRows) {
		return false, nil
	}

	return exists, nil
}

// Update atualiza os dados do usuário. Essa função assume que a camada de serviço
// já fez todos os tratamentos necessários no models.User (verificar se algum dado já é usado, por exemplo).
func (repository UserRepository) Update(id uint64, newData models.User) error {
	stmt, err := repository.db.Prepare(
		"UPDATE userdata SET userEmail = ?, userGroupId = ? WHERE userID = ?",
	)

	if err != nil {
		return fmt.Errorf("update: não foi possível preparar a declaração SQL necessária: [%w]", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(newData.Email, newData.GroupID, id)

	if err != nil {
		return fmt.Errorf("update: não foi possível executar a declaração SQL: [%w]", err)
	}

	return nil
}
