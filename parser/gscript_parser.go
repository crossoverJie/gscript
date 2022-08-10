// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 354,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 7, 4, 80, 10, 4, 12, 4, 14, 4, 83, 11, 4, 3, 5, 3, 5, 3, 5, 5, 5,
	88, 10, 5, 3, 6, 6, 6, 91, 10, 6, 13, 6, 14, 6, 92, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 101, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 120,
	10, 7, 12, 7, 14, 7, 123, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 131, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 139, 10, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 149, 10, 9, 3, 9, 5,
	9, 152, 10, 9, 3, 10, 5, 10, 155, 10, 10, 3, 10, 3, 10, 5, 10, 159, 10,
	10, 3, 10, 3, 10, 5, 10, 163, 10, 10, 3, 11, 3, 11, 5, 11, 167, 10, 11,
	3, 12, 5, 12, 170, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 176, 10,
	12, 12, 12, 14, 12, 179, 11, 12, 3, 12, 3, 12, 5, 12, 183, 10, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 192, 10, 14, 12, 14, 14,
	14, 195, 11, 14, 3, 15, 3, 15, 5, 15, 199, 10, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 7, 16, 206, 10, 16, 12, 16, 14, 16, 209, 11, 16, 3, 16, 3,
	16, 5, 16, 213, 10, 16, 3, 16, 5, 16, 216, 10, 16, 3, 17, 7, 17, 219, 10,
	17, 12, 17, 14, 17, 222, 11, 17, 3, 17, 3, 17, 3, 17, 3, 18, 7, 18, 228,
	10, 18, 12, 18, 14, 18, 231, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 19, 7, 19, 240, 10, 19, 12, 19, 14, 19, 243, 11, 19, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 251, 10, 21, 12, 21, 14, 21, 254,
	11, 21, 3, 22, 3, 22, 3, 22, 5, 22, 259, 10, 22, 3, 23, 3, 23, 3, 23, 7,
	23, 264, 10, 23, 12, 23, 14, 23, 267, 11, 23, 3, 24, 3, 24, 5, 24, 271,
	10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 277, 10, 25, 12, 25, 14, 25,
	280, 11, 25, 3, 25, 5, 25, 283, 10, 25, 5, 25, 285, 10, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 294, 10, 26, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 303, 10, 29, 3, 29, 3, 29, 7, 29,
	307, 10, 29, 12, 29, 14, 29, 310, 11, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 5, 31, 318, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32,
	325, 10, 32, 12, 32, 14, 32, 328, 11, 32, 3, 33, 3, 33, 5, 33, 332, 10,
	33, 3, 34, 3, 34, 3, 34, 7, 34, 337, 10, 34, 12, 34, 14, 34, 340, 11, 34,
	3, 35, 3, 35, 3, 35, 7, 35, 345, 10, 35, 12, 35, 14, 35, 348, 11, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 2, 3, 12, 37, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 9, 3, 2, 22, 23, 3, 2, 26, 27, 3,
	2, 28, 29, 4, 2, 20, 21, 32, 33, 4, 2, 31, 31, 34, 34, 3, 2, 24, 25, 3,
	2, 9, 12, 2, 371, 2, 72, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 81, 3, 2, 2,
	2, 8, 87, 3, 2, 2, 2, 10, 90, 3, 2, 2, 2, 12, 100, 3, 2, 2, 2, 14, 130,
	3, 2, 2, 2, 16, 151, 3, 2, 2, 2, 18, 154, 3, 2, 2, 2, 20, 166, 3, 2, 2,
	2, 22, 169, 3, 2, 2, 2, 24, 186, 3, 2, 2, 2, 26, 188, 3, 2, 2, 2, 28, 196,
	3, 2, 2, 2, 30, 215, 3, 2, 2, 2, 32, 220, 3, 2, 2, 2, 34, 229, 3, 2, 2,
	2, 36, 236, 3, 2, 2, 2, 38, 244, 3, 2, 2, 2, 40, 246, 3, 2, 2, 2, 42, 255,
	3, 2, 2, 2, 44, 260, 3, 2, 2, 2, 46, 270, 3, 2, 2, 2, 48, 272, 3, 2, 2,
	2, 50, 293, 3, 2, 2, 2, 52, 295, 3, 2, 2, 2, 54, 297, 3, 2, 2, 2, 56, 302,
	3, 2, 2, 2, 58, 311, 3, 2, 2, 2, 60, 313, 3, 2, 2, 2, 62, 321, 3, 2, 2,
	2, 64, 331, 3, 2, 2, 2, 66, 333, 3, 2, 2, 2, 68, 341, 3, 2, 2, 2, 70, 349,
	3, 2, 2, 2, 72, 73, 5, 6, 4, 2, 73, 3, 3, 2, 2, 2, 74, 75, 7, 15, 2, 2,
	75, 76, 5, 6, 4, 2, 76, 77, 7, 16, 2, 2, 77, 5, 3, 2, 2, 2, 78, 80, 5,
	8, 5, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81,
	82, 3, 2, 2, 2, 82, 7, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 88, 5, 40, 21,
	2, 85, 88, 5, 16, 9, 2, 86, 88, 5, 22, 12, 2, 87, 84, 3, 2, 2, 2, 87, 85,
	3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 9, 3, 2, 2, 2, 89, 91, 5, 12, 7, 2,
	90, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3,
	2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 7, 2, 2, 3, 95, 11, 3, 2, 2, 2, 96,
	97, 8, 7, 1, 2, 97, 101, 5, 14, 8, 2, 98, 99, 9, 2, 2, 2, 99, 101, 5, 12,
	7, 8, 100, 96, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 121, 3, 2, 2, 2, 102,
	103, 12, 7, 2, 2, 103, 104, 9, 3, 2, 2, 104, 120, 5, 12, 7, 8, 105, 106,
	12, 6, 2, 2, 106, 107, 7, 30, 2, 2, 107, 120, 5, 12, 7, 7, 108, 109, 12,
	5, 2, 2, 109, 110, 9, 4, 2, 2, 110, 120, 5, 12, 7, 6, 111, 112, 12, 4,
	2, 2, 112, 113, 9, 5, 2, 2, 113, 120, 5, 12, 7, 5, 114, 115, 12, 3, 2,
	2, 115, 116, 9, 6, 2, 2, 116, 120, 5, 12, 7, 4, 117, 118, 12, 9, 2, 2,
	118, 120, 9, 7, 2, 2, 119, 102, 3, 2, 2, 2, 119, 105, 3, 2, 2, 2, 119,
	108, 3, 2, 2, 2, 119, 111, 3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 119, 117,
	3, 2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2,
	2, 2, 122, 13, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124, 125, 7, 13, 2, 2,
	125, 126, 5, 12, 7, 2, 126, 127, 7, 14, 2, 2, 127, 131, 3, 2, 2, 2, 128,
	131, 5, 50, 26, 2, 129, 131, 7, 46, 2, 2, 130, 124, 3, 2, 2, 2, 130, 128,
	3, 2, 2, 2, 130, 129, 3, 2, 2, 2, 131, 15, 3, 2, 2, 2, 132, 152, 5, 4,
	3, 2, 133, 134, 7, 38, 2, 2, 134, 135, 5, 70, 36, 2, 135, 138, 5, 16, 9,
	2, 136, 137, 7, 39, 2, 2, 137, 139, 5, 16, 9, 2, 138, 136, 3, 2, 2, 2,
	138, 139, 3, 2, 2, 2, 139, 152, 3, 2, 2, 2, 140, 141, 7, 37, 2, 2, 141,
	142, 7, 13, 2, 2, 142, 143, 5, 18, 10, 2, 143, 144, 7, 14, 2, 2, 144, 145,
	5, 16, 9, 2, 145, 152, 3, 2, 2, 2, 146, 148, 7, 40, 2, 2, 147, 149, 5,
	12, 7, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 152, 3, 2, 2,
	2, 150, 152, 5, 12, 7, 2, 151, 132, 3, 2, 2, 2, 151, 133, 3, 2, 2, 2, 151,
	140, 3, 2, 2, 2, 151, 146, 3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 17, 3,
	2, 2, 2, 153, 155, 5, 20, 11, 2, 154, 153, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 7, 3, 2, 2, 157, 159, 5, 12, 7, 2,
	158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	162, 7, 3, 2, 2, 161, 163, 5, 68, 35, 2, 162, 161, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 19, 3, 2, 2, 2, 164, 167, 5, 40, 21, 2, 165, 167, 5, 68,
	35, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 21, 3, 2, 2, 2,
	168, 170, 5, 64, 33, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	171, 3, 2, 2, 2, 171, 172, 7, 46, 2, 2, 172, 177, 5, 28, 15, 2, 173, 174,
	7, 17, 2, 2, 174, 176, 7, 18, 2, 2, 175, 173, 3, 2, 2, 2, 176, 179, 3,
	2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 182, 3, 2, 2,
	2, 179, 177, 3, 2, 2, 2, 180, 181, 7, 8, 2, 2, 181, 183, 5, 26, 14, 2,
	182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184,
	185, 5, 24, 13, 2, 185, 23, 3, 2, 2, 2, 186, 187, 5, 4, 3, 2, 187, 25,
	3, 2, 2, 2, 188, 193, 5, 36, 19, 2, 189, 190, 7, 4, 2, 2, 190, 192, 5,
	36, 19, 2, 191, 189, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2,
	2, 2, 193, 194, 3, 2, 2, 2, 194, 27, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2,
	196, 198, 7, 13, 2, 2, 197, 199, 5, 30, 16, 2, 198, 197, 3, 2, 2, 2, 198,
	199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 7, 14, 2, 2, 201, 29,
	3, 2, 2, 2, 202, 207, 5, 32, 17, 2, 203, 204, 7, 4, 2, 2, 204, 206, 5,
	32, 17, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2,
	2, 2, 207, 208, 3, 2, 2, 2, 208, 212, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2,
	210, 211, 7, 4, 2, 2, 211, 213, 5, 34, 18, 2, 212, 210, 3, 2, 2, 2, 212,
	213, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 216, 5, 34, 18, 2, 215, 202,
	3, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 31, 3, 2, 2, 2, 217, 219, 5, 38,
	20, 2, 218, 217, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2,
	220, 221, 3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223,
	224, 5, 56, 29, 2, 224, 225, 5, 44, 23, 2, 225, 33, 3, 2, 2, 2, 226, 228,
	5, 38, 20, 2, 227, 226, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3,
	2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232, 3, 2, 2, 2, 231, 229, 3, 2, 2,
	2, 232, 233, 5, 56, 29, 2, 233, 234, 7, 5, 2, 2, 234, 235, 5, 44, 23, 2,
	235, 35, 3, 2, 2, 2, 236, 241, 7, 46, 2, 2, 237, 238, 7, 6, 2, 2, 238,
	240, 7, 46, 2, 2, 239, 237, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239,
	3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 37, 3, 2, 2, 2, 243, 241, 3, 2,
	2, 2, 244, 245, 7, 7, 2, 2, 245, 39, 3, 2, 2, 2, 246, 247, 5, 56, 29, 2,
	247, 252, 5, 42, 22, 2, 248, 249, 7, 4, 2, 2, 249, 251, 5, 42, 22, 2, 250,
	248, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253,
	3, 2, 2, 2, 253, 41, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 258, 5, 44,
	23, 2, 256, 257, 7, 19, 2, 2, 257, 259, 5, 46, 24, 2, 258, 256, 3, 2, 2,
	2, 258, 259, 3, 2, 2, 2, 259, 43, 3, 2, 2, 2, 260, 265, 7, 46, 2, 2, 261,
	262, 7, 17, 2, 2, 262, 264, 7, 18, 2, 2, 263, 261, 3, 2, 2, 2, 264, 267,
	3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 45, 3, 2,
	2, 2, 267, 265, 3, 2, 2, 2, 268, 271, 5, 48, 25, 2, 269, 271, 5, 12, 7,
	2, 270, 268, 3, 2, 2, 2, 270, 269, 3, 2, 2, 2, 271, 47, 3, 2, 2, 2, 272,
	284, 7, 15, 2, 2, 273, 278, 5, 46, 24, 2, 274, 275, 7, 4, 2, 2, 275, 277,
	5, 46, 24, 2, 276, 274, 3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3,
	2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 282, 3, 2, 2, 2, 280, 278, 3, 2, 2,
	2, 281, 283, 7, 4, 2, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283,
	285, 3, 2, 2, 2, 284, 273, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 286,
	3, 2, 2, 2, 286, 287, 7, 16, 2, 2, 287, 49, 3, 2, 2, 2, 288, 294, 7, 44,
	2, 2, 289, 294, 7, 45, 2, 2, 290, 294, 7, 42, 2, 2, 291, 294, 7, 41, 2,
	2, 292, 294, 7, 43, 2, 2, 293, 288, 3, 2, 2, 2, 293, 289, 3, 2, 2, 2, 293,
	290, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293, 292, 3, 2, 2, 2, 294, 51, 3,
	2, 2, 2, 295, 296, 7, 44, 2, 2, 296, 53, 3, 2, 2, 2, 297, 298, 7, 45, 2,
	2, 298, 55, 3, 2, 2, 2, 299, 303, 5, 66, 34, 2, 300, 303, 5, 60, 31, 2,
	301, 303, 5, 58, 30, 2, 302, 299, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302,
	301, 3, 2, 2, 2, 303, 308, 3, 2, 2, 2, 304, 305, 7, 17, 2, 2, 305, 307,
	7, 18, 2, 2, 306, 304, 3, 2, 2, 2, 307, 310, 3, 2, 2, 2, 308, 306, 3, 2,
	2, 2, 308, 309, 3, 2, 2, 2, 309, 57, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2,
	311, 312, 9, 8, 2, 2, 312, 59, 3, 2, 2, 2, 313, 314, 7, 35, 2, 2, 314,
	315, 5, 64, 33, 2, 315, 317, 7, 13, 2, 2, 316, 318, 5, 62, 32, 2, 317,
	316, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 320,
	7, 14, 2, 2, 320, 61, 3, 2, 2, 2, 321, 326, 5, 56, 29, 2, 322, 323, 7,
	4, 2, 2, 323, 325, 5, 56, 29, 2, 324, 322, 3, 2, 2, 2, 325, 328, 3, 2,
	2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 63, 3, 2, 2, 2,
	328, 326, 3, 2, 2, 2, 329, 332, 5, 56, 29, 2, 330, 332, 7, 36, 2, 2, 331,
	329, 3, 2, 2, 2, 331, 330, 3, 2, 2, 2, 332, 65, 3, 2, 2, 2, 333, 338, 7,
	46, 2, 2, 334, 335, 7, 6, 2, 2, 335, 337, 7, 46, 2, 2, 336, 334, 3, 2,
	2, 2, 337, 340, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2,
	339, 67, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 341, 346, 5, 12, 7, 2, 342,
	343, 7, 4, 2, 2, 343, 345, 5, 12, 7, 2, 344, 342, 3, 2, 2, 2, 345, 348,
	3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 69, 3, 2,
	2, 2, 348, 346, 3, 2, 2, 2, 349, 350, 7, 13, 2, 2, 350, 351, 5, 12, 7,
	2, 351, 352, 7, 14, 2, 2, 352, 71, 3, 2, 2, 2, 42, 81, 87, 92, 100, 119,
	121, 130, 138, 148, 151, 154, 158, 162, 166, 169, 177, 182, 193, 198, 207,
	212, 215, 220, 229, 241, 252, 258, 265, 270, 278, 282, 284, 293, 302, 308,
	317, 326, 331, 338, 346,
}
var literalNames = []string{
	"", "';'", "','", "'...'", "'.'", "'final'", "'throws'", "'int'", "'string'",
	"'float'", "'bool'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "'>'",
	"'<'", "'!'", "'~'", "'++'", "'--'", "'*'", "'/'", "'+'", "'-'", "'%'",
	"'=='", "'<='", "'>='", "'!='", "'func'", "'void'", "'for'", "'if'", "'else'",
	"'return'", "", "", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "FINAL", "THROWS", "INT", "STRING", "FLOAT", "BOOLEAN",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "ASSIGN", "GT",
	"LT", "BANG", "TILDE", "INC", "DEC", "MULT", "DIV", "PLUS", "SUB", "MOD",
	"EQUAL", "LE", "GE", "NOTEQUAL", "FUNCTION", "VOID", "FOR", "IF", "ELSE",
	"RETURN", "BOOL_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "DECIMAL_LITERAL",
	"FLOAT_LITERAL", "IDENTIFIER", "SPACES",
}

