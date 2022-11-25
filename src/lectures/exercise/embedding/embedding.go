//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type DashboardComponents struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

type BandwidthUsageAverage interface {
	CalAverageBandwidth() Bytes
}
type MemoryUsageAverage interface {
	CalAverageMemory() Bytes
}
type CpuUsageAverage interface {
	CalAverageCpu() Celcius
}

func (b *BandwidthUsage) CalAverageBandwidth() Bytes {
	var total Bytes = 0
	for _, eachByte := range b.amount {
		total += eachByte
	}
	return total / 5
}
func (m *MemoryUsage) CalAverageMemory() Bytes {
	var total Bytes = 0
	for _, eachByte := range m.amount {
		total += eachByte
	}
	return total / 5
}
func (t *CpuTemp) CalAverageCpu() Celcius {
	var total Celcius = 0
	for _, eachCelc := range t.temp {
		total += eachCelc
	}
	return total / 5
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboardComponents := DashboardComponents{bandwidth, temp, memory}
	fmt.Printf("Average Bandwidth Usage is %v\n", dashboardComponents.CalAverageBandwidth())
	fmt.Printf("Average Memory Usage is %v\n", dashboardComponents.CalAverageMemory())
	fmt.Printf("Average Cpu Temp is %v\n", dashboardComponents.CalAverageCpu())
}
