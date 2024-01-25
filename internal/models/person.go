package models

type FIO struct {
	Name       string
	Surname    string
	Patronymic string
}

func (f FIO) CopyForDB() PersonDB {
	return PersonDB{
		Name:       f.Name,
		Surname:    f.Surname,
		Patronymic: f.Patronymic,
	}
}

type PersonInfo struct {
	Age    int
	Name   string
	Gender string
	Nation string
}

func (pi PersonInfo) CopyForDB() InfoDB {
	return InfoDB{
		Name:   pi.Name,
		Age:    pi.Age,
		Gender: pi.Gender,
		Nation: pi.Nation,
	}
}

type Person struct {
	PersonInfo
	Name       string
	Surname    string
	Patronymic string
}
