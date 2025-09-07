package main

import "fmt"

type BaseTasks interface {
	RunTask()
}

//type TaskConstructor func(message map[string]interface{}) (BaseTasks, error)

type TestTask struct {
	X int
}

func (t *TestTask) RunTask() {
	fmt.Println("This is TestTask")
}

func NewTestTask(message map[string]interface{}) (BaseTasks, error) {
	return &TestTask{X: 1}, nil
}

type ASDFTask struct {
	X int
	Y int
}

func (t *ASDFTask) RunTask() {
	fmt.Println("This is ASDFTask")
}

func NewASDFTask(message map[string]interface{}) (BaseTasks, error) {
	return &ASDFTask{X: 1}, nil
}

func main() {
	m := map[string]func(message map[string]interface{}) (BaseTasks, error){}

	m["test"] = NewTestTask

	m["qwer"] = NewTestTask

	fmt.Println(m)
}
