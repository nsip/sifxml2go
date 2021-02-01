package sifxml


    type AbstractContentElementType struct {
  abstractcontentelementtype `xml:"AbstractContentElementType" json:"AbstractContentElementType"`
}

type abstractcontentelementtype struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      XMLData *AbstractContentElementType_XMLData
      TextData *AbstractContentElementType_TextData
      BinaryData *AbstractContentElementType_BinaryData
      Reference *AbstractContentElementType_Reference
      
      }
    
    type MonetaryAmountType struct {
  monetaryamounttype `xml:"MonetaryAmountType" json:"MonetaryAmountType"`
}

type monetaryamounttype struct {
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
    type ResourceUsageContentTypeType struct {
  resourceusagecontenttypetype `xml:"ResourceUsageContentTypeType" json:"ResourceUsageContentTypeType"`
}

type resourceusagecontenttypetype struct {
        Code *AUCodeSetsResourceUsageContentTypeType `xml:"Code" json:"Code"`
      LocalDescription *String `xml:"LocalDescription,omitempty" json:"LocalDescription,omitempty"`
      
      }
    
    type ResourceReportColumnListType struct {
  resourcereportcolumnlisttype `xml:"ResourceReportColumnListType" json:"ResourceReportColumnListType"`
}

type resourcereportcolumnlisttype struct {
        ResourceReportColumn []ResourceReportColumnType `xml:"ResourceReportColumn" json:"ResourceReportColumn"`
      
      }
    
    type ResourceReportColumnType struct {
  resourcereportcolumntype `xml:"ResourceReportColumnType" json:"ResourceReportColumnType"`
}

type resourcereportcolumntype struct {
        ColumnName *String `xml:"ColumnName" json:"ColumnName"`
      ColumnDescription *String `xml:"ColumnDescription,omitempty" json:"ColumnDescription,omitempty"`
      ColumnDelimiter *String `xml:"ColumnDelimiter,omitempty" json:"ColumnDelimiter,omitempty"`
      
      }
    
    type ResourceReportLineListType struct {
  resourcereportlinelisttype `xml:"ResourceReportLineListType" json:"ResourceReportLineListType"`
}

type resourcereportlinelisttype struct {
        ResourceReportLine []ResourceReportLineType `xml:"ResourceReportLine" json:"ResourceReportLine"`
      
      }
    
    type ReportUserRefIdType struct {
  reportuserrefidtype `xml:"ReportUserRefIdType" json:"ReportUserRefIdType"`
}

type reportuserrefidtype struct {
          SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type ResourceReportLineType struct {
  resourcereportlinetype `xml:"ResourceReportLineType" json:"ResourceReportLineType"`
}

type resourcereportlinetype struct {
        SIF_RefId *ReportUserRefIdType `xml:"SIF_RefId,omitempty" json:"SIF_RefId,omitempty"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      CurrentCost *MonetaryAmountType `xml:"CurrentCost" json:"CurrentCost"`
      ReportRow *String `xml:"ReportRow" json:"ReportRow"`
      
      }
    
    type HomeroomType struct {
  homeroomtype `xml:"HomeroomType" json:"HomeroomType"`
}

type homeroomtype struct {
          SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type StaffRefIdType struct {
  staffrefidtype `xml:"StaffRefIdType" json:"StaffRefIdType"`
}

type staffrefidtype struct {
          SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type SystemRoleObjectRefIdType struct {
  systemroleobjectrefidtype `xml:"SystemRoleObjectRefIdType" json:"SystemRoleObjectRefIdType"`
}

type systemroleobjectrefidtype struct {
          SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type SystemContextListType struct {
  systemcontextlisttype `xml:"SystemContextListType" json:"SystemContextListType"`
}

type systemcontextlisttype struct {
        SystemContext []SystemContextType `xml:"SystemContext" json:"SystemContext"`
      
      }
    
    type SystemContextType struct {
  systemcontexttype `xml:"SystemContextType" json:"SystemContextType"`
}

type systemcontexttype struct {
        SystemId *String `xml:"SystemId,attr" json:"SystemId"`
      RoleList *SystemRoleListType `xml:"RoleList,omitempty" json:"RoleList,omitempty"`
      
      }
    
    type SystemRoleListType struct {
  systemrolelisttype `xml:"SystemRoleListType" json:"SystemRoleListType"`
}

type systemrolelisttype struct {
        Role []SystemRoleElementType `xml:"Role" json:"Role"`
      
      }
    
    type SystemRoleElementType struct {
  systemroleelementtype `xml:"SystemRoleElementType" json:"SystemRoleElementType"`
}

type systemroleelementtype struct {
        RoleId *String `xml:"RoleId,attr" json:"RoleId"`
      RoleScopeList *SystemRoleScopeListType `xml:"RoleScopeList,omitempty" json:"RoleScopeList,omitempty"`
      
      }
    
    type SystemRoleScopeListType struct {
  systemrolescopelisttype `xml:"SystemRoleScopeListType" json:"SystemRoleScopeListType"`
}

type systemrolescopelisttype struct {
        RoleScope []SystemRoleScopeType `xml:"RoleScope" json:"RoleScope"`
      
      }
    
    type SystemRoleScopeRefIdType struct {
  systemrolescoperefidtype `xml:"SystemRoleScopeRefIdType" json:"SystemRoleScopeRefIdType"`
}

type systemrolescoperefidtype struct {
          SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type SystemRoleScopeType struct {
  systemrolescopetype `xml:"SystemRoleScopeType" json:"SystemRoleScopeType"`
}

type systemrolescopetype struct {
        RoleScopeName *String `xml:"RoleScopeName,omitempty" json:"RoleScopeName,omitempty"`
      RoleScopeRefId *SystemRoleScopeRefIdType `xml:"RoleScopeRefId,omitempty" json:"RoleScopeRefId,omitempty"`
      
      }
    
    type TravelDetailsContainerType struct {
  traveldetailscontainertype `xml:"TravelDetailsContainerType" json:"TravelDetailsContainerType"`
}

type traveldetailscontainertype struct {
        ToSchool *SchoolTravelType `xml:"ToSchool,omitempty" json:"ToSchool,omitempty"`
      FromSchool *SchoolTravelType `xml:"FromSchool,omitempty" json:"FromSchool,omitempty"`
      
      }
    
    type SchoolTravelType struct {
  schooltraveltype `xml:"SchoolTravelType" json:"SchoolTravelType"`
}

type schooltraveltype struct {
        TravelMode *AUCodeSetsTravelModeType `xml:"TravelMode,omitempty" json:"TravelMode,omitempty"`
      TravelDetails *String `xml:"TravelDetails,omitempty" json:"TravelDetails,omitempty"`
      TravelAccompaniment *AUCodeSetsAccompanimentType `xml:"TravelAccompaniment,omitempty" json:"TravelAccompaniment,omitempty"`
      
      }
    
    type LibraryTransactionListType struct {
  librarytransactionlisttype `xml:"LibraryTransactionListType" json:"LibraryTransactionListType"`
}

type librarytransactionlisttype struct {
        Transaction []LibraryTransactionType `xml:"Transaction" json:"Transaction"`
      
      }
    
    type LibraryTransactionType struct {
  librarytransactiontype `xml:"LibraryTransactionType" json:"LibraryTransactionType"`
}

type librarytransactiontype struct {
        ItemInfo *LibraryItemInfoType `xml:"ItemInfo,omitempty" json:"ItemInfo,omitempty"`
      CheckoutInfo *CheckoutInfoType `xml:"CheckoutInfo,omitempty" json:"CheckoutInfo,omitempty"`
      FineInfoList *FineInfoListType `xml:"FineInfoList,omitempty" json:"FineInfoList,omitempty"`
      HoldInfoList *HoldInfoListType `xml:"HoldInfoList,omitempty" json:"HoldInfoList,omitempty"`
      
      }
    
    type LibraryItemInfoType struct {
  libraryiteminfotype `xml:"LibraryItemInfoType" json:"LibraryItemInfoType"`
}

type libraryiteminfotype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      Title *String `xml:"Title" json:"Title"`
      Author *String `xml:"Author,omitempty" json:"Author,omitempty"`
      ElectronicId *ElectronicIdType `xml:"ElectronicId,omitempty" json:"ElectronicId,omitempty"`
      CallNumber *String `xml:"CallNumber,omitempty" json:"CallNumber,omitempty"`
      ISBN *String `xml:"ISBN,omitempty" json:"ISBN,omitempty"`
      Cost *MonetaryAmountType `xml:"Cost,omitempty" json:"Cost,omitempty"`
      ReplacementCost *MonetaryAmountType `xml:"ReplacementCost,omitempty" json:"ReplacementCost,omitempty"`
      
      }
    
    type CheckoutInfoType struct {
  checkoutinfotype `xml:"CheckoutInfoType" json:"CheckoutInfoType"`
}

type checkoutinfotype struct {
        CheckedOutOn *String `xml:"CheckedOutOn" json:"CheckedOutOn"`
      ReturnBy *String `xml:"ReturnBy" json:"ReturnBy"`
      RenewalCount *Int `xml:"RenewalCount,omitempty" json:"RenewalCount,omitempty"`
      
      }
    
    type FineInfoListType struct {
  fineinfolisttype `xml:"FineInfoListType" json:"FineInfoListType"`
}

type fineinfolisttype struct {
        FineInfo []FineInfoType `xml:"FineInfo" json:"FineInfo"`
      
      }
    
    type FineInfoType struct {
  fineinfotype `xml:"FineInfoType" json:"FineInfoType"`
}

type fineinfotype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      Assessed *String `xml:"Assessed" json:"Assessed"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      Reference *String `xml:"Reference,omitempty" json:"Reference,omitempty"`
      
      }
    
    type HoldInfoListType struct {
  holdinfolisttype `xml:"HoldInfoListType" json:"HoldInfoListType"`
}

type holdinfolisttype struct {
        HoldInfo []HoldInfoType `xml:"HoldInfo" json:"HoldInfo"`
      
      }
    
    type HoldInfoType struct {
  holdinfotype `xml:"HoldInfoType" json:"HoldInfoType"`
}

type holdinfotype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      DatePlaced *String `xml:"DatePlaced" json:"DatePlaced"`
      DateNeeded *String `xml:"DateNeeded,omitempty" json:"DateNeeded,omitempty"`
      ReservationExpiry *String `xml:"ReservationExpiry,omitempty" json:"ReservationExpiry,omitempty"`
      MadeAvailable *String `xml:"MadeAvailable,omitempty" json:"MadeAvailable,omitempty"`
      Expires *String `xml:"Expires,omitempty" json:"Expires,omitempty"`
      
      }
    
    type LibraryMessageListType struct {
  librarymessagelisttype `xml:"LibraryMessageListType" json:"LibraryMessageListType"`
}

type librarymessagelisttype struct {
        Message []LibraryMessageType `xml:"Message" json:"Message"`
      
      }
    
    type LibraryMessageType struct {
  librarymessagetype `xml:"LibraryMessageType" json:"LibraryMessageType"`
}

type librarymessagetype struct {
        Priority *String `xml:"Priority,attr" json:"Priority"`
      PriorityCodeset *String `xml:"PriorityCodeset,attr" json:"PriorityCodeset"`
      Sent *String `xml:"Sent,omitempty" json:"Sent,omitempty"`
      Text *String `xml:"Text" json:"Text"`
      
      }
    
    type StudentAttendanceCollectionReportingListType struct {
  studentattendancecollectionreportinglisttype `xml:"StudentAttendanceCollectionReportingListType" json:"StudentAttendanceCollectionReportingListType"`
}

type studentattendancecollectionreportinglisttype struct {
        StudentAttendanceCollectionReporting []StudentAttendanceCollectionReportingType `xml:"StudentAttendanceCollectionReporting" json:"StudentAttendanceCollectionReporting"`
      
      }
    
    type StudentAttendanceCollectionReportingType struct {
  studentattendancecollectionreportingtype `xml:"StudentAttendanceCollectionReportingType" json:"StudentAttendanceCollectionReportingType"`
}

type studentattendancecollectionreportingtype struct {
        CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName,omitempty"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      StatsCohortYearLevelList *StatsCohortYearLevelListType `xml:"StatsCohortYearLevelList,omitempty" json:"StatsCohortYearLevelList,omitempty"`
      
      }
    
    type StatsCohortYearLevelListType struct {
  statscohortyearlevellisttype `xml:"StatsCohortYearLevelListType" json:"StatsCohortYearLevelListType"`
}

type statscohortyearlevellisttype struct {
        StatsCohortYearLevel []StatsCohortYearLevelType `xml:"StatsCohortYearLevel" json:"StatsCohortYearLevel"`
      
      }
    
    type StatsCohortYearLevelType struct {
  statscohortyearleveltype `xml:"StatsCohortYearLevelType" json:"StatsCohortYearLevelType"`
}

type statscohortyearleveltype struct {
        CohortYearLevel *YearLevelType `xml:"CohortYearLevel" json:"CohortYearLevel"`
      StatsCohortList *StatsCohortListType `xml:"StatsCohortList" json:"StatsCohortList"`
      
      }
    
    type StatsCohortListType struct {
  statscohortlisttype `xml:"StatsCohortListType" json:"StatsCohortListType"`
}

type statscohortlisttype struct {
        StatsCohort []StatsCohortType `xml:"StatsCohort" json:"StatsCohort"`
      
      }
    
    type StatsCohortType struct {
  statscohorttype `xml:"StatsCohortType" json:"StatsCohortType"`
}

type statscohorttype struct {
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
  addresscollectionreportinglisttype `xml:"AddressCollectionReportingListType" json:"AddressCollectionReportingListType"`
}

type addresscollectionreportinglisttype struct {
        AddressCollectionReporting []AddressCollectionReportingType `xml:"AddressCollectionReporting" json:"AddressCollectionReporting"`
      
      }
    
    type AddressCollectionReportingType struct {
  addresscollectionreportingtype `xml:"AddressCollectionReportingType" json:"AddressCollectionReportingType"`
}

type addresscollectionreportingtype struct {
        CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName,omitempty"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      AGContextualQuestionList *AGContextualQuestionListType `xml:"AGContextualQuestionList,omitempty" json:"AGContextualQuestionList,omitempty"`
      AddressCollectionStudentList *AddressCollectionStudentListType `xml:"AddressCollectionStudentList,omitempty" json:"AddressCollectionStudentList,omitempty"`
      
      }
    
    type AddressCollectionStudentListType struct {
  addresscollectionstudentlisttype `xml:"AddressCollectionStudentListType" json:"AddressCollectionStudentListType"`
}

type addresscollectionstudentlisttype struct {
        AddressCollectionStudent []AddressCollectionStudentType `xml:"AddressCollectionStudent" json:"AddressCollectionStudent"`
      
      }
    
    type AddressCollectionStudentType struct {
  addresscollectionstudenttype `xml:"AddressCollectionStudentType" json:"AddressCollectionStudentType"`
}

type addresscollectionstudenttype struct {
        LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      EducationLevel *AUCodeSetsEducationLevelType `xml:"EducationLevel" json:"EducationLevel"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus,omitempty"`
      ReportingParent2 *String `xml:"ReportingParent2" json:"ReportingParent2"`
      StudentAddress *AddressType `xml:"StudentAddress" json:"StudentAddress"`
      Parent1 *AGParentType `xml:"Parent1" json:"Parent1"`
      Parent2 *AGParentType `xml:"Parent2,omitempty" json:"Parent2,omitempty"`
      
      }
    
    type AGParentType struct {
  agparenttype `xml:"AGParentType" json:"AGParentType"`
}

type agparenttype struct {
        ParentName *NameOfRecordType `xml:"ParentName" json:"ParentName"`
      AddressSameAsStudent *AUCodeSetsYesOrNoCategoryType `xml:"AddressSameAsStudent" json:"AddressSameAsStudent"`
      ParentAddress *AddressType `xml:"ParentAddress" json:"ParentAddress"`
      
      }
    
    type AGRoundListType struct {
  agroundlisttype `xml:"AGRoundListType" json:"AGRoundListType"`
}

type agroundlisttype struct {
        AGRound []AGRoundType `xml:"AGRound" json:"AGRound"`
      
      }
    
    type AGRoundType struct {
  agroundtype `xml:"AGRoundType" json:"AGRoundType"`
}

type agroundtype struct {
        RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      RoundName *String `xml:"RoundName" json:"RoundName"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      DueDate *String `xml:"DueDate" json:"DueDate"`
      EndDate *String `xml:"EndDate" json:"EndDate"`
      
      }
    
    type AGContextualQuestionListType struct {
  agcontextualquestionlisttype `xml:"AGContextualQuestionListType" json:"AGContextualQuestionListType"`
}

type agcontextualquestionlisttype struct {
        AGContextualQuestion []AGContextualQuestionType `xml:"AGContextualQuestion" json:"AGContextualQuestion"`
      
      }
    
    type AGContextualQuestionType struct {
  agcontextualquestiontype `xml:"AGContextualQuestionType" json:"AGContextualQuestionType"`
}

type agcontextualquestiontype struct {
        AGContextCode *AUCodeSetsAGContextQuestionType `xml:"AGContextCode" json:"AGContextCode"`
      AGAnswer *String `xml:"AGAnswer" json:"AGAnswer"`
      
      }
    
    type CensusReportingListType struct {
  censusreportinglisttype `xml:"CensusReportingListType" json:"CensusReportingListType"`
}

type censusreportinglisttype struct {
        CensusReporting []CensusReportingType `xml:"CensusReporting" json:"CensusReporting"`
      
      }
    
    type CensusReportingType struct {
  censusreportingtype `xml:"CensusReportingType" json:"CensusReportingType"`
}

type censusreportingtype struct {
        EntityLevel *String `xml:"EntityLevel,omitempty" json:"EntityLevel,omitempty"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName,omitempty"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      CensusStaffList *CensusStaffListType `xml:"CensusStaffList,omitempty" json:"CensusStaffList,omitempty"`
      CensusStudentList *CensusStudentListType `xml:"CensusStudentList,omitempty" json:"CensusStudentList,omitempty"`
      
      }
    
    type CensusStaffListType struct {
  censusstafflisttype `xml:"CensusStaffListType" json:"CensusStaffListType"`
}

type censusstafflisttype struct {
        CensusStaff []CensusStaffType `xml:"CensusStaff" json:"CensusStaff"`
      
      }
    
    type CensusStaffType struct {
  censusstafftype `xml:"CensusStaffType" json:"CensusStaffType"`
}

type censusstafftype struct {
        StaffCohortId *LocalIdType `xml:"StaffCohortId" json:"StaffCohortId"`
      StaffActivity *StaffActivityExtensionType `xml:"StaffActivity" json:"StaffActivity"`
      CohortGender *String `xml:"CohortGender" json:"CohortGender"`
      CohortIndigenousType *String `xml:"CohortIndigenousType" json:"CohortIndigenousType"`
      PrimaryFTE *Float `xml:"PrimaryFTE,omitempty" json:"PrimaryFTE,omitempty"`
      SecondaryFTE *Float `xml:"SecondaryFTE,omitempty" json:"SecondaryFTE,omitempty"`
      JobFTE *Float `xml:"JobFTE,omitempty" json:"JobFTE,omitempty"`
      Headcount *Int `xml:"Headcount" json:"Headcount"`
      
      }
    
    type StaffAssignmentMostRecentContainerType struct {
  staffassignmentmostrecentcontainertype `xml:"StaffAssignmentMostRecentContainerType" json:"StaffAssignmentMostRecentContainerType"`
}

type staffassignmentmostrecentcontainertype struct {
        PrimaryFTE *Float `xml:"PrimaryFTE,omitempty" json:"PrimaryFTE,omitempty"`
      SecondaryFTE *Float `xml:"SecondaryFTE,omitempty" json:"SecondaryFTE,omitempty"`
      
      }
    
    type CensusStudentListType struct {
  censusstudentlisttype `xml:"CensusStudentListType" json:"CensusStudentListType"`
}

type censusstudentlisttype struct {
        CensusStudent []CensusStudentType `xml:"CensusStudent" json:"CensusStudent"`
      
      }
    
    type CensusStudentType struct {
  censusstudenttype `xml:"CensusStudentType" json:"CensusStudentType"`
}

type censusstudenttype struct {
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
  agreportingobjectresponselisttype `xml:"AGReportingObjectResponseListType" json:"AGReportingObjectResponseListType"`
}

type agreportingobjectresponselisttype struct {
        AGReportingObjectResponse []AGReportingObjectResponseType `xml:"AGReportingObjectResponse" json:"AGReportingObjectResponse"`
      
      }
    
    type AGReportingObjectResponseType struct {
  agreportingobjectresponsetype `xml:"AGReportingObjectResponseType" json:"AGReportingObjectResponseType"`
}

type agreportingobjectresponsetype struct {
        SubmittedRefId *String `xml:"SubmittedRefId,omitempty" json:"SubmittedRefId,omitempty"`
      SIFRefId *String `xml:"SIFRefId,omitempty" json:"SIFRefId,omitempty"`
      HTTPStatusCode *String `xml:"HTTPStatusCode,omitempty" json:"HTTPStatusCode,omitempty"`
      ErrorText *String `xml:"ErrorText,omitempty" json:"ErrorText,omitempty"`
      CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName,omitempty"`
      AGSubmissionStatusCode *AUCodeSetsAGSubmissionStatusType `xml:"AGSubmissionStatusCode" json:"AGSubmissionStatusCode"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList,omitempty"`
      
      }
    
    type FQReportingListType struct {
  fqreportinglisttype `xml:"FQReportingListType" json:"FQReportingListType"`
}

type fqreportinglisttype struct {
        FQReporting []FQReportingType `xml:"FQReporting" json:"FQReporting"`
      
      }
    
    type FQReportingType struct {
  fqreportingtype `xml:"FQReportingType" json:"FQReportingType"`
}

type fqreportingtype struct {
        CommonwealthId *String `xml:"CommonwealthId" json:"CommonwealthId"`
      EntityName *String `xml:"EntityName,omitempty" json:"EntityName,omitempty"`
      EntityContact *EntityContactInfoType `xml:"EntityContact" json:"EntityContact"`
      FQContextualQuestionList *FQContextualQuestionListType `xml:"FQContextualQuestionList,omitempty" json:"FQContextualQuestionList,omitempty"`
      FQItemList *FQItemListType `xml:"FQItemList" json:"FQItemList"`
      AGRuleList *AGRuleListType `xml:"AGRuleList,omitempty" json:"AGRuleList,omitempty"`
      
      }
    
    type FQContextualQuestionListType struct {
  fqcontextualquestionlisttype `xml:"FQContextualQuestionListType" json:"FQContextualQuestionListType"`
}

type fqcontextualquestionlisttype struct {
        FQContextualQuestion []FQContextualQuestionType `xml:"FQContextualQuestion" json:"FQContextualQuestion"`
      
      }
    
    type FQContextualQuestionType struct {
  fqcontextualquestiontype `xml:"FQContextualQuestionType" json:"FQContextualQuestionType"`
}

type fqcontextualquestiontype struct {
        FQContext *String `xml:"FQContext" json:"FQContext"`
      FQAnswer *String `xml:"FQAnswer" json:"FQAnswer"`
      
      }
    
    type FQItemListType struct {
  fqitemlisttype `xml:"FQItemListType" json:"FQItemListType"`
}

type fqitemlisttype struct {
        FQItem []FQItemType `xml:"FQItem" json:"FQItem"`
      
      }
    
    type FQItemType struct {
  fqitemtype `xml:"FQItemType" json:"FQItemType"`
}

type fqitemtype struct {
        FQItemCode *String `xml:"FQItemCode" json:"FQItemCode"`
      TuitionAmount *Float `xml:"TuitionAmount,omitempty" json:"TuitionAmount,omitempty"`
      BoardingAmount *Float `xml:"BoardingAmount,omitempty" json:"BoardingAmount,omitempty"`
      SystemAmount *Float `xml:"SystemAmount,omitempty" json:"SystemAmount,omitempty"`
      DioceseAmount *Float `xml:"DioceseAmount,omitempty" json:"DioceseAmount,omitempty"`
      FQComments *String `xml:"FQComments,omitempty" json:"FQComments,omitempty"`
      
      }
    
    type AGRuleListType struct {
  agrulelisttype `xml:"AGRuleListType" json:"AGRuleListType"`
}

type agrulelisttype struct {
        AGRule []AGRuleType `xml:"AGRule" json:"AGRule"`
      
      }
    
    type AGRuleType struct {
  agruletype `xml:"AGRuleType" json:"AGRuleType"`
}

type agruletype struct {
        AGRuleCode *String `xml:"AGRuleCode,omitempty" json:"AGRuleCode,omitempty"`
      AGRuleComment *String `xml:"AGRuleComment,omitempty" json:"AGRuleComment,omitempty"`
      AGRuleResponse *String `xml:"AGRuleResponse,omitempty" json:"AGRuleResponse,omitempty"`
      AGRuleStatus *String `xml:"AGRuleStatus,omitempty" json:"AGRuleStatus,omitempty"`
      
      }
    
    type SoftwareVendorInfoContainerType struct {
  softwarevendorinfocontainertype `xml:"SoftwareVendorInfoContainerType" json:"SoftwareVendorInfoContainerType"`
}

type softwarevendorinfocontainertype struct {
        SoftwareProduct *String `xml:"SoftwareProduct" json:"SoftwareProduct"`
      SoftwareVersion *String `xml:"SoftwareVersion" json:"SoftwareVersion"`
      
      }
    
    type TimeTableScheduleType struct {
  timetablescheduletype `xml:"TimeTableScheduleType" json:"TimeTableScheduleType"`
}

type timetablescheduletype struct {
        SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      Title *String `xml:"Title" json:"Title"`
      DaysPerCycle *Int `xml:"DaysPerCycle" json:"DaysPerCycle"`
      PeriodsPerDay *Int `xml:"PeriodsPerDay" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *Int `xml:"TeachingPeriodsPerDay,omitempty" json:"TeachingPeriodsPerDay,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolName *String `xml:"SchoolName,omitempty" json:"SchoolName,omitempty"`
      TimeTableCreationDate *String `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate,omitempty"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList" json:"TimeTableDayList"`
      
      }
    
    type TimeTableScheduleCellListType struct {
  timetableschedulecelllisttype `xml:"TimeTableScheduleCellListType" json:"TimeTableScheduleCellListType"`
}

type timetableschedulecelllisttype struct {
        TimeTableScheduleCell []TimeTableScheduleCellType `xml:"TimeTableScheduleCell" json:"TimeTableScheduleCell"`
      
      }
    
    type TimeTableScheduleCellType struct {
  timetableschedulecelltype `xml:"TimeTableScheduleCellType" json:"TimeTableScheduleCellType"`
}

type timetableschedulecelltype struct {
        TimeTableScheduleCellLocalId *LocalIdType `xml:"TimeTableScheduleCellLocalId" json:"TimeTableScheduleCellLocalId"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      TeachingGroupGUID *String `xml:"TeachingGroupGUID" json:"TeachingGroupGUID"`
      RoomInfoRefId *String `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId,omitempty"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      TimeTableLocalId *LocalIdType `xml:"TimeTableLocalId,omitempty" json:"TimeTableLocalId,omitempty"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId,omitempty"`
      TeachingGroupLocalId *LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId,omitempty"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber,omitempty"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId,omitempty"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      CellType *String `xml:"CellType" json:"CellType"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList,omitempty"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList,omitempty"`
      
      }
    
    type TeachingGroupScheduleListType struct {
  teachinggroupschedulelisttype `xml:"TeachingGroupScheduleListType" json:"TeachingGroupScheduleListType"`
}

type teachinggroupschedulelisttype struct {
        TeachingGroupSchedule []TeachingGroupScheduleType `xml:"TeachingGroupSchedule" json:"TeachingGroupSchedule"`
      
      }
    
    type TeachingGroupScheduleType struct {
  teachinggroupscheduletype `xml:"TeachingGroupScheduleType" json:"TeachingGroupScheduleType"`
}

type teachinggroupscheduletype struct {
        EditorGUID *RefIdType `xml:"EditorGUID" json:"EditorGUID"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      ShortName *String `xml:"ShortName" json:"ShortName"`
      LongName *String `xml:"LongName,omitempty" json:"LongName,omitempty"`
      GroupType *String `xml:"GroupType,omitempty" json:"GroupType,omitempty"`
      Set *String `xml:"Set,omitempty" json:"Set,omitempty"`
      Block *String `xml:"Block,omitempty" json:"Block,omitempty"`
      CurriculumLevel *String `xml:"CurriculumLevel,omitempty" json:"CurriculumLevel,omitempty"`
      SchoolInfoRefId *RefIdType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolCourseInfoRefId *RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId,omitempty"`
      SchoolCourseLocalId *LocalIdType `xml:"SchoolCourseLocalId,omitempty" json:"SchoolCourseLocalId,omitempty"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      TimeTableSubjectLocalId *LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId,omitempty"`
      Semester *Int `xml:"Semester,omitempty" json:"Semester,omitempty"`
      StudentList *StudentListType `xml:"StudentList,omitempty" json:"StudentList,omitempty"`
      TeacherList *TeacherListType `xml:"TeacherList,omitempty" json:"TeacherList,omitempty"`
      MinClassSize *Int `xml:"MinClassSize,omitempty" json:"MinClassSize,omitempty"`
      MaxClassSize *Int `xml:"MaxClassSize,omitempty" json:"MaxClassSize,omitempty"`
      TeachingGroupPeriodList *TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList,omitempty" json:"TeachingGroupPeriodList,omitempty"`
      
      }
    
    type LocalCodeListType struct {
  localcodelisttype `xml:"LocalCodeListType" json:"LocalCodeListType"`
}

type localcodelisttype struct {
        LocalCode []LocalCodeType `xml:"LocalCode" json:"LocalCode"`
      
      }
    
    type LocalCodeType struct {
  localcodetype `xml:"LocalCodeType" json:"LocalCodeType"`
}

type localcodetype struct {
        LocalisedCode *String `xml:"LocalisedCode" json:"LocalisedCode"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      Element *String `xml:"Element,omitempty" json:"Element,omitempty"`
      ListIndex *Int `xml:"ListIndex,omitempty" json:"ListIndex,omitempty"`
      
      }
    
    type StudentGroupListType struct {
  studentgrouplisttype `xml:"StudentGroupListType" json:"StudentGroupListType"`
}

type studentgrouplisttype struct {
        StudentGroup []StudentGroupType `xml:"StudentGroup" json:"StudentGroup"`
      
      }
    
    type StudentGroupType struct {
  studentgrouptype `xml:"StudentGroupType" json:"StudentGroupType"`
}

type studentgrouptype struct {
        GroupCategory *AUCodeSetsGroupCategoryCodeType `xml:"GroupCategory" json:"GroupCategory"`
      GroupLocalId *LocalIdType `xml:"GroupLocalId" json:"GroupLocalId"`
      GroupDescription *String `xml:"GroupDescription,omitempty" json:"GroupDescription,omitempty"`
      
      }
    
    type PublishingPermissionListType struct {
  publishingpermissionlisttype `xml:"PublishingPermissionListType" json:"PublishingPermissionListType"`
}

type publishingpermissionlisttype struct {
        PublishingPermission []PublishingPermissionType `xml:"PublishingPermission" json:"PublishingPermission"`
      
      }
    
    type PublishingPermissionType struct {
  publishingpermissiontype `xml:"PublishingPermissionType" json:"PublishingPermissionType"`
}

type publishingpermissiontype struct {
        PermissionCategory *AUCodeSetsPermissionCategoryCodeType `xml:"PermissionCategory" json:"PermissionCategory"`
      PermissionValue *AUCodeSetsYesOrNoCategoryType `xml:"PermissionValue" json:"PermissionValue"`
      
      }
    
    type EntityContactInfoType struct {
  entitycontactinfotype `xml:"EntityContactInfoType" json:"EntityContactInfoType"`
}

type entitycontactinfotype struct {
        Name *NameType `xml:"Name" json:"Name"`
      PositionTitle *String `xml:"PositionTitle,omitempty" json:"PositionTitle,omitempty"`
      Role *String `xml:"Role,omitempty" json:"Role,omitempty"`
      RegistrationDetails *String `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails,omitempty"`
      Qualifications *String `xml:"Qualifications,omitempty" json:"Qualifications,omitempty"`
      Address *AddressType `xml:"Address,omitempty" json:"Address,omitempty"`
      Email *EmailType `xml:"Email" json:"Email"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type CopyRightContainerType struct {
  copyrightcontainertype `xml:"CopyRightContainerType" json:"CopyRightContainerType"`
}

type copyrightcontainertype struct {
        Date *String `xml:"Date,omitempty" json:"Date,omitempty"`
      Holder *String `xml:"Holder,omitempty" json:"Holder,omitempty"`
      
      }
    
    type StandardsSettingBodyType struct {
  standardssettingbodytype `xml:"StandardsSettingBodyType" json:"StandardsSettingBodyType"`
}

type standardssettingbodytype struct {
        Country *CountryType `xml:"Country,omitempty" json:"Country,omitempty"`
      StateProvince *StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince,omitempty"`
      SettingBodyName *String `xml:"SettingBodyName,omitempty" json:"SettingBodyName,omitempty"`
      
      }
    
    type StandardHierarchyLevelType struct {
  standardhierarchyleveltype `xml:"StandardHierarchyLevelType" json:"StandardHierarchyLevelType"`
}

type standardhierarchyleveltype struct {
        Number *Int `xml:"Number" json:"Number"`
      Description *String `xml:"Description" json:"Description"`
      
      }
    
    type StandardIdentifierType struct {
  standardidentifiertype `xml:"StandardIdentifierType" json:"StandardIdentifierType"`
}

type standardidentifiertype struct {
        YearCreated *String `xml:"YearCreated" json:"YearCreated"`
      ACStrandSubjectArea *ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      StandardNumber *String `xml:"StandardNumber" json:"StandardNumber"`
      YearLevels *YearLevelsType `xml:"YearLevels" json:"YearLevels"`
      Benchmark *String `xml:"Benchmark,omitempty" json:"Benchmark,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel,omitempty"`
      IndicatorNumber *String `xml:"IndicatorNumber,omitempty" json:"IndicatorNumber,omitempty"`
      AlternateIdentificationCodes *AlternateIdentificationCodeListType `xml:"AlternateIdentificationCodes,omitempty" json:"AlternateIdentificationCodes,omitempty"`
      Organization *String `xml:"Organization" json:"Organization"`
      
      }
    
    type AlternateIdentificationCodeListType struct {
  alternateidentificationcodelisttype `xml:"AlternateIdentificationCodeListType" json:"AlternateIdentificationCodeListType"`
}

type alternateidentificationcodelisttype struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode" json:"AlternateIdentificationCode"`
      
      }
    
    type RelatedLearningStandardItemRefIdListType struct {
  relatedlearningstandarditemrefidlisttype `xml:"RelatedLearningStandardItemRefIdListType" json:"RelatedLearningStandardItemRefIdListType"`
}

