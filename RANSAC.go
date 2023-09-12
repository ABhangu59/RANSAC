/*
Ali Raza Bhangu
300234254
2023-03-11
*/

package main

//Importing required utilities
import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Creating the needed objects
type Point3D struct {
	X float64
	Y float64
	Z float64
}
type Plane3D struct {
	A float64
	B float64
	C float64
	D float64
}
type Plane3DwSupport struct {
	Plane3D
	SupportSize int
}

// ReadXYZ Function -> Extracts XYZ from the specified file
func ReadXYZ(filename string) []Point3D {
	//Making the needed variables
	listOfPoints := []Point3D{}
	readFile, err := os.Open(filename)

	//Error Handling for the filename
	if err != nil {
		log.Fatal(err)
	}

	//Defering the closing of the file
	defer readFile.Close()

	//New Scanner, using bufio to parse the file.
	scanner := bufio.NewScanner(readFile)
	//For Loop that runs until Scanner.Scan is false.
	for scanner.Scan() {
		line := scanner.Text()
		indexPoints := strings.Split(line, "	")
		//Skipping the first line so that it only reads the actual information.
		//Using .Index function, it returns -1 if the substring is not found inside
		if strings.Index(indexPoints[0], "x") == -1 {
			//Assigning values to x,y,z respectively
			xVal, _ := strconv.ParseFloat(indexPoints[0], 64)
			yVal, _ := strconv.ParseFloat(indexPoints[1], 64)
			zVal, _ := strconv.ParseFloat(indexPoints[2], 64)
			//Appending the new point to the list.
			listOfPoints = append(listOfPoints, Point3D{xVal, yVal, zVal})
		}
	}

	return listOfPoints
}

// SaveXYZ Function - Writes the information into a new file
func SaveXYZ(filename string, points []Point3D) {
	//Creating the file with os package.
	file, err := os.Create(filename)

	//Nil Check with Err.
	if err != nil {
		fmt.Println(err)
	}

	//Deferring the close of the file.
	defer file.Close()

	//Writing the file header
	file.WriteString("X, Y, Z\n")
	//Looping thru all the points and writing them in the specific format
	for index := range points {
		file.WriteString(fmt.Sprintf("%g", points[index].X) + ", " + fmt.Sprintf("%g", points[index].Y) + ", " + fmt.Sprintf("%g", points[index].Z))
		file.WriteString("\n") //for spacing
	}

}

// GetDistance Function - Calculates the distance of p1 and p2
func (p1 *Point3D) GetDistance(p2 *Point3D) float64 {
	//Formula: √((x2 - x1)^2 + (y2 - y1)^2 + (z2 - z1)^2)
	//Using Math.Sqrt to Square Root and Math.Pow to square the differences.
	return math.Sqrt(math.Pow((p2.X-p1.X), 2) + math.Pow((p2.Y-p1.Y), 2) + math.Pow((p2.Z-p1.Z), 2))
}

// GetNumberOfIterations function - similar to that seen in Part 1 of the Project
func GetNumberOfIterations(confidence float64, percentageOfPointsOnPlane float64) int {
	num := math.Log(1-confidence) / math.Log(1-math.Pow(percentageOfPointsOnPlane, 3))
	return int(num)
}

// Calculating the plane equation of the 3 points sent. Similar to the one used in Part 1 of the Project
func GetPlane(points [3]Point3D) Plane3D {
	//Seperating the 3 Points from the collection
	point1 := points[0]
	point2 := points[1]
	point3 := points[2]

	//Equation to make an vector, same as in Part 1
	vec1 := Point3D{point2.X - point1.X, point2.Y - point1.Y, point2.Z - point1.Z}
	vec2 := Point3D{point3.X - point1.X, point3.Y - point1.Y, point3.Z - point1.Z}

	//Cross Product - Split Up for Ease of Reading
	constantA := (vec1.Y * vec2.Z) - (vec1.Z * vec2.Y)
	constantB := (vec1.Z * vec2.X) - (vec2.Z * vec1.X)
	constantC := (vec1.X * vec2.Y) - (vec2.X * vec1.Y)

	//The D Value
	constantD := -1 * ((constantA * point1.X) + (constantB * point1.Y) + (constantC * point1.Z))

	//Returning the Plane with its constants.
	return Plane3D{A: constantA, B: constantB, C: constantC, D: constantD}
}