var ruleNames = []string{
	"prog", "block", "blockStatements", "blockStatement", "parse", "expr",
	"primary", "statement", "forControl", "forInit", "functionDeclaration",
	"functionBody", "qualifiedNameList", "formalParameters", "formalParameterList",
	"formalParameter", "lastFormalParameter", "qualifiedName", "variableModifier",
	"variableDeclarators", "variableDeclarator", "variableDeclaratorId", "variableInitializer",
	"arrayInitializer", "literal", "integerLiteral", "floatLiteral", "typeType",
	"primitiveType", "functionType", "typeList", "typeTypeOrVoid", "classOrInterfaceType",
	"expressionList", "parExpression",
}

type GScriptParser struct {
	*antlr.BaseParser
}

// NewGScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGScriptParser(input antlr.TokenStream) *GScriptParser {
	this := new(GScriptParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GScript.g4"

	return this
}

// GScriptParser tokens.
const (
	GScriptParserEOF             = antlr.TokenEOF
	GScriptParserT__0            = 1
	GScriptParserT__1            = 2
	GScriptParserT__2            = 3
	GScriptParserT__3            = 4
	GScriptParserFINAL           = 5
	GScriptParserTHROWS          = 6
	GScriptParserINT             = 7
	GScriptParserSTRING          = 8
	GScriptParserFLOAT           = 9
	GScriptParserBOOLEAN         = 10
	GScriptParserLPAREN          = 11
	GScriptParserRPAREN          = 12
	GScriptParserLBRACE          = 13
	GScriptParserRBRACE          = 14
	GScriptParserLBRACK          = 15
	GScriptParserRBRACK          = 16
	GScriptParserASSIGN          = 17
	GScriptParserGT              = 18
	GScriptParserLT              = 19
	GScriptParserBANG            = 20
	GScriptParserTILDE           = 21
	GScriptParserINC             = 22
	GScriptParserDEC             = 23
	GScriptParserMULT            = 24
	GScriptParserDIV             = 25
	GScriptParserPLUS            = 26
	GScriptParserSUB             = 27
	GScriptParserMOD             = 28
	GScriptParserEQUAL           = 29
	GScriptParserLE              = 30
	GScriptParserGE              = 31
	GScriptParserNOTEQUAL        = 32
	GScriptParserFUNCTION        = 33
	GScriptParserVOID            = 34
	GScriptParserFOR             = 35
	GScriptParserIF              = 36
	GScriptParserELSE            = 37
	GScriptParserRETURN          = 38
	GScriptParserBOOL_LITERAL    = 39
	GScriptParserSTRING_LITERAL  = 40
	GScriptParserNULL_LITERAL    = 41
	GScriptParserDECIMAL_LITERAL = 42
	GScriptParserFLOAT_LITERAL   = 43
	GScriptParserIDENTIFIER      = 44
	GScriptParserSPACES          = 45
)

