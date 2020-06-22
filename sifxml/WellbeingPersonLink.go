package sifxml


    type WellbeingPersonLink struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      WellbeingEventRefId string `xml:"WellbeingEventRefId"`
      WellbeingResponseRefId string `xml:"WellbeingResponseRefId"`
      GroupId LocalIdType `xml:"GroupId"`
      PersonRefId WellbeingPersonLink_PersonRefId `xml:"PersonRefId"`
      ShortName string `xml:"ShortName"`
      HowInvolved string `xml:"HowInvolved"`
      OtherPersonId LocalIdType `xml:"OtherPersonId"`
      OtherPersonContactDetails string `xml:"OtherPersonContactDetails"`
      PersonRole string `xml:"PersonRole"`
      FollowUpActionList FollowUpActionListType `xml:"FollowUpActionList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type WellbeingPersonLink_PersonRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
