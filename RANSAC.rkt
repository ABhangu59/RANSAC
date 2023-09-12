#lang racket
;Ali Raza Bhangu
;Project 3 -> 2023-04-06


;;readXYZ function: Reads a selected file
;;readXYZ input parameters: recieves the filename of a specific file
;;readXYZ output: Reads the point cloud in a file and creates a list of 3D Points
(define (readXYZ fileIn)
  (let ((sL (map (lambda s (string-split (car s)))
                 (cdr (file->lines fileIn)))))
    (map (lambda (L)
           (map (lambda (s)
                  (if (eqv? (string->number s) #f)
                      s
                      (string->number s))) L)) sL)))


;;getPlane function: Get an input of 3 points, from there it determines the plane equation in ax+by+cz+d form
;;getPlane input parameters: Receives 3 points, which are lists containing (x,y,z)
;;getPlane output: Outputs a list of the co-efficents for the plane equation, (a,b,c,d)
(define (getPlane P1 P2 P3)
  ;Defining all variables required for finding co-efficents:
  (let* ((p1-x (car P1))
        (p1-y (cadr P1))
        (p1-z (caddr P1))
        (p2-x (car P2))
        (p2-y (cadr P2))
        (p2-z (caddr P2))
        (p3-x (car P3))
        (p3-y (cadr P3))
        (p3-z (caddr P3))
        (vec1-x (- p2-x p1-x))
        (vec1-y (- p2-y p1-y))
        (vec1-z (- p2-z p1-y))
        (vec2-x (- p3-x p1-x))
        (vec2-y (- p3-y p1-y))
        (vec2-z (- p3-z p1-z))
        ;Creating the co-efficents needed:
        (a (- (* vec1-y vec2-z) (* vec1-z vec2-y)))
        (b (* -1 [- (* vec1-z vec2-x) (* vec1-x vec2-z)]))
        (c (- (* vec1-x vec2-y) (* vec1-y vec2-x)))
        (d (* (+ (* a p1-x) (* b p1-y) (* c p1-z)) -1)))
    ;Combining the co-efficents in a list as output. 
    (list a b c d)))


;;getSupport function: Count the number of support points and the plane parameter in a pair, eps
;;getSupport input parameter: Takes a list of points and a plane equation.
;;getSupport output: Outputs a pair of the support count and the plane equation
(define (getSupport plane points eps)
  (supportHelper plane points eps 0)) ;Calling the supportHelper function to recursibely determine it. 

;;supportHelper function: Helper method to run through the recursion for getSupport
;;supportHelper input parameter: plane, pointcloud, eps, support-count.
;;supportHelper output: Returns the number of supporting points and the plane equation as a pair.
(define (supportHelper plane point-list eps support-count)
  (cond
    ([null? point-list] (cons support-count(list plane)))
    ([< (getDistance (car point-list) plane) eps] (supportHelper plane (cdr point-list) eps (+ support-count 1)))
    [else (supportHelper plane (cdr point-list) eps support-count)]))

;;getDistance function: Determine the distance of the point using a formula, Formula Used: |ax1 + by1 + cz1 + d| / sqrt(a^2+b^2+c^2)
;;getDistance input parameter: Takes a point and a plane equation
;;getDistance output: Returns the distance after calulating.
(define (getDistance point plane)
  (/ (abs (+ (* (car plane) (car point))
              (* (cadr plane) (cadr point))
              (* (caddr plane) (caddr point))
              (cadddr plane)))
     (sqrt (+ (expt (car plane) 2)
              (expt (cadr plane) 2)
              (expt (caddr plane) 2)))))

;;numOfIterations function: determine number of iterations using the formula provided in the PDF, rounded to ensure proper iterations.
;;numOfIterations input parameter: confidence, percentage values
;;numOfIterations output: Outputs the number of iterations
(define (numOfIterations confidence percentage)
  (round(/ (log (- 1 confidence)) (log (- 1 (* (* percentage percentage) percentage))))))

;;dominantPlane function: Returns the plane that has the best support.
;;dominantPlane input parameter: dominantPlane takes a list of points, a count variable (used to determine the indexing), epsilon, and a currentSupport variable.  
;;dominantPlane output: Outputs the plane with the best support. 
(define (dominantPlane Ps k eps currentSupport) 
  (if (= k 0) ;base case: if the 'k' (counter) is 0, we've iterated the entire list and will return the currentSupport 
      currentSupport 
      (let* ((temp-plane (getPlane (list-ref Ps (random (length Ps))) ;Creating a temporary plane with 3 random points (code provided from pdf)
                                    (list-ref Ps (random (length Ps)))
                                    (list-ref Ps (random (length Ps)))))
            (supportingPlane (getSupport temp-plane Ps eps))) ;creating a supportingPlane with the temp plane, basically a pair with the # of supporting points
        (if (> (car supportingPlane) (car currentSupport)) ;checking if the support of the plane is larger than the current support
            (dominantPlane Ps (- k 1) eps supportingPlane) 
            (dominantPlane Ps (- k 1) eps currentSupport)))))



;;planeRANSAC function: The 'engine' of the program, it will basically allow the program to be called and launch everything
;;planeRANSAC input parameters: Given a filename, confidence, percentage, and epsilon value which are used to run the entire program
;;planeRANSAC output: Returns the best and most dominant plane. 
(define (planeRANSAC filename confidence percentage eps)
  (define point-cloud (readXYZ filename)) ;creating a list of points from the file 
  (dominantPlane point-cloud (numOfIterations confidence percentage) eps (cons 0'(0)))) ;running the entire program. 

;End of File :)