package sifxml


type StudentSchoolEnrollments []StudentSchoolEnrollment

    type StudentSchoolEnrollment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      MembershipType *string `xml:"MembershipType,omitempty" json:"MembershipType"`
      TimeFrame *string `xml:"TimeFrame,omitempty" json:"TimeFrame"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      EntryDate *string `xml:"EntryDate,omitempty" json:"EntryDate"`
      EntryType *StudentEntryContainerType `xml:"EntryType,omitempty" json:"EntryType"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      Homeroom *StudentSchoolEnrollment_Homeroom `xml:"Homeroom,omitempty" json:"Homeroom"`
      Advisor *StudentSchoolEnrollment_Advisor `xml:"Advisor,omitempty" json:"Advisor"`
      Counselor *StudentSchoolEnrollment_Counselor `xml:"Counselor,omitempty" json:"Counselor"`
      Homegroup *string `xml:"Homegroup,omitempty" json:"Homegroup"`
      ACARASchoolId *LocalIdType `xml:"ACARASchoolId,omitempty" json:"ACARASchoolId"`
      ClassCode *string `xml:"ClassCode,omitempty" json:"ClassCode"`
      TestLevel *YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      ReportingSchool *string `xml:"ReportingSchool,omitempty" json:"ReportingSchool"`
      House *string `xml:"House,omitempty" json:"House"`
      IndividualLearningPlan *string `xml:"IndividualLearningPlan,omitempty" json:"IndividualLearningPlan"`
      Calendar *StudentSchoolEnrollment_Calendar `xml:"Calendar,omitempty" json:"Calendar"`
      ExitDate *string `xml:"ExitDate,omitempty" json:"ExitDate"`
      ExitStatus *StudentExitStatusContainerType `xml:"ExitStatus,omitempty" json:"ExitStatus"`
      ExitType *StudentExitContainerType `xml:"ExitType,omitempty" json:"ExitType"`
      FTE *float64 `xml:"FTE,omitempty" json:"FTE"`
      FTPTStatus *string `xml:"FTPTStatus,omitempty" json:"FTPTStatus"`
      FFPOS *string `xml:"FFPOS,omitempty" json:"FFPOS"`
      CatchmentStatus *CatchmentStatusContainerType `xml:"CatchmentStatus,omitempty" json:"CatchmentStatus"`
      RecordClosureReason *string `xml:"RecordClosureReason,omitempty" json:"RecordClosureReason"`
      PromotionInfo *PromotionInfoType `xml:"PromotionInfo,omitempty" json:"PromotionInfo"`
      PreviousSchool *LocalIdType `xml:"PreviousSchool,omitempty" json:"PreviousSchool"`
      PreviousSchoolName *string `xml:"PreviousSchoolName,omitempty" json:"PreviousSchoolName"`
      DestinationSchool *LocalIdType `xml:"DestinationSchool,omitempty" json:"DestinationSchool"`
      DestinationSchoolName *string `xml:"DestinationSchoolName,omitempty" json:"DestinationSchoolName"`
      StudentSubjectChoiceList *StudentSubjectChoiceListType `xml:"StudentSubjectChoiceList,omitempty" json:"StudentSubjectChoiceList"`
      StartedAtSchoolDate *string `xml:"StartedAtSchoolDate,omitempty" json:"StartedAtSchoolDate"`
      StudentGroupList *StudentGroupListType `xml:"StudentGroupList,omitempty" json:"StudentGroupList"`
      PublishingPermissionList *PublishingPermissionListType `xml:"PublishingPermissionList,omitempty" json:"PublishingPermissionList"`
      DisabilityLevelOfAdjustment *string `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *string `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory"`
      CensusAge *float64 `xml:"CensusAge,omitempty" json:"CensusAge"`
      DistanceEducationStudent *string `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent"`
      BoardingStatus *string `xml:"BoardingStatus,omitempty" json:"BoardingStatus"`
      InternationalStudent *string `xml:"InternationalStudent,omitempty" json:"InternationalStudent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type StudentSchoolEnrollment_Homeroom struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Advisor struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Counselor struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Calendar struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
