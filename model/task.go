package model


type Task struct{
	Name string `yaml:"name"`
	Vars map[string]string `yaml:"vars"`
	Steps []string `yaml:"steps"`
}