// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"
  var dial = 50
  var password = 0
  print("The dial starts by pointing at \(dial)")
  if let contents = try? String(contentsOfFile: path) {
    contents.enumerateLines { line, _ in
      let op = String(line.prefix(1))
      let num = Int(line.dropFirst())! % 100
      var newDial = 0
      switch op {
      case "L":
        newDial = dial - num
        if newDial < 0 {
          newDial = 100 + newDial
        }
      case "R":
        newDial = (dial + num) % 100
      default:
        print("Unknown operator \(op)")
      }
      dial = newDial
      if dial == 0 {
        password += 1
      }
      print("The dial is rotated \(line) to point at \(dial)")
    }
    print("The password is \(password)")
  }
}
