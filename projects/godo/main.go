package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)


type GoDoTask struct{
    ID int
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
        taskDesc := strings.Join(cliArgs[1:], " ")
        addTask(taskDesc)

    case "update":

        fmt.Println("==>> Update task <<==")
        taskId, err := strconv.Atoi(cliArgs[1])
        if err != nil {
            pError(err)
        }
        taskDesc := strings.Join(cliArgs[1:], " ")
        updateTask(taskId, taskDesc)

    case "delete":
        fmt.Println("==>> Delete task <<==")
        taskId, err := strconv.Atoi(cliArgs[1])
        pError(err)
        delTask(taskId)

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
    if e != nil {
        log.Fatal(e)
    }
}

// Adds new task to tasks file.
func addTask(desc string) {

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

// Updates an existing task at index
func updateTask(tId int, tDesc string) {
    rb, err := os.ReadFile(godos)
    pError(err)
    var l GoDoList
    err = json.Unmarshal(rb, &l)
    pError(err)
    l[tId].Description = tDesc
    l[tId].UpdateAt = time.Now().Format(time.UnixDate)
    mb, err := json.Marshal(l)
    pError(err)
    os.WriteFile(godos, mb, 0644)
}

// Delete task at index
func delTask(tIndex int){
    var l GoDoList
    rb, err := os.ReadFile(godos)
    pError(err)
    err = json.Unmarshal(rb, &l)
    pError(err)
    l = slices.Delete(l, tIndex, tIndex+1)
    mb, err := json.Marshal(l)
    pError(err)
    os.WriteFile(godos, mb, 0664)
}


// List tasks
func (lst *GoDoList) listTasks()  {
    fmt.Println("#", "\tTitle",  "\t\tDescription", "\tComplete")
    // for i, t := range *lst {
        // fmt.Print(i+1, ".\t", t.Title, "\t", t.Description, "\t", t.Completed, "\n")
    // }
}

// Mark task Completed
func (lst *GoDoList) markComplete(i int) {
    l := *lst
    // l[i].Completed = true
    *lst = l
}

