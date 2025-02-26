package request

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// GetSubjectListRequest は科目リスト取得のリクエストを表す構造体です。
type GetSubjectListRequest struct {
	UserID string `json:"-"`
}

// PostSubjectRequest は科目登録のリクエストを表す構造体です。
type PostSubjectRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// DeleteSubjectRequest は科目削除のリクエストを表す構造体です。
type DeleteSubjectRequest struct {
	SubjectID string
}

// ToGetSubjectListRequest は APIGatewayProxyRequest から GetSubjectListRequest に変換します。
func ToGetSubjectListRequest(r events.APIGatewayProxyRequest) *GetSubjectListRequest {
	return &GetSubjectListRequest{UserID: r.PathParameters["user_id"]}
}

// ValidateGetSubjectListRequest は GetSubjectListRequest のバリデーションを行います。
func ValidateGetSubjectListRequest(req *GetSubjectListRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("ユーザーIDを指定してください")
	}
	return nil
}

// ToPostSubjectRequest は APIGatewayProxyRequest から PostSubjectRequest に変換します。
func ToPostSubjectRequest(r events.APIGatewayProxyRequest) (*PostSubjectRequest, error) {
	var req PostSubjectRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// ValidatePostSubjectRequest は PostSubjectRequest のバリデーションを行います。
func ValidatePostSubjectRequest(req *PostSubjectRequest) error {
	if req.Name == "" {
		return fmt.Errorf("科目名を入力してください")
	}
	if req.Color == "" {
		return fmt.Errorf("色を指定してください")
	}
	return nil
}

// ToDeleteSubjectRequest は APIGatewayProxyRequest から DeleteSubjectRequest に変換します。
func ToDeleteSubjectRequest(r events.APIGatewayProxyRequest) *DeleteSubjectRequest {
	return &DeleteSubjectRequest{SubjectID: r.PathParameters["subject_id"]}
}

// ValidateDeleteSubjectRequest は DeleteSubjectRequest のバリデーションを行います。
func ValidateDeleteSubjectRequest(req *DeleteSubjectRequest) error {
	if req.SubjectID == "" {
		return fmt.Errorf("科目IDを指定してください")
	}
	return nil
}
