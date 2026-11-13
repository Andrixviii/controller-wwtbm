package questions

type Questions struct {
	Id             uint
	Question       string
	Option_a       string
	Option_b       string
	Option_c       string
	Option_d       string
	Correct_answer string
	Difficulty     string
}

type PageData struct {
	Questions   []Questions 
	CurrentPage int         
	TotalPages  int         
	PrevPage    int         
	NextPage    int        
}
