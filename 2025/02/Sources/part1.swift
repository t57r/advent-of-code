// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  contents.enumerateLines { line, _ in
    let ranges = line.split(separator: ",")
    var sum = 0
    for range in ranges {
      let minMax = range.split(separator: "-")
      let min = Int(minMax[0])!
      let max = Int(minMax[1])!
      let localSum = sumInvalidsInRange(min: min, max: max)
      sum += localSum
      print("\(min)-\(max) has \(localSum) local sum invalid IDs")
    }
    print("Total sum: \(sum)")
  }
}

func sumInvalidsInRange(min: Int, max: Int) -> Int {
  var sum = 0
  for num in min...max {
    if isRepeatedPattern(String(num)) {
      sum += num
    }
  }
  return sum
}

func isRepeatedPattern(_ s: String) -> Bool {
  guard s.allSatisfy({ $0.isNumber }) else { return false }

  guard s.count >= 2, s.count % 2 == 0 else { return false }

  let midIndex = s.index(s.startIndex, offsetBy: s.count / 2)
  let firstHalf = s[..<midIndex]
  let secondHalf = s[midIndex...]

  return firstHalf == secondHalf
}
