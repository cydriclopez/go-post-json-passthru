# go-post-json-passthru

## Go POST JSON passthru controller

> ***This tutorial requires some knowledge in Linux, Angular, and Go Programming Language.***

### Table of Contents
1. Introduction
2. Why passthru controller
3. Prerequisites
4. Clone this repo
5. Client-side Angular code
6. Server-side Go code
7. Conclusion

### 1. Introduction

Oftentimes in our Angular web app we need to push JSON data to the server. Working with a tree data-structure can be a bit tricky. In this tutorial we will work with a tree component and send its JSON data to the server.

This tutorial builds from the previous tutorials. It is good to go thru them in sequence especially if you are new to Angular, Postgresql, and the Go programming language.

In the current iteration of this code, the tree component data is obtained from a static file [src/client/src/assets/data/files.json](https://github.com/cydriclopez/go-post-json-passthru/blob/main/src/client/src/assets/data/files.json). Our goal is to push this data to a Go controller which, for now, merely prints the JSON data.

In the next tutorial we will call a Postgresql stored-function that walks-thru this JSON tree data and saves it into records in a table. In this tutorial we will cover just the very basics of JSON data processing using the tree component JSON data.

This is now our ***Tree demo*** app. Note that we have enabled the ***Save*** button. This ***Save*** button sends the tree JSON data to our server-side Go controller that for now just prints it on the console.<br/>
<img src="images/primeng-tree-demo2.png" width="650"/>


### 2. Why passthru controller

The Go standard library has excellent JSON processing features. However oftentimes the final destination of our JSON data is the database. Postgresql since version 9.2 has had the JSON data type. In the current version 14.2 Postgresql JSON features have improved a lot. We can just have Postgresql validate and save our JSON data.

It can be redundant to work with JSON in Go and then push it to Postgresql for more JSON processing. Sometimes this might be necessary. However, in this tutorial we will send our JSON data thru a Go controller which, for now, will merely grab the JSON data and print it in the server-side console. This is in preparation for the next tutorial where we will focus on the Postgresql code to walk-thru and pull apart the JSON data and save them as individual records in a table.

Postgresql JSON processing can infer the row parent-child relationships from the JSON structure and use the row id accordingly. We will see this code in the next tutorial. For now, we will follow the JSON data from the Angular component into the service which, using the HttpClient, calls the Go server-side controller.

### 3. Prerequisites

As mentioned before, this tutorial builds from the previous tutorials. I suggest you go thru them in sequence especially if you are new to Angular and the Go programming language.

I assume that you have a working Angular and Go installation. Please checkout the previous tutorials that cover these topics.

### 4. Clone this repo
### 5. Client-side Angular code
### 6. Server-side Go code
### 7. Conclusion

Please pardon my mess. Work in progress! 😊
