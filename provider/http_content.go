package provider

type httpContent interface {
	GetData() ([]byte, error)
	GetBody() interface{}
	SetBody(content interface{}) error
}
