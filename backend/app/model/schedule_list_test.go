package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleList_Sort(t *testing.T) {
	tests := []struct {
		name     string
		schedule ScheduleList
		want     ScheduleList
	}{
		{
			name:     "0個",
			schedule: ScheduleList{},
			want:     ScheduleList{},
		},
		{
			name:     "1個",
			schedule: ScheduleList{{Order: 1}},
			want:     ScheduleList{{Order: 1}},
		},
		{
			name:     "2個 順番通り",
			schedule: ScheduleList{{Order: 1}, {Order: 2}},
			want:     ScheduleList{{Order: 1}, {Order: 2}},
		},
		{
			name:     "2個 逆順",
			schedule: ScheduleList{{Order: 2}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 2}},
		},
		{
			name:     "3個 順番通り",
			schedule: ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}},
		},
		{
			name:     "3個 逆順",
			schedule: ScheduleList{{Order: 3}, {Order: 2}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}},
		},
		{
			name:     "3個 混ざった状態",
			schedule: ScheduleList{{Order: 2}, {Order: 3}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}},
		},
		{
			name:     "10個 順番通り",
			schedule: ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}},
		},
		{
			name:     "10個 逆順",
			schedule: ScheduleList{{Order: 10}, {Order: 9}, {Order: 8}, {Order: 7}, {Order: 6}, {Order: 5}, {Order: 4}, {Order: 3}, {Order: 2}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}},
		},
		{
			name:     "10個 混ざった状態",
			schedule: ScheduleList{{Order: 2}, {Order: 3}, {Order: 1}, {Order: 5}, {Order: 4}, {Order: 6}, {Order: 7}, {Order: 9}, {Order: 8}, {Order: 10}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}},
		},
		{
			name:     "10個 重複",
			schedule: ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}, {Order: 10}, {Order: 9}, {Order: 8}, {Order: 7}, {Order: 6}, {Order: 5}, {Order: 4}, {Order: 3}, {Order: 2}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 1}, {Order: 2}, {Order: 2}, {Order: 3}, {Order: 3}, {Order: 4}, {Order: 4}, {Order: 5}, {Order: 5}, {Order: 6}, {Order: 6}, {Order: 7}, {Order: 7}, {Order: 8}, {Order: 8}, {Order: 9}, {Order: 9}, {Order: 10}, {Order: 10}},
		},
		{
			name:     "50個 順番通り",
			schedule: ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}, {Order: 11}, {Order: 12}, {Order: 13}, {Order: 14}, {Order: 15}, {Order: 16}, {Order: 17}, {Order: 18}, {Order: 19}, {Order: 20}, {Order: 21}, {Order: 22}, {Order: 23}, {Order: 24}, {Order: 25}, {Order: 26}, {Order: 27}, {Order: 28}, {Order: 29}, {Order: 30}, {Order: 31}, {Order: 32}, {Order: 33}, {Order: 34}, {Order: 35}, {Order: 36}, {Order: 37}, {Order: 38}, {Order: 39}, {Order: 40}, {Order: 41}, {Order: 42}, {Order: 43}, {Order: 44}, {Order: 45}, {Order: 46}, {Order: 47}, {Order: 48}, {Order: 49}, {Order: 50}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}, {Order: 11}, {Order: 12}, {Order: 13}, {Order: 14}, {Order: 15}, {Order: 16}, {Order: 17}, {Order: 18}, {Order: 19}, {Order: 20}, {Order: 21}, {Order: 22}, {Order: 23}, {Order: 24}, {Order: 25}, {Order: 26}, {Order: 27}, {Order: 28}, {Order: 29}, {Order: 30}, {Order: 31}, {Order: 32}, {Order: 33}, {Order: 34}, {Order: 35}, {Order: 36}, {Order: 37}, {Order: 38}, {Order: 39}, {Order: 40}, {Order: 41}, {Order: 42}, {Order: 43}, {Order: 44}, {Order: 45}, {Order: 46}, {Order: 47}, {Order: 48}, {Order: 49}, {Order: 50}},
		},
		{
			name:     "50個 逆順",
			schedule: ScheduleList{{Order: 50}, {Order: 49}, {Order: 48}, {Order: 47}, {Order: 46}, {Order: 45}, {Order: 44}, {Order: 43}, {Order: 42}, {Order: 41}, {Order: 40}, {Order: 39}, {Order: 38}, {Order: 37}, {Order: 36}, {Order: 35}, {Order: 34}, {Order: 33}, {Order: 32}, {Order: 31}, {Order: 30}, {Order: 29}, {Order: 28}, {Order: 27}, {Order: 26}, {Order: 25}, {Order: 24}, {Order: 23}, {Order: 22}, {Order: 21}, {Order: 20}, {Order: 19}, {Order: 18}, {Order: 17}, {Order: 16}, {Order: 15}, {Order: 14}, {Order: 13}, {Order: 12}, {Order: 11}, {Order: 10}, {Order: 9}, {Order: 8}, {Order: 7}, {Order: 6}, {Order: 5}, {Order: 4}, {Order: 3}, {Order: 2}, {Order: 1}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}, {Order: 11}, {Order: 12}, {Order: 13}, {Order: 14}, {Order: 15}, {Order: 16}, {Order: 17}, {Order: 18}, {Order: 19}, {Order: 20}, {Order: 21}, {Order: 22}, {Order: 23}, {Order: 24}, {Order: 25}, {Order: 26}, {Order: 27}, {Order: 28}, {Order: 29}, {Order: 30}, {Order: 31}, {Order: 32}, {Order: 33}, {Order: 34}, {Order: 35}, {Order: 36}, {Order: 37}, {Order: 38}, {Order: 39}, {Order: 40}, {Order: 41}, {Order: 42}, {Order: 43}, {Order: 44}, {Order: 45}, {Order: 46}, {Order: 47}, {Order: 48}, {Order: 49}, {Order: 50}},
		},
		{
			name:     "50個 混ざった状態",
			schedule: ScheduleList{{Order: 50}, {Order: 3}, {Order: 1}, {Order: 5}, {Order: 4}, {Order: 6}, {Order: 7}, {Order: 39}, {Order: 8}, {Order: 10}, {Order: 12}, {Order: 11}, {Order: 13}, {Order: 14}, {Order: 15}, {Order: 17}, {Order: 16}, {Order: 18}, {Order: 19}, {Order: 20}, {Order: 22}, {Order: 21}, {Order: 23}, {Order: 24}, {Order: 25}, {Order: 27}, {Order: 26}, {Order: 28}, {Order: 29}, {Order: 30}, {Order: 32}, {Order: 31}, {Order: 33}, {Order: 34}, {Order: 35}, {Order: 37}, {Order: 36}, {Order: 38}, {Order: 9}, {Order: 40}, {Order: 42}, {Order: 41}, {Order: 43}, {Order: 44}, {Order: 45}, {Order: 47}, {Order: 46}, {Order: 48}, {Order: 49}, {Order: 2}},
			want:     ScheduleList{{Order: 1}, {Order: 2}, {Order: 3}, {Order: 4}, {Order: 5}, {Order: 6}, {Order: 7}, {Order: 8}, {Order: 9}, {Order: 10}, {Order: 11}, {Order: 12}, {Order: 13}, {Order: 14}, {Order: 15}, {Order: 16}, {Order: 17}, {Order: 18}, {Order: 19}, {Order: 20}, {Order: 21}, {Order: 22}, {Order: 23}, {Order: 24}, {Order: 25}, {Order: 26}, {Order: 27}, {Order: 28}, {Order: 29}, {Order: 30}, {Order: 31}, {Order: 32}, {Order: 33}, {Order: 34}, {Order: 35}, {Order: 36}, {Order: 37}, {Order: 38}, {Order: 39}, {Order: 40}, {Order: 41}, {Order: 42}, {Order: 43}, {Order: 44}, {Order: 45}, {Order: 46}, {Order: 47}, {Order: 48}, {Order: 49}, {Order: 50}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.schedule.Sort()
			require.Equal(t, len(tt.want), len(tt.schedule))
			assert.Equal(t, tt.want, tt.schedule)
		})
	}
}
