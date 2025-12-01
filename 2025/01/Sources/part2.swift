// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
  let path = "input.txt"
  var dial = 50
  var password = 0
  print("The dial starts by pointing at \(dial)")

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return
  }

  contents.enumerateLines { line, _ in
    let op = String(line.prefix(1))
    let readNum = Int(line.dropFirst())!
    print("roate \(op)\(readNum) from \(dial)")

    for _ in 0..<readNum {
      switch op {
      case "L":
        dial -= 1
        if dial < 0 { dial += 100 }
      case "R":
        dial += 1
        if dial >= 100 { dial -= 100 }
      default:
        fatalError("Unknown op \(op)")
      }

      if dial == 0 {
        password += 1
      }
    }

    print("Now pointing at \(dial)")
  }
  print("The password is \(password)")
}
