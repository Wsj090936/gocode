package structs

type Person struct {
	Age int `json:"age"`
	Tall int
}
// 注意 方法想要被外部调用，必须大写
func (person *Person) SetAge(age int)  {
	person.Age = age
}

func GetNewPerson() Person{
	p := Person{
		Age:  1,
		Tall: 0,
	}
	return p
}


