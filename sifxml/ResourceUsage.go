package sifxml


type ResourceUsages []ResourceUsage

    type ResourceUsage struct {
  resourceusage `xml:"ResourceUsage" json:"ResourceUsage"`
}

type resourceusage struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      ResourceUsageContentType *ResourceUsageContentTypeType `xml:"ResourceUsageContentType" json:"ResourceUsageContentType"`
      ResourceReportColumnList *ResourceReportColumnListType `xml:"ResourceReportColumnList" json:"ResourceReportColumnList"`
      ResourceReportLineList *ResourceReportLineListType `xml:"ResourceReportLineList" json:"ResourceReportLineList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    