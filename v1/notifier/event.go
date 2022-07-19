package notifier

type IEvent interface {
	Token() string
	Title() string
	Body() interface{} // Href
}

// Default implementation of IEvent
type StatusChangedEvent struct {
	_title string
	_body  interface{}
	_token string
}

func (e *StatusChangedEvent) Token() string {
	return e._token
}

func (e *StatusChangedEvent) Title() string {
	return e._title
}

func (e *StatusChangedEvent) Body() interface{} {
	return e._body
}

func NewStatusChangedEvent(_title string, _body interface{}, _token string) IEvent {
	return &StatusChangedEvent{
		_title,
		_body,
		_token,
	}
}

type PushEvent struct {
	_title string
	_body  interface{}
	_token string
}

func (e *PushEvent) Token() string {
	return e._token
}

func (e *PushEvent) Title() string {
	return e._title
}

func (e *PushEvent) Body() interface{} {
	return e._body
}

func NewPushEvent(_title string, _body interface{}, _token string) IEvent {
	return &PushEvent{
		_title,
		_body,
		_token,
	}
}
