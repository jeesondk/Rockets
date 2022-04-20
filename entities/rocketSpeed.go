package entities

type RocketSpeed struct {
	Current float64
	Max     float64
}

func (r *RocketSpeed) Update(delta float64) {
	r.Current += delta

	if r.Current > r.Max {
		r.Max = r.Current
	}
}
