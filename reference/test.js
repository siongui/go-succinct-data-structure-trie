var bw = new BitWriter();
bw.write(3, 2);
console.log(bw.getDebugString(3))
console.log(bw.getData())
bw.write(0, 3);
console.log(bw.getData())
bw.write(2, 2);
console.log(bw.getData())
console.log(bw.getDebugString(3))
