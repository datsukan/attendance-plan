package response

import "github.com/datsukan/attendance-plan/backend/app/port"

// BaseSubjectResponse は科目のレスポンスデータの基本を表す構造体です。
type BaseSubjectResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// GetSubjectListResponse は科目リスト取得のレスポンスを表す構造体です。
type GetSubjectListResponse struct {
	Subjects []BaseSubjectResponse `json:"subjects"`
}

// PostSubjectResponse は科目登録のレスポンスを表す構造体です。
type PostSubjectResponse BaseSubjectResponse

// ToGetSubjectListResponse は科目リスト取得のレスポンスに変換します。
func ToGetSubjectListResponse(output *port.GetSubjectListOutputData) GetSubjectListResponse {
	if output == nil || len(output.Subjects) == 0 {
		return GetSubjectListResponse{Subjects: []BaseSubjectResponse{}}
	}

	var subjects []BaseSubjectResponse
	for _, s := range output.Subjects {
		subjects = append(subjects, BaseSubjectResponse{
			ID:        s.ID,
			UserID:    s.UserID,
			Name:      s.Name,
			Color:     s.Color,
			CreatedAt: s.CreatedAt,
			UpdatedAt: s.UpdatedAt,
		})
	}

	return GetSubjectListResponse{Subjects: subjects}
}

// ToPostSubjectResponse は科目登録のレスポンスに変換します。
func ToPostSubjectResponse(output *port.CreateSubjectOutputData) PostSubjectResponse {
	if output == nil {
		return PostSubjectResponse{}
	}

	return PostSubjectResponse{
		ID:        output.Subject.ID,
		UserID:    output.Subject.UserID,
		Name:      output.Subject.Name,
		Color:     output.Subject.Color,
		CreatedAt: output.Subject.CreatedAt,
		UpdatedAt: output.Subject.UpdatedAt,
	}
}
