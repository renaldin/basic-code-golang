function countDivisible(A, B, K) {
    return Math.floor(B / K) - Math.floor((A - 1) / K);
}

const input = `2
1
10
3
8
20
4`;

const lines = input.split('\n');
let T = parseInt(lines[0]);
let lineIndex = 1;

for (let caseNum = 1; caseNum <= T; caseNum++) {
    const A = parseInt(lines[lineIndex++]);
    const B = parseInt(lines[lineIndex++]);
    const K = parseInt(lines[lineIndex++]);

    const result = countDivisible(A, B, K);
    console.log(`Case ${caseNum}: ${result}`);
}