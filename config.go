package main

type RouteConfiguration struct {
	Path   string `yaml:"path"`
	Layout string `yaml:"layout"`
}

type Configuration struct {
	Layout struct {
		Directories struct {
			Static   string `yaml:"static"`
			Template string `yaml:"template"`
			Cache    string `yaml:"cache"`
		} `yaml:"directories"`
	} `yaml:"layout"`
	Routes []RouteConfiguration `yaml:"routes"`
}
