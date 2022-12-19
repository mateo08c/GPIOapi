package gpio

import (
	"GPIOapi/internal/document"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type GPIO struct {
	PinID string `json:"pinID"`
}

type Export struct {
	Direction string `json:"direction"`
	Value     string `json:"value"`
}

func (e Export) ToJson() any {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func New(pinID string) (*GPIO, error) {
	n := &GPIO{
		PinID: pinID,
	}

	_, err := n.Pin()
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (d GPIO) Value() (string, error) {
	id, _ := d.Pin()
	if d.IsExported() {
		content, err := document.Read(fmt.Sprintf("/sys/class/gpio/gpio%d/value", id))
		if err != nil {
			return "", fmt.Errorf("read value pin %d failed", id)
		}
		return content, nil
	}

	return "", fmt.Errorf("pin %d not exported", id)
}

func (d GPIO) Direction() (string, error) {
	id, _ := d.Pin()
	if d.IsExported() {
		content, err := document.Read(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", id))
		if err != nil {
			return "", fmt.Errorf("read direction pin %d failed", id)
		}
		return content, nil
	}

	return "", fmt.Errorf("pin %d not exported", id)
}

func (d GPIO) Export() error {
	id, _ := d.Pin()
	if !d.IsExported() {
		err := document.Write("/sys/class/gpio/export", fmt.Sprintf("%d", id))
		if err != nil {
			return fmt.Errorf("export pin %d failed", id)
		}
	} else {
		return fmt.Errorf("pin %d already exported", id)
	}

	return nil
}

func (d GPIO) Unexport() error {
	id, _ := d.Pin()
	if d.IsExported() {
		err := document.Write("/sys/class/gpio/unexport", fmt.Sprintf("%d", id))
		if err != nil {
			return fmt.Errorf("unexport pin %d failed", id)
		}
	} else {
		return fmt.Errorf("pin %d not exported", id)
	}

	return nil
}

func (d GPIO) SetValue(v string) error {
	id, _ := d.Pin()
	if d.IsExported() {
		err := document.Write(fmt.Sprintf("/sys/class/gpio/gpio%d/value", id), v)
		if err != nil {
			return fmt.Errorf("set value pin %d failed", id)
		}

		return nil
	}

	return fmt.Errorf("pin %d not exported", id)
}

func (d GPIO) SetDirection(v string) error {
	id, _ := d.Pin()
	if d.IsExported() {
		err := document.Write(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", id), v)
		if err != nil {
			return fmt.Errorf("set direction pin %d failed", id)
		}

		return nil
	}

	return fmt.Errorf("pin direction %d not exported", id)
}

func (d GPIO) Pin() (int, error) {
	v := strings.ToUpper(d.PinID)
	if len(v) >= 2 {
		le := v[1:2]
		li := strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", le)
		ln := v[2:]
		ai, err := strconv.Atoi(ln)
		if err != nil {
			return 0, fmt.Errorf("pin %s not supported", d.PinID)
		}
		res := (li * 32) + ai

		return res, nil
	}
	return 0, fmt.Errorf("pin %s not supported", d.PinID)
}

func (d GPIO) IsExported() bool {
	id, _ := d.Pin()
	return document.DocumentExist(fmt.Sprintf("/sys/class/gpio/gpio%d", id))
}
