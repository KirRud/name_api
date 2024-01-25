package database

import "name_api/internal/models"

const pageSize = 25

type personRepoInterface interface {
	CreatePerson(person models.PersonDB) error
	UpdatePerson(id int, person models.PersonDB) error
	DeletePerson(id int) error
}

func (d *DataBase) CreatePerson(person models.PersonDB) error {
	return d.db.Create(&person).Error
}

func (d *DataBase) UpdatePerson(id int, person models.PersonDB) error {
	return d.db.Where("id = ?", id).
		Updates(models.PersonDB{Name: person.Name, Surname: person.Surname, Patronymic: person.Patronymic}).Error
}

func (d *DataBase) DeletePerson(id int) error {
	return d.db.Where("id = ?", id).Delete(&models.PersonDB{}).Error
}

type infoRepoInterface interface {
	CreateInfo(info models.InfoDB) error
	GetInfoByName(name string) (models.InfoDB, error)
	DeleteInfo(name string) error
}

func (d *DataBase) CreateInfo(info models.InfoDB) error {
	return d.db.Create(&info).Error
}

func (d *DataBase) GetInfoByName(name string) (models.InfoDB, error) {
	var info models.InfoDB
	if err := d.db.Where("name = ?", name).First(&info).Error; err != nil {
		return models.InfoDB{}, err
	}
	return info, nil
}

func (d *DataBase) DeleteInfo(name string) error {
	return d.db.Where("name = ?", name).Delete(&models.InfoDB{}).Error
}

type joinRepoInterface interface {
	GetPersonByParams(params models.RequestParam) ([]models.PersonDB, error)
}

func (d *DataBase) GetPersonByParams(params models.RequestParam) ([]models.PersonDB, error) {
	var people []models.PersonDB
	err := d.db.Where(params.GetSqlWhere(), params.GetParams).Joins("LEFT JOIN info_dbs ON info_dbs.name = person_dbs.name").
		Offset((params.GetPage() - 1) * pageSize).Limit(pageSize).Find(&people).Error
	if err != nil {
		return nil, err
	}
	return people, nil
}
