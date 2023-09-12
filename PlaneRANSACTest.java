/*
 * Code Written By Ali Raza Bhangu
 * Date: 02-04-2023
 */

/***
 * In the event you (the user) have not read the README.txt, to run this class please use: java PlaneRANSACTest [filename] [number of iterations]
 */


import java.util.Scanner;

//Class That Acts as The Main/Testing Class
public class PlaneRANSACTest
{
    public static void main(String[] args) {
        //Creating a scanner object
        Scanner scanner = new Scanner(System.in);

        //Asking for the epsilon value, will store in double.
        System.out.println("Please enter the epsilon value: ");
        double epsilon = scanner.nextDouble();

        //Making a PointCloud that will be sent to PlaneRANSAC
        PointCloud testCloud = new PointCloud(args[0]);

        //Creating a RANSAC object to test the run function.
        PlaneRANSAC tester = new PlaneRANSAC(testCloud);

        //Setting epsilon
        tester.setEps(epsilon);
        //Calling run with the numberOfIterations (args[1]) and the filename (args[0]).
        tester.run(Integer.parseInt(args[1]), args[0]);
    }
}

//EOF
