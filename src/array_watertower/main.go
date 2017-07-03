package main 

import (
    "flag"
    "fmt"
    //"os"
    //"strings"
    //"strconv"
)

const APP_VERSION = "0.1"


// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")



func main() {
	//var err error
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    // aWater :=[]int{5,1,5}     
    //aWater :=[]int{5,3,7,2,6,4,5,9,1,2}  
    //aWater :=[]int{5,3,7,2,6,4,5,6,1,2} 
    //aWater :=[]int{10,9,8,7,6,5,6,7,8,7}  
    //aWater :=[]int{10,9,8,7,6,5,6,7,8,7,6,5,6,7,3,5} 
      aWater :=[]int{5,4,3,2,1,2,1,2}       
                              
     
        
    /*
    9                                |-|
    8                                | |
    7             |-|  x   x   x  x  | |
    6             | |  x  |-|  x  x  | |
    5     | |  x  | |  x  | |  x |-| | |
    4     | |  x  | |  x  | | |-|| | | |
    3     |-| |-| | |  x  | | | || | | |
    2     | | | | | | |-| | | | || | | |  x  |-|
    1     | | | | | | | | | | | || | | | |-| | |
           5   3   7   2   6   4  5   9   1   2
    
    Answer = 14
    
    10    |-|
    9     | | |-|                               
    8     | | | | |-|  x   x   x   x   x  |-| 
    7     | | | | | | |-|  x   x   x  |-| | | |-|  x   x   x  |-|
    6     | | | | | | | | |-|  x  |-| | | | | | | |-|  x  |-| | |
    5     | | | | | | | | | | |-| | | | | | | | | | | |-| | | | |  x  |-|
    4     | | | | | | | | | | | | | | | | | | | | | | | | | | | |  x  | |
    3     | | | | | | | | | | | | | | | | | | | | | | | | | | | | |-| | |
    2     | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |
    1     | | | | | | | | | | | | | | | | |-| | | | | | | | | | | | | | |
          10   9   8   7   6   5   6   7   8   7   6   5   6   7   3   5
    
    Answer = 15
    
    
    Open End Values below
    9                                 
    8                                 
    7             |-|     
    6             | |  x  |-|  x  x  |-|
    5     | |  x  | |  x  | |  x |-| | |
    4     | |  x  | |  x  | | |-|| | | |
    3     |-| |-| | |  x  | | | || | | |
    2     | | | | | | |-| | | | || | | |  x  |-|
    1     | | | | | | | | | | | || | | | |-| | |
           5   3   7   2   6   4  5   6   1   2
    
    answer = 10
    
    6                     |-|
    5             |-|  x  | |
    4             | |  x  | |
    3     |-|  x  | |  x  | |
    2     | | |-| | | |-| | |
    1     | | | | | | | | | | 
    answer = 4
    */
    nAmount := aWaterTowerCompute(aWater)
    
    fmt.Printf("The amount of water is %d\n", nAmount)
    
}

/*
 This function make two passes at the data. First by going left to right and finally right to left. 
 
 The first pass
 calculates the water amount (as it goes) between two towers (left to right) by using the follwoing logic:
    - when it reaches a tower that is equal to or taller than the beginning left tower, then the left 
      tower is reset to beginning left tower and the process starts again until all towers are evaluated.
    - NOTE: a left tower and right tower that are equal heights will be processed on the first pass
 The Second pass 
 calculates the water amount (as it goes) between two towers (right to left) by using the follwoing logic
    - same logic as the first pass with one exception.
    - Exception - The first time it finds a valid condition where it calculates amount of water, then it adds that amount
      to the total water and breaks out of the calculation loop.
      
 After two passes, the total amount of water calculation is complete.     
 NOTE - Reason for two passes. If the final tower is lower than the left tower in the first pass, then water value is 
        not calculated.

*/
func aWaterTowerCompute(aTower[]int) int{
	
	waterReturn,aLeftMin,tempWater := 0,0,0
	openEnd:=true
	// Go from left to right 
	for i,value := range aTower  {	
		aTemp := value
		openEnd=true		
		if(i == 0){
			aLeftMin=aTemp
		}else if(aLeftMin >= aTemp){
		      tempWater += aLeftMin - aTemp	
		}else{
			// set Left to new value
			waterReturn += tempWater 
			aLeftMin = aTemp
			tempWater,openEnd = 0,false	 
		}
	}
	fmt.Printf("First pass total = %d openEnd = %t\n",waterReturn,openEnd)		
	aLeftMin,tempWater= 0,0
	//go from right to left
	for i := len(aTower)-1; i >= 0; i-- {
		aTemp := aTower[i]
		if(i== len(aTower)-1){
			aLeftMin=aTemp	
	    }else if(aLeftMin > aTemp){
		      tempWater += aLeftMin - aTemp	
		}else{
			// Dont count it if both towers were equal height.. was already handled on first pass 
			if(aLeftMin != aTemp){
			   waterReturn += tempWater
			}else if(openEnd){
		       waterReturn += tempWater 		
			}  
			// set Left to new value			  
			aLeftMin = aTemp
			tempWater=0			
		}
	}		
	
	return waterReturn
}

