package sifxml


type ResourceBookings []ResourceBooking

    type ResourceBooking struct {
  resourcebooking `xml:"ResourceBooking" json:"ResourceBooking"`
}

type resourcebooking struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ResourceRefId *ResourceBooking_ResourceRefId
      ResourceLocalId *LocalIdType `xml:"ResourceLocalId" json:"ResourceLocalId"`
      StartDateTime *String `xml:"StartDateTime" json:"StartDateTime"`
      FinishDateTime *String `xml:"FinishDateTime" json:"FinishDateTime"`
      FromPeriod *LocalIdType `xml:"FromPeriod,omitempty" json:"FromPeriod,omitempty"`
      ToPeriod *LocalIdType `xml:"ToPeriod,omitempty" json:"ToPeriod,omitempty"`
      Booker *String `xml:"Booker" json:"Booker"`
      Reason *String `xml:"Reason,omitempty" json:"Reason,omitempty"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId,omitempty"`
      KeepOld *Bool `xml:"KeepOld,omitempty" json:"KeepOld,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type ResourceBooking_ResourceRefId struct {
  resourcebooking_resourcerefid `xml:"ResourceBooking_ResourceRefId" json:"ResourceBooking_ResourceRefId"`
}

type resourcebooking_resourcerefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

