package controllers

import (
	"github.com/mikaponics/mikaponics-thing/internal/models"
	"github.com/mikaponics/mikaponics-thing/internal/services"
)


type MikaponicsThingServer struct{
    DAL *models.DataAccessLayer
	IAM *services.IAMClient
}
