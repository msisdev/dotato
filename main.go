package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Define the struct you want to marshal and unmarshal
type Person struct {
	Name string
	Age  int
}

func main() {
	// Create an instance of the struct
	p1 := Person{Name: "Alice", Age: 30}

	// Create a buffer to hold the encoded data
	var buffer bytes.Buffer

	// Create a new Gob encoder writing to the buffer
	enc := gob.NewEncoder(&buffer)

	// Marshal the struct into the buffer
	err := enc.Encode(p1)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// The buffer now contains the byte slice representing the struct
	byteSlice := buffer.Bytes()
	fmt.Printf("Marshaled byte slice: %v\n", byteSlice)

	// Create a new buffer with the marshaled data
	var readBuffer bytes.Buffer
	readBuffer.Write(byteSlice)

	// Create a new Gob decoder reading from the buffer
	dec := gob.NewDecoder(&readBuffer)

	// Create a new struct to hold the unmarshaled data
	var p2 Person

	// Unmarshal the byte slice into the struct
	err = dec.Decode(&p2)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	// Print the unmarshaled struct
	fmt.Printf("Unmarshaled struct: %+v\n", p2)
}
