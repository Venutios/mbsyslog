package mbsyslog

import "testing"

func TestMessageFacility_String(t *testing.T) {
	tests := []struct {
		name string
		mf   MessageFacility
		want string
	}{
		{"MessageFacilityKernel", MessageFacilityKernel, "MessageFacilityKernel"},
		{"MessageFacilityUser", MessageFacilityUser, "MessageFacilityUser"},
		{"MessageFacilityMail", MessageFacilityMail, "MessageFacilityMail"},
		{"MessageFacilitySystem", MessageFacilitySystem, "MessageFacilitySystem"},
		{"MessageFacilityAuth", MessageFacilityAuth, "MessageFacilityAuth"},
		{"MessageFacilitySyslog", MessageFacilitySyslog, "MessageFacilitySyslog"},
		{"MessageFacilityPrinter", MessageFacilityPrinter, "MessageFacilityPrinter"},
		{"MessageFacilityNews", MessageFacilityNews, "MessageFacilityNews"},
		{"MessageFacilityUUCP", MessageFacilityUUCP, "MessageFacilityUUCP"},
		{"MessageFacilityCron", MessageFacilityCron, "MessageFacilityCron"},
		{"MessageFacilitySecurity", MessageFacilitySecurity, "MessageFacilitySecurity"},
		{"MessageFacilityFTP", MessageFacilityFTP, "MessageFacilityFTP"},
		{"MessageFacilityNTP", MessageFacilityNTP, "MessageFacilityNTP"},
		{"MessageFacilityLogAudit", MessageFacilityLogAudit, "MessageFacilityLogAudit"},
		{"MessageFacilityLogAlert", MessageFacilityLogAlert, "MessageFacilityLogAlert"},
		{"MessageFacilityClock", MessageFacilityClock, "MessageFacilityClock"},
		{"MessageFacilityLocal0", MessageFacilityLocal0, "MessageFacilityLocal0"},
		{"MessageFacilityLocal1", MessageFacilityLocal1, "MessageFacilityLocal1"},
		{"MessageFacilityLocal2", MessageFacilityLocal2, "MessageFacilityLocal2"},
		{"MessageFacilityLocal3", MessageFacilityLocal3, "MessageFacilityLocal3"},
		{"MessageFacilityLocal4", MessageFacilityLocal4, "MessageFacilityLocal4"},
		{"MessageFacilityLocal5", MessageFacilityLocal5, "MessageFacilityLocal5"},
		{"MessageFacilityLocal6", MessageFacilityLocal6, "MessageFacilityLocal6"},
		{"MessageFacilityLocal7", MessageFacilityLocal7, "MessageFacilityLocal7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mf.String(); got != tt.want {
				t.Errorf("MessageFacility.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
