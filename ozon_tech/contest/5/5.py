import json


def isHackFiles(json):
    if "files" not in json:
        return False
    for filename in json["files"]:
        if filename[-5:] == ".hack":
            return True
    return False


def allNestedFiles(json):
    n = 0
    # print(json["dir"], end="")
    if "files" in json:
        n = len(json["files"])
        # print(json["files"], len(json["files"]))
    if "folders" not in json:
        return n
    for folder in json["folders"]:
        n += allNestedFiles(folder)
    return n


def numberOfHackFiles(json):
    n = 0
    if isHackFiles(json):
        n = allNestedFiles(json)
        return n
    if "folders" not in json:
        return n
    for folder in json["folders"]:
        n += numberOfHackFiles(folder)
    return n


if __name__ == "__main__":
    import sys

    def fast_input():
        return sys.stdin.readline().rstrip("\r\n")

    def fast_output(x):
        sys.stdout.write(f"{x}\n")

    n = int(fast_input())
    for i in range(n):
        v = int(fast_input())
        input_data = []
        for j in range(v):
            input_data.append(fast_input())
        # print(input_data)
        data = json.loads("".join(input_data))
        # print(data)
        fast_output(numberOfHackFiles(data))
        # print(allNestedFiles(data))
