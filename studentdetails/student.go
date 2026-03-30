package studentdetails

type Student struct {
	Name string
	Age int
	Marks float32
}

func PrintStudent() map[int]Student{
	students := map[int]Student{
		1: 	{"ravi" , 25, 54.23},
		2:	{"sanji", 23, 50.05},
		3:	{"zoro", 21, 23.12},
	}
	return students
	
} 
	
