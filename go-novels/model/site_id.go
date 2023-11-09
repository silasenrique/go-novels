package model

type SiteId int

var siteId SiteId

func (s *SiteId) New() SiteId {
	*s += 1

	return *s
}

func (s *SiteId) SetAtual(id int) {
	*s = SiteId(id)
}
