package sifxml


    type NAPStudentResponseSet struct {
        RefId RefIdType `xml:"RefId,attr"`
      ReportExclusionFlag string `xml:"ReportExclusionFlag"`
      CalibrationSampleFlag string `xml:"CalibrationSampleFlag"`
      EquatingSampleFlag string `xml:"EquatingSampleFlag"`
      PathTakenForDomain string `xml:"PathTakenForDomain"`
      ParallelTest string `xml:"ParallelTest"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      PlatformStudentIdentifier LocalIdType `xml:"PlatformStudentIdentifier"`
      NAPTestRefId string `xml:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId"`
      DomainScore DomainScoreType `xml:"DomainScore"`
      TestletList NAPStudentResponseTestletListType `xml:"TestletList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    