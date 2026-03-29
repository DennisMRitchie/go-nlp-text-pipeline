package model

// TextRequest — запрос на обработку одного текста
type TextRequest struct {
	Text string `json:"text" binding:"required,min=1,max=5000"`
	Task string `json:"task" binding:"required,oneof=classify sentiment summarize ner"`
}

// BatchRequest — пакетная обработка
type BatchRequest struct {
	Texts []string `json:"texts" binding:"required,min=1,max=100,dive,min=1,max=5000"`
	Task  string   `json:"task" binding:"required,oneof=classify sentiment summarize ner"`
}
