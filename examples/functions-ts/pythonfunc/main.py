# Copyright 2021, Pulumi Corporation.  All rights reserved.

def handler(request): 
    headers = {
        'Content-Type': 'text/plain'
    }

    return ('Hello, World!', 200, headers)
