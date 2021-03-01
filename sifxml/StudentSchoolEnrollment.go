package sifxml


type StudentSchoolEnrollments []StudentSchoolEnrollment

    type StudentSchoolEnrollment struct {
  studentschoolenrollment `xml:"StudentSchoolEnrollment" json:"StudentSchoolEnrollment"`
}

type studentschoolenrollment struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      MembershipType *AUCodeSetsSchoolEnrollmentTypeType `xml:"MembershipType" json:"MembershipType"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      TimeFrame *AUCodeSetsEnrollmentTimeFrameType `xml:"TimeFrame" json:"TimeFrame"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      EntryDate *String `xml:"EntryDate" json:"EntryDate"`
      EntryType *StudentEntryContainerType `xml:"EntryType,omitempty" json:"EntryType,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel,omitempty"`
      Homeroom *HomeroomType `xml:"Homeroom,omitempty" json:"Homeroom,omitempty"`
      Advisor *StaffRefIdType `xml:"Advisor,omitempty" json:"Advisor,omitempty"`
      Counselor *StaffRefIdType `xml:"Counselor,omitempty" json:"Counselor,omitempty"`
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup,omitempty"`
      ACARASchoolId *LocalIdType `xml:"ACARASchoolId,omitempty" json:"ACARASchoolId,omitempty"`
      ClassCode *String `xml:"ClassCode,omitempty" json:"ClassCode,omitempty"`
      TestLevel *YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel,omitempty"`
      ReportingSchool *AUCodeSetsYesOrNoCategoryType `xml:"ReportingSchool,omitempty" json:"ReportingSchool,omitempty"`
      House *String `xml:"House,omitempty" json:"House,omitempty"`
      IndividualLearningPlan *AUCodeSetsYesOrNoCategoryType `xml:"IndividualLearningPlan,omitempty" json:"IndividualLearningPlan,omitempty"`
      Calendar *StudentSchoolEnrollment_Calendar
      ExitDate *String `xml:"ExitDate,omitempty" json:"ExitDate,omitempty"`
      ExitStatus *StudentExitStatusContainerType `xml:"ExitStatus,omitempty" json:"ExitStatus,omitempty"`
      ExitType *StudentExitContainerType `xml:"ExitType,omitempty" json:"ExitType,omitempty"`
      FTE *Float `xml:"FTE,omitempty" json:"FTE,omitempty"`
      FTPTStatus *AUCodeSetsFTPTStatusCodeType `xml:"FTPTStatus,omitempty" json:"FTPTStatus,omitempty"`
      FFPOS *AUCodeSetsFFPOSStatusCodeType `xml:"FFPOS,omitempty" json:"FFPOS,omitempty"`
      CatchmentStatus *CatchmentStatusContainerType `xml:"CatchmentStatus,omitempty" json:"CatchmentStatus,omitempty"`
      RecordClosureReason *String `xml:"RecordClosureReason,omitempty" json:"RecordClosureReason,omitempty"`
      PromotionInfo *PromotionInfoType `xml:"PromotionInfo,omitempty" json:"PromotionInfo,omitempty"`
      PreviousSchool *LocalIdType `xml:"PreviousSchool,omitempty" json:"PreviousSchool,omitempty"`
      PreviousSchoolName *String `xml:"PreviousSchoolName,omitempty" json:"PreviousSchoolName,omitempty"`
      DestinationSchool *LocalIdType `xml:"DestinationSchool,omitempty" json:"DestinationSchool,omitempty"`
      DestinationSchoolName *String `xml:"DestinationSchoolName,omitempty" json:"DestinationSchoolName,omitempty"`
      StudentSubjectChoiceList *StudentSubjectChoiceListType `xml:"StudentSubjectChoiceList,omitempty" json:"StudentSubjectChoiceList,omitempty"`
      StartedAtSchoolDate *String `xml:"StartedAtSchoolDate,omitempty" json:"StartedAtSchoolDate,omitempty"`
      StudentGroupList *StudentGroupListType `xml:"StudentGroupList,omitempty" json:"StudentGroupList,omitempty"`
      PublishingPermissionList *PublishingPermissionListType `xml:"PublishingPermissionList,omitempty" json:"PublishingPermissionList,omitempty"`
      DisabilityLevelOfAdjustment *String `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment,omitempty"`
      DisabilityCategory *String `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory,omitempty"`
      CensusAge *Int `xml:"CensusAge,omitempty" json:"CensusAge,omitempty"`
      DistanceEducationStudent *AUCodeSetsYesOrNoCategoryType `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent,omitempty"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus,omitempty"`
      InternationalStudent *AUCodeSetsYesOrNoCategoryType `xml:"InternationalStudent,omitempty" json:"InternationalStudent,omitempty"`
      TravelDetails *TravelDetailsContainerType `xml:"TravelDetails,omitempty" json:"TravelDetails,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type StudentSchoolEnrollment_Calendar struct {
  studentschoolenrollment_calendar `xml:"StudentSchoolEnrollment_Calendar" json:"StudentSchoolEnrollment_Calendar"`
}

type studentschoolenrollment_calendar struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
