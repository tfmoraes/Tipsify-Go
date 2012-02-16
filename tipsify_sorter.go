package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func BuildAdjacency(V [][3]float64, I [][3]int) (A [][]int, L []int) {
	for i := 0; i < len(V); i++ {
		A = append(A, make([]int, 0))
		L = append(L, 0)
	}

	for i := 0; i < len(I); i++ {
		A[I[i][0]] = append(A[I[i][0]], i)
		A[I[i][1]] = append(A[I[i][1]], i)
		A[I[i][2]] = append(A[I[i][2]], i)

		L[I[i][0]] += 1
		L[I[i][1]] += 1
		L[I[i][2]] += 1
	}
	return
}

func Skip_dead_end(L []int, D *[]int, I [][3]int, i int) (int) {
	D_ := *D
	for ; len(D_) > 0; {
		d := D_[len(D_) - 1]
		D_ = D_[0:len(D_)-1]
		if L[d] > 0 { 
			*D = D_
			return d
		}
	}
	*D = D_

	for ; i < len(L); {
		if L[i] > 0 { 
			return i
		}
		i += 1
	}
	return -1
}

func Get_next_vertex(I [][3]int, i int, k int, N map[int]bool, C []int, s int, L []int, D *[]int) (int) {
	n := -1
	p := -1
	m := 0
	for v, _ := range N {
		if L[v] > 0 {
			p = 0
			if s - C[v] + 2 * L[v] <= k {
				p = s-C[v]
			}
			if p > m {
				m = p
				n = v
			}
		}
	}
	if n == -1 {
		n = Skip_dead_end(L, D, I, i)
	}
	return n
}

func Tipsify(V [][3]float64, I [][3]int, k int) (O [][3]int) {
	A, L := BuildAdjacency(V, I)
	C := make([]int, len(V))
	D := make([]int, 0)
	E := make([]bool, len(I))
	f := 0
	s := k + 1
	i := 1
	nf := 0
	O = make([][3]int, len(I), len(I))
	for ; f>=0; {
		N := make(map[int]bool)
		for _, t := range A[f] {
			if !E[t] {
				for nv, v := range I[t] {
					O[nf][nv] = v
					D = append(D, v)
					N[v] = true
					L[v] = L[v] - 1
					if s - C[v] > k {
						C[v] = s
						s += 1
					}
				}
				E[t] = true
				nf++
			}
		}
		f = Get_next_vertex(I, i, k, N, C, s, L, &D)
	}
	return
}

func WritePly(ply_filename string, vertices [][3]float64, faces [][3]int) {
	ply_file, _ := os.Create(ply_filename)
	// The header
	ply_file.WriteString("ply\n")
	ply_file.WriteString("format ascii 1.0\n")
	ply_file.WriteString(fmt.Sprintf("element vertex %d\n", len(vertices)))
	ply_file.WriteString("property float x\n")
	ply_file.WriteString("property float y\n")
	ply_file.WriteString("property float z\n")
	ply_file.WriteString(fmt.Sprintf("element face %d\n", len(faces)))
	ply_file.WriteString("property list uchar int vertex_indices\n")
	ply_file.WriteString("end_header\n")

	for i:= 0; i < len(vertices); i++ {
		ply_file.WriteString(fmt.Sprintf("%f %f %f\n", vertices[i][0], vertices[i][1], vertices[i][2]))
	}
	for i:= 0; i < len(faces); i++ {
		ply_file.WriteString(fmt.Sprintf("3 %d %d %d\n", faces[i][0], faces[i][1], faces[i][2]))
	}

}

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
		case err == os.EOF:
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

	buffer_size, _ := strconv.Atoi(os.Args[3])
	ply_reader := ReadPly(os.Args[1])
	var n_vertex, n_faces int

	// Reading the header
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
	vertices := make([][3]float64, n_vertex, n_vertex)
	for i := 0; i < n_vertex; i++ {
		status, line := ply_reader()
		if status == 0 {
			return
		}

		vertex_s := strings.Split(line, " ")
		vertices[i][0], _ = strconv.Atof64(vertex_s[0])
		vertices[i][1], _ = strconv.Atof64(vertex_s[1])
		vertices[i][2], _ = strconv.Atof64(vertex_s[2])
	}

	// Reading the faces.
	faces := make([][3]int, n_faces, n_faces)
	for i := 0; i < n_faces; i++ {
		status, line := ply_reader()
		if status == 0 {
			return
		}

		faces_s := strings.Split(line, " ")
		faces[i][0], _ = strconv.Atoi(faces_s[1])
		faces[i][1], _ = strconv.Atoi(faces_s[2])
		faces[i][2], _ = strconv.Atoi(faces_s[3])
	}
	fmt.Println(n_vertex, n_faces)
	O_faces := Tipsify(vertices, faces, buffer_size)
	WritePly(os.Args[2], vertices, O_faces)
}
