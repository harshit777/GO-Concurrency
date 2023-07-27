package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	// Getting System stardard output
	stdOut := os.Stdout

	// Creating own standard output
	r, w, _ := os.Pipe()
	// adding our write to standard output
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	go printSomething("This is pie value :", &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "pie") {
		t.Errorf("Something went wrong expected to contain pie:")
	}
}
