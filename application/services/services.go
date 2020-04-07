package services

import (
	"markdown-version-control/application/model"
	"markdown-version-control/application/repository"
)

var repo repository.MarkdownRepository = make(repository.Datasource)

// MarkdownService - Service
var MarkdownService mardownService = dao{repo}

type mardownService interface {
	Add(content model.Markdown) (string, error)
	FindAll() []model.GitProject
	FindByID(id string) (model.Markdown, error)
	Update(id string, content model.Markdown) error
	Delete(id string) error
}

type dao struct {
	repository repository.MarkdownRepository
}

// Add - todo
func (d dao) Add(content model.Markdown) (string, error) {
	return d.repository.Save(content)
}

// FindAll - todo
func (d dao) FindAll() []model.GitProject {
	return d.repository.FindAll()
}

// FindByID - todo
func (d dao) FindByID(id string) (model.Markdown, error) {
	return d.repository.FindByID(id)
}

// Update - todo
func (d dao) Update(id string, content model.Markdown) error {
	return d.repository.Update(id, content)
}

// Delete - todo
func (d dao) Delete(id string) error {
	return d.repository.Delete(id)
}
