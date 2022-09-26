package models

type Course struct {
	Name            string
	Desc            string
	BackgroundImage string
}
type CourseTeacher struct {
	Name   string
	Desc   string
	Avatar string
	Title  string
}

type CourseInfo struct {
	CourseId      int
	ClickCount    int
	LearnCount    int
	TotalDuration int
	SectionCount  int
	Price         int
	Detail        string
	State         string
}

type CourseChapter struct {
	CourseId int
	Title    string
}

type CourseSection struct {
	CourseId  int
	ChapterId int
	Title     string
	Type      string
	Video     string
}
