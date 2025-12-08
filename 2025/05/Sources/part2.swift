// The Swift Programming Language
// https://docs.swift.org/swift-book

import Foundation

func part2() {
  let path = "input.txt"
  let (ranges, _) = readData(from: path)

  let merged = mergeRanges(ranges)
  var totalFreshIngredients = 0
  for range in merged {
    totalFreshIngredients += range.upperBound - range.lowerBound + 1
  }

  print("Total fresih ingredients: \(totalFreshIngredients)")

}

func mergeRanges(_ ranges: [ClosedRange<Int>]) -> [ClosedRange<Int>] {
  guard !ranges.isEmpty else { return [] }

  let sorted = ranges.sorted { $0.lowerBound < $1.lowerBound }

  var merged: [ClosedRange<Int>] = []
  var current = sorted[0]

  for range in sorted.dropFirst() {
    if range.lowerBound <= current.upperBound {
      current = current.lowerBound...max(current.upperBound, range.upperBound)
    } else {
      merged.append(current)
      current = range
    }
  }

  merged.append(current)
  return merged
}
