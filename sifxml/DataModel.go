package sifxml


    type ReportPackageType AbstractContentPackageType
    type AbstractContentPackageType struct {
        RefId RefIdType `xml:"RefId,attr"`
      XMLData AbstractContentPackageType_XMLData `xml:"XMLData"`
      TextData AbstractContentPackageType_TextData `xml:"TextData"`
      BinaryData AbstractContentPackageType_BinaryData `xml:"BinaryData"`
      Reference AbstractContentPackageType_Reference `xml:"Reference"`
      
      }
    
    type AbstractContentElementType struct {
      XMLData AbstractContentElementType_XMLData `xml:"XMLData"`
      TextData AbstractContentElementType_TextData `xml:"TextData"`
      BinaryData AbstractContentElementType_BinaryData `xml:"BinaryData"`
      Reference AbstractContentElementType_Reference `xml:"Reference"`
      
      }
    
    type MonetaryAmountType struct {
          Currency string `xml:"Currency,attr"`
      
        Value string `xml:",chardata"`
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
        Transaction []LibraryTransactionType `xml:"Transaction"`
      
      }
    
    type LibraryTransactionType struct {
        ItemInfo LibraryItemInfoType `xml:"ItemInfo"`
      CheckoutInfo CheckoutInfoType `xml:"CheckoutInfo"`
      FineInfoList FineInfoListType `xml:"FineInfoList"`
      HoldInfoList HoldInfoListType `xml:"HoldInfoList"`
      
      }
    
    type LibraryItemInfoType struct {
        Type string `xml:"Type,attr"`
      TypeCodeset string `xml:"TypeCodeset,attr"`
      Title string `xml:"Title"`
      Author string `xml:"Author"`
      ElectronicId ElectronicIdType `xml:"ElectronicId"`
      CallNumber string `xml:"CallNumber"`
      Price MonetaryAmountType `xml:"Price"`
      
      }
    
    type CheckoutInfoType struct {
        ReturnBy string `xml:"ReturnBy"`
      
      }
    
    type FineInfoListType struct {
        FineInfo []FineInfoType `xml:"FineInfo"`
      
      }
    
    type FineInfoType struct {
        Type string `xml:"Type,attr"`
      TypeCodeset string `xml:"TypeCodeset,attr"`
      Assessed string `xml:"Assessed"`
      Description string `xml:"Description"`
      Amount MonetaryAmountType `xml:"Amount"`
      Reference string `xml:"Reference"`
      
      }
    
    type HoldInfoListType struct {
        HoldInfo []HoldInfoType `xml:"HoldInfo"`
      
      }
    
    type HoldInfoType struct {
        Type string `xml:"Type,attr"`
      TypeCodeset string `xml:"TypeCodeset,attr"`
      DatePlaced string `xml:"DatePlaced"`
      DateNeeded string `xml:"DateNeeded"`
      MadeAvailable string `xml:"MadeAvailable"`
      Expires string `xml:"Expires"`
      
      }
    
    type LibraryMessageListType struct {
        Message []LibraryMessageType `xml:"Message"`
      
      }
    
    type LibraryMessageType struct {
        Priority string `xml:"Priority,attr"`
      PriorityCodeset string `xml:"PriorityCodeset,attr"`
      Sent string `xml:"Sent"`
      Text string `xml:"Text"`
      
      }
    
    type StudentAttendanceCollectionReportingListType struct {
        StudentAttendanceCollectionReporting []StudentAttendanceCollectionReportingType `xml:"StudentAttendanceCollectionReporting"`
      
      }
    
    type StudentAttendanceCollectionReportingType struct {
        EntityLevel string `xml:"EntityLevel"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      LocalId LocalIdType `xml:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId"`
      CommonwealthId string `xml:"CommonwealthId"`
      ACARAId string `xml:"ACARAId"`
      EntityName string `xml:"EntityName"`
      EntityContact EntityContactInfoType `xml:"EntityContact"`
      StatsCohortYearLevelList StatsCohortYearLevelListType `xml:"StatsCohortYearLevelList"`
      
      }
    
    type StatsCohortYearLevelListType struct {
        StatsCohortYearLevel []StatsCohortYearLevelType `xml:"StatsCohortYearLevel"`
      
      }
    
    type StatsCohortYearLevelType struct {
        CohortYearLevel YearLevelType `xml:"CohortYearLevel"`
      StatsCohortList StatsCohortListType `xml:"StatsCohortList"`
      
      }
    
    type StatsCohortListType struct {
        StatsCohort []StatsCohortType `xml:"StatsCohort"`
      
      }
    
    type StatsCohortType struct {
        StatsCohortId LocalIdType `xml:"StatsCohortId"`
      StatsIndigenousStudentType string `xml:"StatsIndigenousStudentType"`
      CohortGender string `xml:"CohortGender"`
      DaysInReferencePeriod string `xml:"DaysInReferencePeriod"`
      PossibleSchoolDays string `xml:"PossibleSchoolDays"`
      AttendanceDays string `xml:"AttendanceDays"`
      AttendanceLess90Percent string `xml:"AttendanceLess90Percent"`
      AttendanceGTE90Percent string `xml:"AttendanceGTE90Percent"`
      PossibleSchoolDaysGT90PercentAttendance string `xml:"PossibleSchoolDaysGT90PercentAttendance"`
      
      }
    
    type AddressCollectionReportingListType struct {
        AddressCollectionReporting []AddressCollectionReportingType `xml:"AddressCollectionReporting"`
      
      }
    
    type AddressCollectionReportingType struct {
        EntityLevel string `xml:"EntityLevel"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      LocalId LocalIdType `xml:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId"`
      CommonwealthId string `xml:"CommonwealthId"`
      ACARAId string `xml:"ACARAId"`
      EntityName string `xml:"EntityName"`
      EntityContact EntityContactInfoType `xml:"EntityContact"`
      AGContextualQuestionList AGContextualQuestionListType `xml:"AGContextualQuestionList"`
      AddressCollectionStudentList AddressCollectionStudentListType `xml:"AddressCollectionStudentList"`
      
      }
    
    type AddressCollectionStudentListType struct {
        AddressCollectionStudent []AddressCollectionStudentType `xml:"AddressCollectionStudent"`
      
      }
    
    type AddressCollectionStudentType struct {
        LocalId LocalIdType `xml:"LocalId"`
      EducationLevel string `xml:"EducationLevel"`
      BoardingStatus string `xml:"BoardingStatus"`
      ReportingParent2 string `xml:"ReportingParent2"`
      StudentAddress AddressType `xml:"StudentAddress"`
      Parent1 AGParentType `xml:"Parent1"`
      Parent2 AGParentType `xml:"Parent2"`
      
      }
    
    type AGParentType struct {
        ParentName NameOfRecordType `xml:"ParentName"`
      AddressSameAsStudent string `xml:"AddressSameAsStudent"`
      ParentAddress AddressType `xml:"ParentAddress"`
      
      }
    
    type AGRoundListType struct {
        AGRound []AGRoundType `xml:"AGRound"`
      
      }
    
    type AGRoundType struct {
        RoundCode string `xml:"RoundCode"`
      RoundName string `xml:"RoundName"`
      StartDate string `xml:"StartDate"`
      DueDate string `xml:"DueDate"`
      EndDate string `xml:"EndDate"`
      
      }
    
    type AGContextualQuestionListType struct {
        AGContextualQuestion []AGContextualQuestionType `xml:"AGContextualQuestion"`
      
      }
    
    type AGContextualQuestionType struct {
        AGContextCode string `xml:"AGContextCode"`
      AGAnswer string `xml:"AGAnswer"`
      
      }
    
    type CensusReportingListType struct {
        CensusReporting []CensusReportingType `xml:"CensusReporting"`
      
      }
    
    type CensusReportingType struct {
        EntityLevel string `xml:"EntityLevel"`
      CommonwealthId string `xml:"CommonwealthId"`
      EntityName string `xml:"EntityName"`
      EntityContact EntityContactInfoType `xml:"EntityContact"`
      CensusStaffList CensusStaffListType `xml:"CensusStaffList"`
      CensusStudentList CensusStudentListType `xml:"CensusStudentList"`
      
      }
    
    type CensusStaffListType struct {
        CensusStaff []CensusStaffType `xml:"CensusStaff"`
      
      }
    
    type CensusStaffType struct {
        StaffCohortId LocalIdType `xml:"StaffCohortId"`
      StaffActivity StaffActivityExtensionType `xml:"StaffActivity"`
      CohortGender string `xml:"CohortGender"`
      CohortIndigenousType string `xml:"CohortIndigenousType"`
      PrimaryFTE string `xml:"PrimaryFTE"`
      SecondaryFTE string `xml:"SecondaryFTE"`
      JobFTE string `xml:"JobFTE"`
      Headcount string `xml:"Headcount"`
      
      }
    
    type StaffAssignmentMostRecentContainerType struct {
        PrimaryFTE string `xml:"PrimaryFTE"`
      SecondaryFTE string `xml:"SecondaryFTE"`
      
      }
    
    type CensusStudentListType struct {
        CensusStudent []CensusStudentType `xml:"CensusStudent"`
      
      }
    
    type CensusStudentType struct {
        StudentCohortId LocalIdType `xml:"StudentCohortId"`
      YearLevel YearLevelType `xml:"YearLevel"`
      CensusAge string `xml:"CensusAge"`
      CohortGender string `xml:"CohortGender"`
      CohortIndigenousType string `xml:"CohortIndigenousType"`
      EducationMode string `xml:"EducationMode"`
      StudentOnVisa string `xml:"StudentOnVisa"`
      OverseasStudent string `xml:"OverseasStudent"`
      DisabilityLevelOfAdjustment string `xml:"DisabilityLevelOfAdjustment"`
      DisabilityCategory string `xml:"DisabilityCategory"`
      FTE string `xml:"FTE"`
      Headcount string `xml:"Headcount"`
      
      }
    
    type AGReportingObjectResponseListType struct {
        AGReportingObjectResponse []AGReportingObjectResponseType `xml:"AGReportingObjectResponse"`
      
      }
    
    type AGReportingObjectResponseType struct {
        SubmittedRefId string `xml:"SubmittedRefId"`
      SIFRefId string `xml:"SIFRefId"`
      HTTPStatusCode string `xml:"HTTPStatusCode"`
      ErrorText string `xml:"ErrorText"`
      CommonwealthId string `xml:"CommonwealthId"`
      EntityName string `xml:"EntityName"`
      AGSubmissionStatusCode string `xml:"AGSubmissionStatusCode"`
      AGRuleList AGRuleListType `xml:"AGRuleList"`
      
      }
    
    type FQReportingListType struct {
        FQReporting []FQReportingType `xml:"FQReporting"`
      
      }
    
    type FQReportingType struct {
        EntityLevel string `xml:"EntityLevel"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      LocalId LocalIdType `xml:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId"`
      CommonwealthId string `xml:"CommonwealthId"`
      ACARAId string `xml:"ACARAId"`
      EntityName string `xml:"EntityName"`
      EntityContact EntityContactInfoType `xml:"EntityContact"`
      FQContextualQuestionList FQContextualQuestionListType `xml:"FQContextualQuestionList"`
      FQItemList FQItemListType `xml:"FQItemList"`
      AGRuleList AGRuleListType `xml:"AGRuleList"`
      
      }
    
    type FQContextualQuestionListType struct {
        FQContextualQuestion []FQContextualQuestionType `xml:"FQContextualQuestion"`
      
      }
    
    type FQContextualQuestionType struct {
        FQContext string `xml:"FQContext"`
      FQAnswer string `xml:"FQAnswer"`
      
      }
    
    type FQItemListType struct {
        FQItem []FQItemType `xml:"FQItem"`
      
      }
    
    type FQItemType struct {
        FQItemCode string `xml:"FQItemCode"`
      TuitionAmount string `xml:"TuitionAmount"`
      BoardingAmount string `xml:"BoardingAmount"`
      SystemAmount string `xml:"SystemAmount"`
      DioceseAmount string `xml:"DioceseAmount"`
      FQComments string `xml:"FQComments"`
      
      }
    
    type AGRuleListType struct {
        AGRule []AGRuleType `xml:"AGRule"`
      
      }
    
    type AGRuleType struct {
        AGRuleCode string `xml:"AGRuleCode"`
      AGRuleComment string `xml:"AGRuleComment"`
      AGRuleResponse string `xml:"AGRuleResponse"`
      AGRuleStatus string `xml:"AGRuleStatus"`
      
      }
    
    type SoftwareVendorInfoContainerType struct {
        SoftwareProduct string `xml:"SoftwareProduct"`
      SoftwareVersion string `xml:"SoftwareVersion"`
      
      }
    
    type TimeTableScheduleType struct {
        SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId"`
      Title string `xml:"Title"`
      DaysPerCycle string `xml:"DaysPerCycle"`
      PeriodsPerDay string `xml:"PeriodsPerDay"`
      TeachingPeriodsPerDay string `xml:"TeachingPeriodsPerDay"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      SchoolName string `xml:"SchoolName"`
      TimeTableCreationDate string `xml:"TimeTableCreationDate"`
      StartDate string `xml:"StartDate"`
      EndDate string `xml:"EndDate"`
      TimeTableDayList TimeTableDayListType `xml:"TimeTableDayList"`
      
      }
    
    type TimeTableScheduleCellListType struct {
        TimeTableScheduleCell []TimeTableScheduleCellType `xml:"TimeTableScheduleCell"`
      
      }
    
    type TimeTableScheduleCellType struct {
        TimeTableScheduleCellLocalId LocalIdType `xml:"TimeTableScheduleCellLocalId"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId"`
      TeachingGroupGUID string `xml:"TeachingGroupGUID"`
      RoomInfoRefId string `xml:"RoomInfoRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      TimeTableLocalId LocalIdType `xml:"TimeTableLocalId"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId"`
      TeachingGroupLocalId LocalIdType `xml:"TeachingGroupLocalId"`
      RoomNumber HomeroomNumberType `xml:"RoomNumber"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      DayId LocalIdType `xml:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId"`
      CellType string `xml:"CellType"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList"`
      RoomList RoomListType `xml:"RoomList"`
      
      }
    
    type TeachingGroupScheduleListType struct {
        TeachingGroupSchedule []TeachingGroupScheduleType `xml:"TeachingGroupSchedule"`
      
      }
    
    type TeachingGroupScheduleType struct {
        EditorGUID RefIdType `xml:"EditorGUID"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId"`
      ShortName string `xml:"ShortName"`
      LongName string `xml:"LongName"`
      GroupType string `xml:"GroupType"`
      Set string `xml:"Set"`
      Block string `xml:"Block"`
      CurriculumLevel string `xml:"CurriculumLevel"`
      SchoolInfoRefId RefIdType `xml:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      SchoolCourseInfoRefId RefIdType `xml:"SchoolCourseInfoRefId"`
      SchoolCourseLocalId LocalIdType `xml:"SchoolCourseLocalId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId"`
      TimeTableSubjectLocalId LocalIdType `xml:"TimeTableSubjectLocalId"`
      Semester string `xml:"Semester"`
      StudentList StudentListType `xml:"StudentList"`
      TeacherList TeacherListType `xml:"TeacherList"`
      MinClassSize string `xml:"MinClassSize"`
      MaxClassSize string `xml:"MaxClassSize"`
      TeachingGroupPeriodList TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList"`
      
      }
    
    type LocalCodeListType struct {
        LocalCode []LocalCodeType `xml:"LocalCode"`
      
      }
    
    type LocalCodeType struct {
        LocalisedCode string `xml:"LocalisedCode"`
      Description string `xml:"Description"`
      Element string `xml:"Element"`
      ListIndex string `xml:"ListIndex"`
      
      }
    
    type StudentGroupListType struct {
        StudentGroup []StudentGroupType `xml:"StudentGroup"`
      
      }
    
    type StudentGroupType struct {
        GroupCategory string `xml:"GroupCategory"`
      GroupLocalId LocalIdType `xml:"GroupLocalId"`
      GroupDescription string `xml:"GroupDescription"`
      
      }
    
    type PublishingPermissionListType struct {
        PublishingPermission []PublishingPermissionType `xml:"PublishingPermission"`
      
      }
    
    type PublishingPermissionType struct {
        PermissionCategory string `xml:"PermissionCategory"`
      PermissionValue string `xml:"PermissionValue"`
      
      }
    
    type EntityContactInfoType struct {
        Name NameType `xml:"Name"`
      PositionTitle string `xml:"PositionTitle"`
      Role string `xml:"Role"`
      RegistrationDetails string `xml:"RegistrationDetails"`
      Qualifications string `xml:"Qualifications"`
      Address AddressType `xml:"Address"`
      Email EmailType `xml:"Email"`
      PhoneNumber PhoneNumberType `xml:"PhoneNumber"`
      
      }
    
    type CopyRightContainerType struct {
        Date string `xml:"Date"`
      Holder string `xml:"Holder"`
      
      }
    
    type StandardsSettingBodyType struct {
        Country CountryType `xml:"Country"`
      StateProvince StateProvinceType `xml:"StateProvince"`
      SettingBodyName string `xml:"SettingBodyName"`
      
      }
    
    type StandardHierarchyLevelType struct {
        Number string `xml:"Number"`
      Description string `xml:"Description"`
      
      }
    
    type StandardIdentifierType struct {
        YearCreated string `xml:"YearCreated"`
      ACStrandSubjectArea ACStrandSubjectAreaType `xml:"ACStrandSubjectArea"`
      StandardNumber string `xml:"StandardNumber"`
      YearLevels YearLevelsType `xml:"YearLevels"`
      Benchmark string `xml:"Benchmark"`
      YearLevel YearLevelType `xml:"YearLevel"`
      IndicatorNumber string `xml:"IndicatorNumber"`
      AlternateIdentificationCodes AlternateIdentificationCodeListType `xml:"AlternateIdentificationCodes"`
      Organization string `xml:"Organization"`
      
      }
    
    type AlternateIdentificationCodeListType struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode"`
      
      }
    
    type RelatedLearningStandardItemRefIdListType struct {
        LearningStandardItemRefId []RelatedLearningStandardItemRefIdType `xml:"LearningStandardItemRefId"`
      
      }
    
    type RelatedLearningStandardItemRefIdType struct {
          RelationshipType string `xml:"RelationshipType,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type ValidLetterMarkListType struct {
        ValidLetterMark []ValidLetterMarkType `xml:"ValidLetterMark"`
      
      }
    
    type ValidLetterMarkType struct {
        Code string `xml:"Code"`
      NumericEquivalent string `xml:"NumericEquivalent"`
      Description string `xml:"Description"`
      
      }
    
    type StudentGradeMarkersListType struct {
        Marker []MarkerType `xml:"Marker"`
      
      }
    
    type MarkerType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      Role string `xml:"Role"`
      
      }
    
    type GradingScoreListType struct {
        GradingAssignmentScore []AssignmentScoreType `xml:"GradingAssignmentScore"`
      
      }
    
    type AssignmentScoreType struct {
        GradingAssignmentScoreRefId string `xml:"GradingAssignmentScoreRefId"`
      Weight string `xml:"Weight"`
      
      }
    
    type GradeType struct {
        Percentage string `xml:"Percentage"`
      Numeric string `xml:"Numeric"`
      Letter string `xml:"Letter"`
      Narrative string `xml:"Narrative"`
      MarkInfoRefId string `xml:"MarkInfoRefId"`
      
      }
    
    type LearningStandardListType struct {
        LearningStandard []LearningStandardType `xml:"LearningStandard"`
      
      }
    
    type LearningStandardType struct {
        LearningStandardItemRefId string `xml:"LearningStandardItemRefId"`
      LearningStandardURL string `xml:"LearningStandardURL"`
      LearningStandardLocalId LocalIdType `xml:"LearningStandardLocalId"`
      
      }
    
    type AssignmentListType struct {
        GradingAssignmentRefId []string `xml:"GradingAssignmentRefId"`
      
      }
    
    type GenericRubricType struct {
        RubricType string `xml:"RubricType"`
      ScoreList ScoreListType `xml:"ScoreList"`
      Descriptor string `xml:"Descriptor"`
      
      }
    
    type SymptomListType struct {
        Symptom []string `xml:"Symptom"`
      
      }
    
    type MedicationListType struct {
        Medication []MedicationType `xml:"Medication"`
      
      }
    
    type MedicationType struct {
        MedicationName string `xml:"MedicationName"`
      Dosage string `xml:"Dosage"`
      Frequency string `xml:"Frequency"`
      AdministrationInformation string `xml:"AdministrationInformation"`
      Method string `xml:"Method"`
      
      }
    
    type WellbeingEventCategoryListType struct {
        WellbeingEventCategory []WellbeingEventCategoryType `xml:"WellbeingEventCategory"`
      
      }
    
    type WellbeingEventCategoryType struct {
        EventCategory string `xml:"EventCategory"`
      WellbeingEventSubCategoryList WellbeingEventSubCategoryListType `xml:"WellbeingEventSubCategoryList"`
      
      }
    
    type WellbeingEventSubCategoryListType struct {
        WellbeingEventSubCategory []string `xml:"WellbeingEventSubCategory"`
      
      }
    
    type WellbeingEventLocationDetailsType struct {
        EventLocation string `xml:"EventLocation"`
      Class string `xml:"Class"`
      FurtherLocationNotes string `xml:"FurtherLocationNotes"`
      
      }
    
    type FollowUpActionListType struct {
        FollowUpAction []FollowUpActionType `xml:"FollowUpAction"`
      
      }
    
    type FollowUpActionType struct {
        WellbeingResponseRefId string `xml:"WellbeingResponseRefId"`
      FollowUpDetails string `xml:"FollowUpDetails"`
      FollowUpActionCategory string `xml:"FollowUpActionCategory"`
      
      }
    
    type PersonInvolvementListType struct {
        PersonInvolvement []PersonInvolvementType `xml:"PersonInvolvement"`
      
      }
    
    type PersonInvolvementType struct {
      PersonRefId PersonInvolvementType_PersonRefId `xml:"PersonRefId"`
      ShortName string `xml:"ShortName"`
      HowInvolved string `xml:"HowInvolved"`
      
      }
    
    type WithdrawalTimeListType struct {
        Withdrawal []WithdrawalType `xml:"Withdrawal"`
      
      }
    
    type WithdrawalType struct {
        WithdrawalDate string `xml:"WithdrawalDate"`
      WithdrawalStartTime string `xml:"WithdrawalStartTime"`
      WithdrawalEndTime string `xml:"WithdrawalEndTime"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId"`
      
      }
    
    type SuspensionContainerType struct {
        SuspensionCategory string `xml:"SuspensionCategory"`
      SuspensionNotes string `xml:"SuspensionNotes"`
      WithdrawalTimeList WithdrawalTimeListType `xml:"WithdrawalTimeList"`
      Duration string `xml:"Duration"`
      AdvisementDate string `xml:"AdvisementDate"`
      ResolutionMeetingTime string `xml:"ResolutionMeetingTime"`
      ResolutionNotes string `xml:"ResolutionNotes"`
      EarlyReturnDate string `xml:"EarlyReturnDate"`
      Status string `xml:"Status"`
      
      }
    
    type DetentionContainerType struct {
        DetentionCategory string `xml:"DetentionCategory"`
      DetentionDate string `xml:"DetentionDate"`
      DetentionLocation string `xml:"DetentionLocation"`
      DetentionNotes string `xml:"DetentionNotes"`
      Status string `xml:"Status"`
      
      }
    
    type PlanRequiredContainerType struct {
        PlanRequiredList PlanRequiredListType `xml:"PlanRequiredList"`
      Status string `xml:"Status"`
      
      }
    
    type PlanRequiredListType struct {
        Plan []WellbeingPlanType `xml:"Plan"`
      
      }
    
    type WellbeingPlanType struct {
        PersonalisedPlanRefId string `xml:"PersonalisedPlanRefId"`
      PlanNotes string `xml:"PlanNotes"`
      
      }
    
    type AwardContainerType struct {
        AwardDate string `xml:"AwardDate"`
      AwardType string `xml:"AwardType"`
      AwardDescription string `xml:"AwardDescription"`
      AwardNotes string `xml:"AwardNotes"`
      Status string `xml:"Status"`
      
      }
    
    type OtherWellbeingResponseContainerType struct {
        OtherResponseDate string `xml:"OtherResponseDate"`
      OtherResponseType string `xml:"OtherResponseType"`
      OtherResponseDescription string `xml:"OtherResponseDescription"`
      OtherResponseNotes string `xml:"OtherResponseNotes"`
      Status string `xml:"Status"`
      
      }
    
    type WellbeingDocumentListType struct {
        Document []WellbeingDocumentType `xml:"Document"`
      
      }
    
    type WellbeingDocumentType struct {
        Location string `xml:"Location"`
      Sensitivity string `xml:"Sensitivity"`
      URL string `xml:"URL"`
      DocumentType string `xml:"DocumentType"`
      DocumentReviewDate string `xml:"DocumentReviewDate"`
      DocumentDescription string `xml:"DocumentDescription"`
      
      }
    
    type NAPTestItemContentType struct {
        NAPTestItemLocalId LocalIdType `xml:"NAPTestItemLocalId"`
      ItemName string `xml:"ItemName"`
      ItemType string `xml:"ItemType"`
      Subdomain string `xml:"Subdomain"`
      WritingGenre string `xml:"WritingGenre"`
      ItemDescriptor string `xml:"ItemDescriptor"`
      ReleasedStatus string `xml:"ReleasedStatus"`
      MarkingType string `xml:"MarkingType"`
      MultipleChoiceOptionCount string `xml:"MultipleChoiceOptionCount"`
      CorrectAnswer string `xml:"CorrectAnswer"`
      MaximumScore string `xml:"MaximumScore"`
      ItemDifficulty string `xml:"ItemDifficulty"`
      ItemDifficultyLogit5 string `xml:"ItemDifficultyLogit5"`
      ItemDifficultyLogit62 string `xml:"ItemDifficultyLogit62"`
      ItemDifficultyLogit5SE string `xml:"ItemDifficultyLogit5SE"`
      ItemDifficultyLogit62SE string `xml:"ItemDifficultyLogit62SE"`
      ItemProficiencyBand string `xml:"ItemProficiencyBand"`
      ItemProficiencyLevel string `xml:"ItemProficiencyLevel"`
      ExemplarURL string `xml:"ExemplarURL"`
      ItemSubstitutedForList SubstituteItemListType `xml:"ItemSubstitutedForList"`
      ContentDescriptionList ContentDescriptionListType `xml:"ContentDescriptionList"`
      StimulusList StimulusListType `xml:"StimulusList"`
      NAPWritingRubricList NAPWritingRubricListType `xml:"NAPWritingRubricList"`
      
      }
    
    type NAPTestletContentType struct {
        NAPTestletLocalId LocalIdType `xml:"NAPTestletLocalId"`
      TestletName string `xml:"TestletName"`
      Node string `xml:"Node"`
      LocationInStage string `xml:"LocationInStage"`
      TestletMaximumScore string `xml:"TestletMaximumScore"`
      
      }
    
    type NAPTestContentType struct {
        NAPTestLocalId LocalIdType `xml:"NAPTestLocalId"`
      TestName string `xml:"TestName"`
      TestLevel YearLevelType `xml:"TestLevel"`
      TestType string `xml:"TestType"`
      Domain string `xml:"Domain"`
      TestYear SchoolYearType `xml:"TestYear"`
      StagesCount string `xml:"StagesCount"`
      DomainBands DomainBandsContainerType `xml:"DomainBands"`
      DomainProficiency DomainProficiencyContainerType `xml:"DomainProficiency"`
      
      }
    
    type PlausibleScaledValueListType struct {
        PlausibleScaledValue []string `xml:"PlausibleScaledValue"`
      
      }
    
    type SubstituteItemListType struct {
        SubstituteItem []SubstituteItemType `xml:"SubstituteItem"`
      
      }
    
    type SubstituteItemType struct {
        SubstituteItemRefId string `xml:"SubstituteItemRefId"`
      SubstituteItemLocalId LocalIdType `xml:"SubstituteItemLocalId"`
      PNPCodeList PNPCodeListType `xml:"PNPCodeList"`
      
      }
    
    type CodeFrameTestItemListType struct {
        TestItem []CodeFrameTestItemType `xml:"TestItem"`
      
      }
    
    type CodeFrameTestItemType struct {
        TestItemRefId string `xml:"TestItemRefId"`
      SequenceNumber string `xml:"SequenceNumber"`
      TestItemContent NAPTestItemContentType `xml:"TestItemContent"`
      
      }
    
    type StimulusLocalIdListType struct {
        StimulusLocalId []LocalIdType `xml:"StimulusLocalId"`
      
      }
    
    type DomainBandsContainerType struct {
        Band1Lower string `xml:"Band1Lower"`
      Band1Upper string `xml:"Band1Upper"`
      Band2Lower string `xml:"Band2Lower"`
      Band2Upper string `xml:"Band2Upper"`
      Band3Lower string `xml:"Band3Lower"`
      Band3Upper string `xml:"Band3Upper"`
      Band4Lower string `xml:"Band4Lower"`
      Band4Upper string `xml:"Band4Upper"`
      Band5Lower string `xml:"Band5Lower"`
      Band5Upper string `xml:"Band5Upper"`
      Band6Lower string `xml:"Band6Lower"`
      Band6Upper string `xml:"Band6Upper"`
      Band7Lower string `xml:"Band7Lower"`
      Band7Upper string `xml:"Band7Upper"`
      Band8Lower string `xml:"Band8Lower"`
      Band8Upper string `xml:"Band8Upper"`
      Band9Lower string `xml:"Band9Lower"`
      Band9Upper string `xml:"Band9Upper"`
      Band10Lower string `xml:"Band10Lower"`
      Band10Upper string `xml:"Band10Upper"`
      
      }
    
    type DomainProficiencyContainerType struct {
        Level1Lower string `xml:"Level1Lower"`
      Level1Upper string `xml:"Level1Upper"`
      Level2Lower string `xml:"Level2Lower"`
      Level2Upper string `xml:"Level2Upper"`
      Level3Lower string `xml:"Level3Lower"`
      Level3Upper string `xml:"Level3Upper"`
      Level4Lower string `xml:"Level4Lower"`
      Level4Upper string `xml:"Level4Upper"`
      
      }
    
    type NAPTestItemListType struct {
        TestItem []NAPTestItem2Type `xml:"TestItem"`
      
      }
    
    type NAPTestItem2Type struct {
        TestItemRefId string `xml:"TestItemRefId"`
      TestItemLocalId LocalIdType `xml:"TestItemLocalId"`
      SequenceNumber string `xml:"SequenceNumber"`
      
      }
    
    type NAPCodeFrameTestletListType struct {
        Testlet []NAPTestletCodeFrameType `xml:"Testlet"`
      
      }
    
    type NAPTestletCodeFrameType struct {
        NAPTestletRefId string `xml:"NAPTestletRefId"`
      TestletContent NAPTestletContentType `xml:"TestletContent"`
      TestItemList CodeFrameTestItemListType `xml:"TestItemList"`
      
      }
    
    type NAPStudentResponseTestletListType struct {
        Testlet []NAPTestletResponseType `xml:"Testlet"`
      
      }
    
    type NAPTestletResponseType struct {
        NAPTestletRefId string `xml:"NAPTestletRefId"`
      NAPTestletLocalId LocalIdType `xml:"NAPTestletLocalId"`
      TestletSubScore string `xml:"TestletSubScore"`
      ItemResponseList NAPTestletItemResponseListType `xml:"ItemResponseList"`
      
      }
    
    type NAPTestletItemResponseListType struct {
        ItemResponse []NAPTestletResponseItemType `xml:"ItemResponse"`
      
      }
    
    type NAPTestletResponseItemType struct {
        NAPTestItemRefId string `xml:"NAPTestItemRefId"`
      NAPTestItemLocalId LocalIdType `xml:"NAPTestItemLocalId"`
      Response string `xml:"Response"`
      ResponseCorrectness string `xml:"ResponseCorrectness"`
      Score string `xml:"Score"`
      LapsedTimeItem string `xml:"LapsedTimeItem"`
      SequenceNumber string `xml:"SequenceNumber"`
      ItemWeight string `xml:"ItemWeight"`
      SubscoreList NAPSubscoreListType `xml:"SubscoreList"`
      
      }
    
    type NAPSubscoreListType struct {
        Subscore []NAPSubscoreType `xml:"Subscore"`
      
      }
    
    type NAPSubscoreType struct {
        SubscoreType string `xml:"SubscoreType"`
      SubscoreValue string `xml:"SubscoreValue"`
      
      }
    
    type DomainScoreType struct {
        RawScore string `xml:"RawScore"`
      ScaledScoreValue string `xml:"ScaledScoreValue"`
      ScaledScoreLogitValue string `xml:"ScaledScoreLogitValue"`
      ScaledScoreStandardError string `xml:"ScaledScoreStandardError"`
      ScaledScoreLogitStandardError string `xml:"ScaledScoreLogitStandardError"`
      StudentDomainBand string `xml:"StudentDomainBand"`
      StudentProficiency string `xml:"StudentProficiency"`
      PlausibleScaledValueList PlausibleScaledValueListType `xml:"PlausibleScaledValueList"`
      
      }
    
    type NAPWritingRubricListType struct {
        NAPWritingRubric []NAPWritingRubricType `xml:"NAPWritingRubric"`
      
      }
    
    type NAPWritingRubricType struct {
        RubricType string `xml:"RubricType"`
      ScoreList ScoreListType `xml:"ScoreList"`
      Descriptor string `xml:"Descriptor"`
      
      }
    
    type ScoreListType struct {
        Score []ScoreType `xml:"Score"`
      
      }
    
    type ScoreType struct {
        MaxScoreValue string `xml:"MaxScoreValue"`
      ScoreDescriptionList ScoreDescriptionListType `xml:"ScoreDescriptionList"`
      
      }
    
    type ScoreDescriptionListType struct {
        ScoreDescription []ScoreDescriptionType `xml:"ScoreDescription"`
      
      }
    
    type ScoreDescriptionType struct {
        ScoreValue string `xml:"ScoreValue"`
      Descriptor string `xml:"Descriptor"`
      
      }
    
    type StimulusListType struct {
        Stimulus []StimulusType `xml:"Stimulus"`
      
      }
    
    type StimulusType struct {
        StimulusLocalId LocalIdType `xml:"StimulusLocalId"`
      TextGenre string `xml:"TextGenre"`
      TextType string `xml:"TextType"`
      WordCount string `xml:"WordCount"`
      TextDescriptor string `xml:"TextDescriptor"`
      Content string `xml:"Content"`
      
      }
    
    type ContentDescriptionListType struct {
        ContentDescription []string `xml:"ContentDescription"`
      
      }
    
    type PNPCodeListType struct {
        PNPCode []string `xml:"PNPCode"`
      
      }
    
    type AdjustmentContainerType struct {
        PNPCodeList PNPCodeListType `xml:"PNPCodeList"`
      BookletType string `xml:"BookletType"`
      
      }
    
    type TestDisruptionListType struct {
        TestDisruption []TestDisruptionType `xml:"TestDisruption"`
      
      }
    
    type TestDisruptionType struct {
        Event string `xml:"Event"`
      
      }
    
    type CalendarSummaryListType struct {
        CalendarSummaryRefId []string `xml:"CalendarSummaryRefId"`
      
      }
    
    type VisaSubClassType struct {
        Code VisaSubClassCodeType `xml:"Code"`
      VisaExpiryDate string `xml:"VisaExpiryDate"`
      ATEExpiryDate string `xml:"ATEExpiryDate"`
      ATEStartDate string `xml:"ATEStartDate"`
      VisaStatisticalCode string `xml:"VisaStatisticalCode"`
      
      }
    
    type VisaSubClassListType struct {
        VisaSubClass []VisaSubClassType `xml:"VisaSubClass"`
      
      }
    
    type VisaSubClassCodeType string
    type LanguageBaseType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      LanguageType string `xml:"LanguageType"`
      Dialect string `xml:"Dialect"`
      
      }
    
    type ReligiousEventListType struct {
        ReligiousEvent []ReligiousEventType `xml:"ReligiousEvent"`
      
      }
    
    type ReligiousEventType struct {
        Type string `xml:"Type"`
      Date string `xml:"Date"`
      
      }
    
    type ReligionType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type DwellingArrangementType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type CountryListType struct {
        CountryOfCitizenship []CountryType `xml:"CountryOfCitizenship"`
      
      }
    
    type CountryList2Type struct {
        CountryOfResidency []CountryType `xml:"CountryOfResidency"`
      
      }
    
    type DebitOrCreditAmountType struct {
        MonetaryAmountType
          Type string `xml:"Type,attr"`
      
      }
    
    type ScheduledActivityOverrideType struct {
          DateOfOverride string `xml:"DateOfOverride,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type ActivityTimeType struct {
        CreationDate string `xml:"CreationDate"`
      Duration ActivityTimeType_Duration `xml:"Duration"`
      StartDate string `xml:"StartDate"`
      FinishDate string `xml:"FinishDate"`
      DueDate string `xml:"DueDate"`
      
      }
    
    type SchoolCourseInfoOverrideType struct {
        Override string `xml:"Override,attr"`
      CourseCode string `xml:"CourseCode"`
      StateCourseCode string `xml:"StateCourseCode"`
      DistrictCourseCode string `xml:"DistrictCourseCode"`
      SubjectArea SubjectAreaType `xml:"SubjectArea"`
      CourseTitle string `xml:"CourseTitle"`
      InstructionalLevel string `xml:"InstructionalLevel"`
      CourseCredits string `xml:"CourseCredits"`
      
      }
    
    type LocationOfInstructionType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type LanguageOfInstructionType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type MediumOfInstructionType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type StudentActivityType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type ContactFlagsType struct {
        ParentLegalGuardian string `xml:"ParentLegalGuardian"`
      PickupRights string `xml:"PickupRights"`
      LivesWith string `xml:"LivesWith"`
      AccessToRecords string `xml:"AccessToRecords"`
      ReceivesAssessmentReport string `xml:"ReceivesAssessmentReport"`
      EmergencyContact string `xml:"EmergencyContact"`
      HasCustody string `xml:"HasCustody"`
      DisciplinaryContact string `xml:"DisciplinaryContact"`
      AttendanceContact string `xml:"AttendanceContact"`
      PrimaryCareProvider string `xml:"PrimaryCareProvider"`
      FeesBilling string `xml:"FeesBilling"`
      FeesAccess string `xml:"FeesAccess"`
      FamilyMail string `xml:"FamilyMail"`
      InterventionOrder string `xml:"InterventionOrder"`
      
      }
    
    type AgencyType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type YearRangeType struct {
        Start YearLevelType `xml:"Start"`
      End YearLevelType `xml:"End"`
      
      }
    
    type CreationUserType struct {
        Type string `xml:"Type,attr"`
      UserId string `xml:"UserId"`
      
      }
    
    type AuditInfoType struct {
        CreationUser CreationUserType `xml:"CreationUser"`
      CreationDateTime string `xml:"CreationDateTime"`
      
      }
    
    type AttendanceInfoType struct {
        CountsTowardAttendance string `xml:"CountsTowardAttendance"`
      AttendanceValue string `xml:"AttendanceValue"`
      
      }
    
    type CalendarDateInfoType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type ProgramAvailabilityType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type ReferralSourceType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type PromotionInfoType struct {
        PromotionStatus string `xml:"PromotionStatus"`
      
      }
    
    type CatchmentStatusContainerType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type StudentExitStatusContainerType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type StudentExitContainerType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type StudentEntryContainerType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type StudentMostRecentContainerType struct {
        SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      HomeroomLocalId LocalIdType `xml:"HomeroomLocalId"`
      YearLevel YearLevelType `xml:"YearLevel"`
      FTE string `xml:"FTE"`
      Parent1Language string `xml:"Parent1Language"`
      Parent2Language string `xml:"Parent2Language"`
      Parent1EmploymentType string `xml:"Parent1EmploymentType"`
      Parent2EmploymentType string `xml:"Parent2EmploymentType"`
      Parent1SchoolEducationLevel string `xml:"Parent1SchoolEducationLevel"`
      Parent2SchoolEducationLevel string `xml:"Parent2SchoolEducationLevel"`
      Parent1NonSchoolEducation string `xml:"Parent1NonSchoolEducation"`
      Parent2NonSchoolEducation string `xml:"Parent2NonSchoolEducation"`
      LocalCampusId LocalIdType `xml:"LocalCampusId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId"`
      TestLevel YearLevelType `xml:"TestLevel"`
      Homegroup string `xml:"Homegroup"`
      ClassCode string `xml:"ClassCode"`
      MembershipType string `xml:"MembershipType"`
      FFPOS string `xml:"FFPOS"`
      ReportingSchoolId LocalIdType `xml:"ReportingSchoolId"`
      OtherEnrollmentSchoolACARAId LocalIdType `xml:"OtherEnrollmentSchoolACARAId"`
      OtherSchoolName string `xml:"OtherSchoolName"`
      DisabilityLevelOfAdjustment string `xml:"DisabilityLevelOfAdjustment"`
      DisabilityCategory string `xml:"DisabilityCategory"`
      CensusAge string `xml:"CensusAge"`
      DistanceEducationStudent string `xml:"DistanceEducationStudent"`
      BoardingStatus string `xml:"BoardingStatus"`
      
      }
    
    type StaffMostRecentContainerType struct {
        SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId"`
      LocalCampusId LocalIdType `xml:"LocalCampusId"`
      NAPLANClassList NAPLANClassListType `xml:"NAPLANClassList"`
      HomeGroup string `xml:"HomeGroup"`
      
      }
    
    type StaffActivityExtensionType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type TotalEnrollmentsType struct {
        Girls string `xml:"Girls"`
      Boys string `xml:"Boys"`
      TotalStudents string `xml:"TotalStudents"`
      
      }
    
    type CampusContainerType struct {
        ParentSchoolId string `xml:"ParentSchoolId"`
      SchoolCampusId string `xml:"SchoolCampusId"`
      CampusType string `xml:"CampusType"`
      AdminStatus string `xml:"AdminStatus"`
      
      }
    
    type HouseholdContactInfoListType struct {
        HouseholdContactInfo []HouseholdContactInfoType `xml:"HouseholdContactInfo"`
      
      }
    
    type HouseholdContactInfoType struct {
        PreferenceNumber string `xml:"PreferenceNumber"`
      HouseholdContactId LocalIdType `xml:"HouseholdContactId"`
      HouseholdSalutation string `xml:"HouseholdSalutation"`
      AddressList AddressListType `xml:"AddressList"`
      EmailList EmailListType `xml:"EmailList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList"`
      
      }
    
    type StatementCodesType struct {
        StatementCode []string `xml:"StatementCode"`
      
      }
    
    type StatementsType struct {
        Statement []string `xml:"Statement"`
      
      }
    
    type ProgramFundingSourcesType struct {
        ProgramFundingSource []ProgramFundingSourceType `xml:"ProgramFundingSource"`
      
      }
    
    type ProgramFundingSourceType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type AttendanceTimesType struct {
        AttendanceTime []AttendanceTimeType `xml:"AttendanceTime"`
      
      }
    
    type AttendanceTimeType struct {
        AttendanceType string `xml:"AttendanceType"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus"`
      StartTime string `xml:"StartTime"`
      EndTime string `xml:"EndTime"`
      DurationValue string `xml:"DurationValue"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId"`
      AttendanceNote string `xml:"AttendanceNote"`
      
      }
    
    type PeriodAttendancesType struct {
        PeriodAttendance []PeriodAttendanceType `xml:"PeriodAttendance"`
      
      }
    
    type PeriodAttendanceType struct {
        AttendanceType string `xml:"AttendanceType"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus"`
      Date string `xml:"Date"`
      SessionInfoRefId string `xml:"SessionInfoRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId"`
      TimetablePeriod string `xml:"TimetablePeriod"`
      DayId LocalIdType `xml:"DayId"`
      StartTime string `xml:"StartTime"`
      EndTime string `xml:"EndTime"`
      TimeIn string `xml:"TimeIn"`
      TimeOut string `xml:"TimeOut"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList"`
      RoomList RoomListType `xml:"RoomList"`
      AttendanceNote string `xml:"AttendanceNote"`
      
      }
    
    type StaffSubjectListType struct {
        StaffSubject []StaffSubjectType `xml:"StaffSubject"`
      
      }
    
    type StaffSubjectType struct {
        PreferenceNumber string `xml:"PreferenceNumber"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId"`
      
      }
    
    type TeachingGroupListType struct {
        TeachingGroupRefId []string `xml:"TeachingGroupRefId"`
      
      }
    
    type ScheduledTeacherListType struct {
        TeacherCover []TeacherCoverType `xml:"TeacherCover"`
      
      }
    
    type TeacherCoverType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      StartTime string `xml:"StartTime"`
      FinishTime string `xml:"FinishTime"`
      Credit string `xml:"Credit"`
      Supervision string `xml:"Supervision"`
      Weighting string `xml:"Weighting"`
      
      }
    
    type RoomListType struct {
        RoomInfoRefId []string `xml:"RoomInfoRefId"`
      
      }
    
    type StaffListType struct {
        StaffPersonalRefId []string `xml:"StaffPersonalRefId"`
      
      }
    
    type AuthorsType struct {
        Author []string `xml:"Author"`
      
      }
    
    type OrganizationsType struct {
        Organization []string `xml:"Organization"`
      
      }
    
    type PurchasingItemsType struct {
        PurchasingItem []PurchasingItemType `xml:"PurchasingItem"`
      
      }
    
    type PurchasingItemType struct {
        ItemNumber string `xml:"ItemNumber"`
      ItemDescription string `xml:"ItemDescription"`
      LocalItemId LocalIdType `xml:"LocalItemId"`
      Quantity string `xml:"Quantity"`
      UnitCost MonetaryAmountType `xml:"UnitCost"`
      TotalCost MonetaryAmountType `xml:"TotalCost"`
      QuantityDelivered string `xml:"QuantityDelivered"`
      CancelledOrder string `xml:"CancelledOrder"`
      TaxRate string `xml:"TaxRate"`
      ExpenseAccounts ExpenseAccountsType `xml:"ExpenseAccounts"`
      
      }
    
    type ExpenseAccountsType struct {
        ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount"`
      
      }
    
    type ExpenseAccountType struct {
        AccountCode string `xml:"AccountCode"`
      Amount MonetaryAmountType `xml:"Amount"`
      FinancialAccountRefId string `xml:"FinancialAccountRefId"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod"`
      
      }
    
    type SchoolProgramListType struct {
        Program []SchoolProgramType `xml:"Program"`
      
      }
    
    type SchoolProgramType struct {
        Category string `xml:"Category"`
      Type string `xml:"Type"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type LearningObjectivesType struct {
        LearningObjective []string `xml:"LearningObjective"`
      
      }
    
    type RecognitionListType struct {
        Recognition []string `xml:"Recognition"`
      
      }
    
    type LResourcesType struct {
        LearningResourceRefId []ResourcesType `xml:"LearningResourceRefId"`
      
      }
    
    type ResourcesType struct {
          ResourceType string `xml:"ResourceType,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type SourceObjectsType struct {
      SourceObject []SourceObjectsType_SourceObject `xml:"SourceObject"`
      
      }
    
    type StudentsType struct {
        StudentPersonalRefId []string `xml:"StudentPersonalRefId"`
      
      }
    
    type PrerequisitesType struct {
        Prerequisite []string `xml:"Prerequisite"`
      
      }
    
    type EssentialMaterialsType struct {
        EssentialMaterial []string `xml:"EssentialMaterial"`
      
      }
    
    type TechnicalRequirementsType struct {
        TechnicalRequirement string `xml:"TechnicalRequirement"`
      
      }
    
    type SoftwareRequirementListType struct {
        SoftwareRequirement []SoftwareRequirementType `xml:"SoftwareRequirement"`
      
      }
    
    type SoftwareRequirementType struct {
        SoftwareTitle string `xml:"SoftwareTitle"`
      Version string `xml:"Version"`
      Vendor string `xml:"Vendor"`
      OS string `xml:"OS"`
      
      }
    
    type HouseholdListType struct {
        Household []LocalIdType `xml:"Household"`
      
      }
    
    type StudentSubjectChoiceListType struct {
        StudentSubjectChoice []StudentSubjectChoiceType `xml:"StudentSubjectChoice"`
      
      }
    
    type StudentSubjectChoiceType struct {
        PreferenceNumber string `xml:"PreferenceNumber"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId"`
      StudyDescription SubjectAreaType `xml:"StudyDescription"`
      OtherSchoolLocalId LocalIdType `xml:"OtherSchoolLocalId"`
      
      }
    
    type IdentityAssertionsType struct {
      IdentityAssertion []IdentityAssertionsType_IdentityAssertion `xml:"IdentityAssertion"`
      
      }
    
    type LearningStandardsType struct {
        LearningStandardItemRefId []string `xml:"LearningStandardItemRefId"`
      
      }
    
    type LearningResourcesType struct {
        LearningResourceRefId []string `xml:"LearningResourceRefId"`
      
      }
    
    type LearningStandardsDocumentType struct {
        LearningStandardDocumentRefId []string `xml:"LearningStandardDocumentRefId"`
      
      }
    
    type ComponentsType struct {
        Component []ComponentType `xml:"Component"`
      
      }
    
    type ComponentType struct {
        Name string `xml:"Name"`
      Reference string `xml:"Reference"`
      Description string `xml:"Description"`
      Strategies StrategiesType `xml:"Strategies"`
      AssociatedObjects AssociatedObjectsType `xml:"AssociatedObjects"`
      
      }
    
    type StrategiesType struct {
        Strategy []string `xml:"Strategy"`
      
      }
    
    type AssociatedObjectsType struct {
      AssociatedObject []AssociatedObjectsType_AssociatedObject `xml:"AssociatedObject"`
      
      }
    
    type EvaluationsType struct {
        Evaluation []EvaluationType `xml:"Evaluation"`
      
      }
    
    type EvaluationType struct {
        RefId RefIdType `xml:"RefId,attr"`
      Description string `xml:"Description"`
      Date string `xml:"Date"`
      Name NameType `xml:"Name"`
      
      }
    
    type ApprovalsType struct {
        Approval []ApprovalType `xml:"Approval"`
      
      }
    
    type ApprovalType struct {
        Organization string `xml:"Organization"`
      Date string `xml:"Date"`
      
      }
    
    type MediaTypesType struct {
        MediaType []string `xml:"MediaType"`
      
      }
    
    type LEAContactListType struct {
        LEAContact []LEAContactType `xml:"LEAContact"`
      
      }
    
    type LEAContactType struct {
        PublishInDirectory PublishInDirectoryType `xml:"PublishInDirectory"`
      ContactInfo ContactInfoType `xml:"ContactInfo"`
      
      }
    
    type FinancialAccountRefIdListType struct {
        FinancialAccountRefId []string `xml:"FinancialAccountRefId"`
      
      }
    
    type AccountCodeListType struct {
        AccountCode []string `xml:"AccountCode"`
      
      }
    
    type JournalAdjustmentListType struct {
        JournalAdjustment []JournalAdjustmentType `xml:"JournalAdjustment"`
      
      }
    
    type JournalAdjustmentType struct {
        DebitFinancialAccountRefId string `xml:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId string `xml:"CreditFinancialAccountRefId"`
      DebitAccountCode string `xml:"DebitAccountCode"`
      CreditAccountCode string `xml:"CreditAccountCode"`
      GSTCodeOriginal string `xml:"GSTCodeOriginal"`
      GSTCodeReplacement string `xml:"GSTCodeReplacement"`
      LineAdjustmentAmount MonetaryAmountType `xml:"LineAdjustmentAmount"`
      
      }
    
    type PaymentReceiptLineListType struct {
        PaymentReceiptLine []PaymentReceiptLineType `xml:"PaymentReceiptLine"`
      
      }
    
    type PaymentReceiptLineType struct {
        InvoiceRefId string `xml:"InvoiceRefId"`
      LocalId LocalIdType `xml:"LocalId"`
      LocalPaymentReceiptLineId LocalIdType `xml:"LocalPaymentReceiptLineId"`
      TransactionAmount DebitOrCreditAmountType `xml:"TransactionAmount"`
      FinancialAccountRefId string `xml:"FinancialAccountRefId"`
      AccountCode string `xml:"AccountCode"`
      TransactionDescription string `xml:"TransactionDescription"`
      TaxRate string `xml:"TaxRate"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount"`
      
      }
    
    type PasswordListType struct {
      Password []PasswordListType_Password `xml:"Password"`
      
      }
    
    type ExclusionRulesType struct {
        ExclusionRule []ExclusionRuleType `xml:"ExclusionRule"`
      
      }
    
    type ExclusionRuleType struct {
          Type string `xml:"Type,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type CharacteristicsType struct {
        AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId"`
      
      }
    
    type ContactsType struct {
        Contact []ContactType `xml:"Contact"`
      
      }
    
    type ContactType struct {
        Name NameType `xml:"Name"`
      Address AddressType `xml:"Address"`
      PhoneNumber PhoneNumberType `xml:"PhoneNumber"`
      Email EmailType `xml:"Email"`
      
      }
    
    type TeachingGroupPeriodListType struct {
        TeachingGroupPeriod []TeachingGroupPeriodType `xml:"TeachingGroupPeriod"`
      
      }
    
    type TeachingGroupPeriodType struct {
        TimeTableCellRefId string `xml:"TimeTableCellRefId"`
      RoomNumber HomeroomNumberType `xml:"RoomNumber"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      DayId LocalIdType `xml:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId"`
      StartTime string `xml:"StartTime"`
      CellType string `xml:"CellType"`
      
      }
    
    type TeacherListType struct {
        TeachingGroupTeacher []TeachingGroupTeacherType `xml:"TeachingGroupTeacher"`
      
      }
    
    type TeachingGroupTeacherType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      Name NameOfRecordType `xml:"Name"`
      Association string `xml:"Association"`
      
      }
    
    type StudentListType struct {
        TeachingGroupStudent []TeachingGroupStudentType `xml:"TeachingGroupStudent"`
      
      }
    
    type TeachingGroupStudentType struct {
        StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      StudentLocalId LocalIdType `xml:"StudentLocalId"`
      Name NameOfRecordType `xml:"Name"`
      
      }
    
    type TimeTableDayListType struct {
        TimeTableDay []TimeTableDayType `xml:"TimeTableDay"`
      
      }
    
    type TimeTableDayType struct {
        DayId LocalIdType `xml:"DayId"`
      DayTitle string `xml:"DayTitle"`
      TimeTablePeriodList TimeTablePeriodListType `xml:"TimeTablePeriodList"`
      
      }
    
    type TimeTablePeriodListType struct {
        TimeTablePeriod []TimeTablePeriodType `xml:"TimeTablePeriod"`
      
      }
    
    type TimeTablePeriodType struct {
        PeriodId LocalIdType `xml:"PeriodId"`
      PeriodTitle string `xml:"PeriodTitle"`
      BellPeriod string `xml:"BellPeriod"`
      StartTime string `xml:"StartTime"`
      EndTime string `xml:"EndTime"`
      RegularSchoolPeriod string `xml:"RegularSchoolPeriod"`
      InstructionalMinutes string `xml:"InstructionalMinutes"`
      UseInAttendanceCalculations string `xml:"UseInAttendanceCalculations"`
      
      }
    
    type NAPLANClassListType struct {
        ClassCode []string `xml:"ClassCode"`
      
      }
    
    type SchoolGroupListType struct {
        SchoolGroup []LocalIdType `xml:"SchoolGroup"`
      
      }
    
    type YearLevelEnrollmentListType struct {
        YearLevelEnrollment []YearLevelEnrollmentType `xml:"YearLevelEnrollment"`
      
      }
    
    type YearLevelEnrollmentType struct {
        Year string `xml:"Year"`
      Enrollment string `xml:"Enrollment"`
      
      }
    
    type SchoolFocusListType struct {
        SchoolFocus []string `xml:"SchoolFocus"`
      
      }
    
    type AlertMessagesType struct {
        AlertMessage []AlertMessageType `xml:"AlertMessage"`
      
      }
    
    type AlertMessageType struct {
          Type string `xml:"Type,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type MedicalAlertMessagesType struct {
        MedicalAlertMessage []MedicalAlertMessageType `xml:"MedicalAlertMessage"`
      
      }
    
    type MedicalAlertMessageType struct {
          Severity string `xml:"Severity,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type OtherIdListType struct {
        OtherId []OtherIdType `xml:"OtherId"`
      
      }
    
    type OtherIdType struct {
          Type string `xml:"Type,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type BaseNameType struct {
        Title string `xml:"Title"`
      FamilyName string `xml:"FamilyName"`
      GivenName string `xml:"GivenName"`
      MiddleName string `xml:"MiddleName"`
      FamilyNameFirst string `xml:"FamilyNameFirst"`
      PreferredFamilyName string `xml:"PreferredFamilyName"`
      PreferredFamilyNameFirst string `xml:"PreferredFamilyNameFirst"`
      PreferredGivenName string `xml:"PreferredGivenName"`
      Suffix string `xml:"Suffix"`
      FullName string `xml:"FullName"`
      
      }
    
    type NameOfRecordType struct {
        BaseNameType
          Type string `xml:"Type,attr"`
      
      }
    
    type OtherNameType struct {
        BaseNameType
          Type string `xml:"Type,attr"`
      
      }
    
    type PartialDateType string
    type LocalIdType string
    type LocationType struct {
        Type string `xml:"Type,attr"`
      LocationName string `xml:"LocationName"`
      LocationRefId LocationType_LocationRefId `xml:"LocationRefId"`
      
      }
    
    type StateProvinceIdType string
    type AttendanceCodeType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type YearLevelType struct {
        Code string `xml:"Code"`
      
      }
    
    type PersonInfoType struct {
        Name NameOfRecordType `xml:"Name"`
      OtherNames OtherNamesType `xml:"OtherNames"`
      Demographics DemographicsType `xml:"Demographics"`
      AddressList AddressListType `xml:"AddressList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList"`
      EmailList EmailListType `xml:"EmailList"`
      HouseholdContactInfoList HouseholdContactInfoListType `xml:"HouseholdContactInfoList"`
      
      }
    
    type YearLevelsType struct {
        YearLevel []YearLevelType `xml:"YearLevel"`
      
      }
    
    type SchoolURLType string
    type PrincipalInfoType struct {
        ContactName NameOfRecordType `xml:"ContactName"`
      ContactTitle string `xml:"ContactTitle"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList"`
      EmailList EmailListType `xml:"EmailList"`
      
      }
    
    type SchoolContactType struct {
        PublishInDirectory PublishInDirectoryType `xml:"PublishInDirectory"`
      ContactInfo ContactInfoType `xml:"ContactInfo"`
      
      }
    
    type SchoolContactListType struct {
        SchoolContact []SchoolContactType `xml:"SchoolContact"`
      
      }
    
    type PublishInDirectoryType string
    type ContactInfoType struct {
        Name NameType `xml:"Name"`
      PositionTitle string `xml:"PositionTitle"`
      Role string `xml:"Role"`
      RegistrationDetails string `xml:"RegistrationDetails"`
      Qualifications string `xml:"Qualifications"`
      Address AddressType `xml:"Address"`
      EmailList EmailListType `xml:"EmailList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList"`
      
      }
    
    type AddressStreetType struct {
        Line1 string `xml:"Line1"`
      Line2 string `xml:"Line2"`
      Line3 string `xml:"Line3"`
      Complex string `xml:"Complex"`
      StreetNumber string `xml:"StreetNumber"`
      StreetPrefix string `xml:"StreetPrefix"`
      StreetName string `xml:"StreetName"`
      StreetType string `xml:"StreetType"`
      StreetSuffix string `xml:"StreetSuffix"`
      ApartmentType string `xml:"ApartmentType"`
      ApartmentNumberPrefix string `xml:"ApartmentNumberPrefix"`
      ApartmentNumber string `xml:"ApartmentNumber"`
      ApartmentNumberSuffix string `xml:"ApartmentNumberSuffix"`
      
      }
    
    type AddressType struct {
        Type string `xml:"Type,attr"`
      Role string `xml:"Role,attr"`
      EffectiveFromDate string `xml:"EffectiveFromDate"`
      EffectiveToDate string `xml:"EffectiveToDate"`
      Street AddressStreetType `xml:"Street"`
      City string `xml:"City"`
      StateProvince StateProvinceType `xml:"StateProvince"`
      Country CountryType `xml:"Country"`
      PostalCode string `xml:"PostalCode"`
      GridLocation GridLocationType `xml:"GridLocation"`
      MapReference MapReferenceType `xml:"MapReference"`
      RadioContact string `xml:"RadioContact"`
      Community string `xml:"Community"`
      LocalId LocalIdType `xml:"LocalId"`
      AddressGlobalUID GUIDType `xml:"AddressGlobalUID"`
      StatisticalAreas StatisticalAreasType `xml:"StatisticalAreas"`
      
      }
    
    type MapReferenceType struct {
        Type string `xml:"Type,attr"`
      XCoordinate string `xml:"XCoordinate"`
      YCoordinate string `xml:"YCoordinate"`
      MapNumber string `xml:"MapNumber"`
      
      }
    
    type StatisticalAreasType struct {
        StatisticalArea []StatisticalAreaType `xml:"StatisticalArea"`
      
      }
    
    type StatisticalAreaType struct {
          SpatialUnitType string `xml:"SpatialUnitType,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type AddressListType struct {
        Address []AddressType `xml:"Address"`
      
      }
    
    type EmailListType struct {
        Email []EmailType `xml:"Email"`
      
      }
    
    type EmailType struct {
          Type string `xml:"Type,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type PhoneNumberListType struct {
        PhoneNumber []PhoneNumberType `xml:"PhoneNumber"`
      
      }
    
    type PhoneNumberType struct {
        Type string `xml:"Type,attr"`
      Number string `xml:"Number"`
      Extension string `xml:"Extension"`
      ListedStatus string `xml:"ListedStatus"`
      Preference string `xml:"Preference"`
      
      }
    
    type CountryType string
    type GridLocationType struct {
        Latitude string `xml:"Latitude"`
      Longitude string `xml:"Longitude"`
      
      }
    
    type OperationalStatusType string
    type StateProvinceType string
    type SchoolYearType string
    type ElectronicIdListType struct {
        ElectronicId []ElectronicIdType `xml:"ElectronicId"`
      
      }
    
    type ElectronicIdType struct {
          Type string `xml:"Type,attr"`
      
        Value string `xml:",chardata"`
      }
    
    type OtherNamesType struct {
        Name []OtherNameType `xml:"Name"`
      
      }
    
    type DemographicsType struct {
        IndigenousStatus string `xml:"IndigenousStatus"`
      Sex string `xml:"Sex"`
      BirthDate BirthDateType `xml:"BirthDate"`
      DateOfDeath string `xml:"DateOfDeath"`
      BirthDateVerification string `xml:"BirthDateVerification"`
      PlaceOfBirth string `xml:"PlaceOfBirth"`
      StateOfBirth StateProvinceType `xml:"StateOfBirth"`
      CountryOfBirth CountryType `xml:"CountryOfBirth"`
      CountriesOfCitizenship CountryListType `xml:"CountriesOfCitizenship"`
      CountriesOfResidency CountryList2Type `xml:"CountriesOfResidency"`
      CountryArrivalDate string `xml:"CountryArrivalDate"`
      AustralianCitizenshipStatus string `xml:"AustralianCitizenshipStatus"`
      EnglishProficiency EnglishProficiencyType `xml:"EnglishProficiency"`
      LanguageList LanguageListType `xml:"LanguageList"`
      DwellingArrangement DwellingArrangementType `xml:"DwellingArrangement"`
      Religion ReligionType `xml:"Religion"`
      ReligiousEventList ReligiousEventListType `xml:"ReligiousEventList"`
      ReligiousRegion string `xml:"ReligiousRegion"`
      PermanentResident string `xml:"PermanentResident"`
      VisaSubClass VisaSubClassCodeType `xml:"VisaSubClass"`
      VisaStatisticalCode string `xml:"VisaStatisticalCode"`
      VisaExpiryDate string `xml:"VisaExpiryDate"`
      VisaSubClassList VisaSubClassListType `xml:"VisaSubClassList"`
      LBOTE string `xml:"LBOTE"`
      InterpreterRequired string `xml:"InterpreterRequired"`
      ImmunisationCertificateStatus string `xml:"ImmunisationCertificateStatus"`
      CulturalBackground string `xml:"CulturalBackground"`
      MaritalStatus string `xml:"MaritalStatus"`
      MedicareNumber string `xml:"MedicareNumber"`
      
      }
    
    type EnglishProficiencyType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type LanguageListType struct {
        Language []LanguageBaseType `xml:"Language"`
      
      }
    
    type BirthDateType string
    type ProjectedGraduationYearType string
    type OnTimeGraduationYearType string
    type RelationshipType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type EducationalLevelType string
    type GraduationDateType PartialDateType
    type NameType struct {
        BaseNameType
          Type string `xml:"Type,attr"`
      
      }
    
    type HomeroomNumberType string
    type TimeElementType struct {
        Type string `xml:"Type"`
      Code string `xml:"Code"`
      Name string `xml:"Name"`
      Value string `xml:"Value"`
      StartDateTime string `xml:"StartDateTime"`
      EndDateTime string `xml:"EndDateTime"`
      SpanGaps TimeElementType_SpanGaps `xml:"SpanGaps"`
      IsCurrent string `xml:"IsCurrent"`
      
      }
    
    type LifeCycleType struct {
      Created LifeCycleType_Created `xml:"Created"`
      ModificationHistory LifeCycleType_ModificationHistory `xml:"ModificationHistory"`
      TimeElements LifeCycleType_TimeElements `xml:"TimeElements"`
      
      }
    
    type OtherCodeListType struct {
      OtherCode []OtherCodeListType_OtherCode `xml:"OtherCode"`
      
      }
    
    type ProgramStatusType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type SubjectAreaListType struct {
        SubjectArea []SubjectAreaType `xml:"SubjectArea"`
      
      }
    
    type ACStrandAreaListType struct {
        ACStrandSubjectArea []ACStrandSubjectAreaType `xml:"ACStrandSubjectArea"`
      
      }
    
    type SubjectAreaType struct {
        Code string `xml:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList"`
      
      }
    
    type ACStrandSubjectAreaType struct {
        ACStrand string `xml:"ACStrand"`
      SubjectArea SubjectAreaType `xml:"SubjectArea"`
      
      }
    
    type EducationFilterType struct {
        LearningStandardItems LearningStandardsType `xml:"LearningStandardItems"`
      
      }
    
    type SIF_ExtendedElementsType struct {
      SIF_ExtendedElement []SIF_ExtendedElementsType_SIF_ExtendedElement `xml:"SIF_ExtendedElement"`
      
      }
    
    type SIF_MetadataType struct {
      TimeElements SIF_MetadataType_TimeElements `xml:"TimeElements"`
      LifeCycle LifeCycleType `xml:"LifeCycle"`
      EducationFilter EducationFilterType `xml:"EducationFilter"`
      
      }
    type AbstractContentPackageType_XMLData struct {
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentPackageType_TextData struct {
      MIMEType string `xml:"MIMEType,attr"`
      FileName string `xml:"FileName,attr"`
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentPackageType_BinaryData struct {
      MIMEType string `xml:"MIMEType,attr"`
      FileName string `xml:"FileName,attr"`
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentPackageType_Reference struct {
      MIMEType string `xml:"MIMEType,attr"`
      Description string `xml:"Description,attr"`
       URL string `xml:"URL"`
}
type AbstractContentElementType_XMLData struct {
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentElementType_TextData struct {
      MIMEType string `xml:"MIMEType,attr"`
      FileName string `xml:"FileName,attr"`
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentElementType_BinaryData struct {
      MIMEType string `xml:"MIMEType,attr"`
      FileName string `xml:"FileName,attr"`
      Description string `xml:"Description,attr"`
      Value string `xml:",chardata"`
}
type AbstractContentElementType_Reference struct {
      MIMEType string `xml:"MIMEType,attr"`
      Description string `xml:"Description,attr"`
       URL string `xml:"URL"`
}
type PersonInvolvementType_PersonRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type ActivityTimeType_Duration struct {
      Units string `xml:"Units,attr"`
      Value string `xml:",chardata"`
}
type SourceObjectsType_SourceObject struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type IdentityAssertionsType_IdentityAssertion struct {
      SchemaName string `xml:"SchemaName,attr"`
      Value string `xml:",chardata"`
}
type AssociatedObjectsType_AssociatedObject struct {
      SIF_RefObject ObjectNameType `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type PasswordListType_Password struct {
      Algorithm string `xml:"Algorithm,attr"`
      KeyName string `xml:"KeyName,attr"`
      Value string `xml:",chardata"`
}
type LocationType_LocationRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
type TimeElementType_SpanGaps struct {
      SpanGap []TimeElementType_SpanGap `xml:"SpanGap"`
}
type LifeCycleType_Created struct {
       DateTime string `xml:"DateTime"`
      Creators LifeCycleType_Creators `xml:"Creators"`
}
type LifeCycleType_ModificationHistory struct {
      Modified []LifeCycleType_Modified `xml:"Modified"`
}
type LifeCycleType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement"`
}
type OtherCodeListType_OtherCode struct {
      Codeset string `xml:"Codeset,attr"`
      Value string `xml:",chardata"`
}
type SIF_ExtendedElementsType_SIF_ExtendedElement struct {
      Name string `xml:"Name,attr"`
      Type string `xml:"Type,attr"`
      SIF_Action string `xml:"SIF_Action,attr"`
      Value ExtendedContentType `xml:",chardata"`
}
type SIF_MetadataType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement"`
}
type TimeElementType_SpanGap struct {
       Type string `xml:"Type"`
       Code string `xml:"Code"`
       Name string `xml:"Name"`
       Value string `xml:"Value"`
       StartDateTime string `xml:"StartDateTime"`
       EndDateTime string `xml:"EndDateTime"`
}
type LifeCycleType_Creators struct {
      Creator []LifeCycleType_Creator `xml:"Creator"`
}
type LifeCycleType_Modified struct {
       By string `xml:"By"`
       DateTime string `xml:"DateTime"`
       Description string `xml:"Description"`
}
type LifeCycleType_Creator struct {
       Name string `xml:"Name"`
       ID string `xml:"ID"`
}
