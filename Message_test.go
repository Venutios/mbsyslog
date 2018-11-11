package mbsyslog_test

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/venutios/mbsyslog"
)

func TestMessage_Source(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want net.UDPAddr
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Source(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Source() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Valid(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want bool
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), true},
		{"SimpleInvalid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151The quick brown fox jumps over the lazy dog")), false},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), true},
		{"RFC3164Invalid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), false},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), true},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), true},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), true},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Valid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Format(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want mbsyslog.MessageFormat
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFormatSimple},
		{"SimpleInvalid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFormatInvalid},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFormatRFC3164},
		{"RFC3164Invalid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFormatInvalid},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), mbsyslog.MessageFormatRFC5424},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), mbsyslog.MessageFormatRFC5424},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), mbsyslog.MessageFormatRFC5424},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), mbsyslog.MessageFormatRFC5424},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Format(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Priority(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want int
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), 151},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), 3},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), 34},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), 165},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), 165},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), 165},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Priority(); got != tt.want {
				t.Errorf("Message.Priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Facility(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want mbsyslog.MessageFacility
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFacilityLocal2},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), mbsyslog.MessageFacilityKernel},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), mbsyslog.MessageFacilityAuth},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), mbsyslog.MessageFacilityLocal4},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), mbsyslog.MessageFacilityLocal4},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), mbsyslog.MessageFacilityLocal4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Facility(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Facility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Severity(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want mbsyslog.MessageSeverity
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), mbsyslog.MessageSeverityDebug},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), mbsyslog.MessageSeverityError},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), mbsyslog.MessageSeverityCritical},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), mbsyslog.MessageSeverityNotice},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), mbsyslog.MessageSeverityNotice},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), mbsyslog.MessageSeverityNotice},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Severity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Severity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Version(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want int
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), -1},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), -1},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), 1},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), 1},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), 1},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Version(); got != tt.want {
				t.Errorf("Message.Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Date(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want time.Time
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), time.Time{}},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), time.Date(time.Now().Year(), time.November, 10, 14, 38, 52, 0, time.UTC)},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), time.Date(2003, time.October, 11, 22, 14, 15, 3000000, time.UTC)},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), func() time.Time { date, _ := time.Parse(time.RFC3339, "2003-08-24T05:14:15.000003-07:00"); return date }()},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), time.Date(2003, time.October, 11, 22, 14, 15, 3000000, time.UTC)},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), time.Date(2003, time.October, 11, 22, 14, 15, 3000000, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Date(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Hostname(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want string
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), ""},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), "machineName"},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), "mymachine.example.com"},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), "192.0.2.1"},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), "mymachine.example.com"},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), "mymachine.example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Hostname(); got != tt.want {
				t.Errorf("Message.Hostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Application(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want string
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), ""},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), "appName"},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), "su"},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), "myproc"},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), "evntslog"},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), "evntslog"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Application(); got != tt.want {
				t.Errorf("Message.Application() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_ProcessID(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want int
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), -1},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), -1},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), -1},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), 8710},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), -1},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ProcessID(); got != tt.want {
				t.Errorf("Message.ProcessID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_MessageID(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want string
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), ""},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), ""},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), "ID47"},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), ""},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), "ID47"},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), "ID47"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MessageID(); got != tt.want {
				t.Errorf("Message.MessageID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Content(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want string
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), "The quick brown fox jumps over the lazy dog"},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), "The quick brown fox jumps over the lazy dog"},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), "'su root' failed for lonvick on /dev/pts/8"},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), "%% It's time to make the do-nuts."},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), "An application event log entry..."},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Content(); got != tt.want {
				t.Errorf("Message.Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_String(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want string
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), "127.0.0.1 " + "<151>The quick brown fox jumps over the lazy dog"},
		{"RFC3164Valid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog")), "127.0.0.1 " + "<3>Nov 10 14:38:52 machineName appName The quick brown fox jumps over the lazy dog"},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), "127.0.0.1 " + "<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), "127.0.0.1 " + "<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts."},
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), "127.0.0.1 " + "<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry..."},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), "127.0.0.1 " + "<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Message.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_StructuredDataEmpty(t *testing.T) {
	tests := []struct {
		name string
		m    mbsyslog.Message
		want mbsyslog.StructuredData
	}{
		{"SimpleValid", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<151>The quick brown fox jumps over the lazy dog")), mbsyslog.StructuredData{}},
		{"RFC5424Valid1", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8")), mbsyslog.StructuredData{}},
		{"RFC5424Valid2", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts.")), mbsyslog.StructuredData{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.StructuredData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.StructuredData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_StructuredDataPopulated(t *testing.T) {
	tests := []struct {
		name            string
		m               mbsyslog.Message
		elementCount    int
		ids             []string
		parameterCounts []int
		parameters      [][][]string
	}{
		{"RFC5424Valid3", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"] BOMAn application event log entry...")), 1, []string{"exampleSDID@32473"}, []int{3}, [][][]string{{{"iut", "eventSource", "eventID"}, {"3", "Application", "1011"}}}},
		{"RFC5424Valid4", *mbsyslog.NewMessage(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte("<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut=\"3\" eventSource=\"Application\" eventID=\"1011\"][examplePriority@32473 class=\"high\"]")), 2, []string{"exampleSDID@32473", "examplePriority@32473"}, []int{3, 1}, [][][]string{{{"iut", "eventSource", "eventID"}, {"3", "Application", "1011"}}, {{"class"}, {"high"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.StructuredData().Count(); !reflect.DeepEqual(got, tt.elementCount) {
				t.Errorf("Message.StructuredData().Count() = %d, want %d", got, tt.elementCount)
			}

			for elementIndex := 0; elementIndex < tt.elementCount; elementIndex++ {
				if got := tt.m.StructuredData().Element(elementIndex).ID(); !reflect.DeepEqual(got, tt.ids[elementIndex]) {
					t.Errorf("Message.StructuredData().Element().ID() = %s, want %s", got, tt.ids[elementIndex])
				}
				if got := tt.m.StructuredData().Element(elementIndex).Count(); !reflect.DeepEqual(got, tt.parameterCounts[elementIndex]) {
					t.Errorf("Message.StructuredData().Element().Count() = %d, want %d", got, tt.parameterCounts[elementIndex])
				}
				for parameterIndex := 0; parameterIndex < tt.m.StructuredData().Element(elementIndex).Count(); parameterIndex++ {
					if got := tt.m.StructuredData().Element(elementIndex).Parameter(parameterIndex).Name(); !reflect.DeepEqual(got, tt.parameters[elementIndex][0][parameterIndex]) {
						t.Errorf("Message.StructuredData().Element().Parameter().Name() = %s, want %s", got, tt.parameters[elementIndex][0][parameterIndex])
					}
					if got := tt.m.StructuredData().Element(elementIndex).Parameter(parameterIndex).Value(); !reflect.DeepEqual(got, tt.parameters[elementIndex][1][parameterIndex]) {
						t.Errorf("Message.StructuredData().Element().Parameter().Value() = %s, want %s", got, tt.parameters[elementIndex][1][parameterIndex])
					}
				}
			}
		})
	}
}
