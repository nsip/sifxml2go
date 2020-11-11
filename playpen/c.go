package main

import (
	"../sifxml"
	"encoding/xml"
	"fmt"
	"log"
)

func errcheck(err error) {
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

func main() {
	var err error
	new := sifxml.StudentAttendanceCollection{}
	output, err := xml.MarshalIndent(new, "", "  ")
	errcheck(err)
	fmt.Println(string(output))

	new = sifxml.StudentAttendanceCollection{}
	new.Set("RefId", "ABC")
	new.Set("RoundCode", "DEF")
	new.Set("ReportingAuthorityCommonwealthId", "GHI")
	/*
		new.Characteristics = new.Characteristics.Add("XYZ")
		new.Characteristics = new.Characteristics.Add("UVW")
	*/

	s := sifxml.StatsCohortType{}
	s.Set("StatsCohortId", "1")
	s.Set("DaysInReferencePeriod", 5)
	s_arr := sifxml.StatsCohortListType{}
	s_arr.Add(s)
	y := sifxml.YearLevelType{}
	y.Set("Code", "12")
	s1 := sifxml.StatsCohortYearLevelType{}
	s1.Set("CohortYearLevel", y)
	s1.Set("StatsCohortList", s_arr)
	sa := sifxml.StudentAttendanceCollectionReportingType{}
	sa.StatsCohortYearLevelList = sa.StatsCohortYearLevelList.Add(s1)
	sa_arr := sifxml.StudentAttendanceCollectionReportingListType{}
	sa_arr.Add(sa)
	new.Set("StudentAttendanceCollectionReportingList", sa_arr)

	new.StudentAttendanceCollectionReportingList = new.StudentAttendanceCollectionReportingList.AddNew()
	new.StudentAttendanceCollectionReportingList.Last().Set("EntityLevel", "2222")
	new.StudentAttendanceCollectionReportingList.Last().EntityContact = new.StudentAttendanceCollectionReportingList.Last().EntityContact.New()
	new.StudentAttendanceCollectionReportingList.Last().EntityContact.Set("PositionTitle", "Teacher")
	new.StudentAttendanceCollectionReportingList.Last().EntityContact.Name = new.StudentAttendanceCollectionReportingList.Last().EntityContact.Name.New()
	new.StudentAttendanceCollectionReportingList.Last().EntityContact.Name.Set("Type", "LGL")
	new.StudentAttendanceCollectionReportingList.Last().EntityContact.Name.Set("FamilyName", "Jackson")
	new.StudentAttendanceCollectionReportingList.Last().EntityContact.Name.Set("FamilyNameFirst", "Y")
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList = new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.AddNew()
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().CohortYearLevel = new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().CohortYearLevel.New()
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().CohortYearLevel.Set("Code", "11")
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().StatsCohortList = new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().StatsCohortList.AddNew()
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().StatsCohortList.Last().Set("StatsCohortId", "22")
	output, err = xml.MarshalIndent(new, "", "  ")
	errcheck(err)
	fmt.Println(string(output))

	/*
		StudentAttendanceCollectionReportingList :
		  EntityLevel: "2222"
		  EntityContact:
		     PositionTitle: Teacher
		     Name:
		       Type: "LGL"
	*/

	/* OUTPUT GENERATED:

	<StudentAttendanceCollection RefId="ABC">
	  <RoundCode>DEF</RoundCode>
	  <ReportingAuthorityCommonwealthId>GHI</ReportingAuthorityCommonwealthId>
	  <StudentAttendanceCollectionReportingList>
	    <StudentAttendanceCollectionReporting>
	      <StatsCohortYearLevelList>
	        <StatsCohortYearLevel>
	          <CohortYearLevel>
	            <Code>12</Code>
	          </CohortYearLevel>
	          <StatsCohortList>
	            <StatsCohort>
	              <StatsCohortId>1</StatsCohortId>
	              <DaysInReferencePeriod>5</DaysInReferencePeriod>
	            </StatsCohort>
	          </StatsCohortList>
	        </StatsCohortYearLevel>
	      </StatsCohortYearLevelList>
	    </StudentAttendanceCollectionReporting>
	    <StudentAttendanceCollectionReporting>
	      <EntityLevel>2222</EntityLevel>
	      <EntityContact>
	        <Name Type="LGL">
	          <FamilyName>Jackson</FamilyName>
	          <FamilyNameFirst>Y</FamilyNameFirst>
	        </Name>
	        <PositionTitle>Teacher</PositionTitle>
	      </EntityContact>
	      <StatsCohortYearLevelList>
	        <StatsCohortYearLevel>
	          <CohortYearLevel>
	            <Code>11</Code>
	          </CohortYearLevel>
	          <StatsCohortList>
	            <StatsCohort>
	              <StatsCohortId>22</StatsCohortId>
	            </StatsCohort>
	          </StatsCohortList>
	        </StatsCohortYearLevel>
	      </StatsCohortYearLevelList>
	    </StudentAttendanceCollectionReporting>
	  </StudentAttendanceCollectionReportingList>
	</StudentAttendanceCollection>

	*/
	new.StudentAttendanceCollectionReportingList.Last().StatsCohortYearLevelList.Last().CohortYearLevel.Set("Code", "FRED")
	/* crashes */
}