// GScriptParser rules.
const (
	GScriptParserRULE_prog                 = 0
	GScriptParserRULE_block                = 1
	GScriptParserRULE_blockStatements      = 2
	GScriptParserRULE_blockStatement       = 3
	GScriptParserRULE_parse                = 4
	GScriptParserRULE_expr                 = 5
	GScriptParserRULE_primary              = 6
	GScriptParserRULE_statement            = 7
	GScriptParserRULE_forControl           = 8
	GScriptParserRULE_forInit              = 9
	GScriptParserRULE_functionDeclaration  = 10
	GScriptParserRULE_functionBody         = 11
	GScriptParserRULE_qualifiedNameList    = 12
	GScriptParserRULE_formalParameters     = 13
	GScriptParserRULE_formalParameterList  = 14
	GScriptParserRULE_formalParameter      = 15
	GScriptParserRULE_lastFormalParameter  = 16
	GScriptParserRULE_qualifiedName        = 17
	GScriptParserRULE_variableModifier     = 18
	GScriptParserRULE_variableDeclarators  = 19
	GScriptParserRULE_variableDeclarator   = 20
	GScriptParserRULE_variableDeclaratorId = 21
	GScriptParserRULE_variableInitializer  = 22
	GScriptParserRULE_arrayInitializer     = 23
	GScriptParserRULE_literal              = 24
	GScriptParserRULE_integerLiteral       = 25
	GScriptParserRULE_floatLiteral         = 26
	GScriptParserRULE_typeType             = 27
	GScriptParserRULE_primitiveType        = 28
	GScriptParserRULE_functionType         = 29
	GScriptParserRULE_typeList             = 30
	GScriptParserRULE_typeTypeOrVoid       = 31
	GScriptParserRULE_classOrInterfaceType = 32
	GScriptParserRULE_expressionList       = 33
	GScriptParserRULE_parExpression        = 34
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GScriptParserRULE_prog)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.BlockStatements()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GScriptParserLBRACE, 0)
}

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GScriptParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(GScriptParserLBRACE)
	}
	{
		p.SetState(73)
		p.BlockStatements()
	}
	{
		p.SetState(74)
		p.Match(GScriptParserRBRACE)
	}

	return localctx
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_blockStatements
	return p
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) CopyFrom(ctx *BlockStatementsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockStmsContext struct {
	*BlockStatementsContext
}

func NewBlockStmsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmsContext {
	var p = new(BlockStmsContext)

	p.BlockStatementsContext = NewEmptyBlockStatementsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementsContext))

	return p
}

