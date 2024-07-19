#include <iostream>
#include <string>
#include <vector>

using namespace std;

string long2flag(long flag) {
    long num1 = flag >> 26;
    long num2 = flag >> 6 & 0x3ffff;
    char symbols[3];
    int index = 0;
    for (int i = 0; i < 6; i += 2) {
        symbols[index] = static_cast<char>(1 << 6 | (flag >> i & 0x3));
        index++;
    }
    return "2GIS.CTF{" + std::to_string(num1) + "-" + symbols[2] + "-" + symbols[1] + "-" + symbols[0] + "-" + std::to_string(num2) + "}";
}

int main() {
    std::cout << long2flag(3207649565579002) << std::endl;
    return 0;
}
