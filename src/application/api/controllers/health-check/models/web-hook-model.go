package healthCheckController

type WebHookModel struct {
	Method  string `json:"method"`
	Content any    `json:"content"`
}
