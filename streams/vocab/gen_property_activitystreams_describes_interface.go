package vocab

import "net/url"

// On a Profile object, the describes property identifies the object described by
// the Profile.
//
// Example 141 (https://www.w3.org/TR/activitystreams-vocabulary/#ex185-jsonld):
//   {
//     "describes": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "summary": "Sally's profile",
//     "type": "Profile",
//     "url": "http://sally.example.org"
//   }
type ActivityStreamsDescribesProperty interface {
	// Clear ensures no value of this property is set. Calling HasAny or any
	// of the 'Is' methods afterwards will return false.
	Clear()
	// GetActivityStreamsAccept returns the value of this property. When
	// IsActivityStreamsAccept returns false, GetActivityStreamsAccept
	// will return an arbitrary value.
	GetActivityStreamsAccept() ActivityStreamsAccept
	// GetActivityStreamsActivity returns the value of this property. When
	// IsActivityStreamsActivity returns false, GetActivityStreamsActivity
	// will return an arbitrary value.
	GetActivityStreamsActivity() ActivityStreamsActivity
	// GetActivityStreamsAdd returns the value of this property. When
	// IsActivityStreamsAdd returns false, GetActivityStreamsAdd will
	// return an arbitrary value.
	GetActivityStreamsAdd() ActivityStreamsAdd
	// GetActivityStreamsAnnounce returns the value of this property. When
	// IsActivityStreamsAnnounce returns false, GetActivityStreamsAnnounce
	// will return an arbitrary value.
	GetActivityStreamsAnnounce() ActivityStreamsAnnounce
	// GetActivityStreamsApplication returns the value of this property. When
	// IsActivityStreamsApplication returns false,
	// GetActivityStreamsApplication will return an arbitrary value.
	GetActivityStreamsApplication() ActivityStreamsApplication
	// GetActivityStreamsArrive returns the value of this property. When
	// IsActivityStreamsArrive returns false, GetActivityStreamsArrive
	// will return an arbitrary value.
	GetActivityStreamsArrive() ActivityStreamsArrive
	// GetActivityStreamsArticle returns the value of this property. When
	// IsActivityStreamsArticle returns false, GetActivityStreamsArticle
	// will return an arbitrary value.
	GetActivityStreamsArticle() ActivityStreamsArticle
	// GetActivityStreamsAudio returns the value of this property. When
	// IsActivityStreamsAudio returns false, GetActivityStreamsAudio will
	// return an arbitrary value.
	GetActivityStreamsAudio() ActivityStreamsAudio
	// GetActivityStreamsBlock returns the value of this property. When
	// IsActivityStreamsBlock returns false, GetActivityStreamsBlock will
	// return an arbitrary value.
	GetActivityStreamsBlock() ActivityStreamsBlock
	// GetActivityStreamsCollection returns the value of this property. When
	// IsActivityStreamsCollection returns false,
	// GetActivityStreamsCollection will return an arbitrary value.
	GetActivityStreamsCollection() ActivityStreamsCollection
	// GetActivityStreamsCollectionPage returns the value of this property.
	// When IsActivityStreamsCollectionPage returns false,
	// GetActivityStreamsCollectionPage will return an arbitrary value.
	GetActivityStreamsCollectionPage() ActivityStreamsCollectionPage
	// GetActivityStreamsCreate returns the value of this property. When
	// IsActivityStreamsCreate returns false, GetActivityStreamsCreate
	// will return an arbitrary value.
	GetActivityStreamsCreate() ActivityStreamsCreate
	// GetActivityStreamsDelete returns the value of this property. When
	// IsActivityStreamsDelete returns false, GetActivityStreamsDelete
	// will return an arbitrary value.
	GetActivityStreamsDelete() ActivityStreamsDelete
	// GetActivityStreamsDislike returns the value of this property. When
	// IsActivityStreamsDislike returns false, GetActivityStreamsDislike
	// will return an arbitrary value.
	GetActivityStreamsDislike() ActivityStreamsDislike
	// GetActivityStreamsDocument returns the value of this property. When
	// IsActivityStreamsDocument returns false, GetActivityStreamsDocument
	// will return an arbitrary value.
	GetActivityStreamsDocument() ActivityStreamsDocument
	// GetActivityStreamsEvent returns the value of this property. When
	// IsActivityStreamsEvent returns false, GetActivityStreamsEvent will
	// return an arbitrary value.
	GetActivityStreamsEvent() ActivityStreamsEvent
	// GetActivityStreamsFlag returns the value of this property. When
	// IsActivityStreamsFlag returns false, GetActivityStreamsFlag will
	// return an arbitrary value.
	GetActivityStreamsFlag() ActivityStreamsFlag
	// GetActivityStreamsFollow returns the value of this property. When
	// IsActivityStreamsFollow returns false, GetActivityStreamsFollow
	// will return an arbitrary value.
	GetActivityStreamsFollow() ActivityStreamsFollow
	// GetActivityStreamsGroup returns the value of this property. When
	// IsActivityStreamsGroup returns false, GetActivityStreamsGroup will
	// return an arbitrary value.
	GetActivityStreamsGroup() ActivityStreamsGroup
	// GetActivityStreamsIgnore returns the value of this property. When
	// IsActivityStreamsIgnore returns false, GetActivityStreamsIgnore
	// will return an arbitrary value.
	GetActivityStreamsIgnore() ActivityStreamsIgnore
	// GetActivityStreamsImage returns the value of this property. When
	// IsActivityStreamsImage returns false, GetActivityStreamsImage will
	// return an arbitrary value.
	GetActivityStreamsImage() ActivityStreamsImage
	// GetActivityStreamsIntransitiveActivity returns the value of this
	// property. When IsActivityStreamsIntransitiveActivity returns false,
	// GetActivityStreamsIntransitiveActivity will return an arbitrary
	// value.
	GetActivityStreamsIntransitiveActivity() ActivityStreamsIntransitiveActivity
	// GetActivityStreamsInvite returns the value of this property. When
	// IsActivityStreamsInvite returns false, GetActivityStreamsInvite
	// will return an arbitrary value.
	GetActivityStreamsInvite() ActivityStreamsInvite
	// GetActivityStreamsJoin returns the value of this property. When
	// IsActivityStreamsJoin returns false, GetActivityStreamsJoin will
	// return an arbitrary value.
	GetActivityStreamsJoin() ActivityStreamsJoin
	// GetActivityStreamsLeave returns the value of this property. When
	// IsActivityStreamsLeave returns false, GetActivityStreamsLeave will
	// return an arbitrary value.
	GetActivityStreamsLeave() ActivityStreamsLeave
	// GetActivityStreamsLike returns the value of this property. When
	// IsActivityStreamsLike returns false, GetActivityStreamsLike will
	// return an arbitrary value.
	GetActivityStreamsLike() ActivityStreamsLike
	// GetActivityStreamsListen returns the value of this property. When
	// IsActivityStreamsListen returns false, GetActivityStreamsListen
	// will return an arbitrary value.
	GetActivityStreamsListen() ActivityStreamsListen
	// GetActivityStreamsMove returns the value of this property. When
	// IsActivityStreamsMove returns false, GetActivityStreamsMove will
	// return an arbitrary value.
	GetActivityStreamsMove() ActivityStreamsMove
	// GetActivityStreamsNote returns the value of this property. When
	// IsActivityStreamsNote returns false, GetActivityStreamsNote will
	// return an arbitrary value.
	GetActivityStreamsNote() ActivityStreamsNote
	// GetActivityStreamsObject returns the value of this property. When
	// IsActivityStreamsObject returns false, GetActivityStreamsObject
	// will return an arbitrary value.
	GetActivityStreamsObject() ActivityStreamsObject
	// GetActivityStreamsOffer returns the value of this property. When
	// IsActivityStreamsOffer returns false, GetActivityStreamsOffer will
	// return an arbitrary value.
	GetActivityStreamsOffer() ActivityStreamsOffer
	// GetActivityStreamsOrderedCollection returns the value of this property.
	// When IsActivityStreamsOrderedCollection returns false,
	// GetActivityStreamsOrderedCollection will return an arbitrary value.
	GetActivityStreamsOrderedCollection() ActivityStreamsOrderedCollection
	// GetActivityStreamsOrderedCollectionPage returns the value of this
	// property. When IsActivityStreamsOrderedCollectionPage returns
	// false, GetActivityStreamsOrderedCollectionPage will return an
	// arbitrary value.
	GetActivityStreamsOrderedCollectionPage() ActivityStreamsOrderedCollectionPage
	// GetActivityStreamsOrganization returns the value of this property. When
	// IsActivityStreamsOrganization returns false,
	// GetActivityStreamsOrganization will return an arbitrary value.
	GetActivityStreamsOrganization() ActivityStreamsOrganization
	// GetActivityStreamsPage returns the value of this property. When
	// IsActivityStreamsPage returns false, GetActivityStreamsPage will
	// return an arbitrary value.
	GetActivityStreamsPage() ActivityStreamsPage
	// GetActivityStreamsPerson returns the value of this property. When
	// IsActivityStreamsPerson returns false, GetActivityStreamsPerson
	// will return an arbitrary value.
	GetActivityStreamsPerson() ActivityStreamsPerson
	// GetActivityStreamsPlace returns the value of this property. When
	// IsActivityStreamsPlace returns false, GetActivityStreamsPlace will
	// return an arbitrary value.
	GetActivityStreamsPlace() ActivityStreamsPlace
	// GetActivityStreamsProfile returns the value of this property. When
	// IsActivityStreamsProfile returns false, GetActivityStreamsProfile
	// will return an arbitrary value.
	GetActivityStreamsProfile() ActivityStreamsProfile
	// GetActivityStreamsQuestion returns the value of this property. When
	// IsActivityStreamsQuestion returns false, GetActivityStreamsQuestion
	// will return an arbitrary value.
	GetActivityStreamsQuestion() ActivityStreamsQuestion
	// GetActivityStreamsRead returns the value of this property. When
	// IsActivityStreamsRead returns false, GetActivityStreamsRead will
	// return an arbitrary value.
	GetActivityStreamsRead() ActivityStreamsRead
	// GetActivityStreamsReject returns the value of this property. When
	// IsActivityStreamsReject returns false, GetActivityStreamsReject
	// will return an arbitrary value.
	GetActivityStreamsReject() ActivityStreamsReject
	// GetActivityStreamsRelationship returns the value of this property. When
	// IsActivityStreamsRelationship returns false,
	// GetActivityStreamsRelationship will return an arbitrary value.
	GetActivityStreamsRelationship() ActivityStreamsRelationship
	// GetActivityStreamsRemove returns the value of this property. When
	// IsActivityStreamsRemove returns false, GetActivityStreamsRemove
	// will return an arbitrary value.
	GetActivityStreamsRemove() ActivityStreamsRemove
	// GetActivityStreamsService returns the value of this property. When
	// IsActivityStreamsService returns false, GetActivityStreamsService
	// will return an arbitrary value.
	GetActivityStreamsService() ActivityStreamsService
	// GetActivityStreamsTentativeAccept returns the value of this property.
	// When IsActivityStreamsTentativeAccept returns false,
	// GetActivityStreamsTentativeAccept will return an arbitrary value.
	GetActivityStreamsTentativeAccept() ActivityStreamsTentativeAccept
	// GetActivityStreamsTentativeReject returns the value of this property.
	// When IsActivityStreamsTentativeReject returns false,
	// GetActivityStreamsTentativeReject will return an arbitrary value.
	GetActivityStreamsTentativeReject() ActivityStreamsTentativeReject
	// GetActivityStreamsTombstone returns the value of this property. When
	// IsActivityStreamsTombstone returns false,
	// GetActivityStreamsTombstone will return an arbitrary value.
	GetActivityStreamsTombstone() ActivityStreamsTombstone
	// GetActivityStreamsTravel returns the value of this property. When
	// IsActivityStreamsTravel returns false, GetActivityStreamsTravel
	// will return an arbitrary value.
	GetActivityStreamsTravel() ActivityStreamsTravel
	// GetActivityStreamsUndo returns the value of this property. When
	// IsActivityStreamsUndo returns false, GetActivityStreamsUndo will
	// return an arbitrary value.
	GetActivityStreamsUndo() ActivityStreamsUndo
	// GetActivityStreamsUpdate returns the value of this property. When
	// IsActivityStreamsUpdate returns false, GetActivityStreamsUpdate
	// will return an arbitrary value.
	GetActivityStreamsUpdate() ActivityStreamsUpdate
	// GetActivityStreamsVideo returns the value of this property. When
	// IsActivityStreamsVideo returns false, GetActivityStreamsVideo will
	// return an arbitrary value.
	GetActivityStreamsVideo() ActivityStreamsVideo
	// GetActivityStreamsView returns the value of this property. When
	// IsActivityStreamsView returns false, GetActivityStreamsView will
	// return an arbitrary value.
	GetActivityStreamsView() ActivityStreamsView
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsAccept returns true if this property has a type of
	// "Accept". When true, use the GetActivityStreamsAccept and
	// SetActivityStreamsAccept methods to access and set this property.
	IsActivityStreamsAccept() bool
	// IsActivityStreamsActivity returns true if this property has a type of
	// "Activity". When true, use the GetActivityStreamsActivity and
	// SetActivityStreamsActivity methods to access and set this property.
	IsActivityStreamsActivity() bool
	// IsActivityStreamsAdd returns true if this property has a type of "Add".
	// When true, use the GetActivityStreamsAdd and SetActivityStreamsAdd
	// methods to access and set this property.
	IsActivityStreamsAdd() bool
	// IsActivityStreamsAnnounce returns true if this property has a type of
	// "Announce". When true, use the GetActivityStreamsAnnounce and
	// SetActivityStreamsAnnounce methods to access and set this property.
	IsActivityStreamsAnnounce() bool
	// IsActivityStreamsApplication returns true if this property has a type
	// of "Application". When true, use the GetActivityStreamsApplication
	// and SetActivityStreamsApplication methods to access and set this
	// property.
	IsActivityStreamsApplication() bool
	// IsActivityStreamsArrive returns true if this property has a type of
	// "Arrive". When true, use the GetActivityStreamsArrive and
	// SetActivityStreamsArrive methods to access and set this property.
	IsActivityStreamsArrive() bool
	// IsActivityStreamsArticle returns true if this property has a type of
	// "Article". When true, use the GetActivityStreamsArticle and
	// SetActivityStreamsArticle methods to access and set this property.
	IsActivityStreamsArticle() bool
	// IsActivityStreamsAudio returns true if this property has a type of
	// "Audio". When true, use the GetActivityStreamsAudio and
	// SetActivityStreamsAudio methods to access and set this property.
	IsActivityStreamsAudio() bool
	// IsActivityStreamsBlock returns true if this property has a type of
	// "Block". When true, use the GetActivityStreamsBlock and
	// SetActivityStreamsBlock methods to access and set this property.
	IsActivityStreamsBlock() bool
	// IsActivityStreamsCollection returns true if this property has a type of
	// "Collection". When true, use the GetActivityStreamsCollection and
	// SetActivityStreamsCollection methods to access and set this
	// property.
	IsActivityStreamsCollection() bool
	// IsActivityStreamsCollectionPage returns true if this property has a
	// type of "CollectionPage". When true, use the
	// GetActivityStreamsCollectionPage and
	// SetActivityStreamsCollectionPage methods to access and set this
	// property.
	IsActivityStreamsCollectionPage() bool
	// IsActivityStreamsCreate returns true if this property has a type of
	// "Create". When true, use the GetActivityStreamsCreate and
	// SetActivityStreamsCreate methods to access and set this property.
	IsActivityStreamsCreate() bool
	// IsActivityStreamsDelete returns true if this property has a type of
	// "Delete". When true, use the GetActivityStreamsDelete and
	// SetActivityStreamsDelete methods to access and set this property.
	IsActivityStreamsDelete() bool
	// IsActivityStreamsDislike returns true if this property has a type of
	// "Dislike". When true, use the GetActivityStreamsDislike and
	// SetActivityStreamsDislike methods to access and set this property.
	IsActivityStreamsDislike() bool
	// IsActivityStreamsDocument returns true if this property has a type of
	// "Document". When true, use the GetActivityStreamsDocument and
	// SetActivityStreamsDocument methods to access and set this property.
	IsActivityStreamsDocument() bool
	// IsActivityStreamsEvent returns true if this property has a type of
	// "Event". When true, use the GetActivityStreamsEvent and
	// SetActivityStreamsEvent methods to access and set this property.
	IsActivityStreamsEvent() bool
	// IsActivityStreamsFlag returns true if this property has a type of
	// "Flag". When true, use the GetActivityStreamsFlag and
	// SetActivityStreamsFlag methods to access and set this property.
	IsActivityStreamsFlag() bool
	// IsActivityStreamsFollow returns true if this property has a type of
	// "Follow". When true, use the GetActivityStreamsFollow and
	// SetActivityStreamsFollow methods to access and set this property.
	IsActivityStreamsFollow() bool
	// IsActivityStreamsGroup returns true if this property has a type of
	// "Group". When true, use the GetActivityStreamsGroup and
	// SetActivityStreamsGroup methods to access and set this property.
	IsActivityStreamsGroup() bool
	// IsActivityStreamsIgnore returns true if this property has a type of
	// "Ignore". When true, use the GetActivityStreamsIgnore and
	// SetActivityStreamsIgnore methods to access and set this property.
	IsActivityStreamsIgnore() bool
	// IsActivityStreamsImage returns true if this property has a type of
	// "Image". When true, use the GetActivityStreamsImage and
	// SetActivityStreamsImage methods to access and set this property.
	IsActivityStreamsImage() bool
	// IsActivityStreamsIntransitiveActivity returns true if this property has
	// a type of "IntransitiveActivity". When true, use the
	// GetActivityStreamsIntransitiveActivity and
	// SetActivityStreamsIntransitiveActivity methods to access and set
	// this property.
	IsActivityStreamsIntransitiveActivity() bool
	// IsActivityStreamsInvite returns true if this property has a type of
	// "Invite". When true, use the GetActivityStreamsInvite and
	// SetActivityStreamsInvite methods to access and set this property.
	IsActivityStreamsInvite() bool
	// IsActivityStreamsJoin returns true if this property has a type of
	// "Join". When true, use the GetActivityStreamsJoin and
	// SetActivityStreamsJoin methods to access and set this property.
	IsActivityStreamsJoin() bool
	// IsActivityStreamsLeave returns true if this property has a type of
	// "Leave". When true, use the GetActivityStreamsLeave and
	// SetActivityStreamsLeave methods to access and set this property.
	IsActivityStreamsLeave() bool
	// IsActivityStreamsLike returns true if this property has a type of
	// "Like". When true, use the GetActivityStreamsLike and
	// SetActivityStreamsLike methods to access and set this property.
	IsActivityStreamsLike() bool
	// IsActivityStreamsListen returns true if this property has a type of
	// "Listen". When true, use the GetActivityStreamsListen and
	// SetActivityStreamsListen methods to access and set this property.
	IsActivityStreamsListen() bool
	// IsActivityStreamsMove returns true if this property has a type of
	// "Move". When true, use the GetActivityStreamsMove and
	// SetActivityStreamsMove methods to access and set this property.
	IsActivityStreamsMove() bool
	// IsActivityStreamsNote returns true if this property has a type of
	// "Note". When true, use the GetActivityStreamsNote and
	// SetActivityStreamsNote methods to access and set this property.
	IsActivityStreamsNote() bool
	// IsActivityStreamsObject returns true if this property has a type of
	// "Object". When true, use the GetActivityStreamsObject and
	// SetActivityStreamsObject methods to access and set this property.
	IsActivityStreamsObject() bool
	// IsActivityStreamsOffer returns true if this property has a type of
	// "Offer". When true, use the GetActivityStreamsOffer and
	// SetActivityStreamsOffer methods to access and set this property.
	IsActivityStreamsOffer() bool
	// IsActivityStreamsOrderedCollection returns true if this property has a
	// type of "OrderedCollection". When true, use the
	// GetActivityStreamsOrderedCollection and
	// SetActivityStreamsOrderedCollection methods to access and set this
	// property.
	IsActivityStreamsOrderedCollection() bool
	// IsActivityStreamsOrderedCollectionPage returns true if this property
	// has a type of "OrderedCollectionPage". When true, use the
	// GetActivityStreamsOrderedCollectionPage and
	// SetActivityStreamsOrderedCollectionPage methods to access and set
	// this property.
	IsActivityStreamsOrderedCollectionPage() bool
	// IsActivityStreamsOrganization returns true if this property has a type
	// of "Organization". When true, use the
	// GetActivityStreamsOrganization and SetActivityStreamsOrganization
	// methods to access and set this property.
	IsActivityStreamsOrganization() bool
	// IsActivityStreamsPage returns true if this property has a type of
	// "Page". When true, use the GetActivityStreamsPage and
	// SetActivityStreamsPage methods to access and set this property.
	IsActivityStreamsPage() bool
	// IsActivityStreamsPerson returns true if this property has a type of
	// "Person". When true, use the GetActivityStreamsPerson and
	// SetActivityStreamsPerson methods to access and set this property.
	IsActivityStreamsPerson() bool
	// IsActivityStreamsPlace returns true if this property has a type of
	// "Place". When true, use the GetActivityStreamsPlace and
	// SetActivityStreamsPlace methods to access and set this property.
	IsActivityStreamsPlace() bool
	// IsActivityStreamsProfile returns true if this property has a type of
	// "Profile". When true, use the GetActivityStreamsProfile and
	// SetActivityStreamsProfile methods to access and set this property.
	IsActivityStreamsProfile() bool
	// IsActivityStreamsQuestion returns true if this property has a type of
	// "Question". When true, use the GetActivityStreamsQuestion and
	// SetActivityStreamsQuestion methods to access and set this property.
	IsActivityStreamsQuestion() bool
	// IsActivityStreamsRead returns true if this property has a type of
	// "Read". When true, use the GetActivityStreamsRead and
	// SetActivityStreamsRead methods to access and set this property.
	IsActivityStreamsRead() bool
	// IsActivityStreamsReject returns true if this property has a type of
	// "Reject". When true, use the GetActivityStreamsReject and
	// SetActivityStreamsReject methods to access and set this property.
	IsActivityStreamsReject() bool
	// IsActivityStreamsRelationship returns true if this property has a type
	// of "Relationship". When true, use the
	// GetActivityStreamsRelationship and SetActivityStreamsRelationship
	// methods to access and set this property.
	IsActivityStreamsRelationship() bool
	// IsActivityStreamsRemove returns true if this property has a type of
	// "Remove". When true, use the GetActivityStreamsRemove and
	// SetActivityStreamsRemove methods to access and set this property.
	IsActivityStreamsRemove() bool
	// IsActivityStreamsService returns true if this property has a type of
	// "Service". When true, use the GetActivityStreamsService and
	// SetActivityStreamsService methods to access and set this property.
	IsActivityStreamsService() bool
	// IsActivityStreamsTentativeAccept returns true if this property has a
	// type of "TentativeAccept". When true, use the
	// GetActivityStreamsTentativeAccept and
	// SetActivityStreamsTentativeAccept methods to access and set this
	// property.
	IsActivityStreamsTentativeAccept() bool
	// IsActivityStreamsTentativeReject returns true if this property has a
	// type of "TentativeReject". When true, use the
	// GetActivityStreamsTentativeReject and
	// SetActivityStreamsTentativeReject methods to access and set this
	// property.
	IsActivityStreamsTentativeReject() bool
	// IsActivityStreamsTombstone returns true if this property has a type of
	// "Tombstone". When true, use the GetActivityStreamsTombstone and
	// SetActivityStreamsTombstone methods to access and set this property.
	IsActivityStreamsTombstone() bool
	// IsActivityStreamsTravel returns true if this property has a type of
	// "Travel". When true, use the GetActivityStreamsTravel and
	// SetActivityStreamsTravel methods to access and set this property.
	IsActivityStreamsTravel() bool
	// IsActivityStreamsUndo returns true if this property has a type of
	// "Undo". When true, use the GetActivityStreamsUndo and
	// SetActivityStreamsUndo methods to access and set this property.
	IsActivityStreamsUndo() bool
	// IsActivityStreamsUpdate returns true if this property has a type of
	// "Update". When true, use the GetActivityStreamsUpdate and
	// SetActivityStreamsUpdate methods to access and set this property.
	IsActivityStreamsUpdate() bool
	// IsActivityStreamsVideo returns true if this property has a type of
	// "Video". When true, use the GetActivityStreamsVideo and
	// SetActivityStreamsVideo methods to access and set this property.
	IsActivityStreamsVideo() bool
	// IsActivityStreamsView returns true if this property has a type of
	// "View". When true, use the GetActivityStreamsView and
	// SetActivityStreamsView methods to access and set this property.
	IsActivityStreamsView() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsDescribesProperty) bool
	// Name returns the name of this property: "describes".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsAccept sets the value of this property. Calling
	// IsActivityStreamsAccept afterwards returns true.
	SetActivityStreamsAccept(v ActivityStreamsAccept)
	// SetActivityStreamsActivity sets the value of this property. Calling
	// IsActivityStreamsActivity afterwards returns true.
	SetActivityStreamsActivity(v ActivityStreamsActivity)
	// SetActivityStreamsAdd sets the value of this property. Calling
	// IsActivityStreamsAdd afterwards returns true.
	SetActivityStreamsAdd(v ActivityStreamsAdd)
	// SetActivityStreamsAnnounce sets the value of this property. Calling
	// IsActivityStreamsAnnounce afterwards returns true.
	SetActivityStreamsAnnounce(v ActivityStreamsAnnounce)
	// SetActivityStreamsApplication sets the value of this property. Calling
	// IsActivityStreamsApplication afterwards returns true.
	SetActivityStreamsApplication(v ActivityStreamsApplication)
	// SetActivityStreamsArrive sets the value of this property. Calling
	// IsActivityStreamsArrive afterwards returns true.
	SetActivityStreamsArrive(v ActivityStreamsArrive)
	// SetActivityStreamsArticle sets the value of this property. Calling
	// IsActivityStreamsArticle afterwards returns true.
	SetActivityStreamsArticle(v ActivityStreamsArticle)
	// SetActivityStreamsAudio sets the value of this property. Calling
	// IsActivityStreamsAudio afterwards returns true.
	SetActivityStreamsAudio(v ActivityStreamsAudio)
	// SetActivityStreamsBlock sets the value of this property. Calling
	// IsActivityStreamsBlock afterwards returns true.
	SetActivityStreamsBlock(v ActivityStreamsBlock)
	// SetActivityStreamsCollection sets the value of this property. Calling
	// IsActivityStreamsCollection afterwards returns true.
	SetActivityStreamsCollection(v ActivityStreamsCollection)
	// SetActivityStreamsCollectionPage sets the value of this property.
	// Calling IsActivityStreamsCollectionPage afterwards returns true.
	SetActivityStreamsCollectionPage(v ActivityStreamsCollectionPage)
	// SetActivityStreamsCreate sets the value of this property. Calling
	// IsActivityStreamsCreate afterwards returns true.
	SetActivityStreamsCreate(v ActivityStreamsCreate)
	// SetActivityStreamsDelete sets the value of this property. Calling
	// IsActivityStreamsDelete afterwards returns true.
	SetActivityStreamsDelete(v ActivityStreamsDelete)
	// SetActivityStreamsDislike sets the value of this property. Calling
	// IsActivityStreamsDislike afterwards returns true.
	SetActivityStreamsDislike(v ActivityStreamsDislike)
	// SetActivityStreamsDocument sets the value of this property. Calling
	// IsActivityStreamsDocument afterwards returns true.
	SetActivityStreamsDocument(v ActivityStreamsDocument)
	// SetActivityStreamsEvent sets the value of this property. Calling
	// IsActivityStreamsEvent afterwards returns true.
	SetActivityStreamsEvent(v ActivityStreamsEvent)
	// SetActivityStreamsFlag sets the value of this property. Calling
	// IsActivityStreamsFlag afterwards returns true.
	SetActivityStreamsFlag(v ActivityStreamsFlag)
	// SetActivityStreamsFollow sets the value of this property. Calling
	// IsActivityStreamsFollow afterwards returns true.
	SetActivityStreamsFollow(v ActivityStreamsFollow)
	// SetActivityStreamsGroup sets the value of this property. Calling
	// IsActivityStreamsGroup afterwards returns true.
	SetActivityStreamsGroup(v ActivityStreamsGroup)
	// SetActivityStreamsIgnore sets the value of this property. Calling
	// IsActivityStreamsIgnore afterwards returns true.
	SetActivityStreamsIgnore(v ActivityStreamsIgnore)
	// SetActivityStreamsImage sets the value of this property. Calling
	// IsActivityStreamsImage afterwards returns true.
	SetActivityStreamsImage(v ActivityStreamsImage)
	// SetActivityStreamsIntransitiveActivity sets the value of this property.
	// Calling IsActivityStreamsIntransitiveActivity afterwards returns
	// true.
	SetActivityStreamsIntransitiveActivity(v ActivityStreamsIntransitiveActivity)
	// SetActivityStreamsInvite sets the value of this property. Calling
	// IsActivityStreamsInvite afterwards returns true.
	SetActivityStreamsInvite(v ActivityStreamsInvite)
	// SetActivityStreamsJoin sets the value of this property. Calling
	// IsActivityStreamsJoin afterwards returns true.
	SetActivityStreamsJoin(v ActivityStreamsJoin)
	// SetActivityStreamsLeave sets the value of this property. Calling
	// IsActivityStreamsLeave afterwards returns true.
	SetActivityStreamsLeave(v ActivityStreamsLeave)
	// SetActivityStreamsLike sets the value of this property. Calling
	// IsActivityStreamsLike afterwards returns true.
	SetActivityStreamsLike(v ActivityStreamsLike)
	// SetActivityStreamsListen sets the value of this property. Calling
	// IsActivityStreamsListen afterwards returns true.
	SetActivityStreamsListen(v ActivityStreamsListen)
	// SetActivityStreamsMove sets the value of this property. Calling
	// IsActivityStreamsMove afterwards returns true.
	SetActivityStreamsMove(v ActivityStreamsMove)
	// SetActivityStreamsNote sets the value of this property. Calling
	// IsActivityStreamsNote afterwards returns true.
	SetActivityStreamsNote(v ActivityStreamsNote)
	// SetActivityStreamsObject sets the value of this property. Calling
	// IsActivityStreamsObject afterwards returns true.
	SetActivityStreamsObject(v ActivityStreamsObject)
	// SetActivityStreamsOffer sets the value of this property. Calling
	// IsActivityStreamsOffer afterwards returns true.
	SetActivityStreamsOffer(v ActivityStreamsOffer)
	// SetActivityStreamsOrderedCollection sets the value of this property.
	// Calling IsActivityStreamsOrderedCollection afterwards returns true.
	SetActivityStreamsOrderedCollection(v ActivityStreamsOrderedCollection)
	// SetActivityStreamsOrderedCollectionPage sets the value of this
	// property. Calling IsActivityStreamsOrderedCollectionPage afterwards
	// returns true.
	SetActivityStreamsOrderedCollectionPage(v ActivityStreamsOrderedCollectionPage)
	// SetActivityStreamsOrganization sets the value of this property. Calling
	// IsActivityStreamsOrganization afterwards returns true.
	SetActivityStreamsOrganization(v ActivityStreamsOrganization)
	// SetActivityStreamsPage sets the value of this property. Calling
	// IsActivityStreamsPage afterwards returns true.
	SetActivityStreamsPage(v ActivityStreamsPage)
	// SetActivityStreamsPerson sets the value of this property. Calling
	// IsActivityStreamsPerson afterwards returns true.
	SetActivityStreamsPerson(v ActivityStreamsPerson)
	// SetActivityStreamsPlace sets the value of this property. Calling
	// IsActivityStreamsPlace afterwards returns true.
	SetActivityStreamsPlace(v ActivityStreamsPlace)
	// SetActivityStreamsProfile sets the value of this property. Calling
	// IsActivityStreamsProfile afterwards returns true.
	SetActivityStreamsProfile(v ActivityStreamsProfile)
	// SetActivityStreamsQuestion sets the value of this property. Calling
	// IsActivityStreamsQuestion afterwards returns true.
	SetActivityStreamsQuestion(v ActivityStreamsQuestion)
	// SetActivityStreamsRead sets the value of this property. Calling
	// IsActivityStreamsRead afterwards returns true.
	SetActivityStreamsRead(v ActivityStreamsRead)
	// SetActivityStreamsReject sets the value of this property. Calling
	// IsActivityStreamsReject afterwards returns true.
	SetActivityStreamsReject(v ActivityStreamsReject)
	// SetActivityStreamsRelationship sets the value of this property. Calling
	// IsActivityStreamsRelationship afterwards returns true.
	SetActivityStreamsRelationship(v ActivityStreamsRelationship)
	// SetActivityStreamsRemove sets the value of this property. Calling
	// IsActivityStreamsRemove afterwards returns true.
	SetActivityStreamsRemove(v ActivityStreamsRemove)
	// SetActivityStreamsService sets the value of this property. Calling
	// IsActivityStreamsService afterwards returns true.
	SetActivityStreamsService(v ActivityStreamsService)
	// SetActivityStreamsTentativeAccept sets the value of this property.
	// Calling IsActivityStreamsTentativeAccept afterwards returns true.
	SetActivityStreamsTentativeAccept(v ActivityStreamsTentativeAccept)
	// SetActivityStreamsTentativeReject sets the value of this property.
	// Calling IsActivityStreamsTentativeReject afterwards returns true.
	SetActivityStreamsTentativeReject(v ActivityStreamsTentativeReject)
	// SetActivityStreamsTombstone sets the value of this property. Calling
	// IsActivityStreamsTombstone afterwards returns true.
	SetActivityStreamsTombstone(v ActivityStreamsTombstone)
	// SetActivityStreamsTravel sets the value of this property. Calling
	// IsActivityStreamsTravel afterwards returns true.
	SetActivityStreamsTravel(v ActivityStreamsTravel)
	// SetActivityStreamsUndo sets the value of this property. Calling
	// IsActivityStreamsUndo afterwards returns true.
	SetActivityStreamsUndo(v ActivityStreamsUndo)
	// SetActivityStreamsUpdate sets the value of this property. Calling
	// IsActivityStreamsUpdate afterwards returns true.
	SetActivityStreamsUpdate(v ActivityStreamsUpdate)
	// SetActivityStreamsVideo sets the value of this property. Calling
	// IsActivityStreamsVideo afterwards returns true.
	SetActivityStreamsVideo(v ActivityStreamsVideo)
	// SetActivityStreamsView sets the value of this property. Calling
	// IsActivityStreamsView afterwards returns true.
	SetActivityStreamsView(v ActivityStreamsView)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
}
