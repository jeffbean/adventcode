package main

import (
	"container/ring"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jeffbean/adventcode/2017/common"
)

// Knot is a has knot ring that tracks information of a ring
type Knot struct {
	*ring.Ring
	skipSize int
}

func NewKnot(l int) *Knot {
	r := ring.New(l)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	return &Knot{Ring: r, skipSize: 0}
}

func (k *Knot) String() string {
	var r string
	k.Do(func(x interface{}) {
		r += fmt.Sprintf("%v,", x)
	})

	return r
}
func (k *Knot) TwistHash(length int) {
	// fmt.Printf("playing with ring %v\n", k)

	orig := k.Ring
	nums := []int{}
	for ; length > 0; length-- {
		nums = append(nums, k.Value.(int))
		k.Ring = k.Next()
	}
	// fmt.Println()
	k.Ring = orig
	// fmt.Println(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	// fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		k.Value = nums[i]
		k.Ring = k.Next()
	}
	for j := 0; j < k.skipSize; j++ {
		k.Ring = k.Next()
	}
}

func (k *Knot) twitshHashLengths(lengths []int) {
	for _, val := range lengths {
		k.TwistHash(val)
		// increase skip size each time we twist
		k.skipSize++
		// fmt.Println(k)
	}
	// we are starting at "current position" and the array is one ahead of that
	k.Ring = k.Next()
	// fmt.Println(k)
}

func part1(data []byte) *Knot {
	k := NewKnot(256)
	var nums []int
	lengths := strings.Split(string(data), ",")
	for _, length := range lengths {
		num, err := strconv.Atoi(strings.TrimSpace(length))
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	k.twitshHashLengths(nums)
	return k
}
func part2() int {
	return 0
}

func main() {
	data := common.GetInputFile()
	// log.Println(data)
	fmt.Printf("%v\n", part1(data))
	// fmt.Printf("%v\n", part2())
}
