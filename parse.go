package parseobj3d

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"github.com/alanmacleod/math3d"
)

type FaceIndex struct {
	I1 int
	I2 int
	I3 int
}


func Parse(file string) ([]math3d.Vector3, []math3d.Vector2, []FaceIndex, []FaceIndex, error) {

	lines, err := readLines(file)

	vertexPositions := []math3d.Vector3{}
	vertexUVs := []math3d.Vector2{}
	facePositions := []FaceIndex{}
	faceUVs := []FaceIndex{}

	for i, line := range lines {

		s := strings.Split(line, " ")

		switch s[0] {
		case "v":

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

			vertexPositions = append(vertexPositions, math3d.Vector3{x, y, z})

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
			// 1/2/3 4/5/6 7/8/9
			// ^ we want first two numbers in each group of ^

			v1 := strings.Split(s[1], "/")
			v2 := strings.Split(s[2], "/")
			v3 := strings.Split(s[3], "/")

			posIndex1, err := strconv.Atoi(v1[0])
			if err != nil {
				fmt.Println(i, "Error parsing Int")
			}

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

			facePositions = append(facePositions,FaceIndex{posIndex1, posIndex2, posIndex3})
			faceUVs = append(faceUVs,FaceIndex{uvIndex1, uvIndex2, uvIndex3})


		//case "g":
			//fmt.Println(i, "Group: ", s[1])

		}

	}

	//obj := Object3d{&vertex, &vertexCoord}

	return vertexPositions, vertexUVs, facePositions, faceUVs, err
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