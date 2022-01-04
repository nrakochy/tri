package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if(err != nil){
		return err
	}

	writeErr := ioutil.WriteFile(filename, b, 0644)
        if(writeErr != nil){
		return err
	}

	//fmt.Println(string(b))
	return nil
}

func ReadItems(filename string)([]Item, error){
	b, err := ioutil.ReadFile(filename)
	if(err != nil){
	  return []Item{}, nil
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
	  return []Item{}, nil
	}

	return items, nil
}