func GetSupport(plane Plane3D, points []Point3D, eps float64) Plane3DwSupport {
	//This will count the size of the Plane3DwSupport
	var supportCounter int

	for i := 0; i < len(points); i++ {
		//Using formula for Distance Between A Point and Plane
		topOfDenom := math.Abs(plane.A*points[i].X + plane.B*points[i].Y + plane.C*points[i].Z + plane.D)
		botOfDenom := math.Sqrt(math.Pow(plane.A, 2) + math.Pow(plane.B, 2) + math.Pow(plane.C, 2))

		distance := topOfDenom / botOfDenom

		//Comparing the plane point to the epislon. if less than epsilon, increment the supportCounter
		if distance < eps {
			supportCounter++
		}
	}

	//Returning Plane3DwSupport
	planeWithSupport := Plane3DwSupport{Plane3D: plane, SupportSize: supportCounter}
	return planeWithSupport
}

// This extracts the points that support thegiven plane and returns them as a slice of points
// Reused some code from GetSupport as it has overlap in functionality.
func GetSupportingPoints(plane Plane3D, points []Point3D, eps float64) []Point3D {

	//Making a slice that will be used to determine the supportSize for the Plane3DwSupport
	supportPoints := make([]Point3D, 0)

	for i := 0; i < len(points); i++ {
		//Using formula for Distance Between A Point and Plane
		topOfDenom := math.Abs(plane.A*points[i].X + plane.B*points[i].Y + plane.C*points[i].Z + plane.D)
		botOfDenom := math.Sqrt(math.Pow(plane.A, 2) + math.Pow(plane.B, 2) + math.Pow(plane.C, 2))

		distance := topOfDenom / botOfDenom

		//Same Comparison as in the function above, checking to see the distance.
		if distance < eps {
			//Appending the points to SupportPoints
			supportPoints = append(supportPoints, points[i])
		}
	}
	return supportPoints
}

// RemovePlane Function: Removes the points on the plane that support the plane
func RemovePlane(plane Plane3D, points []Point3D, eps float64) []Point3D {
	//Variable that will returned, hence named retSlice
	var retSlice []Point3D
	//Determining what supports the plane
	onPlanePoints := GetSupportingPoints(plane, points, eps)

	//For Loop -> Will run through all the points
	for i := 0; i < len(points); i++ {
		//This variable will be used to determine what supports the plane
		var exists bool
		//Nested For Loop -> Checks if
		for j := 0; j < len(onPlanePoints); j++ {
			//If Statement to determine if point supports
			if points[i] == onPlanePoints[j] {
				//If Supports, set exists to true and break this mini loop
				exists = true
				//Breaking the mini loop once exists is true so that the other part can be ran
				break
			}
		}

		//If Statement to basically add the points that do not support the plane to the new returned slice
		if !exists {
			retSlice = append(retSlice, points[i])
		}
	}
	return retSlice

}

/*
Start of Pipeline:
Basically it's set so each function will be provided with two channels, 1 as an input channel and 1 as a stop.
first method (normal parameters, start channel, stop)
second method (normal parameters, output channel of first as input, new output method, stop)
n method (normal parameters, n-1 output channel as input, n output as output, stop)
*/

