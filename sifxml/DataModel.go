package sifxml


    type ReportPackageType AbstractContentPackageType
    type AbstractContentPackageType struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      XMLData *AbstractContentPackageType_XMLData `xml:"XMLData" json:"XMLData"`
      TextData *AbstractContentPackageType_TextData `xml:"TextData" json:"TextData"`
      BinaryData *AbstractContentPackageType_BinaryData `xml:"BinaryData" json:"BinaryData"`
      Reference *AbstractContentPackageType_Reference `xml:"Reference" json:"Reference"`
      
      }
    
    type AbstractContentElementType struct {
      XMLData *AbstractContentElementType_XMLData `xml:"XMLData" json:"XMLData"`
      TextData *AbstractContentElementType_TextData `xml:"TextData" json:"TextData"`
      BinaryData *AbstractContentElementType_BinaryData `xml:"BinaryData" json:"BinaryData"`
      Reference *AbstractContentElementType_Reference `xml:"Reference" json:"Reference"`
      
      }
    
    type MonetaryAmountType struct {
          Currency *ISO4217CurrencyNamesAndCodeElementsType `xml:"Currency,attr" json:"Currency"`
      
        Value *float64 `xml:",chardata" json:"value"`
      }
    
    type ObjectNameType string
    type ServiceNameType string
    type ObjectType string
    type ReportDataObjectType string
    type URIOrBinaryType string
    type GUIDType string
    type MsgIdType GUIDType
    type RefIdType GUIDType
    type IdRefType RefIdType
    type VersionType string
    type VersionWithWildcardsType string
    type DefinedProtocolsType string
    type ExtendedContentType string
    type SelectedContentType string
    type LibraryTransactionListType struct {
        Transaction []LibraryTransactionType `xml:"Transaction" json:"Transaction"`
      
      }
    
    type LibraryTransactionType struct {
        ItemInfo *LibraryItemInfoType `xml:"ItemInfo,omitempty" json:"ItemInfo"`
      CheckoutInfo *CheckoutInfoType `xml:"CheckoutInfo,omitempty" json:"CheckoutInfo"`
      FineInfoList *FineInfoListType `xml:"FineInfoList,omitempty" json:"FineInfoList"`
      HoldInfoList *HoldInfoListType `xml:"HoldInfoList,omitempty" json:"HoldInfoList"`
      
      }
    
    type LibraryItemInfoType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      Title *string `xml:"Title" json:"Title"`
      Author *string `xml:"Author,omitempty" json:"Author"`
      ElectronicId *ElectronicIdType `xml:"ElectronicId,omitempty" json:"ElectronicId"`
      CallNumber *string `xml:"CallNumber,omitempty" json:"CallNumber"`
      ISBN *string `xml:"ISBN,omitempty" json:"ISBN"`
      Cost *MonetaryAmountType `xml:"Cost,omitempty" json:"Cost"`
      ReplacementCost *MonetaryAmountType `xml:"ReplacementCost,omitempty" json:"ReplacementCost"`
      
      }
    
    type CheckoutInfoType struct {
        CheckedOutOn *string `xml:"CheckedOutOn" json:"CheckedOutOn"`
      ReturnBy *string `xml:"ReturnBy" json:"ReturnBy"`
      RenewalCount *int `xml:"RenewalCount" json:"RenewalCount"`
      
      }
    
    type FineInfoListType struct {
        FineInfo []FineInfoType `xml:"FineInfo" json:"FineInfo"`
      
      }
    
    type FineInfoType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      Assessed *string `xml:"Assessed" json:"Assessed"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      Reference *string `xml:"Reference,omitempty" json:"Reference"`
      
      }
    
    type HoldInfoListType struct {
        HoldInfo []HoldInfoType `xml:"HoldInfo" json:"HoldInfo"`
      
      }
    
    type HoldInfoType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      DatePlaced *string `xml:"DatePlaced" json:"DatePlaced"`
      DateNeeded *string `xml:"DateNeeded,omitempty" json:"DateNeeded"`
      ReservationExpiry *string `xml:"ReservationExpiry,omitempty" json:"ReservationExpiry"`
      MadeAvailable *string `xml:"MadeAvailable,omitempty" json:"MadeAvailable"`
      Expires *string `xml:"Expires,omitempty" json:"Expires"`
      
      }
    
    type LibraryMessageListType struct {
        Message []LibraryMessageType `xml:"Message" json:"Message"`
      
      }
    
    type LibraryMessageType struct {
        Priority *string `xml:"Priority,attr" json:"Priority"`
      PriorityCodeset *string `xml:"PriorityCodeset,attr" json:"PriorityCodeset"`
      Sent *string `xml:"Sent,omitempty" json:"Sent"`
      Text *string `xml:"Text" json:"Text"`
      
      }
    
    type StudentAttendanceCollectionReportingListType struct {
        StudentAttendanceCollectionReporting []StudentAttendanceCollectionReportingType `xml:"StudentAttendanceCollectionReporting" json:"StudentAttendanceCollectionReporting"`
      
      }
    
    type StudentAttendanceCollectionReportingType struct {
        EntityLevel *string `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *string `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *string `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *string `xml:"EntityName,omitempty" json:"EntityName"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      StatsCohortYearLevelList *StatsCohortYearLevelListType `xml:"StatsCohortYearLevelList,omitempty" json:"StatsCohortYearLevelList"`
      
      }
    
    type StatsCohortYearLevelListType struct {
        StatsCohortYearLevel []StatsCohortYearLevelType `xml:"StatsCohortYearLevel" json:"StatsCohortYearLevel"`
      
      }
    
    type StatsCohortYearLevelType struct {
        CohortYearLevel *YearLevelType `xml:"CohortYearLevel" json:"CohortYearLevel"`
      StatsCohortList *StatsCohortListType `xml:"StatsCohortList" json:"StatsCohortList"`
      
      }
    
    type StatsCohortListType struct {
        StatsCohort []StatsCohortType `xml:"StatsCohort" json:"StatsCohort"`
      
      }
    
    type StatsCohortType struct {
        StatsCohortId *LocalIdType `xml:"StatsCohortId" json:"StatsCohortId"`
      StatsIndigenousStudentType *string `xml:"StatsIndigenousStudentType" json:"StatsIndigenousStudentType"`
      CohortGender *string `xml:"CohortGender" json:"CohortGender"`
      DaysInReferencePeriod *int `xml:"DaysInReferencePeriod" json:"DaysInReferencePeriod"`
      PossibleSchoolDays *int `xml:"PossibleSchoolDays" json:"PossibleSchoolDays"`
      AttendanceDays *float64 `xml:"AttendanceDays" json:"AttendanceDays"`
      AttendanceLess90Percent *int `xml:"AttendanceLess90Percent" json:"AttendanceLess90Percent"`
      AttendanceGTE90Percent *int `xml:"AttendanceGTE90Percent" json:"AttendanceGTE90Percent"`
      PossibleSchoolDaysGT90PercentAttendance *int `xml:"PossibleSchoolDaysGT90PercentAttendance" json:"PossibleSchoolDaysGT90PercentAttendance"`
      
      }
    
    type AddressCollectionReportingListType struct {
        AddressCollectionReporting []AddressCollectionReportingType `xml:"AddressCollectionReporting" json:"AddressCollectionReporting"`
      
      }
    
    type AddressCollectionReportingType struct {
        EntityLevel *string `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *string `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *string `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *string `xml:"EntityName,omitempty" json:"EntityName"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      AGContextualQuestionList *AGContextualQuestionListType `xml:"AGContextualQuestionList,omitempty" json:"AGContextualQuestionList"`
      AddressCollectionStudentList *AddressCollectionStudentListType `xml:"AddressCollectionStudentList,omitempty" json:"AddressCollectionStudentList"`
      
      }
    
    type AddressCollectionStudentListType struct {
        AddressCollectionStudent []AddressCollectionStudentType `xml:"AddressCollectionStudent" json:"AddressCollectionStudent"`
      
      }
    
    type AddressCollectionStudentType struct {
        LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      EducationLevel *AUCodeSetsEducationLevelType `xml:"EducationLevel" json:"EducationLevel"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus"`
      ReportingParent2 *string `xml:"ReportingParent2" json:"ReportingParent2"`
      StudentAddress *AddressType `xml:"StudentAddress" json:"StudentAddress"`
      Parent1 *AGParentType `xml:"Parent1" json:"Parent1"`
      Parent2 *AGParentType `xml:"Parent2,omitempty" json:"Parent2"`
      
      }
    
    type AGParentType struct {
        ParentName *NameOfRecordType `xml:"ParentName" json:"ParentName"`
      AddressSameAsStudent *AUCodeSetsYesOrNoCategoryType `xml:"AddressSameAsStudent" json:"AddressSameAsStudent"`
      ParentAddress *AddressType `xml:"ParentAddress" json:"ParentAddress"`
      
      }
    
    type AGRoundListType struct {
        AGRound []AGRoundType `xml:"AGRound" json:"AGRound"`
      
      }
    
    type AGRoundType struct {
        RoundCode *string `xml:"RoundCode" json:"RoundCode"`
      RoundName *string `xml:"RoundName" json:"RoundName"`
      StartDate *string `xml:"StartDate" json:"StartDate"`
      DueDate *string `xml:"DueDate" json:"DueDate"`
      EndDate *string `xml:"EndDate" json:"EndDate"`
      
      }
    
    type AGContextualQuestionListType struct {
        AGContextualQuestion []AGContextualQuestionType `xml:"AGContextualQuestion" json:"AGContextualQuestion"`
      
      }
    
    type AGContextualQuestionType struct {
        AGContextCode *AUCodeSetsAGContextQuestionType `xml:"AGContextCode" json:"AGContextCode"`
      AGAnswer *string `xml:"AGAnswer" json:"AGAnswer"`
      
      }
    
    type CensusReportingListType struct {
        CensusReporting []CensusReportingType `xml:"CensusReporting" json:"CensusReporting"`
      
      }
    
    type CensusReportingType struct {
        EntityLevel *string `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      CommonwealthId *string `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *string `xml:"EntityName,omitempty" json:"EntityName"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      CensusStaffList *CensusStaffListType `xml:"CensusStaffList,omitempty" json:"CensusStaffList"`
      CensusStudentList *CensusStudentListType `xml:"CensusStudentList,omitempty" json:"CensusStudentList"`
      
      }
    
    type CensusStaffListType struct {
        CensusStaff []CensusStaffType `xml:"CensusStaff" json:"CensusStaff"`
      
      }
    
    type CensusStaffType struct {
        StaffCohortId *LocalIdType `xml:"StaffCohortId" json:"StaffCohortId"`
      StaffActivity *StaffActivityExtensionType `xml:"StaffActivity" json:"StaffActivity"`
      CohortGender *string `xml:"CohortGender" json:"CohortGender"`
      CohortIndigenousType *string `xml:"CohortIndigenousType" json:"CohortIndigenousType"`
      PrimaryFTE *float64 `xml:"PrimaryFTE" json:"PrimaryFTE"`
      SecondaryFTE *float64 `xml:"SecondaryFTE" json:"SecondaryFTE"`
      JobFTE *float64 `xml:"JobFTE" json:"JobFTE"`
      Headcount *int `xml:"Headcount" json:"Headcount"`
      
      }
    
    type StaffAssignmentMostRecentContainerType struct {
        PrimaryFTE *float64 `xml:"PrimaryFTE" json:"PrimaryFTE"`
      SecondaryFTE *float64 `xml:"SecondaryFTE" json:"SecondaryFTE"`
      
      }
    
    type CensusStudentListType struct {
        CensusStudent []CensusStudentType `xml:"CensusStudent" json:"CensusStudent"`
      
      }
    
    type CensusStudentType struct {
        StudentCohortId *LocalIdType `xml:"StudentCohortId" json:"StudentCohortId"`
      YearLevel *YearLevelType `xml:"YearLevel" json:"YearLevel"`
      CensusAge *int `xml:"CensusAge" json:"CensusAge"`
      CohortGender *string `xml:"CohortGender" json:"CohortGender"`
      CohortIndigenousType *string `xml:"CohortIndigenousType" json:"CohortIndigenousType"`
      EducationMode *string `xml:"EducationMode" json:"EducationMode"`
      StudentOnVisa *string `xml:"StudentOnVisa" json:"StudentOnVisa"`
      OverseasStudent *string `xml:"OverseasStudent" json:"OverseasStudent"`
      DisabilityLevelOfAdjustment *string `xml:"DisabilityLevelOfAdjustment" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *string `xml:"DisabilityCategory" json:"DisabilityCategory"`
      FTE *float64 `xml:"FTE" json:"FTE"`
      Headcount *int `xml:"Headcount" json:"Headcount"`
      
      }
    
    type AGReportingObjectResponseListType struct {
        AGReportingObjectResponse []AGReportingObjectResponseType `xml:"AGReportingObjectResponse" json:"AGReportingObjectResponse"`
      
      }
    
    type AGReportingObjectResponseType struct {
        SubmittedRefId *string `xml:"SubmittedRefId,omitempty" json:"SubmittedRefId"`
      SIFRefId *string `xml:"SIFRefId,omitempty" json:"SIFRefId"`
      HTTPStatusCode *string `xml:"HTTPStatusCode,omitempty" json:"HTTPStatusCode"`
      ErrorText *string `xml:"ErrorText,omitempty" json:"ErrorText"`
      CommonwealthId *string `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *string `xml:"EntityName,omitempty" json:"EntityName"`
      AGSubmissionStatusCode *AUCodeSetsAGSubmissionStatusType `xml:"AGSubmissionStatusCode" json:"AGSubmissionStatusCode"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList"`
      
      }
    
    type FQReportingListType struct {
        FQReporting []FQReportingType `xml:"FQReporting" json:"FQReporting"`
      
      }
    
    type FQReportingType struct {
        EntityLevel *string `xml:"EntityLevel" json:"EntityLevel"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *string `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *string `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *string `xml:"EntityName,omitempty" json:"EntityName"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      FQContextualQuestionList *FQContextualQuestionListType `xml:"FQContextualQuestionList,omitempty" json:"FQContextualQuestionList"`
      FQItemList *FQItemListType `xml:"FQItemList" json:"FQItemList"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList"`
      
      }
    
    type FQContextualQuestionListType struct {
        FQContextualQuestion []FQContextualQuestionType `xml:"FQContextualQuestion" json:"FQContextualQuestion"`
      
      }
    
    type FQContextualQuestionType struct {
        FQContext *string `xml:"FQContext" json:"FQContext"`
      FQAnswer *string `xml:"FQAnswer" json:"FQAnswer"`
      
      }
    
    type FQItemListType struct {
        FQItem []FQItemType `xml:"FQItem" json:"FQItem"`
      
      }
    
    type FQItemType struct {
        FQItemCode *string `xml:"FQItemCode" json:"FQItemCode"`
      TuitionAmount *float64 `xml:"TuitionAmount" json:"TuitionAmount"`
      BoardingAmount *float64 `xml:"BoardingAmount" json:"BoardingAmount"`
      SystemAmount *float64 `xml:"SystemAmount" json:"SystemAmount"`
      DioceseAmount *float64 `xml:"DioceseAmount" json:"DioceseAmount"`
      FQComments *string `xml:"FQComments,omitempty" json:"FQComments"`
      
      }
    
    type AGRuleListType struct {
        AGRule []AGRuleType `xml:"AGRule" json:"AGRule"`
      
      }
    
    type AGRuleType struct {
        AGRuleCode *string `xml:"AGRuleCode,omitempty" json:"AGRuleCode"`
      AGRuleComment *string `xml:"AGRuleComment,omitempty" json:"AGRuleComment"`
      AGRuleResponse *string `xml:"AGRuleResponse,omitempty" json:"AGRuleResponse"`
      AGRuleStatus *string `xml:"AGRuleStatus,omitempty" json:"AGRuleStatus"`
      
      }
    
    type SoftwareVendorInfoContainerType struct {
        SoftwareProduct *string `xml:"SoftwareProduct" json:"SoftwareProduct"`
      SoftwareVersion *string `xml:"SoftwareVersion" json:"SoftwareVersion"`
      
      }
    
    type TimeTableScheduleType struct {
        SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Title *string `xml:"Title" json:"Title"`
      DaysPerCycle *int `xml:"DaysPerCycle" json:"DaysPerCycle"`
      PeriodsPerDay *int `xml:"PeriodsPerDay" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *int `xml:"TeachingPeriodsPerDay" json:"TeachingPeriodsPerDay"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolName *string `xml:"SchoolName,omitempty" json:"SchoolName"`
      TimeTableCreationDate *string `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate"`
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList" json:"TimeTableDayList"`
      
      }
    
    type TimeTableScheduleCellListType struct {
        TimeTableScheduleCell []TimeTableScheduleCellType `xml:"TimeTableScheduleCell" json:"TimeTableScheduleCell"`
      
      }
    
    type TimeTableScheduleCellType struct {
        TimeTableScheduleCellLocalId *LocalIdType `xml:"TimeTableScheduleCellLocalId" json:"TimeTableScheduleCellLocalId"`
      TimeTableSubjectRefId *string `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeachingGroupGUID *string `xml:"TeachingGroupGUID" json:"TeachingGroupGUID"`
      RoomInfoRefId *string `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId"`
      StaffPersonalRefId *string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      TimeTableLocalId *LocalIdType `xml:"TimeTableLocalId,omitempty" json:"TimeTableLocalId"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TeachingGroupLocalId *LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      CellType *string `xml:"CellType" json:"CellType"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      
      }
    
    type TeachingGroupScheduleListType struct {
        TeachingGroupSchedule []TeachingGroupScheduleType `xml:"TeachingGroupSchedule" json:"TeachingGroupSchedule"`
      
      }
    
    type TeachingGroupScheduleType struct {
        EditorGUID *RefIdType `xml:"EditorGUID" json:"EditorGUID"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      ShortName *string `xml:"ShortName" json:"ShortName"`
      LongName *string `xml:"LongName,omitempty" json:"LongName"`
      GroupType *string `xml:"GroupType,omitempty" json:"GroupType"`
      Sett *string `xml:"Set,omitempty" json:"Set"`
      Block *string `xml:"Block,omitempty" json:"Block"`
      CurriculumLevel *string `xml:"CurriculumLevel,omitempty" json:"CurriculumLevel"`
      SchoolInfoRefId *RefIdType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      SchoolCourseLocalId *LocalIdType `xml:"SchoolCourseLocalId,omitempty" json:"SchoolCourseLocalId"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TimeTableSubjectLocalId *LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId"`
      Semester *int `xml:"Semester" json:"Semester"`
      StudentList *StudentListType `xml:"StudentList,omitempty" json:"StudentList"`
      TeacherList *TeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      MinClassSize *int `xml:"MinClassSize" json:"MinClassSize"`
      MaxClassSize *int `xml:"MaxClassSize" json:"MaxClassSize"`
      TeachingGroupPeriodList *TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList,omitempty" json:"TeachingGroupPeriodList"`
      
      }
    
    type LocalCodeListType struct {
        LocalCode []LocalCodeType `xml:"LocalCode" json:"LocalCode"`
      
      }
    
    type LocalCodeType struct {
        LocalisedCode *string `xml:"LocalisedCode" json:"LocalisedCode"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      Element *string `xml:"Element,omitempty" json:"Element"`
      ListIndex *int `xml:"ListIndex" json:"ListIndex"`
      
      }
    
    type StudentGroupListType struct {
        StudentGroup []StudentGroupType `xml:"StudentGroup" json:"StudentGroup"`
      
      }
    
    type StudentGroupType struct {
        GroupCategory *AUCodeSetsGroupCategoryCodeType `xml:"GroupCategory" json:"GroupCategory"`
      GroupLocalId *LocalIdType `xml:"GroupLocalId" json:"GroupLocalId"`
      GroupDescription *string `xml:"GroupDescription,omitempty" json:"GroupDescription"`
      
      }
    
    type PublishingPermissionListType struct {
        PublishingPermission []PublishingPermissionType `xml:"PublishingPermission" json:"PublishingPermission"`
      
      }
    
    type PublishingPermissionType struct {
        PermissionCategory *AUCodeSetsPermissionCategoryCodeType `xml:"PermissionCategory" json:"PermissionCategory"`
      PermissionValue *AUCodeSetsYesOrNoCategoryType `xml:"PermissionValue" json:"PermissionValue"`
      
      }
    
    type EntityContactInfoType struct {
        Name *NameType `xml:"Name" json:"Name"`
      PositionTitle *string `xml:"PositionTitle,omitempty" json:"PositionTitle"`
      Role *string `xml:"Role,omitempty" json:"Role"`
      RegistrationDetails *string `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails"`
      Qualifications *string `xml:"Qualifications,omitempty" json:"Qualifications"`
      Address *AddressType `xml:"Address,omitempty" json:"Address"`
      Email *EmailType `xml:"Email" json:"Email"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type CopyRightContainerType struct {
        Date *string `xml:"Date" json:"Date"`
      Holder *string `xml:"Holder" json:"Holder"`
      
      }
    
    type StandardsSettingBodyType struct {
        Country *CountryType `xml:"Country,omitempty" json:"Country"`
      StateProvince *StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince"`
      SettingBodyName *string `xml:"SettingBodyName,omitempty" json:"SettingBodyName"`
      
      }
    
    type StandardHierarchyLevelType struct {
        Number *int `xml:"Number" json:"Number"`
      Description *string `xml:"Description" json:"Description"`
      
      }
    
    type StandardIdentifierType struct {
        YearCreated *string `xml:"YearCreated" json:"YearCreated"`
      ACStrandSubjectArea *ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      StandardNumber *string `xml:"StandardNumber" json:"StandardNumber"`
      YearLevels *YearLevelsType `xml:"YearLevels" json:"YearLevels"`
      Benchmark *string `xml:"Benchmark,omitempty" json:"Benchmark"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      IndicatorNumber *string `xml:"IndicatorNumber,omitempty" json:"IndicatorNumber"`
      AlternateIdentificationCodes *AlternateIdentificationCodeListType `xml:"AlternateIdentificationCodes,omitempty" json:"AlternateIdentificationCodes"`
      Organization *string `xml:"Organization" json:"Organization"`
      
      }
    
    type AlternateIdentificationCodeListType struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode" json:"AlternateIdentificationCode"`
      
      }
    
    type RelatedLearningStandardItemRefIdListType struct {
        LearningStandardItemRefId []RelatedLearningStandardItemRefIdType `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      
      }
    
    type RelatedLearningStandardItemRefIdType struct {
          RelationshipType *string `xml:"RelationshipType,attr" json:"RelationshipType"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type ValidLetterMarkListType struct {
        ValidLetterMark []ValidLetterMarkType `xml:"ValidLetterMark" json:"ValidLetterMark"`
      
      }
    
    type ValidLetterMarkType struct {
        Code *string `xml:"Code" json:"Code"`
      NumericEquivalent *float64 `xml:"NumericEquivalent" json:"NumericEquivalent"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      
      }
    
    type StudentGradeMarkersListType struct {
        Marker []MarkerType `xml:"Marker" json:"Marker"`
      
      }
    
    type MarkerType struct {
        StaffPersonalRefId *string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Role *string `xml:"Role,omitempty" json:"Role"`
      
      }
    
    type GradingScoreListType struct {
        GradingAssignmentScore []AssignmentScoreType `xml:"GradingAssignmentScore" json:"GradingAssignmentScore"`
      
      }
    
    type AssignmentScoreType struct {
        GradingAssignmentScoreRefId *string `xml:"GradingAssignmentScoreRefId,omitempty" json:"GradingAssignmentScoreRefId"`
      Weight *float64 `xml:"Weight" json:"Weight"`
      
      }
    
    type GradeType struct {
        Percentage *float64 `xml:"Percentage" json:"Percentage"`
      Numeric *float64 `xml:"Numeric" json:"Numeric"`
      Letter *string `xml:"Letter" json:"Letter"`
      Narrative *string `xml:"Narrative,omitempty" json:"Narrative"`
      MarkInfoRefId *string `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      
      }
    
    type LearningStandardListType struct {
        LearningStandard []LearningStandardType `xml:"LearningStandard" json:"LearningStandard"`
      
      }
    
    type LearningStandardType struct {
        LearningStandardItemRefId *string `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      LearningStandardURL *string `xml:"LearningStandardURL,omitempty" json:"LearningStandardURL"`
      LearningStandardLocalId *LocalIdType `xml:"LearningStandardLocalId,omitempty" json:"LearningStandardLocalId"`
      
      }
    
    type AssignmentListType struct {
        GradingAssignmentRefId []string `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      
      }
    
    type GenericRubricType struct {
        RubricType *string `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type SymptomListType struct {
        Symptom []string `xml:"Symptom" json:"Symptom"`
      
      }
    
    type MedicationListType struct {
        Medication []MedicationType `xml:"Medication" json:"Medication"`
      
      }
    
    type MedicationType struct {
        MedicationName *string `xml:"MedicationName,omitempty" json:"MedicationName"`
      Dosage *string `xml:"Dosage,omitempty" json:"Dosage"`
      Frequency *string `xml:"Frequency,omitempty" json:"Frequency"`
      AdministrationInformation *string `xml:"AdministrationInformation,omitempty" json:"AdministrationInformation"`
      Method *string `xml:"Method,omitempty" json:"Method"`
      
      }
    
    type WellbeingEventCategoryListType struct {
        WellbeingEventCategory []WellbeingEventCategoryType `xml:"WellbeingEventCategory" json:"WellbeingEventCategory"`
      
      }
    
    type WellbeingEventCategoryType struct {
        EventCategory *string `xml:"EventCategory" json:"EventCategory"`
      WellbeingEventSubCategoryList *WellbeingEventSubCategoryListType `xml:"WellbeingEventSubCategoryList" json:"WellbeingEventSubCategoryList"`
      
      }
    
    type WellbeingEventSubCategoryListType struct {
        WellbeingEventSubCategory []string `xml:"WellbeingEventSubCategory" json:"WellbeingEventSubCategory"`
      
      }
    
    type WellbeingEventLocationDetailsType struct {
        EventLocation *AUCodeSetsWellbeingEventLocationType `xml:"EventLocation" json:"EventLocation"`
      Class *string `xml:"Class,omitempty" json:"Class"`
      FurtherLocationNotes *string `xml:"FurtherLocationNotes,omitempty" json:"FurtherLocationNotes"`
      
      }
    
    type FollowUpActionListType struct {
        FollowUpAction []FollowUpActionType `xml:"FollowUpAction" json:"FollowUpAction"`
      
      }
    
    type FollowUpActionType struct {
        WellbeingResponseRefId *string `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      FollowUpDetails *string `xml:"FollowUpDetails,omitempty" json:"FollowUpDetails"`
      FollowUpActionCategory *string `xml:"FollowUpActionCategory,omitempty" json:"FollowUpActionCategory"`
      
      }
    
    type PersonInvolvementListType struct {
        PersonInvolvement []PersonInvolvementType `xml:"PersonInvolvement" json:"PersonInvolvement"`
      
      }
    
    type PersonInvolvementType struct {
      PersonRefId *PersonInvolvementType_PersonRefId
      ShortName *string `xml:"ShortName,omitempty" json:"ShortName"`
      HowInvolved *string `xml:"HowInvolved,omitempty" json:"HowInvolved"`
      
      }
    
    type WithdrawalTimeListType struct {
        Withdrawal []WithdrawalType `xml:"Withdrawal" json:"Withdrawal"`
      
      }
    
    type WithdrawalType struct {
        WithdrawalDate *string `xml:"WithdrawalDate" json:"WithdrawalDate"`
      WithdrawalStartTime *string `xml:"WithdrawalStartTime,omitempty" json:"WithdrawalStartTime"`
      WithdrawalEndTime *string `xml:"WithdrawalEndTime,omitempty" json:"WithdrawalEndTime"`
      TimeTableSubjectRefId *string `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      ScheduledActivityRefId *string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimeTableCellRefId *string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      
      }
    
    type SuspensionContainerType struct {
        SuspensionCategory *AUCodeSetsSuspensionCategoryType `xml:"SuspensionCategory" json:"SuspensionCategory"`
      SuspensionNotes *string `xml:"SuspensionNotes,omitempty" json:"SuspensionNotes"`
      WithdrawalTimeList *WithdrawalTimeListType `xml:"WithdrawalTimeList,omitempty" json:"WithdrawalTimeList"`
      Duration *float64 `xml:"Duration" json:"Duration"`
      AdvisementDate *string `xml:"AdvisementDate,omitempty" json:"AdvisementDate"`
      ResolutionMeetingTime *string `xml:"ResolutionMeetingTime,omitempty" json:"ResolutionMeetingTime"`
      ResolutionNotes *string `xml:"ResolutionNotes,omitempty" json:"ResolutionNotes"`
      EarlyReturnDate *string `xml:"EarlyReturnDate,omitempty" json:"EarlyReturnDate"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type DetentionContainerType struct {
        DetentionCategory *AUCodeSetsDetentionCategoryType `xml:"DetentionCategory" json:"DetentionCategory"`
      DetentionDate *string `xml:"DetentionDate,omitempty" json:"DetentionDate"`
      DetentionLocation *string `xml:"DetentionLocation,omitempty" json:"DetentionLocation"`
      DetentionNotes *string `xml:"DetentionNotes,omitempty" json:"DetentionNotes"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type PlanRequiredContainerType struct {
        PlanRequiredList *PlanRequiredListType `xml:"PlanRequiredList,omitempty" json:"PlanRequiredList"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type PlanRequiredListType struct {
        Plan []WellbeingPlanType `xml:"Plan" json:"Plan"`
      
      }
    
    type WellbeingPlanType struct {
        PersonalisedPlanRefId *string `xml:"PersonalisedPlanRefId,omitempty" json:"PersonalisedPlanRefId"`
      PlanNotes *string `xml:"PlanNotes,omitempty" json:"PlanNotes"`
      
      }
    
    type AwardContainerType struct {
        AwardDate *string `xml:"AwardDate,omitempty" json:"AwardDate"`
      AwardType *string `xml:"AwardType,omitempty" json:"AwardType"`
      AwardDescription *string `xml:"AwardDescription,omitempty" json:"AwardDescription"`
      AwardNotes *string `xml:"AwardNotes,omitempty" json:"AwardNotes"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type OtherWellbeingResponseContainerType struct {
        OtherResponseDate *string `xml:"OtherResponseDate,omitempty" json:"OtherResponseDate"`
      OtherResponseType *string `xml:"OtherResponseType,omitempty" json:"OtherResponseType"`
      OtherResponseDescription *string `xml:"OtherResponseDescription,omitempty" json:"OtherResponseDescription"`
      OtherResponseNotes *string `xml:"OtherResponseNotes,omitempty" json:"OtherResponseNotes"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type WellbeingDocumentListType struct {
        Document []WellbeingDocumentType `xml:"Document" json:"Document"`
      
      }
    
    type WellbeingDocumentType struct {
        Location *string `xml:"Location" json:"Location"`
      Sensitivity *string `xml:"Sensitivity,omitempty" json:"Sensitivity"`
      URL *string `xml:"URL,omitempty" json:"URL"`
      DocumentType *string `xml:"DocumentType,omitempty" json:"DocumentType"`
      DocumentReviewDate *string `xml:"DocumentReviewDate,omitempty" json:"DocumentReviewDate"`
      DocumentDescription *string `xml:"DocumentDescription,omitempty" json:"DocumentDescription"`
      
      }
    
    type NAPTestItemContentType struct {
        NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      ItemName *string `xml:"ItemName" json:"ItemName"`
      ItemType *AUCodeSetsNAPTestItemTypeType `xml:"ItemType" json:"ItemType"`
      Subdomain *string `xml:"Subdomain" json:"Subdomain"`
      WritingGenre *AUCodeSetsNAPWritingGenreType `xml:"WritingGenre,omitempty" json:"WritingGenre"`
      ItemDescriptor *string `xml:"ItemDescriptor" json:"ItemDescriptor"`
      ReleasedStatus *bool `xml:"ReleasedStatus" json:"ReleasedStatus"`
      MarkingType *AUCodeSetsNAPTestItemMarkingTypeType `xml:"MarkingType" json:"MarkingType"`
      MultipleChoiceOptionCount *int `xml:"MultipleChoiceOptionCount" json:"MultipleChoiceOptionCount"`
      CorrectAnswer *string `xml:"CorrectAnswer,omitempty" json:"CorrectAnswer"`
      MaximumScore *float64 `xml:"MaximumScore" json:"MaximumScore"`
      ItemDifficulty *float64 `xml:"ItemDifficulty" json:"ItemDifficulty"`
      ItemDifficultyLogit5 *float64 `xml:"ItemDifficultyLogit5" json:"ItemDifficultyLogit5"`
      ItemDifficultyLogit62 *float64 `xml:"ItemDifficultyLogit62" json:"ItemDifficultyLogit62"`
      ItemDifficultyLogit5SE *float64 `xml:"ItemDifficultyLogit5SE" json:"ItemDifficultyLogit5SE"`
      ItemDifficultyLogit62SE *float64 `xml:"ItemDifficultyLogit62SE" json:"ItemDifficultyLogit62SE"`
      ItemProficiencyBand *int `xml:"ItemProficiencyBand" json:"ItemProficiencyBand"`
      ItemProficiencyLevel *string `xml:"ItemProficiencyLevel,omitempty" json:"ItemProficiencyLevel"`
      ExemplarURL *string `xml:"ExemplarURL,omitempty" json:"ExemplarURL"`
      ItemSubstitutedForList *SubstituteItemListType `xml:"ItemSubstitutedForList,omitempty" json:"ItemSubstitutedForList"`
      ContentDescriptionList *ContentDescriptionListType `xml:"ContentDescriptionList,omitempty" json:"ContentDescriptionList"`
      StimulusList *StimulusListType `xml:"StimulusList,omitempty" json:"StimulusList"`
      NAPWritingRubricList *NAPWritingRubricListType `xml:"NAPWritingRubricList,omitempty" json:"NAPWritingRubricList"`
      
      }
    
    type NAPTestletContentType struct {
        NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletName *string `xml:"TestletName,omitempty" json:"TestletName"`
      Node *string `xml:"Node,omitempty" json:"Node"`
      LocationInStage *int `xml:"LocationInStage" json:"LocationInStage"`
      TestletMaximumScore *float64 `xml:"TestletMaximumScore" json:"TestletMaximumScore"`
      
      }
    
    type NAPTestContentType struct {
        NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      TestName *string `xml:"TestName" json:"TestName"`
      TestLevel *YearLevelType `xml:"TestLevel" json:"TestLevel"`
      TestType *AUCodeSetsNAPTestTypeType `xml:"TestType" json:"TestType"`
      Domain *AUCodeSetsNAPTestDomainType `xml:"Domain" json:"Domain"`
      TestYear *SchoolYearType `xml:"TestYear" json:"TestYear"`
      StagesCount *int `xml:"StagesCount" json:"StagesCount"`
      DomainBands *DomainBandsContainerType `xml:"DomainBands" json:"DomainBands"`
      DomainProficiency *DomainProficiencyContainerType `xml:"DomainProficiency" json:"DomainProficiency"`
      
      }
    
    type PlausibleScaledValueListType struct {
        PlausibleScaledValue []float64 `xml:"PlausibleScaledValue" json:"PlausibleScaledValue"`
      
      }
    
    type SubstituteItemListType struct {
        SubstituteItem []SubstituteItemType `xml:"SubstituteItem" json:"SubstituteItem"`
      
      }
    
    type SubstituteItemType struct {
        SubstituteItemRefId *string `xml:"SubstituteItemRefId" json:"SubstituteItemRefId"`
      SubstituteItemLocalId *LocalIdType `xml:"SubstituteItemLocalId,omitempty" json:"SubstituteItemLocalId"`
      PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      
      }
    
    type CodeFrameTestItemListType struct {
        TestItem []CodeFrameTestItemType `xml:"TestItem" json:"TestItem"`
      
      }
    
    type CodeFrameTestItemType struct {
        TestItemRefId *string `xml:"TestItemRefId" json:"TestItemRefId"`
      SequenceNumber *int `xml:"SequenceNumber" json:"SequenceNumber"`
      TestItemContent *NAPTestItemContentType `xml:"TestItemContent" json:"TestItemContent"`
      
      }
    
    type StimulusLocalIdListType struct {
        StimulusLocalId []LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      
      }
    
    type DomainBandsContainerType struct {
        Band1Lower *float64 `xml:"Band1Lower" json:"Band1Lower"`
      Band1Upper *float64 `xml:"Band1Upper" json:"Band1Upper"`
      Band2Lower *float64 `xml:"Band2Lower" json:"Band2Lower"`
      Band2Upper *float64 `xml:"Band2Upper" json:"Band2Upper"`
      Band3Lower *float64 `xml:"Band3Lower" json:"Band3Lower"`
      Band3Upper *float64 `xml:"Band3Upper" json:"Band3Upper"`
      Band4Lower *float64 `xml:"Band4Lower" json:"Band4Lower"`
      Band4Upper *float64 `xml:"Band4Upper" json:"Band4Upper"`
      Band5Lower *float64 `xml:"Band5Lower" json:"Band5Lower"`
      Band5Upper *float64 `xml:"Band5Upper" json:"Band5Upper"`
      Band6Lower *float64 `xml:"Band6Lower" json:"Band6Lower"`
      Band6Upper *float64 `xml:"Band6Upper" json:"Band6Upper"`
      Band7Lower *float64 `xml:"Band7Lower" json:"Band7Lower"`
      Band7Upper *float64 `xml:"Band7Upper" json:"Band7Upper"`
      Band8Lower *float64 `xml:"Band8Lower" json:"Band8Lower"`
      Band8Upper *float64 `xml:"Band8Upper" json:"Band8Upper"`
      Band9Lower *float64 `xml:"Band9Lower" json:"Band9Lower"`
      Band9Upper *float64 `xml:"Band9Upper" json:"Band9Upper"`
      Band10Lower *float64 `xml:"Band10Lower" json:"Band10Lower"`
      Band10Upper *float64 `xml:"Band10Upper" json:"Band10Upper"`
      
      }
    
    type DomainProficiencyContainerType struct {
        Level1Lower *float64 `xml:"Level1Lower" json:"Level1Lower"`
      Level1Upper *float64 `xml:"Level1Upper" json:"Level1Upper"`
      Level2Lower *float64 `xml:"Level2Lower" json:"Level2Lower"`
      Level2Upper *float64 `xml:"Level2Upper" json:"Level2Upper"`
      Level3Lower *float64 `xml:"Level3Lower" json:"Level3Lower"`
      Level3Upper *float64 `xml:"Level3Upper" json:"Level3Upper"`
      Level4Lower *float64 `xml:"Level4Lower" json:"Level4Lower"`
      Level4Upper *float64 `xml:"Level4Upper" json:"Level4Upper"`
      
      }
    
    type NAPTestItemListType struct {
        TestItem []NAPTestItem2Type `xml:"TestItem" json:"TestItem"`
      
      }
    
    type NAPTestItem2Type struct {
        TestItemRefId *string `xml:"TestItemRefId" json:"TestItemRefId"`
      TestItemLocalId *LocalIdType `xml:"TestItemLocalId" json:"TestItemLocalId"`
      SequenceNumber *int `xml:"SequenceNumber" json:"SequenceNumber"`
      
      }
    
    type NAPCodeFrameTestletListType struct {
        Testlet []NAPTestletCodeFrameType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletCodeFrameType struct {
        NAPTestletRefId *string `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      TestletContent *NAPTestletContentType `xml:"TestletContent" json:"TestletContent"`
      TestItemList *CodeFrameTestItemListType `xml:"TestItemList" json:"TestItemList"`
      
      }
    
    type NAPStudentResponseTestletListType struct {
        Testlet []NAPTestletResponseType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletResponseType struct {
        NAPTestletRefId *string `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletSubScore *float64 `xml:"TestletSubScore" json:"TestletSubScore"`
      ItemResponseList *NAPTestletItemResponseListType `xml:"ItemResponseList" json:"ItemResponseList"`
      
      }
    
    type NAPTestletItemResponseListType struct {
        ItemResponse []NAPTestletResponseItemType `xml:"ItemResponse" json:"ItemResponse"`
      
      }
    
    type NAPTestletResponseItemType struct {
        NAPTestItemRefId *string `xml:"NAPTestItemRefId,omitempty" json:"NAPTestItemRefId"`
      NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      Response *string `xml:"Response,omitempty" json:"Response"`
      ResponseCorrectness *AUCodeSetsNAPResponseCorrectnessType `xml:"ResponseCorrectness" json:"ResponseCorrectness"`
      Score *float64 `xml:"Score" json:"Score"`
      LapsedTimeItem *string `xml:"LapsedTimeItem,omitempty" json:"LapsedTimeItem"`
      SequenceNumber *int `xml:"SequenceNumber" json:"SequenceNumber"`
      ItemWeight *float64 `xml:"ItemWeight" json:"ItemWeight"`
      SubscoreList *NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      
      }
    
    type NAPSubscoreListType struct {
        Subscore []NAPSubscoreType `xml:"Subscore" json:"Subscore"`
      
      }
    
    type NAPSubscoreType struct {
        SubscoreType *string `xml:"SubscoreType" json:"SubscoreType"`
      SubscoreValue *float64 `xml:"SubscoreValue" json:"SubscoreValue"`
      
      }
    
    type DomainScoreType struct {
        RawScore *float64 `xml:"RawScore" json:"RawScore"`
      ScaledScoreValue *float64 `xml:"ScaledScoreValue" json:"ScaledScoreValue"`
      ScaledScoreLogitValue *float64 `xml:"ScaledScoreLogitValue" json:"ScaledScoreLogitValue"`
      ScaledScoreStandardError *float64 `xml:"ScaledScoreStandardError" json:"ScaledScoreStandardError"`
      ScaledScoreLogitStandardError *float64 `xml:"ScaledScoreLogitStandardError" json:"ScaledScoreLogitStandardError"`
      StudentDomainBand *int `xml:"StudentDomainBand" json:"StudentDomainBand"`
      StudentProficiency *string `xml:"StudentProficiency,omitempty" json:"StudentProficiency"`
      PlausibleScaledValueList *PlausibleScaledValueListType `xml:"PlausibleScaledValueList" json:"PlausibleScaledValueList"`
      
      }
    
    type NAPWritingRubricListType struct {
        NAPWritingRubric []NAPWritingRubricType `xml:"NAPWritingRubric" json:"NAPWritingRubric"`
      
      }
    
    type NAPWritingRubricType struct {
        RubricType *string `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type ScoreListType struct {
        Score []ScoreType `xml:"Score" json:"Score"`
      
      }
    
    type ScoreType struct {
        MaxScoreValue *float64 `xml:"MaxScoreValue" json:"MaxScoreValue"`
      ScoreDescriptionList *ScoreDescriptionListType `xml:"ScoreDescriptionList" json:"ScoreDescriptionList"`
      
      }
    
    type ScoreDescriptionListType struct {
        ScoreDescription []ScoreDescriptionType `xml:"ScoreDescription" json:"ScoreDescription"`
      
      }
    
    type ScoreDescriptionType struct {
        ScoreValue *float64 `xml:"ScoreValue" json:"ScoreValue"`
      Descriptor *string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type StimulusListType struct {
        Stimulus []StimulusType `xml:"Stimulus" json:"Stimulus"`
      
      }
    
    type StimulusType struct {
        StimulusLocalId *LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      TextGenre *string `xml:"TextGenre,omitempty" json:"TextGenre"`
      TextType *string `xml:"TextType,omitempty" json:"TextType"`
      WordCount *int `xml:"WordCount" json:"WordCount"`
      TextDescriptor *string `xml:"TextDescriptor,omitempty" json:"TextDescriptor"`
      Content *string `xml:"Content" json:"Content"`
      
      }
    
    type ContentDescriptionListType struct {
        ContentDescription []string `xml:"ContentDescription" json:"ContentDescription"`
      
      }
    
    type PNPCodeListType struct {
        PNPCode []AUCodeSetsPNPCodeType `xml:"PNPCode" json:"PNPCode"`
      
      }
    
    type AdjustmentContainerType struct {
        PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      BookletType *string `xml:"BookletType,omitempty" json:"BookletType"`
      
      }
    
    type TestDisruptionListType struct {
        TestDisruption []TestDisruptionType `xml:"TestDisruption" json:"TestDisruption"`
      
      }
    
    type TestDisruptionType struct {
        Event *string `xml:"Event" json:"Event"`
      
      }
    
    type CalendarSummaryListType struct {
        CalendarSummaryRefId []string `xml:"CalendarSummaryRefId" json:"CalendarSummaryRefId"`
      
      }
    
    type VisaSubClassType struct {
        Code *VisaSubClassCodeType `xml:"Code" json:"Code"`
      VisaExpiryDate *string `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      ATEExpiryDate *string `xml:"ATEExpiryDate,omitempty" json:"ATEExpiryDate"`
      ATEStartDate *string `xml:"ATEStartDate,omitempty" json:"ATEStartDate"`
      VisaStatisticalCode *string `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      
      }
    
    type VisaSubClassListType struct {
        VisaSubClass []VisaSubClassType `xml:"VisaSubClass" json:"VisaSubClass"`
      
      }
    
    type VisaSubClassCodeType string
    type LanguageBaseType struct {
        Code *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      LanguageType *AUCodeSetsLanguageTypeType `xml:"LanguageType,omitempty" json:"LanguageType"`
      Dialect *string `xml:"Dialect,omitempty" json:"Dialect"`
      
      }
    
    type ReligiousEventListType struct {
        ReligiousEvent []ReligiousEventType `xml:"ReligiousEvent" json:"ReligiousEvent"`
      
      }
    
    type ReligiousEventType struct {
        Type *string `xml:"Type" json:"Type"`
      Date *string `xml:"Date" json:"Date"`
      
      }
    
    type ReligionType struct {
        Code *AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type DwellingArrangementType struct {
        Code *AUCodeSetsDwellingArrangementType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type CountryListType struct {
        CountryOfCitizenship []CountryType `xml:"CountryOfCitizenship" json:"CountryOfCitizenship"`
      
      }
    
    type CountryList2Type struct {
        CountryOfResidency []CountryType `xml:"CountryOfResidency" json:"CountryOfResidency"`
      
      }
    
    type DebitOrCreditAmountType struct {
        MonetaryAmountType
          Type *string `xml:"Type,attr" json:"Type"`
      
      }
    
    type ScheduledActivityOverrideType struct {
          DateOfOverride *string `xml:"DateOfOverride,attr" json:"DateOfOverride"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type ActivityTimeType struct {
        CreationDate *string `xml:"CreationDate" json:"CreationDate"`
      Duration *ActivityTimeType_Duration
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      FinishDate *string `xml:"FinishDate,omitempty" json:"FinishDate"`
      DueDate *string `xml:"DueDate,omitempty" json:"DueDate"`
      
      }
    
    type SchoolCourseInfoOverrideType struct {
        Override *string `xml:"Override,attr" json:"Override"`
      CourseCode *string `xml:"CourseCode,omitempty" json:"CourseCode"`
      StateCourseCode *string `xml:"StateCourseCode,omitempty" json:"StateCourseCode"`
      DistrictCourseCode *string `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      CourseTitle *string `xml:"CourseTitle,omitempty" json:"CourseTitle"`
      InstructionalLevel *string `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel"`
      CourseCredits *string `xml:"CourseCredits,omitempty" json:"CourseCredits"`
      
      }
    
    type LocationOfInstructionType struct {
        Code *AUCodeSetsReceivingLocationOfInstructionType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LanguageOfInstructionType struct {
        Code *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type MediumOfInstructionType struct {
        Code *AUCodeSetsMediumOfInstructionType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentActivityType struct {
        Code *AUCodeSetsActivityInvolvementCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ContactFlagsType struct {
        ParentLegalGuardian *AUCodeSetsYesOrNoCategoryType `xml:"ParentLegalGuardian" json:"ParentLegalGuardian"`
      PickupRights *AUCodeSetsYesOrNoCategoryType `xml:"PickupRights" json:"PickupRights"`
      LivesWith *AUCodeSetsYesOrNoCategoryType `xml:"LivesWith" json:"LivesWith"`
      AccessToRecords *AUCodeSetsYesOrNoCategoryType `xml:"AccessToRecords" json:"AccessToRecords"`
      ReceivesAssessmentReport *AUCodeSetsYesOrNoCategoryType `xml:"ReceivesAssessmentReport" json:"ReceivesAssessmentReport"`
      EmergencyContact *AUCodeSetsYesOrNoCategoryType `xml:"EmergencyContact" json:"EmergencyContact"`
      HasCustody *AUCodeSetsYesOrNoCategoryType `xml:"HasCustody" json:"HasCustody"`
      DisciplinaryContact *AUCodeSetsYesOrNoCategoryType `xml:"DisciplinaryContact" json:"DisciplinaryContact"`
      AttendanceContact *AUCodeSetsYesOrNoCategoryType `xml:"AttendanceContact" json:"AttendanceContact"`
      PrimaryCareProvider *AUCodeSetsYesOrNoCategoryType `xml:"PrimaryCareProvider" json:"PrimaryCareProvider"`
      FeesBilling *AUCodeSetsYesOrNoCategoryType `xml:"FeesBilling" json:"FeesBilling"`
      FeesAccess *AUCodeSetsYesOrNoCategoryType `xml:"FeesAccess" json:"FeesAccess"`
      FamilyMail *AUCodeSetsYesOrNoCategoryType `xml:"FamilyMail" json:"FamilyMail"`
      InterventionOrder *AUCodeSetsYesOrNoCategoryType `xml:"InterventionOrder" json:"InterventionOrder"`
      
      }
    
    type AgencyType struct {
        Code *AUCodeSetsEducationAgencyTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type YearRangeType struct {
        Start *YearLevelType `xml:"Start" json:"Start"`
      End *YearLevelType `xml:"End" json:"End"`
      
      }
    
    type CreationUserType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      UserId *string `xml:"UserId" json:"UserId"`
      
      }
    
    type AuditInfoType struct {
        CreationUser *CreationUserType `xml:"CreationUser" json:"CreationUser"`
      CreationDateTime *string `xml:"CreationDateTime" json:"CreationDateTime"`
      
      }
    
    type AttendanceInfoType struct {
        CountsTowardAttendance *string `xml:"CountsTowardAttendance" json:"CountsTowardAttendance"`
      AttendanceValue *float64 `xml:"AttendanceValue" json:"AttendanceValue"`
      
      }
    
    type CalendarDateInfoType struct {
        Code *AUCodeSetsCalendarEventType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ProgramAvailabilityType struct {
        Code *AUCodeSets0211ProgramAvailabilityType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ReferralSourceType struct {
        Code *AUCodeSets0792IdentificationProcedureType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type PromotionInfoType struct {
        PromotionStatus *string `xml:"PromotionStatus,omitempty" json:"PromotionStatus"`
      
      }
    
    type CatchmentStatusContainerType struct {
        Code *AUCodeSetsPublicSchoolCatchmentStatusType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentExitStatusContainerType struct {
        Code *AUCodeSetsExitWithdrawalStatusType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentExitContainerType struct {
        Code *AUCodeSetsExitWithdrawalTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentEntryContainerType struct {
        Code *AUCodeSetsEntryTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentMostRecentContainerType struct {
        SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      HomeroomLocalId *LocalIdType `xml:"HomeroomLocalId,omitempty" json:"HomeroomLocalId"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      FTE *float64 `xml:"FTE" json:"FTE"`
      Parent1Language *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Parent1Language,omitempty" json:"Parent1Language"`
      Parent2Language *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Parent2Language,omitempty" json:"Parent2Language"`
      Parent1EmploymentType *AUCodeSetsEmploymentTypeType `xml:"Parent1EmploymentType,omitempty" json:"Parent1EmploymentType"`
      Parent2EmploymentType *AUCodeSetsEmploymentTypeType `xml:"Parent2EmploymentType,omitempty" json:"Parent2EmploymentType"`
      Parent1SchoolEducationLevel *AUCodeSetsSchoolEducationLevelTypeType `xml:"Parent1SchoolEducationLevel,omitempty" json:"Parent1SchoolEducationLevel"`
      Parent2SchoolEducationLevel *AUCodeSetsSchoolEducationLevelTypeType `xml:"Parent2SchoolEducationLevel,omitempty" json:"Parent2SchoolEducationLevel"`
      Parent1NonSchoolEducation *AUCodeSetsNonSchoolEducationType `xml:"Parent1NonSchoolEducation,omitempty" json:"Parent1NonSchoolEducation"`
      Parent2NonSchoolEducation *AUCodeSetsNonSchoolEducationType `xml:"Parent2NonSchoolEducation,omitempty" json:"Parent2NonSchoolEducation"`
      LocalCampusId *LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      TestLevel *YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      Homegroup *string `xml:"Homegroup,omitempty" json:"Homegroup"`
      ClassCode *string `xml:"ClassCode,omitempty" json:"ClassCode"`
      MembershipType *AUCodeSetsSchoolEnrollmentTypeType `xml:"MembershipType,omitempty" json:"MembershipType"`
      FFPOS *AUCodeSetsFFPOSStatusCodeType `xml:"FFPOS,omitempty" json:"FFPOS"`
      ReportingSchoolId *LocalIdType `xml:"ReportingSchoolId,omitempty" json:"ReportingSchoolId"`
      OtherEnrollmentSchoolACARAId *LocalIdType `xml:"OtherEnrollmentSchoolACARAId,omitempty" json:"OtherEnrollmentSchoolACARAId"`
      OtherSchoolName *string `xml:"OtherSchoolName,omitempty" json:"OtherSchoolName"`
      DisabilityLevelOfAdjustment *string `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *string `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory"`
      CensusAge *int `xml:"CensusAge" json:"CensusAge"`
      DistanceEducationStudent *AUCodeSetsYesOrNoCategoryType `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus"`
      
      }
    
    type StaffMostRecentContainerType struct {
        SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      LocalCampusId *LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId"`
      NAPLANClassList *NAPLANClassListType `xml:"NAPLANClassList,omitempty" json:"NAPLANClassList"`
      HomeGroup *string `xml:"HomeGroup,omitempty" json:"HomeGroup"`
      
      }
    
    type StaffActivityExtensionType struct {
        Code *AUCodeSetsStaffActivityType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type TotalEnrollmentsType struct {
        Girls *string `xml:"Girls,omitempty" json:"Girls"`
      Boys *string `xml:"Boys,omitempty" json:"Boys"`
      TotalStudents *string `xml:"TotalStudents,omitempty" json:"TotalStudents"`
      
      }
    
    type CampusContainerType struct {
        ParentSchoolId *string `xml:"ParentSchoolId" json:"ParentSchoolId"`
      SchoolCampusId *string `xml:"SchoolCampusId" json:"SchoolCampusId"`
      CampusType *AUCodeSetsSchoolLevelType `xml:"CampusType,omitempty" json:"CampusType"`
      AdminStatus *AUCodeSetsYesOrNoCategoryType `xml:"AdminStatus" json:"AdminStatus"`
      
      }
    
    type HouseholdContactInfoListType struct {
        HouseholdContactInfo []HouseholdContactInfoType `xml:"HouseholdContactInfo" json:"HouseholdContactInfo"`
      
      }
    
    type HouseholdContactInfoType struct {
        PreferenceNumber *int `xml:"PreferenceNumber" json:"PreferenceNumber"`
      HouseholdContactId *LocalIdType `xml:"HouseholdContactId,omitempty" json:"HouseholdContactId"`
      HouseholdSalutation *string `xml:"HouseholdSalutation,omitempty" json:"HouseholdSalutation"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      
      }
    
    type StatementCodesType struct {
        StatementCode []string `xml:"StatementCode" json:"StatementCode"`
      
      }
    
    type StatementsType struct {
        Statement []string `xml:"Statement" json:"Statement"`
      
      }
    
    type ProgramFundingSourcesType struct {
        ProgramFundingSource []ProgramFundingSourceType `xml:"ProgramFundingSource" json:"ProgramFundingSource"`
      
      }
    
    type ProgramFundingSourceType struct {
        Code *AUCodeSetsProgramFundingSourceCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type AttendanceTimesType struct {
        AttendanceTime []AttendanceTimeType `xml:"AttendanceTime" json:"AttendanceTime"`
      
      }
    
    type AttendanceTimeType struct {
        AttendanceType *string `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      StartTime *string `xml:"StartTime" json:"StartTime"`
      EndTime *string `xml:"EndTime" json:"EndTime"`
      DurationValue *float64 `xml:"DurationValue" json:"DurationValue"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      AttendanceNote *string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type PeriodAttendancesType struct {
        PeriodAttendance []PeriodAttendanceType `xml:"PeriodAttendance" json:"PeriodAttendance"`
      
      }
    
    type PeriodAttendanceType struct {
        AttendanceType *string `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      Date *string `xml:"Date" json:"Date"`
      SessionInfoRefId *string `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId *string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod *string `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      DayId *LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      StartTime *string `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime *string `xml:"EndTime,omitempty" json:"EndTime"`
      TimeIn *string `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut *string `xml:"TimeOut,omitempty" json:"TimeOut"`
      TimeTableCellRefId *string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      AttendanceNote *string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type StaffSubjectListType struct {
        StaffSubject []StaffSubjectType `xml:"StaffSubject" json:"StaffSubject"`
      
      }
    
    type StaffSubjectType struct {
        PreferenceNumber *int `xml:"PreferenceNumber" json:"PreferenceNumber"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      
      }
    
    type TeachingGroupListType struct {
        TeachingGroupRefId []string `xml:"TeachingGroupRefId" json:"TeachingGroupRefId"`
      
      }
    
    type ScheduledTeacherListType struct {
        TeacherCover []TeacherCoverType `xml:"TeacherCover" json:"TeacherCover"`
      
      }
    
    type TeacherCoverType struct {
        StaffPersonalRefId *string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      StartTime *string `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime *string `xml:"FinishTime,omitempty" json:"FinishTime"`
      Credit *AUCodeSetsTeacherCoverCreditType `xml:"Credit,omitempty" json:"Credit"`
      Supervision *AUCodeSetsTeacherCoverSupervisionType `xml:"Supervision,omitempty" json:"Supervision"`
      Weighting *float64 `xml:"Weighting" json:"Weighting"`
      
      }
    
    type RoomListType struct {
        RoomInfoRefId []string `xml:"RoomInfoRefId" json:"RoomInfoRefId"`
      
      }
    
    type StaffListType struct {
        StaffPersonalRefId []string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      
      }
    
    type AuthorsType struct {
        Author []string `xml:"Author" json:"Author"`
      
      }
    
    type OrganizationsType struct {
        Organization []string `xml:"Organization" json:"Organization"`
      
      }
    
    type PurchasingItemsType struct {
        PurchasingItem []PurchasingItemType `xml:"PurchasingItem" json:"PurchasingItem"`
      
      }
    
    type PurchasingItemType struct {
        ItemNumber *string `xml:"ItemNumber,omitempty" json:"ItemNumber"`
      ItemDescription *string `xml:"ItemDescription" json:"ItemDescription"`
      LocalItemId *LocalIdType `xml:"LocalItemId,omitempty" json:"LocalItemId"`
      Quantity *string `xml:"Quantity,omitempty" json:"Quantity"`
      UnitCost *MonetaryAmountType `xml:"UnitCost,omitempty" json:"UnitCost"`
      TotalCost *MonetaryAmountType `xml:"TotalCost,omitempty" json:"TotalCost"`
      QuantityDelivered *string `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered"`
      CancelledOrder *AUCodeSetsYesOrNoCategoryType `xml:"CancelledOrder,omitempty" json:"CancelledOrder"`
      TaxRate *float64 `xml:"TaxRate" json:"TaxRate"`
      ExpenseAccounts *ExpenseAccountsType `xml:"ExpenseAccounts,omitempty" json:"ExpenseAccounts"`
      
      }
    
    type ExpenseAccountsType struct {
        ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount" json:"ExpenseAccount"`
      
      }
    
    type ExpenseAccountType struct {
        AccountCode *string `xml:"AccountCode" json:"AccountCode"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      FinancialAccountRefId *string `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      
      }
    
    type SchoolProgramListType struct {
        Program []SchoolProgramType `xml:"Program" json:"Program"`
      
      }
    
    type SchoolProgramType struct {
        Category *string `xml:"Category,omitempty" json:"Category"`
      Type *string `xml:"Type" json:"Type"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LearningObjectivesType struct {
        LearningObjective []string `xml:"LearningObjective" json:"LearningObjective"`
      
      }
    
    type RecognitionListType struct {
        Recognition []string `xml:"Recognition" json:"Recognition"`
      
      }
    
    type LResourcesType struct {
        LearningResourceRefId []ResourcesType `xml:"LearningResourceRefId" json:"LearningResourceRefId"`
      
      }
    
    type ResourcesType struct {
          ResourceType *string `xml:"ResourceType,attr" json:"ResourceType"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type SourceObjectsType struct {
      SourceObject []SourceObjectsType_SourceObject `xml:"SourceObject" json:"SourceObject"`
      
      }
    
    type StudentsType struct {
        StudentPersonalRefId []string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      
      }
    
    type PrerequisitesType struct {
        Prerequisite []string `xml:"Prerequisite" json:"Prerequisite"`
      
      }
    
    type EssentialMaterialsType struct {
        EssentialMaterial []string `xml:"EssentialMaterial" json:"EssentialMaterial"`
      
      }
    
    type TechnicalRequirementsType struct {
        TechnicalRequirement *string `xml:"TechnicalRequirement,omitempty" json:"TechnicalRequirement"`
      
      }
    
    type SoftwareRequirementListType struct {
        SoftwareRequirement []SoftwareRequirementType `xml:"SoftwareRequirement" json:"SoftwareRequirement"`
      
      }
    
    type SoftwareRequirementType struct {
        SoftwareTitle *string `xml:"SoftwareTitle" json:"SoftwareTitle"`
      Version *string `xml:"Version,omitempty" json:"Version"`
      Vendor *string `xml:"Vendor,omitempty" json:"Vendor"`
      OS *string `xml:"OS,omitempty" json:"OS"`
      
      }
    
    type HouseholdListType struct {
        Household []LocalIdType `xml:"Household" json:"Household"`
      
      }
    
    type StudentSubjectChoiceListType struct {
        StudentSubjectChoice []StudentSubjectChoiceType `xml:"StudentSubjectChoice" json:"StudentSubjectChoice"`
      
      }
    
    type StudentSubjectChoiceType struct {
        PreferenceNumber *int `xml:"PreferenceNumber" json:"PreferenceNumber"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId" json:"SubjectLocalId"`
      StudyDescription *SubjectAreaType `xml:"StudyDescription,omitempty" json:"StudyDescription"`
      OtherSchoolLocalId *LocalIdType `xml:"OtherSchoolLocalId,omitempty" json:"OtherSchoolLocalId"`
      
      }
    
    type IdentityAssertionsType struct {
      IdentityAssertion []IdentityAssertionsType_IdentityAssertion `xml:"IdentityAssertion" json:"IdentityAssertion"`
      
      }
    
    type LearningStandardsType struct {
        LearningStandardItemRefId []string `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      
      }
    
    type LearningResourcesType struct {
        LearningResourceRefId []string `xml:"LearningResourceRefId" json:"LearningResourceRefId"`
      
      }
    
    type LearningStandardsDocumentType struct {
        LearningStandardDocumentRefId []string `xml:"LearningStandardDocumentRefId" json:"LearningStandardDocumentRefId"`
      
      }
    
    type ComponentsType struct {
        Component []ComponentType `xml:"Component" json:"Component"`
      
      }
    
    type ComponentType struct {
        Name *string `xml:"Name" json:"Name"`
      Reference *string `xml:"Reference" json:"Reference"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      Strategies *StrategiesType `xml:"Strategies,omitempty" json:"Strategies"`
      AssociatedObjects *AssociatedObjectsType `xml:"AssociatedObjects,omitempty" json:"AssociatedObjects"`
      
      }
    
    type StrategiesType struct {
        Strategy []string `xml:"Strategy" json:"Strategy"`
      
      }
    
    type AssociatedObjectsType struct {
      AssociatedObject []AssociatedObjectsType_AssociatedObject `xml:"AssociatedObject" json:"AssociatedObject"`
      
      }
    
    type EvaluationsType struct {
        Evaluation []EvaluationType `xml:"Evaluation" json:"Evaluation"`
      
      }
    
    type EvaluationType struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      Date *string `xml:"Date,omitempty" json:"Date"`
      Name *NameType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type ApprovalsType struct {
        Approval []ApprovalType `xml:"Approval" json:"Approval"`
      
      }
    
    type ApprovalType struct {
        Organization *string `xml:"Organization" json:"Organization"`
      Date *string `xml:"Date" json:"Date"`
      
      }
    
    type MediaTypesType struct {
        MediaType []string `xml:"MediaType" json:"MediaType"`
      
      }
    
    type LEAContactListType struct {
        LEAContact []LEAContactType `xml:"LEAContact" json:"LEAContact"`
      
      }
    
    type LEAContactType struct {
        PublishInDirectory *PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory"`
      ContactInfo *ContactInfoType `xml:"ContactInfo" json:"ContactInfo"`
      
      }
    
    type FinancialAccountRefIdListType struct {
        FinancialAccountRefId []string `xml:"FinancialAccountRefId" json:"FinancialAccountRefId"`
      
      }
    
    type AccountCodeListType struct {
        AccountCode []string `xml:"AccountCode" json:"AccountCode"`
      
      }
    
    type JournalAdjustmentListType struct {
        JournalAdjustment []JournalAdjustmentType `xml:"JournalAdjustment" json:"JournalAdjustment"`
      
      }
    
    type JournalAdjustmentType struct {
        DebitFinancialAccountRefId *string `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId *string `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId"`
      DebitAccountCode *string `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode"`
      CreditAccountCode *string `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode"`
      GSTCodeOriginal *string `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal"`
      GSTCodeReplacement *string `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement"`
      LineAdjustmentAmount *MonetaryAmountType `xml:"LineAdjustmentAmount" json:"LineAdjustmentAmount"`
      
      }
    
    type PaymentReceiptLineListType struct {
        PaymentReceiptLine []PaymentReceiptLineType `xml:"PaymentReceiptLine" json:"PaymentReceiptLine"`
      
      }
    
    type PaymentReceiptLineType struct {
        InvoiceRefId *string `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      LocalPaymentReceiptLineId *LocalIdType `xml:"LocalPaymentReceiptLineId,omitempty" json:"LocalPaymentReceiptLineId"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      FinancialAccountRefId *string `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      AccountCode *string `xml:"AccountCode,omitempty" json:"AccountCode"`
      TransactionDescription *string `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      TaxRate *float64 `xml:"TaxRate" json:"TaxRate"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      
      }
    
    type PasswordListType struct {
      Password []PasswordListType_Password `xml:"Password" json:"Password"`
      
      }
    
    type ExclusionRulesType struct {
        ExclusionRule []ExclusionRuleType `xml:"ExclusionRule" json:"ExclusionRule"`
      
      }
    
    type ExclusionRuleType struct {
          Type *string `xml:"Type,attr" json:"Type"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type CharacteristicsType struct {
        AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId" json:"AggregateCharacteristicInfoRefId"`
      
      }
    
    type ContactsType struct {
        Contact []ContactType `xml:"Contact" json:"Contact"`
      
      }
    
    type ContactType struct {
        Name *NameType `xml:"Name,omitempty" json:"Name"`
      Address *AddressType `xml:"Address,omitempty" json:"Address"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      Email *EmailType `xml:"Email,omitempty" json:"Email"`
      
      }
    
    type TeachingGroupPeriodListType struct {
        TeachingGroupPeriod []TeachingGroupPeriodType `xml:"TeachingGroupPeriod" json:"TeachingGroupPeriod"`
      
      }
    
    type TeachingGroupPeriodType struct {
        TimeTableCellRefId *string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      StartTime *string `xml:"StartTime" json:"StartTime"`
      CellType *string `xml:"CellType,omitempty" json:"CellType"`
      
      }
    
    type TeacherListType struct {
        TeachingGroupTeacher []TeachingGroupTeacherType `xml:"TeachingGroupTeacher" json:"TeachingGroupTeacher"`
      
      }
    
    type TeachingGroupTeacherType struct {
        StaffPersonalRefId *string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId" json:"StaffLocalId"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      Association *string `xml:"Association" json:"Association"`
      
      }
    
    type StudentListType struct {
        TeachingGroupStudent []TeachingGroupStudentType `xml:"TeachingGroupStudent" json:"TeachingGroupStudent"`
      
      }
    
    type TeachingGroupStudentType struct {
        StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId" json:"StudentLocalId"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type TimeTableDayListType struct {
        TimeTableDay []TimeTableDayType `xml:"TimeTableDay" json:"TimeTableDay"`
      
      }
    
    type TimeTableDayType struct {
        DayId *LocalIdType `xml:"DayId" json:"DayId"`
      DayTitle *string `xml:"DayTitle" json:"DayTitle"`
      TimeTablePeriodList *TimeTablePeriodListType `xml:"TimeTablePeriodList" json:"TimeTablePeriodList"`
      
      }
    
    type TimeTablePeriodListType struct {
        TimeTablePeriod []TimeTablePeriodType `xml:"TimeTablePeriod" json:"TimeTablePeriod"`
      
      }
    
    type TimeTablePeriodType struct {
        PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      PeriodTitle *string `xml:"PeriodTitle" json:"PeriodTitle"`
      BellPeriod *string `xml:"BellPeriod,omitempty" json:"BellPeriod"`
      StartTime *string `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime *string `xml:"EndTime,omitempty" json:"EndTime"`
      RegularSchoolPeriod *string `xml:"RegularSchoolPeriod,omitempty" json:"RegularSchoolPeriod"`
      InstructionalMinutes *int `xml:"InstructionalMinutes" json:"InstructionalMinutes"`
      UseInAttendanceCalculations *string `xml:"UseInAttendanceCalculations,omitempty" json:"UseInAttendanceCalculations"`
      
      }
    
    type NAPLANClassListType struct {
        ClassCode []string `xml:"ClassCode" json:"ClassCode"`
      
      }
    
    type SchoolGroupListType struct {
        SchoolGroup []LocalIdType `xml:"SchoolGroup" json:"SchoolGroup"`
      
      }
    
    type YearLevelEnrollmentListType struct {
        YearLevelEnrollment []YearLevelEnrollmentType `xml:"YearLevelEnrollment" json:"YearLevelEnrollment"`
      
      }
    
    type YearLevelEnrollmentType struct {
        Year *AUCodeSetsYearLevelCodeType `xml:"Year" json:"Year"`
      Enrollment *string `xml:"Enrollment" json:"Enrollment"`
      
      }
    
    type SchoolFocusListType struct {
        SchoolFocus []AUCodeSetsSchoolFocusCodeType `xml:"SchoolFocus" json:"SchoolFocus"`
      
      }
    
    type AlertMessagesType struct {
        AlertMessage []AlertMessageType `xml:"AlertMessage" json:"AlertMessage"`
      
      }
    
    type AlertMessageType struct {
          Type *string `xml:"Type,attr" json:"Type"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type MedicalAlertMessagesType struct {
        MedicalAlertMessage []MedicalAlertMessageType `xml:"MedicalAlertMessage" json:"MedicalAlertMessage"`
      
      }
    
    type MedicalAlertMessageType struct {
          Severity *string `xml:"Severity,attr" json:"Severity"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type OtherIdListType struct {
        OtherId []OtherIdType `xml:"OtherId" json:"OtherId"`
      
      }
    
    type OtherIdType struct {
          Type *string `xml:"Type,attr" json:"Type"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type BaseNameType struct {
        Title *string `xml:"Title,omitempty" json:"Title"`
      FamilyName *string `xml:"FamilyName" json:"FamilyName"`
      GivenName *string `xml:"GivenName" json:"GivenName"`
      MiddleName *string `xml:"MiddleName,omitempty" json:"MiddleName"`
      FamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"FamilyNameFirst,omitempty" json:"FamilyNameFirst"`
      PreferredFamilyName *string `xml:"PreferredFamilyName,omitempty" json:"PreferredFamilyName"`
      PreferredFamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"PreferredFamilyNameFirst,omitempty" json:"PreferredFamilyNameFirst"`
      PreferredGivenName *string `xml:"PreferredGivenName,omitempty" json:"PreferredGivenName"`
      Suffix *string `xml:"Suffix,omitempty" json:"Suffix"`
      FullName *string `xml:"FullName" json:"FullName"`
      
      }
    
    type NameOfRecordType struct {
        BaseNameType
          Type *string `xml:"Type,attr" json:"Type"`
      
      }
    
    type OtherNameType struct {
        BaseNameType
          Type *AUCodeSetsNameUsageTypeType `xml:"Type,attr" json:"Type"`
      
      }
    
    type PartialDateType string
    type LocalIdType string
    type LocationType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      LocationName *string `xml:"LocationName,omitempty" json:"LocationName"`
      LocationRefId *LocationType_LocationRefId
      
      }
    
    type StateProvinceIdType string
    type AttendanceCodeType struct {
        Code *AUCodeSetsAttendanceCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type YearLevelType struct {
        Code *AUCodeSetsYearLevelCodeType `xml:"Code" json:"Code"`
      
      }
    
    type PersonInfoType struct {
        Name *NameOfRecordType `xml:"Name" json:"Name"`
      OtherNames *OtherNamesType `xml:"OtherNames,omitempty" json:"OtherNames"`
      Demographics *DemographicsType `xml:"Demographics,omitempty" json:"Demographics"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      HouseholdContactInfoList *HouseholdContactInfoListType `xml:"HouseholdContactInfoList,omitempty" json:"HouseholdContactInfoList"`
      
      }
    
    type YearLevelsType struct {
        YearLevel []YearLevelType `xml:"YearLevel" json:"YearLevel"`
      
      }
    
    type SchoolURLType string
    type PrincipalInfoType struct {
        ContactName *NameOfRecordType `xml:"ContactName" json:"ContactName"`
      ContactTitle *string `xml:"ContactTitle,omitempty" json:"ContactTitle"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      
      }
    
    type SchoolContactType struct {
        PublishInDirectory *PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory"`
      ContactInfo *ContactInfoType `xml:"ContactInfo" json:"ContactInfo"`
      
      }
    
    type SchoolContactListType struct {
        SchoolContact []SchoolContactType `xml:"SchoolContact" json:"SchoolContact"`
      
      }
    
    type PublishInDirectoryType AUCodeSetsYesOrNoCategoryType
    type ContactInfoType struct {
        Name *NameType `xml:"Name" json:"Name"`
      PositionTitle *string `xml:"PositionTitle,omitempty" json:"PositionTitle"`
      Role *string `xml:"Role,omitempty" json:"Role"`
      RegistrationDetails *string `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails"`
      Qualifications *string `xml:"Qualifications,omitempty" json:"Qualifications"`
      Address *AddressType `xml:"Address,omitempty" json:"Address"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      
      }
    
    type AddressStreetType struct {
        Line1 *string `xml:"Line1,omitempty" json:"Line1"`
      Line2 *string `xml:"Line2,omitempty" json:"Line2"`
      Line3 *string `xml:"Line3,omitempty" json:"Line3"`
      Complex *string `xml:"Complex,omitempty" json:"Complex"`
      StreetNumber *string `xml:"StreetNumber,omitempty" json:"StreetNumber"`
      StreetPrefix *string `xml:"StreetPrefix,omitempty" json:"StreetPrefix"`
      StreetName *string `xml:"StreetName,omitempty" json:"StreetName"`
      StreetType *string `xml:"StreetType,omitempty" json:"StreetType"`
      StreetSuffix *string `xml:"StreetSuffix,omitempty" json:"StreetSuffix"`
      ApartmentType *string `xml:"ApartmentType,omitempty" json:"ApartmentType"`
      ApartmentNumberPrefix *string `xml:"ApartmentNumberPrefix,omitempty" json:"ApartmentNumberPrefix"`
      ApartmentNumber *string `xml:"ApartmentNumber,omitempty" json:"ApartmentNumber"`
      ApartmentNumberSuffix *string `xml:"ApartmentNumberSuffix,omitempty" json:"ApartmentNumberSuffix"`
      
      }
    
    type AddressType struct {
        Type *AUCodeSetsAddressTypeType `xml:"Type,attr" json:"Type"`
      Role *AUCodeSetsAddressRoleType `xml:"Role,attr" json:"Role"`
      EffectiveFromDate *string `xml:"EffectiveFromDate,omitempty" json:"EffectiveFromDate"`
      EffectiveToDate *string `xml:"EffectiveToDate,omitempty" json:"EffectiveToDate"`
      Street *AddressStreetType `xml:"Street" json:"Street"`
      City *string `xml:"City" json:"City"`
      StateProvince *StateProvinceType `xml:"StateProvince" json:"StateProvince"`
      Country *CountryType `xml:"Country,omitempty" json:"Country"`
      PostalCode *string `xml:"PostalCode" json:"PostalCode"`
      GridLocation *GridLocationType `xml:"GridLocation,omitempty" json:"GridLocation"`
      MapReference *MapReferenceType `xml:"MapReference,omitempty" json:"MapReference"`
      RadioContact *string `xml:"RadioContact,omitempty" json:"RadioContact"`
      Community *string `xml:"Community,omitempty" json:"Community"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      AddressGlobalUID *GUIDType `xml:"AddressGlobalUID,omitempty" json:"AddressGlobalUID"`
      StatisticalAreas *StatisticalAreasType `xml:"StatisticalAreas,omitempty" json:"StatisticalAreas"`
      
      }
    
    type MapReferenceType struct {
        Type *string `xml:"Type,attr" json:"Type"`
      XCoordinate *string `xml:"XCoordinate" json:"XCoordinate"`
      YCoordinate *string `xml:"YCoordinate" json:"YCoordinate"`
      MapNumber *string `xml:"MapNumber,omitempty" json:"MapNumber"`
      
      }
    
    type StatisticalAreasType struct {
        StatisticalArea []StatisticalAreaType `xml:"StatisticalArea" json:"StatisticalArea"`
      
      }
    
    type StatisticalAreaType struct {
          SpatialUnitType *string `xml:"SpatialUnitType,attr" json:"SpatialUnitType"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type AddressListType struct {
        Address []AddressType `xml:"Address" json:"Address"`
      
      }
    
    type EmailListType struct {
        Email []EmailType `xml:"Email" json:"Email"`
      
      }
    
    type EmailType struct {
          Type *AUCodeSetsEmailTypeType `xml:"Type,attr" json:"Type"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type PhoneNumberListType struct {
        PhoneNumber []PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type PhoneNumberType struct {
        Type *AUCodeSetsTelephoneNumberTypeType `xml:"Type,attr" json:"Type"`
      Number *string `xml:"Number" json:"Number"`
      Extension *string `xml:"Extension,omitempty" json:"Extension"`
      ListedStatus *AUCodeSetsYesOrNoCategoryType `xml:"ListedStatus,omitempty" json:"ListedStatus"`
      Preference *int `xml:"Preference" json:"Preference"`
      
      }
    
    type CountryType AUCodeSetsStandardAustralianClassificationOfCountriesSACCType
    type GridLocationType struct {
        Latitude *float64 `xml:"Latitude" json:"Latitude"`
      Longitude *float64 `xml:"Longitude" json:"Longitude"`
      
      }
    
    type OperationalStatusType AUCodeSetsOperationalStatusType
    type StateProvinceType string
    type SchoolYearType string
    type ElectronicIdListType struct {
        ElectronicId []ElectronicIdType `xml:"ElectronicId" json:"ElectronicId"`
      
      }
    
    type ElectronicIdType struct {
          Type *AUCodeSetsElectronicIdTypeType `xml:"Type,attr" json:"Type"`
      
        Value *string `xml:",chardata" json:"value"`
      }
    
    type OtherNamesType struct {
        Name []OtherNameType `xml:"Name" json:"Name"`
      
      }
    
    type DemographicsType struct {
        IndigenousStatus *AUCodeSetsIndigenousStatusType `xml:"IndigenousStatus,omitempty" json:"IndigenousStatus"`
      Sex *AUCodeSetsSexCodeType `xml:"Sex,omitempty" json:"Sex"`
      BirthDate *BirthDateType `xml:"BirthDate,omitempty" json:"BirthDate"`
      DateOfDeath *string `xml:"DateOfDeath,omitempty" json:"DateOfDeath"`
      BirthDateVerification *AUCodeSetsBirthdateVerificationType `xml:"BirthDateVerification,omitempty" json:"BirthDateVerification"`
      PlaceOfBirth *string `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth"`
      StateOfBirth *StateProvinceType `xml:"StateOfBirth,omitempty" json:"StateOfBirth"`
      CountryOfBirth *CountryType `xml:"CountryOfBirth,omitempty" json:"CountryOfBirth"`
      CountriesOfCitizenship *CountryListType `xml:"CountriesOfCitizenship,omitempty" json:"CountriesOfCitizenship"`
      CountriesOfResidency *CountryList2Type `xml:"CountriesOfResidency,omitempty" json:"CountriesOfResidency"`
      CountryArrivalDate *string `xml:"CountryArrivalDate,omitempty" json:"CountryArrivalDate"`
      AustralianCitizenshipStatus *AUCodeSetsAustralianCitizenshipStatusType `xml:"AustralianCitizenshipStatus,omitempty" json:"AustralianCitizenshipStatus"`
      EnglishProficiency *EnglishProficiencyType `xml:"EnglishProficiency,omitempty" json:"EnglishProficiency"`
      LanguageList *LanguageListType `xml:"LanguageList,omitempty" json:"LanguageList"`
      DwellingArrangement *DwellingArrangementType `xml:"DwellingArrangement,omitempty" json:"DwellingArrangement"`
      Religion *ReligionType `xml:"Religion,omitempty" json:"Religion"`
      ReligiousEventList *ReligiousEventListType `xml:"ReligiousEventList,omitempty" json:"ReligiousEventList"`
      ReligiousRegion *string `xml:"ReligiousRegion,omitempty" json:"ReligiousRegion"`
      PermanentResident *AUCodeSetsPermanentResidentStatusType `xml:"PermanentResident,omitempty" json:"PermanentResident"`
      VisaSubClass *VisaSubClassCodeType `xml:"VisaSubClass,omitempty" json:"VisaSubClass"`
      VisaStatisticalCode *string `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      VisaExpiryDate *string `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      VisaSubClassList *VisaSubClassListType `xml:"VisaSubClassList,omitempty" json:"VisaSubClassList"`
      LBOTE *AUCodeSetsYesOrNoCategoryType `xml:"LBOTE,omitempty" json:"LBOTE"`
      InterpreterRequired *AUCodeSetsYesOrNoCategoryType `xml:"InterpreterRequired,omitempty" json:"InterpreterRequired"`
      ImmunisationCertificateStatus *AUCodeSetsImmunisationCertificateStatusType `xml:"ImmunisationCertificateStatus,omitempty" json:"ImmunisationCertificateStatus"`
      CulturalBackground *AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType `xml:"CulturalBackground,omitempty" json:"CulturalBackground"`
      MaritalStatus *AUCodeSetsMaritalStatusAIHWType `xml:"MaritalStatus,omitempty" json:"MaritalStatus"`
      MedicareNumber *string `xml:"MedicareNumber,omitempty" json:"MedicareNumber"`
      
      }
    
    type EnglishProficiencyType struct {
        Code *AUCodeSetsEnglishProficiencyType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LanguageListType struct {
        Language []LanguageBaseType `xml:"Language" json:"Language"`
      
      }
    
    type BirthDateType string
    type ProjectedGraduationYearType string
    type OnTimeGraduationYearType string
    type RelationshipType struct {
        Code *AUCodeSetsRelationshipToStudentType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type EducationalLevelType AUCodeSetsSchoolEducationLevelTypeType
    type GraduationDateType PartialDateType
    type NameType struct {
        BaseNameType
          Type *AUCodeSetsNameUsageTypeType `xml:"Type,attr" json:"Type"`
      
      }
    
    type HomeroomNumberType string
    type TimeElementType struct {
        Type *string `xml:"Type" json:"Type"`
      Code *string `xml:"Code" json:"Code"`
      Name *string `xml:"Name" json:"Name"`
      Value *string `xml:"Value" json:"Value"`
      StartDateTime *string `xml:"StartDateTime,omitempty" json:"StartDateTime"`
      EndDateTime *string `xml:"EndDateTime,omitempty" json:"EndDateTime"`
      SpanGaps *TimeElementType_SpanGaps
      IsCurrent *bool `xml:"IsCurrent" json:"IsCurrent"`
      
      }
    
    type LifeCycleType struct {
      Created *LifeCycleType_Created
      ModificationHistory *LifeCycleType_ModificationHistory
      TimeElements *LifeCycleType_TimeElements
      
      }
    
    type OtherCodeListType struct {
      OtherCode []OtherCodeListType_OtherCode `xml:"OtherCode" json:"OtherCode"`
      
      }
    
    type ProgramStatusType struct {
        Code *string `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type SubjectAreaListType struct {
        SubjectArea []SubjectAreaType `xml:"SubjectArea" json:"SubjectArea"`
      
      }
    
    type ACStrandAreaListType struct {
        ACStrandSubjectArea []ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      
      }
    
    type SubjectAreaType struct {
        Code *string `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ACStrandSubjectAreaType struct {
        ACStrand *AUCodeSetsACStrandType `xml:"ACStrand" json:"ACStrand"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      
      }
    
    type EducationFilterType struct {
        LearningStandardItems *LearningStandardsType `xml:"LearningStandardItems,omitempty" json:"LearningStandardItems"`
      
      }
    
    type SIF_ExtendedElementsType struct {
      SIF_ExtendedElement []SIF_ExtendedElementsType_SIF_ExtendedElement `xml:"SIF_ExtendedElement" json:"SIF_ExtendedElement"`
      
      }
    
    type SIF_MetadataType struct {
      TimeElements *SIF_MetadataType_TimeElements
      LifeCycle *LifeCycleType `xml:"LifeCycle,omitempty" json:"LifeCycle"`
      EducationFilter *EducationFilterType `xml:"EducationFilter,omitempty" json:"EducationFilter"`
      
      }
    type AbstractContentPackageType_XMLData struct {
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_TextData struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *string `xml:"FileName,attr" json:"FileName"`
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_BinaryData struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *string `xml:"FileName,attr" json:"FileName"`
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_Reference struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      Description *string `xml:"Description,attr" json:"Description"`
       URL *string `xml:"URL" json:"URL"`
}
type AbstractContentElementType_XMLData struct {
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentElementType_TextData struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *string `xml:"FileName,attr" json:"FileName"`
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentElementType_BinaryData struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *string `xml:"FileName,attr" json:"FileName"`
      Description *string `xml:"Description,attr" json:"Description"`
      Value *string `xml:",chardata" json:"value"`
}
type AbstractContentElementType_Reference struct {
      MIMEType *string `xml:"MIMEType,attr" json:"MIMEType"`
      Description *string `xml:"Description,attr" json:"Description"`
       URL *string `xml:"URL" json:"URL"`
}
type PersonInvolvementType_PersonRefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type ActivityTimeType_Duration struct {
      Units *string `xml:"Units,attr" json:"Units"`
      Value *int `xml:",chardata" json:"value"`
}
type SourceObjectsType_SourceObject struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type IdentityAssertionsType_IdentityAssertion struct {
      SchemaName *string `xml:"SchemaName,attr" json:"SchemaName"`
      Value *string `xml:",chardata" json:"value"`
}
type AssociatedObjectsType_AssociatedObject struct {
      SIF_RefObject *ObjectNameType `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type PasswordListType_Password struct {
      Algorithm *string `xml:"Algorithm,attr" json:"Algorithm"`
      KeyName *string `xml:"KeyName,attr" json:"KeyName"`
      Value *string `xml:",chardata" json:"value"`
}
type LocationType_LocationRefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
type TimeElementType_SpanGaps struct {
      SpanGap []TimeElementType_SpanGap `xml:"SpanGap" json:"SpanGap"`
}
type LifeCycleType_Created struct {
       DateTime *string `xml:"DateTime" json:"DateTime"`
      Creators *LifeCycleType_Creators
}
type LifeCycleType_ModificationHistory struct {
      Modified []LifeCycleType_Modified `xml:"Modified" json:"Modified"`
}
type LifeCycleType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type OtherCodeListType_OtherCode struct {
      Codeset *string `xml:"Codeset,attr" json:"Codeset"`
      Value *string `xml:",chardata" json:"value"`
}
type SIF_ExtendedElementsType_SIF_ExtendedElement struct {
      Name *string `xml:"Name,attr" json:"Name"`
      Type *string `xml:"Type,attr" json:"Type"`
      SIF_Action *string `xml:"SIF_Action,attr" json:"SIF_Action"`
      Value *ExtendedContentType `xml:",chardata" json:"value"`
}
type SIF_MetadataType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type TimeElementType_SpanGap struct {
       Type *string `xml:"Type" json:"Type"`
       Code *string `xml:"Code" json:"Code"`
       Name *string `xml:"Name" json:"Name"`
       Value *string `xml:"Value" json:"Value"`
       StartDateTime *string `xml:"StartDateTime,omitempty" json:"StartDateTime"`
       EndDateTime *string `xml:"EndDateTime,omitempty" json:"EndDateTime"`
}
type LifeCycleType_Creators struct {
      Creator []LifeCycleType_Creator `xml:"Creator" json:"Creator"`
}
type LifeCycleType_Modified struct {
       By *string `xml:"By" json:"By"`
       DateTime *string `xml:"DateTime" json:"DateTime"`
       Description *string `xml:"Description,omitempty" json:"Description"`
}
type LifeCycleType_Creator struct {
       Name *string `xml:"Name" json:"Name"`
       ID *string `xml:"ID" json:"ID"`
}
