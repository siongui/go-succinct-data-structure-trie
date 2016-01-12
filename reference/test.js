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

TestBitString()
