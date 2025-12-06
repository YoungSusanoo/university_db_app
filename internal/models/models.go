package models

type User struct {
	Name    string
	IsAdmin bool
}

type Credentials struct {
	Id         int64
	FirstName  string
	LastName   string
	FatherName string
}

type Student struct {
	Id      int64
	Name    Credentials
	GroupId int
}

type Teacher struct {
	Id   int64
	Name Credentials
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
