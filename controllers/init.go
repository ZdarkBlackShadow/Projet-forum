package controllers

import "html/template"

var Templates *template.Template

func Init() error {
	var err error
	Templates, err = template.ParseGlob("views/*.html")
	return err
}