package controller

import (
	log "github.com/sirupsen/logrus"
	"name_api/internal/benification"
	"name_api/internal/database"
	"name_api/internal/models"
	"net/http"
	"reflect"
)

type ControllerInterface interface {
	CreateEntity(person models.FIO) error
	UpdateEntity(id int, person models.FIO) error
	DeleteEntity(id int) error
	GetInfo(params models.RequestParam) ([]models.PersonDB, error)
}

type Controller struct {
	db database.DataBaseInterface
}

func NewController(db database.DataBaseInterface) ControllerInterface {
	return Controller{db: db}
}

func (c Controller) CreateEntity(person models.FIO) error {
	log.Infof("Create new Person: %v", person)
	info, err := c.db.GetInfoByName(person.Name)
	log.Debugf("controller.CreateEntity; Person Info in DB: %v", info)
	if reflect.DeepEqual(info, models.InfoDB{}) || err != nil {
		client := http.Client{}
		provider := benification.NewProvider(&client)

		newInfo := provider.GetInfo(person.Name)
		newInfoDB := newInfo.CopyForDB()
		log.Debugf("controller.CreateEntity; New Person Info for DB: %v", newInfoDB)
		if err := c.db.CreateInfo(newInfoDB); err != nil {
			return err
		}
	}

	personDB := person.CopyForDB()
	return c.db.CreatePerson(personDB)
}

func (c Controller) UpdateEntity(id int, person models.FIO) error {
	log.Infof("Update Person with id %d", id)
	personDB := person.CopyForDB()
	personDB.ID = id
	log.Debugf("controller.UpdateEntity; Person for DB:%v", personDB)

	return c.db.UpdatePerson(id, personDB)
}

func (c Controller) DeleteEntity(id int) error {
	log.Infof("Delete Person with id: %d", id)
	return c.db.DeletePerson(id)
}

func (c Controller) GetInfo(params models.RequestParam) ([]models.PersonDB, error) {
	people, err := c.db.GetPersonByParams(params)
	log.Debugf("controller.GetInfo; count of data recieved: %d", len(people))
	if err != nil {
		return nil, err
	}

	return people, nil
}
