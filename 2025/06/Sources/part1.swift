// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  var numbers: [[Int]] = []
  contents.enumerateLines { line, _ in
    let isFirstCharNumber = line.first(where: { !$0.isWhitespace })?.isNumber == true
    if isFirstCharNumber {
      let lineNumbers =
        line
        .split(whereSeparator: \.isWhitespace)
        .compactMap { Int($0) }
      numbers.append(lineNumbers)
    } else {
      // it's a line of operations
      let operations: [Character] = line.filter { !$0.isWhitespace }
      calcSum(numbers: numbers, operations: operations)
    }
  }

}

func calcSum(numbers: [[Int]], operations: [Character]) {
  var sum = 0
  let columns = numbers[0].count
  for columnIndex in 0..<columns {
    let isMultiply = operations[columnIndex] == "*"
    var columnTotal = isMultiply ? 1 : 0
    for rowIndex in 0..<numbers.count {
      if isMultiply {
        columnTotal *= numbers[rowIndex][columnIndex]
      } else {
        columnTotal += numbers[rowIndex][columnIndex]
      }
    }
    sum += columnTotal
  }

  print("Sum \(sum)")
}
