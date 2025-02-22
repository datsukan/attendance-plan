package usecase

const (
	MsgInternalServerError    = "サーバーエラーが発生しました。再試行してください。再試行しても解決しない場合は、管理者にお問い合わせください。"
	MsgEmailOrPasswordInvalid = "メールアドレスまたはパスワードが間違っています"
	MsgEmailAlreadyExists     = "入力されたメールアドレスはすでに登録されています"
	MsgEmailNotFound          = "入力されたメールアドレスは登録されていません"
	MsgTokenInvalid           = "トークンが無効もしくは期限切れです"
	MsgScheduleNotFound       = "指定されたスケジュールは存在しません"
	MsgFormatInvalid          = "%sの形式が正しくありません"
	MsgUnauthorized           = "ログインしてください"
	MsgUserNotFound           = "ユーザーが見つかりません"
	MsgRequestFormatInvalid   = "リクエストの形式が正しくありません"
	MsgEmailIsSame            = "新しいメールアドレスが現在と同じです"
)
