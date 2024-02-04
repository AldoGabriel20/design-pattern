package factory_pattern

type Bazooka struct {
	Gun
}

func newBazooka() IGun {
	return &Bazooka{
		Gun: Gun{
			name:  "Bazooka",
			power: 200,
		},
	}
}
