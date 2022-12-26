package elevator

type Config struct {
	Id       string
	Floors   int
	MaxSpeed float64
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithId(id string) *Config {
	c.Id = id
	return c
}

func (c *Config) WithFloors(floors int) *Config {
	c.Floors = floors
	return c
}

func (c *Config) WithMaxSpeed(maxSpeed float64) *Config {
	c.MaxSpeed = maxSpeed
	return c
}
