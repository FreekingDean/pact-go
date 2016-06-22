package provider

import "errors"

type plainTextContent struct {
	data string
}

func (c *plainTextContent) GetData() ([]byte, error) {
	return []byte(c.data), nil
}

func (c *plainTextContent) GetBody() interface{} {
	return c.data
}

func (c *plainTextContent) SetBody(content interface{}) error {
	if v, ok := content.(string); ok {
		c.data = v
		return nil
	}
	return errors.New("content is not valid text")
}
