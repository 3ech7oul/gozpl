package zpl

import "errors"

var zplCommnads = map[string]string{
	"^FD": "",
	"^XA": "",
	"^FS": "",
	"^GB": "",
	"^FO": "",
	"^CF": "",
}

type ZplComm interface {
	hasParameters() bool
}

type Command struct {
	ZplComm    string
	Buffer     []byte
	Parameters []CommandParameter
}

type CommandParameter struct {
	Parameter string
	Value     string
}

func CreateCommand(command string) (Command, error) {
	var result Command

	_, ok := zplCommnads[command]
	if ok {
		result = Command{
			ZplComm: command,
		}

		return result, nil
	}

	return result, errors.New("Not a command")
}

func (c *Command) AddToBuffer(b byte) {
	c.Buffer = append(c.Buffer, b)
}
