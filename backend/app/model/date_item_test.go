package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDateItemList_Sort(t *testing.T) {
	tests := []struct {
		name      string
		dateItems DateItemList
		want      DateItemList
	}{
		{
			name:      "0個",
			dateItems: DateItemList{},
			want:      DateItemList{},
		},
		{
			name:      "1個",
			dateItems: DateItemList{{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}},
			want:      DateItemList{{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)}},
		},
		{
			name: "2個 順番通り",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "2個 逆順",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "3個 順番通り",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "3個 逆順",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "3個 混ざった状態",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "5個 順番通り",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "5個 逆順",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "5個 混ざった状態",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "10個 順番通り",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "10個 逆順",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "10個 混ざった状態",
			dateItems: DateItemList{
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
			want: DateItemList{
				{Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 8, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 9, 0, 0, 0, 0, time.UTC)},
				{Date: time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dateItems.Sort()
			require.Equal(t, tt.want, tt.dateItems)
			assert.Equal(t, tt.want, tt.dateItems)
		})
	}
}
