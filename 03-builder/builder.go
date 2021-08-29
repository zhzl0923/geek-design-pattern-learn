package builder

const (
	defaultMaxTotal = 8
	defaultMaxIdle  = 8
	defaultMinIdle  = 0
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewResourcePoolConfig(builder *Builder) *ResourcePoolConfig {
	return &ResourcePoolConfig{
		name:     builder.name,
		maxTotal: builder.maxTotal,
		maxIdle:  builder.maxIdle,
		minIdle:  builder.minIdle,
	}
}

type Builder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (b *Builder) Build() *ResourcePoolConfig {
	if b.name == "" {
		panic("name is empty")
	}

	if b.maxTotal <= 0 {
		b.maxTotal = defaultMaxTotal
	}

	if b.maxIdle <= 0 {
		b.maxIdle = defaultMaxIdle
	}

	if b.minIdle <= 0 {
		b.minIdle = defaultMinIdle
	}

	if b.maxIdle > b.maxTotal {
		panic("max idle cannot > max total")
	}
	if b.minIdle > b.maxTotal || b.minIdle > b.maxIdle {
		panic("min idle must < max total max idle")
	}
	return NewResourcePoolConfig(b)
}

func (b *Builder) SetName(name string) *Builder {
	if name == "" {
		panic("name is empty")
	}
	b.name = name
	return b
}

func (b *Builder) SetMaxTotal(maxTotal int) *Builder {
	if maxTotal <= 0 {
		panic("max total must > 0")
	}
	b.maxTotal = maxTotal
	return b
}

func (b *Builder) SetMaxIdle(maxIdle int) *Builder {
	if maxIdle < 0 {
		panic("max idle must >= 0")
	}
	b.maxIdle = maxIdle
	return b
}

func (b *Builder) SetMinIdle(minIdle int) *Builder {
	if minIdle < 0 {
		panic("min idle must >= 0")
	}
	b.minIdle = minIdle
	return b
}
