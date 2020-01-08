package Bencode
import "bytes"
func Parse(str string) string{
	var buffer = bytes.NewBufferString(str);
	buffer.ReadByte()
	return ""
}