func (s *BlockStmsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmsContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockStmsContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStmsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlockStms(s)
	}
}

func (s *BlockStmsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlockStms(s)
	}
}

func (s *BlockStmsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockStms(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) BlockStatements() (localctx IBlockStatementsContext) {
	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GScriptParserRULE_blockStatements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewBlockStmsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserINT)|(1<<GScriptParserSTRING)|(1<<GScriptParserFLOAT)|(1<<GScriptParserBOOLEAN)|(1<<GScriptParserLPAREN)|(1<<GScriptParserLBRACE)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GScriptParserFUNCTION-33))|(1<<(GScriptParserVOID-33))|(1<<(GScriptParserFOR-33))|(1<<(GScriptParserIF-33))|(1<<(GScriptParserRETURN-33))|(1<<(GScriptParserBOOL_LITERAL-33))|(1<<(GScriptParserSTRING_LITERAL-33))|(1<<(GScriptParserNULL_LITERAL-33))|(1<<(GScriptParserDECIMAL_LITERAL-33))|(1<<(GScriptParserFLOAT_LITERAL-33))|(1<<(GScriptParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(76)
			p.BlockStatement()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) CopyFrom(ctx *BlockStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockStmContext struct {
	*BlockStatementContext
}

func NewBlockStmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmContext {
	var p = new(BlockStmContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockStmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlockStm(s)
	}
}

func (s *BlockStmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlockStm(s)
	}
}

func (s *BlockStmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockStm(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockFuncContext struct {
	*BlockStatementContext
}

func NewBlockFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockFuncContext {
	var p = new(BlockFuncContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockFuncContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *BlockFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlockFunc(s)
	}
}

func (s *BlockFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlockFunc(s)
	}
}

func (s *BlockFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockVarDeclarContext struct {
	*BlockStatementContext
}

func NewBlockVarDeclarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockVarDeclarContext {
	var p = new(BlockVarDeclarContext)

	p.BlockStatementContext = NewEmptyBlockStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BlockStatementContext))

	return p
}

func (s *BlockVarDeclarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockVarDeclarContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *BlockVarDeclarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlockVarDeclar(s)
	}
}

func (s *BlockVarDeclarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlockVarDeclar(s)
	}
}

func (s *BlockVarDeclarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockVarDeclar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GScriptParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBlockVarDeclarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.VariableDeclarators()
		}

	case 2:
		localctx = NewBlockStmContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Statement()
		}

	case 3:
		localctx = NewBlockFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.FunctionDeclaration()
		}

	}

	return localctx
}

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetExpr_list returns the expr_list rule context list.
	GetExpr_list() []IExprContext

	// SetExpr_list sets the expr_list rule context list.
	SetExpr_list([]IExprContext)

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_expr     IExprContext
	expr_list []IExprContext
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Get_expr() IExprContext { return s._expr }

func (s *ParseContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ParseContext) GetExpr_list() []IExprContext { return s.expr_list }

func (s *ParseContext) SetExpr_list(v []IExprContext) { s.expr_list = v }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(GScriptParserEOF, 0)
}

