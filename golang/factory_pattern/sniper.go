package factory_pattern

type Sniper struct {
	Gun
}

func newSniper() IGun {
	return &Sniper{
		Gun: Gun{
			name:  "Sniper",
			power: 400,
		},
	}
}
