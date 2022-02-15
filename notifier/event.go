package notifier

type IEvent interface {
	Token() string
	Title() string
	Body() string // Href
}

// Default implementation of IEvent
type StatusChangedEvent struct {
	_title string
	_body  string
	_token string
}

func (e *StatusChangedEvent) Token() string {
	return e._token
}

func (e *StatusChangedEvent) Title() string {
	return e._title
}

func (e *StatusChangedEvent) Body() string {
	return e._body
}

func NewStatusChangedEvent(_title, _body, _token string) IEvent {
	return &StatusChangedEvent{
		_title,
		_body,
		_token,
	}
}
