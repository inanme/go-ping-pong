package rndom

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func printMe(x time.Weekday) {
	println(x)
}

func Test1(t *testing.T) {
	println(time.Now().Weekday())
	printMe(3)
	f := hex.Dumper(os.Stdout)
	fmt.Fprintln(f, "mert")
}

type Person struct {
	Name, Surname string
}

func Test_xml(t *testing.T) {
	person := Person{"mert", "inan"}
	p, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	println(string(p))
}

func Test_json(t *testing.T) {
	person := Person{"mert", "inan"}
	p, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	println(string(p))
}

//(1) The interface
type Mutable interface {
	mutate(newValue string) error
}

//(2) Struct
type Data struct {
	name string
}
type Data1 struct {
	name string
}

//(3) Implements the interface with a pointer receiver
func (d *Data) mutate(newValue string) error {
	d.name = newValue
	return nil
}
func (d Data1) mutate(newValue string) error {
	d.name = newValue
	return nil
}

//(4) Function that accepts the interface
func mutator(mute Mutable) error {
	switch mute.(type) {
	case *Data:
		return mute.mutate("mutate")
	case Data1:
		return mute.mutate("mutate1")
	default:
		return mute.mutate("unknown")
	}
}
func Test_interface(t *testing.T) {
	d, d1 := Data{"fresh"}, Data1{"fresh1"}
	fmt.Println(d.name) //fresh
	fmt.Println(d1.name)
	//(5) pass as a pointer
	mutator(&d)
	mutator(d1)
	fmt.Println(d.name) //mutate
	fmt.Println(d1.name)
}
