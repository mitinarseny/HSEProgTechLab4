package custom

// linear congruent PRNG
type Custom struct {
	mod  uint32
	k    uint32
	b    uint32
	last uint32
}

// next = (k * prev + b) mod M
func (g *Custom) Gen() uint32 {
	g.last = (g.k*g.last + g.b) % g.mod
	return g.last
}

func New(mod, k, b, init uint32) *Custom {
	return &Custom{
		mod:  mod,
		k:    k,
		b:    b,
		last: init,
	}
}
