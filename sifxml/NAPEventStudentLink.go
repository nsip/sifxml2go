package sifxml


type NAPEventStudentLinks []NAPEventStudentLink

    type NAPEventStudentLink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier *LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      SchoolSector *AUCodeSetsSchoolSectorCodeType `xml:"SchoolSector" json:"SchoolSector"`
      System *AUCodeSetsSchoolSystemType `xml:"System,omitempty" json:"System"`
      SchoolGeolocation *AUCodeSetsSchoolLocationType `xml:"SchoolGeolocation,omitempty" json:"SchoolGeolocation"`
      ReportingSchoolName *String `xml:"ReportingSchoolName,omitempty" json:"ReportingSchoolName"`
      NAPJurisdiction *AUCodeSetsNAPJurisdictionType `xml:"NAPJurisdiction,omitempty" json:"NAPJurisdiction"`
      ParticipationCode *AUCodeSetsNAPParticipationCodeType `xml:"ParticipationCode" json:"ParticipationCode"`
      ParticipationText *String `xml:"ParticipationText" json:"ParticipationText"`
      Device *String `xml:"Device,omitempty" json:"Device"`
      Date *String `xml:"Date,omitempty" json:"Date"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      LapsedTimeTest *String `xml:"LapsedTimeTest,omitempty" json:"LapsedTimeTest"`
      ExemptionReason *String `xml:"ExemptionReason,omitempty" json:"ExemptionReason"`
      PersonalDetailsChanged *Bool `xml:"PersonalDetailsChanged,omitempty" json:"PersonalDetailsChanged"`
      PSIOtherIdMatch *Bool `xml:"PSIOtherIdMatch,omitempty" json:"PSIOtherIdMatch"`
      PossibleDuplicate *Bool `xml:"PossibleDuplicate,omitempty" json:"PossibleDuplicate"`
      DOBRange *Bool `xml:"DOBRange,omitempty" json:"DOBRange"`
      TestDisruptionList *TestDisruptionListType `xml:"TestDisruptionList,omitempty" json:"TestDisruptionList"`
      Adjustment *AdjustmentContainerType `xml:"Adjustment,omitempty" json:"Adjustment"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    