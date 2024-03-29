# Calendar
Perpetual calendar written in Go

## Examples

### Monthly
`$ ./cal`
```
       Jun 2023
 Su Mo Tu We Th Fr Sa
              1  2  3
  4  5  6  7  8  9 10
 11 12 13 14 15 16 17
 18 19 20 21 22 23 24
 25 26 27 28 29 30
```
### Quarterly
`$ ./cal .`
```
                                2023
         May                    Jun                    Jul          
 Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa 
     1  2  3  4  5  6                1  2  3                      1 
  7  8  9 10 11 12 13    4  5  6  7  8  9 10    2  3  4  5  6  7  8 
 14 15 16 17 18 19 20   11 12 13 14 15 16 17    9 10 11 12 13 14 15 
 21 22 23 24 25 26 27   18 19 20 21 22 23 24   16 17 18 19 20 21 22 
 28 29 30 31            25 26 27 28 29 30      23 24 25 26 27 28 29 
                                               30 31                
```
### Yearly
`$ ./cal +`
```
                                2023
         Jan                    Feb                    Mar           
 Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa  
  1  2  3  4  5  6  7             1  2  3  4             1  2  3  4  
  8  9 10 11 12 13 14    5  6  7  8  9 10 11    5  6  7  8  9 10 11  
 15 16 17 18 19 20 21   12 13 14 15 16 17 18   12 13 14 15 16 17 18  
 22 23 24 25 26 27 28   19 20 21 22 23 24 25   19 20 21 22 23 24 25  
 29 30 31               26 27 28               26 27 28 29 30 31     
                                                                     

         Apr                    May                    Jun           
 Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa  
                    1       1  2  3  4  5  6                1  2  3  
  2  3  4  5  6  7  8    7  8  9 10 11 12 13    4  5  6  7  8  9 10  
  9 10 11 12 13 14 15   14 15 16 17 18 19 20   11 12 13 14 15 16 17  
 16 17 18 19 20 21 22   21 22 23 24 25 26 27   18 19 20 21 22 23 24  
 23 24 25 26 27 28 29   28 29 30 31            25 26 27 28 29 30     
 30                                                                  

         Jul                    Aug                    Sep           
 Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa  
                    1          1  2  3  4  5                   1  2  
  2  3  4  5  6  7  8    6  7  8  9 10 11 12    3  4  5  6  7  8  9  
  9 10 11 12 13 14 15   13 14 15 16 17 18 19   10 11 12 13 14 15 16  
 16 17 18 19 20 21 22   20 21 22 23 24 25 26   17 18 19 20 21 22 23  
 23 24 25 26 27 28 29   27 28 29 30 31         24 25 26 27 28 29 30  
 30 31                                                               

         Oct                    Nov                    Dec           
 Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa  
  1  2  3  4  5  6  7             1  2  3  4                   1  2  
  8  9 10 11 12 13 14    5  6  7  8  9 10 11    3  4  5  6  7  8  9  
 15 16 17 18 19 20 21   12 13 14 15 16 17 18   10 11 12 13 14 15 16  
 22 23 24 25 26 27 28   19 20 21 22 23 24 25   17 18 19 20 21 22 23  
 29 30 31               26 27 28 29 30         24 25 26 27 28 29 30  
                                               31                    
```

Copyright (c) 1582 Pope Gregory XIII 
