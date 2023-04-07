//package encoder
package main

import "fmt"

//    - NOTE: 2^9 = 28^8 * 2
//    - polynomial long division
//    - XOR (^) if decimal is >= 256 then XOR with 285 -> recursive
//    - When adding exponents, if the exponent becomes greater than or equal to 256, simply apply modulo 255
//       - 2^170 * 2^164 = 2^(170+164) = 2^334 -> 2^334%255 = 2^79
//    - https://www.thonky.com/qr-code-tutorial/log-antilog-table
//
//
// Generator Polynomial for 2 Error Correction Codewords!
// 
//    - TODO:
//       - create message polynomial: <NUMBER>^len-(index+1) + <NUMBER>^len-(index+1), if ^n == 0 -> return just number
//       - generator polynomial: x^10 + α^251x^9 + α^67x^8 + α^46x^7 + α^61x^6 + α^118x^5 + α^70x^4 + α^64x^3 + α^94x^2 + α^32x + α^45 
//
//       - division:
//          1) multiply message x^10 -> <NUMBER>x^(exp + 10)
//          2) multiply generator by x^15, to have same exponentional number as message, if one "component" do not have a^n, add a^0
//          3) now you can divide
//       - Multiply generator, repeat 16x
//          1) multiply generator by index[0] from message, like its 32x^25, 32 in table == 5 so it's a^5
//          2) if exponent >255, do % 255
//          3) now convert it to numbers, where a == 2, so output will be 32x^25 + 2x^24 + 101x^23 + 10x^22 + 97x^21 + 197x^20 + 15x^19 + 47x^18 + 134x^17 + 74x^16 + 5x^15
//          4) do XOR with message a generator, (NUM_MESSAGE XOR NUM_GENERATOR)x^GENRATOR_EXP
//          4) NOTE: if something after it is 0x^n, remove it

//          5) If it's 16x done (16 = len message), remove exponents and X == result


func main() {
	// Define the message polynomial
//	msg := []int{32, 91, 11, 120, 209, 114, 220, 77, 67, 64, 236, 17, 236, 17, 236, 17}


	// Divide message polynomial by generator polynomial to get the codeword

	// Print the codeword
	fmt.Println(256 ^ 285)
}

