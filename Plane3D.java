/*
 * Code Written By Ali Raza Bhangu
 * Date: 01-29-2023
*/

//Importing utilities
import java.util.*;

//Plane3D Class
public class Plane3D {

    //Initializing variables that will be needed
    private double a,b,c,d;

    //First Constructor, transforming 3 given points into a plane equation.
    public Plane3D(Point3D p1, Point3D p2, Point3D p3)
    {
        //Equation to make an vector
        Double[] vec1= {p2.getX() - p1.getX(), p2.getY() - p1.getY(), p2.getZ() - p1.getZ()};
        Double[] vec2 = {p3.getX() - p1.getX(), p3.getY() - p1.getY(), p3.getZ() - p1.getZ()};



        //Added this code here due to issues arising without it, originally attempted to utilize the secondary constructor, however this step was required.
        //Cross Product
        this.a = (vec1[1]*vec2[2]) - (vec2[1]*vec1[2]);
        this.b = (vec1[2]*vec2[0]) - (vec2[2]*vec1[0]);
        this.c = (vec1[0]*vec2[1]) - (vec2[0]*vec1[1]);
        this.d =((a*p1.getX()) + (b*p1.getY()) + (c* p1.getZ()))*-1;

        //Sending the newly obtained plane equation to the secondary constructor.
        new Plane3D((vec1[1]*vec2[2]) - (vec2[1]*vec1[2]),(vec1[2]*vec2[0]) - (vec2[2]*vec1[0]),(vec1[0]*vec2[1]) - (vec2[0]*vec1[1]),((a*p1.getX()) + (b*p1.getY()) + (c* p1.getZ()))*-1);
    }

    //Second Constructor that saves the coefficient values
    public Plane3D(double a, double b, double c, double d) {
        this.a = a;
        this.b = b;
        this.c = c;
        this.d = d;
    }

    //Method that gets the distance of the plane and a selected point. Uses a formula.
    public double getDistance(Point3D pt) {
        //Formula: |ax1 + by1 + cz1 + d| / sqrt(a^2+b^2+c^2)
        double distance = 0;

        //Splitting up the formula to make it less error-prone and more readable.
        double distFormulaTop = pt.getX()*this.a + pt.getY()*this.b + pt.getZ()*this.c + this.d;
        double distFormulaBot = Math.sqrt((a*a)+(b*b)+(c*c));


        //Calculating the final distance and returning it.
        distance = distFormulaTop/distFormulaBot;
        return Math.abs(distance);

    }


}

//EOF
