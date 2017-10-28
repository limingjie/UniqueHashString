package main

import "fmt"

var randomBase64 = []byte("Nz746LU-BCcolIygTV9Z0GaeX8puRKO5PEisvWDt3qbnrdFhf1wAMkHxQ_2jYmSJ")

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

var unRandomBase64 = make(map[byte]uint64)

func encode(value uint64) (code []byte) {
	var accumulate, remainder, position uint64

	for {
		accumulate += remainder
		remainder = value & 0x3f
		value >>= 6
		position = (accumulate + remainder) & 0x3f
		code = append(code, randomBase64[position])

		if value == 0 {
			break
		}
	}

	return
}

func decode(code []byte) (value uint64) {
	var accumulate, remainder, position uint64

	size := len(code)
	for i := 0; i < size; i++ {
		position = unRandomBase64[code[i]]
		remainder = (position + 64 - accumulate) & 0x3f
		accumulate += remainder
		value += remainder << uint64(6*i)
	}

	return
}

func main() {
	for k, v := range randomBase64 {
		unRandomBase64[v] = uint64(k)
	}

	// fmt.Println("Value -> encode() -> decode()")

	count := 0
	for i := uint64(16345678912345678900); i < uint64(16345678912345678900+65536000); i++ {
		code := encode(i)
		value := decode(code[0:])
		// fmt.Println(i, "->", string(code), "->", value)

		if i != value {
			count++
		}

		if i%65536 == 0 {
			fmt.Print(".")
		}
	}

	fmt.Println("Decode Error Count =", count)
}
