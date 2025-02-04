package repository

import (
	"database/sql"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models"
	"github.com/wincentrtz/gobase/utils"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		Conn: db,
	}
}

func (ur *userRepository) FetchUserById(userId int) (*models.User, error) {
	var id int
	var name string
	var email string

	query := utils.NewQueryBuilder().
		Table("users").
		Select("id,name,email").
		Where("id", "=", strconv.Itoa(userId)).
		Build()

	err := ur.Conn.QueryRow(query).Scan(
		&id,
		&name,
		&email,
	)

	if err != nil {
		panic(err)
		return nil, err
	}

	user := &models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return user, nil
}
