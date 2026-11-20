package categorycontroller

import (
	questions "go-web/entitas"
	"go-web/models/question_model"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	currentPage, err := strconv.Atoi(pageStr)
	if err != nil || currentPage < 1 {
		currentPage = 1
	}

	search := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("search")))
	difficulty := r.URL.Query().Get("difficulty")

	limit := 5
	allQuestions := question_model.GetAll()

	// Filter search + difficulty (in-memory, pagination sudah in-memory juga)
	if search != "" || difficulty != "" {
		var filtered []questions.Questions
		for _, q := range allQuestions {
			if difficulty != "" && !strings.EqualFold(q.Difficulty, difficulty) {
				continue
			}
			if search != "" && !strings.Contains(strings.ToLower(q.Question), search) {
				continue
			}
			filtered = append(filtered, q)
		}
		allQuestions = filtered
	}
	totalItems := len(allQuestions)

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

	pagedQuestions := allQuestions
	if totalItems > 0 {
		pagedQuestions = allQuestions[offset:end]
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
		"Search":      search,
		"Difficulty":  difficulty,
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
	idStr := r.URL.Query().Get("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if r.Method == "GET" {
		question := question_model.GetById(id)

		temp, err := template.ParseFiles("view/actions/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, question)
		return
	}

	if r.Method == "POST" {
		var q questions.Questions

		q.Id = uint(id)
		q.Question = r.FormValue("question")
		q.Option_a = r.FormValue("option_a")
		q.Option_b = r.FormValue("option_b")
		q.Option_c = r.FormValue("option_c")
		q.Option_d = r.FormValue("option_d")
		q.Correct_answer = r.FormValue("correct_answer")
		q.Difficulty = r.FormValue("difficulty")

		err := question_model.Update(q)
		if err != nil {
			http.Error(w, "Gagal mengupdate data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/actions", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("Id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = question_model.Delete(id)
	if err != nil {
		http.Error(w, "Gagal menghapus data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/actions", http.StatusSeeOther)
}
