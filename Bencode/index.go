package Bencode

import (
	"bytes"
	"fmt"
)
func Parse(str string) string{
	var produce Produce;
	produce.NewProduce(str);
	r,err := produce.Next();
	if(err == nil){
		fmt.Print(r);
	}

	//if r,err := produce.Next();err==nil {
	//	fmt.Print(int(r))
	//}
	return ""
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