package response

import "github.com/datsukan/attendance-plan/backend/app/port"

// UserUsageSubjectResponse は科目のレスポンスデータを表す構造体です。
type UserUsageSubjectResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// UserUsageResponse はユーザー利用状況のレスポンスデータを表す構造体です。
type UserUsageResponse struct {
	ID           string                     `json:"id"`
	Email        string                     `json:"email"`
	Name         string                     `json:"name"`
	RegisteredAt string                     `json:"registered_at"`
	LastUsedAt   string                     `json:"last_used_at"`
	Subjects     []UserUsageSubjectResponse `json:"subjects"`
}

// GetUserUsageListResponse はユーザー利用状況リスト取得のレスポンスを表す構造体です。
type GetUserUsageListResponse struct {
	Total int                 `json:"total"`
	Users []UserUsageResponse `json:"users"`
}

// ToGetUserUsageListResponse はユーザー利用状況リスト取得のレスポンスに変換します。
func ToGetUserUsageListResponse(output *port.GetUserUsageListOutputData) GetUserUsageListResponse {
	if output == nil {
		return GetUserUsageListResponse{Total: 0, Users: []UserUsageResponse{}}
	}

	users := make([]UserUsageResponse, 0, len(output.Users))
	for _, u := range output.Users {
		subjects := make([]UserUsageSubjectResponse, 0, len(u.Subjects))
		for _, s := range u.Subjects {
			subjects = append(subjects, UserUsageSubjectResponse{
				ID:        s.ID,
				Name:      s.Name,
				Color:     s.Color,
				CreatedAt: s.CreatedAt,
				UpdatedAt: s.UpdatedAt,
			})
		}
		users = append(users, UserUsageResponse{
			ID:           u.ID,
			Email:        u.Email,
			Name:         u.Name,
			RegisteredAt: u.RegisteredAt,
			LastUsedAt:   u.LastUsedAt,
			Subjects:     subjects,
		})
	}

	return GetUserUsageListResponse{Total: output.Total, Users: users}
}
