package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	api "github.com/nsip/sifxml2go/funcapi"
)

func main() {

	// new object
	asf := &api.AggregateStatisticFact{}

	asf.Characteristics().
		Append().SetValue("one"). // add array members
		Append().SetValue("two"). // shorthand
		Append().SetValue("three").
		Append().AggregateCharacteristicInfoRefId().SetValue("four"). // fully qualified
		AggregateCharacteristicInfoRefId().Append().SetValue("five")  // alternative ordering, same output

	//
	// properties can be set independently of objects
	// all values are interfaces, managed by setters
	var bob string = "bob"
	var rid api.RefIdType = "abc-def-hij"
	props := []api.Prop{
		api.Prop{"Excluded", "bob"}, // ok, will be coerced to pointer
		api.Prop{"Excluded", 3},     // fails, generates warning, can't coerce
		api.Prop{"Excluded", bob},   // ok, converts to poiner
		api.Prop{"Excluded", &bob},  // ok, is accepted
		api.Prop{"Value", 21.8},     // ok, float pointer type
		api.Prop{"RefId", rid},      // ok, will be coerced
		api.Prop{"RefId", &rid},     // ok, is accepted
	}
	asf.SetProperties(props...). // set properties in bulk, any not relevant will be discarded with warnings
					SetProperty("RefId", "ovverride-refid") // but can override specific

	// set internal object properties
	asf.StatsCohortYearLevelList().Append().CohortYearLevel().
		SetProperty("Code", "12")
	asf.StatsCohortYearLevelList().Append().CohortYearLevel().
		SetProperty("Code", "11")
	asf.StatsCohortYearLevelList().Last().CohortYearLevel().
		SetProperty("Code", "23") // fail, gnerates warning as is invalid codeset value

	// a full set of object properties
	scprops := []api.Prop{
		api.Prop{"StatsCohortId", "localid-a"},
		api.Prop{"StatsIndigenousStudentType", "type-1"},
		api.Prop{"CohortGender", "gender-x"},
		api.Prop{"DaysInReferencePeriod", 7},
		api.Prop{"PossibleSchoolDays", 25},
		api.Prop{"AttendanceDays", 12.5},
		api.Prop{"AttendanceLess90Percent", 6},
		api.Prop{"AttendanceGTE90Percent", 5},
		api.Prop{"PossibleSchoolDaysGT90PercentAttendance", 4},
	}
	asf.StatsCohortYearLevelList().Last().StatsCohortList().Append(). // add a new member
										StatsCohort().SetProperties(scprops...). // set in bulk
										SetProperty("AttendanceDays", 100.2)     // then override
	asf.StatsCohortYearLevelList().Last().StatsCohortList().Last(). // refer to same object
									SetProperties(scprops...).             // set in bulk..
									SetProperty("AttendanceDays", 10001.1) // then override
	asf.StatsCohortYearLevelList().Last().StatsCohortList().Append(). // add a new member
										StatsCohort().SetProperties(scprops...) // set in bulk
	//
	// render built object in json
	//
	b, err := json.Marshal(asf)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)

}
