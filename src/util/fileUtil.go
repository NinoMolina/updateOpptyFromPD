package util

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/csv"
)

func CreateFile(path string, override bool) {
	// detect if file exists
	var _, err = os.Stat(path)
	var createFile = false;
	// create file if not exists
	if override && err == nil {
		DeleteFile(path)
		createFile = true
	}

	if os.IsNotExist(err) || createFile {
		var file, err = os.Create(path)
		if IsError(err) { return }
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func ReadFile(path string) {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if IsError(err) { return }
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			IsError(err)
			break
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
}

func DeleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if IsError(err) { return }

	fmt.Println("==> done deleting file")
}

func WriteFile(path string, content string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if IsError(err) { return }
	defer file.Close()

	_, err = file.WriteString(content)
	if IsError(err) { return }

	// save changes
	err = file.Sync()
	if IsError(err) { return }

	fmt.Println("==> done writing to file")
}

func AppendStringToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func ReadJsonFile(jsonFile string) []byte {
	raw, err := ioutil.ReadFile(jsonFile)
	CheckErr(err, "")
	return raw
}

func ReadCsvFile(file string) (*csv.Reader, *os.File) {
	f, err := os.Open(file)
	CheckErr(err, "")
	return csv.NewReader(f), f
}

func CloseCsvFile(file *os.File) {
	defer file.Close()
}
