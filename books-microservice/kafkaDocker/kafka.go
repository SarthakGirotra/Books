package kafkaDocker

import (
	"b/container"
	"b/models"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context, books []models.Books) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "Users",
	})
	msgBytes, errJ := json.Marshal(books)
	if errJ != nil {
		fmt.Println(errJ)
	}

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(1)),

		Value: msgBytes,
	})
	if err != nil {
		fmt.Println("Could not write msg", err)
	}
}

func Consume(ctx context.Context, container container.Container) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "Users",
		GroupID: "my-group",
	})
	for {

		msg, err := r.ReadMessage(ctx)
		msgJson := &[]models.Books{}
		json.Unmarshal(msg.Value, msgJson)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		collection := container.GetDB().Db.Collection("books")
		var bi []interface{}
		for _, b := range *msgJson {
			bi = append(bi, b)
		}
		_, errI := collection.InsertMany(context.Background(), bi)
		if errI != nil {
			fmt.Println(errI)
		}
		fmt.Println("Added to DB")
	}
}
