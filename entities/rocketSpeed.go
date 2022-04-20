package entities

type RocketSpeed struct {
	Current int
	Max     int
}

func (r *RocketSpeed) Update(delta int) {
	r.Current += delta

	if r.Current > r.Max {
		r.Max = r.Current
	}
}
