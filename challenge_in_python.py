# ## Assume you have sent an HTTP GET Request to an API. The API responds with a JSON payload.
# ## Your job is to determine if that payload matches a given pattern.
# ## The JSON payload is a map with keys of type string and values of various types.
# ## Values can be of type string, float, array, or object.
# ##
# ## The pattern matches the JSON payload if all of the following are true:
# ## - for every field defined in the pattern, the payload also has that field
# ## - The field is equal to the value in the given pattern
# ##

matches = {
    "type": "circle",
    "payload": {
        "x": 150.0,
        "y": 150.0,
        "r": 100.0,
        "text": "Hello circle!"
    }
}

doesntMatch = {
    "type": "circle",
    "payload": {
        "x1": 150.0,
        "y1": 150.0,
        "r": 100.0,
        "text": "Hello bob!"
    }
}

doesntMatch2 = {
    "type": "circle",
    "payload": {
        "x1": 150.0,
        "y1": 150.0,
        "text": "Hello circle!"
    }
}


squareMatches = {
    "type": "square",
    "payload": {
        "metadata": {
            "text": "silly",
            "colors": ["green", "blue"],
            "friends": {
                "bob": False,
                "sally": True
            }
        },
        "coordinates": [1.0,2.0,3.0,4.0]
    }
}

squareDoesntMatch = {
    "type": "square",
    "payload": {
        "y1": 123.0,
        "x2": 123.0,
        "y2": 123.0
    },
    "bob": ["extra data"]
}

squareDoesntMatch2 = {
    "type": "square",
    "payload": {
        "metadata": {
            "text": "silly",
            "colors": ["green", "blue"],
            "friends": {
                "bob": True,
                "sally": False
            }
        },
        "coordinates": [1.0,2.0,3.0,4.0]
    }
}

squareMatches2 = {
    "type": "square",
    "payload": {
        "metadata": {
            "text": "silly",
            "colors": ["green", "blue"],
            "friends": {
                "bob": True,
                "sally": False,
                "red": "blue"
            }
        },
        "coordinates": [1.0,2.0,3.0,4.0]
    }
}

patterns = {
    "circle": {
        "x": 150.0,
        "y": 150.0,
        "r": 100.0,
        "text": "Hello circle!",
    },
    "square": {
        "coordinates": [1.0, 2.0, 3.0, 4.0],
        "metadata": {
            "text":   "silly",
            "colors": ["green", "blue"],
            "friends": {
                "sally": True,
                "bob":   False,
            },
        },
    },
}

def CheckMatches(patterns, matches):
    patternType = matches['type']
    if patternType in patterns:
        for p in matches['payload']:
            if isinstance(patterns[patternType][p],dict):
                matches = CheckMatches(patterns[patternType][p], matches['payload'][p]:
                if not matches:
                    return False

            if p in patterns[patternType] and matches['payload'][p] == patterns[patternType][p]:
                print(p, "ok")
            else:
                return False
        return True
    else:
        return False



assert(CheckMatches(patterns, matches) == True)
assert(CheckMatches(patterns, doesntMatch) == False)
assert(CheckMatches(patterns, doesntMatch2) == False)
assert(CheckMatches(patterns, squareMatches) == True)
assert(CheckMatches(patterns, squareDoesntMatch) == False)
assert(CheckMatches(patterns, squareDoesntMatch2) == False)
assert(CheckMatches(patterns, squareMatches2) == True)
