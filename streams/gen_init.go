package streams

import (
	propertyaccuracy "github.com/go-fed/activity/streams/impl/activitystreams/property_accuracy"
	propertyactor "github.com/go-fed/activity/streams/impl/activitystreams/property_actor"
	propertyaltitude "github.com/go-fed/activity/streams/impl/activitystreams/property_altitude"
	propertyanyof "github.com/go-fed/activity/streams/impl/activitystreams/property_anyof"
	propertyattachment "github.com/go-fed/activity/streams/impl/activitystreams/property_attachment"
	propertyattributedto "github.com/go-fed/activity/streams/impl/activitystreams/property_attributedto"
	propertyaudience "github.com/go-fed/activity/streams/impl/activitystreams/property_audience"
	propertybcc "github.com/go-fed/activity/streams/impl/activitystreams/property_bcc"
	propertybto "github.com/go-fed/activity/streams/impl/activitystreams/property_bto"
	propertycc "github.com/go-fed/activity/streams/impl/activitystreams/property_cc"
	propertyclosed "github.com/go-fed/activity/streams/impl/activitystreams/property_closed"
	propertycontent "github.com/go-fed/activity/streams/impl/activitystreams/property_content"
	propertycontext "github.com/go-fed/activity/streams/impl/activitystreams/property_context"
	propertycurrent "github.com/go-fed/activity/streams/impl/activitystreams/property_current"
	propertydeleted "github.com/go-fed/activity/streams/impl/activitystreams/property_deleted"
	propertydescribes "github.com/go-fed/activity/streams/impl/activitystreams/property_describes"
	propertyduration "github.com/go-fed/activity/streams/impl/activitystreams/property_duration"
	propertyendtime "github.com/go-fed/activity/streams/impl/activitystreams/property_endtime"
	propertyfirst "github.com/go-fed/activity/streams/impl/activitystreams/property_first"
	propertyfollowers "github.com/go-fed/activity/streams/impl/activitystreams/property_followers"
	propertyfollowing "github.com/go-fed/activity/streams/impl/activitystreams/property_following"
	propertyformertype "github.com/go-fed/activity/streams/impl/activitystreams/property_formertype"
	propertygenerator "github.com/go-fed/activity/streams/impl/activitystreams/property_generator"
	propertyheight "github.com/go-fed/activity/streams/impl/activitystreams/property_height"
	propertyhref "github.com/go-fed/activity/streams/impl/activitystreams/property_href"
	propertyhreflang "github.com/go-fed/activity/streams/impl/activitystreams/property_hreflang"
	propertyicon "github.com/go-fed/activity/streams/impl/activitystreams/property_icon"
	propertyid "github.com/go-fed/activity/streams/impl/activitystreams/property_id"
	propertyimage "github.com/go-fed/activity/streams/impl/activitystreams/property_image"
	propertyinbox "github.com/go-fed/activity/streams/impl/activitystreams/property_inbox"
	propertyinreplyto "github.com/go-fed/activity/streams/impl/activitystreams/property_inreplyto"
	propertyinstrument "github.com/go-fed/activity/streams/impl/activitystreams/property_instrument"
	propertyitems "github.com/go-fed/activity/streams/impl/activitystreams/property_items"
	propertylast "github.com/go-fed/activity/streams/impl/activitystreams/property_last"
	propertylatitude "github.com/go-fed/activity/streams/impl/activitystreams/property_latitude"
	propertyliked "github.com/go-fed/activity/streams/impl/activitystreams/property_liked"
	propertylikes "github.com/go-fed/activity/streams/impl/activitystreams/property_likes"
	propertylocation "github.com/go-fed/activity/streams/impl/activitystreams/property_location"
	propertylongitude "github.com/go-fed/activity/streams/impl/activitystreams/property_longitude"
	propertymediatype "github.com/go-fed/activity/streams/impl/activitystreams/property_mediatype"
	propertyname "github.com/go-fed/activity/streams/impl/activitystreams/property_name"
	propertynext "github.com/go-fed/activity/streams/impl/activitystreams/property_next"
	propertyobject "github.com/go-fed/activity/streams/impl/activitystreams/property_object"
	propertyoneof "github.com/go-fed/activity/streams/impl/activitystreams/property_oneof"
	propertyordereditems "github.com/go-fed/activity/streams/impl/activitystreams/property_ordereditems"
	propertyorigin "github.com/go-fed/activity/streams/impl/activitystreams/property_origin"
	propertyoutbox "github.com/go-fed/activity/streams/impl/activitystreams/property_outbox"
	propertypartof "github.com/go-fed/activity/streams/impl/activitystreams/property_partof"
	propertypreferredusername "github.com/go-fed/activity/streams/impl/activitystreams/property_preferredusername"
	propertyprev "github.com/go-fed/activity/streams/impl/activitystreams/property_prev"
	propertypreview "github.com/go-fed/activity/streams/impl/activitystreams/property_preview"
	propertypublished "github.com/go-fed/activity/streams/impl/activitystreams/property_published"
	propertyradius "github.com/go-fed/activity/streams/impl/activitystreams/property_radius"
	propertyrel "github.com/go-fed/activity/streams/impl/activitystreams/property_rel"
	propertyrelationship "github.com/go-fed/activity/streams/impl/activitystreams/property_relationship"
	propertyreplies "github.com/go-fed/activity/streams/impl/activitystreams/property_replies"
	propertyresult "github.com/go-fed/activity/streams/impl/activitystreams/property_result"
	propertyshares "github.com/go-fed/activity/streams/impl/activitystreams/property_shares"
	propertystartindex "github.com/go-fed/activity/streams/impl/activitystreams/property_startindex"
	propertystarttime "github.com/go-fed/activity/streams/impl/activitystreams/property_starttime"
	propertystreams "github.com/go-fed/activity/streams/impl/activitystreams/property_streams"
	propertysubject "github.com/go-fed/activity/streams/impl/activitystreams/property_subject"
	propertysummary "github.com/go-fed/activity/streams/impl/activitystreams/property_summary"
	propertytag "github.com/go-fed/activity/streams/impl/activitystreams/property_tag"
	propertytarget "github.com/go-fed/activity/streams/impl/activitystreams/property_target"
	propertyto "github.com/go-fed/activity/streams/impl/activitystreams/property_to"
	propertytotalitems "github.com/go-fed/activity/streams/impl/activitystreams/property_totalitems"
	propertytype "github.com/go-fed/activity/streams/impl/activitystreams/property_type"
	propertyunits "github.com/go-fed/activity/streams/impl/activitystreams/property_units"
	propertyupdated "github.com/go-fed/activity/streams/impl/activitystreams/property_updated"
	propertyurl "github.com/go-fed/activity/streams/impl/activitystreams/property_url"
	propertywidth "github.com/go-fed/activity/streams/impl/activitystreams/property_width"
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	propertyowner "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_owner"
	propertypublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickey"
	propertypublickeypem "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickeypem"
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
)

