package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Part struct {
	Name  string
	Email string
	Rsvp  string
}
type Meeting struct {
	Id           string
	Title        string
	Participants []Part
	Start_Time   time.Time
	End_Time     time.Time
	Timestamp    time.Time
}

func viewdocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/meeting/")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)

		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")
		fmt.Println(id)
		collection := client.Database("appointy_task").Collection("test")
		filter := bson.D{{"id", string(id)}}
		var result Meeting

		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Found a single document: %+v\n", result)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Make a get request please !"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete request not accepted"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func insertdocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		email := r.FormValue("participant")
		start := r.FormValue("start")
		end := r.FormValue("end")
		if email != "" {
			client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				log.Fatal(err)
			}
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}
			defer client.Disconnect(ctx)

			fmt.Println(email)
			fmt.Println(r)
			collection := client.Database("appointy_task").Collection("test")

			cursor, err := collection.Find(ctx, bson.M{
				"participants.email": string(email),
			})
			if err != nil {
				log.Fatal(err)
			}
			var meetings []bson.M
			if err = cursor.All(ctx, &meetings); err != nil {
				log.Fatal(err)
			}
			fmt.Println(meetings)
			json.NewEncoder(w).Encode(meetings)
		} else if start != "" && end != "" {
			client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				log.Fatal(err)
			}
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}
			defer client.Disconnect(ctx)

			fmt.Println(email)
			fmt.Println(r)
			collection := client.Database("appointy_task").Collection("test")

			cursor, err := collection.Find(ctx, bson.M{
				"start_time": bson.M{"$gt": start}, "end_time": bson.M{"$lt": end},
			})
			if err != nil {
				log.Fatal(err)
			}
			var meetings []bson.M
			if err = cursor.All(ctx, &meetings); err != nil {
				log.Fatal(err)
			}
			fmt.Println(meetings)
			json.NewEncoder(w).Encode(meetings)
			// w.Write([]byte(`{"message": "start and end"}`))
		} else {
			w.Write([]byte(`{"message": "Bad request, get request parameters should be either (email of participants) or (start and end times of meetings)"}`))
		}

	case "POST":
		w.WriteHeader(http.StatusCreated)

		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")
		collection := client.Database("appointy_task").Collection("test")

		fmt.Println("hello world")

		var meet1 Meeting
		_ = json.NewDecoder(r.Body).Decode(&meet1)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, _ := collection.InsertOne(ctx, meet1)
		json.NewEncoder(w).Encode(result)

		// meet2 := Meeting{
		// 	Id:    "1",
		// 	Title: "appointy",
		// 	Participants: []Part{
		// 		{
		// 			Name:  "mukesh",
		// 			Email: "chmukesh1612@gmail.com",
		// 			Rsvp:  "yes",
		// 		},
		// 	},
		// 	Start_Time: time.Date(2020, 11, 14, 8, 45, 16, 0, time.UTC),
		// 	End_Time:   time.Date(2020, 11, 14, 12, 45, 16, 0, time.UTC),
		// 	Timestamp:  time.Now(),
		// }
		// result2, _ := collection.InsertOne(ctx, meet2)
		// json.NewEncoder(w).Encode(result2)

	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Make a post request please !"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete request not accepted"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/meetings", insertdocs)
	http.HandleFunc("/meeting/", viewdocs)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
