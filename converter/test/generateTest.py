import random
import sys


def f(n):
    i = random.getrandbits(n)
    print(i, hex(i)[2:])


original_stdout = sys.stdout
with open('tests.txt', 'w') as file:
    sys.stdout = file
    for i in range(12, 259, 3):
        f(i)
sys.stdout = original_stdout
