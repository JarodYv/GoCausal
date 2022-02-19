package graph

type IAttribute interface {
	GetAllAttributes() map[string]interface{}
	GetAttribute(string) interface{}
	AddAttribute(string, interface{})
	RemoveAttribute(string)
}

type Attribute struct {
	IAttribute
	attributes map[string]interface{}
}

func (attr *Attribute) GetAllAttributes() map[string]interface{} {
	return attr.attributes
}

func (attr *Attribute) GetAttribute(key string) interface{} {
	return attr.attributes[key]
}

func (attr *Attribute) AddAttribute(key string, value interface{}) {
	attr.attributes[key] = value
}

func (attr *Attribute) RemoveAttribute(key string) {
	delete(attr.attributes, key)
}
