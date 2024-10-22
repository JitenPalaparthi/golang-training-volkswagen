- take an array of any

arr3 := [...]any{"Hello", 123, 4343.343, true, 'A',10,123,1232.23,false,'k',4,343.34343,34}

- iterate thru the array 
- if the element is a number (any of uint-int,uint8-uint64,int8-int64,float32,float64,rune,byte) .. sum those numbers
- if the element is of a type bool-> add 1 if true add 0 if false
- print the sum 