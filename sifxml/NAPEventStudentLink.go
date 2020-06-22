package sifxml


    type NAPEventStudentLink struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      PlatformStudentIdentifier LocalIdType `xml:"PlatformStudentIdentifier"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId"`
      NAPTestRefId string `xml:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId"`
      SchoolSector string `xml:"SchoolSector"`
      System string `xml:"System"`
      SchoolGeolocation string `xml:"SchoolGeolocation"`
      ReportingSchoolName string `xml:"ReportingSchoolName"`
      NAPJurisdiction string `xml:"NAPJurisdiction"`
      ParticipationCode string `xml:"ParticipationCode"`
      ParticipationText string `xml:"ParticipationText"`
      Device string `xml:"Device"`
      Date string `xml:"Date"`
      StartTime string `xml:"StartTime"`
      LapsedTimeTest string `xml:"LapsedTimeTest"`
      ExemptionReason string `xml:"ExemptionReason"`
      PersonalDetailsChanged string `xml:"PersonalDetailsChanged"`
      PSIOtherIdMatch string `xml:"PSIOtherIdMatch"`
      PossibleDuplicate string `xml:"PossibleDuplicate"`
      DOBRange string `xml:"DOBRange"`
      TestDisruptionList TestDisruptionListType `xml:"TestDisruptionList"`
      Adjustment AdjustmentContainerType `xml:"Adjustment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    