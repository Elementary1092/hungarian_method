package hungarian_method

type Pair struct {
	value  int64
	weight int64
}

func NewPair(val, weight int64) *Pair {
	return &Pair{
		value: val,
		weight: weight,
	}
}

func (p *Pair) Value() int64 {
	return p.value
}

func (p *Pair) Weight() int64 {
	return p.weight
}

func (p *Pair) ChangeWeight(newWeight int64) {
	p.weight = newWeight
} 