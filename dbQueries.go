package main

var selectStatements SQLSelectStatements

func (statements *SQLSelectStatements) loadStatements() {
	statements.Statements["currentTemp"] = ""
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
