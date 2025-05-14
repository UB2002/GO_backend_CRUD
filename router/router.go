package router

import (
    "net/http"
	"patient_portal/config"
	"patient_portal/handlers"
	"patient_portal/middleware"
	"patient_portal/models"
	"patient_portal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
    r := gin.Default()
// Health check route
r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
        "message": "Patient Portal API is running",
    })
})

    // Auth
    authSvc := services.NewAuthService(db, cfg.JWTSecret)
    authH := handlers.NewAuthHandler(authSvc)
    r.POST("/login", authH.Login)

    // Patients
    patSvc := services.NewPatientService(db)
    patH := handlers.NewPatientHandler(patSvc)

    // Routes accessible by both doctors and receptionists
    bothRoles := middleware.AuthRequired(cfg.JWTSecret, models.RoleReceptionist, models.RoleDoctor)
    sharedGroup := r.Group("/patients").Use(bothRoles)
    sharedGroup.GET("/", patH.List)         // Both can list patients
    sharedGroup.GET("/:id", patH.GetByID)   // Both can view a patient
    sharedGroup.PUT("/:id", patH.Update)    // Both can update a patient

    // Routes accessible only by receptionists
    recepOnly := middleware.AuthRequired(cfg.JWTSecret, models.RoleReceptionist)
    recepGroup := r.Group("/patients").Use(recepOnly)
    recepGroup.POST("/", patH.Create)       // Only receptionists can create
    recepGroup.DELETE("/:id", patH.Delete)  // Only receptionists can delete

    return r
}