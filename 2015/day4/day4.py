import fileinput
import hashlib

def mine_for_santa(input, leading):
    zeros = '0'*leading
    for i in range(10000000):
        msg = input + str(i)
        if hashlib.md5(msg).hexdigest()[:leading] == str(zeros):
            return i
            break
    return None

p1 = None
p2 = None
input = 'iwrupvqb'

p1 = mine_for_santa(input, 5)
p2 = mine_for_santa(input, 6)
print(p1)
print(p2)
