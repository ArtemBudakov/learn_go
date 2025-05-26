package stage_5

import (
	"fmt"
	"math/rand"
)

func StageFive() {
	emergencyContacts, addresses, persons := FillStructures()
	PrintStructure(emergencyContacts)
	PrintStructure(addresses)
	PrintStructure(persons)
}

func CreateEmergencyContacts() []EmergencyContact {
	var emergencyContacts []EmergencyContact
	contactFirst := EmergencyContact{
		Name:  "First Emergency Contact",
		Phone: "0123456789",
	}
	contactSecond := EmergencyContact{
		Name:  "Second Emergency Contact",
		Phone: "0987654321",
	}
	emergencyContacts = append(emergencyContacts, contactFirst)
	emergencyContacts = append(emergencyContacts, contactSecond)

	return emergencyContacts
}

func CreateAddressContacts() []Address {
	var addresses []Address
	addressFirst := Address{
		Street: "Street_1",
		City:   "City_1",
		State:  "State_1",
		Zip:    123456,
	}
	addressSecond := Address{
		Street: "Street_2",
		City:   "City_2",
		State:  "State_2",
		Zip:    234567,
	}
	addresses = append(addresses, addressFirst)
	addresses = append(addresses, addressSecond)

	return addresses
}

func CreatePersons(addresses []Address, emergencyContact []EmergencyContact) []Person {
	var persons []Person

	for index, address := range addresses {
		indexPlusOne := index + 1

		person := Person{
			FirstName:        fmt.Sprintf("First %d", indexPlusOne),
			LastName:         fmt.Sprintf("Last %d", indexPlusOne),
			Age:              (indexPlusOne) * 10,
			Height:           float32(indexPlusOne*10 + indexPlusOne/10),
			IsEmployed:       true,
			Salary:           float32(indexPlusOne*1000 + indexPlusOne/10 + indexPlusOne/100),
			Address:          address,
			Skills:           []string{"some skill", "some another skill"},
			Children:         []Person{},
			EmergencyContact: emergencyContact[rand.Intn(len(emergencyContact))],
		}
		persons = append(persons, person)
	}
	return persons
}

func FillStructures() ([]EmergencyContact, []Address, []Person) {
	emergencyContacts := CreateEmergencyContacts()
	addresses := CreateAddressContacts()
	persons := CreatePersons(addresses, emergencyContacts)

	return emergencyContacts, addresses, persons
}

func PrintStructure[T any](structures []T) {
	for _, structure := range structures {
		fmt.Println(structure)
	}
}
