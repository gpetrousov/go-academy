package main

import (
	"bytes"
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
        fmt.Println("==>> Mark task <<==")
        taskFlag := cliArgs[1]
        taskId, err := strconv.Atoi(cliArgs[2])
        pError(err)
        markTask(taskId, taskFlag)

    case "list":

        fmt.Println("==>> List task <<==")
        listTasks()

    default:
        // Print usage
        fmt.Println(usage)

    }
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

// Change task flag
func  markTask(taskIndex int, taskFlag string) {
    rb, err := os.ReadFile(godos)
    pError(err)
    var l GoDoList
    err = json.Unmarshal(rb, &l)
    pError(err)
    l[taskIndex].Status = taskFlag
    l[taskIndex].UpdateAt = time.Now().Format(time.UnixDate)
    mb, err := json.Marshal(l)
    pError(err)
    os.WriteFile(godos, mb, 0644)
}

// List tasks
func listTasks()  {
    rb, err := os.ReadFile(godos)
    pError(err)
    var out bytes.Buffer
    json.Indent(&out, rb, "", "\t")
    out.WriteTo(os.Stdout)
}
