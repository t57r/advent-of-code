// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  var totalJoltage = 0
  contents.enumerateLines { battery, _ in
    let volts = Array(battery)
    var maxJoltage = 0
    for i in 0..<volts.count {
      for j in i + 1..<volts.count {
        let tens = volts[i].wholeNumberValue!
        let ones = volts[j].wholeNumberValue!
        let joltage = calcJoltage(tens: tens, ones: ones)
        if joltage > maxJoltage {
          maxJoltage = joltage
        }
      }
    }
    totalJoltage += maxJoltage
  }
  print("Total Joltage \(totalJoltage)")
}

func calcJoltage(tens: Int, ones: Int) -> Int {
  return 10 * tens + ones
}
