package sifxml



    type WellbeingEvents struct {
      wellbeingevents `json:"WellbeingEvents"`
    }

    type wellbeingevents struct {
      WellbeingEvent []wellbeingevent `json:"WellbeingEvent"`
    }

    type WellbeingEvent struct {
      wellbeingevent `xml:"WellbeingEvent" json:"WellbeingEvent"`
     }

     type wellbeingevent struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      GroupIndicator *Bool `xml:"GroupIndicator,omitempty" json:"GroupIndicator,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      EventId *LocalIdType `xml:"EventId,omitempty" json:"EventId,omitempty"`
      WellbeingEventNotes *String `xml:"WellbeingEventNotes,omitempty" json:"WellbeingEventNotes,omitempty"`
      WellbeingEventCategoryClass *AUCodeSetsWellbeingEventCategoryClassType `xml:"WellbeingEventCategoryClass" json:"WellbeingEventCategoryClass"`
      WellbeingEventCategoryList *WellbeingEventCategoryListType `xml:"WellbeingEventCategoryList,omitempty" json:"WellbeingEventCategoryList,omitempty"`
      ReportingStaffRefId *String `xml:"ReportingStaffRefId,omitempty" json:"ReportingStaffRefId,omitempty"`
      WellbeingEventLocationDetails *WellbeingEventLocationDetailsType `xml:"WellbeingEventLocationDetails,omitempty" json:"WellbeingEventLocationDetails,omitempty"`
      WellbeingEventCreationTimeStamp *String `xml:"WellbeingEventCreationTimeStamp,omitempty" json:"WellbeingEventCreationTimeStamp,omitempty"`
      WellbeingEventDate *String `xml:"WellbeingEventDate" json:"WellbeingEventDate"`
      WellbeingEventTime *String `xml:"WellbeingEventTime,omitempty" json:"WellbeingEventTime,omitempty"`
      WellbeingEventDescription *String `xml:"WellbeingEventDescription,omitempty" json:"WellbeingEventDescription,omitempty"`
      WellbeingEventTimePeriod *AUCodeSetsWellbeingEventTimePeriodType `xml:"WellbeingEventTimePeriod" json:"WellbeingEventTimePeriod"`
      ConfidentialFlag *AUCodeSetsYesOrNoCategoryType `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag,omitempty"`
      PersonInvolvementList *PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList,omitempty"`
      FollowUpActionList *FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

