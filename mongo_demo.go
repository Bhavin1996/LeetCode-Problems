// Q6oUxOWcubB8fnyg

package leetcode

import (
	// Add required Go packages
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	// Add the MongoDB driver packages
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

username := "sudhanshubawane"

password := "Crossbone@100st"

cluster := "cluster0"

authSource  := "nlgg6dn"

authMechanism := "mongodb.net/?retryWrites=true&w=majority"

uri := "mongodb+srv://" + url.QueryEscape(username) + ":" + 

		url.QueryEscape(password) + "@" + cluster + 

		"/?authSource=" + authSource +

		"&authMechanism=" + authMechanism


// Your MongoDB Atlas Connection String
const uri = "mongodb+srv://sudhanshubawane:Crossbone@100st@cluster0.nlgg6dn.mongodb.net/?retryWrites=true&w=majority"

// A global variable that will hold a reference to the MongoDB client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

// Our implementation logic for connecting to MongoDB
func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

// GET /movies - Get all movies
func getMovies(c *gin.Context) {
	// Find movies
	cursor, err := mongoClient.Database("sample_mflix").Collection("movies").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map results
	var movies []bson.M
	if err = cursor.All(context.TODO(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movies
	c.JSON(http.StatusOK, movies)
}

func main()_mongo {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/movies", getMovies)
	//r.GET("/movies/:id", getMovieByID)
	//r.POST("/movies/aggregations", aggregateMovies)

	r.Run()
}
