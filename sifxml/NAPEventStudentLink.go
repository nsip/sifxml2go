package sifxml

import _ "github.com/creasty/defaults"


type NAPEventStudentLinks []NAPEventStudentLink

    type NAPEventStudentLink struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      SchoolSector AUCodeSetsSchoolSectorCodeType `xml:"SchoolSector" json:"SchoolSector"`
      System AUCodeSetsSchoolSystemType `xml:"System,omitempty" json:"System"`
      SchoolGeolocation AUCodeSetsSchoolLocationType `xml:"SchoolGeolocation,omitempty" json:"SchoolGeolocation"`
      ReportingSchoolName string `xml:"ReportingSchoolName,omitempty" json:"ReportingSchoolName"`
      NAPJurisdiction AUCodeSetsNAPJurisdictionType `xml:"NAPJurisdiction,omitempty" json:"NAPJurisdiction"`
      ParticipationCode AUCodeSetsNAPParticipationCodeType `xml:"ParticipationCode" json:"ParticipationCode"`
      ParticipationText string `xml:"ParticipationText" json:"ParticipationText"`
      Device string `xml:"Device,omitempty" json:"Device"`
      Date string `xml:"Date" json:"Date"`
      StartTime string `xml:"StartTime" json:"StartTime"`
      LapsedTimeTest string `xml:"LapsedTimeTest" json:"LapsedTimeTest"`
      ExemptionReason string `xml:"ExemptionReason,omitempty" json:"ExemptionReason"`
      PersonalDetailsChanged bool `xml:"PersonalDetailsChanged" json:"PersonalDetailsChanged"`
      PSIOtherIdMatch bool `xml:"PSIOtherIdMatch" json:"PSIOtherIdMatch"`
      PossibleDuplicate bool `xml:"PossibleDuplicate" json:"PossibleDuplicate"`
      DOBRange bool `xml:"DOBRange" json:"DOBRange"`
      TestDisruptionList TestDisruptionListType `xml:"TestDisruptionList,omitempty" json:"TestDisruptionList"`
      Adjustment AdjustmentContainerType `xml:"Adjustment,omitempty" json:"Adjustment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    