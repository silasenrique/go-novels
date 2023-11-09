package model

import (
	"fmt"
	"time"
)

type Site struct {
	Id           SiteId `json:"id"`
	Name         string `json:"name"`
	Domain       string `json:"domain"`
	CreationDate string `json:"creationDate"`
}

func NewSite(name, domain string) *Site {
	return &Site{
		Id:           siteId.New(),
		Name:         name,
		Domain:       domain,
		CreationDate: time.Now().Local().String(),
	}
}

func (s *Site) Validate() (*Site, error) {
	var listErr []error

	errDomainIsNull := s.DomainIsNull()
	if errDomainIsNull != nil {
		listErr = append(listErr, errDomainIsNull)
	}

	errNameIsNull := s.NameIsNull()
	if errNameIsNull != nil {
		listErr = append(listErr, errNameIsNull)
	}

	if len(listErr) > 0 {

	}

	return s, nil
}

func (s *Site) NameIsNull() error {
	if s.Name == "" {
		return fmt.Errorf("o nome não pode ser nulo! informe o nome do site")
	}

	return nil
}

func (s *Site) DomainIsNull() error {
	if s.Domain == "" {
		return fmt.Errorf("o dominio nao pode ser nulo! informe o nome do dominio")
	}

	return nil
}
