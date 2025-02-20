package main

import (
	"os"
	"context"
	"path/filepath"
)

type App struct {
	ctx context.Context
	filePath string
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context, filePath string) {
	a.ctx = ctx
	if filePath != "" {
		a.OpenFile(filePath)
	}
}

func (a *App) OpenFile(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "Arquivo não encontrado."
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "Erro ao resolver caminho do arquivo."
	}
	
	content, err := os.ReadFile(absPath)
	if err != nil {
		return "Erro ao abrir o arquivo."
	}

	a.filePath = absPath
	return string(content)
}

func (a *App) GetMarkdown() string {
	if a.filePath == "" {
		return ""
	}

	content, err := os.ReadFile(a.filePath)
	if err != nil {
		println("Erro ao ler o arquivo:", err)
		return "Erro ao abrir o arquivo."
	}

	return string(content)
}