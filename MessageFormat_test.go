package mbsyslog

import "testing"

func TestMessageFormat_String(t *testing.T) {
	tests := []struct {
		name string
		mf   MessageFormat
		want string
	}{
		{"MessageFormatUnknown", MessageFormatUnknown, "MessageFormatUnknown"},
		{"MessageFormatRFC3164", MessageFormatUnknown, "MessageFormatUnknown"},
		{"MessageFormatRFC5424", MessageFormatUnknown, "MessageFormatUnknown"},
		{"MessageFormatSimple", MessageFormatUnknown, "MessageFormatUnknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mf.String(); got != tt.want {
				t.Errorf("MessageFormat.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
