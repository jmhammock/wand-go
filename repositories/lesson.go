package repositories

import (
	"database/sql"
	"sort"

	"github.com/jmhammock/wand-go/models"
)

type ILessonRepo interface {
	Get(id int64) (*models.Lesson, error)
	List() (*models.Lessons, error)
}

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{
		db: db,
	}
}

func (r *LessonRepo) Get(id int64) (*models.Lesson, error) {
	query := `SELECT
			l.id,
			l.title,
			l.created_at,
			q.id,
			q.text,
			q.created_at,
			o.id,
			o.option_type,
			o.text,
			o.created_at  
		FROM lessons l
		JOIN questions q ON l.id = q.lesson_id
		JOIN options o ON q.id = o.question_id
		WHERE l.id = ?
		ORDER BY q.id, o.id;`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lesson := &models.Lesson{}
	questions := make(map[int64]*models.Question)
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

		question.LessonId = lesson.Id

		if _, ok := questions[question.Id]; !ok {
			questions[question.Id] = question
		}

		questions[question.Id].Options = append(questions[question.Id].Options, option)
	}

	for _, question := range questions {
		lesson.Questions = append(lesson.Questions, question)
	}

	sort.Slice(lesson.Questions, func(i, j int) bool {
		return lesson.Questions[i].Id < lesson.Questions[j].Id
	})

	return lesson, nil
}

func (r *LessonRepo) List() (*models.Lessons, error) {
	query := `SELECT
			id,
			title,
			created_at
		FROM lessons
		ORDER BY id;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lessons := make(models.Lessons, 0)
	for rows.Next() {
		lesson := &models.Lesson{}
		if err := rows.Scan(&lesson.Id, &lesson.Title, &lesson.CreatedAt); err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	return &lessons, nil
}
