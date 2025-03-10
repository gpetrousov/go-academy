package main

import (
	"fmt"
	"slices"
)


type GoDoTask struct{
    title string
    description string
    completed bool
}

type GoDoList []GoDoTask


var taskList GoDoList


func main()  {
    
    taskList = append(taskList, GoDoTask{ title: "First task", description: "One, Two, Three", })
    taskList = append(taskList, GoDoTask{ title: "Second task", description: "One, Two, Three", })
    // fmt.Println(taskList)

    // Append new task
    taskList.newTask("New task", "Step one, Step two")

    // List tasks
    taskList.listTasks()

    // Delete last task
    taskList.delTask(2)

    // Complete first task
    taskList.markComplete(0)

    // List tasks
    taskList.listTasks()
}

// Add new task
func (lst *GoDoList) newTask(title string, desc string) {
    l := *lst
    l = append(l, GoDoTask{ title : title, description: desc})
    *lst = l
}

// List tasks
func (lst *GoDoList) listTasks()  {
    fmt.Println("#", "\tTitle",  "\t\tDescription", "\tComplete")
    for i, t := range *lst {
        fmt.Print(i+1, ".\t", t.title, "\t", t.description, "\t", t.completed, "\n")
    }
}

// Delete task
func (lst *GoDoList) delTask(i int){
    l := *lst
    l = slices.Delete(l, i, i+1)
    *lst = l
}

// Mark task completed
func (lst *GoDoList) markComplete(i int) {
    l := *lst
    l[i].completed = true
    *lst = l
}
