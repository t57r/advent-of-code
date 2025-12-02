// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
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
      let localSum = sumInvalidsInRange2(min: min, max: max)
      sum += localSum
      print("\(min)-\(max) has \(localSum) local sum invalid IDs")
    }
    print("Total sum: \(sum)")
  }
}

func sumInvalidsInRange2(min: Int, max: Int) -> Int {
  var sum = 0
  for num in min...max {
    if isRepeatedPattern2(String(num)) {
      sum += num
    }
  }
  return sum
}

func isRepeatedPattern2(_ s: String) -> Bool {
  let n = s.count
  if n < 2 { return false }  // need at least 2 chars to be a repetition

  let chars = Array(s)  // easier indexing

  // Try all possible pattern lengths
  for len in 1...(n / 2) {
    if n % len != 0 { continue }  // pattern must divide length

    // pattern: first `len` characters
    let pattern = chars[0..<len]
    var valid = true

    // Check each chunk of length `len`
    var index = len
    while index < n {
      let segment = chars[index..<index + len]
      if segment != pattern {
        valid = false
        break
      }
      index += len
    }

    if valid {
      return true
    }
  }

  return false
}
