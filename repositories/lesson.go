package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmhammock/wand-go/models"
)

type ILessonRepo interface {
	Get(id uuid.UUID) (*models.Lesson, error)
	List() (*models.Lessons, error)
}

type LessonRepo struct {
	db *sql.DB
}

func (r *LessonRepo) Get(id uuid.UUID) (*models.Lesson, error) {
	query := `SELECT
			l.ld,
			l.title,
			l.created_at,
			q.id,
			q.title,
			q.created_at,
			o.id,
			o.option_type,
			o.created_at  
		FROM lessons l
		JOIN questions q ON l.id = q.lesson_id
		JOIN options o ON q.id = o.question_id
		WHERE id = ?;`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lesson := &models.Lesson{}
	for rows.Next() {
		question := &models.Question{}
		option := &models.Option{}
		args := []any{
			&lesson.Id,
			&lesson.Title,
			&lesson.CreatedAt,
			&question.Id,
			&question.Text,
			&question.CreatedAt,
			&option.Id,
			&option.OptionType,
			&option.Text,
			&option.CreatedAt,
		}

		if err := rows.Scan(args...); err != nil {
			return nil, err
		}

		question.Options = append(question.Options, option)
		lesson.Questions = append(lesson.Questions, question)
	}

	return lesson, nil
}

func (r *LessonRepo) List() (models.Lessons, error) {
	query := `SELECT
			id,
			title,
			created_at
		FROM lessons;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lessons := make(models.Lessons, 0)
	for rows.Next() {
		var lesson *models.Lesson
		if err := rows.Scan(&lesson.Id, &lesson.Title, &lesson.CreatedAt); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}
