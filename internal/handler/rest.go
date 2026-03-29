package handler

import (
    "net/http"

    "github.com/DennisMRitchie/go-nlp-text-pipeline/internal/model"
    "github.com/DennisMRitchie/go-nlp-text-pipeline/internal/service"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    processor *service.Processor
}

func NewHandler(p *service.Processor) *Handler {
    return &Handler{processor: p}
}

func (h *Handler) ProcessText(c *gin.Context) {
    var req model.TextRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.processor.Process(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}

func (h *Handler) BatchProcess(c *gin.Context) {
    var req model.BatchRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.processor.BatchProcess(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}