// randomPointGen -> Piece 1 of the Pipeline.
// This function receieves points, and only two channels since its the first method.
func randomPointGen(points []Point3D, inputRPG chan Point3D, stopRPG chan bool) {
	//Flag Method for For Loop
	flag := true
	//For Loop that runs until Flag is False
	for flag {
		//Selecting an individual random point
		point := points[rand.Intn(len(points))]
		//Select Statement for different cases
		select {
		//Case 1: Stop The Program
		case <-stopRPG:
			flag = false
			break
		//Case 2: Keep Making Points
		case inputRPG <- point:
		}
	}

	return
}

// triplePointGenerator -> Piece 2 of the Pipeline
func triplePointGenerator(inputTPG chan Point3D, outputTPG chan [3]Point3D, stopTPG chan bool) {
	//Making a size of 3 list, as only 3 points are being generated.
	points := make([]Point3D, 3)
	//Counter called i
	i := 0
	for {
		//Select Statement for different outcomes
		select {
		case point := <-inputTPG:
			points[i] = point
			if i == 2 {
				outputTPG <- [3]Point3D{points[0], points[1], points[2]}
				i = 0
			} else {
				i++
			}
		case <-stopTPG:
			return
		}
	}
}

// takeN -> Piece 3 of the pipeline
func takeN(inputTKN chan [3]Point3D, outputTKN chan [3]Point3D, numOfArrays int, stopTKN chan bool) {
	count := 0                 // Counter to track number of arrays taken
	var placeHolder [3]Point3D //Placeholder array to hold the incoming
	for {
		select {
		case <-stopTKN: //If stop signal is received, end this function
			return
		case placeHolder = <-inputTKN: //input the incoming array to the placeholder
			if count != numOfArrays { //if not right number of arrays have been taken yet
				outputTKN <- placeHolder // send the array to the output channel
				count++                  //increment count
			}
			if count == numOfArrays { //If we have the right number of count, then stop.
				stopTKN <- true
			}
		}
	}
}

// planeEstimator -> Piece 4 of the pipeline
func planeEstimator(inputPE chan [3]Point3D, outputPE chan Plane3D, stopPE chan bool) {
	var placeHolderPoints [3]Point3D
	for {
		select {
		case placeHolderPoints = <-inputPE: //Sending the input to points
			outputPE <- GetPlane(placeHolderPoints)
		case <-stopPE: //If stop signal is sent, stop.
			return
		}
	}
}

// supportingPointFinder -> Piece 5 of the pipeline
func supportingPointFinder(inputSPF chan Plane3D, outputSPF chan Plane3DwSupport, supportingPoints []Point3D, eps float64, stopSPF chan bool) {
	var placeHolderPlane Plane3D //placeholder so I can transfer the incoming info to output

	//Continously listen for Plane3D values from inputSPF and send the results to outputSPF
	for {
		select {
		case placeHolderPlane = <-inputSPF:
			//Using getSupport to find the Plane3DWithSupport and send that to the output channel
			supportPlane := GetSupport(placeHolderPlane, supportingPoints, eps)
			outputSPF <- supportPlane
		//If a stop signal is recieved end the function.
		case <-stopSPF:
			return
		}
	}
}

// FanIn -> Piece 6 of the pipeline
func fanIn(inputFI chan Plane3DwSupport, outputFI chan Plane3DwSupport, stopFI chan bool) {
	for {
		select {
		case placeHolderPlanes := <-inputFI: //transferring the info from the input to the output
			outputFI <- placeHolderPlanes
		case <-stopFI: //this will stop the function if the need arises.
			return
		}
	}
}

// DominantPlanes -> Last piece of the pipeline
func dominantPlaneIdentifier(mainPlane *Plane3DwSupport, inputDPI chan Plane3DwSupport, stopDPI chan bool) {
	for {
		select {
		case comparisonPlane := <-inputDPI:
			//Checking the most dominant plane
			if mainPlane.SupportSize < comparisonPlane.SupportSize {
				*mainPlane = comparisonPlane
			}

		//This is the stop case, if the channel is true it'll break the loop and return.
		case <-stopDPI:
			return
		}
	}
}

