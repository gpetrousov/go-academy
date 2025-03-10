package main

import (
	"encoding/json"
	"fmt"
	"slices"
    "os"
)


type GoDoTask struct{
    Title string
    Description string
    Completed bool
}

type GoDoList []GoDoTask


var taskList GoDoList


func main()  {
    
    taskList = append(taskList, GoDoTask{ Title: "First task", Description: "One, Two, Three", })
    taskList = append(taskList, GoDoTask{ Title: "Second task", Description: "One, Two, Three", })
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

    // Convert to JSON
    taskList.toJSON()
}

// Add new task
func (lst *GoDoList) newTask(Title string, desc string) {
    l := *lst
    l = append(l, GoDoTask{ Title : Title, Description: desc})
    *lst = l
}

// List tasks
func (lst *GoDoList) listTasks()  {
    fmt.Println("#", "\tTitle",  "\t\tDescription", "\tComplete")
    for i, t := range *lst {
        fmt.Print(i+1, ".\t", t.Title, "\t", t.Description, "\t", t.Completed, "\n")
    }
}

// Delete task
func (lst *GoDoList) delTask(i int){
    l := *lst
    l = slices.Delete(l, i, i+1)
    *lst = l
}

// Mark task Completed
func (lst *GoDoList) markComplete(i int) {
    l := *lst
    l[i].Completed = true
    *lst = l
}

// Convert to JSON
func (lst *GoDoList) toJSON() {
    l := *lst
    b, err := json.Marshal(l)
    if err != nil {
        fmt.Println(err)
    }
    os.Stdout.Write(b)
}
