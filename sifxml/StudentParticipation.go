package sifxml


type StudentParticipations []StudentParticipation

    type StudentParticipation struct {
  studentparticipation `xml:"StudentParticipation" json:"StudentParticipation"`
}

type studentparticipation struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentParticipationAsOfDate *String `xml:"StudentParticipationAsOfDate" json:"StudentParticipationAsOfDate"`
      ProgramType *AUCodeSetsStudentFamilyProgramTypeType `xml:"ProgramType,omitempty" json:"ProgramType,omitempty"`
      ProgramFundingSources *ProgramFundingSourcesType `xml:"ProgramFundingSources,omitempty" json:"ProgramFundingSources,omitempty"`
      ManagingSchool *StudentParticipation_ManagingSchool `xml:"ManagingSchool" json:"ManagingSchool"`
      ReferralDate *String `xml:"ReferralDate,omitempty" json:"ReferralDate,omitempty"`
      ReferralSource *ReferralSourceType `xml:"ReferralSource,omitempty" json:"ReferralSource,omitempty"`
      ProgramStatus *ProgramStatusType `xml:"ProgramStatus,omitempty" json:"ProgramStatus,omitempty"`
      GiftedEligibilityCriteria *AUCodeSetsYesOrNoCategoryType `xml:"GiftedEligibilityCriteria,omitempty" json:"GiftedEligibilityCriteria,omitempty"`
      EvaluationParentalConsentDate *String `xml:"EvaluationParentalConsentDate,omitempty" json:"EvaluationParentalConsentDate,omitempty"`
      EvaluationDate *String `xml:"EvaluationDate,omitempty" json:"EvaluationDate,omitempty"`
      EvaluationExtensionDate *String `xml:"EvaluationExtensionDate,omitempty" json:"EvaluationExtensionDate,omitempty"`
      ExtensionComments *String `xml:"ExtensionComments,omitempty" json:"ExtensionComments,omitempty"`
      ReevaluationDate *String `xml:"ReevaluationDate,omitempty" json:"ReevaluationDate,omitempty"`
      ProgramEligibilityDate *String `xml:"ProgramEligibilityDate,omitempty" json:"ProgramEligibilityDate,omitempty"`
      ProgramPlanDate *String `xml:"ProgramPlanDate,omitempty" json:"ProgramPlanDate,omitempty"`
      ProgramPlanEffectiveDate *String `xml:"ProgramPlanEffectiveDate,omitempty" json:"ProgramPlanEffectiveDate,omitempty"`
      NOREPDate *String `xml:"NOREPDate,omitempty" json:"NOREPDate,omitempty"`
      PlacementParentalConsentDate *String `xml:"PlacementParentalConsentDate,omitempty" json:"PlacementParentalConsentDate,omitempty"`
      ProgramPlacementDate *String `xml:"ProgramPlacementDate,omitempty" json:"ProgramPlacementDate,omitempty"`
      ExtendedSchoolYear *Bool `xml:"ExtendedSchoolYear,omitempty" json:"ExtendedSchoolYear,omitempty"`
      ExtendedDay *Bool `xml:"ExtendedDay,omitempty" json:"ExtendedDay,omitempty"`
      ProgramAvailability *ProgramAvailabilityType `xml:"ProgramAvailability,omitempty" json:"ProgramAvailability,omitempty"`
      EntryPerson *String `xml:"EntryPerson,omitempty" json:"EntryPerson,omitempty"`
      StudentSpecialEducationFTE *FTEType `xml:"StudentSpecialEducationFTE,omitempty" json:"StudentSpecialEducationFTE,omitempty"`
      ParticipationContact *String `xml:"ParticipationContact,omitempty" json:"ParticipationContact,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type StudentParticipation_ManagingSchool struct {
  studentparticipation_managingschool `xml:"StudentParticipation_ManagingSchool" json:"StudentParticipation_ManagingSchool"`
}

type studentparticipation_managingschool struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

