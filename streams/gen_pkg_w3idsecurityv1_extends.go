package streams

import (
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// W3IDSecurityV1W3IDSecurityV1PublicKeyExtends returns true if PublicKey extends
// from the other's type.
func W3IDSecurityV1W3IDSecurityV1PublicKeyExtends(other vocab.Type) bool {
	return typepublickey.W3IDSecurityV1PublicKeyExtends(other)
}
