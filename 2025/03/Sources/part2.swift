// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
  let path = "input.txt"

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  var totalJoltage = 0
  contents.enumerateLines { battery, _ in
    totalJoltage += best12DigitJoltage(from: battery)
  }
  print("Total Joltage \(totalJoltage)")
}

func best12DigitJoltage(from battery: String) -> Int {
  let digits = Array(battery)
  let k = 12
  // If line is shorter than 12, just use all digits
  let keep = min(k, digits.count)
  var toRemove = max(0, digits.count - keep)

  var stack = [Character]()

  for d in digits {
    while toRemove > 0, let last = stack.last, last < d {
      stack.removeLast()
      toRemove -= 1
    }
    stack.append(d)
  }

  // If we still have more than keep digits, trim from the end
  if stack.count > keep {
    stack.removeLast(stack.count - keep)
  }

  // Convert the final 12-digit sequence to Int
  let resultString = String(stack)
  return Int(resultString)!
}
