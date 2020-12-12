# frozen_string_literal: true

# Map
class Map
  TREE = '#'

  def initialize(pattern)
    @pattern = pattern
    @width = pattern.first.size
  end

  def [](x_coord, y_coord)
    @pattern[y_coord]&.[](x_coord % @width)
  end

  def collisions(dx, dy)
    x = 0
    y = 0
    count = 0

    loop do
      x += dx
      y += dy

      case self[x, y]
      when TREE then count += 1
      when nil then break
      end
    end

    count
  end
end

file_path = File.expand_path('data', __dir__)
INPUT = File.read(file_path)

PATTERN = INPUT.split("\n").map { |line| line.split('') }
SLOPES = [
  [1, 1],
  [3, 1],
  [5, 1],
  [7, 1],
  [1, 2]
].freeze

map = Map.new(PATTERN)
p "part a: #{map.collisions(3, 1)}"
p "part b: #{SLOPES.map { |(dx, dy)| map.collisions(dx, dy) }.inject(:*)}"