func (s *ParseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ParseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GScriptParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserLPAREN)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(GScriptParserBOOL_LITERAL-39))|(1<<(GScriptParserSTRING_LITERAL-39))|(1<<(GScriptParserNULL_LITERAL-39))|(1<<(GScriptParserDECIMAL_LITERAL-39))|(1<<(GScriptParserFLOAT_LITERAL-39))|(1<<(GScriptParserIDENTIFIER-39)))) != 0) {
		{
			p.SetState(87)

			var _x = p.expr(0)

			localctx.(*ParseContext)._expr = _x
		}
		localctx.(*ParseContext).expr_list = append(localctx.(*ParseContext).expr_list, localctx.(*ParseContext)._expr)

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(GScriptParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultDivExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewMultDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultDivExprContext {
	var p = new(MultDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultDivExprContext) GetBop() antlr.Token { return s.bop }

func (s *MultDivExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *MultDivExprContext) GetLhs() IExprContext { return s.lhs }

func (s *MultDivExprContext) GetRhs() IExprContext { return s.rhs }

func (s *MultDivExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *MultDivExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *MultDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultDivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultDivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultDivExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(GScriptParserMULT, 0)
}

func (s *MultDivExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(GScriptParserDIV, 0)
}

func (s *MultDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterMultDivExpr(s)
	}
}

func (s *MultDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitMultDivExpr(s)
	}
}

func (s *MultDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitMultDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixExprContext struct {
	*ExprContext
	lhs     IExprContext
	postfix antlr.Token
}

func NewPostfixExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixExprContext {
	var p = new(PostfixExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PostfixExprContext) GetPostfix() antlr.Token { return s.postfix }

func (s *PostfixExprContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *PostfixExprContext) GetLhs() IExprContext { return s.lhs }

func (s *PostfixExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *PostfixExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PostfixExprContext) INC() antlr.TerminalNode {
	return s.GetToken(GScriptParserINC, 0)
}

func (s *PostfixExprContext) DEC() antlr.TerminalNode {
	return s.GetToken(GScriptParserDEC, 0)
}

func (s *PostfixExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterPostfixExpr(s)
	}
}

func (s *PostfixExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitPostfixExpr(s)
	}
}

func (s *PostfixExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPostfixExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GLeContext struct {
	*ExprContext
	bop antlr.Token
}

func NewGLeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GLeContext {
	var p = new(GLeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GLeContext) GetBop() antlr.Token { return s.bop }

func (s *GLeContext) SetBop(v antlr.Token) { s.bop = v }

func (s *GLeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GLeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GLeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GLeContext) LE() antlr.TerminalNode {
	return s.GetToken(GScriptParserLE, 0)
}

func (s *GLeContext) GE() antlr.TerminalNode {
	return s.GetToken(GScriptParserGE, 0)
}

func (s *GLeContext) GT() antlr.TerminalNode {
	return s.GetToken(GScriptParserGT, 0)
}

func (s *GLeContext) LT() antlr.TerminalNode {
	return s.GetToken(GScriptParserLT, 0)
}

func (s *GLeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterGLe(s)
	}
}

func (s *GLeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitGLe(s)
	}
}

func (s *GLeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitGLe(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusSubExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewPlusSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusSubExprContext {
	var p = new(PlusSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PlusSubExprContext) GetBop() antlr.Token { return s.bop }

func (s *PlusSubExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *PlusSubExprContext) GetLhs() IExprContext { return s.lhs }

func (s *PlusSubExprContext) GetRhs() IExprContext { return s.rhs }

func (s *PlusSubExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *PlusSubExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *PlusSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PlusSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PlusSubExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GScriptParserPLUS, 0)
}

func (s *PlusSubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(GScriptParserSUB, 0)
}

func (s *PlusSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterPlusSubExpr(s)
	}
}

func (s *PlusSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitPlusSubExpr(s)
	}
}

func (s *PlusSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPlusSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	*ExprContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExprContext
	prefix antlr.Token
	rhs    IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetPrefix() antlr.Token { return s.prefix }

func (s *NotExprContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *NotExprContext) GetRhs() IExprContext { return s.rhs }

func (s *NotExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(GScriptParserTILDE, 0)
}

func (s *NotExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(GScriptParserBANG, 0)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ModExprContext struct {
	*ExprContext
	lhs IExprContext
	bop antlr.Token
	rhs IExprContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ModExprContext) GetBop() antlr.Token { return s.bop }

func (s *ModExprContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ModExprContext) GetLhs() IExprContext { return s.lhs }

func (s *ModExprContext) GetRhs() IExprContext { return s.rhs }

func (s *ModExprContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ModExprContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ModExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(GScriptParserMOD, 0)
}

func (s *ModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterModExpr(s)
	}
}

func (s *ModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitModExpr(s)
	}
}

func (s *ModExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitModExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqualOrNotContext struct {
	*ExprContext
	bop antlr.Token
}

func NewEqualOrNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualOrNotContext {
	var p = new(EqualOrNotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualOrNotContext) GetBop() antlr.Token { return s.bop }

func (s *EqualOrNotContext) SetBop(v antlr.Token) { s.bop = v }

func (s *EqualOrNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualOrNotContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualOrNotContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualOrNotContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserEQUAL, 0)
}

func (s *EqualOrNotContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserNOTEQUAL, 0)
}

func (s *EqualOrNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterEqualOrNot(s)
	}
}

func (s *EqualOrNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitEqualOrNot(s)
	}
}

func (s *EqualOrNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitEqualOrNot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GScriptParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, GScriptParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserLPAREN, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL, GScriptParserIDENTIFIER:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(95)
			p.Primary()
		}

	case GScriptParserBANG, GScriptParserTILDE:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*NotExprContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GScriptParserBANG || _la == GScriptParserTILDE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*NotExprContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(97)

			var _x = p.expr(6)

			localctx.(*NotExprContext).rhs = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultDivExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(101)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultDivExprContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserMULT || _la == GScriptParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultDivExprContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(102)

					var _x = p.expr(6)

					localctx.(*MultDivExprContext).rhs = _x
				}

			case 2:
				localctx = NewModExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ModExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(104)

					var _m = p.Match(GScriptParserMOD)

					localctx.(*ModExprContext).bop = _m
				}
				{
					p.SetState(105)

					var _x = p.expr(5)

					localctx.(*ModExprContext).rhs = _x
				}

			case 3:
				localctx = NewPlusSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PlusSubExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(107)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PlusSubExprContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserPLUS || _la == GScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PlusSubExprContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(108)

					var _x = p.expr(4)

					localctx.(*PlusSubExprContext).rhs = _x
				}

			case 4:
				localctx = NewGLeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(110)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*GLeContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserGT)|(1<<GScriptParserLT)|(1<<GScriptParserLE)|(1<<GScriptParserGE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*GLeContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(111)
					p.expr(3)
				}

			case 5:
				localctx = NewEqualOrNotContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(113)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualOrNotContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserEQUAL || _la == GScriptParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualOrNotContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(114)
					p.expr(2)
				}

			case 6:
				localctx = NewPostfixExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PostfixExprContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GScriptParserRULE_expr)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(116)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PostfixExprContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GScriptParserINC || _la == GScriptParserDEC) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PostfixExprContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyFrom(ctx *PrimaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprPrimaryContext struct {
	*PrimaryContext
}

func NewExprPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPrimaryContext {
	var p = new(ExprPrimaryContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *ExprPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserLPAREN, 0)
}

func (s *ExprPrimaryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprPrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRPAREN, 0)
}

func (s *ExprPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterExprPrimary(s)
	}
}

func (s *ExprPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitExprPrimary(s)
	}
}

