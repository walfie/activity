// Package vocab contains the interfaces for the W3IDSecurityV1 vocabulary. All
// applications are strongly encouraged to use these interface types instead
// of the concrete definitions contained in the implementation subpackage.
// These interfaces allow applications to consume only the types and
// properties needed and be independent of the go-fed implementation if
// another alternative implementation is created. This package is
// code-generated and subject to the same license as the go-fed tool used to
// generate it.
//
// Type interfaces contain "Get" and "Set" methods for its properties. Types
// also have a "Serialize" method to convert the type into an interface map
// for use with the json package. There is a convenience "IsExtending" method
// on each types which helps with the ActivityStreams hierarchy, which is not
// the same as object oriented inheritance. While types also have a "LessThan"
// method, it is an arbitrary sort. Do not use it if needing to sort on
// specific properties, such as publish time. It is best used for normalizing
// the type. Lastly, do not use the "GetUnknownProperties" method in an
// application. Instead, use the go-fed tool to code generate the property
// needed.
//
// Properties come in two flavors: functional and non-functional. Functional
// means that a property can have at most one value, while non-functional
// means a property could have zero, one, or more values. Any property value
// may also be an IRI, in which case the application will need to make a HTTP
// request to fetch the property value.
//
// Functional properties have "Get", "Is", and "Set" methods for determining
// what kind of value the property is, fetching that value, or setting that
// value. There is also a "Serialize" method which converts the property into
// an interface type, but applications should not typically use a property's
// "Serialize" and instead should use a type's "Serialize" instead. Like
// types, properties have an arbitrary "LessThan" comparison function that
// should not be used if needing to sort on specific values. Finally,
// applications should not use the "KindIndex" method as it is a comparison
// mechanism only for those looking to write an alternate implementation.
//
// Non-functional properties can have more than one value, so it has  "Len"
// for getting its length, "At" for getting an iterator pointing to an
// element, "Append" and "Prepend" for adding values, "Remove" for removing a
// value, "Set" for overwriting a value, and "Swap" for swapping two values'
// indices. Note that a non-functional property satisfies the sort interface,
// but it results in an arbitrary but stable ordering best used as a
// normalized form. A non-functional property's iterator looks like a
// functional property with "Next" and "Previous" methods. Applications should
// not use the "KindIndex" methods as it is a comparison mechanism only for
// those looking to write an alternate implementation of this library.
//
// Types and properties have a "JSONLDContext" method that returns a mapping
// of vocabulary URIs to aliases that are required in the JSON-LD @context
// when serializing this value. The aliases used by this library when
// serializing objects is done at code-generation time, unless a different
// alias was used to deserialize the type or property.
//
// Types, functional properties, and non-functional properties are not
// designed for concurrent usage by two or more goroutines. Also, certain
// methods on a non-functional property will invalidate iterators and possibly
// cause unexpected behaviors. To avoid this, re-obtain an iterator after
// modifying a non-functional property.
package vocab
