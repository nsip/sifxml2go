package sifxml


type WellbeingPersonLinks []WellbeingPersonLink

    type WellbeingPersonLink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      WellbeingEventRefId *string `xml:"WellbeingEventRefId" json:"WellbeingEventRefId"`
      WellbeingResponseRefId *string `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      GroupId *LocalIdType `xml:"GroupId" json:"GroupId"`
      PersonRefId *WellbeingPersonLink_PersonRefId
      ShortName *string `xml:"ShortName,omitempty" json:"ShortName"`
      HowInvolved *string `xml:"HowInvolved,omitempty" json:"HowInvolved"`
      OtherPersonId *LocalIdType `xml:"OtherPersonId" json:"OtherPersonId"`
      OtherPersonContactDetails *string `xml:"OtherPersonContactDetails" json:"OtherPersonContactDetails"`
      PersonRole *string `xml:"PersonRole" json:"PersonRole"`
      FollowUpActionList *FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type WellbeingPersonLink_PersonRefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
