package proto

import "encoding/json"

type ESDSDKTarget = int

const (
	HardwareAndSoftware ESDSDKTarget = iota
	HardwareOnly
	SoftwareOnly
)

type MessageHeader struct {
	Event      string `json:"event"`
	Action     string `json:"action"`
	Context    string `json:"context"`
	ActionInfo json.RawMessage
}

type Info struct {
	Application struct {
		Font            string `json:"font"`
		Language        string `json:"language"`
		Platform        string `json:"platform"`
		PlatformVersion string `json:"platformVersion"`
		Version         string `json:"version"`
	} `json:"application"`
	Colors struct {
		ButtonMouseOverBackgroundColor string `json:"buttonMouseOverBackgroundColor"`
		ButtonPressedBackgroundColor   string `json:"buttonPressedBackgroundColor"`
		ButtonPressedBorderColor       string `json:"buttonPressedBorderColor"`
		ButtonPressedTextColor         string `json:"buttonPressedTextColor"`
		HighlightColor                 string `json:"highlightColor"`
	} `json:"colors"`
	DevicePixelRatio int `json:"devicePixelRatio"`
	Devices          []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Size struct {
			Columns int `json:"columns"`
			Rows    int `json:"rows"`
		} `json:"size"`
		Type int `json:"type"`
	} `json:"devices"`
	Plugin struct {
		Uuid    string `json:"uuid"`
		Version string `json:"version"`
	} `json:"plugin"`
}