var mgr *Manager

// init handles the 'magic' of creating a Manager and dependency-injecting it into
// every other code-generated package. This gives the implementations access
// to create any type needed to deserialize, without relying on the other
// specific concrete implementations. In order to replace a go-fed created
// type with your own, be sure to have the manager call your own
// implementation's deserialize functions instead of the built-in type.
// Finally, each implementation views the Manager as an interface with only a
// subset of funcitons available. This means this Manager implements the union
// of those interfaces.
func init() {
	mgr = &Manager{}
	propertyaccuracy.SetManager(mgr)
	propertyactor.SetManager(mgr)
	propertyaltitude.SetManager(mgr)
	propertyanyof.SetManager(mgr)
	propertyattachment.SetManager(mgr)
	propertyattributedto.SetManager(mgr)
	propertyaudience.SetManager(mgr)
	propertybcc.SetManager(mgr)
	propertybto.SetManager(mgr)
	propertycc.SetManager(mgr)
	propertyclosed.SetManager(mgr)
	propertycontent.SetManager(mgr)
	propertycontext.SetManager(mgr)
	propertycurrent.SetManager(mgr)
	propertydeleted.SetManager(mgr)
	propertydescribes.SetManager(mgr)
	propertyduration.SetManager(mgr)
	propertyendtime.SetManager(mgr)
	propertyfirst.SetManager(mgr)
	propertyfollowers.SetManager(mgr)
	propertyfollowing.SetManager(mgr)
	propertyformertype.SetManager(mgr)
	propertygenerator.SetManager(mgr)
	propertyheight.SetManager(mgr)
	propertyhref.SetManager(mgr)
	propertyhreflang.SetManager(mgr)
	propertyicon.SetManager(mgr)
	propertyid.SetManager(mgr)
	propertyimage.SetManager(mgr)
	propertyinbox.SetManager(mgr)
	propertyinreplyto.SetManager(mgr)
	propertyinstrument.SetManager(mgr)
	propertyitems.SetManager(mgr)
	propertylast.SetManager(mgr)
	propertylatitude.SetManager(mgr)
	propertyliked.SetManager(mgr)
	propertylikes.SetManager(mgr)
	propertylocation.SetManager(mgr)
	propertylongitude.SetManager(mgr)
	propertymediatype.SetManager(mgr)
	propertyname.SetManager(mgr)
	propertynext.SetManager(mgr)
	propertyobject.SetManager(mgr)
	propertyoneof.SetManager(mgr)
	propertyordereditems.SetManager(mgr)
	propertyorigin.SetManager(mgr)
	propertyoutbox.SetManager(mgr)
	propertypartof.SetManager(mgr)
	propertypreferredusername.SetManager(mgr)
	propertyprev.SetManager(mgr)
	propertypreview.SetManager(mgr)
	propertypublished.SetManager(mgr)
	propertyradius.SetManager(mgr)
	propertyrel.SetManager(mgr)
	propertyrelationship.SetManager(mgr)
	propertyreplies.SetManager(mgr)
	propertyresult.SetManager(mgr)
	propertyshares.SetManager(mgr)
	propertystartindex.SetManager(mgr)
	propertystarttime.SetManager(mgr)
	propertystreams.SetManager(mgr)
	propertysubject.SetManager(mgr)
	propertysummary.SetManager(mgr)
	propertytag.SetManager(mgr)
	propertytarget.SetManager(mgr)
	propertyto.SetManager(mgr)
	propertytotalitems.SetManager(mgr)
	propertytype.SetManager(mgr)
	propertyunits.SetManager(mgr)
	propertyupdated.SetManager(mgr)
	propertyurl.SetManager(mgr)
	propertywidth.SetManager(mgr)
	typeaccept.SetManager(mgr)
	typeactivity.SetManager(mgr)
	typeadd.SetManager(mgr)
	typeannounce.SetManager(mgr)
	typeapplication.SetManager(mgr)
	typearrive.SetManager(mgr)
	typearticle.SetManager(mgr)
	typeaudio.SetManager(mgr)
	typeblock.SetManager(mgr)
	typecollection.SetManager(mgr)
	typecollectionpage.SetManager(mgr)
	typecreate.SetManager(mgr)
	typedelete.SetManager(mgr)
	typedislike.SetManager(mgr)
	typedocument.SetManager(mgr)
	typeevent.SetManager(mgr)
	typeflag.SetManager(mgr)
	typefollow.SetManager(mgr)
	typegroup.SetManager(mgr)
	typeignore.SetManager(mgr)
	typeimage.SetManager(mgr)
	typeintransitiveactivity.SetManager(mgr)
	typeinvite.SetManager(mgr)
	typejoin.SetManager(mgr)
	typeleave.SetManager(mgr)
	typelike.SetManager(mgr)
	typelink.SetManager(mgr)
	typelisten.SetManager(mgr)
	typemention.SetManager(mgr)
	typemove.SetManager(mgr)
	typenote.SetManager(mgr)
	typeobject.SetManager(mgr)
	typeoffer.SetManager(mgr)
	typeorderedcollection.SetManager(mgr)
	typeorderedcollectionpage.SetManager(mgr)
	typeorganization.SetManager(mgr)
	typepage.SetManager(mgr)
	typeperson.SetManager(mgr)
	typeplace.SetManager(mgr)
	typeprofile.SetManager(mgr)
	typequestion.SetManager(mgr)
	typeread.SetManager(mgr)
	typereject.SetManager(mgr)
	typerelationship.SetManager(mgr)
	typeremove.SetManager(mgr)
	typeservice.SetManager(mgr)
	typetentativeaccept.SetManager(mgr)
	typetentativereject.SetManager(mgr)
	typetombstone.SetManager(mgr)
	typetravel.SetManager(mgr)
	typeundo.SetManager(mgr)
	typeupdate.SetManager(mgr)
	typevideo.SetManager(mgr)
	typeview.SetManager(mgr)
	propertyowner.SetManager(mgr)
	propertypublickey.SetManager(mgr)
	propertypublickeypem.SetManager(mgr)
	typepublickey.SetManager(mgr)
	typeaccept.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeactivity.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeadd.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeannounce.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeapplication.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typearrive.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typearticle.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeaudio.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeblock.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typecollection.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typecollectionpage.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typecreate.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typedelete.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typedislike.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typedocument.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeevent.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeflag.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typefollow.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typegroup.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeignore.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeimage.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeintransitiveactivity.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeinvite.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typejoin.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeleave.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typelike.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typelink.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typelisten.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typemention.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typemove.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typenote.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeobject.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeoffer.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeorderedcollection.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeorderedcollectionpage.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeorganization.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typepage.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeperson.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeplace.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeprofile.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typequestion.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeread.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typereject.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typerelationship.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeremove.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeservice.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typetentativeaccept.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typetentativereject.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typetombstone.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typetravel.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeundo.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeupdate.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typevideo.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typeview.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
	typepublickey.SetTypePropertyConstructor(NewActivityStreamsTypeProperty)
}
