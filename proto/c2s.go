package proto

type Register struct {
	UUID  string `json:"uuid"`
	Event string `json:"event"`
}

//func NewRegister() *Register {
//	return &Register{}
//}

type SetSettings struct {
	Event   string                 `json:"event"`
	Context string                 `json:"context"`
	Payload map[string]interface{} `json:"payload"`
}

func NewSetSettings() *SetSettings {
	return &SetSettings{
		Event: "setSettings",
	}
}

type GetSettings struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

func NewGetSettings() *GetSettings {
	return &GetSettings{
		Event: "getSettings",
	}
}

type SetGlobalSettings struct {
	Event   string                 `json:"event"`
	Context string                 `json:"context"`
	Payload map[string]interface{} `json:"payload"`
}

func NewSetGlobalSettings() *SetGlobalSettings {
	return &SetGlobalSettings{
		Event: "setGlobalSettings",
	}
}

type GetGlobalSettings struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

func NewGetGlobalSettings() *GetGlobalSettings {
	return &GetGlobalSettings{
		Event: "getGlobalSettings",
	}
}

type OpenUrl struct {
	Event   string `json:"event"`
	Payload struct {
		Url string `json:"url"`
	} `json:"payload"`
}

func NewOpenUrl() *OpenUrl {
	return &OpenUrl{
		Event: "openUrl",
	}
}

type LogMessage struct {
	Event   string `json:"event"`
	Payload struct {
		Message string `json:"message"`
	} `json:"payload"`
}

func NewLogMessage() *LogMessage {
	return &LogMessage{
		Event: "logMessage",
	}
}

type SetTitle struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		Title  string       `json:"title"`
		Target ESDSDKTarget `json:"target"` //Specifies whether the title should be displayed on hardware and software (0), only on hardware (1), only on software (2), default is 0
		State  int          `json:"state"`  //An integer value starting from 0, representing the state of an action with multiple states. If not specified, the title will be set for all states
	} `json:"payload"`
}

func NewSetTitle() *SetTitle {
	return &SetTitle{
		Event: "setTitle",
	}
}

type SetImage struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		Image  string `json:"image"`  //The base64-encoded image to set
		Target int    `json:"target"` //Specifies whether the title should be displayed on hardware and software (0), only on hardware (1), only on software (2), default is 0
		State  int    `json:"state"`  //An integer value starting from 0, representing the state of an action with multiple states. If not specified, the title will be set for all states
	} `json:"payload"`
}

func NewSetImage() *SetImage {
	return &SetImage{
		Event: "setImage",
	}
}

type ShowAlert struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

func NewShowAlert() *ShowAlert {
	return &ShowAlert{
		Event: "showAlert",
	}
}

type ShowOk struct {
	Event   string `json:"event"`
	Context string `json:"context"`
}

func NewShowOk() *ShowOk {
	return &ShowOk{
		Event: "showOk",
	}
}

type SetState struct {
	Event   string `json:"event"`
	Context string `json:"context"`
	Payload struct {
		State int `json:"state"`
	} `json:"payload"`
}

func NewSetState() *SetState {
	return &SetState{
		Event: "setState",
	}
}

type SendToPropertyInspector struct {
	Action  string                 `json:"action"`
	Event   string                 `json:"event"`
	Context string                 `json:"context"`
	Payload map[string]interface{} `json:"payload"`
}

func NewSendToPropertyInspector() *SendToPropertyInspector {
	return &SendToPropertyInspector{
		Event: "sendToPropertyInspector",
	}
}

func (m *SendToPropertyInspector) GetEvent() string {
	return m.Event
}

func (m *SendToPropertyInspector) GetAction() string {
	return m.Action
}

func (m *SendToPropertyInspector) GetContext() string {
	return m.Context
}

type SendToPlugin struct {
	Action  string                 `json:"action"`
	Event   string                 `json:"event"`
	Context string                 `json:"context"`
	Payload map[string]interface{} `json:"payload"`
}

func NewSendToPlugin() *SendToPlugin {
	return &SendToPlugin{
		Event: "sendToPlugin",
	}
}

func (m *SendToPlugin) GetEvent() string {
	return m.Event
}

func (m *SendToPlugin) GetAction() string {
	return m.Action
}

func (m *SendToPlugin) GetContext() string {
	return m.Context
}
