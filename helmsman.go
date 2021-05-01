package helmsman

type Helmsman struct{}

func New() Helmsman {
	return Helmsman{}
}

func (hm *Helmsman) LoadConfiguration() error {
	return nil
}
