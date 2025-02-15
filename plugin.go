package cli

func New(apps ...*Application) *Plugin {
	if len(apps) == 0 {
		app := &Application{}
		app.Init()
		return &Plugin{Instance: app}
	}
	return &Plugin{Instance: apps[0]}
}

type Plugin struct {
	Instance *Application
}

func (receiver *Plugin) Init() {
}

func (receiver *Plugin) Register() {
	if receiver.Instance == nil {
		app := &Application{}
		app.Init()
		receiver.Instance = app
	}
}

func (receiver *Plugin) Name() string {
	return "Console"
}

func (receiver *Plugin) DependsOn() []string {
	return []string{}
}

func (receiver *Plugin) Close() error {
	return nil
}
