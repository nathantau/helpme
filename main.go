package main

import (
    "fmt"
    "os"
    "github.com/nathantau/textrn"
)

var (
    client = &textrn.Client{os.Getenv("TEXTRN_CONNECT_SID"), os.Getenv("TEXTRN_USERNAME")}
    phone = os.Getenv("TEXTRN_PHONE")
)


func main() {

    if len(os.Args) != 3 {
        fmt.Println("Usage: helpme <cmd> <body>")
        fmt.Println("cmd: remind")
        return
    }

    switch os.Args[1] {
    case "remind":
        client.SendMessage(phone, os.Args[2])
    default:
        fmt.Println("Invalid command")
    }
}