func (s *ExprPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitExprPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiterPrimaryContext struct {
	*PrimaryContext
}

func NewLiterPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiterPrimaryContext {
	var p = new(LiterPrimaryContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *LiterPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiterPrimaryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiterPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterLiterPrimary(s)
	}
}

func (s *LiterPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitLiterPrimary(s)
	}
}

func (s *LiterPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitLiterPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierPrimaryContext struct {
	*PrimaryContext
}

func NewIdentifierPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierPrimaryContext {
	var p = new(IdentifierPrimaryContext)

	p.PrimaryContext = NewEmptyPrimaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryContext))

	return p
}

func (s *IdentifierPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, 0)
}

func (s *IdentifierPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterIdentifierPrimary(s)
	}
}

func (s *IdentifierPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitIdentifierPrimary(s)
	}
}

func (s *IdentifierPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitIdentifierPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GScriptParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserLPAREN:
		localctx = NewExprPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(GScriptParserLPAREN)
		}
		{
			p.SetState(123)
			p.expr(0)
		}
		{
			p.SetState(124)
			p.Match(GScriptParserRPAREN)
		}

	case GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL:
		localctx = NewLiterPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Literal()
		}

	case GScriptParserIDENTIFIER:
		localctx = NewIdentifierPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.Match(GScriptParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnContext struct {
	*StatementContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRETURN, 0)
}

func (s *ReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseContext struct {
	*StatementContext
}

func NewIfElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseContext {
	var p = new(IfElseContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseContext) IF() antlr.TerminalNode {
	return s.GetToken(GScriptParserIF, 0)
}

func (s *IfElseContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *IfElseContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfElseContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GScriptParserELSE, 0)
}

func (s *IfElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterIfElse(s)
	}
}

func (s *IfElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitIfElse(s)
	}
}

func (s *IfElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitIfElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForContext struct {
	*StatementContext
}

func NewForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForContext {
	var p = new(ForContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForContext) FOR() antlr.TerminalNode {
	return s.GetToken(GScriptParserFOR, 0)
}

func (s *ForContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserLPAREN, 0)
}

func (s *ForContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *ForContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRPAREN, 0)
}

func (s *ForContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFor(s)
	}
}

func (s *ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFor(s)
	}
}

func (s *ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmExprContext struct {
	*StatementContext
	statementExpression IExprContext
}

func NewStmExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmExprContext {
	var p = new(StmExprContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *StmExprContext) GetStatementExpression() IExprContext { return s.statementExpression }

func (s *StmExprContext) SetStatementExpression(v IExprContext) { s.statementExpression = v }

func (s *StmExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterStmExpr(s)
	}
}

func (s *StmExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitStmExpr(s)
	}
}

func (s *StmExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitStmExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockLabelContext struct {
	*StatementContext
	blockLabel IBlockContext
}

func NewBlockLabelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockLabelContext {
	var p = new(BlockLabelContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockLabelContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *BlockLabelContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *BlockLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLabelContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBlockLabel(s)
	}
}

func (s *BlockLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBlockLabel(s)
	}
}

func (s *BlockLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBlockLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GScriptParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserLBRACE:
		localctx = NewBlockLabelContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)

			var _x = p.Block()

			localctx.(*BlockLabelContext).blockLabel = _x
		}

	case GScriptParserIF:
		localctx = NewIfElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Match(GScriptParserIF)
		}
		{
			p.SetState(132)
			p.ParExpression()
		}
		{
			p.SetState(133)
			p.Statement()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(134)
				p.Match(GScriptParserELSE)
			}
			{
				p.SetState(135)
				p.Statement()
			}

		}

	case GScriptParserFOR:
		localctx = NewForContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)
			p.Match(GScriptParserFOR)
		}
		{
			p.SetState(139)
			p.Match(GScriptParserLPAREN)
		}
		{
			p.SetState(140)
			p.ForControl()
		}
		{
			p.SetState(141)
			p.Match(GScriptParserRPAREN)
		}
		{
			p.SetState(142)
			p.Statement()
		}

	case GScriptParserRETURN:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
			p.Match(GScriptParserRETURN)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(145)
				p.expr(0)
			}

		}

	case GScriptParserLPAREN, GScriptParserBANG, GScriptParserTILDE, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL, GScriptParserIDENTIFIER:
		localctx = NewStmExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)

			var _x = p.expr(0)

			localctx.(*StmExprContext).statementExpression = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForUpdate returns the forUpdate rule contexts.
	GetForUpdate() IExpressionListContext

	// SetForUpdate sets the forUpdate rule contexts.
	SetForUpdate(IExpressionListContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	forUpdate IExpressionListContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForUpdate() IExpressionListContext { return s.forUpdate }

func (s *ForControlContext) SetForUpdate(v IExpressionListContext) { s.forUpdate = v }

func (s *ForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForControlContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GScriptParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserINT)|(1<<GScriptParserSTRING)|(1<<GScriptParserFLOAT)|(1<<GScriptParserBOOLEAN)|(1<<GScriptParserLPAREN)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GScriptParserFUNCTION-33))|(1<<(GScriptParserBOOL_LITERAL-33))|(1<<(GScriptParserSTRING_LITERAL-33))|(1<<(GScriptParserNULL_LITERAL-33))|(1<<(GScriptParserDECIMAL_LITERAL-33))|(1<<(GScriptParserFLOAT_LITERAL-33))|(1<<(GScriptParserIDENTIFIER-33)))) != 0) {
		{
			p.SetState(151)
			p.ForInit()
		}

	}
	{
		p.SetState(154)
		p.Match(GScriptParserT__0)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserLPAREN)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(GScriptParserBOOL_LITERAL-39))|(1<<(GScriptParserSTRING_LITERAL-39))|(1<<(GScriptParserNULL_LITERAL-39))|(1<<(GScriptParserDECIMAL_LITERAL-39))|(1<<(GScriptParserFLOAT_LITERAL-39))|(1<<(GScriptParserIDENTIFIER-39)))) != 0) {
		{
			p.SetState(155)
			p.expr(0)
		}

	}
	{
		p.SetState(158)
		p.Match(GScriptParserT__0)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserLPAREN)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(GScriptParserBOOL_LITERAL-39))|(1<<(GScriptParserSTRING_LITERAL-39))|(1<<(GScriptParserNULL_LITERAL-39))|(1<<(GScriptParserDECIMAL_LITERAL-39))|(1<<(GScriptParserFLOAT_LITERAL-39))|(1<<(GScriptParserIDENTIFIER-39)))) != 0) {
		{
			p.SetState(159)

			var _x = p.ExpressionList()

			localctx.(*ForControlContext).forUpdate = _x
		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GScriptParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.VariableDeclarators()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.ExpressionList()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionDeclarationContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserLBRACK)
}

