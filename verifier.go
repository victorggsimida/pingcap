package main

import (
	"context"
	"fmt"
)

// Verifier verify whether Consumer have merged versioned data
// NOTE: the implementation for Verifier is optional
type Verifier struct {
	ch <-chan *VersionedData
}

// NewVerifier creates a new Verifier
func NewVerifier(ch <-chan *VersionedData) *Verifier {
	v := &Verifier{
		ch: ch,
	}
	return v
}

// Start starts verifying
func (v *Verifier) Start(ctx context.Context) {
	// TODO: implement this to verify data from v.ch
	versionId := int64(0)
	for {
		data := <- v.ch
		if data.VersionID == versionId{
			continue
		} else if data.VersionID < versionId{
			fmt.Println("******verify is not pass********")
			fmt.Println(data)
		} else {
			versionId = data.VersionID
		}

	}
}
