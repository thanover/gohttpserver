# github.com/thanover/gohttpserver

## Overview

This is a small project that I created to facilitate learning go and as well as http. It is by no means meant to be a complete http server solution and probably has no practical purpose beside being a learning aid.

## Goals

Create a HTTP server that can:

1. handle ANY HTTP request, even if it just responds with IDK
2. able to serve resources for a limited list, including:
   1. **GET /** --- html page with not much in it. It will then need to handle subsequent request for anything specified as needed in that HTML page (e.g. css, js, etc.)
   2. **POST /submitquery** -- basic script that returns some message given a input

## Rules

1. I can only use the built-in go libraries (besides http of course), as well as any other modules I've published myself
2. I can't copy/paste code from other place, besides simple one liners
3. I can use AI, but only to explain things and NOT to generate code

## Steps

1. completed a [Go fundamentals](https://www.pluralsight.com/courses/go-fundamentals) course on Plural Sight
2. read through [A Tour of Go](https://go.dev/tour/list)
3. read over the [RFC 9110 HTTP Semantics](https://www.rfc-editor.org/rfc/rfc9110). I didn't read everything line-by-line, but I did read the first three sections and then skimmed the rest
4. read over this document: [HTTP Made Really Easy](https://www.jmarshall.com/easy/http/)
5. started coding...

⚠️ Side Quest: While I started to code, I decided I wanted to log what's happening to the console as well as to a file. I figured this is a fun way to learn some more basics of go. I decided to create and publish a module to do this: [https://github.com/thanover/gologger](https://github.com/thanover/gologger)
