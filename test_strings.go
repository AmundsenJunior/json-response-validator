package main

var matches = []string{
	`{
  "type": "circle",
  "payload": {
    "x": 150.0,
    "y": 150.0,
    "r": 100.0,
    "text": "Hello circle!"
  }
}`,
	`{
  "type": "square",
  "payload": {
    "metadata": {
      "text": "silly",
      "colors": ["green", "blue"],
      "friends": {
        "bob": false,
        "sally": true
      }
    },
    "coordinates": [1.0,2.0,3.0,4.0]
  }
}`,
}

var nonmatches = []string{
	`{
  "type": "circle",
  "payload": {
    "x1": 150.0,
    "y1": 150.0,
    "r": 100.0,
    "text": "Hello bob!"
  }
}`,
	`{
  "type": "circle",
  "payload": {
    "x1": 150.0,
    "y1": 150.0,
    "text": "Hello circle!"
  }
}`,
	`{
  "type": "square",
  "payload": {
    "y1": 123.0,
    "x2": 123.0,
    "y2": 123.0
  },
  "bob": ["extra data"]
}`,
	`{
  "type": "square",
  "payload": {
    "metadata": {
      "text": "silly",
      "colors": ["green", "blue"],
      "friends": {
        "bob": true,
        "sally": false,
        "red": "blue"
      }
    },
    "coordinates": [1.0,2.0,3.0,4.0]
  }
}`,
	`{
  "type": "square",
  "payload": {
    "metadata": {
      "text": "silly",
      "colors": ["green", "blue"],
      "friends": {
        "bob": true,
        "sally": false
      }
    },
    "coordinates": [1.0,2.0,3.0,4.0]
  }
}`,
}
