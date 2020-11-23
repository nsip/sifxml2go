package sifxml


type WellbeingPersonLinks []WellbeingPersonLink

    type WellbeingPersonLink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      WellbeingEventRefId *String `xml:"WellbeingEventRefId" json:"WellbeingEventRefId"`
      WellbeingResponseRefId *String `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      GroupId *LocalIdType `xml:"GroupId" json:"GroupId"`
      PersonRefId *WellbeingPersonLink_PersonRefId
      ShortName *String `xml:"ShortName,omitempty" json:"ShortName"`
      HowInvolved *String `xml:"HowInvolved,omitempty" json:"HowInvolved"`
      OtherPersonId *LocalIdType `xml:"OtherPersonId,omitempty" json:"OtherPersonId"`
      OtherPersonContactDetails *String `xml:"OtherPersonContactDetails,omitempty" json:"OtherPersonContactDetails"`
      PersonRole *String `xml:"PersonRole,omitempty" json:"PersonRole"`
      FollowUpActionList *FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type WellbeingPersonLink_PersonRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
