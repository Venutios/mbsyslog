package mbsyslog

//MessageFacility are the sources of syslog messages
type MessageFacility int

const (
	//MessageFacilityKernel is for kernel messages
	MessageFacilityKernel MessageFacility = iota
	//MessageFacilityUser is for user-level messages
	MessageFacilityUser
	//MessageFacilityMail is for the mail system
	MessageFacilityMail
	//MessageFacilitySystem is for system daemons
	MessageFacilitySystem
	//MessageFacilityAuth is for security/authorization messages
	MessageFacilityAuth
	//MessageFacilitySyslog is for messages generated internally by syslogd
	MessageFacilitySyslog
	//MessageFacilityPrinter is for the line printer subsystem
	MessageFacilityPrinter
	//MessageFacilityNews if for the network news subsystem
	MessageFacilityNews
	//MessageFacilityUUCP is for the UUCP subsystem
	MessageFacilityUUCP
	//MessageFacilityCron is for the clock daemon
	MessageFacilityCron
	//MessageFacilitySecurity is for security/authorization messages
	MessageFacilitySecurity
	//MessageFacilityFTP is for the FTP daemon
	MessageFacilityFTP
	//MessageFacilityNTP is for the NTP subsystem
	MessageFacilityNTP
	//MessageFacilityLogAudit is for log audit messages
	MessageFacilityLogAudit
	//MessageFacilityLogAlert is for log alert messages
	MessageFacilityLogAlert
	//MessageFacilityClock is for the clock daemon
	MessageFacilityClock
	//MessageFacilityLocal0 is for local use 0
	MessageFacilityLocal0
	//MessageFacilityLocal1 is for local use 1
	MessageFacilityLocal1
	//MessageFacilityLocal2 is for local use 2
	MessageFacilityLocal2
	//MessageFacilityLocal3 is for local use 3
	MessageFacilityLocal3
	//MessageFacilityLocal4 is for local use 4
	MessageFacilityLocal4
	//MessageFacilityLocal5 is for local use 5
	MessageFacilityLocal5
	//MessageFacilityLocal6 is for local use 6
	MessageFacilityLocal6
	//MessageFacilityLocal7 is for local use 7
	MessageFacilityLocal7
)

func (mf MessageFacility) String() string {
	switch mf {
	case MessageFacilityKernel:
		return "MessageFacilityKernel"
	case MessageFacilityUser:
		return "MessageFacilityUser"
	case MessageFacilityMail:
		return "MessageFacilityMail"
	case MessageFacilitySystem:
		return "MessageFacilitySystem"
	case MessageFacilityAuth:
		return "MessageFacilityAuth"
	case MessageFacilitySyslog:
		return "MessageFacilitySyslog"
	case MessageFacilityPrinter:
		return "MessageFacilityPrinter"
	case MessageFacilityNews:
		return "MessageFacilityNews"
	case MessageFacilityUUCP:
		return "MessageFacilityUUCP"
	case MessageFacilityCron:
		return "MessageFacilityCron"
	case MessageFacilitySecurity:
		return "MessageFacilitySecurity"
	case MessageFacilityFTP:
		return "MessageFacilityFTP"
	case MessageFacilityNTP:
		return "MessageFacilityNTP"
	case MessageFacilityLogAudit:
		return "MessageFacilityLogAudit"
	case MessageFacilityLogAlert:
		return "MessageFacilityLogAlert"
	case MessageFacilityClock:
		return "MessageFacilityClock"
	case MessageFacilityLocal0:
		return "MessageFacilityLocal0"
	case MessageFacilityLocal1:
		return "MessageFacilityLocal1"
	case MessageFacilityLocal2:
		return "MessageFacilityLocal2"
	case MessageFacilityLocal3:
		return "MessageFacilityLocal3"
	case MessageFacilityLocal4:
		return "MessageFacilityLocal4"
	case MessageFacilityLocal5:
		return "MessageFacilityLocal5"
	case MessageFacilityLocal6:
		return "MessageFacilityLocal6"
	case MessageFacilityLocal7:
		return "MessageFacilityLocal7"
	default:
		return "Unknown"
	}
}
