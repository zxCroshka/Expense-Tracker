package expenses

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func CreateIfNotExists(name string) error {
	if flag, err := exists(name); err == nil {
		if !flag {
			file, err := os.Create(name)
			if err != nil {
				log.Fatal(err)
			}
			if err := file.Close(); err != nil {
				log.Fatal(err)
			}
		}

	}
	return nil
}

func JSONtoStruct(res any) error {
	file, err := os.Open("expenses.json")
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewDecoder(file).Decode(&res); err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return nil
}
func StructToJSON(res any) error {
	file, err := os.OpenFile("expenses.json", os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = json.NewEncoder(file).Encode(res); err != nil {
		log.Fatal(err)
	}
	return nil
}