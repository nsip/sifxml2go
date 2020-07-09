package sifxml


type NAPEventStudentLinks []NAPEventStudentLink

    type NAPEventStudentLink struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier *LocalIdType `xml:"PlatformStudentIdentifier,omitempty" json:"PlatformStudentIdentifier"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      NAPTestRefId *string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId,omitempty" json:"NAPTestLocalId"`
      SchoolSector *string `xml:"SchoolSector,omitempty" json:"SchoolSector"`
      System *string `xml:"System,omitempty" json:"System"`
      SchoolGeolocation *string `xml:"SchoolGeolocation,omitempty" json:"SchoolGeolocation"`
      ReportingSchoolName *string `xml:"ReportingSchoolName,omitempty" json:"ReportingSchoolName"`
      NAPJurisdiction *string `xml:"NAPJurisdiction,omitempty" json:"NAPJurisdiction"`
      ParticipationCode *string `xml:"ParticipationCode,omitempty" json:"ParticipationCode"`
      ParticipationText *string `xml:"ParticipationText,omitempty" json:"ParticipationText"`
      Device *string `xml:"Device,omitempty" json:"Device"`
      Date *string `xml:"Date,omitempty" json:"Date"`
      StartTime *string `xml:"StartTime,omitempty" json:"StartTime"`
      LapsedTimeTest *string `xml:"LapsedTimeTest,omitempty" json:"LapsedTimeTest"`
      ExemptionReason *string `xml:"ExemptionReason,omitempty" json:"ExemptionReason"`
      PersonalDetailsChanged *bool `xml:"PersonalDetailsChanged,omitempty" json:"PersonalDetailsChanged"`
      PSIOtherIdMatch *bool `xml:"PSIOtherIdMatch,omitempty" json:"PSIOtherIdMatch"`
      PossibleDuplicate *bool `xml:"PossibleDuplicate,omitempty" json:"PossibleDuplicate"`
      DOBRange *bool `xml:"DOBRange,omitempty" json:"DOBRange"`
      TestDisruptionList *TestDisruptionListType `xml:"TestDisruptionList,omitempty" json:"TestDisruptionList"`
      Adjustment *AdjustmentContainerType `xml:"Adjustment,omitempty" json:"Adjustment"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    