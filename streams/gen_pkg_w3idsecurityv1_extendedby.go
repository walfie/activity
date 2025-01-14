package streams

import (
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// W3IDSecurityV1PublicKeyIsExtendedBy returns true if the other's type extends
// from PublicKey. Note that it returns false if the types are the same; see
// the "IsOrExtends" variant instead.
func W3IDSecurityV1PublicKeyIsExtendedBy(other vocab.Type) bool {
	return typepublickey.PublicKeyIsExtendedBy(other)
}
