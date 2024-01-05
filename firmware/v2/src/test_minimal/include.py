#!/sbin/python

# very basic preprocessor
# only supports one level of ifdef because I'm lazy

import os
import sys

included = []
defines = []
program = ""

SYMBOL_INCLUDE = 0
SYMBOL_IFDEF = 1
SYMBOL_ENDIF = 2
SYMBOL_DEFINE = 3

def get_symbol(line):
    l = len(line)
    if l > 9 and line[0:8] == "#include":
        return (SYMBOL_INCLUDE, line[9:].strip("\n"))
    elif l > 7 and line[0:6] == "#ifdef":
        return (SYMBOL_IFDEF, line[7:].strip("\n"))
    elif l >= 6 and line[0:6] == "#endif":
        return (SYMBOL_ENDIF, None)
    elif l > 8 and line[0:7] == "#define":
        return (SYMBOL_DEFINE, line[8:].strip("\n"))
    return (None, None)

def process(path):
    print(f"processing {path}")
    global included, defines, program
    ifdef_current = False
    ifdef_ok = False

    program += f"\n\n//------------ BEGIN {path} ------------\n"
    with open(path, "r") as f:
        data = f.readlines()
    for i in data:
        s, d = get_symbol(i)
        if s == SYMBOL_INCLUDE and d not in included:
            included.append(d)
            dir = os.path.dirname(path)
            process(os.path.join(dir, d))
        elif s == SYMBOL_IFDEF:
            ifdef_current = True
            ifdef_ok = d in defines
        elif s == SYMBOL_ENDIF:
            ifdef_current = False
        elif s == SYMBOL_DEFINE:
            defines.append(d)
        else:
            if not ifdef_current or ifdef_ok:
                program += i
    program +=  f"\n//------------ END {path} ------------\n"

if __name__ == "__main__":
    process(sys.argv[1])
    with open(sys.argv[2], "w+") as f:
        f.write(program)

