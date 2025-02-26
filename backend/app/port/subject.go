package port

// BaseSubjectData は科目の基本データを表す構造体です。
type BaseSubjectData struct {
	ID        string
	UserID    string
	Name      string
	Color     string
	CreatedAt string
	UpdatedAt string
}

// GetSubjectListInputData は科目リスト取得の入力データを表す構造体です。
type GetSubjectListInputData struct {
	UserID string
}

// GetSubjectListOutputData は科目リスト取得の出力データを表す構造体です。
type GetSubjectListOutputData struct {
	Subjects []BaseSubjectData
}

// CreateSubjectInputData は科目作成の入力データを表す構造体です。
type CreateSubjectInputData struct {
	UserID string
	Name   string
	Color  string
}

// CreateSubjectOutputData は科目作成の出力データを表す構造体です。
type CreateSubjectOutputData struct {
	Subject BaseSubjectData
}

// DeleteSubjectInputData は科目削除の入力データを表す構造体です。
type DeleteSubjectInputData struct {
	SubjectID string
}

// DeleteSubjectOutputData は科目削除の出力データを表す構造体です。
type DeleteSubjectOutputData struct{}

// SubjectInputPort は科目のユースケースを表すインターフェースです。
type SubjectInputPort interface {
	GetSubjectList(inputData GetSubjectListInputData)
	CreateSubject(inputData CreateSubjectInputData)
	DeleteSubject(inputData DeleteSubjectInputData)
}

// SubjectOutputPort は科目のユースケースの外部出力を表すインターフェースです。
type SubjectOutputPort interface {
	GetResponse() (int, string)
	SetResponseGetSubjectList(outputData *GetSubjectListOutputData, result Result)
	SetResponseCreateSubject(outputData *CreateSubjectOutputData, result Result)
	SetResponseDeleteSubject(outputData *DeleteSubjectOutputData, result Result)
}
