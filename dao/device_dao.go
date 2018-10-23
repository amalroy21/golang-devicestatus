package dao

import(
		"log"
		mgo "gopkg.in/mgo.v2"
		//"gopkg.in/mgo.v2/bson"
		."DcmStatusReceiver/models"
)

type DeviceDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "device"
)

//Establish a connection to database
func (d *DeviceDAO) Connect(){
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}

//Update an existing device
func (d *DeviceDAO) Update(device Device) error {
	err := db.C(COLLECTION).UpdateId(device.ID, &device)
	return err
}

//Insert new records
func (m *DeviceDAO) Insert(device Device) error {
	err := db.C(COLLECTION).Insert(&device)
	return err
}