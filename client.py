#!/usr/bin/env python3

import json
import requests


def rpc(method, **args):
    RPC = 'http://localhost:8888/api/v1/rpc'
    r = requests.post(RPC, json=dict(jsonrpc="2.0", method=method,
                                     params=args, id=1))
    return r.json()['result']


def main():
    uid = rpc("nmap.scan",
              hosts=["factory.sh", "blog.garambrogne.net"],
              ports=[22, 80, 443])
    n = 0
    while True:
        result = rpc("longrun.next", id=uid, n=n)
        stop = False
        for r in result:
            print(json.dumps(r, indent=2))
            stop |= r['state'] in ['success', 'canceled', 'error']
        if stop:
            break
        n += len(result)


if __name__ == "__main__":
    main()
