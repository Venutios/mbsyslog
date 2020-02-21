package mbsyslog

import "testing"

func TestMessageSeverity_String(t *testing.T) {
	tests := []struct {
		name string
		ms   MessageSeverity
		want string
	}{
		{"MessageSeverityEmergency", MessageSeverityEmergency, "MessageSeverityEmergency"},
		{"MessageSeverityAlert", MessageSeverityAlert, "MessageSeverityAlert"},
		{"MessageSeverityCritical", MessageSeverityCritical, "MessageSeverityCritical"},
		{"MessageSeverityError", MessageSeverityError, "MessageSeverityError"},
		{"MessageSeverityWarning", MessageSeverityWarning, "MessageSeverityWarning"},
		{"MessageSeverityNotice", MessageSeverityNotice, "MessageSeverityNotice"},
		{"MessageSeverityInformational", MessageSeverityInformational, "MessageSeverityInformational"},
		{"MessageSeverityDebug", MessageSeverityDebug, "MessageSeverityDebug"},
		{"MessageSeverityUnknown", 392, "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ms.String(); got != tt.want {
				t.Errorf("MessageSeverity.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
