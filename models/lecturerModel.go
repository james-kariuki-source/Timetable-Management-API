package models

type Lecturer struct {
	Picture     string `json:"Picture" gorm:"type:varchar(100)"`
	LecturerId  string `json:"student_id" gorm:"type:varchar(100)"`
	Nationality string `json:"nationality"`
	ID_number   string `json:"id_number"`
	DOB         string `json:"birthday"`
	City        string `json:"city"`
	Address     string `json:"address" gorm:"type:varchar(100)"`
	Fname       string `json:"first_name"`
	Mname       string `json:"middle_name"`
	Lname       string `json:"last_name"`
	Disability  string `json:"disability"`
	Department  string `json:"department"`
	LecPass     string `json:"password"`
}
