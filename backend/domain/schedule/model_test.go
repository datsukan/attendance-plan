package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScheduleType_String(t *testing.T) {
	tests := []struct {
		name       string
		st         ScheduleType
		wantString string
	}{
		{
			name:       "master",
			st:         ScheduleTypeMaster,
			wantString: "master",
		},
		{
			name:       "custom",
			st:         ScheduleTypeCustom,
			wantString: "custom",
		},
		{
			name:       "empty",
			st:         "",
			wantString: "",
		},
		{
			name:       "unknown",
			st:         "unknown",
			wantString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.st.String()
			assert.Equal(t, tt.wantString, got)
		})
	}
}
