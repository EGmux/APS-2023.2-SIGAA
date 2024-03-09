package data

type Direction int

// ENUMS
const (
	SQLTOSTRUCT Direction = iota
	STRUCTTOSQL
)

// Abs value for int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// Inplace, pass &, Convert a list of struct pointer to list of struct SQL, assuming both are equivalent
func MapSQLStruct[U SQLTablesPtrs, V TablePtrs](
	listSrcType *[]*U,
	listDestType *[]*V,
	mapping func(*V, *U, Direction) error,
	direction Direction,
	interval ...[]int,
) error {
	// NOTE: the subsequent functions expect empty listDestType
	if interval != nil {
		for _, pos := range interval {
			for _, tpos := range pos {
				_ = mapping((*listDestType)[tpos], (*listSrcType)[tpos], direction)
			}
		}
		return nil
	}
	for pos := range *listSrcType {
		_ = mapping((*listDestType)[pos], (*listSrcType)[pos], direction)
	}
	return nil
}
