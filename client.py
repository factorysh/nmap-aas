#!/usr/bin/env python3

import requests


def main():
    RPC = 'http://localhost:8888/api/v1/rpc'
    r = requests.post(RPC,
                      json=dict(jsonrpc="2.0",
                                method="nmap.scan",
                                params=dict(hosts=["bearstech.com"],
                                            ports=[22, 80, 443]),
                                id=1))
    uid = r.json()['result']
    n = 0
    while True:
        r = requests.post(RPC,
                        json=dict(jsonrpc="2.0",
                                    method="longrun.next",
                                    params=dict(id=uid, n=n),
                                    id=1))
        result = r.json()['result']
        stop = False
        for r in result:
            print(r)
            stop |= r['state'] in ['success', 'canceled', 'error']
        if stop:
            break
        n += len(result)


if __name__ == "__main__":
    main()
