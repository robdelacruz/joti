package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func main() {
	content :=
		`Paragraph 1

Paragraph 2

Paragraph 3 with break
Line 2`

	// Replace line break with two spaces + line break
	re := regexp.MustCompile("(\\S)\n(\\S)")
	//content = re.ReplaceAllString(content, "$1__\n$2")
	content = re.ReplaceAllString(content, "$1  \n$2")

	fmt.Printf("content:\n%s\n", content)
}

func test_time() {
	t := time.Now()
	fmt.Printf("t: '%s'\n", isodate(t))

	t = t.Add(60 * time.Hour * 24)
	fmt.Printf("t + 60 days: '%s'\n", isodate(t))

	t = t.Add(-60 * time.Hour * 24)
	fmt.Printf("t - 60 days: '%s'\n", isodate(t))

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			newtime := <-ticker.C
			fmt.Printf("ticker: newtime: '%s'\n", isodate(newtime))
			t := newtime.Add(days_to_duration(-22))
			fmt.Printf("ticker: t-22 days: '%s'\n", isodate(t))
		}
	}()

	for {
	}
}

func test_md() {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	bs, err := os.ReadFile("test.md")

	var buf bytes.Buffer
	err = md.Convert(bs, &buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf.String())
}
