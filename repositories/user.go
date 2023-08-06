package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmhammock/wand-go/models"
)

type IUserRepo interface {
	Get(id uuid.UUID) (*models.User, error)
	List() (models.Users, error)
}

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) Get(id uuid.UUID) (*models.User, error) {
	query := `SELECT
			id,
			first_name,
			last_name,
			email,
			created_at
		FROM users
		WHERE id = ?;`
	row := r.db.QueryRow(query, id)

	var user *models.User
	args := []any{
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
	}
	if err := row.Scan(args...); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) List() (models.Users, error) {
	query := `SELECT
			id,
			first_name,
			last_name,
			email,
			created_at
		FROM users;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make(models.Users, 0)
	for rows.Next() {
		var user *models.User
		args := []any{
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
		}
		if err := rows.Scan(args...); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
