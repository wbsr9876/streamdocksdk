package proto

func NewMessage[T any]() interface{} {
	return new(T)
}

var MessageFactory = map[string]func() interface{}{
	"didReceiveSettings":            NewMessage[DidReceiveSettings],
	"didReceiveGlobalSettings":      NewMessage[DidReceiveGlobalSettings],
	"dialDown":                      NewMessage[DialDown],
	"dialUp":                        NewMessage[DialUp],
	"dialRotate":                    NewMessage[DialRotate],
	"keyDown":                       NewMessage[KeyDown],
	"keyUp":                         NewMessage[KeyUp],
	"willAppear":                    NewMessage[WillAppear],
	"willDisappear":                 NewMessage[WillDisappear],
	"titleParametersDidChange":      NewMessage[TitleParametersDidChange],
	"deviceDidConnect":              NewMessage[DeviceDidConnect],
	"deviceDidDisconnect":           NewMessage[DeviceDidDisconnect],
	"applicationDidLaunch":          NewMessage[ApplicationDidLaunch],
	"applicationDidTerminate":       NewMessage[ApplicationDidTerminate],
	"systemDidWakeUp":               NewMessage[SystemDidWakeUp],
	"propertyInspectorDidAppear":    NewMessage[PropertyInspectorDidAppear],
	"propertyInspectorDidDisappear": NewMessage[PropertyInspectorDidDisappear],
	"sendToPlugin":                  NewMessage[SendToPlugin],
	"sendToPropertyInspector":       NewMessage[SendToPropertyInspector],
}

type DidReceiveSettings struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

func (m *DidReceiveSettings) GetEvent() string {
	return m.Event
}

func (m *DidReceiveSettings) GetAction() string {
	return m.Action
}

func (m *DidReceiveSettings) GetContext() string {
	return m.Context
}

func (m *DidReceiveSettings) GetDevice() string {
	return m.Device
}

type DidReceiveGlobalSettings struct {
	Event   string `json:"event"`
	Payload struct {
		Settings interface{} `json:"settings"`
	} `json:"payload"`
}

func (m *DidReceiveGlobalSettings) GetEvent() string {
	return m.Event
}

type DialDown struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Controller  string      `json:"controller"`
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
	} `json:"payload"`
}

func (m *DialDown) GetEvent() string {
	return m.Event
}

func (m *DialDown) GetAction() string {
	return m.Action
}

func (m *DialDown) GetContext() string {
	return m.Context
}

func (m *DialDown) GetDevice() string {
	return m.Device
}

type DialUp struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Controller  string      `json:"controller"`
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
	} `json:"payload"`
}

func (m *DialUp) GetEvent() string {
	return m.Event
}

func (m *DialUp) GetAction() string {
	return m.Action
}

func (m *DialUp) GetContext() string {
	return m.Context
}

func (m *DialUp) GetDevice() string {
	return m.Device
}

type DialRotate struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Controller  string      `json:"controller"`
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		Ticks   int  `json:"ticks"`
		Pressed bool `json:"pressed"`
	} `json:"payload"`
}

func (m *DialRotate) GetEvent() string {
	return m.Event
}

func (m *DialRotate) GetAction() string {
	return m.Action
}

func (m *DialRotate) GetContext() string {
	return m.Context
}

func (m *DialRotate) GetDevice() string {
	return m.Device
}

