package hypolaslogger

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// Test if log is wheel write.
func TestLogs(t *testing.T) {
	os.Setenv("HYPOLAS_LOGS_FILE", "test/test.log")

	fpath := os.Getenv("HYPOLAS_LOGS_FILE")
	// Remove existing test file
	os.Remove(fpath)

	log := NewLogger("")
	log.Info.Println("Info")
	log.Warn.Println("Warn")
	log.Err.Println("Err")

	// Test variable
	var tesvar os.Process
	log.VarDebug(tesvar, "tesvar")

	readFile, err := os.Open(fpath)

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
