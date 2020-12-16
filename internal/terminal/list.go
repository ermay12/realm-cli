package terminal

import (
	"fmt"
	"strings"
)

var (
	listFields = []string{logFieldMessage, logFieldData}
)

type list struct {
	message string
	data    []string
}

func newList(message string, data []interface{}) list {
	l := list{
		message: message,
		data:    make([]string, 0, len(data)),
	}
	for _, item := range data {
		l.data = append(l.data, parseValue(item))
	}
	return l
}

func (l list) Message() (string, error) {
	return fmt.Sprintf("%s\n%s\n", l.message, l.dataString()), nil
}

func (l list) Payload() ([]string, map[string]interface{}, error) {
	return listFields, map[string]interface{}{
		logFieldMessage: l.message,
		logFieldData:    l.data,
	}, nil
}

func (l list) dataString() string {
	data := make([]string, 0, len(l.data))
	for _, item := range l.data {
		data = append(data, "  "+item)
	}
	return strings.Join(data, "\n")
}
