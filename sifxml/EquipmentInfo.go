package sifxml


type EquipmentInfos []EquipmentInfo

    type EquipmentInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      AssetNumber *LocalIdType `xml:"AssetNumber,omitempty" json:"AssetNumber"`
      InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      PurchaseOrderRefId *String `xml:"PurchaseOrderRefId,omitempty" json:"PurchaseOrderRefId"`
      EquipmentType *String `xml:"EquipmentType,omitempty" json:"EquipmentType"`
      SIF_RefId *EquipmentInfo_SIF_RefId
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type EquipmentInfo_SIF_RefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