// Execution Function -> this runs the entire pipeline, this is like the engine of the program
func execution(points []Point3D, iterations int, eps float64, threads int) Plane3D {

	//This is the plane that will be altered after the pipelines ran
	bestPlane := Plane3DwSupport{Plane3D{0, 0, 0, 0}, 0}
	bestPointer := &bestPlane

	//Initalizing the output channels
	RPGoutput := make(chan Point3D, 300)
	TPGoutput := make(chan [3]Point3D, 300)
	TKNoutput := make(chan [3]Point3D, 300)
	PEoutput := make(chan Plane3D, 300)
	SPFoutput := make(chan Plane3DwSupport, 300)
	FIoutput := make(chan Plane3DwSupport, 300)

	// Initialize the stop channels
	RPGstop := make(chan bool)
	TPGstop := make(chan bool)
	TKNstop := make(chan bool)
	PEstop := make(chan bool)
	SPFstop := make([]chan bool, threads)
	FIstop := make(chan bool)
	DPIstop := make(chan bool)

	// Start the pipelines with goroutines
	go func() {
		defer close(RPGoutput)
		randomPointGen(points, RPGoutput, RPGstop)
	}()
	go func() {
		defer close(TPGoutput)
		triplePointGenerator(RPGoutput, TPGoutput, TPGstop)
	}()
	go func() {
		defer close(TKNoutput)
		takeN(TPGoutput, TKNoutput, 3, TKNstop)
	}()
	go func() {
		defer close(PEoutput)
		planeEstimator(TKNoutput, PEoutput, PEstop)
	}()

	for i := 0; i < threads; i++ {
		SPFstop[i] = make(chan bool)
		go supportingPointFinder(PEoutput, SPFoutput, points, eps, SPFstop[i])
	}

	go func() {
		defer close(FIoutput)
		fanIn(SPFoutput, FIoutput, FIstop)
	}()
	go func() {
		dominantPlaneIdentifier(bestPointer, FIoutput, DPIstop)
	}()

	//Stopping the routine.
	<-TKNstop
	RPGstop <- true
	TPGstop <- true
	PEstop <- true

	//For Loop to accurately stop.
	for c := 0; c < threads; c++ {
		SPFstop[c] <- true
	}
	FIstop <- true
	DPIstop <- true

	//Returning the Plane
	return bestPlane.Plane3D

}

/*
Main Function For Program:
args[1] - filename
args[2] - confidence
args[3] - percentage
args[4] - epsilon

os.Args for command line reading.
*/
func main() {
	//Start of Runtime Clock
	start := time.Now()

	//Initalizing the information provided in the command line during run.
	points := ReadXYZ(os.Args[1])
	confidence, _ := strconv.ParseFloat(os.Args[2], 64)
	percentage, _ := strconv.ParseFloat(os.Args[3], 64)
	eps, _ := strconv.ParseFloat(os.Args[4], 64)

	//Number of threads for testing
	threads := 2

	//Finding number of iterations needed.
	totIterations := GetNumberOfIterations(confidence, percentage)

	for i := 0; i < 3; i++ {
		pipeLine := execution(points, totIterations, eps, threads)

		//This is to make sure the fileName will save for any of the pointCloud files.
		tempString := strings.TrimSuffix(os.Args[1], ".xyz")
		newFileName := tempString + "_p" + fmt.Sprintf("%d", i) + ".xyz"

		//Saving the file that will have the result of the pipeline.
		SaveXYZ(newFileName, GetSupportingPoints(pipeLine, points, eps))

		removedPoints := RemovePlane(pipeLine, points, eps)
		//Saving the file with the removed plane
		SaveXYZ("test"+"_p"+fmt.Sprintf("%d", i)+".xyz", removedPoints)
	}

	//This is here for the testing.
	currentTime := time.Now()

	timePassed := currentTime.Sub(start)

	fmt.Println(timePassed)
}

//End of File
