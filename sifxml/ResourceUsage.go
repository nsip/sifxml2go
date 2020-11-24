package sifxml


type ResourceUsages []ResourceUsage

    type ResourceUsage struct {
  resourceusage `xml:"ResourceUsage" json:"ResourceUsage"`
}

type resourceusage struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      ResourceUsageContentType *ResourceUsage_ResourceUsageContentType `xml:"ResourceUsageContentType" json:"ResourceUsageContentType"`
      ResourceReportColumnList *ResourceUsage_ResourceReportColumnList `xml:"ResourceReportColumnList" json:"ResourceReportColumnList"`
      ResourceReportLineList *ResourceUsage_ResourceReportLineList `xml:"ResourceReportLineList" json:"ResourceReportLineList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type ResourceUsage_ResourceUsageContentType struct {
  resourceusage_resourceusagecontenttype `xml:"ResourceUsage_ResourceUsageContentType" json:"ResourceUsage_ResourceUsageContentType"`
}

type resourceusage_resourceusagecontenttype struct {
       Code *AUCodeSetsResourceUsageContentTypeType `xml:"Code" json:"Code"`
       LocalDescription *String `xml:"LocalDescription,omitempty" json:"LocalDescription,omitempty"`
}
type ResourceUsage_ResourceReportColumnList struct {
  resourceusage_resourcereportcolumnlist `xml:"ResourceUsage_ResourceReportColumnList" json:"ResourceUsage_ResourceReportColumnList"`
}

type resourceusage_resourcereportcolumnlist struct {
      ResourceReportColumn []ResourceUsage_ResourceReportColumn `xml:"ResourceReportColumn" json:"ResourceReportColumn"`
}
type ResourceUsage_ResourceReportLineList struct {
  resourceusage_resourcereportlinelist `xml:"ResourceUsage_ResourceReportLineList" json:"ResourceUsage_ResourceReportLineList"`
}

type resourceusage_resourcereportlinelist struct {
      ResourceReportLine []ResourceUsage_ResourceReportLine `xml:"ResourceReportLine" json:"ResourceReportLine"`
}
type ResourceUsage_ResourceReportColumn struct {
  resourceusage_resourcereportcolumn `xml:"ResourceUsage_ResourceReportColumn" json:"ResourceUsage_ResourceReportColumn"`
}

type resourceusage_resourcereportcolumn struct {
       ColumnName *String `xml:"ColumnName" json:"ColumnName"`
       ColumnDescription *String `xml:"ColumnDescription,omitempty" json:"ColumnDescription,omitempty"`
       ColumnDelimiter *String `xml:"ColumnDelimiter,omitempty" json:"ColumnDelimiter,omitempty"`
}
type ResourceUsage_ResourceReportLine struct {
  resourceusage_resourcereportline `xml:"ResourceUsage_ResourceReportLine" json:"ResourceUsage_ResourceReportLine"`
}

type resourceusage_resourcereportline struct {
      SIF_RefId *ResourceUsage_SIF_RefId
       StartDate *String `xml:"StartDate" json:"StartDate"`
       EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
       CurrentCost *MonetaryAmountType `xml:"CurrentCost" json:"CurrentCost"`
       ReportRow *String `xml:"ReportRow" json:"ReportRow"`
}
type ResourceUsage_SIF_RefId struct {
  resourceusage_sif_refid `xml:"ResourceUsage_SIF_RefId" json:"ResourceUsage_SIF_RefId"`
}

type resourceusage_sif_refid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
