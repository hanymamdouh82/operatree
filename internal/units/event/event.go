package event

type Event struct {
	Type         string   `yaml:"type"`
	Name         string   `yaml:"name"`
	Date         string   `yaml:"time"`
	Location     string   `yaml:"location"`
	Participants []string `yaml:"participants"`
	Tags         []string `yaml:"tags"`
	Notes        string   `yaml:"notes"`
}
