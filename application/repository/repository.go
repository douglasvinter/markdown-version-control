package repository

import (
	"fmt"
	"markdown-version-control/application/model"

	uuid "github.com/satori/go.uuid"
)

// Datasource - Generic datasource
type Datasource map[string]model.Markdown

// MarkdownRepository - MarkdownRepository interface
type MarkdownRepository interface {
	FindAll() []model.GitProject
	FindByID(id string) (model.Markdown, error)
	Save(content model.Markdown) (string, error)
	Update(id string, content model.Markdown) error
	Exists(content model.Markdown) bool
	Delete(id string) error
}

// FindAll - todo
func (ds Datasource) FindAll() []model.GitProject {
	var response []model.GitProject

	for projectID, markdown := range ds {
		response = append(response, model.GitProject{projectID, markdown})
	}

	return response
}

// FindByID - todo
func (ds Datasource) FindByID(id string) (model.Markdown, error) {
	if mardown, exists := ds[id]; exists {
		return mardown, nil
	}

	return model.Markdown{}, fmt.Errorf("Markdown{id=%s} not found", id)
}

// Save - todo
func (ds Datasource) Save(content model.Markdown) (string, error) {
	if ds.Exists(content) {
		return "", fmt.Errorf("Markdown{id=%s} already exists", content.Commit)
	}

	contentID := uuid.Must(uuid.NewV4()).String()

	ds[contentID] = content

	return contentID, nil
}

// Update - todo
func (ds Datasource) Update(id string, content model.Markdown) error {
	if _, err := ds.FindByID(id); err != nil {
		return fmt.Errorf("Markdown{id=%s} already exists", content.Commit)
	}

	ds[id] = content

	return nil
}

// Exists - todo
func (ds Datasource) Exists(content model.Markdown) bool {
	for _, markdown := range ds {
		if markdown == content {
			return true
		}
	}

	return false
}

// Delete - todo
func (ds Datasource) Delete(id string) error {
	if _, err := ds.FindByID(id); err != nil {
		return fmt.Errorf("Markdown{id=%s} not found", id)
	}

	delete(ds, id)

	return nil
}
