package mbsyslog

//Parameter is a name/value pair in syslog structured data
type Parameter struct {
	name  string
	value string
}

//NewParameter creates a new parameter for syslog structured data
func NewParameter(name, value string) *Parameter {
	p := new(Parameter)
	p.name = name
	p.value = value
	return p
}

//Name returns the parameter name
func (p Parameter) Name() string {
	return p.name
}

//Value returns the parameter value
func (p Parameter) Value() string {
	return p.value
}
