package models

var bitcoinEncoding = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func join(bs ...byte) string {
	return string(bs)
}
