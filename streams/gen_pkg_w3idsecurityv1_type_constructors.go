package streams

import (
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// NewW3IDSecurityV1PublicKey creates a new W3IDSecurityV1PublicKey
func NewW3IDSecurityV1PublicKey() vocab.W3IDSecurityV1PublicKey {
	return typepublickey.NewW3IDSecurityV1PublicKey()
}
