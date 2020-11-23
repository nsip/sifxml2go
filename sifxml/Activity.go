package sifxml


type Activitys []Activity

    type Activity struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *String `xml:"Title,omitempty" json:"Title"`
      Preamble *String `xml:"Preamble,omitempty" json:"Preamble"`
      TechnicalRequirements *TechnicalRequirementsType `xml:"TechnicalRequirements,omitempty" json:"TechnicalRequirements"`
      SoftwareRequirementList *SoftwareRequirementListType `xml:"SoftwareRequirementList,omitempty" json:"SoftwareRequirementList"`
      EssentialMaterials *EssentialMaterialsType `xml:"EssentialMaterials,omitempty" json:"EssentialMaterials"`
      LearningObjectives *LearningObjectivesType `xml:"LearningObjectives,omitempty" json:"LearningObjectives"`
      LearningStandards *LearningStandardsType `xml:"LearningStandards,omitempty" json:"LearningStandards"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      Prerequisites *PrerequisitesType `xml:"Prerequisites,omitempty" json:"Prerequisites"`
      Students *StudentsType `xml:"Students,omitempty" json:"Students"`
      SourceObjects *SourceObjectsType `xml:"SourceObjects,omitempty" json:"SourceObjects"`
      Points *Int `xml:"Points,omitempty" json:"Points"`
      ActivityTime *ActivityTimeType `xml:"ActivityTime" json:"ActivityTime"`
      AssessmentRefId *String `xml:"AssessmentRefId,omitempty" json:"AssessmentRefId"`
      MaxAttemptsAllowed *Int `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed"`
      ActivityWeight *Float `xml:"ActivityWeight,omitempty" json:"ActivityWeight"`
      Evaluation *Activity_Evaluation
      LearningResources *LearningResourcesType `xml:"LearningResources,omitempty" json:"LearningResources"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Activity_Evaluation struct {
      EvaluationType *String `xml:"EvaluationType,attr" json:"EvaluationType"`
       Description *String `xml:"Description,omitempty" json:"Description"`
}
