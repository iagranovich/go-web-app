package storage

import (
	"os"
	"simple-webapp/page/model"
)

func SavePage(p *model.Page) error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Data, 0600)
}

func GetPage(title string) (*model.Page, error) {
	filename := title + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &model.Page{
		Title: title,
		Data:  data,
	}, nil
}
