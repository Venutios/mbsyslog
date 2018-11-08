package mbsyslog

import (
	"errors"
	"strings"
)

//Element is a single element in the syslog structured data, which may contain
//multiple parameters
type Element struct {
	id         string
	parameters []*Parameter
}

//NewElement initialize the elements, parsing the string, and creating all parameters
//found. If the raw string can't be parsed nil is returned with an error.
func NewElement(raw string) (*Element, error) {
	result := new(Element)
	index := strings.Index(raw, " ")

	//parse the id from the string
	if index == -1 {
		return nil, errors.New("Id not found in element")
	}
	result.id = raw[0:index]
	index++

	//Continuing parsing the next parameter until the string is consumed
	for index < len(raw) {
		//parameters take the form name="value" and separated by spaces
		equals := strings.Index(raw[index:], "=")
		openQuote := -1

		if equals != -1 {
			openQuote = strings.Index(raw[equals:], "\"")
		}
		//quotes can be escaped, search for an unescaped closing
		endQuote := findEndQuote(raw, openQuote+1)

		//if any piece wasn't found, the parameter and element is malformed
		if equals == -1 || openQuote == -1 || endQuote == -1 {
			return nil, errors.New("Element is malformed and not parseable")
		}

		//Extract the parameter, and skip the space
		result.parameters = append(result.parameters, NewParameter(raw[index:equals], raw[openQuote+1:endQuote]))
		index = endQuote + 2
	}

	return result, nil
}

func findEndQuote(raw string, index int) int {
	//continue searching while it isn't the end of the string, and the
	//current and previous chars don't equal \"
	for index != -1 && !(raw[index] == '"' && raw[index-1] == '\'') {
		index = strings.Index(raw[index+1:], "\"")
	}

	return index
}

//ID returns the id of the element
func (e Element) ID() string {
	return e.id
}

//Parameter returns a parameter from the element
func (e Element) Parameter(index int) Parameter {
	return *e.parameters[index]
}

//Count returns the number of parameters in the element
func (e Element) Count() int {
	return len(e.parameters)
}
