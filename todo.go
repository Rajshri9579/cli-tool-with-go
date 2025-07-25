package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	completed   bool
	CreatedAt   time.Time
	CompletedAT *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		completed: true,
		CompletedAT: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error{
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	*todos = append(t[:index], t[index + 1:]...)

	return nil
}

func (todos *Todos) toggle(index int); error{
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	isCompleted := t[index].Completed

	if isCompleted {
		completionTime := time.Now()
		t[index].CompletedAT = &completionTime
	}

	t[indexd].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string); error{
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	t[indexd].Title = !isCompleted

	return nil
}
func (todos *Todos) print()