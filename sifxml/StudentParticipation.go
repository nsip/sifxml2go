package sifxml


type StudentParticipations []StudentParticipation

    type StudentParticipation struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentParticipationAsOfDate *String `xml:"StudentParticipationAsOfDate" json:"StudentParticipationAsOfDate"`
      ProgramType *AUCodeSetsStudentFamilyProgramTypeType `xml:"ProgramType,omitempty" json:"ProgramType"`
      ProgramFundingSources *ProgramFundingSourcesType `xml:"ProgramFundingSources,omitempty" json:"ProgramFundingSources"`
      ManagingSchool *StudentParticipation_ManagingSchool `xml:"ManagingSchool" json:"ManagingSchool"`
      ReferralDate *String `xml:"ReferralDate,omitempty" json:"ReferralDate"`
      ReferralSource *ReferralSourceType `xml:"ReferralSource,omitempty" json:"ReferralSource"`
      ProgramStatus *ProgramStatusType `xml:"ProgramStatus,omitempty" json:"ProgramStatus"`
      GiftedEligibilityCriteria *AUCodeSetsYesOrNoCategoryType `xml:"GiftedEligibilityCriteria,omitempty" json:"GiftedEligibilityCriteria"`
      EvaluationParentalConsentDate *String `xml:"EvaluationParentalConsentDate,omitempty" json:"EvaluationParentalConsentDate"`
      EvaluationDate *String `xml:"EvaluationDate,omitempty" json:"EvaluationDate"`
      EvaluationExtensionDate *String `xml:"EvaluationExtensionDate,omitempty" json:"EvaluationExtensionDate"`
      ExtensionComments *String `xml:"ExtensionComments,omitempty" json:"ExtensionComments"`
      ReevaluationDate *String `xml:"ReevaluationDate,omitempty" json:"ReevaluationDate"`
      ProgramEligibilityDate *String `xml:"ProgramEligibilityDate,omitempty" json:"ProgramEligibilityDate"`
      ProgramPlanDate *String `xml:"ProgramPlanDate,omitempty" json:"ProgramPlanDate"`
      ProgramPlanEffectiveDate *String `xml:"ProgramPlanEffectiveDate,omitempty" json:"ProgramPlanEffectiveDate"`
      NOREPDate *String `xml:"NOREPDate,omitempty" json:"NOREPDate"`
      PlacementParentalConsentDate *String `xml:"PlacementParentalConsentDate,omitempty" json:"PlacementParentalConsentDate"`
      ProgramPlacementDate *String `xml:"ProgramPlacementDate,omitempty" json:"ProgramPlacementDate"`
      ExtendedSchoolYear *Bool `xml:"ExtendedSchoolYear,omitempty" json:"ExtendedSchoolYear"`
      ExtendedDay *Bool `xml:"ExtendedDay,omitempty" json:"ExtendedDay"`
      ProgramAvailability *ProgramAvailabilityType `xml:"ProgramAvailability,omitempty" json:"ProgramAvailability"`
      EntryPerson *String `xml:"EntryPerson,omitempty" json:"EntryPerson"`
      StudentSpecialEducationFTE *Float `xml:"StudentSpecialEducationFTE,omitempty" json:"StudentSpecialEducationFTE"`
      ParticipationContact *String `xml:"ParticipationContact,omitempty" json:"ParticipationContact"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type StudentParticipation_ManagingSchool struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
