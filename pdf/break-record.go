import (
	"fmt"
)

func breakingRecords(scores []int32) []int32 {
    // Write your code here
    lowestR := scores[0]
    highestR := scores[0]
    var countLow int32 = 0
    var countHigh int32 = 0
    
    for _, x := range scores {
        if x > highestR {
            countHigh++
            highestR = x
        } else if x < lowestR {
            countLow++
            lowestR = x
        }
    }
    
    return []int32{countHigh, countLow}
}


func main() {
	result := breakingRecords([]int32{10, 6, 10, 20, 12, 30, 25, 4})
	fmt.Println(result)
}