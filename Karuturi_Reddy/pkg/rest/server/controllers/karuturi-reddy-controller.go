package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/models"
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type KaruturiReddyController struct {
	karuturiReddyService *services.KaruturiReddyService
}

func NewKaruturiReddyController() (*KaruturiReddyController, error) {
	karuturiReddyService, err := services.NewKaruturiReddyService()
	if err != nil {
		return nil, err
	}
	return &KaruturiReddyController{
		karuturiReddyService: karuturiReddyService,
	}, nil
}

func (karuturiReddyController *KaruturiReddyController) CreateKaruturiReddy(context *gin.Context) {
	// validate input
	var input models.KaruturiReddy
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiReddy creation
	if _, err := karuturiReddyController.karuturiReddyService.CreateKaruturiReddy(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "KaruturiReddy created successfully"})
}

func (karuturiReddyController *KaruturiReddyController) UpdateKaruturiReddy(context *gin.Context) {
	// validate input
	var input models.KaruturiReddy
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiReddy update
	if _, err := karuturiReddyController.karuturiReddyService.UpdateKaruturiReddy(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "KaruturiReddy updated successfully"})
}

func (karuturiReddyController *KaruturiReddyController) FetchKaruturiReddy(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiReddy fetching
	karuturiReddy, err := karuturiReddyController.karuturiReddyService.GetKaruturiReddy(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, karuturiReddy)
}

func (karuturiReddyController *KaruturiReddyController) DeleteKaruturiReddy(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger karuturiReddy deletion
	if err := karuturiReddyController.karuturiReddyService.DeleteKaruturiReddy(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "KaruturiReddy deleted successfully",
	})
}

func (karuturiReddyController *KaruturiReddyController) ListKaruturiReddies(context *gin.Context) {
	// trigger all karuturiReddies fetching
	karuturiReddies, err := karuturiReddyController.karuturiReddyService.ListKaruturiReddies()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, karuturiReddies)
}

func (*KaruturiReddyController) PatchKaruturiReddy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*KaruturiReddyController) OptionsKaruturiReddy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*KaruturiReddyController) HeadKaruturiReddy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
