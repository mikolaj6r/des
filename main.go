package main
import ("strconv";
        "fmt" )



func permutate(src string, array []int) string {
    new_block := make([]rune, len(array));
    for index, element := range array{
        new_block[index] = rune(src[element - 1]);
    }
    return string(new_block);
}

func string_xor( a string , b string) string {    
    var c string;
    
    for index, element := range b {
        left := int8(a[index] - '0');
        right := int8(element - '0');
        temp:= strconv.Itoa(int( left ^ right));
        c += (temp);
    }
    return string(c);
}                      

func bits4(num uint64) string {
    bits := strconv.FormatUint(num, 2);
    for i:= len([]rune(bits)); i < 4; i++ {
         bits = "0" + bits;
    }
    return bits;
}

func main(){

    var startPermutation = []int{
                    58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20, 12, 4,
                    62, 54, 46, 38, 30, 22, 14, 6, 64, 56, 48, 40, 32, 24, 16, 8,
                    57, 49, 41, 33, 25, 17, 9,  1, 59, 51, 43, 35, 27, 19, 11, 3,
                    61, 53, 45, 37, 29, 21, 13, 5, 63, 55, 47, 39, 31, 23, 15, 7};

    var  endPermutation = []int{ 
                    40, 8, 48, 16, 56, 24, 64, 32, 39, 7, 47, 15, 55, 23, 63, 31,
                    38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29,
                    36, 4, 44, 12, 52, 20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27,
                    34, 2, 42, 10, 50, 18, 58, 26, 33, 1, 41, 9,  49, 17, 57, 25};

    var  keyPermutation = []int{
                    57, 49, 41, 33, 25, 17,  9,  1, 58, 50, 42, 34, 26, 18,
                    10,  2, 59, 51, 43, 35, 27, 19, 11,  3, 60, 52, 44, 36,
                    63, 55, 47, 39, 31, 23, 15,  7, 62, 54, 46, 38, 30, 22,
                    14,  6, 61, 53, 45, 37, 29, 21, 13,  5, 28, 20, 12,  4};

    var  roundShift = []int{1,1,2,2,2,2,2,2,1,2,2,2,2,2,2,1};

    var  compressPermutation= []int{
                    14, 17, 11, 24,  1,  5,  3, 28, 15,  6 ,21, 10,
                    23, 19, 12,  4, 26,  8, 16,  7, 27, 20, 13,  2,
                    41, 52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48,
                    44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32};

    var  extendPermutation= []int{
            32,  1,  2,  3,  4,  5,  4,  5,  6,  7,  8,  9,
            8,  9, 10, 11, 12, 13, 12, 13, 14, 15, 16, 17,
            16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25,
            24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32,  1};

    var  P_box= []int{
            16, 7, 20, 21, 29, 12, 28, 17,  1, 15, 23, 26,  5, 18, 31, 10,
            2, 8, 24, 14, 32, 27,  3,  9, 19, 13, 30,  6, 22, 11,  4, 25};

    var  S_box= [][][]int{
            {
                []int{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
                []int{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
                []int{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
                []int{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},},{
                []int{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
                []int{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
                []int{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
                []int{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},},{
                []int{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
                []int{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
                []int{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
                []int{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},},{
                []int{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
                []int{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
                []int{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
                []int{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},},{
                []int{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
                []int{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
                []int{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
                []int{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},},{
                []int{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
                []int{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
                []int{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
                []int{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},},{
                []int{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
                []int{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
                []int{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
                []int{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},},{
                []int{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
                []int{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
                []int{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
                []int{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11}}};


    // keyString must come from generator
    key := "0000111000110010100100100011001011101010011011010000110101110011";
    //64 bitowy blok
    input := "1000011110000111100001111000011110000111100001111000011110000111";  

    // generate underkeys
    // permutation KP
    keysKP := permutate(key, keyPermutation);
    //split keys
    keysLeft := keysKP[:28];
    keysRight := keysKP[28:];
    
    // input permutation
    inputP := permutate(input, startPermutation);
    // split message
    inputLeft := inputP[:32];
    inputRight := inputP[32:];
    
    for i:=0; i<16; i++{
        // shifting half keys
        for j := 0; j < roundShift[i]; j++ {
           keysLeft = keysLeft[1:28] + string(keysLeft[0]); 
           keysRight = keysRight[1:28] + string(keysRight[0]);
        }  
        // permutation with compression CP: 56-> 48 bits
        keys := permutate(keysLeft + keysRight, compressPermutation);
        // extend permutation EP: 32 -> 48 bits
        inputRightExt := permutate(inputRight, extendPermutation);
        // xor with 48 bits of underkey
        inputRightExt = string_xor(inputRightExt, keys);
        var f_result string;
        
        // send to 8 S-box to get 32 bits -> each sbox gives 4 bits
        for j:=0; j<8; j++{
            // 6 bits of input
            f_input := inputRightExt[j*6:j*6+6];
            // f-function
            row, _ := strconv.ParseInt(string(f_input[0]) + string(f_input[5]), 2, 64);
            col, _ := strconv.ParseInt(string(f_input[1:5]), 2, 64); 
            
            var temp uint64 = uint64(S_box[j][row][col]);
            f_result += bits4(temp)//4 bits
        }
        f_result = permutate(f_result, P_box);
        newRight := string_xor(inputLeft, f_result);
        inputLeft = inputRight;
        inputRight = newRight;
    }

    output :=  inputRight + inputLeft;
    encrypted := permutate(output, endPermutation);
    
    fmt.Println("input:  ", input, "\nkey:      ", keyString, "\nencrypted:", encrypted, "\n");

}