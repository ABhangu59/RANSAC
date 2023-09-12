/*
 * Code Written By Ali Raza Bhangu
 * Date: 01-29-2023
 */

import java.awt.*;
import java.util.*;
import java.io.*;

//PointCloud Class
public class PointCloud
{

    //Creating an ArrayList of Point3Ds.
    private ArrayList<Point3D> pointCloud;

    /***
     * Constructor that reads the file and parses the needed info.
     * @param filename A filename to be entered
     */
    PointCloud(String filename)
    {
        this.pointCloud = new ArrayList<Point3D>();

        //Try and Catch for the read function
        try
        {
            //Creating a new scanner object, with the filename specified in the method call.
            Scanner scanner = new Scanner(new File(filename));

            //Skipping past X,Y,Z on the DB
            scanner.nextLine(); //Skipping a line.

            //While the scanner has a next line, the csv file will be split at commas and the respective x,y,z values will be sent to a new point.
            while (scanner.hasNext())
            {
                String s = scanner.next();

                double tempX = Double.parseDouble(s);
                double tempY = Double.parseDouble(scanner.next());
                double tempZ = Double.parseDouble(scanner.next());

                pointCloud.add(new Point3D(tempX,tempY,tempZ));
            }
            scanner.close(); //After the while loop, closes the scanner object.
        }
        catch (FileNotFoundException e)
        {
            System.out.println("File Not Found");
        }
    }

    //Second Constructor: Instantiates the ArrayList
    PointCloud()
    {
        this.pointCloud = new ArrayList<Point3D>();
    }

    /***
     *  Function to add a submitted point to the existing cloud.
     * @param pt Requires a point of form Point3D.
     */
    public void addPoint(Point3D pt)
    {
        pointCloud.add(pt);
    }

    /***
     * Function to return a random Point within the PointCloud
     * @return A random Point3D Object
     */
    public Point3D getPoint()
    {
        //Creating a new random object that will be used to call random methods.
        Random r = new Random();

        //Getting a random integer within the size of the pointCloud
        int randomInt = r.nextInt(pointCloud.size());


        //returning a random point.
        return pointCloud.get(randomInt);
    }


    /***
     * Function to save the selected information as a new file of .xyz
     * @param filename User entered filename
     */
    public void save(String filename)
    {//Creating a new file object
        File file = new File(filename);
        FileWriter fr = null; //Creating a new file-writer object, setting as null for now

        try
        {
            fr = new FileWriter(file);
            //Starting off with writing the first line which specifies the positioning of the information
            fr.write("x,y,z\n");

            //For Loop - Runs through the length of the list and just writes the information of the different points.
            for(int m = 0; m < pointCloud.size(); m++)
            {
                fr.write(pointCloud.get(m).getX() + "\t" + pointCloud.get(m).getY() + "\t" +pointCloud.get(m).getZ()+"\n");
            }
        }
        //Catches an IOException in the event it happens - this code was needed for it to run
        catch (IOException e)
        {
            System.out.println("Error");
        }
    }

    //GetSize method, returns the size of the point. Added this to make it easier to get the size in PlaneRANSAC.java
    public int getSize()
    {
        return pointCloud.size();
    }


    //Iterator Function, formatted so that the iterator returns an instance of an iterator of point3D.
    public Iterator<Point3D> iterator()
    {
        return new Iterator<Point3D>() {
            //Setting the index to 0
            int index = 0;

            //Has Next method that just determines if there is another element.
            @Override
            public boolean hasNext() {
                return index < pointCloud.size();
            }
            //Next method that continues to grab a new point and increment the index.
            @Override
            public Point3D next() {
                return pointCloud.get(index++);
            }

            //Remove method that will remove the said index, proceeds to decrease index after removal to ensure there is no bound error.
            public void remove()
            {
                pointCloud.remove(index);
                index= index - 1;
            }
        };
    }
}

//EOF


//Ali Bhangu :)

