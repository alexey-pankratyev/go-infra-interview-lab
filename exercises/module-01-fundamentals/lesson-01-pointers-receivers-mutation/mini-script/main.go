package main

import "fmt"

type Config struct {
	Retries int
	Labels  map[string]string
}

// TODO: choose the correct receiver type and implement Normalize.
// Rules:
// - Retries must be at least 1.
// - Labels must be initialized when nil.
// - Labels["env"] must default to "dev" when missing.
func (c Config) Normalize() {
	// TODO: implement me.
}

func main() {
	first := Config{}
	first.Normalize()
	fmt.Printf("retries=%d env=%s\n", first.Retries, first.Labels["env"])

	second := Config{
		Retries: 3,
		Labels:  map[string]string{"env": "prod"},
	}
	second.Normalize()
	fmt.Printf("retries=%d env=%s\n", second.Retries, second.Labels["env"])
}
