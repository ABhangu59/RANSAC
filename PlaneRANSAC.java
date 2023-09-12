/*
 * Code Written By Ali Raza Bhangu
 * Date: 02-01-2023
 */



import java.awt.*;
import java.util.*;
import java.io.*;
import java.util.ArrayList;

public class PlaneRANSAC
{
    //Declaring objects/variables that will be used throughout.
    private double eps;
    private PointCloud pointCloud;

    //Constructor
    public PlaneRANSAC(PointCloud pc)
    {
        this.pointCloud = pc;
        this.eps = 0;
    }

    //Setter & Getter for Epsilon value.
    public void setEps(double eps)
    {
        this.eps = eps;
    }
    public double getEps()
    {
        return eps;
    }

    /***
     * Function to determine the numberOfIterations via confidence and percentageOfPointsOnPlane
     * @param confidence
     * @param percentageOfPointsOnPlane
     * @return calculated amount of iterations, through provided formula.
     */
    public int getNumberOfIterations(double confidence, double percentageOfPointsOnPlane)
    {
        double temp = Math.log(1-confidence)/Math.log(1-Math.pow(percentageOfPointsOnPlane, 3));
        return (int) Math.ceil(temp);
    }


    /***
     * Primary function that will execute the RANSAC Algorithm
     * @param numberOfIterations - number of iterations for the algorithm
     * @param filename - filename that will be parsed and have variations made.
     */
    public void run(int numberOfIterations, String filename)
    {
        //Starting off with initializing needed variables.
        int currentSupport = 0;

        //Initializing these outside of the loop, these are the 3 points that will be used to create the comparison plane.
        Point3D p1;
        Point3D p2;
        Point3D p3;
        Plane3D currentPlane; //Initializing the comparisonPlane outside the loop similar to the 3 points above.

        //Creating the dominantCloud.
        PointCloud dominantCloud = new PointCloud();


        //Initializing these outside of the for loop.
        Iterator<Point3D> iterator;
        PointCloud temporaryCloud;


        //For Loop -> Iterating through the number of iterations.
        for (int m = 0; m < numberOfIterations; m++)
        {
            int counter = 0;

            //Randomly drawing 3 points from the point cloud
            p1 = this.pointCloud.getPoint();
            p2 = this.pointCloud.getPoint();
            p3 = this.pointCloud.getPoint();

            //Formulating a plane with chose three points
            currentPlane = new Plane3D(p1,p2,p3);

            //Officially assigning values to the temporaryCloud and iterator that were created on line 72 and 71 respectively.
            temporaryCloud = new PointCloud();
            iterator = pointCloud.iterator();

            //While Loop - Utilizing the iterator method hasNext to keep the loop running.
            while (iterator.hasNext()) {

                //Creating a point that acts as a counter of sorts by being the current point that is iterated on.
                Point3D counterPoint = iterator.next();

                //Step 4 of the Algorithm, just comparing the distance of the plane and the point to the epsilon.
                if (currentPlane.getDistance(counterPoint) < getEps())
                {
                    //If the point is indeed larger, than the point is added.
                    temporaryCloud.addPoint(counterPoint);
                    counter++; //Incrementing counter.
                }
            }

            //Step 5 of the algorithm.
            if (temporaryCloud.getSize() > currentSupport)
            {
                //Setting currentSupport equal to the counter.
                currentSupport = counter;
                //Storing tempCloud into the dominant cloud
                dominantCloud = temporaryCloud;
            }


            //Getting rid of the .xyz and adding in the required things.
            String tempString = filename.substring(0, filename.length()-4);
            tempString = tempString+"_p"+(m+1)+".xyz";
            temporaryCloud.save(tempString);
        }
    }
}


//EOF
