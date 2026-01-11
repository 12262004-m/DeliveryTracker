package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
	DB   string
}

func loadConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DB:   getEnv("DB_URL", "./data/order.db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func loadTemplates(router *gin.Engine) error {
	functions := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"json": func(v interface{}) template.JS {
			b, _ := json.Marshal(v)
			return template.JS(b)
		},
	}

	tmpl, err := template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(tmpl)
	return nil
}
