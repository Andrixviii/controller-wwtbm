package question_model

import (
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


