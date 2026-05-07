package port

// UserUsageSubjectData はユーザー利用状況の科目データを表す構造体です。
type UserUsageSubjectData struct {
	ID        string
	Name      string
	Color     string
	CreatedAt string
	UpdatedAt string
}

// UserUsageData はユーザー1件の利用状況データを表す構造体です。
type UserUsageData struct {
	ID           string
	Email        string
	Name         string
	RegisteredAt string
	LastUsedAt   string
	Subjects     []UserUsageSubjectData
}

// GetUserUsageListInputData はユーザー利用状況リスト取得の入力データを表す構造体です。
type GetUserUsageListInputData struct {
	RequesterUserID string
}

// GetUserUsageListOutputData はユーザー利用状況リスト取得の出力データを表す構造体です。
type GetUserUsageListOutputData struct {
	Total int
	Users []UserUsageData
}

// UserUsageInputPort はユーザー利用状況ユースケースを表すインターフェースです。
type UserUsageInputPort interface {
	GetUserUsageList(inputData GetUserUsageListInputData)
}

// UserUsageOutputPort はユーザー利用状況ユースケースの外部出力を表すインターフェースです。
type UserUsageOutputPort interface {
	GetResponse() (int, string)
	SetResponseGetUserUsageList(outputData *GetUserUsageListOutputData, result Result)
}
