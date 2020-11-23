package sifxml


type ResourceBookings []ResourceBooking

    type ResourceBooking struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ResourceRefId *ResourceBooking_ResourceRefId
      ResourceLocalId *LocalIdType `xml:"ResourceLocalId" json:"ResourceLocalId"`
      StartDateTime *String `xml:"StartDateTime" json:"StartDateTime"`
      FinishDateTime *String `xml:"FinishDateTime" json:"FinishDateTime"`
      FromPeriod *LocalIdType `xml:"FromPeriod,omitempty" json:"FromPeriod"`
      ToPeriod *LocalIdType `xml:"ToPeriod,omitempty" json:"ToPeriod"`
      Booker *String `xml:"Booker" json:"Booker"`
      Reason *String `xml:"Reason,omitempty" json:"Reason"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      KeepOld *Bool `xml:"KeepOld,omitempty" json:"KeepOld"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceBooking_ResourceRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
