package main

import "fmt"

func main()  {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	//两个for
	for i:=0; i<5; i++ {
		for j:=0; j<10; j++ {
			println(j)
		}
	}

	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	//基于条件判断
	var i int = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)  //0 0 0 0 0
		v = 5
	}

	//标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	/*
	i is: 0, and j is: 1
i is: 0, and j is: 2
i is: 0, and j is: 3
i is: 1, and j is: 0
i is: 1, and j is: 1
i is: 1, and j is: 2
i is: 1, and j is: 3
i is: 2, and j is: 0
i is: 2, and j is: 1
i is: 2, and j is: 2
i is: 2, and j is: 3
i is: 3, and j is: 0
i is: 3, and j is: 1
i is: 3, and j is: 2
i is: 3, and j is: 3
i is: 4, and j is: 0
i is: 4, and j is: 1
i is: 4, and j is: 2
i is: 4, and j is: 3
i is: 5, and j is: 0
i is: 5, and j is: 1
i is: 5, and j is: 2
i is: 5, and j is: 3
	 */
}
/*
0
This is the 0 iteration
1
This is the 1 iteration
2
This is the 2 iteration
3
This is the 3 iteration
4
This is the 4 iteration
5
The length of str is: 27
6
Character on position 0 is: G
7
8
Character on position 1 is: o
9
Character on position 2 is:
0
Character on position 3 is: i
1
2
Character on position 4 is: s
Character on position 5 is:
3
Character on position 6 is: a
4
Character on position 7 is:
5
Character on position 8 is: b
6
Character on position 9 is: e
7
Character on position 10 is: a
8
Character on position 11 is: u
9
Character on position 12 is: t
0
Character on position 13 is: i
1
Character on position 14 is: f
2
Character on position 15 is: u
3
Character on position 16 is: l
4
Character on position 17 is:
5
Character on position 18 is: l
6
Character on position 19 is: a
7
Character on position 20 is: n
8
Character on position 21 is: g
9
Character on position 22 is: u
0
Character on position 23 is: a
1
Character on position 24 is: g
2
Character on position 25 is: e
3
Character on position 26 is: !
4
The length of str2 is: 9
5
Character on position 0 is: æ
6
Character on position 1 is: 
7
Character on position 2 is: ¥
8
Character on position 3 is: æ
9
Character on position 4 is: 
0
Character on position 5 is: ¬
1
Character on position 6 is: è
2
Character on position 7 is: ª
3
Character on position 8 is: 
4
5
6
7
8
9
 */