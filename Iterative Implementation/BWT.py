#simple iterative barrows wheeler transform written in python

def bwt(s):
    s += '$'
    m = []
    for _ in s:
        s = s[-1] + s[0:len(s) - 1]
        m.append(s)
    m.sort()
    t = ""
    for i in m:
        t += i[-1]
    return t



# hello world$
# $hello world
# d$hello worl
# ld$hello wor
# rld$hello wo
# orld$hello w
# world$hello
#  world$hello
# o world$hell
# lo world$hel
# llo world$he
# ello world$h


def bwtRestore(s):
    m = [""] * len(s)
    for i in range(len(s)):
        m[i] += s[i]

    for i in range(len(s) - 1):
        m.sort()
        m = [s[i] + m[i] for i in range(len(s))]

    print(m)

    for i in m:
        if i[-1] =="$":
            return i


def main():
    user = input("Enter character string to compress: ")
    transform = bwt(user)
    print("Burrows-Wheeler Transform: " + str(transform))

    restored = bwtRestore(transform)
    print("Restored: " + str(restored))

    user += "$"
    print("Are original and restored the same? " + str(user == restored))


if __name__ == "__main__":
    main()
