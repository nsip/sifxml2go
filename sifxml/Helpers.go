package sifxml

import (
  "log"
  )

func IntCreate(x int) *int {
  return &x
}

func FloatCreate(x float64) *float64 {
  return &x
}

func StringCreate(x string) *string {
  return &x
}

func BoolCreate(x bool) *bool {
  return &x
}

func HouseholdContactInfoTypeCreate(x HouseholdContactInfoType) *HouseholdContactInfoType {
        return &x
}

func (n *HouseholdContactInfoType) New() *HouseholdContactInfoType {
        return HouseholdContactInfoTypeCreate(HouseholdContactInfoType{})
}

func SchoolContactListTypeCreate(x SchoolContactListType) *SchoolContactListType {
        return &x
}

func (n *SchoolContactListType) New() *SchoolContactListType {
        return SchoolContactListTypeCreate(SchoolContactListType{})
}

func TimeElementType_SpanGapsCreate(x TimeElementType_SpanGaps) *TimeElementType_SpanGaps {
        return &x
}

func (n *TimeElementType_SpanGaps) New() *TimeElementType_SpanGaps {
        return TimeElementType_SpanGapsCreate(TimeElementType_SpanGaps{})
}

func PasswordListType_PasswordCreate(x PasswordListType_Password) *PasswordListType_Password {
        return &x
}

func (n *PasswordListType_Password) New() *PasswordListType_Password {
        return PasswordListType_PasswordCreate(PasswordListType_Password{})
}

func ResourcesTypeCreate(x ResourcesType) *ResourcesType {
        return &x
}

func (n *ResourcesType) New() *ResourcesType {
        return ResourcesTypeCreate(ResourcesType{})
}

func NAPLANClassListTypeCreate(x NAPLANClassListType) *NAPLANClassListType {
        return &x
}

func (n *NAPLANClassListType) New() *NAPLANClassListType {
        return NAPLANClassListTypeCreate(NAPLANClassListType{})
}

func StudentActivityInfoCreate(x StudentActivityInfo) *StudentActivityInfo {
        return &x
}

func (n *StudentActivityInfo) New() *StudentActivityInfo {
        return StudentActivityInfoCreate(StudentActivityInfo{})
}

func CensusCollectionCreate(x CensusCollection) *CensusCollection {
        return &x
}

func (n *CensusCollection) New() *CensusCollection {
        return CensusCollectionCreate(CensusCollection{})
}

func LifeCycleType_CreatedCreate(x LifeCycleType_Created) *LifeCycleType_Created {
        return &x
}

func (n *LifeCycleType_Created) New() *LifeCycleType_Created {
        return LifeCycleType_CreatedCreate(LifeCycleType_Created{})
}

func MonetaryAmountTypeCreate(x MonetaryAmountType) *MonetaryAmountType {
        return &x
}

func (n *MonetaryAmountType) New() *MonetaryAmountType {
        return MonetaryAmountTypeCreate(MonetaryAmountType{})
}

func StaffPersonalCreate(x StaffPersonal) *StaffPersonal {
        return &x
}

func (n *StaffPersonal) New() *StaffPersonal {
        return StaffPersonalCreate(StaffPersonal{})
}

func AwardContainerTypeCreate(x AwardContainerType) *AwardContainerType {
        return &x
}

func (n *AwardContainerType) New() *AwardContainerType {
        return AwardContainerTypeCreate(AwardContainerType{})
}

func StudentSchoolEnrollment_HomeroomCreate(x StudentSchoolEnrollment_Homeroom) *StudentSchoolEnrollment_Homeroom {
        return &x
}

func (n *StudentSchoolEnrollment_Homeroom) New() *StudentSchoolEnrollment_Homeroom {
        return StudentSchoolEnrollment_HomeroomCreate(StudentSchoolEnrollment_Homeroom{})
}

func YearRangeTypeCreate(x YearRangeType) *YearRangeType {
        return &x
}

func (n *YearRangeType) New() *YearRangeType {
        return YearRangeTypeCreate(YearRangeType{})
}

func AbstractContentElementType_TextDataCreate(x AbstractContentElementType_TextData) *AbstractContentElementType_TextData {
        return &x
}

func (n *AbstractContentElementType_TextData) New() *AbstractContentElementType_TextData {
        return AbstractContentElementType_TextDataCreate(AbstractContentElementType_TextData{})
}

func ResourceUsage_ResourceReportLineListCreate(x ResourceUsage_ResourceReportLineList) *ResourceUsage_ResourceReportLineList {
        return &x
}

func (n *ResourceUsage_ResourceReportLineList) New() *ResourceUsage_ResourceReportLineList {
        return ResourceUsage_ResourceReportLineListCreate(ResourceUsage_ResourceReportLineList{})
}

func LEAContactTypeCreate(x LEAContactType) *LEAContactType {
        return &x
}

func (n *LEAContactType) New() *LEAContactType {
        return LEAContactTypeCreate(LEAContactType{})
}

func TimeTableDayListTypeCreate(x TimeTableDayListType) *TimeTableDayListType {
        return &x
}

func (n *TimeTableDayListType) New() *TimeTableDayListType {
        return TimeTableDayListTypeCreate(TimeTableDayListType{})
}

func NAPTestletCodeFrameTypeCreate(x NAPTestletCodeFrameType) *NAPTestletCodeFrameType {
        return &x
}

func (n *NAPTestletCodeFrameType) New() *NAPTestletCodeFrameType {
        return NAPTestletCodeFrameTypeCreate(NAPTestletCodeFrameType{})
}

func EducationFilterTypeCreate(x EducationFilterType) *EducationFilterType {
        return &x
}

func (n *EducationFilterType) New() *EducationFilterType {
        return EducationFilterTypeCreate(EducationFilterType{})
}

func ComponentsTypeCreate(x ComponentsType) *ComponentsType {
        return &x
}

func (n *ComponentsType) New() *ComponentsType {
        return ComponentsTypeCreate(ComponentsType{})
}

func LibraryMessageTypeCreate(x LibraryMessageType) *LibraryMessageType {
        return &x
}

func (n *LibraryMessageType) New() *LibraryMessageType {
        return LibraryMessageTypeCreate(LibraryMessageType{})
}

func CalendarSummaryListTypeCreate(x CalendarSummaryListType) *CalendarSummaryListType {
        return &x
}

func (n *CalendarSummaryListType) New() *CalendarSummaryListType {
        return CalendarSummaryListTypeCreate(CalendarSummaryListType{})
}

func ExpenseAccountTypeCreate(x ExpenseAccountType) *ExpenseAccountType {
        return &x
}

func (n *ExpenseAccountType) New() *ExpenseAccountType {
        return ExpenseAccountTypeCreate(ExpenseAccountType{})
}

func AbstractContentElementType_XMLDataCreate(x AbstractContentElementType_XMLData) *AbstractContentElementType_XMLData {
        return &x
}

func (n *AbstractContentElementType_XMLData) New() *AbstractContentElementType_XMLData {
        return AbstractContentElementType_XMLDataCreate(AbstractContentElementType_XMLData{})
}

func FQItemListTypeCreate(x FQItemListType) *FQItemListType {
        return &x
}

func (n *FQItemListType) New() *FQItemListType {
        return FQItemListTypeCreate(FQItemListType{})
}

func MedicalAlertMessageTypeCreate(x MedicalAlertMessageType) *MedicalAlertMessageType {
        return &x
}

func (n *MedicalAlertMessageType) New() *MedicalAlertMessageType {
        return MedicalAlertMessageTypeCreate(MedicalAlertMessageType{})
}

func CountryListTypeCreate(x CountryListType) *CountryListType {
        return &x
}

func (n *CountryListType) New() *CountryListType {
        return CountryListTypeCreate(CountryListType{})
}

func StaffActivityExtensionTypeCreate(x StaffActivityExtensionType) *StaffActivityExtensionType {
        return &x
}

func (n *StaffActivityExtensionType) New() *StaffActivityExtensionType {
        return StaffActivityExtensionTypeCreate(StaffActivityExtensionType{})
}

func AccountCodeListTypeCreate(x AccountCodeListType) *AccountCodeListType {
        return &x
}

func (n *AccountCodeListType) New() *AccountCodeListType {
        return AccountCodeListTypeCreate(AccountCodeListType{})
}

func StudentSchoolEnrollment_CounselorCreate(x StudentSchoolEnrollment_Counselor) *StudentSchoolEnrollment_Counselor {
        return &x
}

func (n *StudentSchoolEnrollment_Counselor) New() *StudentSchoolEnrollment_Counselor {
        return StudentSchoolEnrollment_CounselorCreate(StudentSchoolEnrollment_Counselor{})
}

func PaymentReceiptLineListTypeCreate(x PaymentReceiptLineListType) *PaymentReceiptLineListType {
        return &x
}

func (n *PaymentReceiptLineListType) New() *PaymentReceiptLineListType {
        return PaymentReceiptLineListTypeCreate(PaymentReceiptLineListType{})
}

func EmailListTypeCreate(x EmailListType) *EmailListType {
        return &x
}

func (n *EmailListType) New() *EmailListType {
        return EmailListTypeCreate(EmailListType{})
}

func InvoiceCreate(x Invoice) *Invoice {
        return &x
}

func (n *Invoice) New() *Invoice {
        return InvoiceCreate(Invoice{})
}

func CatchmentStatusContainerTypeCreate(x CatchmentStatusContainerType) *CatchmentStatusContainerType {
        return &x
}

func (n *CatchmentStatusContainerType) New() *CatchmentStatusContainerType {
        return CatchmentStatusContainerTypeCreate(CatchmentStatusContainerType{})
}

func StimulusLocalIdListTypeCreate(x StimulusLocalIdListType) *StimulusLocalIdListType {
        return &x
}

func (n *StimulusLocalIdListType) New() *StimulusLocalIdListType {
        return StimulusLocalIdListTypeCreate(StimulusLocalIdListType{})
}

func StudentParticipation_ManagingSchoolCreate(x StudentParticipation_ManagingSchool) *StudentParticipation_ManagingSchool {
        return &x
}

func (n *StudentParticipation_ManagingSchool) New() *StudentParticipation_ManagingSchool {
        return StudentParticipation_ManagingSchoolCreate(StudentParticipation_ManagingSchool{})
}

func AGReportingObjectResponseListTypeCreate(x AGReportingObjectResponseListType) *AGReportingObjectResponseListType {
        return &x
}

func (n *AGReportingObjectResponseListType) New() *AGReportingObjectResponseListType {
        return AGReportingObjectResponseListTypeCreate(AGReportingObjectResponseListType{})
}

func AGRuleListTypeCreate(x AGRuleListType) *AGRuleListType {
        return &x
}

func (n *AGRuleListType) New() *AGRuleListType {
        return AGRuleListTypeCreate(AGRuleListType{})
}

func SubjectAreaTypeCreate(x SubjectAreaType) *SubjectAreaType {
        return &x
}

func (n *SubjectAreaType) New() *SubjectAreaType {
        return SubjectAreaTypeCreate(SubjectAreaType{})
}

func StudentParticipationCreate(x StudentParticipation) *StudentParticipation {
        return &x
}

func (n *StudentParticipation) New() *StudentParticipation {
        return StudentParticipationCreate(StudentParticipation{})
}

func AssignmentListTypeCreate(x AssignmentListType) *AssignmentListType {
        return &x
}

func (n *AssignmentListType) New() *AssignmentListType {
        return AssignmentListTypeCreate(AssignmentListType{})
}

func AddressCollectionStudentTypeCreate(x AddressCollectionStudentType) *AddressCollectionStudentType {
        return &x
}

func (n *AddressCollectionStudentType) New() *AddressCollectionStudentType {
        return AddressCollectionStudentTypeCreate(AddressCollectionStudentType{})
}

func WellbeingEventCategoryListTypeCreate(x WellbeingEventCategoryListType) *WellbeingEventCategoryListType {
        return &x
}

func (n *WellbeingEventCategoryListType) New() *WellbeingEventCategoryListType {
        return WellbeingEventCategoryListTypeCreate(WellbeingEventCategoryListType{})
}

func ReferralSourceTypeCreate(x ReferralSourceType) *ReferralSourceType {
        return &x
}

func (n *ReferralSourceType) New() *ReferralSourceType {
        return ReferralSourceTypeCreate(ReferralSourceType{})
}

func OtherCodeListTypeCreate(x OtherCodeListType) *OtherCodeListType {
        return &x
}

func (n *OtherCodeListType) New() *OtherCodeListType {
        return OtherCodeListTypeCreate(OtherCodeListType{})
}

func StudentSubjectChoiceTypeCreate(x StudentSubjectChoiceType) *StudentSubjectChoiceType {
        return &x
}

func (n *StudentSubjectChoiceType) New() *StudentSubjectChoiceType {
        return StudentSubjectChoiceTypeCreate(StudentSubjectChoiceType{})
}

func StudentsTypeCreate(x StudentsType) *StudentsType {
        return &x
}

func (n *StudentsType) New() *StudentsType {
        return StudentsTypeCreate(StudentsType{})
}

func LifeCycleTypeCreate(x LifeCycleType) *LifeCycleType {
        return &x
}

func (n *LifeCycleType) New() *LifeCycleType {
        return LifeCycleTypeCreate(LifeCycleType{})
}

func LifeCycleType_ModificationHistoryCreate(x LifeCycleType_ModificationHistory) *LifeCycleType_ModificationHistory {
        return &x
}

func (n *LifeCycleType_ModificationHistory) New() *LifeCycleType_ModificationHistory {
        return LifeCycleType_ModificationHistoryCreate(LifeCycleType_ModificationHistory{})
}

func DomainScoreTypeCreate(x DomainScoreType) *DomainScoreType {
        return &x
}

func (n *DomainScoreType) New() *DomainScoreType {
        return DomainScoreTypeCreate(DomainScoreType{})
}

func LocationType_LocationRefIdCreate(x LocationType_LocationRefId) *LocationType_LocationRefId {
        return &x
}

func (n *LocationType_LocationRefId) New() *LocationType_LocationRefId {
        return LocationType_LocationRefIdCreate(LocationType_LocationRefId{})
}

func CampusContainerTypeCreate(x CampusContainerType) *CampusContainerType {
        return &x
}

func (n *CampusContainerType) New() *CampusContainerType {
        return CampusContainerTypeCreate(CampusContainerType{})
}

func AggregateStatisticInfoCreate(x AggregateStatisticInfo) *AggregateStatisticInfo {
        return &x
}

func (n *AggregateStatisticInfo) New() *AggregateStatisticInfo {
        return AggregateStatisticInfoCreate(AggregateStatisticInfo{})
}

func StatsCohortListTypeCreate(x StatsCohortListType) *StatsCohortListType {
        return &x
}

func (n *StatsCohortListType) New() *StatsCohortListType {
        return StatsCohortListTypeCreate(StatsCohortListType{})
}

func LifeCycleType_TimeElementsCreate(x LifeCycleType_TimeElements) *LifeCycleType_TimeElements {
        return &x
}

func (n *LifeCycleType_TimeElements) New() *LifeCycleType_TimeElements {
        return LifeCycleType_TimeElementsCreate(LifeCycleType_TimeElements{})
}

func CollectionStatusCreate(x CollectionStatus) *CollectionStatus {
        return &x
}

func (n *CollectionStatus) New() *CollectionStatus {
        return CollectionStatusCreate(CollectionStatus{})
}

func PersonalisedPlanCreate(x PersonalisedPlan) *PersonalisedPlan {
        return &x
}

func (n *PersonalisedPlan) New() *PersonalisedPlan {
        return PersonalisedPlanCreate(PersonalisedPlan{})
}

func SIF_ExtendedElementsType_SIF_ExtendedElementCreate(x SIF_ExtendedElementsType_SIF_ExtendedElement) *SIF_ExtendedElementsType_SIF_ExtendedElement {
        return &x
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) New() *SIF_ExtendedElementsType_SIF_ExtendedElement {
        return SIF_ExtendedElementsType_SIF_ExtendedElementCreate(SIF_ExtendedElementsType_SIF_ExtendedElement{})
}

func FinancialQuestionnaireCollectionCreate(x FinancialQuestionnaireCollection) *FinancialQuestionnaireCollection {
        return &x
}

func (n *FinancialQuestionnaireCollection) New() *FinancialQuestionnaireCollection {
        return FinancialQuestionnaireCollectionCreate(FinancialQuestionnaireCollection{})
}

func JournalAdjustmentTypeCreate(x JournalAdjustmentType) *JournalAdjustmentType {
        return &x
}

func (n *JournalAdjustmentType) New() *JournalAdjustmentType {
        return JournalAdjustmentTypeCreate(JournalAdjustmentType{})
}

func WellbeingPersonLink_PersonRefIdCreate(x WellbeingPersonLink_PersonRefId) *WellbeingPersonLink_PersonRefId {
        return &x
}

func (n *WellbeingPersonLink_PersonRefId) New() *WellbeingPersonLink_PersonRefId {
        return WellbeingPersonLink_PersonRefIdCreate(WellbeingPersonLink_PersonRefId{})
}

func HoldInfoTypeCreate(x HoldInfoType) *HoldInfoType {
        return &x
}

func (n *HoldInfoType) New() *HoldInfoType {
        return HoldInfoTypeCreate(HoldInfoType{})
}

func TimeTableCreate(x TimeTable) *TimeTable {
        return &x
}

func (n *TimeTable) New() *TimeTable {
        return TimeTableCreate(TimeTable{})
}

func StudentGradeCreate(x StudentGrade) *StudentGrade {
        return &x
}

func (n *StudentGrade) New() *StudentGrade {
        return StudentGradeCreate(StudentGrade{})
}

func StandardsSettingBodyTypeCreate(x StandardsSettingBodyType) *StandardsSettingBodyType {
        return &x
}

func (n *StandardsSettingBodyType) New() *StandardsSettingBodyType {
        return StandardsSettingBodyTypeCreate(StandardsSettingBodyType{})
}

func StudentGroupListTypeCreate(x StudentGroupListType) *StudentGroupListType {
        return &x
}

func (n *StudentGroupListType) New() *StudentGroupListType {
        return StudentGroupListTypeCreate(StudentGroupListType{})
}

func OtherNamesTypeCreate(x OtherNamesType) *OtherNamesType {
        return &x
}

func (n *OtherNamesType) New() *OtherNamesType {
        return OtherNamesTypeCreate(OtherNamesType{})
}

func NameTypeCreate(x NameType) *NameType {
        return &x
}

func (n *NameType) New() *NameType {
        return NameTypeCreate(NameType{})
}

func CollectionRoundCreate(x CollectionRound) *CollectionRound {
        return &x
}

func (n *CollectionRound) New() *CollectionRound {
        return CollectionRoundCreate(CollectionRound{})
}

func CountryList2TypeCreate(x CountryList2Type) *CountryList2Type {
        return &x
}

func (n *CountryList2Type) New() *CountryList2Type {
        return CountryList2TypeCreate(CountryList2Type{})
}

func ResourceBooking_ResourceRefIdCreate(x ResourceBooking_ResourceRefId) *ResourceBooking_ResourceRefId {
        return &x
}

func (n *ResourceBooking_ResourceRefId) New() *ResourceBooking_ResourceRefId {
        return ResourceBooking_ResourceRefIdCreate(ResourceBooking_ResourceRefId{})
}

func AbstractContentElementType_BinaryDataCreate(x AbstractContentElementType_BinaryData) *AbstractContentElementType_BinaryData {
        return &x
}

func (n *AbstractContentElementType_BinaryData) New() *AbstractContentElementType_BinaryData {
        return AbstractContentElementType_BinaryDataCreate(AbstractContentElementType_BinaryData{})
}

func AdjustmentContainerTypeCreate(x AdjustmentContainerType) *AdjustmentContainerType {
        return &x
}

func (n *AdjustmentContainerType) New() *AdjustmentContainerType {
        return AdjustmentContainerTypeCreate(AdjustmentContainerType{})
}

func PurchaseOrderCreate(x PurchaseOrder) *PurchaseOrder {
        return &x
}

func (n *PurchaseOrder) New() *PurchaseOrder {
        return PurchaseOrderCreate(PurchaseOrder{})
}

func TeachingGroupPeriodListTypeCreate(x TeachingGroupPeriodListType) *TeachingGroupPeriodListType {
        return &x
}

func (n *TeachingGroupPeriodListType) New() *TeachingGroupPeriodListType {
        return TeachingGroupPeriodListTypeCreate(TeachingGroupPeriodListType{})
}

func PublishingPermissionListTypeCreate(x PublishingPermissionListType) *PublishingPermissionListType {
        return &x
}

func (n *PublishingPermissionListType) New() *PublishingPermissionListType {
        return PublishingPermissionListTypeCreate(PublishingPermissionListType{})
}

func TestDisruptionListTypeCreate(x TestDisruptionListType) *TestDisruptionListType {
        return &x
}

func (n *TestDisruptionListType) New() *TestDisruptionListType {
        return TestDisruptionListTypeCreate(TestDisruptionListType{})
}

func StudentEntryContainerTypeCreate(x StudentEntryContainerType) *StudentEntryContainerType {
        return &x
}

func (n *StudentEntryContainerType) New() *StudentEntryContainerType {
        return StudentEntryContainerTypeCreate(StudentEntryContainerType{})
}

func ExpenseAccountsTypeCreate(x ExpenseAccountsType) *ExpenseAccountsType {
        return &x
}

func (n *ExpenseAccountsType) New() *ExpenseAccountsType {
        return ExpenseAccountsTypeCreate(ExpenseAccountsType{})
}

func FQReportingListTypeCreate(x FQReportingListType) *FQReportingListType {
        return &x
}

func (n *FQReportingListType) New() *FQReportingListType {
        return FQReportingListTypeCreate(FQReportingListType{})
}

func EvaluationTypeCreate(x EvaluationType) *EvaluationType {
        return &x
}

func (n *EvaluationType) New() *EvaluationType {
        return EvaluationTypeCreate(EvaluationType{})
}

func NAPCodeFrameCreate(x NAPCodeFrame) *NAPCodeFrame {
        return &x
}

func (n *NAPCodeFrame) New() *NAPCodeFrame {
        return NAPCodeFrameCreate(NAPCodeFrame{})
}

func SectionInfoCreate(x SectionInfo) *SectionInfo {
        return &x
}

func (n *SectionInfo) New() *SectionInfo {
        return SectionInfoCreate(SectionInfo{})
}

func CodeFrameTestItemListTypeCreate(x CodeFrameTestItemListType) *CodeFrameTestItemListType {
        return &x
}

func (n *CodeFrameTestItemListType) New() *CodeFrameTestItemListType {
        return CodeFrameTestItemListTypeCreate(CodeFrameTestItemListType{})
}

func CharacteristicsTypeCreate(x CharacteristicsType) *CharacteristicsType {
        return &x
}

func (n *CharacteristicsType) New() *CharacteristicsType {
        return CharacteristicsTypeCreate(CharacteristicsType{})
}

func AttendanceCodeTypeCreate(x AttendanceCodeType) *AttendanceCodeType {
        return &x
}

func (n *AttendanceCodeType) New() *AttendanceCodeType {
        return AttendanceCodeTypeCreate(AttendanceCodeType{})
}

func EssentialMaterialsTypeCreate(x EssentialMaterialsType) *EssentialMaterialsType {
        return &x
}

func (n *EssentialMaterialsType) New() *EssentialMaterialsType {
        return EssentialMaterialsTypeCreate(EssentialMaterialsType{})
}

func ChargedLocationInfoCreate(x ChargedLocationInfo) *ChargedLocationInfo {
        return &x
}

func (n *ChargedLocationInfo) New() *ChargedLocationInfo {
        return ChargedLocationInfoCreate(ChargedLocationInfo{})
}

func StudentDailyAttendanceCreate(x StudentDailyAttendance) *StudentDailyAttendance {
        return &x
}

func (n *StudentDailyAttendance) New() *StudentDailyAttendance {
        return StudentDailyAttendanceCreate(StudentDailyAttendance{})
}

func YearLevelsTypeCreate(x YearLevelsType) *YearLevelsType {
        return &x
}

func (n *YearLevelsType) New() *YearLevelsType {
        return YearLevelsTypeCreate(YearLevelsType{})
}

func DomainProficiencyContainerTypeCreate(x DomainProficiencyContainerType) *DomainProficiencyContainerType {
        return &x
}

func (n *DomainProficiencyContainerType) New() *DomainProficiencyContainerType {
        return DomainProficiencyContainerTypeCreate(DomainProficiencyContainerType{})
}

func DwellingArrangementTypeCreate(x DwellingArrangementType) *DwellingArrangementType {
        return &x
}

func (n *DwellingArrangementType) New() *DwellingArrangementType {
        return DwellingArrangementTypeCreate(DwellingArrangementType{})
}

func AlertMessageTypeCreate(x AlertMessageType) *AlertMessageType {
        return &x
}

func (n *AlertMessageType) New() *AlertMessageType {
        return AlertMessageTypeCreate(AlertMessageType{})
}

func ActivityCreate(x Activity) *Activity {
        return &x
}

func (n *Activity) New() *Activity {
        return ActivityCreate(Activity{})
}

func WellbeingDocumentTypeCreate(x WellbeingDocumentType) *WellbeingDocumentType {
        return &x
}

func (n *WellbeingDocumentType) New() *WellbeingDocumentType {
        return WellbeingDocumentTypeCreate(WellbeingDocumentType{})
}

func StudentAttendanceTimeListCreate(x StudentAttendanceTimeList) *StudentAttendanceTimeList {
        return &x
}

func (n *StudentAttendanceTimeList) New() *StudentAttendanceTimeList {
        return StudentAttendanceTimeListCreate(StudentAttendanceTimeList{})
}

func PrerequisitesTypeCreate(x PrerequisitesType) *PrerequisitesType {
        return &x
}

func (n *PrerequisitesType) New() *PrerequisitesType {
        return PrerequisitesTypeCreate(PrerequisitesType{})
}

func StudentSchoolEnrollment_AdvisorCreate(x StudentSchoolEnrollment_Advisor) *StudentSchoolEnrollment_Advisor {
        return &x
}

func (n *StudentSchoolEnrollment_Advisor) New() *StudentSchoolEnrollment_Advisor {
        return StudentSchoolEnrollment_AdvisorCreate(StudentSchoolEnrollment_Advisor{})
}

func NAPTestItem2TypeCreate(x NAPTestItem2Type) *NAPTestItem2Type {
        return &x
}

func (n *NAPTestItem2Type) New() *NAPTestItem2Type {
        return NAPTestItem2TypeCreate(NAPTestItem2Type{})
}

func ScheduledActivityOverrideTypeCreate(x ScheduledActivityOverrideType) *ScheduledActivityOverrideType {
        return &x
}

func (n *ScheduledActivityOverrideType) New() *ScheduledActivityOverrideType {
        return ScheduledActivityOverrideTypeCreate(ScheduledActivityOverrideType{})
}

func VisaSubClassTypeCreate(x VisaSubClassType) *VisaSubClassType {
        return &x
}

func (n *VisaSubClassType) New() *VisaSubClassType {
        return VisaSubClassTypeCreate(VisaSubClassType{})
}

func AGParentTypeCreate(x AGParentType) *AGParentType {
        return &x
}

func (n *AGParentType) New() *AGParentType {
        return AGParentTypeCreate(AGParentType{})
}

func FollowUpActionTypeCreate(x FollowUpActionType) *FollowUpActionType {
        return &x
}

func (n *FollowUpActionType) New() *FollowUpActionType {
        return FollowUpActionTypeCreate(FollowUpActionType{})
}

func SIF_MetadataTypeCreate(x SIF_MetadataType) *SIF_MetadataType {
        return &x
}

func (n *SIF_MetadataType) New() *SIF_MetadataType {
        return SIF_MetadataTypeCreate(SIF_MetadataType{})
}

func RelatedLearningStandardItemRefIdListTypeCreate(x RelatedLearningStandardItemRefIdListType) *RelatedLearningStandardItemRefIdListType {
        return &x
}

func (n *RelatedLearningStandardItemRefIdListType) New() *RelatedLearningStandardItemRefIdListType {
        return RelatedLearningStandardItemRefIdListTypeCreate(RelatedLearningStandardItemRefIdListType{})
}

func CensusStaffTypeCreate(x CensusStaffType) *CensusStaffType {
        return &x
}

func (n *CensusStaffType) New() *CensusStaffType {
        return CensusStaffTypeCreate(CensusStaffType{})
}

func GradeTypeCreate(x GradeType) *GradeType {
        return &x
}

func (n *GradeType) New() *GradeType {
        return GradeTypeCreate(GradeType{})
}

func GradingAssignmentCreate(x GradingAssignment) *GradingAssignment {
        return &x
}

func (n *GradingAssignment) New() *GradingAssignment {
        return GradingAssignmentCreate(GradingAssignment{})
}

func FQContextualQuestionTypeCreate(x FQContextualQuestionType) *FQContextualQuestionType {
        return &x
}

func (n *FQContextualQuestionType) New() *FQContextualQuestionType {
        return FQContextualQuestionTypeCreate(FQContextualQuestionType{})
}

func PublishingPermissionTypeCreate(x PublishingPermissionType) *PublishingPermissionType {
        return &x
}

func (n *PublishingPermissionType) New() *PublishingPermissionType {
        return PublishingPermissionTypeCreate(PublishingPermissionType{})
}

func WellbeingResponseCreate(x WellbeingResponse) *WellbeingResponse {
        return &x
}

func (n *WellbeingResponse) New() *WellbeingResponse {
        return WellbeingResponseCreate(WellbeingResponse{})
}

func AddressCollectionReportingTypeCreate(x AddressCollectionReportingType) *AddressCollectionReportingType {
        return &x
}

func (n *AddressCollectionReportingType) New() *AddressCollectionReportingType {
        return AddressCollectionReportingTypeCreate(AddressCollectionReportingType{})
}

func SymptomListTypeCreate(x SymptomListType) *SymptomListType {
        return &x
}

func (n *SymptomListType) New() *SymptomListType {
        return SymptomListTypeCreate(SymptomListType{})
}

func AbstractContentPackageType_ReferenceCreate(x AbstractContentPackageType_Reference) *AbstractContentPackageType_Reference {
        return &x
}

func (n *AbstractContentPackageType_Reference) New() *AbstractContentPackageType_Reference {
        return AbstractContentPackageType_ReferenceCreate(AbstractContentPackageType_Reference{})
}

func DetentionContainerTypeCreate(x DetentionContainerType) *DetentionContainerType {
        return &x
}

func (n *DetentionContainerType) New() *DetentionContainerType {
        return DetentionContainerTypeCreate(DetentionContainerType{})
}

func PersonInvolvementListTypeCreate(x PersonInvolvementListType) *PersonInvolvementListType {
        return &x
}

func (n *PersonInvolvementListType) New() *PersonInvolvementListType {
        return PersonInvolvementListTypeCreate(PersonInvolvementListType{})
}

func FQContextualQuestionListTypeCreate(x FQContextualQuestionListType) *FQContextualQuestionListType {
        return &x
}

func (n *FQContextualQuestionListType) New() *FQContextualQuestionListType {
        return FQContextualQuestionListTypeCreate(FQContextualQuestionListType{})
}

func AddressCollectionReportingListTypeCreate(x AddressCollectionReportingListType) *AddressCollectionReportingListType {
        return &x
}

func (n *AddressCollectionReportingListType) New() *AddressCollectionReportingListType {
        return AddressCollectionReportingListTypeCreate(AddressCollectionReportingListType{})
}

func TeachingGroupPeriodTypeCreate(x TeachingGroupPeriodType) *TeachingGroupPeriodType {
        return &x
}

func (n *TeachingGroupPeriodType) New() *TeachingGroupPeriodType {
        return TeachingGroupPeriodTypeCreate(TeachingGroupPeriodType{})
}

func StimulusListTypeCreate(x StimulusListType) *StimulusListType {
        return &x
}

func (n *StimulusListType) New() *StimulusListType {
        return StimulusListTypeCreate(StimulusListType{})
}

func NAPStudentResponseTestletListTypeCreate(x NAPStudentResponseTestletListType) *NAPStudentResponseTestletListType {
        return &x
}

func (n *NAPStudentResponseTestletListType) New() *NAPStudentResponseTestletListType {
        return NAPStudentResponseTestletListTypeCreate(NAPStudentResponseTestletListType{})
}

func CensusReportingTypeCreate(x CensusReportingType) *CensusReportingType {
        return &x
}

func (n *CensusReportingType) New() *CensusReportingType {
        return CensusReportingTypeCreate(CensusReportingType{})
}

func PurchasingItemsTypeCreate(x PurchasingItemsType) *PurchasingItemsType {
        return &x
}

func (n *PurchasingItemsType) New() *PurchasingItemsType {
        return PurchasingItemsTypeCreate(PurchasingItemsType{})
}

func HouseholdContactInfoListTypeCreate(x HouseholdContactInfoListType) *HouseholdContactInfoListType {
        return &x
}

func (n *HouseholdContactInfoListType) New() *HouseholdContactInfoListType {
        return HouseholdContactInfoListTypeCreate(HouseholdContactInfoListType{})
}

func AggregateCharacteristicInfoCreate(x AggregateCharacteristicInfo) *AggregateCharacteristicInfo {
        return &x
}

func (n *AggregateCharacteristicInfo) New() *AggregateCharacteristicInfo {
        return AggregateCharacteristicInfoCreate(AggregateCharacteristicInfo{})
}

func AttendanceTimeTypeCreate(x AttendanceTimeType) *AttendanceTimeType {
        return &x
}

func (n *AttendanceTimeType) New() *AttendanceTimeType {
        return AttendanceTimeTypeCreate(AttendanceTimeType{})
}

func SoftwareVendorInfoContainerTypeCreate(x SoftwareVendorInfoContainerType) *SoftwareVendorInfoContainerType {
        return &x
}

func (n *SoftwareVendorInfoContainerType) New() *SoftwareVendorInfoContainerType {
        return SoftwareVendorInfoContainerTypeCreate(SoftwareVendorInfoContainerType{})
}

func NAPSubscoreListTypeCreate(x NAPSubscoreListType) *NAPSubscoreListType {
        return &x
}

func (n *NAPSubscoreListType) New() *NAPSubscoreListType {
        return NAPSubscoreListTypeCreate(NAPSubscoreListType{})
}

func LEAInfoCreate(x LEAInfo) *LEAInfo {
        return &x
}

func (n *LEAInfo) New() *LEAInfo {
        return LEAInfoCreate(LEAInfo{})
}

func AggregateStatisticFactCreate(x AggregateStatisticFact) *AggregateStatisticFact {
        return &x
}

func (n *AggregateStatisticFact) New() *AggregateStatisticFact {
        return AggregateStatisticFactCreate(AggregateStatisticFact{})
}

func YearLevelEnrollmentListTypeCreate(x YearLevelEnrollmentListType) *YearLevelEnrollmentListType {
        return &x
}

func (n *YearLevelEnrollmentListType) New() *YearLevelEnrollmentListType {
        return YearLevelEnrollmentListTypeCreate(YearLevelEnrollmentListType{})
}

func LocationOfInstructionTypeCreate(x LocationOfInstructionType) *LocationOfInstructionType {
        return &x
}

func (n *LocationOfInstructionType) New() *LocationOfInstructionType {
        return LocationOfInstructionTypeCreate(LocationOfInstructionType{})
}

func ReligiousEventListTypeCreate(x ReligiousEventListType) *ReligiousEventListType {
        return &x
}

func (n *ReligiousEventListType) New() *ReligiousEventListType {
        return ReligiousEventListTypeCreate(ReligiousEventListType{})
}

func GradingAssignmentScoreCreate(x GradingAssignmentScore) *GradingAssignmentScore {
        return &x
}

func (n *GradingAssignmentScore) New() *GradingAssignmentScore {
        return GradingAssignmentScoreCreate(GradingAssignmentScore{})
}

func CreationUserTypeCreate(x CreationUserType) *CreationUserType {
        return &x
}

func (n *CreationUserType) New() *CreationUserType {
        return CreationUserTypeCreate(CreationUserType{})
}

func LibraryTransactionTypeCreate(x LibraryTransactionType) *LibraryTransactionType {
        return &x
}

func (n *LibraryTransactionType) New() *LibraryTransactionType {
        return LibraryTransactionTypeCreate(LibraryTransactionType{})
}

func PaymentReceiptCreate(x PaymentReceipt) *PaymentReceipt {
        return &x
}

func (n *PaymentReceipt) New() *PaymentReceipt {
        return PaymentReceiptCreate(PaymentReceipt{})
}

func ContentDescriptionListTypeCreate(x ContentDescriptionListType) *ContentDescriptionListType {
        return &x
}

func (n *ContentDescriptionListType) New() *ContentDescriptionListType {
        return ContentDescriptionListTypeCreate(ContentDescriptionListType{})
}

func TimeTableSubjectCreate(x TimeTableSubject) *TimeTableSubject {
        return &x
}

func (n *TimeTableSubject) New() *TimeTableSubject {
        return TimeTableSubjectCreate(TimeTableSubject{})
}

func SourceObjectsType_SourceObjectCreate(x SourceObjectsType_SourceObject) *SourceObjectsType_SourceObject {
        return &x
}

func (n *SourceObjectsType_SourceObject) New() *SourceObjectsType_SourceObject {
        return SourceObjectsType_SourceObjectCreate(SourceObjectsType_SourceObject{})
}

func CalendarSummaryCreate(x CalendarSummary) *CalendarSummary {
        return &x
}

func (n *CalendarSummary) New() *CalendarSummary {
        return CalendarSummaryCreate(CalendarSummary{})
}

func LocationTypeCreate(x LocationType) *LocationType {
        return &x
}

func (n *LocationType) New() *LocationType {
        return LocationTypeCreate(LocationType{})
}

func StudentContactPersonalCreate(x StudentContactPersonal) *StudentContactPersonal {
        return &x
}

func (n *StudentContactPersonal) New() *StudentContactPersonal {
        return StudentContactPersonalCreate(StudentContactPersonal{})
}

func ValidLetterMarkTypeCreate(x ValidLetterMarkType) *ValidLetterMarkType {
        return &x
}

func (n *ValidLetterMarkType) New() *ValidLetterMarkType {
        return ValidLetterMarkTypeCreate(ValidLetterMarkType{})
}

func ElectronicIdListTypeCreate(x ElectronicIdListType) *ElectronicIdListType {
        return &x
}

func (n *ElectronicIdListType) New() *ElectronicIdListType {
        return ElectronicIdListTypeCreate(ElectronicIdListType{})
}

func ActivityTimeTypeCreate(x ActivityTimeType) *ActivityTimeType {
        return &x
}

func (n *ActivityTimeType) New() *ActivityTimeType {
        return ActivityTimeTypeCreate(ActivityTimeType{})
}

func TeachingGroupCreate(x TeachingGroup) *TeachingGroup {
        return &x
}

func (n *TeachingGroup) New() *TeachingGroup {
        return TeachingGroupCreate(TeachingGroup{})
}

func StudentSectionEnrollmentCreate(x StudentSectionEnrollment) *StudentSectionEnrollment {
        return &x
}

func (n *StudentSectionEnrollment) New() *StudentSectionEnrollment {
        return StudentSectionEnrollmentCreate(StudentSectionEnrollment{})
}

func StudentAttendanceCollectionReportingListTypeCreate(x StudentAttendanceCollectionReportingListType) *StudentAttendanceCollectionReportingListType {
        return &x
}

func (n *StudentAttendanceCollectionReportingListType) New() *StudentAttendanceCollectionReportingListType {
        return StudentAttendanceCollectionReportingListTypeCreate(StudentAttendanceCollectionReportingListType{})
}

func MediumOfInstructionTypeCreate(x MediumOfInstructionType) *MediumOfInstructionType {
        return &x
}

func (n *MediumOfInstructionType) New() *MediumOfInstructionType {
        return MediumOfInstructionTypeCreate(MediumOfInstructionType{})
}

func StatisticalAreaTypeCreate(x StatisticalAreaType) *StatisticalAreaType {
        return &x
}

func (n *StatisticalAreaType) New() *StatisticalAreaType {
        return StatisticalAreaTypeCreate(StatisticalAreaType{})
}

func ResourceUsage_ResourceReportColumnListCreate(x ResourceUsage_ResourceReportColumnList) *ResourceUsage_ResourceReportColumnList {
        return &x
}

func (n *ResourceUsage_ResourceReportColumnList) New() *ResourceUsage_ResourceReportColumnList {
        return ResourceUsage_ResourceReportColumnListCreate(ResourceUsage_ResourceReportColumnList{})
}

func StudentPeriodAttendanceCreate(x StudentPeriodAttendance) *StudentPeriodAttendance {
        return &x
}

func (n *StudentPeriodAttendance) New() *StudentPeriodAttendance {
        return StudentPeriodAttendanceCreate(StudentPeriodAttendance{})
}

func ComponentTypeCreate(x ComponentType) *ComponentType {
        return &x
}

func (n *ComponentType) New() *ComponentType {
        return ComponentTypeCreate(ComponentType{})
}

func StaffSubjectListTypeCreate(x StaffSubjectListType) *StaffSubjectListType {
        return &x
}

func (n *StaffSubjectListType) New() *StaffSubjectListType {
        return StaffSubjectListTypeCreate(StaffSubjectListType{})
}

func CensusStudentListTypeCreate(x CensusStudentListType) *CensusStudentListType {
        return &x
}

func (n *CensusStudentListType) New() *CensusStudentListType {
        return CensusStudentListTypeCreate(CensusStudentListType{})
}

func ContactInfoTypeCreate(x ContactInfoType) *ContactInfoType {
        return &x
}

func (n *ContactInfoType) New() *ContactInfoType {
        return ContactInfoTypeCreate(ContactInfoType{})
}

func DomainBandsContainerTypeCreate(x DomainBandsContainerType) *DomainBandsContainerType {
        return &x
}

func (n *DomainBandsContainerType) New() *DomainBandsContainerType {
        return DomainBandsContainerTypeCreate(DomainBandsContainerType{})
}

func AddressStreetTypeCreate(x AddressStreetType) *AddressStreetType {
        return &x
}

func (n *AddressStreetType) New() *AddressStreetType {
        return AddressStreetTypeCreate(AddressStreetType{})
}

func TimeTablePeriodListTypeCreate(x TimeTablePeriodListType) *TimeTablePeriodListType {
        return &x
}

func (n *TimeTablePeriodListType) New() *TimeTablePeriodListType {
        return TimeTablePeriodListTypeCreate(TimeTablePeriodListType{})
}

func LifeCycleType_CreatorsCreate(x LifeCycleType_Creators) *LifeCycleType_Creators {
        return &x
}

func (n *LifeCycleType_Creators) New() *LifeCycleType_Creators {
        return LifeCycleType_CreatorsCreate(LifeCycleType_Creators{})
}

func LanguageOfInstructionTypeCreate(x LanguageOfInstructionType) *LanguageOfInstructionType {
        return &x
}

func (n *LanguageOfInstructionType) New() *LanguageOfInstructionType {
        return LanguageOfInstructionTypeCreate(LanguageOfInstructionType{})
}

func FineInfoListTypeCreate(x FineInfoListType) *FineInfoListType {
        return &x
}

func (n *FineInfoListType) New() *FineInfoListType {
        return FineInfoListTypeCreate(FineInfoListType{})
}

func YearLevelTypeCreate(x YearLevelType) *YearLevelType {
        return &x
}

func (n *YearLevelType) New() *YearLevelType {
        return YearLevelTypeCreate(YearLevelType{})
}

func SubstituteItemListTypeCreate(x SubstituteItemListType) *SubstituteItemListType {
        return &x
}

func (n *SubstituteItemListType) New() *SubstituteItemListType {
        return SubstituteItemListTypeCreate(SubstituteItemListType{})
}

func PaymentReceiptLineTypeCreate(x PaymentReceiptLineType) *PaymentReceiptLineType {
        return &x
}

func (n *PaymentReceiptLineType) New() *PaymentReceiptLineType {
        return PaymentReceiptLineTypeCreate(PaymentReceiptLineType{})
}

func IdentityCreate(x Identity) *Identity {
        return &x
}

func (n *Identity) New() *Identity {
        return IdentityCreate(Identity{})
}

func SchoolInfo_OtherLEACreate(x SchoolInfo_OtherLEA) *SchoolInfo_OtherLEA {
        return &x
}

func (n *SchoolInfo_OtherLEA) New() *SchoolInfo_OtherLEA {
        return SchoolInfo_OtherLEACreate(SchoolInfo_OtherLEA{})
}

func Identity_SIF_RefIdCreate(x Identity_SIF_RefId) *Identity_SIF_RefId {
        return &x
}

func (n *Identity_SIF_RefId) New() *Identity_SIF_RefId {
        return Identity_SIF_RefIdCreate(Identity_SIF_RefId{})
}

func SchoolInfoCreate(x SchoolInfo) *SchoolInfo {
        return &x
}

func (n *SchoolInfo) New() *SchoolInfo {
        return SchoolInfoCreate(SchoolInfo{})
}

func StaffAssignmentCreate(x StaffAssignment) *StaffAssignment {
        return &x
}

func (n *StaffAssignment) New() *StaffAssignment {
        return StaffAssignmentCreate(StaffAssignment{})
}

func SchoolProgramTypeCreate(x SchoolProgramType) *SchoolProgramType {
        return &x
}

func (n *SchoolProgramType) New() *SchoolProgramType {
        return SchoolProgramTypeCreate(SchoolProgramType{})
}

func ApprovalTypeCreate(x ApprovalType) *ApprovalType {
        return &x
}

func (n *ApprovalType) New() *ApprovalType {
        return ApprovalTypeCreate(ApprovalType{})
}

func EmailTypeCreate(x EmailType) *EmailType {
        return &x
}

func (n *EmailType) New() *EmailType {
        return EmailTypeCreate(EmailType{})
}

func StudentAttendanceCollectionReportingTypeCreate(x StudentAttendanceCollectionReportingType) *StudentAttendanceCollectionReportingType {
        return &x
}

func (n *StudentAttendanceCollectionReportingType) New() *StudentAttendanceCollectionReportingType {
        return StudentAttendanceCollectionReportingTypeCreate(StudentAttendanceCollectionReportingType{})
}

func DebtorCreate(x Debtor) *Debtor {
        return &x
}

func (n *Debtor) New() *Debtor {
        return DebtorCreate(Debtor{})
}

func AgencyTypeCreate(x AgencyType) *AgencyType {
        return &x
}

func (n *AgencyType) New() *AgencyType {
        return AgencyTypeCreate(AgencyType{})
}

func StaffAssignmentMostRecentContainerTypeCreate(x StaffAssignmentMostRecentContainerType) *StaffAssignmentMostRecentContainerType {
        return &x
}

func (n *StaffAssignmentMostRecentContainerType) New() *StaffAssignmentMostRecentContainerType {
        return StaffAssignmentMostRecentContainerTypeCreate(StaffAssignmentMostRecentContainerType{})
}

func AGRoundTypeCreate(x AGRoundType) *AGRoundType {
        return &x
}

func (n *AGRoundType) New() *AGRoundType {
        return AGRoundTypeCreate(AGRoundType{})
}

func AGContextualQuestionTypeCreate(x AGContextualQuestionType) *AGContextualQuestionType {
        return &x
}

func (n *AGContextualQuestionType) New() *AGContextualQuestionType {
        return AGContextualQuestionTypeCreate(AGContextualQuestionType{})
}

func StudentListTypeCreate(x StudentListType) *StudentListType {
        return &x
}

func (n *StudentListType) New() *StudentListType {
        return StudentListTypeCreate(StudentListType{})
}

func AlertMessagesTypeCreate(x AlertMessagesType) *AlertMessagesType {
        return &x
}

func (n *AlertMessagesType) New() *AlertMessagesType {
        return AlertMessagesTypeCreate(AlertMessagesType{})
}

func AttendanceTimesTypeCreate(x AttendanceTimesType) *AttendanceTimesType {
        return &x
}

func (n *AttendanceTimesType) New() *AttendanceTimesType {
        return AttendanceTimesTypeCreate(AttendanceTimesType{})
}

func PurchasingItemTypeCreate(x PurchasingItemType) *PurchasingItemType {
        return &x
}

func (n *PurchasingItemType) New() *PurchasingItemType {
        return PurchasingItemTypeCreate(PurchasingItemType{})
}

func AGRoundListTypeCreate(x AGRoundListType) *AGRoundListType {
        return &x
}

func (n *AGRoundListType) New() *AGRoundListType {
        return AGRoundListTypeCreate(AGRoundListType{})
}

func TimeElementTypeCreate(x TimeElementType) *TimeElementType {
        return &x
}

func (n *TimeElementType) New() *TimeElementType {
        return TimeElementTypeCreate(TimeElementType{})
}

func LEAContactListTypeCreate(x LEAContactListType) *LEAContactListType {
        return &x
}

func (n *LEAContactListType) New() *LEAContactListType {
        return LEAContactListTypeCreate(LEAContactListType{})
}

func ResourceUsage_ResourceUsageContentTypeCreate(x ResourceUsage_ResourceUsageContentType) *ResourceUsage_ResourceUsageContentType {
        return &x
}

func (n *ResourceUsage_ResourceUsageContentType) New() *ResourceUsage_ResourceUsageContentType {
        return ResourceUsage_ResourceUsageContentTypeCreate(ResourceUsage_ResourceUsageContentType{})
}

func AggregateStatisticInfo_CalculationRuleCreate(x AggregateStatisticInfo_CalculationRule) *AggregateStatisticInfo_CalculationRule {
        return &x
}

func (n *AggregateStatisticInfo_CalculationRule) New() *AggregateStatisticInfo_CalculationRule {
        return AggregateStatisticInfo_CalculationRuleCreate(AggregateStatisticInfo_CalculationRule{})
}

func SchoolProgramsCreate(x SchoolPrograms) *SchoolPrograms {
        return &x
}

func (n *SchoolPrograms) New() *SchoolPrograms {
        return SchoolProgramsCreate(SchoolPrograms{})
}

func ExclusionRuleTypeCreate(x ExclusionRuleType) *ExclusionRuleType {
        return &x
}

func (n *ExclusionRuleType) New() *ExclusionRuleType {
        return ExclusionRuleTypeCreate(ExclusionRuleType{})
}

func SystemRole_SystemContextCreate(x SystemRole_SystemContext) *SystemRole_SystemContext {
        return &x
}

func (n *SystemRole_SystemContext) New() *SystemRole_SystemContext {
        return SystemRole_SystemContextCreate(SystemRole_SystemContext{})
}

func MedicalAlertMessagesTypeCreate(x MedicalAlertMessagesType) *MedicalAlertMessagesType {
        return &x
}

func (n *MedicalAlertMessagesType) New() *MedicalAlertMessagesType {
        return MedicalAlertMessagesTypeCreate(MedicalAlertMessagesType{})
}

func MapReferenceTypeCreate(x MapReferenceType) *MapReferenceType {
        return &x
}

func (n *MapReferenceType) New() *MapReferenceType {
        return MapReferenceTypeCreate(MapReferenceType{})
}

func WellbeingCharacteristicCreate(x WellbeingCharacteristic) *WellbeingCharacteristic {
        return &x
}

func (n *WellbeingCharacteristic) New() *WellbeingCharacteristic {
        return WellbeingCharacteristicCreate(WellbeingCharacteristic{})
}

func LearningResource_LocationCreate(x LearningResource_Location) *LearningResource_Location {
        return &x
}

func (n *LearningResource_Location) New() *LearningResource_Location {
        return LearningResource_LocationCreate(LearningResource_Location{})
}

func YearLevelEnrollmentTypeCreate(x YearLevelEnrollmentType) *YearLevelEnrollmentType {
        return &x
}

func (n *YearLevelEnrollmentType) New() *YearLevelEnrollmentType {
        return YearLevelEnrollmentTypeCreate(YearLevelEnrollmentType{})
}

func SystemRole_RoleCreate(x SystemRole_Role) *SystemRole_Role {
        return &x
}

func (n *SystemRole_Role) New() *SystemRole_Role {
        return SystemRole_RoleCreate(SystemRole_Role{})
}

func EntityContactInfoTypeCreate(x EntityContactInfoType) *EntityContactInfoType {
        return &x
}

func (n *EntityContactInfoType) New() *EntityContactInfoType {
        return EntityContactInfoTypeCreate(EntityContactInfoType{})
}

func AbstractContentPackageType_XMLDataCreate(x AbstractContentPackageType_XMLData) *AbstractContentPackageType_XMLData {
        return &x
}

func (n *AbstractContentPackageType_XMLData) New() *AbstractContentPackageType_XMLData {
        return AbstractContentPackageType_XMLDataCreate(AbstractContentPackageType_XMLData{})
}

func MarkValueInfoCreate(x MarkValueInfo) *MarkValueInfo {
        return &x
}

func (n *MarkValueInfo) New() *MarkValueInfo {
        return MarkValueInfoCreate(MarkValueInfo{})
}

func Invoice_InvoicedEntityCreate(x Invoice_InvoicedEntity) *Invoice_InvoicedEntity {
        return &x
}

func (n *Invoice_InvoicedEntity) New() *Invoice_InvoicedEntity {
        return Invoice_InvoicedEntityCreate(Invoice_InvoicedEntity{})
}

func PhoneNumberTypeCreate(x PhoneNumberType) *PhoneNumberType {
        return &x
}

func (n *PhoneNumberType) New() *PhoneNumberType {
        return PhoneNumberTypeCreate(PhoneNumberType{})
}

func AGContextualQuestionListTypeCreate(x AGContextualQuestionListType) *AGContextualQuestionListType {
        return &x
}

func (n *AGContextualQuestionListType) New() *AGContextualQuestionListType {
        return AGContextualQuestionListTypeCreate(AGContextualQuestionListType{})
}

func WithdrawalTimeListTypeCreate(x WithdrawalTimeListType) *WithdrawalTimeListType {
        return &x
}

func (n *WithdrawalTimeListType) New() *WithdrawalTimeListType {
        return WithdrawalTimeListTypeCreate(WithdrawalTimeListType{})
}

func ProgramFundingSourceTypeCreate(x ProgramFundingSourceType) *ProgramFundingSourceType {
        return &x
}

func (n *ProgramFundingSourceType) New() *ProgramFundingSourceType {
        return ProgramFundingSourceTypeCreate(ProgramFundingSourceType{})
}

func OtherNameTypeCreate(x OtherNameType) *OtherNameType {
        return &x
}

func (n *OtherNameType) New() *OtherNameType {
        return OtherNameTypeCreate(OtherNameType{})
}

func DemographicsTypeCreate(x DemographicsType) *DemographicsType {
        return &x
}

func (n *DemographicsType) New() *DemographicsType {
        return DemographicsTypeCreate(DemographicsType{})
}

func WellbeingEventSubCategoryListTypeCreate(x WellbeingEventSubCategoryListType) *WellbeingEventSubCategoryListType {
        return &x
}

func (n *WellbeingEventSubCategoryListType) New() *WellbeingEventSubCategoryListType {
        return WellbeingEventSubCategoryListTypeCreate(WellbeingEventSubCategoryListType{})
}

func StrategiesTypeCreate(x StrategiesType) *StrategiesType {
        return &x
}

func (n *StrategiesType) New() *StrategiesType {
        return StrategiesTypeCreate(StrategiesType{})
}

func LearningStandardsDocumentTypeCreate(x LearningStandardsDocumentType) *LearningStandardsDocumentType {
        return &x
}

func (n *LearningStandardsDocumentType) New() *LearningStandardsDocumentType {
        return LearningStandardsDocumentTypeCreate(LearningStandardsDocumentType{})
}

func LifeCycleType_CreatorCreate(x LifeCycleType_Creator) *LifeCycleType_Creator {
        return &x
}

func (n *LifeCycleType_Creator) New() *LifeCycleType_Creator {
        return LifeCycleType_CreatorCreate(LifeCycleType_Creator{})
}

func PlausibleScaledValueListTypeCreate(x PlausibleScaledValueListType) *PlausibleScaledValueListType {
        return &x
}

func (n *PlausibleScaledValueListType) New() *PlausibleScaledValueListType {
        return PlausibleScaledValueListTypeCreate(PlausibleScaledValueListType{})
}

func SubjectAreaListTypeCreate(x SubjectAreaListType) *SubjectAreaListType {
        return &x
}

func (n *SubjectAreaListType) New() *SubjectAreaListType {
        return SubjectAreaListTypeCreate(SubjectAreaListType{})
}

func LearningObjectivesTypeCreate(x LearningObjectivesType) *LearningObjectivesType {
        return &x
}

func (n *LearningObjectivesType) New() *LearningObjectivesType {
        return LearningObjectivesTypeCreate(LearningObjectivesType{})
}

func LearningResourcesTypeCreate(x LearningResourcesType) *LearningResourcesType {
        return &x
}

func (n *LearningResourcesType) New() *LearningResourcesType {
        return LearningResourcesTypeCreate(LearningResourcesType{})
}

func SuspensionContainerTypeCreate(x SuspensionContainerType) *SuspensionContainerType {
        return &x
}

func (n *SuspensionContainerType) New() *SuspensionContainerType {
        return SuspensionContainerTypeCreate(SuspensionContainerType{})
}

func StaffListTypeCreate(x StaffListType) *StaffListType {
        return &x
}

func (n *StaffListType) New() *StaffListType {
        return StaffListTypeCreate(StaffListType{})
}

func StudentMostRecentContainerTypeCreate(x StudentMostRecentContainerType) *StudentMostRecentContainerType {
        return &x
}

func (n *StudentMostRecentContainerType) New() *StudentMostRecentContainerType {
        return StudentMostRecentContainerTypeCreate(StudentMostRecentContainerType{})
}

func EvaluationsTypeCreate(x EvaluationsType) *EvaluationsType {
        return &x
}

func (n *EvaluationsType) New() *EvaluationsType {
        return EvaluationsTypeCreate(EvaluationsType{})
}

func LibraryPatronStatusCreate(x LibraryPatronStatus) *LibraryPatronStatus {
        return &x
}

func (n *LibraryPatronStatus) New() *LibraryPatronStatus {
        return LibraryPatronStatusCreate(LibraryPatronStatus{})
}

func StimulusTypeCreate(x StimulusType) *StimulusType {
        return &x
}

func (n *StimulusType) New() *StimulusType {
        return StimulusTypeCreate(StimulusType{})
}

func LocalCodeTypeCreate(x LocalCodeType) *LocalCodeType {
        return &x
}

func (n *LocalCodeType) New() *LocalCodeType {
        return LocalCodeTypeCreate(LocalCodeType{})
}

func GenericRubricTypeCreate(x GenericRubricType) *GenericRubricType {
        return &x
}

func (n *GenericRubricType) New() *GenericRubricType {
        return GenericRubricTypeCreate(GenericRubricType{})
}

func AbstractContentElementType_ReferenceCreate(x AbstractContentElementType_Reference) *AbstractContentElementType_Reference {
        return &x
}

func (n *AbstractContentElementType_Reference) New() *AbstractContentElementType_Reference {
        return AbstractContentElementType_ReferenceCreate(AbstractContentElementType_Reference{})
}

func RecognitionListTypeCreate(x RecognitionListType) *RecognitionListType {
        return &x
}

func (n *RecognitionListType) New() *RecognitionListType {
        return RecognitionListTypeCreate(RecognitionListType{})
}

func PersonPicture_PictureSourceCreate(x PersonPicture_PictureSource) *PersonPicture_PictureSource {
        return &x
}

func (n *PersonPicture_PictureSource) New() *PersonPicture_PictureSource {
        return PersonPicture_PictureSourceCreate(PersonPicture_PictureSource{})
}

func SystemRole_RoleScopeListCreate(x SystemRole_RoleScopeList) *SystemRole_RoleScopeList {
        return &x
}

func (n *SystemRole_RoleScopeList) New() *SystemRole_RoleScopeList {
        return SystemRole_RoleScopeListCreate(SystemRole_RoleScopeList{})
}

func PNPCodeListTypeCreate(x PNPCodeListType) *PNPCodeListType {
        return &x
}

func (n *PNPCodeListType) New() *PNPCodeListType {
        return PNPCodeListTypeCreate(PNPCodeListType{})
}

func PlanRequiredListTypeCreate(x PlanRequiredListType) *PlanRequiredListType {
        return &x
}

func (n *PlanRequiredListType) New() *PlanRequiredListType {
        return PlanRequiredListTypeCreate(PlanRequiredListType{})
}

func PeriodAttendanceTypeCreate(x PeriodAttendanceType) *PeriodAttendanceType {
        return &x
}

func (n *PeriodAttendanceType) New() *PeriodAttendanceType {
        return PeriodAttendanceTypeCreate(PeriodAttendanceType{})
}

func ScoreListTypeCreate(x ScoreListType) *ScoreListType {
        return &x
}

func (n *ScoreListType) New() *ScoreListType {
        return ScoreListTypeCreate(ScoreListType{})
}

func StudentActivityTypeCreate(x StudentActivityType) *StudentActivityType {
        return &x
}

func (n *StudentActivityType) New() *StudentActivityType {
        return StudentActivityTypeCreate(StudentActivityType{})
}

func NAPTestItemContentTypeCreate(x NAPTestItemContentType) *NAPTestItemContentType {
        return &x
}

func (n *NAPTestItemContentType) New() *NAPTestItemContentType {
        return NAPTestItemContentTypeCreate(NAPTestItemContentType{})
}

func SystemRole_SystemContextListCreate(x SystemRole_SystemContextList) *SystemRole_SystemContextList {
        return &x
}

func (n *SystemRole_SystemContextList) New() *SystemRole_SystemContextList {
        return SystemRole_SystemContextListCreate(SystemRole_SystemContextList{})
}

func TermInfoCreate(x TermInfo) *TermInfo {
        return &x
}

func (n *TermInfo) New() *TermInfo {
        return TermInfoCreate(TermInfo{})
}

func WellbeingEventCategoryTypeCreate(x WellbeingEventCategoryType) *WellbeingEventCategoryType {
        return &x
}

func (n *WellbeingEventCategoryType) New() *WellbeingEventCategoryType {
        return WellbeingEventCategoryTypeCreate(WellbeingEventCategoryType{})
}

func AddressCollectionCreate(x AddressCollection) *AddressCollection {
        return &x
}

func (n *AddressCollection) New() *AddressCollection {
        return AddressCollectionCreate(AddressCollection{})
}

func SystemRole_RoleListCreate(x SystemRole_RoleList) *SystemRole_RoleList {
        return &x
}

func (n *SystemRole_RoleList) New() *SystemRole_RoleList {
        return SystemRole_RoleListCreate(SystemRole_RoleList{})
}

func WithdrawalTypeCreate(x WithdrawalType) *WithdrawalType {
        return &x
}

func (n *WithdrawalType) New() *WithdrawalType {
        return WithdrawalTypeCreate(WithdrawalType{})
}

func ScoreDescriptionListTypeCreate(x ScoreDescriptionListType) *ScoreDescriptionListType {
        return &x
}

func (n *ScoreDescriptionListType) New() *ScoreDescriptionListType {
        return ScoreDescriptionListTypeCreate(ScoreDescriptionListType{})
}

func ACStrandSubjectAreaTypeCreate(x ACStrandSubjectAreaType) *ACStrandSubjectAreaType {
        return &x
}

func (n *ACStrandSubjectAreaType) New() *ACStrandSubjectAreaType {
        return ACStrandSubjectAreaTypeCreate(ACStrandSubjectAreaType{})
}

func LibraryTransactionListTypeCreate(x LibraryTransactionListType) *LibraryTransactionListType {
        return &x
}

func (n *LibraryTransactionListType) New() *LibraryTransactionListType {
        return LibraryTransactionListTypeCreate(LibraryTransactionListType{})
}

func CensusStudentTypeCreate(x CensusStudentType) *CensusStudentType {
        return &x
}

func (n *CensusStudentType) New() *CensusStudentType {
        return CensusStudentTypeCreate(CensusStudentType{})
}

func AddressTypeCreate(x AddressType) *AddressType {
        return &x
}

func (n *AddressType) New() *AddressType {
        return AddressTypeCreate(AddressType{})
}

func NAPEventStudentLinkCreate(x NAPEventStudentLink) *NAPEventStudentLink {
        return &x
}

func (n *NAPEventStudentLink) New() *NAPEventStudentLink {
        return NAPEventStudentLinkCreate(NAPEventStudentLink{})
}

func LibraryMessageListTypeCreate(x LibraryMessageListType) *LibraryMessageListType {
        return &x
}

func (n *LibraryMessageListType) New() *LibraryMessageListType {
        return LibraryMessageListTypeCreate(LibraryMessageListType{})
}

func LifeCycleType_ModifiedCreate(x LifeCycleType_Modified) *LifeCycleType_Modified {
        return &x
}

func (n *LifeCycleType_Modified) New() *LifeCycleType_Modified {
        return LifeCycleType_ModifiedCreate(LifeCycleType_Modified{})
}

func SessionInfoCreate(x SessionInfo) *SessionInfo {
        return &x
}

func (n *SessionInfo) New() *SessionInfo {
        return SessionInfoCreate(SessionInfo{})
}

func NAPTestScoreSummaryCreate(x NAPTestScoreSummary) *NAPTestScoreSummary {
        return &x
}

func (n *NAPTestScoreSummary) New() *NAPTestScoreSummary {
        return NAPTestScoreSummaryCreate(NAPTestScoreSummary{})
}

func TimeTableCellCreate(x TimeTableCell) *TimeTableCell {
        return &x
}

func (n *TimeTableCell) New() *TimeTableCell {
        return TimeTableCellCreate(TimeTableCell{})
}

func ScheduledActivityCreate(x ScheduledActivity) *ScheduledActivity {
        return &x
}

func (n *ScheduledActivity) New() *ScheduledActivity {
        return ScheduledActivityCreate(ScheduledActivity{})
}

func SoftwareRequirementTypeCreate(x SoftwareRequirementType) *SoftwareRequirementType {
        return &x
}

func (n *SoftwareRequirementType) New() *SoftwareRequirementType {
        return SoftwareRequirementTypeCreate(SoftwareRequirementType{})
}

func ScoreDescriptionTypeCreate(x ScoreDescriptionType) *ScoreDescriptionType {
        return &x
}

func (n *ScoreDescriptionType) New() *ScoreDescriptionType {
        return ScoreDescriptionTypeCreate(ScoreDescriptionType{})
}

func NAPTestItemListTypeCreate(x NAPTestItemListType) *NAPTestItemListType {
        return &x
}

func (n *NAPTestItemListType) New() *NAPTestItemListType {
        return NAPTestItemListTypeCreate(NAPTestItemListType{})
}

func AddressListTypeCreate(x AddressListType) *AddressListType {
        return &x
}

func (n *AddressListType) New() *AddressListType {
        return AddressListTypeCreate(AddressListType{})
}

func NAPTestletResponseItemTypeCreate(x NAPTestletResponseItemType) *NAPTestletResponseItemType {
        return &x
}

func (n *NAPTestletResponseItemType) New() *NAPTestletResponseItemType {
        return NAPTestletResponseItemTypeCreate(NAPTestletResponseItemType{})
}

func MedicationListTypeCreate(x MedicationListType) *MedicationListType {
        return &x
}

func (n *MedicationListType) New() *MedicationListType {
        return MedicationListTypeCreate(MedicationListType{})
}

func CalendarDateCreate(x CalendarDate) *CalendarDate {
        return &x
}

func (n *CalendarDate) New() *CalendarDate {
        return CalendarDateCreate(CalendarDate{})
}

func NAPTestletCreate(x NAPTestlet) *NAPTestlet {
        return &x
}

func (n *NAPTestlet) New() *NAPTestlet {
        return NAPTestletCreate(NAPTestlet{})
}

func FineInfoTypeCreate(x FineInfoType) *FineInfoType {
        return &x
}

func (n *FineInfoType) New() *FineInfoType {
        return FineInfoTypeCreate(FineInfoType{})
}

func FinancialAccountRefIdListTypeCreate(x FinancialAccountRefIdListType) *FinancialAccountRefIdListType {
        return &x
}

func (n *FinancialAccountRefIdListType) New() *FinancialAccountRefIdListType {
        return FinancialAccountRefIdListTypeCreate(FinancialAccountRefIdListType{})
}

func CensusStaffListTypeCreate(x CensusStaffListType) *CensusStaffListType {
        return &x
}

func (n *CensusStaffListType) New() *CensusStaffListType {
        return CensusStaffListTypeCreate(CensusStaffListType{})
}

func FQItemTypeCreate(x FQItemType) *FQItemType {
        return &x
}

func (n *FQItemType) New() *FQItemType {
        return FQItemTypeCreate(FQItemType{})
}

func IdentityAssertionsTypeCreate(x IdentityAssertionsType) *IdentityAssertionsType {
        return &x
}

func (n *IdentityAssertionsType) New() *IdentityAssertionsType {
        return IdentityAssertionsTypeCreate(IdentityAssertionsType{})
}

func SoftwareRequirementListTypeCreate(x SoftwareRequirementListType) *SoftwareRequirementListType {
        return &x
}

func (n *SoftwareRequirementListType) New() *SoftwareRequirementListType {
        return SoftwareRequirementListTypeCreate(SoftwareRequirementListType{})
}

func SchoolProgramListTypeCreate(x SchoolProgramListType) *SchoolProgramListType {
        return &x
}

func (n *SchoolProgramListType) New() *SchoolProgramListType {
        return SchoolProgramListTypeCreate(SchoolProgramListType{})
}

func LanguageListTypeCreate(x LanguageListType) *LanguageListType {
        return &x
}

func (n *LanguageListType) New() *LanguageListType {
        return LanguageListTypeCreate(LanguageListType{})
}

func TestDisruptionTypeCreate(x TestDisruptionType) *TestDisruptionType {
        return &x
}

func (n *TestDisruptionType) New() *TestDisruptionType {
        return TestDisruptionTypeCreate(TestDisruptionType{})
}

func StudentSchoolEnrollmentCreate(x StudentSchoolEnrollment) *StudentSchoolEnrollment {
        return &x
}

func (n *StudentSchoolEnrollment) New() *StudentSchoolEnrollment {
        return StudentSchoolEnrollmentCreate(StudentSchoolEnrollment{})
}

func AssociatedObjectsType_AssociatedObjectCreate(x AssociatedObjectsType_AssociatedObject) *AssociatedObjectsType_AssociatedObject {
        return &x
}

func (n *AssociatedObjectsType_AssociatedObject) New() *AssociatedObjectsType_AssociatedObject {
        return AssociatedObjectsType_AssociatedObjectCreate(AssociatedObjectsType_AssociatedObject{})
}

func StudentPersonalCreate(x StudentPersonal) *StudentPersonal {
        return &x
}

func (n *StudentPersonal) New() *StudentPersonal {
        return StudentPersonalCreate(StudentPersonal{})
}

func SystemRoleCreate(x SystemRole) *SystemRole {
        return &x
}

func (n *SystemRole) New() *SystemRole {
        return SystemRoleCreate(SystemRole{})
}

func AGReportingObjectResponseTypeCreate(x AGReportingObjectResponseType) *AGReportingObjectResponseType {
        return &x
}

func (n *AGReportingObjectResponseType) New() *AGReportingObjectResponseType {
        return AGReportingObjectResponseTypeCreate(AGReportingObjectResponseType{})
}

func OtherIdTypeCreate(x OtherIdType) *OtherIdType {
        return &x
}

func (n *OtherIdType) New() *OtherIdType {
        return OtherIdTypeCreate(OtherIdType{})
}

func WellbeingDocumentListTypeCreate(x WellbeingDocumentListType) *WellbeingDocumentListType {
        return &x
}

func (n *WellbeingDocumentListType) New() *WellbeingDocumentListType {
        return WellbeingDocumentListTypeCreate(WellbeingDocumentListType{})
}

func PersonInvolvementTypeCreate(x PersonInvolvementType) *PersonInvolvementType {
        return &x
}

func (n *PersonInvolvementType) New() *PersonInvolvementType {
        return PersonInvolvementTypeCreate(PersonInvolvementType{})
}

func NAPTestCreate(x NAPTest) *NAPTest {
        return &x
}

func (n *NAPTest) New() *NAPTest {
        return NAPTestCreate(NAPTest{})
}

func TechnicalRequirementsTypeCreate(x TechnicalRequirementsType) *TechnicalRequirementsType {
        return &x
}

func (n *TechnicalRequirementsType) New() *TechnicalRequirementsType {
        return TechnicalRequirementsTypeCreate(TechnicalRequirementsType{})
}

func FQReportingTypeCreate(x FQReportingType) *FQReportingType {
        return &x
}

func (n *FQReportingType) New() *FQReportingType {
        return FQReportingTypeCreate(FQReportingType{})
}

func AGRuleTypeCreate(x AGRuleType) *AGRuleType {
        return &x
}

func (n *AGRuleType) New() *AGRuleType {
        return AGRuleTypeCreate(AGRuleType{})
}

func StudentAttendanceSummaryCreate(x StudentAttendanceSummary) *StudentAttendanceSummary {
        return &x
}

func (n *StudentAttendanceSummary) New() *StudentAttendanceSummary {
        return StudentAttendanceSummaryCreate(StudentAttendanceSummary{})
}

func JournalCreate(x Journal) *Journal {
        return &x
}

func (n *Journal) New() *Journal {
        return JournalCreate(Journal{})
}

func SIF_ExtendedElementsTypeCreate(x SIF_ExtendedElementsType) *SIF_ExtendedElementsType {
        return &x
}

func (n *SIF_ExtendedElementsType) New() *SIF_ExtendedElementsType {
        return SIF_ExtendedElementsTypeCreate(SIF_ExtendedElementsType{})
}

func OtherIdListTypeCreate(x OtherIdListType) *OtherIdListType {
        return &x
}

func (n *OtherIdListType) New() *OtherIdListType {
        return OtherIdListTypeCreate(OtherIdListType{})
}

func LearningStandardDocumentCreate(x LearningStandardDocument) *LearningStandardDocument {
        return &x
}

func (n *LearningStandardDocument) New() *LearningStandardDocument {
        return LearningStandardDocumentCreate(LearningStandardDocument{})
}

func LanguageBaseTypeCreate(x LanguageBaseType) *LanguageBaseType {
        return &x
}

func (n *LanguageBaseType) New() *LanguageBaseType {
        return LanguageBaseTypeCreate(LanguageBaseType{})
}

func AbstractContentPackageType_TextDataCreate(x AbstractContentPackageType_TextData) *AbstractContentPackageType_TextData {
        return &x
}

func (n *AbstractContentPackageType_TextData) New() *AbstractContentPackageType_TextData {
        return AbstractContentPackageType_TextDataCreate(AbstractContentPackageType_TextData{})
}

func ProgramAvailabilityTypeCreate(x ProgramAvailabilityType) *ProgramAvailabilityType {
        return &x
}

func (n *ProgramAvailabilityType) New() *ProgramAvailabilityType {
        return ProgramAvailabilityTypeCreate(ProgramAvailabilityType{})
}

func LocalCodeListTypeCreate(x LocalCodeListType) *LocalCodeListType {
        return &x
}

func (n *LocalCodeListType) New() *LocalCodeListType {
        return LocalCodeListTypeCreate(LocalCodeListType{})
}

func ExclusionRulesTypeCreate(x ExclusionRulesType) *ExclusionRulesType {
        return &x
}

func (n *ExclusionRulesType) New() *ExclusionRulesType {
        return ExclusionRulesTypeCreate(ExclusionRulesType{})
}

func TeachingGroupScheduleTypeCreate(x TeachingGroupScheduleType) *TeachingGroupScheduleType {
        return &x
}

func (n *TeachingGroupScheduleType) New() *TeachingGroupScheduleType {
        return TeachingGroupScheduleTypeCreate(TeachingGroupScheduleType{})
}

func WellbeingPlanTypeCreate(x WellbeingPlanType) *WellbeingPlanType {
        return &x
}

func (n *WellbeingPlanType) New() *WellbeingPlanType {
        return WellbeingPlanTypeCreate(WellbeingPlanType{})
}

func EnglishProficiencyTypeCreate(x EnglishProficiencyType) *EnglishProficiencyType {
        return &x
}

func (n *EnglishProficiencyType) New() *EnglishProficiencyType {
        return EnglishProficiencyTypeCreate(EnglishProficiencyType{})
}

func RoomListTypeCreate(x RoomListType) *RoomListType {
        return &x
}

func (n *RoomListType) New() *RoomListType {
        return RoomListTypeCreate(RoomListType{})
}

func SourceObjectsTypeCreate(x SourceObjectsType) *SourceObjectsType {
        return &x
}

func (n *SourceObjectsType) New() *SourceObjectsType {
        return SourceObjectsTypeCreate(SourceObjectsType{})
}

func FollowUpActionListTypeCreate(x FollowUpActionListType) *FollowUpActionListType {
        return &x
}

func (n *FollowUpActionListType) New() *FollowUpActionListType {
        return FollowUpActionListTypeCreate(FollowUpActionListType{})
}

func StaffSubjectTypeCreate(x StaffSubjectType) *StaffSubjectType {
        return &x
}

func (n *StaffSubjectType) New() *StaffSubjectType {
        return StaffSubjectTypeCreate(StaffSubjectType{})
}

func LearningStandardsTypeCreate(x LearningStandardsType) *LearningStandardsType {
        return &x
}

func (n *LearningStandardsType) New() *LearningStandardsType {
        return LearningStandardsTypeCreate(LearningStandardsType{})
}

func StatsCohortTypeCreate(x StatsCohortType) *StatsCohortType {
        return &x
}

func (n *StatsCohortType) New() *StatsCohortType {
        return StatsCohortTypeCreate(StatsCohortType{})
}

func ElectronicIdTypeCreate(x ElectronicIdType) *ElectronicIdType {
        return &x
}

func (n *ElectronicIdType) New() *ElectronicIdType {
        return ElectronicIdTypeCreate(ElectronicIdType{})
}

func NAPCodeFrameTestletListTypeCreate(x NAPCodeFrameTestletListType) *NAPCodeFrameTestletListType {
        return &x
}

func (n *NAPCodeFrameTestletListType) New() *NAPCodeFrameTestletListType {
        return NAPCodeFrameTestletListTypeCreate(NAPCodeFrameTestletListType{})
}

func AuthorsTypeCreate(x AuthorsType) *AuthorsType {
        return &x
}

func (n *AuthorsType) New() *AuthorsType {
        return AuthorsTypeCreate(AuthorsType{})
}

func NAPTestletContentTypeCreate(x NAPTestletContentType) *NAPTestletContentType {
        return &x
}

func (n *NAPTestletContentType) New() *NAPTestletContentType {
        return NAPTestletContentTypeCreate(NAPTestletContentType{})
}

func GradingScoreListTypeCreate(x GradingScoreListType) *GradingScoreListType {
        return &x
}

func (n *GradingScoreListType) New() *GradingScoreListType {
        return GradingScoreListTypeCreate(GradingScoreListType{})
}

func TeachingGroupStudentTypeCreate(x TeachingGroupStudentType) *TeachingGroupStudentType {
        return &x
}

func (n *TeachingGroupStudentType) New() *TeachingGroupStudentType {
        return TeachingGroupStudentTypeCreate(TeachingGroupStudentType{})
}

func NAPTestItemCreate(x NAPTestItem) *NAPTestItem {
        return &x
}

func (n *NAPTestItem) New() *NAPTestItem {
        return NAPTestItemCreate(NAPTestItem{})
}

func TeacherListTypeCreate(x TeacherListType) *TeacherListType {
        return &x
}

func (n *TeacherListType) New() *TeacherListType {
        return TeacherListTypeCreate(TeacherListType{})
}

func ContactTypeCreate(x ContactType) *ContactType {
        return &x
}

func (n *ContactType) New() *ContactType {
        return ContactTypeCreate(ContactType{})
}

func AbstractContentPackageType_BinaryDataCreate(x AbstractContentPackageType_BinaryData) *AbstractContentPackageType_BinaryData {
        return &x
}

func (n *AbstractContentPackageType_BinaryData) New() *AbstractContentPackageType_BinaryData {
        return AbstractContentPackageType_BinaryDataCreate(AbstractContentPackageType_BinaryData{})
}

func StudentActivityParticipationCreate(x StudentActivityParticipation) *StudentActivityParticipation {
        return &x
}

func (n *StudentActivityParticipation) New() *StudentActivityParticipation {
        return StudentActivityParticipationCreate(StudentActivityParticipation{})
}

func SchoolContactTypeCreate(x SchoolContactType) *SchoolContactType {
        return &x
}

func (n *SchoolContactType) New() *SchoolContactType {
        return SchoolContactTypeCreate(SchoolContactType{})
}

func RelatedLearningStandardItemRefIdTypeCreate(x RelatedLearningStandardItemRefIdType) *RelatedLearningStandardItemRefIdType {
        return &x
}

func (n *RelatedLearningStandardItemRefIdType) New() *RelatedLearningStandardItemRefIdType {
        return RelatedLearningStandardItemRefIdTypeCreate(RelatedLearningStandardItemRefIdType{})
}

func ScheduledTeacherListTypeCreate(x ScheduledTeacherListType) *ScheduledTeacherListType {
        return &x
}

func (n *ScheduledTeacherListType) New() *ScheduledTeacherListType {
        return ScheduledTeacherListTypeCreate(ScheduledTeacherListType{})
}

func HouseholdListTypeCreate(x HouseholdListType) *HouseholdListType {
        return &x
}

func (n *HouseholdListType) New() *HouseholdListType {
        return HouseholdListTypeCreate(HouseholdListType{})
}

func PlanRequiredContainerTypeCreate(x PlanRequiredContainerType) *PlanRequiredContainerType {
        return &x
}

func (n *PlanRequiredContainerType) New() *PlanRequiredContainerType {
        return PlanRequiredContainerTypeCreate(PlanRequiredContainerType{})
}

func GridLocationTypeCreate(x GridLocationType) *GridLocationType {
        return &x
}

func (n *GridLocationType) New() *GridLocationType {
        return GridLocationTypeCreate(GridLocationType{})
}

func MedicationTypeCreate(x MedicationType) *MedicationType {
        return &x
}

func (n *MedicationType) New() *MedicationType {
        return MedicationTypeCreate(MedicationType{})
}

func VendorInfoCreate(x VendorInfo) *VendorInfo {
        return &x
}

func (n *VendorInfo) New() *VendorInfo {
        return VendorInfoCreate(VendorInfo{})
}

func AttendanceInfoTypeCreate(x AttendanceInfoType) *AttendanceInfoType {
        return &x
}

func (n *AttendanceInfoType) New() *AttendanceInfoType {
        return AttendanceInfoTypeCreate(AttendanceInfoType{})
}

func SystemRole_SIF_RefIdCreate(x SystemRole_SIF_RefId) *SystemRole_SIF_RefId {
        return &x
}

func (n *SystemRole_SIF_RefId) New() *SystemRole_SIF_RefId {
        return SystemRole_SIF_RefIdCreate(SystemRole_SIF_RefId{})
}

func AlternateIdentificationCodeListTypeCreate(x AlternateIdentificationCodeListType) *AlternateIdentificationCodeListType {
        return &x
}

func (n *AlternateIdentificationCodeListType) New() *AlternateIdentificationCodeListType {
        return AlternateIdentificationCodeListTypeCreate(AlternateIdentificationCodeListType{})
}

func StandardIdentifierTypeCreate(x StandardIdentifierType) *StandardIdentifierType {
        return &x
}

func (n *StandardIdentifierType) New() *StandardIdentifierType {
        return StandardIdentifierTypeCreate(StandardIdentifierType{})
}

func ProgramFundingSourcesTypeCreate(x ProgramFundingSourcesType) *ProgramFundingSourcesType {
        return &x
}

func (n *ProgramFundingSourcesType) New() *ProgramFundingSourcesType {
        return ProgramFundingSourcesTypeCreate(ProgramFundingSourcesType{})
}

func LearningStandardItemCreate(x LearningStandardItem) *LearningStandardItem {
        return &x
}

func (n *LearningStandardItem) New() *LearningStandardItem {
        return LearningStandardItemCreate(LearningStandardItem{})
}

func StudentContactRelationshipCreate(x StudentContactRelationship) *StudentContactRelationship {
        return &x
}

func (n *StudentContactRelationship) New() *StudentContactRelationship {
        return StudentContactRelationshipCreate(StudentContactRelationship{})
}

func ResourceUsage_SIF_RefIdCreate(x ResourceUsage_SIF_RefId) *ResourceUsage_SIF_RefId {
        return &x
}

func (n *ResourceUsage_SIF_RefId) New() *ResourceUsage_SIF_RefId {
        return ResourceUsage_SIF_RefIdCreate(ResourceUsage_SIF_RefId{})
}

func StatsCohortYearLevelListTypeCreate(x StatsCohortYearLevelListType) *StatsCohortYearLevelListType {
        return &x
}

func (n *StatsCohortYearLevelListType) New() *StatsCohortYearLevelListType {
        return StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
}

func StudentGroupTypeCreate(x StudentGroupType) *StudentGroupType {
        return &x
}

func (n *StudentGroupType) New() *StudentGroupType {
        return StudentGroupTypeCreate(StudentGroupType{})
}

func IdentityAssertionsType_IdentityAssertionCreate(x IdentityAssertionsType_IdentityAssertion) *IdentityAssertionsType_IdentityAssertion {
        return &x
}

func (n *IdentityAssertionsType_IdentityAssertion) New() *IdentityAssertionsType_IdentityAssertion {
        return IdentityAssertionsType_IdentityAssertionCreate(IdentityAssertionsType_IdentityAssertion{})
}

func ReligiousEventTypeCreate(x ReligiousEventType) *ReligiousEventType {
        return &x
}

func (n *ReligiousEventType) New() *ReligiousEventType {
        return ReligiousEventTypeCreate(ReligiousEventType{})
}

func AddressCollectionStudentListTypeCreate(x AddressCollectionStudentListType) *AddressCollectionStudentListType {
        return &x
}

func (n *AddressCollectionStudentListType) New() *AddressCollectionStudentListType {
        return AddressCollectionStudentListTypeCreate(AddressCollectionStudentListType{})
}

func FinancialAccountCreate(x FinancialAccount) *FinancialAccount {
        return &x
}

func (n *FinancialAccount) New() *FinancialAccount {
        return FinancialAccountCreate(FinancialAccount{})
}

func CheckoutInfoTypeCreate(x CheckoutInfoType) *CheckoutInfoType {
        return &x
}

func (n *CheckoutInfoType) New() *CheckoutInfoType {
        return CheckoutInfoTypeCreate(CheckoutInfoType{})
}

func CopyRightContainerTypeCreate(x CopyRightContainerType) *CopyRightContainerType {
        return &x
}

func (n *CopyRightContainerType) New() *CopyRightContainerType {
        return CopyRightContainerTypeCreate(CopyRightContainerType{})
}

func StudentGradeMarkersListTypeCreate(x StudentGradeMarkersListType) *StudentGradeMarkersListType {
        return &x
}

func (n *StudentGradeMarkersListType) New() *StudentGradeMarkersListType {
        return StudentGradeMarkersListTypeCreate(StudentGradeMarkersListType{})
}

func StudentExitContainerTypeCreate(x StudentExitContainerType) *StudentExitContainerType {
        return &x
}

func (n *StudentExitContainerType) New() *StudentExitContainerType {
        return StudentExitContainerTypeCreate(StudentExitContainerType{})
}

func StatsCohortYearLevelTypeCreate(x StatsCohortYearLevelType) *StatsCohortYearLevelType {
        return &x
}

func (n *StatsCohortYearLevelType) New() *StatsCohortYearLevelType {
        return StatsCohortYearLevelTypeCreate(StatsCohortYearLevelType{})
}

func LearningResourceCreate(x LearningResource) *LearningResource {
        return &x
}

func (n *LearningResource) New() *LearningResource {
        return LearningResourceCreate(LearningResource{})
}

func AssignmentScoreTypeCreate(x AssignmentScoreType) *AssignmentScoreType {
        return &x
}

func (n *AssignmentScoreType) New() *AssignmentScoreType {
        return AssignmentScoreTypeCreate(AssignmentScoreType{})
}

func ContactsTypeCreate(x ContactsType) *ContactsType {
        return &x
}

func (n *ContactsType) New() *ContactsType {
        return ContactsTypeCreate(ContactsType{})
}

func NAPTestletItemResponseListTypeCreate(x NAPTestletItemResponseListType) *NAPTestletItemResponseListType {
        return &x
}

func (n *NAPTestletItemResponseListType) New() *NAPTestletItemResponseListType {
        return NAPTestletItemResponseListTypeCreate(NAPTestletItemResponseListType{})
}

func StudentAttendanceCollectionCreate(x StudentAttendanceCollection) *StudentAttendanceCollection {
        return &x
}

func (n *StudentAttendanceCollection) New() *StudentAttendanceCollection {
        return StudentAttendanceCollectionCreate(StudentAttendanceCollection{})
}

func AbstractContentPackageTypeCreate(x AbstractContentPackageType) *AbstractContentPackageType {
        return &x
}

func (n *AbstractContentPackageType) New() *AbstractContentPackageType {
        return AbstractContentPackageTypeCreate(AbstractContentPackageType{})
}

func WellbeingAppealCreate(x WellbeingAppeal) *WellbeingAppeal {
        return &x
}

func (n *WellbeingAppeal) New() *WellbeingAppeal {
        return WellbeingAppealCreate(WellbeingAppeal{})
}

func VisaSubClassListTypeCreate(x VisaSubClassListType) *VisaSubClassListType {
        return &x
}

func (n *VisaSubClassListType) New() *VisaSubClassListType {
        return VisaSubClassListTypeCreate(VisaSubClassListType{})
}

func PrincipalInfoTypeCreate(x PrincipalInfoType) *PrincipalInfoType {
        return &x
}

func (n *PrincipalInfoType) New() *PrincipalInfoType {
        return PrincipalInfoTypeCreate(PrincipalInfoType{})
}

func ApprovalsTypeCreate(x ApprovalsType) *ApprovalsType {
        return &x
}

func (n *ApprovalsType) New() *ApprovalsType {
        return ApprovalsTypeCreate(ApprovalsType{})
}

func LResourcesTypeCreate(x LResourcesType) *LResourcesType {
        return &x
}

func (n *LResourcesType) New() *LResourcesType {
        return LResourcesTypeCreate(LResourcesType{})
}

func NAPSubscoreTypeCreate(x NAPSubscoreType) *NAPSubscoreType {
        return &x
}

func (n *NAPSubscoreType) New() *NAPSubscoreType {
        return NAPSubscoreTypeCreate(NAPSubscoreType{})
}

func StudentSchoolEnrollment_CalendarCreate(x StudentSchoolEnrollment_Calendar) *StudentSchoolEnrollment_Calendar {
        return &x
}

func (n *StudentSchoolEnrollment_Calendar) New() *StudentSchoolEnrollment_Calendar {
        return StudentSchoolEnrollment_CalendarCreate(StudentSchoolEnrollment_Calendar{})
}

func EquipmentInfo_SIF_RefIdCreate(x EquipmentInfo_SIF_RefId) *EquipmentInfo_SIF_RefId {
        return &x
}

func (n *EquipmentInfo_SIF_RefId) New() *EquipmentInfo_SIF_RefId {
        return EquipmentInfo_SIF_RefIdCreate(EquipmentInfo_SIF_RefId{})
}

func TotalEnrollmentsTypeCreate(x TotalEnrollmentsType) *TotalEnrollmentsType {
        return &x
}

func (n *TotalEnrollmentsType) New() *TotalEnrollmentsType {
        return TotalEnrollmentsTypeCreate(TotalEnrollmentsType{})
}

func StatementsTypeCreate(x StatementsType) *StatementsType {
        return &x
}

func (n *StatementsType) New() *StatementsType {
        return StatementsTypeCreate(StatementsType{})
}

func MarkerTypeCreate(x MarkerType) *MarkerType {
        return &x
}

func (n *MarkerType) New() *MarkerType {
        return MarkerTypeCreate(MarkerType{})
}

func CalendarDateInfoTypeCreate(x CalendarDateInfoType) *CalendarDateInfoType {
        return &x
}

func (n *CalendarDateInfoType) New() *CalendarDateInfoType {
        return CalendarDateInfoTypeCreate(CalendarDateInfoType{})
}

func Journal_OriginatingTransactionRefIdCreate(x Journal_OriginatingTransactionRefId) *Journal_OriginatingTransactionRefId {
        return &x
}

func (n *Journal_OriginatingTransactionRefId) New() *Journal_OriginatingTransactionRefId {
        return Journal_OriginatingTransactionRefIdCreate(Journal_OriginatingTransactionRefId{})
}

func PromotionInfoTypeCreate(x PromotionInfoType) *PromotionInfoType {
        return &x
}

func (n *PromotionInfoType) New() *PromotionInfoType {
        return PromotionInfoTypeCreate(PromotionInfoType{})
}

func TeachingGroupTeacherTypeCreate(x TeachingGroupTeacherType) *TeachingGroupTeacherType {
        return &x
}

func (n *TeachingGroupTeacherType) New() *TeachingGroupTeacherType {
        return TeachingGroupTeacherTypeCreate(TeachingGroupTeacherType{})
}

func Activity_EvaluationCreate(x Activity_Evaluation) *Activity_Evaluation {
        return &x
}

func (n *Activity_Evaluation) New() *Activity_Evaluation {
        return Activity_EvaluationCreate(Activity_Evaluation{})
}

func StaffMostRecentContainerTypeCreate(x StaffMostRecentContainerType) *StaffMostRecentContainerType {
        return &x
}

func (n *StaffMostRecentContainerType) New() *StaffMostRecentContainerType {
        return StaffMostRecentContainerTypeCreate(StaffMostRecentContainerType{})
}

func SIF_MetadataType_TimeElementsCreate(x SIF_MetadataType_TimeElements) *SIF_MetadataType_TimeElements {
        return &x
}

func (n *SIF_MetadataType_TimeElements) New() *SIF_MetadataType_TimeElements {
        return SIF_MetadataType_TimeElementsCreate(SIF_MetadataType_TimeElements{})
}

func ResourceBookingCreate(x ResourceBooking) *ResourceBooking {
        return &x
}

func (n *ResourceBooking) New() *ResourceBooking {
        return ResourceBookingCreate(ResourceBooking{})
}

func NAPStudentResponseSetCreate(x NAPStudentResponseSet) *NAPStudentResponseSet {
        return &x
}

func (n *NAPStudentResponseSet) New() *NAPStudentResponseSet {
        return NAPStudentResponseSetCreate(NAPStudentResponseSet{})
}

func BaseNameTypeCreate(x BaseNameType) *BaseNameType {
        return &x
}

func (n *BaseNameType) New() *BaseNameType {
        return BaseNameTypeCreate(BaseNameType{})
}

func StandardHierarchyLevelTypeCreate(x StandardHierarchyLevelType) *StandardHierarchyLevelType {
        return &x
}

func (n *StandardHierarchyLevelType) New() *StandardHierarchyLevelType {
        return StandardHierarchyLevelTypeCreate(StandardHierarchyLevelType{})
}

func StatisticalAreasTypeCreate(x StatisticalAreasType) *StatisticalAreasType {
        return &x
}

func (n *StatisticalAreasType) New() *StatisticalAreasType {
        return StatisticalAreasTypeCreate(StatisticalAreasType{})
}

func PeriodAttendancesTypeCreate(x PeriodAttendancesType) *PeriodAttendancesType {
        return &x
}

func (n *PeriodAttendancesType) New() *PeriodAttendancesType {
        return PeriodAttendancesTypeCreate(PeriodAttendancesType{})
}

func StatementCodesTypeCreate(x StatementCodesType) *StatementCodesType {
        return &x
}

func (n *StatementCodesType) New() *StatementCodesType {
        return StatementCodesTypeCreate(StatementCodesType{})
}

func TeacherCoverTypeCreate(x TeacherCoverType) *TeacherCoverType {
        return &x
}

func (n *TeacherCoverType) New() *TeacherCoverType {
        return TeacherCoverTypeCreate(TeacherCoverType{})
}

func TimeTableScheduleCellTypeCreate(x TimeTableScheduleCellType) *TimeTableScheduleCellType {
        return &x
}

func (n *TimeTableScheduleCellType) New() *TimeTableScheduleCellType {
        return TimeTableScheduleCellTypeCreate(TimeTableScheduleCellType{})
}

func SchoolFocusListTypeCreate(x SchoolFocusListType) *SchoolFocusListType {
        return &x
}

func (n *SchoolFocusListType) New() *SchoolFocusListType {
        return SchoolFocusListTypeCreate(SchoolFocusListType{})
}

func StudentExitStatusContainerTypeCreate(x StudentExitStatusContainerType) *StudentExitStatusContainerType {
        return &x
}

func (n *StudentExitStatusContainerType) New() *StudentExitStatusContainerType {
        return StudentExitStatusContainerTypeCreate(StudentExitStatusContainerType{})
}

func TimeTableDayTypeCreate(x TimeTableDayType) *TimeTableDayType {
        return &x
}

func (n *TimeTableDayType) New() *TimeTableDayType {
        return TimeTableDayTypeCreate(TimeTableDayType{})
}

func EquipmentInfoCreate(x EquipmentInfo) *EquipmentInfo {
        return &x
}

func (n *EquipmentInfo) New() *EquipmentInfo {
        return EquipmentInfoCreate(EquipmentInfo{})
}

func StudentSubjectChoiceListTypeCreate(x StudentSubjectChoiceListType) *StudentSubjectChoiceListType {
        return &x
}

func (n *StudentSubjectChoiceListType) New() *StudentSubjectChoiceListType {
        return StudentSubjectChoiceListTypeCreate(StudentSubjectChoiceListType{})
}

func WellbeingEventLocationDetailsTypeCreate(x WellbeingEventLocationDetailsType) *WellbeingEventLocationDetailsType {
        return &x
}

func (n *WellbeingEventLocationDetailsType) New() *WellbeingEventLocationDetailsType {
        return WellbeingEventLocationDetailsTypeCreate(WellbeingEventLocationDetailsType{})
}

func LibraryItemInfoTypeCreate(x LibraryItemInfoType) *LibraryItemInfoType {
        return &x
}

func (n *LibraryItemInfoType) New() *LibraryItemInfoType {
        return LibraryItemInfoTypeCreate(LibraryItemInfoType{})
}

func TimeElementType_SpanGapCreate(x TimeElementType_SpanGap) *TimeElementType_SpanGap {
        return &x
}

func (n *TimeElementType_SpanGap) New() *TimeElementType_SpanGap {
        return TimeElementType_SpanGapCreate(TimeElementType_SpanGap{})
}

func OtherWellbeingResponseContainerTypeCreate(x OtherWellbeingResponseContainerType) *OtherWellbeingResponseContainerType {
        return &x
}

func (n *OtherWellbeingResponseContainerType) New() *OtherWellbeingResponseContainerType {
        return OtherWellbeingResponseContainerTypeCreate(OtherWellbeingResponseContainerType{})
}

func PasswordListTypeCreate(x PasswordListType) *PasswordListType {
        return &x
}

func (n *PasswordListType) New() *PasswordListType {
        return PasswordListTypeCreate(PasswordListType{})
}

func CodeFrameTestItemTypeCreate(x CodeFrameTestItemType) *CodeFrameTestItemType {
        return &x
}

func (n *CodeFrameTestItemType) New() *CodeFrameTestItemType {
        return CodeFrameTestItemTypeCreate(CodeFrameTestItemType{})
}

func LearningStandardTypeCreate(x LearningStandardType) *LearningStandardType {
        return &x
}

func (n *LearningStandardType) New() *LearningStandardType {
        return LearningStandardTypeCreate(LearningStandardType{})
}

func ActivityTimeType_DurationCreate(x ActivityTimeType_Duration) *ActivityTimeType_Duration {
        return &x
}

func (n *ActivityTimeType_Duration) New() *ActivityTimeType_Duration {
        return ActivityTimeType_DurationCreate(ActivityTimeType_Duration{})
}

func NameOfRecordTypeCreate(x NameOfRecordType) *NameOfRecordType {
        return &x
}

func (n *NameOfRecordType) New() *NameOfRecordType {
        return NameOfRecordTypeCreate(NameOfRecordType{})
}

func SubstituteItemTypeCreate(x SubstituteItemType) *SubstituteItemType {
        return &x
}

func (n *SubstituteItemType) New() *SubstituteItemType {
        return SubstituteItemTypeCreate(SubstituteItemType{})
}

func TimeTableScheduleTypeCreate(x TimeTableScheduleType) *TimeTableScheduleType {
        return &x
}

func (n *TimeTableScheduleType) New() *TimeTableScheduleType {
        return TimeTableScheduleTypeCreate(TimeTableScheduleType{})
}

func OrganizationsTypeCreate(x OrganizationsType) *OrganizationsType {
        return &x
}

func (n *OrganizationsType) New() *OrganizationsType {
        return OrganizationsTypeCreate(OrganizationsType{})
}

func ResourceUsage_ResourceReportColumnCreate(x ResourceUsage_ResourceReportColumn) *ResourceUsage_ResourceReportColumn {
        return &x
}

func (n *ResourceUsage_ResourceReportColumn) New() *ResourceUsage_ResourceReportColumn {
        return ResourceUsage_ResourceReportColumnCreate(ResourceUsage_ResourceReportColumn{})
}

func WellbeingAlertCreate(x WellbeingAlert) *WellbeingAlert {
        return &x
}

func (n *WellbeingAlert) New() *WellbeingAlert {
        return WellbeingAlertCreate(WellbeingAlert{})
}

func TeachingGroupListTypeCreate(x TeachingGroupListType) *TeachingGroupListType {
        return &x
}

func (n *TeachingGroupListType) New() *TeachingGroupListType {
        return TeachingGroupListTypeCreate(TeachingGroupListType{})
}

func HoldInfoListTypeCreate(x HoldInfoListType) *HoldInfoListType {
        return &x
}

func (n *HoldInfoListType) New() *HoldInfoListType {
        return HoldInfoListTypeCreate(HoldInfoListType{})
}

func PhoneNumberListTypeCreate(x PhoneNumberListType) *PhoneNumberListType {
        return &x
}

func (n *PhoneNumberListType) New() *PhoneNumberListType {
        return PhoneNumberListTypeCreate(PhoneNumberListType{})
}

func NAPTestContentTypeCreate(x NAPTestContentType) *NAPTestContentType {
        return &x
}

func (n *NAPTestContentType) New() *NAPTestContentType {
        return NAPTestContentTypeCreate(NAPTestContentType{})
}

func WellbeingEventCreate(x WellbeingEvent) *WellbeingEvent {
        return &x
}

func (n *WellbeingEvent) New() *WellbeingEvent {
        return WellbeingEventCreate(WellbeingEvent{})
}

func WellbeingPersonLinkCreate(x WellbeingPersonLink) *WellbeingPersonLink {
        return &x
}

func (n *WellbeingPersonLink) New() *WellbeingPersonLink {
        return WellbeingPersonLinkCreate(WellbeingPersonLink{})
}

func MediaTypesTypeCreate(x MediaTypesType) *MediaTypesType {
        return &x
}

func (n *MediaTypesType) New() *MediaTypesType {
        return MediaTypesTypeCreate(MediaTypesType{})
}

func NAPTestletResponseTypeCreate(x NAPTestletResponseType) *NAPTestletResponseType {
        return &x
}

func (n *NAPTestletResponseType) New() *NAPTestletResponseType {
        return NAPTestletResponseTypeCreate(NAPTestletResponseType{})
}

func LearningStandardListTypeCreate(x LearningStandardListType) *LearningStandardListType {
        return &x
}

func (n *LearningStandardListType) New() *LearningStandardListType {
        return LearningStandardListTypeCreate(LearningStandardListType{})
}

func TimeTableContainerCreate(x TimeTableContainer) *TimeTableContainer {
        return &x
}

func (n *TimeTableContainer) New() *TimeTableContainer {
        return TimeTableContainerCreate(TimeTableContainer{})
}

func CensusReportingListTypeCreate(x CensusReportingListType) *CensusReportingListType {
        return &x
}

func (n *CensusReportingListType) New() *CensusReportingListType {
        return CensusReportingListTypeCreate(CensusReportingListType{})
}

func DebitOrCreditAmountTypeCreate(x DebitOrCreditAmountType) *DebitOrCreditAmountType {
        return &x
}

func (n *DebitOrCreditAmountType) New() *DebitOrCreditAmountType {
        return DebitOrCreditAmountTypeCreate(DebitOrCreditAmountType{})
}

func RoomInfoCreate(x RoomInfo) *RoomInfo {
        return &x
}

func (n *RoomInfo) New() *RoomInfo {
        return RoomInfoCreate(RoomInfo{})
}

func Debtor_BilledEntityCreate(x Debtor_BilledEntity) *Debtor_BilledEntity {
        return &x
}

func (n *Debtor_BilledEntity) New() *Debtor_BilledEntity {
        return Debtor_BilledEntityCreate(Debtor_BilledEntity{})
}

func SystemRole_RoleScopeCreate(x SystemRole_RoleScope) *SystemRole_RoleScope {
        return &x
}

func (n *SystemRole_RoleScope) New() *SystemRole_RoleScope {
        return SystemRole_RoleScopeCreate(SystemRole_RoleScope{})
}

func PersonInfoTypeCreate(x PersonInfoType) *PersonInfoType {
        return &x
}

func (n *PersonInfoType) New() *PersonInfoType {
        return PersonInfoTypeCreate(PersonInfoType{})
}

func ValidLetterMarkListTypeCreate(x ValidLetterMarkListType) *ValidLetterMarkListType {
        return &x
}

func (n *ValidLetterMarkListType) New() *ValidLetterMarkListType {
        return ValidLetterMarkListTypeCreate(ValidLetterMarkListType{})
}

func JournalAdjustmentListTypeCreate(x JournalAdjustmentListType) *JournalAdjustmentListType {
        return &x
}

func (n *JournalAdjustmentListType) New() *JournalAdjustmentListType {
        return JournalAdjustmentListTypeCreate(JournalAdjustmentListType{})
}

func SchoolCourseInfoOverrideTypeCreate(x SchoolCourseInfoOverrideType) *SchoolCourseInfoOverrideType {
        return &x
}

func (n *SchoolCourseInfoOverrideType) New() *SchoolCourseInfoOverrideType {
        return SchoolCourseInfoOverrideTypeCreate(SchoolCourseInfoOverrideType{})
}

func ResourceUsageCreate(x ResourceUsage) *ResourceUsage {
        return &x
}

func (n *ResourceUsage) New() *ResourceUsage {
        return ResourceUsageCreate(ResourceUsage{})
}

func AuditInfoTypeCreate(x AuditInfoType) *AuditInfoType {
        return &x
}

func (n *AuditInfoType) New() *AuditInfoType {
        return AuditInfoTypeCreate(AuditInfoType{})
}

func NAPWritingRubricListTypeCreate(x NAPWritingRubricListType) *NAPWritingRubricListType {
        return &x
}

func (n *NAPWritingRubricListType) New() *NAPWritingRubricListType {
        return NAPWritingRubricListTypeCreate(NAPWritingRubricListType{})
}

func PersonPictureCreate(x PersonPicture) *PersonPicture {
        return &x
}

func (n *PersonPicture) New() *PersonPicture {
        return PersonPictureCreate(PersonPicture{})
}

func RelationshipTypeCreate(x RelationshipType) *RelationshipType {
        return &x
}

func (n *RelationshipType) New() *RelationshipType {
        return RelationshipTypeCreate(RelationshipType{})
}

func ContactFlagsTypeCreate(x ContactFlagsType) *ContactFlagsType {
        return &x
}

func (n *ContactFlagsType) New() *ContactFlagsType {
        return ContactFlagsTypeCreate(ContactFlagsType{})
}

func PersonPicture_ParentObjectRefIdCreate(x PersonPicture_ParentObjectRefId) *PersonPicture_ParentObjectRefId {
        return &x
}

func (n *PersonPicture_ParentObjectRefId) New() *PersonPicture_ParentObjectRefId {
        return PersonPicture_ParentObjectRefIdCreate(PersonPicture_ParentObjectRefId{})
}

func PersonInvolvementType_PersonRefIdCreate(x PersonInvolvementType_PersonRefId) *PersonInvolvementType_PersonRefId {
        return &x
}

func (n *PersonInvolvementType_PersonRefId) New() *PersonInvolvementType_PersonRefId {
        return PersonInvolvementType_PersonRefIdCreate(PersonInvolvementType_PersonRefId{})
}

func StudentScoreJudgementAgainstStandardCreate(x StudentScoreJudgementAgainstStandard) *StudentScoreJudgementAgainstStandard {
        return &x
}

func (n *StudentScoreJudgementAgainstStandard) New() *StudentScoreJudgementAgainstStandard {
        return StudentScoreJudgementAgainstStandardCreate(StudentScoreJudgementAgainstStandard{})
}

func ProgramStatusTypeCreate(x ProgramStatusType) *ProgramStatusType {
        return &x
}

func (n *ProgramStatusType) New() *ProgramStatusType {
        return ProgramStatusTypeCreate(ProgramStatusType{})
}

func SchoolCourseInfoCreate(x SchoolCourseInfo) *SchoolCourseInfo {
        return &x
}

func (n *SchoolCourseInfo) New() *SchoolCourseInfo {
        return SchoolCourseInfoCreate(SchoolCourseInfo{})
}

func TimeTablePeriodTypeCreate(x TimeTablePeriodType) *TimeTablePeriodType {
        return &x
}

func (n *TimeTablePeriodType) New() *TimeTablePeriodType {
        return TimeTablePeriodTypeCreate(TimeTablePeriodType{})
}

func AssociatedObjectsTypeCreate(x AssociatedObjectsType) *AssociatedObjectsType {
        return &x
}

func (n *AssociatedObjectsType) New() *AssociatedObjectsType {
        return AssociatedObjectsTypeCreate(AssociatedObjectsType{})
}

func SystemRole_RoleScopeRefIdCreate(x SystemRole_RoleScopeRefId) *SystemRole_RoleScopeRefId {
        return &x
}

func (n *SystemRole_RoleScopeRefId) New() *SystemRole_RoleScopeRefId {
        return SystemRole_RoleScopeRefIdCreate(SystemRole_RoleScopeRefId{})
}

func TeachingGroupScheduleListTypeCreate(x TeachingGroupScheduleListType) *TeachingGroupScheduleListType {
        return &x
}

func (n *TeachingGroupScheduleListType) New() *TeachingGroupScheduleListType {
        return TeachingGroupScheduleListTypeCreate(TeachingGroupScheduleListType{})
}

func AbstractContentElementTypeCreate(x AbstractContentElementType) *AbstractContentElementType {
        return &x
}

func (n *AbstractContentElementType) New() *AbstractContentElementType {
        return AbstractContentElementTypeCreate(AbstractContentElementType{})
}

func TimeTableScheduleCellListTypeCreate(x TimeTableScheduleCellListType) *TimeTableScheduleCellListType {
        return &x
}

func (n *TimeTableScheduleCellListType) New() *TimeTableScheduleCellListType {
        return TimeTableScheduleCellListTypeCreate(TimeTableScheduleCellListType{})
}

func ResourceUsage_ResourceReportLineCreate(x ResourceUsage_ResourceReportLine) *ResourceUsage_ResourceReportLine {
        return &x
}

func (n *ResourceUsage_ResourceReportLine) New() *ResourceUsage_ResourceReportLine {
        return ResourceUsage_ResourceReportLineCreate(ResourceUsage_ResourceReportLine{})
}

func ScoreTypeCreate(x ScoreType) *ScoreType {
        return &x
}

func (n *ScoreType) New() *ScoreType {
        return ScoreTypeCreate(ScoreType{})
}

func NAPWritingRubricTypeCreate(x NAPWritingRubricType) *NAPWritingRubricType {
        return &x
}

func (n *NAPWritingRubricType) New() *NAPWritingRubricType {
        return NAPWritingRubricTypeCreate(NAPWritingRubricType{})
}

func SchoolGroupListTypeCreate(x SchoolGroupListType) *SchoolGroupListType {
        return &x
}

func (n *SchoolGroupListType) New() *SchoolGroupListType {
        return SchoolGroupListTypeCreate(SchoolGroupListType{})
}

func OtherCodeListType_OtherCodeCreate(x OtherCodeListType_OtherCode) *OtherCodeListType_OtherCode {
        return &x
}

func (n *OtherCodeListType_OtherCode) New() *OtherCodeListType_OtherCode {
        return OtherCodeListType_OtherCodeCreate(OtherCodeListType_OtherCode{})
}

func ReligionTypeCreate(x ReligionType) *ReligionType {
        return &x
}

func (n *ReligionType) New() *ReligionType {
        return ReligionTypeCreate(ReligionType{})
}

func ACStrandAreaListTypeCreate(x ACStrandAreaListType) *ACStrandAreaListType {
        return &x
}

func (n *ACStrandAreaListType) New() *ACStrandAreaListType {
        return ACStrandAreaListTypeCreate(ACStrandAreaListType{})
}

  func (t *StatisticalAreasType) Add(value StatisticalAreaType) *StatisticalAreasType {
    
        if t == nil {
                t = StatisticalAreasTypeCreate(StatisticalAreasType{})
        }
        if t.StatisticalArea == nil {
                t.StatisticalArea = make([]StatisticalAreaType, 0)
        }
        t.StatisticalArea = append(t.StatisticalArea, value)
        return t
}

func (t *StatisticalAreasType) AddNew() *StatisticalAreasType {
        if t == nil {
                t = StatisticalAreasTypeCreate(StatisticalAreasType{})
        }
        if t.StatisticalArea == nil {
                t.StatisticalArea = make([]StatisticalAreaType, 0)
        }
        t.StatisticalArea = append(t.StatisticalArea, StatisticalAreaType{})
        return t
}

func (t *StatisticalAreasType) Last() *StatisticalAreaType {
        return &(t.StatisticalArea[len(t.StatisticalArea)-1])
}

  func (t *WellbeingEventCategoryListType) Add(value WellbeingEventCategoryType) *WellbeingEventCategoryListType {
    
        if t == nil {
                t = WellbeingEventCategoryListTypeCreate(WellbeingEventCategoryListType{})
        }
        if t.WellbeingEventCategory == nil {
                t.WellbeingEventCategory = make([]WellbeingEventCategoryType, 0)
        }
        t.WellbeingEventCategory = append(t.WellbeingEventCategory, value)
        return t
}

func (t *WellbeingEventCategoryListType) AddNew() *WellbeingEventCategoryListType {
        if t == nil {
                t = WellbeingEventCategoryListTypeCreate(WellbeingEventCategoryListType{})
        }
        if t.WellbeingEventCategory == nil {
                t.WellbeingEventCategory = make([]WellbeingEventCategoryType, 0)
        }
        t.WellbeingEventCategory = append(t.WellbeingEventCategory, WellbeingEventCategoryType{})
        return t
}

func (t *WellbeingEventCategoryListType) Last() *WellbeingEventCategoryType {
        return &(t.WellbeingEventCategory[len(t.WellbeingEventCategory)-1])
}

  func (t *CensusStudentListType) Add(value CensusStudentType) *CensusStudentListType {
    
        if t == nil {
                t = CensusStudentListTypeCreate(CensusStudentListType{})
        }
        if t.CensusStudent == nil {
                t.CensusStudent = make([]CensusStudentType, 0)
        }
        t.CensusStudent = append(t.CensusStudent, value)
        return t
}

func (t *CensusStudentListType) AddNew() *CensusStudentListType {
        if t == nil {
                t = CensusStudentListTypeCreate(CensusStudentListType{})
        }
        if t.CensusStudent == nil {
                t.CensusStudent = make([]CensusStudentType, 0)
        }
        t.CensusStudent = append(t.CensusStudent, CensusStudentType{})
        return t
}

func (t *CensusStudentListType) Last() *CensusStudentType {
        return &(t.CensusStudent[len(t.CensusStudent)-1])
}

  func (t *StaffSubjectListType) Add(value StaffSubjectType) *StaffSubjectListType {
    
        if t == nil {
                t = StaffSubjectListTypeCreate(StaffSubjectListType{})
        }
        if t.StaffSubject == nil {
                t.StaffSubject = make([]StaffSubjectType, 0)
        }
        t.StaffSubject = append(t.StaffSubject, value)
        return t
}

func (t *StaffSubjectListType) AddNew() *StaffSubjectListType {
        if t == nil {
                t = StaffSubjectListTypeCreate(StaffSubjectListType{})
        }
        if t.StaffSubject == nil {
                t.StaffSubject = make([]StaffSubjectType, 0)
        }
        t.StaffSubject = append(t.StaffSubject, StaffSubjectType{})
        return t
}

func (t *StaffSubjectListType) Last() *StaffSubjectType {
        return &(t.StaffSubject[len(t.StaffSubject)-1])
}

  func (t *AssignmentListType) Add(value string) *AssignmentListType {
    
        if t == nil {
                t = AssignmentListTypeCreate(AssignmentListType{})
        }
        if t.GradingAssignmentRefId == nil {
                t.GradingAssignmentRefId = make([]string, 0)
        }
        t.GradingAssignmentRefId = append(t.GradingAssignmentRefId, value)
        return t
}

func (t *AssignmentListType) AddNew() *AssignmentListType {
        if t == nil {
                t = AssignmentListTypeCreate(AssignmentListType{})
        }
        if t.GradingAssignmentRefId == nil {
                t.GradingAssignmentRefId = make([]string, 0)
        }
        t.GradingAssignmentRefId = append(t.GradingAssignmentRefId, "")
        return t
}

func (t *AssignmentListType) Last() *string {
        return &(t.GradingAssignmentRefId[len(t.GradingAssignmentRefId)-1])
}

  func (t *AGRuleListType) Add(value AGRuleType) *AGRuleListType {
    
        if t == nil {
                t = AGRuleListTypeCreate(AGRuleListType{})
        }
        if t.AGRule == nil {
                t.AGRule = make([]AGRuleType, 0)
        }
        t.AGRule = append(t.AGRule, value)
        return t
}

func (t *AGRuleListType) AddNew() *AGRuleListType {
        if t == nil {
                t = AGRuleListTypeCreate(AGRuleListType{})
        }
        if t.AGRule == nil {
                t.AGRule = make([]AGRuleType, 0)
        }
        t.AGRule = append(t.AGRule, AGRuleType{})
        return t
}

func (t *AGRuleListType) Last() *AGRuleType {
        return &(t.AGRule[len(t.AGRule)-1])
}

  func (t *AGReportingObjectResponseListType) Add(value AGReportingObjectResponseType) *AGReportingObjectResponseListType {
    
        if t == nil {
                t = AGReportingObjectResponseListTypeCreate(AGReportingObjectResponseListType{})
        }
        if t.AGReportingObjectResponse == nil {
                t.AGReportingObjectResponse = make([]AGReportingObjectResponseType, 0)
        }
        t.AGReportingObjectResponse = append(t.AGReportingObjectResponse, value)
        return t
}

func (t *AGReportingObjectResponseListType) AddNew() *AGReportingObjectResponseListType {
        if t == nil {
                t = AGReportingObjectResponseListTypeCreate(AGReportingObjectResponseListType{})
        }
        if t.AGReportingObjectResponse == nil {
                t.AGReportingObjectResponse = make([]AGReportingObjectResponseType, 0)
        }
        t.AGReportingObjectResponse = append(t.AGReportingObjectResponse, AGReportingObjectResponseType{})
        return t
}

func (t *AGReportingObjectResponseListType) Last() *AGReportingObjectResponseType {
        return &(t.AGReportingObjectResponse[len(t.AGReportingObjectResponse)-1])
}

  func (t *ResourceUsage_ResourceReportColumnList) Add(value ResourceUsage_ResourceReportColumn) *ResourceUsage_ResourceReportColumnList {
    
        if t == nil {
                t = ResourceUsage_ResourceReportColumnListCreate(ResourceUsage_ResourceReportColumnList{})
        }
        if t.ResourceReportColumn == nil {
                t.ResourceReportColumn = make([]ResourceUsage_ResourceReportColumn, 0)
        }
        t.ResourceReportColumn = append(t.ResourceReportColumn, value)
        return t
}

func (t *ResourceUsage_ResourceReportColumnList) AddNew() *ResourceUsage_ResourceReportColumnList {
        if t == nil {
                t = ResourceUsage_ResourceReportColumnListCreate(ResourceUsage_ResourceReportColumnList{})
        }
        if t.ResourceReportColumn == nil {
                t.ResourceReportColumn = make([]ResourceUsage_ResourceReportColumn, 0)
        }
        t.ResourceReportColumn = append(t.ResourceReportColumn, ResourceUsage_ResourceReportColumn{})
        return t
}

func (t *ResourceUsage_ResourceReportColumnList) Last() *ResourceUsage_ResourceReportColumn {
        return &(t.ResourceReportColumn[len(t.ResourceReportColumn)-1])
}

  func (t *LibraryMessageListType) Add(value LibraryMessageType) *LibraryMessageListType {
    
        if t == nil {
                t = LibraryMessageListTypeCreate(LibraryMessageListType{})
        }
        if t.Message == nil {
                t.Message = make([]LibraryMessageType, 0)
        }
        t.Message = append(t.Message, value)
        return t
}

func (t *LibraryMessageListType) AddNew() *LibraryMessageListType {
        if t == nil {
                t = LibraryMessageListTypeCreate(LibraryMessageListType{})
        }
        if t.Message == nil {
                t.Message = make([]LibraryMessageType, 0)
        }
        t.Message = append(t.Message, LibraryMessageType{})
        return t
}

func (t *LibraryMessageListType) Last() *LibraryMessageType {
        return &(t.Message[len(t.Message)-1])
}

  func (t *SIF_MetadataType_TimeElements) Add(value TimeElementType) *SIF_MetadataType_TimeElements {
    
        if t == nil {
                t = SIF_MetadataType_TimeElementsCreate(SIF_MetadataType_TimeElements{})
        }
        if t.TimeElement == nil {
                t.TimeElement = make([]TimeElementType, 0)
        }
        t.TimeElement = append(t.TimeElement, value)
        return t
}

func (t *SIF_MetadataType_TimeElements) AddNew() *SIF_MetadataType_TimeElements {
        if t == nil {
                t = SIF_MetadataType_TimeElementsCreate(SIF_MetadataType_TimeElements{})
        }
        if t.TimeElement == nil {
                t.TimeElement = make([]TimeElementType, 0)
        }
        t.TimeElement = append(t.TimeElement, TimeElementType{})
        return t
}

func (t *SIF_MetadataType_TimeElements) Last() *TimeElementType {
        return &(t.TimeElement[len(t.TimeElement)-1])
}

  func (t *StimulusLocalIdListType) Add(value LocalIdType) *StimulusLocalIdListType {
    
        if t == nil {
                t = StimulusLocalIdListTypeCreate(StimulusLocalIdListType{})
        }
        if t.StimulusLocalId == nil {
                t.StimulusLocalId = make([]LocalIdType, 0)
        }
        t.StimulusLocalId = append(t.StimulusLocalId, value)
        return t
}

func (t *StimulusLocalIdListType) AddNew() *StimulusLocalIdListType {
        if t == nil {
                t = StimulusLocalIdListTypeCreate(StimulusLocalIdListType{})
        }
        if t.StimulusLocalId == nil {
                t.StimulusLocalId = make([]LocalIdType, 0)
        }
        t.StimulusLocalId = append(t.StimulusLocalId, "")
        return t
}

func (t *StimulusLocalIdListType) Last() *LocalIdType {
        return &(t.StimulusLocalId[len(t.StimulusLocalId)-1])
}

  func (t *StudentAttendanceCollectionReportingListType) Add(value StudentAttendanceCollectionReportingType) *StudentAttendanceCollectionReportingListType {
    
        if t == nil {
                t = StudentAttendanceCollectionReportingListTypeCreate(StudentAttendanceCollectionReportingListType{})
        }
        if t.StudentAttendanceCollectionReporting == nil {
                t.StudentAttendanceCollectionReporting = make([]StudentAttendanceCollectionReportingType, 0)
        }
        t.StudentAttendanceCollectionReporting = append(t.StudentAttendanceCollectionReporting, value)
        return t
}

func (t *StudentAttendanceCollectionReportingListType) AddNew() *StudentAttendanceCollectionReportingListType {
        if t == nil {
                t = StudentAttendanceCollectionReportingListTypeCreate(StudentAttendanceCollectionReportingListType{})
        }
        if t.StudentAttendanceCollectionReporting == nil {
                t.StudentAttendanceCollectionReporting = make([]StudentAttendanceCollectionReportingType, 0)
        }
        t.StudentAttendanceCollectionReporting = append(t.StudentAttendanceCollectionReporting, StudentAttendanceCollectionReportingType{})
        return t
}

func (t *StudentAttendanceCollectionReportingListType) Last() *StudentAttendanceCollectionReportingType {
        return &(t.StudentAttendanceCollectionReporting[len(t.StudentAttendanceCollectionReporting)-1])
}

  func (t *StudentsType) Add(value string) *StudentsType {
    
        if t == nil {
                t = StudentsTypeCreate(StudentsType{})
        }
        if t.StudentPersonalRefId == nil {
                t.StudentPersonalRefId = make([]string, 0)
        }
        t.StudentPersonalRefId = append(t.StudentPersonalRefId, value)
        return t
}

func (t *StudentsType) AddNew() *StudentsType {
        if t == nil {
                t = StudentsTypeCreate(StudentsType{})
        }
        if t.StudentPersonalRefId == nil {
                t.StudentPersonalRefId = make([]string, 0)
        }
        t.StudentPersonalRefId = append(t.StudentPersonalRefId, "")
        return t
}

func (t *StudentsType) Last() *string {
        return &(t.StudentPersonalRefId[len(t.StudentPersonalRefId)-1])
}

  func (t *AddressListType) Add(value AddressType) *AddressListType {
    
        if t == nil {
                t = AddressListTypeCreate(AddressListType{})
        }
        if t.Address == nil {
                t.Address = make([]AddressType, 0)
        }
        t.Address = append(t.Address, value)
        return t
}

func (t *AddressListType) AddNew() *AddressListType {
        if t == nil {
                t = AddressListTypeCreate(AddressListType{})
        }
        if t.Address == nil {
                t.Address = make([]AddressType, 0)
        }
        t.Address = append(t.Address, AddressType{})
        return t
}

func (t *AddressListType) Last() *AddressType {
        return &(t.Address[len(t.Address)-1])
}

  func (t *LifeCycleType_Creators) Add(value LifeCycleType_Creator) *LifeCycleType_Creators {
    
        if t == nil {
                t = LifeCycleType_CreatorsCreate(LifeCycleType_Creators{})
        }
        if t.Creator == nil {
                t.Creator = make([]LifeCycleType_Creator, 0)
        }
        t.Creator = append(t.Creator, value)
        return t
}

func (t *LifeCycleType_Creators) AddNew() *LifeCycleType_Creators {
        if t == nil {
                t = LifeCycleType_CreatorsCreate(LifeCycleType_Creators{})
        }
        if t.Creator == nil {
                t.Creator = make([]LifeCycleType_Creator, 0)
        }
        t.Creator = append(t.Creator, LifeCycleType_Creator{})
        return t
}

func (t *LifeCycleType_Creators) Last() *LifeCycleType_Creator {
        return &(t.Creator[len(t.Creator)-1])
}

  func (t *StatementCodesType) Add(value string) *StatementCodesType {
    
        if t == nil {
                t = StatementCodesTypeCreate(StatementCodesType{})
        }
        if t.StatementCode == nil {
                t.StatementCode = make([]string, 0)
        }
        t.StatementCode = append(t.StatementCode, value)
        return t
}

func (t *StatementCodesType) AddNew() *StatementCodesType {
        if t == nil {
                t = StatementCodesTypeCreate(StatementCodesType{})
        }
        if t.StatementCode == nil {
                t.StatementCode = make([]string, 0)
        }
        t.StatementCode = append(t.StatementCode, "")
        return t
}

func (t *StatementCodesType) Last() *string {
        return &(t.StatementCode[len(t.StatementCode)-1])
}

  func (t *PeriodAttendancesType) Add(value PeriodAttendanceType) *PeriodAttendancesType {
    
        if t == nil {
                t = PeriodAttendancesTypeCreate(PeriodAttendancesType{})
        }
        if t.PeriodAttendance == nil {
                t.PeriodAttendance = make([]PeriodAttendanceType, 0)
        }
        t.PeriodAttendance = append(t.PeriodAttendance, value)
        return t
}

func (t *PeriodAttendancesType) AddNew() *PeriodAttendancesType {
        if t == nil {
                t = PeriodAttendancesTypeCreate(PeriodAttendancesType{})
        }
        if t.PeriodAttendance == nil {
                t.PeriodAttendance = make([]PeriodAttendanceType, 0)
        }
        t.PeriodAttendance = append(t.PeriodAttendance, PeriodAttendanceType{})
        return t
}

func (t *PeriodAttendancesType) Last() *PeriodAttendanceType {
        return &(t.PeriodAttendance[len(t.PeriodAttendance)-1])
}

  func (t *NAPTestItemListType) Add(value NAPTestItem2Type) *NAPTestItemListType {
    
        if t == nil {
                t = NAPTestItemListTypeCreate(NAPTestItemListType{})
        }
        if t.TestItem == nil {
                t.TestItem = make([]NAPTestItem2Type, 0)
        }
        t.TestItem = append(t.TestItem, value)
        return t
}

func (t *NAPTestItemListType) AddNew() *NAPTestItemListType {
        if t == nil {
                t = NAPTestItemListTypeCreate(NAPTestItemListType{})
        }
        if t.TestItem == nil {
                t.TestItem = make([]NAPTestItem2Type, 0)
        }
        t.TestItem = append(t.TestItem, NAPTestItem2Type{})
        return t
}

func (t *NAPTestItemListType) Last() *NAPTestItem2Type {
        return &(t.TestItem[len(t.TestItem)-1])
}

  func (t *TimeTablePeriodListType) Add(value TimeTablePeriodType) *TimeTablePeriodListType {
    
        if t == nil {
                t = TimeTablePeriodListTypeCreate(TimeTablePeriodListType{})
        }
        if t.TimeTablePeriod == nil {
                t.TimeTablePeriod = make([]TimeTablePeriodType, 0)
        }
        t.TimeTablePeriod = append(t.TimeTablePeriod, value)
        return t
}

func (t *TimeTablePeriodListType) AddNew() *TimeTablePeriodListType {
        if t == nil {
                t = TimeTablePeriodListTypeCreate(TimeTablePeriodListType{})
        }
        if t.TimeTablePeriod == nil {
                t.TimeTablePeriod = make([]TimeTablePeriodType, 0)
        }
        t.TimeTablePeriod = append(t.TimeTablePeriod, TimeTablePeriodType{})
        return t
}

func (t *TimeTablePeriodListType) Last() *TimeTablePeriodType {
        return &(t.TimeTablePeriod[len(t.TimeTablePeriod)-1])
}

  func (t *OtherCodeListType) Add(value OtherCodeListType_OtherCode) *OtherCodeListType {
    
        if t == nil {
                t = OtherCodeListTypeCreate(OtherCodeListType{})
        }
        if t.OtherCode == nil {
                t.OtherCode = make([]OtherCodeListType_OtherCode, 0)
        }
        t.OtherCode = append(t.OtherCode, value)
        return t
}

func (t *OtherCodeListType) AddNew() *OtherCodeListType {
        if t == nil {
                t = OtherCodeListTypeCreate(OtherCodeListType{})
        }
        if t.OtherCode == nil {
                t.OtherCode = make([]OtherCodeListType_OtherCode, 0)
        }
        t.OtherCode = append(t.OtherCode, OtherCodeListType_OtherCode{})
        return t
}

func (t *OtherCodeListType) Last() *OtherCodeListType_OtherCode {
        return &(t.OtherCode[len(t.OtherCode)-1])
}

  func (t *SystemRole_RoleList) Add(value SystemRole_Role) *SystemRole_RoleList {
    
        if t == nil {
                t = SystemRole_RoleListCreate(SystemRole_RoleList{})
        }
        if t.Role == nil {
                t.Role = make([]SystemRole_Role, 0)
        }
        t.Role = append(t.Role, value)
        return t
}

func (t *SystemRole_RoleList) AddNew() *SystemRole_RoleList {
        if t == nil {
                t = SystemRole_RoleListCreate(SystemRole_RoleList{})
        }
        if t.Role == nil {
                t.Role = make([]SystemRole_Role, 0)
        }
        t.Role = append(t.Role, SystemRole_Role{})
        return t
}

func (t *SystemRole_RoleList) Last() *SystemRole_Role {
        return &(t.Role[len(t.Role)-1])
}

  func (t *ContentDescriptionListType) Add(value string) *ContentDescriptionListType {
    
        if t == nil {
                t = ContentDescriptionListTypeCreate(ContentDescriptionListType{})
        }
        if t.ContentDescription == nil {
                t.ContentDescription = make([]string, 0)
        }
        t.ContentDescription = append(t.ContentDescription, value)
        return t
}

func (t *ContentDescriptionListType) AddNew() *ContentDescriptionListType {
        if t == nil {
                t = ContentDescriptionListTypeCreate(ContentDescriptionListType{})
        }
        if t.ContentDescription == nil {
                t.ContentDescription = make([]string, 0)
        }
        t.ContentDescription = append(t.ContentDescription, "")
        return t
}

func (t *ContentDescriptionListType) Last() *string {
        return &(t.ContentDescription[len(t.ContentDescription)-1])
}

  func (t *CalendarSummaryListType) Add(value string) *CalendarSummaryListType {
    
        if t == nil {
                t = CalendarSummaryListTypeCreate(CalendarSummaryListType{})
        }
        if t.CalendarSummaryRefId == nil {
                t.CalendarSummaryRefId = make([]string, 0)
        }
        t.CalendarSummaryRefId = append(t.CalendarSummaryRefId, value)
        return t
}

func (t *CalendarSummaryListType) AddNew() *CalendarSummaryListType {
        if t == nil {
                t = CalendarSummaryListTypeCreate(CalendarSummaryListType{})
        }
        if t.CalendarSummaryRefId == nil {
                t.CalendarSummaryRefId = make([]string, 0)
        }
        t.CalendarSummaryRefId = append(t.CalendarSummaryRefId, "")
        return t
}

func (t *CalendarSummaryListType) Last() *string {
        return &(t.CalendarSummaryRefId[len(t.CalendarSummaryRefId)-1])
}

  func (t *LResourcesType) Add(value ResourcesType) *LResourcesType {
    
        if t == nil {
                t = LResourcesTypeCreate(LResourcesType{})
        }
        if t.LearningResourceRefId == nil {
                t.LearningResourceRefId = make([]ResourcesType, 0)
        }
        t.LearningResourceRefId = append(t.LearningResourceRefId, value)
        return t
}

func (t *LResourcesType) AddNew() *LResourcesType {
        if t == nil {
                t = LResourcesTypeCreate(LResourcesType{})
        }
        if t.LearningResourceRefId == nil {
                t.LearningResourceRefId = make([]ResourcesType, 0)
        }
        t.LearningResourceRefId = append(t.LearningResourceRefId, ResourcesType{})
        return t
}

func (t *LResourcesType) Last() *ResourcesType {
        return &(t.LearningResourceRefId[len(t.LearningResourceRefId)-1])
}

  func (t *ApprovalsType) Add(value ApprovalType) *ApprovalsType {
    
        if t == nil {
                t = ApprovalsTypeCreate(ApprovalsType{})
        }
        if t.Approval == nil {
                t.Approval = make([]ApprovalType, 0)
        }
        t.Approval = append(t.Approval, value)
        return t
}

func (t *ApprovalsType) AddNew() *ApprovalsType {
        if t == nil {
                t = ApprovalsTypeCreate(ApprovalsType{})
        }
        if t.Approval == nil {
                t.Approval = make([]ApprovalType, 0)
        }
        t.Approval = append(t.Approval, ApprovalType{})
        return t
}

func (t *ApprovalsType) Last() *ApprovalType {
        return &(t.Approval[len(t.Approval)-1])
}

  func (t *VisaSubClassListType) Add(value VisaSubClassType) *VisaSubClassListType {
    
        if t == nil {
                t = VisaSubClassListTypeCreate(VisaSubClassListType{})
        }
        if t.VisaSubClass == nil {
                t.VisaSubClass = make([]VisaSubClassType, 0)
        }
        t.VisaSubClass = append(t.VisaSubClass, value)
        return t
}

func (t *VisaSubClassListType) AddNew() *VisaSubClassListType {
        if t == nil {
                t = VisaSubClassListTypeCreate(VisaSubClassListType{})
        }
        if t.VisaSubClass == nil {
                t.VisaSubClass = make([]VisaSubClassType, 0)
        }
        t.VisaSubClass = append(t.VisaSubClass, VisaSubClassType{})
        return t
}

func (t *VisaSubClassListType) Last() *VisaSubClassType {
        return &(t.VisaSubClass[len(t.VisaSubClass)-1])
}

  func (t *LibraryTransactionListType) Add(value LibraryTransactionType) *LibraryTransactionListType {
    
        if t == nil {
                t = LibraryTransactionListTypeCreate(LibraryTransactionListType{})
        }
        if t.Transaction == nil {
                t.Transaction = make([]LibraryTransactionType, 0)
        }
        t.Transaction = append(t.Transaction, value)
        return t
}

func (t *LibraryTransactionListType) AddNew() *LibraryTransactionListType {
        if t == nil {
                t = LibraryTransactionListTypeCreate(LibraryTransactionListType{})
        }
        if t.Transaction == nil {
                t.Transaction = make([]LibraryTransactionType, 0)
        }
        t.Transaction = append(t.Transaction, LibraryTransactionType{})
        return t
}

func (t *LibraryTransactionListType) Last() *LibraryTransactionType {
        return &(t.Transaction[len(t.Transaction)-1])
}

  func (t *EmailListType) Add(value EmailType) *EmailListType {
    
        if t == nil {
                t = EmailListTypeCreate(EmailListType{})
        }
        if t.Email == nil {
                t.Email = make([]EmailType, 0)
        }
        t.Email = append(t.Email, value)
        return t
}

func (t *EmailListType) AddNew() *EmailListType {
        if t == nil {
                t = EmailListTypeCreate(EmailListType{})
        }
        if t.Email == nil {
                t.Email = make([]EmailType, 0)
        }
        t.Email = append(t.Email, EmailType{})
        return t
}

func (t *EmailListType) Last() *EmailType {
        return &(t.Email[len(t.Email)-1])
}

  func (t *PaymentReceiptLineListType) Add(value PaymentReceiptLineType) *PaymentReceiptLineListType {
    
        if t == nil {
                t = PaymentReceiptLineListTypeCreate(PaymentReceiptLineListType{})
        }
        if t.PaymentReceiptLine == nil {
                t.PaymentReceiptLine = make([]PaymentReceiptLineType, 0)
        }
        t.PaymentReceiptLine = append(t.PaymentReceiptLine, value)
        return t
}

func (t *PaymentReceiptLineListType) AddNew() *PaymentReceiptLineListType {
        if t == nil {
                t = PaymentReceiptLineListTypeCreate(PaymentReceiptLineListType{})
        }
        if t.PaymentReceiptLine == nil {
                t.PaymentReceiptLine = make([]PaymentReceiptLineType, 0)
        }
        t.PaymentReceiptLine = append(t.PaymentReceiptLine, PaymentReceiptLineType{})
        return t
}

func (t *PaymentReceiptLineListType) Last() *PaymentReceiptLineType {
        return &(t.PaymentReceiptLine[len(t.PaymentReceiptLine)-1])
}

  func (t *ElectronicIdListType) Add(value ElectronicIdType) *ElectronicIdListType {
    
        if t == nil {
                t = ElectronicIdListTypeCreate(ElectronicIdListType{})
        }
        if t.ElectronicId == nil {
                t.ElectronicId = make([]ElectronicIdType, 0)
        }
        t.ElectronicId = append(t.ElectronicId, value)
        return t
}

func (t *ElectronicIdListType) AddNew() *ElectronicIdListType {
        if t == nil {
                t = ElectronicIdListTypeCreate(ElectronicIdListType{})
        }
        if t.ElectronicId == nil {
                t.ElectronicId = make([]ElectronicIdType, 0)
        }
        t.ElectronicId = append(t.ElectronicId, ElectronicIdType{})
        return t
}

func (t *ElectronicIdListType) Last() *ElectronicIdType {
        return &(t.ElectronicId[len(t.ElectronicId)-1])
}

  func (t *AccountCodeListType) Add(value string) *AccountCodeListType {
    
        if t == nil {
                t = AccountCodeListTypeCreate(AccountCodeListType{})
        }
        if t.AccountCode == nil {
                t.AccountCode = make([]string, 0)
        }
        t.AccountCode = append(t.AccountCode, value)
        return t
}

func (t *AccountCodeListType) AddNew() *AccountCodeListType {
        if t == nil {
                t = AccountCodeListTypeCreate(AccountCodeListType{})
        }
        if t.AccountCode == nil {
                t.AccountCode = make([]string, 0)
        }
        t.AccountCode = append(t.AccountCode, "")
        return t
}

func (t *AccountCodeListType) Last() *string {
        return &(t.AccountCode[len(t.AccountCode)-1])
}

  func (t *ScoreDescriptionListType) Add(value ScoreDescriptionType) *ScoreDescriptionListType {
    
        if t == nil {
                t = ScoreDescriptionListTypeCreate(ScoreDescriptionListType{})
        }
        if t.ScoreDescription == nil {
                t.ScoreDescription = make([]ScoreDescriptionType, 0)
        }
        t.ScoreDescription = append(t.ScoreDescription, value)
        return t
}

func (t *ScoreDescriptionListType) AddNew() *ScoreDescriptionListType {
        if t == nil {
                t = ScoreDescriptionListTypeCreate(ScoreDescriptionListType{})
        }
        if t.ScoreDescription == nil {
                t.ScoreDescription = make([]ScoreDescriptionType, 0)
        }
        t.ScoreDescription = append(t.ScoreDescription, ScoreDescriptionType{})
        return t
}

func (t *ScoreDescriptionListType) Last() *ScoreDescriptionType {
        return &(t.ScoreDescription[len(t.ScoreDescription)-1])
}

  func (t *CountryListType) Add(value CountryType) *CountryListType {
    
        if t == nil {
                t = CountryListTypeCreate(CountryListType{})
        }
        if t.CountryOfCitizenship == nil {
                t.CountryOfCitizenship = make([]CountryType, 0)
        }
        t.CountryOfCitizenship = append(t.CountryOfCitizenship, value)
        return t
}

func (t *CountryListType) AddNew() *CountryListType {
        if t == nil {
                t = CountryListTypeCreate(CountryListType{})
        }
        if t.CountryOfCitizenship == nil {
                t.CountryOfCitizenship = make([]CountryType, 0)
        }
        t.CountryOfCitizenship = append(t.CountryOfCitizenship, "")
        return t
}

func (t *CountryListType) Last() *CountryType {
        return &(t.CountryOfCitizenship[len(t.CountryOfCitizenship)-1])
}

  func (t *StatementsType) Add(value string) *StatementsType {
    
        if t == nil {
                t = StatementsTypeCreate(StatementsType{})
        }
        if t.Statement == nil {
                t.Statement = make([]string, 0)
        }
        t.Statement = append(t.Statement, value)
        return t
}

func (t *StatementsType) AddNew() *StatementsType {
        if t == nil {
                t = StatementsTypeCreate(StatementsType{})
        }
        if t.Statement == nil {
                t.Statement = make([]string, 0)
        }
        t.Statement = append(t.Statement, "")
        return t
}

func (t *StatementsType) Last() *string {
        return &(t.Statement[len(t.Statement)-1])
}

  func (t *FQItemListType) Add(value FQItemType) *FQItemListType {
    
        if t == nil {
                t = FQItemListTypeCreate(FQItemListType{})
        }
        if t.FQItem == nil {
                t.FQItem = make([]FQItemType, 0)
        }
        t.FQItem = append(t.FQItem, value)
        return t
}

func (t *FQItemListType) AddNew() *FQItemListType {
        if t == nil {
                t = FQItemListTypeCreate(FQItemListType{})
        }
        if t.FQItem == nil {
                t.FQItem = make([]FQItemType, 0)
        }
        t.FQItem = append(t.FQItem, FQItemType{})
        return t
}

func (t *FQItemListType) Last() *FQItemType {
        return &(t.FQItem[len(t.FQItem)-1])
}

  func (t *ResourceUsage_ResourceReportLineList) Add(value ResourceUsage_ResourceReportLine) *ResourceUsage_ResourceReportLineList {
    
        if t == nil {
                t = ResourceUsage_ResourceReportLineListCreate(ResourceUsage_ResourceReportLineList{})
        }
        if t.ResourceReportLine == nil {
                t.ResourceReportLine = make([]ResourceUsage_ResourceReportLine, 0)
        }
        t.ResourceReportLine = append(t.ResourceReportLine, value)
        return t
}

func (t *ResourceUsage_ResourceReportLineList) AddNew() *ResourceUsage_ResourceReportLineList {
        if t == nil {
                t = ResourceUsage_ResourceReportLineListCreate(ResourceUsage_ResourceReportLineList{})
        }
        if t.ResourceReportLine == nil {
                t.ResourceReportLine = make([]ResourceUsage_ResourceReportLine, 0)
        }
        t.ResourceReportLine = append(t.ResourceReportLine, ResourceUsage_ResourceReportLine{})
        return t
}

func (t *ResourceUsage_ResourceReportLineList) Last() *ResourceUsage_ResourceReportLine {
        return &(t.ResourceReportLine[len(t.ResourceReportLine)-1])
}

  func (t *StudentGradeMarkersListType) Add(value MarkerType) *StudentGradeMarkersListType {
    
        if t == nil {
                t = StudentGradeMarkersListTypeCreate(StudentGradeMarkersListType{})
        }
        if t.Marker == nil {
                t.Marker = make([]MarkerType, 0)
        }
        t.Marker = append(t.Marker, value)
        return t
}

func (t *StudentGradeMarkersListType) AddNew() *StudentGradeMarkersListType {
        if t == nil {
                t = StudentGradeMarkersListTypeCreate(StudentGradeMarkersListType{})
        }
        if t.Marker == nil {
                t.Marker = make([]MarkerType, 0)
        }
        t.Marker = append(t.Marker, MarkerType{})
        return t
}

func (t *StudentGradeMarkersListType) Last() *MarkerType {
        return &(t.Marker[len(t.Marker)-1])
}

  func (t *EvaluationsType) Add(value EvaluationType) *EvaluationsType {
    
        if t == nil {
                t = EvaluationsTypeCreate(EvaluationsType{})
        }
        if t.Evaluation == nil {
                t.Evaluation = make([]EvaluationType, 0)
        }
        t.Evaluation = append(t.Evaluation, value)
        return t
}

func (t *EvaluationsType) AddNew() *EvaluationsType {
        if t == nil {
                t = EvaluationsTypeCreate(EvaluationsType{})
        }
        if t.Evaluation == nil {
                t.Evaluation = make([]EvaluationType, 0)
        }
        t.Evaluation = append(t.Evaluation, EvaluationType{})
        return t
}

func (t *EvaluationsType) Last() *EvaluationType {
        return &(t.Evaluation[len(t.Evaluation)-1])
}

  func (t *NAPStudentResponseTestletListType) Add(value NAPTestletResponseType) *NAPStudentResponseTestletListType {
    
        if t == nil {
                t = NAPStudentResponseTestletListTypeCreate(NAPStudentResponseTestletListType{})
        }
        if t.Testlet == nil {
                t.Testlet = make([]NAPTestletResponseType, 0)
        }
        t.Testlet = append(t.Testlet, value)
        return t
}

func (t *NAPStudentResponseTestletListType) AddNew() *NAPStudentResponseTestletListType {
        if t == nil {
                t = NAPStudentResponseTestletListTypeCreate(NAPStudentResponseTestletListType{})
        }
        if t.Testlet == nil {
                t.Testlet = make([]NAPTestletResponseType, 0)
        }
        t.Testlet = append(t.Testlet, NAPTestletResponseType{})
        return t
}

func (t *NAPStudentResponseTestletListType) Last() *NAPTestletResponseType {
        return &(t.Testlet[len(t.Testlet)-1])
}

  func (t *StimulusListType) Add(value StimulusType) *StimulusListType {
    
        if t == nil {
                t = StimulusListTypeCreate(StimulusListType{})
        }
        if t.Stimulus == nil {
                t.Stimulus = make([]StimulusType, 0)
        }
        t.Stimulus = append(t.Stimulus, value)
        return t
}

func (t *StimulusListType) AddNew() *StimulusListType {
        if t == nil {
                t = StimulusListTypeCreate(StimulusListType{})
        }
        if t.Stimulus == nil {
                t.Stimulus = make([]StimulusType, 0)
        }
        t.Stimulus = append(t.Stimulus, StimulusType{})
        return t
}

func (t *StimulusListType) Last() *StimulusType {
        return &(t.Stimulus[len(t.Stimulus)-1])
}

  func (t *AddressCollectionStudentListType) Add(value AddressCollectionStudentType) *AddressCollectionStudentListType {
    
        if t == nil {
                t = AddressCollectionStudentListTypeCreate(AddressCollectionStudentListType{})
        }
        if t.AddressCollectionStudent == nil {
                t.AddressCollectionStudent = make([]AddressCollectionStudentType, 0)
        }
        t.AddressCollectionStudent = append(t.AddressCollectionStudent, value)
        return t
}

func (t *AddressCollectionStudentListType) AddNew() *AddressCollectionStudentListType {
        if t == nil {
                t = AddressCollectionStudentListTypeCreate(AddressCollectionStudentListType{})
        }
        if t.AddressCollectionStudent == nil {
                t.AddressCollectionStudent = make([]AddressCollectionStudentType, 0)
        }
        t.AddressCollectionStudent = append(t.AddressCollectionStudent, AddressCollectionStudentType{})
        return t
}

func (t *AddressCollectionStudentListType) Last() *AddressCollectionStudentType {
        return &(t.AddressCollectionStudent[len(t.AddressCollectionStudent)-1])
}

  func (t *AddressCollectionReportingListType) Add(value AddressCollectionReportingType) *AddressCollectionReportingListType {
    
        if t == nil {
                t = AddressCollectionReportingListTypeCreate(AddressCollectionReportingListType{})
        }
        if t.AddressCollectionReporting == nil {
                t.AddressCollectionReporting = make([]AddressCollectionReportingType, 0)
        }
        t.AddressCollectionReporting = append(t.AddressCollectionReporting, value)
        return t
}

func (t *AddressCollectionReportingListType) AddNew() *AddressCollectionReportingListType {
        if t == nil {
                t = AddressCollectionReportingListTypeCreate(AddressCollectionReportingListType{})
        }
        if t.AddressCollectionReporting == nil {
                t.AddressCollectionReporting = make([]AddressCollectionReportingType, 0)
        }
        t.AddressCollectionReporting = append(t.AddressCollectionReporting, AddressCollectionReportingType{})
        return t
}

func (t *AddressCollectionReportingListType) Last() *AddressCollectionReportingType {
        return &(t.AddressCollectionReporting[len(t.AddressCollectionReporting)-1])
}

  func (t *SystemRole_SystemContextList) Add(value SystemRole_SystemContext) *SystemRole_SystemContextList {
    
        if t == nil {
                t = SystemRole_SystemContextListCreate(SystemRole_SystemContextList{})
        }
        if t.SystemContext == nil {
                t.SystemContext = make([]SystemRole_SystemContext, 0)
        }
        t.SystemContext = append(t.SystemContext, value)
        return t
}

func (t *SystemRole_SystemContextList) AddNew() *SystemRole_SystemContextList {
        if t == nil {
                t = SystemRole_SystemContextListCreate(SystemRole_SystemContextList{})
        }
        if t.SystemContext == nil {
                t.SystemContext = make([]SystemRole_SystemContext, 0)
        }
        t.SystemContext = append(t.SystemContext, SystemRole_SystemContext{})
        return t
}

func (t *SystemRole_SystemContextList) Last() *SystemRole_SystemContext {
        return &(t.SystemContext[len(t.SystemContext)-1])
}

  func (t *NAPTestletItemResponseListType) Add(value NAPTestletResponseItemType) *NAPTestletItemResponseListType {
    
        if t == nil {
                t = NAPTestletItemResponseListTypeCreate(NAPTestletItemResponseListType{})
        }
        if t.ItemResponse == nil {
                t.ItemResponse = make([]NAPTestletResponseItemType, 0)
        }
        t.ItemResponse = append(t.ItemResponse, value)
        return t
}

func (t *NAPTestletItemResponseListType) AddNew() *NAPTestletItemResponseListType {
        if t == nil {
                t = NAPTestletItemResponseListTypeCreate(NAPTestletItemResponseListType{})
        }
        if t.ItemResponse == nil {
                t.ItemResponse = make([]NAPTestletResponseItemType, 0)
        }
        t.ItemResponse = append(t.ItemResponse, NAPTestletResponseItemType{})
        return t
}

func (t *NAPTestletItemResponseListType) Last() *NAPTestletResponseItemType {
        return &(t.ItemResponse[len(t.ItemResponse)-1])
}

  func (t *ComponentsType) Add(value ComponentType) *ComponentsType {
    
        if t == nil {
                t = ComponentsTypeCreate(ComponentsType{})
        }
        if t.Component == nil {
                t.Component = make([]ComponentType, 0)
        }
        t.Component = append(t.Component, value)
        return t
}

func (t *ComponentsType) AddNew() *ComponentsType {
        if t == nil {
                t = ComponentsTypeCreate(ComponentsType{})
        }
        if t.Component == nil {
                t.Component = make([]ComponentType, 0)
        }
        t.Component = append(t.Component, ComponentType{})
        return t
}

func (t *ComponentsType) Last() *ComponentType {
        return &(t.Component[len(t.Component)-1])
}

  func (t *ReligiousEventListType) Add(value ReligiousEventType) *ReligiousEventListType {
    
        if t == nil {
                t = ReligiousEventListTypeCreate(ReligiousEventListType{})
        }
        if t.ReligiousEvent == nil {
                t.ReligiousEvent = make([]ReligiousEventType, 0)
        }
        t.ReligiousEvent = append(t.ReligiousEvent, value)
        return t
}

func (t *ReligiousEventListType) AddNew() *ReligiousEventListType {
        if t == nil {
                t = ReligiousEventListTypeCreate(ReligiousEventListType{})
        }
        if t.ReligiousEvent == nil {
                t.ReligiousEvent = make([]ReligiousEventType, 0)
        }
        t.ReligiousEvent = append(t.ReligiousEvent, ReligiousEventType{})
        return t
}

func (t *ReligiousEventListType) Last() *ReligiousEventType {
        return &(t.ReligiousEvent[len(t.ReligiousEvent)-1])
}

  func (t *ScoreListType) Add(value ScoreType) *ScoreListType {
    
        if t == nil {
                t = ScoreListTypeCreate(ScoreListType{})
        }
        if t.Score == nil {
                t.Score = make([]ScoreType, 0)
        }
        t.Score = append(t.Score, value)
        return t
}

func (t *ScoreListType) AddNew() *ScoreListType {
        if t == nil {
                t = ScoreListTypeCreate(ScoreListType{})
        }
        if t.Score == nil {
                t.Score = make([]ScoreType, 0)
        }
        t.Score = append(t.Score, ScoreType{})
        return t
}

func (t *ScoreListType) Last() *ScoreType {
        return &(t.Score[len(t.Score)-1])
}

  func (t *YearLevelEnrollmentListType) Add(value YearLevelEnrollmentType) *YearLevelEnrollmentListType {
    
        if t == nil {
                t = YearLevelEnrollmentListTypeCreate(YearLevelEnrollmentListType{})
        }
        if t.YearLevelEnrollment == nil {
                t.YearLevelEnrollment = make([]YearLevelEnrollmentType, 0)
        }
        t.YearLevelEnrollment = append(t.YearLevelEnrollment, value)
        return t
}

func (t *YearLevelEnrollmentListType) AddNew() *YearLevelEnrollmentListType {
        if t == nil {
                t = YearLevelEnrollmentListTypeCreate(YearLevelEnrollmentListType{})
        }
        if t.YearLevelEnrollment == nil {
                t.YearLevelEnrollment = make([]YearLevelEnrollmentType, 0)
        }
        t.YearLevelEnrollment = append(t.YearLevelEnrollment, YearLevelEnrollmentType{})
        return t
}

func (t *YearLevelEnrollmentListType) Last() *YearLevelEnrollmentType {
        return &(t.YearLevelEnrollment[len(t.YearLevelEnrollment)-1])
}

  func (t *ContactsType) Add(value ContactType) *ContactsType {
    
        if t == nil {
                t = ContactsTypeCreate(ContactsType{})
        }
        if t.Contact == nil {
                t.Contact = make([]ContactType, 0)
        }
        t.Contact = append(t.Contact, value)
        return t
}

func (t *ContactsType) AddNew() *ContactsType {
        if t == nil {
                t = ContactsTypeCreate(ContactsType{})
        }
        if t.Contact == nil {
                t.Contact = make([]ContactType, 0)
        }
        t.Contact = append(t.Contact, ContactType{})
        return t
}

func (t *ContactsType) Last() *ContactType {
        return &(t.Contact[len(t.Contact)-1])
}

  func (t *PNPCodeListType) Add(value AUCodeSetsPNPCodeType) *PNPCodeListType {
      present := false
  for _, b := range AUCodeSetsPNPCodeType_values {
        if b == string(value) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPNPCodeType_values")
      }

        if t == nil {
                t = PNPCodeListTypeCreate(PNPCodeListType{})
        }
        if t.PNPCode == nil {
                t.PNPCode = make([]AUCodeSetsPNPCodeType, 0)
        }
        t.PNPCode = append(t.PNPCode, value)
        return t
}

func (t *PNPCodeListType) AddNew() *PNPCodeListType {
        if t == nil {
                t = PNPCodeListTypeCreate(PNPCodeListType{})
        }
        if t.PNPCode == nil {
                t.PNPCode = make([]AUCodeSetsPNPCodeType, 0)
        }
        t.PNPCode = append(t.PNPCode, "")
        return t
}

func (t *PNPCodeListType) Last() *AUCodeSetsPNPCodeType {
        return &(t.PNPCode[len(t.PNPCode)-1])
}

  func (t *PlanRequiredListType) Add(value WellbeingPlanType) *PlanRequiredListType {
    
        if t == nil {
                t = PlanRequiredListTypeCreate(PlanRequiredListType{})
        }
        if t.Plan == nil {
                t.Plan = make([]WellbeingPlanType, 0)
        }
        t.Plan = append(t.Plan, value)
        return t
}

func (t *PlanRequiredListType) AddNew() *PlanRequiredListType {
        if t == nil {
                t = PlanRequiredListTypeCreate(PlanRequiredListType{})
        }
        if t.Plan == nil {
                t.Plan = make([]WellbeingPlanType, 0)
        }
        t.Plan = append(t.Plan, WellbeingPlanType{})
        return t
}

func (t *PlanRequiredListType) Last() *WellbeingPlanType {
        return &(t.Plan[len(t.Plan)-1])
}

  func (t *NAPSubscoreListType) Add(value NAPSubscoreType) *NAPSubscoreListType {
    
        if t == nil {
                t = NAPSubscoreListTypeCreate(NAPSubscoreListType{})
        }
        if t.Subscore == nil {
                t.Subscore = make([]NAPSubscoreType, 0)
        }
        t.Subscore = append(t.Subscore, value)
        return t
}

func (t *NAPSubscoreListType) AddNew() *NAPSubscoreListType {
        if t == nil {
                t = NAPSubscoreListTypeCreate(NAPSubscoreListType{})
        }
        if t.Subscore == nil {
                t.Subscore = make([]NAPSubscoreType, 0)
        }
        t.Subscore = append(t.Subscore, NAPSubscoreType{})
        return t
}

func (t *NAPSubscoreListType) Last() *NAPSubscoreType {
        return &(t.Subscore[len(t.Subscore)-1])
}

  func (t *HouseholdContactInfoListType) Add(value HouseholdContactInfoType) *HouseholdContactInfoListType {
    
        if t == nil {
                t = HouseholdContactInfoListTypeCreate(HouseholdContactInfoListType{})
        }
        if t.HouseholdContactInfo == nil {
                t.HouseholdContactInfo = make([]HouseholdContactInfoType, 0)
        }
        t.HouseholdContactInfo = append(t.HouseholdContactInfo, value)
        return t
}

func (t *HouseholdContactInfoListType) AddNew() *HouseholdContactInfoListType {
        if t == nil {
                t = HouseholdContactInfoListTypeCreate(HouseholdContactInfoListType{})
        }
        if t.HouseholdContactInfo == nil {
                t.HouseholdContactInfo = make([]HouseholdContactInfoType, 0)
        }
        t.HouseholdContactInfo = append(t.HouseholdContactInfo, HouseholdContactInfoType{})
        return t
}

func (t *HouseholdContactInfoListType) Last() *HouseholdContactInfoType {
        return &(t.HouseholdContactInfo[len(t.HouseholdContactInfo)-1])
}

  func (t *SystemRole_RoleScopeList) Add(value SystemRole_RoleScope) *SystemRole_RoleScopeList {
    
        if t == nil {
                t = SystemRole_RoleScopeListCreate(SystemRole_RoleScopeList{})
        }
        if t.RoleScope == nil {
                t.RoleScope = make([]SystemRole_RoleScope, 0)
        }
        t.RoleScope = append(t.RoleScope, value)
        return t
}

func (t *SystemRole_RoleScopeList) AddNew() *SystemRole_RoleScopeList {
        if t == nil {
                t = SystemRole_RoleScopeListCreate(SystemRole_RoleScopeList{})
        }
        if t.RoleScope == nil {
                t.RoleScope = make([]SystemRole_RoleScope, 0)
        }
        t.RoleScope = append(t.RoleScope, SystemRole_RoleScope{})
        return t
}

func (t *SystemRole_RoleScopeList) Last() *SystemRole_RoleScope {
        return &(t.RoleScope[len(t.RoleScope)-1])
}

  func (t *PurchasingItemsType) Add(value PurchasingItemType) *PurchasingItemsType {
    
        if t == nil {
                t = PurchasingItemsTypeCreate(PurchasingItemsType{})
        }
        if t.PurchasingItem == nil {
                t.PurchasingItem = make([]PurchasingItemType, 0)
        }
        t.PurchasingItem = append(t.PurchasingItem, value)
        return t
}

func (t *PurchasingItemsType) AddNew() *PurchasingItemsType {
        if t == nil {
                t = PurchasingItemsTypeCreate(PurchasingItemsType{})
        }
        if t.PurchasingItem == nil {
                t.PurchasingItem = make([]PurchasingItemType, 0)
        }
        t.PurchasingItem = append(t.PurchasingItem, PurchasingItemType{})
        return t
}

func (t *PurchasingItemsType) Last() *PurchasingItemType {
        return &(t.PurchasingItem[len(t.PurchasingItem)-1])
}

  func (t *TimeTableDayListType) Add(value TimeTableDayType) *TimeTableDayListType {
    
        if t == nil {
                t = TimeTableDayListTypeCreate(TimeTableDayListType{})
        }
        if t.TimeTableDay == nil {
                t.TimeTableDay = make([]TimeTableDayType, 0)
        }
        t.TimeTableDay = append(t.TimeTableDay, value)
        return t
}

func (t *TimeTableDayListType) AddNew() *TimeTableDayListType {
        if t == nil {
                t = TimeTableDayListTypeCreate(TimeTableDayListType{})
        }
        if t.TimeTableDay == nil {
                t.TimeTableDay = make([]TimeTableDayType, 0)
        }
        t.TimeTableDay = append(t.TimeTableDay, TimeTableDayType{})
        return t
}

func (t *TimeTableDayListType) Last() *TimeTableDayType {
        return &(t.TimeTableDay[len(t.TimeTableDay)-1])
}

  func (t *RecognitionListType) Add(value string) *RecognitionListType {
    
        if t == nil {
                t = RecognitionListTypeCreate(RecognitionListType{})
        }
        if t.Recognition == nil {
                t.Recognition = make([]string, 0)
        }
        t.Recognition = append(t.Recognition, value)
        return t
}

func (t *RecognitionListType) AddNew() *RecognitionListType {
        if t == nil {
                t = RecognitionListTypeCreate(RecognitionListType{})
        }
        if t.Recognition == nil {
                t.Recognition = make([]string, 0)
        }
        t.Recognition = append(t.Recognition, "")
        return t
}

func (t *RecognitionListType) Last() *string {
        return &(t.Recognition[len(t.Recognition)-1])
}

  func (t *TimeElementType_SpanGaps) Add(value TimeElementType_SpanGap) *TimeElementType_SpanGaps {
    
        if t == nil {
                t = TimeElementType_SpanGapsCreate(TimeElementType_SpanGaps{})
        }
        if t.SpanGap == nil {
                t.SpanGap = make([]TimeElementType_SpanGap, 0)
        }
        t.SpanGap = append(t.SpanGap, value)
        return t
}

func (t *TimeElementType_SpanGaps) AddNew() *TimeElementType_SpanGaps {
        if t == nil {
                t = TimeElementType_SpanGapsCreate(TimeElementType_SpanGaps{})
        }
        if t.SpanGap == nil {
                t.SpanGap = make([]TimeElementType_SpanGap, 0)
        }
        t.SpanGap = append(t.SpanGap, TimeElementType_SpanGap{})
        return t
}

func (t *TimeElementType_SpanGaps) Last() *TimeElementType_SpanGap {
        return &(t.SpanGap[len(t.SpanGap)-1])
}

  func (t *PlausibleScaledValueListType) Add(value float64) *PlausibleScaledValueListType {
    
        if t == nil {
                t = PlausibleScaledValueListTypeCreate(PlausibleScaledValueListType{})
        }
        if t.PlausibleScaledValue == nil {
                t.PlausibleScaledValue = make([]float64, 0)
        }
        t.PlausibleScaledValue = append(t.PlausibleScaledValue, value)
        return t
}

func (t *PlausibleScaledValueListType) AddNew() *PlausibleScaledValueListType {
        if t == nil {
                t = PlausibleScaledValueListTypeCreate(PlausibleScaledValueListType{})
        }
        if t.PlausibleScaledValue == nil {
                t.PlausibleScaledValue = make([]float64, 0)
        }
        t.PlausibleScaledValue = append(t.PlausibleScaledValue, 0)
        return t
}

func (t *PlausibleScaledValueListType) Last() *float64 {
        return &(t.PlausibleScaledValue[len(t.PlausibleScaledValue)-1])
}

  func (t *LearningStandardsDocumentType) Add(value string) *LearningStandardsDocumentType {
    
        if t == nil {
                t = LearningStandardsDocumentTypeCreate(LearningStandardsDocumentType{})
        }
        if t.LearningStandardDocumentRefId == nil {
                t.LearningStandardDocumentRefId = make([]string, 0)
        }
        t.LearningStandardDocumentRefId = append(t.LearningStandardDocumentRefId, value)
        return t
}

func (t *LearningStandardsDocumentType) AddNew() *LearningStandardsDocumentType {
        if t == nil {
                t = LearningStandardsDocumentTypeCreate(LearningStandardsDocumentType{})
        }
        if t.LearningStandardDocumentRefId == nil {
                t.LearningStandardDocumentRefId = make([]string, 0)
        }
        t.LearningStandardDocumentRefId = append(t.LearningStandardDocumentRefId, "")
        return t
}

func (t *LearningStandardsDocumentType) Last() *string {
        return &(t.LearningStandardDocumentRefId[len(t.LearningStandardDocumentRefId)-1])
}

  func (t *SymptomListType) Add(value string) *SymptomListType {
    
        if t == nil {
                t = SymptomListTypeCreate(SymptomListType{})
        }
        if t.Symptom == nil {
                t.Symptom = make([]string, 0)
        }
        t.Symptom = append(t.Symptom, value)
        return t
}

func (t *SymptomListType) AddNew() *SymptomListType {
        if t == nil {
                t = SymptomListTypeCreate(SymptomListType{})
        }
        if t.Symptom == nil {
                t.Symptom = make([]string, 0)
        }
        t.Symptom = append(t.Symptom, "")
        return t
}

func (t *SymptomListType) Last() *string {
        return &(t.Symptom[len(t.Symptom)-1])
}

  func (t *SchoolContactListType) Add(value SchoolContactType) *SchoolContactListType {
    
        if t == nil {
                t = SchoolContactListTypeCreate(SchoolContactListType{})
        }
        if t.SchoolContact == nil {
                t.SchoolContact = make([]SchoolContactType, 0)
        }
        t.SchoolContact = append(t.SchoolContact, value)
        return t
}

func (t *SchoolContactListType) AddNew() *SchoolContactListType {
        if t == nil {
                t = SchoolContactListTypeCreate(SchoolContactListType{})
        }
        if t.SchoolContact == nil {
                t.SchoolContact = make([]SchoolContactType, 0)
        }
        t.SchoolContact = append(t.SchoolContact, SchoolContactType{})
        return t
}

func (t *SchoolContactListType) Last() *SchoolContactType {
        return &(t.SchoolContact[len(t.SchoolContact)-1])
}

  func (t *StrategiesType) Add(value string) *StrategiesType {
    
        if t == nil {
                t = StrategiesTypeCreate(StrategiesType{})
        }
        if t.Strategy == nil {
                t.Strategy = make([]string, 0)
        }
        t.Strategy = append(t.Strategy, value)
        return t
}

func (t *StrategiesType) AddNew() *StrategiesType {
        if t == nil {
                t = StrategiesTypeCreate(StrategiesType{})
        }
        if t.Strategy == nil {
                t.Strategy = make([]string, 0)
        }
        t.Strategy = append(t.Strategy, "")
        return t
}

func (t *StrategiesType) Last() *string {
        return &(t.Strategy[len(t.Strategy)-1])
}

  func (t *HouseholdListType) Add(value LocalIdType) *HouseholdListType {
    
        if t == nil {
                t = HouseholdListTypeCreate(HouseholdListType{})
        }
        if t.Household == nil {
                t.Household = make([]LocalIdType, 0)
        }
        t.Household = append(t.Household, value)
        return t
}

func (t *HouseholdListType) AddNew() *HouseholdListType {
        if t == nil {
                t = HouseholdListTypeCreate(HouseholdListType{})
        }
        if t.Household == nil {
                t.Household = make([]LocalIdType, 0)
        }
        t.Household = append(t.Household, "")
        return t
}

func (t *HouseholdListType) Last() *LocalIdType {
        return &(t.Household[len(t.Household)-1])
}

  func (t *WellbeingEventSubCategoryListType) Add(value string) *WellbeingEventSubCategoryListType {
    
        if t == nil {
                t = WellbeingEventSubCategoryListTypeCreate(WellbeingEventSubCategoryListType{})
        }
        if t.WellbeingEventSubCategory == nil {
                t.WellbeingEventSubCategory = make([]string, 0)
        }
        t.WellbeingEventSubCategory = append(t.WellbeingEventSubCategory, value)
        return t
}

func (t *WellbeingEventSubCategoryListType) AddNew() *WellbeingEventSubCategoryListType {
        if t == nil {
                t = WellbeingEventSubCategoryListTypeCreate(WellbeingEventSubCategoryListType{})
        }
        if t.WellbeingEventSubCategory == nil {
                t.WellbeingEventSubCategory = make([]string, 0)
        }
        t.WellbeingEventSubCategory = append(t.WellbeingEventSubCategory, "")
        return t
}

func (t *WellbeingEventSubCategoryListType) Last() *string {
        return &(t.WellbeingEventSubCategory[len(t.WellbeingEventSubCategory)-1])
}

  func (t *ScheduledTeacherListType) Add(value TeacherCoverType) *ScheduledTeacherListType {
    
        if t == nil {
                t = ScheduledTeacherListTypeCreate(ScheduledTeacherListType{})
        }
        if t.TeacherCover == nil {
                t.TeacherCover = make([]TeacherCoverType, 0)
        }
        t.TeacherCover = append(t.TeacherCover, value)
        return t
}

func (t *ScheduledTeacherListType) AddNew() *ScheduledTeacherListType {
        if t == nil {
                t = ScheduledTeacherListTypeCreate(ScheduledTeacherListType{})
        }
        if t.TeacherCover == nil {
                t.TeacherCover = make([]TeacherCoverType, 0)
        }
        t.TeacherCover = append(t.TeacherCover, TeacherCoverType{})
        return t
}

func (t *ScheduledTeacherListType) Last() *TeacherCoverType {
        return &(t.TeacherCover[len(t.TeacherCover)-1])
}

  func (t *StaffListType) Add(value string) *StaffListType {
    
        if t == nil {
                t = StaffListTypeCreate(StaffListType{})
        }
        if t.StaffPersonalRefId == nil {
                t.StaffPersonalRefId = make([]string, 0)
        }
        t.StaffPersonalRefId = append(t.StaffPersonalRefId, value)
        return t
}

func (t *StaffListType) AddNew() *StaffListType {
        if t == nil {
                t = StaffListTypeCreate(StaffListType{})
        }
        if t.StaffPersonalRefId == nil {
                t.StaffPersonalRefId = make([]string, 0)
        }
        t.StaffPersonalRefId = append(t.StaffPersonalRefId, "")
        return t
}

func (t *StaffListType) Last() *string {
        return &(t.StaffPersonalRefId[len(t.StaffPersonalRefId)-1])
}

  func (t *FQContextualQuestionListType) Add(value FQContextualQuestionType) *FQContextualQuestionListType {
    
        if t == nil {
                t = FQContextualQuestionListTypeCreate(FQContextualQuestionListType{})
        }
        if t.FQContextualQuestion == nil {
                t.FQContextualQuestion = make([]FQContextualQuestionType, 0)
        }
        t.FQContextualQuestion = append(t.FQContextualQuestion, value)
        return t
}

func (t *FQContextualQuestionListType) AddNew() *FQContextualQuestionListType {
        if t == nil {
                t = FQContextualQuestionListTypeCreate(FQContextualQuestionListType{})
        }
        if t.FQContextualQuestion == nil {
                t.FQContextualQuestion = make([]FQContextualQuestionType, 0)
        }
        t.FQContextualQuestion = append(t.FQContextualQuestion, FQContextualQuestionType{})
        return t
}

func (t *FQContextualQuestionListType) Last() *FQContextualQuestionType {
        return &(t.FQContextualQuestion[len(t.FQContextualQuestion)-1])
}

  func (t *LearningResourcesType) Add(value string) *LearningResourcesType {
    
        if t == nil {
                t = LearningResourcesTypeCreate(LearningResourcesType{})
        }
        if t.LearningResourceRefId == nil {
                t.LearningResourceRefId = make([]string, 0)
        }
        t.LearningResourceRefId = append(t.LearningResourceRefId, value)
        return t
}

func (t *LearningResourcesType) AddNew() *LearningResourcesType {
        if t == nil {
                t = LearningResourcesTypeCreate(LearningResourcesType{})
        }
        if t.LearningResourceRefId == nil {
                t.LearningResourceRefId = make([]string, 0)
        }
        t.LearningResourceRefId = append(t.LearningResourceRefId, "")
        return t
}

func (t *LearningResourcesType) Last() *string {
        return &(t.LearningResourceRefId[len(t.LearningResourceRefId)-1])
}

  func (t *LearningObjectivesType) Add(value string) *LearningObjectivesType {
    
        if t == nil {
                t = LearningObjectivesTypeCreate(LearningObjectivesType{})
        }
        if t.LearningObjective == nil {
                t.LearningObjective = make([]string, 0)
        }
        t.LearningObjective = append(t.LearningObjective, value)
        return t
}

func (t *LearningObjectivesType) AddNew() *LearningObjectivesType {
        if t == nil {
                t = LearningObjectivesTypeCreate(LearningObjectivesType{})
        }
        if t.LearningObjective == nil {
                t.LearningObjective = make([]string, 0)
        }
        t.LearningObjective = append(t.LearningObjective, "")
        return t
}

func (t *LearningObjectivesType) Last() *string {
        return &(t.LearningObjective[len(t.LearningObjective)-1])
}

  func (t *NAPLANClassListType) Add(value string) *NAPLANClassListType {
    
        if t == nil {
                t = NAPLANClassListTypeCreate(NAPLANClassListType{})
        }
        if t.ClassCode == nil {
                t.ClassCode = make([]string, 0)
        }
        t.ClassCode = append(t.ClassCode, value)
        return t
}

func (t *NAPLANClassListType) AddNew() *NAPLANClassListType {
        if t == nil {
                t = NAPLANClassListTypeCreate(NAPLANClassListType{})
        }
        if t.ClassCode == nil {
                t.ClassCode = make([]string, 0)
        }
        t.ClassCode = append(t.ClassCode, "")
        return t
}

func (t *NAPLANClassListType) Last() *string {
        return &(t.ClassCode[len(t.ClassCode)-1])
}

  func (t *SubjectAreaListType) Add(value SubjectAreaType) *SubjectAreaListType {
    
        if t == nil {
                t = SubjectAreaListTypeCreate(SubjectAreaListType{})
        }
        if t.SubjectArea == nil {
                t.SubjectArea = make([]SubjectAreaType, 0)
        }
        t.SubjectArea = append(t.SubjectArea, value)
        return t
}

func (t *SubjectAreaListType) AddNew() *SubjectAreaListType {
        if t == nil {
                t = SubjectAreaListTypeCreate(SubjectAreaListType{})
        }
        if t.SubjectArea == nil {
                t.SubjectArea = make([]SubjectAreaType, 0)
        }
        t.SubjectArea = append(t.SubjectArea, SubjectAreaType{})
        return t
}

func (t *SubjectAreaListType) Last() *SubjectAreaType {
        return &(t.SubjectArea[len(t.SubjectArea)-1])
}

  func (t *PersonInvolvementListType) Add(value PersonInvolvementType) *PersonInvolvementListType {
    
        if t == nil {
                t = PersonInvolvementListTypeCreate(PersonInvolvementListType{})
        }
        if t.PersonInvolvement == nil {
                t.PersonInvolvement = make([]PersonInvolvementType, 0)
        }
        t.PersonInvolvement = append(t.PersonInvolvement, value)
        return t
}

func (t *PersonInvolvementListType) AddNew() *PersonInvolvementListType {
        if t == nil {
                t = PersonInvolvementListTypeCreate(PersonInvolvementListType{})
        }
        if t.PersonInvolvement == nil {
                t.PersonInvolvement = make([]PersonInvolvementType, 0)
        }
        t.PersonInvolvement = append(t.PersonInvolvement, PersonInvolvementType{})
        return t
}

func (t *PersonInvolvementListType) Last() *PersonInvolvementType {
        return &(t.PersonInvolvement[len(t.PersonInvolvement)-1])
}

  func (t *StatsCohortYearLevelListType) Add(value StatsCohortYearLevelType) *StatsCohortYearLevelListType {
    
        if t == nil {
                t = StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
        }
        if t.StatsCohortYearLevel == nil {
                t.StatsCohortYearLevel = make([]StatsCohortYearLevelType, 0)
        }
        t.StatsCohortYearLevel = append(t.StatsCohortYearLevel, value)
        return t
}

func (t *StatsCohortYearLevelListType) AddNew() *StatsCohortYearLevelListType {
        if t == nil {
                t = StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
        }
        if t.StatsCohortYearLevel == nil {
                t.StatsCohortYearLevel = make([]StatsCohortYearLevelType, 0)
        }
        t.StatsCohortYearLevel = append(t.StatsCohortYearLevel, StatsCohortYearLevelType{})
        return t
}

func (t *StatsCohortYearLevelListType) Last() *StatsCohortYearLevelType {
        return &(t.StatsCohortYearLevel[len(t.StatsCohortYearLevel)-1])
}

  func (t *ProgramFundingSourcesType) Add(value ProgramFundingSourceType) *ProgramFundingSourcesType {
    
        if t == nil {
                t = ProgramFundingSourcesTypeCreate(ProgramFundingSourcesType{})
        }
        if t.ProgramFundingSource == nil {
                t.ProgramFundingSource = make([]ProgramFundingSourceType, 0)
        }
        t.ProgramFundingSource = append(t.ProgramFundingSource, value)
        return t
}

func (t *ProgramFundingSourcesType) AddNew() *ProgramFundingSourcesType {
        if t == nil {
                t = ProgramFundingSourcesTypeCreate(ProgramFundingSourcesType{})
        }
        if t.ProgramFundingSource == nil {
                t.ProgramFundingSource = make([]ProgramFundingSourceType, 0)
        }
        t.ProgramFundingSource = append(t.ProgramFundingSource, ProgramFundingSourceType{})
        return t
}

func (t *ProgramFundingSourcesType) Last() *ProgramFundingSourceType {
        return &(t.ProgramFundingSource[len(t.ProgramFundingSource)-1])
}

  func (t *AlternateIdentificationCodeListType) Add(value string) *AlternateIdentificationCodeListType {
    
        if t == nil {
                t = AlternateIdentificationCodeListTypeCreate(AlternateIdentificationCodeListType{})
        }
        if t.AlternateIdentificationCode == nil {
                t.AlternateIdentificationCode = make([]string, 0)
        }
        t.AlternateIdentificationCode = append(t.AlternateIdentificationCode, value)
        return t
}

func (t *AlternateIdentificationCodeListType) AddNew() *AlternateIdentificationCodeListType {
        if t == nil {
                t = AlternateIdentificationCodeListTypeCreate(AlternateIdentificationCodeListType{})
        }
        if t.AlternateIdentificationCode == nil {
                t.AlternateIdentificationCode = make([]string, 0)
        }
        t.AlternateIdentificationCode = append(t.AlternateIdentificationCode, "")
        return t
}

func (t *AlternateIdentificationCodeListType) Last() *string {
        return &(t.AlternateIdentificationCode[len(t.AlternateIdentificationCode)-1])
}

  func (t *AuthorsType) Add(value string) *AuthorsType {
    
        if t == nil {
                t = AuthorsTypeCreate(AuthorsType{})
        }
        if t.Author == nil {
                t.Author = make([]string, 0)
        }
        t.Author = append(t.Author, value)
        return t
}

func (t *AuthorsType) AddNew() *AuthorsType {
        if t == nil {
                t = AuthorsTypeCreate(AuthorsType{})
        }
        if t.Author == nil {
                t.Author = make([]string, 0)
        }
        t.Author = append(t.Author, "")
        return t
}

func (t *AuthorsType) Last() *string {
        return &(t.Author[len(t.Author)-1])
}

  func (t *NAPCodeFrameTestletListType) Add(value NAPTestletCodeFrameType) *NAPCodeFrameTestletListType {
    
        if t == nil {
                t = NAPCodeFrameTestletListTypeCreate(NAPCodeFrameTestletListType{})
        }
        if t.Testlet == nil {
                t.Testlet = make([]NAPTestletCodeFrameType, 0)
        }
        t.Testlet = append(t.Testlet, value)
        return t
}

func (t *NAPCodeFrameTestletListType) AddNew() *NAPCodeFrameTestletListType {
        if t == nil {
                t = NAPCodeFrameTestletListTypeCreate(NAPCodeFrameTestletListType{})
        }
        if t.Testlet == nil {
                t.Testlet = make([]NAPTestletCodeFrameType, 0)
        }
        t.Testlet = append(t.Testlet, NAPTestletCodeFrameType{})
        return t
}

func (t *NAPCodeFrameTestletListType) Last() *NAPTestletCodeFrameType {
        return &(t.Testlet[len(t.Testlet)-1])
}

  func (t *PrerequisitesType) Add(value string) *PrerequisitesType {
    
        if t == nil {
                t = PrerequisitesTypeCreate(PrerequisitesType{})
        }
        if t.Prerequisite == nil {
                t.Prerequisite = make([]string, 0)
        }
        t.Prerequisite = append(t.Prerequisite, value)
        return t
}

func (t *PrerequisitesType) AddNew() *PrerequisitesType {
        if t == nil {
                t = PrerequisitesTypeCreate(PrerequisitesType{})
        }
        if t.Prerequisite == nil {
                t.Prerequisite = make([]string, 0)
        }
        t.Prerequisite = append(t.Prerequisite, "")
        return t
}

func (t *PrerequisitesType) Last() *string {
        return &(t.Prerequisite[len(t.Prerequisite)-1])
}

  func (t *WithdrawalTimeListType) Add(value WithdrawalType) *WithdrawalTimeListType {
    
        if t == nil {
                t = WithdrawalTimeListTypeCreate(WithdrawalTimeListType{})
        }
        if t.Withdrawal == nil {
                t.Withdrawal = make([]WithdrawalType, 0)
        }
        t.Withdrawal = append(t.Withdrawal, value)
        return t
}

func (t *WithdrawalTimeListType) AddNew() *WithdrawalTimeListType {
        if t == nil {
                t = WithdrawalTimeListTypeCreate(WithdrawalTimeListType{})
        }
        if t.Withdrawal == nil {
                t.Withdrawal = make([]WithdrawalType, 0)
        }
        t.Withdrawal = append(t.Withdrawal, WithdrawalType{})
        return t
}

func (t *WithdrawalTimeListType) Last() *WithdrawalType {
        return &(t.Withdrawal[len(t.Withdrawal)-1])
}

  func (t *AssociatedObjectsType) Add(value AssociatedObjectsType_AssociatedObject) *AssociatedObjectsType {
    
        if t == nil {
                t = AssociatedObjectsTypeCreate(AssociatedObjectsType{})
        }
        if t.AssociatedObject == nil {
                t.AssociatedObject = make([]AssociatedObjectsType_AssociatedObject, 0)
        }
        t.AssociatedObject = append(t.AssociatedObject, value)
        return t
}

func (t *AssociatedObjectsType) AddNew() *AssociatedObjectsType {
        if t == nil {
                t = AssociatedObjectsTypeCreate(AssociatedObjectsType{})
        }
        if t.AssociatedObject == nil {
                t.AssociatedObject = make([]AssociatedObjectsType_AssociatedObject, 0)
        }
        t.AssociatedObject = append(t.AssociatedObject, AssociatedObjectsType_AssociatedObject{})
        return t
}

func (t *AssociatedObjectsType) Last() *AssociatedObjectsType_AssociatedObject {
        return &(t.AssociatedObject[len(t.AssociatedObject)-1])
}

  func (t *ACStrandAreaListType) Add(value ACStrandSubjectAreaType) *ACStrandAreaListType {
    
        if t == nil {
                t = ACStrandAreaListTypeCreate(ACStrandAreaListType{})
        }
        if t.ACStrandSubjectArea == nil {
                t.ACStrandSubjectArea = make([]ACStrandSubjectAreaType, 0)
        }
        t.ACStrandSubjectArea = append(t.ACStrandSubjectArea, value)
        return t
}

func (t *ACStrandAreaListType) AddNew() *ACStrandAreaListType {
        if t == nil {
                t = ACStrandAreaListTypeCreate(ACStrandAreaListType{})
        }
        if t.ACStrandSubjectArea == nil {
                t.ACStrandSubjectArea = make([]ACStrandSubjectAreaType, 0)
        }
        t.ACStrandSubjectArea = append(t.ACStrandSubjectArea, ACStrandSubjectAreaType{})
        return t
}

func (t *ACStrandAreaListType) Last() *ACStrandSubjectAreaType {
        return &(t.ACStrandSubjectArea[len(t.ACStrandSubjectArea)-1])
}

  func (t *SchoolGroupListType) Add(value LocalIdType) *SchoolGroupListType {
    
        if t == nil {
                t = SchoolGroupListTypeCreate(SchoolGroupListType{})
        }
        if t.SchoolGroup == nil {
                t.SchoolGroup = make([]LocalIdType, 0)
        }
        t.SchoolGroup = append(t.SchoolGroup, value)
        return t
}

func (t *SchoolGroupListType) AddNew() *SchoolGroupListType {
        if t == nil {
                t = SchoolGroupListTypeCreate(SchoolGroupListType{})
        }
        if t.SchoolGroup == nil {
                t.SchoolGroup = make([]LocalIdType, 0)
        }
        t.SchoolGroup = append(t.SchoolGroup, "")
        return t
}

func (t *SchoolGroupListType) Last() *LocalIdType {
        return &(t.SchoolGroup[len(t.SchoolGroup)-1])
}

  func (t *RelatedLearningStandardItemRefIdListType) Add(value RelatedLearningStandardItemRefIdType) *RelatedLearningStandardItemRefIdListType {
    
        if t == nil {
                t = RelatedLearningStandardItemRefIdListTypeCreate(RelatedLearningStandardItemRefIdListType{})
        }
        if t.LearningStandardItemRefId == nil {
                t.LearningStandardItemRefId = make([]RelatedLearningStandardItemRefIdType, 0)
        }
        t.LearningStandardItemRefId = append(t.LearningStandardItemRefId, value)
        return t
}

func (t *RelatedLearningStandardItemRefIdListType) AddNew() *RelatedLearningStandardItemRefIdListType {
        if t == nil {
                t = RelatedLearningStandardItemRefIdListTypeCreate(RelatedLearningStandardItemRefIdListType{})
        }
        if t.LearningStandardItemRefId == nil {
                t.LearningStandardItemRefId = make([]RelatedLearningStandardItemRefIdType, 0)
        }
        t.LearningStandardItemRefId = append(t.LearningStandardItemRefId, RelatedLearningStandardItemRefIdType{})
        return t
}

func (t *RelatedLearningStandardItemRefIdListType) Last() *RelatedLearningStandardItemRefIdType {
        return &(t.LearningStandardItemRefId[len(t.LearningStandardItemRefId)-1])
}

  func (t *TeacherListType) Add(value TeachingGroupTeacherType) *TeacherListType {
    
        if t == nil {
                t = TeacherListTypeCreate(TeacherListType{})
        }
        if t.TeachingGroupTeacher == nil {
                t.TeachingGroupTeacher = make([]TeachingGroupTeacherType, 0)
        }
        t.TeachingGroupTeacher = append(t.TeachingGroupTeacher, value)
        return t
}

func (t *TeacherListType) AddNew() *TeacherListType {
        if t == nil {
                t = TeacherListTypeCreate(TeacherListType{})
        }
        if t.TeachingGroupTeacher == nil {
                t.TeachingGroupTeacher = make([]TeachingGroupTeacherType, 0)
        }
        t.TeachingGroupTeacher = append(t.TeachingGroupTeacher, TeachingGroupTeacherType{})
        return t
}

func (t *TeacherListType) Last() *TeachingGroupTeacherType {
        return &(t.TeachingGroupTeacher[len(t.TeachingGroupTeacher)-1])
}

  func (t *TimeTableScheduleCellListType) Add(value TimeTableScheduleCellType) *TimeTableScheduleCellListType {
    
        if t == nil {
                t = TimeTableScheduleCellListTypeCreate(TimeTableScheduleCellListType{})
        }
        if t.TimeTableScheduleCell == nil {
                t.TimeTableScheduleCell = make([]TimeTableScheduleCellType, 0)
        }
        t.TimeTableScheduleCell = append(t.TimeTableScheduleCell, value)
        return t
}

func (t *TimeTableScheduleCellListType) AddNew() *TimeTableScheduleCellListType {
        if t == nil {
                t = TimeTableScheduleCellListTypeCreate(TimeTableScheduleCellListType{})
        }
        if t.TimeTableScheduleCell == nil {
                t.TimeTableScheduleCell = make([]TimeTableScheduleCellType, 0)
        }
        t.TimeTableScheduleCell = append(t.TimeTableScheduleCell, TimeTableScheduleCellType{})
        return t
}

func (t *TimeTableScheduleCellListType) Last() *TimeTableScheduleCellType {
        return &(t.TimeTableScheduleCell[len(t.TimeTableScheduleCell)-1])
}

  func (t *GradingScoreListType) Add(value AssignmentScoreType) *GradingScoreListType {
    
        if t == nil {
                t = GradingScoreListTypeCreate(GradingScoreListType{})
        }
        if t.GradingAssignmentScore == nil {
                t.GradingAssignmentScore = make([]AssignmentScoreType, 0)
        }
        t.GradingAssignmentScore = append(t.GradingAssignmentScore, value)
        return t
}

func (t *GradingScoreListType) AddNew() *GradingScoreListType {
        if t == nil {
                t = GradingScoreListTypeCreate(GradingScoreListType{})
        }
        if t.GradingAssignmentScore == nil {
                t.GradingAssignmentScore = make([]AssignmentScoreType, 0)
        }
        t.GradingAssignmentScore = append(t.GradingAssignmentScore, AssignmentScoreType{})
        return t
}

func (t *GradingScoreListType) Last() *AssignmentScoreType {
        return &(t.GradingAssignmentScore[len(t.GradingAssignmentScore)-1])
}

  func (t *TeachingGroupScheduleListType) Add(value TeachingGroupScheduleType) *TeachingGroupScheduleListType {
    
        if t == nil {
                t = TeachingGroupScheduleListTypeCreate(TeachingGroupScheduleListType{})
        }
        if t.TeachingGroupSchedule == nil {
                t.TeachingGroupSchedule = make([]TeachingGroupScheduleType, 0)
        }
        t.TeachingGroupSchedule = append(t.TeachingGroupSchedule, value)
        return t
}

func (t *TeachingGroupScheduleListType) AddNew() *TeachingGroupScheduleListType {
        if t == nil {
                t = TeachingGroupScheduleListTypeCreate(TeachingGroupScheduleListType{})
        }
        if t.TeachingGroupSchedule == nil {
                t.TeachingGroupSchedule = make([]TeachingGroupScheduleType, 0)
        }
        t.TeachingGroupSchedule = append(t.TeachingGroupSchedule, TeachingGroupScheduleType{})
        return t
}

func (t *TeachingGroupScheduleListType) Last() *TeachingGroupScheduleType {
        return &(t.TeachingGroupSchedule[len(t.TeachingGroupSchedule)-1])
}

  func (t *CharacteristicsType) Add(value string) *CharacteristicsType {
    
        if t == nil {
                t = CharacteristicsTypeCreate(CharacteristicsType{})
        }
        if t.AggregateCharacteristicInfoRefId == nil {
                t.AggregateCharacteristicInfoRefId = make([]string, 0)
        }
        t.AggregateCharacteristicInfoRefId = append(t.AggregateCharacteristicInfoRefId, value)
        return t
}

func (t *CharacteristicsType) AddNew() *CharacteristicsType {
        if t == nil {
                t = CharacteristicsTypeCreate(CharacteristicsType{})
        }
        if t.AggregateCharacteristicInfoRefId == nil {
                t.AggregateCharacteristicInfoRefId = make([]string, 0)
        }
        t.AggregateCharacteristicInfoRefId = append(t.AggregateCharacteristicInfoRefId, "")
        return t
}

func (t *CharacteristicsType) Last() *string {
        return &(t.AggregateCharacteristicInfoRefId[len(t.AggregateCharacteristicInfoRefId)-1])
}

  func (t *RoomListType) Add(value string) *RoomListType {
    
        if t == nil {
                t = RoomListTypeCreate(RoomListType{})
        }
        if t.RoomInfoRefId == nil {
                t.RoomInfoRefId = make([]string, 0)
        }
        t.RoomInfoRefId = append(t.RoomInfoRefId, value)
        return t
}

func (t *RoomListType) AddNew() *RoomListType {
        if t == nil {
                t = RoomListTypeCreate(RoomListType{})
        }
        if t.RoomInfoRefId == nil {
                t.RoomInfoRefId = make([]string, 0)
        }
        t.RoomInfoRefId = append(t.RoomInfoRefId, "")
        return t
}

func (t *RoomListType) Last() *string {
        return &(t.RoomInfoRefId[len(t.RoomInfoRefId)-1])
}

  func (t *JournalAdjustmentListType) Add(value JournalAdjustmentType) *JournalAdjustmentListType {
    
        if t == nil {
                t = JournalAdjustmentListTypeCreate(JournalAdjustmentListType{})
        }
        if t.JournalAdjustment == nil {
                t.JournalAdjustment = make([]JournalAdjustmentType, 0)
        }
        t.JournalAdjustment = append(t.JournalAdjustment, value)
        return t
}

func (t *JournalAdjustmentListType) AddNew() *JournalAdjustmentListType {
        if t == nil {
                t = JournalAdjustmentListTypeCreate(JournalAdjustmentListType{})
        }
        if t.JournalAdjustment == nil {
                t.JournalAdjustment = make([]JournalAdjustmentType, 0)
        }
        t.JournalAdjustment = append(t.JournalAdjustment, JournalAdjustmentType{})
        return t
}

func (t *JournalAdjustmentListType) Last() *JournalAdjustmentType {
        return &(t.JournalAdjustment[len(t.JournalAdjustment)-1])
}

  func (t *CodeFrameTestItemListType) Add(value CodeFrameTestItemType) *CodeFrameTestItemListType {
    
        if t == nil {
                t = CodeFrameTestItemListTypeCreate(CodeFrameTestItemListType{})
        }
        if t.TestItem == nil {
                t.TestItem = make([]CodeFrameTestItemType, 0)
        }
        t.TestItem = append(t.TestItem, value)
        return t
}

func (t *CodeFrameTestItemListType) AddNew() *CodeFrameTestItemListType {
        if t == nil {
                t = CodeFrameTestItemListTypeCreate(CodeFrameTestItemListType{})
        }
        if t.TestItem == nil {
                t.TestItem = make([]CodeFrameTestItemType, 0)
        }
        t.TestItem = append(t.TestItem, CodeFrameTestItemType{})
        return t
}

func (t *CodeFrameTestItemListType) Last() *CodeFrameTestItemType {
        return &(t.TestItem[len(t.TestItem)-1])
}

  func (t *ValidLetterMarkListType) Add(value ValidLetterMarkType) *ValidLetterMarkListType {
    
        if t == nil {
                t = ValidLetterMarkListTypeCreate(ValidLetterMarkListType{})
        }
        if t.ValidLetterMark == nil {
                t.ValidLetterMark = make([]ValidLetterMarkType, 0)
        }
        t.ValidLetterMark = append(t.ValidLetterMark, value)
        return t
}

func (t *ValidLetterMarkListType) AddNew() *ValidLetterMarkListType {
        if t == nil {
                t = ValidLetterMarkListTypeCreate(ValidLetterMarkListType{})
        }
        if t.ValidLetterMark == nil {
                t.ValidLetterMark = make([]ValidLetterMarkType, 0)
        }
        t.ValidLetterMark = append(t.ValidLetterMark, ValidLetterMarkType{})
        return t
}

func (t *ValidLetterMarkListType) Last() *ValidLetterMarkType {
        return &(t.ValidLetterMark[len(t.ValidLetterMark)-1])
}

  func (t *MedicalAlertMessagesType) Add(value MedicalAlertMessageType) *MedicalAlertMessagesType {
    
        if t == nil {
                t = MedicalAlertMessagesTypeCreate(MedicalAlertMessagesType{})
        }
        if t.MedicalAlertMessage == nil {
                t.MedicalAlertMessage = make([]MedicalAlertMessageType, 0)
        }
        t.MedicalAlertMessage = append(t.MedicalAlertMessage, value)
        return t
}

func (t *MedicalAlertMessagesType) AddNew() *MedicalAlertMessagesType {
        if t == nil {
                t = MedicalAlertMessagesTypeCreate(MedicalAlertMessagesType{})
        }
        if t.MedicalAlertMessage == nil {
                t.MedicalAlertMessage = make([]MedicalAlertMessageType, 0)
        }
        t.MedicalAlertMessage = append(t.MedicalAlertMessage, MedicalAlertMessageType{})
        return t
}

func (t *MedicalAlertMessagesType) Last() *MedicalAlertMessageType {
        return &(t.MedicalAlertMessage[len(t.MedicalAlertMessage)-1])
}

  func (t *ExclusionRulesType) Add(value ExclusionRuleType) *ExclusionRulesType {
    
        if t == nil {
                t = ExclusionRulesTypeCreate(ExclusionRulesType{})
        }
        if t.ExclusionRule == nil {
                t.ExclusionRule = make([]ExclusionRuleType, 0)
        }
        t.ExclusionRule = append(t.ExclusionRule, value)
        return t
}

func (t *ExclusionRulesType) AddNew() *ExclusionRulesType {
        if t == nil {
                t = ExclusionRulesTypeCreate(ExclusionRulesType{})
        }
        if t.ExclusionRule == nil {
                t.ExclusionRule = make([]ExclusionRuleType, 0)
        }
        t.ExclusionRule = append(t.ExclusionRule, ExclusionRuleType{})
        return t
}

func (t *ExclusionRulesType) Last() *ExclusionRuleType {
        return &(t.ExclusionRule[len(t.ExclusionRule)-1])
}

  func (t *FQReportingListType) Add(value FQReportingType) *FQReportingListType {
    
        if t == nil {
                t = FQReportingListTypeCreate(FQReportingListType{})
        }
        if t.FQReporting == nil {
                t.FQReporting = make([]FQReportingType, 0)
        }
        t.FQReporting = append(t.FQReporting, value)
        return t
}

func (t *FQReportingListType) AddNew() *FQReportingListType {
        if t == nil {
                t = FQReportingListTypeCreate(FQReportingListType{})
        }
        if t.FQReporting == nil {
                t.FQReporting = make([]FQReportingType, 0)
        }
        t.FQReporting = append(t.FQReporting, FQReportingType{})
        return t
}

func (t *FQReportingListType) Last() *FQReportingType {
        return &(t.FQReporting[len(t.FQReporting)-1])
}

  func (t *ExpenseAccountsType) Add(value ExpenseAccountType) *ExpenseAccountsType {
    
        if t == nil {
                t = ExpenseAccountsTypeCreate(ExpenseAccountsType{})
        }
        if t.ExpenseAccount == nil {
                t.ExpenseAccount = make([]ExpenseAccountType, 0)
        }
        t.ExpenseAccount = append(t.ExpenseAccount, value)
        return t
}

func (t *ExpenseAccountsType) AddNew() *ExpenseAccountsType {
        if t == nil {
                t = ExpenseAccountsTypeCreate(ExpenseAccountsType{})
        }
        if t.ExpenseAccount == nil {
                t.ExpenseAccount = make([]ExpenseAccountType, 0)
        }
        t.ExpenseAccount = append(t.ExpenseAccount, ExpenseAccountType{})
        return t
}

func (t *ExpenseAccountsType) Last() *ExpenseAccountType {
        return &(t.ExpenseAccount[len(t.ExpenseAccount)-1])
}

  func (t *AGContextualQuestionListType) Add(value AGContextualQuestionType) *AGContextualQuestionListType {
    
        if t == nil {
                t = AGContextualQuestionListTypeCreate(AGContextualQuestionListType{})
        }
        if t.AGContextualQuestion == nil {
                t.AGContextualQuestion = make([]AGContextualQuestionType, 0)
        }
        t.AGContextualQuestion = append(t.AGContextualQuestion, value)
        return t
}

func (t *AGContextualQuestionListType) AddNew() *AGContextualQuestionListType {
        if t == nil {
                t = AGContextualQuestionListTypeCreate(AGContextualQuestionListType{})
        }
        if t.AGContextualQuestion == nil {
                t.AGContextualQuestion = make([]AGContextualQuestionType, 0)
        }
        t.AGContextualQuestion = append(t.AGContextualQuestion, AGContextualQuestionType{})
        return t
}

func (t *AGContextualQuestionListType) Last() *AGContextualQuestionType {
        return &(t.AGContextualQuestion[len(t.AGContextualQuestion)-1])
}

  func (t *NAPWritingRubricListType) Add(value NAPWritingRubricType) *NAPWritingRubricListType {
    
        if t == nil {
                t = NAPWritingRubricListTypeCreate(NAPWritingRubricListType{})
        }
        if t.NAPWritingRubric == nil {
                t.NAPWritingRubric = make([]NAPWritingRubricType, 0)
        }
        t.NAPWritingRubric = append(t.NAPWritingRubric, value)
        return t
}

func (t *NAPWritingRubricListType) AddNew() *NAPWritingRubricListType {
        if t == nil {
                t = NAPWritingRubricListTypeCreate(NAPWritingRubricListType{})
        }
        if t.NAPWritingRubric == nil {
                t.NAPWritingRubric = make([]NAPWritingRubricType, 0)
        }
        t.NAPWritingRubric = append(t.NAPWritingRubric, NAPWritingRubricType{})
        return t
}

func (t *NAPWritingRubricListType) Last() *NAPWritingRubricType {
        return &(t.NAPWritingRubric[len(t.NAPWritingRubric)-1])
}

  func (t *LearningStandardsType) Add(value string) *LearningStandardsType {
    
        if t == nil {
                t = LearningStandardsTypeCreate(LearningStandardsType{})
        }
        if t.LearningStandardItemRefId == nil {
                t.LearningStandardItemRefId = make([]string, 0)
        }
        t.LearningStandardItemRefId = append(t.LearningStandardItemRefId, value)
        return t
}

func (t *LearningStandardsType) AddNew() *LearningStandardsType {
        if t == nil {
                t = LearningStandardsTypeCreate(LearningStandardsType{})
        }
        if t.LearningStandardItemRefId == nil {
                t.LearningStandardItemRefId = make([]string, 0)
        }
        t.LearningStandardItemRefId = append(t.LearningStandardItemRefId, "")
        return t
}

func (t *LearningStandardsType) Last() *string {
        return &(t.LearningStandardItemRefId[len(t.LearningStandardItemRefId)-1])
}

  func (t *YearLevelsType) Add(value YearLevelType) *YearLevelsType {
    
        if t == nil {
                t = YearLevelsTypeCreate(YearLevelsType{})
        }
        if t.YearLevel == nil {
                t.YearLevel = make([]YearLevelType, 0)
        }
        t.YearLevel = append(t.YearLevel, value)
        return t
}

func (t *YearLevelsType) AddNew() *YearLevelsType {
        if t == nil {
                t = YearLevelsTypeCreate(YearLevelsType{})
        }
        if t.YearLevel == nil {
                t.YearLevel = make([]YearLevelType, 0)
        }
        t.YearLevel = append(t.YearLevel, YearLevelType{})
        return t
}

func (t *YearLevelsType) Last() *YearLevelType {
        return &(t.YearLevel[len(t.YearLevel)-1])
}

  func (t *SourceObjectsType) Add(value SourceObjectsType_SourceObject) *SourceObjectsType {
    
        if t == nil {
                t = SourceObjectsTypeCreate(SourceObjectsType{})
        }
        if t.SourceObject == nil {
                t.SourceObject = make([]SourceObjectsType_SourceObject, 0)
        }
        t.SourceObject = append(t.SourceObject, value)
        return t
}

func (t *SourceObjectsType) AddNew() *SourceObjectsType {
        if t == nil {
                t = SourceObjectsTypeCreate(SourceObjectsType{})
        }
        if t.SourceObject == nil {
                t.SourceObject = make([]SourceObjectsType_SourceObject, 0)
        }
        t.SourceObject = append(t.SourceObject, SourceObjectsType_SourceObject{})
        return t
}

func (t *SourceObjectsType) Last() *SourceObjectsType_SourceObject {
        return &(t.SourceObject[len(t.SourceObject)-1])
}

  func (t *FollowUpActionListType) Add(value FollowUpActionType) *FollowUpActionListType {
    
        if t == nil {
                t = FollowUpActionListTypeCreate(FollowUpActionListType{})
        }
        if t.FollowUpAction == nil {
                t.FollowUpAction = make([]FollowUpActionType, 0)
        }
        t.FollowUpAction = append(t.FollowUpAction, value)
        return t
}

func (t *FollowUpActionListType) AddNew() *FollowUpActionListType {
        if t == nil {
                t = FollowUpActionListTypeCreate(FollowUpActionListType{})
        }
        if t.FollowUpAction == nil {
                t.FollowUpAction = make([]FollowUpActionType, 0)
        }
        t.FollowUpAction = append(t.FollowUpAction, FollowUpActionType{})
        return t
}

func (t *FollowUpActionListType) Last() *FollowUpActionType {
        return &(t.FollowUpAction[len(t.FollowUpAction)-1])
}

  func (t *EssentialMaterialsType) Add(value string) *EssentialMaterialsType {
    
        if t == nil {
                t = EssentialMaterialsTypeCreate(EssentialMaterialsType{})
        }
        if t.EssentialMaterial == nil {
                t.EssentialMaterial = make([]string, 0)
        }
        t.EssentialMaterial = append(t.EssentialMaterial, value)
        return t
}

func (t *EssentialMaterialsType) AddNew() *EssentialMaterialsType {
        if t == nil {
                t = EssentialMaterialsTypeCreate(EssentialMaterialsType{})
        }
        if t.EssentialMaterial == nil {
                t.EssentialMaterial = make([]string, 0)
        }
        t.EssentialMaterial = append(t.EssentialMaterial, "")
        return t
}

func (t *EssentialMaterialsType) Last() *string {
        return &(t.EssentialMaterial[len(t.EssentialMaterial)-1])
}

  func (t *AlertMessagesType) Add(value AlertMessageType) *AlertMessagesType {
    
        if t == nil {
                t = AlertMessagesTypeCreate(AlertMessagesType{})
        }
        if t.AlertMessage == nil {
                t.AlertMessage = make([]AlertMessageType, 0)
        }
        t.AlertMessage = append(t.AlertMessage, value)
        return t
}

func (t *AlertMessagesType) AddNew() *AlertMessagesType {
        if t == nil {
                t = AlertMessagesTypeCreate(AlertMessagesType{})
        }
        if t.AlertMessage == nil {
                t.AlertMessage = make([]AlertMessageType, 0)
        }
        t.AlertMessage = append(t.AlertMessage, AlertMessageType{})
        return t
}

func (t *AlertMessagesType) Last() *AlertMessageType {
        return &(t.AlertMessage[len(t.AlertMessage)-1])
}

  func (t *OtherNamesType) Add(value OtherNameType) *OtherNamesType {
    
        if t == nil {
                t = OtherNamesTypeCreate(OtherNamesType{})
        }
        if t.Name == nil {
                t.Name = make([]OtherNameType, 0)
        }
        t.Name = append(t.Name, value)
        return t
}

func (t *OtherNamesType) AddNew() *OtherNamesType {
        if t == nil {
                t = OtherNamesTypeCreate(OtherNamesType{})
        }
        if t.Name == nil {
                t.Name = make([]OtherNameType, 0)
        }
        t.Name = append(t.Name, OtherNameType{})
        return t
}

func (t *OtherNamesType) Last() *OtherNameType {
        return &(t.Name[len(t.Name)-1])
}

  func (t *StudentListType) Add(value TeachingGroupStudentType) *StudentListType {
    
        if t == nil {
                t = StudentListTypeCreate(StudentListType{})
        }
        if t.TeachingGroupStudent == nil {
                t.TeachingGroupStudent = make([]TeachingGroupStudentType, 0)
        }
        t.TeachingGroupStudent = append(t.TeachingGroupStudent, value)
        return t
}

func (t *StudentListType) AddNew() *StudentListType {
        if t == nil {
                t = StudentListTypeCreate(StudentListType{})
        }
        if t.TeachingGroupStudent == nil {
                t.TeachingGroupStudent = make([]TeachingGroupStudentType, 0)
        }
        t.TeachingGroupStudent = append(t.TeachingGroupStudent, TeachingGroupStudentType{})
        return t
}

func (t *StudentListType) Last() *TeachingGroupStudentType {
        return &(t.TeachingGroupStudent[len(t.TeachingGroupStudent)-1])
}

  func (t *PhoneNumberListType) Add(value PhoneNumberType) *PhoneNumberListType {
    
        if t == nil {
                t = PhoneNumberListTypeCreate(PhoneNumberListType{})
        }
        if t.PhoneNumber == nil {
                t.PhoneNumber = make([]PhoneNumberType, 0)
        }
        t.PhoneNumber = append(t.PhoneNumber, value)
        return t
}

func (t *PhoneNumberListType) AddNew() *PhoneNumberListType {
        if t == nil {
                t = PhoneNumberListTypeCreate(PhoneNumberListType{})
        }
        if t.PhoneNumber == nil {
                t.PhoneNumber = make([]PhoneNumberType, 0)
        }
        t.PhoneNumber = append(t.PhoneNumber, PhoneNumberType{})
        return t
}

func (t *PhoneNumberListType) Last() *PhoneNumberType {
        return &(t.PhoneNumber[len(t.PhoneNumber)-1])
}

  func (t *StudentGroupListType) Add(value StudentGroupType) *StudentGroupListType {
    
        if t == nil {
                t = StudentGroupListTypeCreate(StudentGroupListType{})
        }
        if t.StudentGroup == nil {
                t.StudentGroup = make([]StudentGroupType, 0)
        }
        t.StudentGroup = append(t.StudentGroup, value)
        return t
}

func (t *StudentGroupListType) AddNew() *StudentGroupListType {
        if t == nil {
                t = StudentGroupListTypeCreate(StudentGroupListType{})
        }
        if t.StudentGroup == nil {
                t.StudentGroup = make([]StudentGroupType, 0)
        }
        t.StudentGroup = append(t.StudentGroup, StudentGroupType{})
        return t
}

func (t *StudentGroupListType) Last() *StudentGroupType {
        return &(t.StudentGroup[len(t.StudentGroup)-1])
}

  func (t *HoldInfoListType) Add(value HoldInfoType) *HoldInfoListType {
    
        if t == nil {
                t = HoldInfoListTypeCreate(HoldInfoListType{})
        }
        if t.HoldInfo == nil {
                t.HoldInfo = make([]HoldInfoType, 0)
        }
        t.HoldInfo = append(t.HoldInfo, value)
        return t
}

func (t *HoldInfoListType) AddNew() *HoldInfoListType {
        if t == nil {
                t = HoldInfoListTypeCreate(HoldInfoListType{})
        }
        if t.HoldInfo == nil {
                t.HoldInfo = make([]HoldInfoType, 0)
        }
        t.HoldInfo = append(t.HoldInfo, HoldInfoType{})
        return t
}

func (t *HoldInfoListType) Last() *HoldInfoType {
        return &(t.HoldInfo[len(t.HoldInfo)-1])
}

  func (t *TeachingGroupListType) Add(value string) *TeachingGroupListType {
    
        if t == nil {
                t = TeachingGroupListTypeCreate(TeachingGroupListType{})
        }
        if t.TeachingGroupRefId == nil {
                t.TeachingGroupRefId = make([]string, 0)
        }
        t.TeachingGroupRefId = append(t.TeachingGroupRefId, value)
        return t
}

func (t *TeachingGroupListType) AddNew() *TeachingGroupListType {
        if t == nil {
                t = TeachingGroupListTypeCreate(TeachingGroupListType{})
        }
        if t.TeachingGroupRefId == nil {
                t.TeachingGroupRefId = make([]string, 0)
        }
        t.TeachingGroupRefId = append(t.TeachingGroupRefId, "")
        return t
}

func (t *TeachingGroupListType) Last() *string {
        return &(t.TeachingGroupRefId[len(t.TeachingGroupRefId)-1])
}

  func (t *WellbeingDocumentListType) Add(value WellbeingDocumentType) *WellbeingDocumentListType {
    
        if t == nil {
                t = WellbeingDocumentListTypeCreate(WellbeingDocumentListType{})
        }
        if t.Document == nil {
                t.Document = make([]WellbeingDocumentType, 0)
        }
        t.Document = append(t.Document, value)
        return t
}

func (t *WellbeingDocumentListType) AddNew() *WellbeingDocumentListType {
        if t == nil {
                t = WellbeingDocumentListTypeCreate(WellbeingDocumentListType{})
        }
        if t.Document == nil {
                t.Document = make([]WellbeingDocumentType, 0)
        }
        t.Document = append(t.Document, WellbeingDocumentType{})
        return t
}

func (t *WellbeingDocumentListType) Last() *WellbeingDocumentType {
        return &(t.Document[len(t.Document)-1])
}

  func (t *OrganizationsType) Add(value string) *OrganizationsType {
    
        if t == nil {
                t = OrganizationsTypeCreate(OrganizationsType{})
        }
        if t.Organization == nil {
                t.Organization = make([]string, 0)
        }
        t.Organization = append(t.Organization, value)
        return t
}

func (t *OrganizationsType) AddNew() *OrganizationsType {
        if t == nil {
                t = OrganizationsTypeCreate(OrganizationsType{})
        }
        if t.Organization == nil {
                t.Organization = make([]string, 0)
        }
        t.Organization = append(t.Organization, "")
        return t
}

func (t *OrganizationsType) Last() *string {
        return &(t.Organization[len(t.Organization)-1])
}

  func (t *LocalCodeListType) Add(value LocalCodeType) *LocalCodeListType {
    
        if t == nil {
                t = LocalCodeListTypeCreate(LocalCodeListType{})
        }
        if t.LocalCode == nil {
                t.LocalCode = make([]LocalCodeType, 0)
        }
        t.LocalCode = append(t.LocalCode, value)
        return t
}

func (t *LocalCodeListType) AddNew() *LocalCodeListType {
        if t == nil {
                t = LocalCodeListTypeCreate(LocalCodeListType{})
        }
        if t.LocalCode == nil {
                t.LocalCode = make([]LocalCodeType, 0)
        }
        t.LocalCode = append(t.LocalCode, LocalCodeType{})
        return t
}

func (t *LocalCodeListType) Last() *LocalCodeType {
        return &(t.LocalCode[len(t.LocalCode)-1])
}

  func (t *TestDisruptionListType) Add(value TestDisruptionType) *TestDisruptionListType {
    
        if t == nil {
                t = TestDisruptionListTypeCreate(TestDisruptionListType{})
        }
        if t.TestDisruption == nil {
                t.TestDisruption = make([]TestDisruptionType, 0)
        }
        t.TestDisruption = append(t.TestDisruption, value)
        return t
}

func (t *TestDisruptionListType) AddNew() *TestDisruptionListType {
        if t == nil {
                t = TestDisruptionListTypeCreate(TestDisruptionListType{})
        }
        if t.TestDisruption == nil {
                t.TestDisruption = make([]TestDisruptionType, 0)
        }
        t.TestDisruption = append(t.TestDisruption, TestDisruptionType{})
        return t
}

func (t *TestDisruptionListType) Last() *TestDisruptionType {
        return &(t.TestDisruption[len(t.TestDisruption)-1])
}

  func (t *PublishingPermissionListType) Add(value PublishingPermissionType) *PublishingPermissionListType {
    
        if t == nil {
                t = PublishingPermissionListTypeCreate(PublishingPermissionListType{})
        }
        if t.PublishingPermission == nil {
                t.PublishingPermission = make([]PublishingPermissionType, 0)
        }
        t.PublishingPermission = append(t.PublishingPermission, value)
        return t
}

func (t *PublishingPermissionListType) AddNew() *PublishingPermissionListType {
        if t == nil {
                t = PublishingPermissionListTypeCreate(PublishingPermissionListType{})
        }
        if t.PublishingPermission == nil {
                t.PublishingPermission = make([]PublishingPermissionType, 0)
        }
        t.PublishingPermission = append(t.PublishingPermission, PublishingPermissionType{})
        return t
}

func (t *PublishingPermissionListType) Last() *PublishingPermissionType {
        return &(t.PublishingPermission[len(t.PublishingPermission)-1])
}

  func (t *TeachingGroupPeriodListType) Add(value TeachingGroupPeriodType) *TeachingGroupPeriodListType {
    
        if t == nil {
                t = TeachingGroupPeriodListTypeCreate(TeachingGroupPeriodListType{})
        }
        if t.TeachingGroupPeriod == nil {
                t.TeachingGroupPeriod = make([]TeachingGroupPeriodType, 0)
        }
        t.TeachingGroupPeriod = append(t.TeachingGroupPeriod, value)
        return t
}

func (t *TeachingGroupPeriodListType) AddNew() *TeachingGroupPeriodListType {
        if t == nil {
                t = TeachingGroupPeriodListTypeCreate(TeachingGroupPeriodListType{})
        }
        if t.TeachingGroupPeriod == nil {
                t.TeachingGroupPeriod = make([]TeachingGroupPeriodType, 0)
        }
        t.TeachingGroupPeriod = append(t.TeachingGroupPeriod, TeachingGroupPeriodType{})
        return t
}

func (t *TeachingGroupPeriodListType) Last() *TeachingGroupPeriodType {
        return &(t.TeachingGroupPeriod[len(t.TeachingGroupPeriod)-1])
}

  func (t *LEAContactListType) Add(value LEAContactType) *LEAContactListType {
    
        if t == nil {
                t = LEAContactListTypeCreate(LEAContactListType{})
        }
        if t.LEAContact == nil {
                t.LEAContact = make([]LEAContactType, 0)
        }
        t.LEAContact = append(t.LEAContact, value)
        return t
}

func (t *LEAContactListType) AddNew() *LEAContactListType {
        if t == nil {
                t = LEAContactListTypeCreate(LEAContactListType{})
        }
        if t.LEAContact == nil {
                t.LEAContact = make([]LEAContactType, 0)
        }
        t.LEAContact = append(t.LEAContact, LEAContactType{})
        return t
}

func (t *LEAContactListType) Last() *LEAContactType {
        return &(t.LEAContact[len(t.LEAContact)-1])
}

  func (t *OtherIdListType) Add(value OtherIdType) *OtherIdListType {
    
        if t == nil {
                t = OtherIdListTypeCreate(OtherIdListType{})
        }
        if t.OtherId == nil {
                t.OtherId = make([]OtherIdType, 0)
        }
        t.OtherId = append(t.OtherId, value)
        return t
}

func (t *OtherIdListType) AddNew() *OtherIdListType {
        if t == nil {
                t = OtherIdListTypeCreate(OtherIdListType{})
        }
        if t.OtherId == nil {
                t.OtherId = make([]OtherIdType, 0)
        }
        t.OtherId = append(t.OtherId, OtherIdType{})
        return t
}

func (t *OtherIdListType) Last() *OtherIdType {
        return &(t.OtherId[len(t.OtherId)-1])
}

  func (t *SIF_ExtendedElementsType) Add(value SIF_ExtendedElementsType_SIF_ExtendedElement) *SIF_ExtendedElementsType {
    
        if t == nil {
                t = SIF_ExtendedElementsTypeCreate(SIF_ExtendedElementsType{})
        }
        if t.SIF_ExtendedElement == nil {
                t.SIF_ExtendedElement = make([]SIF_ExtendedElementsType_SIF_ExtendedElement, 0)
        }
        t.SIF_ExtendedElement = append(t.SIF_ExtendedElement, value)
        return t
}

func (t *SIF_ExtendedElementsType) AddNew() *SIF_ExtendedElementsType {
        if t == nil {
                t = SIF_ExtendedElementsTypeCreate(SIF_ExtendedElementsType{})
        }
        if t.SIF_ExtendedElement == nil {
                t.SIF_ExtendedElement = make([]SIF_ExtendedElementsType_SIF_ExtendedElement, 0)
        }
        t.SIF_ExtendedElement = append(t.SIF_ExtendedElement, SIF_ExtendedElementsType_SIF_ExtendedElement{})
        return t
}

func (t *SIF_ExtendedElementsType) Last() *SIF_ExtendedElementsType_SIF_ExtendedElement {
        return &(t.SIF_ExtendedElement[len(t.SIF_ExtendedElement)-1])
}

  func (t *CensusReportingListType) Add(value CensusReportingType) *CensusReportingListType {
    
        if t == nil {
                t = CensusReportingListTypeCreate(CensusReportingListType{})
        }
        if t.CensusReporting == nil {
                t.CensusReporting = make([]CensusReportingType, 0)
        }
        t.CensusReporting = append(t.CensusReporting, value)
        return t
}

func (t *CensusReportingListType) AddNew() *CensusReportingListType {
        if t == nil {
                t = CensusReportingListTypeCreate(CensusReportingListType{})
        }
        if t.CensusReporting == nil {
                t.CensusReporting = make([]CensusReportingType, 0)
        }
        t.CensusReporting = append(t.CensusReporting, CensusReportingType{})
        return t
}

func (t *CensusReportingListType) Last() *CensusReportingType {
        return &(t.CensusReporting[len(t.CensusReporting)-1])
}

  func (t *AGRoundListType) Add(value AGRoundType) *AGRoundListType {
    
        if t == nil {
                t = AGRoundListTypeCreate(AGRoundListType{})
        }
        if t.AGRound == nil {
                t.AGRound = make([]AGRoundType, 0)
        }
        t.AGRound = append(t.AGRound, value)
        return t
}

func (t *AGRoundListType) AddNew() *AGRoundListType {
        if t == nil {
                t = AGRoundListTypeCreate(AGRoundListType{})
        }
        if t.AGRound == nil {
                t.AGRound = make([]AGRoundType, 0)
        }
        t.AGRound = append(t.AGRound, AGRoundType{})
        return t
}

func (t *AGRoundListType) Last() *AGRoundType {
        return &(t.AGRound[len(t.AGRound)-1])
}

  func (t *LearningStandardListType) Add(value LearningStandardType) *LearningStandardListType {
    
        if t == nil {
                t = LearningStandardListTypeCreate(LearningStandardListType{})
        }
        if t.LearningStandard == nil {
                t.LearningStandard = make([]LearningStandardType, 0)
        }
        t.LearningStandard = append(t.LearningStandard, value)
        return t
}

func (t *LearningStandardListType) AddNew() *LearningStandardListType {
        if t == nil {
                t = LearningStandardListTypeCreate(LearningStandardListType{})
        }
        if t.LearningStandard == nil {
                t.LearningStandard = make([]LearningStandardType, 0)
        }
        t.LearningStandard = append(t.LearningStandard, LearningStandardType{})
        return t
}

func (t *LearningStandardListType) Last() *LearningStandardType {
        return &(t.LearningStandard[len(t.LearningStandard)-1])
}

  func (t *MediaTypesType) Add(value string) *MediaTypesType {
    
        if t == nil {
                t = MediaTypesTypeCreate(MediaTypesType{})
        }
        if t.MediaType == nil {
                t.MediaType = make([]string, 0)
        }
        t.MediaType = append(t.MediaType, value)
        return t
}

func (t *MediaTypesType) AddNew() *MediaTypesType {
        if t == nil {
                t = MediaTypesTypeCreate(MediaTypesType{})
        }
        if t.MediaType == nil {
                t.MediaType = make([]string, 0)
        }
        t.MediaType = append(t.MediaType, "")
        return t
}

func (t *MediaTypesType) Last() *string {
        return &(t.MediaType[len(t.MediaType)-1])
}

  func (t *AttendanceTimesType) Add(value AttendanceTimeType) *AttendanceTimesType {
    
        if t == nil {
                t = AttendanceTimesTypeCreate(AttendanceTimesType{})
        }
        if t.AttendanceTime == nil {
                t.AttendanceTime = make([]AttendanceTimeType, 0)
        }
        t.AttendanceTime = append(t.AttendanceTime, value)
        return t
}

func (t *AttendanceTimesType) AddNew() *AttendanceTimesType {
        if t == nil {
                t = AttendanceTimesTypeCreate(AttendanceTimesType{})
        }
        if t.AttendanceTime == nil {
                t.AttendanceTime = make([]AttendanceTimeType, 0)
        }
        t.AttendanceTime = append(t.AttendanceTime, AttendanceTimeType{})
        return t
}

func (t *AttendanceTimesType) Last() *AttendanceTimeType {
        return &(t.AttendanceTime[len(t.AttendanceTime)-1])
}

  func (t *CountryList2Type) Add(value CountryType) *CountryList2Type {
    
        if t == nil {
                t = CountryList2TypeCreate(CountryList2Type{})
        }
        if t.CountryOfResidency == nil {
                t.CountryOfResidency = make([]CountryType, 0)
        }
        t.CountryOfResidency = append(t.CountryOfResidency, value)
        return t
}

func (t *CountryList2Type) AddNew() *CountryList2Type {
        if t == nil {
                t = CountryList2TypeCreate(CountryList2Type{})
        }
        if t.CountryOfResidency == nil {
                t.CountryOfResidency = make([]CountryType, 0)
        }
        t.CountryOfResidency = append(t.CountryOfResidency, "")
        return t
}

func (t *CountryList2Type) Last() *CountryType {
        return &(t.CountryOfResidency[len(t.CountryOfResidency)-1])
}

  func (t *IdentityAssertionsType) Add(value IdentityAssertionsType_IdentityAssertion) *IdentityAssertionsType {
    
        if t == nil {
                t = IdentityAssertionsTypeCreate(IdentityAssertionsType{})
        }
        if t.IdentityAssertion == nil {
                t.IdentityAssertion = make([]IdentityAssertionsType_IdentityAssertion, 0)
        }
        t.IdentityAssertion = append(t.IdentityAssertion, value)
        return t
}

func (t *IdentityAssertionsType) AddNew() *IdentityAssertionsType {
        if t == nil {
                t = IdentityAssertionsTypeCreate(IdentityAssertionsType{})
        }
        if t.IdentityAssertion == nil {
                t.IdentityAssertion = make([]IdentityAssertionsType_IdentityAssertion, 0)
        }
        t.IdentityAssertion = append(t.IdentityAssertion, IdentityAssertionsType_IdentityAssertion{})
        return t
}

func (t *IdentityAssertionsType) Last() *IdentityAssertionsType_IdentityAssertion {
        return &(t.IdentityAssertion[len(t.IdentityAssertion)-1])
}

  func (t *CensusStaffListType) Add(value CensusStaffType) *CensusStaffListType {
    
        if t == nil {
                t = CensusStaffListTypeCreate(CensusStaffListType{})
        }
        if t.CensusStaff == nil {
                t.CensusStaff = make([]CensusStaffType, 0)
        }
        t.CensusStaff = append(t.CensusStaff, value)
        return t
}

func (t *CensusStaffListType) AddNew() *CensusStaffListType {
        if t == nil {
                t = CensusStaffListTypeCreate(CensusStaffListType{})
        }
        if t.CensusStaff == nil {
                t.CensusStaff = make([]CensusStaffType, 0)
        }
        t.CensusStaff = append(t.CensusStaff, CensusStaffType{})
        return t
}

func (t *CensusStaffListType) Last() *CensusStaffType {
        return &(t.CensusStaff[len(t.CensusStaff)-1])
}

  func (t *FinancialAccountRefIdListType) Add(value string) *FinancialAccountRefIdListType {
    
        if t == nil {
                t = FinancialAccountRefIdListTypeCreate(FinancialAccountRefIdListType{})
        }
        if t.FinancialAccountRefId == nil {
                t.FinancialAccountRefId = make([]string, 0)
        }
        t.FinancialAccountRefId = append(t.FinancialAccountRefId, value)
        return t
}

func (t *FinancialAccountRefIdListType) AddNew() *FinancialAccountRefIdListType {
        if t == nil {
                t = FinancialAccountRefIdListTypeCreate(FinancialAccountRefIdListType{})
        }
        if t.FinancialAccountRefId == nil {
                t.FinancialAccountRefId = make([]string, 0)
        }
        t.FinancialAccountRefId = append(t.FinancialAccountRefId, "")
        return t
}

func (t *FinancialAccountRefIdListType) Last() *string {
        return &(t.FinancialAccountRefId[len(t.FinancialAccountRefId)-1])
}

  func (t *StudentSubjectChoiceListType) Add(value StudentSubjectChoiceType) *StudentSubjectChoiceListType {
    
        if t == nil {
                t = StudentSubjectChoiceListTypeCreate(StudentSubjectChoiceListType{})
        }
        if t.StudentSubjectChoice == nil {
                t.StudentSubjectChoice = make([]StudentSubjectChoiceType, 0)
        }
        t.StudentSubjectChoice = append(t.StudentSubjectChoice, value)
        return t
}

func (t *StudentSubjectChoiceListType) AddNew() *StudentSubjectChoiceListType {
        if t == nil {
                t = StudentSubjectChoiceListTypeCreate(StudentSubjectChoiceListType{})
        }
        if t.StudentSubjectChoice == nil {
                t.StudentSubjectChoice = make([]StudentSubjectChoiceType, 0)
        }
        t.StudentSubjectChoice = append(t.StudentSubjectChoice, StudentSubjectChoiceType{})
        return t
}

func (t *StudentSubjectChoiceListType) Last() *StudentSubjectChoiceType {
        return &(t.StudentSubjectChoice[len(t.StudentSubjectChoice)-1])
}

  func (t *SubstituteItemListType) Add(value SubstituteItemType) *SubstituteItemListType {
    
        if t == nil {
                t = SubstituteItemListTypeCreate(SubstituteItemListType{})
        }
        if t.SubstituteItem == nil {
                t.SubstituteItem = make([]SubstituteItemType, 0)
        }
        t.SubstituteItem = append(t.SubstituteItem, value)
        return t
}

func (t *SubstituteItemListType) AddNew() *SubstituteItemListType {
        if t == nil {
                t = SubstituteItemListTypeCreate(SubstituteItemListType{})
        }
        if t.SubstituteItem == nil {
                t.SubstituteItem = make([]SubstituteItemType, 0)
        }
        t.SubstituteItem = append(t.SubstituteItem, SubstituteItemType{})
        return t
}

func (t *SubstituteItemListType) Last() *SubstituteItemType {
        return &(t.SubstituteItem[len(t.SubstituteItem)-1])
}

  func (t *LifeCycleType_ModificationHistory) Add(value LifeCycleType_Modified) *LifeCycleType_ModificationHistory {
    
        if t == nil {
                t = LifeCycleType_ModificationHistoryCreate(LifeCycleType_ModificationHistory{})
        }
        if t.Modified == nil {
                t.Modified = make([]LifeCycleType_Modified, 0)
        }
        t.Modified = append(t.Modified, value)
        return t
}

func (t *LifeCycleType_ModificationHistory) AddNew() *LifeCycleType_ModificationHistory {
        if t == nil {
                t = LifeCycleType_ModificationHistoryCreate(LifeCycleType_ModificationHistory{})
        }
        if t.Modified == nil {
                t.Modified = make([]LifeCycleType_Modified, 0)
        }
        t.Modified = append(t.Modified, LifeCycleType_Modified{})
        return t
}

func (t *LifeCycleType_ModificationHistory) Last() *LifeCycleType_Modified {
        return &(t.Modified[len(t.Modified)-1])
}

  func (t *MedicationListType) Add(value MedicationType) *MedicationListType {
    
        if t == nil {
                t = MedicationListTypeCreate(MedicationListType{})
        }
        if t.Medication == nil {
                t.Medication = make([]MedicationType, 0)
        }
        t.Medication = append(t.Medication, value)
        return t
}

func (t *MedicationListType) AddNew() *MedicationListType {
        if t == nil {
                t = MedicationListTypeCreate(MedicationListType{})
        }
        if t.Medication == nil {
                t.Medication = make([]MedicationType, 0)
        }
        t.Medication = append(t.Medication, MedicationType{})
        return t
}

func (t *MedicationListType) Last() *MedicationType {
        return &(t.Medication[len(t.Medication)-1])
}

  func (t *FineInfoListType) Add(value FineInfoType) *FineInfoListType {
    
        if t == nil {
                t = FineInfoListTypeCreate(FineInfoListType{})
        }
        if t.FineInfo == nil {
                t.FineInfo = make([]FineInfoType, 0)
        }
        t.FineInfo = append(t.FineInfo, value)
        return t
}

func (t *FineInfoListType) AddNew() *FineInfoListType {
        if t == nil {
                t = FineInfoListTypeCreate(FineInfoListType{})
        }
        if t.FineInfo == nil {
                t.FineInfo = make([]FineInfoType, 0)
        }
        t.FineInfo = append(t.FineInfo, FineInfoType{})
        return t
}

func (t *FineInfoListType) Last() *FineInfoType {
        return &(t.FineInfo[len(t.FineInfo)-1])
}

  func (t *SchoolFocusListType) Add(value AUCodeSetsSchoolFocusCodeType) *SchoolFocusListType {
      present := false
  for _, b := range AUCodeSetsSchoolFocusCodeType_values {
        if b == string(value) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolFocusCodeType_values")
      }

        if t == nil {
                t = SchoolFocusListTypeCreate(SchoolFocusListType{})
        }
        if t.SchoolFocus == nil {
                t.SchoolFocus = make([]AUCodeSetsSchoolFocusCodeType, 0)
        }
        t.SchoolFocus = append(t.SchoolFocus, value)
        return t
}

func (t *SchoolFocusListType) AddNew() *SchoolFocusListType {
        if t == nil {
                t = SchoolFocusListTypeCreate(SchoolFocusListType{})
        }
        if t.SchoolFocus == nil {
                t.SchoolFocus = make([]AUCodeSetsSchoolFocusCodeType, 0)
        }
        t.SchoolFocus = append(t.SchoolFocus, "")
        return t
}

func (t *SchoolFocusListType) Last() *AUCodeSetsSchoolFocusCodeType {
        return &(t.SchoolFocus[len(t.SchoolFocus)-1])
}

  func (t *LanguageListType) Add(value LanguageBaseType) *LanguageListType {
    
        if t == nil {
                t = LanguageListTypeCreate(LanguageListType{})
        }
        if t.Language == nil {
                t.Language = make([]LanguageBaseType, 0)
        }
        t.Language = append(t.Language, value)
        return t
}

func (t *LanguageListType) AddNew() *LanguageListType {
        if t == nil {
                t = LanguageListTypeCreate(LanguageListType{})
        }
        if t.Language == nil {
                t.Language = make([]LanguageBaseType, 0)
        }
        t.Language = append(t.Language, LanguageBaseType{})
        return t
}

func (t *LanguageListType) Last() *LanguageBaseType {
        return &(t.Language[len(t.Language)-1])
}

  func (t *LifeCycleType_TimeElements) Add(value TimeElementType) *LifeCycleType_TimeElements {
    
        if t == nil {
                t = LifeCycleType_TimeElementsCreate(LifeCycleType_TimeElements{})
        }
        if t.TimeElement == nil {
                t.TimeElement = make([]TimeElementType, 0)
        }
        t.TimeElement = append(t.TimeElement, value)
        return t
}

func (t *LifeCycleType_TimeElements) AddNew() *LifeCycleType_TimeElements {
        if t == nil {
                t = LifeCycleType_TimeElementsCreate(LifeCycleType_TimeElements{})
        }
        if t.TimeElement == nil {
                t.TimeElement = make([]TimeElementType, 0)
        }
        t.TimeElement = append(t.TimeElement, TimeElementType{})
        return t
}

func (t *LifeCycleType_TimeElements) Last() *TimeElementType {
        return &(t.TimeElement[len(t.TimeElement)-1])
}

  func (t *StatsCohortListType) Add(value StatsCohortType) *StatsCohortListType {
    
        if t == nil {
                t = StatsCohortListTypeCreate(StatsCohortListType{})
        }
        if t.StatsCohort == nil {
                t.StatsCohort = make([]StatsCohortType, 0)
        }
        t.StatsCohort = append(t.StatsCohort, value)
        return t
}

func (t *StatsCohortListType) AddNew() *StatsCohortListType {
        if t == nil {
                t = StatsCohortListTypeCreate(StatsCohortListType{})
        }
        if t.StatsCohort == nil {
                t.StatsCohort = make([]StatsCohortType, 0)
        }
        t.StatsCohort = append(t.StatsCohort, StatsCohortType{})
        return t
}

func (t *StatsCohortListType) Last() *StatsCohortType {
        return &(t.StatsCohort[len(t.StatsCohort)-1])
}

  func (t *SchoolProgramListType) Add(value SchoolProgramType) *SchoolProgramListType {
    
        if t == nil {
                t = SchoolProgramListTypeCreate(SchoolProgramListType{})
        }
        if t.Program == nil {
                t.Program = make([]SchoolProgramType, 0)
        }
        t.Program = append(t.Program, value)
        return t
}

func (t *SchoolProgramListType) AddNew() *SchoolProgramListType {
        if t == nil {
                t = SchoolProgramListTypeCreate(SchoolProgramListType{})
        }
        if t.Program == nil {
                t.Program = make([]SchoolProgramType, 0)
        }
        t.Program = append(t.Program, SchoolProgramType{})
        return t
}

func (t *SchoolProgramListType) Last() *SchoolProgramType {
        return &(t.Program[len(t.Program)-1])
}

  func (t *PasswordListType) Add(value PasswordListType_Password) *PasswordListType {
    
        if t == nil {
                t = PasswordListTypeCreate(PasswordListType{})
        }
        if t.Password == nil {
                t.Password = make([]PasswordListType_Password, 0)
        }
        t.Password = append(t.Password, value)
        return t
}

func (t *PasswordListType) AddNew() *PasswordListType {
        if t == nil {
                t = PasswordListTypeCreate(PasswordListType{})
        }
        if t.Password == nil {
                t.Password = make([]PasswordListType_Password, 0)
        }
        t.Password = append(t.Password, PasswordListType_Password{})
        return t
}

func (t *PasswordListType) Last() *PasswordListType_Password {
        return &(t.Password[len(t.Password)-1])
}

  func (t *SoftwareRequirementListType) Add(value SoftwareRequirementType) *SoftwareRequirementListType {
    
        if t == nil {
                t = SoftwareRequirementListTypeCreate(SoftwareRequirementListType{})
        }
        if t.SoftwareRequirement == nil {
                t.SoftwareRequirement = make([]SoftwareRequirementType, 0)
        }
        t.SoftwareRequirement = append(t.SoftwareRequirement, value)
        return t
}

func (t *SoftwareRequirementListType) AddNew() *SoftwareRequirementListType {
        if t == nil {
                t = SoftwareRequirementListTypeCreate(SoftwareRequirementListType{})
        }
        if t.SoftwareRequirement == nil {
                t.SoftwareRequirement = make([]SoftwareRequirementType, 0)
        }
        t.SoftwareRequirement = append(t.SoftwareRequirement, SoftwareRequirementType{})
        return t
}

func (t *SoftwareRequirementListType) Last() *SoftwareRequirementType {
        return &(t.SoftwareRequirement[len(t.SoftwareRequirement)-1])
}

func (n *HouseholdContactInfoType) Set(key string, value interface{}) *HouseholdContactInfoType {
        if n == nil {
                n = HouseholdContactInfoTypeCreate(HouseholdContactInfoType{})
        }
        switch key {
    case "HouseholdSalutation":
                      n.HouseholdSalutation = StringCreate(value.(string))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "HouseholdContactId":
    
                      n.HouseholdContactId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PasswordListType_Password) Set(key string, value interface{}) *PasswordListType_Password {
        if n == nil {
                n = PasswordListType_PasswordCreate(PasswordListType_Password{})
        }
        switch key {
    case "Algorithm":
                      n.Algorithm = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "KeyName":
                      n.KeyName = StringCreate(value.(string))
        }
        return n
}

func (n *ResourcesType) Set(key string, value interface{}) *ResourcesType {
        if n == nil {
                n = ResourcesTypeCreate(ResourcesType{})
        }
        switch key {
    case "ResourceType":
                      n.ResourceType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentActivityInfo) Set(key string, value interface{}) *StudentActivityInfo {
        if n == nil {
                n = StudentActivityInfoCreate(StudentActivityInfo{})
        }
        switch key {
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Location":
                      n.Location = LocationTypeCreate(value.(LocationType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "StudentActivityType":
                      n.StudentActivityType = StudentActivityTypeCreate(value.(StudentActivityType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CurricularStatus":
      present := false
  for _, b := range AUCodeSetsActivityTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsActivityTypeType_values")
      }

                      n.CurricularStatus = ((*AUCodeSetsActivityTypeType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "StudentActivityLevel":
                      n.StudentActivityLevel = StringCreate(value.(string))
        }
        return n
}

func (n *CensusCollection) Set(key string, value interface{}) *CensusCollection {
        if n == nil {
                n = CensusCollectionCreate(CensusCollection{})
        }
        switch key {
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "CensusReportingList":
                      n.CensusReportingList = CensusReportingListTypeCreate(value.(CensusReportingListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CensusYear":
    
                      n.CensusYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
        }
        return n
}

func (n *LifeCycleType_Created) Set(key string, value interface{}) *LifeCycleType_Created {
        if n == nil {
                n = LifeCycleType_CreatedCreate(LifeCycleType_Created{})
        }
        switch key {
    case "DateTime":
                      n.DateTime = StringCreate(value.(string))
    case "Creators":
                      n.Creators = LifeCycleType_CreatorsCreate(value.(LifeCycleType_Creators))
        }
        return n
}

func (n *MonetaryAmountType) Set(key string, value interface{}) *MonetaryAmountType {
        if n == nil {
                n = MonetaryAmountTypeCreate(MonetaryAmountType{})
        }
        switch key {
    case "Value":
                      n.Value = FloatCreate(value.(float64))
    case "Currency":
      present := false
  for _, b := range ISO4217CurrencyNamesAndCodeElementsType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "ISO4217CurrencyNamesAndCodeElementsType_values")
      }

                      n.Currency = ((*ISO4217CurrencyNamesAndCodeElementsType)(StringCreate(value.(string))))
        }
        return n
}

func (n *StaffPersonal) Set(key string, value interface{}) *StaffPersonal {
        if n == nil {
                n = StaffPersonalCreate(StaffPersonal{})
        }
        switch key {
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "EmploymentStatus":
      present := false
  for _, b := range AUCodeSetsStaffStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStaffStatusType_values")
      }

                      n.EmploymentStatus = ((*AUCodeSetsStaffStatusType)(StringCreate(value.(string))))
    case "MostRecent":
                      n.MostRecent = StaffMostRecentContainerTypeCreate(value.(StaffMostRecentContainerType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *AwardContainerType) Set(key string, value interface{}) *AwardContainerType {
        if n == nil {
                n = AwardContainerTypeCreate(AwardContainerType{})
        }
        switch key {
    case "AwardNotes":
                      n.AwardNotes = StringCreate(value.(string))
    case "AwardDate":
                      n.AwardDate = StringCreate(value.(string))
    case "AwardType":
                      n.AwardType = StringCreate(value.(string))
    case "AwardDescription":
                      n.AwardDescription = StringCreate(value.(string))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
        }
        return n
}

func (n *StudentSchoolEnrollment_Homeroom) Set(key string, value interface{}) *StudentSchoolEnrollment_Homeroom {
        if n == nil {
                n = StudentSchoolEnrollment_HomeroomCreate(StudentSchoolEnrollment_Homeroom{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *YearRangeType) Set(key string, value interface{}) *YearRangeType {
        if n == nil {
                n = YearRangeTypeCreate(YearRangeType{})
        }
        switch key {
    case "End":
                      n.End = YearLevelTypeCreate(value.(YearLevelType))
    case "Start":
                      n.Start = YearLevelTypeCreate(value.(YearLevelType))
        }
        return n
}

func (n *AbstractContentElementType_TextData) Set(key string, value interface{}) *AbstractContentElementType_TextData {
        if n == nil {
                n = AbstractContentElementType_TextDataCreate(AbstractContentElementType_TextData{})
        }
        switch key {
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
        }
        return n
}

func (n *LEAContactType) Set(key string, value interface{}) *LEAContactType {
        if n == nil {
                n = LEAContactTypeCreate(LEAContactType{})
        }
        switch key {
    case "PublishInDirectory":
    
                      n.PublishInDirectory = ((*PublishInDirectoryType)(StringCreate(value.(string))))
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
        }
        return n
}

func (n *NAPTestletCodeFrameType) Set(key string, value interface{}) *NAPTestletCodeFrameType {
        if n == nil {
                n = NAPTestletCodeFrameTypeCreate(NAPTestletCodeFrameType{})
        }
        switch key {
    case "TestItemList":
                      n.TestItemList = CodeFrameTestItemListTypeCreate(value.(CodeFrameTestItemListType))
    case "TestletContent":
                      n.TestletContent = NAPTestletContentTypeCreate(value.(NAPTestletContentType))
    case "NAPTestletRefId":
                      n.NAPTestletRefId = StringCreate(value.(string))
        }
        return n
}

func (n *EducationFilterType) Set(key string, value interface{}) *EducationFilterType {
        if n == nil {
                n = EducationFilterTypeCreate(EducationFilterType{})
        }
        switch key {
    case "LearningStandardItems":
                      n.LearningStandardItems = LearningStandardsTypeCreate(value.(LearningStandardsType))
        }
        return n
}

func (n *LibraryMessageType) Set(key string, value interface{}) *LibraryMessageType {
        if n == nil {
                n = LibraryMessageTypeCreate(LibraryMessageType{})
        }
        switch key {
    case "PriorityCodeset":
                      n.PriorityCodeset = StringCreate(value.(string))
    case "Text":
                      n.Text = StringCreate(value.(string))
    case "Sent":
                      n.Sent = StringCreate(value.(string))
    case "Priority":
                      n.Priority = StringCreate(value.(string))
        }
        return n
}

func (n *ExpenseAccountType) Set(key string, value interface{}) *ExpenseAccountType {
        if n == nil {
                n = ExpenseAccountTypeCreate(ExpenseAccountType{})
        }
        switch key {
    case "AccountingPeriod":
    
                      n.AccountingPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
    case "FinancialAccountRefId":
                      n.FinancialAccountRefId = StringCreate(value.(string))
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
        }
        return n
}

func (n *AbstractContentElementType_XMLData) Set(key string, value interface{}) *AbstractContentElementType_XMLData {
        if n == nil {
                n = AbstractContentElementType_XMLDataCreate(AbstractContentElementType_XMLData{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *MedicalAlertMessageType) Set(key string, value interface{}) *MedicalAlertMessageType {
        if n == nil {
                n = MedicalAlertMessageTypeCreate(MedicalAlertMessageType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Severity":
                      n.Severity = StringCreate(value.(string))
        }
        return n
}

func (n *StaffActivityExtensionType) Set(key string, value interface{}) *StaffActivityExtensionType {
        if n == nil {
                n = StaffActivityExtensionTypeCreate(StaffActivityExtensionType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsStaffActivityType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStaffActivityType_values")
      }

                      n.Code = ((*AUCodeSetsStaffActivityType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *StudentSchoolEnrollment_Counselor) Set(key string, value interface{}) *StudentSchoolEnrollment_Counselor {
        if n == nil {
                n = StudentSchoolEnrollment_CounselorCreate(StudentSchoolEnrollment_Counselor{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *Invoice) Set(key string, value interface{}) *Invoice {
        if n == nil {
                n = InvoiceCreate(Invoice{})
        }
        switch key {
    case "ItemDetail":
                      n.ItemDetail = StringCreate(value.(string))
    case "Voluntary":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.Voluntary = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "BilledAmount":
                      n.BilledAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "BillingDate":
                      n.BillingDate = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AccountCodeList":
                      n.AccountCodeList = AccountCodeListTypeCreate(value.(AccountCodeListType))
    case "CreatedBy":
                      n.CreatedBy = StringCreate(value.(string))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Ledger":
                      n.Ledger = StringCreate(value.(string))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "AccountingPeriod":
    
                      n.AccountingPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "FinancialAccountRefIdList":
                      n.FinancialAccountRefIdList = FinancialAccountRefIdListTypeCreate(value.(FinancialAccountRefIdListType))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "InvoicedEntity":
                      n.InvoicedEntity = Invoice_InvoicedEntityCreate(value.(Invoice_InvoicedEntity))
    case "PurchasingItems":
                      n.PurchasingItems = PurchasingItemsTypeCreate(value.(PurchasingItemsType))
    case "RelatedPurchaseOrderRefId":
                      n.RelatedPurchaseOrderRefId = StringCreate(value.(string))
    case "FormNumber":
    
                      n.FormNumber = ((*LocalIdType)(StringCreate(value.(string))))
    case "ApprovedBy":
                      n.ApprovedBy = StringCreate(value.(string))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TaxType":
                      n.TaxType = StringCreate(value.(string))
    case "NetAmount":
                      n.NetAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
        }
        return n
}

func (n *CatchmentStatusContainerType) Set(key string, value interface{}) *CatchmentStatusContainerType {
        if n == nil {
                n = CatchmentStatusContainerTypeCreate(CatchmentStatusContainerType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsPublicSchoolCatchmentStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPublicSchoolCatchmentStatusType_values")
      }

                      n.Code = ((*AUCodeSetsPublicSchoolCatchmentStatusType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *StudentParticipation_ManagingSchool) Set(key string, value interface{}) *StudentParticipation_ManagingSchool {
        if n == nil {
                n = StudentParticipation_ManagingSchoolCreate(StudentParticipation_ManagingSchool{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *SubjectAreaType) Set(key string, value interface{}) *SubjectAreaType {
        if n == nil {
                n = SubjectAreaTypeCreate(SubjectAreaType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
                      n.Code = StringCreate(value.(string))
        }
        return n
}

func (n *StudentParticipation) Set(key string, value interface{}) *StudentParticipation {
        if n == nil {
                n = StudentParticipationCreate(StudentParticipation{})
        }
        switch key {
    case "ReferralDate":
                      n.ReferralDate = StringCreate(value.(string))
    case "ProgramPlanEffectiveDate":
                      n.ProgramPlanEffectiveDate = StringCreate(value.(string))
    case "ProgramPlacementDate":
                      n.ProgramPlacementDate = StringCreate(value.(string))
    case "ProgramEligibilityDate":
                      n.ProgramEligibilityDate = StringCreate(value.(string))
    case "GiftedEligibilityCriteria":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.GiftedEligibilityCriteria = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ManagingSchool":
                      n.ManagingSchool = StudentParticipation_ManagingSchoolCreate(value.(StudentParticipation_ManagingSchool))
    case "ProgramFundingSources":
                      n.ProgramFundingSources = ProgramFundingSourcesTypeCreate(value.(ProgramFundingSourcesType))
    case "EvaluationParentalConsentDate":
                      n.EvaluationParentalConsentDate = StringCreate(value.(string))
    case "EntryPerson":
                      n.EntryPerson = StringCreate(value.(string))
    case "StudentParticipationAsOfDate":
                      n.StudentParticipationAsOfDate = StringCreate(value.(string))
    case "ExtendedSchoolYear":
                      n.ExtendedSchoolYear = BoolCreate(value.(bool))
    case "ExtensionComments":
                      n.ExtensionComments = StringCreate(value.(string))
    case "ProgramAvailability":
                      n.ProgramAvailability = ProgramAvailabilityTypeCreate(value.(ProgramAvailabilityType))
    case "ParticipationContact":
                      n.ParticipationContact = StringCreate(value.(string))
    case "ProgramType":
      present := false
  for _, b := range AUCodeSetsStudentFamilyProgramTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStudentFamilyProgramTypeType_values")
      }

                      n.ProgramType = ((*AUCodeSetsStudentFamilyProgramTypeType)(StringCreate(value.(string))))
    case "ReferralSource":
                      n.ReferralSource = ReferralSourceTypeCreate(value.(ReferralSourceType))
    case "NOREPDate":
                      n.NOREPDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "EvaluationDate":
                      n.EvaluationDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "ProgramPlanDate":
                      n.ProgramPlanDate = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ReevaluationDate":
                      n.ReevaluationDate = StringCreate(value.(string))
    case "ExtendedDay":
                      n.ExtendedDay = BoolCreate(value.(bool))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StudentSpecialEducationFTE":
                      n.StudentSpecialEducationFTE = FloatCreate(value.(float64))
    case "PlacementParentalConsentDate":
                      n.PlacementParentalConsentDate = StringCreate(value.(string))
    case "EvaluationExtensionDate":
                      n.EvaluationExtensionDate = StringCreate(value.(string))
    case "ProgramStatus":
                      n.ProgramStatus = ProgramStatusTypeCreate(value.(ProgramStatusType))
        }
        return n
}

func (n *AddressCollectionStudentType) Set(key string, value interface{}) *AddressCollectionStudentType {
        if n == nil {
                n = AddressCollectionStudentTypeCreate(AddressCollectionStudentType{})
        }
        switch key {
    case "Parent1":
                      n.Parent1 = AGParentTypeCreate(value.(AGParentType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentAddress":
                      n.StudentAddress = AddressTypeCreate(value.(AddressType))
    case "Parent2":
                      n.Parent2 = AGParentTypeCreate(value.(AGParentType))
    case "ReportingParent2":
                      n.ReportingParent2 = StringCreate(value.(string))
    case "BoardingStatus":
      present := false
  for _, b := range AUCodeSetsBoardingType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsBoardingType_values")
      }

                      n.BoardingStatus = ((*AUCodeSetsBoardingType)(StringCreate(value.(string))))
    case "EducationLevel":
      present := false
  for _, b := range AUCodeSetsEducationLevelType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEducationLevelType_values")
      }

                      n.EducationLevel = ((*AUCodeSetsEducationLevelType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ReferralSourceType) Set(key string, value interface{}) *ReferralSourceType {
        if n == nil {
                n = ReferralSourceTypeCreate(ReferralSourceType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSets0792IdentificationProcedureType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSets0792IdentificationProcedureType_values")
      }

                      n.Code = ((*AUCodeSets0792IdentificationProcedureType)(StringCreate(value.(string))))
        }
        return n
}

func (n *StudentSubjectChoiceType) Set(key string, value interface{}) *StudentSubjectChoiceType {
        if n == nil {
                n = StudentSubjectChoiceTypeCreate(StudentSubjectChoiceType{})
        }
        switch key {
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
    case "OtherSchoolLocalId":
    
                      n.OtherSchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudyDescription":
                      n.StudyDescription = SubjectAreaTypeCreate(value.(SubjectAreaType))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *LifeCycleType) Set(key string, value interface{}) *LifeCycleType {
        if n == nil {
                n = LifeCycleTypeCreate(LifeCycleType{})
        }
        switch key {
    case "Created":
                      n.Created = LifeCycleType_CreatedCreate(value.(LifeCycleType_Created))
    case "ModificationHistory":
                      n.ModificationHistory = LifeCycleType_ModificationHistoryCreate(value.(LifeCycleType_ModificationHistory))
    case "TimeElements":
                      n.TimeElements = LifeCycleType_TimeElementsCreate(value.(LifeCycleType_TimeElements))
        }
        return n
}

func (n *DomainScoreType) Set(key string, value interface{}) *DomainScoreType {
        if n == nil {
                n = DomainScoreTypeCreate(DomainScoreType{})
        }
        switch key {
    case "RawScore":
                      n.RawScore = FloatCreate(value.(float64))
    case "ScaledScoreLogitValue":
                      n.ScaledScoreLogitValue = FloatCreate(value.(float64))
    case "StudentProficiency":
                      n.StudentProficiency = StringCreate(value.(string))
    case "ScaledScoreLogitStandardError":
                      n.ScaledScoreLogitStandardError = FloatCreate(value.(float64))
    case "ScaledScoreValue":
                      n.ScaledScoreValue = FloatCreate(value.(float64))
    case "PlausibleScaledValueList":
                      n.PlausibleScaledValueList = PlausibleScaledValueListTypeCreate(value.(PlausibleScaledValueListType))
    case "StudentDomainBand":
                      n.StudentDomainBand = IntCreate(value.(int))
    case "ScaledScoreStandardError":
                      n.ScaledScoreStandardError = FloatCreate(value.(float64))
        }
        return n
}

func (n *LocationType_LocationRefId) Set(key string, value interface{}) *LocationType_LocationRefId {
        if n == nil {
                n = LocationType_LocationRefIdCreate(LocationType_LocationRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *CampusContainerType) Set(key string, value interface{}) *CampusContainerType {
        if n == nil {
                n = CampusContainerTypeCreate(CampusContainerType{})
        }
        switch key {
    case "CampusType":
      present := false
  for _, b := range AUCodeSetsSchoolLevelType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolLevelType_values")
      }

                      n.CampusType = ((*AUCodeSetsSchoolLevelType)(StringCreate(value.(string))))
    case "SchoolCampusId":
                      n.SchoolCampusId = StringCreate(value.(string))
    case "AdminStatus":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AdminStatus = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ParentSchoolId":
                      n.ParentSchoolId = StringCreate(value.(string))
        }
        return n
}

func (n *AggregateStatisticInfo) Set(key string, value interface{}) *AggregateStatisticInfo {
        if n == nil {
                n = AggregateStatisticInfoCreate(AggregateStatisticInfo{})
        }
        switch key {
    case "ExclusionRules":
                      n.ExclusionRules = ExclusionRulesTypeCreate(value.(ExclusionRulesType))
    case "Source":
                      n.Source = StringCreate(value.(string))
    case "EffectiveDate":
                      n.EffectiveDate = StringCreate(value.(string))
    case "ExpirationDate":
                      n.ExpirationDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ApprovalDate":
                      n.ApprovalDate = StringCreate(value.(string))
    case "Location":
                      n.Location = LocationTypeCreate(value.(LocationType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Measure":
                      n.Measure = StringCreate(value.(string))
    case "StatisticName":
                      n.StatisticName = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "DiscontinueDate":
                      n.DiscontinueDate = StringCreate(value.(string))
    case "CalculationRule":
                      n.CalculationRule = AggregateStatisticInfo_CalculationRuleCreate(value.(AggregateStatisticInfo_CalculationRule))
        }
        return n
}

func (n *CollectionStatus) Set(key string, value interface{}) *CollectionStatus {
        if n == nil {
                n = CollectionStatusCreate(CollectionStatus{})
        }
        switch key {
    case "AGCollection":
                      n.AGCollection = StringCreate(value.(string))
    case "AGReportingObjectResponseList":
                      n.AGReportingObjectResponseList = AGReportingObjectResponseListTypeCreate(value.(AGReportingObjectResponseListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ReportingAuthoritySystem":
                      n.ReportingAuthoritySystem = StringCreate(value.(string))
    case "SubmissionTimestamp":
                      n.SubmissionTimestamp = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SubmittedBy":
                      n.SubmittedBy = StringCreate(value.(string))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "CollectionYear":
    
                      n.CollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "ReportingAuthority":
                      n.ReportingAuthority = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *PersonalisedPlan) Set(key string, value interface{}) *PersonalisedPlan {
        if n == nil {
                n = PersonalisedPlanCreate(PersonalisedPlan{})
        }
        switch key {
    case "PersonalisedPlanStartDate":
                      n.PersonalisedPlanStartDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PersonalisedPlanReviewDate":
                      n.PersonalisedPlanReviewDate = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PersonalisedPlanCategory":
      present := false
  for _, b := range AUCodeSetsPersonalisedPlanType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPersonalisedPlanType_values")
      }

                      n.PersonalisedPlanCategory = ((*AUCodeSetsPersonalisedPlanType)(StringCreate(value.(string))))
    case "PersonalisedPlanEndDate":
                      n.PersonalisedPlanEndDate = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PersonalisedPlanNotes":
                      n.PersonalisedPlanNotes = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "AssociatedAttachment":
                      n.AssociatedAttachment = StringCreate(value.(string))
        }
        return n
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) Set(key string, value interface{}) *SIF_ExtendedElementsType_SIF_ExtendedElement {
        if n == nil {
                n = SIF_ExtendedElementsType_SIF_ExtendedElementCreate(SIF_ExtendedElementsType_SIF_ExtendedElement{})
        }
        switch key {
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "SIF_Action":
                      n.SIF_Action = StringCreate(value.(string))
    case "Value":
    
                      n.Value = ((*ExtendedContentType)(StringCreate(value.(string))))
        }
        return n
}

func (n *FinancialQuestionnaireCollection) Set(key string, value interface{}) *FinancialQuestionnaireCollection {
        if n == nil {
                n = FinancialQuestionnaireCollectionCreate(FinancialQuestionnaireCollection{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "FQYear":
    
                      n.FQYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "FQReportingList":
                      n.FQReportingList = FQReportingListTypeCreate(value.(FQReportingListType))
        }
        return n
}

func (n *JournalAdjustmentType) Set(key string, value interface{}) *JournalAdjustmentType {
        if n == nil {
                n = JournalAdjustmentTypeCreate(JournalAdjustmentType{})
        }
        switch key {
    case "DebitFinancialAccountRefId":
                      n.DebitFinancialAccountRefId = StringCreate(value.(string))
    case "DebitAccountCode":
                      n.DebitAccountCode = StringCreate(value.(string))
    case "GSTCodeOriginal":
                      n.GSTCodeOriginal = StringCreate(value.(string))
    case "CreditFinancialAccountRefId":
                      n.CreditFinancialAccountRefId = StringCreate(value.(string))
    case "LineAdjustmentAmount":
                      n.LineAdjustmentAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "CreditAccountCode":
                      n.CreditAccountCode = StringCreate(value.(string))
    case "GSTCodeReplacement":
                      n.GSTCodeReplacement = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingPersonLink_PersonRefId) Set(key string, value interface{}) *WellbeingPersonLink_PersonRefId {
        if n == nil {
                n = WellbeingPersonLink_PersonRefIdCreate(WellbeingPersonLink_PersonRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *HoldInfoType) Set(key string, value interface{}) *HoldInfoType {
        if n == nil {
                n = HoldInfoTypeCreate(HoldInfoType{})
        }
        switch key {
    case "MadeAvailable":
                      n.MadeAvailable = StringCreate(value.(string))
    case "Expires":
                      n.Expires = StringCreate(value.(string))
    case "DateNeeded":
                      n.DateNeeded = StringCreate(value.(string))
    case "DatePlaced":
                      n.DatePlaced = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "ReservationExpiry":
                      n.ReservationExpiry = StringCreate(value.(string))
        }
        return n
}

func (n *TimeTable) Set(key string, value interface{}) *TimeTable {
        if n == nil {
                n = TimeTableCreate(TimeTable{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TimeTableDayList":
                      n.TimeTableDayList = TimeTableDayListTypeCreate(value.(TimeTableDayListType))
    case "TeachingPeriodsPerDay":
                      n.TeachingPeriodsPerDay = IntCreate(value.(int))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableCreationDate":
                      n.TimeTableCreationDate = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "PeriodsPerDay":
                      n.PeriodsPerDay = IntCreate(value.(int))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "DaysPerCycle":
                      n.DaysPerCycle = IntCreate(value.(int))
    case "Title":
                      n.Title = StringCreate(value.(string))
        }
        return n
}

func (n *StudentGrade) Set(key string, value interface{}) *StudentGrade {
        if n == nil {
                n = StudentGradeCreate(StudentGrade{})
        }
        switch key {
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Markers":
                      n.Markers = StudentGradeMarkersListTypeCreate(value.(StudentGradeMarkersListType))
    case "LearningArea":
                      n.LearningArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TeachingGroupShortName":
                      n.TeachingGroupShortName = StringCreate(value.(string))
    case "GradingScoreList":
                      n.GradingScoreList = GradingScoreListTypeCreate(value.(GradingScoreListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Grade":
                      n.Grade = GradeTypeCreate(value.(GradeType))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "TeacherJudgement":
                      n.TeacherJudgement = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
        }
        return n
}

func (n *StandardsSettingBodyType) Set(key string, value interface{}) *StandardsSettingBodyType {
        if n == nil {
                n = StandardsSettingBodyTypeCreate(StandardsSettingBodyType{})
        }
        switch key {
    case "Country":
    
                      n.Country = ((*CountryType)(StringCreate(value.(string))))
    case "StateProvince":
    
                      n.StateProvince = ((*StateProvinceType)(StringCreate(value.(string))))
    case "SettingBodyName":
                      n.SettingBodyName = StringCreate(value.(string))
        }
        return n
}

func (n *NameType) Set(key string, value interface{}) *NameType {
        if n == nil {
                n = NameTypeCreate(NameType{})
        }
        switch key {
    case "Type":
      present := false
  for _, b := range AUCodeSetsNameUsageTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNameUsageTypeType_values")
      }

                      n.Type = ((*AUCodeSetsNameUsageTypeType)(StringCreate(value.(string))))
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "FamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PreferredFamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PreferredFamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
        }
        return n
}

func (n *CollectionRound) Set(key string, value interface{}) *CollectionRound {
        if n == nil {
                n = CollectionRoundCreate(CollectionRound{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CollectionYear":
    
                      n.CollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "AGRoundList":
                      n.AGRoundList = AGRoundListTypeCreate(value.(AGRoundListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AGCollection":
      present := false
  for _, b := range AUCodeSetsAGCollectionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAGCollectionType_values")
      }

                      n.AGCollection = ((*AUCodeSetsAGCollectionType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ResourceBooking_ResourceRefId) Set(key string, value interface{}) *ResourceBooking_ResourceRefId {
        if n == nil {
                n = ResourceBooking_ResourceRefIdCreate(ResourceBooking_ResourceRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *AbstractContentElementType_BinaryData) Set(key string, value interface{}) *AbstractContentElementType_BinaryData {
        if n == nil {
                n = AbstractContentElementType_BinaryDataCreate(AbstractContentElementType_BinaryData{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
        }
        return n
}

func (n *AdjustmentContainerType) Set(key string, value interface{}) *AdjustmentContainerType {
        if n == nil {
                n = AdjustmentContainerTypeCreate(AdjustmentContainerType{})
        }
        switch key {
    case "PNPCodeList":
                      n.PNPCodeList = PNPCodeListTypeCreate(value.(PNPCodeListType))
    case "BookletType":
                      n.BookletType = StringCreate(value.(string))
        }
        return n
}

func (n *PurchaseOrder) Set(key string, value interface{}) *PurchaseOrder {
        if n == nil {
                n = PurchaseOrderCreate(PurchaseOrder{})
        }
        switch key {
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "PurchasingItems":
                      n.PurchasingItems = PurchasingItemsTypeCreate(value.(PurchasingItemsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TotalAmount":
                      n.TotalAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "OriginalPurchaseOrderRefId":
                      n.OriginalPurchaseOrderRefId = StringCreate(value.(string))
    case "VendorInfoRefId":
                      n.VendorInfoRefId = StringCreate(value.(string))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "UpdateDate":
                      n.UpdateDate = StringCreate(value.(string))
    case "FullyDelivered":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FullyDelivered = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FormNumber":
                      n.FormNumber = StringCreate(value.(string))
    case "EmployeePersonalRefId":
                      n.EmployeePersonalRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        }
        return n
}

func (n *StudentEntryContainerType) Set(key string, value interface{}) *StudentEntryContainerType {
        if n == nil {
                n = StudentEntryContainerTypeCreate(StudentEntryContainerType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsEntryTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEntryTypeType_values")
      }

                      n.Code = ((*AUCodeSetsEntryTypeType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *EvaluationType) Set(key string, value interface{}) *EvaluationType {
        if n == nil {
                n = EvaluationTypeCreate(EvaluationType{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
    case "Date":
                      n.Date = StringCreate(value.(string))
        }
        return n
}

func (n *NAPCodeFrame) Set(key string, value interface{}) *NAPCodeFrame {
        if n == nil {
                n = NAPCodeFrameCreate(NAPCodeFrame{})
        }
        switch key {
    case "TestContent":
                      n.TestContent = NAPTestContentTypeCreate(value.(NAPTestContentType))
    case "TestletList":
                      n.TestletList = NAPCodeFrameTestletListTypeCreate(value.(NAPCodeFrameTestletListType))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        }
        return n
}

func (n *SectionInfo) Set(key string, value interface{}) *SectionInfo {
        if n == nil {
                n = SectionInfoCreate(SectionInfo{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LocationOfInstruction":
                      n.LocationOfInstruction = LocationOfInstructionTypeCreate(value.(LocationOfInstructionType))
    case "CountForAttendance":
                      n.CountForAttendance = StringCreate(value.(string))
    case "SchoolCourseInfoRefId":
                      n.SchoolCourseInfoRefId = StringCreate(value.(string))
    case "SummerSchool":
                      n.SummerSchool = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "MediumOfInstruction":
                      n.MediumOfInstruction = MediumOfInstructionTypeCreate(value.(MediumOfInstructionType))
    case "SchoolCourseInfoOverride":
                      n.SchoolCourseInfoOverride = SchoolCourseInfoOverrideTypeCreate(value.(SchoolCourseInfoOverrideType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SectionCode":
                      n.SectionCode = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LanguageOfInstruction":
                      n.LanguageOfInstruction = LanguageOfInstructionTypeCreate(value.(LanguageOfInstructionType))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CourseSectionCode":
                      n.CourseSectionCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *AttendanceCodeType) Set(key string, value interface{}) *AttendanceCodeType {
        if n == nil {
                n = AttendanceCodeTypeCreate(AttendanceCodeType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsAttendanceCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAttendanceCodeType_values")
      }

                      n.Code = ((*AUCodeSetsAttendanceCodeType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *ChargedLocationInfo) Set(key string, value interface{}) *ChargedLocationInfo {
        if n == nil {
                n = ChargedLocationInfoCreate(ChargedLocationInfo{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LocationType":
                      n.LocationType = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SiteCategory":
                      n.SiteCategory = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "ParentChargedLocationInfoRefId":
                      n.ParentChargedLocationInfoRefId = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        }
        return n
}

func (n *StudentDailyAttendance) Set(key string, value interface{}) *StudentDailyAttendance {
        if n == nil {
                n = StudentDailyAttendanceCreate(StudentDailyAttendance{})
        }
        switch key {
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "AbsenceValue":
                      n.AbsenceValue = FloatCreate(value.(float64))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DayValue":
      present := false
  for _, b := range AUCodeSetsDayValueCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsDayValueCodeType_values")
      }

                      n.DayValue = ((*AUCodeSetsDayValueCodeType)(StringCreate(value.(string))))
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "AttendanceStatus":
      present := false
  for _, b := range AUCodeSetsAttendanceStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAttendanceStatusType_values")
      }

                      n.AttendanceStatus = ((*AUCodeSetsAttendanceStatusType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
        }
        return n
}

func (n *DomainProficiencyContainerType) Set(key string, value interface{}) *DomainProficiencyContainerType {
        if n == nil {
                n = DomainProficiencyContainerTypeCreate(DomainProficiencyContainerType{})
        }
        switch key {
    case "Level1Upper":
                      n.Level1Upper = FloatCreate(value.(float64))
    case "Level2Upper":
                      n.Level2Upper = FloatCreate(value.(float64))
    case "Level2Lower":
                      n.Level2Lower = FloatCreate(value.(float64))
    case "Level1Lower":
                      n.Level1Lower = FloatCreate(value.(float64))
    case "Level4Upper":
                      n.Level4Upper = FloatCreate(value.(float64))
    case "Level4Lower":
                      n.Level4Lower = FloatCreate(value.(float64))
    case "Level3Upper":
                      n.Level3Upper = FloatCreate(value.(float64))
    case "Level3Lower":
                      n.Level3Lower = FloatCreate(value.(float64))
        }
        return n
}

func (n *DwellingArrangementType) Set(key string, value interface{}) *DwellingArrangementType {
        if n == nil {
                n = DwellingArrangementTypeCreate(DwellingArrangementType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsDwellingArrangementType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsDwellingArrangementType_values")
      }

                      n.Code = ((*AUCodeSetsDwellingArrangementType)(StringCreate(value.(string))))
        }
        return n
}

func (n *AlertMessageType) Set(key string, value interface{}) *AlertMessageType {
        if n == nil {
                n = AlertMessageTypeCreate(AlertMessageType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *Activity) Set(key string, value interface{}) *Activity {
        if n == nil {
                n = ActivityCreate(Activity{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "EssentialMaterials":
                      n.EssentialMaterials = EssentialMaterialsTypeCreate(value.(EssentialMaterialsType))
    case "LearningObjectives":
                      n.LearningObjectives = LearningObjectivesTypeCreate(value.(LearningObjectivesType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ActivityWeight":
                      n.ActivityWeight = FloatCreate(value.(float64))
    case "Preamble":
                      n.Preamble = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TechnicalRequirements":
                      n.TechnicalRequirements = TechnicalRequirementsTypeCreate(value.(TechnicalRequirementsType))
    case "Prerequisites":
                      n.Prerequisites = PrerequisitesTypeCreate(value.(PrerequisitesType))
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
    case "SoftwareRequirementList":
                      n.SoftwareRequirementList = SoftwareRequirementListTypeCreate(value.(SoftwareRequirementListType))
    case "Evaluation":
                      n.Evaluation = Activity_EvaluationCreate(value.(Activity_Evaluation))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "MaxAttemptsAllowed":
                      n.MaxAttemptsAllowed = IntCreate(value.(int))
    case "SourceObjects":
                      n.SourceObjects = SourceObjectsTypeCreate(value.(SourceObjectsType))
    case "Students":
                      n.Students = StudentsTypeCreate(value.(StudentsType))
    case "ActivityTime":
                      n.ActivityTime = ActivityTimeTypeCreate(value.(ActivityTimeType))
    case "LearningResources":
                      n.LearningResources = LearningResourcesTypeCreate(value.(LearningResourcesType))
    case "AssessmentRefId":
                      n.AssessmentRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Points":
                      n.Points = IntCreate(value.(int))
    case "LearningStandards":
                      n.LearningStandards = LearningStandardsTypeCreate(value.(LearningStandardsType))
        }
        return n
}

func (n *WellbeingDocumentType) Set(key string, value interface{}) *WellbeingDocumentType {
        if n == nil {
                n = WellbeingDocumentTypeCreate(WellbeingDocumentType{})
        }
        switch key {
    case "Location":
                      n.Location = StringCreate(value.(string))
    case "DocumentType":
                      n.DocumentType = StringCreate(value.(string))
    case "Sensitivity":
                      n.Sensitivity = StringCreate(value.(string))
    case "DocumentDescription":
                      n.DocumentDescription = StringCreate(value.(string))
    case "URL":
                      n.URL = StringCreate(value.(string))
    case "DocumentReviewDate":
                      n.DocumentReviewDate = StringCreate(value.(string))
        }
        return n
}

func (n *StudentAttendanceTimeList) Set(key string, value interface{}) *StudentAttendanceTimeList {
        if n == nil {
                n = StudentAttendanceTimeListCreate(StudentAttendanceTimeList{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AttendanceTimes":
                      n.AttendanceTimes = AttendanceTimesTypeCreate(value.(AttendanceTimesType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PeriodAttendances":
                      n.PeriodAttendances = PeriodAttendancesTypeCreate(value.(PeriodAttendancesType))
        }
        return n
}

func (n *StudentSchoolEnrollment_Advisor) Set(key string, value interface{}) *StudentSchoolEnrollment_Advisor {
        if n == nil {
                n = StudentSchoolEnrollment_AdvisorCreate(StudentSchoolEnrollment_Advisor{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *NAPTestItem2Type) Set(key string, value interface{}) *NAPTestItem2Type {
        if n == nil {
                n = NAPTestItem2TypeCreate(NAPTestItem2Type{})
        }
        switch key {
    case "SequenceNumber":
                      n.SequenceNumber = IntCreate(value.(int))
    case "TestItemLocalId":
    
                      n.TestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestItemRefId":
                      n.TestItemRefId = StringCreate(value.(string))
        }
        return n
}

func (n *ScheduledActivityOverrideType) Set(key string, value interface{}) *ScheduledActivityOverrideType {
        if n == nil {
                n = ScheduledActivityOverrideTypeCreate(ScheduledActivityOverrideType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "DateOfOverride":
                      n.DateOfOverride = StringCreate(value.(string))
        }
        return n
}

func (n *VisaSubClassType) Set(key string, value interface{}) *VisaSubClassType {
        if n == nil {
                n = VisaSubClassTypeCreate(VisaSubClassType{})
        }
        switch key {
    case "Code":
    
                      n.Code = ((*VisaSubClassCodeType)(StringCreate(value.(string))))
    case "ATEStartDate":
                      n.ATEStartDate = StringCreate(value.(string))
    case "VisaStatisticalCode":
                      n.VisaStatisticalCode = StringCreate(value.(string))
    case "VisaExpiryDate":
                      n.VisaExpiryDate = StringCreate(value.(string))
    case "ATEExpiryDate":
                      n.ATEExpiryDate = StringCreate(value.(string))
        }
        return n
}

func (n *AGParentType) Set(key string, value interface{}) *AGParentType {
        if n == nil {
                n = AGParentTypeCreate(AGParentType{})
        }
        switch key {
    case "ParentAddress":
                      n.ParentAddress = AddressTypeCreate(value.(AddressType))
    case "ParentName":
                      n.ParentName = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "AddressSameAsStudent":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AddressSameAsStudent = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
        }
        return n
}

func (n *FollowUpActionType) Set(key string, value interface{}) *FollowUpActionType {
        if n == nil {
                n = FollowUpActionTypeCreate(FollowUpActionType{})
        }
        switch key {
    case "FollowUpDetails":
                      n.FollowUpDetails = StringCreate(value.(string))
    case "FollowUpActionCategory":
                      n.FollowUpActionCategory = StringCreate(value.(string))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
        }
        return n
}

func (n *SIF_MetadataType) Set(key string, value interface{}) *SIF_MetadataType {
        if n == nil {
                n = SIF_MetadataTypeCreate(SIF_MetadataType{})
        }
        switch key {
    case "TimeElements":
                      n.TimeElements = SIF_MetadataType_TimeElementsCreate(value.(SIF_MetadataType_TimeElements))
    case "LifeCycle":
                      n.LifeCycle = LifeCycleTypeCreate(value.(LifeCycleType))
    case "EducationFilter":
                      n.EducationFilter = EducationFilterTypeCreate(value.(EducationFilterType))
        }
        return n
}

func (n *CensusStaffType) Set(key string, value interface{}) *CensusStaffType {
        if n == nil {
                n = CensusStaffTypeCreate(CensusStaffType{})
        }
        switch key {
    case "SecondaryFTE":
                      n.SecondaryFTE = FloatCreate(value.(float64))
    case "Headcount":
                      n.Headcount = IntCreate(value.(int))
    case "CohortIndigenousType":
                      n.CohortIndigenousType = StringCreate(value.(string))
    case "StaffCohortId":
    
                      n.StaffCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "JobFTE":
                      n.JobFTE = FloatCreate(value.(float64))
    case "StaffActivity":
                      n.StaffActivity = StaffActivityExtensionTypeCreate(value.(StaffActivityExtensionType))
    case "PrimaryFTE":
                      n.PrimaryFTE = FloatCreate(value.(float64))
        }
        return n
}

func (n *GradeType) Set(key string, value interface{}) *GradeType {
        if n == nil {
                n = GradeTypeCreate(GradeType{})
        }
        switch key {
    case "MarkInfoRefId":
                      n.MarkInfoRefId = StringCreate(value.(string))
    case "Letter":
                      n.Letter = StringCreate(value.(string))
    case "Percentage":
                      n.Percentage = FloatCreate(value.(float64))
    case "Narrative":
                      n.Narrative = StringCreate(value.(string))
    case "Numeric":
                      n.Numeric = FloatCreate(value.(float64))
        }
        return n
}

func (n *GradingAssignment) Set(key string, value interface{}) *GradingAssignment {
        if n == nil {
                n = GradingAssignmentCreate(GradingAssignment{})
        }
        switch key {
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "RubricScoringGuide":
                      n.RubricScoringGuide = GenericRubricTypeCreate(value.(GenericRubricType))
    case "PrerequisiteList":
                      n.PrerequisiteList = PrerequisitesTypeCreate(value.(PrerequisitesType))
    case "DetailedDescriptionBinary":
                      n.DetailedDescriptionBinary = StringCreate(value.(string))
    case "DetailedDescriptionURL":
                      n.DetailedDescriptionURL = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "MaxAttemptsAllowed":
                      n.MaxAttemptsAllowed = IntCreate(value.(int))
    case "LevelAssessed":
                      n.LevelAssessed = StringCreate(value.(string))
    case "AssessmentType":
                      n.AssessmentType = StringCreate(value.(string))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "StudentPersonalRefIdList":
                      n.StudentPersonalRefIdList = StudentsTypeCreate(value.(StudentsType))
    case "PointsPossible":
                      n.PointsPossible = IntCreate(value.(int))
    case "AssignmentPurpose":
                      n.AssignmentPurpose = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "GradingCategory":
                      n.GradingCategory = StringCreate(value.(string))
    case "Weight":
                      n.Weight = FloatCreate(value.(float64))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "SubAssignmentList":
                      n.SubAssignmentList = AssignmentListTypeCreate(value.(AssignmentListType))
    case "CreateDate":
                      n.CreateDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
        }
        return n
}

func (n *FQContextualQuestionType) Set(key string, value interface{}) *FQContextualQuestionType {
        if n == nil {
                n = FQContextualQuestionTypeCreate(FQContextualQuestionType{})
        }
        switch key {
    case "FQContext":
                      n.FQContext = StringCreate(value.(string))
    case "FQAnswer":
                      n.FQAnswer = StringCreate(value.(string))
        }
        return n
}

func (n *PublishingPermissionType) Set(key string, value interface{}) *PublishingPermissionType {
        if n == nil {
                n = PublishingPermissionTypeCreate(PublishingPermissionType{})
        }
        switch key {
    case "PermissionValue":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PermissionValue = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PermissionCategory":
      present := false
  for _, b := range AUCodeSetsPermissionCategoryCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPermissionCategoryCodeType_values")
      }

                      n.PermissionCategory = ((*AUCodeSetsPermissionCategoryCodeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *WellbeingResponse) Set(key string, value interface{}) *WellbeingResponse {
        if n == nil {
                n = WellbeingResponseCreate(WellbeingResponse{})
        }
        switch key {
    case "OtherResponse":
                      n.OtherResponse = OtherWellbeingResponseContainerTypeCreate(value.(OtherWellbeingResponseContainerType))
    case "WellbeingResponseCategory":
      present := false
  for _, b := range AUCodeSetsWellbeingResponseCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingResponseCategoryType_values")
      }

                      n.WellbeingResponseCategory = ((*AUCodeSetsWellbeingResponseCategoryType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PersonInvolvementList":
                      n.PersonInvolvementList = PersonInvolvementListTypeCreate(value.(PersonInvolvementListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "WellbeingResponseEndDate":
                      n.WellbeingResponseEndDate = StringCreate(value.(string))
    case "Detention":
                      n.Detention = DetentionContainerTypeCreate(value.(DetentionContainerType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "WellbeingResponseStartDate":
                      n.WellbeingResponseStartDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "WellbeingResponseNotes":
                      n.WellbeingResponseNotes = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "Suspension":
                      n.Suspension = SuspensionContainerTypeCreate(value.(SuspensionContainerType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Award":
                      n.Award = AwardContainerTypeCreate(value.(AwardContainerType))
    case "PlanRequired":
                      n.PlanRequired = PlanRequiredContainerTypeCreate(value.(PlanRequiredContainerType))
        }
        return n
}

func (n *AddressCollectionReportingType) Set(key string, value interface{}) *AddressCollectionReportingType {
        if n == nil {
                n = AddressCollectionReportingTypeCreate(AddressCollectionReportingType{})
        }
        switch key {
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "AddressCollectionStudentList":
                      n.AddressCollectionStudentList = AddressCollectionStudentListTypeCreate(value.(AddressCollectionStudentListType))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AGContextualQuestionList":
                      n.AGContextualQuestionList = AGContextualQuestionListTypeCreate(value.(AGContextualQuestionListType))
        }
        return n
}

func (n *AbstractContentPackageType_Reference) Set(key string, value interface{}) *AbstractContentPackageType_Reference {
        if n == nil {
                n = AbstractContentPackageType_ReferenceCreate(AbstractContentPackageType_Reference{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "URL":
                      n.URL = StringCreate(value.(string))
        }
        return n
}

func (n *DetentionContainerType) Set(key string, value interface{}) *DetentionContainerType {
        if n == nil {
                n = DetentionContainerTypeCreate(DetentionContainerType{})
        }
        switch key {
    case "DetentionNotes":
                      n.DetentionNotes = StringCreate(value.(string))
    case "DetentionCategory":
      present := false
  for _, b := range AUCodeSetsDetentionCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsDetentionCategoryType_values")
      }

                      n.DetentionCategory = ((*AUCodeSetsDetentionCategoryType)(StringCreate(value.(string))))
    case "DetentionLocation":
                      n.DetentionLocation = StringCreate(value.(string))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
    case "DetentionDate":
                      n.DetentionDate = StringCreate(value.(string))
        }
        return n
}

func (n *TeachingGroupPeriodType) Set(key string, value interface{}) *TeachingGroupPeriodType {
        if n == nil {
                n = TeachingGroupPeriodTypeCreate(TeachingGroupPeriodType{})
        }
        switch key {
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *CensusReportingType) Set(key string, value interface{}) *CensusReportingType {
        if n == nil {
                n = CensusReportingTypeCreate(CensusReportingType{})
        }
        switch key {
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "CensusStaffList":
                      n.CensusStaffList = CensusStaffListTypeCreate(value.(CensusStaffListType))
    case "CensusStudentList":
                      n.CensusStudentList = CensusStudentListTypeCreate(value.(CensusStudentListType))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
        }
        return n
}

func (n *AggregateCharacteristicInfo) Set(key string, value interface{}) *AggregateCharacteristicInfo {
        if n == nil {
                n = AggregateCharacteristicInfoCreate(AggregateCharacteristicInfo{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ElementName":
                      n.ElementName = StringCreate(value.(string))
    case "Definition":
                      n.Definition = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *AttendanceTimeType) Set(key string, value interface{}) *AttendanceTimeType {
        if n == nil {
                n = AttendanceTimeTypeCreate(AttendanceTimeType{})
        }
        switch key {
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "EndTime":
                      n.EndTime = StringCreate(value.(string))
    case "AttendanceType":
                      n.AttendanceType = StringCreate(value.(string))
    case "AttendanceStatus":
      present := false
  for _, b := range AUCodeSetsAttendanceStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAttendanceStatusType_values")
      }

                      n.AttendanceStatus = ((*AUCodeSetsAttendanceStatusType)(StringCreate(value.(string))))
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "DurationValue":
                      n.DurationValue = FloatCreate(value.(float64))
        }
        return n
}

func (n *SoftwareVendorInfoContainerType) Set(key string, value interface{}) *SoftwareVendorInfoContainerType {
        if n == nil {
                n = SoftwareVendorInfoContainerTypeCreate(SoftwareVendorInfoContainerType{})
        }
        switch key {
    case "SoftwareVersion":
                      n.SoftwareVersion = StringCreate(value.(string))
    case "SoftwareProduct":
                      n.SoftwareProduct = StringCreate(value.(string))
        }
        return n
}

func (n *LEAInfo) Set(key string, value interface{}) *LEAInfo {
        if n == nil {
                n = LEAInfoCreate(LEAInfo{})
        }
        switch key {
    case "SLA":
      present := false
  for _, b := range AUCodeSetsAustralianStandardGeographicalClassificationASGCType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardGeographicalClassificationASGCType_values")
      }

                      n.SLA = ((*AUCodeSetsAustralianStandardGeographicalClassificationASGCType)(StringCreate(value.(string))))
    case "EducationAgencyType":
                      n.EducationAgencyType = AgencyTypeCreate(value.(AgencyType))
    case "LEAName":
                      n.LEAName = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "JurisdictionLowerHouse":
                      n.JurisdictionLowerHouse = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "OperationalStatus":
    
                      n.OperationalStatus = ((*OperationalStatusType)(StringCreate(value.(string))))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LEAURL":
                      n.LEAURL = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LEAContactList":
                      n.LEAContactList = LEAContactListTypeCreate(value.(LEAContactListType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
        }
        return n
}

func (n *AggregateStatisticFact) Set(key string, value interface{}) *AggregateStatisticFact {
        if n == nil {
                n = AggregateStatisticFactCreate(AggregateStatisticFact{})
        }
        switch key {
    case "Value":
                      n.Value = FloatCreate(value.(float64))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Excluded":
                      n.Excluded = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AggregateStatisticInfoRefId":
                      n.AggregateStatisticInfoRefId = StringCreate(value.(string))
    case "Characteristics":
                      n.Characteristics = CharacteristicsTypeCreate(value.(CharacteristicsType))
        }
        return n
}

func (n *LocationOfInstructionType) Set(key string, value interface{}) *LocationOfInstructionType {
        if n == nil {
                n = LocationOfInstructionTypeCreate(LocationOfInstructionType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsReceivingLocationOfInstructionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsReceivingLocationOfInstructionType_values")
      }

                      n.Code = ((*AUCodeSetsReceivingLocationOfInstructionType)(StringCreate(value.(string))))
        }
        return n
}

func (n *GradingAssignmentScore) Set(key string, value interface{}) *GradingAssignmentScore {
        if n == nil {
                n = GradingAssignmentScoreCreate(GradingAssignmentScore{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "GradingAssignmentRefId":
                      n.GradingAssignmentRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SubscoreList":
                      n.SubscoreList = NAPSubscoreListTypeCreate(value.(NAPSubscoreListType))
    case "StudentPersonalLocalId":
    
                      n.StudentPersonalLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AssignmentScoreIteration":
                      n.AssignmentScoreIteration = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "DateGraded":
                      n.DateGraded = StringCreate(value.(string))
    case "MarkInfoRefId":
                      n.MarkInfoRefId = StringCreate(value.(string))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "TeacherJudgement":
                      n.TeacherJudgement = StringCreate(value.(string))
    case "ExpectedScore":
                      n.ExpectedScore = BoolCreate(value.(bool))
    case "ScorePoints":
                      n.ScorePoints = IntCreate(value.(int))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "ScoreDescription":
                      n.ScoreDescription = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ScorePercent":
                      n.ScorePercent = FloatCreate(value.(float64))
    case "ScoreLetter":
                      n.ScoreLetter = StringCreate(value.(string))
        }
        return n
}

func (n *CreationUserType) Set(key string, value interface{}) *CreationUserType {
        if n == nil {
                n = CreationUserTypeCreate(CreationUserType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "UserId":
                      n.UserId = StringCreate(value.(string))
        }
        return n
}

func (n *LibraryTransactionType) Set(key string, value interface{}) *LibraryTransactionType {
        if n == nil {
                n = LibraryTransactionTypeCreate(LibraryTransactionType{})
        }
        switch key {
    case "FineInfoList":
                      n.FineInfoList = FineInfoListTypeCreate(value.(FineInfoListType))
    case "HoldInfoList":
                      n.HoldInfoList = HoldInfoListTypeCreate(value.(HoldInfoListType))
    case "ItemInfo":
                      n.ItemInfo = LibraryItemInfoTypeCreate(value.(LibraryItemInfoType))
    case "CheckoutInfo":
                      n.CheckoutInfo = CheckoutInfoTypeCreate(value.(CheckoutInfoType))
        }
        return n
}

func (n *PaymentReceipt) Set(key string, value interface{}) *PaymentReceipt {
        if n == nil {
                n = PaymentReceiptCreate(PaymentReceipt{})
        }
        switch key {
    case "DebtorRefId":
                      n.DebtorRefId = StringCreate(value.(string))
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
    case "TransactionNote":
                      n.TransactionNote = StringCreate(value.(string))
    case "AccountingPeriod":
    
                      n.AccountingPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "FinancialAccountRefIdList":
                      n.FinancialAccountRefIdList = FinancialAccountRefIdListTypeCreate(value.(FinancialAccountRefIdListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AccountCodeList":
                      n.AccountCodeList = AccountCodeListTypeCreate(value.(AccountCodeListType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "ChequeNumber":
                      n.ChequeNumber = StringCreate(value.(string))
    case "PaymentReceiptLineList":
                      n.PaymentReceiptLineList = PaymentReceiptLineListTypeCreate(value.(PaymentReceiptLineListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TransactionMethod":
                      n.TransactionMethod = StringCreate(value.(string))
    case "TransactionAmount":
                      n.TransactionAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "ReceivedTransactionId":
                      n.ReceivedTransactionId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TransactionDate":
                      n.TransactionDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "VendorInfoRefId":
                      n.VendorInfoRefId = StringCreate(value.(string))
    case "TransactionType":
                      n.TransactionType = StringCreate(value.(string))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
        }
        return n
}

func (n *TimeTableSubject) Set(key string, value interface{}) *TimeTableSubject {
        if n == nil {
                n = TimeTableSubjectCreate(TimeTableSubject{})
        }
        switch key {
    case "SubjectLongName":
                      n.SubjectLongName = StringCreate(value.(string))
    case "CourseLocalId":
    
                      n.CourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "AcademicYear":
                      n.AcademicYear = YearLevelTypeCreate(value.(YearLevelType))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SubjectShortName":
                      n.SubjectShortName = StringCreate(value.(string))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "ProposedMinClassSize":
                      n.ProposedMinClassSize = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Faculty":
                      n.Faculty = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ProposedMaxClassSize":
                      n.ProposedMaxClassSize = FloatCreate(value.(float64))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AcademicYearRange":
                      n.AcademicYearRange = YearRangeTypeCreate(value.(YearRangeType))
    case "SubjectType":
                      n.SubjectType = StringCreate(value.(string))
        }
        return n
}

func (n *SourceObjectsType_SourceObject) Set(key string, value interface{}) *SourceObjectsType_SourceObject {
        if n == nil {
                n = SourceObjectsType_SourceObjectCreate(SourceObjectsType_SourceObject{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *CalendarSummary) Set(key string, value interface{}) *CalendarSummary {
        if n == nil {
                n = CalendarSummaryCreate(CalendarSummary{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "InstructionalMinutes":
                      n.InstructionalMinutes = IntCreate(value.(int))
    case "GraduationDate":
    
                      n.GraduationDate = ((*GraduationDateType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "LastInstructionDate":
                      n.LastInstructionDate = StringCreate(value.(string))
    case "FirstInstructionDate":
                      n.FirstInstructionDate = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "MinutesPerDay":
                      n.MinutesPerDay = IntCreate(value.(int))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "DaysInSession":
                      n.DaysInSession = IntCreate(value.(int))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *LocationType) Set(key string, value interface{}) *LocationType {
        if n == nil {
                n = LocationTypeCreate(LocationType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "LocationName":
                      n.LocationName = StringCreate(value.(string))
    case "LocationRefId":
                      n.LocationRefId = LocationType_LocationRefIdCreate(value.(LocationType_LocationRefId))
        }
        return n
}

func (n *StudentContactPersonal) Set(key string, value interface{}) *StudentContactPersonal {
        if n == nil {
                n = StudentContactPersonalCreate(StudentContactPersonal{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "NonSchoolEducation":
      present := false
  for _, b := range AUCodeSetsNonSchoolEducationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNonSchoolEducationType_values")
      }

                      n.NonSchoolEducation = ((*AUCodeSetsNonSchoolEducationType)(StringCreate(value.(string))))
    case "SchoolEducationalLevel":
    
                      n.SchoolEducationalLevel = ((*EducationalLevelType)(StringCreate(value.(string))))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "EmploymentType":
      present := false
  for _, b := range AUCodeSetsEmploymentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEmploymentTypeType_values")
      }

                      n.EmploymentType = ((*AUCodeSetsEmploymentTypeType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ValidLetterMarkType) Set(key string, value interface{}) *ValidLetterMarkType {
        if n == nil {
                n = ValidLetterMarkTypeCreate(ValidLetterMarkType{})
        }
        switch key {
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "NumericEquivalent":
                      n.NumericEquivalent = FloatCreate(value.(float64))
        }
        return n
}

func (n *ActivityTimeType) Set(key string, value interface{}) *ActivityTimeType {
        if n == nil {
                n = ActivityTimeTypeCreate(ActivityTimeType{})
        }
        switch key {
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "Duration":
                      n.Duration = ActivityTimeType_DurationCreate(value.(ActivityTimeType_Duration))
    case "FinishDate":
                      n.FinishDate = StringCreate(value.(string))
        }
        return n
}

func (n *TeachingGroup) Set(key string, value interface{}) *TeachingGroup {
        if n == nil {
                n = TeachingGroupCreate(TeachingGroup{})
        }
        switch key {
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "LongName":
                      n.LongName = StringCreate(value.(string))
    case "KeyLearningArea":
      present := false
  for _, b := range AUCodeSetsACStrandType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsACStrandType_values")
      }

                      n.KeyLearningArea = ((*AUCodeSetsACStrandType)(StringCreate(value.(string))))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
    case "StudentList":
                      n.StudentList = StudentListTypeCreate(value.(StudentListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SchoolCourseLocalId":
    
                      n.SchoolCourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TeachingGroupPeriodList":
                      n.TeachingGroupPeriodList = TeachingGroupPeriodListTypeCreate(value.(TeachingGroupPeriodListType))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TeacherList":
                      n.TeacherList = TeacherListTypeCreate(value.(TeacherListType))
    case "GroupType":
                      n.GroupType = StringCreate(value.(string))
    case "Sett":
                      n.Sett = StringCreate(value.(string))
    case "SchoolInfoRefId":
    
                      n.SchoolInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "MaxClassSize":
                      n.MaxClassSize = IntCreate(value.(int))
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "MinClassSize":
                      n.MinClassSize = IntCreate(value.(int))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CurriculumLevel":
                      n.CurriculumLevel = StringCreate(value.(string))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Block":
                      n.Block = StringCreate(value.(string))
        }
        return n
}

func (n *StudentSectionEnrollment) Set(key string, value interface{}) *StudentSectionEnrollment {
        if n == nil {
                n = StudentSectionEnrollmentCreate(StudentSectionEnrollment{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ExitDate":
                      n.ExitDate = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "EntryDate":
                      n.EntryDate = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "SectionInfoRefId":
                      n.SectionInfoRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *MediumOfInstructionType) Set(key string, value interface{}) *MediumOfInstructionType {
        if n == nil {
                n = MediumOfInstructionTypeCreate(MediumOfInstructionType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsMediumOfInstructionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsMediumOfInstructionType_values")
      }

                      n.Code = ((*AUCodeSetsMediumOfInstructionType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *StatisticalAreaType) Set(key string, value interface{}) *StatisticalAreaType {
        if n == nil {
                n = StatisticalAreaTypeCreate(StatisticalAreaType{})
        }
        switch key {
    case "SpatialUnitType":
                      n.SpatialUnitType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentPeriodAttendance) Set(key string, value interface{}) *StudentPeriodAttendance {
        if n == nil {
                n = StudentPeriodAttendanceCreate(StudentPeriodAttendance{})
        }
        switch key {
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
    case "SessionInfoRefId":
                      n.SessionInfoRefId = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TimetablePeriod":
                      n.TimetablePeriod = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "AttendanceComment":
                      n.AttendanceComment = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "AttendanceStatus":
      present := false
  for _, b := range AUCodeSetsAttendanceStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAttendanceStatusType_values")
      }

                      n.AttendanceStatus = ((*AUCodeSetsAttendanceStatusType)(StringCreate(value.(string))))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AuditInfo":
                      n.AuditInfo = AuditInfoTypeCreate(value.(AuditInfoType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ComponentType) Set(key string, value interface{}) *ComponentType {
        if n == nil {
                n = ComponentTypeCreate(ComponentType{})
        }
        switch key {
    case "Strategies":
                      n.Strategies = StrategiesTypeCreate(value.(StrategiesType))
    case "AssociatedObjects":
                      n.AssociatedObjects = AssociatedObjectsTypeCreate(value.(AssociatedObjectsType))
    case "Reference":
                      n.Reference = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
        }
        return n
}

func (n *ContactInfoType) Set(key string, value interface{}) *ContactInfoType {
        if n == nil {
                n = ContactInfoTypeCreate(ContactInfoType{})
        }
        switch key {
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "PositionTitle":
                      n.PositionTitle = StringCreate(value.(string))
    case "Qualifications":
                      n.Qualifications = StringCreate(value.(string))
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "RegistrationDetails":
                      n.RegistrationDetails = StringCreate(value.(string))
        }
        return n
}

func (n *DomainBandsContainerType) Set(key string, value interface{}) *DomainBandsContainerType {
        if n == nil {
                n = DomainBandsContainerTypeCreate(DomainBandsContainerType{})
        }
        switch key {
    case "Band5Lower":
                      n.Band5Lower = FloatCreate(value.(float64))
    case "Band5Upper":
                      n.Band5Upper = FloatCreate(value.(float64))
    case "Band10Lower":
                      n.Band10Lower = FloatCreate(value.(float64))
    case "Band2Lower":
                      n.Band2Lower = FloatCreate(value.(float64))
    case "Band2Upper":
                      n.Band2Upper = FloatCreate(value.(float64))
    case "Band1Lower":
                      n.Band1Lower = FloatCreate(value.(float64))
    case "Band1Upper":
                      n.Band1Upper = FloatCreate(value.(float64))
    case "Band4Upper":
                      n.Band4Upper = FloatCreate(value.(float64))
    case "Band4Lower":
                      n.Band4Lower = FloatCreate(value.(float64))
    case "Band9Lower":
                      n.Band9Lower = FloatCreate(value.(float64))
    case "Band7Upper":
                      n.Band7Upper = FloatCreate(value.(float64))
    case "Band7Lower":
                      n.Band7Lower = FloatCreate(value.(float64))
    case "Band9Upper":
                      n.Band9Upper = FloatCreate(value.(float64))
    case "Band6Upper":
                      n.Band6Upper = FloatCreate(value.(float64))
    case "Band10Upper":
                      n.Band10Upper = FloatCreate(value.(float64))
    case "Band6Lower":
                      n.Band6Lower = FloatCreate(value.(float64))
    case "Band3Upper":
                      n.Band3Upper = FloatCreate(value.(float64))
    case "Band8Lower":
                      n.Band8Lower = FloatCreate(value.(float64))
    case "Band3Lower":
                      n.Band3Lower = FloatCreate(value.(float64))
    case "Band8Upper":
                      n.Band8Upper = FloatCreate(value.(float64))
        }
        return n
}

func (n *AddressStreetType) Set(key string, value interface{}) *AddressStreetType {
        if n == nil {
                n = AddressStreetTypeCreate(AddressStreetType{})
        }
        switch key {
    case "Line3":
                      n.Line3 = StringCreate(value.(string))
    case "Complex":
                      n.Complex = StringCreate(value.(string))
    case "StreetName":
                      n.StreetName = StringCreate(value.(string))
    case "ApartmentNumber":
                      n.ApartmentNumber = StringCreate(value.(string))
    case "StreetType":
                      n.StreetType = StringCreate(value.(string))
    case "StreetPrefix":
                      n.StreetPrefix = StringCreate(value.(string))
    case "Line1":
                      n.Line1 = StringCreate(value.(string))
    case "StreetNumber":
                      n.StreetNumber = StringCreate(value.(string))
    case "StreetSuffix":
                      n.StreetSuffix = StringCreate(value.(string))
    case "ApartmentType":
                      n.ApartmentType = StringCreate(value.(string))
    case "ApartmentNumberSuffix":
                      n.ApartmentNumberSuffix = StringCreate(value.(string))
    case "ApartmentNumberPrefix":
                      n.ApartmentNumberPrefix = StringCreate(value.(string))
    case "Line2":
                      n.Line2 = StringCreate(value.(string))
        }
        return n
}

func (n *LanguageOfInstructionType) Set(key string, value interface{}) *LanguageOfInstructionType {
        if n == nil {
                n = LanguageOfInstructionTypeCreate(LanguageOfInstructionType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values")
      }

                      n.Code = ((*AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *YearLevelType) Set(key string, value interface{}) *YearLevelType {
        if n == nil {
                n = YearLevelTypeCreate(YearLevelType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsYearLevelCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYearLevelCodeType_values")
      }

                      n.Code = ((*AUCodeSetsYearLevelCodeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PaymentReceiptLineType) Set(key string, value interface{}) *PaymentReceiptLineType {
        if n == nil {
                n = PaymentReceiptLineTypeCreate(PaymentReceiptLineType{})
        }
        switch key {
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "LocalPaymentReceiptLineId":
    
                      n.LocalPaymentReceiptLineId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TransactionAmount":
                      n.TransactionAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
    case "FinancialAccountRefId":
                      n.FinancialAccountRefId = StringCreate(value.(string))
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
        }
        return n
}

func (n *Identity) Set(key string, value interface{}) *Identity {
        if n == nil {
                n = IdentityCreate(Identity{})
        }
        switch key {
    case "AuthenticationSource":
                      n.AuthenticationSource = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AuthenticationSourceGlobalUID":
                      n.AuthenticationSourceGlobalUID = StringCreate(value.(string))
    case "SIF_RefId":
                      n.SIF_RefId = Identity_SIF_RefIdCreate(value.(Identity_SIF_RefId))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "PasswordList":
                      n.PasswordList = PasswordListTypeCreate(value.(PasswordListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "IdentityAssertions":
                      n.IdentityAssertions = IdentityAssertionsTypeCreate(value.(IdentityAssertionsType))
        }
        return n
}

func (n *SchoolInfo_OtherLEA) Set(key string, value interface{}) *SchoolInfo_OtherLEA {
        if n == nil {
                n = SchoolInfo_OtherLEACreate(SchoolInfo_OtherLEA{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
    
                      n.Value = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *Identity_SIF_RefId) Set(key string, value interface{}) *Identity_SIF_RefId {
        if n == nil {
                n = Identity_SIF_RefIdCreate(Identity_SIF_RefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *SchoolInfo) Set(key string, value interface{}) *SchoolInfo {
        if n == nil {
                n = SchoolInfoCreate(SchoolInfo{})
        }
        switch key {
    case "SchoolURL":
    
                      n.SchoolURL = ((*SchoolURLType)(StringCreate(value.(string))))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "SchoolTimeZone":
      present := false
  for _, b := range AUCodeSetsAustralianTimeZoneType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianTimeZoneType_values")
      }

                      n.SchoolTimeZone = ((*AUCodeSetsAustralianTimeZoneType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LEAInfoRefId":
    
                      n.LEAInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolContactList":
                      n.SchoolContactList = SchoolContactListTypeCreate(value.(SchoolContactListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
    case "Campus":
                      n.Campus = CampusContainerTypeCreate(value.(CampusContainerType))
    case "JurisdictionLowerHouse":
                      n.JurisdictionLowerHouse = StringCreate(value.(string))
    case "PrincipalInfo":
                      n.PrincipalInfo = PrincipalInfoTypeCreate(value.(PrincipalInfoType))
    case "TotalEnrollments":
                      n.TotalEnrollments = TotalEnrollmentsTypeCreate(value.(TotalEnrollmentsType))
    case "SLA":
      present := false
  for _, b := range AUCodeSetsAustralianStandardGeographicalClassificationASGCType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardGeographicalClassificationASGCType_values")
      }

                      n.SLA = ((*AUCodeSetsAustralianStandardGeographicalClassificationASGCType)(StringCreate(value.(string))))
    case "SchoolDistrictLocalId":
    
                      n.SchoolDistrictLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ARIA":
                      n.ARIA = FloatCreate(value.(float64))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "SchoolFocusList":
                      n.SchoolFocusList = SchoolFocusListTypeCreate(value.(SchoolFocusListType))
    case "Entity_Open":
                      n.Entity_Open = StringCreate(value.(string))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "OperationalStatus":
    
                      n.OperationalStatus = ((*OperationalStatusType)(StringCreate(value.(string))))
    case "SchoolCoEdStatus":
      present := false
  for _, b := range AUCodeSetsSchoolCoEdStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolCoEdStatusType_values")
      }

                      n.SchoolCoEdStatus = ((*AUCodeSetsSchoolCoEdStatusType)(StringCreate(value.(string))))
    case "System":
      present := false
  for _, b := range AUCodeSetsSchoolSystemType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolSystemType_values")
      }

                      n.System = ((*AUCodeSetsSchoolSystemType)(StringCreate(value.(string))))
    case "SessionType":
      present := false
  for _, b := range AUCodeSetsSessionTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSessionTypeType_values")
      }

                      n.SessionType = ((*AUCodeSetsSessionTypeType)(StringCreate(value.(string))))
    case "IndependentSchool":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.IndependentSchool = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "OtherLEA":
                      n.OtherLEA = SchoolInfo_OtherLEACreate(value.(SchoolInfo_OtherLEA))
    case "FederalElectorate":
      present := false
  for _, b := range AUCodeSetsFederalElectorateType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsFederalElectorateType_values")
      }

                      n.FederalElectorate = ((*AUCodeSetsFederalElectorateType)(StringCreate(value.(string))))
    case "SchoolType":
      present := false
  for _, b := range AUCodeSetsSchoolLevelType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolLevelType_values")
      }

                      n.SchoolType = ((*AUCodeSetsSchoolLevelType)(StringCreate(value.(string))))
    case "SchoolEmailList":
                      n.SchoolEmailList = EmailListTypeCreate(value.(EmailListType))
    case "YearLevelEnrollmentList":
                      n.YearLevelEnrollmentList = YearLevelEnrollmentListTypeCreate(value.(YearLevelEnrollmentListType))
    case "SchoolDistrict":
                      n.SchoolDistrict = StringCreate(value.(string))
    case "SchoolGeographicLocation":
      present := false
  for _, b := range AUCodeSetsSchoolLocationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolLocationType_values")
      }

                      n.SchoolGeographicLocation = ((*AUCodeSetsSchoolLocationType)(StringCreate(value.(string))))
    case "SchoolSector":
      present := false
  for _, b := range AUCodeSetsSchoolSectorCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolSectorCodeType_values")
      }

                      n.SchoolSector = ((*AUCodeSetsSchoolSectorCodeType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "BoardingSchoolStatus":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.BoardingSchoolStatus = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SchoolGroupList":
                      n.SchoolGroupList = SchoolGroupListTypeCreate(value.(SchoolGroupListType))
    case "NonGovSystemicStatus":
      present := false
  for _, b := range AUCodeSetsSystemicStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSystemicStatusType_values")
      }

                      n.NonGovSystemicStatus = ((*AUCodeSetsSystemicStatusType)(StringCreate(value.(string))))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ReligiousAffiliation":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType_values")
      }

                      n.ReligiousAffiliation = ((*AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType)(StringCreate(value.(string))))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Entity_Close":
                      n.Entity_Close = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LocalGovernmentArea":
                      n.LocalGovernmentArea = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *StaffAssignment) Set(key string, value interface{}) *StaffAssignment {
        if n == nil {
                n = StaffAssignmentCreate(StaffAssignment{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "JobStartDate":
                      n.JobStartDate = StringCreate(value.(string))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "JobFunction":
                      n.JobFunction = StringCreate(value.(string))
    case "StaffSubjectList":
                      n.StaffSubjectList = StaffSubjectListTypeCreate(value.(StaffSubjectListType))
    case "AvailableForTimetable":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AvailableForTimetable = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "House":
                      n.House = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "JobEndDate":
                      n.JobEndDate = StringCreate(value.(string))
    case "PreviousSchoolName":
                      n.PreviousSchoolName = StringCreate(value.(string))
    case "CalendarSummaryList":
                      n.CalendarSummaryList = CalendarSummaryListTypeCreate(value.(CalendarSummaryListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "JobFTE":
                      n.JobFTE = FloatCreate(value.(float64))
    case "PrimaryAssignment":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PrimaryAssignment = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "StaffActivity":
                      n.StaffActivity = StaffActivityExtensionTypeCreate(value.(StaffActivityExtensionType))
    case "CasualReliefTeacher":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.CasualReliefTeacher = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "EmploymentStatus":
      present := false
  for _, b := range AUCodeSetsStaffStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStaffStatusType_values")
      }

                      n.EmploymentStatus = ((*AUCodeSetsStaffStatusType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *SchoolProgramType) Set(key string, value interface{}) *SchoolProgramType {
        if n == nil {
                n = SchoolProgramTypeCreate(SchoolProgramType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Category":
                      n.Category = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *ApprovalType) Set(key string, value interface{}) *ApprovalType {
        if n == nil {
                n = ApprovalTypeCreate(ApprovalType{})
        }
        switch key {
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "Organization":
                      n.Organization = StringCreate(value.(string))
        }
        return n
}

func (n *EmailType) Set(key string, value interface{}) *EmailType {
        if n == nil {
                n = EmailTypeCreate(EmailType{})
        }
        switch key {
    case "Type":
      present := false
  for _, b := range AUCodeSetsEmailTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEmailTypeType_values")
      }

                      n.Type = ((*AUCodeSetsEmailTypeType)(StringCreate(value.(string))))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentAttendanceCollectionReportingType) Set(key string, value interface{}) *StudentAttendanceCollectionReportingType {
        if n == nil {
                n = StudentAttendanceCollectionReportingTypeCreate(StudentAttendanceCollectionReportingType{})
        }
        switch key {
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "StatsCohortYearLevelList":
                      n.StatsCohortYearLevelList = StatsCohortYearLevelListTypeCreate(value.(StatsCohortYearLevelListType))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
        }
        return n
}

func (n *Debtor) Set(key string, value interface{}) *Debtor {
        if n == nil {
                n = DebtorCreate(Debtor{})
        }
        switch key {
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "BillingName":
                      n.BillingName = StringCreate(value.(string))
    case "BillingNote":
                      n.BillingNote = StringCreate(value.(string))
    case "BilledEntity":
                      n.BilledEntity = Debtor_BilledEntityCreate(value.(Debtor_BilledEntity))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Discount":
                      n.Discount = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        }
        return n
}

func (n *AgencyType) Set(key string, value interface{}) *AgencyType {
        if n == nil {
                n = AgencyTypeCreate(AgencyType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsEducationAgencyTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEducationAgencyTypeType_values")
      }

                      n.Code = ((*AUCodeSetsEducationAgencyTypeType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *StaffAssignmentMostRecentContainerType) Set(key string, value interface{}) *StaffAssignmentMostRecentContainerType {
        if n == nil {
                n = StaffAssignmentMostRecentContainerTypeCreate(StaffAssignmentMostRecentContainerType{})
        }
        switch key {
    case "SecondaryFTE":
                      n.SecondaryFTE = FloatCreate(value.(float64))
    case "PrimaryFTE":
                      n.PrimaryFTE = FloatCreate(value.(float64))
        }
        return n
}

func (n *AGRoundType) Set(key string, value interface{}) *AGRoundType {
        if n == nil {
                n = AGRoundTypeCreate(AGRoundType{})
        }
        switch key {
    case "RoundName":
                      n.RoundName = StringCreate(value.(string))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
        }
        return n
}

func (n *AGContextualQuestionType) Set(key string, value interface{}) *AGContextualQuestionType {
        if n == nil {
                n = AGContextualQuestionTypeCreate(AGContextualQuestionType{})
        }
        switch key {
    case "AGAnswer":
                      n.AGAnswer = StringCreate(value.(string))
    case "AGContextCode":
      present := false
  for _, b := range AUCodeSetsAGContextQuestionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAGContextQuestionType_values")
      }

                      n.AGContextCode = ((*AUCodeSetsAGContextQuestionType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PurchasingItemType) Set(key string, value interface{}) *PurchasingItemType {
        if n == nil {
                n = PurchasingItemTypeCreate(PurchasingItemType{})
        }
        switch key {
    case "ItemNumber":
                      n.ItemNumber = StringCreate(value.(string))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "ExpenseAccounts":
                      n.ExpenseAccounts = ExpenseAccountsTypeCreate(value.(ExpenseAccountsType))
    case "ItemDescription":
                      n.ItemDescription = StringCreate(value.(string))
    case "QuantityDelivered":
                      n.QuantityDelivered = StringCreate(value.(string))
    case "Quantity":
                      n.Quantity = StringCreate(value.(string))
    case "UnitCost":
                      n.UnitCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TotalCost":
                      n.TotalCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "CancelledOrder":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.CancelledOrder = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "LocalItemId":
    
                      n.LocalItemId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *TimeElementType) Set(key string, value interface{}) *TimeElementType {
        if n == nil {
                n = TimeElementTypeCreate(TimeElementType{})
        }
        switch key {
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "SpanGaps":
                      n.SpanGaps = TimeElementType_SpanGapsCreate(value.(TimeElementType_SpanGaps))
    case "EndDateTime":
                      n.EndDateTime = StringCreate(value.(string))
    case "IsCurrent":
                      n.IsCurrent = BoolCreate(value.(bool))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *ResourceUsage_ResourceUsageContentType) Set(key string, value interface{}) *ResourceUsage_ResourceUsageContentType {
        if n == nil {
                n = ResourceUsage_ResourceUsageContentTypeCreate(ResourceUsage_ResourceUsageContentType{})
        }
        switch key {
    case "LocalDescription":
                      n.LocalDescription = StringCreate(value.(string))
    case "Code":
      present := false
  for _, b := range AUCodeSetsResourceUsageContentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsResourceUsageContentTypeType_values")
      }

                      n.Code = ((*AUCodeSetsResourceUsageContentTypeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *AggregateStatisticInfo_CalculationRule) Set(key string, value interface{}) *AggregateStatisticInfo_CalculationRule {
        if n == nil {
                n = AggregateStatisticInfo_CalculationRuleCreate(AggregateStatisticInfo_CalculationRule{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *SchoolPrograms) Set(key string, value interface{}) *SchoolPrograms {
        if n == nil {
                n = SchoolProgramsCreate(SchoolPrograms{})
        }
        switch key {
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolProgramList":
                      n.SchoolProgramList = SchoolProgramListTypeCreate(value.(SchoolProgramListType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        }
        return n
}

func (n *ExclusionRuleType) Set(key string, value interface{}) *ExclusionRuleType {
        if n == nil {
                n = ExclusionRuleTypeCreate(ExclusionRuleType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *SystemRole_SystemContext) Set(key string, value interface{}) *SystemRole_SystemContext {
        if n == nil {
                n = SystemRole_SystemContextCreate(SystemRole_SystemContext{})
        }
        switch key {
    case "RoleList":
                      n.RoleList = SystemRole_RoleListCreate(value.(SystemRole_RoleList))
    case "SystemId":
                      n.SystemId = StringCreate(value.(string))
        }
        return n
}

func (n *MapReferenceType) Set(key string, value interface{}) *MapReferenceType {
        if n == nil {
                n = MapReferenceTypeCreate(MapReferenceType{})
        }
        switch key {
    case "XCoordinate":
                      n.XCoordinate = StringCreate(value.(string))
    case "YCoordinate":
                      n.YCoordinate = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "MapNumber":
                      n.MapNumber = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingCharacteristic) Set(key string, value interface{}) *WellbeingCharacteristic {
        if n == nil {
                n = WellbeingCharacteristicCreate(WellbeingCharacteristic{})
        }
        switch key {
    case "EmergencyResponsePlan":
                      n.EmergencyResponsePlan = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "WellbeingCharacteristicEndDate":
                      n.WellbeingCharacteristicEndDate = StringCreate(value.(string))
    case "CharacteristicSeverity":
                      n.CharacteristicSeverity = StringCreate(value.(string))
    case "WellbeingCharacteristicClassification":
      present := false
  for _, b := range AUCodeSetsWellbeingCharacteristicClassificationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingCharacteristicClassificationType_values")
      }

                      n.WellbeingCharacteristicClassification = ((*AUCodeSetsWellbeingCharacteristicClassificationType)(StringCreate(value.(string))))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "Alert":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.Alert = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Trigger":
                      n.Trigger = StringCreate(value.(string))
    case "MedicationList":
                      n.MedicationList = MedicationListTypeCreate(value.(MedicationListType))
    case "WellbeingCharacteristicReviewDate":
                      n.WellbeingCharacteristicReviewDate = StringCreate(value.(string))
    case "DailyManagement":
                      n.DailyManagement = StringCreate(value.(string))
    case "ConfidentialFlag":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ConfidentialFlag = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EmergencyManagement":
                      n.EmergencyManagement = StringCreate(value.(string))
    case "WellbeingCharacteristicCategory":
                      n.WellbeingCharacteristicCategory = StringCreate(value.(string))
    case "LocalCharacteristicCode":
    
                      n.LocalCharacteristicCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "WellbeingCharacteristicStartDate":
                      n.WellbeingCharacteristicStartDate = StringCreate(value.(string))
    case "WellbeingCharacteristicNotes":
                      n.WellbeingCharacteristicNotes = StringCreate(value.(string))
    case "WellbeingCharacteristicSubCategory":
                      n.WellbeingCharacteristicSubCategory = StringCreate(value.(string))
    case "SymptomList":
                      n.SymptomList = SymptomListTypeCreate(value.(SymptomListType))
        }
        return n
}

func (n *LearningResource_Location) Set(key string, value interface{}) *LearningResource_Location {
        if n == nil {
                n = LearningResource_LocationCreate(LearningResource_Location{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "ReferenceType":
                      n.ReferenceType = StringCreate(value.(string))
        }
        return n
}

func (n *YearLevelEnrollmentType) Set(key string, value interface{}) *YearLevelEnrollmentType {
        if n == nil {
                n = YearLevelEnrollmentTypeCreate(YearLevelEnrollmentType{})
        }
        switch key {
    case "Enrollment":
                      n.Enrollment = StringCreate(value.(string))
    case "Year":
      present := false
  for _, b := range AUCodeSetsYearLevelCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYearLevelCodeType_values")
      }

                      n.Year = ((*AUCodeSetsYearLevelCodeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *SystemRole_Role) Set(key string, value interface{}) *SystemRole_Role {
        if n == nil {
                n = SystemRole_RoleCreate(SystemRole_Role{})
        }
        switch key {
    case "RoleScopeList":
                      n.RoleScopeList = SystemRole_RoleScopeListCreate(value.(SystemRole_RoleScopeList))
    case "RoleId":
                      n.RoleId = StringCreate(value.(string))
        }
        return n
}

func (n *EntityContactInfoType) Set(key string, value interface{}) *EntityContactInfoType {
        if n == nil {
                n = EntityContactInfoTypeCreate(EntityContactInfoType{})
        }
        switch key {
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "Email":
                      n.Email = EmailTypeCreate(value.(EmailType))
    case "RegistrationDetails":
                      n.RegistrationDetails = StringCreate(value.(string))
    case "PositionTitle":
                      n.PositionTitle = StringCreate(value.(string))
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
    case "Qualifications":
                      n.Qualifications = StringCreate(value.(string))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
        }
        return n
}

func (n *AbstractContentPackageType_XMLData) Set(key string, value interface{}) *AbstractContentPackageType_XMLData {
        if n == nil {
                n = AbstractContentPackageType_XMLDataCreate(AbstractContentPackageType_XMLData{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
        }
        return n
}

func (n *MarkValueInfo) Set(key string, value interface{}) *MarkValueInfo {
        if n == nil {
                n = MarkValueInfoCreate(MarkValueInfo{})
        }
        switch key {
    case "Narrative":
                      n.Narrative = StringCreate(value.(string))
    case "PercentagePassingGrade":
                      n.PercentagePassingGrade = FloatCreate(value.(float64))
    case "NumericLow":
                      n.NumericLow = FloatCreate(value.(float64))
    case "PercentageMinimum":
                      n.PercentageMinimum = FloatCreate(value.(float64))
    case "NumericScale":
                      n.NumericScale = IntCreate(value.(int))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "NumericHigh":
                      n.NumericHigh = FloatCreate(value.(float64))
    case "NarrativeMaximumSize":
                      n.NarrativeMaximumSize = IntCreate(value.(int))
    case "ValidLetterMarkList":
                      n.ValidLetterMarkList = ValidLetterMarkListTypeCreate(value.(ValidLetterMarkListType))
    case "PercentageMaximum":
                      n.PercentageMaximum = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "NumericPrecision":
                      n.NumericPrecision = IntCreate(value.(int))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "NumericPassingGrade":
                      n.NumericPassingGrade = FloatCreate(value.(float64))
        }
        return n
}

func (n *Invoice_InvoicedEntity) Set(key string, value interface{}) *Invoice_InvoicedEntity {
        if n == nil {
                n = Invoice_InvoicedEntityCreate(Invoice_InvoicedEntity{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *PhoneNumberType) Set(key string, value interface{}) *PhoneNumberType {
        if n == nil {
                n = PhoneNumberTypeCreate(PhoneNumberType{})
        }
        switch key {
    case "Type":
      present := false
  for _, b := range AUCodeSetsTelephoneNumberTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsTelephoneNumberTypeType_values")
      }

                      n.Type = ((*AUCodeSetsTelephoneNumberTypeType)(StringCreate(value.(string))))
    case "Preference":
                      n.Preference = IntCreate(value.(int))
    case "ListedStatus":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ListedStatus = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Extension":
                      n.Extension = StringCreate(value.(string))
    case "Number":
                      n.Number = StringCreate(value.(string))
        }
        return n
}

func (n *ProgramFundingSourceType) Set(key string, value interface{}) *ProgramFundingSourceType {
        if n == nil {
                n = ProgramFundingSourceTypeCreate(ProgramFundingSourceType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsProgramFundingSourceCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsProgramFundingSourceCodeType_values")
      }

                      n.Code = ((*AUCodeSetsProgramFundingSourceCodeType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *OtherNameType) Set(key string, value interface{}) *OtherNameType {
        if n == nil {
                n = OtherNameTypeCreate(OtherNameType{})
        }
        switch key {
    case "Type":
      present := false
  for _, b := range AUCodeSetsNameUsageTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNameUsageTypeType_values")
      }

                      n.Type = ((*AUCodeSetsNameUsageTypeType)(StringCreate(value.(string))))
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "FamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PreferredFamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PreferredFamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
        }
        return n
}

func (n *DemographicsType) Set(key string, value interface{}) *DemographicsType {
        if n == nil {
                n = DemographicsTypeCreate(DemographicsType{})
        }
        switch key {
    case "LBOTE":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.LBOTE = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PlaceOfBirth":
                      n.PlaceOfBirth = StringCreate(value.(string))
    case "ReligiousRegion":
                      n.ReligiousRegion = StringCreate(value.(string))
    case "BirthDate":
    
                      n.BirthDate = ((*BirthDateType)(StringCreate(value.(string))))
    case "DateOfDeath":
                      n.DateOfDeath = StringCreate(value.(string))
    case "ReligiousEventList":
                      n.ReligiousEventList = ReligiousEventListTypeCreate(value.(ReligiousEventListType))
    case "VisaSubClassList":
                      n.VisaSubClassList = VisaSubClassListTypeCreate(value.(VisaSubClassListType))
    case "MaritalStatus":
      present := false
  for _, b := range AUCodeSetsMaritalStatusAIHWType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsMaritalStatusAIHWType_values")
      }

                      n.MaritalStatus = ((*AUCodeSetsMaritalStatusAIHWType)(StringCreate(value.(string))))
    case "Religion":
                      n.Religion = ReligionTypeCreate(value.(ReligionType))
    case "StateOfBirth":
    
                      n.StateOfBirth = ((*StateProvinceType)(StringCreate(value.(string))))
    case "CountryOfBirth":
    
                      n.CountryOfBirth = ((*CountryType)(StringCreate(value.(string))))
    case "VisaStatisticalCode":
                      n.VisaStatisticalCode = StringCreate(value.(string))
    case "CountryArrivalDate":
                      n.CountryArrivalDate = StringCreate(value.(string))
    case "IndigenousStatus":
      present := false
  for _, b := range AUCodeSetsIndigenousStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsIndigenousStatusType_values")
      }

                      n.IndigenousStatus = ((*AUCodeSetsIndigenousStatusType)(StringCreate(value.(string))))
    case "CountriesOfResidency":
                      n.CountriesOfResidency = CountryList2TypeCreate(value.(CountryList2Type))
    case "LanguageList":
                      n.LanguageList = LanguageListTypeCreate(value.(LanguageListType))
    case "VisaExpiryDate":
                      n.VisaExpiryDate = StringCreate(value.(string))
    case "CountriesOfCitizenship":
                      n.CountriesOfCitizenship = CountryListTypeCreate(value.(CountryListType))
    case "CulturalBackground":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType_values")
      }

                      n.CulturalBackground = ((*AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType)(StringCreate(value.(string))))
    case "BirthDateVerification":
      present := false
  for _, b := range AUCodeSetsBirthdateVerificationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsBirthdateVerificationType_values")
      }

                      n.BirthDateVerification = ((*AUCodeSetsBirthdateVerificationType)(StringCreate(value.(string))))
    case "AustralianCitizenshipStatus":
      present := false
  for _, b := range AUCodeSetsAustralianCitizenshipStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianCitizenshipStatusType_values")
      }

                      n.AustralianCitizenshipStatus = ((*AUCodeSetsAustralianCitizenshipStatusType)(StringCreate(value.(string))))
    case "ImmunisationCertificateStatus":
      present := false
  for _, b := range AUCodeSetsImmunisationCertificateStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsImmunisationCertificateStatusType_values")
      }

                      n.ImmunisationCertificateStatus = ((*AUCodeSetsImmunisationCertificateStatusType)(StringCreate(value.(string))))
    case "MedicareNumber":
                      n.MedicareNumber = StringCreate(value.(string))
    case "Sex":
      present := false
  for _, b := range AUCodeSetsSexCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSexCodeType_values")
      }

                      n.Sex = ((*AUCodeSetsSexCodeType)(StringCreate(value.(string))))
    case "PermanentResident":
      present := false
  for _, b := range AUCodeSetsPermanentResidentStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPermanentResidentStatusType_values")
      }

                      n.PermanentResident = ((*AUCodeSetsPermanentResidentStatusType)(StringCreate(value.(string))))
    case "DwellingArrangement":
                      n.DwellingArrangement = DwellingArrangementTypeCreate(value.(DwellingArrangementType))
    case "InterpreterRequired":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.InterpreterRequired = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "VisaSubClass":
    
                      n.VisaSubClass = ((*VisaSubClassCodeType)(StringCreate(value.(string))))
    case "EnglishProficiency":
                      n.EnglishProficiency = EnglishProficiencyTypeCreate(value.(EnglishProficiencyType))
        }
        return n
}

func (n *LifeCycleType_Creator) Set(key string, value interface{}) *LifeCycleType_Creator {
        if n == nil {
                n = LifeCycleType_CreatorCreate(LifeCycleType_Creator{})
        }
        switch key {
    case "ID":
                      n.ID = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
        }
        return n
}

func (n *SuspensionContainerType) Set(key string, value interface{}) *SuspensionContainerType {
        if n == nil {
                n = SuspensionContainerTypeCreate(SuspensionContainerType{})
        }
        switch key {
    case "ResolutionNotes":
                      n.ResolutionNotes = StringCreate(value.(string))
    case "AdvisementDate":
                      n.AdvisementDate = StringCreate(value.(string))
    case "EarlyReturnDate":
                      n.EarlyReturnDate = StringCreate(value.(string))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
    case "WithdrawalTimeList":
                      n.WithdrawalTimeList = WithdrawalTimeListTypeCreate(value.(WithdrawalTimeListType))
    case "ResolutionMeetingTime":
                      n.ResolutionMeetingTime = StringCreate(value.(string))
    case "SuspensionCategory":
      present := false
  for _, b := range AUCodeSetsSuspensionCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSuspensionCategoryType_values")
      }

                      n.SuspensionCategory = ((*AUCodeSetsSuspensionCategoryType)(StringCreate(value.(string))))
    case "SuspensionNotes":
                      n.SuspensionNotes = StringCreate(value.(string))
    case "Duration":
                      n.Duration = FloatCreate(value.(float64))
        }
        return n
}

func (n *StudentMostRecentContainerType) Set(key string, value interface{}) *StudentMostRecentContainerType {
        if n == nil {
                n = StudentMostRecentContainerTypeCreate(StudentMostRecentContainerType{})
        }
        switch key {
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "Parent2SchoolEducationLevel":
      present := false
  for _, b := range AUCodeSetsSchoolEducationLevelTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolEducationLevelTypeType_values")
      }

                      n.Parent2SchoolEducationLevel = ((*AUCodeSetsSchoolEducationLevelTypeType)(StringCreate(value.(string))))
    case "BoardingStatus":
      present := false
  for _, b := range AUCodeSetsBoardingType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsBoardingType_values")
      }

                      n.BoardingStatus = ((*AUCodeSetsBoardingType)(StringCreate(value.(string))))
    case "Parent1NonSchoolEducation":
      present := false
  for _, b := range AUCodeSetsNonSchoolEducationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNonSchoolEducationType_values")
      }

                      n.Parent1NonSchoolEducation = ((*AUCodeSetsNonSchoolEducationType)(StringCreate(value.(string))))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "ClassCode":
                      n.ClassCode = StringCreate(value.(string))
    case "FFPOS":
      present := false
  for _, b := range AUCodeSetsFFPOSStatusCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsFFPOSStatusCodeType_values")
      }

                      n.FFPOS = ((*AUCodeSetsFFPOSStatusCodeType)(StringCreate(value.(string))))
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DistanceEducationStudent":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.DistanceEducationStudent = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Parent1EmploymentType":
      present := false
  for _, b := range AUCodeSetsEmploymentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEmploymentTypeType_values")
      }

                      n.Parent1EmploymentType = ((*AUCodeSetsEmploymentTypeType)(StringCreate(value.(string))))
    case "MembershipType":
      present := false
  for _, b := range AUCodeSetsSchoolEnrollmentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolEnrollmentTypeType_values")
      }

                      n.MembershipType = ((*AUCodeSetsSchoolEnrollmentTypeType)(StringCreate(value.(string))))
    case "Parent1Language":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values")
      }

                      n.Parent1Language = ((*AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType)(StringCreate(value.(string))))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ReportingSchoolId":
    
                      n.ReportingSchoolId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CensusAge":
                      n.CensusAge = IntCreate(value.(int))
    case "Parent2NonSchoolEducation":
      present := false
  for _, b := range AUCodeSetsNonSchoolEducationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNonSchoolEducationType_values")
      }

                      n.Parent2NonSchoolEducation = ((*AUCodeSetsNonSchoolEducationType)(StringCreate(value.(string))))
    case "OtherEnrollmentSchoolACARAId":
    
                      n.OtherEnrollmentSchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "Parent1SchoolEducationLevel":
      present := false
  for _, b := range AUCodeSetsSchoolEducationLevelTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolEducationLevelTypeType_values")
      }

                      n.Parent1SchoolEducationLevel = ((*AUCodeSetsSchoolEducationLevelTypeType)(StringCreate(value.(string))))
    case "HomeroomLocalId":
    
                      n.HomeroomLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Parent2EmploymentType":
      present := false
  for _, b := range AUCodeSetsEmploymentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEmploymentTypeType_values")
      }

                      n.Parent2EmploymentType = ((*AUCodeSetsEmploymentTypeType)(StringCreate(value.(string))))
    case "Parent2Language":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values")
      }

                      n.Parent2Language = ((*AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType)(StringCreate(value.(string))))
    case "OtherSchoolName":
                      n.OtherSchoolName = StringCreate(value.(string))
    case "LocalCampusId":
    
                      n.LocalCampusId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
        }
        return n
}

func (n *LibraryPatronStatus) Set(key string, value interface{}) *LibraryPatronStatus {
        if n == nil {
                n = LibraryPatronStatusCreate(LibraryPatronStatus{})
        }
        switch key {
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
    case "PatronLocalId":
    
                      n.PatronLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FineAmount":
                      n.FineAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PatronName":
                      n.PatronName = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "TransactionList":
                      n.TransactionList = LibraryTransactionListTypeCreate(value.(LibraryTransactionListType))
    case "NumberOfOverdues":
                      n.NumberOfOverdues = IntCreate(value.(int))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "NumberOfFines":
                      n.NumberOfFines = IntCreate(value.(int))
    case "NumberOfCheckouts":
                      n.NumberOfCheckouts = IntCreate(value.(int))
    case "MessageList":
                      n.MessageList = LibraryMessageListTypeCreate(value.(LibraryMessageListType))
    case "PatronRefObject":
                      n.PatronRefObject = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefundAmount":
                      n.RefundAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "NumberOfHoldItems":
                      n.NumberOfHoldItems = IntCreate(value.(int))
    case "PatronRefId":
                      n.PatronRefId = StringCreate(value.(string))
    case "NumberOfRefunds":
                      n.NumberOfRefunds = IntCreate(value.(int))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LibraryType":
                      n.LibraryType = StringCreate(value.(string))
        }
        return n
}

func (n *StimulusType) Set(key string, value interface{}) *StimulusType {
        if n == nil {
                n = StimulusTypeCreate(StimulusType{})
        }
        switch key {
    case "Content":
                      n.Content = StringCreate(value.(string))
    case "TextGenre":
                      n.TextGenre = StringCreate(value.(string))
    case "WordCount":
                      n.WordCount = IntCreate(value.(int))
    case "TextDescriptor":
                      n.TextDescriptor = StringCreate(value.(string))
    case "StimulusLocalId":
    
                      n.StimulusLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TextType":
                      n.TextType = StringCreate(value.(string))
        }
        return n
}

func (n *LocalCodeType) Set(key string, value interface{}) *LocalCodeType {
        if n == nil {
                n = LocalCodeTypeCreate(LocalCodeType{})
        }
        switch key {
    case "Element":
                      n.Element = StringCreate(value.(string))
    case "LocalisedCode":
                      n.LocalisedCode = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "ListIndex":
                      n.ListIndex = IntCreate(value.(int))
        }
        return n
}

func (n *GenericRubricType) Set(key string, value interface{}) *GenericRubricType {
        if n == nil {
                n = GenericRubricTypeCreate(GenericRubricType{})
        }
        switch key {
    case "Descriptor":
                      n.Descriptor = StringCreate(value.(string))
    case "RubricType":
                      n.RubricType = StringCreate(value.(string))
    case "ScoreList":
                      n.ScoreList = ScoreListTypeCreate(value.(ScoreListType))
        }
        return n
}

func (n *AbstractContentElementType_Reference) Set(key string, value interface{}) *AbstractContentElementType_Reference {
        if n == nil {
                n = AbstractContentElementType_ReferenceCreate(AbstractContentElementType_Reference{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "URL":
                      n.URL = StringCreate(value.(string))
        }
        return n
}

func (n *PersonPicture_PictureSource) Set(key string, value interface{}) *PersonPicture_PictureSource {
        if n == nil {
                n = PersonPicture_PictureSourceCreate(PersonPicture_PictureSource{})
        }
        switch key {
    case "Value":
    
                      n.Value = ((*URIOrBinaryType)(StringCreate(value.(string))))
    case "Type":
      present := false
  for _, b := range AUCodeSetsPictureSourceType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPictureSourceType_values")
      }

                      n.Type = ((*AUCodeSetsPictureSourceType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PeriodAttendanceType) Set(key string, value interface{}) *PeriodAttendanceType {
        if n == nil {
                n = PeriodAttendanceTypeCreate(PeriodAttendanceType{})
        }
        switch key {
    case "AttendanceStatus":
      present := false
  for _, b := range AUCodeSetsAttendanceStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAttendanceStatusType_values")
      }

                      n.AttendanceStatus = ((*AUCodeSetsAttendanceStatusType)(StringCreate(value.(string))))
    case "AttendanceType":
                      n.AttendanceType = StringCreate(value.(string))
    case "EndTime":
                      n.EndTime = StringCreate(value.(string))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "SessionInfoRefId":
                      n.SessionInfoRefId = StringCreate(value.(string))
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "TimetablePeriod":
                      n.TimetablePeriod = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
        }
        return n
}

func (n *StudentActivityType) Set(key string, value interface{}) *StudentActivityType {
        if n == nil {
                n = StudentActivityTypeCreate(StudentActivityType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsActivityInvolvementCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsActivityInvolvementCodeType_values")
      }

                      n.Code = ((*AUCodeSetsActivityInvolvementCodeType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *NAPTestItemContentType) Set(key string, value interface{}) *NAPTestItemContentType {
        if n == nil {
                n = NAPTestItemContentTypeCreate(NAPTestItemContentType{})
        }
        switch key {
    case "MarkingType":
      present := false
  for _, b := range AUCodeSetsNAPTestItemMarkingTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPTestItemMarkingTypeType_values")
      }

                      n.MarkingType = ((*AUCodeSetsNAPTestItemMarkingTypeType)(StringCreate(value.(string))))
    case "CorrectAnswer":
                      n.CorrectAnswer = StringCreate(value.(string))
    case "Subdomain":
                      n.Subdomain = StringCreate(value.(string))
    case "StimulusList":
                      n.StimulusList = StimulusListTypeCreate(value.(StimulusListType))
    case "ItemName":
                      n.ItemName = StringCreate(value.(string))
    case "ExemplarURL":
                      n.ExemplarURL = StringCreate(value.(string))
    case "ItemDifficultyLogit5SE":
                      n.ItemDifficultyLogit5SE = FloatCreate(value.(float64))
    case "ItemDescriptor":
                      n.ItemDescriptor = StringCreate(value.(string))
    case "ReleasedStatus":
                      n.ReleasedStatus = BoolCreate(value.(bool))
    case "ItemDifficultyLogit62":
                      n.ItemDifficultyLogit62 = FloatCreate(value.(float64))
    case "MaximumScore":
                      n.MaximumScore = FloatCreate(value.(float64))
    case "ItemDifficulty":
                      n.ItemDifficulty = FloatCreate(value.(float64))
    case "ItemProficiencyBand":
                      n.ItemProficiencyBand = IntCreate(value.(int))
    case "NAPTestItemLocalId":
    
                      n.NAPTestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPWritingRubricList":
                      n.NAPWritingRubricList = NAPWritingRubricListTypeCreate(value.(NAPWritingRubricListType))
    case "ItemSubstitutedForList":
                      n.ItemSubstitutedForList = SubstituteItemListTypeCreate(value.(SubstituteItemListType))
    case "ItemProficiencyLevel":
                      n.ItemProficiencyLevel = StringCreate(value.(string))
    case "ItemDifficultyLogit5":
                      n.ItemDifficultyLogit5 = FloatCreate(value.(float64))
    case "MultipleChoiceOptionCount":
                      n.MultipleChoiceOptionCount = IntCreate(value.(int))
    case "ContentDescriptionList":
                      n.ContentDescriptionList = ContentDescriptionListTypeCreate(value.(ContentDescriptionListType))
    case "ItemType":
      present := false
  for _, b := range AUCodeSetsNAPTestItemTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPTestItemTypeType_values")
      }

                      n.ItemType = ((*AUCodeSetsNAPTestItemTypeType)(StringCreate(value.(string))))
    case "WritingGenre":
      present := false
  for _, b := range AUCodeSetsNAPWritingGenreType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPWritingGenreType_values")
      }

                      n.WritingGenre = ((*AUCodeSetsNAPWritingGenreType)(StringCreate(value.(string))))
    case "ItemDifficultyLogit62SE":
                      n.ItemDifficultyLogit62SE = FloatCreate(value.(float64))
        }
        return n
}

func (n *TermInfo) Set(key string, value interface{}) *TermInfo {
        if n == nil {
                n = TermInfoCreate(TermInfo{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SchedulingTerm":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.SchedulingTerm = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Track":
                      n.Track = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "MarkingTerm":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.MarkingTerm = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "TermSpan":
      present := false
  for _, b := range AUCodeSetsSessionTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSessionTypeType_values")
      }

                      n.TermSpan = ((*AUCodeSetsSessionTypeType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "TermCode":
                      n.TermCode = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RelativeDuration":
                      n.RelativeDuration = FloatCreate(value.(float64))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AttendanceTerm":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AttendanceTerm = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingEventCategoryType) Set(key string, value interface{}) *WellbeingEventCategoryType {
        if n == nil {
                n = WellbeingEventCategoryTypeCreate(WellbeingEventCategoryType{})
        }
        switch key {
    case "WellbeingEventSubCategoryList":
                      n.WellbeingEventSubCategoryList = WellbeingEventSubCategoryListTypeCreate(value.(WellbeingEventSubCategoryListType))
    case "EventCategory":
                      n.EventCategory = StringCreate(value.(string))
        }
        return n
}

func (n *AddressCollection) Set(key string, value interface{}) *AddressCollection {
        if n == nil {
                n = AddressCollectionCreate(AddressCollection{})
        }
        switch key {
    case "AddressCollectionReportingList":
                      n.AddressCollectionReportingList = AddressCollectionReportingListTypeCreate(value.(AddressCollectionReportingListType))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AddressCollectionYear":
    
                      n.AddressCollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
        }
        return n
}

func (n *WithdrawalType) Set(key string, value interface{}) *WithdrawalType {
        if n == nil {
                n = WithdrawalTypeCreate(WithdrawalType{})
        }
        switch key {
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "WithdrawalEndTime":
                      n.WithdrawalEndTime = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "WithdrawalDate":
                      n.WithdrawalDate = StringCreate(value.(string))
    case "WithdrawalStartTime":
                      n.WithdrawalStartTime = StringCreate(value.(string))
        }
        return n
}

func (n *ACStrandSubjectAreaType) Set(key string, value interface{}) *ACStrandSubjectAreaType {
        if n == nil {
                n = ACStrandSubjectAreaTypeCreate(ACStrandSubjectAreaType{})
        }
        switch key {
    case "ACStrand":
      present := false
  for _, b := range AUCodeSetsACStrandType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsACStrandType_values")
      }

                      n.ACStrand = ((*AUCodeSetsACStrandType)(StringCreate(value.(string))))
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
        }
        return n
}

func (n *CensusStudentType) Set(key string, value interface{}) *CensusStudentType {
        if n == nil {
                n = CensusStudentTypeCreate(CensusStudentType{})
        }
        switch key {
    case "EducationMode":
                      n.EducationMode = StringCreate(value.(string))
    case "OverseasStudent":
                      n.OverseasStudent = StringCreate(value.(string))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "Headcount":
                      n.Headcount = IntCreate(value.(int))
    case "StudentCohortId":
    
                      n.StudentCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CohortIndigenousType":
                      n.CohortIndigenousType = StringCreate(value.(string))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "StudentOnVisa":
                      n.StudentOnVisa = StringCreate(value.(string))
    case "CensusAge":
                      n.CensusAge = IntCreate(value.(int))
        }
        return n
}

func (n *AddressType) Set(key string, value interface{}) *AddressType {
        if n == nil {
                n = AddressTypeCreate(AddressType{})
        }
        switch key {
    case "GridLocation":
                      n.GridLocation = GridLocationTypeCreate(value.(GridLocationType))
    case "MapReference":
                      n.MapReference = MapReferenceTypeCreate(value.(MapReferenceType))
    case "StatisticalAreas":
                      n.StatisticalAreas = StatisticalAreasTypeCreate(value.(StatisticalAreasType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Role":
      present := false
  for _, b := range AUCodeSetsAddressRoleType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAddressRoleType_values")
      }

                      n.Role = ((*AUCodeSetsAddressRoleType)(StringCreate(value.(string))))
    case "Country":
    
                      n.Country = ((*CountryType)(StringCreate(value.(string))))
    case "StateProvince":
    
                      n.StateProvince = ((*StateProvinceType)(StringCreate(value.(string))))
    case "EffectiveFromDate":
                      n.EffectiveFromDate = StringCreate(value.(string))
    case "PostalCode":
                      n.PostalCode = StringCreate(value.(string))
    case "EffectiveToDate":
                      n.EffectiveToDate = StringCreate(value.(string))
    case "Street":
                      n.Street = AddressStreetTypeCreate(value.(AddressStreetType))
    case "Community":
                      n.Community = StringCreate(value.(string))
    case "RadioContact":
                      n.RadioContact = StringCreate(value.(string))
    case "AddressGlobalUID":
    
                      n.AddressGlobalUID = ((*GUIDType)(StringCreate(value.(string))))
    case "City":
                      n.City = StringCreate(value.(string))
    case "Type":
      present := false
  for _, b := range AUCodeSetsAddressTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAddressTypeType_values")
      }

                      n.Type = ((*AUCodeSetsAddressTypeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *NAPEventStudentLink) Set(key string, value interface{}) *NAPEventStudentLink {
        if n == nil {
                n = NAPEventStudentLinkCreate(NAPEventStudentLink{})
        }
        switch key {
    case "ParticipationText":
                      n.ParticipationText = StringCreate(value.(string))
    case "Adjustment":
                      n.Adjustment = AdjustmentContainerTypeCreate(value.(AdjustmentContainerType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LapsedTimeTest":
                      n.LapsedTimeTest = StringCreate(value.(string))
    case "PossibleDuplicate":
                      n.PossibleDuplicate = BoolCreate(value.(bool))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "System":
      present := false
  for _, b := range AUCodeSetsSchoolSystemType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolSystemType_values")
      }

                      n.System = ((*AUCodeSetsSchoolSystemType)(StringCreate(value.(string))))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "ReportingSchoolName":
                      n.ReportingSchoolName = StringCreate(value.(string))
    case "ParticipationCode":
      present := false
  for _, b := range AUCodeSetsNAPParticipationCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPParticipationCodeType_values")
      }

                      n.ParticipationCode = ((*AUCodeSetsNAPParticipationCodeType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TestDisruptionList":
                      n.TestDisruptionList = TestDisruptionListTypeCreate(value.(TestDisruptionListType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "PersonalDetailsChanged":
                      n.PersonalDetailsChanged = BoolCreate(value.(bool))
    case "NAPJurisdiction":
      present := false
  for _, b := range AUCodeSetsNAPJurisdictionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPJurisdictionType_values")
      }

                      n.NAPJurisdiction = ((*AUCodeSetsNAPJurisdictionType)(StringCreate(value.(string))))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ExemptionReason":
                      n.ExemptionReason = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Device":
                      n.Device = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "PSIOtherIdMatch":
                      n.PSIOtherIdMatch = BoolCreate(value.(bool))
    case "PlatformStudentIdentifier":
    
                      n.PlatformStudentIdentifier = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolGeolocation":
      present := false
  for _, b := range AUCodeSetsSchoolLocationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolLocationType_values")
      }

                      n.SchoolGeolocation = ((*AUCodeSetsSchoolLocationType)(StringCreate(value.(string))))
    case "SchoolSector":
      present := false
  for _, b := range AUCodeSetsSchoolSectorCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolSectorCodeType_values")
      }

                      n.SchoolSector = ((*AUCodeSetsSchoolSectorCodeType)(StringCreate(value.(string))))
    case "DOBRange":
                      n.DOBRange = BoolCreate(value.(bool))
        }
        return n
}

func (n *LifeCycleType_Modified) Set(key string, value interface{}) *LifeCycleType_Modified {
        if n == nil {
                n = LifeCycleType_ModifiedCreate(LifeCycleType_Modified{})
        }
        switch key {
    case "DateTime":
                      n.DateTime = StringCreate(value.(string))
    case "By":
                      n.By = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
        }
        return n
}

func (n *SessionInfo) Set(key string, value interface{}) *SessionInfo {
        if n == nil {
                n = SessionInfoCreate(SessionInfo{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
                      n.RoomNumber = StringCreate(value.(string))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RollMarked":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.RollMarked = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "StaffPersonalLocalId":
    
                      n.StaffPersonalLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SessionDate":
                      n.SessionDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "FinishTime":
                      n.FinishTime = StringCreate(value.(string))
        }
        return n
}

func (n *NAPTestScoreSummary) Set(key string, value interface{}) *NAPTestScoreSummary {
        if n == nil {
                n = NAPTestScoreSummaryCreate(NAPTestScoreSummary{})
        }
        switch key {
    case "DomainJurisdictionAverage":
                      n.DomainJurisdictionAverage = FloatCreate(value.(float64))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DomainNationalAverage":
                      n.DomainNationalAverage = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "DomainTopNational60Percent":
                      n.DomainTopNational60Percent = FloatCreate(value.(float64))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "DomainBottomNational60Percent":
                      n.DomainBottomNational60Percent = FloatCreate(value.(float64))
    case "DomainSchoolAverage":
                      n.DomainSchoolAverage = FloatCreate(value.(float64))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *TimeTableCell) Set(key string, value interface{}) *TimeTableCell {
        if n == nil {
                n = TimeTableCellCreate(TimeTableCell{})
        }
        switch key {
    case "RoomInfoRefId":
                      n.RoomInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableLocalId":
    
                      n.TimeTableLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableRefId":
                      n.TimeTableRefId = StringCreate(value.(string))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ScheduledActivity) Set(key string, value interface{}) *ScheduledActivity {
        if n == nil {
                n = ScheduledActivityCreate(ScheduledActivity{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "TimeTableRefId":
                      n.TimeTableRefId = StringCreate(value.(string))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ActivityType":
      present := false
  for _, b := range AUCodeSetsScheduledActivityTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsScheduledActivityTypeType_values")
      }

                      n.ActivityType = ((*AUCodeSetsScheduledActivityTypeType)(StringCreate(value.(string))))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "Location":
                      n.Location = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "Override":
                      n.Override = ScheduledActivityOverrideTypeCreate(value.(ScheduledActivityOverrideType))
    case "StudentList":
                      n.StudentList = StudentsTypeCreate(value.(StudentsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ActivityDate":
                      n.ActivityDate = StringCreate(value.(string))
    case "TeachingGroupList":
                      n.TeachingGroupList = TeachingGroupListTypeCreate(value.(TeachingGroupListType))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "ActivityName":
                      n.ActivityName = StringCreate(value.(string))
    case "FinishTime":
                      n.FinishTime = StringCreate(value.(string))
    case "ActivityComment":
                      n.ActivityComment = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
        }
        return n
}

func (n *SoftwareRequirementType) Set(key string, value interface{}) *SoftwareRequirementType {
        if n == nil {
                n = SoftwareRequirementTypeCreate(SoftwareRequirementType{})
        }
        switch key {
    case "Vendor":
                      n.Vendor = StringCreate(value.(string))
    case "OS":
                      n.OS = StringCreate(value.(string))
    case "Version":
                      n.Version = StringCreate(value.(string))
    case "SoftwareTitle":
                      n.SoftwareTitle = StringCreate(value.(string))
        }
        return n
}

func (n *ScoreDescriptionType) Set(key string, value interface{}) *ScoreDescriptionType {
        if n == nil {
                n = ScoreDescriptionTypeCreate(ScoreDescriptionType{})
        }
        switch key {
    case "Descriptor":
                      n.Descriptor = StringCreate(value.(string))
    case "ScoreValue":
                      n.ScoreValue = FloatCreate(value.(float64))
        }
        return n
}

func (n *NAPTestletResponseItemType) Set(key string, value interface{}) *NAPTestletResponseItemType {
        if n == nil {
                n = NAPTestletResponseItemTypeCreate(NAPTestletResponseItemType{})
        }
        switch key {
    case "LapsedTimeItem":
                      n.LapsedTimeItem = StringCreate(value.(string))
    case "Response":
                      n.Response = StringCreate(value.(string))
    case "Score":
                      n.Score = FloatCreate(value.(float64))
    case "NAPTestItemRefId":
                      n.NAPTestItemRefId = StringCreate(value.(string))
    case "SubscoreList":
                      n.SubscoreList = NAPSubscoreListTypeCreate(value.(NAPSubscoreListType))
    case "ResponseCorrectness":
      present := false
  for _, b := range AUCodeSetsNAPResponseCorrectnessType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPResponseCorrectnessType_values")
      }

                      n.ResponseCorrectness = ((*AUCodeSetsNAPResponseCorrectnessType)(StringCreate(value.(string))))
    case "SequenceNumber":
                      n.SequenceNumber = IntCreate(value.(int))
    case "ItemWeight":
                      n.ItemWeight = FloatCreate(value.(float64))
    case "NAPTestItemLocalId":
    
                      n.NAPTestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *CalendarDate) Set(key string, value interface{}) *CalendarDate {
        if n == nil {
                n = CalendarDateCreate(CalendarDate{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "TeacherAttendance":
                      n.TeacherAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "CalendarDateNumber":
                      n.CalendarDateNumber = IntCreate(value.(int))
    case "AdministratorAttendance":
                      n.AdministratorAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CalendarSummaryRefId":
                      n.CalendarSummaryRefId = StringCreate(value.(string))
    case "CalendarDateType":
                      n.CalendarDateType = CalendarDateInfoTypeCreate(value.(CalendarDateInfoType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "CalendarDateRefId":
                      n.CalendarDateRefId = StringCreate(value.(string))
    case "StudentAttendance":
                      n.StudentAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        }
        return n
}

func (n *NAPTestlet) Set(key string, value interface{}) *NAPTestlet {
        if n == nil {
                n = NAPTestletCreate(NAPTestlet{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TestletContent":
                      n.TestletContent = NAPTestletContentTypeCreate(value.(NAPTestletContentType))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TestItemList":
                      n.TestItemList = NAPTestItemListTypeCreate(value.(NAPTestItemListType))
        }
        return n
}

func (n *FineInfoType) Set(key string, value interface{}) *FineInfoType {
        if n == nil {
                n = FineInfoTypeCreate(FineInfoType{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Reference":
                      n.Reference = StringCreate(value.(string))
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "Assessed":
                      n.Assessed = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *FQItemType) Set(key string, value interface{}) *FQItemType {
        if n == nil {
                n = FQItemTypeCreate(FQItemType{})
        }
        switch key {
    case "DioceseAmount":
                      n.DioceseAmount = FloatCreate(value.(float64))
    case "FQComments":
                      n.FQComments = StringCreate(value.(string))
    case "BoardingAmount":
                      n.BoardingAmount = FloatCreate(value.(float64))
    case "SystemAmount":
                      n.SystemAmount = FloatCreate(value.(float64))
    case "FQItemCode":
                      n.FQItemCode = StringCreate(value.(string))
    case "TuitionAmount":
                      n.TuitionAmount = FloatCreate(value.(float64))
        }
        return n
}

func (n *TestDisruptionType) Set(key string, value interface{}) *TestDisruptionType {
        if n == nil {
                n = TestDisruptionTypeCreate(TestDisruptionType{})
        }
        switch key {
    case "Event":
                      n.Event = StringCreate(value.(string))
        }
        return n
}

func (n *StudentSchoolEnrollment) Set(key string, value interface{}) *StudentSchoolEnrollment {
        if n == nil {
                n = StudentSchoolEnrollmentCreate(StudentSchoolEnrollment{})
        }
        switch key {
    case "RecordClosureReason":
                      n.RecordClosureReason = StringCreate(value.(string))
    case "InternationalStudent":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.InternationalStudent = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "StartedAtSchoolDate":
                      n.StartedAtSchoolDate = StringCreate(value.(string))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "BoardingStatus":
      present := false
  for _, b := range AUCodeSetsBoardingType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsBoardingType_values")
      }

                      n.BoardingStatus = ((*AUCodeSetsBoardingType)(StringCreate(value.(string))))
    case "FFPOS":
      present := false
  for _, b := range AUCodeSetsFFPOSStatusCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsFFPOSStatusCodeType_values")
      }

                      n.FFPOS = ((*AUCodeSetsFFPOSStatusCodeType)(StringCreate(value.(string))))
    case "ExitStatus":
                      n.ExitStatus = StudentExitStatusContainerTypeCreate(value.(StudentExitStatusContainerType))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "PublishingPermissionList":
                      n.PublishingPermissionList = PublishingPermissionListTypeCreate(value.(PublishingPermissionListType))
    case "DestinationSchoolName":
                      n.DestinationSchoolName = StringCreate(value.(string))
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
    case "PreviousSchool":
    
                      n.PreviousSchool = ((*LocalIdType)(StringCreate(value.(string))))
    case "ExitType":
                      n.ExitType = StudentExitContainerTypeCreate(value.(StudentExitContainerType))
    case "CensusAge":
                      n.CensusAge = FloatCreate(value.(float64))
    case "Calendar":
                      n.Calendar = StudentSchoolEnrollment_CalendarCreate(value.(StudentSchoolEnrollment_Calendar))
    case "Homeroom":
                      n.Homeroom = StudentSchoolEnrollment_HomeroomCreate(value.(StudentSchoolEnrollment_Homeroom))
    case "PreviousSchoolName":
                      n.PreviousSchoolName = StringCreate(value.(string))
    case "PromotionInfo":
                      n.PromotionInfo = PromotionInfoTypeCreate(value.(PromotionInfoType))
    case "TimeFrame":
      present := false
  for _, b := range AUCodeSetsEnrollmentTimeFrameType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEnrollmentTimeFrameType_values")
      }

                      n.TimeFrame = ((*AUCodeSetsEnrollmentTimeFrameType)(StringCreate(value.(string))))
    case "Advisor":
                      n.Advisor = StudentSchoolEnrollment_AdvisorCreate(value.(StudentSchoolEnrollment_Advisor))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "EntryDate":
                      n.EntryDate = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "DestinationSchool":
    
                      n.DestinationSchool = ((*LocalIdType)(StringCreate(value.(string))))
    case "ClassCode":
                      n.ClassCode = StringCreate(value.(string))
    case "IndividualLearningPlan":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.IndividualLearningPlan = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "House":
                      n.House = StringCreate(value.(string))
    case "CatchmentStatus":
                      n.CatchmentStatus = CatchmentStatusContainerTypeCreate(value.(CatchmentStatusContainerType))
    case "StudentGroupList":
                      n.StudentGroupList = StudentGroupListTypeCreate(value.(StudentGroupListType))
    case "EntryType":
                      n.EntryType = StudentEntryContainerTypeCreate(value.(StudentEntryContainerType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "MembershipType":
      present := false
  for _, b := range AUCodeSetsSchoolEnrollmentTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolEnrollmentTypeType_values")
      }

                      n.MembershipType = ((*AUCodeSetsSchoolEnrollmentTypeType)(StringCreate(value.(string))))
    case "FTPTStatus":
      present := false
  for _, b := range AUCodeSetsFTPTStatusCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsFTPTStatusCodeType_values")
      }

                      n.FTPTStatus = ((*AUCodeSetsFTPTStatusCodeType)(StringCreate(value.(string))))
    case "ACARASchoolId":
    
                      n.ACARASchoolId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DistanceEducationStudent":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.DistanceEducationStudent = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ReportingSchool":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ReportingSchool = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "ExitDate":
                      n.ExitDate = StringCreate(value.(string))
    case "Counselor":
                      n.Counselor = StudentSchoolEnrollment_CounselorCreate(value.(StudentSchoolEnrollment_Counselor))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StudentSubjectChoiceList":
                      n.StudentSubjectChoiceList = StudentSubjectChoiceListTypeCreate(value.(StudentSubjectChoiceListType))
        }
        return n
}

func (n *AssociatedObjectsType_AssociatedObject) Set(key string, value interface{}) *AssociatedObjectsType_AssociatedObject {
        if n == nil {
                n = AssociatedObjectsType_AssociatedObjectCreate(AssociatedObjectsType_AssociatedObject{})
        }
        switch key {
    case "SIF_RefObject":
    
                      n.SIF_RefObject = ((*ObjectNameType)(StringCreate(value.(string))))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentPersonal) Set(key string, value interface{}) *StudentPersonal {
        if n == nil {
                n = StudentPersonalCreate(StudentPersonal{})
        }
        switch key {
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "Disability":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.Disability = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PrePrimaryEducation":
                      n.PrePrimaryEducation = StringCreate(value.(string))
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "OnTimeGraduationYear":
    
                      n.OnTimeGraduationYear = ((*OnTimeGraduationYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AlertMessages":
                      n.AlertMessages = AlertMessagesTypeCreate(value.(AlertMessagesType))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "EconomicDisadvantage":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.EconomicDisadvantage = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ESLDateAssessed":
                      n.ESLDateAssessed = StringCreate(value.(string))
    case "ESLSupport":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ESLSupport = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Sensitive":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.Sensitive = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "GraduationDate":
    
                      n.GraduationDate = ((*GraduationDateType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "GiftedTalented":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.GiftedTalented = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "MedicalAlertMessages":
                      n.MedicalAlertMessages = MedicalAlertMessagesTypeCreate(value.(MedicalAlertMessagesType))
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
    case "ProjectedGraduationYear":
    
                      n.ProjectedGraduationYear = ((*ProjectedGraduationYearType)(StringCreate(value.(string))))
    case "HomeSchooledStudent":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.HomeSchooledStudent = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "IntegrationAide":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.IntegrationAide = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PrePrimaryEducationHours":
      present := false
  for _, b := range AUCodeSetsPrePrimaryHoursType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsPrePrimaryHoursType_values")
      }

                      n.PrePrimaryEducationHours = ((*AUCodeSetsPrePrimaryHoursType)(StringCreate(value.(string))))
    case "ESL":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ESL = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "MostRecent":
                      n.MostRecent = StudentMostRecentContainerTypeCreate(value.(StudentMostRecentContainerType))
    case "FirstAUSchoolEnrollment":
                      n.FirstAUSchoolEnrollment = StringCreate(value.(string))
    case "OfflineDelivery":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.OfflineDelivery = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "YoungCarersRole":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.YoungCarersRole = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AcceptableUsePolicy":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AcceptableUsePolicy = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "EducationSupport":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.EducationSupport = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
        }
        return n
}

func (n *SystemRole) Set(key string, value interface{}) *SystemRole {
        if n == nil {
                n = SystemRoleCreate(SystemRole{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_RefId":
                      n.SIF_RefId = SystemRole_SIF_RefIdCreate(value.(SystemRole_SIF_RefId))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SystemContextList":
                      n.SystemContextList = SystemRole_SystemContextListCreate(value.(SystemRole_SystemContextList))
        }
        return n
}

func (n *AGReportingObjectResponseType) Set(key string, value interface{}) *AGReportingObjectResponseType {
        if n == nil {
                n = AGReportingObjectResponseTypeCreate(AGReportingObjectResponseType{})
        }
        switch key {
    case "AGRuleList":
                      n.AGRuleList = AGRuleListTypeCreate(value.(AGRuleListType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "AGSubmissionStatusCode":
      present := false
  for _, b := range AUCodeSetsAGSubmissionStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAGSubmissionStatusType_values")
      }

                      n.AGSubmissionStatusCode = ((*AUCodeSetsAGSubmissionStatusType)(StringCreate(value.(string))))
    case "HTTPStatusCode":
                      n.HTTPStatusCode = StringCreate(value.(string))
    case "SubmittedRefId":
                      n.SubmittedRefId = StringCreate(value.(string))
    case "ErrorText":
                      n.ErrorText = StringCreate(value.(string))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "SIFRefId":
                      n.SIFRefId = StringCreate(value.(string))
        }
        return n
}

func (n *OtherIdType) Set(key string, value interface{}) *OtherIdType {
        if n == nil {
                n = OtherIdTypeCreate(OtherIdType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *PersonInvolvementType) Set(key string, value interface{}) *PersonInvolvementType {
        if n == nil {
                n = PersonInvolvementTypeCreate(PersonInvolvementType{})
        }
        switch key {
    case "PersonRefId":
                      n.PersonRefId = PersonInvolvementType_PersonRefIdCreate(value.(PersonInvolvementType_PersonRefId))
    case "HowInvolved":
                      n.HowInvolved = StringCreate(value.(string))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
        }
        return n
}

func (n *NAPTest) Set(key string, value interface{}) *NAPTest {
        if n == nil {
                n = NAPTestCreate(NAPTest{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TestContent":
                      n.TestContent = NAPTestContentTypeCreate(value.(NAPTestContentType))
        }
        return n
}

func (n *TechnicalRequirementsType) Set(key string, value interface{}) *TechnicalRequirementsType {
        if n == nil {
                n = TechnicalRequirementsTypeCreate(TechnicalRequirementsType{})
        }
        switch key {
    case "TechnicalRequirement":
                      n.TechnicalRequirement = StringCreate(value.(string))
        }
        return n
}

func (n *FQReportingType) Set(key string, value interface{}) *FQReportingType {
        if n == nil {
                n = FQReportingTypeCreate(FQReportingType{})
        }
        switch key {
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "FQItemList":
                      n.FQItemList = FQItemListTypeCreate(value.(FQItemListType))
    case "FQContextualQuestionList":
                      n.FQContextualQuestionList = FQContextualQuestionListTypeCreate(value.(FQContextualQuestionListType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "AGRuleList":
                      n.AGRuleList = AGRuleListTypeCreate(value.(AGRuleListType))
        }
        return n
}

func (n *AGRuleType) Set(key string, value interface{}) *AGRuleType {
        if n == nil {
                n = AGRuleTypeCreate(AGRuleType{})
        }
        switch key {
    case "AGRuleStatus":
                      n.AGRuleStatus = StringCreate(value.(string))
    case "AGRuleResponse":
                      n.AGRuleResponse = StringCreate(value.(string))
    case "AGRuleCode":
                      n.AGRuleCode = StringCreate(value.(string))
    case "AGRuleComment":
                      n.AGRuleComment = StringCreate(value.(string))
        }
        return n
}

func (n *StudentAttendanceSummary) Set(key string, value interface{}) *StudentAttendanceSummary {
        if n == nil {
                n = StudentAttendanceSummaryCreate(StudentAttendanceSummary{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DaysAttended":
                      n.DaysAttended = FloatCreate(value.(float64))
    case "StudentAttendanceSummaryRefId":
                      n.StudentAttendanceSummaryRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "ExcusedAbsences":
                      n.ExcusedAbsences = FloatCreate(value.(float64))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "DaysInMembership":
                      n.DaysInMembership = FloatCreate(value.(float64))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "StartDay":
                      n.StartDay = IntCreate(value.(int))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "DaysTardy":
                      n.DaysTardy = FloatCreate(value.(float64))
    case "UnexcusedAbsences":
                      n.UnexcusedAbsences = FloatCreate(value.(float64))
    case "EndDay":
                      n.EndDay = IntCreate(value.(int))
        }
        return n
}

func (n *Journal) Set(key string, value interface{}) *Journal {
        if n == nil {
                n = JournalCreate(Journal{})
        }
        switch key {
    case "OriginatingTransactionRefId":
                      n.OriginatingTransactionRefId = Journal_OriginatingTransactionRefIdCreate(value.(Journal_OriginatingTransactionRefId))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CreatedDate":
                      n.CreatedDate = StringCreate(value.(string))
    case "ApprovedBy":
                      n.ApprovedBy = StringCreate(value.(string))
    case "GSTCodeReplacement":
                      n.GSTCodeReplacement = StringCreate(value.(string))
    case "CreditAccountCode":
                      n.CreditAccountCode = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ApprovedDate":
                      n.ApprovedDate = StringCreate(value.(string))
    case "Note":
                      n.Note = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "JournalAdjustmentList":
                      n.JournalAdjustmentList = JournalAdjustmentListTypeCreate(value.(JournalAdjustmentListType))
    case "CreatedBy":
                      n.CreatedBy = StringCreate(value.(string))
    case "DebitFinancialAccountRefId":
                      n.DebitFinancialAccountRefId = StringCreate(value.(string))
    case "CreditFinancialAccountRefId":
                      n.CreditFinancialAccountRefId = StringCreate(value.(string))
    case "GSTCodeOriginal":
                      n.GSTCodeOriginal = StringCreate(value.(string))
    case "DebitAccountCode":
                      n.DebitAccountCode = StringCreate(value.(string))
        }
        return n
}

func (n *LearningStandardDocument) Set(key string, value interface{}) *LearningStandardDocument {
        if n == nil {
                n = LearningStandardDocumentCreate(LearningStandardDocument{})
        }
        switch key {
    case "Authors":
                      n.Authors = AuthorsTypeCreate(value.(AuthorsType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Copyright":
                      n.Copyright = CopyRightContainerTypeCreate(value.(CopyRightContainerType))
    case "RelatedLearningStandards":
                      n.RelatedLearningStandards = LearningStandardsDocumentTypeCreate(value.(LearningStandardsDocumentType))
    case "RepositoryDate":
                      n.RepositoryDate = StringCreate(value.(string))
    case "RichDescription":
                      n.RichDescription = AbstractContentElementTypeCreate(value.(AbstractContentElementType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "LocalAdoptionDate":
                      n.LocalAdoptionDate = StringCreate(value.(string))
    case "LocalArchiveDate":
                      n.LocalArchiveDate = StringCreate(value.(string))
    case "Source":
                      n.Source = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "EndOfLifeDate":
                      n.EndOfLifeDate = StringCreate(value.(string))
    case "OrganizationContactPoint":
                      n.OrganizationContactPoint = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "DocumentStatus":
                      n.DocumentStatus = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "DocumentDate":
                      n.DocumentDate = StringCreate(value.(string))
    case "Organizations":
                      n.Organizations = OrganizationsTypeCreate(value.(OrganizationsType))
    case "LearningStandardItemRefId":
                      n.LearningStandardItemRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SubjectAreas":
                      n.SubjectAreas = ACStrandAreaListTypeCreate(value.(ACStrandAreaListType))
        }
        return n
}

func (n *LanguageBaseType) Set(key string, value interface{}) *LanguageBaseType {
        if n == nil {
                n = LanguageBaseTypeCreate(LanguageBaseType{})
        }
        switch key {
    case "Dialect":
                      n.Dialect = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "LanguageType":
      present := false
  for _, b := range AUCodeSetsLanguageTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsLanguageTypeType_values")
      }

                      n.LanguageType = ((*AUCodeSetsLanguageTypeType)(StringCreate(value.(string))))
    case "Code":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType_values")
      }

                      n.Code = ((*AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType)(StringCreate(value.(string))))
        }
        return n
}

func (n *AbstractContentPackageType_TextData) Set(key string, value interface{}) *AbstractContentPackageType_TextData {
        if n == nil {
                n = AbstractContentPackageType_TextDataCreate(AbstractContentPackageType_TextData{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *ProgramAvailabilityType) Set(key string, value interface{}) *ProgramAvailabilityType {
        if n == nil {
                n = ProgramAvailabilityTypeCreate(ProgramAvailabilityType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSets0211ProgramAvailabilityType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSets0211ProgramAvailabilityType_values")
      }

                      n.Code = ((*AUCodeSets0211ProgramAvailabilityType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *TeachingGroupScheduleType) Set(key string, value interface{}) *TeachingGroupScheduleType {
        if n == nil {
                n = TeachingGroupScheduleTypeCreate(TeachingGroupScheduleType{})
        }
        switch key {
    case "SchoolCourseLocalId":
    
                      n.SchoolCourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "MinClassSize":
                      n.MinClassSize = IntCreate(value.(int))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupPeriodList":
                      n.TeachingGroupPeriodList = TeachingGroupPeriodListTypeCreate(value.(TeachingGroupPeriodListType))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "TeacherList":
                      n.TeacherList = TeacherListTypeCreate(value.(TeacherListType))
    case "Block":
                      n.Block = StringCreate(value.(string))
    case "CurriculumLevel":
                      n.CurriculumLevel = StringCreate(value.(string))
    case "EditorGUID":
    
                      n.EditorGUID = ((*RefIdType)(StringCreate(value.(string))))
    case "MaxClassSize":
                      n.MaxClassSize = IntCreate(value.(int))
    case "GroupType":
                      n.GroupType = StringCreate(value.(string))
    case "Sett":
                      n.Sett = StringCreate(value.(string))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "SchoolInfoRefId":
    
                      n.SchoolInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LongName":
                      n.LongName = StringCreate(value.(string))
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentList":
                      n.StudentList = StudentListTypeCreate(value.(StudentListType))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
        }
        return n
}

func (n *WellbeingPlanType) Set(key string, value interface{}) *WellbeingPlanType {
        if n == nil {
                n = WellbeingPlanTypeCreate(WellbeingPlanType{})
        }
        switch key {
    case "PersonalisedPlanRefId":
                      n.PersonalisedPlanRefId = StringCreate(value.(string))
    case "PlanNotes":
                      n.PlanNotes = StringCreate(value.(string))
        }
        return n
}

func (n *EnglishProficiencyType) Set(key string, value interface{}) *EnglishProficiencyType {
        if n == nil {
                n = EnglishProficiencyTypeCreate(EnglishProficiencyType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsEnglishProficiencyType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsEnglishProficiencyType_values")
      }

                      n.Code = ((*AUCodeSetsEnglishProficiencyType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *StaffSubjectType) Set(key string, value interface{}) *StaffSubjectType {
        if n == nil {
                n = StaffSubjectTypeCreate(StaffSubjectType{})
        }
        switch key {
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *StatsCohortType) Set(key string, value interface{}) *StatsCohortType {
        if n == nil {
                n = StatsCohortTypeCreate(StatsCohortType{})
        }
        switch key {
    case "DaysInReferencePeriod":
                      n.DaysInReferencePeriod = IntCreate(value.(int))
    case "StatsIndigenousStudentType":
                      n.StatsIndigenousStudentType = StringCreate(value.(string))
    case "PossibleSchoolDays":
                      n.PossibleSchoolDays = IntCreate(value.(int))
    case "AttendanceLess90Percent":
                      n.AttendanceLess90Percent = IntCreate(value.(int))
    case "PossibleSchoolDaysGT90PercentAttendance":
                      n.PossibleSchoolDaysGT90PercentAttendance = IntCreate(value.(int))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "StatsCohortId":
    
                      n.StatsCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AttendanceGTE90Percent":
                      n.AttendanceGTE90Percent = IntCreate(value.(int))
    case "AttendanceDays":
                      n.AttendanceDays = FloatCreate(value.(float64))
        }
        return n
}

func (n *ElectronicIdType) Set(key string, value interface{}) *ElectronicIdType {
        if n == nil {
                n = ElectronicIdTypeCreate(ElectronicIdType{})
        }
        switch key {
    case "Type":
      present := false
  for _, b := range AUCodeSetsElectronicIdTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsElectronicIdTypeType_values")
      }

                      n.Type = ((*AUCodeSetsElectronicIdTypeType)(StringCreate(value.(string))))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *NAPTestletContentType) Set(key string, value interface{}) *NAPTestletContentType {
        if n == nil {
                n = NAPTestletContentTypeCreate(NAPTestletContentType{})
        }
        switch key {
    case "TestletMaximumScore":
                      n.TestletMaximumScore = FloatCreate(value.(float64))
    case "TestletName":
                      n.TestletName = StringCreate(value.(string))
    case "NAPTestletLocalId":
    
                      n.NAPTestletLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocationInStage":
                      n.LocationInStage = IntCreate(value.(int))
    case "Node":
                      n.Node = StringCreate(value.(string))
        }
        return n
}

func (n *TeachingGroupStudentType) Set(key string, value interface{}) *TeachingGroupStudentType {
        if n == nil {
                n = TeachingGroupStudentTypeCreate(TeachingGroupStudentType{})
        }
        switch key {
    case "StudentLocalId":
    
                      n.StudentLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "Name":
                      n.Name = NameOfRecordTypeCreate(value.(NameOfRecordType))
        }
        return n
}

func (n *NAPTestItem) Set(key string, value interface{}) *NAPTestItem {
        if n == nil {
                n = NAPTestItemCreate(NAPTestItem{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TestItemContent":
                      n.TestItemContent = NAPTestItemContentTypeCreate(value.(NAPTestItemContentType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        }
        return n
}

func (n *ContactType) Set(key string, value interface{}) *ContactType {
        if n == nil {
                n = ContactTypeCreate(ContactType{})
        }
        switch key {
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
    case "Email":
                      n.Email = EmailTypeCreate(value.(EmailType))
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
        }
        return n
}

func (n *AbstractContentPackageType_BinaryData) Set(key string, value interface{}) *AbstractContentPackageType_BinaryData {
        if n == nil {
                n = AbstractContentPackageType_BinaryDataCreate(AbstractContentPackageType_BinaryData{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentActivityParticipation) Set(key string, value interface{}) *StudentActivityParticipation {
        if n == nil {
                n = StudentActivityParticipationCreate(StudentActivityParticipation{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "RecognitionList":
                      n.RecognitionList = RecognitionListTypeCreate(value.(RecognitionListType))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "StudentActivityInfoRefId":
                      n.StudentActivityInfoRefId = StringCreate(value.(string))
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ParticipationComment":
                      n.ParticipationComment = StringCreate(value.(string))
        }
        return n
}

func (n *SchoolContactType) Set(key string, value interface{}) *SchoolContactType {
        if n == nil {
                n = SchoolContactTypeCreate(SchoolContactType{})
        }
        switch key {
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
    case "PublishInDirectory":
    
                      n.PublishInDirectory = ((*PublishInDirectoryType)(StringCreate(value.(string))))
        }
        return n
}

func (n *RelatedLearningStandardItemRefIdType) Set(key string, value interface{}) *RelatedLearningStandardItemRefIdType {
        if n == nil {
                n = RelatedLearningStandardItemRefIdTypeCreate(RelatedLearningStandardItemRefIdType{})
        }
        switch key {
    case "RelationshipType":
                      n.RelationshipType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *PlanRequiredContainerType) Set(key string, value interface{}) *PlanRequiredContainerType {
        if n == nil {
                n = PlanRequiredContainerTypeCreate(PlanRequiredContainerType{})
        }
        switch key {
    case "PlanRequiredList":
                      n.PlanRequiredList = PlanRequiredListTypeCreate(value.(PlanRequiredListType))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
        }
        return n
}

func (n *GridLocationType) Set(key string, value interface{}) *GridLocationType {
        if n == nil {
                n = GridLocationTypeCreate(GridLocationType{})
        }
        switch key {
    case "Longitude":
                      n.Longitude = FloatCreate(value.(float64))
    case "Latitude":
                      n.Latitude = FloatCreate(value.(float64))
        }
        return n
}

func (n *MedicationType) Set(key string, value interface{}) *MedicationType {
        if n == nil {
                n = MedicationTypeCreate(MedicationType{})
        }
        switch key {
    case "Method":
                      n.Method = StringCreate(value.(string))
    case "MedicationName":
                      n.MedicationName = StringCreate(value.(string))
    case "Dosage":
                      n.Dosage = StringCreate(value.(string))
    case "AdministrationInformation":
                      n.AdministrationInformation = StringCreate(value.(string))
    case "Frequency":
                      n.Frequency = StringCreate(value.(string))
        }
        return n
}

func (n *VendorInfo) Set(key string, value interface{}) *VendorInfo {
        if n == nil {
                n = VendorInfoCreate(VendorInfo{})
        }
        switch key {
    case "PaymentTerms":
                      n.PaymentTerms = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "BPay":
                      n.BPay = StringCreate(value.(string))
    case "ABN":
                      n.ABN = StringCreate(value.(string))
    case "CustomerId":
                      n.CustomerId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "BSB":
                      n.BSB = StringCreate(value.(string))
    case "RegisteredForGST":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.RegisteredForGST = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "AccountName":
                      n.AccountName = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AccountNumber":
                      n.AccountNumber = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
        }
        return n
}

func (n *AttendanceInfoType) Set(key string, value interface{}) *AttendanceInfoType {
        if n == nil {
                n = AttendanceInfoTypeCreate(AttendanceInfoType{})
        }
        switch key {
    case "CountsTowardAttendance":
                      n.CountsTowardAttendance = StringCreate(value.(string))
    case "AttendanceValue":
                      n.AttendanceValue = FloatCreate(value.(float64))
        }
        return n
}

func (n *SystemRole_SIF_RefId) Set(key string, value interface{}) *SystemRole_SIF_RefId {
        if n == nil {
                n = SystemRole_SIF_RefIdCreate(SystemRole_SIF_RefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *StandardIdentifierType) Set(key string, value interface{}) *StandardIdentifierType {
        if n == nil {
                n = StandardIdentifierTypeCreate(StandardIdentifierType{})
        }
        switch key {
    case "StandardNumber":
                      n.StandardNumber = StringCreate(value.(string))
    case "YearCreated":
                      n.YearCreated = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "Organization":
                      n.Organization = StringCreate(value.(string))
    case "ACStrandSubjectArea":
                      n.ACStrandSubjectArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "Benchmark":
                      n.Benchmark = StringCreate(value.(string))
    case "IndicatorNumber":
                      n.IndicatorNumber = StringCreate(value.(string))
    case "AlternateIdentificationCodes":
                      n.AlternateIdentificationCodes = AlternateIdentificationCodeListTypeCreate(value.(AlternateIdentificationCodeListType))
        }
        return n
}

func (n *LearningStandardItem) Set(key string, value interface{}) *LearningStandardItem {
        if n == nil {
                n = LearningStandardItemCreate(LearningStandardItem{})
        }
        switch key {
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "Statements":
                      n.Statements = StatementsTypeCreate(value.(StatementsType))
    case "StandardSettingBody":
                      n.StandardSettingBody = StandardsSettingBodyTypeCreate(value.(StandardsSettingBodyType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Level5":
                      n.Level5 = StringCreate(value.(string))
    case "LearningStandardDocumentRefId":
                      n.LearningStandardDocumentRefId = StringCreate(value.(string))
    case "StatementCodes":
                      n.StatementCodes = StatementCodesTypeCreate(value.(StatementCodesType))
    case "StandardHierarchyLevel":
                      n.StandardHierarchyLevel = StandardHierarchyLevelTypeCreate(value.(StandardHierarchyLevelType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StandardIdentifier":
                      n.StandardIdentifier = StandardIdentifierTypeCreate(value.(StandardIdentifierType))
    case "Level4":
                      n.Level4 = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PredecessorItems":
                      n.PredecessorItems = LearningStandardsTypeCreate(value.(LearningStandardsType))
    case "ACStrandSubjectArea":
                      n.ACStrandSubjectArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Resources":
                      n.Resources = LResourcesTypeCreate(value.(LResourcesType))
    case "RelatedLearningStandardItems":
                      n.RelatedLearningStandardItems = RelatedLearningStandardItemRefIdListTypeCreate(value.(RelatedLearningStandardItemRefIdListType))
        }
        return n
}

func (n *StudentContactRelationship) Set(key string, value interface{}) *StudentContactRelationship {
        if n == nil {
                n = StudentContactRelationshipCreate(StudentContactRelationship{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "ParentRelationshipStatus":
                      n.ParentRelationshipStatus = StringCreate(value.(string))
    case "HouseholdList":
                      n.HouseholdList = HouseholdListTypeCreate(value.(HouseholdListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ContactSequence":
                      n.ContactSequence = IntCreate(value.(int))
    case "StudentContactPersonalRefId":
    
                      n.StudentContactPersonalRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "MainlySpeaksEnglishAtHome":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.MainlySpeaksEnglishAtHome = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Relationship":
                      n.Relationship = RelationshipTypeCreate(value.(RelationshipType))
    case "ContactFlags":
                      n.ContactFlags = ContactFlagsTypeCreate(value.(ContactFlagsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ContactSequenceSource":
      present := false
  for _, b := range AUCodeSetsSourceCodeTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSourceCodeTypeType_values")
      }

                      n.ContactSequenceSource = ((*AUCodeSetsSourceCodeTypeType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentContactRelationshipRefId":
                      n.StudentContactRelationshipRefId = StringCreate(value.(string))
    case "StudentPersonalRefId":
    
                      n.StudentPersonalRefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ResourceUsage_SIF_RefId) Set(key string, value interface{}) *ResourceUsage_SIF_RefId {
        if n == nil {
                n = ResourceUsage_SIF_RefIdCreate(ResourceUsage_SIF_RefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *StudentGroupType) Set(key string, value interface{}) *StudentGroupType {
        if n == nil {
                n = StudentGroupTypeCreate(StudentGroupType{})
        }
        switch key {
    case "GroupDescription":
                      n.GroupDescription = StringCreate(value.(string))
    case "GroupLocalId":
    
                      n.GroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "GroupCategory":
      present := false
  for _, b := range AUCodeSetsGroupCategoryCodeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsGroupCategoryCodeType_values")
      }

                      n.GroupCategory = ((*AUCodeSetsGroupCategoryCodeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *IdentityAssertionsType_IdentityAssertion) Set(key string, value interface{}) *IdentityAssertionsType_IdentityAssertion {
        if n == nil {
                n = IdentityAssertionsType_IdentityAssertionCreate(IdentityAssertionsType_IdentityAssertion{})
        }
        switch key {
    case "SchemaName":
                      n.SchemaName = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *ReligiousEventType) Set(key string, value interface{}) *ReligiousEventType {
        if n == nil {
                n = ReligiousEventTypeCreate(ReligiousEventType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
        }
        return n
}

func (n *FinancialAccount) Set(key string, value interface{}) *FinancialAccount {
        if n == nil {
                n = FinancialAccountCreate(FinancialAccount{})
        }
        switch key {
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "CreationTime":
                      n.CreationTime = StringCreate(value.(string))
    case "AccountNumber":
                      n.AccountNumber = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ParentAccountRefId":
                      n.ParentAccountRefId = StringCreate(value.(string))
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "ClassType":
                      n.ClassType = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
        }
        return n
}

func (n *CheckoutInfoType) Set(key string, value interface{}) *CheckoutInfoType {
        if n == nil {
                n = CheckoutInfoTypeCreate(CheckoutInfoType{})
        }
        switch key {
    case "RenewalCount":
                      n.RenewalCount = IntCreate(value.(int))
    case "CheckedOutOn":
                      n.CheckedOutOn = StringCreate(value.(string))
    case "ReturnBy":
                      n.ReturnBy = StringCreate(value.(string))
        }
        return n
}

func (n *CopyRightContainerType) Set(key string, value interface{}) *CopyRightContainerType {
        if n == nil {
                n = CopyRightContainerTypeCreate(CopyRightContainerType{})
        }
        switch key {
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "Holder":
                      n.Holder = StringCreate(value.(string))
        }
        return n
}

func (n *StudentExitContainerType) Set(key string, value interface{}) *StudentExitContainerType {
        if n == nil {
                n = StudentExitContainerTypeCreate(StudentExitContainerType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsExitWithdrawalTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsExitWithdrawalTypeType_values")
      }

                      n.Code = ((*AUCodeSetsExitWithdrawalTypeType)(StringCreate(value.(string))))
        }
        return n
}

func (n *StatsCohortYearLevelType) Set(key string, value interface{}) *StatsCohortYearLevelType {
        if n == nil {
                n = StatsCohortYearLevelTypeCreate(StatsCohortYearLevelType{})
        }
        switch key {
    case "StatsCohortList":
                      n.StatsCohortList = StatsCohortListTypeCreate(value.(StatsCohortListType))
    case "CohortYearLevel":
                      n.CohortYearLevel = YearLevelTypeCreate(value.(YearLevelType))
        }
        return n
}

func (n *LearningResource) Set(key string, value interface{}) *LearningResource {
        if n == nil {
                n = LearningResourceCreate(LearningResource{})
        }
        switch key {
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "Author":
                      n.Author = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "MediaTypes":
                      n.MediaTypes = MediaTypesTypeCreate(value.(MediaTypesType))
    case "Approvals":
                      n.Approvals = ApprovalsTypeCreate(value.(ApprovalsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Components":
                      n.Components = ComponentsTypeCreate(value.(ComponentsType))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Location":
                      n.Location = LearningResource_LocationCreate(value.(LearningResource_Location))
    case "LearningStandards":
                      n.LearningStandards = LearningStandardsTypeCreate(value.(LearningStandardsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Status":
                      n.Status = StringCreate(value.(string))
    case "Contacts":
                      n.Contacts = ContactsTypeCreate(value.(ContactsType))
    case "LearningResourcePackageRefId":
                      n.LearningResourcePackageRefId = StringCreate(value.(string))
    case "UseAgreement":
                      n.UseAgreement = StringCreate(value.(string))
    case "Evaluations":
                      n.Evaluations = EvaluationsTypeCreate(value.(EvaluationsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SubjectAreas":
                      n.SubjectAreas = ACStrandAreaListTypeCreate(value.(ACStrandAreaListType))
    case "AgreementDate":
                      n.AgreementDate = StringCreate(value.(string))
        }
        return n
}

func (n *AssignmentScoreType) Set(key string, value interface{}) *AssignmentScoreType {
        if n == nil {
                n = AssignmentScoreTypeCreate(AssignmentScoreType{})
        }
        switch key {
    case "Weight":
                      n.Weight = FloatCreate(value.(float64))
    case "GradingAssignmentScoreRefId":
                      n.GradingAssignmentScoreRefId = StringCreate(value.(string))
        }
        return n
}

func (n *StudentAttendanceCollection) Set(key string, value interface{}) *StudentAttendanceCollection {
        if n == nil {
                n = StudentAttendanceCollectionCreate(StudentAttendanceCollection{})
        }
        switch key {
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StudentAttendanceCollectionYear":
    
                      n.StudentAttendanceCollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StudentAttendanceCollectionReportingList":
                      n.StudentAttendanceCollectionReportingList = StudentAttendanceCollectionReportingListTypeCreate(value.(StudentAttendanceCollectionReportingListType))
        }
        return n
}

func (n *AbstractContentPackageType) Set(key string, value interface{}) *AbstractContentPackageType {
        if n == nil {
                n = AbstractContentPackageTypeCreate(AbstractContentPackageType{})
        }
        switch key {
    case "TextData":
                      n.TextData = AbstractContentPackageType_TextDataCreate(value.(AbstractContentPackageType_TextData))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Reference":
                      n.Reference = AbstractContentPackageType_ReferenceCreate(value.(AbstractContentPackageType_Reference))
    case "XMLData":
                      n.XMLData = AbstractContentPackageType_XMLDataCreate(value.(AbstractContentPackageType_XMLData))
    case "BinaryData":
                      n.BinaryData = AbstractContentPackageType_BinaryDataCreate(value.(AbstractContentPackageType_BinaryData))
        }
        return n
}

func (n *WellbeingAppeal) Set(key string, value interface{}) *WellbeingAppeal {
        if n == nil {
                n = WellbeingAppealCreate(WellbeingAppeal{})
        }
        switch key {
    case "AppealOutcome":
                      n.AppealOutcome = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "AppealNotes":
                      n.AppealNotes = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "AppealStatusCode":
      present := false
  for _, b := range AUCodeSetsWellbeingAppealStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingAppealStatusType_values")
      }

                      n.AppealStatusCode = ((*AUCodeSetsWellbeingAppealStatusType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "LocalAppealId":
    
                      n.LocalAppealId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PrincipalInfoType) Set(key string, value interface{}) *PrincipalInfoType {
        if n == nil {
                n = PrincipalInfoTypeCreate(PrincipalInfoType{})
        }
        switch key {
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "ContactTitle":
                      n.ContactTitle = StringCreate(value.(string))
    case "ContactName":
                      n.ContactName = NameOfRecordTypeCreate(value.(NameOfRecordType))
        }
        return n
}

func (n *NAPSubscoreType) Set(key string, value interface{}) *NAPSubscoreType {
        if n == nil {
                n = NAPSubscoreTypeCreate(NAPSubscoreType{})
        }
        switch key {
    case "SubscoreValue":
                      n.SubscoreValue = FloatCreate(value.(float64))
    case "SubscoreType":
                      n.SubscoreType = StringCreate(value.(string))
        }
        return n
}

func (n *StudentSchoolEnrollment_Calendar) Set(key string, value interface{}) *StudentSchoolEnrollment_Calendar {
        if n == nil {
                n = StudentSchoolEnrollment_CalendarCreate(StudentSchoolEnrollment_Calendar{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *EquipmentInfo_SIF_RefId) Set(key string, value interface{}) *EquipmentInfo_SIF_RefId {
        if n == nil {
                n = EquipmentInfo_SIF_RefIdCreate(EquipmentInfo_SIF_RefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *TotalEnrollmentsType) Set(key string, value interface{}) *TotalEnrollmentsType {
        if n == nil {
                n = TotalEnrollmentsTypeCreate(TotalEnrollmentsType{})
        }
        switch key {
    case "Girls":
                      n.Girls = StringCreate(value.(string))
    case "Boys":
                      n.Boys = StringCreate(value.(string))
    case "TotalStudents":
                      n.TotalStudents = StringCreate(value.(string))
        }
        return n
}

func (n *MarkerType) Set(key string, value interface{}) *MarkerType {
        if n == nil {
                n = MarkerTypeCreate(MarkerType{})
        }
        switch key {
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "Role":
                      n.Role = StringCreate(value.(string))
        }
        return n
}

func (n *CalendarDateInfoType) Set(key string, value interface{}) *CalendarDateInfoType {
        if n == nil {
                n = CalendarDateInfoTypeCreate(CalendarDateInfoType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsCalendarEventType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsCalendarEventType_values")
      }

                      n.Code = ((*AUCodeSetsCalendarEventType)(StringCreate(value.(string))))
        }
        return n
}

func (n *Journal_OriginatingTransactionRefId) Set(key string, value interface{}) *Journal_OriginatingTransactionRefId {
        if n == nil {
                n = Journal_OriginatingTransactionRefIdCreate(Journal_OriginatingTransactionRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *PromotionInfoType) Set(key string, value interface{}) *PromotionInfoType {
        if n == nil {
                n = PromotionInfoTypeCreate(PromotionInfoType{})
        }
        switch key {
    case "PromotionStatus":
                      n.PromotionStatus = StringCreate(value.(string))
        }
        return n
}

func (n *TeachingGroupTeacherType) Set(key string, value interface{}) *TeachingGroupTeacherType {
        if n == nil {
                n = TeachingGroupTeacherTypeCreate(TeachingGroupTeacherType{})
        }
        switch key {
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "Association":
                      n.Association = StringCreate(value.(string))
    case "Name":
                      n.Name = NameOfRecordTypeCreate(value.(NameOfRecordType))
        }
        return n
}

func (n *Activity_Evaluation) Set(key string, value interface{}) *Activity_Evaluation {
        if n == nil {
                n = Activity_EvaluationCreate(Activity_Evaluation{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "EvaluationType":
                      n.EvaluationType = StringCreate(value.(string))
        }
        return n
}

func (n *StaffMostRecentContainerType) Set(key string, value interface{}) *StaffMostRecentContainerType {
        if n == nil {
                n = StaffMostRecentContainerTypeCreate(StaffMostRecentContainerType{})
        }
        switch key {
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPLANClassList":
                      n.NAPLANClassList = NAPLANClassListTypeCreate(value.(NAPLANClassListType))
    case "LocalCampusId":
    
                      n.LocalCampusId = ((*LocalIdType)(StringCreate(value.(string))))
    case "HomeGroup":
                      n.HomeGroup = StringCreate(value.(string))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ResourceBooking) Set(key string, value interface{}) *ResourceBooking {
        if n == nil {
                n = ResourceBookingCreate(ResourceBooking{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "ResourceRefId":
                      n.ResourceRefId = ResourceBooking_ResourceRefIdCreate(value.(ResourceBooking_ResourceRefId))
    case "Booker":
                      n.Booker = StringCreate(value.(string))
    case "ToPeriod":
    
                      n.ToPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "FromPeriod":
    
                      n.FromPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "ResourceLocalId":
    
                      n.ResourceLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "KeepOld":
                      n.KeepOld = BoolCreate(value.(bool))
    case "Reason":
                      n.Reason = StringCreate(value.(string))
    case "FinishDateTime":
                      n.FinishDateTime = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        }
        return n
}

func (n *NAPStudentResponseSet) Set(key string, value interface{}) *NAPStudentResponseSet {
        if n == nil {
                n = NAPStudentResponseSetCreate(NAPStudentResponseSet{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "EquatingSampleFlag":
                      n.EquatingSampleFlag = StringCreate(value.(string))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ReportExclusionFlag":
                      n.ReportExclusionFlag = BoolCreate(value.(bool))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CalibrationSampleFlag":
                      n.CalibrationSampleFlag = StringCreate(value.(string))
    case "PathTakenForDomain":
                      n.PathTakenForDomain = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "PlatformStudentIdentifier":
    
                      n.PlatformStudentIdentifier = ((*LocalIdType)(StringCreate(value.(string))))
    case "DomainScore":
                      n.DomainScore = DomainScoreTypeCreate(value.(DomainScoreType))
    case "ParallelTest":
                      n.ParallelTest = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TestletList":
                      n.TestletList = NAPStudentResponseTestletListTypeCreate(value.(NAPStudentResponseTestletListType))
        }
        return n
}

func (n *BaseNameType) Set(key string, value interface{}) *BaseNameType {
        if n == nil {
                n = BaseNameTypeCreate(BaseNameType{})
        }
        switch key {
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "FamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PreferredFamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PreferredFamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
        }
        return n
}

func (n *StandardHierarchyLevelType) Set(key string, value interface{}) *StandardHierarchyLevelType {
        if n == nil {
                n = StandardHierarchyLevelTypeCreate(StandardHierarchyLevelType{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Number":
                      n.Number = IntCreate(value.(int))
        }
        return n
}

func (n *TeacherCoverType) Set(key string, value interface{}) *TeacherCoverType {
        if n == nil {
                n = TeacherCoverTypeCreate(TeacherCoverType{})
        }
        switch key {
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Supervision":
      present := false
  for _, b := range AUCodeSetsTeacherCoverSupervisionType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsTeacherCoverSupervisionType_values")
      }

                      n.Supervision = ((*AUCodeSetsTeacherCoverSupervisionType)(StringCreate(value.(string))))
    case "FinishTime":
                      n.FinishTime = StringCreate(value.(string))
    case "Credit":
      present := false
  for _, b := range AUCodeSetsTeacherCoverCreditType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsTeacherCoverCreditType_values")
      }

                      n.Credit = ((*AUCodeSetsTeacherCoverCreditType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "Weighting":
                      n.Weighting = FloatCreate(value.(float64))
        }
        return n
}

func (n *TimeTableScheduleCellType) Set(key string, value interface{}) *TimeTableScheduleCellType {
        if n == nil {
                n = TimeTableScheduleCellTypeCreate(TimeTableScheduleCellType{})
        }
        switch key {
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableLocalId":
    
                      n.TimeTableLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableScheduleCellLocalId":
    
                      n.TimeTableScheduleCellLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomInfoRefId":
                      n.RoomInfoRefId = StringCreate(value.(string))
    case "TeachingGroupGUID":
                      n.TeachingGroupGUID = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
        }
        return n
}

func (n *StudentExitStatusContainerType) Set(key string, value interface{}) *StudentExitStatusContainerType {
        if n == nil {
                n = StudentExitStatusContainerTypeCreate(StudentExitStatusContainerType{})
        }
        switch key {
    case "Code":
      present := false
  for _, b := range AUCodeSetsExitWithdrawalStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsExitWithdrawalStatusType_values")
      }

                      n.Code = ((*AUCodeSetsExitWithdrawalStatusType)(StringCreate(value.(string))))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *TimeTableDayType) Set(key string, value interface{}) *TimeTableDayType {
        if n == nil {
                n = TimeTableDayTypeCreate(TimeTableDayType{})
        }
        switch key {
    case "TimeTablePeriodList":
                      n.TimeTablePeriodList = TimeTablePeriodListTypeCreate(value.(TimeTablePeriodListType))
    case "DayTitle":
                      n.DayTitle = StringCreate(value.(string))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *EquipmentInfo) Set(key string, value interface{}) *EquipmentInfo {
        if n == nil {
                n = EquipmentInfoCreate(EquipmentInfo{})
        }
        switch key {
    case "EquipmentType":
                      n.EquipmentType = StringCreate(value.(string))
    case "PurchaseOrderRefId":
                      n.PurchaseOrderRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "AssetNumber":
    
                      n.AssetNumber = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SIF_RefId":
                      n.SIF_RefId = EquipmentInfo_SIF_RefIdCreate(value.(EquipmentInfo_SIF_RefId))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingEventLocationDetailsType) Set(key string, value interface{}) *WellbeingEventLocationDetailsType {
        if n == nil {
                n = WellbeingEventLocationDetailsTypeCreate(WellbeingEventLocationDetailsType{})
        }
        switch key {
    case "Class":
                      n.Class = StringCreate(value.(string))
    case "EventLocation":
      present := false
  for _, b := range AUCodeSetsWellbeingEventLocationType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingEventLocationType_values")
      }

                      n.EventLocation = ((*AUCodeSetsWellbeingEventLocationType)(StringCreate(value.(string))))
    case "FurtherLocationNotes":
                      n.FurtherLocationNotes = StringCreate(value.(string))
        }
        return n
}

func (n *LibraryItemInfoType) Set(key string, value interface{}) *LibraryItemInfoType {
        if n == nil {
                n = LibraryItemInfoTypeCreate(LibraryItemInfoType{})
        }
        switch key {
    case "Cost":
                      n.Cost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "CallNumber":
                      n.CallNumber = StringCreate(value.(string))
    case "Author":
                      n.Author = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "ISBN":
                      n.ISBN = StringCreate(value.(string))
    case "ElectronicId":
                      n.ElectronicId = ElectronicIdTypeCreate(value.(ElectronicIdType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "ReplacementCost":
                      n.ReplacementCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
        }
        return n
}

func (n *TimeElementType_SpanGap) Set(key string, value interface{}) *TimeElementType_SpanGap {
        if n == nil {
                n = TimeElementType_SpanGapCreate(TimeElementType_SpanGap{})
        }
        switch key {
    case "EndDateTime":
                      n.EndDateTime = StringCreate(value.(string))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        }
        return n
}

func (n *OtherWellbeingResponseContainerType) Set(key string, value interface{}) *OtherWellbeingResponseContainerType {
        if n == nil {
                n = OtherWellbeingResponseContainerTypeCreate(OtherWellbeingResponseContainerType{})
        }
        switch key {
    case "OtherResponseType":
                      n.OtherResponseType = StringCreate(value.(string))
    case "OtherResponseNotes":
                      n.OtherResponseNotes = StringCreate(value.(string))
    case "OtherResponseDate":
                      n.OtherResponseDate = StringCreate(value.(string))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
    case "OtherResponseDescription":
                      n.OtherResponseDescription = StringCreate(value.(string))
        }
        return n
}

func (n *CodeFrameTestItemType) Set(key string, value interface{}) *CodeFrameTestItemType {
        if n == nil {
                n = CodeFrameTestItemTypeCreate(CodeFrameTestItemType{})
        }
        switch key {
    case "SequenceNumber":
                      n.SequenceNumber = IntCreate(value.(int))
    case "TestItemRefId":
                      n.TestItemRefId = StringCreate(value.(string))
    case "TestItemContent":
                      n.TestItemContent = NAPTestItemContentTypeCreate(value.(NAPTestItemContentType))
        }
        return n
}

func (n *LearningStandardType) Set(key string, value interface{}) *LearningStandardType {
        if n == nil {
                n = LearningStandardTypeCreate(LearningStandardType{})
        }
        switch key {
    case "LearningStandardLocalId":
    
                      n.LearningStandardLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LearningStandardURL":
                      n.LearningStandardURL = StringCreate(value.(string))
    case "LearningStandardItemRefId":
                      n.LearningStandardItemRefId = StringCreate(value.(string))
        }
        return n
}

func (n *ActivityTimeType_Duration) Set(key string, value interface{}) *ActivityTimeType_Duration {
        if n == nil {
                n = ActivityTimeType_DurationCreate(ActivityTimeType_Duration{})
        }
        switch key {
    case "Units":
                      n.Units = StringCreate(value.(string))
    case "Value":
                      n.Value = IntCreate(value.(int))
        }
        return n
}

func (n *NameOfRecordType) Set(key string, value interface{}) *NameOfRecordType {
        if n == nil {
                n = NameOfRecordTypeCreate(NameOfRecordType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "FamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PreferredFamilyNameFirst":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PreferredFamilyNameFirst = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
        }
        return n
}

func (n *SubstituteItemType) Set(key string, value interface{}) *SubstituteItemType {
        if n == nil {
                n = SubstituteItemTypeCreate(SubstituteItemType{})
        }
        switch key {
    case "PNPCodeList":
                      n.PNPCodeList = PNPCodeListTypeCreate(value.(PNPCodeListType))
    case "SubstituteItemRefId":
                      n.SubstituteItemRefId = StringCreate(value.(string))
    case "SubstituteItemLocalId":
    
                      n.SubstituteItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *TimeTableScheduleType) Set(key string, value interface{}) *TimeTableScheduleType {
        if n == nil {
                n = TimeTableScheduleTypeCreate(TimeTableScheduleType{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "TeachingPeriodsPerDay":
                      n.TeachingPeriodsPerDay = IntCreate(value.(int))
    case "TimeTableDayList":
                      n.TimeTableDayList = TimeTableDayListTypeCreate(value.(TimeTableDayListType))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "TimeTableCreationDate":
                      n.TimeTableCreationDate = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "DaysPerCycle":
                      n.DaysPerCycle = IntCreate(value.(int))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "PeriodsPerDay":
                      n.PeriodsPerDay = IntCreate(value.(int))
        }
        return n
}

func (n *ResourceUsage_ResourceReportColumn) Set(key string, value interface{}) *ResourceUsage_ResourceReportColumn {
        if n == nil {
                n = ResourceUsage_ResourceReportColumnCreate(ResourceUsage_ResourceReportColumn{})
        }
        switch key {
    case "ColumnName":
                      n.ColumnName = StringCreate(value.(string))
    case "ColumnDescription":
                      n.ColumnDescription = StringCreate(value.(string))
    case "ColumnDelimiter":
                      n.ColumnDelimiter = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingAlert) Set(key string, value interface{}) *WellbeingAlert {
        if n == nil {
                n = WellbeingAlertCreate(WellbeingAlert{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "WellbeingAlertEndDate":
                      n.WellbeingAlertEndDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "AlertKeyContact":
                      n.AlertKeyContact = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "WellbeingAlertStartDate":
                      n.WellbeingAlertStartDate = StringCreate(value.(string))
    case "WellbeingAlertDescription":
                      n.WellbeingAlertDescription = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EnrolmentRestricted":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.EnrolmentRestricted = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "WellbeingAlertCategory":
      present := false
  for _, b := range AUCodeSetsWellbeingAlertCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingAlertCategoryType_values")
      }

                      n.WellbeingAlertCategory = ((*AUCodeSetsWellbeingAlertCategoryType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "AlertSeverity":
                      n.AlertSeverity = StringCreate(value.(string))
    case "AlertAudience":
                      n.AlertAudience = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *NAPTestContentType) Set(key string, value interface{}) *NAPTestContentType {
        if n == nil {
                n = NAPTestContentTypeCreate(NAPTestContentType{})
        }
        switch key {
    case "DomainBands":
                      n.DomainBands = DomainBandsContainerTypeCreate(value.(DomainBandsContainerType))
    case "TestType":
      present := false
  for _, b := range AUCodeSetsNAPTestTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPTestTypeType_values")
      }

                      n.TestType = ((*AUCodeSetsNAPTestTypeType)(StringCreate(value.(string))))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "StagesCount":
                      n.StagesCount = IntCreate(value.(int))
    case "DomainProficiency":
                      n.DomainProficiency = DomainProficiencyContainerTypeCreate(value.(DomainProficiencyContainerType))
    case "TestYear":
    
                      n.TestYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Domain":
      present := false
  for _, b := range AUCodeSetsNAPTestDomainType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsNAPTestDomainType_values")
      }

                      n.Domain = ((*AUCodeSetsNAPTestDomainType)(StringCreate(value.(string))))
    case "TestName":
                      n.TestName = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingEvent) Set(key string, value interface{}) *WellbeingEvent {
        if n == nil {
                n = WellbeingEventCreate(WellbeingEvent{})
        }
        switch key {
    case "EventId":
    
                      n.EventId = ((*LocalIdType)(StringCreate(value.(string))))
    case "WellbeingEventCategoryList":
                      n.WellbeingEventCategoryList = WellbeingEventCategoryListTypeCreate(value.(WellbeingEventCategoryListType))
    case "WellbeingEventDescription":
                      n.WellbeingEventDescription = StringCreate(value.(string))
    case "PersonInvolvementList":
                      n.PersonInvolvementList = PersonInvolvementListTypeCreate(value.(PersonInvolvementListType))
    case "Status":
      present := false
  for _, b := range AUCodeSetsWellbeingStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingStatusType_values")
      }

                      n.Status = ((*AUCodeSetsWellbeingStatusType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "GroupIndicator":
                      n.GroupIndicator = BoolCreate(value.(bool))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "FollowUpActionList":
                      n.FollowUpActionList = FollowUpActionListTypeCreate(value.(FollowUpActionListType))
    case "WellbeingEventLocationDetails":
                      n.WellbeingEventLocationDetails = WellbeingEventLocationDetailsTypeCreate(value.(WellbeingEventLocationDetailsType))
    case "ReportingStaffRefId":
                      n.ReportingStaffRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "WellbeingEventCategoryClass":
      present := false
  for _, b := range AUCodeSetsWellbeingEventCategoryClassType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingEventCategoryClassType_values")
      }

                      n.WellbeingEventCategoryClass = ((*AUCodeSetsWellbeingEventCategoryClassType)(StringCreate(value.(string))))
    case "WellbeingEventCreationTimeStamp":
                      n.WellbeingEventCreationTimeStamp = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "WellbeingEventTimePeriod":
      present := false
  for _, b := range AUCodeSetsWellbeingEventTimePeriodType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsWellbeingEventTimePeriodType_values")
      }

                      n.WellbeingEventTimePeriod = ((*AUCodeSetsWellbeingEventTimePeriodType)(StringCreate(value.(string))))
    case "ConfidentialFlag":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ConfidentialFlag = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "WellbeingEventDate":
                      n.WellbeingEventDate = StringCreate(value.(string))
    case "WellbeingEventNotes":
                      n.WellbeingEventNotes = StringCreate(value.(string))
    case "WellbeingEventTime":
                      n.WellbeingEventTime = StringCreate(value.(string))
        }
        return n
}

func (n *WellbeingPersonLink) Set(key string, value interface{}) *WellbeingPersonLink {
        if n == nil {
                n = WellbeingPersonLinkCreate(WellbeingPersonLink{})
        }
        switch key {
    case "OtherPersonContactDetails":
                      n.OtherPersonContactDetails = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "PersonRole":
                      n.PersonRole = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "WellbeingEventRefId":
                      n.WellbeingEventRefId = StringCreate(value.(string))
    case "GroupId":
    
                      n.GroupId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FollowUpActionList":
                      n.FollowUpActionList = FollowUpActionListTypeCreate(value.(FollowUpActionListType))
    case "OtherPersonId":
    
                      n.OtherPersonId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "HowInvolved":
                      n.HowInvolved = StringCreate(value.(string))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PersonRefId":
                      n.PersonRefId = WellbeingPersonLink_PersonRefIdCreate(value.(WellbeingPersonLink_PersonRefId))
        }
        return n
}

func (n *NAPTestletResponseType) Set(key string, value interface{}) *NAPTestletResponseType {
        if n == nil {
                n = NAPTestletResponseTypeCreate(NAPTestletResponseType{})
        }
        switch key {
    case "TestletSubScore":
                      n.TestletSubScore = FloatCreate(value.(float64))
    case "NAPTestletLocalId":
    
                      n.NAPTestletLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPTestletRefId":
                      n.NAPTestletRefId = StringCreate(value.(string))
    case "ItemResponseList":
                      n.ItemResponseList = NAPTestletItemResponseListTypeCreate(value.(NAPTestletItemResponseListType))
        }
        return n
}

func (n *TimeTableContainer) Set(key string, value interface{}) *TimeTableContainer {
        if n == nil {
                n = TimeTableContainerCreate(TimeTableContainer{})
        }
        switch key {
    case "TimeTableScheduleCellList":
                      n.TimeTableScheduleCellList = TimeTableScheduleCellListTypeCreate(value.(TimeTableScheduleCellListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TimeTableSchedule":
                      n.TimeTableSchedule = TimeTableScheduleTypeCreate(value.(TimeTableScheduleType))
    case "TeachingGroupScheduleList":
                      n.TeachingGroupScheduleList = TeachingGroupScheduleTypeCreate(value.(TeachingGroupScheduleType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *DebitOrCreditAmountType) Set(key string, value interface{}) *DebitOrCreditAmountType {
        if n == nil {
                n = DebitOrCreditAmountTypeCreate(DebitOrCreditAmountType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = FloatCreate(value.(float64))
    case "Currency":
      present := false
  for _, b := range ISO4217CurrencyNamesAndCodeElementsType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "ISO4217CurrencyNamesAndCodeElementsType_values")
      }

                      n.Currency = ((*ISO4217CurrencyNamesAndCodeElementsType)(StringCreate(value.(string))))
        }
        return n
}

func (n *RoomInfo) Set(key string, value interface{}) *RoomInfo {
        if n == nil {
                n = RoomInfoCreate(RoomInfo{})
        }
        switch key {
    case "Building":
                      n.Building = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StaffList":
                      n.StaffList = StaffListTypeCreate(value.(StaffListType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RoomNumber":
                      n.RoomNumber = StringCreate(value.(string))
    case "HomeroomNumber":
                      n.HomeroomNumber = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Capacity":
                      n.Capacity = IntCreate(value.(int))
    case "RoomType":
                      n.RoomType = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
    case "AvailableForTimetable":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AvailableForTimetable = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "Size":
                      n.Size = FloatCreate(value.(float64))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *Debtor_BilledEntity) Set(key string, value interface{}) *Debtor_BilledEntity {
        if n == nil {
                n = Debtor_BilledEntityCreate(Debtor_BilledEntity{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *SystemRole_RoleScope) Set(key string, value interface{}) *SystemRole_RoleScope {
        if n == nil {
                n = SystemRole_RoleScopeCreate(SystemRole_RoleScope{})
        }
        switch key {
    case "RoleScopeName":
                      n.RoleScopeName = StringCreate(value.(string))
    case "RoleScopeRefId":
                      n.RoleScopeRefId = SystemRole_RoleScopeRefIdCreate(value.(SystemRole_RoleScopeRefId))
        }
        return n
}

func (n *PersonInfoType) Set(key string, value interface{}) *PersonInfoType {
        if n == nil {
                n = PersonInfoTypeCreate(PersonInfoType{})
        }
        switch key {
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "HouseholdContactInfoList":
                      n.HouseholdContactInfoList = HouseholdContactInfoListTypeCreate(value.(HouseholdContactInfoListType))
    case "Demographics":
                      n.Demographics = DemographicsTypeCreate(value.(DemographicsType))
    case "OtherNames":
                      n.OtherNames = OtherNamesTypeCreate(value.(OtherNamesType))
    case "Name":
                      n.Name = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
        }
        return n
}

func (n *SchoolCourseInfoOverrideType) Set(key string, value interface{}) *SchoolCourseInfoOverrideType {
        if n == nil {
                n = SchoolCourseInfoOverrideTypeCreate(SchoolCourseInfoOverrideType{})
        }
        switch key {
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
    case "CourseTitle":
                      n.CourseTitle = StringCreate(value.(string))
    case "Override":
                      n.Override = StringCreate(value.(string))
    case "StateCourseCode":
                      n.StateCourseCode = StringCreate(value.(string))
    case "DistrictCourseCode":
                      n.DistrictCourseCode = StringCreate(value.(string))
    case "CourseCode":
                      n.CourseCode = StringCreate(value.(string))
    case "CourseCredits":
                      n.CourseCredits = StringCreate(value.(string))
    case "InstructionalLevel":
                      n.InstructionalLevel = StringCreate(value.(string))
        }
        return n
}

func (n *ResourceUsage) Set(key string, value interface{}) *ResourceUsage {
        if n == nil {
                n = ResourceUsageCreate(ResourceUsage{})
        }
        switch key {
    case "ResourceReportLineList":
                      n.ResourceReportLineList = ResourceUsage_ResourceReportLineListCreate(value.(ResourceUsage_ResourceReportLineList))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ResourceUsageContentType":
                      n.ResourceUsageContentType = ResourceUsage_ResourceUsageContentTypeCreate(value.(ResourceUsage_ResourceUsageContentType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ResourceReportColumnList":
                      n.ResourceReportColumnList = ResourceUsage_ResourceReportColumnListCreate(value.(ResourceUsage_ResourceReportColumnList))
        }
        return n
}

func (n *AuditInfoType) Set(key string, value interface{}) *AuditInfoType {
        if n == nil {
                n = AuditInfoTypeCreate(AuditInfoType{})
        }
        switch key {
    case "CreationUser":
                      n.CreationUser = CreationUserTypeCreate(value.(CreationUserType))
    case "CreationDateTime":
                      n.CreationDateTime = StringCreate(value.(string))
        }
        return n
}

func (n *PersonPicture) Set(key string, value interface{}) *PersonPicture {
        if n == nil {
                n = PersonPictureCreate(PersonPicture{})
        }
        switch key {
    case "ParentObjectRefId":
                      n.ParentObjectRefId = PersonPicture_ParentObjectRefIdCreate(value.(PersonPicture_ParentObjectRefId))
    case "PictureSource":
                      n.PictureSource = PersonPicture_PictureSourceCreate(value.(PersonPicture_PictureSource))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PublishingPermissionList":
                      n.PublishingPermissionList = PublishingPermissionListTypeCreate(value.(PublishingPermissionListType))
    case "OKToPublish":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.OKToPublish = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
        }
        return n
}

func (n *RelationshipType) Set(key string, value interface{}) *RelationshipType {
        if n == nil {
                n = RelationshipTypeCreate(RelationshipType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsRelationshipToStudentType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsRelationshipToStudentType_values")
      }

                      n.Code = ((*AUCodeSetsRelationshipToStudentType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ContactFlagsType) Set(key string, value interface{}) *ContactFlagsType {
        if n == nil {
                n = ContactFlagsTypeCreate(ContactFlagsType{})
        }
        switch key {
    case "FeesAccess":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FeesAccess = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ReceivesAssessmentReport":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ReceivesAssessmentReport = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "AccessToRecords":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AccessToRecords = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "DisciplinaryContact":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.DisciplinaryContact = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "InterventionOrder":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.InterventionOrder = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PickupRights":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PickupRights = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "AttendanceContact":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.AttendanceContact = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "FeesBilling":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FeesBilling = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "ParentLegalGuardian":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.ParentLegalGuardian = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "FamilyMail":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.FamilyMail = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "EmergencyContact":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.EmergencyContact = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "LivesWith":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.LivesWith = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "HasCustody":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.HasCustody = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "PrimaryCareProvider":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PrimaryCareProvider = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
        }
        return n
}

func (n *PersonPicture_ParentObjectRefId) Set(key string, value interface{}) *PersonPicture_ParentObjectRefId {
        if n == nil {
                n = PersonPicture_ParentObjectRefIdCreate(PersonPicture_ParentObjectRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        }
        return n
}

func (n *PersonInvolvementType_PersonRefId) Set(key string, value interface{}) *PersonInvolvementType_PersonRefId {
        if n == nil {
                n = PersonInvolvementType_PersonRefIdCreate(PersonInvolvementType_PersonRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *StudentScoreJudgementAgainstStandard) Set(key string, value interface{}) *StudentScoreJudgementAgainstStandard {
        if n == nil {
                n = StudentScoreJudgementAgainstStandardCreate(StudentScoreJudgementAgainstStandard{})
        }
        switch key {
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "SpecialCircumstanceLocalCode":
    
                      n.SpecialCircumstanceLocalCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CurriculumNodeCode":
    
                      n.CurriculumNodeCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "ManagedPathwayLocalCode":
    
                      n.ManagedPathwayLocalCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalTermCode":
    
                      n.LocalTermCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentStateProvinceId":
    
                      n.StudentStateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "Score":
                      n.Score = StringCreate(value.(string))
    case "StudentLocalId":
    
                      n.StudentLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolCommonwealthId":
                      n.SchoolCommonwealthId = StringCreate(value.(string))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "ClassLocalId":
                      n.ClassLocalId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CurriculumCode":
    
                      n.CurriculumCode = ((*LocalIdType)(StringCreate(value.(string))))
        }
        return n
}

func (n *ProgramStatusType) Set(key string, value interface{}) *ProgramStatusType {
        if n == nil {
                n = ProgramStatusTypeCreate(ProgramStatusType{})
        }
        switch key {
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        }
        return n
}

func (n *SchoolCourseInfo) Set(key string, value interface{}) *SchoolCourseInfo {
        if n == nil {
                n = SchoolCourseInfoCreate(SchoolCourseInfo{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "CourseCredits":
                      n.CourseCredits = StringCreate(value.(string))
    case "Department":
                      n.Department = StringCreate(value.(string))
    case "StateCourseCode":
                      n.StateCourseCode = StringCreate(value.(string))
    case "GraduationRequirement":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.GraduationRequirement = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "SubjectAreaList":
                      n.SubjectAreaList = SubjectAreaListTypeCreate(value.(SubjectAreaListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CourseCode":
                      n.CourseCode = StringCreate(value.(string))
    case "CourseContent":
                      n.CourseContent = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "CoreAcademicCourse":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.CoreAcademicCourse = ((*AUCodeSetsYesOrNoCategoryType)(StringCreate(value.(string))))
    case "InstructionalLevel":
                      n.InstructionalLevel = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "CourseTitle":
                      n.CourseTitle = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DistrictCourseCode":
                      n.DistrictCourseCode = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
        }
        return n
}

func (n *TimeTablePeriodType) Set(key string, value interface{}) *TimeTablePeriodType {
        if n == nil {
                n = TimeTablePeriodTypeCreate(TimeTablePeriodType{})
        }
        switch key {
    case "UseInAttendanceCalculations":
                      n.UseInAttendanceCalculations = StringCreate(value.(string))
    case "RegularSchoolPeriod":
                      n.RegularSchoolPeriod = StringCreate(value.(string))
    case "EndTime":
                      n.EndTime = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PeriodTitle":
                      n.PeriodTitle = StringCreate(value.(string))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "BellPeriod":
                      n.BellPeriod = StringCreate(value.(string))
    case "InstructionalMinutes":
                      n.InstructionalMinutes = IntCreate(value.(int))
        }
        return n
}

func (n *SystemRole_RoleScopeRefId) Set(key string, value interface{}) *SystemRole_RoleScopeRefId {
        if n == nil {
                n = SystemRole_RoleScopeRefIdCreate(SystemRole_RoleScopeRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        }
        return n
}

func (n *AbstractContentElementType) Set(key string, value interface{}) *AbstractContentElementType {
        if n == nil {
                n = AbstractContentElementTypeCreate(AbstractContentElementType{})
        }
        switch key {
    case "TextData":
                      n.TextData = AbstractContentElementType_TextDataCreate(value.(AbstractContentElementType_TextData))
    case "XMLData":
                      n.XMLData = AbstractContentElementType_XMLDataCreate(value.(AbstractContentElementType_XMLData))
    case "Reference":
                      n.Reference = AbstractContentElementType_ReferenceCreate(value.(AbstractContentElementType_Reference))
    case "BinaryData":
                      n.BinaryData = AbstractContentElementType_BinaryDataCreate(value.(AbstractContentElementType_BinaryData))
        }
        return n
}

func (n *ResourceUsage_ResourceReportLine) Set(key string, value interface{}) *ResourceUsage_ResourceReportLine {
        if n == nil {
                n = ResourceUsage_ResourceReportLineCreate(ResourceUsage_ResourceReportLine{})
        }
        switch key {
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "CurrentCost":
                      n.CurrentCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "SIF_RefId":
                      n.SIF_RefId = ResourceUsage_SIF_RefIdCreate(value.(ResourceUsage_SIF_RefId))
    case "ReportRow":
                      n.ReportRow = StringCreate(value.(string))
        }
        return n
}

func (n *ScoreType) Set(key string, value interface{}) *ScoreType {
        if n == nil {
                n = ScoreTypeCreate(ScoreType{})
        }
        switch key {
    case "MaxScoreValue":
                      n.MaxScoreValue = FloatCreate(value.(float64))
    case "ScoreDescriptionList":
                      n.ScoreDescriptionList = ScoreDescriptionListTypeCreate(value.(ScoreDescriptionListType))
        }
        return n
}

func (n *NAPWritingRubricType) Set(key string, value interface{}) *NAPWritingRubricType {
        if n == nil {
                n = NAPWritingRubricTypeCreate(NAPWritingRubricType{})
        }
        switch key {
    case "RubricType":
                      n.RubricType = StringCreate(value.(string))
    case "ScoreList":
                      n.ScoreList = ScoreListTypeCreate(value.(ScoreListType))
    case "Descriptor":
                      n.Descriptor = StringCreate(value.(string))
        }
        return n
}

func (n *OtherCodeListType_OtherCode) Set(key string, value interface{}) *OtherCodeListType_OtherCode {
        if n == nil {
                n = OtherCodeListType_OtherCodeCreate(OtherCodeListType_OtherCode{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "Codeset":
                      n.Codeset = StringCreate(value.(string))
        }
        return n
}

func (n *ReligionType) Set(key string, value interface{}) *ReligionType {
        if n == nil {
                n = ReligionTypeCreate(ReligionType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
      present := false
  for _, b := range AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType_values")
      }

                      n.Code = ((*AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType)(StringCreate(value.(string))))
        }
        return n
}

