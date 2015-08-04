package goingup

// AppOptions is the container for all the global application settings
type Config struct {
	Port            int
	TemplateDir     string
	ContentDir      string
	StaticAssetsDir string
	StaticAssetsURL string
}

func NewConfig() *Config {
    return &Config{
        Port:            80,
        TemplateDir:     "templates",
        ContentDir:      "content",
        StaticAssetsDir: "static/",
        StaticAssetsURL: "/static/",
    }
}