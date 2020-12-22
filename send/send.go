package send

type Send interface {
	Send(subject string, content string, to string) error
	SendToArray(subject string, content string, to ...string) error
}
