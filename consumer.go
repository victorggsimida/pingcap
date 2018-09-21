package main

import (
	"context"
)

// Consumer consumes versioned data which produced by multi Producer
// NOTE: the implementation for Consumer is necessary
type Consumer struct {
	dataChannels []<-chan *VersionedData // producer output channels
	streamCount  int64                   // stream count for every producer output chan
	ch           chan *VersionedData     // merged result chan
}

// NewConsumer creates a new Consumer
func NewConsumer(channels []<-chan *VersionedData, streamCount int64) *Consumer {
	c := &Consumer{
		dataChannels: channels,
		streamCount:  streamCount,
		ch:           make(chan *VersionedData),
	}
	return c
}

// Start starts consuming versioned data
func (c *Consumer) Start(ctx context.Context) {
	// TODO: implement this to consume data from c.dataChannels, and merged result should send to c.ch
}

// MergedChan returns a chan which can be used to receive merged versioned data
func (c *Consumer) MergedChan() <-chan *VersionedData {
	return c.ch
}
