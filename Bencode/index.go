package Bencode

import (
	"bytes"
	"fmt"
	"strings"
)

func Parse(str string) string{
	var produce Produce;
	produce.NewProduce(str);
	context := &Context{
		Index:-1,
		CurChar:rune(0),
		CurType:BenTypeUnknow{},
		Buffer:&strings.Builder{},
		Parent:nil,
	}
	var customer Customer;

	index := int32(-1)
	for{
		r,err := produce.Next();
		if(err != nil){
			// todo 判断错误类型
			fmt.Print(r)
			break;
		}
		parseError := customer.eat(context,r);
		if(parseError != nil){
			// todo 抛出错误栈
			break;
		}

		index+=1
		//context.Index = index;
		//context.CurChar = r;
	}
	return ""
}
// 解析上下文
type Context struct {
	Index int32
	CurChar rune
	CurType BenType
	Buffer *strings.Builder
	Parent *Context
}

//func (p *Context) CopyConext(){
//
//}

type Customer struct {}

func (c Customer) eat(context *Context,chr rune) error{
	switch chr {
	case ':':
		return nil


	}
	return nil
}

type Produce struct{
	buffer *bytes.Buffer
	index int;
}

func (p *Produce) NewProduce(str string) {
	p.buffer = bytes.NewBufferString(str);
	p.index=0;
	//return &p
}

func (p *Produce) Next() (rune,error){
	var r,_,err = p.buffer.ReadRune()
	p.index +=1;
	if(err == nil){
		return r,nil
	}
	return rune(0),err
}

type  Iterator interface{
	Next()rune
}

