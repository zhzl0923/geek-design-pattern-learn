package decorator

type Drink interface {
	Desc() string
	Cost() int
}

type Coffee struct{}

func (c Coffee) Cost() int {
	return 1
}
func (c Coffee) Desc() string {
	return "咖啡"
}

type Milk struct {
	drink Drink
}

func (m Milk) Cost() int {
	return m.drink.Cost() + 1
}

func (m Milk) Desc() string {
	return "牛奶" + m.drink.Desc()
}
