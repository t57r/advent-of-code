// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"
  let grid = readGrid(from: path)

  var acceptableRolls = 0
  for i in 0..<grid.count {
    for j in 0..<grid[i].count {
      let count = countRolls(i: i, j: j, grid: grid)
      if grid[i][j] == "@" && count < 4 {
        acceptableRolls += 1
      }
    }
  }

  print("Acceptable rolls: \(acceptableRolls)")
}

func countRolls(i: Int, j: Int, grid: [[Character]]) -> Int {
  let n = grid.count
  var count = 0
  count += i - 1 >= 0 && j - 1 >= 0 && grid[i - 1][j - 1] == "@" ? 1 : 0  // left top
  count += i - 1 >= 0 && grid[i - 1][j] == "@" ? 1 : 0  // left
  count += i + 1 < n && j - 1 >= 0 && grid[i + 1][j - 1] == "@" ? 1 : 0  // left bottom
  count += j - 1 >= 0 && grid[i][j - 1] == "@" ? 1 : 0  // top
  count += i - 1 >= 0 && j + 1 < n && grid[i - 1][j + 1] == "@" ? 1 : 0  // right top
  count += j + 1 < n && grid[i][j + 1] == "@" ? 1 : 0  // right
  count += i + 1 < n && j + 1 < n && grid[i + 1][j + 1] == "@" ? 1 : 0  // right bottom
  count += i + 1 < n && grid[i + 1][j] == "@" ? 1 : 0  // bottom
  return count
}

func readGrid(from path: String) -> [[Character]] {
  var grid: [[Character]] = []

  guard let contents = try? String(contentsOfFile: path) else {
    print("Failed to read file")
    return grid
  }

  contents.enumerateLines { line, _ in
    grid.append(Array(line))
  }

  return grid
}
