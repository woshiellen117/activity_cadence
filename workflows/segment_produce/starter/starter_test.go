package main

import (
	"testing"
)

func BenchmarkStartWorkflowSegmentProduce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		err := StartWorkflowSegmentProduce()
		if err != nil {
			panic(err)
		}

		//oneMs,_ :=time.ParseDuration("100000000ns")
		//time.Sleep(oneMs)
	}
}
