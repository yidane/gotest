/*
http://blog.csdn.net/david_jett/article/details/44750687

有个直角三角形三边为A,B,C三个整数。已知C为最长边。现告诉你A的长度，求一组B,C，使得B和C最接近

a^2+b^2=c^2

a^2=(c+b)(c-b)    -- c-b=1
a^2=2b+1
b=(a^2-1)/2
b=(a+1)(a-1)/2         -- 当a为奇数时,b为整数,c=b+1为整数

a^2=(c+b)(c-b)         -- c-b=2
a^2=(b+2+b)*2
a^2=4b+4
b=(a^2-4)/4
b=(a/2+1)(a/2-1)       -- 当a为偶数时，b为整数，c=b+2为整数

结论:
在边长a已知，且c为斜边情况下。只需要判断a是奇数还是偶数，即可求出b和c的值。
1、当a为奇数时，c和b最小差值为1
	b=(a^2-1)/2
	c=b+1
2、当a为偶数时，c和b最小差值为2
	b=(a^2-4)/4
	c=b+2
*/

package main

func rightTriangle(a int) (b, c int) {
	switch a % 2 {
	case 0:
		b = (a*a - 4) / 4
		c = b + 2
	case 1:
		b = (a*a - 1) / 2
		c = b + 1
	}

	if b <= 0 {
		c = 0
		b = 0
	}

	return
}
