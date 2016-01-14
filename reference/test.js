function TestBitWriter() {
  var bw = new BitWriter();
  bw.write(3, 2);
  console.log(bw.getDebugString(3));
  console.log(bw.getData());
  bw.write(0, 3);
  console.log(bw.getData());
  bw.write(2, 2);
  console.log(bw.getData());
  console.log(bw.getDebugString(3));
}

function TestBitString() {
  var bs = new BitString("88kj5w_6phb");
  console.log(bs);
  console.log(bs.rank(5));
  console.log(bs.rank(24));
  console.log(bs.rank(37));
  console.log(bs.rank(55));
  console.log(bs.rank(65));
  console.log(bs.get(5,7));
  console.log(bs.get(7,13));
  console.log(bs.get(0,5));
  console.log(bs.get(3,3));
  console.log(bs.get(33,17));
  console.log(bs.count(0,17));
  console.log(bs.count(7,2));
  console.log(bs.count(56,9));
  console.log(bs.count(12,1));
  console.log(bs.count(5,7));
}

function TestRankDirectory() {
  var rd = RankDirectory.Create("1wnc2bxhbx7mkbgnpwq7vtlub7p6pkls42lvie9j1ekcpt0zytrdl67enescolwex7aumq4imywstrpktbvxy0rp61nnonj9grdf", 400, L1, L2);
  console.log(rd);
  console.log(rd.rank(1, 200))
  console.log(rd.rank(0, 100))
  console.log(rd.select(1, 134))
  console.log(rd.select(0, 77))
}

function TestTrie() {
  var te = new Trie();
  te.insert("apple");
  te.insert("orange");
  te.insert("alphapha");
  te.insert("lamp");
  te.insert("hello");
  te.insert("jello");
  te.insert("quiz");
  var teData = te.encode();
  console.log(teData);
  console.log(te.getNodeCount());
  var rd = RankDirectory.Create(teData, te.getNodeCount() * 2 + 1, L1, L2);
  console.log(rd.getData());

  var ftrie = new FrozenTrie( teData, rd.getData(), te.getNodeCount());
  console.log(ftrie.lookup("alphapha"))
}

TestTrie();
