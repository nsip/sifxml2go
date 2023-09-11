package sifxml



    type EquipmentInfos struct {
      equipmentinfos `json:"EquipmentInfos"`
    }

    type equipmentinfos struct {
      EquipmentInfo []equipmentinfo `json:"EquipmentInfo"`
    }

    type EquipmentInfo struct {
      equipmentinfo `xml:"EquipmentInfo" json:"EquipmentInfo"`
     }

     type equipmentinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      AssetNumber *LocalIdType `xml:"AssetNumber,omitempty" json:"AssetNumber,omitempty"`
      InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId,omitempty"`
      PurchaseOrderRefId *String `xml:"PurchaseOrderRefId,omitempty" json:"PurchaseOrderRefId,omitempty"`
      EquipmentType *String `xml:"EquipmentType,omitempty" json:"EquipmentType,omitempty"`
      SIF_RefId *EquipmentInfo_SIF_RefId
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type EquipmentInfo_SIF_RefId struct {
  equipmentinfo_sif_refid `xml:"EquipmentInfo_SIF_RefId" json:"EquipmentInfo_SIF_RefId"`
}

type equipmentinfo_sif_refid struct {

      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

