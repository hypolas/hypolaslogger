package hypolaslogger

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const (
	testfile = "test/test.txt"
)

// Test if log is wheel write.
func TestLogs(t *testing.T) {
	// Remove existing test file
	os.Remove(testfile)

	log := NewLogger(testfile)
	log.Info.Println("Info")
	log.Warn.Println("Warn")
	log.Err.Println("Err")
	log.VarDebug(testfile, "testfile")

	readFile, err := os.Open(testfile)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
