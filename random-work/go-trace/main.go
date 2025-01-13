package main

import (
    "fmt"
    "os"
    "strings"
    "sync"
    "net/http"
    "net/url"
    "golang.org/x/net/html"
)


var visited = make(map[string]bool)
var mu sync.Mutex
var wg sync.WaitGroup


func isSameDomain(baseURL, link string) bool {
    base, err := url.Parse(baseURL)
    if err != nil {
        return false
    }

    parsedLink, err := url.Parse(link)
    if err != nil {
        return false
    }

    return base.Host == parsedLink.Host
}

func resolveURL(base, href string) string {
    parsedBase, err := url.Parse(base)
    if err != nil {
        return href
    }

    parsedHref, err := url.Parse(href)
    if err != nil {
        return href
    }

    return parsedBase.ResolveReference(parsedHref).String()
}

func extractLinks(baseURL string, body *html.Node) []string {
    var links []string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    absoluteURL := resolveURL(baseURL, attr.Val)
                    if strings.HasPrefix(absoluteURL, "mailto:") {
                        fmt.Printf("[email] - %s\n", absoluteURL)
                        continue
                    }

                    links = append(links, absoluteURL)
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(body)
    return links
}

func checkLink(link string, rootDomain string) {
    defer wg.Done()

    mu.Lock()
    if visited[link] {
        mu.Unlock()
        return
    }
    visited[link] = true
    mu.Unlock()

    resp, err := http.Get(link)
    if err != nil {
        fmt.Printf("Error fetching %s: %s\n", link, err)
        return
    }
    defer resp.Body.Close()

    fmt.Print("[url] - ")

    if resp.StatusCode >= 400 && resp.StatusCode < 600 {
        fmt.Printf("\033[31m[Status: %d] : Dead link found: %s\033[0m\n", resp.StatusCode, link)
    } else if resp.StatusCode >= 600 {
        fmt.Printf("\033[33m[Status: %d] : Blocked link found: %s\033[0m\n", resp.StatusCode, link)
    } else {
        fmt.Printf("[Status: %d] : Valid Link found: %s\n", resp.StatusCode, link)
    }

    if isSameDomain(rootDomain, link) && resp.StatusCode == http.StatusOK {
        doc, err := html.Parse(resp.Body)
        if err != nil {
            fmt.Printf("Error parsing HTML for %s: %s\n", link, err)
            return
        }

        for _, href := range extractLinks(link, doc) {
            wg.Add(1)
            go checkLink(href, rootDomain)
        }
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: %s <URL>\n", os.Args[0])
        return
    }

    wg.Add(1)
    go checkLink(os.Args[1], os.Args[1])
    wg.Wait()

    fmt.Println("Finished checking all links.")
}

// valid INTERNAL/EXTERNAL/OUTBIUT link found
// --ignore-errors, --success --invalid, --email, -o
//          ANSII coloring not good for output
// fix stdout concurrency formatting
// maybe add feature to say WHERE the link is?
