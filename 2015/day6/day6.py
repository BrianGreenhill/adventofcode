import fileinput

p1 = 0
p2 = 0

lines = list(fileinput.input())
lines.append('')
for line in lines:
    line = line.strip()
    if line:
        words = line.split()
        if words[0] == "turn":
            action, start, end = words[1], words[2], words[4]
            print(action, start, end)
        if words[0] == "toggle":
            start, end = words[1], words[3]
            print(start, end)

print(p1)
print(p2)
