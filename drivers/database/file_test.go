package database

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile_InitItem(t *testing.T) {
	const filename = "testfile.txt"

	defer func() {
		os.Remove(filename)
	}()

	// Assert NewFileStore creates a file
	NewFileStore(filename)
	if _, err := os.Stat(filename); err != nil {
		t.Error("Unable to stat file")
	}
}

func TestFile_AddItem(t *testing.T) {
	const filename = "testfile.txt"
	const expectedAddItem = "{\"name\":\"Test\",\"price\":1399}\n"
	const expectedAddItemTwice = "{\"name\":\"Test\",\"price\":1399}\n{\"name\":\"Test2\",\"price\":14000}\n"
	addItemStruct := map[string]interface{}{
		"name":  "Test",
		"price": 1399,
	}
	addItemTwiceStruct := map[string]interface{}{
		"name":  "Test2",
		"price": 14000,
	}

	defer func() {
		os.Remove(filename)
	}()

	file := NewFileStore(filename)

	// Assert AddItem works
	file.AddItem(addItemStruct)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Unable to read file for comparison")
	}
	if strings.Compare(string(b), expectedAddItem) != 0 {
		t.Error("AddItem does not produce correct string")
	}

	// Assert AddItem works twice
	file.AddItem(addItemTwiceStruct)

	b, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Unable to read file for comparison")
	}
	if strings.Compare(string(b), expectedAddItemTwice) != 0 {
		t.Error("AddItem twice does not produce correct string")
	}

}
