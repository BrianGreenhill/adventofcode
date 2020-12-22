import fileinput

floor = 0
p1 = None
p2 = None
parens = list(fileinput.input())
for paren in parens:
    if paren:
        P = list(paren)
        for p in range(len(P)):
            if floor == -1 and p2 is None:
                p2 = p
            if P[p] == '(':
                floor += 1
            if P[p] == ')':
                floor -= 1
p1 = floor
print("Part 1: " + str(p1))
print("Part 2: " + str(p2))
