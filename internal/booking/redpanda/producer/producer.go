package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/vitaliysev/mts_go_project/internal/models"
)

type Producer struct {
	client *kgo.Client
	topic  string
}

func New(brokers []string, topic string) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, err
	}
	return &Producer{client: client, topic: topic}, nil
}
func (p *Producer) SendMessage(msg models.CreateBookingResponse) {
	ctx := context.Background()
	b, _ := json.Marshal(msg)
	fmt.Println(b)
	p.client.Produce(ctx, &kgo.Record{Topic: p.topic, Value: b}, nil)
}
func (p *Producer) Close() {
	p.client.Close()
}
