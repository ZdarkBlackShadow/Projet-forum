package controllers

import "html/template"

var Tmpl *template.Template

func Init() error {
	var err error
	Tmpl, err = template.ParseGlob("views/*.html")
	return err
}