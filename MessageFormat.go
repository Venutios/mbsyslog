package mbsyslog

//MessageFormat are the types returned from the syslog server
type MessageFormat int

const (
	//MessageFormatInvalid is an invalid message
	MessageFormatInvalid MessageFormat = iota
	//MessageFormatRFC3164 is an RFC 3164 compliant message
	MessageFormatRFC3164
	//MessageFormatRFC5424 is an RFC 5424 compliant message
	MessageFormatRFC5424
	//MessageFormatSimple is a simple message that consists of a priority and content only
	MessageFormatSimple
)

//String returns the string representation of the MessageFormat
func (mf MessageFormat) String() string {
	switch mf {
	case MessageFormatInvalid:
		return "MessageFormatInvalid"
	case MessageFormatRFC3164:
		return "MessageFormatRFC3164"
	case MessageFormatRFC5424:
		return "MessageFormatRFC5424"
	case MessageFormatSimple:
		return "MessageFormatSimple"
	default:
		return "Unknown"
	}
}
