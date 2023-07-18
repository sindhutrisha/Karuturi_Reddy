package services

import (
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/daos"
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/models"
)

type KaruturiReddyService struct {
	karuturiReddyDao *daos.KaruturiReddyDao
}

func NewKaruturiReddyService() (*KaruturiReddyService, error) {
	karuturiReddyDao, err := daos.NewKaruturiReddyDao()
	if err != nil {
		return nil, err
	}
	return &KaruturiReddyService{
		karuturiReddyDao: karuturiReddyDao,
	}, nil
}

func (karuturiReddyService *KaruturiReddyService) CreateKaruturiReddy(karuturiReddy *models.KaruturiReddy) (*models.KaruturiReddy, error) {
	return karuturiReddyService.karuturiReddyDao.CreateKaruturiReddy(karuturiReddy)
}

func (karuturiReddyService *KaruturiReddyService) UpdateKaruturiReddy(id int64, karuturiReddy *models.KaruturiReddy) (*models.KaruturiReddy, error) {
	return karuturiReddyService.karuturiReddyDao.UpdateKaruturiReddy(id, karuturiReddy)
}

func (karuturiReddyService *KaruturiReddyService) DeleteKaruturiReddy(id int64) error {
	return karuturiReddyService.karuturiReddyDao.DeleteKaruturiReddy(id)
}

func (karuturiReddyService *KaruturiReddyService) ListKaruturiReddies() ([]*models.KaruturiReddy, error) {
	return karuturiReddyService.karuturiReddyDao.ListKaruturiReddies()
}

func (karuturiReddyService *KaruturiReddyService) GetKaruturiReddy(id int64) (*models.KaruturiReddy, error) {
	return karuturiReddyService.karuturiReddyDao.GetKaruturiReddy(id)
}
