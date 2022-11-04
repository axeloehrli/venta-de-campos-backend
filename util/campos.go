package util

import (
	"fmt"
	"reflect"
	"strings"
)

// Converts filters struct to map
func FiltersStructToMap(myStruct any) map[string]any {
	filtersMap := make(map[string]any)

	v := reflect.ValueOf(myStruct)
	typeOfS := v.Type()

	// Iterates through struct's fields
	for i := 0; i < v.NumField(); i++ {
		// Struct field name
		fieldName := typeOfS.Field(i).Tag.Get("form")
		fmt.Println("FORM TAG: ", typeOfS.Field(i).Tag.Get("form"))
		// Struct field value
		fieldValue := v.Field(i).Interface()
		// Checks if fields are not LIMIT (PageSize) or OFFSET (PageID)
		// and whether the field's value is nil
		if fieldName != "page_id" && fieldName != "page_size" && !v.Field(i).IsNil() {

			// Checks if field value is a string pointer
			if strPointer, ok := fieldValue.(*string); ok {
				filtersMap[fieldName] = *strPointer
			}

			// Checks if field value is an int pointer
			if intPointer, ok := fieldValue.(*int64); ok {
				filtersMap[fieldName] = *intPointer
			}

		}

	}
	return filtersMap
}

func BuildDynamicCamposQuery(filtersMap map[string]any, limit int32, offset int32) string {

	count := 0

	// Slice of conditions for the dynamic query
	var conditions []string
	// Loops over filters map
	for key, value := range filtersMap {
		// If it's the first filter, the condition should start with 'WHERE'
		if count == 0 {

			// Checks if the filter is a min or max field, in which case
			// it will use the < or > operator instead of ==
			if strings.Contains(key, "min") {
				// trims the _min part from the filter key
				trimmedKey := strings.ReplaceAll(key, "_min", "")
				condition := fmt.Sprintf("WHERE %v > %v", trimmedKey, value)
				conditions = append(conditions, condition)
				count++
				continue
			}
			if strings.Contains(key, "max") {
				// trims the _max part from the filter key
				trimmedKey := strings.ReplaceAll(key, "_max", "")
				condition := fmt.Sprintf("WHERE %v < %v", trimmedKey, value)
				conditions = append(conditions, condition)
				count++
				continue
			}

			// Checks if the filter's value is a string, if it is,
			// surrounds it with single quotes
			if str, ok := value.(string); ok {
				condition := fmt.Sprintf("WHERE %v = '%v'", key, str)
				conditions = append(conditions, condition)
				count++
			} else {
				condition := fmt.Sprintf("WHERE %v = %v", key, value)
				conditions = append(conditions, condition)
				count++
			}
		} else {
			// If the count is not 0, it means it's not the first filter,
			// so the condition will start with AND instead of WHERE

			// Checks if filter is a min or max field, in case it
			// will use the > or < operators instead of ==
			if strings.Contains(key, "min") {
				// trims the _min part from the filter's key
				trimmedKey := strings.ReplaceAll(key, "_min", "")
				condition := fmt.Sprintf("AND %v > %v", trimmedKey, value)
				conditions = append(conditions, condition)
				count++
				continue
			}
			if strings.Contains(key, "max") {
				// trims the _max part from the filter's key
				trimmedKey := strings.ReplaceAll(key, "_max", "")
				condition := fmt.Sprintf("AND %v < %v", trimmedKey, value)
				conditions = append(conditions, condition)
				count++
				continue
			}
			// Checks if the filter's value is a string, if it is,
			// surrounds it with single quotes
			if str, ok := value.(string); ok {
				condition := fmt.Sprintf("AND %v = '%v'", key, str)
				conditions = append(conditions, condition)
				count++
			} else {
				condition := fmt.Sprintf("AND %v = %v", key, value)
				conditions = append(conditions, condition)
				count++
			}
		}
	}
	fullCondition := fmt.Sprintf("SELECT * FROM campos %v LIMIT %v OFFSET %v", strings.Join(conditions, " "), limit, offset)
	return fullCondition
}
