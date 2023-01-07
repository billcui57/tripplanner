package types

type ISite struct {
	Location IGeoCode `json:"location"`
}

// ISite builder pattern code
type ISiteBuilder struct {
	iSite *ISite
}

func NewISiteBuilder() *ISiteBuilder {
	iSite := &ISite{}
	b := &ISiteBuilder{iSite: iSite}
	return b
}

func (b *ISiteBuilder) Location(location IGeoCode) *ISiteBuilder {
	b.iSite.Location = location
	return b
}

func (b *ISiteBuilder) Build() *ISite {
	return b.iSite
}
