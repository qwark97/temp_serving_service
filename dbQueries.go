package main

var selectStatements SQLSelectStatements

func (statements *SQLSelectStatements) loadStatements() {
	// todo lack of specifying from the what source temperature measurements should be fetchetched (sensor1, sensor2, etc.)
	statements.Statements["currentTemp"] = "SELECT temperature, unit FROM temperatures ORDER BY temp_id DESC LIMIT 1"
	statements.Statements["averageDay"] = ""
	statements.Statements["average7Days"] = ""
	statements.Statements["average30Days"] = ""
}

func (statements *SQLSelectStatements) getStatement(category string) string {
	val := statements.Statements[category]
	if val == "" {
		panic("Invalid SQL statement category")
	}
	return val
}
