package port

import "time"

// StartSessionInputDate はセッション開始の入力データを表す構造体です。
type StartSessionInputDate struct {
	UserID string
}

// StartSessionOutputData はセッション開始の出力データを表す構造体です。
type StartSessionOutputData struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
	Error     error
}

// EndSessionInputDate はセッション終了の入力データを表す構造体です。
type EndSessionInputDate struct {
	UserID string
}

// EndSessionOutputData はセッション終了の出力データを表す構造体です。
type EndSessionOutputData struct {
	Error error
}

// SessionInputPort はセッションのユースケースを表すインターフェースです。
type SessionInputPort interface {
	StartSession(input StartSessionInputDate)
	EndSession(input EndSessionInputDate)
}

// SessionOutputPort はセッションのユースケースの外部出力を表すインターフェースです。
type SessionOutputPort interface {
	SetReturnStartSession(output *StartSessionOutputData)
	SetReturnEndSession(output *EndSessionOutputData)
}
