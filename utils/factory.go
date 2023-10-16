package utils

import (
	"regexp"
	"runtime/debug"
	"strings"
	"time"
)

func Path(path string) string {
	compiled := regexp.MustCompile(`\{(\w+:\w+)\}`)
	compiledChild := regexp.MustCompile(`(\w+):(\w+)`)
	matches := compiled.FindAllStringSubmatch(path, -1)
	for _, match := range matches {
		if m := compiledChild.FindStringSubmatch(match[1]); len(m) > 0 {
			switch m[1] {
			case "date":
				path = strings.Replace(path, "{"+match[1]+"}", transDate(m[2]), 1)
			}
		}
	}
	return path
}

func transDate(format string) string {
	return time.Now().Format(format)
}

func Recover() {
	if err := recover(); err != nil {
		debug.PrintStack()
		return
	}
}
