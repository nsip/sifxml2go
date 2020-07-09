package sifxml


type ResourceBookings []ResourceBooking

    type ResourceBooking struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ResourceRefId *ResourceBooking_ResourceRefId `xml:"ResourceRefId,omitempty" json:"ResourceRefId"`
      ResourceLocalId *LocalIdType `xml:"ResourceLocalId,omitempty" json:"ResourceLocalId"`
      StartDateTime *string `xml:"StartDateTime,omitempty" json:"StartDateTime"`
      FinishDateTime *string `xml:"FinishDateTime,omitempty" json:"FinishDateTime"`
      FromPeriod *LocalIdType `xml:"FromPeriod,omitempty" json:"FromPeriod"`
      ToPeriod *LocalIdType `xml:"ToPeriod,omitempty" json:"ToPeriod"`
      Booker *string `xml:"Booker,omitempty" json:"Booker"`
      Reason *string `xml:"Reason,omitempty" json:"Reason"`
      ScheduledActivityRefId *string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      KeepOld *bool `xml:"KeepOld,omitempty" json:"KeepOld"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceBooking_ResourceRefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
