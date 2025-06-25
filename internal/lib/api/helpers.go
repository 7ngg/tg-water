package api

type ReponseBase struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Code   int    `json:"code"`
}

func Error(msg string, code int) *ReponseBase {
	return &ReponseBase{
		Status: "Error",
		Error:  msg,
		Code:   code,
	}
}

func Ok(code int) *ReponseBase {
	return &ReponseBase{
		Status: "OK",
		Code:   code,
	}
}
