package helper

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToPrimitiveObj(id string) (primitive.ObjectID, error) {

	obj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return obj, nil
}
