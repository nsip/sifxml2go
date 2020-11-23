package sifxml


type ResourceUsages []ResourceUsage

    type ResourceUsage struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      ResourceUsageContentType *ResourceUsage_ResourceUsageContentType `xml:"ResourceUsageContentType" json:"ResourceUsageContentType"`
      ResourceReportColumnList *ResourceUsage_ResourceReportColumnList `xml:"ResourceReportColumnList" json:"ResourceReportColumnList"`
      ResourceReportLineList *ResourceUsage_ResourceReportLineList `xml:"ResourceReportLineList" json:"ResourceReportLineList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceUsage_ResourceUsageContentType struct {
       Code *AUCodeSetsResourceUsageContentTypeType `xml:"Code" json:"Code"`
       LocalDescription *String `xml:"LocalDescription,omitempty" json:"LocalDescription"`
}
type ResourceUsage_ResourceReportColumnList struct {
      ResourceReportColumn []ResourceUsage_ResourceReportColumn `xml:"ResourceReportColumn" json:"ResourceReportColumn"`
}
type ResourceUsage_ResourceReportLineList struct {
      ResourceReportLine []ResourceUsage_ResourceReportLine `xml:"ResourceReportLine" json:"ResourceReportLine"`
}
type ResourceUsage_ResourceReportColumn struct {
       ColumnName *String `xml:"ColumnName" json:"ColumnName"`
       ColumnDescription *String `xml:"ColumnDescription,omitempty" json:"ColumnDescription"`
       ColumnDelimiter *String `xml:"ColumnDelimiter,omitempty" json:"ColumnDelimiter"`
}
type ResourceUsage_ResourceReportLine struct {
      SIF_RefId *ResourceUsage_SIF_RefId
       StartDate *String `xml:"StartDate" json:"StartDate"`
       EndDate *String `xml:"EndDate,omitempty" json:"EndDate"`
       CurrentCost *MonetaryAmountType `xml:"CurrentCost" json:"CurrentCost"`
       ReportRow *String `xml:"ReportRow" json:"ReportRow"`
}
type ResourceUsage_SIF_RefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
