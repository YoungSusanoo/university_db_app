package models

type User struct {
	Name    string
	IsAdmin bool
}

type Student struct {
	Id         int64
	FirstName  string
	LastName   string
	FatherName string
	GroupId    int64
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

type Marks struct {
	Id        int64
	StudentId int64
	SubjectId int64
	TeacherId int64
	Value     int
}