type KeyDown struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State           int  `json:"state"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

func (m *KeyDown) GetEvent() string {
	return m.Event
}

func (m *KeyDown) GetAction() string {
	return m.Action
}

func (m *KeyDown) GetContext() string {
	return m.Context
}

func (m *KeyDown) GetDevice() string {
	return m.Device
}

type KeyUp struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		State           int  `json:"state"`
		IsInMultiAction bool `json:"isInMultiAction"`
	} `json:"payload"`
}

func (m *KeyUp) GetEvent() string {
	return m.Event
}

func (m *KeyUp) GetAction() string {
	return m.Action
}

func (m *KeyUp) GetContext() string {
	return m.Context
}

func (m *KeyUp) GetDevice() string {
	return m.Device
}

type WillAppear struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		Controller      string `json:"controller"`
		State           int    `json:"state"`
		IsInMultiAction bool   `json:"isInMultiAction"`
	} `json:"payload"`
}

func (m *WillAppear) GetEvent() string {
	return m.Event
}

func (m *WillAppear) GetAction() string {
	return m.Action
}

func (m *WillAppear) GetContext() string {
	return m.Context
}

func (m *WillAppear) GetDevice() string {
	return m.Device
}

type WillDisappear struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Settings    interface{} `json:"settings"`
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		Controller      string `json:"controller"`
		State           int    `json:"state"`
		IsInMultiAction bool   `json:"isInMultiAction"`
	} `json:"payload"`
}

func (m *WillDisappear) GetEvent() string {
	return m.Event
}

func (m *WillDisappear) GetAction() string {
	return m.Action
}

func (m *WillDisappear) GetContext() string {
	return m.Context
}

func (m *WillDisappear) GetDevice() string {
	return m.Device
}

type TitleParametersDidChange struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
	Payload struct {
		Coordinates struct {
			Column int `json:"column"`
			Row    int `json:"row"`
		} `json:"coordinates"`
		Settings        interface{} `json:"settings"`
		State           int         `json:"state"`
		Title           string      `json:"title"`
		TitleParameters struct {
			FontFamily     string `json:"fontFamily"`
			FontSize       int    `json:"fontSize"`
			FontStyle      string `json:"fontStyle"`
			FontUnderline  bool   `json:"fontUnderline"`
			ShowTitle      bool   `json:"showTitle"`
			TitleAlignment string `json:"titleAlignment"`
			TitleColor     string `json:"titleColor"`
		} `json:"titleParameters"`
	} `json:"payload"`
}

func (m *TitleParametersDidChange) GetEvent() string {
	return m.Event
}

func (m *TitleParametersDidChange) GetAction() string {
	return m.Action
}

func (m *TitleParametersDidChange) GetContext() string {
	return m.Context
}

func (m *TitleParametersDidChange) GetDevice() string {
	return m.Device
}

type DeviceDidConnect struct {
	Event      string `json:"event"`
	Device     string `json:"device"`
	DeviceInfo struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		Size struct {
			Columns int `json:"columns"`
			Rows    int `json:"rows"`
		} `json:"size"`
	} `json:"deviceInfo"`
}

func (m *DeviceDidConnect) GetEvent() string {
	return m.Event
}

type DeviceDidDisconnect struct {
	Event  string `json:"event"`
	Device string `json:"device"`
}

func (m *DeviceDidDisconnect) GetEvent() string {
	return m.Event
}

type ApplicationDidLaunch struct {
	Event   string `json:"event"`
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

func (m *ApplicationDidLaunch) GetEvent() string {
	return m.Event
}

type ApplicationDidTerminate struct {
	Event   string `json:"event"`
	Payload struct {
		Application string `json:"application"`
	} `json:"payload"`
}

func (m *ApplicationDidTerminate) GetEvent() string {
	return m.Event
}

type SystemDidWakeUp struct {
	Event string `json:"event"`
}

func (m *SystemDidWakeUp) GetEvent() string {
	return m.Event
}

type PropertyInspectorDidAppear struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
}

func (m *PropertyInspectorDidAppear) GetEvent() string {
	return m.Event
}

func (m *PropertyInspectorDidAppear) GetAction() string {
	return m.Action
}

func (m *PropertyInspectorDidAppear) GetContext() string {
	return m.Context
}

func (m *PropertyInspectorDidAppear) GetDevice() string {
	return m.Device
}

type PropertyInspectorDidDisappear struct {
	Action  string `json:"action"`
	Event   string `json:"event"`
	Context string `json:"context"`
	Device  string `json:"device"`
}

func (m *PropertyInspectorDidDisappear) GetEvent() string {
	return m.Event
}

func (m *PropertyInspectorDidDisappear) GetAction() string {
	return m.Action
}

func (m *PropertyInspectorDidDisappear) GetContext() string {
	return m.Context
}

func (m *PropertyInspectorDidDisappear) GetDevice() string {
	return m.Device
}
