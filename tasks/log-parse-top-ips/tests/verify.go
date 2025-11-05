package main

import (
    "bufio"
    "fmt"
    "os"
    "reflect"
)

var expectedLog = []string{
    "192.168.0.10 - - [01/Jan/2025:00:00:01 +0000] \"GET /index.html HTTP/1.1\" 200 123",
    "10.0.0.5 - - [01/Jan/2025:00:00:02 +0000] \"GET /about HTTP/1.1\" 200 231",
    "192.168.0.10 - - [01/Jan/2025:00:00:03 +0000] \"POST /api/v1/login HTTP/1.1\" 200 77",
    "172.16.0.2 - - [01/Jan/2025:00:00:04 +0000] \"GET /contact HTTP/1.1\" 404 12",
    "10.0.0.5 - - [01/Jan/2025:00:00:05 +0000] \"GET /home HTTP/1.1\" 200 198",
    "203.0.113.9 - - [01/Jan/2025:00:00:06 +0000] \"GET /index.html HTTP/1.1\" 200 123",
    "192.168.0.10 - - [01/Jan/2025:00:00:07 +0000] \"GET /docs HTTP/1.1\" 200 321",
    "10.0.0.5 - - [01/Jan/2025:00:00:08 +0000] \"POST /api/v1/order HTTP/1.1\" 201 65",
    "198.51.100.7 - - [01/Jan/2025:00:00:09 +0000] \"GET /index.html HTTP/1.1\" 200 123",
    "192.0.2.4 - - [01/Jan/2025:00:00:10 +0000] \"GET /help HTTP/1.1\" 200 54",
    "10.0.0.5 - - [01/Jan/2025:00:00:11 +0000] \"GET /pricing HTTP/1.1\" 200 87",
    "192.168.0.10 - - [01/Jan/2025:00:00:12 +0000] \"GET /blog HTTP/1.1\" 200 90",
}

var expectedTop = []string{
    "10.0.0.5 4",
    "192.168.0.10 4",
    "192.0.2.4 1",
}

func readLines(path string) ([]string, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    var lines []string
    s := bufio.NewScanner(f)
    for s.Scan() {
        lines = append(lines, s.Text())
    }
    return lines, s.Err()
}

func main() {
    
    if _, err := os.Stat("/app/web.log"); err != nil {
        fmt.Println("Missing /app/web.log:", err)
        os.Exit(1)
    }
    if _, err := os.Stat("/app/top_ips.txt"); err != nil {
        fmt.Println("Missing /app/top_ips.txt:", err)
        os.Exit(1)
    }

    logLines, err := readLines("/app/web.log")
    if err != nil {
        fmt.Println("Error reading web.log:", err)
        os.Exit(1)
    }
    if len(logLines) != 12 {
        fmt.Printf("web.log must have exactly 12 lines, got %d\n", len(logLines))
        os.Exit(1)
    }
    if !reflect.DeepEqual(logLines, expectedLog) {
        fmt.Println("web.log content does not match the required lines")
        os.Exit(1)
    }

    topLines, err := readLines("/app/top_ips.txt")
    if err != nil {
        fmt.Println("Error reading top_ips.txt:", err)
        os.Exit(1)
    }
    if len(topLines) != 3 {
        fmt.Printf("top_ips.txt must have exactly 3 lines, got %d\n", len(topLines))
        os.Exit(1)
    }
    if !reflect.DeepEqual(topLines, expectedTop) {
        fmt.Println("top_ips.txt has incorrect ranking or format")
        os.Exit(1)
    }

    fmt.Println("Go verifier passed")
}


