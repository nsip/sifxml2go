package sifxml


type StudentSchoolEnrollments []StudentSchoolEnrollment

    type StudentSchoolEnrollment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      MembershipType *AUCodeSetsSchoolEnrollmentTypeType `xml:"MembershipType" json:"MembershipType"`
      TimeFrame *AUCodeSetsEnrollmentTimeFrameType `xml:"TimeFrame" json:"TimeFrame"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      EntryDate *String `xml:"EntryDate" json:"EntryDate"`
      EntryType *StudentEntryContainerType `xml:"EntryType,omitempty" json:"EntryType"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      Homeroom *StudentSchoolEnrollment_Homeroom
      Advisor *StudentSchoolEnrollment_Advisor
      Counselor *StudentSchoolEnrollment_Counselor
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup"`
      ACARASchoolId *LocalIdType `xml:"ACARASchoolId,omitempty" json:"ACARASchoolId"`
      ClassCode *String `xml:"ClassCode,omitempty" json:"ClassCode"`
      TestLevel *YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      ReportingSchool *AUCodeSetsYesOrNoCategoryType `xml:"ReportingSchool,omitempty" json:"ReportingSchool"`
      House *String `xml:"House,omitempty" json:"House"`
      IndividualLearningPlan *AUCodeSetsYesOrNoCategoryType `xml:"IndividualLearningPlan,omitempty" json:"IndividualLearningPlan"`
      Calendar *StudentSchoolEnrollment_Calendar
      ExitDate *String `xml:"ExitDate,omitempty" json:"ExitDate"`
      ExitStatus *StudentExitStatusContainerType `xml:"ExitStatus,omitempty" json:"ExitStatus"`
      ExitType *StudentExitContainerType `xml:"ExitType,omitempty" json:"ExitType"`
      FTE *Float `xml:"FTE,omitempty" json:"FTE"`
      FTPTStatus *AUCodeSetsFTPTStatusCodeType `xml:"FTPTStatus,omitempty" json:"FTPTStatus"`
      FFPOS *AUCodeSetsFFPOSStatusCodeType `xml:"FFPOS,omitempty" json:"FFPOS"`
      CatchmentStatus *CatchmentStatusContainerType `xml:"CatchmentStatus,omitempty" json:"CatchmentStatus"`
      RecordClosureReason *String `xml:"RecordClosureReason,omitempty" json:"RecordClosureReason"`
      PromotionInfo *PromotionInfoType `xml:"PromotionInfo,omitempty" json:"PromotionInfo"`
      PreviousSchool *LocalIdType `xml:"PreviousSchool,omitempty" json:"PreviousSchool"`
      PreviousSchoolName *String `xml:"PreviousSchoolName,omitempty" json:"PreviousSchoolName"`
      DestinationSchool *LocalIdType `xml:"DestinationSchool,omitempty" json:"DestinationSchool"`
      DestinationSchoolName *String `xml:"DestinationSchoolName,omitempty" json:"DestinationSchoolName"`
      StudentSubjectChoiceList *StudentSubjectChoiceListType `xml:"StudentSubjectChoiceList,omitempty" json:"StudentSubjectChoiceList"`
      StartedAtSchoolDate *String `xml:"StartedAtSchoolDate,omitempty" json:"StartedAtSchoolDate"`
      StudentGroupList *StudentGroupListType `xml:"StudentGroupList,omitempty" json:"StudentGroupList"`
      PublishingPermissionList *PublishingPermissionListType `xml:"PublishingPermissionList,omitempty" json:"PublishingPermissionList"`
      DisabilityLevelOfAdjustment *String `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *String `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory"`
      CensusAge *Float `xml:"CensusAge,omitempty" json:"CensusAge"`
      DistanceEducationStudent *AUCodeSetsYesOrNoCategoryType `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus"`
      InternationalStudent *AUCodeSetsYesOrNoCategoryType `xml:"InternationalStudent,omitempty" json:"InternationalStudent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type StudentSchoolEnrollment_Homeroom struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Advisor struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Counselor struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type StudentSchoolEnrollment_Calendar struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
