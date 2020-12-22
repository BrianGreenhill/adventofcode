import fileinput

p1 = 0
p2 = 0
sqft = 0
r = 0

lines = list(fileinput.input())
lines.append('')
for line in lines:
    line = line.strip()
    if line:
        measurements = line.split('x')
        l = int(measurements[0])
        w = int(measurements[1])
        h = int(measurements[2])
        s1 = l*w
        s2 = w*h
        s3 = h*l

        sa = 2*s1+2*s2+2*s3
        wp = sa + min(s1, s2, s3)
        sqft += wp

        p = min(l+l+w+w,l+l+h+h,w+w+h+h)
        b = l*w*h
        r += p+b

p1 = sqft
p2 = r
print("Part 1: {}".format(p1))
print("Part 2: {}".format(p2))
