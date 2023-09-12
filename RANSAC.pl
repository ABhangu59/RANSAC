/*
*Ali Raza Bhangu
*Prolog RANSAC
*2023-04-08
*/


/*
Predicate that reads the point cloud in a file and creates a list of 3D points.
*/
read_xyz_file(File, Points) :-
 open(File, read, Stream),
read_xyz_points(Stream,Points),
 close(Stream).

read_xyz_points(Stream, []) :-
 at_end_of_stream(Stream).

read_xyz_points(Stream, [Point|Points]) :-
 \+ at_end_of_stream(Stream),
read_line_to_string(Stream,L), split_string(L, "\t", "\s\t\n",
XYZ), convert_to_float(XYZ,Point),
 read_xyz_points(Stream, Points).
convert_to_float([],[]).

convert_to_float([H|T],[HH|TT]) :-
 atom_number(H, HH),
 convert_to_float(T,TT).

/*
random3points/2 predicate: Given a list of points, determines a triplet.
*/

random3points(Points, Point3) :-
    [P1,P2,P3] = Point3,
    [X1,Y1,Z1] = P1,
    [X2,Y2,Z2] = P2,
    [X3,Y3,Z3] = P3,
    % Using random member to get a random member from points
    random_member([X1,Y1,Z1], Points),
    random_member([X2,Y2,Z2], Points),
    random_member([X3,Y3,Z3], Points).
/*
plane/2 predicate: Given Point3, should be true if the plane is determined.
Finding the plane using this formula: ax+by+cz=d.
*/

plane(Point3, Plane) :-
    [A,B,C,D] = Plane,
    [P1,P2,P3] = Point3,
    [X1,Y1,Z1] = P1,
    [X2,Y2,Z2] = P2,
    [X3,Y3,Z3] = P3,

    % Creating the vectors V1 and V2:
    V1X is X2-X1,
    V1Y is Y2-Y1,
    V1Z is Z2-Z1,
    V2X is X3-X1,
    V2Y is Y3-Y1,
    V2Z is Z3-Z1,

    % Calculating co-efficents for the plane EQN
    ANew is (V1Y * V2Z)-(V1Z * V2Y),
    BNew is (V1Z * V2X)-(V2Z*V1Z),
    CNew is (V1X*V2Y)-(V2X*V1Y),
    DNew is (ANew*X1)+(BNew*Y1)+(CNew*Z1) * -1,

    % Formulating plane EQN:
    Plane = [ANew, BNew, CNew, DNew].

/*
support/4 predicate:
*/
support(Plane, Points, Eps, N) :-
    support_helper(Plane, Points, Eps, N, 0).

/*
Helper Functions for Support:
These predicates will help with recursion for support and thus increase readability. 
*/
support_helper(_, [], _, N, N).

support_helper(Plane, [H|T], Eps, OriginalN, Counter) :-
    % Checking if distance < eps and then recursively running through the entire point cloud
    distanceChecker(H, Plane, Eps),
    NewCount is Counter + 1,
    support_helper(Plane, T, Eps, OriginalN, NewCount).

support_helper(Plane, [_|T],Eps, OriginalN, Counter) :-
    support_helper(Plane, T, Eps, OriginalN, Counter).

/*
Helper Function #1 for Support:
distance/3 predicate: Calculates the distance between a point and plane using the forumula.
*/
distance(Point, Plane, Dist) :-
    % Simple distance function
    [X,Y,Z] = Point,
    [A,B,C,D] = Plane,
    Dist is abs(A*X + B*Y + C*Z + D) / sqrt(A*A + B*B + C*C).

/*
Helper Function #2 for Support:
distanceChecker predicate:Will return true if the point lies on the plane within eps distance
*/
distanceChecker(Point, Plane, Eps) :-
    % Function to make reading of code neater.
    distance(Point, Plane, Dist),
    Dist < Eps.

/*
ransac-number-of-iterations/3 predicate: compares N to the desired numbers of iterations.
*/
ransac-number-of-iterations(Confidence, Percentage, N) :-
    K is log(1 - Confidence) / log(1 - Percentage**3),
    N is round(K).


/*
Test Cases - 1 Per Text File Provided (Excluding test cases for number of iterations):
*/
% Random Points Test Cases:

testR3P(random3points, 1, TripSet) :- 
    read_xyz_file('Point_Cloud_1_No_Road_Reduced.xyz', Points),
	random3points(Points, TripSet).

testR3P(random3points, 2, TripSet) :- 
    read_xyz_file('Point_Cloud_2_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet).


testR3P(random3points, 3, TripSet) :-
    read_xyz_file('Point_Cloud_3_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet).

% Plane Test Cases:
testPL(plane, 1, EQN) :- read_xyz_file('Point_Cloud_1_No_Road_Reduced.xyz', Points),
						  	random3points(Points, TripSet),
						  	plane(TripSet, EQN).

testPL(plane, 2, EQN) :-
    read_xyz_file('Point_Cloud_2_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet),
    plane(TripSet, EQN).

testPL(plane, 3, EQN) :-
    read_xyz_file('Point_Cloud_3_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet),
    plane(TripSet, EQN).


% Support Test Cases:
testSUP(support, 1, N) :-
    read_xyz_file('Point_Cloud_1_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet),
    plane(TripSet, EQN),
    support(EQN, Points, 0.5, N).
    
testSUP(support, 2, N) :-
    read_xyz_file('Point_Cloud_2_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet),
    plane(TripSet, EQN),
    support(EQN, Points, 0.5, N).

testSUP(support, 3, N) :-
    read_xyz_file('Point_Cloud_3_No_Road_Reduced.xyz', Points),
    random3points(Points, TripSet),
    plane(TripSet, EQN),
    support(EQN, Points, 0.5, N).


% Number of Iterations Test Cases:
testNUM(ransac-number-of-iterations, 1, t) :-
    ransac-number-of-iterations(0.5, 0.5, t).

testNUM(ransac-number-of-iterations, 2, t) :-
    ransac-number-of-iterations(0.25, 0.65, t).

testNUM(ransac-number-of-iterations, 3, t) :-
    ransac-number-of-iterations(0.99, 0.73, t).

% End of File