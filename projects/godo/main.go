package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)


type GoDoTask struct{
    ID int
    Title string
    Description string
    Status string
    CreatedAt string
    UpdateAt string

}

type GoDoList []GoDoTask

var usage = "Usage godo [options] [parameters]"
var godos = "godos.json"


func main()  {

    if len(os.Args) < 2 {
        fmt.Println(usage)
        return
    }

    cliArgs := os.Args[1:]
    switch cliArgs[0] {
    case "add":

        fmt.Println("==>> Add new task <<==")
        taskTitle := cliArgs[1]
        taskDesc := strings.Join(cliArgs[2:], " ")
        addTask(taskTitle, taskDesc)

    case "update":
        fmt.Println("Update task")

    case "delete":
        fmt.Println("Delete task")

    case "mark":
        fmt.Println("mark task")

    case "list":

        switch cliArgs[1] {
        case "all":
            fmt.Println("==>> List all tasks <<==")

        case "done":
            fmt.Println("List all done tasks")

        case "todo":
            fmt.Println("List all todo tasks")

        case "in-progress":
            fmt.Println("List all on going tasks")
        }

    default:
        // Print usage
        fmt.Println(usage)

    }



    
    // taskList = append(taskList, GoDoTask{ Title: "First task", Description: "One, Two, Three", })
    // taskList = append(taskList, GoDoTask{ Title: "Second task", Description: "One, Two, Three", })
    // // fmt.Println(taskList)

    // // List tasks
    // taskList.listTasks()
    //
    // // Delete last task
    // taskList.delTask(2)
    //
    // // Complete first task
    // taskList.markComplete(0)
    //
    // // List tasks
    // taskList.listTasks()
    //
    // // Convert to JSON
    // taskList.toJSON()
}


func pError(e error) {
    log.Fatal(e)
}

// Adds new task to tasks file.
func addTask(title, desc string) {

    rb, err := os.ReadFile(godos)
    if err != nil {
        pError(err)
    }

    id := 0
    var l GoDoList
    if len(rb) != 0 {
        json.Unmarshal(rb, &l)
        id = len(l)
    }

    newTask := GoDoTask {
        ID : id,
        Title : title,
        Description: desc,
        Status : "godo",
        CreatedAt: time.Now().Format(time.UnixDate),
    }

    l = append(l, newTask)
    mb, err := json.Marshal(l)
    if err != nil {
        pError(err)
    }
    os.WriteFile(godos, mb, 0664)
}

// List tasks
func (lst *GoDoList) listTasks()  {
    fmt.Println("#", "\tTitle",  "\t\tDescription", "\tComplete")
    // for i, t := range *lst {
        // fmt.Print(i+1, ".\t", t.Title, "\t", t.Description, "\t", t.Completed, "\n")
    // }
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
    // l[i].Completed = true
    *lst = l
}