func (s *FunctionDeclarationContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserLBRACK, i)
}

func (s *FunctionDeclarationContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserRBRACK)
}

func (s *FunctionDeclarationContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserRBRACK, i)
}

func (s *FunctionDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(GScriptParserTHROWS, 0)
}

func (s *FunctionDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GScriptParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(166)
			p.TypeTypeOrVoid()
		}

	}
	{
		p.SetState(169)
		p.Match(GScriptParserIDENTIFIER)
	}
	{
		p.SetState(170)
		p.FormalParameters()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserLBRACK {
		{
			p.SetState(171)
			p.Match(GScriptParserLBRACK)
		}
		{
			p.SetState(172)
			p.Match(GScriptParserRBRACK)
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GScriptParserTHROWS {
		{
			p.SetState(178)
			p.Match(GScriptParserTHROWS)
		}
		{
			p.SetState(179)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(182)
		p.FunctionBody()
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GScriptParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Block()
	}

	return localctx
}

// IQualifiedNameListContext is an interface to support dynamic dispatch.
type IQualifiedNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameListContext differentiates from other interfaces.
	IsQualifiedNameListContext()
}

type QualifiedNameListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameListContext() *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_qualifiedNameList
	return p
}

func (*QualifiedNameListContext) IsQualifiedNameListContext() {}

func NewQualifiedNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_qualifiedNameList

	return p
}

func (s *QualifiedNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameListContext) AllQualifiedName() []IQualifiedNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem())
	var tst = make([]IQualifiedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualifiedNameContext)
		}
	}

	return tst
}

func (s *QualifiedNameListContext) QualifiedName(i int) IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *QualifiedNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitQualifiedNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) QualifiedNameList() (localctx IQualifiedNameListContext) {
	localctx = NewQualifiedNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GScriptParserRULE_qualifiedNameList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.QualifiedName()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__1 {
		{
			p.SetState(187)
			p.Match(GScriptParserT__1)
		}
		{
			p.SetState(188)
			p.QualifiedName()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserLPAREN, 0)
}

func (s *FormalParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRPAREN, 0)
}

func (s *FormalParametersContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GScriptParserRULE_formalParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(GScriptParserLPAREN)
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserFINAL)|(1<<GScriptParserINT)|(1<<GScriptParserSTRING)|(1<<GScriptParserFLOAT)|(1<<GScriptParserBOOLEAN))) != 0) || _la == GScriptParserFUNCTION || _la == GScriptParserIDENTIFIER {
		{
			p.SetState(195)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(198)
		p.Match(GScriptParserRPAREN)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameter() []IFormalParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem())
	var tst = make([]IFormalParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalParameterContext)
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameter(i int) IFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterContext)
}

func (s *FormalParameterListContext) LastFormalParameter() ILastFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastFormalParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastFormalParameterContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GScriptParserRULE_formalParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.FormalParameter()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(201)
					p.Match(GScriptParserT__1)
				}
				{
					p.SetState(202)
					p.FormalParameter()
				}

			}
			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GScriptParserT__1 {
			{
				p.SetState(208)
				p.Match(GScriptParserT__1)
			}
			{
				p.SetState(209)
				p.LastFormalParameter()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.LastFormalParameter()
		}

	}

	return localctx
}

// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_formalParameter
	return p
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *FormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *FormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *FormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}

func (s *FormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FormalParameter() (localctx IFormalParameterContext) {
	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GScriptParserRULE_formalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserFINAL {
		{
			p.SetState(215)
			p.VariableModifier()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.TypeType()
	}
	{
		p.SetState(222)
		p.VariableDeclaratorId()
	}

	return localctx
}

// ILastFormalParameterContext is an interface to support dynamic dispatch.
type ILastFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastFormalParameterContext differentiates from other interfaces.
	IsLastFormalParameterContext()
}

type LastFormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastFormalParameterContext() *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_lastFormalParameter
	return p
}

func (*LastFormalParameterContext) IsLastFormalParameterContext() {}

func NewLastFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_lastFormalParameter

	return p
}

