/**
 * @Author: LiuShuXin
 * @Description:
 * @File:  CalculateWrapper.c
 * Software: Goland
 * @Date: 2025/2/13 10:04
 */

#include <windows.h>

typedef int (*AddFunc)(int, int);

int Add(int a, int b) {
    HINSTANCE hDLL = LoadLibrary("Calculate.dll");
    if (hDLL != NULL) {
        AddFunc add = (AddFunc)GetProcAddress(hDLL, "Add");
        if (add != NULL) {
            return add(a, b);
        }
        FreeLibrary(hDLL);
    }
    return -1; // 错误处理
}
