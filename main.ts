/**
 * @author Joshua Adeyemi
 * @version 1.0.0
 * @date 2025-12-14
 * @fileOverview This program calculates the average of student marks and gives performance feedback.
 */

// Hardcoded input for portfolio / linter purposes
const numMarks: number = 3;
const marks: number[] = [75.8, 90.2, 80];

let sum: number = 0;

for (let i = 0; i < numMarks; i++) {
  sum += marks[i];
}

const average: number = sum / numMarks;

console.log(`You have entered ${numMarks} marks. The student's average is ${average.toFixed(0)}%.`);

if (average <= 49) {
  console.log("The student is failing.");
} else if (average <= 69) {
  console.log("The student's performance is just under average.");
} else if (average <= 79) {
  console.log("The student's performance is average.");
} else {
  console.log("The student is on the honour roll.");
}

console.log("\nDone.");
