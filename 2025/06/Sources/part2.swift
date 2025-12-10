// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

struct ColumnProblem {
  let op: Character
  let operands: [Int]
}

func parseProblems(from lines: [String], linesCount: Int) -> [ColumnProblem] {
  precondition(
    lines.count == linesCount, "Expected \(linesCount) (\(linesCount - 1) numbers + 1 ops)")

  // Convert each line to [Character] for easy indexing
  let numberRows = lines[0..<(linesCount - 1)].map { Array($0) }
  let opRow = Array(lines[linesCount - 1])

  let rowCount = numberRows.count
  let width = opRow.count

  // 1. Find all operator positions
  var opPositions: [(index: Int, op: Character)] = []
  for (idx, ch) in opRow.enumerated() where ch == "+" || ch == "*" {
    opPositions.append((idx, ch))
  }

  var problems: [ColumnProblem] = []

  // 2. For each operator, define its block of columns and collect operands
  for (k, opInfo) in opPositions.enumerated() {
    let start = opInfo.index
    let end = (k + 1 < opPositions.count) ? opPositions[k + 1].index - 1 : width - 1

    var operands: [Int] = []

    // 3. Walk columns in this block from right to left
    for col in stride(from: end, through: start, by: -1) {
      var digits: [Character] = []

      // 4. Collect digits top-to-bottom in this column
      for row in 0..<rowCount {
        let ch = numberRows[row][col]
        if ch.isNumber {
          digits.append(ch)
        }
      }

      // If at least one digit, from a number
      if !digits.isEmpty {
        let value = Int(String(digits))!
        operands.append(value)
      }
    }

    problems.append(ColumnProblem(op: opInfo.op, operands: operands))
  }

  return problems
}

func part2() {
  let path = "input.txt"

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  var lines: [String] = []
  contents.enumerateLines { line, _ in
    lines.append(line)
  }

  let problems = parseProblems(from: lines, linesCount: lines.count)

  var numbers: [[Int]] = []
  var operations: [Character] = []

  // Print them in the format from the description:
  for (idx, p) in problems.enumerated().reversed() {  // reversed so we describe from rightmost
    let expr = p.operands
      .map(String.init)
      .joined(separator: " \(p.op) ")
    print("Problem \(idx): \(expr)")

    numbers.append(p.operands)
    operations.append(p.op)
  }

  print("Numbers \(numbers)")
  print("Operations \(operations)")

  calcSum2(numbers: numbers, operations: operations)
}

func calcSum2(numbers: [[Int]], operations: [Character]) {
  var sum = 0
  for i in 0..<numbers.count {
    let isMultiply = operations[i] == "*"
    var rowTotal = isMultiply ? 1 : 0
    for j in 0..<numbers[i].count {
      if isMultiply {
        rowTotal *= numbers[i][j]
      } else {
        rowTotal += numbers[i][j]
      }
    }
    sum += rowTotal
  }

  print("Sum \(sum)")
}
