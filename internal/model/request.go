package model

type TextRequest struct {
    Text string `json:"text" binding:"required"`
    Task string `json:"task" binding:"required,oneof=classify sentiment summarize ner"`
}

type BatchRequest struct {
    Texts []string `json:"texts" binding:"required,min=1,max=100"`
    Task  string   `json:"task" binding:"required,oneof=classify sentiment summarize ner"`
}
