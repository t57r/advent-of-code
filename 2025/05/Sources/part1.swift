// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"
  let (ranges, ingredientIds) = readData(from: path)

  var count = 0
  for id in ingredientIds {
    if isInAnyRange(id, ranges: ranges) {
      count += 1
    }
  }

  print("Count \(count)")

}

func isInAnyRange(_ x: Int, ranges: [ClosedRange<Int>]) -> Bool {
  ranges.contains { $0.contains(x) }
}

func readData(from path: String) -> ([ClosedRange<Int>], [Int]) {
  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return ([], [])
  }

  var ranges = [ClosedRange<Int>]()
  var ingredientIds = [Int]()
  var rangesFinished = false

  contents.enumerateLines { line, _ in
    if line.count == 0 {
      rangesFinished = true
      return
    }

    if rangesFinished {
      ingredientIds.append(Int(line)!)
    } else {
      let minMax = line.split(separator: "-")
      let min = Int(minMax[0])!
      let max = Int(minMax[1])!
      ranges.append(min...max)
    }
  }

  return (ranges, ingredientIds)
}
