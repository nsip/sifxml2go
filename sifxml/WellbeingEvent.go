package sifxml


type WellbeingEvents []WellbeingEvent

    type WellbeingEvent struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      GroupIndicator *Bool `xml:"GroupIndicator,omitempty" json:"GroupIndicator"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      EventId *LocalIdType `xml:"EventId,omitempty" json:"EventId"`
      WellbeingEventNotes *String `xml:"WellbeingEventNotes,omitempty" json:"WellbeingEventNotes"`
      WellbeingEventCategoryClass *AUCodeSetsWellbeingEventCategoryClassType `xml:"WellbeingEventCategoryClass" json:"WellbeingEventCategoryClass"`
      WellbeingEventCategoryList *WellbeingEventCategoryListType `xml:"WellbeingEventCategoryList,omitempty" json:"WellbeingEventCategoryList"`
      ReportingStaffRefId *String `xml:"ReportingStaffRefId,omitempty" json:"ReportingStaffRefId"`
      WellbeingEventLocationDetails *WellbeingEventLocationDetailsType `xml:"WellbeingEventLocationDetails,omitempty" json:"WellbeingEventLocationDetails"`
      WellbeingEventCreationTimeStamp *String `xml:"WellbeingEventCreationTimeStamp,omitempty" json:"WellbeingEventCreationTimeStamp"`
      WellbeingEventDate *String `xml:"WellbeingEventDate" json:"WellbeingEventDate"`
      WellbeingEventTime *String `xml:"WellbeingEventTime,omitempty" json:"WellbeingEventTime"`
      WellbeingEventDescription *String `xml:"WellbeingEventDescription,omitempty" json:"WellbeingEventDescription"`
      WellbeingEventTimePeriod *AUCodeSetsWellbeingEventTimePeriodType `xml:"WellbeingEventTimePeriod" json:"WellbeingEventTimePeriod"`
      ConfidentialFlag *AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag"`
      PersonInvolvementList *PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList"`
      FollowUpActionList *FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    