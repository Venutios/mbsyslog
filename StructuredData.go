package mbsyslog

//StructuredData is an optional part of the syslog message that holds a
//sequence of elements, and each element is made up of multiple parameters.
//Example with two elements, and different number of parameters:
//
//[elementID param1="value1" param2="value2"][anotherElementID param1="value1"]
type StructuredData struct {
	elements []*Element
}

//NewStructuredData creates and initializes a new structured data object to
//hold syslog data
func NewStructuredData() *StructuredData {
	result := new(StructuredData)
	result.elements = make([]*Element, 0)
	return result
}

func (sd StructuredData) addElement(raw string) bool {
	e, err := NewElement(raw)
	if err != nil {
		sd.elements = append(sd.elements, e)
		return true
	}

	return false
}

//Count returns the number of elements in the structured data
func (sd StructuredData) Count() int {
	return len(sd.elements)
}

//Element returns an element from the structured data
func (sd StructuredData) Element(index int) *Element {
	return sd.elements[index]
}
