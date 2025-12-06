// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
  let path = "input.txt"
  var grid = readGrid(from: path)

  var totalAcceptableRolls = 0

  while true {
    var acceptableRolls = 0
    var removableCoords: [(Int, Int)] = []
    for i in 0..<grid.count {
      for j in 0..<grid[i].count {
        let count = countRolls(i: i, j: j, grid: grid)
        if grid[i][j] == "@" && count < 4 {
          removableCoords.append((i, j))
          acceptableRolls += 1
        }
      }
    }

    if acceptableRolls == 0 {
      break
    }
    totalAcceptableRolls += acceptableRolls

    for (i, j) in removableCoords {
      grid[i][j] = "."
    }
  }

  print("Total acceptable rolls: \(totalAcceptableRolls)")
}
