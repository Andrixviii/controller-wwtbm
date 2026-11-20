package question_model

import (
	"database/sql"
	"go-web/config"
	questions "go-web/entitas"
)

func GetAll() []questions.Questions {
	rows, err := config.DB.Query(`SELECT * FROM questionsnew`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var questionsnew []questions.Questions

	for rows.Next() {
		var q questions.Questions
		if err := rows.Scan(&q.Id, &q.Question, &q.Option_a, &q.Option_b, &q.Option_c, &q.Option_d, &q.Correct_answer, &q.Difficulty); err != nil {
			panic(err)
		}
		questionsnew = append(questionsnew, q)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return questionsnew
}

func GetById(id int) questions.Questions {
	row := config.DB.QueryRow(`SELECT * FROM questionsnew WHERE id = $1`, id)
	var q questions.Questions
	if err := row.Scan(&q.Id, &q.Question, &q.Option_a, &q.Option_b, &q.Option_c, &q.Option_d, &q.Correct_answer, &q.Difficulty); err != nil {
		panic(err)
	}
	return q
}

func Create(q questions.Questions) (int64, error) {
	var id int64
	err := config.DB.QueryRow(
		`INSERT INTO questionsnew
		(question, option_a, option_b, option_c, option_d, correct_answer, difficulty)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		q.Question, q.Option_a, q.Option_b, q.Option_c, q.Option_d, q.Correct_answer, q.Difficulty,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func Update(q questions.Questions) error {
	_, err := config.DB.Exec(
		`UPDATE questionsnew
		SET question = $1, option_a = $2, option_b = $3, option_c = $4, option_d = $5,
		    correct_answer = $6, difficulty = $7
		WHERE id = $8`,
		q.Question, q.Option_a, q.Option_b, q.Option_c, q.Option_d, q.Correct_answer, q.Difficulty, q.Id,
	)
	return err
}

func Delete(id int) error {
	result, err := config.DB.Exec(`DELETE FROM questionsnew WHERE id = $1`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

