package sifxml

import _ "github.com/creasty/defaults"


type ResourceBookings []ResourceBooking

    type ResourceBooking struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      ResourceRefId ResourceBooking_ResourceRefId
      ResourceLocalId LocalIdType `xml:"ResourceLocalId" json:"ResourceLocalId"`
      StartDateTime string `xml:"StartDateTime" json:"StartDateTime"`
      FinishDateTime string `xml:"FinishDateTime" json:"FinishDateTime"`
      FromPeriod LocalIdType `xml:"FromPeriod,omitempty" json:"FromPeriod"`
      ToPeriod LocalIdType `xml:"ToPeriod,omitempty" json:"ToPeriod"`
      Booker string `xml:"Booker" json:"Booker"`
      Reason string `xml:"Reason,omitempty" json:"Reason"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      KeepOld bool `xml:"KeepOld" json:"KeepOld"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceBooking_ResourceRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value string `xml:",chardata" json:"value"`
}
