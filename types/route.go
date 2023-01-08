package types

type IRoute struct {
	Steps []IStep    `json:"steps"`
	Path  []IGeoCode `json:"path"`
}

// IRoute builder pattern code
type IRouteBuilder struct {
	iRoute *IRoute
}

func NewIRouteBuilder() *IRouteBuilder {
	iRoute := &IRoute{}
	b := &IRouteBuilder{iRoute: iRoute}
	return b
}

func (b *IRouteBuilder) Steps(steps []IStep) *IRouteBuilder {
	b.iRoute.Steps = steps
	return b
}

func (b *IRouteBuilder) Path(path []IGeoCode) *IRouteBuilder {
	b.iRoute.Path = path
	return b
}

func (b *IRouteBuilder) Build() *IRoute {
	return b.iRoute
}

type ILeanRoute struct {
	Path []IGeoCode `json:"path"`
}

// ILeanRoute builder pattern code
type ILeanRouteBuilder struct {
	iLeanRoute *ILeanRoute
}

func NewILeanRouteBuilder() *ILeanRouteBuilder {
	iLeanRoute := &ILeanRoute{}
	b := &ILeanRouteBuilder{iLeanRoute: iLeanRoute}
	return b
}

func (b *ILeanRouteBuilder) Path(path []IGeoCode) *ILeanRouteBuilder {
	b.iLeanRoute.Path = path
	return b
}

func (b *ILeanRouteBuilder) Build() *ILeanRoute {
	return b.iLeanRoute
}
