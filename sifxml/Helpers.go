package sifxml

import (
  "fmt"
  "log"
  "reflect"

      "github.com/ulule/deepcopier"
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

func IdentityAssertionsTypeCreate(x IdentityAssertionsType) *IdentityAssertionsType {
        return &x
}

func (n *IdentityAssertionsType) New() *IdentityAssertionsType {
        return IdentityAssertionsTypeCreate(IdentityAssertionsType{})
}

func (n *IdentityAssertionsType) Clone() IdentityAssertionsType {
  resource := &IdentityAssertionsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressStreetTypeCreate(x AddressStreetType) *AddressStreetType {
        return &x
}

func (n *AddressStreetType) New() *AddressStreetType {
        return AddressStreetTypeCreate(AddressStreetType{})
}

func (n *AddressStreetType) Clone() AddressStreetType {
  resource := &AddressStreetType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScheduledTeacherListTypeCreate(x ScheduledTeacherListType) *ScheduledTeacherListType {
        return &x
}

func (n *ScheduledTeacherListType) New() *ScheduledTeacherListType {
        return ScheduledTeacherListTypeCreate(ScheduledTeacherListType{})
}

func (n *ScheduledTeacherListType) Clone() ScheduledTeacherListType {
  resource := &ScheduledTeacherListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingAlertCreate(x WellbeingAlert) *WellbeingAlert {
        return &x
}

func (n *WellbeingAlert) New() *WellbeingAlert {
        return WellbeingAlertCreate(WellbeingAlert{})
}

func (n *WellbeingAlert) Clone() WellbeingAlert {
  resource := &WellbeingAlert{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LocationOfInstructionTypeCreate(x LocationOfInstructionType) *LocationOfInstructionType {
        return &x
}

func (n *LocationOfInstructionType) New() *LocationOfInstructionType {
        return LocationOfInstructionTypeCreate(LocationOfInstructionType{})
}

func (n *LocationOfInstructionType) Clone() LocationOfInstructionType {
  resource := &LocationOfInstructionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatementCodesTypeCreate(x StatementCodesType) *StatementCodesType {
        return &x
}

func (n *StatementCodesType) New() *StatementCodesType {
        return StatementCodesTypeCreate(StatementCodesType{})
}

func (n *StatementCodesType) Clone() StatementCodesType {
  resource := &StatementCodesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PrerequisitesTypeCreate(x PrerequisitesType) *PrerequisitesType {
        return &x
}

func (n *PrerequisitesType) New() *PrerequisitesType {
        return PrerequisitesTypeCreate(PrerequisitesType{})
}

func (n *PrerequisitesType) Clone() PrerequisitesType {
  resource := &PrerequisitesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentDailyAttendanceCreate(x StudentDailyAttendance) *StudentDailyAttendance {
        return &x
}

func (n *StudentDailyAttendance) New() *StudentDailyAttendance {
        return StudentDailyAttendanceCreate(StudentDailyAttendance{})
}

func (n *StudentDailyAttendance) Clone() StudentDailyAttendance {
  resource := &StudentDailyAttendance{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingEventCategoryTypeCreate(x WellbeingEventCategoryType) *WellbeingEventCategoryType {
        return &x
}

func (n *WellbeingEventCategoryType) New() *WellbeingEventCategoryType {
        return WellbeingEventCategoryTypeCreate(WellbeingEventCategoryType{})
}

func (n *WellbeingEventCategoryType) Clone() WellbeingEventCategoryType {
  resource := &WellbeingEventCategoryType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestScoreSummaryCreate(x NAPTestScoreSummary) *NAPTestScoreSummary {
        return &x
}

func (n *NAPTestScoreSummary) New() *NAPTestScoreSummary {
        return NAPTestScoreSummaryCreate(NAPTestScoreSummary{})
}

func (n *NAPTestScoreSummary) Clone() NAPTestScoreSummary {
  resource := &NAPTestScoreSummary{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPSubscoreListTypeCreate(x NAPSubscoreListType) *NAPSubscoreListType {
        return &x
}

func (n *NAPSubscoreListType) New() *NAPSubscoreListType {
        return NAPSubscoreListTypeCreate(NAPSubscoreListType{})
}

func (n *NAPSubscoreListType) Clone() NAPSubscoreListType {
  resource := &NAPSubscoreListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AttendanceTimeTypeCreate(x AttendanceTimeType) *AttendanceTimeType {
        return &x
}

func (n *AttendanceTimeType) New() *AttendanceTimeType {
        return AttendanceTimeTypeCreate(AttendanceTimeType{})
}

func (n *AttendanceTimeType) Clone() AttendanceTimeType {
  resource := &AttendanceTimeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RelationshipTypeCreate(x RelationshipType) *RelationshipType {
        return &x
}

func (n *RelationshipType) New() *RelationshipType {
        return RelationshipTypeCreate(RelationshipType{})
}

func (n *RelationshipType) Clone() RelationshipType {
  resource := &RelationshipType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ActivityTimeType_DurationCreate(x ActivityTimeType_Duration) *ActivityTimeType_Duration {
        return &x
}

func (n *ActivityTimeType_Duration) New() *ActivityTimeType_Duration {
        return ActivityTimeType_DurationCreate(ActivityTimeType_Duration{})
}

func (n *ActivityTimeType_Duration) Clone() ActivityTimeType_Duration {
  resource := &ActivityTimeType_Duration{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScoreDescriptionTypeCreate(x ScoreDescriptionType) *ScoreDescriptionType {
        return &x
}

func (n *ScoreDescriptionType) New() *ScoreDescriptionType {
        return ScoreDescriptionTypeCreate(ScoreDescriptionType{})
}

func (n *ScoreDescriptionType) Clone() ScoreDescriptionType {
  resource := &ScoreDescriptionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffActivityExtensionTypeCreate(x StaffActivityExtensionType) *StaffActivityExtensionType {
        return &x
}

func (n *StaffActivityExtensionType) New() *StaffActivityExtensionType {
        return StaffActivityExtensionTypeCreate(StaffActivityExtensionType{})
}

func (n *StaffActivityExtensionType) Clone() StaffActivityExtensionType {
  resource := &StaffActivityExtensionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryMessageListTypeCreate(x LibraryMessageListType) *LibraryMessageListType {
        return &x
}

func (n *LibraryMessageListType) New() *LibraryMessageListType {
        return LibraryMessageListTypeCreate(LibraryMessageListType{})
}

func (n *LibraryMessageListType) Clone() LibraryMessageListType {
  resource := &LibraryMessageListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentExitStatusContainerTypeCreate(x StudentExitStatusContainerType) *StudentExitStatusContainerType {
        return &x
}

func (n *StudentExitStatusContainerType) New() *StudentExitStatusContainerType {
        return StudentExitStatusContainerTypeCreate(StudentExitStatusContainerType{})
}

func (n *StudentExitStatusContainerType) Clone() StudentExitStatusContainerType {
  resource := &StudentExitStatusContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SuspensionContainerTypeCreate(x SuspensionContainerType) *SuspensionContainerType {
        return &x
}

func (n *SuspensionContainerType) New() *SuspensionContainerType {
        return SuspensionContainerTypeCreate(SuspensionContainerType{})
}

func (n *SuspensionContainerType) Clone() SuspensionContainerType {
  resource := &SuspensionContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardDocumentCreate(x LearningStandardDocument) *LearningStandardDocument {
        return &x
}

func (n *LearningStandardDocument) New() *LearningStandardDocument {
        return LearningStandardDocumentCreate(LearningStandardDocument{})
}

func (n *LearningStandardDocument) Clone() LearningStandardDocument {
  resource := &LearningStandardDocument{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ContentDescriptionListTypeCreate(x ContentDescriptionListType) *ContentDescriptionListType {
        return &x
}

func (n *ContentDescriptionListType) New() *ContentDescriptionListType {
        return ContentDescriptionListTypeCreate(ContentDescriptionListType{})
}

func (n *ContentDescriptionListType) Clone() ContentDescriptionListType {
  resource := &ContentDescriptionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SoftwareVendorInfoContainerTypeCreate(x SoftwareVendorInfoContainerType) *SoftwareVendorInfoContainerType {
        return &x
}

func (n *SoftwareVendorInfoContainerType) New() *SoftwareVendorInfoContainerType {
        return SoftwareVendorInfoContainerTypeCreate(SoftwareVendorInfoContainerType{})
}

func (n *SoftwareVendorInfoContainerType) Clone() SoftwareVendorInfoContainerType {
  resource := &SoftwareVendorInfoContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SourceObjectsTypeCreate(x SourceObjectsType) *SourceObjectsType {
        return &x
}

func (n *SourceObjectsType) New() *SourceObjectsType {
        return SourceObjectsTypeCreate(SourceObjectsType{})
}

func (n *SourceObjectsType) Clone() SourceObjectsType {
  resource := &SourceObjectsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryMessageTypeCreate(x LibraryMessageType) *LibraryMessageType {
        return &x
}

func (n *LibraryMessageType) New() *LibraryMessageType {
        return LibraryMessageTypeCreate(LibraryMessageType{})
}

func (n *LibraryMessageType) Clone() LibraryMessageType {
  resource := &LibraryMessageType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ElectronicIdListTypeCreate(x ElectronicIdListType) *ElectronicIdListType {
        return &x
}

func (n *ElectronicIdListType) New() *ElectronicIdListType {
        return ElectronicIdListTypeCreate(ElectronicIdListType{})
}

func (n *ElectronicIdListType) Clone() ElectronicIdListType {
  resource := &ElectronicIdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_SIF_RefIdCreate(x SystemRole_SIF_RefId) *SystemRole_SIF_RefId {
        return &x
}

func (n *SystemRole_SIF_RefId) New() *SystemRole_SIF_RefId {
        return SystemRole_SIF_RefIdCreate(SystemRole_SIF_RefId{})
}

func (n *SystemRole_SIF_RefId) Clone() SystemRole_SIF_RefId {
  resource := &SystemRole_SIF_RefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AlertMessageTypeCreate(x AlertMessageType) *AlertMessageType {
        return &x
}

func (n *AlertMessageType) New() *AlertMessageType {
        return AlertMessageTypeCreate(AlertMessageType{})
}

func (n *AlertMessageType) Clone() AlertMessageType {
  resource := &AlertMessageType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressCollectionStudentTypeCreate(x AddressCollectionStudentType) *AddressCollectionStudentType {
        return &x
}

func (n *AddressCollectionStudentType) New() *AddressCollectionStudentType {
        return AddressCollectionStudentTypeCreate(AddressCollectionStudentType{})
}

func (n *AddressCollectionStudentType) Clone() AddressCollectionStudentType {
  resource := &AddressCollectionStudentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_CreatedCreate(x LifeCycleType_Created) *LifeCycleType_Created {
        return &x
}

func (n *LifeCycleType_Created) New() *LifeCycleType_Created {
        return LifeCycleType_CreatedCreate(LifeCycleType_Created{})
}

func (n *LifeCycleType_Created) Clone() LifeCycleType_Created {
  resource := &LifeCycleType_Created{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MedicalAlertMessageTypeCreate(x MedicalAlertMessageType) *MedicalAlertMessageType {
        return &x
}

func (n *MedicalAlertMessageType) New() *MedicalAlertMessageType {
        return MedicalAlertMessageTypeCreate(MedicalAlertMessageType{})
}

func (n *MedicalAlertMessageType) Clone() MedicalAlertMessageType {
  resource := &MedicalAlertMessageType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ComponentTypeCreate(x ComponentType) *ComponentType {
        return &x
}

func (n *ComponentType) New() *ComponentType {
        return ComponentTypeCreate(ComponentType{})
}

func (n *ComponentType) Clone() ComponentType {
  resource := &ComponentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DomainProficiencyContainerTypeCreate(x DomainProficiencyContainerType) *DomainProficiencyContainerType {
        return &x
}

func (n *DomainProficiencyContainerType) New() *DomainProficiencyContainerType {
        return DomainProficiencyContainerTypeCreate(DomainProficiencyContainerType{})
}

func (n *DomainProficiencyContainerType) Clone() DomainProficiencyContainerType {
  resource := &DomainProficiencyContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherIdTypeCreate(x OtherIdType) *OtherIdType {
        return &x
}

func (n *OtherIdType) New() *OtherIdType {
        return OtherIdTypeCreate(OtherIdType{})
}

func (n *OtherIdType) Clone() OtherIdType {
  resource := &OtherIdType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableDayListTypeCreate(x TimeTableDayListType) *TimeTableDayListType {
        return &x
}

func (n *TimeTableDayListType) New() *TimeTableDayListType {
        return TimeTableDayListTypeCreate(TimeTableDayListType{})
}

func (n *TimeTableDayListType) Clone() TimeTableDayListType {
  resource := &TimeTableDayListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PeriodAttendanceTypeCreate(x PeriodAttendanceType) *PeriodAttendanceType {
        return &x
}

func (n *PeriodAttendanceType) New() *PeriodAttendanceType {
        return PeriodAttendanceTypeCreate(PeriodAttendanceType{})
}

func (n *PeriodAttendanceType) Clone() PeriodAttendanceType {
  resource := &PeriodAttendanceType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherNameTypeCreate(x OtherNameType) *OtherNameType {
        return &x
}

func (n *OtherNameType) New() *OtherNameType {
        return OtherNameTypeCreate(OtherNameType{})
}

func (n *OtherNameType) Clone() OtherNameType {
  resource := &OtherNameType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PNPCodeListTypeCreate(x PNPCodeListType) *PNPCodeListType {
        return &x
}

func (n *PNPCodeListType) New() *PNPCodeListType {
        return PNPCodeListTypeCreate(PNPCodeListType{})
}

func (n *PNPCodeListType) Clone() PNPCodeListType {
  resource := &PNPCodeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffAssignmentMostRecentContainerTypeCreate(x StaffAssignmentMostRecentContainerType) *StaffAssignmentMostRecentContainerType {
        return &x
}

func (n *StaffAssignmentMostRecentContainerType) New() *StaffAssignmentMostRecentContainerType {
        return StaffAssignmentMostRecentContainerTypeCreate(StaffAssignmentMostRecentContainerType{})
}

func (n *StaffAssignmentMostRecentContainerType) Clone() StaffAssignmentMostRecentContainerType {
  resource := &StaffAssignmentMostRecentContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SessionInfoCreate(x SessionInfo) *SessionInfo {
        return &x
}

func (n *SessionInfo) New() *SessionInfo {
        return SessionInfoCreate(SessionInfo{})
}

func (n *SessionInfo) Clone() SessionInfo {
  resource := &SessionInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSchoolEnrollment_CalendarCreate(x StudentSchoolEnrollment_Calendar) *StudentSchoolEnrollment_Calendar {
        return &x
}

func (n *StudentSchoolEnrollment_Calendar) New() *StudentSchoolEnrollment_Calendar {
        return StudentSchoolEnrollment_CalendarCreate(StudentSchoolEnrollment_Calendar{})
}

func (n *StudentSchoolEnrollment_Calendar) Clone() StudentSchoolEnrollment_Calendar {
  resource := &StudentSchoolEnrollment_Calendar{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherIdListTypeCreate(x OtherIdListType) *OtherIdListType {
        return &x
}

func (n *OtherIdListType) New() *OtherIdListType {
        return OtherIdListTypeCreate(OtherIdListType{})
}

func (n *OtherIdListType) Clone() OtherIdListType {
  resource := &OtherIdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusStudentTypeCreate(x CensusStudentType) *CensusStudentType {
        return &x
}

func (n *CensusStudentType) New() *CensusStudentType {
        return CensusStudentTypeCreate(CensusStudentType{})
}

func (n *CensusStudentType) Clone() CensusStudentType {
  resource := &CensusStudentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupCreate(x TeachingGroup) *TeachingGroup {
        return &x
}

func (n *TeachingGroup) New() *TeachingGroup {
        return TeachingGroupCreate(TeachingGroup{})
}

func (n *TeachingGroup) Clone() TeachingGroup {
  resource := &TeachingGroup{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonPictureCreate(x PersonPicture) *PersonPicture {
        return &x
}

func (n *PersonPicture) New() *PersonPicture {
        return PersonPictureCreate(PersonPicture{})
}

func (n *PersonPicture) Clone() PersonPicture {
  resource := &PersonPicture{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardsDocumentTypeCreate(x LearningStandardsDocumentType) *LearningStandardsDocumentType {
        return &x
}

func (n *LearningStandardsDocumentType) New() *LearningStandardsDocumentType {
        return LearningStandardsDocumentTypeCreate(LearningStandardsDocumentType{})
}

func (n *LearningStandardsDocumentType) Clone() LearningStandardsDocumentType {
  resource := &LearningStandardsDocumentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonInvolvementTypeCreate(x PersonInvolvementType) *PersonInvolvementType {
        return &x
}

func (n *PersonInvolvementType) New() *PersonInvolvementType {
        return PersonInvolvementTypeCreate(PersonInvolvementType{})
}

func (n *PersonInvolvementType) Clone() PersonInvolvementType {
  resource := &PersonInvolvementType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ProgramStatusTypeCreate(x ProgramStatusType) *ProgramStatusType {
        return &x
}

func (n *ProgramStatusType) New() *ProgramStatusType {
        return ProgramStatusTypeCreate(ProgramStatusType{})
}

func (n *ProgramStatusType) Clone() ProgramStatusType {
  resource := &ProgramStatusType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AggregateCharacteristicInfoCreate(x AggregateCharacteristicInfo) *AggregateCharacteristicInfo {
        return &x
}

func (n *AggregateCharacteristicInfo) New() *AggregateCharacteristicInfo {
        return AggregateCharacteristicInfoCreate(AggregateCharacteristicInfo{})
}

func (n *AggregateCharacteristicInfo) Clone() AggregateCharacteristicInfo {
  resource := &AggregateCharacteristicInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PasswordListType_PasswordCreate(x PasswordListType_Password) *PasswordListType_Password {
        return &x
}

func (n *PasswordListType_Password) New() *PasswordListType_Password {
        return PasswordListType_PasswordCreate(PasswordListType_Password{})
}

func (n *PasswordListType_Password) Clone() PasswordListType_Password {
  resource := &PasswordListType_Password{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ElectronicIdTypeCreate(x ElectronicIdType) *ElectronicIdType {
        return &x
}

func (n *ElectronicIdType) New() *ElectronicIdType {
        return ElectronicIdTypeCreate(ElectronicIdType{})
}

func (n *ElectronicIdType) Clone() ElectronicIdType {
  resource := &ElectronicIdType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentElementType_XMLDataCreate(x AbstractContentElementType_XMLData) *AbstractContentElementType_XMLData {
        return &x
}

func (n *AbstractContentElementType_XMLData) New() *AbstractContentElementType_XMLData {
        return AbstractContentElementType_XMLDataCreate(AbstractContentElementType_XMLData{})
}

func (n *AbstractContentElementType_XMLData) Clone() AbstractContentElementType_XMLData {
  resource := &AbstractContentElementType_XMLData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CreationUserTypeCreate(x CreationUserType) *CreationUserType {
        return &x
}

func (n *CreationUserType) New() *CreationUserType {
        return CreationUserTypeCreate(CreationUserType{})
}

func (n *CreationUserType) Clone() CreationUserType {
  resource := &CreationUserType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_ModifiedCreate(x LifeCycleType_Modified) *LifeCycleType_Modified {
        return &x
}

func (n *LifeCycleType_Modified) New() *LifeCycleType_Modified {
        return LifeCycleType_ModifiedCreate(LifeCycleType_Modified{})
}

func (n *LifeCycleType_Modified) Clone() LifeCycleType_Modified {
  resource := &LifeCycleType_Modified{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentContactPersonalCreate(x StudentContactPersonal) *StudentContactPersonal {
        return &x
}

func (n *StudentContactPersonal) New() *StudentContactPersonal {
        return StudentContactPersonalCreate(StudentContactPersonal{})
}

func (n *StudentContactPersonal) Clone() StudentContactPersonal {
  resource := &StudentContactPersonal{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PurchaseOrderCreate(x PurchaseOrder) *PurchaseOrder {
        return &x
}

func (n *PurchaseOrder) New() *PurchaseOrder {
        return PurchaseOrderCreate(PurchaseOrder{})
}

func (n *PurchaseOrder) Clone() PurchaseOrder {
  resource := &PurchaseOrder{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableDayTypeCreate(x TimeTableDayType) *TimeTableDayType {
        return &x
}

func (n *TimeTableDayType) New() *TimeTableDayType {
        return TimeTableDayTypeCreate(TimeTableDayType{})
}

func (n *TimeTableDayType) Clone() TimeTableDayType {
  resource := &TimeTableDayType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CalendarDateCreate(x CalendarDate) *CalendarDate {
        return &x
}

func (n *CalendarDate) New() *CalendarDate {
        return CalendarDateCreate(CalendarDate{})
}

func (n *CalendarDate) Clone() CalendarDate {
  resource := &CalendarDate{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func IdentityCreate(x Identity) *Identity {
        return &x
}

func (n *Identity) New() *Identity {
        return IdentityCreate(Identity{})
}

func (n *Identity) Clone() Identity {
  resource := &Identity{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentExitContainerTypeCreate(x StudentExitContainerType) *StudentExitContainerType {
        return &x
}

func (n *StudentExitContainerType) New() *StudentExitContainerType {
        return StudentExitContainerTypeCreate(StudentExitContainerType{})
}

func (n *StudentExitContainerType) Clone() StudentExitContainerType {
  resource := &StudentExitContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolInfo_OtherLEACreate(x SchoolInfo_OtherLEA) *SchoolInfo_OtherLEA {
        return &x
}

func (n *SchoolInfo_OtherLEA) New() *SchoolInfo_OtherLEA {
        return SchoolInfo_OtherLEACreate(SchoolInfo_OtherLEA{})
}

func (n *SchoolInfo_OtherLEA) Clone() SchoolInfo_OtherLEA {
  resource := &SchoolInfo_OtherLEA{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func Activity_EvaluationCreate(x Activity_Evaluation) *Activity_Evaluation {
        return &x
}

func (n *Activity_Evaluation) New() *Activity_Evaluation {
        return Activity_EvaluationCreate(Activity_Evaluation{})
}

func (n *Activity_Evaluation) Clone() Activity_Evaluation {
  resource := &Activity_Evaluation{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CollectionRoundCreate(x CollectionRound) *CollectionRound {
        return &x
}

func (n *CollectionRound) New() *CollectionRound {
        return CollectionRoundCreate(CollectionRound{})
}

func (n *CollectionRound) Clone() CollectionRound {
  resource := &CollectionRound{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSchoolEnrollmentCreate(x StudentSchoolEnrollment) *StudentSchoolEnrollment {
        return &x
}

func (n *StudentSchoolEnrollment) New() *StudentSchoolEnrollment {
        return StudentSchoolEnrollmentCreate(StudentSchoolEnrollment{})
}

func (n *StudentSchoolEnrollment) Clone() StudentSchoolEnrollment {
  resource := &StudentSchoolEnrollment{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentElementTypeCreate(x AbstractContentElementType) *AbstractContentElementType {
        return &x
}

func (n *AbstractContentElementType) New() *AbstractContentElementType {
        return AbstractContentElementTypeCreate(AbstractContentElementType{})
}

func (n *AbstractContentElementType) Clone() AbstractContentElementType {
  resource := &AbstractContentElementType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ValidLetterMarkTypeCreate(x ValidLetterMarkType) *ValidLetterMarkType {
        return &x
}

func (n *ValidLetterMarkType) New() *ValidLetterMarkType {
        return ValidLetterMarkTypeCreate(ValidLetterMarkType{})
}

func (n *ValidLetterMarkType) Clone() ValidLetterMarkType {
  resource := &ValidLetterMarkType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TotalEnrollmentsTypeCreate(x TotalEnrollmentsType) *TotalEnrollmentsType {
        return &x
}

func (n *TotalEnrollmentsType) New() *TotalEnrollmentsType {
        return TotalEnrollmentsTypeCreate(TotalEnrollmentsType{})
}

func (n *TotalEnrollmentsType) Clone() TotalEnrollmentsType {
  resource := &TotalEnrollmentsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupPeriodListTypeCreate(x TeachingGroupPeriodListType) *TeachingGroupPeriodListType {
        return &x
}

func (n *TeachingGroupPeriodListType) New() *TeachingGroupPeriodListType {
        return TeachingGroupPeriodListTypeCreate(TeachingGroupPeriodListType{})
}

func (n *TeachingGroupPeriodListType) Clone() TeachingGroupPeriodListType {
  resource := &TeachingGroupPeriodListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GradingAssignmentScoreCreate(x GradingAssignmentScore) *GradingAssignmentScore {
        return &x
}

func (n *GradingAssignmentScore) New() *GradingAssignmentScore {
        return GradingAssignmentScoreCreate(GradingAssignmentScore{})
}

func (n *GradingAssignmentScore) Clone() GradingAssignmentScore {
  resource := &GradingAssignmentScore{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ActivityCreate(x Activity) *Activity {
        return &x
}

func (n *Activity) New() *Activity {
        return ActivityCreate(Activity{})
}

func (n *Activity) Clone() Activity {
  resource := &Activity{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ValidLetterMarkListTypeCreate(x ValidLetterMarkListType) *ValidLetterMarkListType {
        return &x
}

func (n *ValidLetterMarkListType) New() *ValidLetterMarkListType {
        return ValidLetterMarkListTypeCreate(ValidLetterMarkListType{})
}

func (n *ValidLetterMarkListType) Clone() ValidLetterMarkListType {
  resource := &ValidLetterMarkListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentPeriodAttendanceCreate(x StudentPeriodAttendance) *StudentPeriodAttendance {
        return &x
}

func (n *StudentPeriodAttendance) New() *StudentPeriodAttendance {
        return StudentPeriodAttendanceCreate(StudentPeriodAttendance{})
}

func (n *StudentPeriodAttendance) Clone() StudentPeriodAttendance {
  resource := &StudentPeriodAttendance{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPStudentResponseSetCreate(x NAPStudentResponseSet) *NAPStudentResponseSet {
        return &x
}

func (n *NAPStudentResponseSet) New() *NAPStudentResponseSet {
        return NAPStudentResponseSetCreate(NAPStudentResponseSet{})
}

func (n *NAPStudentResponseSet) Clone() NAPStudentResponseSet {
  resource := &NAPStudentResponseSet{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LanguageListTypeCreate(x LanguageListType) *LanguageListType {
        return &x
}

func (n *LanguageListType) New() *LanguageListType {
        return LanguageListTypeCreate(LanguageListType{})
}

func (n *LanguageListType) Clone() LanguageListType {
  resource := &LanguageListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPStudentResponseTestletListTypeCreate(x NAPStudentResponseTestletListType) *NAPStudentResponseTestletListType {
        return &x
}

func (n *NAPStudentResponseTestletListType) New() *NAPStudentResponseTestletListType {
        return NAPStudentResponseTestletListTypeCreate(NAPStudentResponseTestletListType{})
}

func (n *NAPStudentResponseTestletListType) Clone() NAPStudentResponseTestletListType {
  resource := &NAPStudentResponseTestletListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentParticipationCreate(x StudentParticipation) *StudentParticipation {
        return &x
}

func (n *StudentParticipation) New() *StudentParticipation {
        return StudentParticipationCreate(StudentParticipation{})
}

func (n *StudentParticipation) Clone() StudentParticipation {
  resource := &StudentParticipation{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSchoolEnrollment_CounselorCreate(x StudentSchoolEnrollment_Counselor) *StudentSchoolEnrollment_Counselor {
        return &x
}

func (n *StudentSchoolEnrollment_Counselor) New() *StudentSchoolEnrollment_Counselor {
        return StudentSchoolEnrollment_CounselorCreate(StudentSchoolEnrollment_Counselor{})
}

func (n *StudentSchoolEnrollment_Counselor) Clone() StudentSchoolEnrollment_Counselor {
  resource := &StudentSchoolEnrollment_Counselor{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableContainerCreate(x TimeTableContainer) *TimeTableContainer {
        return &x
}

func (n *TimeTableContainer) New() *TimeTableContainer {
        return TimeTableContainerCreate(TimeTableContainer{})
}

func (n *TimeTableContainer) Clone() TimeTableContainer {
  resource := &TimeTableContainer{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FinancialAccountCreate(x FinancialAccount) *FinancialAccount {
        return &x
}

func (n *FinancialAccount) New() *FinancialAccount {
        return FinancialAccountCreate(FinancialAccount{})
}

func (n *FinancialAccount) Clone() FinancialAccount {
  resource := &FinancialAccount{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ContactsTypeCreate(x ContactsType) *ContactsType {
        return &x
}

func (n *ContactsType) New() *ContactsType {
        return ContactsTypeCreate(ContactsType{})
}

func (n *ContactsType) Clone() ContactsType {
  resource := &ContactsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingAppealCreate(x WellbeingAppeal) *WellbeingAppeal {
        return &x
}

func (n *WellbeingAppeal) New() *WellbeingAppeal {
        return WellbeingAppealCreate(WellbeingAppeal{})
}

func (n *WellbeingAppeal) Clone() WellbeingAppeal {
  resource := &WellbeingAppeal{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentElementType_BinaryDataCreate(x AbstractContentElementType_BinaryData) *AbstractContentElementType_BinaryData {
        return &x
}

func (n *AbstractContentElementType_BinaryData) New() *AbstractContentElementType_BinaryData {
        return AbstractContentElementType_BinaryDataCreate(AbstractContentElementType_BinaryData{})
}

func (n *AbstractContentElementType_BinaryData) Clone() AbstractContentElementType_BinaryData {
  resource := &AbstractContentElementType_BinaryData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentAttendanceCollectionReportingListTypeCreate(x StudentAttendanceCollectionReportingListType) *StudentAttendanceCollectionReportingListType {
        return &x
}

func (n *StudentAttendanceCollectionReportingListType) New() *StudentAttendanceCollectionReportingListType {
        return StudentAttendanceCollectionReportingListTypeCreate(StudentAttendanceCollectionReportingListType{})
}

func (n *StudentAttendanceCollectionReportingListType) Clone() StudentAttendanceCollectionReportingListType {
  resource := &StudentAttendanceCollectionReportingListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolProgramTypeCreate(x SchoolProgramType) *SchoolProgramType {
        return &x
}

func (n *SchoolProgramType) New() *SchoolProgramType {
        return SchoolProgramTypeCreate(SchoolProgramType{})
}

func (n *SchoolProgramType) Clone() SchoolProgramType {
  resource := &SchoolProgramType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleTypeCreate(x LifeCycleType) *LifeCycleType {
        return &x
}

func (n *LifeCycleType) New() *LifeCycleType {
        return LifeCycleTypeCreate(LifeCycleType{})
}

func (n *LifeCycleType) Clone() LifeCycleType {
  resource := &LifeCycleType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ContactFlagsTypeCreate(x ContactFlagsType) *ContactFlagsType {
        return &x
}

func (n *ContactFlagsType) New() *ContactFlagsType {
        return ContactFlagsTypeCreate(ContactFlagsType{})
}

func (n *ContactFlagsType) Clone() ContactFlagsType {
  resource := &ContactFlagsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupScheduleTypeCreate(x TeachingGroupScheduleType) *TeachingGroupScheduleType {
        return &x
}

func (n *TeachingGroupScheduleType) New() *TeachingGroupScheduleType {
        return TeachingGroupScheduleTypeCreate(TeachingGroupScheduleType{})
}

func (n *TeachingGroupScheduleType) Clone() TeachingGroupScheduleType {
  resource := &TeachingGroupScheduleType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func BaseNameTypeCreate(x BaseNameType) *BaseNameType {
        return &x
}

func (n *BaseNameType) New() *BaseNameType {
        return BaseNameTypeCreate(BaseNameType{})
}

func (n *BaseNameType) Clone() BaseNameType {
  resource := &BaseNameType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AuditInfoTypeCreate(x AuditInfoType) *AuditInfoType {
        return &x
}

func (n *AuditInfoType) New() *AuditInfoType {
        return AuditInfoTypeCreate(AuditInfoType{})
}

func (n *AuditInfoType) Clone() AuditInfoType {
  resource := &AuditInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentAttendanceSummaryCreate(x StudentAttendanceSummary) *StudentAttendanceSummary {
        return &x
}

func (n *StudentAttendanceSummary) New() *StudentAttendanceSummary {
        return StudentAttendanceSummaryCreate(StudentAttendanceSummary{})
}

func (n *StudentAttendanceSummary) Clone() StudentAttendanceSummary {
  resource := &StudentAttendanceSummary{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletResponseItemTypeCreate(x NAPTestletResponseItemType) *NAPTestletResponseItemType {
        return &x
}

func (n *NAPTestletResponseItemType) New() *NAPTestletResponseItemType {
        return NAPTestletResponseItemTypeCreate(NAPTestletResponseItemType{})
}

func (n *NAPTestletResponseItemType) Clone() NAPTestletResponseItemType {
  resource := &NAPTestletResponseItemType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ExpenseAccountsTypeCreate(x ExpenseAccountsType) *ExpenseAccountsType {
        return &x
}

func (n *ExpenseAccountsType) New() *ExpenseAccountsType {
        return ExpenseAccountsTypeCreate(ExpenseAccountsType{})
}

func (n *ExpenseAccountsType) Clone() ExpenseAccountsType {
  resource := &ExpenseAccountsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EvaluationsTypeCreate(x EvaluationsType) *EvaluationsType {
        return &x
}

func (n *EvaluationsType) New() *EvaluationsType {
        return EvaluationsTypeCreate(EvaluationsType{})
}

func (n *EvaluationsType) Clone() EvaluationsType {
  resource := &EvaluationsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentScoreJudgementAgainstStandardCreate(x StudentScoreJudgementAgainstStandard) *StudentScoreJudgementAgainstStandard {
        return &x
}

func (n *StudentScoreJudgementAgainstStandard) New() *StudentScoreJudgementAgainstStandard {
        return StudentScoreJudgementAgainstStandardCreate(StudentScoreJudgementAgainstStandard{})
}

func (n *StudentScoreJudgementAgainstStandard) Clone() StudentScoreJudgementAgainstStandard {
  resource := &StudentScoreJudgementAgainstStandard{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NameTypeCreate(x NameType) *NameType {
        return &x
}

func (n *NameType) New() *NameType {
        return NameTypeCreate(NameType{})
}

func (n *NameType) Clone() NameType {
  resource := &NameType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EnglishProficiencyTypeCreate(x EnglishProficiencyType) *EnglishProficiencyType {
        return &x
}

func (n *EnglishProficiencyType) New() *EnglishProficiencyType {
        return EnglishProficiencyTypeCreate(EnglishProficiencyType{})
}

func (n *EnglishProficiencyType) Clone() EnglishProficiencyType {
  resource := &EnglishProficiencyType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CountryList2TypeCreate(x CountryList2Type) *CountryList2Type {
        return &x
}

func (n *CountryList2Type) New() *CountryList2Type {
        return CountryList2TypeCreate(CountryList2Type{})
}

func (n *CountryList2Type) Clone() CountryList2Type {
  resource := &CountryList2Type{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LEAContactListTypeCreate(x LEAContactListType) *LEAContactListType {
        return &x
}

func (n *LEAContactListType) New() *LEAContactListType {
        return LEAContactListTypeCreate(LEAContactListType{})
}

func (n *LEAContactListType) Clone() LEAContactListType {
  resource := &LEAContactListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DebitOrCreditAmountTypeCreate(x DebitOrCreditAmountType) *DebitOrCreditAmountType {
        return &x
}

func (n *DebitOrCreditAmountType) New() *DebitOrCreditAmountType {
        return DebitOrCreditAmountTypeCreate(DebitOrCreditAmountType{})
}

func (n *DebitOrCreditAmountType) Clone() DebitOrCreditAmountType {
  resource := &DebitOrCreditAmountType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CalendarDateInfoTypeCreate(x CalendarDateInfoType) *CalendarDateInfoType {
        return &x
}

func (n *CalendarDateInfoType) New() *CalendarDateInfoType {
        return CalendarDateInfoTypeCreate(CalendarDateInfoType{})
}

func (n *CalendarDateInfoType) Clone() CalendarDateInfoType {
  resource := &CalendarDateInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AdjustmentContainerTypeCreate(x AdjustmentContainerType) *AdjustmentContainerType {
        return &x
}

func (n *AdjustmentContainerType) New() *AdjustmentContainerType {
        return AdjustmentContainerTypeCreate(AdjustmentContainerType{})
}

func (n *AdjustmentContainerType) Clone() AdjustmentContainerType {
  resource := &AdjustmentContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSchoolEnrollment_AdvisorCreate(x StudentSchoolEnrollment_Advisor) *StudentSchoolEnrollment_Advisor {
        return &x
}

func (n *StudentSchoolEnrollment_Advisor) New() *StudentSchoolEnrollment_Advisor {
        return StudentSchoolEnrollment_AdvisorCreate(StudentSchoolEnrollment_Advisor{})
}

func (n *StudentSchoolEnrollment_Advisor) Clone() StudentSchoolEnrollment_Advisor {
  resource := &StudentSchoolEnrollment_Advisor{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSectionEnrollmentCreate(x StudentSectionEnrollment) *StudentSectionEnrollment {
        return &x
}

func (n *StudentSectionEnrollment) New() *StudentSectionEnrollment {
        return StudentSectionEnrollmentCreate(StudentSectionEnrollment{})
}

func (n *StudentSectionEnrollment) Clone() StudentSectionEnrollment {
  resource := &StudentSectionEnrollment{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeElementType_SpanGapsCreate(x TimeElementType_SpanGaps) *TimeElementType_SpanGaps {
        return &x
}

func (n *TimeElementType_SpanGaps) New() *TimeElementType_SpanGaps {
        return TimeElementType_SpanGapsCreate(TimeElementType_SpanGaps{})
}

func (n *TimeElementType_SpanGaps) Clone() TimeElementType_SpanGaps {
  resource := &TimeElementType_SpanGaps{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StandardHierarchyLevelTypeCreate(x StandardHierarchyLevelType) *StandardHierarchyLevelType {
        return &x
}

func (n *StandardHierarchyLevelType) New() *StandardHierarchyLevelType {
        return StandardHierarchyLevelTypeCreate(StandardHierarchyLevelType{})
}

func (n *StandardHierarchyLevelType) Clone() StandardHierarchyLevelType {
  resource := &StandardHierarchyLevelType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_CreatorCreate(x LifeCycleType_Creator) *LifeCycleType_Creator {
        return &x
}

func (n *LifeCycleType_Creator) New() *LifeCycleType_Creator {
        return LifeCycleType_CreatorCreate(LifeCycleType_Creator{})
}

func (n *LifeCycleType_Creator) Clone() LifeCycleType_Creator {
  resource := &LifeCycleType_Creator{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SubjectAreaTypeCreate(x SubjectAreaType) *SubjectAreaType {
        return &x
}

func (n *SubjectAreaType) New() *SubjectAreaType {
        return SubjectAreaTypeCreate(SubjectAreaType{})
}

func (n *SubjectAreaType) Clone() SubjectAreaType {
  resource := &SubjectAreaType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EmailTypeCreate(x EmailType) *EmailType {
        return &x
}

func (n *EmailType) New() *EmailType {
        return EmailTypeCreate(EmailType{})
}

func (n *EmailType) Clone() EmailType {
  resource := &EmailType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupTeacherTypeCreate(x TeachingGroupTeacherType) *TeachingGroupTeacherType {
        return &x
}

func (n *TeachingGroupTeacherType) New() *TeachingGroupTeacherType {
        return TeachingGroupTeacherTypeCreate(TeachingGroupTeacherType{})
}

func (n *TeachingGroupTeacherType) Clone() TeachingGroupTeacherType {
  resource := &TeachingGroupTeacherType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func JournalAdjustmentTypeCreate(x JournalAdjustmentType) *JournalAdjustmentType {
        return &x
}

func (n *JournalAdjustmentType) New() *JournalAdjustmentType {
        return JournalAdjustmentTypeCreate(JournalAdjustmentType{})
}

func (n *JournalAdjustmentType) Clone() JournalAdjustmentType {
  resource := &JournalAdjustmentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffSubjectTypeCreate(x StaffSubjectType) *StaffSubjectType {
        return &x
}

func (n *StaffSubjectType) New() *StaffSubjectType {
        return StaffSubjectTypeCreate(StaffSubjectType{})
}

func (n *StaffSubjectType) Clone() StaffSubjectType {
  resource := &StaffSubjectType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SIF_MetadataType_TimeElementsCreate(x SIF_MetadataType_TimeElements) *SIF_MetadataType_TimeElements {
        return &x
}

func (n *SIF_MetadataType_TimeElements) New() *SIF_MetadataType_TimeElements {
        return SIF_MetadataType_TimeElementsCreate(SIF_MetadataType_TimeElements{})
}

func (n *SIF_MetadataType_TimeElements) Clone() SIF_MetadataType_TimeElements {
  resource := &SIF_MetadataType_TimeElements{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffMostRecentContainerTypeCreate(x StaffMostRecentContainerType) *StaffMostRecentContainerType {
        return &x
}

func (n *StaffMostRecentContainerType) New() *StaffMostRecentContainerType {
        return StaffMostRecentContainerTypeCreate(StaffMostRecentContainerType{})
}

func (n *StaffMostRecentContainerType) Clone() StaffMostRecentContainerType {
  resource := &StaffMostRecentContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RelatedLearningStandardItemRefIdListTypeCreate(x RelatedLearningStandardItemRefIdListType) *RelatedLearningStandardItemRefIdListType {
        return &x
}

func (n *RelatedLearningStandardItemRefIdListType) New() *RelatedLearningStandardItemRefIdListType {
        return RelatedLearningStandardItemRefIdListTypeCreate(RelatedLearningStandardItemRefIdListType{})
}

func (n *RelatedLearningStandardItemRefIdListType) Clone() RelatedLearningStandardItemRefIdListType {
  resource := &RelatedLearningStandardItemRefIdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LanguageBaseTypeCreate(x LanguageBaseType) *LanguageBaseType {
        return &x
}

func (n *LanguageBaseType) New() *LanguageBaseType {
        return LanguageBaseTypeCreate(LanguageBaseType{})
}

func (n *LanguageBaseType) Clone() LanguageBaseType {
  resource := &LanguageBaseType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScheduledActivityCreate(x ScheduledActivity) *ScheduledActivity {
        return &x
}

func (n *ScheduledActivity) New() *ScheduledActivity {
        return ScheduledActivityCreate(ScheduledActivity{})
}

func (n *ScheduledActivity) Clone() ScheduledActivity {
  resource := &ScheduledActivity{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AssignmentScoreTypeCreate(x AssignmentScoreType) *AssignmentScoreType {
        return &x
}

func (n *AssignmentScoreType) New() *AssignmentScoreType {
        return AssignmentScoreTypeCreate(AssignmentScoreType{})
}

func (n *AssignmentScoreType) Clone() AssignmentScoreType {
  resource := &AssignmentScoreType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func HouseholdContactInfoListTypeCreate(x HouseholdContactInfoListType) *HouseholdContactInfoListType {
        return &x
}

func (n *HouseholdContactInfoListType) New() *HouseholdContactInfoListType {
        return HouseholdContactInfoListTypeCreate(HouseholdContactInfoListType{})
}

func (n *HouseholdContactInfoListType) Clone() HouseholdContactInfoListType {
  resource := &HouseholdContactInfoListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScoreTypeCreate(x ScoreType) *ScoreType {
        return &x
}

func (n *ScoreType) New() *ScoreType {
        return ScoreTypeCreate(ScoreType{})
}

func (n *ScoreType) Clone() ScoreType {
  resource := &ScoreType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ReligiousEventListTypeCreate(x ReligiousEventListType) *ReligiousEventListType {
        return &x
}

func (n *ReligiousEventListType) New() *ReligiousEventListType {
        return ReligiousEventListTypeCreate(ReligiousEventListType{})
}

func (n *ReligiousEventListType) Clone() ReligiousEventListType {
  resource := &ReligiousEventListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EvaluationTypeCreate(x EvaluationType) *EvaluationType {
        return &x
}

func (n *EvaluationType) New() *EvaluationType {
        return EvaluationTypeCreate(EvaluationType{})
}

func (n *EvaluationType) Clone() EvaluationType {
  resource := &EvaluationType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPWritingRubricTypeCreate(x NAPWritingRubricType) *NAPWritingRubricType {
        return &x
}

func (n *NAPWritingRubricType) New() *NAPWritingRubricType {
        return NAPWritingRubricTypeCreate(NAPWritingRubricType{})
}

func (n *NAPWritingRubricType) Clone() NAPWritingRubricType {
  resource := &NAPWritingRubricType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourcesTypeCreate(x ResourcesType) *ResourcesType {
        return &x
}

func (n *ResourcesType) New() *ResourcesType {
        return ResourcesTypeCreate(ResourcesType{})
}

func (n *ResourcesType) Clone() ResourcesType {
  resource := &ResourcesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func HoldInfoListTypeCreate(x HoldInfoListType) *HoldInfoListType {
        return &x
}

func (n *HoldInfoListType) New() *HoldInfoListType {
        return HoldInfoListTypeCreate(HoldInfoListType{})
}

func (n *HoldInfoListType) Clone() HoldInfoListType {
  resource := &HoldInfoListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PaymentReceiptLineListTypeCreate(x PaymentReceiptLineListType) *PaymentReceiptLineListType {
        return &x
}

func (n *PaymentReceiptLineListType) New() *PaymentReceiptLineListType {
        return PaymentReceiptLineListTypeCreate(PaymentReceiptLineListType{})
}

func (n *PaymentReceiptLineListType) Clone() PaymentReceiptLineListType {
  resource := &PaymentReceiptLineListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatsCohortListTypeCreate(x StatsCohortListType) *StatsCohortListType {
        return &x
}

func (n *StatsCohortListType) New() *StatsCohortListType {
        return StatsCohortListTypeCreate(StatsCohortListType{})
}

func (n *StatsCohortListType) Clone() StatsCohortListType {
  resource := &StatsCohortListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_ModificationHistoryCreate(x LifeCycleType_ModificationHistory) *LifeCycleType_ModificationHistory {
        return &x
}

func (n *LifeCycleType_ModificationHistory) New() *LifeCycleType_ModificationHistory {
        return LifeCycleType_ModificationHistoryCreate(LifeCycleType_ModificationHistory{})
}

func (n *LifeCycleType_ModificationHistory) Clone() LifeCycleType_ModificationHistory {
  resource := &LifeCycleType_ModificationHistory{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MediumOfInstructionTypeCreate(x MediumOfInstructionType) *MediumOfInstructionType {
        return &x
}

func (n *MediumOfInstructionType) New() *MediumOfInstructionType {
        return MediumOfInstructionTypeCreate(MediumOfInstructionType{})
}

func (n *MediumOfInstructionType) Clone() MediumOfInstructionType {
  resource := &MediumOfInstructionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func HoldInfoTypeCreate(x HoldInfoType) *HoldInfoType {
        return &x
}

func (n *HoldInfoType) New() *HoldInfoType {
        return HoldInfoTypeCreate(HoldInfoType{})
}

func (n *HoldInfoType) Clone() HoldInfoType {
  resource := &HoldInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatisticalAreasTypeCreate(x StatisticalAreasType) *StatisticalAreasType {
        return &x
}

func (n *StatisticalAreasType) New() *StatisticalAreasType {
        return StatisticalAreasTypeCreate(StatisticalAreasType{})
}

func (n *StatisticalAreasType) Clone() StatisticalAreasType {
  resource := &StatisticalAreasType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressCollectionCreate(x AddressCollection) *AddressCollection {
        return &x
}

func (n *AddressCollection) New() *AddressCollection {
        return AddressCollectionCreate(AddressCollection{})
}

func (n *AddressCollection) Clone() AddressCollection {
  resource := &AddressCollection{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusReportingTypeCreate(x CensusReportingType) *CensusReportingType {
        return &x
}

func (n *CensusReportingType) New() *CensusReportingType {
        return CensusReportingTypeCreate(CensusReportingType{})
}

func (n *CensusReportingType) Clone() CensusReportingType {
  resource := &CensusReportingType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupPeriodTypeCreate(x TeachingGroupPeriodType) *TeachingGroupPeriodType {
        return &x
}

func (n *TeachingGroupPeriodType) New() *TeachingGroupPeriodType {
        return TeachingGroupPeriodTypeCreate(TeachingGroupPeriodType{})
}

func (n *TeachingGroupPeriodType) Clone() TeachingGroupPeriodType {
  resource := &TeachingGroupPeriodType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AggregateStatisticFactCreate(x AggregateStatisticFact) *AggregateStatisticFact {
        return &x
}

func (n *AggregateStatisticFact) New() *AggregateStatisticFact {
        return AggregateStatisticFactCreate(AggregateStatisticFact{})
}

func (n *AggregateStatisticFact) Clone() AggregateStatisticFact {
  resource := &AggregateStatisticFact{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningResourceCreate(x LearningResource) *LearningResource {
        return &x
}

func (n *LearningResource) New() *LearningResource {
        return LearningResourceCreate(LearningResource{})
}

func (n *LearningResource) Clone() LearningResource {
  resource := &LearningResource{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PhoneNumberListTypeCreate(x PhoneNumberListType) *PhoneNumberListType {
        return &x
}

func (n *PhoneNumberListType) New() *PhoneNumberListType {
        return PhoneNumberListTypeCreate(PhoneNumberListType{})
}

func (n *PhoneNumberListType) Clone() PhoneNumberListType {
  resource := &PhoneNumberListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CountryListTypeCreate(x CountryListType) *CountryListType {
        return &x
}

func (n *CountryListType) New() *CountryListType {
        return CountryListTypeCreate(CountryListType{})
}

func (n *CountryListType) Clone() CountryListType {
  resource := &CountryListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentGroupTypeCreate(x StudentGroupType) *StudentGroupType {
        return &x
}

func (n *StudentGroupType) New() *StudentGroupType {
        return StudentGroupTypeCreate(StudentGroupType{})
}

func (n *StudentGroupType) Clone() StudentGroupType {
  resource := &StudentGroupType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StimulusLocalIdListTypeCreate(x StimulusLocalIdListType) *StimulusLocalIdListType {
        return &x
}

func (n *StimulusLocalIdListType) New() *StimulusLocalIdListType {
        return StimulusLocalIdListTypeCreate(StimulusLocalIdListType{})
}

func (n *StimulusLocalIdListType) Clone() StimulusLocalIdListType {
  resource := &StimulusLocalIdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ReligiousEventTypeCreate(x ReligiousEventType) *ReligiousEventType {
        return &x
}

func (n *ReligiousEventType) New() *ReligiousEventType {
        return ReligiousEventTypeCreate(ReligiousEventType{})
}

func (n *ReligiousEventType) Clone() ReligiousEventType {
  resource := &ReligiousEventType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LEAInfoCreate(x LEAInfo) *LEAInfo {
        return &x
}

func (n *LEAInfo) New() *LEAInfo {
        return LEAInfoCreate(LEAInfo{})
}

func (n *LEAInfo) Clone() LEAInfo {
  resource := &LEAInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PublishingPermissionListTypeCreate(x PublishingPermissionListType) *PublishingPermissionListType {
        return &x
}

func (n *PublishingPermissionListType) New() *PublishingPermissionListType {
        return PublishingPermissionListTypeCreate(PublishingPermissionListType{})
}

func (n *PublishingPermissionListType) Clone() PublishingPermissionListType {
  resource := &PublishingPermissionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MedicalAlertMessagesTypeCreate(x MedicalAlertMessagesType) *MedicalAlertMessagesType {
        return &x
}

func (n *MedicalAlertMessagesType) New() *MedicalAlertMessagesType {
        return MedicalAlertMessagesTypeCreate(MedicalAlertMessagesType{})
}

func (n *MedicalAlertMessagesType) Clone() MedicalAlertMessagesType {
  resource := &MedicalAlertMessagesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AgencyTypeCreate(x AgencyType) *AgencyType {
        return &x
}

func (n *AgencyType) New() *AgencyType {
        return AgencyTypeCreate(AgencyType{})
}

func (n *AgencyType) Clone() AgencyType {
  resource := &AgencyType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestItem2TypeCreate(x NAPTestItem2Type) *NAPTestItem2Type {
        return &x
}

func (n *NAPTestItem2Type) New() *NAPTestItem2Type {
        return NAPTestItem2TypeCreate(NAPTestItem2Type{})
}

func (n *NAPTestItem2Type) Clone() NAPTestItem2Type {
  resource := &NAPTestItem2Type{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SIF_ExtendedElementsTypeCreate(x SIF_ExtendedElementsType) *SIF_ExtendedElementsType {
        return &x
}

func (n *SIF_ExtendedElementsType) New() *SIF_ExtendedElementsType {
        return SIF_ExtendedElementsTypeCreate(SIF_ExtendedElementsType{})
}

func (n *SIF_ExtendedElementsType) Clone() SIF_ExtendedElementsType {
  resource := &SIF_ExtendedElementsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQItemTypeCreate(x FQItemType) *FQItemType {
        return &x
}

func (n *FQItemType) New() *FQItemType {
        return FQItemTypeCreate(FQItemType{})
}

func (n *FQItemType) Clone() FQItemType {
  resource := &FQItemType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusReportingListTypeCreate(x CensusReportingListType) *CensusReportingListType {
        return &x
}

func (n *CensusReportingListType) New() *CensusReportingListType {
        return CensusReportingListTypeCreate(CensusReportingListType{})
}

func (n *CensusReportingListType) Clone() CensusReportingListType {
  resource := &CensusReportingListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQReportingListTypeCreate(x FQReportingListType) *FQReportingListType {
        return &x
}

func (n *FQReportingListType) New() *FQReportingListType {
        return FQReportingListTypeCreate(FQReportingListType{})
}

func (n *FQReportingListType) Clone() FQReportingListType {
  resource := &FQReportingListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressCollectionStudentListTypeCreate(x AddressCollectionStudentListType) *AddressCollectionStudentListType {
        return &x
}

func (n *AddressCollectionStudentListType) New() *AddressCollectionStudentListType {
        return AddressCollectionStudentListTypeCreate(AddressCollectionStudentListType{})
}

func (n *AddressCollectionStudentListType) Clone() AddressCollectionStudentListType {
  resource := &AddressCollectionStudentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherCodeListTypeCreate(x OtherCodeListType) *OtherCodeListType {
        return &x
}

func (n *OtherCodeListType) New() *OtherCodeListType {
        return OtherCodeListTypeCreate(OtherCodeListType{})
}

func (n *OtherCodeListType) Clone() OtherCodeListType {
  resource := &OtherCodeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CalendarSummaryListTypeCreate(x CalendarSummaryListType) *CalendarSummaryListType {
        return &x
}

func (n *CalendarSummaryListType) New() *CalendarSummaryListType {
        return CalendarSummaryListTypeCreate(CalendarSummaryListType{})
}

func (n *CalendarSummaryListType) Clone() CalendarSummaryListType {
  resource := &CalendarSummaryListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatisticalAreaTypeCreate(x StatisticalAreaType) *StatisticalAreaType {
        return &x
}

func (n *StatisticalAreaType) New() *StatisticalAreaType {
        return StatisticalAreaTypeCreate(StatisticalAreaType{})
}

func (n *StatisticalAreaType) Clone() StatisticalAreaType {
  resource := &StatisticalAreaType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPCodeFrameCreate(x NAPCodeFrame) *NAPCodeFrame {
        return &x
}

func (n *NAPCodeFrame) New() *NAPCodeFrame {
        return NAPCodeFrameCreate(NAPCodeFrame{})
}

func (n *NAPCodeFrame) Clone() NAPCodeFrame {
  resource := &NAPCodeFrame{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EmailListTypeCreate(x EmailListType) *EmailListType {
        return &x
}

func (n *EmailListType) New() *EmailListType {
        return EmailListTypeCreate(EmailListType{})
}

func (n *EmailListType) Clone() EmailListType {
  resource := &EmailListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GradeTypeCreate(x GradeType) *GradeType {
        return &x
}

func (n *GradeType) New() *GradeType {
        return GradeTypeCreate(GradeType{})
}

func (n *GradeType) Clone() GradeType {
  resource := &GradeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LanguageOfInstructionTypeCreate(x LanguageOfInstructionType) *LanguageOfInstructionType {
        return &x
}

func (n *LanguageOfInstructionType) New() *LanguageOfInstructionType {
        return LanguageOfInstructionTypeCreate(LanguageOfInstructionType{})
}

func (n *LanguageOfInstructionType) Clone() LanguageOfInstructionType {
  resource := &LanguageOfInstructionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ActivityTimeTypeCreate(x ActivityTimeType) *ActivityTimeType {
        return &x
}

func (n *ActivityTimeType) New() *ActivityTimeType {
        return ActivityTimeTypeCreate(ActivityTimeType{})
}

func (n *ActivityTimeType) Clone() ActivityTimeType {
  resource := &ActivityTimeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolInfoCreate(x SchoolInfo) *SchoolInfo {
        return &x
}

func (n *SchoolInfo) New() *SchoolInfo {
        return SchoolInfoCreate(SchoolInfo{})
}

func (n *SchoolInfo) Clone() SchoolInfo {
  resource := &SchoolInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_SystemContextListCreate(x SystemRole_SystemContextList) *SystemRole_SystemContextList {
        return &x
}

func (n *SystemRole_SystemContextList) New() *SystemRole_SystemContextList {
        return SystemRole_SystemContextListCreate(SystemRole_SystemContextList{})
}

func (n *SystemRole_SystemContextList) Clone() SystemRole_SystemContextList {
  resource := &SystemRole_SystemContextList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffPersonalCreate(x StaffPersonal) *StaffPersonal {
        return &x
}

func (n *StaffPersonal) New() *StaffPersonal {
        return StaffPersonalCreate(StaffPersonal{})
}

func (n *StaffPersonal) Clone() StaffPersonal {
  resource := &StaffPersonal{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PrincipalInfoTypeCreate(x PrincipalInfoType) *PrincipalInfoType {
        return &x
}

func (n *PrincipalInfoType) New() *PrincipalInfoType {
        return PrincipalInfoTypeCreate(PrincipalInfoType{})
}

func (n *PrincipalInfoType) Clone() PrincipalInfoType {
  resource := &PrincipalInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletResponseTypeCreate(x NAPTestletResponseType) *NAPTestletResponseType {
        return &x
}

func (n *NAPTestletResponseType) New() *NAPTestletResponseType {
        return NAPTestletResponseTypeCreate(NAPTestletResponseType{})
}

func (n *NAPTestletResponseType) Clone() NAPTestletResponseType {
  resource := &NAPTestletResponseType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingEventSubCategoryListTypeCreate(x WellbeingEventSubCategoryListType) *WellbeingEventSubCategoryListType {
        return &x
}

func (n *WellbeingEventSubCategoryListType) New() *WellbeingEventSubCategoryListType {
        return WellbeingEventSubCategoryListTypeCreate(WellbeingEventSubCategoryListType{})
}

func (n *WellbeingEventSubCategoryListType) Clone() WellbeingEventSubCategoryListType {
  resource := &WellbeingEventSubCategoryListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQItemListTypeCreate(x FQItemListType) *FQItemListType {
        return &x
}

func (n *FQItemListType) New() *FQItemListType {
        return FQItemListTypeCreate(FQItemListType{})
}

func (n *FQItemListType) Clone() FQItemListType {
  resource := &FQItemListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusStaffListTypeCreate(x CensusStaffListType) *CensusStaffListType {
        return &x
}

func (n *CensusStaffListType) New() *CensusStaffListType {
        return CensusStaffListTypeCreate(CensusStaffListType{})
}

func (n *CensusStaffListType) Clone() CensusStaffListType {
  resource := &CensusStaffListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGParentTypeCreate(x AGParentType) *AGParentType {
        return &x
}

func (n *AGParentType) New() *AGParentType {
        return AGParentTypeCreate(AGParentType{})
}

func (n *AGParentType) Clone() AGParentType {
  resource := &AGParentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScoreListTypeCreate(x ScoreListType) *ScoreListType {
        return &x
}

func (n *ScoreListType) New() *ScoreListType {
        return ScoreListTypeCreate(ScoreListType{})
}

func (n *ScoreListType) Clone() ScoreListType {
  resource := &ScoreListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingEventLocationDetailsTypeCreate(x WellbeingEventLocationDetailsType) *WellbeingEventLocationDetailsType {
        return &x
}

func (n *WellbeingEventLocationDetailsType) New() *WellbeingEventLocationDetailsType {
        return WellbeingEventLocationDetailsTypeCreate(WellbeingEventLocationDetailsType{})
}

func (n *WellbeingEventLocationDetailsType) Clone() WellbeingEventLocationDetailsType {
  resource := &WellbeingEventLocationDetailsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableCellCreate(x TimeTableCell) *TimeTableCell {
        return &x
}

func (n *TimeTableCell) New() *TimeTableCell {
        return TimeTableCellCreate(TimeTableCell{})
}

func (n *TimeTableCell) Clone() TimeTableCell {
  resource := &TimeTableCell{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolContactTypeCreate(x SchoolContactType) *SchoolContactType {
        return &x
}

func (n *SchoolContactType) New() *SchoolContactType {
        return SchoolContactTypeCreate(SchoolContactType{})
}

func (n *SchoolContactType) Clone() SchoolContactType {
  resource := &SchoolContactType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_RoleListCreate(x SystemRole_RoleList) *SystemRole_RoleList {
        return &x
}

func (n *SystemRole_RoleList) New() *SystemRole_RoleList {
        return SystemRole_RoleListCreate(SystemRole_RoleList{})
}

func (n *SystemRole_RoleList) Clone() SystemRole_RoleList {
  resource := &SystemRole_RoleList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRoleCreate(x SystemRole) *SystemRole {
        return &x
}

func (n *SystemRole) New() *SystemRole {
        return SystemRoleCreate(SystemRole{})
}

func (n *SystemRole) Clone() SystemRole {
  resource := &SystemRole{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentGroupListTypeCreate(x StudentGroupListType) *StudentGroupListType {
        return &x
}

func (n *StudentGroupListType) New() *StudentGroupListType {
        return StudentGroupListTypeCreate(StudentGroupListType{})
}

func (n *StudentGroupListType) Clone() StudentGroupListType {
  resource := &StudentGroupListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGRuleListTypeCreate(x AGRuleListType) *AGRuleListType {
        return &x
}

func (n *AGRuleListType) New() *AGRuleListType {
        return AGRuleListTypeCreate(AGRuleListType{})
}

func (n *AGRuleListType) Clone() AGRuleListType {
  resource := &AGRuleListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolCourseInfoCreate(x SchoolCourseInfo) *SchoolCourseInfo {
        return &x
}

func (n *SchoolCourseInfo) New() *SchoolCourseInfo {
        return SchoolCourseInfoCreate(SchoolCourseInfo{})
}

func (n *SchoolCourseInfo) Clone() SchoolCourseInfo {
  resource := &SchoolCourseInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_RoleCreate(x SystemRole_Role) *SystemRole_Role {
        return &x
}

func (n *SystemRole_Role) New() *SystemRole_Role {
        return SystemRole_RoleCreate(SystemRole_Role{})
}

func (n *SystemRole_Role) Clone() SystemRole_Role {
  resource := &SystemRole_Role{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScheduledActivityOverrideTypeCreate(x ScheduledActivityOverrideType) *ScheduledActivityOverrideType {
        return &x
}

func (n *ScheduledActivityOverrideType) New() *ScheduledActivityOverrideType {
        return ScheduledActivityOverrideTypeCreate(ScheduledActivityOverrideType{})
}

func (n *ScheduledActivityOverrideType) Clone() ScheduledActivityOverrideType {
  resource := &ScheduledActivityOverrideType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ContactTypeCreate(x ContactType) *ContactType {
        return &x
}

func (n *ContactType) New() *ContactType {
        return ContactTypeCreate(ContactType{})
}

func (n *ContactType) Clone() ContactType {
  resource := &ContactType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func Debtor_BilledEntityCreate(x Debtor_BilledEntity) *Debtor_BilledEntity {
        return &x
}

func (n *Debtor_BilledEntity) New() *Debtor_BilledEntity {
        return Debtor_BilledEntityCreate(Debtor_BilledEntity{})
}

func (n *Debtor_BilledEntity) Clone() Debtor_BilledEntity {
  resource := &Debtor_BilledEntity{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeElementType_SpanGapCreate(x TimeElementType_SpanGap) *TimeElementType_SpanGap {
        return &x
}

func (n *TimeElementType_SpanGap) New() *TimeElementType_SpanGap {
        return TimeElementType_SpanGapCreate(TimeElementType_SpanGap{})
}

func (n *TimeElementType_SpanGap) Clone() TimeElementType_SpanGap {
  resource := &TimeElementType_SpanGap{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonalisedPlanCreate(x PersonalisedPlan) *PersonalisedPlan {
        return &x
}

func (n *PersonalisedPlan) New() *PersonalisedPlan {
        return PersonalisedPlanCreate(PersonalisedPlan{})
}

func (n *PersonalisedPlan) Clone() PersonalisedPlan {
  resource := &PersonalisedPlan{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonPicture_PictureSourceCreate(x PersonPicture_PictureSource) *PersonPicture_PictureSource {
        return &x
}

func (n *PersonPicture_PictureSource) New() *PersonPicture_PictureSource {
        return PersonPicture_PictureSourceCreate(PersonPicture_PictureSource{})
}

func (n *PersonPicture_PictureSource) Clone() PersonPicture_PictureSource {
  resource := &PersonPicture_PictureSource{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardTypeCreate(x LearningStandardType) *LearningStandardType {
        return &x
}

func (n *LearningStandardType) New() *LearningStandardType {
        return LearningStandardTypeCreate(LearningStandardType{})
}

func (n *LearningStandardType) Clone() LearningStandardType {
  resource := &LearningStandardType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGRuleTypeCreate(x AGRuleType) *AGRuleType {
        return &x
}

func (n *AGRuleType) New() *AGRuleType {
        return AGRuleTypeCreate(AGRuleType{})
}

func (n *AGRuleType) Clone() AGRuleType {
  resource := &AGRuleType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NameOfRecordTypeCreate(x NameOfRecordType) *NameOfRecordType {
        return &x
}

func (n *NameOfRecordType) New() *NameOfRecordType {
        return NameOfRecordTypeCreate(NameOfRecordType{})
}

func (n *NameOfRecordType) Clone() NameOfRecordType {
  resource := &NameOfRecordType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonInvolvementListTypeCreate(x PersonInvolvementListType) *PersonInvolvementListType {
        return &x
}

func (n *PersonInvolvementListType) New() *PersonInvolvementListType {
        return PersonInvolvementListTypeCreate(PersonInvolvementListType{})
}

func (n *PersonInvolvementListType) Clone() PersonInvolvementListType {
  resource := &PersonInvolvementListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CodeFrameTestItemListTypeCreate(x CodeFrameTestItemListType) *CodeFrameTestItemListType {
        return &x
}

func (n *CodeFrameTestItemListType) New() *CodeFrameTestItemListType {
        return CodeFrameTestItemListTypeCreate(CodeFrameTestItemListType{})
}

func (n *CodeFrameTestItemListType) Clone() CodeFrameTestItemListType {
  resource := &CodeFrameTestItemListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentPackageType_BinaryDataCreate(x AbstractContentPackageType_BinaryData) *AbstractContentPackageType_BinaryData {
        return &x
}

func (n *AbstractContentPackageType_BinaryData) New() *AbstractContentPackageType_BinaryData {
        return AbstractContentPackageType_BinaryDataCreate(AbstractContentPackageType_BinaryData{})
}

func (n *AbstractContentPackageType_BinaryData) Clone() AbstractContentPackageType_BinaryData {
  resource := &AbstractContentPackageType_BinaryData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RelatedLearningStandardItemRefIdTypeCreate(x RelatedLearningStandardItemRefIdType) *RelatedLearningStandardItemRefIdType {
        return &x
}

func (n *RelatedLearningStandardItemRefIdType) New() *RelatedLearningStandardItemRefIdType {
        return RelatedLearningStandardItemRefIdTypeCreate(RelatedLearningStandardItemRefIdType{})
}

func (n *RelatedLearningStandardItemRefIdType) Clone() RelatedLearningStandardItemRefIdType {
  resource := &RelatedLearningStandardItemRefIdType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPCodeFrameTestletListTypeCreate(x NAPCodeFrameTestletListType) *NAPCodeFrameTestletListType {
        return &x
}

func (n *NAPCodeFrameTestletListType) New() *NAPCodeFrameTestletListType {
        return NAPCodeFrameTestletListTypeCreate(NAPCodeFrameTestletListType{})
}

func (n *NAPCodeFrameTestletListType) Clone() NAPCodeFrameTestletListType {
  resource := &NAPCodeFrameTestletListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ContactInfoTypeCreate(x ContactInfoType) *ContactInfoType {
        return &x
}

func (n *ContactInfoType) New() *ContactInfoType {
        return ContactInfoTypeCreate(ContactInfoType{})
}

func (n *ContactInfoType) Clone() ContactInfoType {
  resource := &ContactInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TestDisruptionListTypeCreate(x TestDisruptionListType) *TestDisruptionListType {
        return &x
}

func (n *TestDisruptionListType) New() *TestDisruptionListType {
        return TestDisruptionListTypeCreate(TestDisruptionListType{})
}

func (n *TestDisruptionListType) Clone() TestDisruptionListType {
  resource := &TestDisruptionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeacherCoverTypeCreate(x TeacherCoverType) *TeacherCoverType {
        return &x
}

func (n *TeacherCoverType) New() *TeacherCoverType {
        return TeacherCoverTypeCreate(TeacherCoverType{})
}

func (n *TeacherCoverType) Clone() TeacherCoverType {
  resource := &TeacherCoverType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentPackageType_TextDataCreate(x AbstractContentPackageType_TextData) *AbstractContentPackageType_TextData {
        return &x
}

func (n *AbstractContentPackageType_TextData) New() *AbstractContentPackageType_TextData {
        return AbstractContentPackageType_TextDataCreate(AbstractContentPackageType_TextData{})
}

func (n *AbstractContentPackageType_TextData) Clone() AbstractContentPackageType_TextData {
  resource := &AbstractContentPackageType_TextData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffSubjectListTypeCreate(x StaffSubjectListType) *StaffSubjectListType {
        return &x
}

func (n *StaffSubjectListType) New() *StaffSubjectListType {
        return StaffSubjectListTypeCreate(StaffSubjectListType{})
}

func (n *StaffSubjectListType) Clone() StaffSubjectListType {
  resource := &StaffSubjectListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TestDisruptionTypeCreate(x TestDisruptionType) *TestDisruptionType {
        return &x
}

func (n *TestDisruptionType) New() *TestDisruptionType {
        return TestDisruptionTypeCreate(TestDisruptionType{})
}

func (n *TestDisruptionType) Clone() TestDisruptionType {
  resource := &TestDisruptionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableScheduleCellTypeCreate(x TimeTableScheduleCellType) *TimeTableScheduleCellType {
        return &x
}

func (n *TimeTableScheduleCellType) New() *TimeTableScheduleCellType {
        return TimeTableScheduleCellTypeCreate(TimeTableScheduleCellType{})
}

func (n *TimeTableScheduleCellType) Clone() TimeTableScheduleCellType {
  resource := &TimeTableScheduleCellType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LEAContactTypeCreate(x LEAContactType) *LEAContactType {
        return &x
}

func (n *LEAContactType) New() *LEAContactType {
        return LEAContactTypeCreate(LEAContactType{})
}

func (n *LEAContactType) Clone() LEAContactType {
  resource := &LEAContactType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceBooking_ResourceRefIdCreate(x ResourceBooking_ResourceRefId) *ResourceBooking_ResourceRefId {
        return &x
}

func (n *ResourceBooking_ResourceRefId) New() *ResourceBooking_ResourceRefId {
        return ResourceBooking_ResourceRefIdCreate(ResourceBooking_ResourceRefId{})
}

func (n *ResourceBooking_ResourceRefId) Clone() ResourceBooking_ResourceRefId {
  resource := &ResourceBooking_ResourceRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressCollectionReportingListTypeCreate(x AddressCollectionReportingListType) *AddressCollectionReportingListType {
        return &x
}

func (n *AddressCollectionReportingListType) New() *AddressCollectionReportingListType {
        return AddressCollectionReportingListTypeCreate(AddressCollectionReportingListType{})
}

func (n *AddressCollectionReportingListType) Clone() AddressCollectionReportingListType {
  resource := &AddressCollectionReportingListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonPicture_ParentObjectRefIdCreate(x PersonPicture_ParentObjectRefId) *PersonPicture_ParentObjectRefId {
        return &x
}

func (n *PersonPicture_ParentObjectRefId) New() *PersonPicture_ParentObjectRefId {
        return PersonPicture_ParentObjectRefIdCreate(PersonPicture_ParentObjectRefId{})
}

func (n *PersonPicture_ParentObjectRefId) Clone() PersonPicture_ParentObjectRefId {
  resource := &PersonPicture_ParentObjectRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSchoolEnrollment_HomeroomCreate(x StudentSchoolEnrollment_Homeroom) *StudentSchoolEnrollment_Homeroom {
        return &x
}

func (n *StudentSchoolEnrollment_Homeroom) New() *StudentSchoolEnrollment_Homeroom {
        return StudentSchoolEnrollment_HomeroomCreate(StudentSchoolEnrollment_Homeroom{})
}

func (n *StudentSchoolEnrollment_Homeroom) Clone() StudentSchoolEnrollment_Homeroom {
  resource := &StudentSchoolEnrollment_Homeroom{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGReportingObjectResponseTypeCreate(x AGReportingObjectResponseType) *AGReportingObjectResponseType {
        return &x
}

func (n *AGReportingObjectResponseType) New() *AGReportingObjectResponseType {
        return AGReportingObjectResponseTypeCreate(AGReportingObjectResponseType{})
}

func (n *AGReportingObjectResponseType) Clone() AGReportingObjectResponseType {
  resource := &AGReportingObjectResponseType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingCharacteristicCreate(x WellbeingCharacteristic) *WellbeingCharacteristic {
        return &x
}

func (n *WellbeingCharacteristic) New() *WellbeingCharacteristic {
        return WellbeingCharacteristicCreate(WellbeingCharacteristic{})
}

func (n *WellbeingCharacteristic) Clone() WellbeingCharacteristic {
  resource := &WellbeingCharacteristic{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingEventCategoryListTypeCreate(x WellbeingEventCategoryListType) *WellbeingEventCategoryListType {
        return &x
}

func (n *WellbeingEventCategoryListType) New() *WellbeingEventCategoryListType {
        return WellbeingEventCategoryListTypeCreate(WellbeingEventCategoryListType{})
}

func (n *WellbeingEventCategoryListType) Clone() WellbeingEventCategoryListType {
  resource := &WellbeingEventCategoryListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LResourcesTypeCreate(x LResourcesType) *LResourcesType {
        return &x
}

func (n *LResourcesType) New() *LResourcesType {
        return LResourcesTypeCreate(LResourcesType{})
}

func (n *LResourcesType) Clone() LResourcesType {
  resource := &LResourcesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryItemInfoTypeCreate(x LibraryItemInfoType) *LibraryItemInfoType {
        return &x
}

func (n *LibraryItemInfoType) New() *LibraryItemInfoType {
        return LibraryItemInfoTypeCreate(LibraryItemInfoType{})
}

func (n *LibraryItemInfoType) Clone() LibraryItemInfoType {
  resource := &LibraryItemInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func Identity_SIF_RefIdCreate(x Identity_SIF_RefId) *Identity_SIF_RefId {
        return &x
}

func (n *Identity_SIF_RefId) New() *Identity_SIF_RefId {
        return Identity_SIF_RefIdCreate(Identity_SIF_RefId{})
}

func (n *Identity_SIF_RefId) Clone() Identity_SIF_RefId {
  resource := &Identity_SIF_RefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQContextualQuestionTypeCreate(x FQContextualQuestionType) *FQContextualQuestionType {
        return &x
}

func (n *FQContextualQuestionType) New() *FQContextualQuestionType {
        return FQContextualQuestionTypeCreate(FQContextualQuestionType{})
}

func (n *FQContextualQuestionType) Clone() FQContextualQuestionType {
  resource := &FQContextualQuestionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WithdrawalTimeListTypeCreate(x WithdrawalTimeListType) *WithdrawalTimeListType {
        return &x
}

func (n *WithdrawalTimeListType) New() *WithdrawalTimeListType {
        return WithdrawalTimeListTypeCreate(WithdrawalTimeListType{})
}

func (n *WithdrawalTimeListType) Clone() WithdrawalTimeListType {
  resource := &WithdrawalTimeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func YearLevelsTypeCreate(x YearLevelsType) *YearLevelsType {
        return &x
}

func (n *YearLevelsType) New() *YearLevelsType {
        return YearLevelsTypeCreate(YearLevelsType{})
}

func (n *YearLevelsType) Clone() YearLevelsType {
  resource := &YearLevelsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AttendanceInfoTypeCreate(x AttendanceInfoType) *AttendanceInfoType {
        return &x
}

func (n *AttendanceInfoType) New() *AttendanceInfoType {
        return AttendanceInfoTypeCreate(AttendanceInfoType{})
}

func (n *AttendanceInfoType) Clone() AttendanceInfoType {
  resource := &AttendanceInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeacherListTypeCreate(x TeacherListType) *TeacherListType {
        return &x
}

func (n *TeacherListType) New() *TeacherListType {
        return TeacherListTypeCreate(TeacherListType{})
}

func (n *TeacherListType) Clone() TeacherListType {
  resource := &TeacherListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPLANClassListTypeCreate(x NAPLANClassListType) *NAPLANClassListType {
        return &x
}

func (n *NAPLANClassListType) New() *NAPLANClassListType {
        return NAPLANClassListTypeCreate(NAPLANClassListType{})
}

func (n *NAPLANClassListType) Clone() NAPLANClassListType {
  resource := &NAPLANClassListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonInfoTypeCreate(x PersonInfoType) *PersonInfoType {
        return &x
}

func (n *PersonInfoType) New() *PersonInfoType {
        return PersonInfoTypeCreate(PersonInfoType{})
}

func (n *PersonInfoType) Clone() PersonInfoType {
  resource := &PersonInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardItemCreate(x LearningStandardItem) *LearningStandardItem {
        return &x
}

func (n *LearningStandardItem) New() *LearningStandardItem {
        return LearningStandardItemCreate(LearningStandardItem{})
}

func (n *LearningStandardItem) Clone() LearningStandardItem {
  resource := &LearningStandardItem{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SIF_MetadataTypeCreate(x SIF_MetadataType) *SIF_MetadataType {
        return &x
}

func (n *SIF_MetadataType) New() *SIF_MetadataType {
        return SIF_MetadataTypeCreate(SIF_MetadataType{})
}

func (n *SIF_MetadataType) Clone() SIF_MetadataType {
  resource := &SIF_MetadataType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingDocumentTypeCreate(x WellbeingDocumentType) *WellbeingDocumentType {
        return &x
}

func (n *WellbeingDocumentType) New() *WellbeingDocumentType {
        return WellbeingDocumentTypeCreate(WellbeingDocumentType{})
}

func (n *WellbeingDocumentType) Clone() WellbeingDocumentType {
  resource := &WellbeingDocumentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PurchasingItemTypeCreate(x PurchasingItemType) *PurchasingItemType {
        return &x
}

func (n *PurchasingItemType) New() *PurchasingItemType {
        return PurchasingItemTypeCreate(PurchasingItemType{})
}

func (n *PurchasingItemType) Clone() PurchasingItemType {
  resource := &PurchasingItemType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GradingAssignmentCreate(x GradingAssignment) *GradingAssignment {
        return &x
}

func (n *GradingAssignment) New() *GradingAssignment {
        return GradingAssignmentCreate(GradingAssignment{})
}

func (n *GradingAssignment) Clone() GradingAssignment {
  resource := &GradingAssignment{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func JournalCreate(x Journal) *Journal {
        return &x
}

func (n *Journal) New() *Journal {
        return JournalCreate(Journal{})
}

func (n *Journal) Clone() Journal {
  resource := &Journal{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardListTypeCreate(x LearningStandardListType) *LearningStandardListType {
        return &x
}

func (n *LearningStandardListType) New() *LearningStandardListType {
        return LearningStandardListTypeCreate(LearningStandardListType{})
}

func (n *LearningStandardListType) Clone() LearningStandardListType {
  resource := &LearningStandardListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTablePeriodListTypeCreate(x TimeTablePeriodListType) *TimeTablePeriodListType {
        return &x
}

func (n *TimeTablePeriodListType) New() *TimeTablePeriodListType {
        return TimeTablePeriodListTypeCreate(TimeTablePeriodListType{})
}

func (n *TimeTablePeriodListType) Clone() TimeTablePeriodListType {
  resource := &TimeTablePeriodListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AuthorsTypeCreate(x AuthorsType) *AuthorsType {
        return &x
}

func (n *AuthorsType) New() *AuthorsType {
        return AuthorsTypeCreate(AuthorsType{})
}

func (n *AuthorsType) Clone() AuthorsType {
  resource := &AuthorsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSubjectChoiceListTypeCreate(x StudentSubjectChoiceListType) *StudentSubjectChoiceListType {
        return &x
}

func (n *StudentSubjectChoiceListType) New() *StudentSubjectChoiceListType {
        return StudentSubjectChoiceListTypeCreate(StudentSubjectChoiceListType{})
}

func (n *StudentSubjectChoiceListType) Clone() StudentSubjectChoiceListType {
  resource := &StudentSubjectChoiceListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatsCohortYearLevelListTypeCreate(x StatsCohortYearLevelListType) *StatsCohortYearLevelListType {
        return &x
}

func (n *StatsCohortYearLevelListType) New() *StatsCohortYearLevelListType {
        return StatsCohortYearLevelListTypeCreate(StatsCohortYearLevelListType{})
}

func (n *StatsCohortYearLevelListType) Clone() StatsCohortYearLevelListType {
  resource := &StatsCohortYearLevelListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DemographicsTypeCreate(x DemographicsType) *DemographicsType {
        return &x
}

func (n *DemographicsType) New() *DemographicsType {
        return DemographicsTypeCreate(DemographicsType{})
}

func (n *DemographicsType) Clone() DemographicsType {
  resource := &DemographicsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTablePeriodTypeCreate(x TimeTablePeriodType) *TimeTablePeriodType {
        return &x
}

func (n *TimeTablePeriodType) New() *TimeTablePeriodType {
        return TimeTablePeriodTypeCreate(TimeTablePeriodType{})
}

func (n *TimeTablePeriodType) Clone() TimeTablePeriodType {
  resource := &TimeTablePeriodType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningStandardsTypeCreate(x LearningStandardsType) *LearningStandardsType {
        return &x
}

func (n *LearningStandardsType) New() *LearningStandardsType {
        return LearningStandardsTypeCreate(LearningStandardsType{})
}

func (n *LearningStandardsType) Clone() LearningStandardsType {
  resource := &LearningStandardsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolGroupListTypeCreate(x SchoolGroupListType) *SchoolGroupListType {
        return &x
}

func (n *SchoolGroupListType) New() *SchoolGroupListType {
        return SchoolGroupListTypeCreate(SchoolGroupListType{})
}

func (n *SchoolGroupListType) Clone() SchoolGroupListType {
  resource := &SchoolGroupListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressTypeCreate(x AddressType) *AddressType {
        return &x
}

func (n *AddressType) New() *AddressType {
        return AddressTypeCreate(AddressType{})
}

func (n *AddressType) Clone() AddressType {
  resource := &AddressType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGContextualQuestionListTypeCreate(x AGContextualQuestionListType) *AGContextualQuestionListType {
        return &x
}

func (n *AGContextualQuestionListType) New() *AGContextualQuestionListType {
        return AGContextualQuestionListTypeCreate(AGContextualQuestionListType{})
}

func (n *AGContextualQuestionListType) Clone() AGContextualQuestionListType {
  resource := &AGContextualQuestionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DetentionContainerTypeCreate(x DetentionContainerType) *DetentionContainerType {
        return &x
}

func (n *DetentionContainerType) New() *DetentionContainerType {
        return DetentionContainerTypeCreate(DetentionContainerType{})
}

func (n *DetentionContainerType) Clone() DetentionContainerType {
  resource := &DetentionContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingEventCreate(x WellbeingEvent) *WellbeingEvent {
        return &x
}

func (n *WellbeingEvent) New() *WellbeingEvent {
        return WellbeingEventCreate(WellbeingEvent{})
}

func (n *WellbeingEvent) Clone() WellbeingEvent {
  resource := &WellbeingEvent{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressListTypeCreate(x AddressListType) *AddressListType {
        return &x
}

func (n *AddressListType) New() *AddressListType {
        return AddressListTypeCreate(AddressListType{})
}

func (n *AddressListType) Clone() AddressListType {
  resource := &AddressListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQContextualQuestionListTypeCreate(x FQContextualQuestionListType) *FQContextualQuestionListType {
        return &x
}

func (n *FQContextualQuestionListType) New() *FQContextualQuestionListType {
        return FQContextualQuestionListTypeCreate(FQContextualQuestionListType{})
}

func (n *FQContextualQuestionListType) Clone() FQContextualQuestionListType {
  resource := &FQContextualQuestionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PersonInvolvementType_PersonRefIdCreate(x PersonInvolvementType_PersonRefId) *PersonInvolvementType_PersonRefId {
        return &x
}

func (n *PersonInvolvementType_PersonRefId) New() *PersonInvolvementType_PersonRefId {
        return PersonInvolvementType_PersonRefIdCreate(PersonInvolvementType_PersonRefId{})
}

func (n *PersonInvolvementType_PersonRefId) Clone() PersonInvolvementType_PersonRefId {
  resource := &PersonInvolvementType_PersonRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingResponseCreate(x WellbeingResponse) *WellbeingResponse {
        return &x
}

func (n *WellbeingResponse) New() *WellbeingResponse {
        return WellbeingResponseCreate(WellbeingResponse{})
}

func (n *WellbeingResponse) Clone() WellbeingResponse {
  resource := &WellbeingResponse{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TermInfoCreate(x TermInfo) *TermInfo {
        return &x
}

func (n *TermInfo) New() *TermInfo {
        return TermInfoCreate(TermInfo{})
}

func (n *TermInfo) Clone() TermInfo {
  resource := &TermInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolProgramListTypeCreate(x SchoolProgramListType) *SchoolProgramListType {
        return &x
}

func (n *SchoolProgramListType) New() *SchoolProgramListType {
        return SchoolProgramListTypeCreate(SchoolProgramListType{})
}

func (n *SchoolProgramListType) Clone() SchoolProgramListType {
  resource := &SchoolProgramListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ChargedLocationInfoCreate(x ChargedLocationInfo) *ChargedLocationInfo {
        return &x
}

func (n *ChargedLocationInfo) New() *ChargedLocationInfo {
        return ChargedLocationInfoCreate(ChargedLocationInfo{})
}

func (n *ChargedLocationInfo) Clone() ChargedLocationInfo {
  resource := &ChargedLocationInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PlanRequiredListTypeCreate(x PlanRequiredListType) *PlanRequiredListType {
        return &x
}

func (n *PlanRequiredListType) New() *PlanRequiredListType {
        return PlanRequiredListTypeCreate(PlanRequiredListType{})
}

func (n *PlanRequiredListType) Clone() PlanRequiredListType {
  resource := &PlanRequiredListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentsTypeCreate(x StudentsType) *StudentsType {
        return &x
}

func (n *StudentsType) New() *StudentsType {
        return StudentsTypeCreate(StudentsType{})
}

func (n *StudentsType) Clone() StudentsType {
  resource := &StudentsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FollowUpActionTypeCreate(x FollowUpActionType) *FollowUpActionType {
        return &x
}

func (n *FollowUpActionType) New() *FollowUpActionType {
        return FollowUpActionTypeCreate(FollowUpActionType{})
}

func (n *FollowUpActionType) Clone() FollowUpActionType {
  resource := &FollowUpActionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffListTypeCreate(x StaffListType) *StaffListType {
        return &x
}

func (n *StaffListType) New() *StaffListType {
        return StaffListTypeCreate(StaffListType{})
}

func (n *StaffListType) Clone() StaffListType {
  resource := &StaffListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_ResourceUsageContentTypeCreate(x ResourceUsage_ResourceUsageContentType) *ResourceUsage_ResourceUsageContentType {
        return &x
}

func (n *ResourceUsage_ResourceUsageContentType) New() *ResourceUsage_ResourceUsageContentType {
        return ResourceUsage_ResourceUsageContentTypeCreate(ResourceUsage_ResourceUsageContentType{})
}

func (n *ResourceUsage_ResourceUsageContentType) Clone() ResourceUsage_ResourceUsageContentType {
  resource := &ResourceUsage_ResourceUsageContentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WithdrawalTypeCreate(x WithdrawalType) *WithdrawalType {
        return &x
}

func (n *WithdrawalType) New() *WithdrawalType {
        return WithdrawalTypeCreate(WithdrawalType{})
}

func (n *WithdrawalType) Clone() WithdrawalType {
  resource := &WithdrawalType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SymptomListTypeCreate(x SymptomListType) *SymptomListType {
        return &x
}

func (n *SymptomListType) New() *SymptomListType {
        return SymptomListTypeCreate(SymptomListType{})
}

func (n *SymptomListType) Clone() SymptomListType {
  resource := &SymptomListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func Invoice_InvoicedEntityCreate(x Invoice_InvoicedEntity) *Invoice_InvoicedEntity {
        return &x
}

func (n *Invoice_InvoicedEntity) New() *Invoice_InvoicedEntity {
        return Invoice_InvoicedEntityCreate(Invoice_InvoicedEntity{})
}

func (n *Invoice_InvoicedEntity) Clone() Invoice_InvoicedEntity {
  resource := &Invoice_InvoicedEntity{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingPersonLink_PersonRefIdCreate(x WellbeingPersonLink_PersonRefId) *WellbeingPersonLink_PersonRefId {
        return &x
}

func (n *WellbeingPersonLink_PersonRefId) New() *WellbeingPersonLink_PersonRefId {
        return WellbeingPersonLink_PersonRefIdCreate(WellbeingPersonLink_PersonRefId{})
}

func (n *WellbeingPersonLink_PersonRefId) Clone() WellbeingPersonLink_PersonRefId {
  resource := &WellbeingPersonLink_PersonRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FollowUpActionListTypeCreate(x FollowUpActionListType) *FollowUpActionListType {
        return &x
}

func (n *FollowUpActionListType) New() *FollowUpActionListType {
        return FollowUpActionListTypeCreate(FollowUpActionListType{})
}

func (n *FollowUpActionListType) Clone() FollowUpActionListType {
  resource := &FollowUpActionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AwardContainerTypeCreate(x AwardContainerType) *AwardContainerType {
        return &x
}

func (n *AwardContainerType) New() *AwardContainerType {
        return AwardContainerTypeCreate(AwardContainerType{})
}

func (n *AwardContainerType) Clone() AwardContainerType {
  resource := &AwardContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ExpenseAccountTypeCreate(x ExpenseAccountType) *ExpenseAccountType {
        return &x
}

func (n *ExpenseAccountType) New() *ExpenseAccountType {
        return ExpenseAccountTypeCreate(ExpenseAccountType{})
}

func (n *ExpenseAccountType) Clone() ExpenseAccountType {
  resource := &ExpenseAccountType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MedicationTypeCreate(x MedicationType) *MedicationType {
        return &x
}

func (n *MedicationType) New() *MedicationType {
        return MedicationTypeCreate(MedicationType{})
}

func (n *MedicationType) Clone() MedicationType {
  resource := &MedicationType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_RoleScopeRefIdCreate(x SystemRole_RoleScopeRefId) *SystemRole_RoleScopeRefId {
        return &x
}

func (n *SystemRole_RoleScopeRefId) New() *SystemRole_RoleScopeRefId {
        return SystemRole_RoleScopeRefIdCreate(SystemRole_RoleScopeRefId{})
}

func (n *SystemRole_RoleScopeRefId) Clone() SystemRole_RoleScopeRefId {
  resource := &SystemRole_RoleScopeRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_ResourceReportLineCreate(x ResourceUsage_ResourceReportLine) *ResourceUsage_ResourceReportLine {
        return &x
}

func (n *ResourceUsage_ResourceReportLine) New() *ResourceUsage_ResourceReportLine {
        return ResourceUsage_ResourceReportLineCreate(ResourceUsage_ResourceReportLine{})
}

func (n *ResourceUsage_ResourceReportLine) Clone() ResourceUsage_ResourceReportLine {
  resource := &ResourceUsage_ResourceReportLine{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ProgramFundingSourcesTypeCreate(x ProgramFundingSourcesType) *ProgramFundingSourcesType {
        return &x
}

func (n *ProgramFundingSourcesType) New() *ProgramFundingSourcesType {
        return ProgramFundingSourcesTypeCreate(ProgramFundingSourcesType{})
}

func (n *ProgramFundingSourcesType) Clone() ProgramFundingSourcesType {
  resource := &ProgramFundingSourcesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StrategiesTypeCreate(x StrategiesType) *StrategiesType {
        return &x
}

func (n *StrategiesType) New() *StrategiesType {
        return StrategiesTypeCreate(StrategiesType{})
}

func (n *StrategiesType) Clone() StrategiesType {
  resource := &StrategiesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func YearLevelEnrollmentListTypeCreate(x YearLevelEnrollmentListType) *YearLevelEnrollmentListType {
        return &x
}

func (n *YearLevelEnrollmentListType) New() *YearLevelEnrollmentListType {
        return YearLevelEnrollmentListTypeCreate(YearLevelEnrollmentListType{})
}

func (n *YearLevelEnrollmentListType) Clone() YearLevelEnrollmentListType {
  resource := &YearLevelEnrollmentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_RoleScopeListCreate(x SystemRole_RoleScopeList) *SystemRole_RoleScopeList {
        return &x
}

func (n *SystemRole_RoleScopeList) New() *SystemRole_RoleScopeList {
        return SystemRole_RoleScopeListCreate(SystemRole_RoleScopeList{})
}

func (n *SystemRole_RoleScopeList) Clone() SystemRole_RoleScopeList {
  resource := &SystemRole_RoleScopeList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletContentTypeCreate(x NAPTestletContentType) *NAPTestletContentType {
        return &x
}

func (n *NAPTestletContentType) New() *NAPTestletContentType {
        return NAPTestletContentTypeCreate(NAPTestletContentType{})
}

func (n *NAPTestletContentType) Clone() NAPTestletContentType {
  resource := &NAPTestletContentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatementsTypeCreate(x StatementsType) *StatementsType {
        return &x
}

func (n *StatementsType) New() *StatementsType {
        return StatementsTypeCreate(StatementsType{})
}

func (n *StatementsType) Clone() StatementsType {
  resource := &StatementsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletCodeFrameTypeCreate(x NAPTestletCodeFrameType) *NAPTestletCodeFrameType {
        return &x
}

func (n *NAPTestletCodeFrameType) New() *NAPTestletCodeFrameType {
        return NAPTestletCodeFrameTypeCreate(NAPTestletCodeFrameType{})
}

func (n *NAPTestletCodeFrameType) Clone() NAPTestletCodeFrameType {
  resource := &NAPTestletCodeFrameType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentElementType_TextDataCreate(x AbstractContentElementType_TextData) *AbstractContentElementType_TextData {
        return &x
}

func (n *AbstractContentElementType_TextData) New() *AbstractContentElementType_TextData {
        return AbstractContentElementType_TextDataCreate(AbstractContentElementType_TextData{})
}

func (n *AbstractContentElementType_TextData) Clone() AbstractContentElementType_TextData {
  resource := &AbstractContentElementType_TextData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ACStrandSubjectAreaTypeCreate(x ACStrandSubjectAreaType) *ACStrandSubjectAreaType {
        return &x
}

func (n *ACStrandSubjectAreaType) New() *ACStrandSubjectAreaType {
        return ACStrandSubjectAreaTypeCreate(ACStrandSubjectAreaType{})
}

func (n *ACStrandSubjectAreaType) Clone() ACStrandSubjectAreaType {
  resource := &ACStrandSubjectAreaType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentElementType_ReferenceCreate(x AbstractContentElementType_Reference) *AbstractContentElementType_Reference {
        return &x
}

func (n *AbstractContentElementType_Reference) New() *AbstractContentElementType_Reference {
        return AbstractContentElementType_ReferenceCreate(AbstractContentElementType_Reference{})
}

func (n *AbstractContentElementType_Reference) Clone() AbstractContentElementType_Reference {
  resource := &AbstractContentElementType_Reference{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningObjectivesTypeCreate(x LearningObjectivesType) *LearningObjectivesType {
        return &x
}

func (n *LearningObjectivesType) New() *LearningObjectivesType {
        return LearningObjectivesTypeCreate(LearningObjectivesType{})
}

func (n *LearningObjectivesType) Clone() LearningObjectivesType {
  resource := &LearningObjectivesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusCollectionCreate(x CensusCollection) *CensusCollection {
        return &x
}

func (n *CensusCollection) New() *CensusCollection {
        return CensusCollectionCreate(CensusCollection{})
}

func (n *CensusCollection) Clone() CensusCollection {
  resource := &CensusCollection{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentActivityParticipationCreate(x StudentActivityParticipation) *StudentActivityParticipation {
        return &x
}

func (n *StudentActivityParticipation) New() *StudentActivityParticipation {
        return StudentActivityParticipationCreate(StudentActivityParticipation{})
}

func (n *StudentActivityParticipation) Clone() StudentActivityParticipation {
  resource := &StudentActivityParticipation{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPSubscoreTypeCreate(x NAPSubscoreType) *NAPSubscoreType {
        return &x
}

func (n *NAPSubscoreType) New() *NAPSubscoreType {
        return NAPSubscoreTypeCreate(NAPSubscoreType{})
}

func (n *NAPSubscoreType) Clone() NAPSubscoreType {
  resource := &NAPSubscoreType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherWellbeingResponseContainerTypeCreate(x OtherWellbeingResponseContainerType) *OtherWellbeingResponseContainerType {
        return &x
}

func (n *OtherWellbeingResponseContainerType) New() *OtherWellbeingResponseContainerType {
        return OtherWellbeingResponseContainerTypeCreate(OtherWellbeingResponseContainerType{})
}

func (n *OtherWellbeingResponseContainerType) Clone() OtherWellbeingResponseContainerType {
  resource := &OtherWellbeingResponseContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingDocumentListTypeCreate(x WellbeingDocumentListType) *WellbeingDocumentListType {
        return &x
}

func (n *WellbeingDocumentListType) New() *WellbeingDocumentListType {
        return WellbeingDocumentListTypeCreate(WellbeingDocumentListType{})
}

func (n *WellbeingDocumentListType) Clone() WellbeingDocumentListType {
  resource := &WellbeingDocumentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DomainScoreTypeCreate(x DomainScoreType) *DomainScoreType {
        return &x
}

func (n *DomainScoreType) New() *DomainScoreType {
        return DomainScoreTypeCreate(DomainScoreType{})
}

func (n *DomainScoreType) Clone() DomainScoreType {
  resource := &DomainScoreType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AssignmentListTypeCreate(x AssignmentListType) *AssignmentListType {
        return &x
}

func (n *AssignmentListType) New() *AssignmentListType {
        return AssignmentListTypeCreate(AssignmentListType{})
}

func (n *AssignmentListType) Clone() AssignmentListType {
  resource := &AssignmentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AggregateStatisticInfoCreate(x AggregateStatisticInfo) *AggregateStatisticInfo {
        return &x
}

func (n *AggregateStatisticInfo) New() *AggregateStatisticInfo {
        return AggregateStatisticInfoCreate(AggregateStatisticInfo{})
}

func (n *AggregateStatisticInfo) Clone() AggregateStatisticInfo {
  resource := &AggregateStatisticInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SubstituteItemTypeCreate(x SubstituteItemType) *SubstituteItemType {
        return &x
}

func (n *SubstituteItemType) New() *SubstituteItemType {
        return SubstituteItemTypeCreate(SubstituteItemType{})
}

func (n *SubstituteItemType) Clone() SubstituteItemType {
  resource := &SubstituteItemType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SIF_ExtendedElementsType_SIF_ExtendedElementCreate(x SIF_ExtendedElementsType_SIF_ExtendedElement) *SIF_ExtendedElementsType_SIF_ExtendedElement {
        return &x
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) New() *SIF_ExtendedElementsType_SIF_ExtendedElement {
        return SIF_ExtendedElementsType_SIF_ExtendedElementCreate(SIF_ExtendedElementsType_SIF_ExtendedElement{})
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) Clone() SIF_ExtendedElementsType_SIF_ExtendedElement {
  resource := &SIF_ExtendedElementsType_SIF_ExtendedElement{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatsCohortYearLevelTypeCreate(x StatsCohortYearLevelType) *StatsCohortYearLevelType {
        return &x
}

func (n *StatsCohortYearLevelType) New() *StatsCohortYearLevelType {
        return StatsCohortYearLevelTypeCreate(StatsCohortYearLevelType{})
}

func (n *StatsCohortYearLevelType) Clone() StatsCohortYearLevelType {
  resource := &StatsCohortYearLevelType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingPersonLinkCreate(x WellbeingPersonLink) *WellbeingPersonLink {
        return &x
}

func (n *WellbeingPersonLink) New() *WellbeingPersonLink {
        return WellbeingPersonLinkCreate(WellbeingPersonLink{})
}

func (n *WellbeingPersonLink) Clone() WellbeingPersonLink {
  resource := &WellbeingPersonLink{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentContactRelationshipCreate(x StudentContactRelationship) *StudentContactRelationship {
        return &x
}

func (n *StudentContactRelationship) New() *StudentContactRelationship {
        return StudentContactRelationshipCreate(StudentContactRelationship{})
}

func (n *StudentContactRelationship) Clone() StudentContactRelationship {
  resource := &StudentContactRelationship{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func IdentityAssertionsType_IdentityAssertionCreate(x IdentityAssertionsType_IdentityAssertion) *IdentityAssertionsType_IdentityAssertion {
        return &x
}

func (n *IdentityAssertionsType_IdentityAssertion) New() *IdentityAssertionsType_IdentityAssertion {
        return IdentityAssertionsType_IdentityAssertionCreate(IdentityAssertionsType_IdentityAssertion{})
}

func (n *IdentityAssertionsType_IdentityAssertion) Clone() IdentityAssertionsType_IdentityAssertion {
  resource := &IdentityAssertionsType_IdentityAssertion{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FineInfoListTypeCreate(x FineInfoListType) *FineInfoListType {
        return &x
}

func (n *FineInfoListType) New() *FineInfoListType {
        return FineInfoListTypeCreate(FineInfoListType{})
}

func (n *FineInfoListType) Clone() FineInfoListType {
  resource := &FineInfoListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GenericRubricTypeCreate(x GenericRubricType) *GenericRubricType {
        return &x
}

func (n *GenericRubricType) New() *GenericRubricType {
        return GenericRubricTypeCreate(GenericRubricType{})
}

func (n *GenericRubricType) Clone() GenericRubricType {
  resource := &GenericRubricType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGRoundListTypeCreate(x AGRoundListType) *AGRoundListType {
        return &x
}

func (n *AGRoundListType) New() *AGRoundListType {
        return AGRoundListTypeCreate(AGRoundListType{})
}

func (n *AGRoundListType) Clone() AGRoundListType {
  resource := &AGRoundListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusStudentListTypeCreate(x CensusStudentListType) *CensusStudentListType {
        return &x
}

func (n *CensusStudentListType) New() *CensusStudentListType {
        return CensusStudentListTypeCreate(CensusStudentListType{})
}

func (n *CensusStudentListType) Clone() CensusStudentListType {
  resource := &CensusStudentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MarkerTypeCreate(x MarkerType) *MarkerType {
        return &x
}

func (n *MarkerType) New() *MarkerType {
        return MarkerTypeCreate(MarkerType{})
}

func (n *MarkerType) Clone() MarkerType {
  resource := &MarkerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeElementTypeCreate(x TimeElementType) *TimeElementType {
        return &x
}

func (n *TimeElementType) New() *TimeElementType {
        return TimeElementTypeCreate(TimeElementType{})
}

func (n *TimeElementType) Clone() TimeElementType {
  resource := &TimeElementType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ComponentsTypeCreate(x ComponentsType) *ComponentsType {
        return &x
}

func (n *ComponentsType) New() *ComponentsType {
        return ComponentsTypeCreate(ComponentsType{})
}

func (n *ComponentsType) Clone() ComponentsType {
  resource := &ComponentsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherCodeListType_OtherCodeCreate(x OtherCodeListType_OtherCode) *OtherCodeListType_OtherCode {
        return &x
}

func (n *OtherCodeListType_OtherCode) New() *OtherCodeListType_OtherCode {
        return OtherCodeListType_OtherCodeCreate(OtherCodeListType_OtherCode{})
}

func (n *OtherCodeListType_OtherCode) Clone() OtherCodeListType_OtherCode {
  resource := &OtherCodeListType_OtherCode{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentEntryContainerTypeCreate(x StudentEntryContainerType) *StudentEntryContainerType {
        return &x
}

func (n *StudentEntryContainerType) New() *StudentEntryContainerType {
        return StudentEntryContainerTypeCreate(StudentEntryContainerType{})
}

func (n *StudentEntryContainerType) Clone() StudentEntryContainerType {
  resource := &StudentEntryContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableScheduleCellListTypeCreate(x TimeTableScheduleCellListType) *TimeTableScheduleCellListType {
        return &x
}

func (n *TimeTableScheduleCellListType) New() *TimeTableScheduleCellListType {
        return TimeTableScheduleCellListTypeCreate(TimeTableScheduleCellListType{})
}

func (n *TimeTableScheduleCellListType) Clone() TimeTableScheduleCellListType {
  resource := &TimeTableScheduleCellListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DwellingArrangementTypeCreate(x DwellingArrangementType) *DwellingArrangementType {
        return &x
}

func (n *DwellingArrangementType) New() *DwellingArrangementType {
        return DwellingArrangementTypeCreate(DwellingArrangementType{})
}

func (n *DwellingArrangementType) Clone() DwellingArrangementType {
  resource := &DwellingArrangementType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGRoundTypeCreate(x AGRoundType) *AGRoundType {
        return &x
}

func (n *AGRoundType) New() *AGRoundType {
        return AGRoundTypeCreate(AGRoundType{})
}

func (n *AGRoundType) Clone() AGRoundType {
  resource := &AGRoundType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_RoleScopeCreate(x SystemRole_RoleScope) *SystemRole_RoleScope {
        return &x
}

func (n *SystemRole_RoleScope) New() *SystemRole_RoleScope {
        return SystemRole_RoleScopeCreate(SystemRole_RoleScope{})
}

func (n *SystemRole_RoleScope) Clone() SystemRole_RoleScope {
  resource := &SystemRole_RoleScope{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ProgramFundingSourceTypeCreate(x ProgramFundingSourceType) *ProgramFundingSourceType {
        return &x
}

func (n *ProgramFundingSourceType) New() *ProgramFundingSourceType {
        return ProgramFundingSourceTypeCreate(ProgramFundingSourceType{})
}

func (n *ProgramFundingSourceType) Clone() ProgramFundingSourceType {
  resource := &ProgramFundingSourceType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OtherNamesTypeCreate(x OtherNamesType) *OtherNamesType {
        return &x
}

func (n *OtherNamesType) New() *OtherNamesType {
        return OtherNamesTypeCreate(OtherNamesType{})
}

func (n *OtherNamesType) Clone() OtherNamesType {
  resource := &OtherNamesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolCourseInfoOverrideTypeCreate(x SchoolCourseInfoOverrideType) *SchoolCourseInfoOverrideType {
        return &x
}

func (n *SchoolCourseInfoOverrideType) New() *SchoolCourseInfoOverrideType {
        return SchoolCourseInfoOverrideTypeCreate(SchoolCourseInfoOverrideType{})
}

func (n *SchoolCourseInfoOverrideType) Clone() SchoolCourseInfoOverrideType {
  resource := &SchoolCourseInfoOverrideType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestCreate(x NAPTest) *NAPTest {
        return &x
}

func (n *NAPTest) New() *NAPTest {
        return NAPTestCreate(NAPTest{})
}

func (n *NAPTest) Clone() NAPTest {
  resource := &NAPTest{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AddressCollectionReportingTypeCreate(x AddressCollectionReportingType) *AddressCollectionReportingType {
        return &x
}

func (n *AddressCollectionReportingType) New() *AddressCollectionReportingType {
        return AddressCollectionReportingTypeCreate(AddressCollectionReportingType{})
}

func (n *AddressCollectionReportingType) Clone() AddressCollectionReportingType {
  resource := &AddressCollectionReportingType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPWritingRubricListTypeCreate(x NAPWritingRubricListType) *NAPWritingRubricListType {
        return &x
}

func (n *NAPWritingRubricListType) New() *NAPWritingRubricListType {
        return NAPWritingRubricListTypeCreate(NAPWritingRubricListType{})
}

func (n *NAPWritingRubricListType) Clone() NAPWritingRubricListType {
  resource := &NAPWritingRubricListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestItemListTypeCreate(x NAPTestItemListType) *NAPTestItemListType {
        return &x
}

func (n *NAPTestItemListType) New() *NAPTestItemListType {
        return NAPTestItemListTypeCreate(NAPTestItemListType{})
}

func (n *NAPTestItemListType) Clone() NAPTestItemListType {
  resource := &NAPTestItemListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupScheduleListTypeCreate(x TeachingGroupScheduleListType) *TeachingGroupScheduleListType {
        return &x
}

func (n *TeachingGroupScheduleListType) New() *TeachingGroupScheduleListType {
        return TeachingGroupScheduleListTypeCreate(TeachingGroupScheduleListType{})
}

func (n *TeachingGroupScheduleListType) Clone() TeachingGroupScheduleListType {
  resource := &TeachingGroupScheduleListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AttendanceCodeTypeCreate(x AttendanceCodeType) *AttendanceCodeType {
        return &x
}

func (n *AttendanceCodeType) New() *AttendanceCodeType {
        return AttendanceCodeTypeCreate(AttendanceCodeType{})
}

func (n *AttendanceCodeType) Clone() AttendanceCodeType {
  resource := &AttendanceCodeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AttendanceTimesTypeCreate(x AttendanceTimesType) *AttendanceTimesType {
        return &x
}

func (n *AttendanceTimesType) New() *AttendanceTimesType {
        return AttendanceTimesTypeCreate(AttendanceTimesType{})
}

func (n *AttendanceTimesType) Clone() AttendanceTimesType {
  resource := &AttendanceTimesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentAttendanceCollectionReportingTypeCreate(x StudentAttendanceCollectionReportingType) *StudentAttendanceCollectionReportingType {
        return &x
}

func (n *StudentAttendanceCollectionReportingType) New() *StudentAttendanceCollectionReportingType {
        return StudentAttendanceCollectionReportingTypeCreate(StudentAttendanceCollectionReportingType{})
}

func (n *StudentAttendanceCollectionReportingType) Clone() StudentAttendanceCollectionReportingType {
  resource := &StudentAttendanceCollectionReportingType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StaffAssignmentCreate(x StaffAssignment) *StaffAssignment {
        return &x
}

func (n *StaffAssignment) New() *StaffAssignment {
        return StaffAssignmentCreate(StaffAssignment{})
}

func (n *StaffAssignment) Clone() StaffAssignment {
  resource := &StaffAssignment{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RoomListTypeCreate(x RoomListType) *RoomListType {
        return &x
}

func (n *RoomListType) New() *RoomListType {
        return RoomListTypeCreate(RoomListType{})
}

func (n *RoomListType) Clone() RoomListType {
  resource := &RoomListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CheckoutInfoTypeCreate(x CheckoutInfoType) *CheckoutInfoType {
        return &x
}

func (n *CheckoutInfoType) New() *CheckoutInfoType {
        return CheckoutInfoTypeCreate(CheckoutInfoType{})
}

func (n *CheckoutInfoType) Clone() CheckoutInfoType {
  resource := &CheckoutInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_ResourceReportLineListCreate(x ResourceUsage_ResourceReportLineList) *ResourceUsage_ResourceReportLineList {
        return &x
}

func (n *ResourceUsage_ResourceReportLineList) New() *ResourceUsage_ResourceReportLineList {
        return ResourceUsage_ResourceReportLineListCreate(ResourceUsage_ResourceReportLineList{})
}

func (n *ResourceUsage_ResourceReportLineList) Clone() ResourceUsage_ResourceReportLineList {
  resource := &ResourceUsage_ResourceReportLineList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RoomInfoCreate(x RoomInfo) *RoomInfo {
        return &x
}

func (n *RoomInfo) New() *RoomInfo {
        return RoomInfoCreate(RoomInfo{})
}

func (n *RoomInfo) Clone() RoomInfo {
  resource := &RoomInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AccountCodeListTypeCreate(x AccountCodeListType) *AccountCodeListType {
        return &x
}

func (n *AccountCodeListType) New() *AccountCodeListType {
        return AccountCodeListTypeCreate(AccountCodeListType{})
}

func (n *AccountCodeListType) Clone() AccountCodeListType {
  resource := &AccountCodeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_TimeElementsCreate(x LifeCycleType_TimeElements) *LifeCycleType_TimeElements {
        return &x
}

func (n *LifeCycleType_TimeElements) New() *LifeCycleType_TimeElements {
        return LifeCycleType_TimeElementsCreate(LifeCycleType_TimeElements{})
}

func (n *LifeCycleType_TimeElements) Clone() LifeCycleType_TimeElements {
  resource := &LifeCycleType_TimeElements{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SubstituteItemListTypeCreate(x SubstituteItemListType) *SubstituteItemListType {
        return &x
}

func (n *SubstituteItemListType) New() *SubstituteItemListType {
        return SubstituteItemListTypeCreate(SubstituteItemListType{})
}

func (n *SubstituteItemListType) Clone() SubstituteItemListType {
  resource := &SubstituteItemListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GridLocationTypeCreate(x GridLocationType) *GridLocationType {
        return &x
}

func (n *GridLocationType) New() *GridLocationType {
        return GridLocationTypeCreate(GridLocationType{})
}

func (n *GridLocationType) Clone() GridLocationType {
  resource := &GridLocationType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentPackageType_XMLDataCreate(x AbstractContentPackageType_XMLData) *AbstractContentPackageType_XMLData {
        return &x
}

func (n *AbstractContentPackageType_XMLData) New() *AbstractContentPackageType_XMLData {
        return AbstractContentPackageType_XMLDataCreate(AbstractContentPackageType_XMLData{})
}

func (n *AbstractContentPackageType_XMLData) Clone() AbstractContentPackageType_XMLData {
  resource := &AbstractContentPackageType_XMLData{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EntityContactInfoTypeCreate(x EntityContactInfoType) *EntityContactInfoType {
        return &x
}

func (n *EntityContactInfoType) New() *EntityContactInfoType {
        return EntityContactInfoTypeCreate(EntityContactInfoType{})
}

func (n *EntityContactInfoType) Clone() EntityContactInfoType {
  resource := &EntityContactInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGContextualQuestionTypeCreate(x AGContextualQuestionType) *AGContextualQuestionType {
        return &x
}

func (n *AGContextualQuestionType) New() *AGContextualQuestionType {
        return AGContextualQuestionTypeCreate(AGContextualQuestionType{})
}

func (n *AGContextualQuestionType) Clone() AGContextualQuestionType {
  resource := &AGContextualQuestionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FineInfoTypeCreate(x FineInfoType) *FineInfoType {
        return &x
}

func (n *FineInfoType) New() *FineInfoType {
        return FineInfoTypeCreate(FineInfoType{})
}

func (n *FineInfoType) Clone() FineInfoType {
  resource := &FineInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningResourcesTypeCreate(x LearningResourcesType) *LearningResourcesType {
        return &x
}

func (n *LearningResourcesType) New() *LearningResourcesType {
        return LearningResourcesTypeCreate(LearningResourcesType{})
}

func (n *LearningResourcesType) Clone() LearningResourcesType {
  resource := &LearningResourcesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentSubjectChoiceTypeCreate(x StudentSubjectChoiceType) *StudentSubjectChoiceType {
        return &x
}

func (n *StudentSubjectChoiceType) New() *StudentSubjectChoiceType {
        return StudentSubjectChoiceTypeCreate(StudentSubjectChoiceType{})
}

func (n *StudentSubjectChoiceType) Clone() StudentSubjectChoiceType {
  resource := &StudentSubjectChoiceType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolFocusListTypeCreate(x SchoolFocusListType) *SchoolFocusListType {
        return &x
}

func (n *SchoolFocusListType) New() *SchoolFocusListType {
        return SchoolFocusListTypeCreate(SchoolFocusListType{})
}

func (n *SchoolFocusListType) Clone() SchoolFocusListType {
  resource := &SchoolFocusListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func YearRangeTypeCreate(x YearRangeType) *YearRangeType {
        return &x
}

func (n *YearRangeType) New() *YearRangeType {
        return YearRangeTypeCreate(YearRangeType{})
}

func (n *YearRangeType) Clone() YearRangeType {
  resource := &YearRangeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryTransactionTypeCreate(x LibraryTransactionType) *LibraryTransactionType {
        return &x
}

func (n *LibraryTransactionType) New() *LibraryTransactionType {
        return LibraryTransactionTypeCreate(LibraryTransactionType{})
}

func (n *LibraryTransactionType) Clone() LibraryTransactionType {
  resource := &LibraryTransactionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PhoneNumberTypeCreate(x PhoneNumberType) *PhoneNumberType {
        return &x
}

func (n *PhoneNumberType) New() *PhoneNumberType {
        return PhoneNumberTypeCreate(PhoneNumberType{})
}

func (n *PhoneNumberType) Clone() PhoneNumberType {
  resource := &PhoneNumberType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func InvoiceCreate(x Invoice) *Invoice {
        return &x
}

func (n *Invoice) New() *Invoice {
        return InvoiceCreate(Invoice{})
}

func (n *Invoice) Clone() Invoice {
  resource := &Invoice{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentAttendanceCollectionCreate(x StudentAttendanceCollection) *StudentAttendanceCollection {
        return &x
}

func (n *StudentAttendanceCollection) New() *StudentAttendanceCollection {
        return StudentAttendanceCollectionCreate(StudentAttendanceCollection{})
}

func (n *StudentAttendanceCollection) Clone() StudentAttendanceCollection {
  resource := &StudentAttendanceCollection{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FinancialAccountRefIdListTypeCreate(x FinancialAccountRefIdListType) *FinancialAccountRefIdListType {
        return &x
}

func (n *FinancialAccountRefIdListType) New() *FinancialAccountRefIdListType {
        return FinancialAccountRefIdListTypeCreate(FinancialAccountRefIdListType{})
}

func (n *FinancialAccountRefIdListType) Clone() FinancialAccountRefIdListType {
  resource := &FinancialAccountRefIdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableCreate(x TimeTable) *TimeTable {
        return &x
}

func (n *TimeTable) New() *TimeTable {
        return TimeTableCreate(TimeTable{})
}

func (n *TimeTable) Clone() TimeTable {
  resource := &TimeTable{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AlertMessagesTypeCreate(x AlertMessagesType) *AlertMessagesType {
        return &x
}

func (n *AlertMessagesType) New() *AlertMessagesType {
        return AlertMessagesTypeCreate(AlertMessagesType{})
}

func (n *AlertMessagesType) Clone() AlertMessagesType {
  resource := &AlertMessagesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPEventStudentLinkCreate(x NAPEventStudentLink) *NAPEventStudentLink {
        return &x
}

func (n *NAPEventStudentLink) New() *NAPEventStudentLink {
        return NAPEventStudentLinkCreate(NAPEventStudentLink{})
}

func (n *NAPEventStudentLink) Clone() NAPEventStudentLink {
  resource := &NAPEventStudentLink{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FQReportingTypeCreate(x FQReportingType) *FQReportingType {
        return &x
}

func (n *FQReportingType) New() *FQReportingType {
        return FQReportingTypeCreate(FQReportingType{})
}

func (n *FQReportingType) Clone() FQReportingType {
  resource := &FQReportingType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceBookingCreate(x ResourceBooking) *ResourceBooking {
        return &x
}

func (n *ResourceBooking) New() *ResourceBooking {
        return ResourceBookingCreate(ResourceBooking{})
}

func (n *ResourceBooking) Clone() ResourceBooking {
  resource := &ResourceBooking{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func VisaSubClassListTypeCreate(x VisaSubClassListType) *VisaSubClassListType {
        return &x
}

func (n *VisaSubClassListType) New() *VisaSubClassListType {
        return VisaSubClassListTypeCreate(VisaSubClassListType{})
}

func (n *VisaSubClassListType) Clone() VisaSubClassListType {
  resource := &VisaSubClassListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PublishingPermissionTypeCreate(x PublishingPermissionType) *PublishingPermissionType {
        return &x
}

func (n *PublishingPermissionType) New() *PublishingPermissionType {
        return PublishingPermissionTypeCreate(PublishingPermissionType{})
}

func (n *PublishingPermissionType) Clone() PublishingPermissionType {
  resource := &PublishingPermissionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func YearLevelTypeCreate(x YearLevelType) *YearLevelType {
        return &x
}

func (n *YearLevelType) New() *YearLevelType {
        return YearLevelTypeCreate(YearLevelType{})
}

func (n *YearLevelType) Clone() YearLevelType {
  resource := &YearLevelType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletCreate(x NAPTestlet) *NAPTestlet {
        return &x
}

func (n *NAPTestlet) New() *NAPTestlet {
        return NAPTestletCreate(NAPTestlet{})
}

func (n *NAPTestlet) Clone() NAPTestlet {
  resource := &NAPTestlet{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentParticipation_ManagingSchoolCreate(x StudentParticipation_ManagingSchool) *StudentParticipation_ManagingSchool {
        return &x
}

func (n *StudentParticipation_ManagingSchool) New() *StudentParticipation_ManagingSchool {
        return StudentParticipation_ManagingSchoolCreate(StudentParticipation_ManagingSchool{})
}

func (n *StudentParticipation_ManagingSchool) Clone() StudentParticipation_ManagingSchool {
  resource := &StudentParticipation_ManagingSchool{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LocationType_LocationRefIdCreate(x LocationType_LocationRefId) *LocationType_LocationRefId {
        return &x
}

func (n *LocationType_LocationRefId) New() *LocationType_LocationRefId {
        return LocationType_LocationRefIdCreate(LocationType_LocationRefId{})
}

func (n *LocationType_LocationRefId) Clone() LocationType_LocationRefId {
  resource := &LocationType_LocationRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func YearLevelEnrollmentTypeCreate(x YearLevelEnrollmentType) *YearLevelEnrollmentType {
        return &x
}

func (n *YearLevelEnrollmentType) New() *YearLevelEnrollmentType {
        return YearLevelEnrollmentTypeCreate(YearLevelEnrollmentType{})
}

func (n *YearLevelEnrollmentType) Clone() YearLevelEnrollmentType {
  resource := &YearLevelEnrollmentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LocationTypeCreate(x LocationType) *LocationType {
        return &x
}

func (n *LocationType) New() *LocationType {
        return LocationTypeCreate(LocationType{})
}

func (n *LocationType) Clone() LocationType {
  resource := &LocationType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentPackageType_ReferenceCreate(x AbstractContentPackageType_Reference) *AbstractContentPackageType_Reference {
        return &x
}

func (n *AbstractContentPackageType_Reference) New() *AbstractContentPackageType_Reference {
        return AbstractContentPackageType_ReferenceCreate(AbstractContentPackageType_Reference{})
}

func (n *AbstractContentPackageType_Reference) Clone() AbstractContentPackageType_Reference {
  resource := &AbstractContentPackageType_Reference{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func FinancialQuestionnaireCollectionCreate(x FinancialQuestionnaireCollection) *FinancialQuestionnaireCollection {
        return &x
}

func (n *FinancialQuestionnaireCollection) New() *FinancialQuestionnaireCollection {
        return FinancialQuestionnaireCollectionCreate(FinancialQuestionnaireCollection{})
}

func (n *FinancialQuestionnaireCollection) Clone() FinancialQuestionnaireCollection {
  resource := &FinancialQuestionnaireCollection{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CensusStaffTypeCreate(x CensusStaffType) *CensusStaffType {
        return &x
}

func (n *CensusStaffType) New() *CensusStaffType {
        return CensusStaffTypeCreate(CensusStaffType{})
}

func (n *CensusStaffType) Clone() CensusStaffType {
  resource := &CensusStaffType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SubjectAreaListTypeCreate(x SubjectAreaListType) *SubjectAreaListType {
        return &x
}

func (n *SubjectAreaListType) New() *SubjectAreaListType {
        return SubjectAreaListTypeCreate(SubjectAreaListType{})
}

func (n *SubjectAreaListType) Clone() SubjectAreaListType {
  resource := &SubjectAreaListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PasswordListTypeCreate(x PasswordListType) *PasswordListType {
        return &x
}

func (n *PasswordListType) New() *PasswordListType {
        return PasswordListTypeCreate(PasswordListType{})
}

func (n *PasswordListType) Clone() PasswordListType {
  resource := &PasswordListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestContentTypeCreate(x NAPTestContentType) *NAPTestContentType {
        return &x
}

func (n *NAPTestContentType) New() *NAPTestContentType {
        return NAPTestContentTypeCreate(NAPTestContentType{})
}

func (n *NAPTestContentType) Clone() NAPTestContentType {
  resource := &NAPTestContentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AggregateStatisticInfo_CalculationRuleCreate(x AggregateStatisticInfo_CalculationRule) *AggregateStatisticInfo_CalculationRule {
        return &x
}

func (n *AggregateStatisticInfo_CalculationRule) New() *AggregateStatisticInfo_CalculationRule {
        return AggregateStatisticInfo_CalculationRuleCreate(AggregateStatisticInfo_CalculationRule{})
}

func (n *AggregateStatisticInfo_CalculationRule) Clone() AggregateStatisticInfo_CalculationRule {
  resource := &AggregateStatisticInfo_CalculationRule{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentListTypeCreate(x StudentListType) *StudentListType {
        return &x
}

func (n *StudentListType) New() *StudentListType {
        return StudentListTypeCreate(StudentListType{})
}

func (n *StudentListType) Clone() StudentListType {
  resource := &StudentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestItemCreate(x NAPTestItem) *NAPTestItem {
        return &x
}

func (n *NAPTestItem) New() *NAPTestItem {
        return NAPTestItemCreate(NAPTestItem{})
}

func (n *NAPTestItem) Clone() NAPTestItem {
  resource := &NAPTestItem{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MapReferenceTypeCreate(x MapReferenceType) *MapReferenceType {
        return &x
}

func (n *MapReferenceType) New() *MapReferenceType {
        return MapReferenceTypeCreate(MapReferenceType{})
}

func (n *MapReferenceType) Clone() MapReferenceType {
  resource := &MapReferenceType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AssociatedObjectsType_AssociatedObjectCreate(x AssociatedObjectsType_AssociatedObject) *AssociatedObjectsType_AssociatedObject {
        return &x
}

func (n *AssociatedObjectsType_AssociatedObject) New() *AssociatedObjectsType_AssociatedObject {
        return AssociatedObjectsType_AssociatedObjectCreate(AssociatedObjectsType_AssociatedObject{})
}

func (n *AssociatedObjectsType_AssociatedObject) Clone() AssociatedObjectsType_AssociatedObject {
  resource := &AssociatedObjectsType_AssociatedObject{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentGradeMarkersListTypeCreate(x StudentGradeMarkersListType) *StudentGradeMarkersListType {
        return &x
}

func (n *StudentGradeMarkersListType) New() *StudentGradeMarkersListType {
        return StudentGradeMarkersListTypeCreate(StudentGradeMarkersListType{})
}

func (n *StudentGradeMarkersListType) Clone() StudentGradeMarkersListType {
  resource := &StudentGradeMarkersListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ACStrandAreaListTypeCreate(x ACStrandAreaListType) *ACStrandAreaListType {
        return &x
}

func (n *ACStrandAreaListType) New() *ACStrandAreaListType {
        return ACStrandAreaListTypeCreate(ACStrandAreaListType{})
}

func (n *ACStrandAreaListType) Clone() ACStrandAreaListType {
  resource := &ACStrandAreaListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MonetaryAmountTypeCreate(x MonetaryAmountType) *MonetaryAmountType {
        return &x
}

func (n *MonetaryAmountType) New() *MonetaryAmountType {
        return MonetaryAmountTypeCreate(MonetaryAmountType{})
}

func (n *MonetaryAmountType) Clone() MonetaryAmountType {
  resource := &MonetaryAmountType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func JournalAdjustmentListTypeCreate(x JournalAdjustmentListType) *JournalAdjustmentListType {
        return &x
}

func (n *JournalAdjustmentListType) New() *JournalAdjustmentListType {
        return JournalAdjustmentListTypeCreate(JournalAdjustmentListType{})
}

func (n *JournalAdjustmentListType) Clone() JournalAdjustmentListType {
  resource := &JournalAdjustmentListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EquipmentInfo_SIF_RefIdCreate(x EquipmentInfo_SIF_RefId) *EquipmentInfo_SIF_RefId {
        return &x
}

func (n *EquipmentInfo_SIF_RefId) New() *EquipmentInfo_SIF_RefId {
        return EquipmentInfo_SIF_RefIdCreate(EquipmentInfo_SIF_RefId{})
}

func (n *EquipmentInfo_SIF_RefId) Clone() EquipmentInfo_SIF_RefId {
  resource := &EquipmentInfo_SIF_RefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupStudentTypeCreate(x TeachingGroupStudentType) *TeachingGroupStudentType {
        return &x
}

func (n *TeachingGroupStudentType) New() *TeachingGroupStudentType {
        return TeachingGroupStudentTypeCreate(TeachingGroupStudentType{})
}

func (n *TeachingGroupStudentType) Clone() TeachingGroupStudentType {
  resource := &TeachingGroupStudentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SourceObjectsType_SourceObjectCreate(x SourceObjectsType_SourceObject) *SourceObjectsType_SourceObject {
        return &x
}

func (n *SourceObjectsType_SourceObject) New() *SourceObjectsType_SourceObject {
        return SourceObjectsType_SourceObjectCreate(SourceObjectsType_SourceObject{})
}

func (n *SourceObjectsType_SourceObject) Clone() SourceObjectsType_SourceObject {
  resource := &SourceObjectsType_SourceObject{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentActivityTypeCreate(x StudentActivityType) *StudentActivityType {
        return &x
}

func (n *StudentActivityType) New() *StudentActivityType {
        return StudentActivityTypeCreate(StudentActivityType{})
}

func (n *StudentActivityType) Clone() StudentActivityType {
  resource := &StudentActivityType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AssociatedObjectsTypeCreate(x AssociatedObjectsType) *AssociatedObjectsType {
        return &x
}

func (n *AssociatedObjectsType) New() *AssociatedObjectsType {
        return AssociatedObjectsTypeCreate(AssociatedObjectsType{})
}

func (n *AssociatedObjectsType) Clone() AssociatedObjectsType {
  resource := &AssociatedObjectsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ExclusionRulesTypeCreate(x ExclusionRulesType) *ExclusionRulesType {
        return &x
}

func (n *ExclusionRulesType) New() *ExclusionRulesType {
        return ExclusionRulesTypeCreate(ExclusionRulesType{})
}

func (n *ExclusionRulesType) Clone() ExclusionRulesType {
  resource := &ExclusionRulesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func WellbeingPlanTypeCreate(x WellbeingPlanType) *WellbeingPlanType {
        return &x
}

func (n *WellbeingPlanType) New() *WellbeingPlanType {
        return WellbeingPlanTypeCreate(WellbeingPlanType{})
}

func (n *WellbeingPlanType) Clone() WellbeingPlanType {
  resource := &WellbeingPlanType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PurchasingItemsTypeCreate(x PurchasingItemsType) *PurchasingItemsType {
        return &x
}

func (n *PurchasingItemsType) New() *PurchasingItemsType {
        return PurchasingItemsTypeCreate(PurchasingItemsType{})
}

func (n *PurchasingItemsType) Clone() PurchasingItemsType {
  resource := &PurchasingItemsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SectionInfoCreate(x SectionInfo) *SectionInfo {
        return &x
}

func (n *SectionInfo) New() *SectionInfo {
        return SectionInfoCreate(SectionInfo{})
}

func (n *SectionInfo) Clone() SectionInfo {
  resource := &SectionInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SoftwareRequirementListTypeCreate(x SoftwareRequirementListType) *SoftwareRequirementListType {
        return &x
}

func (n *SoftwareRequirementListType) New() *SoftwareRequirementListType {
        return SoftwareRequirementListTypeCreate(SoftwareRequirementListType{})
}

func (n *SoftwareRequirementListType) Clone() SoftwareRequirementListType {
  resource := &SoftwareRequirementListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EducationFilterTypeCreate(x EducationFilterType) *EducationFilterType {
        return &x
}

func (n *EducationFilterType) New() *EducationFilterType {
        return EducationFilterTypeCreate(EducationFilterType{})
}

func (n *EducationFilterType) Clone() EducationFilterType {
  resource := &EducationFilterType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StatsCohortTypeCreate(x StatsCohortType) *StatsCohortType {
        return &x
}

func (n *StatsCohortType) New() *StatsCohortType {
        return StatsCohortTypeCreate(StatsCohortType{})
}

func (n *StatsCohortType) Clone() StatsCohortType {
  resource := &StatsCohortType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MedicationListTypeCreate(x MedicationListType) *MedicationListType {
        return &x
}

func (n *MedicationListType) New() *MedicationListType {
        return MedicationListTypeCreate(MedicationListType{})
}

func (n *MedicationListType) Clone() MedicationListType {
  resource := &MedicationListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentGradeCreate(x StudentGrade) *StudentGrade {
        return &x
}

func (n *StudentGrade) New() *StudentGrade {
        return StudentGradeCreate(StudentGrade{})
}

func (n *StudentGrade) Clone() StudentGrade {
  resource := &StudentGrade{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AGReportingObjectResponseListTypeCreate(x AGReportingObjectResponseListType) *AGReportingObjectResponseListType {
        return &x
}

func (n *AGReportingObjectResponseListType) New() *AGReportingObjectResponseListType {
        return AGReportingObjectResponseListTypeCreate(AGReportingObjectResponseListType{})
}

func (n *AGReportingObjectResponseListType) Clone() AGReportingObjectResponseListType {
  resource := &AGReportingObjectResponseListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_SIF_RefIdCreate(x ResourceUsage_SIF_RefId) *ResourceUsage_SIF_RefId {
        return &x
}

func (n *ResourceUsage_SIF_RefId) New() *ResourceUsage_SIF_RefId {
        return ResourceUsage_SIF_RefIdCreate(ResourceUsage_SIF_RefId{})
}

func (n *ResourceUsage_SIF_RefId) Clone() ResourceUsage_SIF_RefId {
  resource := &ResourceUsage_SIF_RefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func HouseholdContactInfoTypeCreate(x HouseholdContactInfoType) *HouseholdContactInfoType {
        return &x
}

func (n *HouseholdContactInfoType) New() *HouseholdContactInfoType {
        return HouseholdContactInfoTypeCreate(HouseholdContactInfoType{})
}

func (n *HouseholdContactInfoType) Clone() HouseholdContactInfoType {
  resource := &HouseholdContactInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DomainBandsContainerTypeCreate(x DomainBandsContainerType) *DomainBandsContainerType {
        return &x
}

func (n *DomainBandsContainerType) New() *DomainBandsContainerType {
        return DomainBandsContainerTypeCreate(DomainBandsContainerType{})
}

func (n *DomainBandsContainerType) Clone() DomainBandsContainerType {
  resource := &DomainBandsContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolProgramsCreate(x SchoolPrograms) *SchoolPrograms {
        return &x
}

func (n *SchoolPrograms) New() *SchoolPrograms {
        return SchoolProgramsCreate(SchoolPrograms{})
}

func (n *SchoolPrograms) Clone() SchoolPrograms {
  resource := &SchoolPrograms{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CopyRightContainerTypeCreate(x CopyRightContainerType) *CopyRightContainerType {
        return &x
}

func (n *CopyRightContainerType) New() *CopyRightContainerType {
        return CopyRightContainerTypeCreate(CopyRightContainerType{})
}

func (n *CopyRightContainerType) Clone() CopyRightContainerType {
  resource := &CopyRightContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentPersonalCreate(x StudentPersonal) *StudentPersonal {
        return &x
}

func (n *StudentPersonal) New() *StudentPersonal {
        return StudentPersonalCreate(StudentPersonal{})
}

func (n *StudentPersonal) Clone() StudentPersonal {
  resource := &StudentPersonal{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestletItemResponseListTypeCreate(x NAPTestletItemResponseListType) *NAPTestletItemResponseListType {
        return &x
}

func (n *NAPTestletItemResponseListType) New() *NAPTestletItemResponseListType {
        return NAPTestletItemResponseListTypeCreate(NAPTestletItemResponseListType{})
}

func (n *NAPTestletItemResponseListType) Clone() NAPTestletItemResponseListType {
  resource := &NAPTestletItemResponseListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CalendarSummaryCreate(x CalendarSummary) *CalendarSummary {
        return &x
}

func (n *CalendarSummary) New() *CalendarSummary {
        return CalendarSummaryCreate(CalendarSummary{})
}

func (n *CalendarSummary) Clone() CalendarSummary {
  resource := &CalendarSummary{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ApprovalsTypeCreate(x ApprovalsType) *ApprovalsType {
        return &x
}

func (n *ApprovalsType) New() *ApprovalsType {
        return ApprovalsTypeCreate(ApprovalsType{})
}

func (n *ApprovalsType) Clone() ApprovalsType {
  resource := &ApprovalsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func VendorInfoCreate(x VendorInfo) *VendorInfo {
        return &x
}

func (n *VendorInfo) New() *VendorInfo {
        return VendorInfoCreate(VendorInfo{})
}

func (n *VendorInfo) Clone() VendorInfo {
  resource := &VendorInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AlternateIdentificationCodeListTypeCreate(x AlternateIdentificationCodeListType) *AlternateIdentificationCodeListType {
        return &x
}

func (n *AlternateIdentificationCodeListType) New() *AlternateIdentificationCodeListType {
        return AlternateIdentificationCodeListTypeCreate(AlternateIdentificationCodeListType{})
}

func (n *AlternateIdentificationCodeListType) Clone() AlternateIdentificationCodeListType {
  resource := &AlternateIdentificationCodeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func RecognitionListTypeCreate(x RecognitionListType) *RecognitionListType {
        return &x
}

func (n *RecognitionListType) New() *RecognitionListType {
        return RecognitionListTypeCreate(RecognitionListType{})
}

func (n *RecognitionListType) Clone() RecognitionListType {
  resource := &RecognitionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LocalCodeTypeCreate(x LocalCodeType) *LocalCodeType {
        return &x
}

func (n *LocalCodeType) New() *LocalCodeType {
        return LocalCodeTypeCreate(LocalCodeType{})
}

func (n *LocalCodeType) Clone() LocalCodeType {
  resource := &LocalCodeType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableScheduleTypeCreate(x TimeTableScheduleType) *TimeTableScheduleType {
        return &x
}

func (n *TimeTableScheduleType) New() *TimeTableScheduleType {
        return TimeTableScheduleTypeCreate(TimeTableScheduleType{})
}

func (n *TimeTableScheduleType) Clone() TimeTableScheduleType {
  resource := &TimeTableScheduleType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SchoolContactListTypeCreate(x SchoolContactListType) *SchoolContactListType {
        return &x
}

func (n *SchoolContactListType) New() *SchoolContactListType {
        return SchoolContactListTypeCreate(SchoolContactListType{})
}

func (n *SchoolContactListType) Clone() SchoolContactListType {
  resource := &SchoolContactListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentMostRecentContainerTypeCreate(x StudentMostRecentContainerType) *StudentMostRecentContainerType {
        return &x
}

func (n *StudentMostRecentContainerType) New() *StudentMostRecentContainerType {
        return StudentMostRecentContainerTypeCreate(StudentMostRecentContainerType{})
}

func (n *StudentMostRecentContainerType) Clone() StudentMostRecentContainerType {
  resource := &StudentMostRecentContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EquipmentInfoCreate(x EquipmentInfo) *EquipmentInfo {
        return &x
}

func (n *EquipmentInfo) New() *EquipmentInfo {
        return EquipmentInfoCreate(EquipmentInfo{})
}

func (n *EquipmentInfo) Clone() EquipmentInfo {
  resource := &EquipmentInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CatchmentStatusContainerTypeCreate(x CatchmentStatusContainerType) *CatchmentStatusContainerType {
        return &x
}

func (n *CatchmentStatusContainerType) New() *CatchmentStatusContainerType {
        return CatchmentStatusContainerTypeCreate(CatchmentStatusContainerType{})
}

func (n *CatchmentStatusContainerType) Clone() CatchmentStatusContainerType {
  resource := &CatchmentStatusContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PaymentReceiptCreate(x PaymentReceipt) *PaymentReceipt {
        return &x
}

func (n *PaymentReceipt) New() *PaymentReceipt {
        return PaymentReceiptCreate(PaymentReceipt{})
}

func (n *PaymentReceipt) Clone() PaymentReceipt {
  resource := &PaymentReceipt{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StandardsSettingBodyTypeCreate(x StandardsSettingBodyType) *StandardsSettingBodyType {
        return &x
}

func (n *StandardsSettingBodyType) New() *StandardsSettingBodyType {
        return StandardsSettingBodyTypeCreate(StandardsSettingBodyType{})
}

func (n *StandardsSettingBodyType) Clone() StandardsSettingBodyType {
  resource := &StandardsSettingBodyType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SoftwareRequirementTypeCreate(x SoftwareRequirementType) *SoftwareRequirementType {
        return &x
}

func (n *SoftwareRequirementType) New() *SoftwareRequirementType {
        return SoftwareRequirementTypeCreate(SoftwareRequirementType{})
}

func (n *SoftwareRequirementType) Clone() SoftwareRequirementType {
  resource := &SoftwareRequirementType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PlausibleScaledValueListTypeCreate(x PlausibleScaledValueListType) *PlausibleScaledValueListType {
        return &x
}

func (n *PlausibleScaledValueListType) New() *PlausibleScaledValueListType {
        return PlausibleScaledValueListTypeCreate(PlausibleScaledValueListType{})
}

func (n *PlausibleScaledValueListType) Clone() PlausibleScaledValueListType {
  resource := &PlausibleScaledValueListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryPatronStatusCreate(x LibraryPatronStatus) *LibraryPatronStatus {
        return &x
}

func (n *LibraryPatronStatus) New() *LibraryPatronStatus {
        return LibraryPatronStatusCreate(LibraryPatronStatus{})
}

func (n *LibraryPatronStatus) Clone() LibraryPatronStatus {
  resource := &LibraryPatronStatus{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PromotionInfoTypeCreate(x PromotionInfoType) *PromotionInfoType {
        return &x
}

func (n *PromotionInfoType) New() *PromotionInfoType {
        return PromotionInfoTypeCreate(PromotionInfoType{})
}

func (n *PromotionInfoType) Clone() PromotionInfoType {
  resource := &PromotionInfoType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PaymentReceiptLineTypeCreate(x PaymentReceiptLineType) *PaymentReceiptLineType {
        return &x
}

func (n *PaymentReceiptLineType) New() *PaymentReceiptLineType {
        return PaymentReceiptLineTypeCreate(PaymentReceiptLineType{})
}

func (n *PaymentReceiptLineType) Clone() PaymentReceiptLineType {
  resource := &PaymentReceiptLineType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StimulusListTypeCreate(x StimulusListType) *StimulusListType {
        return &x
}

func (n *StimulusListType) New() *StimulusListType {
        return StimulusListTypeCreate(StimulusListType{})
}

func (n *StimulusListType) Clone() StimulusListType {
  resource := &StimulusListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ExclusionRuleTypeCreate(x ExclusionRuleType) *ExclusionRuleType {
        return &x
}

func (n *ExclusionRuleType) New() *ExclusionRuleType {
        return ExclusionRuleTypeCreate(ExclusionRuleType{})
}

func (n *ExclusionRuleType) Clone() ExclusionRuleType {
  resource := &ExclusionRuleType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func NAPTestItemContentTypeCreate(x NAPTestItemContentType) *NAPTestItemContentType {
        return &x
}

func (n *NAPTestItemContentType) New() *NAPTestItemContentType {
        return NAPTestItemContentTypeCreate(NAPTestItemContentType{})
}

func (n *NAPTestItemContentType) Clone() NAPTestItemContentType {
  resource := &NAPTestItemContentType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ScoreDescriptionListTypeCreate(x ScoreDescriptionListType) *ScoreDescriptionListType {
        return &x
}

func (n *ScoreDescriptionListType) New() *ScoreDescriptionListType {
        return ScoreDescriptionListTypeCreate(ScoreDescriptionListType{})
}

func (n *ScoreDescriptionListType) Clone() ScoreDescriptionListType {
  resource := &ScoreDescriptionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MarkValueInfoCreate(x MarkValueInfo) *MarkValueInfo {
        return &x
}

func (n *MarkValueInfo) New() *MarkValueInfo {
        return MarkValueInfoCreate(MarkValueInfo{})
}

func (n *MarkValueInfo) Clone() MarkValueInfo {
  resource := &MarkValueInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LifeCycleType_CreatorsCreate(x LifeCycleType_Creators) *LifeCycleType_Creators {
        return &x
}

func (n *LifeCycleType_Creators) New() *LifeCycleType_Creators {
        return LifeCycleType_CreatorsCreate(LifeCycleType_Creators{})
}

func (n *LifeCycleType_Creators) Clone() LifeCycleType_Creators {
  resource := &LifeCycleType_Creators{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func DebtorCreate(x Debtor) *Debtor {
        return &x
}

func (n *Debtor) New() *Debtor {
        return DebtorCreate(Debtor{})
}

func (n *Debtor) Clone() Debtor {
  resource := &Debtor{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func EssentialMaterialsTypeCreate(x EssentialMaterialsType) *EssentialMaterialsType {
        return &x
}

func (n *EssentialMaterialsType) New() *EssentialMaterialsType {
        return EssentialMaterialsTypeCreate(EssentialMaterialsType{})
}

func (n *EssentialMaterialsType) Clone() EssentialMaterialsType {
  resource := &EssentialMaterialsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsageCreate(x ResourceUsage) *ResourceUsage {
        return &x
}

func (n *ResourceUsage) New() *ResourceUsage {
        return ResourceUsageCreate(ResourceUsage{})
}

func (n *ResourceUsage) Clone() ResourceUsage {
  resource := &ResourceUsage{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func OrganizationsTypeCreate(x OrganizationsType) *OrganizationsType {
        return &x
}

func (n *OrganizationsType) New() *OrganizationsType {
        return OrganizationsTypeCreate(OrganizationsType{})
}

func (n *OrganizationsType) Clone() OrganizationsType {
  resource := &OrganizationsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PlanRequiredContainerTypeCreate(x PlanRequiredContainerType) *PlanRequiredContainerType {
        return &x
}

func (n *PlanRequiredContainerType) New() *PlanRequiredContainerType {
        return PlanRequiredContainerTypeCreate(PlanRequiredContainerType{})
}

func (n *PlanRequiredContainerType) Clone() PlanRequiredContainerType {
  resource := &PlanRequiredContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StandardIdentifierTypeCreate(x StandardIdentifierType) *StandardIdentifierType {
        return &x
}

func (n *StandardIdentifierType) New() *StandardIdentifierType {
        return StandardIdentifierTypeCreate(StandardIdentifierType{})
}

func (n *StandardIdentifierType) Clone() StandardIdentifierType {
  resource := &StandardIdentifierType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func MediaTypesTypeCreate(x MediaTypesType) *MediaTypesType {
        return &x
}

func (n *MediaTypesType) New() *MediaTypesType {
        return MediaTypesTypeCreate(MediaTypesType{})
}

func (n *MediaTypesType) Clone() MediaTypesType {
  resource := &MediaTypesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func HouseholdListTypeCreate(x HouseholdListType) *HouseholdListType {
        return &x
}

func (n *HouseholdListType) New() *HouseholdListType {
        return HouseholdListTypeCreate(HouseholdListType{})
}

func (n *HouseholdListType) Clone() HouseholdListType {
  resource := &HouseholdListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentAttendanceTimeListCreate(x StudentAttendanceTimeList) *StudentAttendanceTimeList {
        return &x
}

func (n *StudentAttendanceTimeList) New() *StudentAttendanceTimeList {
        return StudentAttendanceTimeListCreate(StudentAttendanceTimeList{})
}

func (n *StudentAttendanceTimeList) Clone() StudentAttendanceTimeList {
  resource := &StudentAttendanceTimeList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CodeFrameTestItemTypeCreate(x CodeFrameTestItemType) *CodeFrameTestItemType {
        return &x
}

func (n *CodeFrameTestItemType) New() *CodeFrameTestItemType {
        return CodeFrameTestItemTypeCreate(CodeFrameTestItemType{})
}

func (n *CodeFrameTestItemType) Clone() CodeFrameTestItemType {
  resource := &CodeFrameTestItemType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_ResourceReportColumnCreate(x ResourceUsage_ResourceReportColumn) *ResourceUsage_ResourceReportColumn {
        return &x
}

func (n *ResourceUsage_ResourceReportColumn) New() *ResourceUsage_ResourceReportColumn {
        return ResourceUsage_ResourceReportColumnCreate(ResourceUsage_ResourceReportColumn{})
}

func (n *ResourceUsage_ResourceReportColumn) Clone() ResourceUsage_ResourceReportColumn {
  resource := &ResourceUsage_ResourceReportColumn{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TeachingGroupListTypeCreate(x TeachingGroupListType) *TeachingGroupListType {
        return &x
}

func (n *TeachingGroupListType) New() *TeachingGroupListType {
        return TeachingGroupListTypeCreate(TeachingGroupListType{})
}

func (n *TeachingGroupListType) Clone() TeachingGroupListType {
  resource := &TeachingGroupListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StimulusTypeCreate(x StimulusType) *StimulusType {
        return &x
}

func (n *StimulusType) New() *StimulusType {
        return StimulusTypeCreate(StimulusType{})
}

func (n *StimulusType) Clone() StimulusType {
  resource := &StimulusType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func AbstractContentPackageTypeCreate(x AbstractContentPackageType) *AbstractContentPackageType {
        return &x
}

func (n *AbstractContentPackageType) New() *AbstractContentPackageType {
        return AbstractContentPackageTypeCreate(AbstractContentPackageType{})
}

func (n *AbstractContentPackageType) Clone() AbstractContentPackageType {
  resource := &AbstractContentPackageType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func SystemRole_SystemContextCreate(x SystemRole_SystemContext) *SystemRole_SystemContext {
        return &x
}

func (n *SystemRole_SystemContext) New() *SystemRole_SystemContext {
        return SystemRole_SystemContextCreate(SystemRole_SystemContext{})
}

func (n *SystemRole_SystemContext) Clone() SystemRole_SystemContext {
  resource := &SystemRole_SystemContext{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CollectionStatusCreate(x CollectionStatus) *CollectionStatus {
        return &x
}

func (n *CollectionStatus) New() *CollectionStatus {
        return CollectionStatusCreate(CollectionStatus{})
}

func (n *CollectionStatus) Clone() CollectionStatus {
  resource := &CollectionStatus{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ReferralSourceTypeCreate(x ReferralSourceType) *ReferralSourceType {
        return &x
}

func (n *ReferralSourceType) New() *ReferralSourceType {
        return ReferralSourceTypeCreate(ReferralSourceType{})
}

func (n *ReferralSourceType) Clone() ReferralSourceType {
  resource := &ReferralSourceType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ProgramAvailabilityTypeCreate(x ProgramAvailabilityType) *ProgramAvailabilityType {
        return &x
}

func (n *ProgramAvailabilityType) New() *ProgramAvailabilityType {
        return ProgramAvailabilityTypeCreate(ProgramAvailabilityType{})
}

func (n *ProgramAvailabilityType) Clone() ProgramAvailabilityType {
  resource := &ProgramAvailabilityType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func StudentActivityInfoCreate(x StudentActivityInfo) *StudentActivityInfo {
        return &x
}

func (n *StudentActivityInfo) New() *StudentActivityInfo {
        return StudentActivityInfoCreate(StudentActivityInfo{})
}

func (n *StudentActivityInfo) Clone() StudentActivityInfo {
  resource := &StudentActivityInfo{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func Journal_OriginatingTransactionRefIdCreate(x Journal_OriginatingTransactionRefId) *Journal_OriginatingTransactionRefId {
        return &x
}

func (n *Journal_OriginatingTransactionRefId) New() *Journal_OriginatingTransactionRefId {
        return Journal_OriginatingTransactionRefIdCreate(Journal_OriginatingTransactionRefId{})
}

func (n *Journal_OriginatingTransactionRefId) Clone() Journal_OriginatingTransactionRefId {
  resource := &Journal_OriginatingTransactionRefId{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TechnicalRequirementsTypeCreate(x TechnicalRequirementsType) *TechnicalRequirementsType {
        return &x
}

func (n *TechnicalRequirementsType) New() *TechnicalRequirementsType {
        return TechnicalRequirementsTypeCreate(TechnicalRequirementsType{})
}

func (n *TechnicalRequirementsType) Clone() TechnicalRequirementsType {
  resource := &TechnicalRequirementsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ApprovalTypeCreate(x ApprovalType) *ApprovalType {
        return &x
}

func (n *ApprovalType) New() *ApprovalType {
        return ApprovalTypeCreate(ApprovalType{})
}

func (n *ApprovalType) Clone() ApprovalType {
  resource := &ApprovalType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func PeriodAttendancesTypeCreate(x PeriodAttendancesType) *PeriodAttendancesType {
        return &x
}

func (n *PeriodAttendancesType) New() *PeriodAttendancesType {
        return PeriodAttendancesTypeCreate(PeriodAttendancesType{})
}

func (n *PeriodAttendancesType) Clone() PeriodAttendancesType {
  resource := &PeriodAttendancesType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ReligionTypeCreate(x ReligionType) *ReligionType {
        return &x
}

func (n *ReligionType) New() *ReligionType {
        return ReligionTypeCreate(ReligionType{})
}

func (n *ReligionType) Clone() ReligionType {
  resource := &ReligionType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LibraryTransactionListTypeCreate(x LibraryTransactionListType) *LibraryTransactionListType {
        return &x
}

func (n *LibraryTransactionListType) New() *LibraryTransactionListType {
        return LibraryTransactionListTypeCreate(LibraryTransactionListType{})
}

func (n *LibraryTransactionListType) Clone() LibraryTransactionListType {
  resource := &LibraryTransactionListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LocalCodeListTypeCreate(x LocalCodeListType) *LocalCodeListType {
        return &x
}

func (n *LocalCodeListType) New() *LocalCodeListType {
        return LocalCodeListTypeCreate(LocalCodeListType{})
}

func (n *LocalCodeListType) Clone() LocalCodeListType {
  resource := &LocalCodeListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func VisaSubClassTypeCreate(x VisaSubClassType) *VisaSubClassType {
        return &x
}

func (n *VisaSubClassType) New() *VisaSubClassType {
        return VisaSubClassTypeCreate(VisaSubClassType{})
}

func (n *VisaSubClassType) Clone() VisaSubClassType {
  resource := &VisaSubClassType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func ResourceUsage_ResourceReportColumnListCreate(x ResourceUsage_ResourceReportColumnList) *ResourceUsage_ResourceReportColumnList {
        return &x
}

func (n *ResourceUsage_ResourceReportColumnList) New() *ResourceUsage_ResourceReportColumnList {
        return ResourceUsage_ResourceReportColumnListCreate(ResourceUsage_ResourceReportColumnList{})
}

func (n *ResourceUsage_ResourceReportColumnList) Clone() ResourceUsage_ResourceReportColumnList {
  resource := &ResourceUsage_ResourceReportColumnList{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func GradingScoreListTypeCreate(x GradingScoreListType) *GradingScoreListType {
        return &x
}

func (n *GradingScoreListType) New() *GradingScoreListType {
        return GradingScoreListTypeCreate(GradingScoreListType{})
}

func (n *GradingScoreListType) Clone() GradingScoreListType {
  resource := &GradingScoreListType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CharacteristicsTypeCreate(x CharacteristicsType) *CharacteristicsType {
        return &x
}

func (n *CharacteristicsType) New() *CharacteristicsType {
        return CharacteristicsTypeCreate(CharacteristicsType{})
}

func (n *CharacteristicsType) Clone() CharacteristicsType {
  resource := &CharacteristicsType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func LearningResource_LocationCreate(x LearningResource_Location) *LearningResource_Location {
        return &x
}

func (n *LearningResource_Location) New() *LearningResource_Location {
        return LearningResource_LocationCreate(LearningResource_Location{})
}

func (n *LearningResource_Location) Clone() LearningResource_Location {
  resource := &LearningResource_Location{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func TimeTableSubjectCreate(x TimeTableSubject) *TimeTableSubject {
        return &x
}

func (n *TimeTableSubject) New() *TimeTableSubject {
        return TimeTableSubjectCreate(TimeTableSubject{})
}

func (n *TimeTableSubject) Clone() TimeTableSubject {
  resource := &TimeTableSubject{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

func CampusContainerTypeCreate(x CampusContainerType) *CampusContainerType {
        return &x
}

func (n *CampusContainerType) New() *CampusContainerType {
        return CampusContainerTypeCreate(CampusContainerType{})
}

func (n *CampusContainerType) Clone() CampusContainerType {
  resource := &CampusContainerType{}
  deepcopier.Copy(n).To(resource)
  return *resource
  }

  func (t *LearningResourcesType) Append(value string) *LearningResourcesType {
    
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

      func (t *LearningResourcesType) AppendString(value interface{}) *LearningResourcesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *IdentityAssertionsType) Append(value IdentityAssertionsType_IdentityAssertion) *IdentityAssertionsType {
    
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

  func (t *SchoolFocusListType) Append(value AUCodeSetsSchoolFocusCodeType) *SchoolFocusListType {
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

      func (t *SchoolFocusListType) AppendString(value interface{}) *SchoolFocusListType {
        return t.Append((AUCodeSetsSchoolFocusCodeType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *LearningStandardListType) Append(value LearningStandardType) *LearningStandardListType {
    
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

  func (t *ScheduledTeacherListType) Append(value TeacherCoverType) *ScheduledTeacherListType {
    
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

  func (t *TimeTablePeriodListType) Append(value TimeTablePeriodType) *TimeTablePeriodListType {
    
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

  func (t *SIF_MetadataType_TimeElements) Append(value TimeElementType) *SIF_MetadataType_TimeElements {
    
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

  func (t *RelatedLearningStandardItemRefIdListType) Append(value RelatedLearningStandardItemRefIdType) *RelatedLearningStandardItemRefIdListType {
    
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

  func (t *AuthorsType) Append(value string) *AuthorsType {
    
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

      func (t *AuthorsType) AppendString(value interface{}) *AuthorsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *StudentSubjectChoiceListType) Append(value StudentSubjectChoiceType) *StudentSubjectChoiceListType {
    
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

  func (t *StatsCohortYearLevelListType) Append(value StatsCohortYearLevelType) *StatsCohortYearLevelListType {
    
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

  func (t *FinancialAccountRefIdListType) Append(value string) *FinancialAccountRefIdListType {
    
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

      func (t *FinancialAccountRefIdListType) AppendString(value interface{}) *FinancialAccountRefIdListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *HouseholdContactInfoListType) Append(value HouseholdContactInfoType) *HouseholdContactInfoListType {
    
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

  func (t *AlertMessagesType) Append(value AlertMessageType) *AlertMessagesType {
    
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

  func (t *LearningStandardsType) Append(value string) *LearningStandardsType {
    
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

      func (t *LearningStandardsType) AppendString(value interface{}) *LearningStandardsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *SchoolGroupListType) Append(value LocalIdType) *SchoolGroupListType {
    
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

      func (t *SchoolGroupListType) AppendString(value interface{}) *SchoolGroupListType {
        return t.Append((LocalIdType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *AGContextualQuestionListType) Append(value AGContextualQuestionType) *AGContextualQuestionListType {
    
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

  func (t *StatementCodesType) Append(value string) *StatementCodesType {
    
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

      func (t *StatementCodesType) AppendString(value interface{}) *StatementCodesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *PrerequisitesType) Append(value string) *PrerequisitesType {
    
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

      func (t *PrerequisitesType) AppendString(value interface{}) *PrerequisitesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ReligiousEventListType) Append(value ReligiousEventType) *ReligiousEventListType {
    
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

  func (t *VisaSubClassListType) Append(value VisaSubClassType) *VisaSubClassListType {
    
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

  func (t *AddressListType) Append(value AddressType) *AddressListType {
    
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

  func (t *HoldInfoListType) Append(value HoldInfoType) *HoldInfoListType {
    
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

  func (t *FQContextualQuestionListType) Append(value FQContextualQuestionType) *FQContextualQuestionListType {
    
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

  func (t *NAPSubscoreListType) Append(value NAPSubscoreType) *NAPSubscoreListType {
    
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

  func (t *PaymentReceiptLineListType) Append(value PaymentReceiptLineType) *PaymentReceiptLineListType {
    
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

  func (t *SchoolProgramListType) Append(value SchoolProgramType) *SchoolProgramListType {
    
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

  func (t *PlanRequiredListType) Append(value WellbeingPlanType) *PlanRequiredListType {
    
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

  func (t *StudentsType) Append(value string) *StudentsType {
    
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

      func (t *StudentsType) AppendString(value interface{}) *StudentsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *StaffListType) Append(value string) *StaffListType {
    
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

      func (t *StaffListType) AppendString(value interface{}) *StaffListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *StatsCohortListType) Append(value StatsCohortType) *StatsCohortListType {
    
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

  func (t *LibraryMessageListType) Append(value LibraryMessageType) *LibraryMessageListType {
    
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

  func (t *SymptomListType) Append(value string) *SymptomListType {
    
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

      func (t *SymptomListType) AppendString(value interface{}) *SymptomListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ContentDescriptionListType) Append(value string) *ContentDescriptionListType {
    
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

      func (t *ContentDescriptionListType) AppendString(value interface{}) *ContentDescriptionListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *LifeCycleType_ModificationHistory) Append(value LifeCycleType_Modified) *LifeCycleType_ModificationHistory {
    
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

  func (t *SubjectAreaListType) Append(value SubjectAreaType) *SubjectAreaListType {
    
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

  func (t *StatisticalAreasType) Append(value StatisticalAreaType) *StatisticalAreasType {
    
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

  func (t *PasswordListType) Append(value PasswordListType_Password) *PasswordListType {
    
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

  func (t *FollowUpActionListType) Append(value FollowUpActionType) *FollowUpActionListType {
    
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

  func (t *SourceObjectsType) Append(value SourceObjectsType_SourceObject) *SourceObjectsType {
    
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

  func (t *PhoneNumberListType) Append(value PhoneNumberType) *PhoneNumberListType {
    
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

  func (t *ElectronicIdListType) Append(value ElectronicIdType) *ElectronicIdListType {
    
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

  func (t *CountryListType) Append(value CountryType) *CountryListType {
      present := false
  for _, b := range AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values {
        if b == string(value) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values")
      }

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

      func (t *CountryListType) AppendString(value interface{}) *CountryListType {
        return t.Append((CountryType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ProgramFundingSourcesType) Append(value ProgramFundingSourceType) *ProgramFundingSourcesType {
    
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

  func (t *StudentListType) Append(value TeachingGroupStudentType) *StudentListType {
    
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

  func (t *StimulusLocalIdListType) Append(value LocalIdType) *StimulusLocalIdListType {
    
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

      func (t *StimulusLocalIdListType) AppendString(value interface{}) *StimulusLocalIdListType {
        return t.Append((LocalIdType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *StrategiesType) Append(value string) *StrategiesType {
    
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

      func (t *StrategiesType) AppendString(value interface{}) *StrategiesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *YearLevelEnrollmentListType) Append(value YearLevelEnrollmentType) *YearLevelEnrollmentListType {
    
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

  func (t *PublishingPermissionListType) Append(value PublishingPermissionType) *PublishingPermissionListType {
    
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

  func (t *SystemRole_RoleScopeList) Append(value SystemRole_RoleScope) *SystemRole_RoleScopeList {
    
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

  func (t *MedicalAlertMessagesType) Append(value MedicalAlertMessageType) *MedicalAlertMessagesType {
    
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

  func (t *StudentGradeMarkersListType) Append(value MarkerType) *StudentGradeMarkersListType {
    
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

  func (t *StatementsType) Append(value string) *StatementsType {
    
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

      func (t *StatementsType) AppendString(value interface{}) *StatementsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ACStrandAreaListType) Append(value ACStrandSubjectAreaType) *ACStrandAreaListType {
    
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

  func (t *JournalAdjustmentListType) Append(value JournalAdjustmentType) *JournalAdjustmentListType {
    
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

  func (t *CensusReportingListType) Append(value CensusReportingType) *CensusReportingListType {
    
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

  func (t *SIF_ExtendedElementsType) Append(value SIF_ExtendedElementsType_SIF_ExtendedElement) *SIF_ExtendedElementsType {
    
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

  func (t *FQReportingListType) Append(value FQReportingType) *FQReportingListType {
    
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

  func (t *TimeTableDayListType) Append(value TimeTableDayType) *TimeTableDayListType {
    
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

  func (t *AddressCollectionStudentListType) Append(value AddressCollectionStudentType) *AddressCollectionStudentListType {
    
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

  func (t *OtherCodeListType) Append(value OtherCodeListType_OtherCode) *OtherCodeListType {
    
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

  func (t *CalendarSummaryListType) Append(value string) *CalendarSummaryListType {
    
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

      func (t *CalendarSummaryListType) AppendString(value interface{}) *CalendarSummaryListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *EmailListType) Append(value EmailType) *EmailListType {
    
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

  func (t *AssociatedObjectsType) Append(value AssociatedObjectsType_AssociatedObject) *AssociatedObjectsType {
    
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

  func (t *PNPCodeListType) Append(value AUCodeSetsPNPCodeType) *PNPCodeListType {
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

      func (t *PNPCodeListType) AppendString(value interface{}) *PNPCodeListType {
        return t.Append((AUCodeSetsPNPCodeType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ExclusionRulesType) Append(value ExclusionRuleType) *ExclusionRulesType {
    
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

  func (t *SystemRole_SystemContextList) Append(value SystemRole_SystemContext) *SystemRole_SystemContextList {
    
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

  func (t *OtherIdListType) Append(value OtherIdType) *OtherIdListType {
    
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

  func (t *LearningObjectivesType) Append(value string) *LearningObjectivesType {
    
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

      func (t *LearningObjectivesType) AppendString(value interface{}) *LearningObjectivesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *PurchasingItemsType) Append(value PurchasingItemType) *PurchasingItemsType {
    
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

  func (t *WellbeingDocumentListType) Append(value WellbeingDocumentType) *WellbeingDocumentListType {
    
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

  func (t *SoftwareRequirementListType) Append(value SoftwareRequirementType) *SoftwareRequirementListType {
    
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

  func (t *MedicationListType) Append(value MedicationType) *MedicationListType {
    
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

  func (t *AGReportingObjectResponseListType) Append(value AGReportingObjectResponseType) *AGReportingObjectResponseListType {
    
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

  func (t *WellbeingEventSubCategoryListType) Append(value string) *WellbeingEventSubCategoryListType {
    
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

      func (t *WellbeingEventSubCategoryListType) AppendString(value interface{}) *WellbeingEventSubCategoryListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *AssignmentListType) Append(value string) *AssignmentListType {
    
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

      func (t *AssignmentListType) AppendString(value interface{}) *AssignmentListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *FQItemListType) Append(value FQItemType) *FQItemListType {
    
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

  func (t *CensusStaffListType) Append(value CensusStaffType) *CensusStaffListType {
    
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

  func (t *ScoreListType) Append(value ScoreType) *ScoreListType {
    
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

  func (t *LearningStandardsDocumentType) Append(value string) *LearningStandardsDocumentType {
    
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

      func (t *LearningStandardsDocumentType) AppendString(value interface{}) *LearningStandardsDocumentType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *SystemRole_RoleList) Append(value SystemRole_Role) *SystemRole_RoleList {
    
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

  func (t *NAPTestletItemResponseListType) Append(value NAPTestletResponseItemType) *NAPTestletItemResponseListType {
    
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

  func (t *StudentGroupListType) Append(value StudentGroupType) *StudentGroupListType {
    
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

  func (t *AGRuleListType) Append(value AGRuleType) *AGRuleListType {
    
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

  func (t *ApprovalsType) Append(value ApprovalType) *ApprovalsType {
    
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

  func (t *AlternateIdentificationCodeListType) Append(value string) *AlternateIdentificationCodeListType {
    
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

      func (t *AlternateIdentificationCodeListType) AppendString(value interface{}) *AlternateIdentificationCodeListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *RecognitionListType) Append(value string) *RecognitionListType {
    
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

      func (t *RecognitionListType) AppendString(value interface{}) *RecognitionListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *FineInfoListType) Append(value FineInfoType) *FineInfoListType {
    
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

  func (t *SchoolContactListType) Append(value SchoolContactType) *SchoolContactListType {
    
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

  func (t *AGRoundListType) Append(value AGRoundType) *AGRoundListType {
    
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

  func (t *CensusStudentListType) Append(value CensusStudentType) *CensusStudentListType {
    
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

  func (t *ComponentsType) Append(value ComponentType) *ComponentsType {
    
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

  func (t *PlausibleScaledValueListType) Append(value float64) *PlausibleScaledValueListType {
    
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

  func (t *TimeTableScheduleCellListType) Append(value TimeTableScheduleCellType) *TimeTableScheduleCellListType {
    
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

  func (t *StimulusListType) Append(value StimulusType) *StimulusListType {
    
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

  func (t *OtherNamesType) Append(value OtherNameType) *OtherNamesType {
    
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

  func (t *ScoreDescriptionListType) Append(value ScoreDescriptionType) *ScoreDescriptionListType {
    
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

  func (t *TeachingGroupPeriodListType) Append(value TeachingGroupPeriodType) *TeachingGroupPeriodListType {
    
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

  func (t *LifeCycleType_Creators) Append(value LifeCycleType_Creator) *LifeCycleType_Creators {
    
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

  func (t *PersonInvolvementListType) Append(value PersonInvolvementType) *PersonInvolvementListType {
    
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

  func (t *ValidLetterMarkListType) Append(value ValidLetterMarkType) *ValidLetterMarkListType {
    
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

  func (t *CodeFrameTestItemListType) Append(value CodeFrameTestItemType) *CodeFrameTestItemListType {
    
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

  func (t *EssentialMaterialsType) Append(value string) *EssentialMaterialsType {
    
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

      func (t *EssentialMaterialsType) AppendString(value interface{}) *EssentialMaterialsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *NAPWritingRubricListType) Append(value NAPWritingRubricType) *NAPWritingRubricListType {
    
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

  func (t *NAPCodeFrameTestletListType) Append(value NAPTestletCodeFrameType) *NAPCodeFrameTestletListType {
    
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

  func (t *OrganizationsType) Append(value string) *OrganizationsType {
    
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

      func (t *OrganizationsType) AppendString(value interface{}) *OrganizationsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *NAPTestItemListType) Append(value NAPTestItem2Type) *NAPTestItemListType {
    
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

  func (t *TeachingGroupScheduleListType) Append(value TeachingGroupScheduleType) *TeachingGroupScheduleListType {
    
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

  func (t *TestDisruptionListType) Append(value TestDisruptionType) *TestDisruptionListType {
    
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

  func (t *LanguageListType) Append(value LanguageBaseType) *LanguageListType {
    
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

  func (t *AttendanceTimesType) Append(value AttendanceTimeType) *AttendanceTimesType {
    
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

  func (t *NAPStudentResponseTestletListType) Append(value NAPTestletResponseType) *NAPStudentResponseTestletListType {
    
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

  func (t *StaffSubjectListType) Append(value StaffSubjectType) *StaffSubjectListType {
    
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

  func (t *MediaTypesType) Append(value string) *MediaTypesType {
    
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

      func (t *MediaTypesType) AppendString(value interface{}) *MediaTypesType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ContactsType) Append(value ContactType) *ContactsType {
    
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

  func (t *HouseholdListType) Append(value LocalIdType) *HouseholdListType {
    
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

      func (t *HouseholdListType) AppendString(value interface{}) *HouseholdListType {
        return t.Append((LocalIdType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *StudentAttendanceCollectionReportingListType) Append(value StudentAttendanceCollectionReportingType) *StudentAttendanceCollectionReportingListType {
    
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

  func (t *TeachingGroupListType) Append(value string) *TeachingGroupListType {
    
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

      func (t *TeachingGroupListType) AppendString(value interface{}) *TeachingGroupListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *AddressCollectionReportingListType) Append(value AddressCollectionReportingType) *AddressCollectionReportingListType {
    
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

  func (t *RoomListType) Append(value string) *RoomListType {
    
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

      func (t *RoomListType) AppendString(value interface{}) *RoomListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ResourceUsage_ResourceReportLineList) Append(value ResourceUsage_ResourceReportLine) *ResourceUsage_ResourceReportLineList {
    
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

  func (t *ExpenseAccountsType) Append(value ExpenseAccountType) *ExpenseAccountsType {
    
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

  func (t *WellbeingEventCategoryListType) Append(value WellbeingEventCategoryType) *WellbeingEventCategoryListType {
    
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

  func (t *EvaluationsType) Append(value EvaluationType) *EvaluationsType {
    
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

  func (t *LResourcesType) Append(value ResourcesType) *LResourcesType {
    
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

  func (t *AccountCodeListType) Append(value string) *AccountCodeListType {
    
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

      func (t *AccountCodeListType) AppendString(value interface{}) *AccountCodeListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *LifeCycleType_TimeElements) Append(value TimeElementType) *LifeCycleType_TimeElements {
    
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

  func (t *PeriodAttendancesType) Append(value PeriodAttendanceType) *PeriodAttendancesType {
    
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

  func (t *LibraryTransactionListType) Append(value LibraryTransactionType) *LibraryTransactionListType {
    
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

  func (t *WithdrawalTimeListType) Append(value WithdrawalType) *WithdrawalTimeListType {
    
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

  func (t *YearLevelsType) Append(value YearLevelType) *YearLevelsType {
    
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

  func (t *LocalCodeListType) Append(value LocalCodeType) *LocalCodeListType {
    
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

  func (t *CountryList2Type) Append(value CountryType) *CountryList2Type {
      present := false
  for _, b := range AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values {
        if b == string(value) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values")
      }

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

      func (t *CountryList2Type) AppendString(value interface{}) *CountryList2Type {
        return t.Append((CountryType)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *TeacherListType) Append(value TeachingGroupTeacherType) *TeacherListType {
    
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

  func (t *NAPLANClassListType) Append(value string) *NAPLANClassListType {
    
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

      func (t *NAPLANClassListType) AppendString(value interface{}) *NAPLANClassListType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *ResourceUsage_ResourceReportColumnList) Append(value ResourceUsage_ResourceReportColumn) *ResourceUsage_ResourceReportColumnList {
    
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

  func (t *LEAContactListType) Append(value LEAContactType) *LEAContactListType {
    
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

  func (t *GradingScoreListType) Append(value AssignmentScoreType) *GradingScoreListType {
    
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

  func (t *SubstituteItemListType) Append(value SubstituteItemType) *SubstituteItemListType {
    
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

  func (t *CharacteristicsType) Append(value string) *CharacteristicsType {
    
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

      func (t *CharacteristicsType) AppendString(value interface{}) *CharacteristicsType {
        return t.Append((string)(fmt.Sprint(reflect.ValueOf(value).Elem().Interface())))
        }
  func (t *TimeElementType_SpanGaps) Append(value TimeElementType_SpanGap) *TimeElementType_SpanGaps {
    
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

func (t *AUCodeSetsAGCollectionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPTestDomainType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *MsgIdType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianStandardClassificationOfLanguagesASCLType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsFTPTStatusCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsProgressLevelType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPermissionCategoryCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPictureSourceType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingAlertCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *StateProvinceType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *SelectedContentType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolSectorCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsFFPOSStatusCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsRelationshipToStudentType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingAppealStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *IdRefType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingCharacteristicClassificationType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEntryTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *VersionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAttendanceStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEducationAgencyTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *EducationalLevelType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsStaffStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAGSubmissionStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSourceCodeTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsStaffActivityType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ISO4217CurrencyNamesAndCodeElementsType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolLocationType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *PartialDateType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsMaritalStatusAIHWType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPublicSchoolCatchmentStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *URIOrBinaryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsActivityInvolvementCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ObjectType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsDetentionCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *LocalIdType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSexCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *OperationalStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEmailTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianTimeZoneType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPermanentResidentStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPParticipationCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPTestTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsVisaSubClassType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *DefinedProtocolsType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsOperationalStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsYesOrNoCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ExtendedContentType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *GraduationDateType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *SchoolYearType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAttendanceCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingEventCategoryClassType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolCoEdStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ProjectedGraduationYearType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEducationLevelType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsStudentFamilyProgramTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianCitizenshipStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *BirthDateType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsACStrandType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsDwellingArrangementType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsElectronicIdTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsBirthdateVerificationType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ReportDataObjectType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSystemicStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ServiceNameType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *VisaSubClassCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsIndigenousStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEquipmentTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSets0792IdentificationProcedureType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNameUsageTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsYearLevelCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAssessmentReportingMethodType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *PublishInDirectoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsFederalElectorateType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEmploymentTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsTeacherCoverCreditType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *LearningResourcePackage) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolEnrollmentTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPTestItemTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNonSchoolEducationType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsTelephoneNumberTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *VersionWithWildcardsType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEnglishProficiencyType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *OnTimeGraduationYearType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsExitWithdrawalStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *SchoolURLType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingEventLocationType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingCharacteristicSubCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEventSubCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPersonalisedPlanType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPWritingGenreType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingCharacteristicCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolFocusCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsLanguageTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingEventTimePeriodType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsLearningStandardItemRelationshipTypesType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsReceivingLocationOfInstructionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAddressTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEventCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPJurisdictionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPTestItemMarkingTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ObjectNameType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsImmunisationCertificateStatusType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSuspensionCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsMediumOfInstructionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAssessmentTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsCalendarEventType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsScheduledActivityTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolEducationLevelTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAddressRoleType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsTeacherCoverSupervisionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsStandardAustralianClassificationOfCountriesSACCType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *StateProvinceIdType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPNPCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *HomeroomNumberType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSessionTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolSystemType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *ReportPackageType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsDayValueCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *CountryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsSchoolLevelType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsGroupCategoryCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSets0211ProgramAvailabilityType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsNAPResponseCorrectnessType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsPrePrimaryHoursType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAGContextQuestionType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *GUIDType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsBoardingType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianStandardClassificationOfCulturalAndEthnicGroupsASCCEGType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsResourceUsageContentTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *RefIdType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsExitWithdrawalTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsEnrollmentTimeFrameType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsStateTerritoryCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsWellbeingResponseCategoryType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsProgramFundingSourceCodeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsActivityTypeType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }
func (t *AUCodeSetsAustralianStandardGeographicalClassificationASGCType) String() string {
  return fmt.Sprint(reflect.ValueOf(*t))
  }

func (t *AddressStreetType) CopyString(key string, value interface{}) *AddressStreetType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AddressStreetType) Unset(key string) *AddressStreetType {
        switch key {
  case "ApartmentNumber":
   n.ApartmentNumber = nil
  case "Complex":
   n.Complex = nil
  case "StreetSuffix":
   n.StreetSuffix = nil
  case "ApartmentNumberSuffix":
   n.ApartmentNumberSuffix = nil
  case "StreetPrefix":
   n.StreetPrefix = nil
  case "StreetName":
   n.StreetName = nil
  case "Line2":
   n.Line2 = nil
  case "Line3":
   n.Line3 = nil
  case "Line1":
   n.Line1 = nil
  case "ApartmentType":
   n.ApartmentType = nil
  case "ApartmentNumberPrefix":
   n.ApartmentNumberPrefix = nil
  case "StreetType":
   n.StreetType = nil
  case "StreetNumber":
   n.StreetNumber = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressStreetType")
        }
        return n
}

func (n *AddressStreetType) Set(key string, value interface{}) *AddressStreetType {
        if n == nil {
                n = AddressStreetTypeCreate(AddressStreetType{})
        }
        switch key {
    case "ApartmentNumber":
                      n.ApartmentNumber = StringCreate(value.(string))
    case "Complex":
                      n.Complex = StringCreate(value.(string))
    case "StreetSuffix":
                      n.StreetSuffix = StringCreate(value.(string))
    case "ApartmentNumberSuffix":
                      n.ApartmentNumberSuffix = StringCreate(value.(string))
    case "StreetPrefix":
                      n.StreetPrefix = StringCreate(value.(string))
    case "StreetName":
                      n.StreetName = StringCreate(value.(string))
    case "Line2":
                      n.Line2 = StringCreate(value.(string))
    case "Line3":
                      n.Line3 = StringCreate(value.(string))
    case "Line1":
                      n.Line1 = StringCreate(value.(string))
    case "ApartmentType":
                      n.ApartmentType = StringCreate(value.(string))
    case "ApartmentNumberPrefix":
                      n.ApartmentNumberPrefix = StringCreate(value.(string))
    case "StreetType":
                      n.StreetType = StringCreate(value.(string))
    case "StreetNumber":
                      n.StreetNumber = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressStreetType")
        }
        return n
}


func (t *WellbeingAlert) CopyString(key string, value interface{}) *WellbeingAlert {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingAlert) Unset(key string) *WellbeingAlert {
        switch key {
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "WellbeingAlertCategory":
   n.WellbeingAlertCategory = nil
  case "LocalId":
   n.LocalId = nil
  case "AlertKeyContact":
   n.AlertKeyContact = nil
  case "WellbeingAlertStartDate":
   n.WellbeingAlertStartDate = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "WellbeingAlertDescription":
   n.WellbeingAlertDescription = nil
  case "AlertSeverity":
   n.AlertSeverity = nil
  case "WellbeingAlertEndDate":
   n.WellbeingAlertEndDate = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "AlertAudience":
   n.AlertAudience = nil
  case "Date":
   n.Date = nil
  case "EnrolmentRestricted":
   n.EnrolmentRestricted = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingAlert")
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
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AlertKeyContact":
                      n.AlertKeyContact = StringCreate(value.(string))
    case "WellbeingAlertStartDate":
                      n.WellbeingAlertStartDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "WellbeingAlertDescription":
                      n.WellbeingAlertDescription = StringCreate(value.(string))
    case "AlertSeverity":
                      n.AlertSeverity = StringCreate(value.(string))
    case "WellbeingAlertEndDate":
                      n.WellbeingAlertEndDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "AlertAudience":
                      n.AlertAudience = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
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
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingAlert")
        }
        return n
}


func (t *LocationOfInstructionType) CopyString(key string, value interface{}) *LocationOfInstructionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LocationOfInstructionType) Unset(key string) *LocationOfInstructionType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationOfInstructionType")
        }
        return n
}

func (n *LocationOfInstructionType) Set(key string, value interface{}) *LocationOfInstructionType {
        if n == nil {
                n = LocationOfInstructionTypeCreate(LocationOfInstructionType{})
        }
        switch key {
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
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationOfInstructionType")
        }
        return n
}


func (t *StudentDailyAttendance) CopyString(key string, value interface{}) *StudentDailyAttendance {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentDailyAttendance) Unset(key string) *StudentDailyAttendance {
        switch key {
  case "AttendanceCode":
   n.AttendanceCode = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AbsenceValue":
   n.AbsenceValue = nil
  case "TimeOut":
   n.TimeOut = nil
  case "RefId":
   n.RefId = nil
  case "TimeIn":
   n.TimeIn = nil
  case "AttendanceStatus":
   n.AttendanceStatus = nil
  case "Date":
   n.Date = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "DayValue":
   n.DayValue = nil
  case "AttendanceNote":
   n.AttendanceNote = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentDailyAttendance")
        }
        return n
}

func (n *StudentDailyAttendance) Set(key string, value interface{}) *StudentDailyAttendance {
        if n == nil {
                n = StudentDailyAttendanceCreate(StudentDailyAttendance{})
        }
        switch key {
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AbsenceValue":
                      n.AbsenceValue = FloatCreate(value.(float64))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
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
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentDailyAttendance")
        }
        return n
}


func (t *WellbeingEventCategoryType) CopyString(key string, value interface{}) *WellbeingEventCategoryType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingEventCategoryType) Unset(key string) *WellbeingEventCategoryType {
        switch key {
  case "WellbeingEventSubCategoryList":
   n.WellbeingEventSubCategoryList = nil
  case "EventCategory":
   n.EventCategory = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEventCategoryType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEventCategoryType")
        }
        return n
}


func (t *NAPTestScoreSummary) CopyString(key string, value interface{}) *NAPTestScoreSummary {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestScoreSummary) Unset(key string) *NAPTestScoreSummary {
        switch key {
  case "SchoolACARAId":
   n.SchoolACARAId = nil
  case "DomainSchoolAverage":
   n.DomainSchoolAverage = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "DomainBottomNational60Percent":
   n.DomainBottomNational60Percent = nil
  case "DomainNationalAverage":
   n.DomainNationalAverage = nil
  case "DomainJurisdictionAverage":
   n.DomainJurisdictionAverage = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "NAPTestLocalId":
   n.NAPTestLocalId = nil
  case "NAPTestRefId":
   n.NAPTestRefId = nil
  case "DomainTopNational60Percent":
   n.DomainTopNational60Percent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestScoreSummary")
        }
        return n
}

func (n *NAPTestScoreSummary) Set(key string, value interface{}) *NAPTestScoreSummary {
        if n == nil {
                n = NAPTestScoreSummaryCreate(NAPTestScoreSummary{})
        }
        switch key {
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DomainSchoolAverage":
                      n.DomainSchoolAverage = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "DomainBottomNational60Percent":
                      n.DomainBottomNational60Percent = FloatCreate(value.(float64))
    case "DomainNationalAverage":
                      n.DomainNationalAverage = FloatCreate(value.(float64))
    case "DomainJurisdictionAverage":
                      n.DomainJurisdictionAverage = FloatCreate(value.(float64))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "DomainTopNational60Percent":
                      n.DomainTopNational60Percent = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestScoreSummary")
        }
        return n
}


func (t *AttendanceTimeType) CopyString(key string, value interface{}) *AttendanceTimeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AttendanceTimeType) Unset(key string) *AttendanceTimeType {
        switch key {
  case "AttendanceNote":
   n.AttendanceNote = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "AttendanceCode":
   n.AttendanceCode = nil
  case "StartTime":
   n.StartTime = nil
  case "AttendanceStatus":
   n.AttendanceStatus = nil
  case "AttendanceType":
   n.AttendanceType = nil
  case "EndTime":
   n.EndTime = nil
  case "DurationValue":
   n.DurationValue = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceTimeType")
        }
        return n
}

func (n *AttendanceTimeType) Set(key string, value interface{}) *AttendanceTimeType {
        if n == nil {
                n = AttendanceTimeTypeCreate(AttendanceTimeType{})
        }
        switch key {
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
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
    case "DurationValue":
                      n.DurationValue = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceTimeType")
        }
        return n
}


func (t *RelationshipType) CopyString(key string, value interface{}) *RelationshipType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *RelationshipType) Unset(key string) *RelationshipType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RelationshipType")
        }
        return n
}

func (n *RelationshipType) Set(key string, value interface{}) *RelationshipType {
        if n == nil {
                n = RelationshipTypeCreate(RelationshipType{})
        }
        switch key {
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
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RelationshipType")
        }
        return n
}


func (t *ActivityTimeType_Duration) CopyString(key string, value interface{}) *ActivityTimeType_Duration {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ActivityTimeType_Duration) Unset(key string) *ActivityTimeType_Duration {
        switch key {
  case "Units":
   n.Units = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ActivityTimeType_Duration")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ActivityTimeType_Duration")
        }
        return n
}


func (t *ScoreDescriptionType) CopyString(key string, value interface{}) *ScoreDescriptionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ScoreDescriptionType) Unset(key string) *ScoreDescriptionType {
        switch key {
  case "ScoreValue":
   n.ScoreValue = nil
  case "Descriptor":
   n.Descriptor = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScoreDescriptionType")
        }
        return n
}

func (n *ScoreDescriptionType) Set(key string, value interface{}) *ScoreDescriptionType {
        if n == nil {
                n = ScoreDescriptionTypeCreate(ScoreDescriptionType{})
        }
        switch key {
    case "ScoreValue":
                      n.ScoreValue = FloatCreate(value.(float64))
    case "Descriptor":
                      n.Descriptor = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScoreDescriptionType")
        }
        return n
}


func (t *StaffActivityExtensionType) CopyString(key string, value interface{}) *StaffActivityExtensionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffActivityExtensionType) Unset(key string) *StaffActivityExtensionType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffActivityExtensionType")
        }
        return n
}

func (n *StaffActivityExtensionType) Set(key string, value interface{}) *StaffActivityExtensionType {
        if n == nil {
                n = StaffActivityExtensionTypeCreate(StaffActivityExtensionType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffActivityExtensionType")
        }
        return n
}


func (t *StudentExitStatusContainerType) CopyString(key string, value interface{}) *StudentExitStatusContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentExitStatusContainerType) Unset(key string) *StudentExitStatusContainerType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentExitStatusContainerType")
        }
        return n
}

func (n *StudentExitStatusContainerType) Set(key string, value interface{}) *StudentExitStatusContainerType {
        if n == nil {
                n = StudentExitStatusContainerTypeCreate(StudentExitStatusContainerType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentExitStatusContainerType")
        }
        return n
}


func (t *SuspensionContainerType) CopyString(key string, value interface{}) *SuspensionContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SuspensionContainerType) Unset(key string) *SuspensionContainerType {
        switch key {
  case "Duration":
   n.Duration = nil
  case "Status":
   n.Status = nil
  case "EarlyReturnDate":
   n.EarlyReturnDate = nil
  case "WithdrawalTimeList":
   n.WithdrawalTimeList = nil
  case "SuspensionNotes":
   n.SuspensionNotes = nil
  case "ResolutionMeetingTime":
   n.ResolutionMeetingTime = nil
  case "SuspensionCategory":
   n.SuspensionCategory = nil
  case "AdvisementDate":
   n.AdvisementDate = nil
  case "ResolutionNotes":
   n.ResolutionNotes = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SuspensionContainerType")
        }
        return n
}

func (n *SuspensionContainerType) Set(key string, value interface{}) *SuspensionContainerType {
        if n == nil {
                n = SuspensionContainerTypeCreate(SuspensionContainerType{})
        }
        switch key {
    case "Duration":
                      n.Duration = FloatCreate(value.(float64))
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
    case "EarlyReturnDate":
                      n.EarlyReturnDate = StringCreate(value.(string))
    case "WithdrawalTimeList":
                      n.WithdrawalTimeList = WithdrawalTimeListTypeCreate(value.(WithdrawalTimeListType))
    case "SuspensionNotes":
                      n.SuspensionNotes = StringCreate(value.(string))
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
    case "AdvisementDate":
                      n.AdvisementDate = StringCreate(value.(string))
    case "ResolutionNotes":
                      n.ResolutionNotes = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SuspensionContainerType")
        }
        return n
}


func (t *LearningStandardDocument) CopyString(key string, value interface{}) *LearningStandardDocument {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LearningStandardDocument) Unset(key string) *LearningStandardDocument {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Source":
   n.Source = nil
  case "LocalAdoptionDate":
   n.LocalAdoptionDate = nil
  case "LearningStandardItemRefId":
   n.LearningStandardItemRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RichDescription":
   n.RichDescription = nil
  case "RefId":
   n.RefId = nil
  case "SubjectAreas":
   n.SubjectAreas = nil
  case "RepositoryDate":
   n.RepositoryDate = nil
  case "YearLevels":
   n.YearLevels = nil
  case "Authors":
   n.Authors = nil
  case "RelatedLearningStandards":
   n.RelatedLearningStandards = nil
  case "DocumentStatus":
   n.DocumentStatus = nil
  case "Description":
   n.Description = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Copyright":
   n.Copyright = nil
  case "Title":
   n.Title = nil
  case "EndOfLifeDate":
   n.EndOfLifeDate = nil
  case "DocumentDate":
   n.DocumentDate = nil
  case "LocalArchiveDate":
   n.LocalArchiveDate = nil
  case "OrganizationContactPoint":
   n.OrganizationContactPoint = nil
  case "Organizations":
   n.Organizations = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardDocument")
        }
        return n
}

func (n *LearningStandardDocument) Set(key string, value interface{}) *LearningStandardDocument {
        if n == nil {
                n = LearningStandardDocumentCreate(LearningStandardDocument{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Source":
                      n.Source = StringCreate(value.(string))
    case "LocalAdoptionDate":
                      n.LocalAdoptionDate = StringCreate(value.(string))
    case "LearningStandardItemRefId":
                      n.LearningStandardItemRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RichDescription":
                      n.RichDescription = AbstractContentElementTypeCreate(value.(AbstractContentElementType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SubjectAreas":
                      n.SubjectAreas = ACStrandAreaListTypeCreate(value.(ACStrandAreaListType))
    case "RepositoryDate":
                      n.RepositoryDate = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "Authors":
                      n.Authors = AuthorsTypeCreate(value.(AuthorsType))
    case "RelatedLearningStandards":
                      n.RelatedLearningStandards = LearningStandardsDocumentTypeCreate(value.(LearningStandardsDocumentType))
    case "DocumentStatus":
                      n.DocumentStatus = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Copyright":
                      n.Copyright = CopyRightContainerTypeCreate(value.(CopyRightContainerType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "EndOfLifeDate":
                      n.EndOfLifeDate = StringCreate(value.(string))
    case "DocumentDate":
                      n.DocumentDate = StringCreate(value.(string))
    case "LocalArchiveDate":
                      n.LocalArchiveDate = StringCreate(value.(string))
    case "OrganizationContactPoint":
                      n.OrganizationContactPoint = StringCreate(value.(string))
    case "Organizations":
                      n.Organizations = OrganizationsTypeCreate(value.(OrganizationsType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardDocument")
        }
        return n
}


func (t *SoftwareVendorInfoContainerType) CopyString(key string, value interface{}) *SoftwareVendorInfoContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SoftwareVendorInfoContainerType) Unset(key string) *SoftwareVendorInfoContainerType {
        switch key {
  case "SoftwareProduct":
   n.SoftwareProduct = nil
  case "SoftwareVersion":
   n.SoftwareVersion = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SoftwareVendorInfoContainerType")
        }
        return n
}

func (n *SoftwareVendorInfoContainerType) Set(key string, value interface{}) *SoftwareVendorInfoContainerType {
        if n == nil {
                n = SoftwareVendorInfoContainerTypeCreate(SoftwareVendorInfoContainerType{})
        }
        switch key {
    case "SoftwareProduct":
                      n.SoftwareProduct = StringCreate(value.(string))
    case "SoftwareVersion":
                      n.SoftwareVersion = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SoftwareVendorInfoContainerType")
        }
        return n
}


func (t *LibraryMessageType) CopyString(key string, value interface{}) *LibraryMessageType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LibraryMessageType) Unset(key string) *LibraryMessageType {
        switch key {
  case "Priority":
   n.Priority = nil
  case "Text":
   n.Text = nil
  case "Sent":
   n.Sent = nil
  case "PriorityCodeset":
   n.PriorityCodeset = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryMessageType")
        }
        return n
}

func (n *LibraryMessageType) Set(key string, value interface{}) *LibraryMessageType {
        if n == nil {
                n = LibraryMessageTypeCreate(LibraryMessageType{})
        }
        switch key {
    case "Priority":
                      n.Priority = StringCreate(value.(string))
    case "Text":
                      n.Text = StringCreate(value.(string))
    case "Sent":
                      n.Sent = StringCreate(value.(string))
    case "PriorityCodeset":
                      n.PriorityCodeset = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryMessageType")
        }
        return n
}


func (t *SystemRole_SIF_RefId) CopyString(key string, value interface{}) *SystemRole_SIF_RefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole_SIF_RefId) Unset(key string) *SystemRole_SIF_RefId {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_SIF_RefId")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_SIF_RefId")
        }
        return n
}


func (t *AlertMessageType) CopyString(key string, value interface{}) *AlertMessageType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AlertMessageType) Unset(key string) *AlertMessageType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AlertMessageType")
        }
        return n
}

func (n *AlertMessageType) Set(key string, value interface{}) *AlertMessageType {
        if n == nil {
                n = AlertMessageTypeCreate(AlertMessageType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AlertMessageType")
        }
        return n
}


func (t *AddressCollectionStudentType) CopyString(key string, value interface{}) *AddressCollectionStudentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AddressCollectionStudentType) Unset(key string) *AddressCollectionStudentType {
        switch key {
  case "BoardingStatus":
   n.BoardingStatus = nil
  case "Parent2":
   n.Parent2 = nil
  case "EducationLevel":
   n.EducationLevel = nil
  case "StudentAddress":
   n.StudentAddress = nil
  case "LocalId":
   n.LocalId = nil
  case "ReportingParent2":
   n.ReportingParent2 = nil
  case "Parent1":
   n.Parent1 = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollectionStudentType")
        }
        return n
}

func (n *AddressCollectionStudentType) Set(key string, value interface{}) *AddressCollectionStudentType {
        if n == nil {
                n = AddressCollectionStudentTypeCreate(AddressCollectionStudentType{})
        }
        switch key {
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
    case "Parent2":
                      n.Parent2 = AGParentTypeCreate(value.(AGParentType))
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
    case "StudentAddress":
                      n.StudentAddress = AddressTypeCreate(value.(AddressType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ReportingParent2":
                      n.ReportingParent2 = StringCreate(value.(string))
    case "Parent1":
                      n.Parent1 = AGParentTypeCreate(value.(AGParentType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollectionStudentType")
        }
        return n
}


func (t *LifeCycleType_Created) CopyString(key string, value interface{}) *LifeCycleType_Created {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LifeCycleType_Created) Unset(key string) *LifeCycleType_Created {
        switch key {
  case "DateTime":
   n.DateTime = nil
  case "Creators":
   n.Creators = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Created")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Created")
        }
        return n
}


func (t *MedicalAlertMessageType) CopyString(key string, value interface{}) *MedicalAlertMessageType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MedicalAlertMessageType) Unset(key string) *MedicalAlertMessageType {
        switch key {
  case "Severity":
   n.Severity = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MedicalAlertMessageType")
        }
        return n
}

func (n *MedicalAlertMessageType) Set(key string, value interface{}) *MedicalAlertMessageType {
        if n == nil {
                n = MedicalAlertMessageTypeCreate(MedicalAlertMessageType{})
        }
        switch key {
    case "Severity":
                      n.Severity = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MedicalAlertMessageType")
        }
        return n
}


func (t *ComponentType) CopyString(key string, value interface{}) *ComponentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ComponentType) Unset(key string) *ComponentType {
        switch key {
  case "Description":
   n.Description = nil
  case "Reference":
   n.Reference = nil
  case "AssociatedObjects":
   n.AssociatedObjects = nil
  case "Strategies":
   n.Strategies = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ComponentType")
        }
        return n
}

func (n *ComponentType) Set(key string, value interface{}) *ComponentType {
        if n == nil {
                n = ComponentTypeCreate(ComponentType{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Reference":
                      n.Reference = StringCreate(value.(string))
    case "AssociatedObjects":
                      n.AssociatedObjects = AssociatedObjectsTypeCreate(value.(AssociatedObjectsType))
    case "Strategies":
                      n.Strategies = StrategiesTypeCreate(value.(StrategiesType))
    case "Name":
                      n.Name = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ComponentType")
        }
        return n
}


func (t *DomainProficiencyContainerType) CopyString(key string, value interface{}) *DomainProficiencyContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DomainProficiencyContainerType) Unset(key string) *DomainProficiencyContainerType {
        switch key {
  case "Level3Lower":
   n.Level3Lower = nil
  case "Level4Lower":
   n.Level4Lower = nil
  case "Level1Lower":
   n.Level1Lower = nil
  case "Level2Lower":
   n.Level2Lower = nil
  case "Level4Upper":
   n.Level4Upper = nil
  case "Level3Upper":
   n.Level3Upper = nil
  case "Level2Upper":
   n.Level2Upper = nil
  case "Level1Upper":
   n.Level1Upper = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainProficiencyContainerType")
        }
        return n
}

func (n *DomainProficiencyContainerType) Set(key string, value interface{}) *DomainProficiencyContainerType {
        if n == nil {
                n = DomainProficiencyContainerTypeCreate(DomainProficiencyContainerType{})
        }
        switch key {
    case "Level3Lower":
                      n.Level3Lower = FloatCreate(value.(float64))
    case "Level4Lower":
                      n.Level4Lower = FloatCreate(value.(float64))
    case "Level1Lower":
                      n.Level1Lower = FloatCreate(value.(float64))
    case "Level2Lower":
                      n.Level2Lower = FloatCreate(value.(float64))
    case "Level4Upper":
                      n.Level4Upper = FloatCreate(value.(float64))
    case "Level3Upper":
                      n.Level3Upper = FloatCreate(value.(float64))
    case "Level2Upper":
                      n.Level2Upper = FloatCreate(value.(float64))
    case "Level1Upper":
                      n.Level1Upper = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainProficiencyContainerType")
        }
        return n
}


func (t *OtherIdType) CopyString(key string, value interface{}) *OtherIdType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *OtherIdType) Unset(key string) *OtherIdType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherIdType")
        }
        return n
}

func (n *OtherIdType) Set(key string, value interface{}) *OtherIdType {
        if n == nil {
                n = OtherIdTypeCreate(OtherIdType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherIdType")
        }
        return n
}


func (t *PeriodAttendanceType) CopyString(key string, value interface{}) *PeriodAttendanceType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PeriodAttendanceType) Unset(key string) *PeriodAttendanceType {
        switch key {
  case "SessionInfoRefId":
   n.SessionInfoRefId = nil
  case "StartTime":
   n.StartTime = nil
  case "TimetablePeriod":
   n.TimetablePeriod = nil
  case "TimeIn":
   n.TimeIn = nil
  case "AttendanceStatus":
   n.AttendanceStatus = nil
  case "Date":
   n.Date = nil
  case "EndTime":
   n.EndTime = nil
  case "TimeOut":
   n.TimeOut = nil
  case "DayId":
   n.DayId = nil
  case "ScheduledActivityRefId":
   n.ScheduledActivityRefId = nil
  case "AttendanceCode":
   n.AttendanceCode = nil
  case "TimeTableCellRefId":
   n.TimeTableCellRefId = nil
  case "AttendanceType":
   n.AttendanceType = nil
  case "TeacherList":
   n.TeacherList = nil
  case "RoomList":
   n.RoomList = nil
  case "AttendanceNote":
   n.AttendanceNote = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PeriodAttendanceType")
        }
        return n
}

func (n *PeriodAttendanceType) Set(key string, value interface{}) *PeriodAttendanceType {
        if n == nil {
                n = PeriodAttendanceTypeCreate(PeriodAttendanceType{})
        }
        switch key {
    case "SessionInfoRefId":
                      n.SessionInfoRefId = StringCreate(value.(string))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "TimetablePeriod":
                      n.TimetablePeriod = StringCreate(value.(string))
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
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
    case "EndTime":
                      n.EndTime = StringCreate(value.(string))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "AttendanceType":
                      n.AttendanceType = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "AttendanceNote":
                      n.AttendanceNote = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PeriodAttendanceType")
        }
        return n
}


func (t *OtherNameType) CopyString(key string, value interface{}) *OtherNameType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *OtherNameType) Unset(key string) *OtherNameType {
        switch key {
  case "Type":
   n.Type = nil
  case "Suffix":
   n.Suffix = nil
  case "MiddleName":
   n.MiddleName = nil
  case "GivenName":
   n.GivenName = nil
  case "PreferredGivenName":
   n.PreferredGivenName = nil
  case "FamilyNameFirst":
   n.FamilyNameFirst = nil
  case "FullName":
   n.FullName = nil
  case "FamilyName":
   n.FamilyName = nil
  case "Title":
   n.Title = nil
  case "PreferredFamilyNameFirst":
   n.PreferredFamilyNameFirst = nil
  case "PreferredFamilyName":
   n.PreferredFamilyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherNameType")
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
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
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
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Title":
                      n.Title = StringCreate(value.(string))
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
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherNameType")
        }
        return n
}


func (t *StaffAssignmentMostRecentContainerType) CopyString(key string, value interface{}) *StaffAssignmentMostRecentContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffAssignmentMostRecentContainerType) Unset(key string) *StaffAssignmentMostRecentContainerType {
        switch key {
  case "PrimaryFTE":
   n.PrimaryFTE = nil
  case "SecondaryFTE":
   n.SecondaryFTE = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffAssignmentMostRecentContainerType")
        }
        return n
}

func (n *StaffAssignmentMostRecentContainerType) Set(key string, value interface{}) *StaffAssignmentMostRecentContainerType {
        if n == nil {
                n = StaffAssignmentMostRecentContainerTypeCreate(StaffAssignmentMostRecentContainerType{})
        }
        switch key {
    case "PrimaryFTE":
                      n.PrimaryFTE = FloatCreate(value.(float64))
    case "SecondaryFTE":
                      n.SecondaryFTE = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffAssignmentMostRecentContainerType")
        }
        return n
}


func (t *SessionInfo) CopyString(key string, value interface{}) *SessionInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SessionInfo) Unset(key string) *SessionInfo {
        switch key {
  case "FinishTime":
   n.FinishTime = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TeachingGroupLocalId":
   n.TeachingGroupLocalId = nil
  case "RollMarked":
   n.RollMarked = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "DayId":
   n.DayId = nil
  case "RefId":
   n.RefId = nil
  case "TimeTableSubjectLocalId":
   n.TimeTableSubjectLocalId = nil
  case "StartTime":
   n.StartTime = nil
  case "StaffPersonalLocalId":
   n.StaffPersonalLocalId = nil
  case "LocalId":
   n.LocalId = nil
  case "PeriodId":
   n.PeriodId = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SessionDate":
   n.SessionDate = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "RoomNumber":
   n.RoomNumber = nil
  case "TimeTableCellRefId":
   n.TimeTableCellRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SessionInfo")
        }
        return n
}

func (n *SessionInfo) Set(key string, value interface{}) *SessionInfo {
        if n == nil {
                n = SessionInfoCreate(SessionInfo{})
        }
        switch key {
    case "FinishTime":
                      n.FinishTime = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "StaffPersonalLocalId":
    
                      n.StaffPersonalLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SessionDate":
                      n.SessionDate = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
                      n.RoomNumber = StringCreate(value.(string))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SessionInfo")
        }
        return n
}


func (t *StudentSchoolEnrollment_Calendar) CopyString(key string, value interface{}) *StudentSchoolEnrollment_Calendar {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSchoolEnrollment_Calendar) Unset(key string) *StudentSchoolEnrollment_Calendar {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Calendar")
        }
        return n
}

func (n *StudentSchoolEnrollment_Calendar) Set(key string, value interface{}) *StudentSchoolEnrollment_Calendar {
        if n == nil {
                n = StudentSchoolEnrollment_CalendarCreate(StudentSchoolEnrollment_Calendar{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Calendar")
        }
        return n
}


func (t *CensusStudentType) CopyString(key string, value interface{}) *CensusStudentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CensusStudentType) Unset(key string) *CensusStudentType {
        switch key {
  case "FTE":
   n.FTE = nil
  case "DisabilityLevelOfAdjustment":
   n.DisabilityLevelOfAdjustment = nil
  case "CohortIndigenousType":
   n.CohortIndigenousType = nil
  case "YearLevel":
   n.YearLevel = nil
  case "StudentCohortId":
   n.StudentCohortId = nil
  case "Headcount":
   n.Headcount = nil
  case "CensusAge":
   n.CensusAge = nil
  case "CohortGender":
   n.CohortGender = nil
  case "EducationMode":
   n.EducationMode = nil
  case "StudentOnVisa":
   n.StudentOnVisa = nil
  case "DisabilityCategory":
   n.DisabilityCategory = nil
  case "OverseasStudent":
   n.OverseasStudent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusStudentType")
        }
        return n
}

func (n *CensusStudentType) Set(key string, value interface{}) *CensusStudentType {
        if n == nil {
                n = CensusStudentTypeCreate(CensusStudentType{})
        }
        switch key {
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "CohortIndigenousType":
                      n.CohortIndigenousType = StringCreate(value.(string))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "StudentCohortId":
    
                      n.StudentCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Headcount":
                      n.Headcount = IntCreate(value.(int))
    case "CensusAge":
                      n.CensusAge = IntCreate(value.(int))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "EducationMode":
                      n.EducationMode = StringCreate(value.(string))
    case "StudentOnVisa":
                      n.StudentOnVisa = StringCreate(value.(string))
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
    case "OverseasStudent":
                      n.OverseasStudent = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusStudentType")
        }
        return n
}


func (t *TeachingGroup) CopyString(key string, value interface{}) *TeachingGroup {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeachingGroup) Unset(key string) *TeachingGroup {
        switch key {
  case "TimeTableSubjectLocalId":
   n.TimeTableSubjectLocalId = nil
  case "StudentList":
   n.StudentList = nil
  case "Semester":
   n.Semester = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "Sett":
   n.Sett = nil
  case "ShortName":
   n.ShortName = nil
  case "TeachingGroupPeriodList":
   n.TeachingGroupPeriodList = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "LongName":
   n.LongName = nil
  case "Block":
   n.Block = nil
  case "RefId":
   n.RefId = nil
  case "GroupType":
   n.GroupType = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "KeyLearningArea":
   n.KeyLearningArea = nil
  case "SchoolCourseLocalId":
   n.SchoolCourseLocalId = nil
  case "MaxClassSize":
   n.MaxClassSize = nil
  case "SchoolCourseInfoRefId":
   n.SchoolCourseInfoRefId = nil
  case "MinClassSize":
   n.MinClassSize = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "TeacherList":
   n.TeacherList = nil
  case "CurriculumLevel":
   n.CurriculumLevel = nil
  case "LocalId":
   n.LocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroup")
        }
        return n
}

func (n *TeachingGroup) Set(key string, value interface{}) *TeachingGroup {
        if n == nil {
                n = TeachingGroupCreate(TeachingGroup{})
        }
        switch key {
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentList":
                      n.StudentList = StudentListTypeCreate(value.(StudentListType))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Sett":
                      n.Sett = StringCreate(value.(string))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "TeachingGroupPeriodList":
                      n.TeachingGroupPeriodList = TeachingGroupPeriodListTypeCreate(value.(TeachingGroupPeriodListType))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LongName":
                      n.LongName = StringCreate(value.(string))
    case "Block":
                      n.Block = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "GroupType":
                      n.GroupType = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "SchoolCourseLocalId":
    
                      n.SchoolCourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "MaxClassSize":
                      n.MaxClassSize = IntCreate(value.(int))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "MinClassSize":
                      n.MinClassSize = IntCreate(value.(int))
    case "SchoolInfoRefId":
    
                      n.SchoolInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TeacherList":
                      n.TeacherList = TeacherListTypeCreate(value.(TeacherListType))
    case "CurriculumLevel":
                      n.CurriculumLevel = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroup")
        }
        return n
}


func (t *PersonPicture) CopyString(key string, value interface{}) *PersonPicture {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonPicture) Unset(key string) *PersonPicture {
        switch key {
  case "ParentObjectRefId":
   n.ParentObjectRefId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "RefId":
   n.RefId = nil
  case "PictureSource":
   n.PictureSource = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "PublishingPermissionList":
   n.PublishingPermissionList = nil
  case "OKToPublish":
   n.OKToPublish = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture")
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
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PictureSource":
                      n.PictureSource = PersonPicture_PictureSourceCreate(value.(PersonPicture_PictureSource))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture")
        }
        return n
}


func (t *PersonInvolvementType) CopyString(key string, value interface{}) *PersonInvolvementType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonInvolvementType) Unset(key string) *PersonInvolvementType {
        switch key {
  case "HowInvolved":
   n.HowInvolved = nil
  case "PersonRefId":
   n.PersonRefId = nil
  case "ShortName":
   n.ShortName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInvolvementType")
        }
        return n
}

func (n *PersonInvolvementType) Set(key string, value interface{}) *PersonInvolvementType {
        if n == nil {
                n = PersonInvolvementTypeCreate(PersonInvolvementType{})
        }
        switch key {
    case "HowInvolved":
                      n.HowInvolved = StringCreate(value.(string))
    case "PersonRefId":
                      n.PersonRefId = PersonInvolvementType_PersonRefIdCreate(value.(PersonInvolvementType_PersonRefId))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInvolvementType")
        }
        return n
}


func (t *ProgramStatusType) CopyString(key string, value interface{}) *ProgramStatusType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ProgramStatusType) Unset(key string) *ProgramStatusType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramStatusType")
        }
        return n
}

func (n *ProgramStatusType) Set(key string, value interface{}) *ProgramStatusType {
        if n == nil {
                n = ProgramStatusTypeCreate(ProgramStatusType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Code":
                      n.Code = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramStatusType")
        }
        return n
}


func (t *AggregateCharacteristicInfo) CopyString(key string, value interface{}) *AggregateCharacteristicInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AggregateCharacteristicInfo) Unset(key string) *AggregateCharacteristicInfo {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "RefId":
   n.RefId = nil
  case "Description":
   n.Description = nil
  case "Definition":
   n.Definition = nil
  case "ElementName":
   n.ElementName = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateCharacteristicInfo")
        }
        return n
}

func (n *AggregateCharacteristicInfo) Set(key string, value interface{}) *AggregateCharacteristicInfo {
        if n == nil {
                n = AggregateCharacteristicInfoCreate(AggregateCharacteristicInfo{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Definition":
                      n.Definition = StringCreate(value.(string))
    case "ElementName":
                      n.ElementName = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateCharacteristicInfo")
        }
        return n
}


func (t *PasswordListType_Password) CopyString(key string, value interface{}) *PasswordListType_Password {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PasswordListType_Password) Unset(key string) *PasswordListType_Password {
        switch key {
  case "Algorithm":
   n.Algorithm = nil
  case "Value":
   n.Value = nil
  case "KeyName":
   n.KeyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PasswordListType_Password")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PasswordListType_Password")
        }
        return n
}


func (t *ElectronicIdType) CopyString(key string, value interface{}) *ElectronicIdType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ElectronicIdType) Unset(key string) *ElectronicIdType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ElectronicIdType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ElectronicIdType")
        }
        return n
}


func (t *AbstractContentElementType_XMLData) CopyString(key string, value interface{}) *AbstractContentElementType_XMLData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentElementType_XMLData) Unset(key string) *AbstractContentElementType_XMLData {
        switch key {
  case "Description":
   n.Description = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_XMLData")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_XMLData")
        }
        return n
}


func (t *CreationUserType) CopyString(key string, value interface{}) *CreationUserType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CreationUserType) Unset(key string) *CreationUserType {
        switch key {
  case "UserId":
   n.UserId = nil
  case "Type":
   n.Type = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CreationUserType")
        }
        return n
}

func (n *CreationUserType) Set(key string, value interface{}) *CreationUserType {
        if n == nil {
                n = CreationUserTypeCreate(CreationUserType{})
        }
        switch key {
    case "UserId":
                      n.UserId = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CreationUserType")
        }
        return n
}


func (t *LifeCycleType_Modified) CopyString(key string, value interface{}) *LifeCycleType_Modified {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LifeCycleType_Modified) Unset(key string) *LifeCycleType_Modified {
        switch key {
  case "Description":
   n.Description = nil
  case "DateTime":
   n.DateTime = nil
  case "By":
   n.By = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Modified")
        }
        return n
}

func (n *LifeCycleType_Modified) Set(key string, value interface{}) *LifeCycleType_Modified {
        if n == nil {
                n = LifeCycleType_ModifiedCreate(LifeCycleType_Modified{})
        }
        switch key {
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "DateTime":
                      n.DateTime = StringCreate(value.(string))
    case "By":
                      n.By = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Modified")
        }
        return n
}


func (t *StudentContactPersonal) CopyString(key string, value interface{}) *StudentContactPersonal {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentContactPersonal) Unset(key string) *StudentContactPersonal {
        switch key {
  case "EmploymentType":
   n.EmploymentType = nil
  case "LocalId":
   n.LocalId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "OtherIdList":
   n.OtherIdList = nil
  case "PersonInfo":
   n.PersonInfo = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "RefId":
   n.RefId = nil
  case "NonSchoolEducation":
   n.NonSchoolEducation = nil
  case "SchoolEducationalLevel":
   n.SchoolEducationalLevel = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentContactPersonal")
        }
        return n
}

func (n *StudentContactPersonal) Set(key string, value interface{}) *StudentContactPersonal {
        if n == nil {
                n = StudentContactPersonalCreate(StudentContactPersonal{})
        }
        switch key {
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
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
      present := false
  for _, b := range AUCodeSetsSchoolEducationLevelTypeType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsSchoolEducationLevelTypeType_values")
      }

                      n.SchoolEducationalLevel = ((*EducationalLevelType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentContactPersonal")
        }
        return n
}


func (t *PurchaseOrder) CopyString(key string, value interface{}) *PurchaseOrder {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PurchaseOrder) Unset(key string) *PurchaseOrder {
        switch key {
  case "EmployeePersonalRefId":
   n.EmployeePersonalRefId = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ChargedLocationInfoRefId":
   n.ChargedLocationInfoRefId = nil
  case "TaxAmount":
   n.TaxAmount = nil
  case "TaxRate":
   n.TaxRate = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TotalAmount":
   n.TotalAmount = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "CreationDate":
   n.CreationDate = nil
  case "PurchasingItems":
   n.PurchasingItems = nil
  case "FormNumber":
   n.FormNumber = nil
  case "OriginalPurchaseOrderRefId":
   n.OriginalPurchaseOrderRefId = nil
  case "UpdateDate":
   n.UpdateDate = nil
  case "VendorInfoRefId":
   n.VendorInfoRefId = nil
  case "FullyDelivered":
   n.FullyDelivered = nil
  case "LocalId":
   n.LocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PurchaseOrder")
        }
        return n
}

func (n *PurchaseOrder) Set(key string, value interface{}) *PurchaseOrder {
        if n == nil {
                n = PurchaseOrderCreate(PurchaseOrder{})
        }
        switch key {
    case "EmployeePersonalRefId":
                      n.EmployeePersonalRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TotalAmount":
                      n.TotalAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "PurchasingItems":
                      n.PurchasingItems = PurchasingItemsTypeCreate(value.(PurchasingItemsType))
    case "FormNumber":
                      n.FormNumber = StringCreate(value.(string))
    case "OriginalPurchaseOrderRefId":
                      n.OriginalPurchaseOrderRefId = StringCreate(value.(string))
    case "UpdateDate":
                      n.UpdateDate = StringCreate(value.(string))
    case "VendorInfoRefId":
                      n.VendorInfoRefId = StringCreate(value.(string))
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
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PurchaseOrder")
        }
        return n
}


func (t *TimeTableDayType) CopyString(key string, value interface{}) *TimeTableDayType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableDayType) Unset(key string) *TimeTableDayType {
        switch key {
  case "TimeTablePeriodList":
   n.TimeTablePeriodList = nil
  case "DayId":
   n.DayId = nil
  case "DayTitle":
   n.DayTitle = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableDayType")
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
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DayTitle":
                      n.DayTitle = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableDayType")
        }
        return n
}


func (t *CalendarDate) CopyString(key string, value interface{}) *CalendarDate {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CalendarDate) Unset(key string) *CalendarDate {
        switch key {
  case "Date":
   n.Date = nil
  case "TeacherAttendance":
   n.TeacherAttendance = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "CalendarDateRefId":
   n.CalendarDateRefId = nil
  case "AdministratorAttendance":
   n.AdministratorAttendance = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "CalendarDateType":
   n.CalendarDateType = nil
  case "CalendarSummaryRefId":
   n.CalendarSummaryRefId = nil
  case "StudentAttendance":
   n.StudentAttendance = nil
  case "CalendarDateNumber":
   n.CalendarDateNumber = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarDate")
        }
        return n
}

func (n *CalendarDate) Set(key string, value interface{}) *CalendarDate {
        if n == nil {
                n = CalendarDateCreate(CalendarDate{})
        }
        switch key {
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "TeacherAttendance":
                      n.TeacherAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "CalendarDateRefId":
                      n.CalendarDateRefId = StringCreate(value.(string))
    case "AdministratorAttendance":
                      n.AdministratorAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "CalendarDateType":
                      n.CalendarDateType = CalendarDateInfoTypeCreate(value.(CalendarDateInfoType))
    case "CalendarSummaryRefId":
                      n.CalendarSummaryRefId = StringCreate(value.(string))
    case "StudentAttendance":
                      n.StudentAttendance = AttendanceInfoTypeCreate(value.(AttendanceInfoType))
    case "CalendarDateNumber":
                      n.CalendarDateNumber = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarDate")
        }
        return n
}


func (t *Identity) CopyString(key string, value interface{}) *Identity {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Identity) Unset(key string) *Identity {
        switch key {
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "IdentityAssertions":
   n.IdentityAssertions = nil
  case "RefId":
   n.RefId = nil
  case "AuthenticationSource":
   n.AuthenticationSource = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SIF_RefId":
   n.SIF_RefId = nil
  case "AuthenticationSourceGlobalUID":
   n.AuthenticationSourceGlobalUID = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "PasswordList":
   n.PasswordList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Identity")
        }
        return n
}

func (n *Identity) Set(key string, value interface{}) *Identity {
        if n == nil {
                n = IdentityCreate(Identity{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "IdentityAssertions":
                      n.IdentityAssertions = IdentityAssertionsTypeCreate(value.(IdentityAssertionsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AuthenticationSource":
                      n.AuthenticationSource = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_RefId":
                      n.SIF_RefId = Identity_SIF_RefIdCreate(value.(Identity_SIF_RefId))
    case "AuthenticationSourceGlobalUID":
                      n.AuthenticationSourceGlobalUID = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PasswordList":
                      n.PasswordList = PasswordListTypeCreate(value.(PasswordListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Identity")
        }
        return n
}


func (t *StudentExitContainerType) CopyString(key string, value interface{}) *StudentExitContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentExitContainerType) Unset(key string) *StudentExitContainerType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentExitContainerType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentExitContainerType")
        }
        return n
}


func (t *SchoolInfo_OtherLEA) CopyString(key string, value interface{}) *SchoolInfo_OtherLEA {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolInfo_OtherLEA) Unset(key string) *SchoolInfo_OtherLEA {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolInfo_OtherLEA")
        }
        return n
}

func (n *SchoolInfo_OtherLEA) Set(key string, value interface{}) *SchoolInfo_OtherLEA {
        if n == nil {
                n = SchoolInfo_OtherLEACreate(SchoolInfo_OtherLEA{})
        }
        switch key {
    case "Value":
    
                      n.Value = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolInfo_OtherLEA")
        }
        return n
}


func (t *Activity_Evaluation) CopyString(key string, value interface{}) *Activity_Evaluation {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Activity_Evaluation) Unset(key string) *Activity_Evaluation {
        switch key {
  case "EvaluationType":
   n.EvaluationType = nil
  case "Description":
   n.Description = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Activity_Evaluation")
        }
        return n
}

func (n *Activity_Evaluation) Set(key string, value interface{}) *Activity_Evaluation {
        if n == nil {
                n = Activity_EvaluationCreate(Activity_Evaluation{})
        }
        switch key {
    case "EvaluationType":
                      n.EvaluationType = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Activity_Evaluation")
        }
        return n
}


func (t *CollectionRound) CopyString(key string, value interface{}) *CollectionRound {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CollectionRound) Unset(key string) *CollectionRound {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "CollectionYear":
   n.CollectionYear = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AGCollection":
   n.AGCollection = nil
  case "AGRoundList":
   n.AGRoundList = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CollectionRound")
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
    case "CollectionYear":
    
                      n.CollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "AGRoundList":
                      n.AGRoundList = AGRoundListTypeCreate(value.(AGRoundListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CollectionRound")
        }
        return n
}


func (t *StudentSchoolEnrollment) CopyString(key string, value interface{}) *StudentSchoolEnrollment {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSchoolEnrollment) Unset(key string) *StudentSchoolEnrollment {
        switch key {
  case "Counselor":
   n.Counselor = nil
  case "EntryType":
   n.EntryType = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "DistanceEducationStudent":
   n.DistanceEducationStudent = nil
  case "BoardingStatus":
   n.BoardingStatus = nil
  case "DisabilityCategory":
   n.DisabilityCategory = nil
  case "YearLevel":
   n.YearLevel = nil
  case "FTE":
   n.FTE = nil
  case "StudentSubjectChoiceList":
   n.StudentSubjectChoiceList = nil
  case "DisabilityLevelOfAdjustment":
   n.DisabilityLevelOfAdjustment = nil
  case "CensusAge":
   n.CensusAge = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "CatchmentStatus":
   n.CatchmentStatus = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "MembershipType":
   n.MembershipType = nil
  case "ACARASchoolId":
   n.ACARASchoolId = nil
  case "RecordClosureReason":
   n.RecordClosureReason = nil
  case "InternationalStudent":
   n.InternationalStudent = nil
  case "DestinationSchoolName":
   n.DestinationSchoolName = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ExitType":
   n.ExitType = nil
  case "PublishingPermissionList":
   n.PublishingPermissionList = nil
  case "IndividualLearningPlan":
   n.IndividualLearningPlan = nil
  case "EntryDate":
   n.EntryDate = nil
  case "ReportingSchool":
   n.ReportingSchool = nil
  case "FTPTStatus":
   n.FTPTStatus = nil
  case "ClassCode":
   n.ClassCode = nil
  case "ExitStatus":
   n.ExitStatus = nil
  case "DestinationSchool":
   n.DestinationSchool = nil
  case "PromotionInfo":
   n.PromotionInfo = nil
  case "TimeFrame":
   n.TimeFrame = nil
  case "Calendar":
   n.Calendar = nil
  case "House":
   n.House = nil
  case "StartedAtSchoolDate":
   n.StartedAtSchoolDate = nil
  case "Homegroup":
   n.Homegroup = nil
  case "Homeroom":
   n.Homeroom = nil
  case "StudentGroupList":
   n.StudentGroupList = nil
  case "ExitDate":
   n.ExitDate = nil
  case "TestLevel":
   n.TestLevel = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "Advisor":
   n.Advisor = nil
  case "RefId":
   n.RefId = nil
  case "PreviousSchoolName":
   n.PreviousSchoolName = nil
  case "FFPOS":
   n.FFPOS = nil
  case "PreviousSchool":
   n.PreviousSchool = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment")
        }
        return n
}

func (n *StudentSchoolEnrollment) Set(key string, value interface{}) *StudentSchoolEnrollment {
        if n == nil {
                n = StudentSchoolEnrollmentCreate(StudentSchoolEnrollment{})
        }
        switch key {
    case "Counselor":
                      n.Counselor = StudentSchoolEnrollment_CounselorCreate(value.(StudentSchoolEnrollment_Counselor))
    case "EntryType":
                      n.EntryType = StudentEntryContainerTypeCreate(value.(StudentEntryContainerType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
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
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "StudentSubjectChoiceList":
                      n.StudentSubjectChoiceList = StudentSubjectChoiceListTypeCreate(value.(StudentSubjectChoiceListType))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "CensusAge":
                      n.CensusAge = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "CatchmentStatus":
                      n.CatchmentStatus = CatchmentStatusContainerTypeCreate(value.(CatchmentStatusContainerType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
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
    case "ACARASchoolId":
    
                      n.ACARASchoolId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "DestinationSchoolName":
                      n.DestinationSchoolName = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ExitType":
                      n.ExitType = StudentExitContainerTypeCreate(value.(StudentExitContainerType))
    case "PublishingPermissionList":
                      n.PublishingPermissionList = PublishingPermissionListTypeCreate(value.(PublishingPermissionListType))
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
    case "EntryDate":
                      n.EntryDate = StringCreate(value.(string))
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
    case "ClassCode":
                      n.ClassCode = StringCreate(value.(string))
    case "ExitStatus":
                      n.ExitStatus = StudentExitStatusContainerTypeCreate(value.(StudentExitStatusContainerType))
    case "DestinationSchool":
    
                      n.DestinationSchool = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "Calendar":
                      n.Calendar = StudentSchoolEnrollment_CalendarCreate(value.(StudentSchoolEnrollment_Calendar))
    case "House":
                      n.House = StringCreate(value.(string))
    case "StartedAtSchoolDate":
                      n.StartedAtSchoolDate = StringCreate(value.(string))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "Homeroom":
                      n.Homeroom = StudentSchoolEnrollment_HomeroomCreate(value.(StudentSchoolEnrollment_Homeroom))
    case "StudentGroupList":
                      n.StudentGroupList = StudentGroupListTypeCreate(value.(StudentGroupListType))
    case "ExitDate":
                      n.ExitDate = StringCreate(value.(string))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "Advisor":
                      n.Advisor = StudentSchoolEnrollment_AdvisorCreate(value.(StudentSchoolEnrollment_Advisor))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PreviousSchoolName":
                      n.PreviousSchoolName = StringCreate(value.(string))
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
    case "PreviousSchool":
    
                      n.PreviousSchool = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment")
        }
        return n
}


func (t *AbstractContentElementType) CopyString(key string, value interface{}) *AbstractContentElementType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentElementType) Unset(key string) *AbstractContentElementType {
        switch key {
  case "Reference":
   n.Reference = nil
  case "BinaryData":
   n.BinaryData = nil
  case "XMLData":
   n.XMLData = nil
  case "TextData":
   n.TextData = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType")
        }
        return n
}

func (n *AbstractContentElementType) Set(key string, value interface{}) *AbstractContentElementType {
        if n == nil {
                n = AbstractContentElementTypeCreate(AbstractContentElementType{})
        }
        switch key {
    case "Reference":
                      n.Reference = AbstractContentElementType_ReferenceCreate(value.(AbstractContentElementType_Reference))
    case "BinaryData":
                      n.BinaryData = AbstractContentElementType_BinaryDataCreate(value.(AbstractContentElementType_BinaryData))
    case "XMLData":
                      n.XMLData = AbstractContentElementType_XMLDataCreate(value.(AbstractContentElementType_XMLData))
    case "TextData":
                      n.TextData = AbstractContentElementType_TextDataCreate(value.(AbstractContentElementType_TextData))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType")
        }
        return n
}


func (t *ValidLetterMarkType) CopyString(key string, value interface{}) *ValidLetterMarkType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ValidLetterMarkType) Unset(key string) *ValidLetterMarkType {
        switch key {
  case "NumericEquivalent":
   n.NumericEquivalent = nil
  case "Description":
   n.Description = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ValidLetterMarkType")
        }
        return n
}

func (n *ValidLetterMarkType) Set(key string, value interface{}) *ValidLetterMarkType {
        if n == nil {
                n = ValidLetterMarkTypeCreate(ValidLetterMarkType{})
        }
        switch key {
    case "NumericEquivalent":
                      n.NumericEquivalent = FloatCreate(value.(float64))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Code":
                      n.Code = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ValidLetterMarkType")
        }
        return n
}


func (t *TotalEnrollmentsType) CopyString(key string, value interface{}) *TotalEnrollmentsType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TotalEnrollmentsType) Unset(key string) *TotalEnrollmentsType {
        switch key {
  case "Boys":
   n.Boys = nil
  case "Girls":
   n.Girls = nil
  case "TotalStudents":
   n.TotalStudents = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TotalEnrollmentsType")
        }
        return n
}

func (n *TotalEnrollmentsType) Set(key string, value interface{}) *TotalEnrollmentsType {
        if n == nil {
                n = TotalEnrollmentsTypeCreate(TotalEnrollmentsType{})
        }
        switch key {
    case "Boys":
                      n.Boys = StringCreate(value.(string))
    case "Girls":
                      n.Girls = StringCreate(value.(string))
    case "TotalStudents":
                      n.TotalStudents = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TotalEnrollmentsType")
        }
        return n
}


func (t *GradingAssignmentScore) CopyString(key string, value interface{}) *GradingAssignmentScore {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *GradingAssignmentScore) Unset(key string) *GradingAssignmentScore {
        switch key {
  case "ScorePoints":
   n.ScorePoints = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "ExpectedScore":
   n.ExpectedScore = nil
  case "AssignmentScoreIteration":
   n.AssignmentScoreIteration = nil
  case "ScoreDescription":
   n.ScoreDescription = nil
  case "DateGraded":
   n.DateGraded = nil
  case "RefId":
   n.RefId = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "ScorePercent":
   n.ScorePercent = nil
  case "StudentPersonalLocalId":
   n.StudentPersonalLocalId = nil
  case "ScoreLetter":
   n.ScoreLetter = nil
  case "SubscoreList":
   n.SubscoreList = nil
  case "GradingAssignmentRefId":
   n.GradingAssignmentRefId = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "TeacherJudgement":
   n.TeacherJudgement = nil
  case "TeachingGroupRefId":
   n.TeachingGroupRefId = nil
  case "MarkInfoRefId":
   n.MarkInfoRefId = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradingAssignmentScore")
        }
        return n
}

func (n *GradingAssignmentScore) Set(key string, value interface{}) *GradingAssignmentScore {
        if n == nil {
                n = GradingAssignmentScoreCreate(GradingAssignmentScore{})
        }
        switch key {
    case "ScorePoints":
                      n.ScorePoints = IntCreate(value.(int))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ExpectedScore":
                      n.ExpectedScore = BoolCreate(value.(bool))
    case "AssignmentScoreIteration":
                      n.AssignmentScoreIteration = StringCreate(value.(string))
    case "ScoreDescription":
                      n.ScoreDescription = StringCreate(value.(string))
    case "DateGraded":
                      n.DateGraded = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "ScorePercent":
                      n.ScorePercent = FloatCreate(value.(float64))
    case "StudentPersonalLocalId":
    
                      n.StudentPersonalLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ScoreLetter":
                      n.ScoreLetter = StringCreate(value.(string))
    case "SubscoreList":
                      n.SubscoreList = NAPSubscoreListTypeCreate(value.(NAPSubscoreListType))
    case "GradingAssignmentRefId":
                      n.GradingAssignmentRefId = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "TeacherJudgement":
                      n.TeacherJudgement = StringCreate(value.(string))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "MarkInfoRefId":
                      n.MarkInfoRefId = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradingAssignmentScore")
        }
        return n
}


func (t *Activity) CopyString(key string, value interface{}) *Activity {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Activity) Unset(key string) *Activity {
        switch key {
  case "SoftwareRequirementList":
   n.SoftwareRequirementList = nil
  case "LearningResources":
   n.LearningResources = nil
  case "Students":
   n.Students = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Title":
   n.Title = nil
  case "AssessmentRefId":
   n.AssessmentRefId = nil
  case "SubjectArea":
   n.SubjectArea = nil
  case "MaxAttemptsAllowed":
   n.MaxAttemptsAllowed = nil
  case "EssentialMaterials":
   n.EssentialMaterials = nil
  case "Preamble":
   n.Preamble = nil
  case "LearningObjectives":
   n.LearningObjectives = nil
  case "SourceObjects":
   n.SourceObjects = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Prerequisites":
   n.Prerequisites = nil
  case "TechnicalRequirements":
   n.TechnicalRequirements = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ActivityTime":
   n.ActivityTime = nil
  case "Evaluation":
   n.Evaluation = nil
  case "Points":
   n.Points = nil
  case "LearningStandards":
   n.LearningStandards = nil
  case "RefId":
   n.RefId = nil
  case "ActivityWeight":
   n.ActivityWeight = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Activity")
        }
        return n
}

func (n *Activity) Set(key string, value interface{}) *Activity {
        if n == nil {
                n = ActivityCreate(Activity{})
        }
        switch key {
    case "SoftwareRequirementList":
                      n.SoftwareRequirementList = SoftwareRequirementListTypeCreate(value.(SoftwareRequirementListType))
    case "LearningResources":
                      n.LearningResources = LearningResourcesTypeCreate(value.(LearningResourcesType))
    case "Students":
                      n.Students = StudentsTypeCreate(value.(StudentsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "AssessmentRefId":
                      n.AssessmentRefId = StringCreate(value.(string))
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
    case "MaxAttemptsAllowed":
                      n.MaxAttemptsAllowed = IntCreate(value.(int))
    case "EssentialMaterials":
                      n.EssentialMaterials = EssentialMaterialsTypeCreate(value.(EssentialMaterialsType))
    case "Preamble":
                      n.Preamble = StringCreate(value.(string))
    case "LearningObjectives":
                      n.LearningObjectives = LearningObjectivesTypeCreate(value.(LearningObjectivesType))
    case "SourceObjects":
                      n.SourceObjects = SourceObjectsTypeCreate(value.(SourceObjectsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Prerequisites":
                      n.Prerequisites = PrerequisitesTypeCreate(value.(PrerequisitesType))
    case "TechnicalRequirements":
                      n.TechnicalRequirements = TechnicalRequirementsTypeCreate(value.(TechnicalRequirementsType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ActivityTime":
                      n.ActivityTime = ActivityTimeTypeCreate(value.(ActivityTimeType))
    case "Evaluation":
                      n.Evaluation = Activity_EvaluationCreate(value.(Activity_Evaluation))
    case "Points":
                      n.Points = IntCreate(value.(int))
    case "LearningStandards":
                      n.LearningStandards = LearningStandardsTypeCreate(value.(LearningStandardsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ActivityWeight":
                      n.ActivityWeight = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Activity")
        }
        return n
}


func (t *StudentPeriodAttendance) CopyString(key string, value interface{}) *StudentPeriodAttendance {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentPeriodAttendance) Unset(key string) *StudentPeriodAttendance {
        switch key {
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "AuditInfo":
   n.AuditInfo = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "AttendanceCode":
   n.AttendanceCode = nil
  case "ScheduledActivityRefId":
   n.ScheduledActivityRefId = nil
  case "AttendanceComment":
   n.AttendanceComment = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TimeOut":
   n.TimeOut = nil
  case "RefId":
   n.RefId = nil
  case "Date":
   n.Date = nil
  case "TimetablePeriod":
   n.TimetablePeriod = nil
  case "TimeIn":
   n.TimeIn = nil
  case "AttendanceStatus":
   n.AttendanceStatus = nil
  case "SessionInfoRefId":
   n.SessionInfoRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentPeriodAttendance")
        }
        return n
}

func (n *StudentPeriodAttendance) Set(key string, value interface{}) *StudentPeriodAttendance {
        if n == nil {
                n = StudentPeriodAttendanceCreate(StudentPeriodAttendance{})
        }
        switch key {
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "AuditInfo":
                      n.AuditInfo = AuditInfoTypeCreate(value.(AuditInfoType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "AttendanceCode":
                      n.AttendanceCode = AttendanceCodeTypeCreate(value.(AttendanceCodeType))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "AttendanceComment":
                      n.AttendanceComment = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TimeOut":
                      n.TimeOut = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "TimetablePeriod":
                      n.TimetablePeriod = StringCreate(value.(string))
    case "TimeIn":
                      n.TimeIn = StringCreate(value.(string))
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
    case "SessionInfoRefId":
                      n.SessionInfoRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentPeriodAttendance")
        }
        return n
}


func (t *NAPStudentResponseSet) CopyString(key string, value interface{}) *NAPStudentResponseSet {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPStudentResponseSet) Unset(key string) *NAPStudentResponseSet {
        switch key {
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "PathTakenForDomain":
   n.PathTakenForDomain = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "PlatformStudentIdentifier":
   n.PlatformStudentIdentifier = nil
  case "ParallelTest":
   n.ParallelTest = nil
  case "NAPTestRefId":
   n.NAPTestRefId = nil
  case "DomainScore":
   n.DomainScore = nil
  case "NAPTestLocalId":
   n.NAPTestLocalId = nil
  case "EquatingSampleFlag":
   n.EquatingSampleFlag = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "CalibrationSampleFlag":
   n.CalibrationSampleFlag = nil
  case "ReportExclusionFlag":
   n.ReportExclusionFlag = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TestletList":
   n.TestletList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPStudentResponseSet")
        }
        return n
}

func (n *NAPStudentResponseSet) Set(key string, value interface{}) *NAPStudentResponseSet {
        if n == nil {
                n = NAPStudentResponseSetCreate(NAPStudentResponseSet{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PathTakenForDomain":
                      n.PathTakenForDomain = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "PlatformStudentIdentifier":
    
                      n.PlatformStudentIdentifier = ((*LocalIdType)(StringCreate(value.(string))))
    case "ParallelTest":
                      n.ParallelTest = StringCreate(value.(string))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "DomainScore":
                      n.DomainScore = DomainScoreTypeCreate(value.(DomainScoreType))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EquatingSampleFlag":
                      n.EquatingSampleFlag = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "CalibrationSampleFlag":
                      n.CalibrationSampleFlag = StringCreate(value.(string))
    case "ReportExclusionFlag":
                      n.ReportExclusionFlag = BoolCreate(value.(bool))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TestletList":
                      n.TestletList = NAPStudentResponseTestletListTypeCreate(value.(NAPStudentResponseTestletListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPStudentResponseSet")
        }
        return n
}


func (t *StudentParticipation) CopyString(key string, value interface{}) *StudentParticipation {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentParticipation) Unset(key string) *StudentParticipation {
        switch key {
  case "EntryPerson":
   n.EntryPerson = nil
  case "ReevaluationDate":
   n.ReevaluationDate = nil
  case "ExtendedSchoolYear":
   n.ExtendedSchoolYear = nil
  case "ParticipationContact":
   n.ParticipationContact = nil
  case "ProgramPlacementDate":
   n.ProgramPlacementDate = nil
  case "ProgramType":
   n.ProgramType = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StudentSpecialEducationFTE":
   n.StudentSpecialEducationFTE = nil
  case "ExtensionComments":
   n.ExtensionComments = nil
  case "ManagingSchool":
   n.ManagingSchool = nil
  case "ProgramPlanDate":
   n.ProgramPlanDate = nil
  case "ProgramStatus":
   n.ProgramStatus = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "EvaluationDate":
   n.EvaluationDate = nil
  case "EvaluationExtensionDate":
   n.EvaluationExtensionDate = nil
  case "ExtendedDay":
   n.ExtendedDay = nil
  case "ProgramPlanEffectiveDate":
   n.ProgramPlanEffectiveDate = nil
  case "NOREPDate":
   n.NOREPDate = nil
  case "RefId":
   n.RefId = nil
  case "ReferralSource":
   n.ReferralSource = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ProgramEligibilityDate":
   n.ProgramEligibilityDate = nil
  case "StudentParticipationAsOfDate":
   n.StudentParticipationAsOfDate = nil
  case "ReferralDate":
   n.ReferralDate = nil
  case "ProgramFundingSources":
   n.ProgramFundingSources = nil
  case "EvaluationParentalConsentDate":
   n.EvaluationParentalConsentDate = nil
  case "GiftedEligibilityCriteria":
   n.GiftedEligibilityCriteria = nil
  case "PlacementParentalConsentDate":
   n.PlacementParentalConsentDate = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "ProgramAvailability":
   n.ProgramAvailability = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentParticipation")
        }
        return n
}

func (n *StudentParticipation) Set(key string, value interface{}) *StudentParticipation {
        if n == nil {
                n = StudentParticipationCreate(StudentParticipation{})
        }
        switch key {
    case "EntryPerson":
                      n.EntryPerson = StringCreate(value.(string))
    case "ReevaluationDate":
                      n.ReevaluationDate = StringCreate(value.(string))
    case "ExtendedSchoolYear":
                      n.ExtendedSchoolYear = BoolCreate(value.(bool))
    case "ParticipationContact":
                      n.ParticipationContact = StringCreate(value.(string))
    case "ProgramPlacementDate":
                      n.ProgramPlacementDate = StringCreate(value.(string))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentSpecialEducationFTE":
                      n.StudentSpecialEducationFTE = FloatCreate(value.(float64))
    case "ExtensionComments":
                      n.ExtensionComments = StringCreate(value.(string))
    case "ManagingSchool":
                      n.ManagingSchool = StudentParticipation_ManagingSchoolCreate(value.(StudentParticipation_ManagingSchool))
    case "ProgramPlanDate":
                      n.ProgramPlanDate = StringCreate(value.(string))
    case "ProgramStatus":
                      n.ProgramStatus = ProgramStatusTypeCreate(value.(ProgramStatusType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "EvaluationDate":
                      n.EvaluationDate = StringCreate(value.(string))
    case "EvaluationExtensionDate":
                      n.EvaluationExtensionDate = StringCreate(value.(string))
    case "ExtendedDay":
                      n.ExtendedDay = BoolCreate(value.(bool))
    case "ProgramPlanEffectiveDate":
                      n.ProgramPlanEffectiveDate = StringCreate(value.(string))
    case "NOREPDate":
                      n.NOREPDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ReferralSource":
                      n.ReferralSource = ReferralSourceTypeCreate(value.(ReferralSourceType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ProgramEligibilityDate":
                      n.ProgramEligibilityDate = StringCreate(value.(string))
    case "StudentParticipationAsOfDate":
                      n.StudentParticipationAsOfDate = StringCreate(value.(string))
    case "ReferralDate":
                      n.ReferralDate = StringCreate(value.(string))
    case "ProgramFundingSources":
                      n.ProgramFundingSources = ProgramFundingSourcesTypeCreate(value.(ProgramFundingSourcesType))
    case "EvaluationParentalConsentDate":
                      n.EvaluationParentalConsentDate = StringCreate(value.(string))
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
    case "PlacementParentalConsentDate":
                      n.PlacementParentalConsentDate = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "ProgramAvailability":
                      n.ProgramAvailability = ProgramAvailabilityTypeCreate(value.(ProgramAvailabilityType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentParticipation")
        }
        return n
}


func (t *StudentSchoolEnrollment_Counselor) CopyString(key string, value interface{}) *StudentSchoolEnrollment_Counselor {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSchoolEnrollment_Counselor) Unset(key string) *StudentSchoolEnrollment_Counselor {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Counselor")
        }
        return n
}

func (n *StudentSchoolEnrollment_Counselor) Set(key string, value interface{}) *StudentSchoolEnrollment_Counselor {
        if n == nil {
                n = StudentSchoolEnrollment_CounselorCreate(StudentSchoolEnrollment_Counselor{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Counselor")
        }
        return n
}


func (t *TimeTableContainer) CopyString(key string, value interface{}) *TimeTableContainer {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableContainer) Unset(key string) *TimeTableContainer {
        switch key {
  case "TimeTableSchedule":
   n.TimeTableSchedule = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TeachingGroupScheduleList":
   n.TeachingGroupScheduleList = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "RefId":
   n.RefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "TimeTableScheduleCellList":
   n.TimeTableScheduleCellList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableContainer")
        }
        return n
}

func (n *TimeTableContainer) Set(key string, value interface{}) *TimeTableContainer {
        if n == nil {
                n = TimeTableContainerCreate(TimeTableContainer{})
        }
        switch key {
    case "TimeTableSchedule":
                      n.TimeTableSchedule = TimeTableScheduleTypeCreate(value.(TimeTableScheduleType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TeachingGroupScheduleList":
                      n.TeachingGroupScheduleList = TeachingGroupScheduleTypeCreate(value.(TeachingGroupScheduleType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TimeTableScheduleCellList":
                      n.TimeTableScheduleCellList = TimeTableScheduleCellListTypeCreate(value.(TimeTableScheduleCellListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableContainer")
        }
        return n
}


func (t *FinancialAccount) CopyString(key string, value interface{}) *FinancialAccount {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FinancialAccount) Unset(key string) *FinancialAccount {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AccountNumber":
   n.AccountNumber = nil
  case "AccountCode":
   n.AccountCode = nil
  case "ChargedLocationInfoRefId":
   n.ChargedLocationInfoRefId = nil
  case "ClassType":
   n.ClassType = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "RefId":
   n.RefId = nil
  case "LocalId":
   n.LocalId = nil
  case "ParentAccountRefId":
   n.ParentAccountRefId = nil
  case "Name":
   n.Name = nil
  case "CreationTime":
   n.CreationTime = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "CreationDate":
   n.CreationDate = nil
  case "Description":
   n.Description = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FinancialAccount")
        }
        return n
}

func (n *FinancialAccount) Set(key string, value interface{}) *FinancialAccount {
        if n == nil {
                n = FinancialAccountCreate(FinancialAccount{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AccountNumber":
                      n.AccountNumber = StringCreate(value.(string))
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "ClassType":
                      n.ClassType = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ParentAccountRefId":
                      n.ParentAccountRefId = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "CreationTime":
                      n.CreationTime = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FinancialAccount")
        }
        return n
}


func (t *WellbeingAppeal) CopyString(key string, value interface{}) *WellbeingAppeal {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingAppeal) Unset(key string) *WellbeingAppeal {
        switch key {
  case "LocalId":
   n.LocalId = nil
  case "DocumentList":
   n.DocumentList = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "AppealOutcome":
   n.AppealOutcome = nil
  case "WellbeingResponseRefId":
   n.WellbeingResponseRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "LocalAppealId":
   n.LocalAppealId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AppealNotes":
   n.AppealNotes = nil
  case "RefId":
   n.RefId = nil
  case "AppealStatusCode":
   n.AppealStatusCode = nil
  case "Date":
   n.Date = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingAppeal")
        }
        return n
}

func (n *WellbeingAppeal) Set(key string, value interface{}) *WellbeingAppeal {
        if n == nil {
                n = WellbeingAppealCreate(WellbeingAppeal{})
        }
        switch key {
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "AppealOutcome":
                      n.AppealOutcome = StringCreate(value.(string))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LocalAppealId":
    
                      n.LocalAppealId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AppealNotes":
                      n.AppealNotes = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
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
    case "Date":
                      n.Date = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingAppeal")
        }
        return n
}


func (t *AbstractContentElementType_BinaryData) CopyString(key string, value interface{}) *AbstractContentElementType_BinaryData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentElementType_BinaryData) Unset(key string) *AbstractContentElementType_BinaryData {
        switch key {
  case "MIMEType":
   n.MIMEType = nil
  case "FileName":
   n.FileName = nil
  case "Description":
   n.Description = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_BinaryData")
        }
        return n
}

func (n *AbstractContentElementType_BinaryData) Set(key string, value interface{}) *AbstractContentElementType_BinaryData {
        if n == nil {
                n = AbstractContentElementType_BinaryDataCreate(AbstractContentElementType_BinaryData{})
        }
        switch key {
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_BinaryData")
        }
        return n
}


func (t *SchoolProgramType) CopyString(key string, value interface{}) *SchoolProgramType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolProgramType) Unset(key string) *SchoolProgramType {
        switch key {
  case "Category":
   n.Category = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Type":
   n.Type = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolProgramType")
        }
        return n
}

func (n *SchoolProgramType) Set(key string, value interface{}) *SchoolProgramType {
        if n == nil {
                n = SchoolProgramTypeCreate(SchoolProgramType{})
        }
        switch key {
    case "Category":
                      n.Category = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "Type":
                      n.Type = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolProgramType")
        }
        return n
}


func (t *LifeCycleType) CopyString(key string, value interface{}) *LifeCycleType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LifeCycleType) Unset(key string) *LifeCycleType {
        switch key {
  case "Created":
   n.Created = nil
  case "TimeElements":
   n.TimeElements = nil
  case "ModificationHistory":
   n.ModificationHistory = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType")
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
    case "TimeElements":
                      n.TimeElements = LifeCycleType_TimeElementsCreate(value.(LifeCycleType_TimeElements))
    case "ModificationHistory":
                      n.ModificationHistory = LifeCycleType_ModificationHistoryCreate(value.(LifeCycleType_ModificationHistory))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType")
        }
        return n
}


func (t *ContactFlagsType) CopyString(key string, value interface{}) *ContactFlagsType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ContactFlagsType) Unset(key string) *ContactFlagsType {
        switch key {
  case "AccessToRecords":
   n.AccessToRecords = nil
  case "PickupRights":
   n.PickupRights = nil
  case "AttendanceContact":
   n.AttendanceContact = nil
  case "ReceivesAssessmentReport":
   n.ReceivesAssessmentReport = nil
  case "LivesWith":
   n.LivesWith = nil
  case "EmergencyContact":
   n.EmergencyContact = nil
  case "FeesBilling":
   n.FeesBilling = nil
  case "FeesAccess":
   n.FeesAccess = nil
  case "FamilyMail":
   n.FamilyMail = nil
  case "PrimaryCareProvider":
   n.PrimaryCareProvider = nil
  case "HasCustody":
   n.HasCustody = nil
  case "InterventionOrder":
   n.InterventionOrder = nil
  case "DisciplinaryContact":
   n.DisciplinaryContact = nil
  case "ParentLegalGuardian":
   n.ParentLegalGuardian = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactFlagsType")
        }
        return n
}

func (n *ContactFlagsType) Set(key string, value interface{}) *ContactFlagsType {
        if n == nil {
                n = ContactFlagsTypeCreate(ContactFlagsType{})
        }
        switch key {
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactFlagsType")
        }
        return n
}


func (t *TeachingGroupScheduleType) CopyString(key string, value interface{}) *TeachingGroupScheduleType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeachingGroupScheduleType) Unset(key string) *TeachingGroupScheduleType {
        switch key {
  case "GroupType":
   n.GroupType = nil
  case "SchoolCourseLocalId":
   n.SchoolCourseLocalId = nil
  case "MaxClassSize":
   n.MaxClassSize = nil
  case "SchoolCourseInfoRefId":
   n.SchoolCourseInfoRefId = nil
  case "Semester":
   n.Semester = nil
  case "Block":
   n.Block = nil
  case "TimeTableSubjectLocalId":
   n.TimeTableSubjectLocalId = nil
  case "StudentList":
   n.StudentList = nil
  case "TeacherList":
   n.TeacherList = nil
  case "CurriculumLevel":
   n.CurriculumLevel = nil
  case "TeachingGroupPeriodList":
   n.TeachingGroupPeriodList = nil
  case "ShortName":
   n.ShortName = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "LongName":
   n.LongName = nil
  case "LocalId":
   n.LocalId = nil
  case "EditorGUID":
   n.EditorGUID = nil
  case "MinClassSize":
   n.MinClassSize = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "Sett":
   n.Sett = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupScheduleType")
        }
        return n
}

func (n *TeachingGroupScheduleType) Set(key string, value interface{}) *TeachingGroupScheduleType {
        if n == nil {
                n = TeachingGroupScheduleTypeCreate(TeachingGroupScheduleType{})
        }
        switch key {
    case "GroupType":
                      n.GroupType = StringCreate(value.(string))
    case "SchoolCourseLocalId":
    
                      n.SchoolCourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "MaxClassSize":
                      n.MaxClassSize = IntCreate(value.(int))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
    case "Block":
                      n.Block = StringCreate(value.(string))
    case "TimeTableSubjectLocalId":
    
                      n.TimeTableSubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentList":
                      n.StudentList = StudentListTypeCreate(value.(StudentListType))
    case "TeacherList":
                      n.TeacherList = TeacherListTypeCreate(value.(TeacherListType))
    case "CurriculumLevel":
                      n.CurriculumLevel = StringCreate(value.(string))
    case "TeachingGroupPeriodList":
                      n.TeachingGroupPeriodList = TeachingGroupPeriodListTypeCreate(value.(TeachingGroupPeriodListType))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LongName":
                      n.LongName = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EditorGUID":
    
                      n.EditorGUID = ((*RefIdType)(StringCreate(value.(string))))
    case "MinClassSize":
                      n.MinClassSize = IntCreate(value.(int))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Sett":
                      n.Sett = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
    
                      n.SchoolInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupScheduleType")
        }
        return n
}


func (t *BaseNameType) CopyString(key string, value interface{}) *BaseNameType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *BaseNameType) Unset(key string) *BaseNameType {
        switch key {
  case "Suffix":
   n.Suffix = nil
  case "MiddleName":
   n.MiddleName = nil
  case "GivenName":
   n.GivenName = nil
  case "PreferredGivenName":
   n.PreferredGivenName = nil
  case "FamilyNameFirst":
   n.FamilyNameFirst = nil
  case "FullName":
   n.FullName = nil
  case "FamilyName":
   n.FamilyName = nil
  case "Title":
   n.Title = nil
  case "PreferredFamilyNameFirst":
   n.PreferredFamilyNameFirst = nil
  case "PreferredFamilyName":
   n.PreferredFamilyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "BaseNameType")
        }
        return n
}

func (n *BaseNameType) Set(key string, value interface{}) *BaseNameType {
        if n == nil {
                n = BaseNameTypeCreate(BaseNameType{})
        }
        switch key {
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
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
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Title":
                      n.Title = StringCreate(value.(string))
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
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "BaseNameType")
        }
        return n
}


func (t *AuditInfoType) CopyString(key string, value interface{}) *AuditInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AuditInfoType) Unset(key string) *AuditInfoType {
        switch key {
  case "CreationDateTime":
   n.CreationDateTime = nil
  case "CreationUser":
   n.CreationUser = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AuditInfoType")
        }
        return n
}

func (n *AuditInfoType) Set(key string, value interface{}) *AuditInfoType {
        if n == nil {
                n = AuditInfoTypeCreate(AuditInfoType{})
        }
        switch key {
    case "CreationDateTime":
                      n.CreationDateTime = StringCreate(value.(string))
    case "CreationUser":
                      n.CreationUser = CreationUserTypeCreate(value.(CreationUserType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AuditInfoType")
        }
        return n
}


func (t *StudentAttendanceSummary) CopyString(key string, value interface{}) *StudentAttendanceSummary {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentAttendanceSummary) Unset(key string) *StudentAttendanceSummary {
        switch key {
  case "FTE":
   n.FTE = nil
  case "EndDate":
   n.EndDate = nil
  case "DaysTardy":
   n.DaysTardy = nil
  case "StudentAttendanceSummaryRefId":
   n.StudentAttendanceSummaryRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "StartDay":
   n.StartDay = nil
  case "DaysAttended":
   n.DaysAttended = nil
  case "ExcusedAbsences":
   n.ExcusedAbsences = nil
  case "DaysInMembership":
   n.DaysInMembership = nil
  case "UnexcusedAbsences":
   n.UnexcusedAbsences = nil
  case "StartDate":
   n.StartDate = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "EndDay":
   n.EndDay = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceSummary")
        }
        return n
}

func (n *StudentAttendanceSummary) Set(key string, value interface{}) *StudentAttendanceSummary {
        if n == nil {
                n = StudentAttendanceSummaryCreate(StudentAttendanceSummary{})
        }
        switch key {
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "DaysTardy":
                      n.DaysTardy = FloatCreate(value.(float64))
    case "StudentAttendanceSummaryRefId":
                      n.StudentAttendanceSummaryRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "StartDay":
                      n.StartDay = IntCreate(value.(int))
    case "DaysAttended":
                      n.DaysAttended = FloatCreate(value.(float64))
    case "ExcusedAbsences":
                      n.ExcusedAbsences = FloatCreate(value.(float64))
    case "DaysInMembership":
                      n.DaysInMembership = FloatCreate(value.(float64))
    case "UnexcusedAbsences":
                      n.UnexcusedAbsences = FloatCreate(value.(float64))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "EndDay":
                      n.EndDay = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceSummary")
        }
        return n
}


func (t *NAPTestletResponseItemType) CopyString(key string, value interface{}) *NAPTestletResponseItemType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestletResponseItemType) Unset(key string) *NAPTestletResponseItemType {
        switch key {
  case "NAPTestItemLocalId":
   n.NAPTestItemLocalId = nil
  case "ResponseCorrectness":
   n.ResponseCorrectness = nil
  case "SequenceNumber":
   n.SequenceNumber = nil
  case "Score":
   n.Score = nil
  case "NAPTestItemRefId":
   n.NAPTestItemRefId = nil
  case "Response":
   n.Response = nil
  case "LapsedTimeItem":
   n.LapsedTimeItem = nil
  case "SubscoreList":
   n.SubscoreList = nil
  case "ItemWeight":
   n.ItemWeight = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletResponseItemType")
        }
        return n
}

func (n *NAPTestletResponseItemType) Set(key string, value interface{}) *NAPTestletResponseItemType {
        if n == nil {
                n = NAPTestletResponseItemTypeCreate(NAPTestletResponseItemType{})
        }
        switch key {
    case "NAPTestItemLocalId":
    
                      n.NAPTestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "Score":
                      n.Score = FloatCreate(value.(float64))
    case "NAPTestItemRefId":
                      n.NAPTestItemRefId = StringCreate(value.(string))
    case "Response":
                      n.Response = StringCreate(value.(string))
    case "LapsedTimeItem":
                      n.LapsedTimeItem = StringCreate(value.(string))
    case "SubscoreList":
                      n.SubscoreList = NAPSubscoreListTypeCreate(value.(NAPSubscoreListType))
    case "ItemWeight":
                      n.ItemWeight = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletResponseItemType")
        }
        return n
}


func (t *StudentScoreJudgementAgainstStandard) CopyString(key string, value interface{}) *StudentScoreJudgementAgainstStandard {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentScoreJudgementAgainstStandard) Unset(key string) *StudentScoreJudgementAgainstStandard {
        switch key {
  case "CurriculumNodeCode":
   n.CurriculumNodeCode = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TermInfoRefId":
   n.TermInfoRefId = nil
  case "RefId":
   n.RefId = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "CurriculumCode":
   n.CurriculumCode = nil
  case "Score":
   n.Score = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SpecialCircumstanceLocalCode":
   n.SpecialCircumstanceLocalCode = nil
  case "TeachingGroupRefId":
   n.TeachingGroupRefId = nil
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "LearningStandardList":
   n.LearningStandardList = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SchoolCommonwealthId":
   n.SchoolCommonwealthId = nil
  case "LocalTermCode":
   n.LocalTermCode = nil
  case "YearLevel":
   n.YearLevel = nil
  case "StudentStateProvinceId":
   n.StudentStateProvinceId = nil
  case "StudentLocalId":
   n.StudentLocalId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "ManagedPathwayLocalCode":
   n.ManagedPathwayLocalCode = nil
  case "ClassLocalId":
   n.ClassLocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentScoreJudgementAgainstStandard")
        }
        return n
}

func (n *StudentScoreJudgementAgainstStandard) Set(key string, value interface{}) *StudentScoreJudgementAgainstStandard {
        if n == nil {
                n = StudentScoreJudgementAgainstStandardCreate(StudentScoreJudgementAgainstStandard{})
        }
        switch key {
    case "CurriculumNodeCode":
    
                      n.CurriculumNodeCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "CurriculumCode":
    
                      n.CurriculumCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "Score":
                      n.Score = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SpecialCircumstanceLocalCode":
    
                      n.SpecialCircumstanceLocalCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolCommonwealthId":
                      n.SchoolCommonwealthId = StringCreate(value.(string))
    case "LocalTermCode":
    
                      n.LocalTermCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "StudentStateProvinceId":
    
                      n.StudentStateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "StudentLocalId":
    
                      n.StudentLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ManagedPathwayLocalCode":
    
                      n.ManagedPathwayLocalCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "ClassLocalId":
                      n.ClassLocalId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentScoreJudgementAgainstStandard")
        }
        return n
}


func (t *NameType) CopyString(key string, value interface{}) *NameType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NameType) Unset(key string) *NameType {
        switch key {
  case "Type":
   n.Type = nil
  case "Suffix":
   n.Suffix = nil
  case "MiddleName":
   n.MiddleName = nil
  case "GivenName":
   n.GivenName = nil
  case "PreferredGivenName":
   n.PreferredGivenName = nil
  case "FamilyNameFirst":
   n.FamilyNameFirst = nil
  case "FullName":
   n.FullName = nil
  case "FamilyName":
   n.FamilyName = nil
  case "Title":
   n.Title = nil
  case "PreferredFamilyNameFirst":
   n.PreferredFamilyNameFirst = nil
  case "PreferredFamilyName":
   n.PreferredFamilyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NameType")
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
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
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
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Title":
                      n.Title = StringCreate(value.(string))
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
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NameType")
        }
        return n
}


func (t *EnglishProficiencyType) CopyString(key string, value interface{}) *EnglishProficiencyType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EnglishProficiencyType) Unset(key string) *EnglishProficiencyType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EnglishProficiencyType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EnglishProficiencyType")
        }
        return n
}


func (t *DebitOrCreditAmountType) CopyString(key string, value interface{}) *DebitOrCreditAmountType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DebitOrCreditAmountType) Unset(key string) *DebitOrCreditAmountType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
  case "Currency":
   n.Currency = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DebitOrCreditAmountType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DebitOrCreditAmountType")
        }
        return n
}


func (t *CalendarDateInfoType) CopyString(key string, value interface{}) *CalendarDateInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CalendarDateInfoType) Unset(key string) *CalendarDateInfoType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarDateInfoType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarDateInfoType")
        }
        return n
}


func (t *AdjustmentContainerType) CopyString(key string, value interface{}) *AdjustmentContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AdjustmentContainerType) Unset(key string) *AdjustmentContainerType {
        switch key {
  case "BookletType":
   n.BookletType = nil
  case "PNPCodeList":
   n.PNPCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AdjustmentContainerType")
        }
        return n
}

func (n *AdjustmentContainerType) Set(key string, value interface{}) *AdjustmentContainerType {
        if n == nil {
                n = AdjustmentContainerTypeCreate(AdjustmentContainerType{})
        }
        switch key {
    case "BookletType":
                      n.BookletType = StringCreate(value.(string))
    case "PNPCodeList":
                      n.PNPCodeList = PNPCodeListTypeCreate(value.(PNPCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AdjustmentContainerType")
        }
        return n
}


func (t *StudentSchoolEnrollment_Advisor) CopyString(key string, value interface{}) *StudentSchoolEnrollment_Advisor {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSchoolEnrollment_Advisor) Unset(key string) *StudentSchoolEnrollment_Advisor {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Advisor")
        }
        return n
}

func (n *StudentSchoolEnrollment_Advisor) Set(key string, value interface{}) *StudentSchoolEnrollment_Advisor {
        if n == nil {
                n = StudentSchoolEnrollment_AdvisorCreate(StudentSchoolEnrollment_Advisor{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Advisor")
        }
        return n
}


func (t *StudentSectionEnrollment) CopyString(key string, value interface{}) *StudentSectionEnrollment {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSectionEnrollment) Unset(key string) *StudentSectionEnrollment {
        switch key {
  case "EntryDate":
   n.EntryDate = nil
  case "RefId":
   n.RefId = nil
  case "ExitDate":
   n.ExitDate = nil
  case "SectionInfoRefId":
   n.SectionInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSectionEnrollment")
        }
        return n
}

func (n *StudentSectionEnrollment) Set(key string, value interface{}) *StudentSectionEnrollment {
        if n == nil {
                n = StudentSectionEnrollmentCreate(StudentSectionEnrollment{})
        }
        switch key {
    case "EntryDate":
                      n.EntryDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ExitDate":
                      n.ExitDate = StringCreate(value.(string))
    case "SectionInfoRefId":
                      n.SectionInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSectionEnrollment")
        }
        return n
}


func (t *StandardHierarchyLevelType) CopyString(key string, value interface{}) *StandardHierarchyLevelType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StandardHierarchyLevelType) Unset(key string) *StandardHierarchyLevelType {
        switch key {
  case "Number":
   n.Number = nil
  case "Description":
   n.Description = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardHierarchyLevelType")
        }
        return n
}

func (n *StandardHierarchyLevelType) Set(key string, value interface{}) *StandardHierarchyLevelType {
        if n == nil {
                n = StandardHierarchyLevelTypeCreate(StandardHierarchyLevelType{})
        }
        switch key {
    case "Number":
                      n.Number = IntCreate(value.(int))
    case "Description":
                      n.Description = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardHierarchyLevelType")
        }
        return n
}


func (t *LifeCycleType_Creator) CopyString(key string, value interface{}) *LifeCycleType_Creator {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LifeCycleType_Creator) Unset(key string) *LifeCycleType_Creator {
        switch key {
  case "ID":
   n.ID = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Creator")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LifeCycleType_Creator")
        }
        return n
}


func (t *SubjectAreaType) CopyString(key string, value interface{}) *SubjectAreaType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SubjectAreaType) Unset(key string) *SubjectAreaType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SubjectAreaType")
        }
        return n
}

func (n *SubjectAreaType) Set(key string, value interface{}) *SubjectAreaType {
        if n == nil {
                n = SubjectAreaTypeCreate(SubjectAreaType{})
        }
        switch key {
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SubjectAreaType")
        }
        return n
}


func (t *EmailType) CopyString(key string, value interface{}) *EmailType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EmailType) Unset(key string) *EmailType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EmailType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EmailType")
        }
        return n
}


func (t *TeachingGroupTeacherType) CopyString(key string, value interface{}) *TeachingGroupTeacherType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeachingGroupTeacherType) Unset(key string) *TeachingGroupTeacherType {
        switch key {
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "Name":
   n.Name = nil
  case "Association":
   n.Association = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupTeacherType")
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
    case "Name":
                      n.Name = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "Association":
                      n.Association = StringCreate(value.(string))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupTeacherType")
        }
        return n
}


func (t *JournalAdjustmentType) CopyString(key string, value interface{}) *JournalAdjustmentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *JournalAdjustmentType) Unset(key string) *JournalAdjustmentType {
        switch key {
  case "CreditAccountCode":
   n.CreditAccountCode = nil
  case "DebitAccountCode":
   n.DebitAccountCode = nil
  case "CreditFinancialAccountRefId":
   n.CreditFinancialAccountRefId = nil
  case "GSTCodeReplacement":
   n.GSTCodeReplacement = nil
  case "DebitFinancialAccountRefId":
   n.DebitFinancialAccountRefId = nil
  case "GSTCodeOriginal":
   n.GSTCodeOriginal = nil
  case "LineAdjustmentAmount":
   n.LineAdjustmentAmount = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "JournalAdjustmentType")
        }
        return n
}

func (n *JournalAdjustmentType) Set(key string, value interface{}) *JournalAdjustmentType {
        if n == nil {
                n = JournalAdjustmentTypeCreate(JournalAdjustmentType{})
        }
        switch key {
    case "CreditAccountCode":
                      n.CreditAccountCode = StringCreate(value.(string))
    case "DebitAccountCode":
                      n.DebitAccountCode = StringCreate(value.(string))
    case "CreditFinancialAccountRefId":
                      n.CreditFinancialAccountRefId = StringCreate(value.(string))
    case "GSTCodeReplacement":
                      n.GSTCodeReplacement = StringCreate(value.(string))
    case "DebitFinancialAccountRefId":
                      n.DebitFinancialAccountRefId = StringCreate(value.(string))
    case "GSTCodeOriginal":
                      n.GSTCodeOriginal = StringCreate(value.(string))
    case "LineAdjustmentAmount":
                      n.LineAdjustmentAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "JournalAdjustmentType")
        }
        return n
}


func (t *StaffSubjectType) CopyString(key string, value interface{}) *StaffSubjectType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffSubjectType) Unset(key string) *StaffSubjectType {
        switch key {
  case "PreferenceNumber":
   n.PreferenceNumber = nil
  case "SubjectLocalId":
   n.SubjectLocalId = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffSubjectType")
        }
        return n
}

func (n *StaffSubjectType) Set(key string, value interface{}) *StaffSubjectType {
        if n == nil {
                n = StaffSubjectTypeCreate(StaffSubjectType{})
        }
        switch key {
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableSubjectRefId":
    
                      n.TimeTableSubjectRefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffSubjectType")
        }
        return n
}


func (t *StaffMostRecentContainerType) CopyString(key string, value interface{}) *StaffMostRecentContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffMostRecentContainerType) Unset(key string) *StaffMostRecentContainerType {
        switch key {
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "SchoolACARAId":
   n.SchoolACARAId = nil
  case "LocalCampusId":
   n.LocalCampusId = nil
  case "NAPLANClassList":
   n.NAPLANClassList = nil
  case "HomeGroup":
   n.HomeGroup = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffMostRecentContainerType")
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
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalCampusId":
    
                      n.LocalCampusId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPLANClassList":
                      n.NAPLANClassList = NAPLANClassListTypeCreate(value.(NAPLANClassListType))
    case "HomeGroup":
                      n.HomeGroup = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffMostRecentContainerType")
        }
        return n
}


func (t *LanguageBaseType) CopyString(key string, value interface{}) *LanguageBaseType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LanguageBaseType) Unset(key string) *LanguageBaseType {
        switch key {
  case "Dialect":
   n.Dialect = nil
  case "LanguageType":
   n.LanguageType = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LanguageBaseType")
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
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LanguageBaseType")
        }
        return n
}


func (t *ScheduledActivity) CopyString(key string, value interface{}) *ScheduledActivity {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ScheduledActivity) Unset(key string) *ScheduledActivity {
        switch key {
  case "RoomList":
   n.RoomList = nil
  case "PeriodId":
   n.PeriodId = nil
  case "ActivityType":
   n.ActivityType = nil
  case "ActivityName":
   n.ActivityName = nil
  case "ActivityComment":
   n.ActivityComment = nil
  case "TeacherList":
   n.TeacherList = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "Override":
   n.Override = nil
  case "TeachingGroupList":
   n.TeachingGroupList = nil
  case "TimeTableCellRefId":
   n.TimeTableCellRefId = nil
  case "CellType":
   n.CellType = nil
  case "Location":
   n.Location = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "AddressList":
   n.AddressList = nil
  case "DayId":
   n.DayId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ActivityDate":
   n.ActivityDate = nil
  case "FinishTime":
   n.FinishTime = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StudentList":
   n.StudentList = nil
  case "StartTime":
   n.StartTime = nil
  case "YearLevels":
   n.YearLevels = nil
  case "TimeTableRefId":
   n.TimeTableRefId = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScheduledActivity")
        }
        return n
}

func (n *ScheduledActivity) Set(key string, value interface{}) *ScheduledActivity {
        if n == nil {
                n = ScheduledActivityCreate(ScheduledActivity{})
        }
        switch key {
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "ActivityName":
                      n.ActivityName = StringCreate(value.(string))
    case "ActivityComment":
                      n.ActivityComment = StringCreate(value.(string))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "Override":
                      n.Override = ScheduledActivityOverrideTypeCreate(value.(ScheduledActivityOverrideType))
    case "TeachingGroupList":
                      n.TeachingGroupList = TeachingGroupListTypeCreate(value.(TeachingGroupListType))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "Location":
                      n.Location = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ActivityDate":
                      n.ActivityDate = StringCreate(value.(string))
    case "FinishTime":
                      n.FinishTime = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentList":
                      n.StudentList = StudentsTypeCreate(value.(StudentsType))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "TimeTableRefId":
                      n.TimeTableRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScheduledActivity")
        }
        return n
}


func (t *AssignmentScoreType) CopyString(key string, value interface{}) *AssignmentScoreType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AssignmentScoreType) Unset(key string) *AssignmentScoreType {
        switch key {
  case "GradingAssignmentScoreRefId":
   n.GradingAssignmentScoreRefId = nil
  case "Weight":
   n.Weight = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AssignmentScoreType")
        }
        return n
}

func (n *AssignmentScoreType) Set(key string, value interface{}) *AssignmentScoreType {
        if n == nil {
                n = AssignmentScoreTypeCreate(AssignmentScoreType{})
        }
        switch key {
    case "GradingAssignmentScoreRefId":
                      n.GradingAssignmentScoreRefId = StringCreate(value.(string))
    case "Weight":
                      n.Weight = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AssignmentScoreType")
        }
        return n
}


func (t *ScoreType) CopyString(key string, value interface{}) *ScoreType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ScoreType) Unset(key string) *ScoreType {
        switch key {
  case "ScoreDescriptionList":
   n.ScoreDescriptionList = nil
  case "MaxScoreValue":
   n.MaxScoreValue = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScoreType")
        }
        return n
}

func (n *ScoreType) Set(key string, value interface{}) *ScoreType {
        if n == nil {
                n = ScoreTypeCreate(ScoreType{})
        }
        switch key {
    case "ScoreDescriptionList":
                      n.ScoreDescriptionList = ScoreDescriptionListTypeCreate(value.(ScoreDescriptionListType))
    case "MaxScoreValue":
                      n.MaxScoreValue = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScoreType")
        }
        return n
}


func (t *EvaluationType) CopyString(key string, value interface{}) *EvaluationType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EvaluationType) Unset(key string) *EvaluationType {
        switch key {
  case "Date":
   n.Date = nil
  case "Name":
   n.Name = nil
  case "RefId":
   n.RefId = nil
  case "Description":
   n.Description = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EvaluationType")
        }
        return n
}

func (n *EvaluationType) Set(key string, value interface{}) *EvaluationType {
        if n == nil {
                n = EvaluationTypeCreate(EvaluationType{})
        }
        switch key {
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EvaluationType")
        }
        return n
}


func (t *NAPWritingRubricType) CopyString(key string, value interface{}) *NAPWritingRubricType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPWritingRubricType) Unset(key string) *NAPWritingRubricType {
        switch key {
  case "Descriptor":
   n.Descriptor = nil
  case "RubricType":
   n.RubricType = nil
  case "ScoreList":
   n.ScoreList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPWritingRubricType")
        }
        return n
}

func (n *NAPWritingRubricType) Set(key string, value interface{}) *NAPWritingRubricType {
        if n == nil {
                n = NAPWritingRubricTypeCreate(NAPWritingRubricType{})
        }
        switch key {
    case "Descriptor":
                      n.Descriptor = StringCreate(value.(string))
    case "RubricType":
                      n.RubricType = StringCreate(value.(string))
    case "ScoreList":
                      n.ScoreList = ScoreListTypeCreate(value.(ScoreListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPWritingRubricType")
        }
        return n
}


func (t *ResourcesType) CopyString(key string, value interface{}) *ResourcesType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourcesType) Unset(key string) *ResourcesType {
        switch key {
  case "Value":
   n.Value = nil
  case "ResourceType":
   n.ResourceType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourcesType")
        }
        return n
}

func (n *ResourcesType) Set(key string, value interface{}) *ResourcesType {
        if n == nil {
                n = ResourcesTypeCreate(ResourcesType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "ResourceType":
                      n.ResourceType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourcesType")
        }
        return n
}


func (t *MediumOfInstructionType) CopyString(key string, value interface{}) *MediumOfInstructionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MediumOfInstructionType) Unset(key string) *MediumOfInstructionType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MediumOfInstructionType")
        }
        return n
}

func (n *MediumOfInstructionType) Set(key string, value interface{}) *MediumOfInstructionType {
        if n == nil {
                n = MediumOfInstructionTypeCreate(MediumOfInstructionType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MediumOfInstructionType")
        }
        return n
}


func (t *HoldInfoType) CopyString(key string, value interface{}) *HoldInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *HoldInfoType) Unset(key string) *HoldInfoType {
        switch key {
  case "MadeAvailable":
   n.MadeAvailable = nil
  case "Type":
   n.Type = nil
  case "ReservationExpiry":
   n.ReservationExpiry = nil
  case "DateNeeded":
   n.DateNeeded = nil
  case "Expires":
   n.Expires = nil
  case "DatePlaced":
   n.DatePlaced = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "HoldInfoType")
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
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "ReservationExpiry":
                      n.ReservationExpiry = StringCreate(value.(string))
    case "DateNeeded":
                      n.DateNeeded = StringCreate(value.(string))
    case "Expires":
                      n.Expires = StringCreate(value.(string))
    case "DatePlaced":
                      n.DatePlaced = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "HoldInfoType")
        }
        return n
}


func (t *AddressCollection) CopyString(key string, value interface{}) *AddressCollection {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AddressCollection) Unset(key string) *AddressCollection {
        switch key {
  case "AddressCollectionYear":
   n.AddressCollectionYear = nil
  case "SoftwareVendorInfo":
   n.SoftwareVendorInfo = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "RoundCode":
   n.RoundCode = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AddressCollectionReportingList":
   n.AddressCollectionReportingList = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "ReportingAuthorityCommonwealthId":
   n.ReportingAuthorityCommonwealthId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollection")
        }
        return n
}

func (n *AddressCollection) Set(key string, value interface{}) *AddressCollection {
        if n == nil {
                n = AddressCollectionCreate(AddressCollection{})
        }
        switch key {
    case "AddressCollectionYear":
    
                      n.AddressCollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AddressCollectionReportingList":
                      n.AddressCollectionReportingList = AddressCollectionReportingListTypeCreate(value.(AddressCollectionReportingListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollection")
        }
        return n
}


func (t *CensusReportingType) CopyString(key string, value interface{}) *CensusReportingType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CensusReportingType) Unset(key string) *CensusReportingType {
        switch key {
  case "EntityLevel":
   n.EntityLevel = nil
  case "CensusStaffList":
   n.CensusStaffList = nil
  case "EntityContact":
   n.EntityContact = nil
  case "EntityName":
   n.EntityName = nil
  case "CensusStudentList":
   n.CensusStudentList = nil
  case "CommonwealthId":
   n.CommonwealthId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusReportingType")
        }
        return n
}

func (n *CensusReportingType) Set(key string, value interface{}) *CensusReportingType {
        if n == nil {
                n = CensusReportingTypeCreate(CensusReportingType{})
        }
        switch key {
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "CensusStaffList":
                      n.CensusStaffList = CensusStaffListTypeCreate(value.(CensusStaffListType))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "CensusStudentList":
                      n.CensusStudentList = CensusStudentListTypeCreate(value.(CensusStudentListType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusReportingType")
        }
        return n
}


func (t *TeachingGroupPeriodType) CopyString(key string, value interface{}) *TeachingGroupPeriodType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeachingGroupPeriodType) Unset(key string) *TeachingGroupPeriodType {
        switch key {
  case "StartTime":
   n.StartTime = nil
  case "CellType":
   n.CellType = nil
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "DayId":
   n.DayId = nil
  case "RoomNumber":
   n.RoomNumber = nil
  case "PeriodId":
   n.PeriodId = nil
  case "TimeTableCellRefId":
   n.TimeTableCellRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupPeriodType")
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
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableCellRefId":
                      n.TimeTableCellRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupPeriodType")
        }
        return n
}


func (t *AggregateStatisticFact) CopyString(key string, value interface{}) *AggregateStatisticFact {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AggregateStatisticFact) Unset(key string) *AggregateStatisticFact {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "Value":
   n.Value = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Characteristics":
   n.Characteristics = nil
  case "Excluded":
   n.Excluded = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "AggregateStatisticInfoRefId":
   n.AggregateStatisticInfoRefId = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticFact")
        }
        return n
}

func (n *AggregateStatisticFact) Set(key string, value interface{}) *AggregateStatisticFact {
        if n == nil {
                n = AggregateStatisticFactCreate(AggregateStatisticFact{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Value":
                      n.Value = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Characteristics":
                      n.Characteristics = CharacteristicsTypeCreate(value.(CharacteristicsType))
    case "Excluded":
                      n.Excluded = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AggregateStatisticInfoRefId":
                      n.AggregateStatisticInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticFact")
        }
        return n
}


func (t *LearningResource) CopyString(key string, value interface{}) *LearningResource {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LearningResource) Unset(key string) *LearningResource {
        switch key {
  case "LearningResourcePackageRefId":
   n.LearningResourcePackageRefId = nil
  case "Components":
   n.Components = nil
  case "Status":
   n.Status = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Author":
   n.Author = nil
  case "UseAgreement":
   n.UseAgreement = nil
  case "Description":
   n.Description = nil
  case "AgreementDate":
   n.AgreementDate = nil
  case "Name":
   n.Name = nil
  case "Location":
   n.Location = nil
  case "Evaluations":
   n.Evaluations = nil
  case "Contacts":
   n.Contacts = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SubjectAreas":
   n.SubjectAreas = nil
  case "Approvals":
   n.Approvals = nil
  case "LearningStandards":
   n.LearningStandards = nil
  case "RefId":
   n.RefId = nil
  case "YearLevels":
   n.YearLevels = nil
  case "MediaTypes":
   n.MediaTypes = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningResource")
        }
        return n
}

func (n *LearningResource) Set(key string, value interface{}) *LearningResource {
        if n == nil {
                n = LearningResourceCreate(LearningResource{})
        }
        switch key {
    case "LearningResourcePackageRefId":
                      n.LearningResourcePackageRefId = StringCreate(value.(string))
    case "Components":
                      n.Components = ComponentsTypeCreate(value.(ComponentsType))
    case "Status":
                      n.Status = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Author":
                      n.Author = StringCreate(value.(string))
    case "UseAgreement":
                      n.UseAgreement = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "AgreementDate":
                      n.AgreementDate = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "Location":
                      n.Location = LearningResource_LocationCreate(value.(LearningResource_Location))
    case "Evaluations":
                      n.Evaluations = EvaluationsTypeCreate(value.(EvaluationsType))
    case "Contacts":
                      n.Contacts = ContactsTypeCreate(value.(ContactsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SubjectAreas":
                      n.SubjectAreas = ACStrandAreaListTypeCreate(value.(ACStrandAreaListType))
    case "Approvals":
                      n.Approvals = ApprovalsTypeCreate(value.(ApprovalsType))
    case "LearningStandards":
                      n.LearningStandards = LearningStandardsTypeCreate(value.(LearningStandardsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "MediaTypes":
                      n.MediaTypes = MediaTypesTypeCreate(value.(MediaTypesType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningResource")
        }
        return n
}


func (t *StudentGroupType) CopyString(key string, value interface{}) *StudentGroupType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentGroupType) Unset(key string) *StudentGroupType {
        switch key {
  case "GroupLocalId":
   n.GroupLocalId = nil
  case "GroupCategory":
   n.GroupCategory = nil
  case "GroupDescription":
   n.GroupDescription = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentGroupType")
        }
        return n
}

func (n *StudentGroupType) Set(key string, value interface{}) *StudentGroupType {
        if n == nil {
                n = StudentGroupTypeCreate(StudentGroupType{})
        }
        switch key {
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
    case "GroupDescription":
                      n.GroupDescription = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentGroupType")
        }
        return n
}


func (t *ReligiousEventType) CopyString(key string, value interface{}) *ReligiousEventType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ReligiousEventType) Unset(key string) *ReligiousEventType {
        switch key {
  case "Type":
   n.Type = nil
  case "Date":
   n.Date = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReligiousEventType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReligiousEventType")
        }
        return n
}


func (t *LEAInfo) CopyString(key string, value interface{}) *LEAInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LEAInfo) Unset(key string) *LEAInfo {
        switch key {
  case "EducationAgencyType":
   n.EducationAgencyType = nil
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "LEAContactList":
   n.LEAContactList = nil
  case "LocalId":
   n.LocalId = nil
  case "JurisdictionLowerHouse":
   n.JurisdictionLowerHouse = nil
  case "OperationalStatus":
   n.OperationalStatus = nil
  case "LEAURL":
   n.LEAURL = nil
  case "AddressList":
   n.AddressList = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SLA":
   n.SLA = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "LEAName":
   n.LEAName = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LEAInfo")
        }
        return n
}

func (n *LEAInfo) Set(key string, value interface{}) *LEAInfo {
        if n == nil {
                n = LEAInfoCreate(LEAInfo{})
        }
        switch key {
    case "EducationAgencyType":
                      n.EducationAgencyType = AgencyTypeCreate(value.(AgencyType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LEAContactList":
                      n.LEAContactList = LEAContactListTypeCreate(value.(LEAContactListType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "JurisdictionLowerHouse":
                      n.JurisdictionLowerHouse = StringCreate(value.(string))
    case "OperationalStatus":
      present := false
  for _, b := range AUCodeSetsOperationalStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsOperationalStatusType_values")
      }

                      n.OperationalStatus = ((*OperationalStatusType)(StringCreate(value.(string))))
    case "LEAURL":
                      n.LEAURL = StringCreate(value.(string))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LEAName":
                      n.LEAName = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LEAInfo")
        }
        return n
}


func (t *AgencyType) CopyString(key string, value interface{}) *AgencyType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AgencyType) Unset(key string) *AgencyType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AgencyType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AgencyType")
        }
        return n
}


func (t *NAPTestItem2Type) CopyString(key string, value interface{}) *NAPTestItem2Type {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestItem2Type) Unset(key string) *NAPTestItem2Type {
        switch key {
  case "TestItemRefId":
   n.TestItemRefId = nil
  case "SequenceNumber":
   n.SequenceNumber = nil
  case "TestItemLocalId":
   n.TestItemLocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItem2Type")
        }
        return n
}

func (n *NAPTestItem2Type) Set(key string, value interface{}) *NAPTestItem2Type {
        if n == nil {
                n = NAPTestItem2TypeCreate(NAPTestItem2Type{})
        }
        switch key {
    case "TestItemRefId":
                      n.TestItemRefId = StringCreate(value.(string))
    case "SequenceNumber":
                      n.SequenceNumber = IntCreate(value.(int))
    case "TestItemLocalId":
    
                      n.TestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItem2Type")
        }
        return n
}


func (t *FQItemType) CopyString(key string, value interface{}) *FQItemType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FQItemType) Unset(key string) *FQItemType {
        switch key {
  case "SystemAmount":
   n.SystemAmount = nil
  case "TuitionAmount":
   n.TuitionAmount = nil
  case "FQItemCode":
   n.FQItemCode = nil
  case "BoardingAmount":
   n.BoardingAmount = nil
  case "DioceseAmount":
   n.DioceseAmount = nil
  case "FQComments":
   n.FQComments = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQItemType")
        }
        return n
}

func (n *FQItemType) Set(key string, value interface{}) *FQItemType {
        if n == nil {
                n = FQItemTypeCreate(FQItemType{})
        }
        switch key {
    case "SystemAmount":
                      n.SystemAmount = FloatCreate(value.(float64))
    case "TuitionAmount":
                      n.TuitionAmount = FloatCreate(value.(float64))
    case "FQItemCode":
                      n.FQItemCode = StringCreate(value.(string))
    case "BoardingAmount":
                      n.BoardingAmount = FloatCreate(value.(float64))
    case "DioceseAmount":
                      n.DioceseAmount = FloatCreate(value.(float64))
    case "FQComments":
                      n.FQComments = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQItemType")
        }
        return n
}


func (t *StatisticalAreaType) CopyString(key string, value interface{}) *StatisticalAreaType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StatisticalAreaType) Unset(key string) *StatisticalAreaType {
        switch key {
  case "Value":
   n.Value = nil
  case "SpatialUnitType":
   n.SpatialUnitType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatisticalAreaType")
        }
        return n
}

func (n *StatisticalAreaType) Set(key string, value interface{}) *StatisticalAreaType {
        if n == nil {
                n = StatisticalAreaTypeCreate(StatisticalAreaType{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SpatialUnitType":
                      n.SpatialUnitType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatisticalAreaType")
        }
        return n
}


func (t *NAPCodeFrame) CopyString(key string, value interface{}) *NAPCodeFrame {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPCodeFrame) Unset(key string) *NAPCodeFrame {
        switch key {
  case "NAPTestRefId":
   n.NAPTestRefId = nil
  case "RefId":
   n.RefId = nil
  case "TestContent":
   n.TestContent = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TestletList":
   n.TestletList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPCodeFrame")
        }
        return n
}

func (n *NAPCodeFrame) Set(key string, value interface{}) *NAPCodeFrame {
        if n == nil {
                n = NAPCodeFrameCreate(NAPCodeFrame{})
        }
        switch key {
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TestContent":
                      n.TestContent = NAPTestContentTypeCreate(value.(NAPTestContentType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TestletList":
                      n.TestletList = NAPCodeFrameTestletListTypeCreate(value.(NAPCodeFrameTestletListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPCodeFrame")
        }
        return n
}


func (t *GradeType) CopyString(key string, value interface{}) *GradeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *GradeType) Unset(key string) *GradeType {
        switch key {
  case "Numeric":
   n.Numeric = nil
  case "MarkInfoRefId":
   n.MarkInfoRefId = nil
  case "Letter":
   n.Letter = nil
  case "Percentage":
   n.Percentage = nil
  case "Narrative":
   n.Narrative = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradeType")
        }
        return n
}

func (n *GradeType) Set(key string, value interface{}) *GradeType {
        if n == nil {
                n = GradeTypeCreate(GradeType{})
        }
        switch key {
    case "Numeric":
                      n.Numeric = FloatCreate(value.(float64))
    case "MarkInfoRefId":
                      n.MarkInfoRefId = StringCreate(value.(string))
    case "Letter":
                      n.Letter = StringCreate(value.(string))
    case "Percentage":
                      n.Percentage = FloatCreate(value.(float64))
    case "Narrative":
                      n.Narrative = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradeType")
        }
        return n
}


func (t *LanguageOfInstructionType) CopyString(key string, value interface{}) *LanguageOfInstructionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LanguageOfInstructionType) Unset(key string) *LanguageOfInstructionType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LanguageOfInstructionType")
        }
        return n
}

func (n *LanguageOfInstructionType) Set(key string, value interface{}) *LanguageOfInstructionType {
        if n == nil {
                n = LanguageOfInstructionTypeCreate(LanguageOfInstructionType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LanguageOfInstructionType")
        }
        return n
}


func (t *ActivityTimeType) CopyString(key string, value interface{}) *ActivityTimeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ActivityTimeType) Unset(key string) *ActivityTimeType {
        switch key {
  case "CreationDate":
   n.CreationDate = nil
  case "Duration":
   n.Duration = nil
  case "DueDate":
   n.DueDate = nil
  case "StartDate":
   n.StartDate = nil
  case "FinishDate":
   n.FinishDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ActivityTimeType")
        }
        return n
}

func (n *ActivityTimeType) Set(key string, value interface{}) *ActivityTimeType {
        if n == nil {
                n = ActivityTimeTypeCreate(ActivityTimeType{})
        }
        switch key {
    case "CreationDate":
                      n.CreationDate = StringCreate(value.(string))
    case "Duration":
                      n.Duration = ActivityTimeType_DurationCreate(value.(ActivityTimeType_Duration))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "FinishDate":
                      n.FinishDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ActivityTimeType")
        }
        return n
}


func (t *SchoolInfo) CopyString(key string, value interface{}) *SchoolInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolInfo) Unset(key string) *SchoolInfo {
        switch key {
  case "SchoolDistrict":
   n.SchoolDistrict = nil
  case "SchoolGeographicLocation":
   n.SchoolGeographicLocation = nil
  case "SchoolDistrictLocalId":
   n.SchoolDistrictLocalId = nil
  case "Campus":
   n.Campus = nil
  case "ARIA":
   n.ARIA = nil
  case "TotalEnrollments":
   n.TotalEnrollments = nil
  case "BoardingSchoolStatus":
   n.BoardingSchoolStatus = nil
  case "OperationalStatus":
   n.OperationalStatus = nil
  case "SessionType":
   n.SessionType = nil
  case "RefId":
   n.RefId = nil
  case "LEAInfoRefId":
   n.LEAInfoRefId = nil
  case "SchoolCoEdStatus":
   n.SchoolCoEdStatus = nil
  case "System":
   n.System = nil
  case "SchoolGroupList":
   n.SchoolGroupList = nil
  case "ReligiousAffiliation":
   n.ReligiousAffiliation = nil
  case "SchoolEmailList":
   n.SchoolEmailList = nil
  case "Entity_Open":
   n.Entity_Open = nil
  case "LocalId":
   n.LocalId = nil
  case "NonGovSystemicStatus":
   n.NonGovSystemicStatus = nil
  case "SchoolTimeZone":
   n.SchoolTimeZone = nil
  case "SchoolFocusList":
   n.SchoolFocusList = nil
  case "OtherIdList":
   n.OtherIdList = nil
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "SchoolType":
   n.SchoolType = nil
  case "PrincipalInfo":
   n.PrincipalInfo = nil
  case "YearLevels":
   n.YearLevels = nil
  case "FederalElectorate":
   n.FederalElectorate = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "IndependentSchool":
   n.IndependentSchool = nil
  case "SchoolURL":
   n.SchoolURL = nil
  case "SchoolContactList":
   n.SchoolContactList = nil
  case "LocalGovernmentArea":
   n.LocalGovernmentArea = nil
  case "Entity_Close":
   n.Entity_Close = nil
  case "ACARAId":
   n.ACARAId = nil
  case "SchoolName":
   n.SchoolName = nil
  case "YearLevelEnrollmentList":
   n.YearLevelEnrollmentList = nil
  case "SLA":
   n.SLA = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AddressList":
   n.AddressList = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "SchoolSector":
   n.SchoolSector = nil
  case "JurisdictionLowerHouse":
   n.JurisdictionLowerHouse = nil
  case "OtherLEA":
   n.OtherLEA = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolInfo")
        }
        return n
}

func (n *SchoolInfo) Set(key string, value interface{}) *SchoolInfo {
        if n == nil {
                n = SchoolInfoCreate(SchoolInfo{})
        }
        switch key {
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
    case "SchoolDistrictLocalId":
    
                      n.SchoolDistrictLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Campus":
                      n.Campus = CampusContainerTypeCreate(value.(CampusContainerType))
    case "ARIA":
                      n.ARIA = FloatCreate(value.(float64))
    case "TotalEnrollments":
                      n.TotalEnrollments = TotalEnrollmentsTypeCreate(value.(TotalEnrollmentsType))
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
    case "OperationalStatus":
      present := false
  for _, b := range AUCodeSetsOperationalStatusType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsOperationalStatusType_values")
      }

                      n.OperationalStatus = ((*OperationalStatusType)(StringCreate(value.(string))))
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
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LEAInfoRefId":
    
                      n.LEAInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
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
    case "SchoolGroupList":
                      n.SchoolGroupList = SchoolGroupListTypeCreate(value.(SchoolGroupListType))
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
    case "SchoolEmailList":
                      n.SchoolEmailList = EmailListTypeCreate(value.(EmailListType))
    case "Entity_Open":
                      n.Entity_Open = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "SchoolFocusList":
                      n.SchoolFocusList = SchoolFocusListTypeCreate(value.(SchoolFocusListType))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
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
    case "PrincipalInfo":
                      n.PrincipalInfo = PrincipalInfoTypeCreate(value.(PrincipalInfoType))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
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
    case "SchoolURL":
    
                      n.SchoolURL = ((*SchoolURLType)(StringCreate(value.(string))))
    case "SchoolContactList":
                      n.SchoolContactList = SchoolContactListTypeCreate(value.(SchoolContactListType))
    case "LocalGovernmentArea":
                      n.LocalGovernmentArea = StringCreate(value.(string))
    case "Entity_Close":
                      n.Entity_Close = StringCreate(value.(string))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
    case "YearLevelEnrollmentList":
                      n.YearLevelEnrollmentList = YearLevelEnrollmentListTypeCreate(value.(YearLevelEnrollmentListType))
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
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
    case "JurisdictionLowerHouse":
                      n.JurisdictionLowerHouse = StringCreate(value.(string))
    case "OtherLEA":
                      n.OtherLEA = SchoolInfo_OtherLEACreate(value.(SchoolInfo_OtherLEA))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolInfo")
        }
        return n
}


func (t *StaffPersonal) CopyString(key string, value interface{}) *StaffPersonal {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffPersonal) Unset(key string) *StaffPersonal {
        switch key {
  case "RefId":
   n.RefId = nil
  case "MostRecent":
   n.MostRecent = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Title":
   n.Title = nil
  case "PersonInfo":
   n.PersonInfo = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "LocalId":
   n.LocalId = nil
  case "EmploymentStatus":
   n.EmploymentStatus = nil
  case "ElectronicIdList":
   n.ElectronicIdList = nil
  case "OtherIdList":
   n.OtherIdList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffPersonal")
        }
        return n
}

func (n *StaffPersonal) Set(key string, value interface{}) *StaffPersonal {
        if n == nil {
                n = StaffPersonalCreate(StaffPersonal{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "MostRecent":
                      n.MostRecent = StaffMostRecentContainerTypeCreate(value.(StaffMostRecentContainerType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffPersonal")
        }
        return n
}


func (t *PrincipalInfoType) CopyString(key string, value interface{}) *PrincipalInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PrincipalInfoType) Unset(key string) *PrincipalInfoType {
        switch key {
  case "ContactName":
   n.ContactName = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "ContactTitle":
   n.ContactTitle = nil
  case "EmailList":
   n.EmailList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PrincipalInfoType")
        }
        return n
}

func (n *PrincipalInfoType) Set(key string, value interface{}) *PrincipalInfoType {
        if n == nil {
                n = PrincipalInfoTypeCreate(PrincipalInfoType{})
        }
        switch key {
    case "ContactName":
                      n.ContactName = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "ContactTitle":
                      n.ContactTitle = StringCreate(value.(string))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PrincipalInfoType")
        }
        return n
}


func (t *NAPTestletResponseType) CopyString(key string, value interface{}) *NAPTestletResponseType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestletResponseType) Unset(key string) *NAPTestletResponseType {
        switch key {
  case "NAPTestletRefId":
   n.NAPTestletRefId = nil
  case "ItemResponseList":
   n.ItemResponseList = nil
  case "TestletSubScore":
   n.TestletSubScore = nil
  case "NAPTestletLocalId":
   n.NAPTestletLocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletResponseType")
        }
        return n
}

func (n *NAPTestletResponseType) Set(key string, value interface{}) *NAPTestletResponseType {
        if n == nil {
                n = NAPTestletResponseTypeCreate(NAPTestletResponseType{})
        }
        switch key {
    case "NAPTestletRefId":
                      n.NAPTestletRefId = StringCreate(value.(string))
    case "ItemResponseList":
                      n.ItemResponseList = NAPTestletItemResponseListTypeCreate(value.(NAPTestletItemResponseListType))
    case "TestletSubScore":
                      n.TestletSubScore = FloatCreate(value.(float64))
    case "NAPTestletLocalId":
    
                      n.NAPTestletLocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletResponseType")
        }
        return n
}


func (t *AGParentType) CopyString(key string, value interface{}) *AGParentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AGParentType) Unset(key string) *AGParentType {
        switch key {
  case "ParentAddress":
   n.ParentAddress = nil
  case "ParentName":
   n.ParentName = nil
  case "AddressSameAsStudent":
   n.AddressSameAsStudent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGParentType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGParentType")
        }
        return n
}


func (t *WellbeingEventLocationDetailsType) CopyString(key string, value interface{}) *WellbeingEventLocationDetailsType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingEventLocationDetailsType) Unset(key string) *WellbeingEventLocationDetailsType {
        switch key {
  case "EventLocation":
   n.EventLocation = nil
  case "FurtherLocationNotes":
   n.FurtherLocationNotes = nil
  case "Class":
   n.Class = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEventLocationDetailsType")
        }
        return n
}

func (n *WellbeingEventLocationDetailsType) Set(key string, value interface{}) *WellbeingEventLocationDetailsType {
        if n == nil {
                n = WellbeingEventLocationDetailsTypeCreate(WellbeingEventLocationDetailsType{})
        }
        switch key {
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
    case "Class":
                      n.Class = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEventLocationDetailsType")
        }
        return n
}


func (t *TimeTableCell) CopyString(key string, value interface{}) *TimeTableCell {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableCell) Unset(key string) *TimeTableCell {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SubjectLocalId":
   n.SubjectLocalId = nil
  case "TeachingGroupLocalId":
   n.TeachingGroupLocalId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "DayId":
   n.DayId = nil
  case "RefId":
   n.RefId = nil
  case "TimeTableRefId":
   n.TimeTableRefId = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "TimeTableLocalId":
   n.TimeTableLocalId = nil
  case "RoomInfoRefId":
   n.RoomInfoRefId = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "PeriodId":
   n.PeriodId = nil
  case "RoomList":
   n.RoomList = nil
  case "TeacherList":
   n.TeacherList = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "TeachingGroupRefId":
   n.TeachingGroupRefId = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "RoomNumber":
   n.RoomNumber = nil
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "CellType":
   n.CellType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableCell")
        }
        return n
}

func (n *TimeTableCell) Set(key string, value interface{}) *TimeTableCell {
        if n == nil {
                n = TimeTableCellCreate(TimeTableCell{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TimeTableRefId":
                      n.TimeTableRefId = StringCreate(value.(string))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "TimeTableLocalId":
    
                      n.TimeTableLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomInfoRefId":
                      n.RoomInfoRefId = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableCell")
        }
        return n
}


func (t *SchoolContactType) CopyString(key string, value interface{}) *SchoolContactType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolContactType) Unset(key string) *SchoolContactType {
        switch key {
  case "PublishInDirectory":
   n.PublishInDirectory = nil
  case "ContactInfo":
   n.ContactInfo = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolContactType")
        }
        return n
}

func (n *SchoolContactType) Set(key string, value interface{}) *SchoolContactType {
        if n == nil {
                n = SchoolContactTypeCreate(SchoolContactType{})
        }
        switch key {
    case "PublishInDirectory":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PublishInDirectory = ((*PublishInDirectoryType)(StringCreate(value.(string))))
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolContactType")
        }
        return n
}


func (t *SystemRole) CopyString(key string, value interface{}) *SystemRole {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole) Unset(key string) *SystemRole {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SystemContextList":
   n.SystemContextList = nil
  case "RefId":
   n.RefId = nil
  case "SIF_RefId":
   n.SIF_RefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole")
        }
        return n
}

func (n *SystemRole) Set(key string, value interface{}) *SystemRole {
        if n == nil {
                n = SystemRoleCreate(SystemRole{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SystemContextList":
                      n.SystemContextList = SystemRole_SystemContextListCreate(value.(SystemRole_SystemContextList))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_RefId":
                      n.SIF_RefId = SystemRole_SIF_RefIdCreate(value.(SystemRole_SIF_RefId))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole")
        }
        return n
}


func (t *SchoolCourseInfo) CopyString(key string, value interface{}) *SchoolCourseInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolCourseInfo) Unset(key string) *SchoolCourseInfo {
        switch key {
  case "CourseTitle":
   n.CourseTitle = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "Description":
   n.Description = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "DistrictCourseCode":
   n.DistrictCourseCode = nil
  case "CourseCredits":
   n.CourseCredits = nil
  case "InstructionalLevel":
   n.InstructionalLevel = nil
  case "Department":
   n.Department = nil
  case "CourseCode":
   n.CourseCode = nil
  case "SubjectAreaList":
   n.SubjectAreaList = nil
  case "CoreAcademicCourse":
   n.CoreAcademicCourse = nil
  case "CourseContent":
   n.CourseContent = nil
  case "StateCourseCode":
   n.StateCourseCode = nil
  case "RefId":
   n.RefId = nil
  case "TermInfoRefId":
   n.TermInfoRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "GraduationRequirement":
   n.GraduationRequirement = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolCourseInfo")
        }
        return n
}

func (n *SchoolCourseInfo) Set(key string, value interface{}) *SchoolCourseInfo {
        if n == nil {
                n = SchoolCourseInfoCreate(SchoolCourseInfo{})
        }
        switch key {
    case "CourseTitle":
                      n.CourseTitle = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DistrictCourseCode":
                      n.DistrictCourseCode = StringCreate(value.(string))
    case "CourseCredits":
                      n.CourseCredits = StringCreate(value.(string))
    case "InstructionalLevel":
                      n.InstructionalLevel = StringCreate(value.(string))
    case "Department":
                      n.Department = StringCreate(value.(string))
    case "CourseCode":
                      n.CourseCode = StringCreate(value.(string))
    case "SubjectAreaList":
                      n.SubjectAreaList = SubjectAreaListTypeCreate(value.(SubjectAreaListType))
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
    case "CourseContent":
                      n.CourseContent = StringCreate(value.(string))
    case "StateCourseCode":
                      n.StateCourseCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolCourseInfo")
        }
        return n
}


func (t *SystemRole_Role) CopyString(key string, value interface{}) *SystemRole_Role {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole_Role) Unset(key string) *SystemRole_Role {
        switch key {
  case "RoleScopeList":
   n.RoleScopeList = nil
  case "RoleId":
   n.RoleId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_Role")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_Role")
        }
        return n
}


func (t *ScheduledActivityOverrideType) CopyString(key string, value interface{}) *ScheduledActivityOverrideType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ScheduledActivityOverrideType) Unset(key string) *ScheduledActivityOverrideType {
        switch key {
  case "Value":
   n.Value = nil
  case "DateOfOverride":
   n.DateOfOverride = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScheduledActivityOverrideType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ScheduledActivityOverrideType")
        }
        return n
}


func (t *ContactType) CopyString(key string, value interface{}) *ContactType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ContactType) Unset(key string) *ContactType {
        switch key {
  case "Email":
   n.Email = nil
  case "PhoneNumber":
   n.PhoneNumber = nil
  case "Address":
   n.Address = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactType")
        }
        return n
}

func (n *ContactType) Set(key string, value interface{}) *ContactType {
        if n == nil {
                n = ContactTypeCreate(ContactType{})
        }
        switch key {
    case "Email":
                      n.Email = EmailTypeCreate(value.(EmailType))
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactType")
        }
        return n
}


func (t *Debtor_BilledEntity) CopyString(key string, value interface{}) *Debtor_BilledEntity {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Debtor_BilledEntity) Unset(key string) *Debtor_BilledEntity {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Debtor_BilledEntity")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Debtor_BilledEntity")
        }
        return n
}


func (t *TimeElementType_SpanGap) CopyString(key string, value interface{}) *TimeElementType_SpanGap {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeElementType_SpanGap) Unset(key string) *TimeElementType_SpanGap {
        switch key {
  case "Name":
   n.Name = nil
  case "StartDateTime":
   n.StartDateTime = nil
  case "EndDateTime":
   n.EndDateTime = nil
  case "Code":
   n.Code = nil
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeElementType_SpanGap")
        }
        return n
}

func (n *TimeElementType_SpanGap) Set(key string, value interface{}) *TimeElementType_SpanGap {
        if n == nil {
                n = TimeElementType_SpanGapCreate(TimeElementType_SpanGap{})
        }
        switch key {
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "EndDateTime":
                      n.EndDateTime = StringCreate(value.(string))
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeElementType_SpanGap")
        }
        return n
}


func (t *PersonalisedPlan) CopyString(key string, value interface{}) *PersonalisedPlan {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonalisedPlan) Unset(key string) *PersonalisedPlan {
        switch key {
  case "PersonalisedPlanCategory":
   n.PersonalisedPlanCategory = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "LocalId":
   n.LocalId = nil
  case "DocumentList":
   n.DocumentList = nil
  case "PersonalisedPlanEndDate":
   n.PersonalisedPlanEndDate = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "PersonalisedPlanNotes":
   n.PersonalisedPlanNotes = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "PersonalisedPlanStartDate":
   n.PersonalisedPlanStartDate = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "PersonalisedPlanReviewDate":
   n.PersonalisedPlanReviewDate = nil
  case "AssociatedAttachment":
   n.AssociatedAttachment = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonalisedPlan")
        }
        return n
}

func (n *PersonalisedPlan) Set(key string, value interface{}) *PersonalisedPlan {
        if n == nil {
                n = PersonalisedPlanCreate(PersonalisedPlan{})
        }
        switch key {
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
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "PersonalisedPlanEndDate":
                      n.PersonalisedPlanEndDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PersonalisedPlanNotes":
                      n.PersonalisedPlanNotes = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PersonalisedPlanStartDate":
                      n.PersonalisedPlanStartDate = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PersonalisedPlanReviewDate":
                      n.PersonalisedPlanReviewDate = StringCreate(value.(string))
    case "AssociatedAttachment":
                      n.AssociatedAttachment = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonalisedPlan")
        }
        return n
}


func (t *PersonPicture_PictureSource) CopyString(key string, value interface{}) *PersonPicture_PictureSource {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonPicture_PictureSource) Unset(key string) *PersonPicture_PictureSource {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture_PictureSource")
        }
        return n
}

func (n *PersonPicture_PictureSource) Set(key string, value interface{}) *PersonPicture_PictureSource {
        if n == nil {
                n = PersonPicture_PictureSourceCreate(PersonPicture_PictureSource{})
        }
        switch key {
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
    case "Value":
    
                      n.Value = ((*URIOrBinaryType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture_PictureSource")
        }
        return n
}


func (t *LearningStandardType) CopyString(key string, value interface{}) *LearningStandardType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LearningStandardType) Unset(key string) *LearningStandardType {
        switch key {
  case "LearningStandardLocalId":
   n.LearningStandardLocalId = nil
  case "LearningStandardItemRefId":
   n.LearningStandardItemRefId = nil
  case "LearningStandardURL":
   n.LearningStandardURL = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardType")
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
    case "LearningStandardItemRefId":
                      n.LearningStandardItemRefId = StringCreate(value.(string))
    case "LearningStandardURL":
                      n.LearningStandardURL = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardType")
        }
        return n
}


func (t *AGRuleType) CopyString(key string, value interface{}) *AGRuleType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AGRuleType) Unset(key string) *AGRuleType {
        switch key {
  case "AGRuleStatus":
   n.AGRuleStatus = nil
  case "AGRuleComment":
   n.AGRuleComment = nil
  case "AGRuleResponse":
   n.AGRuleResponse = nil
  case "AGRuleCode":
   n.AGRuleCode = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGRuleType")
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
    case "AGRuleComment":
                      n.AGRuleComment = StringCreate(value.(string))
    case "AGRuleResponse":
                      n.AGRuleResponse = StringCreate(value.(string))
    case "AGRuleCode":
                      n.AGRuleCode = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGRuleType")
        }
        return n
}


func (t *NameOfRecordType) CopyString(key string, value interface{}) *NameOfRecordType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NameOfRecordType) Unset(key string) *NameOfRecordType {
        switch key {
  case "Type":
   n.Type = nil
  case "Suffix":
   n.Suffix = nil
  case "MiddleName":
   n.MiddleName = nil
  case "GivenName":
   n.GivenName = nil
  case "PreferredGivenName":
   n.PreferredGivenName = nil
  case "FamilyNameFirst":
   n.FamilyNameFirst = nil
  case "FullName":
   n.FullName = nil
  case "FamilyName":
   n.FamilyName = nil
  case "Title":
   n.Title = nil
  case "PreferredFamilyNameFirst":
   n.PreferredFamilyNameFirst = nil
  case "PreferredFamilyName":
   n.PreferredFamilyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NameOfRecordType")
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
    case "Suffix":
                      n.Suffix = StringCreate(value.(string))
    case "MiddleName":
                      n.MiddleName = StringCreate(value.(string))
    case "GivenName":
                      n.GivenName = StringCreate(value.(string))
    case "PreferredGivenName":
                      n.PreferredGivenName = StringCreate(value.(string))
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
    case "FullName":
                      n.FullName = StringCreate(value.(string))
    case "FamilyName":
                      n.FamilyName = StringCreate(value.(string))
    case "Title":
                      n.Title = StringCreate(value.(string))
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
    case "PreferredFamilyName":
                      n.PreferredFamilyName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NameOfRecordType")
        }
        return n
}


func (t *AbstractContentPackageType_BinaryData) CopyString(key string, value interface{}) *AbstractContentPackageType_BinaryData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentPackageType_BinaryData) Unset(key string) *AbstractContentPackageType_BinaryData {
        switch key {
  case "Description":
   n.Description = nil
  case "FileName":
   n.FileName = nil
  case "Value":
   n.Value = nil
  case "MIMEType":
   n.MIMEType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_BinaryData")
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
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_BinaryData")
        }
        return n
}


func (t *RelatedLearningStandardItemRefIdType) CopyString(key string, value interface{}) *RelatedLearningStandardItemRefIdType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *RelatedLearningStandardItemRefIdType) Unset(key string) *RelatedLearningStandardItemRefIdType {
        switch key {
  case "RelationshipType":
   n.RelationshipType = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RelatedLearningStandardItemRefIdType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RelatedLearningStandardItemRefIdType")
        }
        return n
}


func (t *ContactInfoType) CopyString(key string, value interface{}) *ContactInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ContactInfoType) Unset(key string) *ContactInfoType {
        switch key {
  case "RegistrationDetails":
   n.RegistrationDetails = nil
  case "EmailList":
   n.EmailList = nil
  case "PositionTitle":
   n.PositionTitle = nil
  case "Qualifications":
   n.Qualifications = nil
  case "Address":
   n.Address = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "Role":
   n.Role = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactInfoType")
        }
        return n
}

func (n *ContactInfoType) Set(key string, value interface{}) *ContactInfoType {
        if n == nil {
                n = ContactInfoTypeCreate(ContactInfoType{})
        }
        switch key {
    case "RegistrationDetails":
                      n.RegistrationDetails = StringCreate(value.(string))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "PositionTitle":
                      n.PositionTitle = StringCreate(value.(string))
    case "Qualifications":
                      n.Qualifications = StringCreate(value.(string))
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ContactInfoType")
        }
        return n
}


func (t *TeacherCoverType) CopyString(key string, value interface{}) *TeacherCoverType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeacherCoverType) Unset(key string) *TeacherCoverType {
        switch key {
  case "StartTime":
   n.StartTime = nil
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "Credit":
   n.Credit = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "Weighting":
   n.Weighting = nil
  case "Supervision":
   n.Supervision = nil
  case "FinishTime":
   n.FinishTime = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeacherCoverType")
        }
        return n
}

func (n *TeacherCoverType) Set(key string, value interface{}) *TeacherCoverType {
        if n == nil {
                n = TeacherCoverTypeCreate(TeacherCoverType{})
        }
        switch key {
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "Weighting":
                      n.Weighting = FloatCreate(value.(float64))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeacherCoverType")
        }
        return n
}


func (t *AbstractContentPackageType_TextData) CopyString(key string, value interface{}) *AbstractContentPackageType_TextData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentPackageType_TextData) Unset(key string) *AbstractContentPackageType_TextData {
        switch key {
  case "Description":
   n.Description = nil
  case "FileName":
   n.FileName = nil
  case "Value":
   n.Value = nil
  case "MIMEType":
   n.MIMEType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_TextData")
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
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_TextData")
        }
        return n
}


func (t *TestDisruptionType) CopyString(key string, value interface{}) *TestDisruptionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TestDisruptionType) Unset(key string) *TestDisruptionType {
        switch key {
  case "Event":
   n.Event = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TestDisruptionType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TestDisruptionType")
        }
        return n
}


func (t *TimeTableScheduleCellType) CopyString(key string, value interface{}) *TimeTableScheduleCellType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableScheduleCellType) Unset(key string) *TimeTableScheduleCellType {
        switch key {
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "StaffLocalId":
   n.StaffLocalId = nil
  case "RoomNumber":
   n.RoomNumber = nil
  case "CellType":
   n.CellType = nil
  case "TimeTableScheduleCellLocalId":
   n.TimeTableScheduleCellLocalId = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "RoomInfoRefId":
   n.RoomInfoRefId = nil
  case "RoomList":
   n.RoomList = nil
  case "PeriodId":
   n.PeriodId = nil
  case "TeacherList":
   n.TeacherList = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "TimeTableLocalId":
   n.TimeTableLocalId = nil
  case "TeachingGroupGUID":
   n.TeachingGroupGUID = nil
  case "TeachingGroupLocalId":
   n.TeachingGroupLocalId = nil
  case "SubjectLocalId":
   n.SubjectLocalId = nil
  case "DayId":
   n.DayId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableScheduleCellType")
        }
        return n
}

func (n *TimeTableScheduleCellType) Set(key string, value interface{}) *TimeTableScheduleCellType {
        if n == nil {
                n = TimeTableScheduleCellTypeCreate(TimeTableScheduleCellType{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StaffLocalId":
    
                      n.StaffLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "RoomNumber":
    
                      n.RoomNumber = ((*HomeroomNumberType)(StringCreate(value.(string))))
    case "CellType":
                      n.CellType = StringCreate(value.(string))
    case "TimeTableScheduleCellLocalId":
    
                      n.TimeTableScheduleCellLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "RoomInfoRefId":
                      n.RoomInfoRefId = StringCreate(value.(string))
    case "RoomList":
                      n.RoomList = RoomListTypeCreate(value.(RoomListType))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeacherList":
                      n.TeacherList = ScheduledTeacherListTypeCreate(value.(ScheduledTeacherListType))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "TimeTableLocalId":
    
                      n.TimeTableLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingGroupGUID":
                      n.TeachingGroupGUID = StringCreate(value.(string))
    case "TeachingGroupLocalId":
    
                      n.TeachingGroupLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DayId":
    
                      n.DayId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableScheduleCellType")
        }
        return n
}


func (t *LEAContactType) CopyString(key string, value interface{}) *LEAContactType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LEAContactType) Unset(key string) *LEAContactType {
        switch key {
  case "ContactInfo":
   n.ContactInfo = nil
  case "PublishInDirectory":
   n.PublishInDirectory = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LEAContactType")
        }
        return n
}

func (n *LEAContactType) Set(key string, value interface{}) *LEAContactType {
        if n == nil {
                n = LEAContactTypeCreate(LEAContactType{})
        }
        switch key {
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
    case "PublishInDirectory":
      present := false
  for _, b := range AUCodeSetsYesOrNoCategoryType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsYesOrNoCategoryType_values")
      }

                      n.PublishInDirectory = ((*PublishInDirectoryType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LEAContactType")
        }
        return n
}


func (t *ResourceBooking_ResourceRefId) CopyString(key string, value interface{}) *ResourceBooking_ResourceRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceBooking_ResourceRefId) Unset(key string) *ResourceBooking_ResourceRefId {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceBooking_ResourceRefId")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceBooking_ResourceRefId")
        }
        return n
}


func (t *PersonPicture_ParentObjectRefId) CopyString(key string, value interface{}) *PersonPicture_ParentObjectRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonPicture_ParentObjectRefId) Unset(key string) *PersonPicture_ParentObjectRefId {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture_ParentObjectRefId")
        }
        return n
}

func (n *PersonPicture_ParentObjectRefId) Set(key string, value interface{}) *PersonPicture_ParentObjectRefId {
        if n == nil {
                n = PersonPicture_ParentObjectRefIdCreate(PersonPicture_ParentObjectRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonPicture_ParentObjectRefId")
        }
        return n
}


func (t *StudentSchoolEnrollment_Homeroom) CopyString(key string, value interface{}) *StudentSchoolEnrollment_Homeroom {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSchoolEnrollment_Homeroom) Unset(key string) *StudentSchoolEnrollment_Homeroom {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Homeroom")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSchoolEnrollment_Homeroom")
        }
        return n
}


func (t *AGReportingObjectResponseType) CopyString(key string, value interface{}) *AGReportingObjectResponseType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AGReportingObjectResponseType) Unset(key string) *AGReportingObjectResponseType {
        switch key {
  case "SIFRefId":
   n.SIFRefId = nil
  case "SubmittedRefId":
   n.SubmittedRefId = nil
  case "AGSubmissionStatusCode":
   n.AGSubmissionStatusCode = nil
  case "HTTPStatusCode":
   n.HTTPStatusCode = nil
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "AGRuleList":
   n.AGRuleList = nil
  case "ErrorText":
   n.ErrorText = nil
  case "EntityName":
   n.EntityName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGReportingObjectResponseType")
        }
        return n
}

func (n *AGReportingObjectResponseType) Set(key string, value interface{}) *AGReportingObjectResponseType {
        if n == nil {
                n = AGReportingObjectResponseTypeCreate(AGReportingObjectResponseType{})
        }
        switch key {
    case "SIFRefId":
                      n.SIFRefId = StringCreate(value.(string))
    case "SubmittedRefId":
                      n.SubmittedRefId = StringCreate(value.(string))
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
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "AGRuleList":
                      n.AGRuleList = AGRuleListTypeCreate(value.(AGRuleListType))
    case "ErrorText":
                      n.ErrorText = StringCreate(value.(string))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGReportingObjectResponseType")
        }
        return n
}


func (t *WellbeingCharacteristic) CopyString(key string, value interface{}) *WellbeingCharacteristic {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingCharacteristic) Unset(key string) *WellbeingCharacteristic {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "EmergencyResponsePlan":
   n.EmergencyResponsePlan = nil
  case "EmergencyManagement":
   n.EmergencyManagement = nil
  case "Alert":
   n.Alert = nil
  case "WellbeingCharacteristicCategory":
   n.WellbeingCharacteristicCategory = nil
  case "ConfidentialFlag":
   n.ConfidentialFlag = nil
  case "MedicationList":
   n.MedicationList = nil
  case "WellbeingCharacteristicNotes":
   n.WellbeingCharacteristicNotes = nil
  case "WellbeingCharacteristicStartDate":
   n.WellbeingCharacteristicStartDate = nil
  case "WellbeingCharacteristicSubCategory":
   n.WellbeingCharacteristicSubCategory = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "DailyManagement":
   n.DailyManagement = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCharacteristicCode":
   n.LocalCharacteristicCode = nil
  case "SymptomList":
   n.SymptomList = nil
  case "RefId":
   n.RefId = nil
  case "WellbeingCharacteristicReviewDate":
   n.WellbeingCharacteristicReviewDate = nil
  case "DocumentList":
   n.DocumentList = nil
  case "LocalId":
   n.LocalId = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "Trigger":
   n.Trigger = nil
  case "WellbeingCharacteristicClassification":
   n.WellbeingCharacteristicClassification = nil
  case "CharacteristicSeverity":
   n.CharacteristicSeverity = nil
  case "WellbeingCharacteristicEndDate":
   n.WellbeingCharacteristicEndDate = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingCharacteristic")
        }
        return n
}

func (n *WellbeingCharacteristic) Set(key string, value interface{}) *WellbeingCharacteristic {
        if n == nil {
                n = WellbeingCharacteristicCreate(WellbeingCharacteristic{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "EmergencyResponsePlan":
                      n.EmergencyResponsePlan = StringCreate(value.(string))
    case "EmergencyManagement":
                      n.EmergencyManagement = StringCreate(value.(string))
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
    case "WellbeingCharacteristicCategory":
                      n.WellbeingCharacteristicCategory = StringCreate(value.(string))
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
    case "MedicationList":
                      n.MedicationList = MedicationListTypeCreate(value.(MedicationListType))
    case "WellbeingCharacteristicNotes":
                      n.WellbeingCharacteristicNotes = StringCreate(value.(string))
    case "WellbeingCharacteristicStartDate":
                      n.WellbeingCharacteristicStartDate = StringCreate(value.(string))
    case "WellbeingCharacteristicSubCategory":
                      n.WellbeingCharacteristicSubCategory = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "DailyManagement":
                      n.DailyManagement = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCharacteristicCode":
    
                      n.LocalCharacteristicCode = ((*LocalIdType)(StringCreate(value.(string))))
    case "SymptomList":
                      n.SymptomList = SymptomListTypeCreate(value.(SymptomListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "WellbeingCharacteristicReviewDate":
                      n.WellbeingCharacteristicReviewDate = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "Trigger":
                      n.Trigger = StringCreate(value.(string))
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
    case "CharacteristicSeverity":
                      n.CharacteristicSeverity = StringCreate(value.(string))
    case "WellbeingCharacteristicEndDate":
                      n.WellbeingCharacteristicEndDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingCharacteristic")
        }
        return n
}


func (t *LibraryItemInfoType) CopyString(key string, value interface{}) *LibraryItemInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LibraryItemInfoType) Unset(key string) *LibraryItemInfoType {
        switch key {
  case "ISBN":
   n.ISBN = nil
  case "CallNumber":
   n.CallNumber = nil
  case "ElectronicId":
   n.ElectronicId = nil
  case "Cost":
   n.Cost = nil
  case "ReplacementCost":
   n.ReplacementCost = nil
  case "Title":
   n.Title = nil
  case "Type":
   n.Type = nil
  case "Author":
   n.Author = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryItemInfoType")
        }
        return n
}

func (n *LibraryItemInfoType) Set(key string, value interface{}) *LibraryItemInfoType {
        if n == nil {
                n = LibraryItemInfoTypeCreate(LibraryItemInfoType{})
        }
        switch key {
    case "ISBN":
                      n.ISBN = StringCreate(value.(string))
    case "CallNumber":
                      n.CallNumber = StringCreate(value.(string))
    case "ElectronicId":
                      n.ElectronicId = ElectronicIdTypeCreate(value.(ElectronicIdType))
    case "Cost":
                      n.Cost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "ReplacementCost":
                      n.ReplacementCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Author":
                      n.Author = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryItemInfoType")
        }
        return n
}


func (t *Identity_SIF_RefId) CopyString(key string, value interface{}) *Identity_SIF_RefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Identity_SIF_RefId) Unset(key string) *Identity_SIF_RefId {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Identity_SIF_RefId")
        }
        return n
}

func (n *Identity_SIF_RefId) Set(key string, value interface{}) *Identity_SIF_RefId {
        if n == nil {
                n = Identity_SIF_RefIdCreate(Identity_SIF_RefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Identity_SIF_RefId")
        }
        return n
}


func (t *FQContextualQuestionType) CopyString(key string, value interface{}) *FQContextualQuestionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FQContextualQuestionType) Unset(key string) *FQContextualQuestionType {
        switch key {
  case "FQAnswer":
   n.FQAnswer = nil
  case "FQContext":
   n.FQContext = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQContextualQuestionType")
        }
        return n
}

func (n *FQContextualQuestionType) Set(key string, value interface{}) *FQContextualQuestionType {
        if n == nil {
                n = FQContextualQuestionTypeCreate(FQContextualQuestionType{})
        }
        switch key {
    case "FQAnswer":
                      n.FQAnswer = StringCreate(value.(string))
    case "FQContext":
                      n.FQContext = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQContextualQuestionType")
        }
        return n
}


func (t *AttendanceInfoType) CopyString(key string, value interface{}) *AttendanceInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AttendanceInfoType) Unset(key string) *AttendanceInfoType {
        switch key {
  case "AttendanceValue":
   n.AttendanceValue = nil
  case "CountsTowardAttendance":
   n.CountsTowardAttendance = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceInfoType")
        }
        return n
}

func (n *AttendanceInfoType) Set(key string, value interface{}) *AttendanceInfoType {
        if n == nil {
                n = AttendanceInfoTypeCreate(AttendanceInfoType{})
        }
        switch key {
    case "AttendanceValue":
                      n.AttendanceValue = FloatCreate(value.(float64))
    case "CountsTowardAttendance":
                      n.CountsTowardAttendance = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceInfoType")
        }
        return n
}


func (t *PersonInfoType) CopyString(key string, value interface{}) *PersonInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonInfoType) Unset(key string) *PersonInfoType {
        switch key {
  case "HouseholdContactInfoList":
   n.HouseholdContactInfoList = nil
  case "EmailList":
   n.EmailList = nil
  case "OtherNames":
   n.OtherNames = nil
  case "Demographics":
   n.Demographics = nil
  case "AddressList":
   n.AddressList = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInfoType")
        }
        return n
}

func (n *PersonInfoType) Set(key string, value interface{}) *PersonInfoType {
        if n == nil {
                n = PersonInfoTypeCreate(PersonInfoType{})
        }
        switch key {
    case "HouseholdContactInfoList":
                      n.HouseholdContactInfoList = HouseholdContactInfoListTypeCreate(value.(HouseholdContactInfoListType))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "OtherNames":
                      n.OtherNames = OtherNamesTypeCreate(value.(OtherNamesType))
    case "Demographics":
                      n.Demographics = DemographicsTypeCreate(value.(DemographicsType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "Name":
                      n.Name = NameOfRecordTypeCreate(value.(NameOfRecordType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInfoType")
        }
        return n
}


func (t *LearningStandardItem) CopyString(key string, value interface{}) *LearningStandardItem {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LearningStandardItem) Unset(key string) *LearningStandardItem {
        switch key {
  case "StandardSettingBody":
   n.StandardSettingBody = nil
  case "StatementCodes":
   n.StatementCodes = nil
  case "LearningStandardDocumentRefId":
   n.LearningStandardDocumentRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "Resources":
   n.Resources = nil
  case "PredecessorItems":
   n.PredecessorItems = nil
  case "RefId":
   n.RefId = nil
  case "ACStrandSubjectArea":
   n.ACStrandSubjectArea = nil
  case "YearLevels":
   n.YearLevels = nil
  case "Statements":
   n.Statements = nil
  case "StandardIdentifier":
   n.StandardIdentifier = nil
  case "Level5":
   n.Level5 = nil
  case "RelatedLearningStandardItems":
   n.RelatedLearningStandardItems = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "StandardHierarchyLevel":
   n.StandardHierarchyLevel = nil
  case "Level4":
   n.Level4 = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardItem")
        }
        return n
}

func (n *LearningStandardItem) Set(key string, value interface{}) *LearningStandardItem {
        if n == nil {
                n = LearningStandardItemCreate(LearningStandardItem{})
        }
        switch key {
    case "StandardSettingBody":
                      n.StandardSettingBody = StandardsSettingBodyTypeCreate(value.(StandardsSettingBodyType))
    case "StatementCodes":
                      n.StatementCodes = StatementCodesTypeCreate(value.(StatementCodesType))
    case "LearningStandardDocumentRefId":
                      n.LearningStandardDocumentRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Resources":
                      n.Resources = LResourcesTypeCreate(value.(LResourcesType))
    case "PredecessorItems":
                      n.PredecessorItems = LearningStandardsTypeCreate(value.(LearningStandardsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ACStrandSubjectArea":
                      n.ACStrandSubjectArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "Statements":
                      n.Statements = StatementsTypeCreate(value.(StatementsType))
    case "StandardIdentifier":
                      n.StandardIdentifier = StandardIdentifierTypeCreate(value.(StandardIdentifierType))
    case "Level5":
                      n.Level5 = StringCreate(value.(string))
    case "RelatedLearningStandardItems":
                      n.RelatedLearningStandardItems = RelatedLearningStandardItemRefIdListTypeCreate(value.(RelatedLearningStandardItemRefIdListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StandardHierarchyLevel":
                      n.StandardHierarchyLevel = StandardHierarchyLevelTypeCreate(value.(StandardHierarchyLevelType))
    case "Level4":
                      n.Level4 = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningStandardItem")
        }
        return n
}


func (t *SIF_MetadataType) CopyString(key string, value interface{}) *SIF_MetadataType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SIF_MetadataType) Unset(key string) *SIF_MetadataType {
        switch key {
  case "LifeCycle":
   n.LifeCycle = nil
  case "TimeElements":
   n.TimeElements = nil
  case "EducationFilter":
   n.EducationFilter = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SIF_MetadataType")
        }
        return n
}

func (n *SIF_MetadataType) Set(key string, value interface{}) *SIF_MetadataType {
        if n == nil {
                n = SIF_MetadataTypeCreate(SIF_MetadataType{})
        }
        switch key {
    case "LifeCycle":
                      n.LifeCycle = LifeCycleTypeCreate(value.(LifeCycleType))
    case "TimeElements":
                      n.TimeElements = SIF_MetadataType_TimeElementsCreate(value.(SIF_MetadataType_TimeElements))
    case "EducationFilter":
                      n.EducationFilter = EducationFilterTypeCreate(value.(EducationFilterType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SIF_MetadataType")
        }
        return n
}


func (t *WellbeingDocumentType) CopyString(key string, value interface{}) *WellbeingDocumentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingDocumentType) Unset(key string) *WellbeingDocumentType {
        switch key {
  case "DocumentDescription":
   n.DocumentDescription = nil
  case "Sensitivity":
   n.Sensitivity = nil
  case "URL":
   n.URL = nil
  case "DocumentReviewDate":
   n.DocumentReviewDate = nil
  case "Location":
   n.Location = nil
  case "DocumentType":
   n.DocumentType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingDocumentType")
        }
        return n
}

func (n *WellbeingDocumentType) Set(key string, value interface{}) *WellbeingDocumentType {
        if n == nil {
                n = WellbeingDocumentTypeCreate(WellbeingDocumentType{})
        }
        switch key {
    case "DocumentDescription":
                      n.DocumentDescription = StringCreate(value.(string))
    case "Sensitivity":
                      n.Sensitivity = StringCreate(value.(string))
    case "URL":
                      n.URL = StringCreate(value.(string))
    case "DocumentReviewDate":
                      n.DocumentReviewDate = StringCreate(value.(string))
    case "Location":
                      n.Location = StringCreate(value.(string))
    case "DocumentType":
                      n.DocumentType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingDocumentType")
        }
        return n
}


func (t *PurchasingItemType) CopyString(key string, value interface{}) *PurchasingItemType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PurchasingItemType) Unset(key string) *PurchasingItemType {
        switch key {
  case "TaxRate":
   n.TaxRate = nil
  case "LocalItemId":
   n.LocalItemId = nil
  case "ItemNumber":
   n.ItemNumber = nil
  case "QuantityDelivered":
   n.QuantityDelivered = nil
  case "UnitCost":
   n.UnitCost = nil
  case "CancelledOrder":
   n.CancelledOrder = nil
  case "ExpenseAccounts":
   n.ExpenseAccounts = nil
  case "Quantity":
   n.Quantity = nil
  case "TotalCost":
   n.TotalCost = nil
  case "ItemDescription":
   n.ItemDescription = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PurchasingItemType")
        }
        return n
}

func (n *PurchasingItemType) Set(key string, value interface{}) *PurchasingItemType {
        if n == nil {
                n = PurchasingItemTypeCreate(PurchasingItemType{})
        }
        switch key {
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "LocalItemId":
    
                      n.LocalItemId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ItemNumber":
                      n.ItemNumber = StringCreate(value.(string))
    case "QuantityDelivered":
                      n.QuantityDelivered = StringCreate(value.(string))
    case "UnitCost":
                      n.UnitCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
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
    case "ExpenseAccounts":
                      n.ExpenseAccounts = ExpenseAccountsTypeCreate(value.(ExpenseAccountsType))
    case "Quantity":
                      n.Quantity = StringCreate(value.(string))
    case "TotalCost":
                      n.TotalCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "ItemDescription":
                      n.ItemDescription = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PurchasingItemType")
        }
        return n
}


func (t *GradingAssignment) CopyString(key string, value interface{}) *GradingAssignment {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *GradingAssignment) Unset(key string) *GradingAssignment {
        switch key {
  case "SubAssignmentList":
   n.SubAssignmentList = nil
  case "LevelAssessed":
   n.LevelAssessed = nil
  case "Description":
   n.Description = nil
  case "TeachingGroupRefId":
   n.TeachingGroupRefId = nil
  case "GradingCategory":
   n.GradingCategory = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "LearningStandardList":
   n.LearningStandardList = nil
  case "DueDate":
   n.DueDate = nil
  case "AssignmentPurpose":
   n.AssignmentPurpose = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "AssessmentType":
   n.AssessmentType = nil
  case "RubricScoringGuide":
   n.RubricScoringGuide = nil
  case "DetailedDescriptionBinary":
   n.DetailedDescriptionBinary = nil
  case "Weight":
   n.Weight = nil
  case "DetailedDescriptionURL":
   n.DetailedDescriptionURL = nil
  case "CreateDate":
   n.CreateDate = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "MaxAttemptsAllowed":
   n.MaxAttemptsAllowed = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StudentPersonalRefIdList":
   n.StudentPersonalRefIdList = nil
  case "PrerequisiteList":
   n.PrerequisiteList = nil
  case "PointsPossible":
   n.PointsPossible = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradingAssignment")
        }
        return n
}

func (n *GradingAssignment) Set(key string, value interface{}) *GradingAssignment {
        if n == nil {
                n = GradingAssignmentCreate(GradingAssignment{})
        }
        switch key {
    case "SubAssignmentList":
                      n.SubAssignmentList = AssignmentListTypeCreate(value.(AssignmentListType))
    case "LevelAssessed":
                      n.LevelAssessed = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "GradingCategory":
                      n.GradingCategory = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "AssignmentPurpose":
                      n.AssignmentPurpose = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AssessmentType":
                      n.AssessmentType = StringCreate(value.(string))
    case "RubricScoringGuide":
                      n.RubricScoringGuide = GenericRubricTypeCreate(value.(GenericRubricType))
    case "DetailedDescriptionBinary":
                      n.DetailedDescriptionBinary = StringCreate(value.(string))
    case "Weight":
                      n.Weight = FloatCreate(value.(float64))
    case "DetailedDescriptionURL":
                      n.DetailedDescriptionURL = StringCreate(value.(string))
    case "CreateDate":
                      n.CreateDate = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "MaxAttemptsAllowed":
                      n.MaxAttemptsAllowed = IntCreate(value.(int))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentPersonalRefIdList":
                      n.StudentPersonalRefIdList = StudentsTypeCreate(value.(StudentsType))
    case "PrerequisiteList":
                      n.PrerequisiteList = PrerequisitesTypeCreate(value.(PrerequisitesType))
    case "PointsPossible":
                      n.PointsPossible = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GradingAssignment")
        }
        return n
}


func (t *Journal) CopyString(key string, value interface{}) *Journal {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Journal) Unset(key string) *Journal {
        switch key {
  case "LocalId":
   n.LocalId = nil
  case "CreatedBy":
   n.CreatedBy = nil
  case "CreatedDate":
   n.CreatedDate = nil
  case "CreditFinancialAccountRefId":
   n.CreditFinancialAccountRefId = nil
  case "OriginatingTransactionRefId":
   n.OriginatingTransactionRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "GSTCodeOriginal":
   n.GSTCodeOriginal = nil
  case "ApprovedDate":
   n.ApprovedDate = nil
  case "JournalAdjustmentList":
   n.JournalAdjustmentList = nil
  case "CreditAccountCode":
   n.CreditAccountCode = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Note":
   n.Note = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "DebitAccountCode":
   n.DebitAccountCode = nil
  case "ApprovedBy":
   n.ApprovedBy = nil
  case "RefId":
   n.RefId = nil
  case "Amount":
   n.Amount = nil
  case "GSTCodeReplacement":
   n.GSTCodeReplacement = nil
  case "DebitFinancialAccountRefId":
   n.DebitFinancialAccountRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Journal")
        }
        return n
}

func (n *Journal) Set(key string, value interface{}) *Journal {
        if n == nil {
                n = JournalCreate(Journal{})
        }
        switch key {
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CreatedBy":
                      n.CreatedBy = StringCreate(value.(string))
    case "CreatedDate":
                      n.CreatedDate = StringCreate(value.(string))
    case "CreditFinancialAccountRefId":
                      n.CreditFinancialAccountRefId = StringCreate(value.(string))
    case "OriginatingTransactionRefId":
                      n.OriginatingTransactionRefId = Journal_OriginatingTransactionRefIdCreate(value.(Journal_OriginatingTransactionRefId))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "GSTCodeOriginal":
                      n.GSTCodeOriginal = StringCreate(value.(string))
    case "ApprovedDate":
                      n.ApprovedDate = StringCreate(value.(string))
    case "JournalAdjustmentList":
                      n.JournalAdjustmentList = JournalAdjustmentListTypeCreate(value.(JournalAdjustmentListType))
    case "CreditAccountCode":
                      n.CreditAccountCode = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Note":
                      n.Note = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "DebitAccountCode":
                      n.DebitAccountCode = StringCreate(value.(string))
    case "ApprovedBy":
                      n.ApprovedBy = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "GSTCodeReplacement":
                      n.GSTCodeReplacement = StringCreate(value.(string))
    case "DebitFinancialAccountRefId":
                      n.DebitFinancialAccountRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Journal")
        }
        return n
}


func (t *DemographicsType) CopyString(key string, value interface{}) *DemographicsType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DemographicsType) Unset(key string) *DemographicsType {
        switch key {
  case "Sex":
   n.Sex = nil
  case "ReligiousRegion":
   n.ReligiousRegion = nil
  case "CountryOfBirth":
   n.CountryOfBirth = nil
  case "CountriesOfResidency":
   n.CountriesOfResidency = nil
  case "CountriesOfCitizenship":
   n.CountriesOfCitizenship = nil
  case "LanguageList":
   n.LanguageList = nil
  case "VisaSubClassList":
   n.VisaSubClassList = nil
  case "CountryArrivalDate":
   n.CountryArrivalDate = nil
  case "InterpreterRequired":
   n.InterpreterRequired = nil
  case "StateOfBirth":
   n.StateOfBirth = nil
  case "PermanentResident":
   n.PermanentResident = nil
  case "DwellingArrangement":
   n.DwellingArrangement = nil
  case "DateOfDeath":
   n.DateOfDeath = nil
  case "PlaceOfBirth":
   n.PlaceOfBirth = nil
  case "VisaSubClass":
   n.VisaSubClass = nil
  case "ImmunisationCertificateStatus":
   n.ImmunisationCertificateStatus = nil
  case "IndigenousStatus":
   n.IndigenousStatus = nil
  case "MedicareNumber":
   n.MedicareNumber = nil
  case "MaritalStatus":
   n.MaritalStatus = nil
  case "VisaStatisticalCode":
   n.VisaStatisticalCode = nil
  case "Religion":
   n.Religion = nil
  case "BirthDate":
   n.BirthDate = nil
  case "BirthDateVerification":
   n.BirthDateVerification = nil
  case "EnglishProficiency":
   n.EnglishProficiency = nil
  case "AustralianCitizenshipStatus":
   n.AustralianCitizenshipStatus = nil
  case "CulturalBackground":
   n.CulturalBackground = nil
  case "LBOTE":
   n.LBOTE = nil
  case "VisaExpiryDate":
   n.VisaExpiryDate = nil
  case "ReligiousEventList":
   n.ReligiousEventList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DemographicsType")
        }
        return n
}

func (n *DemographicsType) Set(key string, value interface{}) *DemographicsType {
        if n == nil {
                n = DemographicsTypeCreate(DemographicsType{})
        }
        switch key {
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
    case "ReligiousRegion":
                      n.ReligiousRegion = StringCreate(value.(string))
    case "CountryOfBirth":
      present := false
  for _, b := range AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values")
      }

                      n.CountryOfBirth = ((*CountryType)(StringCreate(value.(string))))
    case "CountriesOfResidency":
                      n.CountriesOfResidency = CountryList2TypeCreate(value.(CountryList2Type))
    case "CountriesOfCitizenship":
                      n.CountriesOfCitizenship = CountryListTypeCreate(value.(CountryListType))
    case "LanguageList":
                      n.LanguageList = LanguageListTypeCreate(value.(LanguageListType))
    case "VisaSubClassList":
                      n.VisaSubClassList = VisaSubClassListTypeCreate(value.(VisaSubClassListType))
    case "CountryArrivalDate":
                      n.CountryArrivalDate = StringCreate(value.(string))
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
    case "StateOfBirth":
    
                      n.StateOfBirth = ((*StateProvinceType)(StringCreate(value.(string))))
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
    case "DateOfDeath":
                      n.DateOfDeath = StringCreate(value.(string))
    case "PlaceOfBirth":
                      n.PlaceOfBirth = StringCreate(value.(string))
    case "VisaSubClass":
    
                      n.VisaSubClass = ((*VisaSubClassCodeType)(StringCreate(value.(string))))
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
    case "MedicareNumber":
                      n.MedicareNumber = StringCreate(value.(string))
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
    case "VisaStatisticalCode":
                      n.VisaStatisticalCode = StringCreate(value.(string))
    case "Religion":
                      n.Religion = ReligionTypeCreate(value.(ReligionType))
    case "BirthDate":
    
                      n.BirthDate = ((*BirthDateType)(StringCreate(value.(string))))
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
    case "EnglishProficiency":
                      n.EnglishProficiency = EnglishProficiencyTypeCreate(value.(EnglishProficiencyType))
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
    case "VisaExpiryDate":
                      n.VisaExpiryDate = StringCreate(value.(string))
    case "ReligiousEventList":
                      n.ReligiousEventList = ReligiousEventListTypeCreate(value.(ReligiousEventListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DemographicsType")
        }
        return n
}


func (t *TimeTablePeriodType) CopyString(key string, value interface{}) *TimeTablePeriodType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTablePeriodType) Unset(key string) *TimeTablePeriodType {
        switch key {
  case "RegularSchoolPeriod":
   n.RegularSchoolPeriod = nil
  case "BellPeriod":
   n.BellPeriod = nil
  case "EndTime":
   n.EndTime = nil
  case "InstructionalMinutes":
   n.InstructionalMinutes = nil
  case "StartTime":
   n.StartTime = nil
  case "UseInAttendanceCalculations":
   n.UseInAttendanceCalculations = nil
  case "PeriodId":
   n.PeriodId = nil
  case "PeriodTitle":
   n.PeriodTitle = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTablePeriodType")
        }
        return n
}

func (n *TimeTablePeriodType) Set(key string, value interface{}) *TimeTablePeriodType {
        if n == nil {
                n = TimeTablePeriodTypeCreate(TimeTablePeriodType{})
        }
        switch key {
    case "RegularSchoolPeriod":
                      n.RegularSchoolPeriod = StringCreate(value.(string))
    case "BellPeriod":
                      n.BellPeriod = StringCreate(value.(string))
    case "EndTime":
                      n.EndTime = StringCreate(value.(string))
    case "InstructionalMinutes":
                      n.InstructionalMinutes = IntCreate(value.(int))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
    case "UseInAttendanceCalculations":
                      n.UseInAttendanceCalculations = StringCreate(value.(string))
    case "PeriodId":
    
                      n.PeriodId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PeriodTitle":
                      n.PeriodTitle = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTablePeriodType")
        }
        return n
}


func (t *AddressType) CopyString(key string, value interface{}) *AddressType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AddressType) Unset(key string) *AddressType {
        switch key {
  case "RadioContact":
   n.RadioContact = nil
  case "AddressGlobalUID":
   n.AddressGlobalUID = nil
  case "City":
   n.City = nil
  case "Role":
   n.Role = nil
  case "Community":
   n.Community = nil
  case "EffectiveFromDate":
   n.EffectiveFromDate = nil
  case "Country":
   n.Country = nil
  case "MapReference":
   n.MapReference = nil
  case "StateProvince":
   n.StateProvince = nil
  case "EffectiveToDate":
   n.EffectiveToDate = nil
  case "LocalId":
   n.LocalId = nil
  case "GridLocation":
   n.GridLocation = nil
  case "PostalCode":
   n.PostalCode = nil
  case "StatisticalAreas":
   n.StatisticalAreas = nil
  case "Type":
   n.Type = nil
  case "Street":
   n.Street = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressType")
        }
        return n
}

func (n *AddressType) Set(key string, value interface{}) *AddressType {
        if n == nil {
                n = AddressTypeCreate(AddressType{})
        }
        switch key {
    case "RadioContact":
                      n.RadioContact = StringCreate(value.(string))
    case "AddressGlobalUID":
    
                      n.AddressGlobalUID = ((*GUIDType)(StringCreate(value.(string))))
    case "City":
                      n.City = StringCreate(value.(string))
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
    case "Community":
                      n.Community = StringCreate(value.(string))
    case "EffectiveFromDate":
                      n.EffectiveFromDate = StringCreate(value.(string))
    case "Country":
      present := false
  for _, b := range AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values")
      }

                      n.Country = ((*CountryType)(StringCreate(value.(string))))
    case "MapReference":
                      n.MapReference = MapReferenceTypeCreate(value.(MapReferenceType))
    case "StateProvince":
    
                      n.StateProvince = ((*StateProvinceType)(StringCreate(value.(string))))
    case "EffectiveToDate":
                      n.EffectiveToDate = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "GridLocation":
                      n.GridLocation = GridLocationTypeCreate(value.(GridLocationType))
    case "PostalCode":
                      n.PostalCode = StringCreate(value.(string))
    case "StatisticalAreas":
                      n.StatisticalAreas = StatisticalAreasTypeCreate(value.(StatisticalAreasType))
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
    case "Street":
                      n.Street = AddressStreetTypeCreate(value.(AddressStreetType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressType")
        }
        return n
}


func (t *DetentionContainerType) CopyString(key string, value interface{}) *DetentionContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DetentionContainerType) Unset(key string) *DetentionContainerType {
        switch key {
  case "DetentionCategory":
   n.DetentionCategory = nil
  case "DetentionLocation":
   n.DetentionLocation = nil
  case "DetentionDate":
   n.DetentionDate = nil
  case "DetentionNotes":
   n.DetentionNotes = nil
  case "Status":
   n.Status = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DetentionContainerType")
        }
        return n
}

func (n *DetentionContainerType) Set(key string, value interface{}) *DetentionContainerType {
        if n == nil {
                n = DetentionContainerTypeCreate(DetentionContainerType{})
        }
        switch key {
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
    case "DetentionDate":
                      n.DetentionDate = StringCreate(value.(string))
    case "DetentionNotes":
                      n.DetentionNotes = StringCreate(value.(string))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DetentionContainerType")
        }
        return n
}


func (t *WellbeingEvent) CopyString(key string, value interface{}) *WellbeingEvent {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingEvent) Unset(key string) *WellbeingEvent {
        switch key {
  case "WellbeingEventTime":
   n.WellbeingEventTime = nil
  case "PersonInvolvementList":
   n.PersonInvolvementList = nil
  case "RefId":
   n.RefId = nil
  case "FollowUpActionList":
   n.FollowUpActionList = nil
  case "ConfidentialFlag":
   n.ConfidentialFlag = nil
  case "WellbeingEventDate":
   n.WellbeingEventDate = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "WellbeingEventCategoryList":
   n.WellbeingEventCategoryList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ReportingStaffRefId":
   n.ReportingStaffRefId = nil
  case "WellbeingEventTimePeriod":
   n.WellbeingEventTimePeriod = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "WellbeingEventCategoryClass":
   n.WellbeingEventCategoryClass = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "WellbeingEventLocationDetails":
   n.WellbeingEventLocationDetails = nil
  case "WellbeingEventDescription":
   n.WellbeingEventDescription = nil
  case "WellbeingEventCreationTimeStamp":
   n.WellbeingEventCreationTimeStamp = nil
  case "WellbeingEventNotes":
   n.WellbeingEventNotes = nil
  case "GroupIndicator":
   n.GroupIndicator = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "DocumentList":
   n.DocumentList = nil
  case "EventId":
   n.EventId = nil
  case "Status":
   n.Status = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEvent")
        }
        return n
}

func (n *WellbeingEvent) Set(key string, value interface{}) *WellbeingEvent {
        if n == nil {
                n = WellbeingEventCreate(WellbeingEvent{})
        }
        switch key {
    case "WellbeingEventTime":
                      n.WellbeingEventTime = StringCreate(value.(string))
    case "PersonInvolvementList":
                      n.PersonInvolvementList = PersonInvolvementListTypeCreate(value.(PersonInvolvementListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "FollowUpActionList":
                      n.FollowUpActionList = FollowUpActionListTypeCreate(value.(FollowUpActionListType))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "WellbeingEventCategoryList":
                      n.WellbeingEventCategoryList = WellbeingEventCategoryListTypeCreate(value.(WellbeingEventCategoryListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ReportingStaffRefId":
                      n.ReportingStaffRefId = StringCreate(value.(string))
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
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
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
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "WellbeingEventLocationDetails":
                      n.WellbeingEventLocationDetails = WellbeingEventLocationDetailsTypeCreate(value.(WellbeingEventLocationDetailsType))
    case "WellbeingEventDescription":
                      n.WellbeingEventDescription = StringCreate(value.(string))
    case "WellbeingEventCreationTimeStamp":
                      n.WellbeingEventCreationTimeStamp = StringCreate(value.(string))
    case "WellbeingEventNotes":
                      n.WellbeingEventNotes = StringCreate(value.(string))
    case "GroupIndicator":
                      n.GroupIndicator = BoolCreate(value.(bool))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "EventId":
    
                      n.EventId = ((*LocalIdType)(StringCreate(value.(string))))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingEvent")
        }
        return n
}


func (t *PersonInvolvementType_PersonRefId) CopyString(key string, value interface{}) *PersonInvolvementType_PersonRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PersonInvolvementType_PersonRefId) Unset(key string) *PersonInvolvementType_PersonRefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInvolvementType_PersonRefId")
        }
        return n
}

func (n *PersonInvolvementType_PersonRefId) Set(key string, value interface{}) *PersonInvolvementType_PersonRefId {
        if n == nil {
                n = PersonInvolvementType_PersonRefIdCreate(PersonInvolvementType_PersonRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PersonInvolvementType_PersonRefId")
        }
        return n
}


func (t *WellbeingResponse) CopyString(key string, value interface{}) *WellbeingResponse {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingResponse) Unset(key string) *WellbeingResponse {
        switch key {
  case "PersonInvolvementList":
   n.PersonInvolvementList = nil
  case "RefId":
   n.RefId = nil
  case "WellbeingResponseStartDate":
   n.WellbeingResponseStartDate = nil
  case "Date":
   n.Date = nil
  case "Award":
   n.Award = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Detention":
   n.Detention = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "Suspension":
   n.Suspension = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "WellbeingResponseNotes":
   n.WellbeingResponseNotes = nil
  case "PlanRequired":
   n.PlanRequired = nil
  case "WellbeingResponseEndDate":
   n.WellbeingResponseEndDate = nil
  case "WellbeingResponseCategory":
   n.WellbeingResponseCategory = nil
  case "OtherResponse":
   n.OtherResponse = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "DocumentList":
   n.DocumentList = nil
  case "LocalId":
   n.LocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingResponse")
        }
        return n
}

func (n *WellbeingResponse) Set(key string, value interface{}) *WellbeingResponse {
        if n == nil {
                n = WellbeingResponseCreate(WellbeingResponse{})
        }
        switch key {
    case "PersonInvolvementList":
                      n.PersonInvolvementList = PersonInvolvementListTypeCreate(value.(PersonInvolvementListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "WellbeingResponseStartDate":
                      n.WellbeingResponseStartDate = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "Award":
                      n.Award = AwardContainerTypeCreate(value.(AwardContainerType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Detention":
                      n.Detention = DetentionContainerTypeCreate(value.(DetentionContainerType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "Suspension":
                      n.Suspension = SuspensionContainerTypeCreate(value.(SuspensionContainerType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "WellbeingResponseNotes":
                      n.WellbeingResponseNotes = StringCreate(value.(string))
    case "PlanRequired":
                      n.PlanRequired = PlanRequiredContainerTypeCreate(value.(PlanRequiredContainerType))
    case "WellbeingResponseEndDate":
                      n.WellbeingResponseEndDate = StringCreate(value.(string))
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
    case "OtherResponse":
                      n.OtherResponse = OtherWellbeingResponseContainerTypeCreate(value.(OtherWellbeingResponseContainerType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "DocumentList":
                      n.DocumentList = WellbeingDocumentListTypeCreate(value.(WellbeingDocumentListType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingResponse")
        }
        return n
}


func (t *TermInfo) CopyString(key string, value interface{}) *TermInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TermInfo) Unset(key string) *TermInfo {
        switch key {
  case "RefId":
   n.RefId = nil
  case "EndDate":
   n.EndDate = nil
  case "AttendanceTerm":
   n.AttendanceTerm = nil
  case "SchedulingTerm":
   n.SchedulingTerm = nil
  case "MarkingTerm":
   n.MarkingTerm = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TermCode":
   n.TermCode = nil
  case "Track":
   n.Track = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "Description":
   n.Description = nil
  case "TermSpan":
   n.TermSpan = nil
  case "StartDate":
   n.StartDate = nil
  case "RelativeDuration":
   n.RelativeDuration = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TermInfo")
        }
        return n
}

func (n *TermInfo) Set(key string, value interface{}) *TermInfo {
        if n == nil {
                n = TermInfoCreate(TermInfo{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
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
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TermCode":
                      n.TermCode = StringCreate(value.(string))
    case "Track":
                      n.Track = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
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
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "RelativeDuration":
                      n.RelativeDuration = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TermInfo")
        }
        return n
}


func (t *ChargedLocationInfo) CopyString(key string, value interface{}) *ChargedLocationInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ChargedLocationInfo) Unset(key string) *ChargedLocationInfo {
        switch key {
  case "Name":
   n.Name = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "Description":
   n.Description = nil
  case "AddressList":
   n.AddressList = nil
  case "SiteCategory":
   n.SiteCategory = nil
  case "ParentChargedLocationInfoRefId":
   n.ParentChargedLocationInfoRefId = nil
  case "LocalId":
   n.LocalId = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "LocationType":
   n.LocationType = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ChargedLocationInfo")
        }
        return n
}

func (n *ChargedLocationInfo) Set(key string, value interface{}) *ChargedLocationInfo {
        if n == nil {
                n = ChargedLocationInfoCreate(ChargedLocationInfo{})
        }
        switch key {
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "SiteCategory":
                      n.SiteCategory = StringCreate(value.(string))
    case "ParentChargedLocationInfoRefId":
                      n.ParentChargedLocationInfoRefId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "LocationType":
                      n.LocationType = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ChargedLocationInfo")
        }
        return n
}


func (t *FollowUpActionType) CopyString(key string, value interface{}) *FollowUpActionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FollowUpActionType) Unset(key string) *FollowUpActionType {
        switch key {
  case "FollowUpActionCategory":
   n.FollowUpActionCategory = nil
  case "FollowUpDetails":
   n.FollowUpDetails = nil
  case "WellbeingResponseRefId":
   n.WellbeingResponseRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FollowUpActionType")
        }
        return n
}

func (n *FollowUpActionType) Set(key string, value interface{}) *FollowUpActionType {
        if n == nil {
                n = FollowUpActionTypeCreate(FollowUpActionType{})
        }
        switch key {
    case "FollowUpActionCategory":
                      n.FollowUpActionCategory = StringCreate(value.(string))
    case "FollowUpDetails":
                      n.FollowUpDetails = StringCreate(value.(string))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FollowUpActionType")
        }
        return n
}


func (t *ResourceUsage_ResourceUsageContentType) CopyString(key string, value interface{}) *ResourceUsage_ResourceUsageContentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceUsage_ResourceUsageContentType) Unset(key string) *ResourceUsage_ResourceUsageContentType {
        switch key {
  case "Code":
   n.Code = nil
  case "LocalDescription":
   n.LocalDescription = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceUsageContentType")
        }
        return n
}

func (n *ResourceUsage_ResourceUsageContentType) Set(key string, value interface{}) *ResourceUsage_ResourceUsageContentType {
        if n == nil {
                n = ResourceUsage_ResourceUsageContentTypeCreate(ResourceUsage_ResourceUsageContentType{})
        }
        switch key {
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
    case "LocalDescription":
                      n.LocalDescription = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceUsageContentType")
        }
        return n
}


func (t *WithdrawalType) CopyString(key string, value interface{}) *WithdrawalType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WithdrawalType) Unset(key string) *WithdrawalType {
        switch key {
  case "ScheduledActivityRefId":
   n.ScheduledActivityRefId = nil
  case "TimeTableCellRefId":
   n.TimeTableCellRefId = nil
  case "WithdrawalStartTime":
   n.WithdrawalStartTime = nil
  case "WithdrawalDate":
   n.WithdrawalDate = nil
  case "TimeTableSubjectRefId":
   n.TimeTableSubjectRefId = nil
  case "WithdrawalEndTime":
   n.WithdrawalEndTime = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WithdrawalType")
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
    case "WithdrawalStartTime":
                      n.WithdrawalStartTime = StringCreate(value.(string))
    case "WithdrawalDate":
                      n.WithdrawalDate = StringCreate(value.(string))
    case "TimeTableSubjectRefId":
                      n.TimeTableSubjectRefId = StringCreate(value.(string))
    case "WithdrawalEndTime":
                      n.WithdrawalEndTime = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WithdrawalType")
        }
        return n
}


func (t *Invoice_InvoicedEntity) CopyString(key string, value interface{}) *Invoice_InvoicedEntity {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Invoice_InvoicedEntity) Unset(key string) *Invoice_InvoicedEntity {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Invoice_InvoicedEntity")
        }
        return n
}

func (n *Invoice_InvoicedEntity) Set(key string, value interface{}) *Invoice_InvoicedEntity {
        if n == nil {
                n = Invoice_InvoicedEntityCreate(Invoice_InvoicedEntity{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Invoice_InvoicedEntity")
        }
        return n
}


func (t *WellbeingPersonLink_PersonRefId) CopyString(key string, value interface{}) *WellbeingPersonLink_PersonRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingPersonLink_PersonRefId) Unset(key string) *WellbeingPersonLink_PersonRefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPersonLink_PersonRefId")
        }
        return n
}

func (n *WellbeingPersonLink_PersonRefId) Set(key string, value interface{}) *WellbeingPersonLink_PersonRefId {
        if n == nil {
                n = WellbeingPersonLink_PersonRefIdCreate(WellbeingPersonLink_PersonRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPersonLink_PersonRefId")
        }
        return n
}


func (t *AwardContainerType) CopyString(key string, value interface{}) *AwardContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AwardContainerType) Unset(key string) *AwardContainerType {
        switch key {
  case "AwardDescription":
   n.AwardDescription = nil
  case "Status":
   n.Status = nil
  case "AwardNotes":
   n.AwardNotes = nil
  case "AwardType":
   n.AwardType = nil
  case "AwardDate":
   n.AwardDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AwardContainerType")
        }
        return n
}

func (n *AwardContainerType) Set(key string, value interface{}) *AwardContainerType {
        if n == nil {
                n = AwardContainerTypeCreate(AwardContainerType{})
        }
        switch key {
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
    case "AwardNotes":
                      n.AwardNotes = StringCreate(value.(string))
    case "AwardType":
                      n.AwardType = StringCreate(value.(string))
    case "AwardDate":
                      n.AwardDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AwardContainerType")
        }
        return n
}


func (t *ExpenseAccountType) CopyString(key string, value interface{}) *ExpenseAccountType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ExpenseAccountType) Unset(key string) *ExpenseAccountType {
        switch key {
  case "AccountingPeriod":
   n.AccountingPeriod = nil
  case "Amount":
   n.Amount = nil
  case "AccountCode":
   n.AccountCode = nil
  case "FinancialAccountRefId":
   n.FinancialAccountRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ExpenseAccountType")
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
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
    case "FinancialAccountRefId":
                      n.FinancialAccountRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ExpenseAccountType")
        }
        return n
}


func (t *MedicationType) CopyString(key string, value interface{}) *MedicationType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MedicationType) Unset(key string) *MedicationType {
        switch key {
  case "Frequency":
   n.Frequency = nil
  case "Method":
   n.Method = nil
  case "MedicationName":
   n.MedicationName = nil
  case "Dosage":
   n.Dosage = nil
  case "AdministrationInformation":
   n.AdministrationInformation = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MedicationType")
        }
        return n
}

func (n *MedicationType) Set(key string, value interface{}) *MedicationType {
        if n == nil {
                n = MedicationTypeCreate(MedicationType{})
        }
        switch key {
    case "Frequency":
                      n.Frequency = StringCreate(value.(string))
    case "Method":
                      n.Method = StringCreate(value.(string))
    case "MedicationName":
                      n.MedicationName = StringCreate(value.(string))
    case "Dosage":
                      n.Dosage = StringCreate(value.(string))
    case "AdministrationInformation":
                      n.AdministrationInformation = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MedicationType")
        }
        return n
}


func (t *SystemRole_RoleScopeRefId) CopyString(key string, value interface{}) *SystemRole_RoleScopeRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole_RoleScopeRefId) Unset(key string) *SystemRole_RoleScopeRefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_RoleScopeRefId")
        }
        return n
}

func (n *SystemRole_RoleScopeRefId) Set(key string, value interface{}) *SystemRole_RoleScopeRefId {
        if n == nil {
                n = SystemRole_RoleScopeRefIdCreate(SystemRole_RoleScopeRefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_RoleScopeRefId")
        }
        return n
}


func (t *ResourceUsage_ResourceReportLine) CopyString(key string, value interface{}) *ResourceUsage_ResourceReportLine {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceUsage_ResourceReportLine) Unset(key string) *ResourceUsage_ResourceReportLine {
        switch key {
  case "ReportRow":
   n.ReportRow = nil
  case "SIF_RefId":
   n.SIF_RefId = nil
  case "CurrentCost":
   n.CurrentCost = nil
  case "StartDate":
   n.StartDate = nil
  case "EndDate":
   n.EndDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceReportLine")
        }
        return n
}

func (n *ResourceUsage_ResourceReportLine) Set(key string, value interface{}) *ResourceUsage_ResourceReportLine {
        if n == nil {
                n = ResourceUsage_ResourceReportLineCreate(ResourceUsage_ResourceReportLine{})
        }
        switch key {
    case "ReportRow":
                      n.ReportRow = StringCreate(value.(string))
    case "SIF_RefId":
                      n.SIF_RefId = ResourceUsage_SIF_RefIdCreate(value.(ResourceUsage_SIF_RefId))
    case "CurrentCost":
                      n.CurrentCost = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceReportLine")
        }
        return n
}


func (t *NAPTestletContentType) CopyString(key string, value interface{}) *NAPTestletContentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestletContentType) Unset(key string) *NAPTestletContentType {
        switch key {
  case "TestletName":
   n.TestletName = nil
  case "NAPTestletLocalId":
   n.NAPTestletLocalId = nil
  case "TestletMaximumScore":
   n.TestletMaximumScore = nil
  case "Node":
   n.Node = nil
  case "LocationInStage":
   n.LocationInStage = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletContentType")
        }
        return n
}

func (n *NAPTestletContentType) Set(key string, value interface{}) *NAPTestletContentType {
        if n == nil {
                n = NAPTestletContentTypeCreate(NAPTestletContentType{})
        }
        switch key {
    case "TestletName":
                      n.TestletName = StringCreate(value.(string))
    case "NAPTestletLocalId":
    
                      n.NAPTestletLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestletMaximumScore":
                      n.TestletMaximumScore = FloatCreate(value.(float64))
    case "Node":
                      n.Node = StringCreate(value.(string))
    case "LocationInStage":
                      n.LocationInStage = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletContentType")
        }
        return n
}


func (t *NAPTestletCodeFrameType) CopyString(key string, value interface{}) *NAPTestletCodeFrameType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestletCodeFrameType) Unset(key string) *NAPTestletCodeFrameType {
        switch key {
  case "TestItemList":
   n.TestItemList = nil
  case "TestletContent":
   n.TestletContent = nil
  case "NAPTestletRefId":
   n.NAPTestletRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletCodeFrameType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestletCodeFrameType")
        }
        return n
}


func (t *AbstractContentElementType_TextData) CopyString(key string, value interface{}) *AbstractContentElementType_TextData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentElementType_TextData) Unset(key string) *AbstractContentElementType_TextData {
        switch key {
  case "Value":
   n.Value = nil
  case "FileName":
   n.FileName = nil
  case "Description":
   n.Description = nil
  case "MIMEType":
   n.MIMEType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_TextData")
        }
        return n
}

func (n *AbstractContentElementType_TextData) Set(key string, value interface{}) *AbstractContentElementType_TextData {
        if n == nil {
                n = AbstractContentElementType_TextDataCreate(AbstractContentElementType_TextData{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "FileName":
                      n.FileName = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_TextData")
        }
        return n
}


func (t *ACStrandSubjectAreaType) CopyString(key string, value interface{}) *ACStrandSubjectAreaType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ACStrandSubjectAreaType) Unset(key string) *ACStrandSubjectAreaType {
        switch key {
  case "SubjectArea":
   n.SubjectArea = nil
  case "ACStrand":
   n.ACStrand = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ACStrandSubjectAreaType")
        }
        return n
}

func (n *ACStrandSubjectAreaType) Set(key string, value interface{}) *ACStrandSubjectAreaType {
        if n == nil {
                n = ACStrandSubjectAreaTypeCreate(ACStrandSubjectAreaType{})
        }
        switch key {
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ACStrandSubjectAreaType")
        }
        return n
}


func (t *AbstractContentElementType_Reference) CopyString(key string, value interface{}) *AbstractContentElementType_Reference {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentElementType_Reference) Unset(key string) *AbstractContentElementType_Reference {
        switch key {
  case "MIMEType":
   n.MIMEType = nil
  case "Description":
   n.Description = nil
  case "URL":
   n.URL = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_Reference")
        }
        return n
}

func (n *AbstractContentElementType_Reference) Set(key string, value interface{}) *AbstractContentElementType_Reference {
        if n == nil {
                n = AbstractContentElementType_ReferenceCreate(AbstractContentElementType_Reference{})
        }
        switch key {
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "URL":
                      n.URL = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentElementType_Reference")
        }
        return n
}


func (t *CensusCollection) CopyString(key string, value interface{}) *CensusCollection {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CensusCollection) Unset(key string) *CensusCollection {
        switch key {
  case "ReportingAuthorityCommonwealthId":
   n.ReportingAuthorityCommonwealthId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "RoundCode":
   n.RoundCode = nil
  case "CensusReportingList":
   n.CensusReportingList = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SoftwareVendorInfo":
   n.SoftwareVendorInfo = nil
  case "CensusYear":
   n.CensusYear = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusCollection")
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "CensusReportingList":
                      n.CensusReportingList = CensusReportingListTypeCreate(value.(CensusReportingListType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "CensusYear":
    
                      n.CensusYear = ((*SchoolYearType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusCollection")
        }
        return n
}


func (t *StudentActivityParticipation) CopyString(key string, value interface{}) *StudentActivityParticipation {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentActivityParticipation) Unset(key string) *StudentActivityParticipation {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RecognitionList":
   n.RecognitionList = nil
  case "RefId":
   n.RefId = nil
  case "EndDate":
   n.EndDate = nil
  case "StudentActivityInfoRefId":
   n.StudentActivityInfoRefId = nil
  case "Role":
   n.Role = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "ParticipationComment":
   n.ParticipationComment = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "StartDate":
   n.StartDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityParticipation")
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RecognitionList":
                      n.RecognitionList = RecognitionListTypeCreate(value.(RecognitionListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "StudentActivityInfoRefId":
                      n.StudentActivityInfoRefId = StringCreate(value.(string))
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "ParticipationComment":
                      n.ParticipationComment = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityParticipation")
        }
        return n
}


func (t *NAPSubscoreType) CopyString(key string, value interface{}) *NAPSubscoreType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPSubscoreType) Unset(key string) *NAPSubscoreType {
        switch key {
  case "SubscoreValue":
   n.SubscoreValue = nil
  case "SubscoreType":
   n.SubscoreType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPSubscoreType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPSubscoreType")
        }
        return n
}


func (t *OtherWellbeingResponseContainerType) CopyString(key string, value interface{}) *OtherWellbeingResponseContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *OtherWellbeingResponseContainerType) Unset(key string) *OtherWellbeingResponseContainerType {
        switch key {
  case "OtherResponseDescription":
   n.OtherResponseDescription = nil
  case "Status":
   n.Status = nil
  case "OtherResponseDate":
   n.OtherResponseDate = nil
  case "OtherResponseNotes":
   n.OtherResponseNotes = nil
  case "OtherResponseType":
   n.OtherResponseType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherWellbeingResponseContainerType")
        }
        return n
}

func (n *OtherWellbeingResponseContainerType) Set(key string, value interface{}) *OtherWellbeingResponseContainerType {
        if n == nil {
                n = OtherWellbeingResponseContainerTypeCreate(OtherWellbeingResponseContainerType{})
        }
        switch key {
    case "OtherResponseDescription":
                      n.OtherResponseDescription = StringCreate(value.(string))
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
    case "OtherResponseDate":
                      n.OtherResponseDate = StringCreate(value.(string))
    case "OtherResponseNotes":
                      n.OtherResponseNotes = StringCreate(value.(string))
    case "OtherResponseType":
                      n.OtherResponseType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherWellbeingResponseContainerType")
        }
        return n
}


func (t *DomainScoreType) CopyString(key string, value interface{}) *DomainScoreType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DomainScoreType) Unset(key string) *DomainScoreType {
        switch key {
  case "PlausibleScaledValueList":
   n.PlausibleScaledValueList = nil
  case "ScaledScoreStandardError":
   n.ScaledScoreStandardError = nil
  case "StudentDomainBand":
   n.StudentDomainBand = nil
  case "RawScore":
   n.RawScore = nil
  case "ScaledScoreLogitStandardError":
   n.ScaledScoreLogitStandardError = nil
  case "StudentProficiency":
   n.StudentProficiency = nil
  case "ScaledScoreLogitValue":
   n.ScaledScoreLogitValue = nil
  case "ScaledScoreValue":
   n.ScaledScoreValue = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainScoreType")
        }
        return n
}

func (n *DomainScoreType) Set(key string, value interface{}) *DomainScoreType {
        if n == nil {
                n = DomainScoreTypeCreate(DomainScoreType{})
        }
        switch key {
    case "PlausibleScaledValueList":
                      n.PlausibleScaledValueList = PlausibleScaledValueListTypeCreate(value.(PlausibleScaledValueListType))
    case "ScaledScoreStandardError":
                      n.ScaledScoreStandardError = FloatCreate(value.(float64))
    case "StudentDomainBand":
                      n.StudentDomainBand = IntCreate(value.(int))
    case "RawScore":
                      n.RawScore = FloatCreate(value.(float64))
    case "ScaledScoreLogitStandardError":
                      n.ScaledScoreLogitStandardError = FloatCreate(value.(float64))
    case "StudentProficiency":
                      n.StudentProficiency = StringCreate(value.(string))
    case "ScaledScoreLogitValue":
                      n.ScaledScoreLogitValue = FloatCreate(value.(float64))
    case "ScaledScoreValue":
                      n.ScaledScoreValue = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainScoreType")
        }
        return n
}


func (t *AggregateStatisticInfo) CopyString(key string, value interface{}) *AggregateStatisticInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AggregateStatisticInfo) Unset(key string) *AggregateStatisticInfo {
        switch key {
  case "Measure":
   n.Measure = nil
  case "StatisticName":
   n.StatisticName = nil
  case "EffectiveDate":
   n.EffectiveDate = nil
  case "ApprovalDate":
   n.ApprovalDate = nil
  case "DiscontinueDate":
   n.DiscontinueDate = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "ExclusionRules":
   n.ExclusionRules = nil
  case "Location":
   n.Location = nil
  case "CalculationRule":
   n.CalculationRule = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Source":
   n.Source = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "ExpirationDate":
   n.ExpirationDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticInfo")
        }
        return n
}

func (n *AggregateStatisticInfo) Set(key string, value interface{}) *AggregateStatisticInfo {
        if n == nil {
                n = AggregateStatisticInfoCreate(AggregateStatisticInfo{})
        }
        switch key {
    case "Measure":
                      n.Measure = StringCreate(value.(string))
    case "StatisticName":
                      n.StatisticName = StringCreate(value.(string))
    case "EffectiveDate":
                      n.EffectiveDate = StringCreate(value.(string))
    case "ApprovalDate":
                      n.ApprovalDate = StringCreate(value.(string))
    case "DiscontinueDate":
                      n.DiscontinueDate = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ExclusionRules":
                      n.ExclusionRules = ExclusionRulesTypeCreate(value.(ExclusionRulesType))
    case "Location":
                      n.Location = LocationTypeCreate(value.(LocationType))
    case "CalculationRule":
                      n.CalculationRule = AggregateStatisticInfo_CalculationRuleCreate(value.(AggregateStatisticInfo_CalculationRule))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Source":
                      n.Source = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ExpirationDate":
                      n.ExpirationDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticInfo")
        }
        return n
}


func (t *SubstituteItemType) CopyString(key string, value interface{}) *SubstituteItemType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SubstituteItemType) Unset(key string) *SubstituteItemType {
        switch key {
  case "PNPCodeList":
   n.PNPCodeList = nil
  case "SubstituteItemRefId":
   n.SubstituteItemRefId = nil
  case "SubstituteItemLocalId":
   n.SubstituteItemLocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SubstituteItemType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SubstituteItemType")
        }
        return n
}


func (t *SIF_ExtendedElementsType_SIF_ExtendedElement) CopyString(key string, value interface{}) *SIF_ExtendedElementsType_SIF_ExtendedElement {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) Unset(key string) *SIF_ExtendedElementsType_SIF_ExtendedElement {
        switch key {
  case "SIF_Action":
   n.SIF_Action = nil
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SIF_ExtendedElementsType_SIF_ExtendedElement")
        }
        return n
}

func (n *SIF_ExtendedElementsType_SIF_ExtendedElement) Set(key string, value interface{}) *SIF_ExtendedElementsType_SIF_ExtendedElement {
        if n == nil {
                n = SIF_ExtendedElementsType_SIF_ExtendedElementCreate(SIF_ExtendedElementsType_SIF_ExtendedElement{})
        }
        switch key {
    case "SIF_Action":
                      n.SIF_Action = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
    
                      n.Value = ((*ExtendedContentType)(StringCreate(value.(string))))
    case "Name":
                      n.Name = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SIF_ExtendedElementsType_SIF_ExtendedElement")
        }
        return n
}


func (t *StatsCohortYearLevelType) CopyString(key string, value interface{}) *StatsCohortYearLevelType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StatsCohortYearLevelType) Unset(key string) *StatsCohortYearLevelType {
        switch key {
  case "CohortYearLevel":
   n.CohortYearLevel = nil
  case "StatsCohortList":
   n.StatsCohortList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatsCohortYearLevelType")
        }
        return n
}

func (n *StatsCohortYearLevelType) Set(key string, value interface{}) *StatsCohortYearLevelType {
        if n == nil {
                n = StatsCohortYearLevelTypeCreate(StatsCohortYearLevelType{})
        }
        switch key {
    case "CohortYearLevel":
                      n.CohortYearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "StatsCohortList":
                      n.StatsCohortList = StatsCohortListTypeCreate(value.(StatsCohortListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatsCohortYearLevelType")
        }
        return n
}


func (t *WellbeingPersonLink) CopyString(key string, value interface{}) *WellbeingPersonLink {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingPersonLink) Unset(key string) *WellbeingPersonLink {
        switch key {
  case "OtherPersonId":
   n.OtherPersonId = nil
  case "LocalId":
   n.LocalId = nil
  case "HowInvolved":
   n.HowInvolved = nil
  case "ShortName":
   n.ShortName = nil
  case "PersonRole":
   n.PersonRole = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "OtherPersonContactDetails":
   n.OtherPersonContactDetails = nil
  case "WellbeingResponseRefId":
   n.WellbeingResponseRefId = nil
  case "PersonRefId":
   n.PersonRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "WellbeingEventRefId":
   n.WellbeingEventRefId = nil
  case "RefId":
   n.RefId = nil
  case "GroupId":
   n.GroupId = nil
  case "FollowUpActionList":
   n.FollowUpActionList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPersonLink")
        }
        return n
}

func (n *WellbeingPersonLink) Set(key string, value interface{}) *WellbeingPersonLink {
        if n == nil {
                n = WellbeingPersonLinkCreate(WellbeingPersonLink{})
        }
        switch key {
    case "OtherPersonId":
    
                      n.OtherPersonId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "HowInvolved":
                      n.HowInvolved = StringCreate(value.(string))
    case "ShortName":
                      n.ShortName = StringCreate(value.(string))
    case "PersonRole":
                      n.PersonRole = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "OtherPersonContactDetails":
                      n.OtherPersonContactDetails = StringCreate(value.(string))
    case "WellbeingResponseRefId":
                      n.WellbeingResponseRefId = StringCreate(value.(string))
    case "PersonRefId":
                      n.PersonRefId = WellbeingPersonLink_PersonRefIdCreate(value.(WellbeingPersonLink_PersonRefId))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "WellbeingEventRefId":
                      n.WellbeingEventRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "GroupId":
    
                      n.GroupId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FollowUpActionList":
                      n.FollowUpActionList = FollowUpActionListTypeCreate(value.(FollowUpActionListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPersonLink")
        }
        return n
}


func (t *StudentContactRelationship) CopyString(key string, value interface{}) *StudentContactRelationship {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentContactRelationship) Unset(key string) *StudentContactRelationship {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "MainlySpeaksEnglishAtHome":
   n.MainlySpeaksEnglishAtHome = nil
  case "ContactSequence":
   n.ContactSequence = nil
  case "ParentRelationshipStatus":
   n.ParentRelationshipStatus = nil
  case "HouseholdList":
   n.HouseholdList = nil
  case "StudentContactPersonalRefId":
   n.StudentContactPersonalRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Relationship":
   n.Relationship = nil
  case "StudentContactRelationshipRefId":
   n.StudentContactRelationshipRefId = nil
  case "ContactFlags":
   n.ContactFlags = nil
  case "ContactSequenceSource":
   n.ContactSequenceSource = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentContactRelationship")
        }
        return n
}

func (n *StudentContactRelationship) Set(key string, value interface{}) *StudentContactRelationship {
        if n == nil {
                n = StudentContactRelationshipCreate(StudentContactRelationship{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "ContactSequence":
                      n.ContactSequence = IntCreate(value.(int))
    case "ParentRelationshipStatus":
                      n.ParentRelationshipStatus = StringCreate(value.(string))
    case "HouseholdList":
                      n.HouseholdList = HouseholdListTypeCreate(value.(HouseholdListType))
    case "StudentContactPersonalRefId":
    
                      n.StudentContactPersonalRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Relationship":
                      n.Relationship = RelationshipTypeCreate(value.(RelationshipType))
    case "StudentContactRelationshipRefId":
                      n.StudentContactRelationshipRefId = StringCreate(value.(string))
    case "ContactFlags":
                      n.ContactFlags = ContactFlagsTypeCreate(value.(ContactFlagsType))
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
    case "StudentPersonalRefId":
    
                      n.StudentPersonalRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentContactRelationship")
        }
        return n
}


func (t *IdentityAssertionsType_IdentityAssertion) CopyString(key string, value interface{}) *IdentityAssertionsType_IdentityAssertion {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *IdentityAssertionsType_IdentityAssertion) Unset(key string) *IdentityAssertionsType_IdentityAssertion {
        switch key {
  case "SchemaName":
   n.SchemaName = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "IdentityAssertionsType_IdentityAssertion")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "IdentityAssertionsType_IdentityAssertion")
        }
        return n
}


func (t *GenericRubricType) CopyString(key string, value interface{}) *GenericRubricType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *GenericRubricType) Unset(key string) *GenericRubricType {
        switch key {
  case "Descriptor":
   n.Descriptor = nil
  case "RubricType":
   n.RubricType = nil
  case "ScoreList":
   n.ScoreList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GenericRubricType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GenericRubricType")
        }
        return n
}


func (t *MarkerType) CopyString(key string, value interface{}) *MarkerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MarkerType) Unset(key string) *MarkerType {
        switch key {
  case "Role":
   n.Role = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MarkerType")
        }
        return n
}

func (n *MarkerType) Set(key string, value interface{}) *MarkerType {
        if n == nil {
                n = MarkerTypeCreate(MarkerType{})
        }
        switch key {
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MarkerType")
        }
        return n
}


func (t *TimeElementType) CopyString(key string, value interface{}) *TimeElementType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeElementType) Unset(key string) *TimeElementType {
        switch key {
  case "Type":
   n.Type = nil
  case "StartDateTime":
   n.StartDateTime = nil
  case "Name":
   n.Name = nil
  case "IsCurrent":
   n.IsCurrent = nil
  case "Value":
   n.Value = nil
  case "SpanGaps":
   n.SpanGaps = nil
  case "Code":
   n.Code = nil
  case "EndDateTime":
   n.EndDateTime = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeElementType")
        }
        return n
}

func (n *TimeElementType) Set(key string, value interface{}) *TimeElementType {
        if n == nil {
                n = TimeElementTypeCreate(TimeElementType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "IsCurrent":
                      n.IsCurrent = BoolCreate(value.(bool))
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SpanGaps":
                      n.SpanGaps = TimeElementType_SpanGapsCreate(value.(TimeElementType_SpanGaps))
    case "Code":
                      n.Code = StringCreate(value.(string))
    case "EndDateTime":
                      n.EndDateTime = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeElementType")
        }
        return n
}


func (t *OtherCodeListType_OtherCode) CopyString(key string, value interface{}) *OtherCodeListType_OtherCode {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *OtherCodeListType_OtherCode) Unset(key string) *OtherCodeListType_OtherCode {
        switch key {
  case "Codeset":
   n.Codeset = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherCodeListType_OtherCode")
        }
        return n
}

func (n *OtherCodeListType_OtherCode) Set(key string, value interface{}) *OtherCodeListType_OtherCode {
        if n == nil {
                n = OtherCodeListType_OtherCodeCreate(OtherCodeListType_OtherCode{})
        }
        switch key {
    case "Codeset":
                      n.Codeset = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "OtherCodeListType_OtherCode")
        }
        return n
}


func (t *StudentEntryContainerType) CopyString(key string, value interface{}) *StudentEntryContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentEntryContainerType) Unset(key string) *StudentEntryContainerType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentEntryContainerType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentEntryContainerType")
        }
        return n
}


func (t *DwellingArrangementType) CopyString(key string, value interface{}) *DwellingArrangementType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DwellingArrangementType) Unset(key string) *DwellingArrangementType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DwellingArrangementType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DwellingArrangementType")
        }
        return n
}


func (t *AGRoundType) CopyString(key string, value interface{}) *AGRoundType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AGRoundType) Unset(key string) *AGRoundType {
        switch key {
  case "DueDate":
   n.DueDate = nil
  case "StartDate":
   n.StartDate = nil
  case "EndDate":
   n.EndDate = nil
  case "RoundCode":
   n.RoundCode = nil
  case "RoundName":
   n.RoundName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGRoundType")
        }
        return n
}

func (n *AGRoundType) Set(key string, value interface{}) *AGRoundType {
        if n == nil {
                n = AGRoundTypeCreate(AGRoundType{})
        }
        switch key {
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "RoundName":
                      n.RoundName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGRoundType")
        }
        return n
}


func (t *SystemRole_RoleScope) CopyString(key string, value interface{}) *SystemRole_RoleScope {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole_RoleScope) Unset(key string) *SystemRole_RoleScope {
        switch key {
  case "RoleScopeName":
   n.RoleScopeName = nil
  case "RoleScopeRefId":
   n.RoleScopeRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_RoleScope")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_RoleScope")
        }
        return n
}


func (t *ProgramFundingSourceType) CopyString(key string, value interface{}) *ProgramFundingSourceType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ProgramFundingSourceType) Unset(key string) *ProgramFundingSourceType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramFundingSourceType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramFundingSourceType")
        }
        return n
}


func (t *SchoolCourseInfoOverrideType) CopyString(key string, value interface{}) *SchoolCourseInfoOverrideType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolCourseInfoOverrideType) Unset(key string) *SchoolCourseInfoOverrideType {
        switch key {
  case "Override":
   n.Override = nil
  case "InstructionalLevel":
   n.InstructionalLevel = nil
  case "CourseCode":
   n.CourseCode = nil
  case "DistrictCourseCode":
   n.DistrictCourseCode = nil
  case "CourseCredits":
   n.CourseCredits = nil
  case "SubjectArea":
   n.SubjectArea = nil
  case "StateCourseCode":
   n.StateCourseCode = nil
  case "CourseTitle":
   n.CourseTitle = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolCourseInfoOverrideType")
        }
        return n
}

func (n *SchoolCourseInfoOverrideType) Set(key string, value interface{}) *SchoolCourseInfoOverrideType {
        if n == nil {
                n = SchoolCourseInfoOverrideTypeCreate(SchoolCourseInfoOverrideType{})
        }
        switch key {
    case "Override":
                      n.Override = StringCreate(value.(string))
    case "InstructionalLevel":
                      n.InstructionalLevel = StringCreate(value.(string))
    case "CourseCode":
                      n.CourseCode = StringCreate(value.(string))
    case "DistrictCourseCode":
                      n.DistrictCourseCode = StringCreate(value.(string))
    case "CourseCredits":
                      n.CourseCredits = StringCreate(value.(string))
    case "SubjectArea":
                      n.SubjectArea = SubjectAreaTypeCreate(value.(SubjectAreaType))
    case "StateCourseCode":
                      n.StateCourseCode = StringCreate(value.(string))
    case "CourseTitle":
                      n.CourseTitle = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolCourseInfoOverrideType")
        }
        return n
}


func (t *NAPTest) CopyString(key string, value interface{}) *NAPTest {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTest) Unset(key string) *NAPTest {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "RefId":
   n.RefId = nil
  case "TestContent":
   n.TestContent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTest")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTest")
        }
        return n
}


func (t *AddressCollectionReportingType) CopyString(key string, value interface{}) *AddressCollectionReportingType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AddressCollectionReportingType) Unset(key string) *AddressCollectionReportingType {
        switch key {
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "AddressCollectionStudentList":
   n.AddressCollectionStudentList = nil
  case "EntityName":
   n.EntityName = nil
  case "LocalId":
   n.LocalId = nil
  case "ACARAId":
   n.ACARAId = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "AGContextualQuestionList":
   n.AGContextualQuestionList = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "EntityContact":
   n.EntityContact = nil
  case "EntityLevel":
   n.EntityLevel = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollectionReportingType")
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
    case "AddressCollectionStudentList":
                      n.AddressCollectionStudentList = AddressCollectionStudentListTypeCreate(value.(AddressCollectionStudentListType))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "AGContextualQuestionList":
                      n.AGContextualQuestionList = AGContextualQuestionListTypeCreate(value.(AGContextualQuestionListType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AddressCollectionReportingType")
        }
        return n
}


func (t *AttendanceCodeType) CopyString(key string, value interface{}) *AttendanceCodeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AttendanceCodeType) Unset(key string) *AttendanceCodeType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceCodeType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AttendanceCodeType")
        }
        return n
}


func (t *StudentAttendanceCollectionReportingType) CopyString(key string, value interface{}) *StudentAttendanceCollectionReportingType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentAttendanceCollectionReportingType) Unset(key string) *StudentAttendanceCollectionReportingType {
        switch key {
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "LocalId":
   n.LocalId = nil
  case "ACARAId":
   n.ACARAId = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "EntityName":
   n.EntityName = nil
  case "StatsCohortYearLevelList":
   n.StatsCohortYearLevelList = nil
  case "EntityLevel":
   n.EntityLevel = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "EntityContact":
   n.EntityContact = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceCollectionReportingType")
        }
        return n
}

func (n *StudentAttendanceCollectionReportingType) Set(key string, value interface{}) *StudentAttendanceCollectionReportingType {
        if n == nil {
                n = StudentAttendanceCollectionReportingTypeCreate(StudentAttendanceCollectionReportingType{})
        }
        switch key {
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "StatsCohortYearLevelList":
                      n.StatsCohortYearLevelList = StatsCohortYearLevelListTypeCreate(value.(StatsCohortYearLevelListType))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceCollectionReportingType")
        }
        return n
}


func (t *StaffAssignment) CopyString(key string, value interface{}) *StaffAssignment {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StaffAssignment) Unset(key string) *StaffAssignment {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "JobFunction":
   n.JobFunction = nil
  case "JobStartDate":
   n.JobStartDate = nil
  case "JobEndDate":
   n.JobEndDate = nil
  case "RefId":
   n.RefId = nil
  case "PreviousSchoolName":
   n.PreviousSchoolName = nil
  case "StaffActivity":
   n.StaffActivity = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "JobFTE":
   n.JobFTE = nil
  case "StaffSubjectList":
   n.StaffSubjectList = nil
  case "Homegroup":
   n.Homegroup = nil
  case "CalendarSummaryList":
   n.CalendarSummaryList = nil
  case "Description":
   n.Description = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "House":
   n.House = nil
  case "YearLevels":
   n.YearLevels = nil
  case "CasualReliefTeacher":
   n.CasualReliefTeacher = nil
  case "EmploymentStatus":
   n.EmploymentStatus = nil
  case "AvailableForTimetable":
   n.AvailableForTimetable = nil
  case "PrimaryAssignment":
   n.PrimaryAssignment = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffAssignment")
        }
        return n
}

func (n *StaffAssignment) Set(key string, value interface{}) *StaffAssignment {
        if n == nil {
                n = StaffAssignmentCreate(StaffAssignment{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "JobFunction":
                      n.JobFunction = StringCreate(value.(string))
    case "JobStartDate":
                      n.JobStartDate = StringCreate(value.(string))
    case "JobEndDate":
                      n.JobEndDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PreviousSchoolName":
                      n.PreviousSchoolName = StringCreate(value.(string))
    case "StaffActivity":
                      n.StaffActivity = StaffActivityExtensionTypeCreate(value.(StaffActivityExtensionType))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "JobFTE":
                      n.JobFTE = FloatCreate(value.(float64))
    case "StaffSubjectList":
                      n.StaffSubjectList = StaffSubjectListTypeCreate(value.(StaffSubjectListType))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "CalendarSummaryList":
                      n.CalendarSummaryList = CalendarSummaryListTypeCreate(value.(CalendarSummaryListType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "House":
                      n.House = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
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
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StaffAssignment")
        }
        return n
}


func (t *CheckoutInfoType) CopyString(key string, value interface{}) *CheckoutInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CheckoutInfoType) Unset(key string) *CheckoutInfoType {
        switch key {
  case "ReturnBy":
   n.ReturnBy = nil
  case "RenewalCount":
   n.RenewalCount = nil
  case "CheckedOutOn":
   n.CheckedOutOn = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CheckoutInfoType")
        }
        return n
}

func (n *CheckoutInfoType) Set(key string, value interface{}) *CheckoutInfoType {
        if n == nil {
                n = CheckoutInfoTypeCreate(CheckoutInfoType{})
        }
        switch key {
    case "ReturnBy":
                      n.ReturnBy = StringCreate(value.(string))
    case "RenewalCount":
                      n.RenewalCount = IntCreate(value.(int))
    case "CheckedOutOn":
                      n.CheckedOutOn = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CheckoutInfoType")
        }
        return n
}


func (t *RoomInfo) CopyString(key string, value interface{}) *RoomInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *RoomInfo) Unset(key string) *RoomInfo {
        switch key {
  case "AvailableForTimetable":
   n.AvailableForTimetable = nil
  case "PhoneNumber":
   n.PhoneNumber = nil
  case "LocalId":
   n.LocalId = nil
  case "StaffList":
   n.StaffList = nil
  case "Description":
   n.Description = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "RoomNumber":
   n.RoomNumber = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Building":
   n.Building = nil
  case "RoomType":
   n.RoomType = nil
  case "HomeroomNumber":
   n.HomeroomNumber = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "Size":
   n.Size = nil
  case "Capacity":
   n.Capacity = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RoomInfo")
        }
        return n
}

func (n *RoomInfo) Set(key string, value interface{}) *RoomInfo {
        if n == nil {
                n = RoomInfoCreate(RoomInfo{})
        }
        switch key {
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
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StaffList":
                      n.StaffList = StaffListTypeCreate(value.(StaffListType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "RoomNumber":
                      n.RoomNumber = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Building":
                      n.Building = StringCreate(value.(string))
    case "RoomType":
                      n.RoomType = StringCreate(value.(string))
    case "HomeroomNumber":
                      n.HomeroomNumber = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Size":
                      n.Size = FloatCreate(value.(float64))
    case "Capacity":
                      n.Capacity = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "RoomInfo")
        }
        return n
}


func (t *GridLocationType) CopyString(key string, value interface{}) *GridLocationType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *GridLocationType) Unset(key string) *GridLocationType {
        switch key {
  case "Longitude":
   n.Longitude = nil
  case "Latitude":
   n.Latitude = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GridLocationType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "GridLocationType")
        }
        return n
}


func (t *AbstractContentPackageType_XMLData) CopyString(key string, value interface{}) *AbstractContentPackageType_XMLData {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentPackageType_XMLData) Unset(key string) *AbstractContentPackageType_XMLData {
        switch key {
  case "Value":
   n.Value = nil
  case "Description":
   n.Description = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_XMLData")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_XMLData")
        }
        return n
}


func (t *EntityContactInfoType) CopyString(key string, value interface{}) *EntityContactInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EntityContactInfoType) Unset(key string) *EntityContactInfoType {
        switch key {
  case "Address":
   n.Address = nil
  case "Name":
   n.Name = nil
  case "Qualifications":
   n.Qualifications = nil
  case "PositionTitle":
   n.PositionTitle = nil
  case "Role":
   n.Role = nil
  case "RegistrationDetails":
   n.RegistrationDetails = nil
  case "Email":
   n.Email = nil
  case "PhoneNumber":
   n.PhoneNumber = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EntityContactInfoType")
        }
        return n
}

func (n *EntityContactInfoType) Set(key string, value interface{}) *EntityContactInfoType {
        if n == nil {
                n = EntityContactInfoTypeCreate(EntityContactInfoType{})
        }
        switch key {
    case "Address":
                      n.Address = AddressTypeCreate(value.(AddressType))
    case "Name":
                      n.Name = NameTypeCreate(value.(NameType))
    case "Qualifications":
                      n.Qualifications = StringCreate(value.(string))
    case "PositionTitle":
                      n.PositionTitle = StringCreate(value.(string))
    case "Role":
                      n.Role = StringCreate(value.(string))
    case "RegistrationDetails":
                      n.RegistrationDetails = StringCreate(value.(string))
    case "Email":
                      n.Email = EmailTypeCreate(value.(EmailType))
    case "PhoneNumber":
                      n.PhoneNumber = PhoneNumberTypeCreate(value.(PhoneNumberType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EntityContactInfoType")
        }
        return n
}


func (t *AGContextualQuestionType) CopyString(key string, value interface{}) *AGContextualQuestionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AGContextualQuestionType) Unset(key string) *AGContextualQuestionType {
        switch key {
  case "AGAnswer":
   n.AGAnswer = nil
  case "AGContextCode":
   n.AGContextCode = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGContextualQuestionType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AGContextualQuestionType")
        }
        return n
}


func (t *FineInfoType) CopyString(key string, value interface{}) *FineInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FineInfoType) Unset(key string) *FineInfoType {
        switch key {
  case "Description":
   n.Description = nil
  case "Amount":
   n.Amount = nil
  case "Reference":
   n.Reference = nil
  case "Type":
   n.Type = nil
  case "Assessed":
   n.Assessed = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FineInfoType")
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
    case "Amount":
                      n.Amount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "Reference":
                      n.Reference = StringCreate(value.(string))
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Assessed":
                      n.Assessed = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FineInfoType")
        }
        return n
}


func (t *StudentSubjectChoiceType) CopyString(key string, value interface{}) *StudentSubjectChoiceType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentSubjectChoiceType) Unset(key string) *StudentSubjectChoiceType {
        switch key {
  case "SubjectLocalId":
   n.SubjectLocalId = nil
  case "PreferenceNumber":
   n.PreferenceNumber = nil
  case "OtherSchoolLocalId":
   n.OtherSchoolLocalId = nil
  case "StudyDescription":
   n.StudyDescription = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSubjectChoiceType")
        }
        return n
}

func (n *StudentSubjectChoiceType) Set(key string, value interface{}) *StudentSubjectChoiceType {
        if n == nil {
                n = StudentSubjectChoiceTypeCreate(StudentSubjectChoiceType{})
        }
        switch key {
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
    case "OtherSchoolLocalId":
    
                      n.OtherSchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StudyDescription":
                      n.StudyDescription = SubjectAreaTypeCreate(value.(SubjectAreaType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentSubjectChoiceType")
        }
        return n
}


func (t *YearRangeType) CopyString(key string, value interface{}) *YearRangeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *YearRangeType) Unset(key string) *YearRangeType {
        switch key {
  case "Start":
   n.Start = nil
  case "End":
   n.End = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearRangeType")
        }
        return n
}

func (n *YearRangeType) Set(key string, value interface{}) *YearRangeType {
        if n == nil {
                n = YearRangeTypeCreate(YearRangeType{})
        }
        switch key {
    case "Start":
                      n.Start = YearLevelTypeCreate(value.(YearLevelType))
    case "End":
                      n.End = YearLevelTypeCreate(value.(YearLevelType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearRangeType")
        }
        return n
}


func (t *LibraryTransactionType) CopyString(key string, value interface{}) *LibraryTransactionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LibraryTransactionType) Unset(key string) *LibraryTransactionType {
        switch key {
  case "FineInfoList":
   n.FineInfoList = nil
  case "CheckoutInfo":
   n.CheckoutInfo = nil
  case "HoldInfoList":
   n.HoldInfoList = nil
  case "ItemInfo":
   n.ItemInfo = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryTransactionType")
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
    case "CheckoutInfo":
                      n.CheckoutInfo = CheckoutInfoTypeCreate(value.(CheckoutInfoType))
    case "HoldInfoList":
                      n.HoldInfoList = HoldInfoListTypeCreate(value.(HoldInfoListType))
    case "ItemInfo":
                      n.ItemInfo = LibraryItemInfoTypeCreate(value.(LibraryItemInfoType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryTransactionType")
        }
        return n
}


func (t *PhoneNumberType) CopyString(key string, value interface{}) *PhoneNumberType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PhoneNumberType) Unset(key string) *PhoneNumberType {
        switch key {
  case "Type":
   n.Type = nil
  case "Number":
   n.Number = nil
  case "Preference":
   n.Preference = nil
  case "Extension":
   n.Extension = nil
  case "ListedStatus":
   n.ListedStatus = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PhoneNumberType")
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
    case "Number":
                      n.Number = StringCreate(value.(string))
    case "Preference":
                      n.Preference = IntCreate(value.(int))
    case "Extension":
                      n.Extension = StringCreate(value.(string))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PhoneNumberType")
        }
        return n
}


func (t *Invoice) CopyString(key string, value interface{}) *Invoice {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Invoice) Unset(key string) *Invoice {
        switch key {
  case "FinancialAccountRefIdList":
   n.FinancialAccountRefIdList = nil
  case "InvoicedEntity":
   n.InvoicedEntity = nil
  case "Ledger":
   n.Ledger = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TaxAmount":
   n.TaxAmount = nil
  case "TaxRate":
   n.TaxRate = nil
  case "ChargedLocationInfoRefId":
   n.ChargedLocationInfoRefId = nil
  case "PurchasingItems":
   n.PurchasingItems = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "FormNumber":
   n.FormNumber = nil
  case "NetAmount":
   n.NetAmount = nil
  case "TransactionDescription":
   n.TransactionDescription = nil
  case "Voluntary":
   n.Voluntary = nil
  case "RefId":
   n.RefId = nil
  case "AccountCodeList":
   n.AccountCodeList = nil
  case "RelatedPurchaseOrderRefId":
   n.RelatedPurchaseOrderRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ApprovedBy":
   n.ApprovedBy = nil
  case "AccountingPeriod":
   n.AccountingPeriod = nil
  case "ItemDetail":
   n.ItemDetail = nil
  case "DueDate":
   n.DueDate = nil
  case "TaxType":
   n.TaxType = nil
  case "BillingDate":
   n.BillingDate = nil
  case "BilledAmount":
   n.BilledAmount = nil
  case "LocalId":
   n.LocalId = nil
  case "CreatedBy":
   n.CreatedBy = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Invoice")
        }
        return n
}

func (n *Invoice) Set(key string, value interface{}) *Invoice {
        if n == nil {
                n = InvoiceCreate(Invoice{})
        }
        switch key {
    case "FinancialAccountRefIdList":
                      n.FinancialAccountRefIdList = FinancialAccountRefIdListTypeCreate(value.(FinancialAccountRefIdListType))
    case "InvoicedEntity":
                      n.InvoicedEntity = Invoice_InvoicedEntityCreate(value.(Invoice_InvoicedEntity))
    case "Ledger":
                      n.Ledger = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "PurchasingItems":
                      n.PurchasingItems = PurchasingItemsTypeCreate(value.(PurchasingItemsType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "FormNumber":
    
                      n.FormNumber = ((*LocalIdType)(StringCreate(value.(string))))
    case "NetAmount":
                      n.NetAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
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
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AccountCodeList":
                      n.AccountCodeList = AccountCodeListTypeCreate(value.(AccountCodeListType))
    case "RelatedPurchaseOrderRefId":
                      n.RelatedPurchaseOrderRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ApprovedBy":
                      n.ApprovedBy = StringCreate(value.(string))
    case "AccountingPeriod":
    
                      n.AccountingPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "ItemDetail":
                      n.ItemDetail = StringCreate(value.(string))
    case "DueDate":
                      n.DueDate = StringCreate(value.(string))
    case "TaxType":
                      n.TaxType = StringCreate(value.(string))
    case "BillingDate":
                      n.BillingDate = StringCreate(value.(string))
    case "BilledAmount":
                      n.BilledAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CreatedBy":
                      n.CreatedBy = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Invoice")
        }
        return n
}


func (t *StudentAttendanceCollection) CopyString(key string, value interface{}) *StudentAttendanceCollection {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentAttendanceCollection) Unset(key string) *StudentAttendanceCollection {
        switch key {
  case "StudentAttendanceCollectionReportingList":
   n.StudentAttendanceCollectionReportingList = nil
  case "StudentAttendanceCollectionYear":
   n.StudentAttendanceCollectionYear = nil
  case "RefId":
   n.RefId = nil
  case "RoundCode":
   n.RoundCode = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SoftwareVendorInfo":
   n.SoftwareVendorInfo = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "ReportingAuthorityCommonwealthId":
   n.ReportingAuthorityCommonwealthId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceCollection")
        }
        return n
}

func (n *StudentAttendanceCollection) Set(key string, value interface{}) *StudentAttendanceCollection {
        if n == nil {
                n = StudentAttendanceCollectionCreate(StudentAttendanceCollection{})
        }
        switch key {
    case "StudentAttendanceCollectionReportingList":
                      n.StudentAttendanceCollectionReportingList = StudentAttendanceCollectionReportingListTypeCreate(value.(StudentAttendanceCollectionReportingListType))
    case "StudentAttendanceCollectionYear":
    
                      n.StudentAttendanceCollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceCollection")
        }
        return n
}


func (t *TimeTable) CopyString(key string, value interface{}) *TimeTable {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTable) Unset(key string) *TimeTable {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TimeTableDayList":
   n.TimeTableDayList = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SchoolName":
   n.SchoolName = nil
  case "DaysPerCycle":
   n.DaysPerCycle = nil
  case "EndDate":
   n.EndDate = nil
  case "RefId":
   n.RefId = nil
  case "PeriodsPerDay":
   n.PeriodsPerDay = nil
  case "LocalId":
   n.LocalId = nil
  case "TeachingPeriodsPerDay":
   n.TeachingPeriodsPerDay = nil
  case "StartDate":
   n.StartDate = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Title":
   n.Title = nil
  case "TimeTableCreationDate":
   n.TimeTableCreationDate = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTable")
        }
        return n
}

func (n *TimeTable) Set(key string, value interface{}) *TimeTable {
        if n == nil {
                n = TimeTableCreate(TimeTable{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TimeTableDayList":
                      n.TimeTableDayList = TimeTableDayListTypeCreate(value.(TimeTableDayListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
    case "DaysPerCycle":
                      n.DaysPerCycle = IntCreate(value.(int))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "PeriodsPerDay":
                      n.PeriodsPerDay = IntCreate(value.(int))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TeachingPeriodsPerDay":
                      n.TeachingPeriodsPerDay = IntCreate(value.(int))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "TimeTableCreationDate":
                      n.TimeTableCreationDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTable")
        }
        return n
}


func (t *NAPEventStudentLink) CopyString(key string, value interface{}) *NAPEventStudentLink {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPEventStudentLink) Unset(key string) *NAPEventStudentLink {
        switch key {
  case "LapsedTimeTest":
   n.LapsedTimeTest = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SchoolACARAId":
   n.SchoolACARAId = nil
  case "NAPTestRefId":
   n.NAPTestRefId = nil
  case "Device":
   n.Device = nil
  case "ParticipationText":
   n.ParticipationText = nil
  case "PSIOtherIdMatch":
   n.PSIOtherIdMatch = nil
  case "DOBRange":
   n.DOBRange = nil
  case "ReportingSchoolName":
   n.ReportingSchoolName = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "NAPJurisdiction":
   n.NAPJurisdiction = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "SchoolGeolocation":
   n.SchoolGeolocation = nil
  case "Adjustment":
   n.Adjustment = nil
  case "Date":
   n.Date = nil
  case "NAPTestLocalId":
   n.NAPTestLocalId = nil
  case "ExemptionReason":
   n.ExemptionReason = nil
  case "StartTime":
   n.StartTime = nil
  case "System":
   n.System = nil
  case "RefId":
   n.RefId = nil
  case "ParticipationCode":
   n.ParticipationCode = nil
  case "TestDisruptionList":
   n.TestDisruptionList = nil
  case "PersonalDetailsChanged":
   n.PersonalDetailsChanged = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "PlatformStudentIdentifier":
   n.PlatformStudentIdentifier = nil
  case "SchoolSector":
   n.SchoolSector = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "PossibleDuplicate":
   n.PossibleDuplicate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPEventStudentLink")
        }
        return n
}

func (n *NAPEventStudentLink) Set(key string, value interface{}) *NAPEventStudentLink {
        if n == nil {
                n = NAPEventStudentLinkCreate(NAPEventStudentLink{})
        }
        switch key {
    case "LapsedTimeTest":
                      n.LapsedTimeTest = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
    case "Device":
                      n.Device = StringCreate(value.(string))
    case "ParticipationText":
                      n.ParticipationText = StringCreate(value.(string))
    case "PSIOtherIdMatch":
                      n.PSIOtherIdMatch = BoolCreate(value.(bool))
    case "DOBRange":
                      n.DOBRange = BoolCreate(value.(bool))
    case "ReportingSchoolName":
                      n.ReportingSchoolName = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "Adjustment":
                      n.Adjustment = AdjustmentContainerTypeCreate(value.(AdjustmentContainerType))
    case "Date":
                      n.Date = StringCreate(value.(string))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ExemptionReason":
                      n.ExemptionReason = StringCreate(value.(string))
    case "StartTime":
                      n.StartTime = StringCreate(value.(string))
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
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
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
    case "TestDisruptionList":
                      n.TestDisruptionList = TestDisruptionListTypeCreate(value.(TestDisruptionListType))
    case "PersonalDetailsChanged":
                      n.PersonalDetailsChanged = BoolCreate(value.(bool))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "PlatformStudentIdentifier":
    
                      n.PlatformStudentIdentifier = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "PossibleDuplicate":
                      n.PossibleDuplicate = BoolCreate(value.(bool))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPEventStudentLink")
        }
        return n
}


func (t *FQReportingType) CopyString(key string, value interface{}) *FQReportingType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FQReportingType) Unset(key string) *FQReportingType {
        switch key {
  case "LocalId":
   n.LocalId = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "EntityName":
   n.EntityName = nil
  case "CommonwealthId":
   n.CommonwealthId = nil
  case "EntityLevel":
   n.EntityLevel = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "FQItemList":
   n.FQItemList = nil
  case "AGRuleList":
   n.AGRuleList = nil
  case "ACARAId":
   n.ACARAId = nil
  case "EntityContact":
   n.EntityContact = nil
  case "FQContextualQuestionList":
   n.FQContextualQuestionList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQReportingType")
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
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "EntityName":
                      n.EntityName = StringCreate(value.(string))
    case "CommonwealthId":
                      n.CommonwealthId = StringCreate(value.(string))
    case "EntityLevel":
                      n.EntityLevel = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "FQItemList":
                      n.FQItemList = FQItemListTypeCreate(value.(FQItemListType))
    case "AGRuleList":
                      n.AGRuleList = AGRuleListTypeCreate(value.(AGRuleListType))
    case "ACARAId":
                      n.ACARAId = StringCreate(value.(string))
    case "EntityContact":
                      n.EntityContact = EntityContactInfoTypeCreate(value.(EntityContactInfoType))
    case "FQContextualQuestionList":
                      n.FQContextualQuestionList = FQContextualQuestionListTypeCreate(value.(FQContextualQuestionListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FQReportingType")
        }
        return n
}


func (t *ResourceBooking) CopyString(key string, value interface{}) *ResourceBooking {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceBooking) Unset(key string) *ResourceBooking {
        switch key {
  case "FinishDateTime":
   n.FinishDateTime = nil
  case "ResourceRefId":
   n.ResourceRefId = nil
  case "RefId":
   n.RefId = nil
  case "ScheduledActivityRefId":
   n.ScheduledActivityRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ResourceLocalId":
   n.ResourceLocalId = nil
  case "Reason":
   n.Reason = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StartDateTime":
   n.StartDateTime = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "ToPeriod":
   n.ToPeriod = nil
  case "KeepOld":
   n.KeepOld = nil
  case "FromPeriod":
   n.FromPeriod = nil
  case "Booker":
   n.Booker = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceBooking")
        }
        return n
}

func (n *ResourceBooking) Set(key string, value interface{}) *ResourceBooking {
        if n == nil {
                n = ResourceBookingCreate(ResourceBooking{})
        }
        switch key {
    case "FinishDateTime":
                      n.FinishDateTime = StringCreate(value.(string))
    case "ResourceRefId":
                      n.ResourceRefId = ResourceBooking_ResourceRefIdCreate(value.(ResourceBooking_ResourceRefId))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ScheduledActivityRefId":
                      n.ScheduledActivityRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ResourceLocalId":
    
                      n.ResourceLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Reason":
                      n.Reason = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StartDateTime":
                      n.StartDateTime = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ToPeriod":
    
                      n.ToPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "KeepOld":
                      n.KeepOld = BoolCreate(value.(bool))
    case "FromPeriod":
    
                      n.FromPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "Booker":
                      n.Booker = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceBooking")
        }
        return n
}


func (t *PublishingPermissionType) CopyString(key string, value interface{}) *PublishingPermissionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PublishingPermissionType) Unset(key string) *PublishingPermissionType {
        switch key {
  case "PermissionValue":
   n.PermissionValue = nil
  case "PermissionCategory":
   n.PermissionCategory = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PublishingPermissionType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PublishingPermissionType")
        }
        return n
}


func (t *YearLevelType) CopyString(key string, value interface{}) *YearLevelType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *YearLevelType) Unset(key string) *YearLevelType {
        switch key {
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearLevelType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearLevelType")
        }
        return n
}


func (t *NAPTestlet) CopyString(key string, value interface{}) *NAPTestlet {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestlet) Unset(key string) *NAPTestlet {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TestItemList":
   n.TestItemList = nil
  case "RefId":
   n.RefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "NAPTestLocalId":
   n.NAPTestLocalId = nil
  case "TestletContent":
   n.TestletContent = nil
  case "NAPTestRefId":
   n.NAPTestRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestlet")
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TestItemList":
                      n.TestItemList = NAPTestItemListTypeCreate(value.(NAPTestItemListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestletContent":
                      n.TestletContent = NAPTestletContentTypeCreate(value.(NAPTestletContentType))
    case "NAPTestRefId":
                      n.NAPTestRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestlet")
        }
        return n
}


func (t *StudentParticipation_ManagingSchool) CopyString(key string, value interface{}) *StudentParticipation_ManagingSchool {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentParticipation_ManagingSchool) Unset(key string) *StudentParticipation_ManagingSchool {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentParticipation_ManagingSchool")
        }
        return n
}

func (n *StudentParticipation_ManagingSchool) Set(key string, value interface{}) *StudentParticipation_ManagingSchool {
        if n == nil {
                n = StudentParticipation_ManagingSchoolCreate(StudentParticipation_ManagingSchool{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentParticipation_ManagingSchool")
        }
        return n
}


func (t *LocationType_LocationRefId) CopyString(key string, value interface{}) *LocationType_LocationRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LocationType_LocationRefId) Unset(key string) *LocationType_LocationRefId {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationType_LocationRefId")
        }
        return n
}

func (n *LocationType_LocationRefId) Set(key string, value interface{}) *LocationType_LocationRefId {
        if n == nil {
                n = LocationType_LocationRefIdCreate(LocationType_LocationRefId{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationType_LocationRefId")
        }
        return n
}


func (t *YearLevelEnrollmentType) CopyString(key string, value interface{}) *YearLevelEnrollmentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *YearLevelEnrollmentType) Unset(key string) *YearLevelEnrollmentType {
        switch key {
  case "Year":
   n.Year = nil
  case "Enrollment":
   n.Enrollment = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearLevelEnrollmentType")
        }
        return n
}

func (n *YearLevelEnrollmentType) Set(key string, value interface{}) *YearLevelEnrollmentType {
        if n == nil {
                n = YearLevelEnrollmentTypeCreate(YearLevelEnrollmentType{})
        }
        switch key {
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
    case "Enrollment":
                      n.Enrollment = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "YearLevelEnrollmentType")
        }
        return n
}


func (t *LocationType) CopyString(key string, value interface{}) *LocationType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LocationType) Unset(key string) *LocationType {
        switch key {
  case "Type":
   n.Type = nil
  case "LocationName":
   n.LocationName = nil
  case "LocationRefId":
   n.LocationRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocationType")
        }
        return n
}


func (t *AbstractContentPackageType_Reference) CopyString(key string, value interface{}) *AbstractContentPackageType_Reference {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentPackageType_Reference) Unset(key string) *AbstractContentPackageType_Reference {
        switch key {
  case "Description":
   n.Description = nil
  case "URL":
   n.URL = nil
  case "MIMEType":
   n.MIMEType = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_Reference")
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
    case "URL":
                      n.URL = StringCreate(value.(string))
    case "MIMEType":
                      n.MIMEType = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType_Reference")
        }
        return n
}


func (t *FinancialQuestionnaireCollection) CopyString(key string, value interface{}) *FinancialQuestionnaireCollection {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *FinancialQuestionnaireCollection) Unset(key string) *FinancialQuestionnaireCollection {
        switch key {
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "FQYear":
   n.FQYear = nil
  case "SoftwareVendorInfo":
   n.SoftwareVendorInfo = nil
  case "RoundCode":
   n.RoundCode = nil
  case "RefId":
   n.RefId = nil
  case "FQReportingList":
   n.FQReportingList = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "ReportingAuthorityCommonwealthId":
   n.ReportingAuthorityCommonwealthId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FinancialQuestionnaireCollection")
        }
        return n
}

func (n *FinancialQuestionnaireCollection) Set(key string, value interface{}) *FinancialQuestionnaireCollection {
        if n == nil {
                n = FinancialQuestionnaireCollectionCreate(FinancialQuestionnaireCollection{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "FQYear":
    
                      n.FQYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SoftwareVendorInfo":
                      n.SoftwareVendorInfo = SoftwareVendorInfoContainerTypeCreate(value.(SoftwareVendorInfoContainerType))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "FQReportingList":
                      n.FQReportingList = FQReportingListTypeCreate(value.(FQReportingListType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "FinancialQuestionnaireCollection")
        }
        return n
}


func (t *CensusStaffType) CopyString(key string, value interface{}) *CensusStaffType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CensusStaffType) Unset(key string) *CensusStaffType {
        switch key {
  case "PrimaryFTE":
   n.PrimaryFTE = nil
  case "JobFTE":
   n.JobFTE = nil
  case "StaffCohortId":
   n.StaffCohortId = nil
  case "Headcount":
   n.Headcount = nil
  case "CohortIndigenousType":
   n.CohortIndigenousType = nil
  case "SecondaryFTE":
   n.SecondaryFTE = nil
  case "CohortGender":
   n.CohortGender = nil
  case "StaffActivity":
   n.StaffActivity = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusStaffType")
        }
        return n
}

func (n *CensusStaffType) Set(key string, value interface{}) *CensusStaffType {
        if n == nil {
                n = CensusStaffTypeCreate(CensusStaffType{})
        }
        switch key {
    case "PrimaryFTE":
                      n.PrimaryFTE = FloatCreate(value.(float64))
    case "JobFTE":
                      n.JobFTE = FloatCreate(value.(float64))
    case "StaffCohortId":
    
                      n.StaffCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Headcount":
                      n.Headcount = IntCreate(value.(int))
    case "CohortIndigenousType":
                      n.CohortIndigenousType = StringCreate(value.(string))
    case "SecondaryFTE":
                      n.SecondaryFTE = FloatCreate(value.(float64))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "StaffActivity":
                      n.StaffActivity = StaffActivityExtensionTypeCreate(value.(StaffActivityExtensionType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CensusStaffType")
        }
        return n
}


func (t *NAPTestContentType) CopyString(key string, value interface{}) *NAPTestContentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestContentType) Unset(key string) *NAPTestContentType {
        switch key {
  case "TestName":
   n.TestName = nil
  case "TestLevel":
   n.TestLevel = nil
  case "DomainProficiency":
   n.DomainProficiency = nil
  case "NAPTestLocalId":
   n.NAPTestLocalId = nil
  case "DomainBands":
   n.DomainBands = nil
  case "Domain":
   n.Domain = nil
  case "TestYear":
   n.TestYear = nil
  case "TestType":
   n.TestType = nil
  case "StagesCount":
   n.StagesCount = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestContentType")
        }
        return n
}

func (n *NAPTestContentType) Set(key string, value interface{}) *NAPTestContentType {
        if n == nil {
                n = NAPTestContentTypeCreate(NAPTestContentType{})
        }
        switch key {
    case "TestName":
                      n.TestName = StringCreate(value.(string))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "DomainProficiency":
                      n.DomainProficiency = DomainProficiencyContainerTypeCreate(value.(DomainProficiencyContainerType))
    case "NAPTestLocalId":
    
                      n.NAPTestLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DomainBands":
                      n.DomainBands = DomainBandsContainerTypeCreate(value.(DomainBandsContainerType))
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
    case "TestYear":
    
                      n.TestYear = ((*SchoolYearType)(StringCreate(value.(string))))
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
    case "StagesCount":
                      n.StagesCount = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestContentType")
        }
        return n
}


func (t *AggregateStatisticInfo_CalculationRule) CopyString(key string, value interface{}) *AggregateStatisticInfo_CalculationRule {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AggregateStatisticInfo_CalculationRule) Unset(key string) *AggregateStatisticInfo_CalculationRule {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticInfo_CalculationRule")
        }
        return n
}

func (n *AggregateStatisticInfo_CalculationRule) Set(key string, value interface{}) *AggregateStatisticInfo_CalculationRule {
        if n == nil {
                n = AggregateStatisticInfo_CalculationRuleCreate(AggregateStatisticInfo_CalculationRule{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AggregateStatisticInfo_CalculationRule")
        }
        return n
}


func (t *NAPTestItem) CopyString(key string, value interface{}) *NAPTestItem {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestItem) Unset(key string) *NAPTestItem {
        switch key {
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "RefId":
   n.RefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TestItemContent":
   n.TestItemContent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItem")
        }
        return n
}

func (n *NAPTestItem) Set(key string, value interface{}) *NAPTestItem {
        if n == nil {
                n = NAPTestItemCreate(NAPTestItem{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TestItemContent":
                      n.TestItemContent = NAPTestItemContentTypeCreate(value.(NAPTestItemContentType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItem")
        }
        return n
}


func (t *MapReferenceType) CopyString(key string, value interface{}) *MapReferenceType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MapReferenceType) Unset(key string) *MapReferenceType {
        switch key {
  case "Type":
   n.Type = nil
  case "YCoordinate":
   n.YCoordinate = nil
  case "XCoordinate":
   n.XCoordinate = nil
  case "MapNumber":
   n.MapNumber = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MapReferenceType")
        }
        return n
}

func (n *MapReferenceType) Set(key string, value interface{}) *MapReferenceType {
        if n == nil {
                n = MapReferenceTypeCreate(MapReferenceType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "YCoordinate":
                      n.YCoordinate = StringCreate(value.(string))
    case "XCoordinate":
                      n.XCoordinate = StringCreate(value.(string))
    case "MapNumber":
                      n.MapNumber = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MapReferenceType")
        }
        return n
}


func (t *AssociatedObjectsType_AssociatedObject) CopyString(key string, value interface{}) *AssociatedObjectsType_AssociatedObject {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AssociatedObjectsType_AssociatedObject) Unset(key string) *AssociatedObjectsType_AssociatedObject {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AssociatedObjectsType_AssociatedObject")
        }
        return n
}

func (n *AssociatedObjectsType_AssociatedObject) Set(key string, value interface{}) *AssociatedObjectsType_AssociatedObject {
        if n == nil {
                n = AssociatedObjectsType_AssociatedObjectCreate(AssociatedObjectsType_AssociatedObject{})
        }
        switch key {
    case "Value":
                      n.Value = StringCreate(value.(string))
    case "SIF_RefObject":
    
                      n.SIF_RefObject = ((*ObjectNameType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AssociatedObjectsType_AssociatedObject")
        }
        return n
}


func (t *MonetaryAmountType) CopyString(key string, value interface{}) *MonetaryAmountType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MonetaryAmountType) Unset(key string) *MonetaryAmountType {
        switch key {
  case "Value":
   n.Value = nil
  case "Currency":
   n.Currency = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MonetaryAmountType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MonetaryAmountType")
        }
        return n
}


func (t *EquipmentInfo_SIF_RefId) CopyString(key string, value interface{}) *EquipmentInfo_SIF_RefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EquipmentInfo_SIF_RefId) Unset(key string) *EquipmentInfo_SIF_RefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EquipmentInfo_SIF_RefId")
        }
        return n
}

func (n *EquipmentInfo_SIF_RefId) Set(key string, value interface{}) *EquipmentInfo_SIF_RefId {
        if n == nil {
                n = EquipmentInfo_SIF_RefIdCreate(EquipmentInfo_SIF_RefId{})
        }
        switch key {
    case "SIF_RefObject":
                      n.SIF_RefObject = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EquipmentInfo_SIF_RefId")
        }
        return n
}


func (t *TeachingGroupStudentType) CopyString(key string, value interface{}) *TeachingGroupStudentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TeachingGroupStudentType) Unset(key string) *TeachingGroupStudentType {
        switch key {
  case "StudentLocalId":
   n.StudentLocalId = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "Name":
   n.Name = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupStudentType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TeachingGroupStudentType")
        }
        return n
}


func (t *SourceObjectsType_SourceObject) CopyString(key string, value interface{}) *SourceObjectsType_SourceObject {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SourceObjectsType_SourceObject) Unset(key string) *SourceObjectsType_SourceObject {
        switch key {
  case "Value":
   n.Value = nil
  case "SIF_RefObject":
   n.SIF_RefObject = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SourceObjectsType_SourceObject")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SourceObjectsType_SourceObject")
        }
        return n
}


func (t *StudentActivityType) CopyString(key string, value interface{}) *StudentActivityType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentActivityType) Unset(key string) *StudentActivityType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityType")
        }
        return n
}


func (t *WellbeingPlanType) CopyString(key string, value interface{}) *WellbeingPlanType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *WellbeingPlanType) Unset(key string) *WellbeingPlanType {
        switch key {
  case "PersonalisedPlanRefId":
   n.PersonalisedPlanRefId = nil
  case "PlanNotes":
   n.PlanNotes = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPlanType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "WellbeingPlanType")
        }
        return n
}


func (t *SectionInfo) CopyString(key string, value interface{}) *SectionInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SectionInfo) Unset(key string) *SectionInfo {
        switch key {
  case "SectionCode":
   n.SectionCode = nil
  case "RefId":
   n.RefId = nil
  case "LocationOfInstruction":
   n.LocationOfInstruction = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "LanguageOfInstruction":
   n.LanguageOfInstruction = nil
  case "SchoolCourseInfoRefId":
   n.SchoolCourseInfoRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TermInfoRefId":
   n.TermInfoRefId = nil
  case "CourseSectionCode":
   n.CourseSectionCode = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "MediumOfInstruction":
   n.MediumOfInstruction = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "Description":
   n.Description = nil
  case "LocalId":
   n.LocalId = nil
  case "CountForAttendance":
   n.CountForAttendance = nil
  case "SchoolCourseInfoOverride":
   n.SchoolCourseInfoOverride = nil
  case "SummerSchool":
   n.SummerSchool = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SectionInfo")
        }
        return n
}

func (n *SectionInfo) Set(key string, value interface{}) *SectionInfo {
        if n == nil {
                n = SectionInfoCreate(SectionInfo{})
        }
        switch key {
    case "SectionCode":
                      n.SectionCode = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocationOfInstruction":
                      n.LocationOfInstruction = LocationOfInstructionTypeCreate(value.(LocationOfInstructionType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LanguageOfInstruction":
                      n.LanguageOfInstruction = LanguageOfInstructionTypeCreate(value.(LanguageOfInstructionType))
    case "SchoolCourseInfoRefId":
                      n.SchoolCourseInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "CourseSectionCode":
                      n.CourseSectionCode = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "MediumOfInstruction":
                      n.MediumOfInstruction = MediumOfInstructionTypeCreate(value.(MediumOfInstructionType))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CountForAttendance":
                      n.CountForAttendance = StringCreate(value.(string))
    case "SchoolCourseInfoOverride":
                      n.SchoolCourseInfoOverride = SchoolCourseInfoOverrideTypeCreate(value.(SchoolCourseInfoOverrideType))
    case "SummerSchool":
                      n.SummerSchool = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SectionInfo")
        }
        return n
}


func (t *EducationFilterType) CopyString(key string, value interface{}) *EducationFilterType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EducationFilterType) Unset(key string) *EducationFilterType {
        switch key {
  case "LearningStandardItems":
   n.LearningStandardItems = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EducationFilterType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EducationFilterType")
        }
        return n
}


func (t *StatsCohortType) CopyString(key string, value interface{}) *StatsCohortType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StatsCohortType) Unset(key string) *StatsCohortType {
        switch key {
  case "StatsCohortId":
   n.StatsCohortId = nil
  case "AttendanceGTE90Percent":
   n.AttendanceGTE90Percent = nil
  case "CohortGender":
   n.CohortGender = nil
  case "AttendanceDays":
   n.AttendanceDays = nil
  case "PossibleSchoolDaysGT90PercentAttendance":
   n.PossibleSchoolDaysGT90PercentAttendance = nil
  case "PossibleSchoolDays":
   n.PossibleSchoolDays = nil
  case "StatsIndigenousStudentType":
   n.StatsIndigenousStudentType = nil
  case "DaysInReferencePeriod":
   n.DaysInReferencePeriod = nil
  case "AttendanceLess90Percent":
   n.AttendanceLess90Percent = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatsCohortType")
        }
        return n
}

func (n *StatsCohortType) Set(key string, value interface{}) *StatsCohortType {
        if n == nil {
                n = StatsCohortTypeCreate(StatsCohortType{})
        }
        switch key {
    case "StatsCohortId":
    
                      n.StatsCohortId = ((*LocalIdType)(StringCreate(value.(string))))
    case "AttendanceGTE90Percent":
                      n.AttendanceGTE90Percent = IntCreate(value.(int))
    case "CohortGender":
                      n.CohortGender = StringCreate(value.(string))
    case "AttendanceDays":
                      n.AttendanceDays = FloatCreate(value.(float64))
    case "PossibleSchoolDaysGT90PercentAttendance":
                      n.PossibleSchoolDaysGT90PercentAttendance = IntCreate(value.(int))
    case "PossibleSchoolDays":
                      n.PossibleSchoolDays = IntCreate(value.(int))
    case "StatsIndigenousStudentType":
                      n.StatsIndigenousStudentType = StringCreate(value.(string))
    case "DaysInReferencePeriod":
                      n.DaysInReferencePeriod = IntCreate(value.(int))
    case "AttendanceLess90Percent":
                      n.AttendanceLess90Percent = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StatsCohortType")
        }
        return n
}


func (t *StudentGrade) CopyString(key string, value interface{}) *StudentGrade {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentGrade) Unset(key string) *StudentGrade {
        switch key {
  case "RefId":
   n.RefId = nil
  case "StaffPersonalRefId":
   n.StaffPersonalRefId = nil
  case "Markers":
   n.Markers = nil
  case "TeachingGroupShortName":
   n.TeachingGroupShortName = nil
  case "LearningArea":
   n.LearningArea = nil
  case "Grade":
   n.Grade = nil
  case "YearLevel":
   n.YearLevel = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "TermInfoRefId":
   n.TermInfoRefId = nil
  case "TeachingGroupRefId":
   n.TeachingGroupRefId = nil
  case "Description":
   n.Description = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "LearningStandardList":
   n.LearningStandardList = nil
  case "TeacherJudgement":
   n.TeacherJudgement = nil
  case "Homegroup":
   n.Homegroup = nil
  case "GradingScoreList":
   n.GradingScoreList = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentGrade")
        }
        return n
}

func (n *StudentGrade) Set(key string, value interface{}) *StudentGrade {
        if n == nil {
                n = StudentGradeCreate(StudentGrade{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "StaffPersonalRefId":
                      n.StaffPersonalRefId = StringCreate(value.(string))
    case "Markers":
                      n.Markers = StudentGradeMarkersListTypeCreate(value.(StudentGradeMarkersListType))
    case "TeachingGroupShortName":
                      n.TeachingGroupShortName = StringCreate(value.(string))
    case "LearningArea":
                      n.LearningArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "Grade":
                      n.Grade = GradeTypeCreate(value.(GradeType))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "TermInfoRefId":
                      n.TermInfoRefId = StringCreate(value.(string))
    case "TeachingGroupRefId":
                      n.TeachingGroupRefId = StringCreate(value.(string))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LearningStandardList":
                      n.LearningStandardList = LearningStandardListTypeCreate(value.(LearningStandardListType))
    case "TeacherJudgement":
                      n.TeacherJudgement = StringCreate(value.(string))
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
    case "GradingScoreList":
                      n.GradingScoreList = GradingScoreListTypeCreate(value.(GradingScoreListType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentGrade")
        }
        return n
}


func (t *ResourceUsage_SIF_RefId) CopyString(key string, value interface{}) *ResourceUsage_SIF_RefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceUsage_SIF_RefId) Unset(key string) *ResourceUsage_SIF_RefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_SIF_RefId")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_SIF_RefId")
        }
        return n
}


func (t *HouseholdContactInfoType) CopyString(key string, value interface{}) *HouseholdContactInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *HouseholdContactInfoType) Unset(key string) *HouseholdContactInfoType {
        switch key {
  case "HouseholdContactId":
   n.HouseholdContactId = nil
  case "EmailList":
   n.EmailList = nil
  case "AddressList":
   n.AddressList = nil
  case "PhoneNumberList":
   n.PhoneNumberList = nil
  case "HouseholdSalutation":
   n.HouseholdSalutation = nil
  case "PreferenceNumber":
   n.PreferenceNumber = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "HouseholdContactInfoType")
        }
        return n
}

func (n *HouseholdContactInfoType) Set(key string, value interface{}) *HouseholdContactInfoType {
        if n == nil {
                n = HouseholdContactInfoTypeCreate(HouseholdContactInfoType{})
        }
        switch key {
    case "HouseholdContactId":
    
                      n.HouseholdContactId = ((*LocalIdType)(StringCreate(value.(string))))
    case "EmailList":
                      n.EmailList = EmailListTypeCreate(value.(EmailListType))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "PhoneNumberList":
                      n.PhoneNumberList = PhoneNumberListTypeCreate(value.(PhoneNumberListType))
    case "HouseholdSalutation":
                      n.HouseholdSalutation = StringCreate(value.(string))
    case "PreferenceNumber":
                      n.PreferenceNumber = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "HouseholdContactInfoType")
        }
        return n
}


func (t *DomainBandsContainerType) CopyString(key string, value interface{}) *DomainBandsContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *DomainBandsContainerType) Unset(key string) *DomainBandsContainerType {
        switch key {
  case "Band1Lower":
   n.Band1Lower = nil
  case "Band7Upper":
   n.Band7Upper = nil
  case "Band2Lower":
   n.Band2Lower = nil
  case "Band3Lower":
   n.Band3Lower = nil
  case "Band6Upper":
   n.Band6Upper = nil
  case "Band10Upper":
   n.Band10Upper = nil
  case "Band9Upper":
   n.Band9Upper = nil
  case "Band4Lower":
   n.Band4Lower = nil
  case "Band5Upper":
   n.Band5Upper = nil
  case "Band8Lower":
   n.Band8Lower = nil
  case "Band3Upper":
   n.Band3Upper = nil
  case "Band2Upper":
   n.Band2Upper = nil
  case "Band7Lower":
   n.Band7Lower = nil
  case "Band10Lower":
   n.Band10Lower = nil
  case "Band1Upper":
   n.Band1Upper = nil
  case "Band8Upper":
   n.Band8Upper = nil
  case "Band5Lower":
   n.Band5Lower = nil
  case "Band9Lower":
   n.Band9Lower = nil
  case "Band4Upper":
   n.Band4Upper = nil
  case "Band6Lower":
   n.Band6Lower = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainBandsContainerType")
        }
        return n
}

func (n *DomainBandsContainerType) Set(key string, value interface{}) *DomainBandsContainerType {
        if n == nil {
                n = DomainBandsContainerTypeCreate(DomainBandsContainerType{})
        }
        switch key {
    case "Band1Lower":
                      n.Band1Lower = FloatCreate(value.(float64))
    case "Band7Upper":
                      n.Band7Upper = FloatCreate(value.(float64))
    case "Band2Lower":
                      n.Band2Lower = FloatCreate(value.(float64))
    case "Band3Lower":
                      n.Band3Lower = FloatCreate(value.(float64))
    case "Band6Upper":
                      n.Band6Upper = FloatCreate(value.(float64))
    case "Band10Upper":
                      n.Band10Upper = FloatCreate(value.(float64))
    case "Band9Upper":
                      n.Band9Upper = FloatCreate(value.(float64))
    case "Band4Lower":
                      n.Band4Lower = FloatCreate(value.(float64))
    case "Band5Upper":
                      n.Band5Upper = FloatCreate(value.(float64))
    case "Band8Lower":
                      n.Band8Lower = FloatCreate(value.(float64))
    case "Band3Upper":
                      n.Band3Upper = FloatCreate(value.(float64))
    case "Band2Upper":
                      n.Band2Upper = FloatCreate(value.(float64))
    case "Band7Lower":
                      n.Band7Lower = FloatCreate(value.(float64))
    case "Band10Lower":
                      n.Band10Lower = FloatCreate(value.(float64))
    case "Band1Upper":
                      n.Band1Upper = FloatCreate(value.(float64))
    case "Band8Upper":
                      n.Band8Upper = FloatCreate(value.(float64))
    case "Band5Lower":
                      n.Band5Lower = FloatCreate(value.(float64))
    case "Band9Lower":
                      n.Band9Lower = FloatCreate(value.(float64))
    case "Band4Upper":
                      n.Band4Upper = FloatCreate(value.(float64))
    case "Band6Lower":
                      n.Band6Lower = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "DomainBandsContainerType")
        }
        return n
}


func (t *SchoolPrograms) CopyString(key string, value interface{}) *SchoolPrograms {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SchoolPrograms) Unset(key string) *SchoolPrograms {
        switch key {
  case "SchoolProgramList":
   n.SchoolProgramList = nil
  case "RefId":
   n.RefId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolPrograms")
        }
        return n
}

func (n *SchoolPrograms) Set(key string, value interface{}) *SchoolPrograms {
        if n == nil {
                n = SchoolProgramsCreate(SchoolPrograms{})
        }
        switch key {
    case "SchoolProgramList":
                      n.SchoolProgramList = SchoolProgramListTypeCreate(value.(SchoolProgramListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SchoolPrograms")
        }
        return n
}


func (t *CopyRightContainerType) CopyString(key string, value interface{}) *CopyRightContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CopyRightContainerType) Unset(key string) *CopyRightContainerType {
        switch key {
  case "Holder":
   n.Holder = nil
  case "Date":
   n.Date = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CopyRightContainerType")
        }
        return n
}

func (n *CopyRightContainerType) Set(key string, value interface{}) *CopyRightContainerType {
        if n == nil {
                n = CopyRightContainerTypeCreate(CopyRightContainerType{})
        }
        switch key {
    case "Holder":
                      n.Holder = StringCreate(value.(string))
    case "Date":
                      n.Date = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CopyRightContainerType")
        }
        return n
}


func (t *StudentPersonal) CopyString(key string, value interface{}) *StudentPersonal {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentPersonal) Unset(key string) *StudentPersonal {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "MostRecent":
   n.MostRecent = nil
  case "PrePrimaryEducationHours":
   n.PrePrimaryEducationHours = nil
  case "GraduationDate":
   n.GraduationDate = nil
  case "ProjectedGraduationYear":
   n.ProjectedGraduationYear = nil
  case "PrePrimaryEducation":
   n.PrePrimaryEducation = nil
  case "OnTimeGraduationYear":
   n.OnTimeGraduationYear = nil
  case "StateProvinceId":
   n.StateProvinceId = nil
  case "ESLDateAssessed":
   n.ESLDateAssessed = nil
  case "EconomicDisadvantage":
   n.EconomicDisadvantage = nil
  case "PersonInfo":
   n.PersonInfo = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AcceptableUsePolicy":
   n.AcceptableUsePolicy = nil
  case "IntegrationAide":
   n.IntegrationAide = nil
  case "MedicalAlertMessages":
   n.MedicalAlertMessages = nil
  case "YoungCarersRole":
   n.YoungCarersRole = nil
  case "HomeSchooledStudent":
   n.HomeSchooledStudent = nil
  case "Sensitive":
   n.Sensitive = nil
  case "RefId":
   n.RefId = nil
  case "ESL":
   n.ESL = nil
  case "OtherIdList":
   n.OtherIdList = nil
  case "GiftedTalented":
   n.GiftedTalented = nil
  case "ElectronicIdList":
   n.ElectronicIdList = nil
  case "EducationSupport":
   n.EducationSupport = nil
  case "LocalId":
   n.LocalId = nil
  case "FirstAUSchoolEnrollment":
   n.FirstAUSchoolEnrollment = nil
  case "ESLSupport":
   n.ESLSupport = nil
  case "AlertMessages":
   n.AlertMessages = nil
  case "OfflineDelivery":
   n.OfflineDelivery = nil
  case "Disability":
   n.Disability = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentPersonal")
        }
        return n
}

func (n *StudentPersonal) Set(key string, value interface{}) *StudentPersonal {
        if n == nil {
                n = StudentPersonalCreate(StudentPersonal{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "MostRecent":
                      n.MostRecent = StudentMostRecentContainerTypeCreate(value.(StudentMostRecentContainerType))
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
    case "GraduationDate":
    
                      n.GraduationDate = ((*GraduationDateType)(StringCreate(value.(string))))
    case "ProjectedGraduationYear":
    
                      n.ProjectedGraduationYear = ((*ProjectedGraduationYearType)(StringCreate(value.(string))))
    case "PrePrimaryEducation":
                      n.PrePrimaryEducation = StringCreate(value.(string))
    case "OnTimeGraduationYear":
    
                      n.OnTimeGraduationYear = ((*OnTimeGraduationYearType)(StringCreate(value.(string))))
    case "StateProvinceId":
    
                      n.StateProvinceId = ((*StateProvinceIdType)(StringCreate(value.(string))))
    case "ESLDateAssessed":
                      n.ESLDateAssessed = StringCreate(value.(string))
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
    case "PersonInfo":
                      n.PersonInfo = PersonInfoTypeCreate(value.(PersonInfoType))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
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
    case "MedicalAlertMessages":
                      n.MedicalAlertMessages = MedicalAlertMessagesTypeCreate(value.(MedicalAlertMessagesType))
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
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
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
    case "OtherIdList":
                      n.OtherIdList = OtherIdListTypeCreate(value.(OtherIdListType))
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
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
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
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "FirstAUSchoolEnrollment":
                      n.FirstAUSchoolEnrollment = StringCreate(value.(string))
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
    case "AlertMessages":
                      n.AlertMessages = AlertMessagesTypeCreate(value.(AlertMessagesType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentPersonal")
        }
        return n
}


func (t *CalendarSummary) CopyString(key string, value interface{}) *CalendarSummary {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CalendarSummary) Unset(key string) *CalendarSummary {
        switch key {
  case "RefId":
   n.RefId = nil
  case "InstructionalMinutes":
   n.InstructionalMinutes = nil
  case "EndDate":
   n.EndDate = nil
  case "LastInstructionDate":
   n.LastInstructionDate = nil
  case "YearLevels":
   n.YearLevels = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "GraduationDate":
   n.GraduationDate = nil
  case "DaysInSession":
   n.DaysInSession = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "MinutesPerDay":
   n.MinutesPerDay = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Description":
   n.Description = nil
  case "StartDate":
   n.StartDate = nil
  case "FirstInstructionDate":
   n.FirstInstructionDate = nil
  case "LocalId":
   n.LocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarSummary")
        }
        return n
}

func (n *CalendarSummary) Set(key string, value interface{}) *CalendarSummary {
        if n == nil {
                n = CalendarSummaryCreate(CalendarSummary{})
        }
        switch key {
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "InstructionalMinutes":
                      n.InstructionalMinutes = IntCreate(value.(int))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "LastInstructionDate":
                      n.LastInstructionDate = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "GraduationDate":
    
                      n.GraduationDate = ((*GraduationDateType)(StringCreate(value.(string))))
    case "DaysInSession":
                      n.DaysInSession = IntCreate(value.(int))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "MinutesPerDay":
                      n.MinutesPerDay = IntCreate(value.(int))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "FirstInstructionDate":
                      n.FirstInstructionDate = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CalendarSummary")
        }
        return n
}


func (t *VendorInfo) CopyString(key string, value interface{}) *VendorInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *VendorInfo) Unset(key string) *VendorInfo {
        switch key {
  case "Name":
   n.Name = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "ABN":
   n.ABN = nil
  case "LocalId":
   n.LocalId = nil
  case "BSB":
   n.BSB = nil
  case "RefId":
   n.RefId = nil
  case "BPay":
   n.BPay = nil
  case "PaymentTerms":
   n.PaymentTerms = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "AccountName":
   n.AccountName = nil
  case "CustomerId":
   n.CustomerId = nil
  case "ContactInfo":
   n.ContactInfo = nil
  case "RegisteredForGST":
   n.RegisteredForGST = nil
  case "AccountNumber":
   n.AccountNumber = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "VendorInfo")
        }
        return n
}

func (n *VendorInfo) Set(key string, value interface{}) *VendorInfo {
        if n == nil {
                n = VendorInfoCreate(VendorInfo{})
        }
        switch key {
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ABN":
                      n.ABN = StringCreate(value.(string))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "BSB":
                      n.BSB = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "BPay":
                      n.BPay = StringCreate(value.(string))
    case "PaymentTerms":
                      n.PaymentTerms = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "AccountName":
                      n.AccountName = StringCreate(value.(string))
    case "CustomerId":
                      n.CustomerId = StringCreate(value.(string))
    case "ContactInfo":
                      n.ContactInfo = ContactInfoTypeCreate(value.(ContactInfoType))
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
    case "AccountNumber":
                      n.AccountNumber = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "VendorInfo")
        }
        return n
}


func (t *LocalCodeType) CopyString(key string, value interface{}) *LocalCodeType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LocalCodeType) Unset(key string) *LocalCodeType {
        switch key {
  case "ListIndex":
   n.ListIndex = nil
  case "Description":
   n.Description = nil
  case "Element":
   n.Element = nil
  case "LocalisedCode":
   n.LocalisedCode = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocalCodeType")
        }
        return n
}

func (n *LocalCodeType) Set(key string, value interface{}) *LocalCodeType {
        if n == nil {
                n = LocalCodeTypeCreate(LocalCodeType{})
        }
        switch key {
    case "ListIndex":
                      n.ListIndex = IntCreate(value.(int))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Element":
                      n.Element = StringCreate(value.(string))
    case "LocalisedCode":
                      n.LocalisedCode = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LocalCodeType")
        }
        return n
}


func (t *TimeTableScheduleType) CopyString(key string, value interface{}) *TimeTableScheduleType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableScheduleType) Unset(key string) *TimeTableScheduleType {
        switch key {
  case "SchoolYear":
   n.SchoolYear = nil
  case "Title":
   n.Title = nil
  case "TimeTableCreationDate":
   n.TimeTableCreationDate = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "StartDate":
   n.StartDate = nil
  case "TeachingPeriodsPerDay":
   n.TeachingPeriodsPerDay = nil
  case "PeriodsPerDay":
   n.PeriodsPerDay = nil
  case "LocalId":
   n.LocalId = nil
  case "DaysPerCycle":
   n.DaysPerCycle = nil
  case "EndDate":
   n.EndDate = nil
  case "TimeTableDayList":
   n.TimeTableDayList = nil
  case "SchoolName":
   n.SchoolName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableScheduleType")
        }
        return n
}

func (n *TimeTableScheduleType) Set(key string, value interface{}) *TimeTableScheduleType {
        if n == nil {
                n = TimeTableScheduleTypeCreate(TimeTableScheduleType{})
        }
        switch key {
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "TimeTableCreationDate":
                      n.TimeTableCreationDate = StringCreate(value.(string))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "StartDate":
                      n.StartDate = StringCreate(value.(string))
    case "TeachingPeriodsPerDay":
                      n.TeachingPeriodsPerDay = IntCreate(value.(int))
    case "PeriodsPerDay":
                      n.PeriodsPerDay = IntCreate(value.(int))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DaysPerCycle":
                      n.DaysPerCycle = IntCreate(value.(int))
    case "EndDate":
                      n.EndDate = StringCreate(value.(string))
    case "TimeTableDayList":
                      n.TimeTableDayList = TimeTableDayListTypeCreate(value.(TimeTableDayListType))
    case "SchoolName":
                      n.SchoolName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableScheduleType")
        }
        return n
}


func (t *StudentMostRecentContainerType) CopyString(key string, value interface{}) *StudentMostRecentContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentMostRecentContainerType) Unset(key string) *StudentMostRecentContainerType {
        switch key {
  case "Parent2Language":
   n.Parent2Language = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "BoardingStatus":
   n.BoardingStatus = nil
  case "DisabilityCategory":
   n.DisabilityCategory = nil
  case "DistanceEducationStudent":
   n.DistanceEducationStudent = nil
  case "Parent1NonSchoolEducation":
   n.Parent1NonSchoolEducation = nil
  case "LocalCampusId":
   n.LocalCampusId = nil
  case "DisabilityLevelOfAdjustment":
   n.DisabilityLevelOfAdjustment = nil
  case "FTE":
   n.FTE = nil
  case "Parent1EmploymentType":
   n.Parent1EmploymentType = nil
  case "HomeroomLocalId":
   n.HomeroomLocalId = nil
  case "YearLevel":
   n.YearLevel = nil
  case "ClassCode":
   n.ClassCode = nil
  case "SchoolACARAId":
   n.SchoolACARAId = nil
  case "CensusAge":
   n.CensusAge = nil
  case "Parent1SchoolEducationLevel":
   n.Parent1SchoolEducationLevel = nil
  case "Parent2SchoolEducationLevel":
   n.Parent2SchoolEducationLevel = nil
  case "Homegroup":
   n.Homegroup = nil
  case "Parent1Language":
   n.Parent1Language = nil
  case "OtherSchoolName":
   n.OtherSchoolName = nil
  case "OtherEnrollmentSchoolACARAId":
   n.OtherEnrollmentSchoolACARAId = nil
  case "TestLevel":
   n.TestLevel = nil
  case "MembershipType":
   n.MembershipType = nil
  case "ReportingSchoolId":
   n.ReportingSchoolId = nil
  case "FFPOS":
   n.FFPOS = nil
  case "Parent2EmploymentType":
   n.Parent2EmploymentType = nil
  case "Parent2NonSchoolEducation":
   n.Parent2NonSchoolEducation = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentMostRecentContainerType")
        }
        return n
}

func (n *StudentMostRecentContainerType) Set(key string, value interface{}) *StudentMostRecentContainerType {
        if n == nil {
                n = StudentMostRecentContainerTypeCreate(StudentMostRecentContainerType{})
        }
        switch key {
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
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
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
    case "DisabilityCategory":
                      n.DisabilityCategory = StringCreate(value.(string))
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
    case "LocalCampusId":
    
                      n.LocalCampusId = ((*LocalIdType)(StringCreate(value.(string))))
    case "DisabilityLevelOfAdjustment":
                      n.DisabilityLevelOfAdjustment = StringCreate(value.(string))
    case "FTE":
                      n.FTE = FloatCreate(value.(float64))
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
    case "HomeroomLocalId":
    
                      n.HomeroomLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
    case "ClassCode":
                      n.ClassCode = StringCreate(value.(string))
    case "SchoolACARAId":
    
                      n.SchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CensusAge":
                      n.CensusAge = IntCreate(value.(int))
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
    case "Homegroup":
                      n.Homegroup = StringCreate(value.(string))
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
    case "OtherSchoolName":
                      n.OtherSchoolName = StringCreate(value.(string))
    case "OtherEnrollmentSchoolACARAId":
    
                      n.OtherEnrollmentSchoolACARAId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TestLevel":
                      n.TestLevel = YearLevelTypeCreate(value.(YearLevelType))
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
    case "ReportingSchoolId":
    
                      n.ReportingSchoolId = ((*LocalIdType)(StringCreate(value.(string))))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentMostRecentContainerType")
        }
        return n
}


func (t *EquipmentInfo) CopyString(key string, value interface{}) *EquipmentInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *EquipmentInfo) Unset(key string) *EquipmentInfo {
        switch key {
  case "InvoiceRefId":
   n.InvoiceRefId = nil
  case "EquipmentType":
   n.EquipmentType = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "Name":
   n.Name = nil
  case "AssetNumber":
   n.AssetNumber = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Description":
   n.Description = nil
  case "PurchaseOrderRefId":
   n.PurchaseOrderRefId = nil
  case "SIF_RefId":
   n.SIF_RefId = nil
  case "LocalId":
   n.LocalId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EquipmentInfo")
        }
        return n
}

func (n *EquipmentInfo) Set(key string, value interface{}) *EquipmentInfo {
        if n == nil {
                n = EquipmentInfoCreate(EquipmentInfo{})
        }
        switch key {
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
    case "EquipmentType":
                      n.EquipmentType = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "AssetNumber":
    
                      n.AssetNumber = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "PurchaseOrderRefId":
                      n.PurchaseOrderRefId = StringCreate(value.(string))
    case "SIF_RefId":
                      n.SIF_RefId = EquipmentInfo_SIF_RefIdCreate(value.(EquipmentInfo_SIF_RefId))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "EquipmentInfo")
        }
        return n
}


func (t *CatchmentStatusContainerType) CopyString(key string, value interface{}) *CatchmentStatusContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CatchmentStatusContainerType) Unset(key string) *CatchmentStatusContainerType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CatchmentStatusContainerType")
        }
        return n
}

func (n *CatchmentStatusContainerType) Set(key string, value interface{}) *CatchmentStatusContainerType {
        if n == nil {
                n = CatchmentStatusContainerTypeCreate(CatchmentStatusContainerType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CatchmentStatusContainerType")
        }
        return n
}


func (t *PaymentReceipt) CopyString(key string, value interface{}) *PaymentReceipt {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PaymentReceipt) Unset(key string) *PaymentReceipt {
        switch key {
  case "AccountingPeriod":
   n.AccountingPeriod = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "LocalId":
   n.LocalId = nil
  case "TransactionMethod":
   n.TransactionMethod = nil
  case "TransactionDescription":
   n.TransactionDescription = nil
  case "PaymentReceiptLineList":
   n.PaymentReceiptLineList = nil
  case "VendorInfoRefId":
   n.VendorInfoRefId = nil
  case "RefId":
   n.RefId = nil
  case "ReceivedTransactionId":
   n.ReceivedTransactionId = nil
  case "TransactionType":
   n.TransactionType = nil
  case "FinancialAccountRefIdList":
   n.FinancialAccountRefIdList = nil
  case "TransactionNote":
   n.TransactionNote = nil
  case "DebtorRefId":
   n.DebtorRefId = nil
  case "ChequeNumber":
   n.ChequeNumber = nil
  case "AccountCodeList":
   n.AccountCodeList = nil
  case "InvoiceRefId":
   n.InvoiceRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "TaxRate":
   n.TaxRate = nil
  case "TaxAmount":
   n.TaxAmount = nil
  case "TransactionAmount":
   n.TransactionAmount = nil
  case "TransactionDate":
   n.TransactionDate = nil
  case "ChargedLocationInfoRefId":
   n.ChargedLocationInfoRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PaymentReceipt")
        }
        return n
}

func (n *PaymentReceipt) Set(key string, value interface{}) *PaymentReceipt {
        if n == nil {
                n = PaymentReceiptCreate(PaymentReceipt{})
        }
        switch key {
    case "AccountingPeriod":
    
                      n.AccountingPeriod = ((*LocalIdType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TransactionMethod":
                      n.TransactionMethod = StringCreate(value.(string))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
    case "PaymentReceiptLineList":
                      n.PaymentReceiptLineList = PaymentReceiptLineListTypeCreate(value.(PaymentReceiptLineListType))
    case "VendorInfoRefId":
                      n.VendorInfoRefId = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "ReceivedTransactionId":
                      n.ReceivedTransactionId = StringCreate(value.(string))
    case "TransactionType":
                      n.TransactionType = StringCreate(value.(string))
    case "FinancialAccountRefIdList":
                      n.FinancialAccountRefIdList = FinancialAccountRefIdListTypeCreate(value.(FinancialAccountRefIdListType))
    case "TransactionNote":
                      n.TransactionNote = StringCreate(value.(string))
    case "DebtorRefId":
                      n.DebtorRefId = StringCreate(value.(string))
    case "ChequeNumber":
                      n.ChequeNumber = StringCreate(value.(string))
    case "AccountCodeList":
                      n.AccountCodeList = AccountCodeListTypeCreate(value.(AccountCodeListType))
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TransactionAmount":
                      n.TransactionAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "TransactionDate":
                      n.TransactionDate = StringCreate(value.(string))
    case "ChargedLocationInfoRefId":
                      n.ChargedLocationInfoRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PaymentReceipt")
        }
        return n
}


func (t *StandardsSettingBodyType) CopyString(key string, value interface{}) *StandardsSettingBodyType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StandardsSettingBodyType) Unset(key string) *StandardsSettingBodyType {
        switch key {
  case "StateProvince":
   n.StateProvince = nil
  case "Country":
   n.Country = nil
  case "SettingBodyName":
   n.SettingBodyName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardsSettingBodyType")
        }
        return n
}

func (n *StandardsSettingBodyType) Set(key string, value interface{}) *StandardsSettingBodyType {
        if n == nil {
                n = StandardsSettingBodyTypeCreate(StandardsSettingBodyType{})
        }
        switch key {
    case "StateProvince":
    
                      n.StateProvince = ((*StateProvinceType)(StringCreate(value.(string))))
    case "Country":
      present := false
  for _, b := range AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values {
        if b == value.(string) {
          present = true
        }
    }
    if (!present) {
      log.Fatalf("%s is not present in %s\n", value, "AUCodeSetsStandardAustralianClassificationOfCountriesSACCType_values")
      }

                      n.Country = ((*CountryType)(StringCreate(value.(string))))
    case "SettingBodyName":
                      n.SettingBodyName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardsSettingBodyType")
        }
        return n
}


func (t *SoftwareRequirementType) CopyString(key string, value interface{}) *SoftwareRequirementType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SoftwareRequirementType) Unset(key string) *SoftwareRequirementType {
        switch key {
  case "Vendor":
   n.Vendor = nil
  case "SoftwareTitle":
   n.SoftwareTitle = nil
  case "OS":
   n.OS = nil
  case "Version":
   n.Version = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SoftwareRequirementType")
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
    case "SoftwareTitle":
                      n.SoftwareTitle = StringCreate(value.(string))
    case "OS":
                      n.OS = StringCreate(value.(string))
    case "Version":
                      n.Version = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SoftwareRequirementType")
        }
        return n
}


func (t *LibraryPatronStatus) CopyString(key string, value interface{}) *LibraryPatronStatus {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LibraryPatronStatus) Unset(key string) *LibraryPatronStatus {
        switch key {
  case "PatronLocalId":
   n.PatronLocalId = nil
  case "NumberOfFines":
   n.NumberOfFines = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "PatronRefObject":
   n.PatronRefObject = nil
  case "NumberOfHoldItems":
   n.NumberOfHoldItems = nil
  case "ElectronicIdList":
   n.ElectronicIdList = nil
  case "FineAmount":
   n.FineAmount = nil
  case "NumberOfRefunds":
   n.NumberOfRefunds = nil
  case "PatronRefId":
   n.PatronRefId = nil
  case "LibraryType":
   n.LibraryType = nil
  case "NumberOfCheckouts":
   n.NumberOfCheckouts = nil
  case "MessageList":
   n.MessageList = nil
  case "RefId":
   n.RefId = nil
  case "TransactionList":
   n.TransactionList = nil
  case "PatronName":
   n.PatronName = nil
  case "RefundAmount":
   n.RefundAmount = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "NumberOfOverdues":
   n.NumberOfOverdues = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryPatronStatus")
        }
        return n
}

func (n *LibraryPatronStatus) Set(key string, value interface{}) *LibraryPatronStatus {
        if n == nil {
                n = LibraryPatronStatusCreate(LibraryPatronStatus{})
        }
        switch key {
    case "PatronLocalId":
    
                      n.PatronLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "NumberOfFines":
                      n.NumberOfFines = IntCreate(value.(int))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "PatronRefObject":
                      n.PatronRefObject = StringCreate(value.(string))
    case "NumberOfHoldItems":
                      n.NumberOfHoldItems = IntCreate(value.(int))
    case "ElectronicIdList":
                      n.ElectronicIdList = ElectronicIdListTypeCreate(value.(ElectronicIdListType))
    case "FineAmount":
                      n.FineAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "NumberOfRefunds":
                      n.NumberOfRefunds = IntCreate(value.(int))
    case "PatronRefId":
                      n.PatronRefId = StringCreate(value.(string))
    case "LibraryType":
                      n.LibraryType = StringCreate(value.(string))
    case "NumberOfCheckouts":
                      n.NumberOfCheckouts = IntCreate(value.(int))
    case "MessageList":
                      n.MessageList = LibraryMessageListTypeCreate(value.(LibraryMessageListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "TransactionList":
                      n.TransactionList = LibraryTransactionListTypeCreate(value.(LibraryTransactionListType))
    case "PatronName":
                      n.PatronName = NameOfRecordTypeCreate(value.(NameOfRecordType))
    case "RefundAmount":
                      n.RefundAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "NumberOfOverdues":
                      n.NumberOfOverdues = IntCreate(value.(int))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LibraryPatronStatus")
        }
        return n
}


func (t *PromotionInfoType) CopyString(key string, value interface{}) *PromotionInfoType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PromotionInfoType) Unset(key string) *PromotionInfoType {
        switch key {
  case "PromotionStatus":
   n.PromotionStatus = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PromotionInfoType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PromotionInfoType")
        }
        return n
}


func (t *PaymentReceiptLineType) CopyString(key string, value interface{}) *PaymentReceiptLineType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PaymentReceiptLineType) Unset(key string) *PaymentReceiptLineType {
        switch key {
  case "AccountCode":
   n.AccountCode = nil
  case "LocalPaymentReceiptLineId":
   n.LocalPaymentReceiptLineId = nil
  case "LocalId":
   n.LocalId = nil
  case "TransactionAmount":
   n.TransactionAmount = nil
  case "TaxAmount":
   n.TaxAmount = nil
  case "TaxRate":
   n.TaxRate = nil
  case "TransactionDescription":
   n.TransactionDescription = nil
  case "InvoiceRefId":
   n.InvoiceRefId = nil
  case "FinancialAccountRefId":
   n.FinancialAccountRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PaymentReceiptLineType")
        }
        return n
}

func (n *PaymentReceiptLineType) Set(key string, value interface{}) *PaymentReceiptLineType {
        if n == nil {
                n = PaymentReceiptLineTypeCreate(PaymentReceiptLineType{})
        }
        switch key {
    case "AccountCode":
                      n.AccountCode = StringCreate(value.(string))
    case "LocalPaymentReceiptLineId":
    
                      n.LocalPaymentReceiptLineId = ((*LocalIdType)(StringCreate(value.(string))))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "TransactionAmount":
                      n.TransactionAmount = DebitOrCreditAmountTypeCreate(value.(DebitOrCreditAmountType))
    case "TaxAmount":
                      n.TaxAmount = MonetaryAmountTypeCreate(value.(MonetaryAmountType))
    case "TaxRate":
                      n.TaxRate = FloatCreate(value.(float64))
    case "TransactionDescription":
                      n.TransactionDescription = StringCreate(value.(string))
    case "InvoiceRefId":
                      n.InvoiceRefId = StringCreate(value.(string))
    case "FinancialAccountRefId":
                      n.FinancialAccountRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PaymentReceiptLineType")
        }
        return n
}


func (t *ExclusionRuleType) CopyString(key string, value interface{}) *ExclusionRuleType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ExclusionRuleType) Unset(key string) *ExclusionRuleType {
        switch key {
  case "Type":
   n.Type = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ExclusionRuleType")
        }
        return n
}

func (n *ExclusionRuleType) Set(key string, value interface{}) *ExclusionRuleType {
        if n == nil {
                n = ExclusionRuleTypeCreate(ExclusionRuleType{})
        }
        switch key {
    case "Type":
                      n.Type = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ExclusionRuleType")
        }
        return n
}


func (t *NAPTestItemContentType) CopyString(key string, value interface{}) *NAPTestItemContentType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *NAPTestItemContentType) Unset(key string) *NAPTestItemContentType {
        switch key {
  case "MultipleChoiceOptionCount":
   n.MultipleChoiceOptionCount = nil
  case "ItemDifficultyLogit62":
   n.ItemDifficultyLogit62 = nil
  case "ReleasedStatus":
   n.ReleasedStatus = nil
  case "ItemSubstitutedForList":
   n.ItemSubstitutedForList = nil
  case "ExemplarURL":
   n.ExemplarURL = nil
  case "ItemProficiencyLevel":
   n.ItemProficiencyLevel = nil
  case "ItemName":
   n.ItemName = nil
  case "ItemDifficultyLogit5SE":
   n.ItemDifficultyLogit5SE = nil
  case "Subdomain":
   n.Subdomain = nil
  case "ItemDifficulty":
   n.ItemDifficulty = nil
  case "CorrectAnswer":
   n.CorrectAnswer = nil
  case "MaximumScore":
   n.MaximumScore = nil
  case "ItemType":
   n.ItemType = nil
  case "NAPTestItemLocalId":
   n.NAPTestItemLocalId = nil
  case "ItemDifficultyLogit5":
   n.ItemDifficultyLogit5 = nil
  case "ContentDescriptionList":
   n.ContentDescriptionList = nil
  case "ItemDifficultyLogit62SE":
   n.ItemDifficultyLogit62SE = nil
  case "NAPWritingRubricList":
   n.NAPWritingRubricList = nil
  case "WritingGenre":
   n.WritingGenre = nil
  case "MarkingType":
   n.MarkingType = nil
  case "StimulusList":
   n.StimulusList = nil
  case "ItemDescriptor":
   n.ItemDescriptor = nil
  case "ItemProficiencyBand":
   n.ItemProficiencyBand = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItemContentType")
        }
        return n
}

func (n *NAPTestItemContentType) Set(key string, value interface{}) *NAPTestItemContentType {
        if n == nil {
                n = NAPTestItemContentTypeCreate(NAPTestItemContentType{})
        }
        switch key {
    case "MultipleChoiceOptionCount":
                      n.MultipleChoiceOptionCount = IntCreate(value.(int))
    case "ItemDifficultyLogit62":
                      n.ItemDifficultyLogit62 = FloatCreate(value.(float64))
    case "ReleasedStatus":
                      n.ReleasedStatus = BoolCreate(value.(bool))
    case "ItemSubstitutedForList":
                      n.ItemSubstitutedForList = SubstituteItemListTypeCreate(value.(SubstituteItemListType))
    case "ExemplarURL":
                      n.ExemplarURL = StringCreate(value.(string))
    case "ItemProficiencyLevel":
                      n.ItemProficiencyLevel = StringCreate(value.(string))
    case "ItemName":
                      n.ItemName = StringCreate(value.(string))
    case "ItemDifficultyLogit5SE":
                      n.ItemDifficultyLogit5SE = FloatCreate(value.(float64))
    case "Subdomain":
                      n.Subdomain = StringCreate(value.(string))
    case "ItemDifficulty":
                      n.ItemDifficulty = FloatCreate(value.(float64))
    case "CorrectAnswer":
                      n.CorrectAnswer = StringCreate(value.(string))
    case "MaximumScore":
                      n.MaximumScore = FloatCreate(value.(float64))
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
    case "NAPTestItemLocalId":
    
                      n.NAPTestItemLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "ItemDifficultyLogit5":
                      n.ItemDifficultyLogit5 = FloatCreate(value.(float64))
    case "ContentDescriptionList":
                      n.ContentDescriptionList = ContentDescriptionListTypeCreate(value.(ContentDescriptionListType))
    case "ItemDifficultyLogit62SE":
                      n.ItemDifficultyLogit62SE = FloatCreate(value.(float64))
    case "NAPWritingRubricList":
                      n.NAPWritingRubricList = NAPWritingRubricListTypeCreate(value.(NAPWritingRubricListType))
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
    case "StimulusList":
                      n.StimulusList = StimulusListTypeCreate(value.(StimulusListType))
    case "ItemDescriptor":
                      n.ItemDescriptor = StringCreate(value.(string))
    case "ItemProficiencyBand":
                      n.ItemProficiencyBand = IntCreate(value.(int))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "NAPTestItemContentType")
        }
        return n
}


func (t *MarkValueInfo) CopyString(key string, value interface{}) *MarkValueInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *MarkValueInfo) Unset(key string) *MarkValueInfo {
        switch key {
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "Name":
   n.Name = nil
  case "NumericPrecision":
   n.NumericPrecision = nil
  case "PercentagePassingGrade":
   n.PercentagePassingGrade = nil
  case "NarrativeMaximumSize":
   n.NarrativeMaximumSize = nil
  case "NumericHigh":
   n.NumericHigh = nil
  case "NumericScale":
   n.NumericScale = nil
  case "ValidLetterMarkList":
   n.ValidLetterMarkList = nil
  case "RefId":
   n.RefId = nil
  case "NumericPassingGrade":
   n.NumericPassingGrade = nil
  case "YearLevels":
   n.YearLevels = nil
  case "NumericLow":
   n.NumericLow = nil
  case "PercentageMaximum":
   n.PercentageMaximum = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "PercentageMinimum":
   n.PercentageMinimum = nil
  case "Narrative":
   n.Narrative = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MarkValueInfo")
        }
        return n
}

func (n *MarkValueInfo) Set(key string, value interface{}) *MarkValueInfo {
        if n == nil {
                n = MarkValueInfoCreate(MarkValueInfo{})
        }
        switch key {
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "Name":
                      n.Name = StringCreate(value.(string))
    case "NumericPrecision":
                      n.NumericPrecision = IntCreate(value.(int))
    case "PercentagePassingGrade":
                      n.PercentagePassingGrade = FloatCreate(value.(float64))
    case "NarrativeMaximumSize":
                      n.NarrativeMaximumSize = IntCreate(value.(int))
    case "NumericHigh":
                      n.NumericHigh = FloatCreate(value.(float64))
    case "NumericScale":
                      n.NumericScale = IntCreate(value.(int))
    case "ValidLetterMarkList":
                      n.ValidLetterMarkList = ValidLetterMarkListTypeCreate(value.(ValidLetterMarkListType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "NumericPassingGrade":
                      n.NumericPassingGrade = FloatCreate(value.(float64))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "NumericLow":
                      n.NumericLow = FloatCreate(value.(float64))
    case "PercentageMaximum":
                      n.PercentageMaximum = FloatCreate(value.(float64))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PercentageMinimum":
                      n.PercentageMinimum = FloatCreate(value.(float64))
    case "Narrative":
                      n.Narrative = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "MarkValueInfo")
        }
        return n
}


func (t *Debtor) CopyString(key string, value interface{}) *Debtor {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Debtor) Unset(key string) *Debtor {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "BillingName":
   n.BillingName = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "LocalId":
   n.LocalId = nil
  case "BilledEntity":
   n.BilledEntity = nil
  case "RefId":
   n.RefId = nil
  case "AddressList":
   n.AddressList = nil
  case "BillingNote":
   n.BillingNote = nil
  case "Discount":
   n.Discount = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Debtor")
        }
        return n
}

func (n *Debtor) Set(key string, value interface{}) *Debtor {
        if n == nil {
                n = DebtorCreate(Debtor{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "BillingName":
                      n.BillingName = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "LocalId":
    
                      n.LocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "BilledEntity":
                      n.BilledEntity = Debtor_BilledEntityCreate(value.(Debtor_BilledEntity))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "AddressList":
                      n.AddressList = AddressListTypeCreate(value.(AddressListType))
    case "BillingNote":
                      n.BillingNote = StringCreate(value.(string))
    case "Discount":
                      n.Discount = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Debtor")
        }
        return n
}


func (t *ResourceUsage) CopyString(key string, value interface{}) *ResourceUsage {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceUsage) Unset(key string) *ResourceUsage {
        switch key {
  case "ResourceReportLineList":
   n.ResourceReportLineList = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "ResourceReportColumnList":
   n.ResourceReportColumnList = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "ResourceUsageContentType":
   n.ResourceUsageContentType = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage")
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
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "ResourceReportColumnList":
                      n.ResourceReportColumnList = ResourceUsage_ResourceReportColumnListCreate(value.(ResourceUsage_ResourceReportColumnList))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "ResourceUsageContentType":
                      n.ResourceUsageContentType = ResourceUsage_ResourceUsageContentTypeCreate(value.(ResourceUsage_ResourceUsageContentType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage")
        }
        return n
}


func (t *PlanRequiredContainerType) CopyString(key string, value interface{}) *PlanRequiredContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *PlanRequiredContainerType) Unset(key string) *PlanRequiredContainerType {
        switch key {
  case "Status":
   n.Status = nil
  case "PlanRequiredList":
   n.PlanRequiredList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PlanRequiredContainerType")
        }
        return n
}

func (n *PlanRequiredContainerType) Set(key string, value interface{}) *PlanRequiredContainerType {
        if n == nil {
                n = PlanRequiredContainerTypeCreate(PlanRequiredContainerType{})
        }
        switch key {
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
    case "PlanRequiredList":
                      n.PlanRequiredList = PlanRequiredListTypeCreate(value.(PlanRequiredListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "PlanRequiredContainerType")
        }
        return n
}


func (t *StandardIdentifierType) CopyString(key string, value interface{}) *StandardIdentifierType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StandardIdentifierType) Unset(key string) *StandardIdentifierType {
        switch key {
  case "AlternateIdentificationCodes":
   n.AlternateIdentificationCodes = nil
  case "StandardNumber":
   n.StandardNumber = nil
  case "IndicatorNumber":
   n.IndicatorNumber = nil
  case "Benchmark":
   n.Benchmark = nil
  case "Organization":
   n.Organization = nil
  case "YearCreated":
   n.YearCreated = nil
  case "YearLevels":
   n.YearLevels = nil
  case "ACStrandSubjectArea":
   n.ACStrandSubjectArea = nil
  case "YearLevel":
   n.YearLevel = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardIdentifierType")
        }
        return n
}

func (n *StandardIdentifierType) Set(key string, value interface{}) *StandardIdentifierType {
        if n == nil {
                n = StandardIdentifierTypeCreate(StandardIdentifierType{})
        }
        switch key {
    case "AlternateIdentificationCodes":
                      n.AlternateIdentificationCodes = AlternateIdentificationCodeListTypeCreate(value.(AlternateIdentificationCodeListType))
    case "StandardNumber":
                      n.StandardNumber = StringCreate(value.(string))
    case "IndicatorNumber":
                      n.IndicatorNumber = StringCreate(value.(string))
    case "Benchmark":
                      n.Benchmark = StringCreate(value.(string))
    case "Organization":
                      n.Organization = StringCreate(value.(string))
    case "YearCreated":
                      n.YearCreated = StringCreate(value.(string))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "ACStrandSubjectArea":
                      n.ACStrandSubjectArea = ACStrandSubjectAreaTypeCreate(value.(ACStrandSubjectAreaType))
    case "YearLevel":
                      n.YearLevel = YearLevelTypeCreate(value.(YearLevelType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StandardIdentifierType")
        }
        return n
}


func (t *StudentAttendanceTimeList) CopyString(key string, value interface{}) *StudentAttendanceTimeList {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentAttendanceTimeList) Unset(key string) *StudentAttendanceTimeList {
        switch key {
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "PeriodAttendances":
   n.PeriodAttendances = nil
  case "StudentPersonalRefId":
   n.StudentPersonalRefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "AttendanceTimes":
   n.AttendanceTimes = nil
  case "Date":
   n.Date = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceTimeList")
        }
        return n
}

func (n *StudentAttendanceTimeList) Set(key string, value interface{}) *StudentAttendanceTimeList {
        if n == nil {
                n = StudentAttendanceTimeListCreate(StudentAttendanceTimeList{})
        }
        switch key {
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "PeriodAttendances":
                      n.PeriodAttendances = PeriodAttendancesTypeCreate(value.(PeriodAttendancesType))
    case "StudentPersonalRefId":
                      n.StudentPersonalRefId = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "AttendanceTimes":
                      n.AttendanceTimes = AttendanceTimesTypeCreate(value.(AttendanceTimesType))
    case "Date":
                      n.Date = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentAttendanceTimeList")
        }
        return n
}


func (t *CodeFrameTestItemType) CopyString(key string, value interface{}) *CodeFrameTestItemType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CodeFrameTestItemType) Unset(key string) *CodeFrameTestItemType {
        switch key {
  case "TestItemContent":
   n.TestItemContent = nil
  case "SequenceNumber":
   n.SequenceNumber = nil
  case "TestItemRefId":
   n.TestItemRefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CodeFrameTestItemType")
        }
        return n
}

func (n *CodeFrameTestItemType) Set(key string, value interface{}) *CodeFrameTestItemType {
        if n == nil {
                n = CodeFrameTestItemTypeCreate(CodeFrameTestItemType{})
        }
        switch key {
    case "TestItemContent":
                      n.TestItemContent = NAPTestItemContentTypeCreate(value.(NAPTestItemContentType))
    case "SequenceNumber":
                      n.SequenceNumber = IntCreate(value.(int))
    case "TestItemRefId":
                      n.TestItemRefId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CodeFrameTestItemType")
        }
        return n
}


func (t *ResourceUsage_ResourceReportColumn) CopyString(key string, value interface{}) *ResourceUsage_ResourceReportColumn {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ResourceUsage_ResourceReportColumn) Unset(key string) *ResourceUsage_ResourceReportColumn {
        switch key {
  case "ColumnDelimiter":
   n.ColumnDelimiter = nil
  case "ColumnDescription":
   n.ColumnDescription = nil
  case "ColumnName":
   n.ColumnName = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceReportColumn")
        }
        return n
}

func (n *ResourceUsage_ResourceReportColumn) Set(key string, value interface{}) *ResourceUsage_ResourceReportColumn {
        if n == nil {
                n = ResourceUsage_ResourceReportColumnCreate(ResourceUsage_ResourceReportColumn{})
        }
        switch key {
    case "ColumnDelimiter":
                      n.ColumnDelimiter = StringCreate(value.(string))
    case "ColumnDescription":
                      n.ColumnDescription = StringCreate(value.(string))
    case "ColumnName":
                      n.ColumnName = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ResourceUsage_ResourceReportColumn")
        }
        return n
}


func (t *StimulusType) CopyString(key string, value interface{}) *StimulusType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StimulusType) Unset(key string) *StimulusType {
        switch key {
  case "TextType":
   n.TextType = nil
  case "WordCount":
   n.WordCount = nil
  case "TextDescriptor":
   n.TextDescriptor = nil
  case "StimulusLocalId":
   n.StimulusLocalId = nil
  case "Content":
   n.Content = nil
  case "TextGenre":
   n.TextGenre = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StimulusType")
        }
        return n
}

func (n *StimulusType) Set(key string, value interface{}) *StimulusType {
        if n == nil {
                n = StimulusTypeCreate(StimulusType{})
        }
        switch key {
    case "TextType":
                      n.TextType = StringCreate(value.(string))
    case "WordCount":
                      n.WordCount = IntCreate(value.(int))
    case "TextDescriptor":
                      n.TextDescriptor = StringCreate(value.(string))
    case "StimulusLocalId":
    
                      n.StimulusLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "Content":
                      n.Content = StringCreate(value.(string))
    case "TextGenre":
                      n.TextGenre = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StimulusType")
        }
        return n
}


func (t *AbstractContentPackageType) CopyString(key string, value interface{}) *AbstractContentPackageType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *AbstractContentPackageType) Unset(key string) *AbstractContentPackageType {
        switch key {
  case "BinaryData":
   n.BinaryData = nil
  case "TextData":
   n.TextData = nil
  case "XMLData":
   n.XMLData = nil
  case "Reference":
   n.Reference = nil
  case "RefId":
   n.RefId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType")
        }
        return n
}

func (n *AbstractContentPackageType) Set(key string, value interface{}) *AbstractContentPackageType {
        if n == nil {
                n = AbstractContentPackageTypeCreate(AbstractContentPackageType{})
        }
        switch key {
    case "BinaryData":
                      n.BinaryData = AbstractContentPackageType_BinaryDataCreate(value.(AbstractContentPackageType_BinaryData))
    case "TextData":
                      n.TextData = AbstractContentPackageType_TextDataCreate(value.(AbstractContentPackageType_TextData))
    case "XMLData":
                      n.XMLData = AbstractContentPackageType_XMLDataCreate(value.(AbstractContentPackageType_XMLData))
    case "Reference":
                      n.Reference = AbstractContentPackageType_ReferenceCreate(value.(AbstractContentPackageType_Reference))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "AbstractContentPackageType")
        }
        return n
}


func (t *SystemRole_SystemContext) CopyString(key string, value interface{}) *SystemRole_SystemContext {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *SystemRole_SystemContext) Unset(key string) *SystemRole_SystemContext {
        switch key {
  case "SystemId":
   n.SystemId = nil
  case "RoleList":
   n.RoleList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_SystemContext")
        }
        return n
}

func (n *SystemRole_SystemContext) Set(key string, value interface{}) *SystemRole_SystemContext {
        if n == nil {
                n = SystemRole_SystemContextCreate(SystemRole_SystemContext{})
        }
        switch key {
    case "SystemId":
                      n.SystemId = StringCreate(value.(string))
    case "RoleList":
                      n.RoleList = SystemRole_RoleListCreate(value.(SystemRole_RoleList))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "SystemRole_SystemContext")
        }
        return n
}


func (t *CollectionStatus) CopyString(key string, value interface{}) *CollectionStatus {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CollectionStatus) Unset(key string) *CollectionStatus {
        switch key {
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "ReportingAuthoritySystem":
   n.ReportingAuthoritySystem = nil
  case "RoundCode":
   n.RoundCode = nil
  case "AGReportingObjectResponseList":
   n.AGReportingObjectResponseList = nil
  case "CollectionYear":
   n.CollectionYear = nil
  case "AGCollection":
   n.AGCollection = nil
  case "ReportingAuthority":
   n.ReportingAuthority = nil
  case "SubmissionTimestamp":
   n.SubmissionTimestamp = nil
  case "SubmittedBy":
   n.SubmittedBy = nil
  case "RefId":
   n.RefId = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "ReportingAuthorityCommonwealthId":
   n.ReportingAuthorityCommonwealthId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CollectionStatus")
        }
        return n
}

func (n *CollectionStatus) Set(key string, value interface{}) *CollectionStatus {
        if n == nil {
                n = CollectionStatusCreate(CollectionStatus{})
        }
        switch key {
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "ReportingAuthoritySystem":
                      n.ReportingAuthoritySystem = StringCreate(value.(string))
    case "RoundCode":
                      n.RoundCode = StringCreate(value.(string))
    case "AGReportingObjectResponseList":
                      n.AGReportingObjectResponseList = AGReportingObjectResponseListTypeCreate(value.(AGReportingObjectResponseListType))
    case "CollectionYear":
    
                      n.CollectionYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "AGCollection":
                      n.AGCollection = StringCreate(value.(string))
    case "ReportingAuthority":
                      n.ReportingAuthority = StringCreate(value.(string))
    case "SubmissionTimestamp":
                      n.SubmissionTimestamp = StringCreate(value.(string))
    case "SubmittedBy":
                      n.SubmittedBy = StringCreate(value.(string))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "ReportingAuthorityCommonwealthId":
                      n.ReportingAuthorityCommonwealthId = StringCreate(value.(string))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CollectionStatus")
        }
        return n
}


func (t *ReferralSourceType) CopyString(key string, value interface{}) *ReferralSourceType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ReferralSourceType) Unset(key string) *ReferralSourceType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReferralSourceType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReferralSourceType")
        }
        return n
}


func (t *ProgramAvailabilityType) CopyString(key string, value interface{}) *ProgramAvailabilityType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ProgramAvailabilityType) Unset(key string) *ProgramAvailabilityType {
        switch key {
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "Code":
   n.Code = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramAvailabilityType")
        }
        return n
}

func (n *ProgramAvailabilityType) Set(key string, value interface{}) *ProgramAvailabilityType {
        if n == nil {
                n = ProgramAvailabilityTypeCreate(ProgramAvailabilityType{})
        }
        switch key {
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ProgramAvailabilityType")
        }
        return n
}


func (t *StudentActivityInfo) CopyString(key string, value interface{}) *StudentActivityInfo {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *StudentActivityInfo) Unset(key string) *StudentActivityInfo {
        switch key {
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "StudentActivityType":
   n.StudentActivityType = nil
  case "CurricularStatus":
   n.CurricularStatus = nil
  case "Location":
   n.Location = nil
  case "YearLevels":
   n.YearLevels = nil
  case "RefId":
   n.RefId = nil
  case "Description":
   n.Description = nil
  case "Title":
   n.Title = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "StudentActivityLevel":
   n.StudentActivityLevel = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityInfo")
        }
        return n
}

func (n *StudentActivityInfo) Set(key string, value interface{}) *StudentActivityInfo {
        if n == nil {
                n = StudentActivityInfoCreate(StudentActivityInfo{})
        }
        switch key {
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "StudentActivityType":
                      n.StudentActivityType = StudentActivityTypeCreate(value.(StudentActivityType))
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
    case "Location":
                      n.Location = LocationTypeCreate(value.(LocationType))
    case "YearLevels":
                      n.YearLevels = YearLevelsTypeCreate(value.(YearLevelsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Description":
                      n.Description = StringCreate(value.(string))
    case "Title":
                      n.Title = StringCreate(value.(string))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "StudentActivityLevel":
                      n.StudentActivityLevel = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "StudentActivityInfo")
        }
        return n
}


func (t *Journal_OriginatingTransactionRefId) CopyString(key string, value interface{}) *Journal_OriginatingTransactionRefId {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *Journal_OriginatingTransactionRefId) Unset(key string) *Journal_OriginatingTransactionRefId {
        switch key {
  case "SIF_RefObject":
   n.SIF_RefObject = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Journal_OriginatingTransactionRefId")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "Journal_OriginatingTransactionRefId")
        }
        return n
}


func (t *TechnicalRequirementsType) CopyString(key string, value interface{}) *TechnicalRequirementsType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TechnicalRequirementsType) Unset(key string) *TechnicalRequirementsType {
        switch key {
  case "TechnicalRequirement":
   n.TechnicalRequirement = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TechnicalRequirementsType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TechnicalRequirementsType")
        }
        return n
}


func (t *ApprovalType) CopyString(key string, value interface{}) *ApprovalType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ApprovalType) Unset(key string) *ApprovalType {
        switch key {
  case "Date":
   n.Date = nil
  case "Organization":
   n.Organization = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ApprovalType")
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
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ApprovalType")
        }
        return n
}


func (t *ReligionType) CopyString(key string, value interface{}) *ReligionType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *ReligionType) Unset(key string) *ReligionType {
        switch key {
  case "Code":
   n.Code = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReligionType")
        }
        return n
}

func (n *ReligionType) Set(key string, value interface{}) *ReligionType {
        if n == nil {
                n = ReligionTypeCreate(ReligionType{})
        }
        switch key {
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
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "ReligionType")
        }
        return n
}


func (t *VisaSubClassType) CopyString(key string, value interface{}) *VisaSubClassType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *VisaSubClassType) Unset(key string) *VisaSubClassType {
        switch key {
  case "VisaStatisticalCode":
   n.VisaStatisticalCode = nil
  case "ATEStartDate":
   n.ATEStartDate = nil
  case "Code":
   n.Code = nil
  case "VisaExpiryDate":
   n.VisaExpiryDate = nil
  case "ATEExpiryDate":
   n.ATEExpiryDate = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "VisaSubClassType")
        }
        return n
}

func (n *VisaSubClassType) Set(key string, value interface{}) *VisaSubClassType {
        if n == nil {
                n = VisaSubClassTypeCreate(VisaSubClassType{})
        }
        switch key {
    case "VisaStatisticalCode":
                      n.VisaStatisticalCode = StringCreate(value.(string))
    case "ATEStartDate":
                      n.ATEStartDate = StringCreate(value.(string))
    case "Code":
    
                      n.Code = ((*VisaSubClassCodeType)(StringCreate(value.(string))))
    case "VisaExpiryDate":
                      n.VisaExpiryDate = StringCreate(value.(string))
    case "ATEExpiryDate":
                      n.ATEExpiryDate = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "VisaSubClassType")
        }
        return n
}


func (t *LearningResource_Location) CopyString(key string, value interface{}) *LearningResource_Location {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *LearningResource_Location) Unset(key string) *LearningResource_Location {
        switch key {
  case "ReferenceType":
   n.ReferenceType = nil
  case "Value":
   n.Value = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningResource_Location")
        }
        return n
}

func (n *LearningResource_Location) Set(key string, value interface{}) *LearningResource_Location {
        if n == nil {
                n = LearningResource_LocationCreate(LearningResource_Location{})
        }
        switch key {
    case "ReferenceType":
                      n.ReferenceType = StringCreate(value.(string))
    case "Value":
                      n.Value = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "LearningResource_Location")
        }
        return n
}


func (t *TimeTableSubject) CopyString(key string, value interface{}) *TimeTableSubject {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *TimeTableSubject) Unset(key string) *TimeTableSubject {
        switch key {
  case "ProposedMaxClassSize":
   n.ProposedMaxClassSize = nil
  case "Faculty":
   n.Faculty = nil
  case "AcademicYearRange":
   n.AcademicYearRange = nil
  case "SubjectLongName":
   n.SubjectLongName = nil
  case "SchoolLocalId":
   n.SchoolLocalId = nil
  case "CourseLocalId":
   n.CourseLocalId = nil
  case "SchoolYear":
   n.SchoolYear = nil
  case "SIF_Metadata":
   n.SIF_Metadata = nil
  case "SchoolInfoRefId":
   n.SchoolInfoRefId = nil
  case "OtherCodeList":
   n.OtherCodeList = nil
  case "SchoolCourseInfoRefId":
   n.SchoolCourseInfoRefId = nil
  case "LocalCodeList":
   n.LocalCodeList = nil
  case "SubjectLocalId":
   n.SubjectLocalId = nil
  case "SubjectType":
   n.SubjectType = nil
  case "SubjectShortName":
   n.SubjectShortName = nil
  case "SIF_ExtendedElements":
   n.SIF_ExtendedElements = nil
  case "RefId":
   n.RefId = nil
  case "Semester":
   n.Semester = nil
  case "AcademicYear":
   n.AcademicYear = nil
  case "ProposedMinClassSize":
   n.ProposedMinClassSize = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableSubject")
        }
        return n
}

func (n *TimeTableSubject) Set(key string, value interface{}) *TimeTableSubject {
        if n == nil {
                n = TimeTableSubjectCreate(TimeTableSubject{})
        }
        switch key {
    case "ProposedMaxClassSize":
                      n.ProposedMaxClassSize = FloatCreate(value.(float64))
    case "Faculty":
                      n.Faculty = StringCreate(value.(string))
    case "AcademicYearRange":
                      n.AcademicYearRange = YearRangeTypeCreate(value.(YearRangeType))
    case "SubjectLongName":
                      n.SubjectLongName = StringCreate(value.(string))
    case "SchoolLocalId":
    
                      n.SchoolLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "CourseLocalId":
    
                      n.CourseLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SchoolYear":
    
                      n.SchoolYear = ((*SchoolYearType)(StringCreate(value.(string))))
    case "SIF_Metadata":
                      n.SIF_Metadata = SIF_MetadataTypeCreate(value.(SIF_MetadataType))
    case "SchoolInfoRefId":
                      n.SchoolInfoRefId = StringCreate(value.(string))
    case "OtherCodeList":
                      n.OtherCodeList = OtherCodeListTypeCreate(value.(OtherCodeListType))
    case "SchoolCourseInfoRefId":
    
                      n.SchoolCourseInfoRefId = ((*RefIdType)(StringCreate(value.(string))))
    case "LocalCodeList":
                      n.LocalCodeList = LocalCodeListTypeCreate(value.(LocalCodeListType))
    case "SubjectLocalId":
    
                      n.SubjectLocalId = ((*LocalIdType)(StringCreate(value.(string))))
    case "SubjectType":
                      n.SubjectType = StringCreate(value.(string))
    case "SubjectShortName":
                      n.SubjectShortName = StringCreate(value.(string))
    case "SIF_ExtendedElements":
                      n.SIF_ExtendedElements = SIF_ExtendedElementsTypeCreate(value.(SIF_ExtendedElementsType))
    case "RefId":
    
                      n.RefId = ((*RefIdType)(StringCreate(value.(string))))
    case "Semester":
                      n.Semester = IntCreate(value.(int))
    case "AcademicYear":
                      n.AcademicYear = YearLevelTypeCreate(value.(YearLevelType))
    case "ProposedMinClassSize":
                      n.ProposedMinClassSize = FloatCreate(value.(float64))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "TimeTableSubject")
        }
        return n
}


func (t *CampusContainerType) CopyString(key string, value interface{}) *CampusContainerType {
  return t.Set(key, fmt.Sprint(reflect.ValueOf(value).Elem().Interface()))
}

func (n *CampusContainerType) Unset(key string) *CampusContainerType {
        switch key {
  case "SchoolCampusId":
   n.SchoolCampusId = nil
  case "AdminStatus":
   n.AdminStatus = nil
  case "CampusType":
   n.CampusType = nil
  case "ParentSchoolId":
   n.ParentSchoolId = nil
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CampusContainerType")
        }
        return n
}

func (n *CampusContainerType) Set(key string, value interface{}) *CampusContainerType {
        if n == nil {
                n = CampusContainerTypeCreate(CampusContainerType{})
        }
        switch key {
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
    case "ParentSchoolId":
                      n.ParentSchoolId = StringCreate(value.(string))
        default:
          log.Fatalf("%s is not a valid element name in %s\n", key, "CampusContainerType")
        }
        return n
}

