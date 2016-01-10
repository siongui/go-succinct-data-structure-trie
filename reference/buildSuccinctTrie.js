/**
 * Bits.js: A Succinct Trie for Javascript By Steve Hanov
 * @see http://stevehanov.ca/blog/index.php?id=120
 * @see http://www.hanovsolutions.com/trie/Bits.js
 * @see also https://github.com/jeresig/trie-js
 * @see also http://stevehanov.ca/blog/index.php?id=115
 * @see also http://ejohn.org/blog/javascript-trie-performance-analysis/
 * @see also http://ejohn.org/blog/revised-javascript-dictionary-search/
 */
var bitsjs = require(require('./variables.js').BitsjsPath);

var words = require("fs").readdirSync(require('./variables.js').dictWordsJsonDir);
words.sort();

// create a trie
var trie = new bitsjs.Trie();

for (var i=0; i < words.length; i++) {
  console.log('( ' + i +  ' ) inserting ' + words[i] + ' to trie ...');
  trie.insert(words[i]);
}

console.log('Trie insertion completed. Start to encode trie ...');
console.log('Encoding may take a long time. Please wait patiently ...');

// encode the trie
var trieData = trie.encode();

console.log('Trie encoding completed!');

// Encode the rank directory
var directory = bitsjs.CreateRankDirectory( trieData, trie.getNodeCount() * 2 + 1);

var output;
    output = '{\n    "nodeCount": ' + trie.getNodeCount() + ",\n";
    output += '    "directory": "' + directory.getData() + '",\n';
    output += '    "trie": "' + trieData + '"\n';
    output += "}\n";

console.log(output);

var jsonData = eval( '(' + output + ")" );

require('fs').writeFileSync(
    require('./variables.js').succinctTrieJsonPath,
    JSON.stringify(jsonData)
);
