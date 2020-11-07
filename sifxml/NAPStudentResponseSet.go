package sifxml

import _ "github.com/creasty/defaults"


type NAPStudentResponseSets []NAPStudentResponseSet

    type NAPStudentResponseSet struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportExclusionFlag bool `xml:"ReportExclusionFlag" json:"ReportExclusionFlag"`
      CalibrationSampleFlag string `xml:"CalibrationSampleFlag,omitempty" json:"CalibrationSampleFlag"`
      EquatingSampleFlag string `xml:"EquatingSampleFlag,omitempty" json:"EquatingSampleFlag"`
      PathTakenForDomain string `xml:"PathTakenForDomain" json:"PathTakenForDomain"`
      ParallelTest string `xml:"ParallelTest" json:"ParallelTest"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      NAPTestRefId string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainScore DomainScoreType `xml:"DomainScore" json:"DomainScore"`
      TestletList NAPStudentResponseTestletListType `xml:"TestletList" json:"TestletList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    