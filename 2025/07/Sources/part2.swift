// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
  let path = "input.txt"
  let grid = readCharGrid(from: path)

  let totalRouteCount = countRoutesToBottom(grid: grid)
  print("Total routes count \(totalRouteCount)")
}

func countRoutesToBottom(grid: [[Character]]) -> Int {
  let rows = grid.count
  guard rows > 0 else { return 0 }
  let cols = grid[0].count

  guard let sIdx = grid[0].firstIndex(of: "S") else {
    fatalError("No S found")
  }
  let sx = grid[0].distance(from: grid[0].startIndex, to: sIdx)

  // ways for current row
  var ways = [Int](repeating: 0, count: cols)
  ways[sx] = 1

  // Propagate down row-by-row (looking at the cell we enter)
  if rows >= 2 {
    for r in 0..<(rows - 1) {
      var next = [Int](repeating: 0, count: cols)

      for x in 0..<cols {
        let w = ways[x]
        if w == 0 { continue }

        if grid[r + 1][x] == "^" {
          if x > 0 { next[x - 1] &+= w }
          if x + 1 < cols { next[x + 1] &+= w }
        } else {
          next[x] &+= w
        }
      }

      ways = next
    }
  }

  // Total routes that make it to the last row (after last split/straight move)
  return ways.reduce(0, &+)
}
