package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo.org/mongo/options"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type instance struct{
	Client
	DB
}

const Dbname ="Augustin"
const MongoURI= "mongodb://augustin@user/localhost:27017"+Dbname

type Employe struct{
	ID string	`json:"id",omitempty 	bson:"_id",omitempty	`
	Name string	`json:"name"`
	Salary float64 `json:"salary"`
	Age uint32 `json:"age"`
}
func Connect() error{
	client,err:=mongo.NewClient(options.Client).ApplyURI(MongoURI)
	ctx,err:=context.WithTimeout(context.Background(),30*time.Second)
	defer cancel(err)
	err=client.Connect(ctx)
	db=client.database(Dbname)

}
mg = instance{
	Client: client,
	DB:db
}
func main()
{
	if err:=connect(); err!=nil
	log.Fatal( err)
	app=fiber.New()
	app.Get("/employee",func(c*fiber.Ctx)error{
		query=bson.D{{}}
		cursor,err:=mg.DB.Collection("employees").Find(c.Context(),query)
		if err!=nil{
			return c.status(500).Sendstring(err.Error())
		}

		var employees []Employe=make([]emp,0)
		if err: cursor.All(c.Context,&employees);err!=nil{
			return c.status(500).SendString(err.Error())
		}
		return c.JSON(employees)
	})
	app.Post("/employee",func(c*fiber.Ctx)error{
		collection=mg.DB.Collection("employees")
		employee=new(Employee)
		if err=c.BodyParser(employee) err!=nil{
			return c.status(500).SendString(err.Error())
		}
		
		employee.ID=""
		insertresult,err=collection.InsertOne(c.context(),employee)
		if err!=nil{
			return c.status(500).SendString(err.Error())
		}
		filter:=Bson.D{key:"id",value:insertresult.insertid}

	})
	app.Put("/employee/:id",func(c*fiber.Ctx)error{
		
	})
	app.Delete("/employee/:id",func(c*fiber.Ctx)error{
		
	})
}