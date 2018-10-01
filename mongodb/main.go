package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// CONNECTIONSTRING Connection to mongodb server
const CONNECTIONSTRING string = "mongodb://localhost:27017"

// Person is the document to insert into DB
type Person struct {
	ID        objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Firstname string            `bson:"firstname" json:"firstname"`
	Lastname  string            `bson:"lastname" json:"lastname"`
}

var people []Person

func showAll(collection *mongo.Collection) {
	// 3. Show all documents
	fmt.Println("Show all documents")
	// Return to us a cursor which we'll have to walk over
	cur, err := collection.Find(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	var elements []interface{}
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	fmt.Println(elements)
}

func deleteAll(collection *mongo.Collection) {
	fmt.Println("Delete all documents")
	res, err := collection.DeleteMany(context.Background(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func main() {

	people = append(people, Person{ID: objectid.New(), Firstname: "Bruce", Lastname: "Wayne"})
	people = append(people, Person{ID: objectid.New(), Firstname: "Clark", Lastname: "Kent"})

	// Setup connection and DB client
	client, err := mongo.NewClient(CONNECTIONSTRING)
	if err != nil {
		log.Fatal(err)
	}

	// Initializes the Client by starting background monitoring goroutines
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//  Database returns a handle for a given database
	// Collection types can be used to access the database
	collection := client.Database("baz").Collection("qux")

	// Insert multiple documents
	fmt.Println("Insert multiple documents")
	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}

	_, err = collection.InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}

	showAll(collection)

	// Insert one document
	fmt.Println("Insert one document")
	person := Person{ID: objectid.New(), Firstname: "Ioannis", Lastname: "Petrousov"}
	_, err = collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}

	showAll(collection)

	// Delete all documents
	deleteAll(collection)
	showAll(collection)

	// Close connection to DB
	client.Disconnect(context.Background())
	fmt.Println("COnnection terminated")
}
