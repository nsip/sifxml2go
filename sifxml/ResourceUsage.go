package sifxml


type ResourceUsages []ResourceUsage

    type ResourceUsage struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      ResourceUsageContentType *ResourceUsage_ResourceUsageContentType `xml:"ResourceUsageContentType,omitempty" json:"ResourceUsageContentType"`
      ResourceReportColumnList *ResourceUsage_ResourceReportColumnList `xml:"ResourceReportColumnList,omitempty" json:"ResourceReportColumnList"`
      ResourceReportLineList *ResourceUsage_ResourceReportLineList `xml:"ResourceReportLineList,omitempty" json:"ResourceReportLineList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceUsage_ResourceUsageContentType struct {
       Code *string `xml:"Code,omitempty" json:"Code"`
       LocalDescription *string `xml:"LocalDescription,omitempty" json:"LocalDescription"`
}
type ResourceUsage_ResourceReportColumnList struct {
      ResourceReportColumn *[]ResourceUsage_ResourceReportColumn `xml:"ResourceReportColumn,omitempty" json:"ResourceReportColumn"`
}
type ResourceUsage_ResourceReportLineList struct {
      ResourceReportLine *[]ResourceUsage_ResourceReportLine `xml:"ResourceReportLine,omitempty" json:"ResourceReportLine"`
}
type ResourceUsage_ResourceReportColumn struct {
       ColumnName *string `xml:"ColumnName,omitempty" json:"ColumnName"`
       ColumnDescription *string `xml:"ColumnDescription,omitempty" json:"ColumnDescription"`
       ColumnDelimiter *string `xml:"ColumnDelimiter,omitempty" json:"ColumnDelimiter"`
}
type ResourceUsage_ResourceReportLine struct {
      SIF_RefId *ResourceUsage_SIF_RefId `xml:"SIF_RefId,omitempty" json:"SIF_RefId"`
       StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
       EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
       CurrentCost *MonetaryAmountType `xml:"CurrentCost,omitempty" json:"CurrentCost"`
       ReportRow *string `xml:"ReportRow,omitempty" json:"ReportRow"`
}
type ResourceUsage_SIF_RefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
