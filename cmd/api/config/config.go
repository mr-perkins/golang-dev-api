package config

// import (
// 	"log"
// 	"os"
// )

// type Config struct {
// }

// func loadEnv() {
// 	// projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
// 	currentWorkDirectory, _ := os.Getwd()
// 	// rootPath := projectName.Find([]byte(currentWorkDirectory))

// 	err := godotenv.Load(string(currentWorkDirectory) + `/.env`)

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

// func GetConfig() *Config {
// loadEnv()

// dbHost := os.Getenv("DB_HOST")
// dbPort := os.Getenv("DB_PORT")
// dbName := os.Getenv("DB_DATABASE")
// dbUsername := os.Getenv("DB_USERNAME")
// dbPassword := os.Getenv("DB_PASSWORD")

// return &Config{
// 	DB: &DBConfig{
// 		Connection: "mysql",
// 		Host:       dbHost,
// 		Port:       dbPort,
// 		Username:   dbUsername,
// 		Password:   dbPassword,
// 		Name:       dbName,
// 		Charset:    "utf8",
// 	},
// }
// }
