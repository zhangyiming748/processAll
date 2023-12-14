package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func RWLock() {
	filePath := "test.txt"
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		writeFile(filePath, "Hello, World!")
	}()

	go func() {
		defer wg.Done()
		readFile(filePath)
	}()

	wg.Wait()
}

func writeFile(filePath string, data string) {
	rwMutex := new(sync.RWMutex)
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		rwMutex.Lock()
		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		rwMutex.Unlock()
	} else {
		fmt.Println("File already exists")
	}
}

func readFile(filePath string) {
	rwMutex := new(sync.RWMutex)
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	} else if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rwMutex.RLock()
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println(string(data))
	rwMutex.RUnlock()
}
