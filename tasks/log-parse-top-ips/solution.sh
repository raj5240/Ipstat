#!/usr/bin/env bash
set -euo pipefail
cd /app

cat > web.log <<'LOG'
192.168.0.10 - - [01/Jan/2025:00:00:01 +0000] "GET /index.html HTTP/1.1" 200 123
10.0.0.5 - - [01/Jan/2025:00:00:02 +0000] "GET /about HTTP/1.1" 200 231
192.168.0.10 - - [01/Jan/2025:00:00:03 +0000] "POST /api/v1/login HTTP/1.1" 200 77
172.16.0.2 - - [01/Jan/2025:00:00:04 +0000] "GET /contact HTTP/1.1" 404 12
10.0.0.5 - - [01/Jan/2025:00:00:05 +0000] "GET /home HTTP/1.1" 200 198
203.0.113.9 - - [01/Jan/2025:00:00:06 +0000] "GET /index.html HTTP/1.1" 200 123
192.168.0.10 - - [01/Jan/2025:00:00:07 +0000] "GET /docs HTTP/1.1" 200 321
10.0.0.5 - - [01/Jan/2025:00:00:08 +0000] "POST /api/v1/order HTTP/1.1" 201 65
198.51.100.7 - - [01/Jan/2025:00:00:09 +0000] "GET /index.html HTTP/1.1" 200 123
192.0.2.4 - - [01/Jan/2025:00:00:10 +0000] "GET /help HTTP/1.1" 200 54
10.0.0.5 - - [01/Jan/2025:00:00:11 +0000] "GET /pricing HTTP/1.1" 200 87
192.168.0.10 - - [01/Jan/2025:00:00:12 +0000] "GET /blog HTTP/1.1" 200 90
LOG

cat > main.go <<'GO'

package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

type ipCount struct {
    ip    string
    count int
}

func main() {
    f, err := os.Open("web.log")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    counts := make(map[string]int)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, " ", 2)
        if len(parts) > 0 && parts[0] != "" {
            counts[parts[0]]++
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }

    items := make([]ipCount, 0, len(counts))
    for ip, c := range counts {
        items = append(items, ipCount{ip: ip, count: c})
    }

    sort.Slice(items, func(i, j int) bool {
        if items[i].count == items[j].count {
            return items[i].ip < items[j].ip
        }
        return items[i].count > items[j].count
    })

    out, err := os.Create("top_ips.txt")
    if err != nil {
        panic(err)
    }
    defer out.Close()

    max := 3
    if len(items) < max {
        max = len(items)
    }
    for i := 0; i < max; i++ {
        fmt.Fprintf(out, "%s %d\n", items[i].ip, items[i].count)
    }
}
GO

go build -o topips main.go
./topips