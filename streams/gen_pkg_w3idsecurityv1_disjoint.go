package streams

import (
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// W3IDSecurityV1PublicKeyIsDisjointWith returns true if PublicKey is disjoint
// with the other's type.
func W3IDSecurityV1PublicKeyIsDisjointWith(other vocab.Type) bool {
	return typepublickey.PublicKeyIsDisjointWith(other)
}
