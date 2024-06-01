#include <stdio.h>

int main() {
	printf("一个汉字占%ld字节\n", sizeof("x")); // 这里字符串包含一个汉字3字节, 加一个结束符\0一字节

        char arr[] = {"你"};  	
	printf("中文占%ld字节\n", sizeof(arr));
	return 0;
}
