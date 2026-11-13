package categorycontroller

import (
	"go-web/entitas"
	"go-web/models/question_model"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	currentPage, err := strconv.Atoi(pageStr)
	if err != nil || currentPage < 1 {
		currentPage = 1
	}

	limit := 5
	questions := question_model.GetAll()
	totalItems := len(questions)

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if totalPages == 0 {
		totalPages = 1
	}

	// Clamp current page supaya tidak lewat batas
	if currentPage > totalPages {
		currentPage = totalPages
	}

	offset := (currentPage - 1) * limit
	end := offset + limit
	if end > totalItems {
		end = totalItems
	}

	pagedQuestions := questions
	if totalItems > 0 {
		pagedQuestions = questions[offset:end]
	}

	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}

	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	data := map[string]any{
		"questions":   pagedQuestions,
		"CurrentPage": currentPage,
		"TotalPages":  totalPages,
		"PrevPage":    prevPage,
		"NextPage":    nextPage,
	}

	tmpl := template.Must(template.ParseFiles("view/actions/index.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("view/actions/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var questions questions.Questions

		questions.Question = r.FormValue("question")
		questions.Option_a = r.FormValue("option_a")
		questions.Option_b = r.FormValue("option_b")
		questions.Option_c = r.FormValue("option_c")
		questions.Option_d = r.FormValue("option_d")
		questions.Correct_answer = r.FormValue("correct_answer")
		questions.Difficulty = r.FormValue("difficulty")

		id, err := question_model.Create(questions)
		if err != nil || id == 0 {
			temp, _ := template.ParseFiles("view/actions/create.html")
			temp.Execute(w, map[string]any{
				"error": "Gagal menyimpan data: " + err.Error(),
			})
			return
		}
		http.Redirect(w, r, "/actions", http.StatusSeeOther)

	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
