package vars

var Config AppConfig

type AppConfig struct {
	URL         string `env:"URL,notEmpty"`
	ProductName string `env:"PRODUCT_NAME,notEmpty"`
}
