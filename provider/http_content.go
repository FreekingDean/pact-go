package provider

type httpContent interface {
	GetData() ([]byte, error)
	GetBody() interface{}
	SetBody(content interface{}) error
}

// type emptyContent struct {
// }

// func (e *emptyContent) GetData() ([]byte, error) {
// 	return nil, nil
// }

// func (e *emptyContent) GetBody() interface{} {
// 	return nil
// }

// func (e *emptyContent) SetBody(content interface{}) error {
// 	return errors.New("cannot set content for empty body")
// }

