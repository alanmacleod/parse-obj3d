package parseobj3d

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"github.com/alanmacleod/math3d"
	"errors"
)

var vertexPositions [][]float64
var vertexUVs  		[]math3d.Vector2
var facePositions   [][]int
var faceUVs 		[]math3d.Triangle

func Parse(file string) ([][]float64, []math3d.Vector2, [][]int, []math3d.Triangle, error) {

	lines, err := readLines(file)

	for i, line := range lines {

		s := strings.Split(line, " ")

		switch s[0] {

		case "v":

			// Read Vertices' XYZ positions

			x, err := strconv.ParseFloat(s[1], 64)
			if err != nil {
				fmt.Println(i, "Bad float X = ", s[1])
			}
			y, err := strconv.ParseFloat(s[2], 64)
			if err != nil {
				fmt.Println(i, "Bad float Y = ", s[2])
			}
			z, err := strconv.ParseFloat(s[3], 64)
			if err != nil {
				fmt.Println(i, "Bad float Z = ", s[3])
			}

			thisVertex := make([]float64, 3)

			thisVertex[0] = x
			thisVertex[1] = y
			thisVertex[2] = z

			vertexPositions = append(vertexPositions, thisVertex)

			//vertexPositions = append(vertexPositions, math3d.Vertex{&math3d.Vector3{x,y,z},0,0})

		case "vt":

			u, err := strconv.ParseFloat(s[1], 64)
			if err != nil {
				fmt.Println(i, "Bad float U = ", s[1])
			}
			v, err := strconv.ParseFloat(s[2], 64)
			if err != nil {
				fmt.Println(i, "Bad float V = ", s[2])
			}

			vertexUVs = append(vertexUVs, math3d.Vector2{u,v})


		case "f":

			// Read face (triangle) indices (index into the vertices list above)

			// Format: 1/2/3 4/5/6 7/8/9
			// 		   ^ we want first two numbers in each group of ^

			v1 := strings.Split(s[1], "/")
			v2 := strings.Split(s[2], "/")
			v3 := strings.Split(s[3], "/")

			numErrs := 0

			posIndex1, err := strconv.Atoi(v1[0])
			if err != nil { numErrs++ }

			posIndex2, err := strconv.Atoi(v2[0])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

			posIndex3, err := strconv.Atoi(v3[0])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

			uvIndex1, err := strconv.Atoi(v1[1])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

			uvIndex2, err := strconv.Atoi(v2[1])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

			uvIndex3, err := strconv.Atoi(v3[1])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

			uvIndex1 = uvIndex1
			uvIndex2 = uvIndex2
			uvIndex3 = uvIndex3

			thisFace := make([]int, 3)

			thisFace[0] = posIndex1-1
			thisFace[1] = posIndex2-1
			thisFace[2] = posIndex3-1

			facePositions = append(facePositions, thisFace)

			/*
			newFace := math3d.Triangle{}

			fv0, err := getFaceVertex(posIndex1-1)
			if err != nil {
				fmt.Println(err)
			}

			fv1, err := getFaceVertex(posIndex2-1)
			if err != nil {
				fmt.Println(err)
			}

			fv2, err := getFaceVertex(posIndex3-1)
			if err != nil {
				fmt.Println(err)
			}

			//newFace.SetVertices(fv0, fv1, fv2)
			fv0 = fv0
			fv1 = fv1
			fv2 = fv2


			//Triangles = append(Triangles, newFace)
			//facePositions = append(facePositions,FaceIndex{posIndex1, posIndex2, posIndex3})
			facePositions = append(facePositions, newFace)
			//faceUVs = append(faceUVs,FaceIndex{uvIndex1, uvIndex2, uvIndex3})
			*/

		}

	}

	return vertexPositions, vertexUVs, facePositions, faceUVs, err
}


func getFaceVertex(index int) (*[]float64, error) {
	if index  >= len(vertexPositions) {
		return nil, errors.New(fmt.Sprintf("Invalid vertex index! -> %d", index))
	}

	return &vertexPositions[index], nil
}


func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()

}