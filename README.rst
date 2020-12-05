================================
`Succinct Data Structure`_ Trie_
================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-succinct-data-structure-trie?status.svg
   :target: https://godoc.org/github.com/siongui/go-succinct-data-structure-trie

.. .. image:: https://api.travis-ci.org/siongui/go-succinct-data-structure-trie.svg?branch=master
   :target: https://travis-ci.org/siongui/go-succinct-data-structure-trie

.. image:: https://github.com/siongui/go-succinct-data-structure-trie/workflows/Test%20Package/badge.svg
    :target: https://github.com/siongui/go-succinct-data-structure-trie/blob/master/.github/workflows/build.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-succinct-data-structure-trie
   :target: https://goreportcard.com/report/github.com/siongui/go-succinct-data-structure-trie

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/go-succinct-data-structure-trie/master/UNLICENSE

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/go-succinct-data-structure-trie.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Implementation of `Succinct Trie`_ [1]_ in Go_.

The trie structure is great for fast lookup of dictionary words, but if the
vocabulary of the dictionary is big, it may takes a lot of space to store the
constructed trie. For this reason, succinct data structure is applied to the
trie strcuture and we can both have fast lookup and small space requirement.


Usage
=====

- Basic example: `basic usage <example/basic/usage.go>`__
- Advanced example: `pali dir <example/pali/>`__

UNLICENSE
=========

Released in public domain. See UNLICENSE_.


References
==========

.. [1] `Succinct Data Structures: Cramming 80,000 words into a Javascript file. <http://stevehanov.ca/blog/?id=120>`_
       (`source code <http://www.hanovsolutions.com/trie/Bits.js>`__)

.. [2] Google Search `succinct data structure <https://www.google.com/search?q=succinct+data+structure>`__

.. [3] Google Search `succinct trie <https://www.google.com/search?q=succinct+trie>`__

.. [4] Google Search `golang const array <https://www.google.com/search?q=golang+const+array>`__

.. [5] Google Search `golang function as argument <https://www.google.com/search?q=golang+function+as+argument>`__

.. [6] Google Search `golang charcodeat <https://www.google.com/search?q=golang+charcodeat>`__

       `string - Go lang's equivalent of charCode() method of JavaScript - Stack Overflow <http://stackoverflow.com/questions/31239330/go-langs-equivalent-of-charcode-method-of-javascript>`_

.. [7] `[Golang] Succinct Trie Implementation <https://siongui.github.io/2016/02/08/go-succinct-trie-implementation/>`_

.. [8] `[JavaScript] Bug in Succinct Trie Implementation of Bits.js <https://siongui.github.io/2016/02/02/javascript-bug-in-succinct-trie-implementation-of-bits-js/>`_

.. _Go: https://golang.org/
.. _UNLICENSE: http://unlicense.org/
.. _Succinct Data Structure: https://www.google.com/search?q=Succinct+Data+Structure
.. _Trie: https://www.google.com/search?q=Trie
.. _Succinct Trie: https://www.google.com/search?q=Succinct+Trie
