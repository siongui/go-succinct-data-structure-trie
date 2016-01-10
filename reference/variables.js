/**
 * Path of files and directories
 */

function getDictWordsJsonDir() {
  return require('path').resolve(__dirname, 
      '../../pylib/paliwords');
}

function getBitsjsPath() {
  return require('path').resolve(__dirname, 'Bits.js');
}

function getSuccinctTrieJsonPath() {
  return require('path').resolve(__dirname, 
      '../../pylib/json/succinct_trie.json');
}

if ( typeof exports !== "undefined" ) {
  exports.dictWordsJsonDir = getDictWordsJsonDir();
  exports.BitsjsPath = getBitsjsPath();
  exports.succinctTrieJsonPath = getSuccinctTrieJsonPath();
}
