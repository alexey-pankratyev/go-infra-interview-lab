package main

import "fmt"

type Config struct {
	Retries int
	Labels  map[string]string
}

func (c Config) SetRetries(n int) {
	c.Retries = n
}

func (c Config) SetLabel(k, v string) {
	c.Labels[k] = v
}

func update(cfg Config) {
	cfg.SetRetries(5)
	cfg.SetLabel("env", "prod")
}

func main() {
	cfg := Config{
		Retries: 3,
		Labels:  map[string]string{"team": "platform"},
	}

	update(cfg)

	fmt.Println(cfg.Retries)
	fmt.Println(cfg.Labels["env"])
}
