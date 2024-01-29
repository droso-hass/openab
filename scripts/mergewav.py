import os
import sys

try:
    os.remove(sys.argv[2])
except:
    pass

cmd = "ffmpeg -y"

n = 0
p = sys.argv[1]
for i in sorted(os.listdir(p)):
    cmd += f" -i {os.path.join(p,i)}"
    n += 1

cmd += " -filter_complex '"

for i in range(n):
    cmd += f"[{i}:0]"
cmd += f"concat=n={n}:v=0:a=1[out]' -map [out] {sys.argv[2]}"

print(cmd)
os.system(cmd)
