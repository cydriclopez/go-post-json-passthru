# go-post-json-passthru

## Go POST JSON passthru controller

> ***This tutorial requires some knowledge in Linux, Docker, Angular, and Go Programming Language.***

### Table of Contents
1. Introduction
2. Why passthru controller
3. Prerequisites
4. Clone this repo and run "npm install"
5. Compile Go server-side code
6. Client-side Angular code
7. Server-side Go code
8. Conclusion

### 1. Introduction

Oftentimes in our Angular web app we need to push JSON data to the server. Working with a tree data-structure can be a bit tricky. In this tutorial we will work with a tree component and send its JSON data to the server.

This tutorial builds on the previous tutorials. It is good to go thru the previous tutorials in sequence especially if you are new to Angular, Postgresql, and the Go programming language.

In the current iteration of this code, the tree component gets its data from a static file [src/client/src/assets/data/files.json](https://github.com/cydriclopez/go-post-json-passthru/blob/main/src/client/src/assets/data/files.json). Our goal, in this tutorial, is to push this data to a Go controller which, for now, merely prints the JSON data.

In the next tutorial we will call a Postgresql stored-function that walks-thru this JSON tree data and saves it into records in a table. In this tutorial we will cover just the very basics of JSON data processing using the tree component JSON data.

This is now our ***Tree demo*** app. Note that we have enabled the ***Save*** button. This ***Save*** button sends the tree JSON data to our server-side Go controller that for now just prints it on the console.<br/>
<kbd><img src="images/primeng-tree-demo2.png" width="650"/></kbd>


### 2. Why passthru controller

The Go standard library has excellent JSON processing features. However oftentimes the final destination of our JSON data is the database. Postgresql since version 9.2 has had the JSON data type. In the current version 14.2 Postgresql JSON features have improved a lot. We can just have Postgresql validate and save our JSON data.

It can be redundant to work with JSON in Go and then push it to Postgresql for more JSON processing. Sometimes this might be necessary. However, in this tutorial we will send our JSON data thru a Go controller which, for now, will merely grab the JSON data and print it in the server-side console. This is in preparation for the next tutorial where we will focus on the Postgresql code to walk-thru and pull apart the JSON data and save them as individual records in a table.

Postgresql JSON processing can infer the row parent-child relationships from the JSON structure and use the row id accordingly. We will see this code in the next tutorial. For now, we will follow the JSON data from the Angular component into the service which, using the HttpClient, calls the Go server-side controller.

### 3. Prerequisites

As mentioned before, this tutorial builds on the previous tutorials. I suggest you go thru them in sequence especially if you are new to Angular and the Go programming language.

I assume that you have a [working Angular](https://github.com/cydriclopez/docker-ng-dev) and [Go installation](https://github.com/cydriclopez/go-static-server#3-install-go). Please checkout the previous tutorials that cover these topics.

### 4. Clone this repo and run "npm install"

#### 4.1. Clone this repo then change into the repo folder.

You can follow the commands below. You may have to adjust according to your own chosen directory structure.

```bash
user1@penguin:~/Projects$
:git clone https://github.com/cydriclopez/go-post-json-passthru.git

# Change folder into the Angular client app
user1@penguin:~/Projects$
:cd go-post-json-passthru/src/client

# Use the pwd output in your ~/.bashrc angular alias
user1@penguin:~/Projects/go-post-json-passthru/src/client$
:pwd
/home/node/ng/go-post-json-passthru/src/client
```

#### 4.2. Modify your ***~/.bashrc*** file

Note the result of the preceding ***pwd***, print working directory, command. You will have to alter the ***alias angular*** command in your ***~/.bashrc*** accordingly by adding another ***-v*** volume mapping using the ***pwd*** path. <ins>**Substitute your own path here if necessary.**</ins> However try maintain the container mapping into the ***:/home/node/ng/go-post-json-passthru*** folder.

```bash
# Setup Docker Angular working environment
alias angular='docker run -it --rm \
-p 4200:4200 -p 9876:9876 \
-v /home/user1/Projects/ng:/home/node/ng \
-v /home/user1/Projects/go-post-json-passthru/src/client\
:/home/node/ng/go-post-json-passthru \
-w /home/node/ng angular /bin/sh'
```

Reload the ***~/.bashrc*** file using the command: ***source ~/.bashrc***.

```bash
:source ~/.bashrc
```

#### 4.3. Run the ***Angular*** alias

Follow the commands below to run the ***angular*** alias. The prompt should change accordingly to notify you that you are in the Angular docker container.

```bash
user1@penguin:~/Projects/go-post-json-passthru$
:pwd
/home/user1/Projects/go-post-json-passthru

user1@penguin:~/Projects/go-post-json-passthru$
:angular

# Here you are now in the Angular docker container
/home/node/ng # ls -l
drwxr-xr-x    1 node     node           272 Aug  7 04:38 go-post-json-passthru
drwxr-xr-x    1 node     node           248 Jul 28 23:41 treemodule-json

/home/node/ng # cd go-post-json-passthru

/home/node/ng/go-post-json-passthru #
```

#### 4.4. Run "npm install"

Run ***npm install*** to install Angular and all requirements in ***node_modules*** folder.

```bash
/home/node/ng/go-post-json-passthru # npm install
...
[truncated Angular messages]
...
/home/node/ng/go-post-json-passthru #

```

#### 4.5. Run "ng build --watch"

Run "ng build --watch" to generate JavaScript static code in folder ***dist/project_name***. This is the folder we will host in our ***webserv*** Go server-side app.

```bash
/home/node/ng/go-post-json-passthru # ng build --watch

✔ Browser application bundle generation complete.
✔ Index html generation complete.

Initial Chunk Files | Names         |      Size
vendor.js           | vendor        |   3.61 MB
styles.css          | styles        | 241.78 kB
polyfills.js        | polyfills     | 216.88 kB
main.js             | main          |  60.28 kB
runtime.js          | runtime       |   6.40 kB

                    | Initial Total |   4.12 MB

Build at: 2022-08-10T21:07:52.958Z - Hash: ad769f193a142bd6 - Time: 9360ms
```

At this point Angular has "compiled" our client-side web app. The static files that Angular generated into the folder ***src/client/dist/primeng-quickstart-cli*** is ready for serving by our Go server-side app.

### 5. Compile and run Go server code

#### 5.1. Change directory into the Go server code

Here we will compile our Go server-side app. Here we will need to open another terminal tab where we can run our Go compiler. You can follow the steps here to [install the Go language compiler](https://github.com/cydriclopez/go-static-server#3-install-go).

With an installed Go compiler we can change directory into the Go server-side code. Follow the steps below.

```bash
user1@penguin:~/Projects/github/go-post-json-passthru$
:cd src/server

# List files in the folder
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:ll

-rw-r--r-- 1 user1 user1   24 Aug  6 12:22 go.mod
drwxr-xr-x 1 user1 user1   18 Aug  2 13:24 params
-rw-r--r-- 1 user1 user1   24 Aug  9 01:56 README.md
drwxr-xr-x 1 user1 user1   22 Jul 31 12:16 treedata
-rw-r--r-- 1 user1 user1 1254 Aug  9 22:14 webserv.go

# Compile Go source code in the current folder
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:go install

# Locate where is our executable
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:which webserv
/home/user1/go/bin/webserv

user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:
```

#### 5.2. Run our Go server app

The above command ***go install*** read our ***webserv*** web server app, compiled it, then generated the executable in the folder ***~/go/bin***.

Make sure that the folder ***~/go/bin*** is in your path as [instructed here](https://github.com/cydriclopez/go-static-server#34-update-your-path). This is the default path where the Go compiler saves the executables generated from compiling your source code.

It used to be that Go required the setting of the ***GOPATH*** environment variable to function properly. This is not anymore the case with the new Go module system. The file ***go.mod*** tags the folder as a module which is a collection of related packages/folders.

Here is running our ***webserv*** app without parameters.

```bash
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:webserv

2022/08/10 20:23:36
Simple static server of Angular compiled dist/project folder.
Run "ng build --watch" then in another terminal
use dist/project folder as parameter for this utility.
Usage:
webserv STATIC_FOLDER_TO_SERVE [PORT]
Default port: 3000
Examples:
webserv .
webserv ~/Projects/ng/ultima12/dist/ultima
webserv ~/Projects/ng/ultima12/dist/ultima 4000

user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:
```

Here we run ***webserv*** and we feed it the relative folder ***../client/dist/primeng-quickstart-cli*** which is the location of the Angular compiled static files. The default port is :3000.

```bash
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:webserv ../client/dist/primeng-quickstart-cli

2022/08/10 19:06:07
Serving static folder: ../client/dist/primeng-quickstart-cli
Listening on port: :3000
Press Ctrl-C to stop server
```

As mentioned earlier, this is now our ***Tree demo*** app. Note that we have enabled the ***Save*** button. This ***Save*** button sends the tree JSON data to our server-side Go controller that for now just prints it on the console.<br/>
<kbd><img src="images/primeng-tree-demo2.png" width="650"/></kbd>

### 6. Client-side Angular code

A key part of connecting the Angular client-side data, to the Go server-side controller, is for the data-structures to match. We are passing JSON as string from the Angular client to the Go server controller.

#### 6.1. Data-structure match

The TypeScript class ***NodeService*** is defined in file ***src/client/src/app/services/nodeservice.ts***. In this file we define the Angular interface to match the Go server-side struct definition (i.e. the following table).

Note that we are not passing JSON structures from the client to the Go controller. We are passing JSON as string from the Angular client to our Go server controller. Our goal is for the Go controller to be a "passthru" controller which will call a Postgresql stored-function to validate and process our JSON data.

#### Table 6.1. Data-structure match between Angular and Go
|    | Angular interface | Go struct |
| ----------- | --- | ----------- |
|   |export interface JsonData {|type JsonData struct {|
|   |&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;data:   string;|&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Data string \`json:"data"\`|
|   |{    |{    |

#### 6.2. TypeScript function postJsonString()

The TypeScript function ***postJsonString()*** below is quite self-explanatory. Before we can send our JSON data from the Primeng tree component we must JSON.stringify with a replacer array parameter to flatten the JSON to prevent circular references.

```typescript
postJsonString() {
    // First recourse thru JSON data to save toexpand = expanded.
    // The field expanded is a Primeng tree component switch that
    // we cannot persist to the db. We need the toexpand field
    // for that purpose. It's why we extended TreeNode with:
    //    export interface TreeNode2 extends TreeNode { }
    this.treeNodes.forEach(node => {
        this.saveToexpand(node);
    });

    // The JSON.stringify replacer array parameter will flatten the json
    // to prevent circular references. Fields 'key', 'parent', and 'leaf'
    // are filtered-out and inferred from the json structure. PostgreSQL
    // can traverse thru the json and fill-in these fields from the structure.
    const json = JSON.stringify(this.treeNodes,
        [
            'label', 'icon', 'expandedIcon', 'collapsedIcon',
            'data', 'children', 'toexpand', 'group_id'
        ]
    );

    const jsonData: JsonData = { data: json }
    this.http.post<any>('/api/postjsonstring', jsonData, httpOptions)
        .subscribe();
}
```

### 7. Server-side Go code

#### 7.1. Go server app in 3 packages

The Go server-side code is really basic. We have refactored the [previous tutorial's code](https://github.com/cydriclopez/go-static-server/blob/main/src/server/stic.go) into 3 packages.

| package   | file | purpose |
| ----------- | --- | ----------- |
| main | [src/server/webserv.go](https://github.com/cydriclopez/go-post-json-passthru/blob/main/src/server/webserv.go) | main "webserv" executable  |
| params | [src/server/params/params.go](https://github.com/cydriclopez/go-post-json-passthru/blob/main/src/server/params/params.go) | process the command-line args |
| treedata | [src/server/treedata/treedata.go](https://github.com/cydriclopez/go-post-json-passthru/blob/main/src/server/treedata/treedata.go) | process the tree JSON data |

#### 7.2. Method saveJsonData() to save data

Here instead of saving our data into Postgresql we merely print it. In the next tutorial we will focus on the Postgresql stored-function to pick-apart the JSON data and save them as records.

```go
// Save json data to db
func (t *tData) saveJsonData() {
	// For now we will just print the data from the client
	log.Println("jsonData:", t.Jdata)
}
```

#### 7.3. Our app in action

As mentioned earlier, this is now our ***Tree demo*** app. Note that we have enabled the ***Save*** button. This ***Save*** button sends the tree JSON data to our server-side Go controller that for now just prints it on the console.

Try click on the ***Save*** button and notice in our Go server app console the JSON data comes across.
<br/>
<kbd><img src="images/primeng-tree-demo2.png" width="650"/></kbd>

Below is our Go server app in action. When you click on the ***Save*** button the JSON data is printed on the console. In the next tutorial we will focus on the Postgresql stored-function to pick-apart the JSON data and save them as records.

```bash
user1@penguin:~/Projects/github/go-post-json-passthru/src/server$
:webserv ../client/dist/primeng-quickstart-cli

2022/08/10 22:31:37
Serving static folder: ../client/dist/primeng-quickstart-cli
Listening on port: :3000
Press Ctrl-C to stop server
```
<code>
2022/08/10 22:31:49 jsonData: {[{"label":"Documents","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Documents Folder","children":[{"label":"Work","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Work Folder","children":[{"label":"Expenses.doc","icon":"pi pi-file","data":"Expenses Document"},{"label":"Resume.doc","icon":"pi pi-file","data":"Resume Document"}]},{"label":"Home","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Home Folder","children":[{"label":"Invoices.txt","icon":"pi pi-file","data":"Invoices for this month"}]}],"toexpand":true},{"label":"Pictures","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Pictures Folder","children":[{"label":"barcelona.jpg","icon":"pi pi-image","data":"Barcelona Photo"},{"label":"logo.jpg","icon":"pi pi-image","data":"PrimeFaces Logo"},{"label":"primeui.png","icon":"pi pi-image","data":"PrimeUI Logo"}],"toexpand":true},{"label":"Movies","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Movies Folder","children":[{"label":"Al Pacino","data":"Pacino Movies","children":[{"label":"Scarface","icon":"pi pi-video","data":"Scarface Movie"},{"label":"Serpico","icon":"pi pi-video","data":"Serpico Movie"}]},{"label":"Robert De Niro","data":"De Niro Movies","children":[{"label":"Goodfellas","icon":"pi pi-video","data":"Goodfellas Movie"},{"label":"Untouchables","icon":"pi pi-video","data":"Untouchables Movie"}]}]}]}
</code>

### 8. Conclusion

In this tutorial we have covered the following.

1. Introduction
2. Why passthru controller
3. Prerequisites
4. Clone this repo and run "npm install"
5. Compile Go server-side code
6. Client-side Angular code
7. Server-side Go code
8. Conclusion

As you can see it is quite simple to create a passthru Go controller for our Angular JSON data. The next tutorial is quite fun. Postgresql is a fun database to work with. It is where we will pick-apart the JSON data and save them as records in Postgresql.

I know storing JSON is en vogue these days, with NoSQL databases, but the classic usual records in tables format still has its place. I like to push and grab data from Postgresql in JSON format but still storing them in plain records in tables format.

Happy coding! 😊

---
