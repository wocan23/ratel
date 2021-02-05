package model

type Config struct {
	Tasks []Task            `yaml:"tasks"`
	Vars  map[string]string `yaml:"vars"`
}

type Task struct {
	Name  string            `yaml:"name"`
	Vars  map[string]string `yaml:"vars"`
	Steps []string          `yaml:"steps"`
}
