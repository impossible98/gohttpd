package env

import (
	// import built-in packages
	"os"
	// import third-party packages
	"github.com/joho/godotenv"
	// import local packages
	"github.com/impossible98/gohttpd/global"
)

func GetEnv(text string) string {
	err := godotenv.Load()
	if err != nil {
		value, ok := os.LookupEnv(text)
		if !ok {
			port := global.PORT
			return port
		}
		port := value
		return port
	}
	port := os.Getenv(text)
	return port
}
