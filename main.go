package main

import "fmt"

// Студент

type Student struct {
	name string
}

func (s Student) Study() string {
	return fmt.Sprintf("%s готовиться к грядущим экзаменам ", s.name)
}

// Выпускник

type Graduate struct {
	name string
}

func (g Graduate) Study() string {
	return fmt.Sprintf("%s работает в WildBerries", g.name)
}

// Адаптер для выпускника

type GraduateAdapter struct {
	graduate Graduate
}

func (ga GraduateAdapter) Study() string {
	return fmt.Sprintf("%s ушел из WildBerries преподовать в университете", ga.graduate.name)
}

func main() {
	student := Student{name: "Дамир"}
	fmt.Println(student.Study())

	graduate := Graduate{name: "Руслан"}
	fmt.Println(graduate.Study())

	graduateAdapter := GraduateAdapter{graduate: graduate}
	fmt.Println(graduateAdapter.Study())
}
