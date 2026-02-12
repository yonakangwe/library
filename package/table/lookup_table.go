package table

type LookupTable uint8

const (
	USERCATEGORY LookupTable = iota
	ROLE
)

var lookupTableName = map[string]LookupTable{
	"user_category": USERCATEGORY,
	"roles":         ROLE,
}

func TableExists(tableName string) bool {
	_, exists := lookupTableName[tableName]
	return exists
}
