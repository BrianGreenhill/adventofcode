import fileinput

def is_nice(line):
    for s in ['ab', 'cd', 'pq', 'xy']:
        if s in line:
            return False
    vowels = (line.count('a') + line.count('e') + line.count('i')
            + line.count('o') + line.count('u'))
    if vowels < 3:
        return False
    if any([line[i] == line[i+1] for i in range(len(line)-1)]):
        return True
    return False

def is_nice_rev(string):
    if not any([string[i] == string[i+2] for i in range(len(string)-2)]):
        return False
    if any ([(string.count(string[i:i+2])>=2) for i in range(len(string)-2)]):
        return True
    return False


NICE = set()
NICE_REV = set()
lines = list(fileinput.input())
for line in lines:
    if is_nice(line):
        NICE.add(line)
    if is_nice_rev(line):
        NICE_REV.add(line)

p1 = len(NICE)
p2 = len(NICE_REV)
print(p1)
print(p2)
