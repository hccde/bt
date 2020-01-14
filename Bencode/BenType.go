package Bencode
type BenType interface {
	GetTypeName()string
}

type BenTypeString struct {

}

func (p BenTypeString) GetTypeName() string{
	return "string"
}

type BenTypeInteger struct {

}

func (p BenTypeInteger) GetTypeName() string{
	return "integer"
}

type BenTypeList struct {

}
func (p BenTypeList) GetTypeName() string{
	return "list"
}
type BenTypeDictionary struct {

}
func (p BenTypeDictionary) GetTypeName() string{
	return "dictionary"
}

type BenTypeUnknow struct {

}
func (p BenTypeUnknow) GetTypeName() string{
	return "unknow"
}