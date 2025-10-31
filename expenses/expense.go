package expenses

import "log"

func (e *Expense) AddTag(newTag string) {
	e.Category = append(e.Category, newTag)
}

func (e *Expense) DeleteTag(tag string) {
	idx := -1
	for i, v := range e.Category {
		if v == tag {
			idx = i
			break
		}
	}
	if idx != -1 {
		e.Category = append(e.Category[:idx],e.Category[idx+1:]... )
	} else {
		log.Println("there is not this tag")
	}
}
