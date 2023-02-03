# Simple Rest Api

## Technologies

Project created with:
* [gorilla/mux](https://github.com/gorilla/mux)
* [sqlite3](https://github.com/mattn/go-sqlite3)
* [json-iterator](https://github.com/json-iterator/go)

## Endpoints

### POST /imports

Imports elements.
* `id` of each element is unique
* `parentId` can be null
* `size` > 0
* len(`info`) ≤ 255
* in one request cannot be two elements with the same `id`
* no cyclic dependencies
```
{
  "items": [
    {
      "id": "element_1_1",
      "info": "exercitation tempor non",
      "parentId": null,
      "size": 123554
    },
    {
      "id": "element_1_2",
      "info": "nostrud laboris ea exercitation",
      "parentId": "element_1_1",
      "size": 41724467
    },
    {
      "id": "element_1_3",
      "info": "nostradamus",
      "parentId": null,
      "size": 4273784
    }
  ]
}
```
### DELETE /delete/{id}

Deletes element with given id and all of its children.

### GET /nodes/{id}

Returns element with given id and all of its children.
```
{
    "id": "element_1_1",
    "info": "exercitation tempor non",
    "parentId": "",
    "size": 123554,
    "children": [
        {
            "id": "element_1_2",
            "info": "nostrud laboris ea exercitation",
            "parentId": "element_1_1",
            "size": 41724467,
            "children": null
        }
    ]
}
```
