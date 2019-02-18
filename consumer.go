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


var versionId = 1


type versionIdData struct {
	produceId []int64
	streamMap map[int64][]int64
	versionId int64
	data []VersionedData
	retry int
}

var tmpDataList map[int64]*versionIdData
// Start starts consuming versioned data
func (c *Consumer) Start(ctx context.Context) {
	// TODO: implement this to consume data from c.dataChannels, and merged result should send to c.ch

	tmpDataList = make(map[int64]*versionIdData)
	versionId := int64(0)
	gap := int64(10)
	for {
		c.BatchProcess(500)
		for i := versionId; i < versionId+gap; i++ {
			//fmt.Println(i)
			if _, ok := tmpDataList[int64(i)]; ok {
				for _, val := range tmpDataList[int64(i)].data {
					c.ch <- &val
				}
			} else {
				break
			}

		}
		versionId = versionId+gap

	}
	}


// MergedChan returns a chan which can be used to receive merged versioned data
func (c *Consumer) MergedChan() <-chan *VersionedData {
	return c.ch
}


func (c *Consumer) BatchProcess(batchSize int) {
	for _,ch := range c.dataChannels {
		tmpData := <- ch
		tmpversionId := tmpData.VersionID
		for i := 0; i < batchSize; {
			if _,ok := tmpDataList[tmpData.VersionID];ok {

				tmpDataList[tmpData.VersionID].data = append(tmpDataList[tmpData.VersionID].data, *tmpData)

			} else {
				productid := []int64{}
				streamMap := map[int64][]int64{}
				versionId := tmpData.VersionID
				data := []VersionedData{*tmpData}
				retry := 0
				tmp := versionIdData{productid,streamMap, versionId,data,retry}
				tmpDataList[tmpData.VersionID] = &tmp

			}

			if tmpData.VersionID != tmpversionId {
				tmpDataList[tmpversionId].produceId = append(tmpDataList[tmpversionId].produceId, tmpData.ProducerID)
				tmpversionId = tmpData.VersionID

			}
			i++
			tmpData = <- ch


		}
	}
}

