from colour import Color

max_color = Color("#00FF00")
steps = 20
t = "100"

ref = max_color.hsl
colors = []

def fmt(x):
    if len(x) == 4:
        return x[1]+x[1]+x[2]+x[2]+x[3]+x[3]
    return x[1:]

for i in range(steps,0,-1):
    colors.append(fmt(Color(
        hsl=(
            ref[0], ref[1], ref[2]/steps*i
        )
    ).get_hex()))

cmd = "03;4;0;"
for i in colors:
    cmd += i+";"+t+";"
cmd += "000000"+";"+t+";"
for i in reversed(colors[1:]):
    cmd += i+";"+t+";"

print(cmd[:-1])

