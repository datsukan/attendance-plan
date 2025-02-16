package model

// ScheduleList はスケジュールのリストを表す構造体です。
type ScheduleList []Schedule

// FilterByType はスケジュールの種類でフィルタリングします。
func (sl ScheduleList) FilterByType(t ScheduleType) ScheduleList {
	schedules := []Schedule{}
	for _, s := range sl {
		if s.Type == t {
			schedules = append(schedules, s)
		}
	}
	return schedules
}

// Sort はスケジュールを Order の昇順で並び替えます。
func (sl ScheduleList) Sort() {
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i].Order > sl[j].Order {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
}

// NextOrder は次の Order を返します。
func (sl ScheduleList) NextOrder() Order {
	if len(sl) == 0 {
		return 1
	}

	max := Order(1)
	for _, s := range sl {
		if s.Order > max {
			max = s.Order
		}
	}

	return max + 1
}

// ToDateItemList はスケジュールリストを日付ごとのリストに変換します。
// リストは日付の昇順、スケジュールの Order の昇順で並び替えられます。
func (sl ScheduleList) ToDateItemList() DateItemList {
	dataMap := make(map[ScheduleType]map[string]DateItem)
	for _, s := range sl {
		date := s.StartsAt.Format(DateFormat)

		if _, ok := dataMap[s.Type]; !ok {
			dataMap[s.Type] = make(map[string]DateItem)
		}

		dataMap[s.Type][date] = DateItem{
			Date:      s.StartsAt,
			Type:      s.Type,
			Schedules: append(dataMap[s.Type][date].Schedules, s),
		}
	}

	dateItems := DateItemList{}
	for _, data := range dataMap {
		for _, di := range data {
			di.Sort()
			dateItems = append(dateItems, di)
		}
	}

	dateItems.Sort()

	return dateItems
}
