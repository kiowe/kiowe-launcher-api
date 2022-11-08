package utils

import (
	"fmt"
	"strings"
)

type Tsk struct {
	column  []string
	element []interface{}
}

func NewTsk() *Tsk {
	return &Tsk{
		column:  make([]string, 0),
		element: make([]interface{}, 0),
	}
}

func (tsk *Tsk) AddColumn(name string, value interface{}) {
	tsk.column = append(tsk.column, fmt.Sprintf("%s = $%d", name, len(tsk.column)+1))
	tsk.element = append(tsk.element, value)
}

func (tsk *Tsk) JoinColEl(idKey string, id int) string {
	tsk.element = append(tsk.element, id)
	return strings.Join(tsk.column, ", ") + fmt.Sprintf(" WHERE %s = $%d", idKey, len(tsk.column)+1)
}

func (tsk *Tsk) GetValue() []interface{} {
	return tsk.element
}
