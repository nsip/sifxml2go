package sifxml


type WellbeingPersonLinks []WellbeingPersonLink

    type WellbeingPersonLink struct {
  wellbeingpersonlink `xml:"WellbeingPersonLink" json:"WellbeingPersonLink"`
}

type wellbeingpersonlink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      WellbeingEventRefId *String `xml:"WellbeingEventRefId" json:"WellbeingEventRefId"`
      WellbeingResponseRefId *String `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId,omitempty"`
      GroupId *LocalIdType `xml:"GroupId" json:"GroupId"`
      PersonRefId *WellbeingPersonLink_PersonRefId
      ShortName *String `xml:"ShortName,omitempty" json:"ShortName,omitempty"`
      HowInvolved *String `xml:"HowInvolved,omitempty" json:"HowInvolved,omitempty"`
      OtherPersonId *LocalIdType `xml:"OtherPersonId,omitempty" json:"OtherPersonId,omitempty"`
      OtherPersonContactDetails *String `xml:"OtherPersonContactDetails,omitempty" json:"OtherPersonContactDetails,omitempty"`
      PersonRole *String `xml:"PersonRole,omitempty" json:"PersonRole,omitempty"`
      FollowUpActionList *FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type WellbeingPersonLink_PersonRefId struct {
  wellbeingpersonlink_personrefid `xml:"WellbeingPersonLink_PersonRefId" json:"WellbeingPersonLink_PersonRefId"`
}

type wellbeingpersonlink_personrefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
