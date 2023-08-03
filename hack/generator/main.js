'use strict';

const fs = require('fs');
const csvSync = require('csv-parse/lib/sync');
const semver = require('semver')

const file = __dirname + '/../input.csv';
let data = fs.readFileSync(file);

let res = csvSync(data, {
        skip_empty_lines: true,
});

console.log("// Code generated by Node.js. DO NOT EDIT.\n");
console.log("package npm\n");
console.log("var (");
console.log("	autoGeneratedTests = []struct {");
console.log("		version    string");
console.log("		constraint string");
console.log("		want       bool");
console.log("	}{");

res.forEach((element, index, array) => {
	let res = semver.satisfies(element[1], element[0])
        console.log(`		{"${element[1]}", "${element[0]}", ${res}},`)
})

console.log("	}");
console.log(")");
