package sifxml


type NAPStudentResponseSets []NAPStudentResponseSet

    type NAPStudentResponseSet struct {
  napstudentresponseset `xml:"NAPStudentResponseSet" json:"NAPStudentResponseSet"`
}

type napstudentresponseset struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportExclusionFlag *Bool `xml:"ReportExclusionFlag" json:"ReportExclusionFlag"`
      CalibrationSampleFlag *String `xml:"CalibrationSampleFlag,omitempty" json:"CalibrationSampleFlag,omitempty"`
      EquatingSampleFlag *String `xml:"EquatingSampleFlag,omitempty" json:"EquatingSampleFlag,omitempty"`
      PathTakenForDomain *String `xml:"PathTakenForDomain,omitempty" json:"PathTakenForDomain,omitempty"`
      ParallelTest *String `xml:"ParallelTest,omitempty" json:"ParallelTest,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId,omitempty"`
      PlatformStudentIdentifier *LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId,omitempty"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainScore *DomainScoreType `xml:"DomainScore,omitempty" json:"DomainScore,omitempty"`
      TestletList *NAPStudentResponseTestletListType `xml:"TestletList,omitempty" json:"TestletList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    