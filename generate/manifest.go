package main

type Manifest struct {
	Actions []struct {
		Icon   string `json:"Icon"`
		Name   string `json:"Name"`
		States []struct {
			Image string `json:"Image"`
		} `json:"States"`
		SupportedInMultiActions bool   `json:"SupportedInMultiActions"`
		Tooltip                 string `json:"Tooltip"`
		UUID                    string `json:"UUID"`
		PropertyInspectorPath   string `json:"PropertyInspectorPath"`
	} `json:"Actions"`
	SDKVersion   int    `json:"SDKVersion"`
	Author       string `json:"Author"`
	Category     string `json:"Category"`
	CategoryIcon string `json:"CategoryIcon"`
	CodePath     string `json:"CodePath"`
	CodePathMac  string `json:"CodePathMac"`
	CodePathWin  string `json:"CodePathWin"`
	Description  string `json:"Description"`
	Name         string `json:"Name"`
	Icon         string `json:"Icon"`
	Version      string `json:"Version"`
	OS           []struct {
		Platform       string `json:"Platform"`
		MinimumVersion string `json:"MinimumVersion"`
	} `json:"OS"`
	Software struct {
		MinimumVersion string `json:"MinimumVersion"`
	} `json:"Software"`
}
