package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Person struct {
	ID        string `bson:"_id" json:"id"`
	Firstname string `bson:"firstname" json:"firstname"`
	Lastname  string `bson:"lastname" json:"lastname"`
}

var people []Person

func main() {

	people = append(people, Person{ID: "1", Firstname: "Bruce", Lastname: "Wayne"})
	people = append(people, Person{ID: "2", Firstname: "Ioannis", Lastname: "Petrousov"})
	person := Person{ID: "2", Firstname: "Ioannis", Lastname: "Petrousov"}

	// Setup client and connection
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//  Database returns a handle for a given database
	// Collection types can be used to access the database
	collection := client.Database("baz").Collection("qux")

	var ppl []interface{}
	for _, p := range people {
		ppl = append(ppl, p)
	}

	// Insert documents to database
	_, err = collection.InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}

	//  Finds the documents matching a model
	cur, err := collection.Find(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	// Get the next result from the cursor.
	// Returns true if there were no errors and there is a next result.
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}

		// do something with element

		fmt.Printf("Keys of document: ")
		fmt.Println(elem.Keys(false))

		fmt.Printf("Length of document: ")
		fmt.Println(elem.Len())

		// Interface returns the Go value of this Value as an empty interface.
		v := elem.Lookup("firstname")
		fmt.Println(v.Interface())

		// Returns the string balue for this element
		fmt.Println(v.StringValue())

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Delete document from DB
	collection.DeleteOne(context.Background(), person)

	// Close connection to DB
	client.Disconnect(context.Background())
	fmt.Println("COnnection terminated")
}