type relatedlearningstandarditemrefidlisttype struct {
        LearningStandardItemRefId []RelatedLearningStandardItemRefIdType `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      
      }
    
    type RelatedLearningStandardItemRefIdType struct {
  relatedlearningstandarditemrefidtype `xml:"RelatedLearningStandardItemRefIdType" json:"RelatedLearningStandardItemRefIdType"`
}

type relatedlearningstandarditemrefidtype struct {
          RelationshipType *String `xml:"RelationshipType,attr" json:"RelationshipType"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type ValidLetterMarkListType struct {
  validlettermarklisttype `xml:"ValidLetterMarkListType" json:"ValidLetterMarkListType"`
}

type validlettermarklisttype struct {
        ValidLetterMark []ValidLetterMarkType `xml:"ValidLetterMark" json:"ValidLetterMark"`
      
      }
    
    type ValidLetterMarkType struct {
  validlettermarktype `xml:"ValidLetterMarkType" json:"ValidLetterMarkType"`
}

type validlettermarktype struct {
        Code *String `xml:"Code" json:"Code"`
      NumericEquivalent *Float `xml:"NumericEquivalent,omitempty" json:"NumericEquivalent,omitempty"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      
      }
    
    type StudentGradeMarkersListType struct {
  studentgrademarkerslisttype `xml:"StudentGradeMarkersListType" json:"StudentGradeMarkersListType"`
}

type studentgrademarkerslisttype struct {
        Marker []MarkerType `xml:"Marker" json:"Marker"`
      
      }
    
    type MarkerType struct {
  markertype `xml:"MarkerType" json:"MarkerType"`
}

type markertype struct {
        StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      Role *String `xml:"Role,omitempty" json:"Role,omitempty"`
      
      }
    
    type GradingScoreListType struct {
  gradingscorelisttype `xml:"GradingScoreListType" json:"GradingScoreListType"`
}

type gradingscorelisttype struct {
        GradingAssignmentScore []AssignmentScoreType `xml:"GradingAssignmentScore" json:"GradingAssignmentScore"`
      
      }
    
    type AssignmentScoreType struct {
  assignmentscoretype `xml:"AssignmentScoreType" json:"AssignmentScoreType"`
}

type assignmentscoretype struct {
        GradingAssignmentScoreRefId *String `xml:"GradingAssignmentScoreRefId,omitempty" json:"GradingAssignmentScoreRefId,omitempty"`
      Weight *Float `xml:"Weight,omitempty" json:"Weight,omitempty"`
      
      }
    
    type GradeType struct {
  gradetype `xml:"GradeType" json:"GradeType"`
}

type gradetype struct {
        Percentage *Float `xml:"Percentage,omitempty" json:"Percentage,omitempty"`
      Numeric *Float `xml:"Numeric,omitempty" json:"Numeric,omitempty"`
      Letter *String `xml:"Letter,omitempty" json:"Letter,omitempty"`
      Narrative *String `xml:"Narrative,omitempty" json:"Narrative,omitempty"`
      MarkInfoRefId *String `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId,omitempty"`
      
      }
    
    type LearningStandardListType struct {
  learningstandardlisttype `xml:"LearningStandardListType" json:"LearningStandardListType"`
}

type learningstandardlisttype struct {
        LearningStandard []LearningStandardType `xml:"LearningStandard" json:"LearningStandard"`
      
      }
    
    type LearningStandardType struct {
  learningstandardtype `xml:"LearningStandardType" json:"LearningStandardType"`
}

type learningstandardtype struct {
        LearningStandardItemRefId *String `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId,omitempty"`
      LearningStandardURL *String `xml:"LearningStandardURL,omitempty" json:"LearningStandardURL,omitempty"`
      LearningStandardLocalId *LocalIdType `xml:"LearningStandardLocalId,omitempty" json:"LearningStandardLocalId,omitempty"`
      
      }
    
    type AssignmentListType struct {
  assignmentlisttype `xml:"AssignmentListType" json:"AssignmentListType"`
}

type assignmentlisttype struct {
        GradingAssignmentRefId []string `xml:"GradingAssignmentRefId" json:"GradingAssignmentRefId"`
      
      }
    
    type GenericRubricType struct {
  genericrubrictype `xml:"GenericRubricType" json:"GenericRubricType"`
}

type genericrubrictype struct {
        RubricType *String `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor,omitempty"`
      
      }
    
    type SymptomListType struct {
  symptomlisttype `xml:"SymptomListType" json:"SymptomListType"`
}

type symptomlisttype struct {
        Symptom []string `xml:"Symptom" json:"Symptom"`
      
      }
    
    type MedicationListType struct {
  medicationlisttype `xml:"MedicationListType" json:"MedicationListType"`
}

type medicationlisttype struct {
        Medication []MedicationType `xml:"Medication" json:"Medication"`
      
      }
    
    type MedicationType struct {
  medicationtype `xml:"MedicationType" json:"MedicationType"`
}

type medicationtype struct {
        MedicationName *String `xml:"MedicationName,omitempty" json:"MedicationName,omitempty"`
      Dosage *String `xml:"Dosage,omitempty" json:"Dosage,omitempty"`
      Frequency *String `xml:"Frequency,omitempty" json:"Frequency,omitempty"`
      AdministrationInformation *String `xml:"AdministrationInformation,omitempty" json:"AdministrationInformation,omitempty"`
      Method *String `xml:"Method,omitempty" json:"Method,omitempty"`
      
      }
    
    type WellbeingEventCategoryListType struct {
  wellbeingeventcategorylisttype `xml:"WellbeingEventCategoryListType" json:"WellbeingEventCategoryListType"`
}

type wellbeingeventcategorylisttype struct {
        WellbeingEventCategory []WellbeingEventCategoryType `xml:"WellbeingEventCategory" json:"WellbeingEventCategory"`
      
      }
    
    type WellbeingEventCategoryType struct {
  wellbeingeventcategorytype `xml:"WellbeingEventCategoryType" json:"WellbeingEventCategoryType"`
}

type wellbeingeventcategorytype struct {
        EventCategory *String `xml:"EventCategory" json:"EventCategory"`
      WellbeingEventSubCategoryList *WellbeingEventSubCategoryListType `xml:"WellbeingEventSubCategoryList,omitempty" json:"WellbeingEventSubCategoryList,omitempty"`
      
      }
    
    type WellbeingEventSubCategoryListType struct {
  wellbeingeventsubcategorylisttype `xml:"WellbeingEventSubCategoryListType" json:"WellbeingEventSubCategoryListType"`
}

type wellbeingeventsubcategorylisttype struct {
        WellbeingEventSubCategory []string `xml:"WellbeingEventSubCategory" json:"WellbeingEventSubCategory"`
      
      }
    
    type WellbeingEventLocationDetailsType struct {
  wellbeingeventlocationdetailstype `xml:"WellbeingEventLocationDetailsType" json:"WellbeingEventLocationDetailsType"`
}

type wellbeingeventlocationdetailstype struct {
        EventLocation *AUCodeSetsWellbeingEventLocationType `xml:"EventLocation" json:"EventLocation"`
      Class *String `xml:"Class,omitempty" json:"Class,omitempty"`
      FurtherLocationNotes *String `xml:"FurtherLocationNotes,omitempty" json:"FurtherLocationNotes,omitempty"`
      
      }
    
    type FollowUpActionListType struct {
  followupactionlisttype `xml:"FollowUpActionListType" json:"FollowUpActionListType"`
}

type followupactionlisttype struct {
        FollowUpAction []FollowUpActionType `xml:"FollowUpAction" json:"FollowUpAction"`
      
      }
    
    type FollowUpActionType struct {
  followupactiontype `xml:"FollowUpActionType" json:"FollowUpActionType"`
}

type followupactiontype struct {
        WellbeingResponseRefId *String `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId,omitempty"`
      FollowUpDetails *String `xml:"FollowUpDetails,omitempty" json:"FollowUpDetails,omitempty"`
      FollowUpActionCategory *String `xml:"FollowUpActionCategory,omitempty" json:"FollowUpActionCategory,omitempty"`
      
      }
    
    type PersonInvolvementListType struct {
  personinvolvementlisttype `xml:"PersonInvolvementListType" json:"PersonInvolvementListType"`
}

type personinvolvementlisttype struct {
        PersonInvolvement []PersonInvolvementType `xml:"PersonInvolvement" json:"PersonInvolvement"`
      
      }
    
    type PersonInvolvementType struct {
  personinvolvementtype `xml:"PersonInvolvementType" json:"PersonInvolvementType"`
}

type personinvolvementtype struct {
      PersonRefId *PersonInvolvementType_PersonRefId
      ShortName *String `xml:"ShortName,omitempty" json:"ShortName,omitempty"`
      HowInvolved *String `xml:"HowInvolved,omitempty" json:"HowInvolved,omitempty"`
      
      }
    
    type WithdrawalTimeListType struct {
  withdrawaltimelisttype `xml:"WithdrawalTimeListType" json:"WithdrawalTimeListType"`
}

type withdrawaltimelisttype struct {
        Withdrawal []WithdrawalType `xml:"Withdrawal" json:"Withdrawal"`
      
      }
    
    type WithdrawalType struct {
  withdrawaltype `xml:"WithdrawalType" json:"WithdrawalType"`
}

type withdrawaltype struct {
        WithdrawalDate *String `xml:"WithdrawalDate" json:"WithdrawalDate"`
      WithdrawalStartTime *String `xml:"WithdrawalStartTime,omitempty" json:"WithdrawalStartTime,omitempty"`
      WithdrawalEndTime *String `xml:"WithdrawalEndTime,omitempty" json:"WithdrawalEndTime,omitempty"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId,omitempty"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId,omitempty"`
      
      }
    
    type SuspensionContainerType struct {
  suspensioncontainertype `xml:"SuspensionContainerType" json:"SuspensionContainerType"`
}

type suspensioncontainertype struct {
        SuspensionCategory *AUCodeSetsSuspensionCategoryType `xml:"SuspensionCategory" json:"SuspensionCategory"`
      SuspensionNotes *String `xml:"SuspensionNotes,omitempty" json:"SuspensionNotes,omitempty"`
      WithdrawalTimeList *WithdrawalTimeListType `xml:"WithdrawalTimeList,omitempty" json:"WithdrawalTimeList,omitempty"`
      Duration *Float `xml:"Duration,omitempty" json:"Duration,omitempty"`
      AdvisementDate *String `xml:"AdvisementDate,omitempty" json:"AdvisementDate,omitempty"`
      ResolutionMeetingTime *String `xml:"ResolutionMeetingTime,omitempty" json:"ResolutionMeetingTime,omitempty"`
      ResolutionNotes *String `xml:"ResolutionNotes,omitempty" json:"ResolutionNotes,omitempty"`
      EarlyReturnDate *String `xml:"EarlyReturnDate,omitempty" json:"EarlyReturnDate,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      
      }
    
    type DetentionContainerType struct {
  detentioncontainertype `xml:"DetentionContainerType" json:"DetentionContainerType"`
}

type detentioncontainertype struct {
        DetentionCategory *AUCodeSetsDetentionCategoryType `xml:"DetentionCategory" json:"DetentionCategory"`
      DetentionDate *String `xml:"DetentionDate,omitempty" json:"DetentionDate,omitempty"`
      DetentionLocation *String `xml:"DetentionLocation,omitempty" json:"DetentionLocation,omitempty"`
      DetentionNotes *String `xml:"DetentionNotes,omitempty" json:"DetentionNotes,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      
      }
    
    type PlanRequiredContainerType struct {
  planrequiredcontainertype `xml:"PlanRequiredContainerType" json:"PlanRequiredContainerType"`
}

type planrequiredcontainertype struct {
        PlanRequiredList *PlanRequiredListType `xml:"PlanRequiredList,omitempty" json:"PlanRequiredList,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      
      }
    
    type PlanRequiredListType struct {
  planrequiredlisttype `xml:"PlanRequiredListType" json:"PlanRequiredListType"`
}

type planrequiredlisttype struct {
        Plan []WellbeingPlanType `xml:"Plan" json:"Plan"`
      
      }
    
    type WellbeingPlanType struct {
  wellbeingplantype `xml:"WellbeingPlanType" json:"WellbeingPlanType"`
}

type wellbeingplantype struct {
        PersonalisedPlanRefId *String `xml:"PersonalisedPlanRefId,omitempty" json:"PersonalisedPlanRefId,omitempty"`
      PlanNotes *String `xml:"PlanNotes,omitempty" json:"PlanNotes,omitempty"`
      
      }
    
    type AwardContainerType struct {
  awardcontainertype `xml:"AwardContainerType" json:"AwardContainerType"`
}

type awardcontainertype struct {
        AwardDate *String `xml:"AwardDate,omitempty" json:"AwardDate,omitempty"`
      AwardType *String `xml:"AwardType,omitempty" json:"AwardType,omitempty"`
      AwardDescription *String `xml:"AwardDescription,omitempty" json:"AwardDescription,omitempty"`
      AwardNotes *String `xml:"AwardNotes,omitempty" json:"AwardNotes,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      
      }
    
    type OtherWellbeingResponseContainerType struct {
  otherwellbeingresponsecontainertype `xml:"OtherWellbeingResponseContainerType" json:"OtherWellbeingResponseContainerType"`
}

type otherwellbeingresponsecontainertype struct {
        OtherResponseDate *String `xml:"OtherResponseDate,omitempty" json:"OtherResponseDate,omitempty"`
      OtherResponseType *String `xml:"OtherResponseType,omitempty" json:"OtherResponseType,omitempty"`
      OtherResponseDescription *String `xml:"OtherResponseDescription,omitempty" json:"OtherResponseDescription,omitempty"`
      OtherResponseNotes *String `xml:"OtherResponseNotes,omitempty" json:"OtherResponseNotes,omitempty"`
      Status *AUCodeSetsWellbeingStatusType `xml:"Status,omitempty" json:"Status,omitempty"`
      
      }
    
    type WellbeingDocumentListType struct {
  wellbeingdocumentlisttype `xml:"WellbeingDocumentListType" json:"WellbeingDocumentListType"`
}

type wellbeingdocumentlisttype struct {
        Document []WellbeingDocumentType `xml:"Document" json:"Document"`
      
      }
    
    type WellbeingDocumentType struct {
  wellbeingdocumenttype `xml:"WellbeingDocumentType" json:"WellbeingDocumentType"`
}

type wellbeingdocumenttype struct {
        Location *String `xml:"Location" json:"Location"`
      Sensitivity *String `xml:"Sensitivity,omitempty" json:"Sensitivity,omitempty"`
      URL *String `xml:"URL,omitempty" json:"URL,omitempty"`
      DocumentType *String `xml:"DocumentType,omitempty" json:"DocumentType,omitempty"`
      DocumentReviewDate *String `xml:"DocumentReviewDate,omitempty" json:"DocumentReviewDate,omitempty"`
      DocumentDescription *String `xml:"DocumentDescription,omitempty" json:"DocumentDescription,omitempty"`
      
      }
    
    type NAPTestItemContentType struct {
  naptestitemcontenttype `xml:"NAPTestItemContentType" json:"NAPTestItemContentType"`
}

type naptestitemcontenttype struct {
        NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      ItemName *String `xml:"ItemName" json:"ItemName"`
      ItemType *AUCodeSetsNAPTestItemTypeType `xml:"ItemType" json:"ItemType"`
      Subdomain *String `xml:"Subdomain" json:"Subdomain"`
      WritingGenre *AUCodeSetsNAPWritingGenreType `xml:"WritingGenre,omitempty" json:"WritingGenre,omitempty"`
      ItemDescriptor *String `xml:"ItemDescriptor" json:"ItemDescriptor"`
      ReleasedStatus *Bool `xml:"ReleasedStatus" json:"ReleasedStatus"`
      MarkingType *AUCodeSetsNAPTestItemMarkingTypeType `xml:"MarkingType" json:"MarkingType"`
      MultipleChoiceOptionCount *Int `xml:"MultipleChoiceOptionCount,omitempty" json:"MultipleChoiceOptionCount,omitempty"`
      CorrectAnswer *String `xml:"CorrectAnswer,omitempty" json:"CorrectAnswer,omitempty"`
      MaximumScore *Float `xml:"MaximumScore" json:"MaximumScore"`
      ItemDifficulty *Float `xml:"ItemDifficulty" json:"ItemDifficulty"`
      ItemDifficultyLogit5 *Float `xml:"ItemDifficultyLogit5" json:"ItemDifficultyLogit5"`
      ItemDifficultyLogit62 *Float `xml:"ItemDifficultyLogit62" json:"ItemDifficultyLogit62"`
      ItemDifficultyLogit5SE *Float `xml:"ItemDifficultyLogit5SE" json:"ItemDifficultyLogit5SE"`
      ItemDifficultyLogit62SE *Float `xml:"ItemDifficultyLogit62SE" json:"ItemDifficultyLogit62SE"`
      ItemProficiencyBand *Int `xml:"ItemProficiencyBand" json:"ItemProficiencyBand"`
      ItemProficiencyLevel *String `xml:"ItemProficiencyLevel,omitempty" json:"ItemProficiencyLevel,omitempty"`
      ExemplarURL *String `xml:"ExemplarURL,omitempty" json:"ExemplarURL,omitempty"`
      ItemSubstitutedForList *SubstituteItemListType `xml:"ItemSubstitutedForList,omitempty" json:"ItemSubstitutedForList,omitempty"`
      ContentDescriptionList *ContentDescriptionListType `xml:"ContentDescriptionList,omitempty" json:"ContentDescriptionList,omitempty"`
      StimulusList *StimulusListType `xml:"StimulusList,omitempty" json:"StimulusList,omitempty"`
      NAPWritingRubricList *NAPWritingRubricListType `xml:"NAPWritingRubricList,omitempty" json:"NAPWritingRubricList,omitempty"`
      
      }
    
    type NAPTestletContentType struct {
  naptestletcontenttype `xml:"NAPTestletContentType" json:"NAPTestletContentType"`
}

type naptestletcontenttype struct {
        NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletName *String `xml:"TestletName,omitempty" json:"TestletName,omitempty"`
      Node *String `xml:"Node,omitempty" json:"Node,omitempty"`
      LocationInStage *Int `xml:"LocationInStage,omitempty" json:"LocationInStage,omitempty"`
      TestletMaximumScore *Float `xml:"TestletMaximumScore" json:"TestletMaximumScore"`
      
      }
    
    type NAPTestContentType struct {
  naptestcontenttype `xml:"NAPTestContentType" json:"NAPTestContentType"`
}

type naptestcontenttype struct {
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
  plausiblescaledvaluelisttype `xml:"PlausibleScaledValueListType" json:"PlausibleScaledValueListType"`
}

type plausiblescaledvaluelisttype struct {
        PlausibleScaledValue []float64 `xml:"PlausibleScaledValue" json:"PlausibleScaledValue"`
      
      }
    
    type SubstituteItemListType struct {
  substituteitemlisttype `xml:"SubstituteItemListType" json:"SubstituteItemListType"`
}

type substituteitemlisttype struct {
        SubstituteItem []SubstituteItemType `xml:"SubstituteItem" json:"SubstituteItem"`
      
      }
    
    type SubstituteItemType struct {
  substituteitemtype `xml:"SubstituteItemType" json:"SubstituteItemType"`
}

type substituteitemtype struct {
        SubstituteItemRefId *String `xml:"SubstituteItemRefId" json:"SubstituteItemRefId"`
      SubstituteItemLocalId *LocalIdType `xml:"SubstituteItemLocalId,omitempty" json:"SubstituteItemLocalId,omitempty"`
      PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      
      }
    
    type CodeFrameTestItemListType struct {
  codeframetestitemlisttype `xml:"CodeFrameTestItemListType" json:"CodeFrameTestItemListType"`
}

type codeframetestitemlisttype struct {
        TestItem []CodeFrameTestItemType `xml:"TestItem" json:"TestItem"`
      
      }
    
    type CodeFrameTestItemType struct {
  codeframetestitemtype `xml:"CodeFrameTestItemType" json:"CodeFrameTestItemType"`
}

type codeframetestitemtype struct {
        TestItemRefId *String `xml:"TestItemRefId" json:"TestItemRefId"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      TestItemContent *NAPTestItemContentType `xml:"TestItemContent" json:"TestItemContent"`
      
      }
    
    type StimulusLocalIdListType struct {
  stimuluslocalidlisttype `xml:"StimulusLocalIdListType" json:"StimulusLocalIdListType"`
}

type stimuluslocalidlisttype struct {
        StimulusLocalId []LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      
      }
    
    type DomainBandsContainerType struct {
  domainbandscontainertype `xml:"DomainBandsContainerType" json:"DomainBandsContainerType"`
}

type domainbandscontainertype struct {
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
  domainproficiencycontainertype `xml:"DomainProficiencyContainerType" json:"DomainProficiencyContainerType"`
}

type domainproficiencycontainertype struct {
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
  naptestitemlisttype `xml:"NAPTestItemListType" json:"NAPTestItemListType"`
}

type naptestitemlisttype struct {
        TestItem []NAPTestItem2Type `xml:"TestItem" json:"TestItem"`
      
      }
    
    type NAPTestItem2Type struct {
  naptestitem2type `xml:"NAPTestItem2Type" json:"NAPTestItem2Type"`
}

type naptestitem2type struct {
        TestItemRefId *String `xml:"TestItemRefId" json:"TestItemRefId"`
      TestItemLocalId *LocalIdType `xml:"TestItemLocalId" json:"TestItemLocalId"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      
      }
    
    type NAPCodeFrameTestletListType struct {
  napcodeframetestletlisttype `xml:"NAPCodeFrameTestletListType" json:"NAPCodeFrameTestletListType"`
}

type napcodeframetestletlisttype struct {
        Testlet []NAPTestletCodeFrameType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletCodeFrameType struct {
  naptestletcodeframetype `xml:"NAPTestletCodeFrameType" json:"NAPTestletCodeFrameType"`
}

type naptestletcodeframetype struct {
        NAPTestletRefId *String `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId,omitempty"`
      TestletContent *NAPTestletContentType `xml:"TestletContent" json:"TestletContent"`
      TestItemList *CodeFrameTestItemListType `xml:"TestItemList" json:"TestItemList"`
      
      }
    
    type NAPStudentResponseTestletListType struct {
  napstudentresponsetestletlisttype `xml:"NAPStudentResponseTestletListType" json:"NAPStudentResponseTestletListType"`
}

type napstudentresponsetestletlisttype struct {
        Testlet []NAPTestletResponseType `xml:"Testlet" json:"Testlet"`
      
      }
    
    type NAPTestletResponseType struct {
  naptestletresponsetype `xml:"NAPTestletResponseType" json:"NAPTestletResponseType"`
}

type naptestletresponsetype struct {
        NAPTestletRefId *String `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId,omitempty"`
      NAPTestletLocalId *LocalIdType `xml:"NAPTestletLocalId" json:"NAPTestletLocalId"`
      TestletSubScore *Float `xml:"TestletSubScore,omitempty" json:"TestletSubScore,omitempty"`
      ItemResponseList *NAPTestletItemResponseListType `xml:"ItemResponseList" json:"ItemResponseList"`
      
      }
    
    type NAPTestletItemResponseListType struct {
  naptestletitemresponselisttype `xml:"NAPTestletItemResponseListType" json:"NAPTestletItemResponseListType"`
}

type naptestletitemresponselisttype struct {
        ItemResponse []NAPTestletResponseItemType `xml:"ItemResponse" json:"ItemResponse"`
      
      }
    
    type NAPTestletResponseItemType struct {
  naptestletresponseitemtype `xml:"NAPTestletResponseItemType" json:"NAPTestletResponseItemType"`
}

type naptestletresponseitemtype struct {
        NAPTestItemRefId *String `xml:"NAPTestItemRefId,omitempty" json:"NAPTestItemRefId,omitempty"`
      NAPTestItemLocalId *LocalIdType `xml:"NAPTestItemLocalId" json:"NAPTestItemLocalId"`
      Response *String `xml:"Response,omitempty" json:"Response,omitempty"`
      ResponseCorrectness *AUCodeSetsNAPResponseCorrectnessType `xml:"ResponseCorrectness" json:"ResponseCorrectness"`
      Score *Float `xml:"Score,omitempty" json:"Score,omitempty"`
      LapsedTimeItem *String `xml:"LapsedTimeItem,omitempty" json:"LapsedTimeItem,omitempty"`
      SequenceNumber *Int `xml:"SequenceNumber" json:"SequenceNumber"`
      ItemWeight *Float `xml:"ItemWeight" json:"ItemWeight"`
      SubscoreList *NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList,omitempty"`
      
      }
    
    type NAPSubscoreListType struct {
  napsubscorelisttype `xml:"NAPSubscoreListType" json:"NAPSubscoreListType"`
}

type napsubscorelisttype struct {
        Subscore []NAPSubscoreType `xml:"Subscore" json:"Subscore"`
      
      }
    
    type NAPSubscoreType struct {
  napsubscoretype `xml:"NAPSubscoreType" json:"NAPSubscoreType"`
}

type napsubscoretype struct {
        SubscoreType *String `xml:"SubscoreType" json:"SubscoreType"`
      SubscoreValue *Float `xml:"SubscoreValue" json:"SubscoreValue"`
      
      }
    
    type DomainScoreType struct {
  domainscoretype `xml:"DomainScoreType" json:"DomainScoreType"`
}

type domainscoretype struct {
        RawScore *Float `xml:"RawScore" json:"RawScore"`
      ScaledScoreValue *Float `xml:"ScaledScoreValue" json:"ScaledScoreValue"`
      ScaledScoreLogitValue *Float `xml:"ScaledScoreLogitValue" json:"ScaledScoreLogitValue"`
      ScaledScoreStandardError *Float `xml:"ScaledScoreStandardError" json:"ScaledScoreStandardError"`
      ScaledScoreLogitStandardError *Float `xml:"ScaledScoreLogitStandardError" json:"ScaledScoreLogitStandardError"`
      StudentDomainBand *Int `xml:"StudentDomainBand" json:"StudentDomainBand"`
      StudentProficiency *String `xml:"StudentProficiency,omitempty" json:"StudentProficiency,omitempty"`
      PlausibleScaledValueList *PlausibleScaledValueListType `xml:"PlausibleScaledValueList" json:"PlausibleScaledValueList"`
      
      }
    
    type NAPWritingRubricListType struct {
  napwritingrubriclisttype `xml:"NAPWritingRubricListType" json:"NAPWritingRubricListType"`
}

type napwritingrubriclisttype struct {
        NAPWritingRubric []NAPWritingRubricType `xml:"NAPWritingRubric" json:"NAPWritingRubric"`
      
      }
    
    type NAPWritingRubricType struct {
  napwritingrubrictype `xml:"NAPWritingRubricType" json:"NAPWritingRubricType"`
}

type napwritingrubrictype struct {
        RubricType *String `xml:"RubricType" json:"RubricType"`
      ScoreList *ScoreListType `xml:"ScoreList" json:"ScoreList"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor,omitempty"`
      
      }
    
    type ScoreListType struct {
  scorelisttype `xml:"ScoreListType" json:"ScoreListType"`
}

type scorelisttype struct {
        Score []ScoreType `xml:"Score" json:"Score"`
      
      }
    
    type ScoreType struct {
  scoretype `xml:"ScoreType" json:"ScoreType"`
}

type scoretype struct {
        MaxScoreValue *Float `xml:"MaxScoreValue" json:"MaxScoreValue"`
      ScoreDescriptionList *ScoreDescriptionListType `xml:"ScoreDescriptionList" json:"ScoreDescriptionList"`
      
      }
    
    type ScoreDescriptionListType struct {
  scoredescriptionlisttype `xml:"ScoreDescriptionListType" json:"ScoreDescriptionListType"`
}

type scoredescriptionlisttype struct {
        ScoreDescription []ScoreDescriptionType `xml:"ScoreDescription" json:"ScoreDescription"`
      
      }
    
    type ScoreDescriptionType struct {
  scoredescriptiontype `xml:"ScoreDescriptionType" json:"ScoreDescriptionType"`
}

type scoredescriptiontype struct {
        ScoreValue *Float `xml:"ScoreValue" json:"ScoreValue"`
      Descriptor *String `xml:"Descriptor,omitempty" json:"Descriptor,omitempty"`
      
      }
    
    type StimulusListType struct {
  stimuluslisttype `xml:"StimulusListType" json:"StimulusListType"`
}

type stimuluslisttype struct {
        Stimulus []StimulusType `xml:"Stimulus" json:"Stimulus"`
      
      }
    
    type StimulusType struct {
  stimulustype `xml:"StimulusType" json:"StimulusType"`
}

type stimulustype struct {
        StimulusLocalId *LocalIdType `xml:"StimulusLocalId" json:"StimulusLocalId"`
      TextGenre *String `xml:"TextGenre,omitempty" json:"TextGenre,omitempty"`
      TextType *String `xml:"TextType,omitempty" json:"TextType,omitempty"`
      WordCount *Int `xml:"WordCount,omitempty" json:"WordCount,omitempty"`
      TextDescriptor *String `xml:"TextDescriptor,omitempty" json:"TextDescriptor,omitempty"`
      Content *String `xml:"Content" json:"Content"`
      
      }
    
    type ContentDescriptionListType struct {
  contentdescriptionlisttype `xml:"ContentDescriptionListType" json:"ContentDescriptionListType"`
}

type contentdescriptionlisttype struct {
        ContentDescription []string `xml:"ContentDescription" json:"ContentDescription"`
      
      }
    
    type PNPCodeListType struct {
  pnpcodelisttype `xml:"PNPCodeListType" json:"PNPCodeListType"`
}

type pnpcodelisttype struct {
        PNPCode []AUCodeSetsPNPCodeType `xml:"PNPCode" json:"PNPCode"`
      
      }
    
    type AdjustmentContainerType struct {
  adjustmentcontainertype `xml:"AdjustmentContainerType" json:"AdjustmentContainerType"`
}

type adjustmentcontainertype struct {
        PNPCodeList *PNPCodeListType `xml:"PNPCodeList" json:"PNPCodeList"`
      BookletType *String `xml:"BookletType,omitempty" json:"BookletType,omitempty"`
      
      }
    
    type TestDisruptionListType struct {
  testdisruptionlisttype `xml:"TestDisruptionListType" json:"TestDisruptionListType"`
}

type testdisruptionlisttype struct {
        TestDisruption []TestDisruptionType `xml:"TestDisruption" json:"TestDisruption"`
      
      }
    
    type TestDisruptionType struct {
  testdisruptiontype `xml:"TestDisruptionType" json:"TestDisruptionType"`
}

type testdisruptiontype struct {
        Event *String `xml:"Event" json:"Event"`
      
      }
    
    type CalendarSummaryListType struct {
  calendarsummarylisttype `xml:"CalendarSummaryListType" json:"CalendarSummaryListType"`
}

type calendarsummarylisttype struct {
        CalendarSummaryRefId []string `xml:"CalendarSummaryRefId" json:"CalendarSummaryRefId"`
      
      }
    
    type VisaSubClassType struct {
  visasubclasstype `xml:"VisaSubClassType" json:"VisaSubClassType"`
}

type visasubclasstype struct {
        Code *VisaSubClassCodeType `xml:"Code" json:"Code"`
      VisaExpiryDate *String `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate,omitempty"`
      ATEExpiryDate *String `xml:"ATEExpiryDate,omitempty" json:"ATEExpiryDate,omitempty"`
      ATEStartDate *String `xml:"ATEStartDate,omitempty" json:"ATEStartDate,omitempty"`
      VisaStatisticalCode *String `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode,omitempty"`
      
      }
    
    type VisaSubClassListType struct {
  visasubclasslisttype `xml:"VisaSubClassListType" json:"VisaSubClassListType"`
}

type visasubclasslisttype struct {
        VisaSubClass []VisaSubClassType `xml:"VisaSubClass" json:"VisaSubClass"`
      
      }
    
    type VisaSubClassCodeType string
    type LanguageBaseType struct {
  languagebasetype `xml:"LanguageBaseType" json:"LanguageBaseType"`
}

type languagebasetype struct {
        Code *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      LanguageType *AUCodeSetsLanguageTypeType `xml:"LanguageType,omitempty" json:"LanguageType,omitempty"`
      Dialect *String `xml:"Dialect,omitempty" json:"Dialect,omitempty"`
      
      }
    
    type ReligiousEventListType struct {
  religiouseventlisttype `xml:"ReligiousEventListType" json:"ReligiousEventListType"`
}

type religiouseventlisttype struct {
        ReligiousEvent []ReligiousEventType `xml:"ReligiousEvent" json:"ReligiousEvent"`
      
      }
    
    type ReligiousEventType struct {
  religiouseventtype `xml:"ReligiousEventType" json:"ReligiousEventType"`
}

type religiouseventtype struct {
        Type *String `xml:"Type" json:"Type"`
      Date *String `xml:"Date" json:"Date"`
      
      }
    
    type ReligionType struct {
  religiontype `xml:"ReligionType" json:"ReligionType"`
}

type religiontype struct {
        Code *AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type DwellingArrangementType struct {
  dwellingarrangementtype `xml:"DwellingArrangementType" json:"DwellingArrangementType"`
}

type dwellingarrangementtype struct {
        Code *AUCodeSetsDwellingArrangementType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type CountryListType struct {
  countrylisttype `xml:"CountryListType" json:"CountryListType"`
}

type countrylisttype struct {
        CountryOfCitizenship []CountryType `xml:"CountryOfCitizenship" json:"CountryOfCitizenship"`
      
      }
    
    type CountryList2Type struct {
  countrylist2type `xml:"CountryList2Type" json:"CountryList2Type"`
}

type countrylist2type struct {
        CountryOfResidency []CountryType `xml:"CountryOfResidency" json:"CountryOfResidency"`
      
      }
    
    type DebitOrCreditAmountType struct {
  debitorcreditamounttype `xml:"DebitOrCreditAmountType" json:"DebitOrCreditAmountType"`
}

type debitorcreditamounttype struct {
        monetaryamounttype
          Type *String `xml:"Type,attr" json:"Type"`
      
      }
    
    type ScheduledActivityOverrideType struct {
  scheduledactivityoverridetype `xml:"ScheduledActivityOverrideType" json:"ScheduledActivityOverrideType"`
}

type scheduledactivityoverridetype struct {
          DateOfOverride *String `xml:"DateOfOverride,attr" json:"DateOfOverride"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type ActivityTimeType struct {
  activitytimetype `xml:"ActivityTimeType" json:"ActivityTimeType"`
}

type activitytimetype struct {
        CreationDate *String `xml:"CreationDate" json:"CreationDate"`
      Duration *ActivityTimeType_Duration
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      FinishDate *String `xml:"FinishDate,omitempty" json:"FinishDate,omitempty"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate,omitempty"`
      
      }
    
    type SchoolCourseInfoOverrideType struct {
  schoolcourseinfooverridetype `xml:"SchoolCourseInfoOverrideType" json:"SchoolCourseInfoOverrideType"`
}

type schoolcourseinfooverridetype struct {
        Override *String `xml:"Override,attr" json:"Override"`
      CourseCode *String `xml:"CourseCode,omitempty" json:"CourseCode,omitempty"`
      StateCourseCode *String `xml:"StateCourseCode,omitempty" json:"StateCourseCode,omitempty"`
      DistrictCourseCode *String `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode,omitempty"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea,omitempty"`
      CourseTitle *String `xml:"CourseTitle,omitempty" json:"CourseTitle,omitempty"`
      InstructionalLevel *String `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel,omitempty"`
      CourseCredits *String `xml:"CourseCredits,omitempty" json:"CourseCredits,omitempty"`
      
      }
    
    type LocationOfInstructionType struct {
  locationofinstructiontype `xml:"LocationOfInstructionType" json:"LocationOfInstructionType"`
}

type locationofinstructiontype struct {
        Code *AUCodeSetsReceivingLocationOfInstructionType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type LanguageOfInstructionType struct {
  languageofinstructiontype `xml:"LanguageOfInstructionType" json:"LanguageOfInstructionType"`
}

type languageofinstructiontype struct {
        Code *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type MediumOfInstructionType struct {
  mediumofinstructiontype `xml:"MediumOfInstructionType" json:"MediumOfInstructionType"`
}

type mediumofinstructiontype struct {
        Code *AUCodeSetsMediumOfInstructionType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type StudentActivityType struct {
  studentactivitytype `xml:"StudentActivityType" json:"StudentActivityType"`
}

type studentactivitytype struct {
        Code *AUCodeSetsActivityInvolvementCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type ContactFlagsType struct {
  contactflagstype `xml:"ContactFlagsType" json:"ContactFlagsType"`
}

type contactflagstype struct {
        ParentLegalGuardian *AUCodeSetsYesOrNoCategoryType `xml:"ParentLegalGuardian,omitempty" json:"ParentLegalGuardian,omitempty"`
      PickupRights *AUCodeSetsYesOrNoCategoryType `xml:"PickupRights,omitempty" json:"PickupRights,omitempty"`
      LivesWith *AUCodeSetsYesOrNoCategoryType `xml:"LivesWith,omitempty" json:"LivesWith,omitempty"`
      AccessToRecords *AUCodeSetsYesOrNoCategoryType `xml:"AccessToRecords,omitempty" json:"AccessToRecords,omitempty"`
      ReceivesAssessmentReport *AUCodeSetsYesOrNoCategoryType `xml:"ReceivesAssessmentReport,omitempty" json:"ReceivesAssessmentReport,omitempty"`
      EmergencyContact *AUCodeSetsYesOrNoCategoryType `xml:"EmergencyContact,omitempty" json:"EmergencyContact,omitempty"`
      HasCustody *AUCodeSetsYesOrNoCategoryType `xml:"HasCustody,omitempty" json:"HasCustody,omitempty"`
      DisciplinaryContact *AUCodeSetsYesOrNoCategoryType `xml:"DisciplinaryContact,omitempty" json:"DisciplinaryContact,omitempty"`
      AttendanceContact *AUCodeSetsYesOrNoCategoryType `xml:"AttendanceContact,omitempty" json:"AttendanceContact,omitempty"`
      PrimaryCareProvider *AUCodeSetsYesOrNoCategoryType `xml:"PrimaryCareProvider,omitempty" json:"PrimaryCareProvider,omitempty"`
      FeesBilling *AUCodeSetsYesOrNoCategoryType `xml:"FeesBilling,omitempty" json:"FeesBilling,omitempty"`
      FeesAccess *AUCodeSetsYesOrNoCategoryType `xml:"FeesAccess,omitempty" json:"FeesAccess,omitempty"`
      FamilyMail *AUCodeSetsYesOrNoCategoryType `xml:"FamilyMail,omitempty" json:"FamilyMail,omitempty"`
      InterventionOrder *AUCodeSetsYesOrNoCategoryType `xml:"InterventionOrder,omitempty" json:"InterventionOrder,omitempty"`
      
      }
    
    type AgencyType struct {
  agencytype `xml:"AgencyType" json:"AgencyType"`
}

type agencytype struct {
        Code *AUCodeSetsEducationAgencyTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type YearRangeType struct {
  yearrangetype `xml:"YearRangeType" json:"YearRangeType"`
}

type yearrangetype struct {
        Start *YearLevelType `xml:"Start" json:"Start"`
      End *YearLevelType `xml:"End" json:"End"`
      
      }
    
    type CreationUserType struct {
  creationusertype `xml:"CreationUserType" json:"CreationUserType"`
}

type creationusertype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      UserId *String `xml:"UserId" json:"UserId"`
      
      }
    
    type AuditInfoType struct {
  auditinfotype `xml:"AuditInfoType" json:"AuditInfoType"`
}

type auditinfotype struct {
        CreationUser *CreationUserType `xml:"CreationUser" json:"CreationUser"`
      CreationDateTime *String `xml:"CreationDateTime" json:"CreationDateTime"`
      
      }
    
    type AttendanceInfoType struct {
  attendanceinfotype `xml:"AttendanceInfoType" json:"AttendanceInfoType"`
}

type attendanceinfotype struct {
        CountsTowardAttendance *String `xml:"CountsTowardAttendance" json:"CountsTowardAttendance"`
      AttendanceValue *Float `xml:"AttendanceValue" json:"AttendanceValue"`
      
      }
    
    type CalendarDateInfoType struct {
  calendardateinfotype `xml:"CalendarDateInfoType" json:"CalendarDateInfoType"`
}

type calendardateinfotype struct {
        Code *AUCodeSetsCalendarEventType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type ProgramAvailabilityType struct {
  programavailabilitytype `xml:"ProgramAvailabilityType" json:"ProgramAvailabilityType"`
}

type programavailabilitytype struct {
        Code *AUCodeSets0211ProgramAvailabilityType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type ReferralSourceType struct {
  referralsourcetype `xml:"ReferralSourceType" json:"ReferralSourceType"`
}

type referralsourcetype struct {
        Code *AUCodeSets0792IdentificationProcedureType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type PromotionInfoType struct {
  promotioninfotype `xml:"PromotionInfoType" json:"PromotionInfoType"`
}

type promotioninfotype struct {
        PromotionStatus *String `xml:"PromotionStatus,omitempty" json:"PromotionStatus,omitempty"`
      
      }
    
    type CatchmentStatusContainerType struct {
  catchmentstatuscontainertype `xml:"CatchmentStatusContainerType" json:"CatchmentStatusContainerType"`
}

type catchmentstatuscontainertype struct {
        Code *AUCodeSetsPublicSchoolCatchmentStatusType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type StudentExitStatusContainerType struct {
  studentexitstatuscontainertype `xml:"StudentExitStatusContainerType" json:"StudentExitStatusContainerType"`
}

type studentexitstatuscontainertype struct {
        Code *AUCodeSetsExitWithdrawalStatusType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type StudentExitContainerType struct {
  studentexitcontainertype `xml:"StudentExitContainerType" json:"StudentExitContainerType"`
}

type studentexitcontainertype struct {
        Code *AUCodeSetsExitWithdrawalTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type StudentEntryContainerType struct {
  studententrycontainertype `xml:"StudentEntryContainerType" json:"StudentEntryContainerType"`
}

type studententrycontainertype struct {
        Code *AUCodeSetsEntryTypeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type StudentMostRecentContainerType struct {
  studentmostrecentcontainertype `xml:"StudentMostRecentContainerType" json:"StudentMostRecentContainerType"`
}

type studentmostrecentcontainertype struct {
        SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      HomeroomLocalId *LocalIdType `xml:"HomeroomLocalId,omitempty" json:"HomeroomLocalId,omitempty"`
      YearLevel *YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel,omitempty"`
      FTE *Float `xml:"FTE,omitempty" json:"FTE,omitempty"`
      Parent1Language *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Parent1Language,omitempty" json:"Parent1Language,omitempty"`
      Parent2Language *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType `xml:"Parent2Language,omitempty" json:"Parent2Language,omitempty"`
      Parent1EmploymentType *AUCodeSetsEmploymentTypeType `xml:"Parent1EmploymentType,omitempty" json:"Parent1EmploymentType,omitempty"`
      Parent2EmploymentType *AUCodeSetsEmploymentTypeType `xml:"Parent2EmploymentType,omitempty" json:"Parent2EmploymentType,omitempty"`
      Parent1SchoolEducationLevel *AUCodeSetsSchoolEducationLevelTypeType `xml:"Parent1SchoolEducationLevel,omitempty" json:"Parent1SchoolEducationLevel,omitempty"`
      Parent2SchoolEducationLevel *AUCodeSetsSchoolEducationLevelTypeType `xml:"Parent2SchoolEducationLevel,omitempty" json:"Parent2SchoolEducationLevel,omitempty"`
      Parent1NonSchoolEducation *AUCodeSetsNonSchoolEducationType `xml:"Parent1NonSchoolEducation,omitempty" json:"Parent1NonSchoolEducation,omitempty"`
      Parent2NonSchoolEducation *AUCodeSetsNonSchoolEducationType `xml:"Parent2NonSchoolEducation,omitempty" json:"Parent2NonSchoolEducation,omitempty"`
      LocalCampusId *LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId,omitempty"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId,omitempty"`
      TestLevel *YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel,omitempty"`
      Homegroup *String `xml:"Homegroup,omitempty" json:"Homegroup,omitempty"`
      ClassCode *String `xml:"ClassCode,omitempty" json:"ClassCode,omitempty"`
      MembershipType *AUCodeSetsSchoolEnrollmentTypeType `xml:"MembershipType,omitempty" json:"MembershipType,omitempty"`
      FFPOS *AUCodeSetsFFPOSStatusCodeType `xml:"FFPOS,omitempty" json:"FFPOS,omitempty"`
      ReportingSchoolId *LocalIdType `xml:"ReportingSchoolId,omitempty" json:"ReportingSchoolId,omitempty"`
      OtherEnrollmentSchoolACARAId *LocalIdType `xml:"OtherEnrollmentSchoolACARAId,omitempty" json:"OtherEnrollmentSchoolACARAId,omitempty"`
      OtherSchoolName *String `xml:"OtherSchoolName,omitempty" json:"OtherSchoolName,omitempty"`
      DisabilityLevelOfAdjustment *String `xml:"DisabilityLevelOfAdjustment,omitempty" json:"DisabilityLevelOfAdjustment,omitempty"`
      DisabilityCategory *String `xml:"DisabilityCategory,omitempty" json:"DisabilityCategory,omitempty"`
      CensusAge *Int `xml:"CensusAge,omitempty" json:"CensusAge,omitempty"`
      DistanceEducationStudent *AUCodeSetsYesOrNoCategoryType `xml:"DistanceEducationStudent,omitempty" json:"DistanceEducationStudent,omitempty"`
      BoardingStatus *AUCodeSetsBoardingType `xml:"BoardingStatus,omitempty" json:"BoardingStatus,omitempty"`
      
      }
    
    type StaffMostRecentContainerType struct {
  staffmostrecentcontainertype `xml:"StaffMostRecentContainerType" json:"StaffMostRecentContainerType"`
}

type staffmostrecentcontainertype struct {
        SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId,omitempty"`
      LocalCampusId *LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId,omitempty"`
      NAPLANClassList *NAPLANClassListType `xml:"NAPLANClassList,omitempty" json:"NAPLANClassList,omitempty"`
      HomeGroup *String `xml:"HomeGroup,omitempty" json:"HomeGroup,omitempty"`
      
      }
    
    type StaffActivityExtensionType struct {
  staffactivityextensiontype `xml:"StaffActivityExtensionType" json:"StaffActivityExtensionType"`
}

type staffactivityextensiontype struct {
        Code *AUCodeSetsStaffActivityType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type TotalEnrollmentsType struct {
  totalenrollmentstype `xml:"TotalEnrollmentsType" json:"TotalEnrollmentsType"`
}

type totalenrollmentstype struct {
        Girls *String `xml:"Girls,omitempty" json:"Girls,omitempty"`
      Boys *String `xml:"Boys,omitempty" json:"Boys,omitempty"`
      TotalStudents *String `xml:"TotalStudents,omitempty" json:"TotalStudents,omitempty"`
      
      }
    
    type CampusContainerType struct {
  campuscontainertype `xml:"CampusContainerType" json:"CampusContainerType"`
}

type campuscontainertype struct {
        ParentSchoolId *String `xml:"ParentSchoolId,omitempty" json:"ParentSchoolId,omitempty"`
      SchoolCampusId *String `xml:"SchoolCampusId" json:"SchoolCampusId"`
      CampusType *AUCodeSetsSchoolLevelType `xml:"CampusType,omitempty" json:"CampusType,omitempty"`
      AdminStatus *AUCodeSetsYesOrNoCategoryType `xml:"AdminStatus" json:"AdminStatus"`
      
      }
    
    type HouseholdContactInfoListType struct {
  householdcontactinfolisttype `xml:"HouseholdContactInfoListType" json:"HouseholdContactInfoListType"`
}

type householdcontactinfolisttype struct {
        HouseholdContactInfo []HouseholdContactInfoType `xml:"HouseholdContactInfo" json:"HouseholdContactInfo"`
      
      }
    
    type HouseholdContactInfoType struct {
  householdcontactinfotype `xml:"HouseholdContactInfoType" json:"HouseholdContactInfoType"`
}

type householdcontactinfotype struct {
        PreferenceNumber *Int `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber,omitempty"`
      HouseholdContactId *LocalIdType `xml:"HouseholdContactId,omitempty" json:"HouseholdContactId,omitempty"`
      HouseholdSalutation *String `xml:"HouseholdSalutation,omitempty" json:"HouseholdSalutation,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      
      }
    
    type StatementCodesType struct {
  statementcodestype `xml:"StatementCodesType" json:"StatementCodesType"`
}

type statementcodestype struct {
        StatementCode []string `xml:"StatementCode" json:"StatementCode"`
      
      }
    
    type StatementsType struct {
  statementstype `xml:"StatementsType" json:"StatementsType"`
}

type statementstype struct {
        Statement []string `xml:"Statement" json:"Statement"`
      
      }
    
    type ProgramFundingSourcesType struct {
  programfundingsourcestype `xml:"ProgramFundingSourcesType" json:"ProgramFundingSourcesType"`
}

type programfundingsourcestype struct {
        ProgramFundingSource []ProgramFundingSourceType `xml:"ProgramFundingSource" json:"ProgramFundingSource"`
      
      }
    
    type ProgramFundingSourceType struct {
  programfundingsourcetype `xml:"ProgramFundingSourceType" json:"ProgramFundingSourceType"`
}

type programfundingsourcetype struct {
        Code *AUCodeSetsProgramFundingSourceCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type AttendanceTimesType struct {
  attendancetimestype `xml:"AttendanceTimesType" json:"AttendanceTimesType"`
}

type attendancetimestype struct {
        AttendanceTime []AttendanceTimeType `xml:"AttendanceTime" json:"AttendanceTime"`
      
      }
    
    type AttendanceTimeType struct {
  attendancetimetype `xml:"AttendanceTimeType" json:"AttendanceTimeType"`
}

type attendancetimetype struct {
        AttendanceType *String `xml:"AttendanceType,omitempty" json:"AttendanceType,omitempty"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      StartTime *String `xml:"StartTime" json:"StartTime"`
      EndTime *String `xml:"EndTime" json:"EndTime"`
      DurationValue *Float `xml:"DurationValue,omitempty" json:"DurationValue,omitempty"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote,omitempty"`
      
      }
    
    type PeriodAttendancesType struct {
  periodattendancestype `xml:"PeriodAttendancesType" json:"PeriodAttendancesType"`
}

type periodattendancestype struct {
        PeriodAttendance []PeriodAttendanceType `xml:"PeriodAttendance" json:"PeriodAttendance"`
      
      }
    
    type PeriodAttendanceType struct {
  periodattendancetype `xml:"PeriodAttendanceType" json:"PeriodAttendanceType"`
}

type periodattendancetype struct {
        AttendanceType *String `xml:"AttendanceType,omitempty" json:"AttendanceType,omitempty"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      Date *String `xml:"Date" json:"Date"`
      SessionInfoRefId *String `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId,omitempty"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId,omitempty"`
      TimetablePeriod *String `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod,omitempty"`
      DayId *LocalIdType `xml:"DayId,omitempty" json:"DayId,omitempty"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime,omitempty"`
      EndTime *String `xml:"EndTime,omitempty" json:"EndTime,omitempty"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn,omitempty"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut,omitempty"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId,omitempty"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList,omitempty"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList,omitempty"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote,omitempty"`
      
      }
    
    type StaffSubjectListType struct {
  staffsubjectlisttype `xml:"StaffSubjectListType" json:"StaffSubjectListType"`
}

type staffsubjectlisttype struct {
        StaffSubject []StaffSubjectType `xml:"StaffSubject" json:"StaffSubject"`
      
      }
    
    type StaffSubjectType struct {
  staffsubjecttype `xml:"StaffSubjectType" json:"StaffSubjectType"`
}

type staffsubjecttype struct {
        PreferenceNumber *Int `xml:"PreferenceNumber" json:"PreferenceNumber"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId,omitempty"`
      TimeTableSubjectRefId *RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      
      }
    
    type TeachingGroupListType struct {
  teachinggrouplisttype `xml:"TeachingGroupListType" json:"TeachingGroupListType"`
}

type teachinggrouplisttype struct {
        TeachingGroupRefId []string `xml:"TeachingGroupRefId" json:"TeachingGroupRefId"`
      
      }
    
    type ScheduledTeacherListType struct {
  scheduledteacherlisttype `xml:"ScheduledTeacherListType" json:"ScheduledTeacherListType"`
}

type scheduledteacherlisttype struct {
        TeacherCover []TeacherCoverType `xml:"TeacherCover" json:"TeacherCover"`
      
      }
    
    type TeacherCoverType struct {
  teachercovertype `xml:"TeacherCoverType" json:"TeacherCoverType"`
}

type teachercovertype struct {
        StaffPersonalRefId *String `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId,omitempty"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime,omitempty"`
      FinishTime *String `xml:"FinishTime,omitempty" json:"FinishTime,omitempty"`
      Credit *AUCodeSetsTeacherCoverCreditType `xml:"Credit,omitempty" json:"Credit,omitempty"`
      Supervision *AUCodeSetsTeacherCoverSupervisionType `xml:"Supervision,omitempty" json:"Supervision,omitempty"`
      Weighting *Float `xml:"Weighting,omitempty" json:"Weighting,omitempty"`
      
      }
    
    type RoomListType struct {
  roomlisttype `xml:"RoomListType" json:"RoomListType"`
}

type roomlisttype struct {
        RoomInfoRefId []string `xml:"RoomInfoRefId" json:"RoomInfoRefId"`
      
      }
    
    type StaffListType struct {
  stafflisttype `xml:"StaffListType" json:"StaffListType"`
}

type stafflisttype struct {
        StaffPersonalRefId []string `xml:"StaffPersonalRefId" json:"StaffPersonalRefId"`
      
      }
    
    type AuthorsType struct {
  authorstype `xml:"AuthorsType" json:"AuthorsType"`
}

type authorstype struct {
        Author []string `xml:"Author" json:"Author"`
      
      }
    
    type OrganizationsType struct {
  organizationstype `xml:"OrganizationsType" json:"OrganizationsType"`
}

type organizationstype struct {
        Organization []string `xml:"Organization" json:"Organization"`
      
      }
    
    type PurchasingItemsType struct {
  purchasingitemstype `xml:"PurchasingItemsType" json:"PurchasingItemsType"`
}

type purchasingitemstype struct {
        PurchasingItem []PurchasingItemType `xml:"PurchasingItem" json:"PurchasingItem"`
      
      }
    
    type PurchasingItemType struct {
  purchasingitemtype `xml:"PurchasingItemType" json:"PurchasingItemType"`
}

type purchasingitemtype struct {
        ItemNumber *String `xml:"ItemNumber,omitempty" json:"ItemNumber,omitempty"`
      ItemDescription *String `xml:"ItemDescription" json:"ItemDescription"`
      LocalItemId *LocalIdType `xml:"LocalItemId,omitempty" json:"LocalItemId,omitempty"`
      Quantity *String `xml:"Quantity,omitempty" json:"Quantity,omitempty"`
      UnitCost *MonetaryAmountType `xml:"UnitCost,omitempty" json:"UnitCost,omitempty"`
      TotalCost *MonetaryAmountType `xml:"TotalCost,omitempty" json:"TotalCost,omitempty"`
      QuantityDelivered *String `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered,omitempty"`
      CancelledOrder *AUCodeSetsYesOrNoCategoryType `xml:"CancelledOrder,omitempty" json:"CancelledOrder,omitempty"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
      ExpenseAccounts *ExpenseAccountsType `xml:"ExpenseAccounts,omitempty" json:"ExpenseAccounts,omitempty"`
      
      }
    
    type ExpenseAccountsType struct {
  expenseaccountstype `xml:"ExpenseAccountsType" json:"ExpenseAccountsType"`
}

type expenseaccountstype struct {
        ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount" json:"ExpenseAccount"`
      
      }
    
    type ExpenseAccountType struct {
  expenseaccounttype `xml:"ExpenseAccountType" json:"ExpenseAccountType"`
}

type expenseaccounttype struct {
        AccountCode *String `xml:"AccountCode" json:"AccountCode"`
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      FinancialAccountRefId *String `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId,omitempty"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod,omitempty"`
      
      }
    
    type SchoolProgramListType struct {
  schoolprogramlisttype `xml:"SchoolProgramListType" json:"SchoolProgramListType"`
}

type schoolprogramlisttype struct {
        Program []SchoolProgramType `xml:"Program" json:"Program"`
      
      }
    
    type SchoolProgramType struct {
  schoolprogramtype `xml:"SchoolProgramType" json:"SchoolProgramType"`
}

type schoolprogramtype struct {
        Category *String `xml:"Category,omitempty" json:"Category,omitempty"`
      Type *String `xml:"Type" json:"Type"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type LearningObjectivesType struct {
  learningobjectivestype `xml:"LearningObjectivesType" json:"LearningObjectivesType"`
}

type learningobjectivestype struct {
        LearningObjective []string `xml:"LearningObjective" json:"LearningObjective"`
      
      }
    
    type RecognitionListType struct {
  recognitionlisttype `xml:"RecognitionListType" json:"RecognitionListType"`
}

type recognitionlisttype struct {
        Recognition []string `xml:"Recognition" json:"Recognition"`
      
      }
    
    type LResourcesType struct {
  lresourcestype `xml:"LResourcesType" json:"LResourcesType"`
}

type lresourcestype struct {
        LearningResourceRefId []ResourcesType `xml:"LearningResourceRefId" json:"LearningResourceRefId"`
      
      }
    
    type ResourcesType struct {
  resourcestype `xml:"ResourcesType" json:"ResourcesType"`
}

type resourcestype struct {
          ResourceType *String `xml:"ResourceType,attr" json:"ResourceType"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type SourceObjectsType struct {
  sourceobjectstype `xml:"SourceObjectsType" json:"SourceObjectsType"`
}

type sourceobjectstype struct {
      SourceObject []SourceObjectsType_SourceObject `xml:"SourceObject" json:"SourceObject"`
      
      }
    
    type StudentsType struct {
  studentstype `xml:"StudentsType" json:"StudentsType"`
}

type studentstype struct {
        StudentPersonalRefId []string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      
      }
    
    type PrerequisitesType struct {
  prerequisitestype `xml:"PrerequisitesType" json:"PrerequisitesType"`
}

type prerequisitestype struct {
        Prerequisite []string `xml:"Prerequisite" json:"Prerequisite"`
      
      }
    
    type EssentialMaterialsType struct {
  essentialmaterialstype `xml:"EssentialMaterialsType" json:"EssentialMaterialsType"`
}

type essentialmaterialstype struct {
        EssentialMaterial []string `xml:"EssentialMaterial" json:"EssentialMaterial"`
      
      }
    
    type TechnicalRequirementsType struct {
  technicalrequirementstype `xml:"TechnicalRequirementsType" json:"TechnicalRequirementsType"`
}

type technicalrequirementstype struct {
        TechnicalRequirement *String `xml:"TechnicalRequirement,omitempty" json:"TechnicalRequirement,omitempty"`
      
      }
    
    type SoftwareRequirementListType struct {
  softwarerequirementlisttype `xml:"SoftwareRequirementListType" json:"SoftwareRequirementListType"`
}

type softwarerequirementlisttype struct {
        SoftwareRequirement []SoftwareRequirementType `xml:"SoftwareRequirement" json:"SoftwareRequirement"`
      
      }
    
    type SoftwareRequirementType struct {
  softwarerequirementtype `xml:"SoftwareRequirementType" json:"SoftwareRequirementType"`
}

type softwarerequirementtype struct {
        SoftwareTitle *String `xml:"SoftwareTitle" json:"SoftwareTitle"`
      Version *String `xml:"Version,omitempty" json:"Version,omitempty"`
      Vendor *String `xml:"Vendor,omitempty" json:"Vendor,omitempty"`
      OS *String `xml:"OS,omitempty" json:"OS,omitempty"`
      
      }
    
    type HouseholdListType struct {
  householdlisttype `xml:"HouseholdListType" json:"HouseholdListType"`
}

type householdlisttype struct {
        Household []LocalIdType `xml:"Household" json:"Household"`
      
      }
    
    type StudentSubjectChoiceListType struct {
  studentsubjectchoicelisttype `xml:"StudentSubjectChoiceListType" json:"StudentSubjectChoiceListType"`
}

type studentsubjectchoicelisttype struct {
        StudentSubjectChoice []StudentSubjectChoiceType `xml:"StudentSubjectChoice" json:"StudentSubjectChoice"`
      
      }
    
    type StudentSubjectChoiceType struct {
  studentsubjectchoicetype `xml:"StudentSubjectChoiceType" json:"StudentSubjectChoiceType"`
}

type studentsubjectchoicetype struct {
        PreferenceNumber *Int `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber,omitempty"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId" json:"SubjectLocalId"`
      StudyDescription *SubjectAreaType `xml:"StudyDescription,omitempty" json:"StudyDescription,omitempty"`
      OtherSchoolLocalId *LocalIdType `xml:"OtherSchoolLocalId,omitempty" json:"OtherSchoolLocalId,omitempty"`
      
      }
    
    type IdentityAssertionsType struct {
  identityassertionstype `xml:"IdentityAssertionsType" json:"IdentityAssertionsType"`
}

type identityassertionstype struct {
      IdentityAssertion []IdentityAssertionsType_IdentityAssertion `xml:"IdentityAssertion" json:"IdentityAssertion"`
      
      }
    
    type LearningStandardsType struct {
  learningstandardstype `xml:"LearningStandardsType" json:"LearningStandardsType"`
}

type learningstandardstype struct {
        LearningStandardItemRefId []string `xml:"LearningStandardItemRefId" json:"LearningStandardItemRefId"`
      
      }
    
    type LearningResourcesType struct {
  learningresourcestype `xml:"LearningResourcesType" json:"LearningResourcesType"`
}

type learningresourcestype struct {
        LearningResourceRefId []string `xml:"LearningResourceRefId" json:"LearningResourceRefId"`
      
      }
    
    type LearningStandardsDocumentType struct {
  learningstandardsdocumenttype `xml:"LearningStandardsDocumentType" json:"LearningStandardsDocumentType"`
}

type learningstandardsdocumenttype struct {
        LearningStandardDocumentRefId []string `xml:"LearningStandardDocumentRefId" json:"LearningStandardDocumentRefId"`
      
      }
    
    type ComponentsType struct {
  componentstype `xml:"ComponentsType" json:"ComponentsType"`
}

type componentstype struct {
        Component []ComponentType `xml:"Component" json:"Component"`
      
      }
    
    type ComponentType struct {
  componenttype `xml:"ComponentType" json:"ComponentType"`
}

type componenttype struct {
        Name *String `xml:"Name" json:"Name"`
      Reference *String `xml:"Reference" json:"Reference"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      Strategies *StrategiesType `xml:"Strategies,omitempty" json:"Strategies,omitempty"`
      AssociatedObjects *AssociatedObjectsType `xml:"AssociatedObjects,omitempty" json:"AssociatedObjects,omitempty"`
      
      }
    
    type StrategiesType struct {
  strategiestype `xml:"StrategiesType" json:"StrategiesType"`
}

type strategiestype struct {
        Strategy []string `xml:"Strategy" json:"Strategy"`
      
      }
    
    type AssociatedObjectsType struct {
  associatedobjectstype `xml:"AssociatedObjectsType" json:"AssociatedObjectsType"`
}

type associatedobjectstype struct {
      AssociatedObject []AssociatedObjectsType_AssociatedObject `xml:"AssociatedObject" json:"AssociatedObject"`
      
      }
    
    type EvaluationsType struct {
  evaluationstype `xml:"EvaluationsType" json:"EvaluationsType"`
}

type evaluationstype struct {
        Evaluation []EvaluationType `xml:"Evaluation" json:"Evaluation"`
      
      }
    
    type EvaluationType struct {
  evaluationtype `xml:"EvaluationType" json:"EvaluationType"`
}

type evaluationtype struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      Date *String `xml:"Date,omitempty" json:"Date,omitempty"`
      Name *NameType `xml:"Name,omitempty" json:"Name,omitempty"`
      
      }
    
    type ApprovalsType struct {
  approvalstype `xml:"ApprovalsType" json:"ApprovalsType"`
}

type approvalstype struct {
        Approval []ApprovalType `xml:"Approval" json:"Approval"`
      
      }
    
    type ApprovalType struct {
  approvaltype `xml:"ApprovalType" json:"ApprovalType"`
}

type approvaltype struct {
        Organization *String `xml:"Organization" json:"Organization"`
      Date *String `xml:"Date" json:"Date"`
      
      }
    
    type MediaTypesType struct {
  mediatypestype `xml:"MediaTypesType" json:"MediaTypesType"`
}

type mediatypestype struct {
        MediaType []string `xml:"MediaType" json:"MediaType"`
      
      }
    
    type LEAContactListType struct {
  leacontactlisttype `xml:"LEAContactListType" json:"LEAContactListType"`
}

type leacontactlisttype struct {
        LEAContact []LEAContactType `xml:"LEAContact" json:"LEAContact"`
      
      }
    
    type LEAContactType struct {
  leacontacttype `xml:"LEAContactType" json:"LEAContactType"`
}

type leacontacttype struct {
        PublishInDirectory *PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory,omitempty"`
      ContactInfo *ContactInfoType `xml:"ContactInfo" json:"ContactInfo"`
      
      }
    
    type FinancialAccountRefIdListType struct {
  financialaccountrefidlisttype `xml:"FinancialAccountRefIdListType" json:"FinancialAccountRefIdListType"`
}

type financialaccountrefidlisttype struct {
        FinancialAccountRefId []string `xml:"FinancialAccountRefId" json:"FinancialAccountRefId"`
      
      }
    
    type AccountCodeListType struct {
  accountcodelisttype `xml:"AccountCodeListType" json:"AccountCodeListType"`
}

type accountcodelisttype struct {
        AccountCode []string `xml:"AccountCode" json:"AccountCode"`
      
      }
    
    type JournalAdjustmentListType struct {
  journaladjustmentlisttype `xml:"JournalAdjustmentListType" json:"JournalAdjustmentListType"`
}

type journaladjustmentlisttype struct {
        JournalAdjustment []JournalAdjustmentType `xml:"JournalAdjustment" json:"JournalAdjustment"`
      
      }
    
    type JournalAdjustmentType struct {
  journaladjustmenttype `xml:"JournalAdjustmentType" json:"JournalAdjustmentType"`
}

type journaladjustmenttype struct {
        DebitFinancialAccountRefId *String `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId,omitempty"`
      CreditFinancialAccountRefId *String `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId,omitempty"`
      DebitAccountCode *String `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode,omitempty"`
      CreditAccountCode *String `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode,omitempty"`
      GSTCodeOriginal *String `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal,omitempty"`
      GSTCodeReplacement *String `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement,omitempty"`
      LineAdjustmentAmount *MonetaryAmountType `xml:"LineAdjustmentAmount" json:"LineAdjustmentAmount"`
      
      }
    
    type PaymentReceiptLineListType struct {
  paymentreceiptlinelisttype `xml:"PaymentReceiptLineListType" json:"PaymentReceiptLineListType"`
}

type paymentreceiptlinelisttype struct {
        PaymentReceiptLine []PaymentReceiptLineType `xml:"PaymentReceiptLine" json:"PaymentReceiptLine"`
      
      }
    
    type PaymentReceiptLineType struct {
  paymentreceiptlinetype `xml:"PaymentReceiptLineType" json:"PaymentReceiptLineType"`
}

type paymentreceiptlinetype struct {
        InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId,omitempty"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      LocalPaymentReceiptLineId *LocalIdType `xml:"LocalPaymentReceiptLineId,omitempty" json:"LocalPaymentReceiptLineId,omitempty"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      FinancialAccountRefId *String `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId,omitempty"`
      AccountCode *String `xml:"AccountCode,omitempty" json:"AccountCode,omitempty"`
      TransactionDescription *String `xml:"TransactionDescription,omitempty" json:"TransactionDescription,omitempty"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount,omitempty"`
      
      }
    
    type PasswordListType struct {
  passwordlisttype `xml:"PasswordListType" json:"PasswordListType"`
}

type passwordlisttype struct {
      Password []PasswordListType_Password `xml:"Password" json:"Password"`
      
      }
    
    type ExclusionRulesType struct {
  exclusionrulestype `xml:"ExclusionRulesType" json:"ExclusionRulesType"`
}

type exclusionrulestype struct {
        ExclusionRule []ExclusionRuleType `xml:"ExclusionRule" json:"ExclusionRule"`
      
      }
    
    type ExclusionRuleType struct {
  exclusionruletype `xml:"ExclusionRuleType" json:"ExclusionRuleType"`
}

type exclusionruletype struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type CharacteristicsType struct {
  characteristicstype `xml:"CharacteristicsType" json:"CharacteristicsType"`
}

type characteristicstype struct {
        AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId" json:"AggregateCharacteristicInfoRefId"`
      
      }
    
    type ContactsType struct {
  contactstype `xml:"ContactsType" json:"ContactsType"`
}

type contactstype struct {
        Contact []ContactType `xml:"Contact" json:"Contact"`
      
      }
    
    type ContactType struct {
  contacttype `xml:"ContactType" json:"ContactType"`
}

type contacttype struct {
        Name *NameType `xml:"Name,omitempty" json:"Name,omitempty"`
      Address *AddressType `xml:"Address,omitempty" json:"Address,omitempty"`
      PhoneNumber *PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
      Email *EmailType `xml:"Email,omitempty" json:"Email,omitempty"`
      
      }
    
    type TeachingGroupPeriodListType struct {
  teachinggroupperiodlisttype `xml:"TeachingGroupPeriodListType" json:"TeachingGroupPeriodListType"`
}

type teachinggroupperiodlisttype struct {
        TeachingGroupPeriod []TeachingGroupPeriodType `xml:"TeachingGroupPeriod" json:"TeachingGroupPeriod"`
      
      }
    
    type TeachingGroupPeriodType struct {
  teachinggroupperiodtype `xml:"TeachingGroupPeriodType" json:"TeachingGroupPeriodType"`
}

type teachinggroupperiodtype struct {
        TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId,omitempty"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber,omitempty"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId,omitempty"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId,omitempty"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime,omitempty"`
      CellType *String `xml:"CellType,omitempty" json:"CellType,omitempty"`
      
      }
    
    type TeacherListType struct {
  teacherlisttype `xml:"TeacherListType" json:"TeacherListType"`
}

type teacherlisttype struct {
        TeachingGroupTeacher []TeachingGroupTeacherType `xml:"TeachingGroupTeacher" json:"TeachingGroupTeacher"`
      
      }
    
    type TeachingGroupTeacherType struct {
  teachinggroupteachertype `xml:"TeachingGroupTeacherType" json:"TeachingGroupTeacherType"`
}

type teachinggroupteachertype struct {
        StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId,omitempty"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId,omitempty"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name,omitempty"`
      Association *String `xml:"Association" json:"Association"`
      
      }
    
    type StudentListType struct {
  studentlisttype `xml:"StudentListType" json:"StudentListType"`
}

type studentlisttype struct {
        TeachingGroupStudent []TeachingGroupStudentType `xml:"TeachingGroupStudent" json:"TeachingGroupStudent"`
      
      }
    
    type TeachingGroupStudentType struct {
  teachinggroupstudenttype `xml:"TeachingGroupStudentType" json:"TeachingGroupStudentType"`
}

type teachinggroupstudenttype struct {
        StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      StudentLocalId *LocalIdType `xml:"StudentLocalId,omitempty" json:"StudentLocalId,omitempty"`
      Name *NameOfRecordType `xml:"Name,omitempty" json:"Name,omitempty"`
      
      }
    
    type TimeTableDayListType struct {
  timetabledaylisttype `xml:"TimeTableDayListType" json:"TimeTableDayListType"`
}

type timetabledaylisttype struct {
        TimeTableDay []TimeTableDayType `xml:"TimeTableDay" json:"TimeTableDay"`
      
      }
    
    type TimeTableDayType struct {
  timetabledaytype `xml:"TimeTableDayType" json:"TimeTableDayType"`
}

type timetabledaytype struct {
        DayId *LocalIdType `xml:"DayId" json:"DayId"`
      DayTitle *String `xml:"DayTitle" json:"DayTitle"`
      TimeTablePeriodList *TimeTablePeriodListType `xml:"TimeTablePeriodList" json:"TimeTablePeriodList"`
      
      }
    
    type TimeTablePeriodListType struct {
  timetableperiodlisttype `xml:"TimeTablePeriodListType" json:"TimeTablePeriodListType"`
}

type timetableperiodlisttype struct {
        TimeTablePeriod []TimeTablePeriodType `xml:"TimeTablePeriod" json:"TimeTablePeriod"`
      
      }
    
    type TimeTablePeriodType struct {
  timetableperiodtype `xml:"TimeTablePeriodType" json:"TimeTablePeriodType"`
}

type timetableperiodtype struct {
        PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      PeriodTitle *String `xml:"PeriodTitle" json:"PeriodTitle"`
      BellPeriod *String `xml:"BellPeriod,omitempty" json:"BellPeriod,omitempty"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime,omitempty"`
      EndTime *String `xml:"EndTime,omitempty" json:"EndTime,omitempty"`
      RegularSchoolPeriod *String `xml:"RegularSchoolPeriod,omitempty" json:"RegularSchoolPeriod,omitempty"`
      InstructionalMinutes *Int `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes,omitempty"`
      UseInAttendanceCalculations *String `xml:"UseInAttendanceCalculations,omitempty" json:"UseInAttendanceCalculations,omitempty"`
      
      }
    
    type NAPLANClassListType struct {
  naplanclasslisttype `xml:"NAPLANClassListType" json:"NAPLANClassListType"`
}

type naplanclasslisttype struct {
        ClassCode []string `xml:"ClassCode" json:"ClassCode"`
      
      }
    
    type SchoolGroupListType struct {
  schoolgrouplisttype `xml:"SchoolGroupListType" json:"SchoolGroupListType"`
}

type schoolgrouplisttype struct {
        SchoolGroup []LocalIdType `xml:"SchoolGroup" json:"SchoolGroup"`
      
      }
    
    type YearLevelEnrollmentListType struct {
  yearlevelenrollmentlisttype `xml:"YearLevelEnrollmentListType" json:"YearLevelEnrollmentListType"`
}

type yearlevelenrollmentlisttype struct {
        YearLevelEnrollment []YearLevelEnrollmentType `xml:"YearLevelEnrollment" json:"YearLevelEnrollment"`
      
      }
    
    type YearLevelEnrollmentType struct {
  yearlevelenrollmenttype `xml:"YearLevelEnrollmentType" json:"YearLevelEnrollmentType"`
}

type yearlevelenrollmenttype struct {
        Year *AUCodeSetsYearLevelCodeType `xml:"Year" json:"Year"`
      Enrollment *String `xml:"Enrollment" json:"Enrollment"`
      
      }
    
    type SchoolFocusListType struct {
  schoolfocuslisttype `xml:"SchoolFocusListType" json:"SchoolFocusListType"`
}

type schoolfocuslisttype struct {
        SchoolFocus []AUCodeSetsSchoolFocusCodeType `xml:"SchoolFocus" json:"SchoolFocus"`
      
      }
    
    type AlertMessagesType struct {
  alertmessagestype `xml:"AlertMessagesType" json:"AlertMessagesType"`
}

type alertmessagestype struct {
        AlertMessage []AlertMessageType `xml:"AlertMessage" json:"AlertMessage"`
      
      }
    
    type AlertMessageType struct {
  alertmessagetype `xml:"AlertMessageType" json:"AlertMessageType"`
}

type alertmessagetype struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type MedicalAlertMessagesType struct {
  medicalalertmessagestype `xml:"MedicalAlertMessagesType" json:"MedicalAlertMessagesType"`
}

type medicalalertmessagestype struct {
        MedicalAlertMessage []MedicalAlertMessageType `xml:"MedicalAlertMessage" json:"MedicalAlertMessage"`
      
      }
    
    type MedicalAlertMessageType struct {
  medicalalertmessagetype `xml:"MedicalAlertMessageType" json:"MedicalAlertMessageType"`
}

type medicalalertmessagetype struct {
          Severity *String `xml:"Severity,attr" json:"Severity"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type OtherIdListType struct {
  otheridlisttype `xml:"OtherIdListType" json:"OtherIdListType"`
}

type otheridlisttype struct {
        OtherId []OtherIdType `xml:"OtherId" json:"OtherId"`
      
      }
    
    type OtherIdType struct {
  otheridtype `xml:"OtherIdType" json:"OtherIdType"`
}

type otheridtype struct {
          Type *String `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type BaseNameType struct {
  basenametype `xml:"BaseNameType" json:"BaseNameType"`
}

type basenametype struct {
        Title *String `xml:"Title,omitempty" json:"Title,omitempty"`
      FamilyName *String `xml:"FamilyName,omitempty" json:"FamilyName,omitempty"`
      GivenName *String `xml:"GivenName,omitempty" json:"GivenName,omitempty"`
      MiddleName *String `xml:"MiddleName,omitempty" json:"MiddleName,omitempty"`
      FamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"FamilyNameFirst,omitempty" json:"FamilyNameFirst,omitempty"`
      PreferredFamilyName *String `xml:"PreferredFamilyName,omitempty" json:"PreferredFamilyName,omitempty"`
      PreferredFamilyNameFirst *AUCodeSetsYesOrNoCategoryType `xml:"PreferredFamilyNameFirst,omitempty" json:"PreferredFamilyNameFirst,omitempty"`
      PreferredGivenName *String `xml:"PreferredGivenName,omitempty" json:"PreferredGivenName,omitempty"`
      Suffix *String `xml:"Suffix,omitempty" json:"Suffix,omitempty"`
      FullName *String `xml:"FullName,omitempty" json:"FullName,omitempty"`
      
      }
    
    type NameOfRecordType struct {
  nameofrecordtype `xml:"NameOfRecordType" json:"NameOfRecordType"`
}

type nameofrecordtype struct {
        basenametype
          Type *String `xml:"Type,attr" json:"Type"`
      
      }
    
    type OtherNameType struct {
  othernametype `xml:"OtherNameType" json:"OtherNameType"`
}

type othernametype struct {
        basenametype
          Type *AUCodeSetsNameUsageTypeType `xml:"Type,attr" json:"Type"`
      
      }
    
    type PartialDateType string
    type LocalIdType string
    type LocationType struct {
  locationtype `xml:"LocationType" json:"LocationType"`
}

type locationtype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      LocationName *String `xml:"LocationName,omitempty" json:"LocationName,omitempty"`
      LocationRefId *LocationType_LocationRefId
      
      }
    
    type StateProvinceIdType string
    type AttendanceCodeType struct {
  attendancecodetype `xml:"AttendanceCodeType" json:"AttendanceCodeType"`
}

type attendancecodetype struct {
        Code *AUCodeSetsAttendanceCodeType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type YearLevelType struct {
  yearleveltype `xml:"YearLevelType" json:"YearLevelType"`
}

type yearleveltype struct {
        Code *AUCodeSetsYearLevelCodeType `xml:"Code" json:"Code"`
      
      }
    
    type PersonInfoType struct {
  personinfotype `xml:"PersonInfoType" json:"PersonInfoType"`
}

type personinfotype struct {
        Name *NameOfRecordType `xml:"Name" json:"Name"`
      OtherNames *OtherNamesType `xml:"OtherNames,omitempty" json:"OtherNames,omitempty"`
      Demographics *DemographicsType `xml:"Demographics,omitempty" json:"Demographics,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList,omitempty"`
      HouseholdContactInfoList *HouseholdContactInfoListType `xml:"HouseholdContactInfoList,omitempty" json:"HouseholdContactInfoList,omitempty"`
      
      }
    
    type YearLevelsType struct {
  yearlevelstype `xml:"YearLevelsType" json:"YearLevelsType"`
}

type yearlevelstype struct {
        YearLevel []YearLevelType `xml:"YearLevel" json:"YearLevel"`
      
      }
    
    type SchoolURLType string
    type PrincipalInfoType struct {
  principalinfotype `xml:"PrincipalInfoType" json:"PrincipalInfoType"`
}

type principalinfotype struct {
        ContactName *NameOfRecordType `xml:"ContactName" json:"ContactName"`
      ContactTitle *String `xml:"ContactTitle,omitempty" json:"ContactTitle,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList,omitempty"`
      
      }
    
    type SchoolContactType struct {
  schoolcontacttype `xml:"SchoolContactType" json:"SchoolContactType"`
}

type schoolcontacttype struct {
        PublishInDirectory *PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory,omitempty"`
      ContactInfo *ContactInfoType `xml:"ContactInfo" json:"ContactInfo"`
      
      }
    
    type SchoolContactListType struct {
  schoolcontactlisttype `xml:"SchoolContactListType" json:"SchoolContactListType"`
}

type schoolcontactlisttype struct {
        SchoolContact []SchoolContactType `xml:"SchoolContact" json:"SchoolContact"`
      
      }
    
    type PublishInDirectoryType AUCodeSetsYesOrNoCategoryType
    type ContactInfoType struct {
  contactinfotype `xml:"ContactInfoType" json:"ContactInfoType"`
}

type contactinfotype struct {
        Name *NameType `xml:"Name" json:"Name"`
      PositionTitle *String `xml:"PositionTitle,omitempty" json:"PositionTitle,omitempty"`
      Role *String `xml:"Role,omitempty" json:"Role,omitempty"`
      RegistrationDetails *String `xml:"RegistrationDetails,omitempty" json:"RegistrationDetails,omitempty"`
      Qualifications *String `xml:"Qualifications,omitempty" json:"Qualifications,omitempty"`
      Address *AddressType `xml:"Address,omitempty" json:"Address,omitempty"`
      EmailList *EmailListType `xml:"EmailList,omitempty" json:"EmailList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      
      }
    
    type AddressStreetType struct {
  addressstreettype `xml:"AddressStreetType" json:"AddressStreetType"`
}

type addressstreettype struct {
        Line1 *String `xml:"Line1,omitempty" json:"Line1,omitempty"`
      Line2 *String `xml:"Line2,omitempty" json:"Line2,omitempty"`
      Line3 *String `xml:"Line3,omitempty" json:"Line3,omitempty"`
      Complex *String `xml:"Complex,omitempty" json:"Complex,omitempty"`
      StreetNumber *String `xml:"StreetNumber,omitempty" json:"StreetNumber,omitempty"`
      StreetPrefix *String `xml:"StreetPrefix,omitempty" json:"StreetPrefix,omitempty"`
      StreetName *String `xml:"StreetName,omitempty" json:"StreetName,omitempty"`
      StreetType *String `xml:"StreetType,omitempty" json:"StreetType,omitempty"`
      StreetSuffix *String `xml:"StreetSuffix,omitempty" json:"StreetSuffix,omitempty"`
      ApartmentType *String `xml:"ApartmentType,omitempty" json:"ApartmentType,omitempty"`
      ApartmentNumberPrefix *String `xml:"ApartmentNumberPrefix,omitempty" json:"ApartmentNumberPrefix,omitempty"`
      ApartmentNumber *String `xml:"ApartmentNumber,omitempty" json:"ApartmentNumber,omitempty"`
      ApartmentNumberSuffix *String `xml:"ApartmentNumberSuffix,omitempty" json:"ApartmentNumberSuffix,omitempty"`
      
      }
    
    type AddressType struct {
  addresstype `xml:"AddressType" json:"AddressType"`
}

type addresstype struct {
        Type *AUCodeSetsAddressTypeType `xml:"Type,attr" json:"Type"`
      Role *AUCodeSetsAddressRoleType `xml:"Role,attr" json:"Role"`
      EffectiveFromDate *String `xml:"EffectiveFromDate,omitempty" json:"EffectiveFromDate,omitempty"`
      EffectiveToDate *String `xml:"EffectiveToDate,omitempty" json:"EffectiveToDate,omitempty"`
      Street *AddressStreetType `xml:"Street" json:"Street"`
      City *String `xml:"City" json:"City"`
      StateProvince *StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince,omitempty"`
      Country *CountryType `xml:"Country,omitempty" json:"Country,omitempty"`
      PostalCode *String `xml:"PostalCode" json:"PostalCode"`
      GridLocation *GridLocationType `xml:"GridLocation,omitempty" json:"GridLocation,omitempty"`
      MapReference *MapReferenceType `xml:"MapReference,omitempty" json:"MapReference,omitempty"`
      RadioContact *String `xml:"RadioContact,omitempty" json:"RadioContact,omitempty"`
      Community *String `xml:"Community,omitempty" json:"Community,omitempty"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      AddressGlobalUID *GUIDType `xml:"AddressGlobalUID,omitempty" json:"AddressGlobalUID,omitempty"`
      StatisticalAreas *StatisticalAreasType `xml:"StatisticalAreas,omitempty" json:"StatisticalAreas,omitempty"`
      
      }
    
    type MapReferenceType struct {
  mapreferencetype `xml:"MapReferenceType" json:"MapReferenceType"`
}

type mapreferencetype struct {
        Type *String `xml:"Type,attr" json:"Type"`
      XCoordinate *String `xml:"XCoordinate" json:"XCoordinate"`
      YCoordinate *String `xml:"YCoordinate" json:"YCoordinate"`
      MapNumber *String `xml:"MapNumber,omitempty" json:"MapNumber,omitempty"`
      
      }
    
    type StatisticalAreasType struct {
  statisticalareastype `xml:"StatisticalAreasType" json:"StatisticalAreasType"`
}

type statisticalareastype struct {
        StatisticalArea []StatisticalAreaType `xml:"StatisticalArea" json:"StatisticalArea"`
      
      }
    
    type StatisticalAreaType struct {
  statisticalareatype `xml:"StatisticalAreaType" json:"StatisticalAreaType"`
}

type statisticalareatype struct {
          SpatialUnitType *String `xml:"SpatialUnitType,attr" json:"SpatialUnitType"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type AddressListType struct {
  addresslisttype `xml:"AddressListType" json:"AddressListType"`
}

type addresslisttype struct {
        Address []AddressType `xml:"Address" json:"Address"`
      
      }
    
    type EmailListType struct {
  emaillisttype `xml:"EmailListType" json:"EmailListType"`
}

type emaillisttype struct {
        Email []EmailType `xml:"Email" json:"Email"`
      
      }
    
    type EmailType struct {
  emailtype `xml:"EmailType" json:"EmailType"`
}

type emailtype struct {
          Type *AUCodeSetsEmailTypeType `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type PhoneNumberListType struct {
  phonenumberlisttype `xml:"PhoneNumberListType" json:"PhoneNumberListType"`
}

type phonenumberlisttype struct {
        PhoneNumber []PhoneNumberType `xml:"PhoneNumber" json:"PhoneNumber"`
      
      }
    
    type PhoneNumberType struct {
  phonenumbertype `xml:"PhoneNumberType" json:"PhoneNumberType"`
}

type phonenumbertype struct {
        Type *AUCodeSetsTelephoneNumberTypeType `xml:"Type,attr" json:"Type"`
      Number *String `xml:"Number" json:"Number"`
      Extension *String `xml:"Extension,omitempty" json:"Extension,omitempty"`
      ListedStatus *AUCodeSetsYesOrNoCategoryType `xml:"ListedStatus,omitempty" json:"ListedStatus,omitempty"`
      Preference *Int `xml:"Preference,omitempty" json:"Preference,omitempty"`
      
      }
    
    type CountryType AUCodeSetsStandardAustralianClassificationOfCountriesSACCType
    type GridLocationType struct {
  gridlocationtype `xml:"GridLocationType" json:"GridLocationType"`
}

type gridlocationtype struct {
        Latitude *Float `xml:"Latitude" json:"Latitude"`
      Longitude *Float `xml:"Longitude" json:"Longitude"`
      
      }
    
    type OperationalStatusType AUCodeSetsOperationalStatusType
    type StateProvinceType string
    type SchoolYearType string
    type ElectronicIdListType struct {
  electronicidlisttype `xml:"ElectronicIdListType" json:"ElectronicIdListType"`
}

type electronicidlisttype struct {
        ElectronicId []ElectronicIdType `xml:"ElectronicId" json:"ElectronicId"`
      
      }
    
    type ElectronicIdType struct {
  electronicidtype `xml:"ElectronicIdType" json:"ElectronicIdType"`
}

type electronicidtype struct {
          Type *AUCodeSetsElectronicIdTypeType `xml:"Type,attr" json:"Type"`
      
        Value *String `xml:",chardata" json:"value"`
      }
    
    type OtherNamesType struct {
  othernamestype `xml:"OtherNamesType" json:"OtherNamesType"`
}

type othernamestype struct {
        Name []OtherNameType `xml:"Name" json:"Name"`
      
      }
    
    type DemographicsType struct {
  demographicstype `xml:"DemographicsType" json:"DemographicsType"`
}

type demographicstype struct {
        IndigenousStatus *AUCodeSetsIndigenousStatusType `xml:"IndigenousStatus,omitempty" json:"IndigenousStatus,omitempty"`
      Sex *AUCodeSetsSexCodeType `xml:"Sex,omitempty" json:"Sex,omitempty"`
      BirthDate *BirthDateType `xml:"BirthDate,omitempty" json:"BirthDate,omitempty"`
      DateOfDeath *String `xml:"DateOfDeath,omitempty" json:"DateOfDeath,omitempty"`
      Deceased *AUCodeSetsYesOrNoCategoryType `xml:"Deceased,omitempty" json:"Deceased,omitempty"`
      BirthDateVerification *AUCodeSetsBirthdateVerificationType `xml:"BirthDateVerification,omitempty" json:"BirthDateVerification,omitempty"`
      PlaceOfBirth *String `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth,omitempty"`
      StateOfBirth *StateProvinceType `xml:"StateOfBirth,omitempty" json:"StateOfBirth,omitempty"`
      CountryOfBirth *CountryType `xml:"CountryOfBirth,omitempty" json:"CountryOfBirth,omitempty"`
      CountriesOfCitizenship *CountryListType `xml:"CountriesOfCitizenship,omitempty" json:"CountriesOfCitizenship,omitempty"`
      CountriesOfResidency *CountryList2Type `xml:"CountriesOfResidency,omitempty" json:"CountriesOfResidency,omitempty"`
      CountryArrivalDate *String `xml:"CountryArrivalDate,omitempty" json:"CountryArrivalDate,omitempty"`
      AustralianCitizenshipStatus *AUCodeSetsAustralianCitizenshipStatusType `xml:"AustralianCitizenshipStatus,omitempty" json:"AustralianCitizenshipStatus,omitempty"`
      EnglishProficiency *EnglishProficiencyType `xml:"EnglishProficiency,omitempty" json:"EnglishProficiency,omitempty"`
      LanguageList *LanguageListType `xml:"LanguageList,omitempty" json:"LanguageList,omitempty"`
      DwellingArrangement *DwellingArrangementType `xml:"DwellingArrangement,omitempty" json:"DwellingArrangement,omitempty"`
      Religion *ReligionType `xml:"Religion,omitempty" json:"Religion,omitempty"`
      ReligiousEventList *ReligiousEventListType `xml:"ReligiousEventList,omitempty" json:"ReligiousEventList,omitempty"`
      ReligiousRegion *String `xml:"ReligiousRegion,omitempty" json:"ReligiousRegion,omitempty"`
      PermanentResident *AUCodeSetsPermanentResidentStatusType `xml:"PermanentResident,omitempty" json:"PermanentResident,omitempty"`
      VisaSubClass *VisaSubClassCodeType `xml:"VisaSubClass,omitempty" json:"VisaSubClass,omitempty"`
      VisaStatisticalCode *String `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode,omitempty"`
      VisaExpiryDate *String `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate,omitempty"`
      VisaSubClassList *VisaSubClassListType `xml:"VisaSubClassList,omitempty" json:"VisaSubClassList,omitempty"`
      LBOTE *AUCodeSetsYesOrNoCategoryType `xml:"LBOTE,omitempty" json:"LBOTE,omitempty"`
      InterpreterRequired *AUCodeSetsYesOrNoCategoryType `xml:"InterpreterRequired,omitempty" json:"InterpreterRequired,omitempty"`
      ImmunisationCertificateStatus *AUCodeSetsImmunisationCertificateStatusType `xml:"ImmunisationCertificateStatus,omitempty" json:"ImmunisationCertificateStatus,omitempty"`
      CulturalBackground *AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType `xml:"CulturalBackground,omitempty" json:"CulturalBackground,omitempty"`
      MaritalStatus *AUCodeSetsMaritalStatusAIHWType `xml:"MaritalStatus,omitempty" json:"MaritalStatus,omitempty"`
      MedicareNumber *String `xml:"MedicareNumber,omitempty" json:"MedicareNumber,omitempty"`
      
      }
    
    type EnglishProficiencyType struct {
  englishproficiencytype `xml:"EnglishProficiencyType" json:"EnglishProficiencyType"`
}

type englishproficiencytype struct {
        Code *AUCodeSetsEnglishProficiencyType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type LanguageListType struct {
  languagelisttype `xml:"LanguageListType" json:"LanguageListType"`
}

type languagelisttype struct {
        Language []LanguageBaseType `xml:"Language" json:"Language"`
      
      }
    
    type BirthDateType string
    type ProjectedGraduationYearType string
    type OnTimeGraduationYearType string
    type RelationshipType struct {
  relationshiptype `xml:"RelationshipType" json:"RelationshipType"`
}

type relationshiptype struct {
        Code *AUCodeSetsRelationshipToStudentType `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type EducationalLevelType AUCodeSetsSchoolEducationLevelTypeType
    type GraduationDateType PartialDateType
    type NameType struct {
  nametype `xml:"NameType" json:"NameType"`
}

type nametype struct {
        basenametype
          Type *AUCodeSetsNameUsageTypeType `xml:"Type,attr" json:"Type"`
      
      }
    
    type HomeroomNumberType string
    type TimeElementType struct {
  timeelementtype `xml:"TimeElementType" json:"TimeElementType"`
}

type timeelementtype struct {
        Type *String `xml:"Type" json:"Type"`
      Code *String `xml:"Code" json:"Code"`
      Name *String `xml:"Name" json:"Name"`
      Value *String `xml:"Value" json:"Value"`
      StartDateTime *String `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty"`
      EndDateTime *String `xml:"EndDateTime,omitempty" json:"EndDateTime,omitempty"`
      SpanGaps *TimeElementType_SpanGaps
      IsCurrent *Bool `xml:"IsCurrent" json:"IsCurrent"`
      
      }
    
    type LifeCycleType struct {
  lifecycletype `xml:"LifeCycleType" json:"LifeCycleType"`
}

type lifecycletype struct {
      Created *LifeCycleType_Created
      ModificationHistory *LifeCycleType_ModificationHistory
      TimeElements *LifeCycleType_TimeElements
      
      }
    
    type OtherCodeListType struct {
  othercodelisttype `xml:"OtherCodeListType" json:"OtherCodeListType"`
}

type othercodelisttype struct {
      OtherCode []OtherCodeListType_OtherCode `xml:"OtherCode" json:"OtherCode"`
      
      }
    
    type ProgramStatusType struct {
  programstatustype `xml:"ProgramStatusType" json:"ProgramStatusType"`
}

type programstatustype struct {
        Code *String `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type SubjectAreaListType struct {
  subjectarealisttype `xml:"SubjectAreaListType" json:"SubjectAreaListType"`
}

type subjectarealisttype struct {
        SubjectArea []SubjectAreaType `xml:"SubjectArea" json:"SubjectArea"`
      
      }
    
    type ACStrandAreaListType struct {
  acstrandarealisttype `xml:"ACStrandAreaListType" json:"ACStrandAreaListType"`
}

type acstrandarealisttype struct {
        ACStrandSubjectArea []ACStrandSubjectAreaType `xml:"ACStrandSubjectArea" json:"ACStrandSubjectArea"`
      
      }
    
    type SubjectAreaType struct {
  subjectareatype `xml:"SubjectAreaType" json:"SubjectAreaType"`
}

type subjectareatype struct {
        Code *String `xml:"Code" json:"Code"`
      OtherCodeList *OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList,omitempty"`
      
      }
    
    type ACStrandSubjectAreaType struct {
  acstrandsubjectareatype `xml:"ACStrandSubjectAreaType" json:"ACStrandSubjectAreaType"`
}

type acstrandsubjectareatype struct {
        ACStrand *AUCodeSetsACStrandType `xml:"ACStrand" json:"ACStrand"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea,omitempty"`
      
      }
    
    type EducationFilterType struct {
  educationfiltertype `xml:"EducationFilterType" json:"EducationFilterType"`
}

type educationfiltertype struct {
        LearningStandardItems *LearningStandardsType `xml:"LearningStandardItems,omitempty" json:"LearningStandardItems,omitempty"`
      
      }
    
    type SIF_ExtendedElementsType struct {
  sif_extendedelementstype `xml:"SIF_ExtendedElementsType" json:"SIF_ExtendedElementsType"`
}

type sif_extendedelementstype struct {
      SIF_ExtendedElement []SIF_ExtendedElementsType_SIF_ExtendedElement `xml:"SIF_ExtendedElement" json:"SIF_ExtendedElement"`
      
      }
    
    type SIF_MetadataType struct {
  sif_metadatatype `xml:"SIF_MetadataType" json:"SIF_MetadataType"`
}

type sif_metadatatype struct {
      TimeElements *SIF_MetadataType_TimeElements
      LifeCycle *LifeCycleType `xml:"LifeCycle,omitempty" json:"LifeCycle,omitempty"`
      EducationFilter *EducationFilterType `xml:"EducationFilter,omitempty" json:"EducationFilter,omitempty"`
      
      }
    type AbstractContentElementType_XMLData struct {
  abstractcontentelementtype_xmldata `xml:"AbstractContentElementType_XMLData" json:"AbstractContentElementType_XMLData"`
}

type abstractcontentelementtype_xmldata struct {
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_TextData struct {
  abstractcontentelementtype_textdata `xml:"AbstractContentElementType_TextData" json:"AbstractContentElementType_TextData"`
}

type abstractcontentelementtype_textdata struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_BinaryData struct {
  abstractcontentelementtype_binarydata `xml:"AbstractContentElementType_BinaryData" json:"AbstractContentElementType_BinaryData"`
}

type abstractcontentelementtype_binarydata struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      FileName *String `xml:"FileName,attr" json:"FileName"`
      Description *String `xml:"Description,attr" json:"Description"`
      Value *String `xml:",chardata" json:"value"`
}
type AbstractContentElementType_Reference struct {
  abstractcontentelementtype_reference `xml:"AbstractContentElementType_Reference" json:"AbstractContentElementType_Reference"`
}

type abstractcontentelementtype_reference struct {
      MIMEType *String `xml:"MIMEType,attr" json:"MIMEType"`
      Description *String `xml:"Description,attr" json:"Description"`
       URL *String `xml:"URL" json:"URL"`
}
type PersonInvolvementType_PersonRefId struct {
  personinvolvementtype_personrefid `xml:"PersonInvolvementType_PersonRefId" json:"PersonInvolvementType_PersonRefId"`
}

type personinvolvementtype_personrefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type ActivityTimeType_Duration struct {
  activitytimetype_duration `xml:"ActivityTimeType_Duration" json:"ActivityTimeType_Duration"`
}

type activitytimetype_duration struct {
      Units *String `xml:"Units,attr" json:"Units"`
      Value *Int `xml:",chardata" json:"value"`
}
type SourceObjectsType_SourceObject struct {
  sourceobjectstype_sourceobject `xml:"SourceObjectsType_SourceObject" json:"SourceObjectsType_SourceObject"`
}

type sourceobjectstype_sourceobject struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type IdentityAssertionsType_IdentityAssertion struct {
  identityassertionstype_identityassertion `xml:"IdentityAssertionsType_IdentityAssertion" json:"IdentityAssertionsType_IdentityAssertion"`
}

type identityassertionstype_identityassertion struct {
      SchemaName *String `xml:"SchemaName,attr" json:"SchemaName"`
      Value *String `xml:",chardata" json:"value"`
}
type AssociatedObjectsType_AssociatedObject struct {
  associatedobjectstype_associatedobject `xml:"AssociatedObjectsType_AssociatedObject" json:"AssociatedObjectsType_AssociatedObject"`
}

type associatedobjectstype_associatedobject struct {
      SIF_RefObject *ObjectNameType `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type PasswordListType_Password struct {
  passwordlisttype_password `xml:"PasswordListType_Password" json:"PasswordListType_Password"`
}

type passwordlisttype_password struct {
      Algorithm *String `xml:"Algorithm,attr" json:"Algorithm"`
      KeyName *String `xml:"KeyName,attr" json:"KeyName"`
      Value *String `xml:",chardata" json:"value"`
}
type LocationType_LocationRefId struct {
  locationtype_locationrefid `xml:"LocationType_LocationRefId" json:"LocationType_LocationRefId"`
}

type locationtype_locationrefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type TimeElementType_SpanGaps struct {
  timeelementtype_spangaps `xml:"TimeElementType_SpanGaps" json:"TimeElementType_SpanGaps"`
}

type timeelementtype_spangaps struct {
      SpanGap []TimeElementType_SpanGap `xml:"SpanGap" json:"SpanGap"`
}
type LifeCycleType_Created struct {
  lifecycletype_created `xml:"LifeCycleType_Created" json:"LifeCycleType_Created"`
}

type lifecycletype_created struct {
       DateTime *String `xml:"DateTime" json:"DateTime"`
      Creators *LifeCycleType_Creators
}
type LifeCycleType_ModificationHistory struct {
  lifecycletype_modificationhistory `xml:"LifeCycleType_ModificationHistory" json:"LifeCycleType_ModificationHistory"`
}

type lifecycletype_modificationhistory struct {
      Modified []LifeCycleType_Modified `xml:"Modified" json:"Modified"`
}
type LifeCycleType_TimeElements struct {
  lifecycletype_timeelements `xml:"LifeCycleType_TimeElements" json:"LifeCycleType_TimeElements"`
}

type lifecycletype_timeelements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type OtherCodeListType_OtherCode struct {
  othercodelisttype_othercode `xml:"OtherCodeListType_OtherCode" json:"OtherCodeListType_OtherCode"`
}

type othercodelisttype_othercode struct {
      Codeset *String `xml:"Codeset,attr" json:"Codeset"`
      Value *String `xml:",chardata" json:"value"`
}
type SIF_ExtendedElementsType_SIF_ExtendedElement struct {
  sif_extendedelementstype_sif_extendedelement `xml:"SIF_ExtendedElementsType_SIF_ExtendedElement" json:"SIF_ExtendedElementsType_SIF_ExtendedElement"`
}

type sif_extendedelementstype_sif_extendedelement struct {
      Name *String `xml:"Name,attr" json:"Name"`
      Type *String `xml:"Type,attr" json:"Type"`
      SIF_Action *String `xml:"SIF_Action,attr" json:"SIF_Action"`
      Value *ExtendedContentType `xml:",chardata" json:"value"`
}
type SIF_MetadataType_TimeElements struct {
  sif_metadatatype_timeelements `xml:"SIF_MetadataType_TimeElements" json:"SIF_MetadataType_TimeElements"`
}

type sif_metadatatype_timeelements struct {
       TimeElement []TimeElementType `xml:"TimeElement" json:"TimeElement"`
}
type TimeElementType_SpanGap struct {
  timeelementtype_spangap `xml:"TimeElementType_SpanGap" json:"TimeElementType_SpanGap"`
}

type timeelementtype_spangap struct {
       Type *String `xml:"Type" json:"Type"`
       Code *String `xml:"Code" json:"Code"`
       Name *String `xml:"Name" json:"Name"`
       Value *String `xml:"Value" json:"Value"`
       StartDateTime *String `xml:"StartDateTime,omitempty" json:"StartDateTime,omitempty"`
       EndDateTime *String `xml:"EndDateTime,omitempty" json:"EndDateTime,omitempty"`
}
type LifeCycleType_Creators struct {
  lifecycletype_creators `xml:"LifeCycleType_Creators" json:"LifeCycleType_Creators"`
}

type lifecycletype_creators struct {
      Creator []LifeCycleType_Creator `xml:"Creator" json:"Creator"`
}
type LifeCycleType_Modified struct {
  lifecycletype_modified `xml:"LifeCycleType_Modified" json:"LifeCycleType_Modified"`
}

type lifecycletype_modified struct {
       By *String `xml:"By" json:"By"`
       DateTime *String `xml:"DateTime" json:"DateTime"`
       Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
}
type LifeCycleType_Creator struct {
  lifecycletype_creator `xml:"LifeCycleType_Creator" json:"LifeCycleType_Creator"`
}

type lifecycletype_creator struct {
       Name *String `xml:"Name" json:"Name"`
       ID *String `xml:"ID" json:"ID"`
}
type String string
type Int int
type Float float64
type Bool bool
