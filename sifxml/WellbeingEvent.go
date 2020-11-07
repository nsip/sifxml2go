package sifxml

import _ "github.com/creasty/defaults"


type WellbeingEvents []WellbeingEvent

    type WellbeingEvent struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      GroupIndicator bool `xml:"GroupIndicator" json:"GroupIndicator"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      EventId LocalIdType `xml:"EventId,omitempty" json:"EventId"`
      WellbeingEventNotes string `xml:"WellbeingEventNotes,omitempty" json:"WellbeingEventNotes"`
      WellbeingEventCategoryClass AUCodeSetsWellbeingEventCategoryClassType `xml:"WellbeingEventCategoryClass" json:"WellbeingEventCategoryClass"`
      WellbeingEventCategoryList WellbeingEventCategoryListType `xml:"WellbeingEventCategoryList,omitempty" json:"WellbeingEventCategoryList"`
      ReportingStaffRefId string `xml:"ReportingStaffRefId,omitempty" json:"ReportingStaffRefId"`
      WellbeingEventLocationDetails WellbeingEventLocationDetailsType `xml:"WellbeingEventLocationDetails,omitempty" json:"WellbeingEventLocationDetails"`
      WellbeingEventCreationTimeStamp string `xml:"WellbeingEventCreationTimeStamp,omitempty" json:"WellbeingEventCreationTimeStamp"`
      WellbeingEventDate string `xml:"WellbeingEventDate" json:"WellbeingEventDate"`
      WellbeingEventTime string `xml:"WellbeingEventTime,omitempty" json:"WellbeingEventTime"`
      WellbeingEventDescription string `xml:"WellbeingEventDescription,omitempty" json:"WellbeingEventDescription"`
      WellbeingEventTimePeriod AUCodeSetsWellbeingEventTimePeriodType `xml:"WellbeingEventTimePeriod" json:"WellbeingEventTimePeriod"`
      ConfidentialFlag AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag"`
      PersonInvolvementList PersonInvolvementListType `xml:"PersonInvolvementList" json:"PersonInvolvementList"`
      FollowUpActionList FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList"`
      Status AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    