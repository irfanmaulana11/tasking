package config

type MySQLConfiguration struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	DBOptions  string
	Locale     string `default:"Asia/Jakarta"`
}
