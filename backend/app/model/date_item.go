package model

import "time"

const DateFormat = "2006-01-02"

type DateItem struct {
	Date      time.Time
	Type      ScheduleType
	Schedules ScheduleList
}

func (di DateItem) Sort() {
	di.Schedules.Sort()
}

type DateItemList []DateItem

// Sort は DateItemList を日付順にソートします。
func (dil DateItemList) Sort() {
	for i := 0; i < len(dil); i++ {
		for j := i + 1; j < len(dil); j++ {
			if dil[i].Date.After(dil[j].Date) {
				dil[i], dil[j] = dil[j], dil[i]
			}
		}
	}
}

// FilterByType は DateItemList を指定された ScheduleType でフィルタリングします。
func (dil DateItemList) FilterByType(t ScheduleType) DateItemList {
	var res DateItemList
	for _, di := range dil {
		if di.Type == t {
			res = append(res, di)
		}
	}
	return res
}

// ToTypeMap は DateItemList を ScheduleType でグループ化した map を返します。
func (dil DateItemList) ToTypeMap() map[ScheduleType]DateItemList {
	m := make(map[ScheduleType]DateItemList)
	for _, di := range dil {
		m[di.Type] = append(m[di.Type], di)
	}
	return m
}
