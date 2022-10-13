package sifxml


type NAPTestScoreSummarys []NAPTestScoreSummary

    type NAPTestScoreSummary struct {
  naptestscoresummary `xml:"NAPTestScoreSummary" json:"NAPTestScoreSummary"`
}

type naptestscoresummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId,omitempty"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainNationalAverage *Float `xml:"DomainNationalAverage,omitempty" json:"DomainNationalAverage,omitempty"`
      DomainSchoolAverage *Float `xml:"DomainSchoolAverage,omitempty" json:"DomainSchoolAverage,omitempty"`
      DomainJurisdictionAverage *Float `xml:"DomainJurisdictionAverage,omitempty" json:"DomainJurisdictionAverage,omitempty"`
      DomainTopNational60Percent *Float `xml:"DomainTopNational60Percent,omitempty" json:"DomainTopNational60Percent,omitempty"`
      DomainBottomNational60Percent *Float `xml:"DomainBottomNational60Percent,omitempty" json:"DomainBottomNational60Percent,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

