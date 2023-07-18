package daos

import (
	"database/sql"
	"errors"
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/daos/clients/sqls"
	"github.com/sindhutrisha/Karuturi_Reddy/karuturi_reddy/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type KaruturiReddyDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateKaruturiReddies(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS karuturiReddies(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Reddy TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewKaruturiReddyDao() (*KaruturiReddyDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateKaruturiReddies(sqlClient)
	if err != nil {
		return nil, err
	}
	return &KaruturiReddyDao{
		sqlClient,
	}, nil
}

func (karuturiReddyDao *KaruturiReddyDao) CreateKaruturiReddy(m *models.KaruturiReddy) (*models.KaruturiReddy, error) {
	insertQuery := "INSERT INTO karuturiReddies(Reddy)values(?)"
	res, err := karuturiReddyDao.sqlClient.DB.Exec(insertQuery, m.Reddy)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("karuturiReddy created")
	return m, nil
}

func (karuturiReddyDao *KaruturiReddyDao) UpdateKaruturiReddy(id int64, m *models.KaruturiReddy) (*models.KaruturiReddy, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	karuturiReddy, err := karuturiReddyDao.GetKaruturiReddy(id)
	if err != nil {
		return nil, err
	}
	if karuturiReddy == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE karuturiReddies SET Reddy = ? WHERE Id = ?"
	res, err := karuturiReddyDao.sqlClient.DB.Exec(updateQuery, m.Reddy, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("karuturiReddy updated")
	return m, nil
}

func (karuturiReddyDao *KaruturiReddyDao) DeleteKaruturiReddy(id int64) error {
	deleteQuery := "DELETE FROM karuturiReddies WHERE Id = ?"
	res, err := karuturiReddyDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("karuturiReddy deleted")
	return nil
}

func (karuturiReddyDao *KaruturiReddyDao) ListKaruturiReddies() ([]*models.KaruturiReddy, error) {
	selectQuery := "SELECT * FROM karuturiReddies"
	rows, err := karuturiReddyDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var karuturiReddies []*models.KaruturiReddy
	for rows.Next() {
		m := models.KaruturiReddy{}
		if err = rows.Scan(&m.Id, &m.Reddy); err != nil {
			return nil, err
		}
		karuturiReddies = append(karuturiReddies, &m)
	}
	if karuturiReddies == nil {
		karuturiReddies = []*models.KaruturiReddy{}
	}

	log.Debugf("karuturiReddy listed")
	return karuturiReddies, nil
}

func (karuturiReddyDao *KaruturiReddyDao) GetKaruturiReddy(id int64) (*models.KaruturiReddy, error) {
	selectQuery := "SELECT * FROM karuturiReddies WHERE Id = ?"
	row := karuturiReddyDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.KaruturiReddy{}
	if err := row.Scan(&m.Id, &m.Reddy); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("karuturiReddy retrieved")
	return &m, nil
}
