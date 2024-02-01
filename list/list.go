package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Item struct {
	ID       int
	Text     string
	Done     bool
	position int
}

type ByPri []Item

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	sort.Sort(ByPri(items))

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func ClearItems(filename string) error {
	items := []Item{}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ToggleChecked(filename string, i int, checked bool) error {
	items, _ := ReadItems(filename)

	if i > 0 && i <= len(items) {
		items[i-1].Done = checked
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		SaveItems(filename, items)
		return nil
	} else {
		return errors.New("no item found in position")
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + " - "
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "☑"
	}
	return "☐"
}

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}

	return s[i].ID < s[j].ID
}
