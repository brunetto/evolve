NBabel.cpp
Date: Dec 2010
Author: Jeroen Bédorf (Sterrewacht Leiden, bedorf@strw.leidenuniv.nl)

Integration scheme: Predictor-corrector leapfrog
Compiler: gcc version 4.4.5 (Ubuntu/Linaro 4.4.4-14ubuntu5)
Operating system: Ubuntu Linux 10.10 (2.6.35-30-generic)
Hardware: Intel Core I7 CPU M640 2.80GHz

Input file: inputx (Plummer distribution of x equal mass particles)
Time step: constant and shared 1.e-3 N-body unit
Integration from t=0 to t=1.0
Performance:
N	Optimizaltion	tCPU	dE/E
16 	-  		0.02 	0.0082104
16 	O4 		0.0 	0.0082104
32 	- 		0.05 	7.33935e-07
32 	O4 		0.01	7.33935e-07
64 	- 		0.16 	3.59673e-06
64 	O4 		0.08 	3.59673e-06
128 	- 		0.64 	1.72732e-06
128	O4 		0.3 	1.72732e-06
256 	- 		2.58 	-0.000668772
256 	O4 		1.17      -0.000668772
512 	- 		10.25 	-0.0422243
512 	O4 		4.56      -0.0422243
1024 	- 		40.84 	-0.021317
1024 	O4 		18.7      -0.021317
2048 	- 		162.73    -0.00857556
2048 	O4 		74.83 	-0.00857556
4096 	- 		654.96    -0.00257437
4096 	O4 		295.77 	-0.00257437