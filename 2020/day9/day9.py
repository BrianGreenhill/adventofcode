import fileinput
import itertools

p1 = 0
i = 0
target = 556543474

lines = list([int(x) for x in fileinput.input()])
p1 = None
for i in range (25, len(lines)):
    Ok = True
    prev = lines[i-25:i]
    for y, z in itertools.combinations(prev, 2):
        if y + z == lines[i]:
            Ok = False
    if Ok and p1 is None:
        p1 = lines[i]

p2 = None
for i in range(len(lines)):
    for j in range(i+1, len(lines)):
        xs = lines[i:j]
        if sum(xs) == target and p2 is None:
            p2 = min(xs) + max(xs)

print("Part 1: {}".format(p1))
print("Part 2: {}".format(p2))
