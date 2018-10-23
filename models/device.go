package models

import "gopkg.in/mgo.v2/bson"

// Represents a device, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document

type Device struct{
	ID          	bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Status        	string        `bson:"status" json:"status"`
	Description   	string        `bson:"description" json:"description"`	

}