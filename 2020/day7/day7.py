import fileinput
from collections import deque, defaultdict

p1 = 0
p2 = 0

PARENTS = defaultdict(list)
CONTAINS = defaultdict(list)
target = 'shinygoldbag'

lines = list(fileinput.input())
lines.append('')
for line in lines:
    line = line.strip()
    if line:
        words = line.split()
        container = words[0]+words[1]+words[2]
        container = container[:-1]
        if words[-3] == "no":
            continue
        else:
            idx = 4
            while idx < len(words):
                bag = words[idx] + words[idx+1] + words[idx+2] + words[idx+3]
                if bag.endswith(','):
                    bag = bag[:-1]
                if bag.endswith('.'):
                    bag = bag[:-1]
                if bag.endswith('s'):
                    bag = bag[:-1]
                n = int(bag[0])
                assert bag[1] not in '0123456789'
                while any([bag.startswith(d) for d in '0123456789']):
                    bag = bag[1:]
                PARENTS[bag].append(container)
                CONTAINS[container].append((n, bag))
                idx+=4
SEEN = set()
Q = deque([target])
while Q:
    x = Q.popleft()
    if x in SEEN:
        continue
    SEEN.add(x)
    for y in PARENTS[x]:
        Q.append(y)

p1 = len(SEEN) - 1

def size(bag):
    ans = 1
    for (n, y) in CONTAINS[bag]:
        ans += n * size(y)
    return ans
p2 = size(target)-1

print(p1)
print(p2)
