package unihash

import (
	"fmt"
	"sync"
)

var randomBase64 = []byte("Nz746LU-BCcolIygTV9Z0GaeX8puRKO5PEisvWDt3qbnrdFhf1wAMkHxQ_2jYmSJ")

var unRandomBase64 = []uint64{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 20, 49,
	58, 40, 3, 31, 4, 2, 25, 18, 0, 0, 0, 0, 0, 0, 0, 51, 8, 9, 38, 33, 46, 21,
	54, 13, 63, 29, 5, 52, 0, 30, 32, 56, 28, 62, 16, 6, 17, 37, 24, 60, 19, 0,
	0, 0, 0, 57, 0, 22, 42, 10, 45, 23, 48, 15, 47, 34, 59, 53, 12, 61, 43, 11,
	26, 41, 44, 35, 39, 27, 36, 50, 55, 14, 1, 0, 0, 0, 0, 0,
}

// Some more random base64 strings
//     "jt1vZpX9Fi6qHnRhrTb3-UuakzK0_JEW47wxeCO8Qf5IlgPsoYScDm2yNLdGAMBV";
//     "PGpN7Ws0gaFR6mvJT1UXl3bHBxtnuiyq-d9fj_wYckV2zSKIoA5rLMOeDC4ZhEQ8";
//     "3mqEds5hZkUyjD269ABplRHgI8iYzr-XOLbwF07ctou1SveV4KGQCMPNaxnTWfJ_";
//     "b9Tm5k27HuB-VyLEl13RIMwSKNGUDYQpnPhsJgavc6OiC4ofjFrAxd0_ztZqWX8e";
//     "Na3gFiQx1sS8LKyOuZrYBpjzwGEDPbomdq654RcIX_0e2C7k-WHnUhVAJlMf9Ttv";
//     "Zw84hDk-pN5uKcPy1_LdqIn0tQCJWBAROm26XSijzegxME7FHVbUaTGlorf3sYv9";
//     "fBe04QGcSkXsLud76gbxIFpOyUHajWiZmYMrEnDhtw5KqCRA8v3lPTz_o12VN9-J";
//     "RoFTY_jOZtbkai8651lp-VqzEgd4rLuDJ2WHBUv3xA9C0m7wKnsPhfMSQecGINXy";
//     "r3fi0dH_6kYyOaQ8s2eUBWucGS7PnNq9moFbTEh4C1xwMXJzIv-VZDljtRgLA5pK";
//     "vjOShxu1Cq8-JBsylNTGoiX5Kpt0cAEZr9VP2HMw3mkzFI4YL_bfRUegDWn7Qa6d";
//
// Convert randomBase64 to unRandomBase64
//	for k, v := range randomBase64 {
//		unRandomBase64[v] = uint64(k)
//	}

// Encode integer to hash string
func Encode(value uint64) (code [11]byte, size int) {
	var accumulate, remainder, position uint64

	for {
		accumulate += remainder
		remainder = value & 0x3f
		value >>= 6
		position = (accumulate + remainder) & 0x3f
		code[size] = randomBase64[position]
		size++

		if value == 0 {
			break
		}
	}

	return
}

// Decode hash string to integer
func Decode(code []byte) (value uint64) {
	var accumulate, remainder, position uint64

	size := len(code)
	for i := 0; i < size; i++ {
		position = unRandomBase64[code[i]]
		remainder = (position + 64 - accumulate) & 0x3f
		accumulate += remainder
		value |= remainder << uint64(6*i)
	}

	return
}

// Task - range of integers
type Task struct {
	Left  uint64
	Right uint64
}

// Worker for parallel
func Worker(id int, wg *sync.WaitGroup, inTask <-chan Task) {
	for t := range inTask {
		for i := t.Left; i < t.Right; i++ {
			code, size := Encode(i)
			Decode(code[:size])

			// Only for debuggin
			// if i != value {
			// 	fmt.Println("Decode Error", i, "->", string(code[:size]), "->", value)
			// }
		}

		fmt.Printf("Worker %d completed calculation of range [%d, %d).\n", id, t.Left, t.Right)
	}

	wg.Done()
}
