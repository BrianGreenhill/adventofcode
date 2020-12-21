import fileinput

p1 = 0
p2 = 0

lines = list(fileinput.input())
lines.append('')
for line in lines:
    line = line.strip()
    if line:
        words = line.split()
        print(words)

print(p1)
print(p2)
