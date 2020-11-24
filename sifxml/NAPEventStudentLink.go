package sifxml


type NAPEventStudentLinks []NAPEventStudentLink

    type NAPEventStudentLink struct {
  napeventstudentlink `xml:"NAPEventStudentLink" json:"NAPEventStudentLink"`
}

type napeventstudentlink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      PlatformStudentIdentifier *LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId,omitempty"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      SchoolSector *AUCodeSetsSchoolSectorCodeType `xml:"SchoolSector" json:"SchoolSector"`
      System *AUCodeSetsSchoolSystemType `xml:"System,omitempty" json:"System,omitempty"`
      SchoolGeolocation *AUCodeSetsSchoolLocationType `xml:"SchoolGeolocation,omitempty" json:"SchoolGeolocation,omitempty"`
      ReportingSchoolName *String `xml:"ReportingSchoolName,omitempty" json:"ReportingSchoolName,omitempty"`
      NAPJurisdiction *AUCodeSetsNAPJurisdictionType `xml:"NAPJurisdiction,omitempty" json:"NAPJurisdiction,omitempty"`
      ParticipationCode *AUCodeSetsNAPParticipationCodeType `xml:"ParticipationCode" json:"ParticipationCode"`
      ParticipationText *String `xml:"ParticipationText" json:"ParticipationText"`
      Device *String `xml:"Device,omitempty" json:"Device,omitempty"`
      Date *String `xml:"Date,omitempty" json:"Date,omitempty"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime,omitempty"`
      LapsedTimeTest *String `xml:"LapsedTimeTest,omitempty" json:"LapsedTimeTest,omitempty"`
      ExemptionReason *String `xml:"ExemptionReason,omitempty" json:"ExemptionReason,omitempty"`
      PersonalDetailsChanged *Bool `xml:"PersonalDetailsChanged,omitempty" json:"PersonalDetailsChanged,omitempty"`
      PSIOtherIdMatch *Bool `xml:"PSIOtherIdMatch,omitempty" json:"PSIOtherIdMatch,omitempty"`
      PossibleDuplicate *Bool `xml:"PossibleDuplicate,omitempty" json:"PossibleDuplicate,omitempty"`
      DOBRange *Bool `xml:"DOBRange,omitempty" json:"DOBRange,omitempty"`
      TestDisruptionList *TestDisruptionListType `xml:"TestDisruptionList,omitempty" json:"TestDisruptionList,omitempty"`
      Adjustment *AdjustmentContainerType `xml:"Adjustment,omitempty" json:"Adjustment,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    