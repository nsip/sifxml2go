package sifxml


type NAPStudentResponseSets []NAPStudentResponseSet

    type NAPStudentResponseSet struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportExclusionFlag *Bool `xml:"ReportExclusionFlag" json:"ReportExclusionFlag"`
      CalibrationSampleFlag *String `xml:"CalibrationSampleFlag,omitempty" json:"CalibrationSampleFlag"`
      EquatingSampleFlag *String `xml:"EquatingSampleFlag,omitempty" json:"EquatingSampleFlag"`
      PathTakenForDomain *String `xml:"PathTakenForDomain,omitempty" json:"PathTakenForDomain"`
      ParallelTest *String `xml:"ParallelTest,omitempty" json:"ParallelTest"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier *LocalIdType `xml:"PlatformStudentIdentifier" json:"PlatformStudentIdentifier"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainScore *DomainScoreType `xml:"DomainScore,omitempty" json:"DomainScore"`
      TestletList *NAPStudentResponseTestletListType `xml:"TestletList,omitempty" json:"TestletList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    