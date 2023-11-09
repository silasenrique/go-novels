package rep

import (
	"database/sql"
	"fmt"

	"go-models/model"
)

type ISiteRepository interface {
	CreateSite(site *model.Site) error
	GetSiteByDomain(domain string) (*model.Site, error)
	GetSites(siteId *model.SiteId) ([]*model.Site, error)
}

type siteRepository struct {
	db *sql.DB
}

func NewSiteRepository(db *sql.DB) ISiteRepository {
	return &siteRepository{
		db: db,
	}
}

func (s *siteRepository) GetSites(siteId *model.SiteId) ([]*model.Site, error) {
	id := 0

	if siteId != nil {
		id = int(*siteId)
	}

	query := "select * from site s where (s.id = ? or ? = 0)"

	rows, err := s.db.Query(query, siteId)
	if err != nil {
		return nil, fmt.Errorf("nao foi possível recuperar o site com o id: %d. err: %s", id, err)
	}

	defer rows.Close()

	sites := []*model.Site{}
	for rows.Next() {
		site := &model.Site{}
		if err = rows.Scan(site); err != nil {
			return nil, fmt.Errorf("nao foi possivel ler o registro! err: %s", err)
		}

		sites = append(sites, site)

	}

	return sites, nil
}

func (s *siteRepository) GetSiteByDomain(domain string) (*model.Site, error) {
	query := "select * from site s where s.domain = ?"

	row := s.db.QueryRow(query, domain)
	if row.Err() != nil {
		return nil, fmt.Errorf("nao foi possível recuperar o dominio: %s. Error: %s", domain, row.Err())
	}

	site := &model.Site{}

	err := row.Scan(site)
	if err != nil {
		return nil, fmt.Errorf("nao foi possível ler o dominio: %s. Error: %s", domain, err)
	}

	return site, nil
}

func (s *siteRepository) CreateSite(site *model.Site) error {
	insert := "insert into site (id, name, url, creation_date) values (?, ?, ?, DATE())"

	result, err := s.db.Exec(
		insert,
		site.Id,
		site.Name,
		site.Domain,
	)

	if err != nil {
		return fmt.Errorf("nao foi possível inserir os dados do site! err: %s", err)
	}

	insertItem, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("nao foi possível checar se o site foi inserido! err: %s", err)
	}

	if insertItem == 0 {
		return fmt.Errorf("nenhum item foi inserido")
	}

	return nil
}
