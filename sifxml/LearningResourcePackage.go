package sifxml



    type LearningResourcePackages struct {
      learningresourcepackages `json:"LearningResourcePackages"`
    }

    type learningresourcepackages struct {
      LearningResourcePackage []learningresourcepackage `json:"LearningResourcePackage"`
    }

    type learningresourcepackage LearningResourcePackage

    type LearningResourcePackage AbstractContentElementType

