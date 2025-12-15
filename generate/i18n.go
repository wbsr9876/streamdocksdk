package main

type I18n struct {
	Description string `json:"Description"`
	Name        string `json:"Name"`
	Category    string `json:"Category"`
}

type I18nAction struct {
	Name    string `json:"Name"`
	Tooltip string `json:"Tooltip"`
}
