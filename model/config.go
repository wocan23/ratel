package model

type Config struct {

	Tasks []Task `yaml:"tasks"`
	Vars map[string]string `yaml:"vars"`

}