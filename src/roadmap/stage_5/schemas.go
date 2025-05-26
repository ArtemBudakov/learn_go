package stage_5

type EmergencyContact struct {
	Name  string
	Phone string
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    int32
}

type Person struct {
	FirstName        string
	LastName         string
	Age              int
	Height           float32
	IsEmployed       bool
	Salary           float32
	Address          Address
	Skills           []string
	Children         []Person
	EmergencyContact EmergencyContact
}
