package main

import (
    "fmt"
    "strings"
    "os"
    "github.com/base2services/golang-jenkins"
    "log"
)

func help() {
    fmt.Printf("\t--jenkins-user=<username>\tJenkins API user\n")
    fmt.Printf("\t--jenkins-apitoken=<token>\tJenkins API token\n")
    fmt.Printf("\t--jenkins-url=<url>\tJenkins URL\n")
    fmt.Printf("\t--queue-check\tRuns a queue check\n")
    return
}

func queuecheck(jenauth *gojenkins.Auth, jenkinsURL string) int {
    jen := gojenkins.NewJenkins(jenauth, jenkinsURL)
    queue, err := jen.GetQueue()
    if err != nil {
        log.Print("CRITICAL: Error getting queue")
        log.Println(err)
        return 2
    }
    count := len(queue.Items)
    switch count {
        case 0:
        fmt.Printf("No files in queue.\n")
        return 0
        case 1:
        fmt.Printf("One files in queue -- keep an eye on it.\n")
        return 1
        default:
        fmt.Printf("%d files in queue, should not happen. Fix.\n", count)
        return 2
    }
    return 2
}

func main() {
    action := "help"
    jenauth := &gojenkins.Auth {
        Username: "",
        ApiToken: "",
    }
    jenkinsUrl := ""
    for _, arg := range os.Args[1:] {
        switch{
            case strings.HasPrefix(arg, "--jenkins-user="):
            jenauth.Username = strings.TrimPrefix(arg, "--jenkins-user=")
            case strings.HasPrefix(arg, "--jenkins-apitoken="):
            jenauth.ApiToken = strings.TrimPrefix(arg, "--jenkins-apitoken=")
            case strings.HasPrefix(arg, "--jenkins-url="):
            jenkinsUrl = strings.TrimPrefix(arg, "--jenkins-url=")
            case arg == "--queue-check":
            action = "queuecheck"
            default:
            fmt.Printf("I don't know what %s is\n", arg)
            os.Exit(2)
        }
    }
    switch action {
        case "help":
        help()
        os.Exit(1)
        case "queuecheck":
        exitCode := queuecheck(jenauth, jenkinsUrl)
        os.Exit(exitCode)
    }
    os.Exit(2)
}
