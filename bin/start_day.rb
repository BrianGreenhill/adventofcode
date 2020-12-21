#!/usr/bin/env ruby
# frozen_string_literal: true

require 'uri'
require 'net/https'
require 'fileutils'

# TODO: create template file for python

# This script will pull the latest input (if available) from
# https://adventofcode.com/<year>/day/<day>/input
# and put it into a <year>/<day>/input file. It then uses a template based on
# an optional flag for the programming language (default: python)
#
# usage: ./start_day.rb <year> <day>
MAX_YEAR = 2020
MAX_DAY = 24
NUM_ARGS = 3
error = false
puts 'missing <year>, <day>, and/or <session_id> arguments' unless ARGV.length == NUM_ARGS
if ARGV[0].to_i > MAX_YEAR
  error = true
  puts 'invalid year'
end
if ARGV[1].to_i > MAX_DAY
  error = true
  puts 'invalid day'
end

exit unless error == false

# parse args
year = ARGV[0]
day = ARGV[1]
session_id = ARGV[2]

input_url = URI("https://adventofcode.com/#{year}/day/#{day}/input")
puts "creating #{year}/day#{day}/day#{day}.py ..."

# pull adventofcode.com/year/day/day_num/input to file
puts "pulling #{input_url} ..."
response = Net::HTTP.start(
  input_url.host,
  input_url.port,
  use_ssl: input_url.scheme == 'https'
) do |http|
  request = Net::HTTP::Get.new input_url
  request['Cookie'] = "session=#{session_id}"
  begin
    response = http.request request
  rescue StandardError
    puts "Failed to get puzzle input. Check the URL #{input_url.host}"
  end
end
puzzle_input = response.body
if response.body.empty?
  puts 'got an empty puzzle input, bailing out ...'
  exit
end

# create year/day/day_num structure (if not exists)
AOC_DIR = ENV['AOC_DIR']
dir_name = "#{AOC_DIR}/#{year}/day#{day}"
file_name = "#{dir_name}/day#{day}.py"

Dir.mkdir(dir_name) unless File.exist?(dir_name)
File.open(file_name, 'w') unless File.exist?(file_name)
File.open("#{dir_name}/input", 'w') { |file| file.write(puzzle_input) }
FileUtils.cp('./templates/puzzle.py', file_name)

puts "Created input and solution file in #{dir_name}. Good luck!"
