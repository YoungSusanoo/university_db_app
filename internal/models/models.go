package models

type Model interface {
	Student | Teacher | Subject | Mark | Group
}

type User struct {
	Name    string
	IsAdmin bool
}

type Student struct {
	Id         int64
	FirstName  string
	LastName   string
	FatherName string
	Group      string
}

type Teacher struct {
	Id         int64
	FirstName  string
	LastName   string
	FatherName string
}

type Subject struct {
	Id   int64
	Name string
}

type Mark struct {
	Id    int64
	Teach Teacher
	Stud  Student
	Subj  Subject
	Value int
}

type Group struct {
	Id   int64
	Name string
}
