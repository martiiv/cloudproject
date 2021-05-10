package webhooks

import (
	"cloud.google.com/go/firestore"
	"cloudproject/extra"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

/**
 * Class database.go
 * Will contain all database related functionality
 * Contains the following functions:
 *									Init() 		For initializing the database connection
 *									Add() 		For adding an entry to the database
 *									Delete() 	For deleting an entry from the database
 *									Update() 	For updating an entry in the database
 * @author Martin Iversen
 * @date 01.05.2021
 * @version 0.1
 */

//Initializing DB
var Ctx context.Context
var Client *firestore.Client

//const Collection = "RouteInformation" //Defining the name of the collection in FireStore

/*
 * Function for initializing the database, will be used when starting the app
 */
func Init() error {
	// Firebase initialisation
	Ctx = context.Background()

	// Authenticate with key file from firebase
	opt := option.WithCredentialsFile("webhooks/trafficmessage.json")
	app, err := firebase.NewApp(Ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing DataBase: %v", err)
	}

	Client, err = app.Firestore(Ctx)
	if err != nil {
		return fmt.Errorf("error occurred initializing Client: %v", err)
	}

	return nil
}

/*
 * Function for adding RouteInformation to the database
 * Returns the ID an object is given when the database creates
 */
func Add(webhook extra.Webhook) (string, error) {
	newEntry, _, err := Client.Collection(Collection).Add(Ctx, webhook) //Adds RouteInformation
	if err != nil {
		return "", errors.New("Error occurred when adding RouteInformation to database: " + err.Error())
	}
	return newEntry.ID, nil //Returns the id of the entry in the database collection
}

/*
 * Function for deleting a webhook from the database
 */
func Delete(id string) error {
	_, err := Client.Collection(Collection).Doc(id).Delete(Ctx) //Deletes from the database

	if err != nil {
		return errors.New("Error occurred when trying to delete entry. Entry ID: " + id)
	}
	return nil
}

/**
 * Function Get
 * Used for selecting a specific DB entry
 */
func Get(id string) (error, map[string]interface{}) {
	dbSnapShot, err := Client.Collection(Collection).Doc(id).Get(Ctx)
	if err != nil {
		return fmt.Errorf("Error occurred There is no document in the db with the id: %v!", id), nil
	}

	entry := dbSnapShot.Data()
	return nil, entry
}

/*
 * Function for getting all entries in a database
 * Source: https://stackoverflow.com/a/61429531
 * Returns an object containing document snapshots of the entries in the database
 */
func GetAll() ([]*firestore.DocumentSnapshot, error) {
	var docs []*firestore.DocumentSnapshot               //Defining object to be returned
	iter := Client.Collection(Collection).Documents(Ctx) //Gets all entries in the database
	for {                                                //Iterates through the database
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil //Returns a list of entries
}

/*
 * Function for updating information on an entry in the database
 */
func Update(id string, data interface{}) error {
	_, err := Client.Collection(Collection).Doc(id).Set(Ctx, data)
	if err != nil {
		return errors.New("Error while updating Route Information entry in the database: " + err.Error())
	}
	return nil
}