func (s *LastFormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *LastFormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *LastFormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *LastFormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *LastFormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *LastFormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastFormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastFormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitLastFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) LastFormalParameter() (localctx ILastFormalParameterContext) {
	localctx = NewLastFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GScriptParserRULE_lastFormalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserFINAL {
		{
			p.SetState(224)
			p.VariableModifier()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(230)
		p.TypeType()
	}
	{
		p.SetState(231)
		p.Match(GScriptParserT__2)
	}
	{
		p.SetState(232)
		p.VariableDeclaratorId()
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserIDENTIFIER)
}

func (s *QualifiedNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GScriptParserRULE_qualifiedName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(GScriptParserIDENTIFIER)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__3 {
		{
			p.SetState(235)
			p.Match(GScriptParserT__3)
		}
		{
			p.SetState(236)
			p.Match(GScriptParserIDENTIFIER)
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableModifierContext is an interface to support dynamic dispatch.
type IVariableModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableModifierContext differentiates from other interfaces.
	IsVariableModifierContext()
}

type VariableModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableModifierContext() *VariableModifierContext {
	var p = new(VariableModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableModifier
	return p
}

func (*VariableModifierContext) IsVariableModifierContext() {}

func NewVariableModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableModifierContext {
	var p = new(VariableModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableModifier

	return p
}

func (s *VariableModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableModifierContext) FINAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserFINAL, 0)
}

func (s *VariableModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterVariableModifier(s)
	}
}

func (s *VariableModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitVariableModifier(s)
	}
}

func (s *VariableModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableModifier() (localctx IVariableModifierContext) {
	localctx = NewVariableModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GScriptParserRULE_variableModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(GScriptParserFINAL)
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GScriptParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.TypeType()
	}
	{
		p.SetState(245)
		p.VariableDeclarator()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__1 {
		{
			p.SetState(246)
			p.Match(GScriptParserT__1)
		}
		{
			p.SetState(247)
			p.VariableDeclarator()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GScriptParserASSIGN, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GScriptParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.VariableDeclaratorId()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GScriptParserASSIGN {
		{
			p.SetState(254)
			p.Match(GScriptParserASSIGN)
		}
		{
			p.SetState(255)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserLBRACK)
}

func (s *VariableDeclaratorIdContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserLBRACK, i)
}

func (s *VariableDeclaratorIdContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserRBRACK)
}

func (s *VariableDeclaratorIdContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserRBRACK, i)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GScriptParserRULE_variableDeclaratorId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(GScriptParserIDENTIFIER)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserLBRACK {
		{
			p.SetState(259)
			p.Match(GScriptParserLBRACK)
		}
		{
			p.SetState(260)
			p.Match(GScriptParserRBRACK)
		}

		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GScriptParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.ArrayInitializer()
		}

	case GScriptParserLPAREN, GScriptParserBANG, GScriptParserTILDE, GScriptParserBOOL_LITERAL, GScriptParserSTRING_LITERAL, GScriptParserNULL_LITERAL, GScriptParserDECIMAL_LITERAL, GScriptParserFLOAT_LITERAL, GScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_arrayInitializer
	return p
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GScriptParserLBRACE, 0)
}

func (s *ArrayInitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GScriptParserRBRACE, 0)
}

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem())
	var tst = make([]IVariableInitializerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableInitializerContext)
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GScriptParserRULE_arrayInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(GScriptParserLBRACE)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserLPAREN)|(1<<GScriptParserLBRACE)|(1<<GScriptParserBANG)|(1<<GScriptParserTILDE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(GScriptParserBOOL_LITERAL-39))|(1<<(GScriptParserSTRING_LITERAL-39))|(1<<(GScriptParserNULL_LITERAL-39))|(1<<(GScriptParserDECIMAL_LITERAL-39))|(1<<(GScriptParserFLOAT_LITERAL-39))|(1<<(GScriptParserIDENTIFIER-39)))) != 0) {
		{
			p.SetState(271)
			p.VariableInitializer()
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(272)
					p.Match(GScriptParserT__1)
				}
				{
					p.SetState(273)
					p.VariableInitializer()
				}

			}
			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GScriptParserT__1 {
			{
				p.SetState(279)
				p.Match(GScriptParserT__1)
			}

		}

	}
	{
		p.SetState(284)
		p.Match(GScriptParserRBRACE)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatContext struct {
	*LiteralContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*LiteralContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserNULL_LITERAL, 0)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*LiteralContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserBOOL_LITERAL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*LiteralContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserSTRING_LITERAL, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*LiteralContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserDECIMAL_LITERAL:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Match(GScriptParserDECIMAL_LITERAL)
		}

	case GScriptParserFLOAT_LITERAL:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(GScriptParserFLOAT_LITERAL)
		}

	case GScriptParserSTRING_LITERAL:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.Match(GScriptParserSTRING_LITERAL)
		}

	case GScriptParserBOOL_LITERAL:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(289)
			p.Match(GScriptParserBOOL_LITERAL)
		}

	case GScriptParserNULL_LITERAL:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(290)
			p.Match(GScriptParserNULL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GScriptParserRULE_integerLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(GScriptParserDECIMAL_LITERAL)
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(GScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GScriptParserRULE_floatLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(GScriptParserFLOAT_LITERAL)
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *TypeTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeTypeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserLBRACK)
}

func (s *TypeTypeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserLBRACK, i)
}

func (s *TypeTypeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserRBRACK)
}

func (s *TypeTypeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserRBRACK, i)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterTypeType(s)
	}
}

func (s *TypeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitTypeType(s)
	}
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeType() (localctx ITypeTypeContext) {
	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GScriptParserRULE_typeType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserIDENTIFIER:
		{
			p.SetState(297)
			p.ClassOrInterfaceType()
		}

	case GScriptParserFUNCTION:
		{
			p.SetState(298)
			p.FunctionType()
		}

	case GScriptParserINT, GScriptParserSTRING, GScriptParserFLOAT, GScriptParserBOOLEAN:
		{
			p.SetState(299)
			p.PrimitiveType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserLBRACK {
		{
			p.SetState(302)
			p.Match(GScriptParserLBRACK)
		}
		{
			p.SetState(303)
			p.Match(GScriptParserRBRACK)
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GScriptParserINT, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(GScriptParserSTRING, 0)
}

func (s *PrimitiveTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GScriptParserFLOAT, 0)
}

func (s *PrimitiveTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(GScriptParserBOOLEAN, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GScriptParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserINT)|(1<<GScriptParserSTRING)|(1<<GScriptParserFLOAT)|(1<<GScriptParserBOOLEAN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(GScriptParserFUNCTION, 0)
}

func (s *FunctionTypeContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserLPAREN, 0)
}

func (s *FunctionTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRPAREN, 0)
}

func (s *FunctionTypeContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GScriptParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(GScriptParserFUNCTION)
	}
	{
		p.SetState(312)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(313)
		p.Match(GScriptParserLPAREN)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GScriptParserINT)|(1<<GScriptParserSTRING)|(1<<GScriptParserFLOAT)|(1<<GScriptParserBOOLEAN))) != 0) || _la == GScriptParserFUNCTION || _la == GScriptParserIDENTIFIER {
		{
			p.SetState(314)
			p.TypeList()
		}

	}
	{
		p.SetState(317)
		p.Match(GScriptParserRPAREN)
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypeType() []ITypeTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem())
	var tst = make([]ITypeTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeTypeContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypeType(i int) ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GScriptParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.TypeType()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__1 {
		{
			p.SetState(320)
			p.Match(GScriptParserT__1)
		}
		{
			p.SetState(321)
			p.TypeType()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(GScriptParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GScriptParserRULE_typeTypeOrVoid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GScriptParserINT, GScriptParserSTRING, GScriptParserFLOAT, GScriptParserBOOLEAN, GScriptParserFUNCTION, GScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.TypeType()
		}

	case GScriptParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)
			p.Match(GScriptParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GScriptParserIDENTIFIER)
}

func (s *ClassOrInterfaceTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GScriptParserIDENTIFIER, i)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitClassOrInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GScriptParserRULE_classOrInterfaceType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Match(GScriptParserIDENTIFIER)
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__3 {
		{
			p.SetState(332)
			p.Match(GScriptParserT__3)
		}
		{
			p.SetState(333)
			p.Match(GScriptParserIDENTIFIER)
		}

		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.expr(0)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GScriptParserT__1 {
		{
			p.SetState(340)
			p.Match(GScriptParserT__1)
		}
		{
			p.SetState(341)
			p.expr(0)
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GScriptParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GScriptParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserLPAREN, 0)
}

func (s *ParExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GScriptParserRPAREN, 0)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.EnterParExpression(s)
	}
}

func (s *ParExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GScriptListener); ok {
		listenerT.ExitParExpression(s)
	}
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GScriptVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GScriptParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GScriptParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(GScriptParserLPAREN)
	}
	{
		p.SetState(348)
		p.expr(0)
	}
	{
		p.SetState(349)
		p.Match(GScriptParserRPAREN)
	}

	return localctx
}

func (p *GScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GScriptParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
