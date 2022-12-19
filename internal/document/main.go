package document

import (
	"os"
	"strings"
)

func DocumentExist(p string) bool {
	_, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func Read(p string) (string, error) {
	body, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}

func Write(p string, content string) error {
	f, err := os.OpenFile(p, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
