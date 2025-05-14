package handlers

import (
	"net/http"
	"patient_portal/models"
	"patient_portal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
    Svc *services.PatientService
}

func NewPatientHandler(s *services.PatientService) *PatientHandler {
    return &PatientHandler{s}
}

func (h *PatientHandler) Create(c *gin.Context) {
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Svc.Create(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, p)
}

func (h *PatientHandler) List(c *gin.Context) {
    list, err := h.Svc.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, list)
}
func (h *PatientHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
        return
    }

    patient, err := h.Svc.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }

    c.JSON(http.StatusOK, patient)
}


func (h *PatientHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var p models.Patient
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    p.ID = uint(id)
    if err := h.Svc.Update(&p); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, p)
}

func (h *PatientHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Svc.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
