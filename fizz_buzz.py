def fb():
    divisor_to_message = {}
    div = [3, 5, 7, 9]
    msg = ["fizz", "buzz", "boff", "kek"]

    for a,b in zip(div, msg):
        divisor_to_message[a] = b

    for i in range(100):
        res = ""
        for div, msg in divisor_to_message.items():
            if i % div == 0:
                res += msg
        if res == "":
            print(i)
        elif res != "":
            print(res)

fb()