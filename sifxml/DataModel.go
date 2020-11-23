package sifxml


    type ReportPackageType AbstractContentPackageType
    type AbstractContentPackageType struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      XMLData *AbstractContentPackageType_XMLData
      TextData *AbstractContentPackageType_TextData
      BinaryData *AbstractContentPackageType_BinaryData
      Reference *AbstractContentPackageType_Reference
      
      }
    
    type AbstractContentElementType struct {
      XMLData *AbstractContentElementType_XMLData
      TextData *AbstractContentElementType_TextData
      BinaryData *AbstractContentElementType_BinaryData
      Reference *AbstractContentElementType_Reference
      
      }
    
    type MonetaryAmountType struct {
          Currency *ISO4217CurrencyNamesAndCodeElementsType `xml:"Currency,attr" json:"Currency"`
      
        Value *Float `xml:",chardata" json:"value"`
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
        Type *String `xml:"Type,attr" json:"Type"`
      Title *String `xml:"Title" json:"Title"`
      Author *String `xml:"Author,omitempty" json:"Author"`
      ElectronicId *ElectronicIdType `xml:"ElectronicId,omitempty" json:"ElectronicId"`
      CallNumber *String `xml:"CallNumber,omitempty" json:"CallNumber"`
      ISBN *String `xml:"ISBN,omitempty" json:"ISBN"`
      Cost *MonetaryAmountType `xml:"Cost,omitempty" json:"Cost"`
      ReplacementCost *MonetaryAmountType `xml:"ReplacementCost,omitempty" json:"ReplacementCost"`
      
      }
    
    type CheckoutInfoType struct {
        CheckedOutOn *String `xml:"CheckedOutOn" json:"CheckedOutOn"`
      ReturnBy *String `xml:"ReturnBy" json:"ReturnBy"`
      RenewalCount *Int `xml:"RenewalCount,omitempty" json:"RenewalCount"`
      
      }
    
    type FineInfoListType struct {
        FineInfo []FineInfoType `xml:"FineInfo" json:"FineInfo"`
      
      }
    
    type FineInfoType struct {
        Type *String `xml:"Type,attr" json:"Type"`
      Assessed *String `xml:"Assessed" json:"Assessed"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      Reference *String `xml:"Reference,omitempty" json:"Reference"`
      
      }
    
    type HoldInfoListType struct {
        HoldInfo []HoldInfoType `xml:"HoldInfo" json:"HoldInfo"`
      
      }
    
    type HoldInfoType struct {
        Type *String `xml:"Type,attr" json:"Type"`
      DatePlaced *String `xml:"DatePlaced" json:"DatePlaced"`
      DateNeeded *String `xml:"DateNeeded,omitempty" json:"DateNeeded"`
      ReservationExpiry *String `xml:"ReservationExpiry,omitempty" json:"ReservationExpiry"`
      MadeAvailable *String `xml:"MadeAvailable,omitempty" json:"MadeAvailable"`
      Expires *String `xml:"Expires,omitempty" json:"Expires"`
      
      }
    
    type LibraryMessageListType struct {
        Message []LibraryMessageType `xml:"Message" json:"Message"`
      
      }
    
    type LibraryMessageType struct {
        Priority *String `xml:"Priority,attr" json:"Priority"`
      PriorityCodeset *String `xml:"PriorityCodeset,attr" json:"PriorityCodeset"`
      Sent *String `xml:"Sent,omitempty" json:"Sent"`
      Text *String `xml:"Text" json:"Text"`
      
      }
    
    type StudentAttendanceCollectionReportingListType struct {
        StudentAttendanceCollectionReporting []StudentAttendanceCollectionReportingType `xml:"StudentAttendanceCollectionReporting" json:"StudentAttendanceCollectionReporting"`
      
      }
    
    type StudentAttendanceCollectionReportingType struct {
        EntityLevel *String `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *String `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName"`
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
      StatsIndigenousStudentType *String `xml:"StatsIndigenousStudentType" json:"StatsIndigenousStudentType"`
      CohortGender *String `xml:"CohortGender" json:"CohortGender"`
      DaysInReferencePeriod *Int `xml:"DaysInReferencePeriod" json:"DaysInReferencePeriod"`
      PossibleSchoolDays *Int `xml:"PossibleSchoolDays" json:"PossibleSchoolDays"`
      AttendanceDays *Float `xml:"AttendanceDays" json:"AttendanceDays"`
      AttendanceLess90Percent *Int `xml:"AttendanceLess90Percent" json:"AttendanceLess90Percent"`
      AttendanceGTE90Percent *Int `xml:"AttendanceGTE90Percent" json:"AttendanceGTE90Percent"`
      PossibleSchoolDaysGT90PercentAttendance *Int `xml:"PossibleSchoolDaysGT90PercentAttendance" json:"PossibleSchoolDaysGT90PercentAttendance"`
      
      }
    
    type AddressCollectionReportingListType struct {
        AddressCollectionReporting []AddressCollectionReportingType `xml:"AddressCollectionReporting" json:"AddressCollectionReporting"`
      
      }
    
    type AddressCollectionReportingType struct {
        EntityLevel *String `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *String `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName"`
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
      ReportingParent2 *String `xml:"ReportingParent2" json:"ReportingParent2"`
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
        RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      RoundName *String `xml:"RoundName" json:"RoundName"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      DueDate *String `xml:"DueDate" json:"DueDate"`
      EndDate *String `xml:"EndDate" json:"EndDate"`
      
      }
    
    type AGContextualQuestionListType struct {
        AGContextualQuestion []AGContextualQuestionType `xml:"AGContextualQuestion" json:"AGContextualQuestion"`
      
      }
    
    type AGContextualQuestionType struct {
        AGContextCode *AUCodeSetsAGContextQuestionType `xml:"AGContextCode" json:"AGContextCode"`
      AGAnswer *String `xml:"AGAnswer" json:"AGAnswer"`
      
      }
    
    type CensusReportingListType struct {
        CensusReporting []CensusReportingType `xml:"CensusReporting" json:"CensusReporting"`
      
      }
    
    type CensusReportingType struct {
        EntityLevel *String `xml:"EntityLevel,omitempty" json:"EntityLevel"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName"`
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
      CohortGender *String `xml:"CohortGender" json:"CohortGender"`
      CohortIndigenousType *String `xml:"CohortIndigenousType" json:"CohortIndigenousType"`
      PrimaryFTE *Float `xml:"PrimaryFTE,omitempty" json:"PrimaryFTE"`
      SecondaryFTE *Float `xml:"SecondaryFTE,omitempty" json:"SecondaryFTE"`
      JobFTE *Float `xml:"JobFTE,omitempty" json:"JobFTE"`
      Headcount *Int `xml:"Headcount" json:"Headcount"`
      
      }
    
    type StaffAssignmentMostRecentContainerType struct {
        PrimaryFTE *Float `xml:"PrimaryFTE,omitempty" json:"PrimaryFTE"`
      SecondaryFTE *Float `xml:"SecondaryFTE,omitempty" json:"SecondaryFTE"`
      
      }
    
    type CensusStudentListType struct {
        CensusStudent []CensusStudentType `xml:"CensusStudent" json:"CensusStudent"`
      
      }
    
    type CensusStudentType struct {
        StudentCohortId *LocalIdType `xml:"StudentCohortId" json:"StudentCohortId"`
      YearLevel *YearLevelType `xml:"YearLevel" json:"YearLevel"`
      CensusAge *Int `xml:"CensusAge" json:"CensusAge"`
      CohortGender *String `xml:"CohortGender" json:"CohortGender"`
      CohortIndigenousType *String `xml:"CohortIndigenousType" json:"CohortIndigenousType"`
      EducationMode *String `xml:"EducationMode" json:"EducationMode"`
      StudentOnVisa *String `xml:"StudentOnVisa" json:"StudentOnVisa"`
      OverseasStudent *String `xml:"OverseasStudent" json:"OverseasStudent"`
      DisabilityLevelOfAdjustment *String `xml:"DisabilityLevelOfAdjustment" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *String `xml:"DisabilityCategory" json:"DisabilityCategory"`
      FTE *Float `xml:"FTE" json:"FTE"`
      Headcount *Int `xml:"Headcount" json:"Headcount"`
      
      }
    
    type AGReportingObjectResponseListType struct {
        AGReportingObjectResponse []AGReportingObjectResponseType `xml:"AGReportingObjectResponse" json:"AGReportingObjectResponse"`
      
      }
    
    type AGReportingObjectResponseType struct {
        SubmittedRefId *String `xml:"SubmittedRefId,omitempty" json:"SubmittedRefId"`
      SIFRefId *String `xml:"SIFRefId,omitempty" json:"SIFRefId"`
      HTTPStatusCode *String `xml:"HTTPStatusCode,omitempty" json:"HTTPStatusCode"`
      ErrorText *String `xml:"ErrorText,omitempty" json:"ErrorText"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName"`
      AGSubmissionStatusCode *AUCodeSetsAGSubmissionStatusType `xml:"AGSubmissionStatusCode" json:"AGSubmissionStatusCode"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList"`
      
      }
    
    type FQReportingListType struct {
        FQReporting []FQReportingType `xml:"FQReporting" json:"FQReporting"`
      
      }
    
    type FQReportingType struct {
        EntityLevel *String `xml:"EntityLevel" json:"EntityLevel"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      ACARAId *String `xml:"ACARAId,omitempty" json:"ACARAId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      FQContextualQuestionList *FQContextualQuestionListType `xml:"FQContextualQuestionList,omitempty" json:"FQContextualQuestionList"`
      FQItemList *FQItemListType `xml:"FQItemList" json:"FQItemList"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList"`
      
      }
    
    type FQContextualQuestionListType struct {
        FQContextualQuestion []FQContextualQuestionType `xml:"FQContextualQuestion" json:"FQContextualQuestion"`
      
      }
    
    type FQContextualQuestionType struct {
        FQContext *String `xml:"FQContext" json:"FQContext"`
      FQAnswer *String `xml:"FQAnswer" json:"FQAnswer"`
      
      }
    
    type FQItemListType struct {
        FQItem []FQItemType `xml:"FQItem" json:"FQItem"`
      
      }
    
    type FQItemType struct {
        FQItemCode *String `xml:"FQItemCode" json:"FQItemCode"`
      TuitionAmount *Float `xml:"TuitionAmount,omitempty" json:"TuitionAmount"`
      BoardingAmount *Float `xml:"BoardingAmount,omitempty" json:"BoardingAmount"`
      SystemAmount *Float `xml:"SystemAmount,omitempty" json:"SystemAmount"`
      DioceseAmount *Float `xml:"DioceseAmount,omitempty" json:"DioceseAmount"`
      FQComments *String `xml:"FQComments,omitempty" json:"FQComments"`
      
      }
    
    type AGRuleListType struct {
        AGRule []AGRuleType `xml:"AGRule" json:"AGRule"`
      
      }
    
    type AGRuleType struct {
        AGRuleCode *String `xml:"AGRuleCode,omitempty" json:"AGRuleCode"`
      AGRuleComment *String `xml:"AGRuleComment,omitempty" json:"AGRuleComment"`
      AGRuleResponse *String `xml:"AGRuleResponse,omitempty" json:"AGRuleResponse"`
      AGRuleStatus *String `xml:"AGRuleStatus,omitempty" json:"AGRuleStatus"`
      
      }
    
    type SoftwareVendorInfoContainerType struct {
        SoftwareProduct *String `xml:"SoftwareProduct" json:"SoftwareProduct"`
      SoftwareVersion *String `xml:"SoftwareVersion" json:"SoftwareVersion"`
      
      }
    
    type TimeTableScheduleType struct {
        SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Title *String `xml:"Title" json:"Title"`
      DaysPerCycle *Int `xml:"DaysPerCycle" json:"DaysPerCycle"`
      PeriodsPerDay *Int `xml:"PeriodsPerDay" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *Int `xml:"TeachingPeriodsPerDay,omitempty" json:"TeachingPeriodsPerDay"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolName *String `xml:"SchoolName,omitempty" json:"SchoolName"`
      TimeTableCreationDate *String `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList" json:"TimeTableDayList"`
      
      }
    
    type TimeTableScheduleCellListType struct {
        TimeTableScheduleCell []TimeTableScheduleCellType `xml:"TimeTableScheduleCell" json:"TimeTableScheduleCell"`
      
      }
    
    type TimeTableScheduleCellType struct {
        TimeTableScheduleCellLocalId *LocalIdType `xml:"TimeTableScheduleCellLocalId" json:"TimeTableScheduleCellLocalId"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeachingGroupGUID *String `xml:"TeachingGroupGUID" json:"TeachingGroupGUID"`
      RoomInfoRefId *String `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      TimeTableLocalId *LocalIdType `xml:"TimeTableLocalId,omitempty" json:"TimeTableLocalId"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TeachingGroupLocalId *LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      CellType *String `xml:"CellType" json:"CellType"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
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
      ShortName *String `xml:"ShortName" json:"ShortName"`
      LongName *String `xml:"LongName,omitempty" json:"LongName"`
      GroupType *String `xml:"GroupType,omitempty" json:"GroupType"`
      Sett *String `xml:"Set,omitempty" json:"Set"`
      Block *String `xml:"Block,omitempty" json:"Block"`
      CurriculumLevel *String `xml:"CurriculumLevel,omitempty" json:"CurriculumLevel"`
      SchoolInfoRefId *RefIdType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      SchoolCourseLocalId *LocalIdType `xml:"SchoolCourseLocalId,omitempty" json:"SchoolCourseLocalId"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TimeTableSubjectLocalId *LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId"`
      Semester *Int `xml:"Semester,omitempty" json:"Semester"`
      StudentList *StudentListType `xml:"StudentList,omitempty" json:"StudentList"`
      TeacherList *TeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      MinClassSize *Int `xml:"MinClassSize,omitempty" json:"MinClassSize"`
      MaxClassSize *Int `xml:"MaxClassSize,omitempty" json:"MaxClassSize"`
      TeachingGroupPeriodList *TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList,omitempty" json:"TeachingGroupPeriodList"`
      
      }
    
    type LocalCodeListType struct {
        LocalCode []LocalCodeType `xml:"LocalCode" json:"LocalCode"`
      
      }
    
    type LocalCodeType struct {
        LocalisedCode *String `xml:"LocalisedCode" json:"LocalisedCode"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      Element *String `xml:"Element,omitempty" json:"Element"`
      ListIndex *Int `xml:"ListIndex,omitempty" json:"ListIndex"`
      
      }
    
    type StudentGroupListType struct {
        StudentGroup []StudentGroupType `xml:"StudentGroup" json:"StudentGroup"`
      
      }
    
    type StudentGroupType struct {
        GroupCategory *AUCodeSetsGroupCategoryCodeType `xml:"GroupCategory" json:"GroupCategory"`
      GroupLocalId *LocalIdType `xml:"GroupLocalId" json:"GroupLocalId"`
      GroupDescription *String `xml:"GroupDescription,omitempty" json:"GroupDescription"`
      
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
      PositionTitle *String `xml:"PositionTitle,omitempty" json:"PositionTitle"`
      Role *String `xml:"Role,omitempty" json:"Role"`
      RegistrationDetails *String `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails"`
      Qualifications *String `xml:"Qualifications,omitempty" json:"Qualifications"`
      Address *AddressType `xml:"Address,omitempty" json:"Address"`
      Email *EmailType `xml:"Email" json:"Email"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type CopyRightContainerType struct {
        Date *String `xml:"Date,omitempty" json:"Date"`
      Holder *String `xml:"Holder,omitempty" json:"Holder"`
      
      }
    
    type StandardsSettingBodyType struct {
        Country *CountryType `xml:"Country,omitempty" json:"Country"`
      StateProvince *StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince"`
      SettingBodyName *String `xml:"SettingBodyName,omitempty" json:"SettingBodyName"`
      
      }
    
    type StandardHierarchyLevelType struct {
        Number *Int `xml:"Number" json:"Number"`
      Description *String `xml:"Description" json:"Description"`
      
      }
    
    type StandardIdentifierType struct {
        YearCreated *String `xml:"YearCreated" json:"YearCreated"`
      ACStrandSubjectArea *ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      StandardNumber *String `xml:"StandardNumber" json:"StandardNumber"`
      YearLevels *YearLevelsType `xml:"YearLevels" json:"YearLevels"`
      Benchmark *String `xml:"Benchmark,omitempty" json:"Benchmark"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      IndicatorNumber *String `xml:"IndicatorNumber,omitempty" json:"IndicatorNumber"`
      AlternateIdentificationCodes *AlternateIdentificationCodeListType `xml:"AlternateIdentificationCodes,omitempty" json:"AlternateIdentificationCodes"`
      Organization *String `xml:"Organization" json:"Organization"`
      
      }
    
    type AlternateIdentificationCodeListType struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode" json:"AlternateIdentificationCode"`
      
      }
    
    type RelatedLearningStandardItemRefIdListType struct {
        LearningStandardItemRefId []RelatedLearningStandardItemRefIdType `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      
      }
    
    type RelatedLearningStandardItemRefIdType struct {
          RelationshipType *String `xml:"RelationshipType,attr" json:"RelationshipType"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type ValidLetterMarkListType struct {
        ValidLetterMark []ValidLetterMarkType `xml:"ValidLetterMark" json:"ValidLetterMark"`
      
      }
    
    type ValidLetterMarkType struct {
        Code *String `xml:"Code" json:"Code"`
      NumericEquivalent *Float `xml:"NumericEquivalent,omitempty" json:"NumericEquivalent"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      
      }
    
    type StudentGradeMarkersListType struct {
        Marker []MarkerType `xml:"Marker" json:"Marker"`
      
      }
    
    type MarkerType struct {
        StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Role *String `xml:"Role,omitempty" json:"Role"`
      
      }
    
    type GradingScoreListType struct {
        GradingAssignmentScore []AssignmentScoreType `xml:"GradingAssignmentScore" json:"GradingAssignmentScore"`
      
      }
    
    type AssignmentScoreType struct {
        GradingAssignmentScoreRefId *String `xml:"GradingAssignmentScoreRefId,omitempty" json:"GradingAssignmentScoreRefId"`
      Weight *Float `xml:"Weight,omitempty" json:"Weight"`
      
      }
    
    type GradeType struct {
        Percentage *Float `xml:"Percentage,omitempty" json:"Percentage"`
      Numeric *Float `xml:"Numeric,omitempty" json:"Numeric"`
      Letter *String `xml:"Letter,omitempty" json:"Letter"`
      Narrative *String `xml:"Narrative,omitempty" json:"Narrative"`
      MarkInfoRefId *String `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      
      }
    
    type LearningStandardListType struct {
        LearningStandard []LearningStandardType `xml:"LearningStandard" json:"LearningStandard"`
      
      }
    
    type LearningStandardType struct {
        LearningStandardItemRefId *String `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      LearningStandardURL *String `xml:"LearningStandardURL,omitempty" json:"LearningStandardURL"`
      LearningStandardLocalId *LocalIdType `xml:"LearningStandardLocalId,omitempty" json:"LearningStandardLocalId"`
      
      }
    
    type AssignmentListType struct {
        GradingAssignmentRefId []string `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      
      }
    
    type GenericRubricType struct {
        RubricType *String `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type SymptomListType struct {
        Symptom []string `xml:"Symptom" json:"Symptom"`
      
      }
    
    type MedicationListType struct {
        Medication []MedicationType `xml:"Medication" json:"Medication"`
      
      }
    
    type MedicationType struct {
        MedicationName *String `xml:"MedicationName,omitempty" json:"MedicationName"`
      Dosage *String `xml:"Dosage,omitempty" json:"Dosage"`
      Frequency *String `xml:"Frequency,omitempty" json:"Frequency"`
      AdministrationInformation *String `xml:"AdministrationInformation,omitempty" json:"AdministrationInformation"`
      Method *String `xml:"Method,omitempty" json:"Method"`
      
      }
    
    type WellbeingEventCategoryListType struct {
        WellbeingEventCategory []WellbeingEventCategoryType `xml:"WellbeingEventCategory" json:"WellbeingEventCategory"`
      
      }
    
    type WellbeingEventCategoryType struct {
        EventCategory *String `xml:"EventCategory" json:"EventCategory"`
      WellbeingEventSubCategoryList *WellbeingEventSubCategoryListType `xml:"WellbeingEventSubCategoryList,omitempty" json:"WellbeingEventSubCategoryList"`
      
      }
    
    type WellbeingEventSubCategoryListType struct {
        WellbeingEventSubCategory []string `xml:"WellbeingEventSubCategory" json:"WellbeingEventSubCategory"`
      
      }
    
    type WellbeingEventLocationDetailsType struct {
        EventLocation *AUCodeSetsWellbeingEventLocationType `xml:"EventLocation" json:"EventLocation"`
      Class *String `xml:"Class,omitempty" json:"Class"`
      FurtherLocationNotes *String `xml:"FurtherLocationNotes,omitempty" json:"FurtherLocationNotes"`
      
      }
    
    type FollowUpActionListType struct {
        FollowUpAction []FollowUpActionType `xml:"FollowUpAction" json:"FollowUpAction"`
      
      }
    
    type FollowUpActionType struct {
        WellbeingResponseRefId *String `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      FollowUpDetails *String `xml:"FollowUpDetails,omitempty" json:"FollowUpDetails"`
      FollowUpActionCategory *String `xml:"FollowUpActionCategory,omitempty" json:"FollowUpActionCategory"`
      
      }
    
    type PersonInvolvementListType struct {
        PersonInvolvement []PersonInvolvementType `xml:"PersonInvolvement" json:"PersonInvolvement"`
      
      }
    
    type PersonInvolvementType struct {
      PersonRefId *PersonInvolvementType_PersonRefId
      ShortName *String `xml:"ShortName,omitempty" json:"ShortName"`
      HowInvolved *String `xml:"HowInvolved,omitempty" json:"HowInvolved"`
      
      }
    
    type WithdrawalTimeListType struct {
        Withdrawal []WithdrawalType `xml:"Withdrawal" json:"Withdrawal"`
      
      }
    
    type WithdrawalType struct {
        WithdrawalDate *String `xml:"WithdrawalDate" json:"WithdrawalDate"`
      WithdrawalStartTime *String `xml:"WithdrawalStartTime,omitempty" json:"WithdrawalStartTime"`
      WithdrawalEndTime *String `xml:"WithdrawalEndTime,omitempty" json:"WithdrawalEndTime"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      
      }
    
    type SuspensionContainerType struct {
        SuspensionCategory *AUCodeSetsSuspensionCategoryType `xml:"SuspensionCategory" json:"SuspensionCategory"`
      SuspensionNotes *String `xml:"SuspensionNotes,omitempty" json:"SuspensionNotes"`
      WithdrawalTimeList *WithdrawalTimeListType `xml:"WithdrawalTimeList,omitempty" json:"WithdrawalTimeList"`
      Duration *Float `xml:"Duration,omitempty" json:"Duration"`
      AdvisementDate *String `xml:"AdvisementDate,omitempty" json:"AdvisementDate"`
      ResolutionMeetingTime *String `xml:"ResolutionMeetingTime,omitempty" json:"ResolutionMeetingTime"`
      ResolutionNotes *String `xml:"ResolutionNotes,omitempty" json:"ResolutionNotes"`
      EarlyReturnDate *String `xml:"EarlyReturnDate,omitempty" json:"EarlyReturnDate"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type DetentionContainerType struct {
        DetentionCategory *AUCodeSetsDetentionCategoryType `xml:"DetentionCategory" json:"DetentionCategory"`
      DetentionDate *String `xml:"DetentionDate,omitempty" json:"DetentionDate"`
      DetentionLocation *String `xml:"DetentionLocation,omitempty" json:"DetentionLocation"`
      DetentionNotes *String `xml:"DetentionNotes,omitempty" json:"DetentionNotes"`
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
        PersonalisedPlanRefId *String `xml:"PersonalisedPlanRefId,omitempty" json:"PersonalisedPlanRefId"`
      PlanNotes *String `xml:"PlanNotes,omitempty" json:"PlanNotes"`
      
      }
    
    type AwardContainerType struct {
        AwardDate *String `xml:"AwardDate,omitempty" json:"AwardDate"`
      AwardType *String `xml:"AwardType,omitempty" json:"AwardType"`
      AwardDescription *String `xml:"AwardDescription,omitempty" json:"AwardDescription"`
      AwardNotes *String `xml:"AwardNotes,omitempty" json:"AwardNotes"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type OtherWellbeingResponseContainerType struct {
        OtherResponseDate *String `xml:"OtherResponseDate,omitempty" json:"OtherResponseDate"`
      OtherResponseType *String `xml:"OtherResponseType,omitempty" json:"OtherResponseType"`
      OtherResponseDescription *String `xml:"OtherResponseDescription,omitempty" json:"OtherResponseDescription"`
      OtherResponseNotes *String `xml:"OtherResponseNotes,omitempty" json:"OtherResponseNotes"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type WellbeingDocumentListType struct {
        Document []WellbeingDocumentType `xml:"Document" json:"Document"`
      
      }
    
    type WellbeingDocumentType struct {
        Location *String `xml:"Location" json:"Location"`
      Sensitivity *String `xml:"Sensitivity,omitempty" json:"Sensitivity"`
      URL *String `xml:"URL,omitempty" json:"URL"`
      DocumentType *String `xml:"DocumentType,omitempty" json:"DocumentType"`
      DocumentReviewDate *String `xml:"DocumentReviewDate,omitempty" json:"DocumentReviewDate"`
      DocumentDescription *String `xml:"DocumentDescription,omitempty" json:"DocumentDescription"`
      
      }
    
    type NAPTestItemContentType struct {
        NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      ItemName *String `xml:"ItemName" json:"ItemName"`
      ItemType *AUCodeSetsNAPTestItemTypeType `xml:"ItemType" json:"ItemType"`
      Subdomain *String `xml:"Subdomain" json:"Subdomain"`
      WritingGenre *AUCodeSetsNAPWritingGenreType `xml:"WritingGenre,omitempty" json:"WritingGenre"`
      ItemDescriptor *String `xml:"ItemDescriptor" json:"ItemDescriptor"`
      ReleasedStatus *Bool `xml:"ReleasedStatus" json:"ReleasedStatus"`
      MarkingType *AUCodeSetsNAPTestItemMarkingTypeType `xml:"MarkingType" json:"MarkingType"`
      MultipleChoiceOptionCount *Int `xml:"MultipleChoiceOptionCount,omitempty" json:"MultipleChoiceOptionCount"`
      CorrectAnswer *String `xml:"CorrectAnswer,omitempty" json:"CorrectAnswer"`
      MaximumScore *Float `xml:"MaximumScore" json:"MaximumScore"`
      ItemDifficulty *Float `xml:"ItemDifficulty" json:"ItemDifficulty"`
      ItemDifficultyLogit5 *Float `xml:"ItemDifficultyLogit5" json:"ItemDifficultyLogit5"`
      ItemDifficultyLogit62 *Float `xml:"ItemDifficultyLogit62" json:"ItemDifficultyLogit62"`
      ItemDifficultyLogit5SE *Float `xml:"ItemDifficultyLogit5SE" json:"ItemDifficultyLogit5SE"`
      ItemDifficultyLogit62SE *Float `xml:"ItemDifficultyLogit62SE" json:"ItemDifficultyLogit62SE"`
      ItemProficiencyBand *Int `xml:"ItemProficiencyBand" json:"ItemProficiencyBand"`
      ItemProficiencyLevel *String `xml:"ItemProficiencyLevel,omitempty" json:"ItemProficiencyLevel"`
      ExemplarURL *String `xml:"ExemplarURL,omitempty" json:"ExemplarURL"`
      ItemSubstitutedForList *SubstituteItemListType `xml:"ItemSubstitutedForList,omitempty" json:"ItemSubstitutedForList"`
      ContentDescriptionList *ContentDescriptionListType `xml:"ContentDescriptionList,omitempty" json:"ContentDescriptionList"`
      StimulusList *StimulusListType `xml:"StimulusList,omitempty" json:"StimulusList"`
      NAPWritingRubricList *NAPWritingRubricListType `xml:"NAPWritingRubricList,omitempty" json:"NAPWritingRubricList"`
      
      }
    
    type NAPTestletContentType struct {
        NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletName *String `xml:"TestletName,omitempty" json:"TestletName"`
      Node *String `xml:"Node,omitempty" json:"Node"`
      LocationInStage *Int `xml:"LocationInStage,omitempty" json:"LocationInStage"`
      TestletMaximumScore *Float `xml:"TestletMaximumScore" json:"TestletMaximumScore"`
      
      }
    
    type NAPTestContentType struct {
        NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      TestName *String `xml:"TestName" json:"TestName"`
      TestLevel *YearLevelType `xml:"TestLevel" json:"TestLevel"`
      TestType *AUCodeSetsNAPTestTypeType `xml:"TestType" json:"TestType"`
      Domain *AUCodeSetsNAPTestDomainType `xml:"Domain" json:"Domain"`
      TestYear *SchoolYearType `xml:"TestYear" json:"TestYear"`
      StagesCount *Int `xml:"StagesCount" json:"StagesCount"`
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
        SubstituteItemRefId *String `xml:"SubstituteItemRefId" json:"SubstituteItemRefId"`
      SubstituteItemLocalId *LocalIdType `xml:"SubstituteItemLocalId,omitempty" json:"SubstituteItemLocalId"`
      PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      
      }
    
    type CodeFrameTestItemListType struct {
        TestItem []CodeFrameTestItemType `xml:"TestItem" json:"TestItem"`
      
      }
    
    type CodeFrameTestItemType struct {
        TestItemRefId *String `xml:"TestItemRefId" json:"TestItemRefId"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      TestItemContent *NAPTestItemContentType `xml:"TestItemContent" json:"TestItemContent"`
      
      }
    
    type StimulusLocalIdListType struct {
        StimulusLocalId []LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      
      }
    
    type DomainBandsContainerType struct {
        Band1Lower *Float `xml:"Band1Lower" json:"Band1Lower"`
      Band1Upper *Float `xml:"Band1Upper" json:"Band1Upper"`
      Band2Lower *Float `xml:"Band2Lower" json:"Band2Lower"`
      Band2Upper *Float `xml:"Band2Upper" json:"Band2Upper"`
      Band3Lower *Float `xml:"Band3Lower" json:"Band3Lower"`
      Band3Upper *Float `xml:"Band3Upper" json:"Band3Upper"`
      Band4Lower *Float `xml:"Band4Lower" json:"Band4Lower"`
      Band4Upper *Float `xml:"Band4Upper" json:"Band4Upper"`
      Band5Lower *Float `xml:"Band5Lower" json:"Band5Lower"`
      Band5Upper *Float `xml:"Band5Upper" json:"Band5Upper"`
      Band6Lower *Float `xml:"Band6Lower" json:"Band6Lower"`
      Band6Upper *Float `xml:"Band6Upper" json:"Band6Upper"`
      Band7Lower *Float `xml:"Band7Lower" json:"Band7Lower"`
      Band7Upper *Float `xml:"Band7Upper" json:"Band7Upper"`
      Band8Lower *Float `xml:"Band8Lower" json:"Band8Lower"`
      Band8Upper *Float `xml:"Band8Upper" json:"Band8Upper"`
      Band9Lower *Float `xml:"Band9Lower" json:"Band9Lower"`
      Band9Upper *Float `xml:"Band9Upper" json:"Band9Upper"`
      Band10Lower *Float `xml:"Band10Lower" json:"Band10Lower"`
      Band10Upper *Float `xml:"Band10Upper" json:"Band10Upper"`
      
      }
    
    type DomainProficiencyContainerType struct {
        Level1Lower *Float `xml:"Level1Lower" json:"Level1Lower"`
      Level1Upper *Float `xml:"Level1Upper" json:"Level1Upper"`
      Level2Lower *Float `xml:"Level2Lower" json:"Level2Lower"`
      Level2Upper *Float `xml:"Level2Upper" json:"Level2Upper"`
      Level3Lower *Float `xml:"Level3Lower" json:"Level3Lower"`
      Level3Upper *Float `xml:"Level3Upper" json:"Level3Upper"`
      Level4Lower *Float `xml:"Level4Lower" json:"Level4Lower"`
      Level4Upper *Float `xml:"Level4Upper" json:"Level4Upper"`
      
      }
    
    type NAPTestItemListType struct {
        TestItem []NAPTestItem2Type `xml:"TestItem" json:"TestItem"`
      
      }
    
    type NAPTestItem2Type struct {
        TestItemRefId *String `xml:"TestItemRefId" json:"TestItemRefId"`
      TestItemLocalId *LocalIdType `xml:"TestItemLocalId" json:"TestItemLocalId"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      
      }
    
    type NAPCodeFrameTestletListType struct {
        Testlet []NAPTestletCodeFrameType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletCodeFrameType struct {
        NAPTestletRefId *String `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      TestletContent *NAPTestletContentType `xml:"TestletContent" json:"TestletContent"`
      TestItemList *CodeFrameTestItemListType `xml:"TestItemList" json:"TestItemList"`
      
      }
    
    type NAPStudentResponseTestletListType struct {
        Testlet []NAPTestletResponseType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletResponseType struct {
        NAPTestletRefId *String `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletSubScore *Float `xml:"TestletSubScore,omitempty" json:"TestletSubScore"`
      ItemResponseList *NAPTestletItemResponseListType `xml:"ItemResponseList" json:"ItemResponseList"`
      
      }
    
    type NAPTestletItemResponseListType struct {
        ItemResponse []NAPTestletResponseItemType `xml:"ItemResponse" json:"ItemResponse"`
      
      }
    
    type NAPTestletResponseItemType struct {
        NAPTestItemRefId *String `xml:"NAPTestItemRefId,omitempty" json:"NAPTestItemRefId"`
      NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      Response *String `xml:"Response,omitempty" json:"Response"`
      ResponseCorrectness *AUCodeSetsNAPResponseCorrectnessType `xml:"ResponseCorrectness" json:"ResponseCorrectness"`
      Score *Float `xml:"Score,omitempty" json:"Score"`
      LapsedTimeItem *String `xml:"LapsedTimeItem,omitempty" json:"LapsedTimeItem"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      ItemWeight *Float `xml:"ItemWeight" json:"ItemWeight"`
      SubscoreList *NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      
      }
    
    type NAPSubscoreListType struct {
        Subscore []NAPSubscoreType `xml:"Subscore" json:"Subscore"`
      
      }
    
    type NAPSubscoreType struct {
        SubscoreType *String `xml:"SubscoreType" json:"SubscoreType"`
      SubscoreValue *Float `xml:"SubscoreValue" json:"SubscoreValue"`
      
      }
    
    type DomainScoreType struct {
        RawScore *Float `xml:"RawScore" json:"RawScore"`
      ScaledScoreValue *Float `xml:"ScaledScoreValue" json:"ScaledScoreValue"`
      ScaledScoreLogitValue *Float `xml:"ScaledScoreLogitValue" json:"ScaledScoreLogitValue"`
      ScaledScoreStandardError *Float `xml:"ScaledScoreStandardError" json:"ScaledScoreStandardError"`
      ScaledScoreLogitStandardError *Float `xml:"ScaledScoreLogitStandardError" json:"ScaledScoreLogitStandardError"`
      StudentDomainBand *Int `xml:"StudentDomainBand" json:"StudentDomainBand"`
      StudentProficiency *String `xml:"StudentProficiency,omitempty" json:"StudentProficiency"`
      PlausibleScaledValueList *PlausibleScaledValueListType `xml:"PlausibleScaledValueList" json:"PlausibleScaledValueList"`
      
      }
    
    type NAPWritingRubricListType struct {
        NAPWritingRubric []NAPWritingRubricType `xml:"NAPWritingRubric" json:"NAPWritingRubric"`
      
      }
    
    type NAPWritingRubricType struct {
        RubricType *String `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type ScoreListType struct {
        Score []ScoreType `xml:"Score" json:"Score"`
      
      }
    
    type ScoreType struct {
        MaxScoreValue *Float `xml:"MaxScoreValue" json:"MaxScoreValue"`
      ScoreDescriptionList *ScoreDescriptionListType `xml:"ScoreDescriptionList" json:"ScoreDescriptionList"`
      
      }
    
    type ScoreDescriptionListType struct {
        ScoreDescription []ScoreDescriptionType `xml:"ScoreDescription" json:"ScoreDescription"`
      
      }
    
    type ScoreDescriptionType struct {
        ScoreValue *Float `xml:"ScoreValue" json:"ScoreValue"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type StimulusListType struct {
        Stimulus []StimulusType `xml:"Stimulus" json:"Stimulus"`
      
      }
    
    type StimulusType struct {
        StimulusLocalId *LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      TextGenre *String `xml:"TextGenre,omitempty" json:"TextGenre"`
      TextType *String `xml:"TextType,omitempty" json:"TextType"`
      WordCount *Int `xml:"WordCount,omitempty" json:"WordCount"`
      TextDescriptor *String `xml:"TextDescriptor,omitempty" json:"TextDescriptor"`
      Content *String `xml:"Content" json:"Content"`
      
      }
    
    type ContentDescriptionListType struct {
        ContentDescription []string `xml:"ContentDescription" json:"ContentDescription"`
      
      }
    
    type PNPCodeListType struct {
        PNPCode []AUCodeSetsPNPCodeType `xml:"PNPCode" json:"PNPCode"`
      
      }
    
    type AdjustmentContainerType struct {
        PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      BookletType *String `xml:"BookletType,omitempty" json:"BookletType"`
      
      }
    
    type TestDisruptionListType struct {
        TestDisruption []TestDisruptionType `xml:"TestDisruption" json:"TestDisruption"`
      
      }
    
    type TestDisruptionType struct {
        Event *String `xml:"Event" json:"Event"`
      
      }
    
    type CalendarSummaryListType struct {
        CalendarSummaryRefId []string `xml:"CalendarSummaryRefId" json:"CalendarSummaryRefId"`
      
      }
    
    type VisaSubClassType struct {
        Code *VisaSubClassCodeType `xml:"Code" json:"Code"`
      VisaExpiryDate *String `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      ATEExpiryDate *String `xml:"ATEExpiryDate,omitempty" json:"ATEExpiryDate"`
      ATEStartDate *String `xml:"ATEStartDate,omitempty" json:"ATEStartDate"`
      VisaStatisticalCode *String `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      
      }
    
    type VisaSubClassListType struct {
        VisaSubClass []VisaSubClassType `xml:"VisaSubClass" json:"VisaSubClass"`
      
      }
    
    type VisaSubClassCodeType string
    type LanguageBaseType struct {
        Code *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      LanguageType *AUCodeSetsLanguageTypeType `xml:"LanguageType,omitempty" json:"LanguageType"`
      Dialect *String `xml:"Dialect,omitempty" json:"Dialect"`
      
      }
    
    type ReligiousEventListType struct {
        ReligiousEvent []ReligiousEventType `xml:"ReligiousEvent" json:"ReligiousEvent"`
      
      }
    
    type ReligiousEventType struct {
        Type *String `xml:"Type" json:"Type"`
      Date *String `xml:"Date" json:"Date"`
      
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
          Type *String `xml:"Type,attr" json:"Type"`
      
      }
    
    type ScheduledActivityOverrideType struct {
          DateOfOverride *String `xml:"DateOfOverride,attr" json:"DateOfOverride"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type ActivityTimeType struct {
        CreationDate *String `xml:"CreationDate" json:"CreationDate"`
      Duration *ActivityTimeType_Duration
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate"`
      FinishDate *String `xml:"FinishDate,omitempty" json:"FinishDate"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate"`
      
      }
    
    type SchoolCourseInfoOverrideType struct {
        Override *String `xml:"Override,attr" json:"Override"`
      CourseCode *String `xml:"CourseCode,omitempty" json:"CourseCode"`
      StateCourseCode *String `xml:"StateCourseCode,omitempty" json:"StateCourseCode"`
      DistrictCourseCode *String `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      CourseTitle *String `xml:"CourseTitle,omitempty" json:"CourseTitle"`
      InstructionalLevel *String `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel"`
      CourseCredits *String `xml:"CourseCredits,omitempty" json:"CourseCredits"`
      
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
        ParentLegalGuardian *AUCodeSetsYesOrNoCategoryType `xml:"ParentLegalGuardian,omitempty" json:"ParentLegalGuardian"`
      PickupRights *AUCodeSetsYesOrNoCategoryType `xml:"PickupRights,omitempty" json:"PickupRights"`
      LivesWith *AUCodeSetsYesOrNoCategoryType `xml:"LivesWith,omitempty" json:"LivesWith"`
      AccessToRecords *AUCodeSetsYesOrNoCategoryType `xml:"AccessToRecords,omitempty" json:"AccessToRecords"`
      ReceivesAssessmentReport *AUCodeSetsYesOrNoCategoryType `xml:"ReceivesAssessmentReport,omitempty" json:"ReceivesAssessmentReport"`
      EmergencyContact *AUCodeSetsYesOrNoCategoryType `xml:"EmergencyContact,omitempty" json:"EmergencyContact"`
      HasCustody *AUCodeSetsYesOrNoCategoryType `xml:"HasCustody,omitempty" json:"HasCustody"`
      DisciplinaryContact *AUCodeSetsYesOrNoCategoryType `xml:"DisciplinaryContact,omitempty" json:"DisciplinaryContact"`
      AttendanceContact *AUCodeSetsYesOrNoCategoryType `xml:"AttendanceContact,omitempty" json:"AttendanceContact"`
      PrimaryCareProvider *AUCodeSetsYesOrNoCategoryType `xml:"PrimaryCareProvider,omitempty" json:"PrimaryCareProvider"`
      FeesBilling *AUCodeSetsYesOrNoCategoryType `xml:"FeesBilling,omitempty" json:"FeesBilling"`
      FeesAccess *AUCodeSetsYesOrNoCategoryType `xml:"FeesAccess,omitempty" json:"FeesAccess"`
      FamilyMail *AUCodeSetsYesOrNoCategoryType `xml:"FamilyMail,omitempty" json:"FamilyMail"`
      InterventionOrder *AUCodeSetsYesOrNoCategoryType `xml:"InterventionOrder,omitempty" json:"InterventionOrder"`
      
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
        Type *String `xml:"Type,attr" json:"Type"`
      UserId *String `xml:"UserId" json:"UserId"`
      
      }
    
    type AuditInfoType struct {
        CreationUser *CreationUserType `xml:"CreationUser" json:"CreationUser"`
      CreationDateTime *String `xml:"CreationDateTime" json:"CreationDateTime"`
      
      }
    
    type AttendanceInfoType struct {
        CountsTowardAttendance *String `xml:"CountsTowardAttendance" json:"CountsTowardAttendance"`
      AttendanceValue *Float `xml:"AttendanceValue" json:"AttendanceValue"`
      
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
        PromotionStatus *String `xml:"PromotionStatus,omitempty" json:"PromotionStatus"`
      
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
      FTE *Float `xml:"FTE,omitempty" json:"FTE"`
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
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup"`
      ClassCode *String `xml:"ClassCode,omitempty" json:"ClassCode"`
      MembershipType *AUCodeSetsSchoolEnrollmentTypeType `xml:"MembershipType,omitempty" json:"MembershipType"`
      FFPOS *AUCodeSetsFFPOSStatusCodeType `xml:"FFPOS,omitempty" json:"FFPOS"`
      ReportingSchoolId *LocalIdType `xml:"ReportingSchoolId,omitempty" json:"ReportingSchoolId"`
      OtherEnrollmentSchoolACARAId *LocalIdType `xml:"OtherEnrollmentSchoolACARAId,omitempty" json:"OtherEnrollmentSchoolACARAId"`
      OtherSchoolName *String `xml:"OtherSchoolName,omitempty" json:"OtherSchoolName"`
      DisabilityLevelOfAdjustment *String `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment"`
      DisabilityCategory *String `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory"`
      CensusAge *Int `xml:"CensusAge,omitempty" json:"CensusAge"`
      DistanceEducationStudent *AUCodeSetsYesOrNoCategoryType `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus"`
      
      }
    
    type StaffMostRecentContainerType struct {
        SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      LocalCampusId *LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId"`
      NAPLANClassList *NAPLANClassListType `xml:"NAPLANClassList,omitempty" json:"NAPLANClassList"`
      HomeGroup *String `xml:"HomeGroup,omitempty" json:"HomeGroup"`
      
      }
    
    type StaffActivityExtensionType struct {
        Code *AUCodeSetsStaffActivityType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type TotalEnrollmentsType struct {
        Girls *String `xml:"Girls,omitempty" json:"Girls"`
      Boys *String `xml:"Boys,omitempty" json:"Boys"`
      TotalStudents *String `xml:"TotalStudents,omitempty" json:"TotalStudents"`
      
      }
    
    type CampusContainerType struct {
        ParentSchoolId *String `xml:"ParentSchoolId,omitempty" json:"ParentSchoolId"`
      SchoolCampusId *String `xml:"SchoolCampusId" json:"SchoolCampusId"`
      CampusType *AUCodeSetsSchoolLevelType `xml:"CampusType,omitempty" json:"CampusType"`
      AdminStatus *AUCodeSetsYesOrNoCategoryType `xml:"AdminStatus" json:"AdminStatus"`
      
      }
    
    type HouseholdContactInfoListType struct {
        HouseholdContactInfo []HouseholdContactInfoType `xml:"HouseholdContactInfo" json:"HouseholdContactInfo"`
      
      }
    
    type HouseholdContactInfoType struct {
        PreferenceNumber *Int `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber"`
      HouseholdContactId *LocalIdType `xml:"HouseholdContactId,omitempty" json:"HouseholdContactId"`
      HouseholdSalutation *String `xml:"HouseholdSalutation,omitempty" json:"HouseholdSalutation"`
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
        AttendanceType *String `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      StartTime *String `xml:"StartTime" json:"StartTime"`
      EndTime *String `xml:"EndTime" json:"EndTime"`
      DurationValue *Float `xml:"DurationValue,omitempty" json:"DurationValue"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type PeriodAttendancesType struct {
        PeriodAttendance []PeriodAttendanceType `xml:"PeriodAttendance" json:"PeriodAttendance"`
      
      }
    
    type PeriodAttendanceType struct {
        AttendanceType *String `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      Date *String `xml:"Date" json:"Date"`
      SessionInfoRefId *String `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod *String `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      DayId *LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime *String `xml:"EndTime,omitempty" json:"EndTime"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type StaffSubjectListType struct {
        StaffSubject []StaffSubjectType `xml:"StaffSubject" json:"StaffSubject"`
      
      }
    
    type StaffSubjectType struct {
        PreferenceNumber *Int `xml:"PreferenceNumber" json:"PreferenceNumber"`
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
        StaffPersonalRefId *String `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime *String `xml:"FinishTime,omitempty" json:"FinishTime"`
      Credit *AUCodeSetsTeacherCoverCreditType `xml:"Credit,omitempty" json:"Credit"`
      Supervision *AUCodeSetsTeacherCoverSupervisionType `xml:"Supervision,omitempty" json:"Supervision"`
      Weighting *Float `xml:"Weighting,omitempty" json:"Weighting"`
      
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
        ItemNumber *String `xml:"ItemNumber,omitempty" json:"ItemNumber"`
      ItemDescription *String `xml:"ItemDescription" json:"ItemDescription"`
      LocalItemId *LocalIdType `xml:"LocalItemId,omitempty" json:"LocalItemId"`
      Quantity *String `xml:"Quantity,omitempty" json:"Quantity"`
      UnitCost *MonetaryAmountType `xml:"UnitCost,omitempty" json:"UnitCost"`
      TotalCost *MonetaryAmountType `xml:"TotalCost,omitempty" json:"TotalCost"`
      QuantityDelivered *String `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered"`
      CancelledOrder *AUCodeSetsYesOrNoCategoryType `xml:"CancelledOrder,omitempty" json:"CancelledOrder"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate"`
      ExpenseAccounts *ExpenseAccountsType `xml:"ExpenseAccounts,omitempty" json:"ExpenseAccounts"`
      
      }
    
    type ExpenseAccountsType struct {
        ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount" json:"ExpenseAccount"`
      
      }
    
    type ExpenseAccountType struct {
        AccountCode *String `xml:"AccountCode" json:"AccountCode"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      FinancialAccountRefId *String `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      
      }
    
    type SchoolProgramListType struct {
        Program []SchoolProgramType `xml:"Program" json:"Program"`
      
      }
    
    type SchoolProgramType struct {
        Category *String `xml:"Category,omitempty" json:"Category"`
      Type *String `xml:"Type" json:"Type"`
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
          ResourceType *String `xml:"ResourceType,attr" json:"ResourceType"`
      
        Value *String `xml:",chardata" json:"value"`
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
        TechnicalRequirement *String `xml:"TechnicalRequirement,omitempty" json:"TechnicalRequirement"`
      
      }
    
    type SoftwareRequirementListType struct {
        SoftwareRequirement []SoftwareRequirementType `xml:"SoftwareRequirement" json:"SoftwareRequirement"`
      
      }
    
    type SoftwareRequirementType struct {
        SoftwareTitle *String `xml:"SoftwareTitle" json:"SoftwareTitle"`
      Version *String `xml:"Version,omitempty" json:"Version"`
      Vendor *String `xml:"Vendor,omitempty" json:"Vendor"`
      OS *String `xml:"OS,omitempty" json:"OS"`
      
      }
    
    type HouseholdListType struct {
        Household []LocalIdType `xml:"Household" json:"Household"`
      
      }
    
    type StudentSubjectChoiceListType struct {
        StudentSubjectChoice []StudentSubjectChoiceType `xml:"StudentSubjectChoice" json:"StudentSubjectChoice"`
      
      }
    
    type StudentSubjectChoiceType struct {
        PreferenceNumber *Int `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber"`
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
        Name *String `xml:"Name" json:"Name"`
      Reference *String `xml:"Reference" json:"Reference"`
      Description *String `xml:"Description,omitempty" json:"Description"`
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
      Description *String `xml:"Description,omitempty" json:"Description"`
      Date *String `xml:"Date,omitempty" json:"Date"`
      Name *NameType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type ApprovalsType struct {
        Approval []ApprovalType `xml:"Approval" json:"Approval"`
      
      }
    
    type ApprovalType struct {
        Organization *String `xml:"Organization" json:"Organization"`
      Date *String `xml:"Date" json:"Date"`
      
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
        DebitFinancialAccountRefId *String `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId *String `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId"`
      DebitAccountCode *String `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode"`
      CreditAccountCode *String `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode"`
      GSTCodeOriginal *String `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal"`
      GSTCodeReplacement *String `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement"`
      LineAdjustmentAmount *MonetaryAmountType `xml:"LineAdjustmentAmount" json:"LineAdjustmentAmount"`
      
      }
    
    type PaymentReceiptLineListType struct {
        PaymentReceiptLine []PaymentReceiptLineType `xml:"PaymentReceiptLine" json:"PaymentReceiptLine"`
      
      }
    
    type PaymentReceiptLineType struct {
        InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      LocalPaymentReceiptLineId *LocalIdType `xml:"LocalPaymentReceiptLineId,omitempty" json:"LocalPaymentReceiptLineId"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      FinancialAccountRefId *String `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      AccountCode *String `xml:"AccountCode,omitempty" json:"AccountCode"`
      TransactionDescription *String `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      
      }
    
    type PasswordListType struct {
      Password []PasswordListType_Password `xml:"Password" json:"Password"`
      
      }
    
    type ExclusionRulesType struct {
        ExclusionRule []ExclusionRuleType `xml:"ExclusionRule" json:"ExclusionRule"`
      
      }
    
    type ExclusionRuleType struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
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
        TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      CellType *String `xml:"CellType,omitempty" json:"CellType"`
      
      }
    
    type TeacherListType struct {
        TeachingGroupTeacher []TeachingGroupTeacherType `xml:"TeachingGroupTeacher" json:"TeachingGroupTeacher"`
      
      }
    
    type TeachingGroupTeacherType struct {
        StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      Association *String `xml:"Association" json:"Association"`
      
      }
    
    type StudentListType struct {
        TeachingGroupStudent []TeachingGroupStudentType `xml:"TeachingGroupStudent" json:"TeachingGroupStudent"`
      
      }
    
    type TeachingGroupStudentType struct {
        StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId,omitempty" json:"StudentLocalId"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type TimeTableDayListType struct {
        TimeTableDay []TimeTableDayType `xml:"TimeTableDay" json:"TimeTableDay"`
      
      }
    
    type TimeTableDayType struct {
        DayId *LocalIdType `xml:"DayId" json:"DayId"`
      DayTitle *String `xml:"DayTitle" json:"DayTitle"`
      TimeTablePeriodList *TimeTablePeriodListType `xml:"TimeTablePeriodList" json:"TimeTablePeriodList"`
      
      }
    
    type TimeTablePeriodListType struct {
        TimeTablePeriod []TimeTablePeriodType `xml:"TimeTablePeriod" json:"TimeTablePeriod"`
      
      }
    
    type TimeTablePeriodType struct {
        PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      PeriodTitle *String `xml:"PeriodTitle" json:"PeriodTitle"`
      BellPeriod *String `xml:"BellPeriod,omitempty" json:"BellPeriod"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime *String `xml:"EndTime,omitempty" json:"EndTime"`
      RegularSchoolPeriod *String `xml:"RegularSchoolPeriod,omitempty" json:"RegularSchoolPeriod"`
      InstructionalMinutes *Int `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes"`
      UseInAttendanceCalculations *String `xml:"UseInAttendanceCalculations,omitempty" json:"UseInAttendanceCalculations"`
      
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
      Enrollment *String `xml:"Enrollment" json:"Enrollment"`
      
      }
    
    type SchoolFocusListType struct {
        SchoolFocus []AUCodeSetsSchoolFocusCodeType `xml:"SchoolFocus" json:"SchoolFocus"`
      
      }
    
    type AlertMessagesType struct {
        AlertMessage []AlertMessageType `xml:"AlertMessage" json:"AlertMessage"`
      
      }
    
    type AlertMessageType struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type MedicalAlertMessagesType struct {
        MedicalAlertMessage []MedicalAlertMessageType `xml:"MedicalAlertMessage" json:"MedicalAlertMessage"`
      
      }
    
    type MedicalAlertMessageType struct {
          Severity *String `xml:"Severity,attr" json:"Severity"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type OtherIdListType struct {
        OtherId []OtherIdType `xml:"OtherId" json:"OtherId"`
      
      }
    
    type OtherIdType struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type BaseNameType struct {
        Title *String `xml:"Title,omitempty" json:"Title"`
      FamilyName *String `xml:"FamilyName,omitempty" json:"FamilyName"`
      GivenName *String `xml:"GivenName,omitempty" json:"GivenName"`
      MiddleName *String `xml:"MiddleName,omitempty" json:"MiddleName"`
      FamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"FamilyNameFirst,omitempty" json:"FamilyNameFirst"`
      PreferredFamilyName *String `xml:"PreferredFamilyName,omitempty" json:"PreferredFamilyName"`
      PreferredFamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"PreferredFamilyNameFirst,omitempty" json:"PreferredFamilyNameFirst"`
      PreferredGivenName *String `xml:"PreferredGivenName,omitempty" json:"PreferredGivenName"`
      Suffix *String `xml:"Suffix,omitempty" json:"Suffix"`
      FullName *String `xml:"FullName,omitempty" json:"FullName"`
      
      }
    
    type NameOfRecordType struct {
        BaseNameType
          Type *String `xml:"Type,attr" json:"Type"`
      
      }
    
    type OtherNameType struct {
        BaseNameType
          Type *AUCodeSetsNameUsageTypeType `xml:"Type,attr" json:"Type"`
      
      }
    
    type PartialDateType string
    type LocalIdType string
    type LocationType struct {
        Type *String `xml:"Type,attr" json:"Type"`
      LocationName *String `xml:"LocationName,omitempty" json:"LocationName"`
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
      ContactTitle *String `xml:"ContactTitle,omitempty" json:"ContactTitle"`
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
      PositionTitle *String `xml:"PositionTitle,omitempty" json:"PositionTitle"`
      Role *String `xml:"Role,omitempty" json:"Role"`
      RegistrationDetails *String `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails"`
      Qualifications *String `xml:"Qualifications,omitempty" json:"Qualifications"`
      Address *AddressType `xml:"Address,omitempty" json:"Address"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      
      }
    
    type AddressStreetType struct {
        Line1 *String `xml:"Line1,omitempty" json:"Line1"`
      Line2 *String `xml:"Line2,omitempty" json:"Line2"`
      Line3 *String `xml:"Line3,omitempty" json:"Line3"`
      Complex *String `xml:"Complex,omitempty" json:"Complex"`
      StreetNumber *String `xml:"StreetNumber,omitempty" json:"StreetNumber"`
      StreetPrefix *String `xml:"StreetPrefix,omitempty" json:"StreetPrefix"`
      StreetName *String `xml:"StreetName,omitempty" json:"StreetName"`
      StreetType *String `xml:"StreetType,omitempty" json:"StreetType"`
      StreetSuffix *String `xml:"StreetSuffix,omitempty" json:"StreetSuffix"`
      ApartmentType *String `xml:"ApartmentType,omitempty" json:"ApartmentType"`
      ApartmentNumberPrefix *String `xml:"ApartmentNumberPrefix,omitempty" json:"ApartmentNumberPrefix"`
      ApartmentNumber *String `xml:"ApartmentNumber,omitempty" json:"ApartmentNumber"`
      ApartmentNumberSuffix *String `xml:"ApartmentNumberSuffix,omitempty" json:"ApartmentNumberSuffix"`
      
      }
    
    type AddressType struct {
        Type *AUCodeSetsAddressTypeType `xml:"Type,attr" json:"Type"`
      Role *AUCodeSetsAddressRoleType `xml:"Role,attr" json:"Role"`
      EffectiveFromDate *String `xml:"EffectiveFromDate,omitempty" json:"EffectiveFromDate"`
      EffectiveToDate *String `xml:"EffectiveToDate,omitempty" json:"EffectiveToDate"`
      Street *AddressStreetType `xml:"Street" json:"Street"`
      City *String `xml:"City" json:"City"`
      StateProvince *StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince"`
      Country *CountryType `xml:"Country,omitempty" json:"Country"`
      PostalCode *String `xml:"PostalCode" json:"PostalCode"`
      GridLocation *GridLocationType `xml:"GridLocation,omitempty" json:"GridLocation"`
      MapReference *MapReferenceType `xml:"MapReference,omitempty" json:"MapReference"`
      RadioContact *String `xml:"RadioContact,omitempty" json:"RadioContact"`
      Community *String `xml:"Community,omitempty" json:"Community"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      AddressGlobalUID *GUIDType `xml:"AddressGlobalUID,omitempty" json:"AddressGlobalUID"`
      StatisticalAreas *StatisticalAreasType `xml:"StatisticalAreas,omitempty" json:"StatisticalAreas"`
      
      }
    
    type MapReferenceType struct {
        Type *String `xml:"Type,attr" json:"Type"`
      XCoordinate *String `xml:"XCoordinate" json:"XCoordinate"`
      YCoordinate *String `xml:"YCoordinate" json:"YCoordinate"`
      MapNumber *String `xml:"MapNumber,omitempty" json:"MapNumber"`
      
      }
    
    type StatisticalAreasType struct {
        StatisticalArea []StatisticalAreaType `xml:"StatisticalArea" json:"StatisticalArea"`
      
      }
    
    type StatisticalAreaType struct {
          SpatialUnitType *String `xml:"SpatialUnitType,attr" json:"SpatialUnitType"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type AddressListType struct {
        Address []AddressType `xml:"Address" json:"Address"`
      
      }
    
    type EmailListType struct {
        Email []EmailType `xml:"Email" json:"Email"`
      
      }
    
    type EmailType struct {
          Type *AUCodeSetsEmailTypeType `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type PhoneNumberListType struct {
        PhoneNumber []PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type PhoneNumberType struct {
        Type *AUCodeSetsTelephoneNumberTypeType `xml:"Type,attr" json:"Type"`
      Number *String `xml:"Number" json:"Number"`
      Extension *String `xml:"Extension,omitempty" json:"Extension"`
      ListedStatus *AUCodeSetsYesOrNoCategoryType `xml:"ListedStatus,omitempty" json:"ListedStatus"`
      Preference *Int `xml:"Preference,omitempty" json:"Preference"`
      
      }
    
    type CountryType AUCodeSetsStandardAustralianClassificationOfCountriesSACCType
    type GridLocationType struct {
        Latitude *Float `xml:"Latitude" json:"Latitude"`
      Longitude *Float `xml:"Longitude" json:"Longitude"`
      
      }
    
    type OperationalStatusType AUCodeSetsOperationalStatusType
    type StateProvinceType string
    type SchoolYearType string
    type ElectronicIdListType struct {
        ElectronicId []ElectronicIdType `xml:"ElectronicId" json:"ElectronicId"`
      
      }
    
    type ElectronicIdType struct {
          Type *AUCodeSetsElectronicIdTypeType `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type OtherNamesType struct {
        Name []OtherNameType `xml:"Name" json:"Name"`
      
      }
    
    type DemographicsType struct {
        IndigenousStatus *AUCodeSetsIndigenousStatusType `xml:"IndigenousStatus,omitempty" json:"IndigenousStatus"`
      Sex *AUCodeSetsSexCodeType `xml:"Sex,omitempty" json:"Sex"`
      BirthDate *BirthDateType `xml:"BirthDate,omitempty" json:"BirthDate"`
      DateOfDeath *String `xml:"DateOfDeath,omitempty" json:"DateOfDeath"`
      BirthDateVerification *AUCodeSetsBirthdateVerificationType `xml:"BirthDateVerification,omitempty" json:"BirthDateVerification"`
      PlaceOfBirth *String `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth"`
      StateOfBirth *StateProvinceType `xml:"StateOfBirth,omitempty" json:"StateOfBirth"`
      CountryOfBirth *CountryType `xml:"CountryOfBirth,omitempty" json:"CountryOfBirth"`
      CountriesOfCitizenship *CountryListType `xml:"CountriesOfCitizenship,omitempty" json:"CountriesOfCitizenship"`
      CountriesOfResidency *CountryList2Type `xml:"CountriesOfResidency,omitempty" json:"CountriesOfResidency"`
      CountryArrivalDate *String `xml:"CountryArrivalDate,omitempty" json:"CountryArrivalDate"`
      AustralianCitizenshipStatus *AUCodeSetsAustralianCitizenshipStatusType `xml:"AustralianCitizenshipStatus,omitempty" json:"AustralianCitizenshipStatus"`
      EnglishProficiency *EnglishProficiencyType `xml:"EnglishProficiency,omitempty" json:"EnglishProficiency"`
      LanguageList *LanguageListType `xml:"LanguageList,omitempty" json:"LanguageList"`
      DwellingArrangement *DwellingArrangementType `xml:"DwellingArrangement,omitempty" json:"DwellingArrangement"`
      Religion *ReligionType `xml:"Religion,omitempty" json:"Religion"`
      ReligiousEventList *ReligiousEventListType `xml:"ReligiousEventList,omitempty" json:"ReligiousEventList"`
      ReligiousRegion *String `xml:"ReligiousRegion,omitempty" json:"ReligiousRegion"`
      PermanentResident *AUCodeSetsPermanentResidentStatusType `xml:"PermanentResident,omitempty" json:"PermanentResident"`
      VisaSubClass *VisaSubClassCodeType `xml:"VisaSubClass,omitempty" json:"VisaSubClass"`
      VisaStatisticalCode *String `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      VisaExpiryDate *String `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      VisaSubClassList *VisaSubClassListType `xml:"VisaSubClassList,omitempty" json:"VisaSubClassList"`
      LBOTE *AUCodeSetsYesOrNoCategoryType `xml:"LBOTE,omitempty" json:"LBOTE"`
      InterpreterRequired *AUCodeSetsYesOrNoCategoryType `xml:"InterpreterRequired,omitempty" json:"InterpreterRequired"`
      ImmunisationCertificateStatus *AUCodeSetsImmunisationCertificateStatusType `xml:"ImmunisationCertificateStatus,omitempty" json:"ImmunisationCertificateStatus"`
      CulturalBackground *AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType `xml:"CulturalBackground,omitempty" json:"CulturalBackground"`
      MaritalStatus *AUCodeSetsMaritalStatusAIHWType `xml:"MaritalStatus,omitempty" json:"MaritalStatus"`
      MedicareNumber *String `xml:"MedicareNumber,omitempty" json:"MedicareNumber"`
      
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
        Type *String `xml:"Type" json:"Type"`
      Code *String `xml:"Code" json:"Code"`
      Name *String `xml:"Name" json:"Name"`
      Value *String `xml:"Value" json:"Value"`
      StartDateTime *String `xml:"StartDateTime,omitempty" json:"StartDateTime"`
      EndDateTime *String `xml:"EndDateTime,omitempty" json:"EndDateTime"`
      SpanGaps *TimeElementType_SpanGaps
      IsCurrent *Bool `xml:"IsCurrent" json:"IsCurrent"`
      
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
        Code *String `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type SubjectAreaListType struct {
        SubjectArea []SubjectAreaType `xml:"SubjectArea" json:"SubjectArea"`
      
      }
    
    type ACStrandAreaListType struct {
        ACStrandSubjectArea []ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      
      }
    
    type SubjectAreaType struct {
        Code *String `xml:"Code" json:"Code"`
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
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_TextData struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_BinaryData struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentPackageType_Reference struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      Description *String `xml:"Description,attr" json:"Description"`
       URL *String `xml:"URL" json:"URL"`
}
type AbstractContentElementType_XMLData struct {
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_TextData struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_BinaryData struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_Reference struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      Description *String `xml:"Description,attr" json:"Description"`
       URL *String `xml:"URL" json:"URL"`
}
type PersonInvolvementType_PersonRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type ActivityTimeType_Duration struct {
      Units *String `xml:"Units,attr" json:"Units"`
      Value *Int `xml:",chardata" json:"value"`
}
type SourceObjectsType_SourceObject struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type IdentityAssertionsType_IdentityAssertion struct {
      SchemaName *String `xml:"SchemaName,attr" json:"SchemaName"`
      Value *String `xml:",chardata" json:"value"`
}
type AssociatedObjectsType_AssociatedObject struct {
      SIF_RefObject *ObjectNameType `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type PasswordListType_Password struct {
      Algorithm *String `xml:"Algorithm,attr" json:"Algorithm"`
      KeyName *String `xml:"KeyName,attr" json:"KeyName"`
      Value *String `xml:",chardata" json:"value"`
}
type LocationType_LocationRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type TimeElementType_SpanGaps struct {
      SpanGap []TimeElementType_SpanGap `xml:"SpanGap" json:"SpanGap"`
}
type LifeCycleType_Created struct {
       DateTime *String `xml:"DateTime" json:"DateTime"`
      Creators *LifeCycleType_Creators
}
type LifeCycleType_ModificationHistory struct {
      Modified []LifeCycleType_Modified `xml:"Modified" json:"Modified"`
}
type LifeCycleType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type OtherCodeListType_OtherCode struct {
      Codeset *String `xml:"Codeset,attr" json:"Codeset"`
      Value *String `xml:",chardata" json:"value"`
}
type SIF_ExtendedElementsType_SIF_ExtendedElement struct {
      Name *String `xml:"Name,attr" json:"Name"`
      Type *String `xml:"Type,attr" json:"Type"`
      SIF_Action *String `xml:"SIF_Action,attr" json:"SIF_Action"`
      Value *ExtendedContentType `xml:",chardata" json:"value"`
}
type SIF_MetadataType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type TimeElementType_SpanGap struct {
       Type *String `xml:"Type" json:"Type"`
       Code *String `xml:"Code" json:"Code"`
       Name *String `xml:"Name" json:"Name"`
       Value *String `xml:"Value" json:"Value"`
       StartDateTime *String `xml:"StartDateTime,omitempty" json:"StartDateTime"`
       EndDateTime *String `xml:"EndDateTime,omitempty" json:"EndDateTime"`
}
type LifeCycleType_Creators struct {
      Creator []LifeCycleType_Creator `xml:"Creator" json:"Creator"`
}
type LifeCycleType_Modified struct {
       By *String `xml:"By" json:"By"`
       DateTime *String `xml:"DateTime" json:"DateTime"`
       Description *String `xml:"Description,omitempty" json:"Description"`
}
type LifeCycleType_Creator struct {
       Name *String `xml:"Name" json:"Name"`
       ID *String `xml:"ID" json:"ID"`
}
type String string
type Int int
type Float float64
type Bool bool
