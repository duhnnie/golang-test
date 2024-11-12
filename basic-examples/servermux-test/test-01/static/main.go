package static

import (
	"embed"
	"log"
)

//go:embed templates
var templatesDir embed.FS

func GetIndexPage() []byte {
	bytes, err := templatesDir.ReadFile("templates/index.html")

	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
