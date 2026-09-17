package service

import (
	"github.com/yoshihio/go-memps/config"
	"github.com/yoshihio/go-memps/entity"
)

type QueueService interface {
	Enqueue(data entity.Data)
	Dequeue() entity.Data
	IsEmpty() bool
}

type Queue struct {
	data []entity.Data
}

func NewQueue(cfg config.Config) QueueService {
	return &Queue{
		data: make([]entity.Data, cfg.QueueSize),
	}
}

func (q *Queue) Enqueue(data entity.Data) {
	q.data = append(q.data, data)
}

func (q *Queue) Dequeue() entity.Data {
	data := q.data[0]
	q.data = q.data[1:]
	return data
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
