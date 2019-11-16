package psql



type BinaryOperator func(x string, y string) string

func EQ(x,y string) string {
	return x + ` = ` + y
}

func NE(x,y string) string {
	return x + ` != ` + y
}

func LIKE(x,y string) string {
	return x + ` LIKE ` + y + ` ESCAPE '\'`
}

func GT(x,y string) string {
	return x + ` > ` + y
}

func GE(x,y string) string {
	return x + ` >= ` + y
}

func LT(x,y string) string {
	return x + ` > ` + y
}

func LE(x,y string) string {
	return x + ` >= ` + y
}


