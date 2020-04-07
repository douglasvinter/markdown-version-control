# Simple REST markdown versioning tool using Go & Gin

Created a simple REST API with Go and Gin for learning purposes after a 3 days learning go language and 1 day learning the gin framework

####Creating a markdown repository

**POST** 

localhost:8080/api/markdown/web-hook/

**body**

	{ "commitId": "1928376198", "buildNumber": "12rc", "payload": "##MARKDOWN 2" }

**header response**
Will return a header with key "Location" and value POST URI + UUID

	{"key":"Location","value":": /api/markdown/web-hook/d4dacd42-1c05-4921-b166-180ec4a11b88"}

####List all markdowns

**GET**

localhost:8080/api/markdown/web-hook/

**response**

	{
    "data": [
        {
            "projectId": "d4dacd42-1c05-4921-b166-180ec4a11b88",
            "markdown": {
                "commitId": "1928376198",
                "buildNumber": "12rc",
                "payload": "##MARKDOWN 2"
            }
        },
        {
            "projectId": "c84bcdf7-f6ea-48c2-8ae2-b3cabca1662f",
            "markdown": {
                "commitId": "1928376198",
                "buildNumber": "12rc",
                "payload": "##MARKDOWN 4"
            }
        }
    ]
}

####List aal markdowns for a given managed project

**GET**

localhost:8080/api/markdown/web-hook/d4dacd42-1c05-4921-b166-180ec4a11b88

**response**

	{
    "data": {
        "commitId": "1928376198",
        "buildNumber": "12rc",
        "payload": "##MARKDOWN 2"
    }
}

####Updating a markdown

**PUT**

localhost:8080/api/markdown/web-hook/d4dacd42-1c05-4921-b166-180ec4a11b88

**body**

	{ "commitId": "192837619823719823", "buildNumber": "1273rc", "payload": "##MARKDOWN TEXT UPDATED 2" }

**response**

 {
    "data": {
        "commitId": "192837619823719823",
        "buildNumber": "1273rc",
        "payload": "##MARKDOWN TEXT UPDATED 2"
    }
}


#### Deleting a markdown

**DELETE**

localhost:8080/api/markdown/web-hook/d4dacd42-1c05-4921-b166-180ec4a11b88