package dummy

type Dummy struct {
	last uint32
}

func (g *Dummy) Gen() uint32 {
	g.last++
	return g.last
}

func New(seed uint32) *Dummy {
	return &Dummy{
		last: seed,
	}
}
