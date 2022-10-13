package sifxml


type Activitys []Activity

    type Activity struct {
  activity `xml:"Activity" json:"Activity"`
}

type activity struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *String `xml:"Title,omitempty" json:"Title,omitempty"`
      Preamble *String `xml:"Preamble,omitempty" json:"Preamble,omitempty"`
      TechnicalRequirements *TechnicalRequirementsType `xml:"TechnicalRequirements,omitempty" json:"TechnicalRequirements,omitempty"`
      SoftwareRequirementList *SoftwareRequirementListType `xml:"SoftwareRequirementList,omitempty" json:"SoftwareRequirementList,omitempty"`
      EssentialMaterials *EssentialMaterialsType `xml:"EssentialMaterials,omitempty" json:"EssentialMaterials,omitempty"`
      LearningObjectives *LearningObjectivesType `xml:"LearningObjectives,omitempty" json:"LearningObjectives,omitempty"`
      LearningStandards *LearningStandardsType `xml:"LearningStandards,omitempty" json:"LearningStandards,omitempty"`
      SubjectArea *SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea,omitempty"`
      Prerequisites *PrerequisitesType `xml:"Prerequisites,omitempty" json:"Prerequisites,omitempty"`
      Students *StudentsType `xml:"Students,omitempty" json:"Students,omitempty"`
      SourceObjects *SourceObjectsType `xml:"SourceObjects,omitempty" json:"SourceObjects,omitempty"`
      Points *Int `xml:"Points,omitempty" json:"Points,omitempty"`
      ActivityTime *ActivityTimeType `xml:"ActivityTime" json:"ActivityTime"`
      AssessmentRefId *String `xml:"AssessmentRefId,omitempty" json:"AssessmentRefId,omitempty"`
      MaxAttemptsAllowed *Int `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed,omitempty"`
      ActivityWeight *Float `xml:"ActivityWeight,omitempty" json:"ActivityWeight,omitempty"`
      Evaluation *Activity_Evaluation
      LearningResources *LearningResourcesType `xml:"LearningResources,omitempty" json:"LearningResources,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type Activity_Evaluation struct {
  activity_evaluation `xml:"Activity_Evaluation" json:"Activity_Evaluation"`
}

type activity_evaluation struct {
      EvaluationType *String `xml:"EvaluationType,attr" json:"EvaluationType"`
       Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
}

