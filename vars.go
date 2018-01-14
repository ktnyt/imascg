package imascg

import uuid "github.com/satori/go.uuid"

var bitcoinEncoding = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var apiNamespace = uuid.NewV5(uuid.NamespaceURL, "api.imascg.moe")
