package mbsyslog

//MessageSeverity are the severity levels of syslog messages
type MessageSeverity int

const (
	//MessageSeverityEmergency indicates the system is unusable
	MessageSeverityEmergency MessageSeverity = iota
	//MessageSeverityAlert indicates action must be taken immediately
	MessageSeverityAlert
	//MessageSeverityCritical indicates critical conditions
	MessageSeverityCritical
	//MessageSeverityError indicates error conditions
	MessageSeverityError
	//MessageSeverityWarning indicates warning conditions
	MessageSeverityWarning
	//MessageSeverityNotice indicates normal but significant condition
	MessageSeverityNotice
	//MessageSeverityInformational indicates informational messages
	MessageSeverityInformational
	//MessageSeverityDebug indicates debug-level messages
	MessageSeverityDebug
)

//String returns the string representation of the MessageSeverity
func (ms MessageSeverity) String() string {
	switch ms {
	case MessageSeverityEmergency:
		return "MessageSeverityEmergency"
	case MessageSeverityAlert:
		return "MessageSeverityAlert"
	case MessageSeverityCritical:
		return "MessageSeverityCritical"
	case MessageSeverityError:
		return "MessageSeverityError"
	case MessageSeverityWarning:
		return "MessageSeverityWarning"
	case MessageSeverityNotice:
		return "MessageSeverityNotice"
	case MessageSeverityInformational:
		return "MessageSeverityInformational"
	case MessageSeverityDebug:
		return "MessageSeverityDebug"
	default:
		return "Unknown"
	}
}
