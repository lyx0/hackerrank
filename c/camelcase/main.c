#include <stdio.h>

char input[] = "oneTwoThreeFour";

int main() {
    int found = 0;

    for (int i = 0; i <= sizeof(input); i++) {
        if ((int)input[i] <= 96) {
            found++;
        }
    }
    printf("There were %d words in the input string.", found-1);
}
