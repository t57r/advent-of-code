// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part1() {
  let path = "input.txt"
  let grid = readCharGrid(from: path)

  let splitCount = countSplits(grid: grid)
  print("Split count: \(splitCount)")
}

func readCharGrid(from path: String) -> [[Character]] {
  guard let contents = try? String(contentsOfFile: path) else {
    fatalError("Failed to read file")
  }

  return
    contents
    .split(whereSeparator: \.isNewline)
    .map { Array($0) }
}

func countSplits(grid: [[Character]]) -> Int {
  let rows = grid.count
  let cols = grid.first?.count ?? 0
  guard rows > 0, cols > 0 else { return 0 }

  guard let sIndex = grid[0].firstIndex(of: "S") else {
    fatalError("No S found on first row")
  }
  let startX = grid[0].distance(from: grid[0].startIndex, to: sIndex)

  var beams: Set<Int> = [startX]
  var splitCount = 0

  for r in 1..<rows {
    var nextBeams = Set<Int>()

    for x in beams {
      if grid[r][x] == "^" {
        splitCount += 1
        if x - 1 >= 0 { nextBeams.insert(x - 1) }
        if x + 1 < cols { nextBeams.insert(x + 1) }
      } else {
        nextBeams.insert(x)
      }
    }

    beams = nextBeams

    if beams.isEmpty { break }
  }

  return splitCount
}
