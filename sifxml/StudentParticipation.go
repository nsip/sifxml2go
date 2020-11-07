package sifxml


type StudentParticipations []StudentParticipation

    type StudentParticipation struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentParticipationAsOfDate *string `xml:"StudentParticipationAsOfDate" json:"StudentParticipationAsOfDate"`
      ProgramType *AUCodeSetsStudentFamilyProgramTypeType `xml:"ProgramType,omitempty" json:"ProgramType"`
      ProgramFundingSources *ProgramFundingSourcesType `xml:"ProgramFundingSources,omitempty" json:"ProgramFundingSources"`
      ManagingSchool *StudentParticipation_ManagingSchool `xml:"ManagingSchool" json:"ManagingSchool"`
      ReferralDate *string `xml:"ReferralDate,omitempty" json:"ReferralDate"`
      ReferralSource *ReferralSourceType `xml:"ReferralSource,omitempty" json:"ReferralSource"`
      ProgramStatus *ProgramStatusType `xml:"ProgramStatus,omitempty" json:"ProgramStatus"`
      GiftedEligibilityCriteria *AUCodeSetsYesOrNoCategoryType `xml:"GiftedEligibilityCriteria,omitempty" json:"GiftedEligibilityCriteria"`
      EvaluationParentalConsentDate *string `xml:"EvaluationParentalConsentDate,omitempty" json:"EvaluationParentalConsentDate"`
      EvaluationDate *string `xml:"EvaluationDate,omitempty" json:"EvaluationDate"`
      EvaluationExtensionDate *string `xml:"EvaluationExtensionDate,omitempty" json:"EvaluationExtensionDate"`
      ExtensionComments *string `xml:"ExtensionComments,omitempty" json:"ExtensionComments"`
      ReevaluationDate *string `xml:"ReevaluationDate,omitempty" json:"ReevaluationDate"`
      ProgramEligibilityDate *string `xml:"ProgramEligibilityDate,omitempty" json:"ProgramEligibilityDate"`
      ProgramPlanDate *string `xml:"ProgramPlanDate,omitempty" json:"ProgramPlanDate"`
      ProgramPlanEffectiveDate *string `xml:"ProgramPlanEffectiveDate,omitempty" json:"ProgramPlanEffectiveDate"`
      NOREPDate *string `xml:"NOREPDate,omitempty" json:"NOREPDate"`
      PlacementParentalConsentDate *string `xml:"PlacementParentalConsentDate,omitempty" json:"PlacementParentalConsentDate"`
      ProgramPlacementDate *string `xml:"ProgramPlacementDate,omitempty" json:"ProgramPlacementDate"`
      ExtendedSchoolYear *bool `xml:"ExtendedSchoolYear" json:"ExtendedSchoolYear"`
      ExtendedDay *bool `xml:"ExtendedDay" json:"ExtendedDay"`
      ProgramAvailability *ProgramAvailabilityType `xml:"ProgramAvailability,omitempty" json:"ProgramAvailability"`
      EntryPerson *string `xml:"EntryPerson,omitempty" json:"EntryPerson"`
      StudentSpecialEducationFTE *float64 `xml:"StudentSpecialEducationFTE" json:"StudentSpecialEducationFTE"`
      ParticipationContact *string `xml:"ParticipationContact,omitempty" json:"ParticipationContact"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type StudentParticipation_ManagingSchool struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
