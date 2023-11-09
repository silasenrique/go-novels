package service

import (
	"database/sql"
	"fmt"

	"go-models/model"
	rep "go-models/repository"
)

type SiteInsertServie struct {
	repository rep.ISiteRepository
}

func NewSiteInsertService(db *sql.DB) *SiteInsertServie {
	return &SiteInsertServie{
		repository: rep.NewSiteRepository(db),
	}
}

func (s *SiteInsertServie) Insert(site *model.Site) (*model.Site, error) {
	site, err := model.NewSite(site.Name, site.Domain).Validate()
	if err != nil {
		return nil, err
	}

	domainSite, err := s.repository.GetSiteByDomain(site.Domain)
	if err != nil {
		return nil, err
	}

	if domainSite != nil {
		return nil, fmt.Errorf("o dominio não pode ser cadastrado pois já existe outro site com o mesmo dominio. Site com o mesmo dominio: id %d", domainSite.Id)
	}

	err = s.repository.CreateSite(site)
	if err != nil {
		return nil, err
	}

	return site, nil
}
