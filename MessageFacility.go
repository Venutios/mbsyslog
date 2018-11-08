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
