package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"io"
	"strings"
)

func ReadPly(ply_filename string) func() (int, string) {
	ply_file, f_error := os.Open(ply_filename)
	reader := bufio.NewReader(ply_file)

	return func() (int, string) {
		if f_error != nil {
			return -1, ""
		}
		var status int
		var output string
		line, isPrefix, err := reader.ReadLine()
		switch {
		case err == io.EOF:
			status = 0
			output = ""
		case isPrefix:
			status = 1
			output = string(line)
			for isPrefix {
				line, isPrefix, err = reader.ReadLine()
				output = output + string(line)
			}
		default:
			status = 1
			output = string(line)
		}
		return status, output
	}
}

func main() {
	var n_vertex, n_faces,  misses int
	ply_filename := os.Args[1]
	ply_reader := ReadPly(ply_filename)

	buffer_size, _ := strconv.Atoi(os.Args[2])
	buffer := NewBuffer(buffer_size)

	for {
		status, line := ply_reader()
		if strings.HasPrefix(line, "element vertex") {
			n_vertex, _ = strconv.Atoi(strings.Split(line, " ")[2])
		} else if strings.HasPrefix(line, "element face") {
			n_faces, _ = strconv.Atoi(strings.Split(line, " ")[2])
		} else if strings.HasPrefix(line, "end_header") {
			break
		}
		if status == 0 {
			return
		}
	}

	// Reading the vertices
	for i := 0; i < n_vertex; i++ {
		status, _ := ply_reader()
		if status == 0 {
			return
		}
	}

	// Reading the faces.
	for i := 0; i < n_faces; i++ {
		status, line := ply_reader()
		if status == 0 {
			return
		}

		face := strings.Split(strings.TrimSpace(line), " ")[1:]
		for _, vs := range face {
			v, _ := strconv.Atoi(vs)
			if buffer.Push(v) {
				misses++
			}
		}
	}
	fmt.Printf("Number of vertex: %d\n", n_vertex)
	fmt.Printf("Number of faces: %d\n", n_faces)
	fmt.Printf("Number of misses: %d\n", misses)
	fmt.Printf("ACMR: %f\n", float64(misses) / float64(n_faces))
}
