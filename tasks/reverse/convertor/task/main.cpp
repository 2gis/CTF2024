#include <iostream>
#include <string>
#include <vector>

using namespace std;

long flag2long(const string& flag) {
    string flag_content = flag.substr(5, flag.length() - 1);
    vector<string> parts;
    size_t pos = 0;
    while ((pos = flag_content.find('-')) != string::npos) {
        parts.push_back(flag_content.substr(0, pos));
        flag_content.erase(0, pos + 1);
    }
    parts.push_back(flag_content);

    long num1 = stol(parts[0]);
    char symbols[3] = {parts[1][0], parts[2][0], parts[3][0]};
    long num2 = stol(parts[4]);

    long nums = (num1 << 20) | num2;
    long result = 0L;
    int index = 0;
    for(long res = nums << 2; index <= 2; res <<= 2){
        res |= symbols[index] & 0x3;
        index++;
        if(index == 3) result = res;
    }
    return result;
}

int main() {
    string flag;
    std::cin >> flag;
    std::cout << flag2long(flag) << std::endl;
    return 0;
}